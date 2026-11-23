/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import (
	"container/list"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"abate/kernel/mutation"
	"abate/kernel/verify"
)

// stagePoolCap bounds how many mutant workspaces the session keeps
// materialized at once. The kernel judges points concurrently, and each
// in-flight point needs its own staged tree; the pool admits every active
// point but evicts idle trees least-recently-used first, so peak /tmp
// usage stays proportional to the judging concurrency instead of the
// mutant count.
const stagePoolCap = 8

// compileTimeout is the hard deadline for one Compile callback. The
// backend driver's Compile method receives a context with this timeout
// so a hung toolchain process (e.g. kotlinc deadlocked on a malformed
// mutant) is terminated by the OS instead of blocking the worker forever.
// Compile is a narrow syntax/type check that completes in sub-second
// wall time; 60 s leaves an order-of-magnitude headroom.
const compileTimeout = 60 * time.Second

// runTestTimeout is the hard deadline for one RunTest callback, applied
// at the session layer around the per-backend RunTest invocation. Each
// backend also applies its own per-test timeout (cfg.TestTimeout) inside
// RunTest, so this outer bound only fires when the inner timeout fails
// to cancel the context tree—it is a safety net, not a tuning knob.
const runTestTimeout = 300 * time.Second

// stagedEntry is one materialized mutant workspace in the pool. refs
// counts the evaluations currently running a toolchain against dir: an
// entry with refs > 0 must never be deleted from under a live run, or the
// runner would fail for a filesystem reason and fabricate a kill.
type stagedEntry struct {
	dir      string
	compiles bool
	refs     int
	elem     *list.Element
}

// Session is the owner of one real-workspace mutation session, shared by
// every language backend. It exposes the two kernel ports through separate
// adapter values, because both ports declare Fingerprint and a single
// method set cannot answer two different digests. Close releases the
// staged mutant workspaces; a session that ran zero evaluations owns
// nothing and closes cleanly.
type Session struct {
	core *core
}

// core is the shared session state behind both port adapters. It is never
// copied after construction: the token table, test list, coverage matrix,
// outcome cache and staged-mutant bookkeeping all belong to one session.
type core struct {
	drv     Driver
	root    string
	tokens  []Token
	byNode  map[string]Token
	tests   []TestRef
	matrix  *CoverMatrix
	cache   map[string]verify.Outcome
	tokenFP [32]byte
	testFP  [32]byte

	// staged is the pool of materialized mutant workspaces, keyed by
	// point fingerprint. Concurrent evaluations of one point share its
	// tree, different points each get their own, and idle trees are
	// evicted least-recently-used first under stagePoolCap. stageMu
	// guards both the map and the LRU list; toolchain runs never hold it.
	stageMu  sync.Mutex
	staged   map[[32]byte]*stagedEntry
	stageLRU *list.List

	// cacheMu guards the outcome cache so concurrent workers can record
	// verdicts without racing the map.
	cacheMu sync.Mutex

	// sandboxRoot is the pristine copy of the project tree in native
	// mode. When non-empty, stagedAcquire uses CopyWorkspaceNative
	// instead of CopyWorkspace so heavy dependency directories are
	// shared through symlinks rather than copied per mutant.
	sandboxRoot   string
	sandboxPolicy CopyPolicy
	compiler      NativeCompiler
}

// New builds a session over the given workspace root using the language
// driver: it scans mutable tokens, discovers every test of the tree and
// measures the per-test coverage on the pristine snapshot. Any toolchain
// failure aborts construction, so no mutant is ever judged against a
// partially probed baseline.
func New(ctx context.Context, drv Driver, root string) (*Session, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidRoot, err)
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("%w: %s", ErrInvalidRoot, root)
	}
	tokens, err := drv.ScanTokens(abs)
	if err != nil {
		return nil, err
	}
	tests, err := drv.DiscoverTests(ctx, abs)
	if err != nil {
		return nil, err
	}
	matrix, err := drv.MeasureCoverage(ctx, abs, tests)
	if err != nil {
		return nil, err
	}
	byNode := make(map[string]Token, len(tokens))
	for _, tok := range tokens {
		byNode[tok.Node()] = tok
	}
	return &Session{core: &core{
		drv: drv, root: abs,
		tokens: tokens, byNode: byNode, tests: tests, matrix: matrix,
		cache:    make(map[string]verify.Outcome, len(tokens)*2),
		staged:   make(map[[32]byte]*stagedEntry),
		stageLRU: list.New(),
		tokenFP:  FingerprintTokens(tokens),
		testFP:   FingerprintTests(tests, matrix.Files()),
	}}, nil
}

