/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/mutcore"
	"bytes"
	"container/list"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// stagePoolCap bounds how many mutant workspaces the session keeps
// materialized at once. The kernel judges points concurrently, and each
// in-flight point needs its own staged tree; the pool admits every active
// point but evicts idle trees least-recently-used first, so peak /tmp usage
// stays proportional to the judging concurrency instead of the mutant count.
const stagePoolCap = 8

// stagedEntry is one materialized mutant workspace in the pool. refs counts
// the evaluations currently running a toolchain against dir: an entry with
// refs > 0 must never be deleted from under a live `go test`, or the runner
// would fail for a filesystem reason and fabricate a kill.
type stagedEntry struct {
	dir      string
	compiles bool
	refs     int
	elem     *list.Element
}

// evaluate answers one (point, test) query. Outcomes are cached under the
// point fingerprint plus the test identifier: the mutation kernel calls
// Evaluate for one point across the sorted test list and again inside the
// killing-test minimization, and the cache makes every repeated query a
// pure lookup, which the Target contract requires. The cache mutex keeps
// concurrent workers safe without serializing the toolchain runs.
func (c *core) evaluate(point mutation.MutationPoint, test string) verify.Outcome {
	key := string(point.Fingerprint[:]) + "|" + test
	c.cacheMu.Lock()
	outcome, ok := c.cache[key]
	c.cacheMu.Unlock()
	if ok {
		return outcome
	}
	outcome = c.evaluateUncached(point, test)
	c.cacheMu.Lock()
	c.cache[key] = outcome
	c.cacheMu.Unlock()
	return outcome
}

// evaluateUncached acquires the mutant workspace for the point from the
// staged pool (building it on first use), then runs exactly the requested
// test against the mutated sources.
func (c *core) evaluateUncached(point mutation.MutationPoint, test string) verify.Outcome {
	tok, ok := c.byNode[point.Node]
	if !ok {
		return verify.OutcomeInconclusive
	}
	testRef, ok := c.testByID(test)
	if !ok {
		return verify.OutcomeInconclusive
	}
	if !c.matrix.testCovers(test, tok.Path, tok.Line) {
		// Coverage-guided pruning: a test whose own measured coverprofile
		// never reaches the mutated line cannot witness the mutation, so
		// the standard mutation-testing verdict for it is HELD without
		// staging a workspace at all.
		return verify.OutcomeHeld
	}
	stagedRoot, compileOK, err := c.stagedAcquire(point, tok)
	if err != nil {
		return verify.OutcomeViolated
	}
	defer c.stagedRelease(point.Fingerprint)
	if !compileOK {
		return verify.OutcomeViolated
	}
	rel, err := filepath.Rel(c.root, testRef.Dir)
	if err != nil {
		return verify.OutcomeViolated
	}
	passed, _, runErr := runGoTest(context.Background(), c.cfg,
		filepath.Join(stagedRoot, rel), "-run", "^"+testRef.Name+"$")
	if runErr != nil {
		// The toolchain itself failed (not an ExitError): the mutant was
		// never actually judged, so the honest answer is inconclusive.
		return verify.OutcomeInconclusive
	}
	if !passed {
		return verify.OutcomeViolated
	}
	return verify.OutcomeHeld
}

// stagedAcquire returns the directory of the mutant workspace materialized
// for the given point, building it on first use, and takes one reference on
// the pool entry; the caller must release it with stagedRelease. The copy,
// mutation and build run outside the pool lock so concurrent workers never
// serialize on each other's toolchain time; only the index updates are
// locked. The second return value reports whether the mutated tree compiles:
// a mutant that does not build is killed by the toolchain itself, and every
// test query for it answers VIOLATED without spawning further doomed runs.
func (c *core) stagedAcquire(point mutation.MutationPoint, tok Token) (string, bool, error) {
	c.stageMu.Lock()
	if entry, ok := c.staged[point.Fingerprint]; ok {
		entry.refs++
		c.stageLRU.MoveToFront(entry.elem)
		dir, compiles := entry.dir, entry.compiles
		c.stageMu.Unlock()
		return dir, compiles, nil
	}
	c.stageMu.Unlock()

	dir, err := os.MkdirTemp("", "abate-gomutate-mut-")
	if err != nil {
		return "", false, fmt.Errorf("%w: %v", ErrWorkspace, err)
	}
	// Sandbox mode: materialize only Always files from the pristine
	// sandbox root and symlink Once directories, skipping the full
	// copyWorkspace pass. Legacy mode keeps the byte-identical
	// copyWorkspace path.
	if c.sandboxRoot != "" {
		if err := mutcore.CopyWorkspaceNative(c.sandboxRoot, dir, c.sandboxPolicy, goRequiredFile); err != nil {
			_ = os.RemoveAll(dir)
			return "", false, err
		}
	} else {
		if err := copyWorkspace(c.root, dir); err != nil {
			_ = os.RemoveAll(dir)
			return "", false, err
		}
	}
	if err := applyMutation(dir, tok, point.Mutated); err != nil {
		_ = os.RemoveAll(dir)
		return "", false, err
	}
	// Two-level gating in native sandbox mode: a per-file syntax check
	// kills broken rewrites for free before the package build pays for a
	// toolchain invocation. The legacy path keeps its single build gate
	// byte-identical.
	var compiles bool
	if c.compiler != nil {
		compiles = c.compiler.FastCompile(dir, tok.Path) &&
			c.compiler.PackageBuild(dir, tok.Path)
	} else {
		compiles = compileTree(c.cfg, dir, tok)
	}

	c.stageMu.Lock()
	defer c.stageMu.Unlock()
	if entry, ok := c.staged[point.Fingerprint]; ok {
		// Another worker staged the same point first: adopt its tree and
		// discard ours, so every evaluation of one point shares exactly one
		// materialized workspace.
		_ = os.RemoveAll(dir)
		entry.refs++
		c.stageLRU.MoveToFront(entry.elem)
		return entry.dir, entry.compiles, nil
	}
	entry := &stagedEntry{dir: dir, compiles: compiles, refs: 1}
	entry.elem = c.stageLRU.PushFront(point.Fingerprint)
	c.staged[point.Fingerprint] = entry
	c.evictLocked()
	return dir, compiles, nil
}

