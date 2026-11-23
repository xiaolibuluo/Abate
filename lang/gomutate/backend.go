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
	"container/list"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Backend is the owner of one real-workspace mutation session. It exposes
// the two kernel ports through separate adapter values, because both ports
// declare Fingerprint and a single method set cannot answer two different
// digests. Close releases the staged mutant workspace; a session that ran
// zero evaluations owns nothing and closes cleanly.
type Backend struct {
	core *core
}

// core is the shared session state behind both port adapters. It is never
// copied after construction: the token table, test list, coverage matrix,
// outcome cache and staged-mutant bookkeeping all belong to one session.
type core struct {
	// root is the module root: the directory holding go.mod. Every staged
	// copy, build and test run is anchored here so a mutate target may be
	// any subtree of the module without losing the build context.
	root    string
	cfg     runConfig
	module  string
	tokens  []Token
	byNode  map[string]Token
	tests   []testRef
	matrix  *coverMatrix
	cache   map[string]verify.Outcome
	tokenFP [32]byte
	testFP  [32]byte
	// sandboxRoot is the sandbox root directory created by
	// InitSandbox. When non-empty, stagedAcquire copies the source
	// tree from the project root and symlinks Once directories from
	// this root instead of duplicating the dependency tree per mutant.
	sandboxRoot   string
	sandboxPolicy mutcore.CopyPolicy
	// compiler is the optional two-level gating accelerator. It is nil
	// on the legacy path and set by NewNative, so stagedAcquire can run
	// the cheap per-file syntax check before the package build.
	compiler mutcore.NativeCompiler
	// staged is the pool of materialized mutant workspaces, keyed by point
	// fingerprint. The kernel judges points concurrently, so the single
	// staged tree of the sequential backend is replaced by a reference
	// counted pool: concurrent evaluations of one point share its tree,
	// different points each get their own, and idle trees are evicted
	// least-recently-used first under stagePoolCap. stageMu guards both the
	// map and the LRU list; toolchain runs never hold it.
	stageMu  sync.Mutex
	staged   map[[32]byte]*stagedEntry
	stageLRU *list.List
	// cacheMu guards the outcome cache so concurrent workers can record
	// verdicts without racing the map.
	cacheMu sync.Mutex
}

// New builds a session over the module containing the given directory:
// it locates go.mod by walking upward, scans mutable tokens of the target
// subtree with module-relative paths, discovers every TestXxx of the module
// through the toolchain and measures the per-test coverage on the pristine
// tree. Anchoring everything at the module root is what keeps a staged
// mutant workspace buildable and its coverprofile paths comparable. Any
// toolchain failure aborts construction with the package sentinel, so no
// mutant is ever judged against a partially probed baseline.
func New(ctx context.Context, root string) (*Backend, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	cfg := defaultConfig()
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidRoot, err)
	}
	modRoot, err := findModuleRoot(abs)
	if err != nil {
		return nil, err
	}
	tokens, err := newScanner(modRoot, abs).scan()
	if err != nil {
		return nil, err
	}
	tests, err := discoverTests(ctx, cfg, modRoot, abs)
	if err != nil {
		return nil, err
	}
	module, err := modulePath(ctx, cfg, modRoot)
	if err != nil {
		return nil, err
	}
	matrix, err := measureCoverage(ctx, cfg, modRoot, module, tests)
	if err != nil {
		return nil, err
	}
	byNode := make(map[string]Token, len(tokens))
	for _, tok := range tokens {
		byNode[tok.Node()] = tok
	}
	return &Backend{core: &core{
		root: modRoot, cfg: cfg, module: module,
		tokens: tokens, byNode: byNode, tests: tests, matrix: matrix,
		cache:    make(map[string]verify.Outcome, len(tokens)*2),
		staged:   make(map[[32]byte]*stagedEntry),
		stageLRU: list.New(),
		tokenFP:  fingerprintTokens(tokens),
		testFP:   fingerprintTests(tests, matrix),
	}}, nil
}

// findModuleRoot walks upward from dir until it finds the directory holding
// go.mod. A mutate target may be any subtree of a module; the staged
// workspace must still be copied from the module root, or the go toolchain
// would run outside module context and every build would fail.
func findModuleRoot(dir string) (string, error) {
	current := dir
	for {
		if info, err := os.Stat(filepath.Join(current, "go.mod")); err == nil && !info.IsDir() {
			return current, nil
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", fmt.Errorf("%w: no go.mod found at or above %s", ErrInvalidRoot, dir)
		}
		current = parent
	}
}

// Target returns the mutation.Target view of this session.
func (b *Backend) Target() mutation.Target {
	return targetAdapter{core: b.core}
}

// Suite returns the mutation.Suite view of this session.
func (b *Backend) Suite() mutation.Suite {
	return suiteAdapter{core: b.core}
}

// Close removes the staged mutant workspace, if any. It is idempotent.
func (b *Backend) Close() error {
	b.core.unstage()
	if b.core.sandboxRoot != "" {
		_ = os.RemoveAll(b.core.sandboxRoot)
	}
	return nil
}

// Root exposes the absolute workspace root the session mutates.
func (b *Backend) Root() string {
	return b.core.root
}

// ModulePath exposes the resolved module path of the workspace.
func (b *Backend) ModulePath() string {
	return b.core.module
}

// TokenCount reports how many mutable tokens the scanner discovered.
func (b *Backend) TokenCount() int {
	return len(b.core.tokens)
}

// TestCount reports how many tests the toolchain discovered.
func (b *Backend) TestCount() int {
	return len(b.core.tests)
}