// NewNative builds a session in native (sandbox) mode over the project at
// root. It validates the copy policy, creates a sandbox root with the Once
// directories copied from the project, then runs the same discovery steps
// as New. Per-mutant workspaces created by stagedAcquire reference the
// sandbox root through CopyWorkspaceNative instead of copying the full
// tree from root, so heavy dependency directories are shared through
// symlinks.
func NewNative(ctx context.Context, drv Driver, root string, policy CopyPolicy) (*Session, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidRoot, err)
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("%w: %s", ErrInvalidRoot, root)
	}
	sandRoot, err := InitSandbox(abs, policy)
	if err != nil {
		return nil, err
	}
	tokens, err := drv.ScanTokens(abs)
	if err != nil {
		_ = os.RemoveAll(sandRoot)
		return nil, err
	}
	tests, err := drv.DiscoverTests(ctx, abs)
	if err != nil {
		_ = os.RemoveAll(sandRoot)
		return nil, err
	}
	matrix, err := drv.MeasureCoverage(ctx, abs, tests)
	if err != nil {
		_ = os.RemoveAll(sandRoot)
		return nil, err
	}
	byNode := make(map[string]Token, len(tokens))
	for _, tok := range tokens {
		byNode[tok.Node()] = tok
	}
	var compiler NativeCompiler
	if provider, ok := drv.(NativeCompilerProvider); ok {
		compiler = provider.NativeCompiler()
	}
	return &Session{core: &core{
		drv: drv, root: abs,
		tokens: tokens, byNode: byNode, tests: tests, matrix: matrix,
		cache:         make(map[string]verify.Outcome, len(tokens)*2),
		staged:        make(map[[32]byte]*stagedEntry),
		stageLRU:      list.New(),
		tokenFP:       FingerprintTokens(tokens),
		testFP:        FingerprintTests(tests, matrix.Files()),
		sandboxRoot:   sandRoot,
		sandboxPolicy: policy,
		compiler:      compiler,
	}}, nil
}

// Target returns the mutation.Target view of this session.
func (s *Session) Target() mutation.Target {
	return targetAdapter{core: s.core}
}

// Suite returns the mutation.Suite view of this session.
func (s *Session) Suite() mutation.Suite {
	return suiteAdapter{core: s.core}
}

// Close removes every staged mutant workspace and the sandbox root, if
// any. It is idempotent and safe to call after the judging workers have
// all finished.
func (s *Session) Close() error {
	s.core.unstage()
	if s.core.sandboxRoot != "" {
		_ = os.RemoveAll(s.core.sandboxRoot)
	}
	return nil
}

// Root exposes the absolute workspace root the session mutates.
func (s *Session) Root() string {
	return s.core.root
}

// Language exposes the canonical tag of the driver behind the session.
func (s *Session) Language() string {
	return s.core.drv.Language()
}

// TokenCount reports how many mutable tokens the scanner discovered.
func (s *Session) TokenCount() int {
	return len(s.core.tokens)
}

// TestCount reports how many tests the runner discovered.
func (s *Session) TestCount() int {
	return len(s.core.tests)
}

// evaluate answers one (point, test) query. Outcomes are cached under the
// point fingerprint plus the test identifier: the mutation kernel calls
// Evaluate for one point across the sorted test list and again inside the
// killing-test minimization, and the cache makes every repeated query a
// pure lookup, which the Target contract requires.
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
	if !c.matrix.TestCovers(test, tok.Path, tok.Line) {
		// Coverage-guided pruning: a test whose own measured coverage
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
	runCtx, runCancel := context.WithTimeout(context.Background(), runTestTimeout)
	passed, runErr := c.drv.RunTest(runCtx, stagedRoot, testRef)
	runCancel()
	if runErr != nil {
		// The toolchain itself failed (not a test failure): the mutant
		// was never actually judged, so the honest answer is inconclusive.
		return verify.OutcomeInconclusive
	}
	if !passed {
		return verify.OutcomeViolated
	}
	return verify.OutcomeHeld
}

// stagedAcquire returns the directory of the mutant workspace materialized
// for the given point, building it on first use, and takes one reference
// on the pool entry; the caller must release it with stagedRelease. The
// copy, mutation and compile run outside the pool lock so concurrent
// workers never serialize on each other's toolchain time; only the index
// updates are locked. The second return value reports whether the mutated
// tree compiles: a mutant that does not build is killed by the toolchain
// itself, and every test query for it answers VIOLATED without spawning
// further doomed runs.
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
	dir, err := os.MkdirTemp("", "abate-"+c.drv.Language()+"-mut-")
	if err != nil {
		return "", false, fmt.Errorf("%w: %v", ErrWorkspaceFailed, err)
	}
	if c.sandboxRoot != "" {
		if err := CopyWorkspaceNative(c.sandboxRoot, dir, c.sandboxPolicy, c.drv.RequiredFile); err != nil {
			_ = os.RemoveAll(dir)
			return "", false, err
		}
	} else if err := CopyWorkspace(c.root, dir, c.drv.RequiredFile); err != nil {
		_ = os.RemoveAll(dir)
		return "", false, err
	}
	if err := ApplyMutation(dir, tok, point.Mutated); err != nil {
		_ = os.RemoveAll(dir)
		return "", false, err
	}
	var compiles bool
	if c.compiler != nil {
		compiles = c.compiler.FastCompile(dir, tok.Path) &&
			c.compiler.PackageBuild(dir, tok.Path)
	} else {
		compileCtx, compileCancel := context.WithTimeout(context.Background(), compileTimeout)
		compiles = c.drv.Compile(compileCtx, dir, tok)
		compileCancel()
	}
	c.stageMu.Lock()
	defer c.stageMu.Unlock()
	if entry, ok := c.staged[point.Fingerprint]; ok {
		// Another worker staged the same point first: adopt its tree and
		// discard ours, so every evaluation of one point shares exactly
		// one materialized workspace.
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
// deleting a directory from under a running toolchain invocation.
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

// testByID locates one discovered test by its fully qualified identifier.
func (c *core) testByID(id string) (TestRef, bool) {
	for _, test := range c.tests {
		if test.ID == id {
			return test, true
		}
	}
	return TestRef{}, false
}