// stagedRelease drops one reference acquired by stagedAcquire.
func (c *core) stagedRelease(fingerprint [32]byte) {
	c.stageMu.Lock()
	defer c.stageMu.Unlock()
	if entry, ok := c.staged[fingerprint]; ok && entry.refs > 0 {
		entry.refs--
	}
}

// evictLocked shrinks the pool to stagePoolCap by removing idle entries in
// least-recently-used order. Entries still referenced by a live toolchain
// run are skipped; if every entry is in use the pool stays over capacity
// until those runs finish, which keeps peak usage honest without ever
// deleting a directory from under a running `go test`.
func (c *core) evictLocked() {
	for len(c.staged) > stagePoolCap {
		evicted := false
		for elem := c.stageLRU.Back(); elem != nil; elem = elem.Prev() {
			fp, _ := elem.Value.([32]byte)
			entry, ok := c.staged[fp]
			if !ok || entry.refs > 0 {
				continue
			}
			delete(c.staged, fp)
			c.stageLRU.Remove(elem)
			_ = os.RemoveAll(entry.dir)
			evicted = true
			break
		}
		if !evicted {
			return
		}
	}
}

// compileTree reports whether the mutated workspace still compiles, judged
// by `go build` of the package holding the mutated file. Every operator of
// the registry rewrites a token inside one function body and can never
// change an exported API, so the package-level build is the complete and
// honest build verdict while costing a fraction of a whole-module build.
func compileTree(cfg runConfig, dir string, tok Token) bool {
	return buildPackage(cfg, dir, tok.Path)
}

// buildPackage reports whether the package holding file compiles inside
// the staged workspace dir. It is the shared build gate of the legacy and
// native pipelines, so a mutant killed here receives the identical verdict
// under both modes.
func buildPackage(cfg runConfig, dir, file string) bool {
	pkg := "./" + filepath.ToSlash(filepath.Dir(file))
	if pkg == "./." {
		pkg = "."
	}
	cmd := exec.Command(cfg.GoBin, "build", pkg)
	cmd.Dir = dir
	cmd.Env = cfg.goEnv()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// unstage removes every staged mutant workspace, if any. It is idempotent
// and safe to call after the judging workers have all finished, which the
// kernel guarantees because Engine.Run returns only after the pool joins.
func (c *core) unstage() {
	c.stageMu.Lock()
	entries := c.staged
	c.staged = make(map[[32]byte]*stagedEntry)
	c.stageLRU.Init()
	c.stageMu.Unlock()
	for _, entry := range entries {
		_ = os.RemoveAll(entry.dir)
	}
}

// applyMutation rewrites exactly one token of the staged workspace to the
// mutated text. The rewrite is verified against the snapshot discovered by
// the scanner: if the staged file does not carry the expected original text
// at the recorded byte interval, the workspace drifted and the evaluation
// aborts instead of judging the wrong source.
func applyMutation(dir string, tok Token, mutated string) error {
	target := filepath.Join(dir, filepath.FromSlash(tok.Path))
	content, err := os.ReadFile(target) //nolint:gosec // target is derived from the scanned snapshot inside our own staging tree.
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspace, err)
	}
	rewritten, err := rewriteToken(content, tok, mutated)
	if err != nil {
		return err
	}
	if err := os.WriteFile(target, rewritten, 0o644); err != nil {
		return fmt.Errorf("%w: %v", ErrWorkspace, err)
	}
	return nil
}

// testByID locates one discovered test by its fully qualified identifier.
func (c *core) testByID(id string) (testRef, bool) {
	for _, test := range c.tests {
		if test.ID == id {
			return test, true
		}
	}
	return testRef{}, false
}
