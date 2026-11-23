/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gomutate

import (
	"abate/kernel/verify"
	"abate/lang/mutcore"
	"container/list"
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// GoSandboxPolicy returns the CopyPolicy for a native Go project: every
// .go file and the two module manifests are copied per mutant, while the
// vendor directory (when present) is linked from the sandbox root so a
// large dependency tree incurs a one-time cost. Build outputs and VCS
// directories are excluded at the session layer (skippedDirs).
func GoSandboxPolicy() mutcore.CopyPolicy {
	return mutcore.CopyPolicy{
		Always:  []string{"*.go", "go.mod", "go.sum"},
		Once:    []string{"vendor"},
		Exclude: []string{"go.work", "go.work.sum"},
	}
}

// GoSandboxPolicyWithCustom returns a CopyPolicy that adds the caller's
// extra directories to Once and paths to Exclude on top of the defaults.
func GoSandboxPolicyWithCustom(extraOnce, extraExclude []string) mutcore.CopyPolicy {
	base := GoSandboxPolicy()
	base.Once = append(append([]string{}, base.Once...), extraOnce...)
	base.Exclude = append(append([]string{}, base.Exclude...), extraExclude...)
	return base
}

// NewNative builds a native-sandbox mutation session over a real Go
// project. It locates the module root, validates the CopyPolicy, creates a
// sandbox root with the Once directories materialized once, and then
// follows the same scan-discover-measure-assemble flow as New. The
// returned Backend's stagedAcquire copies only Always files per mutant and
// symlinks the Once directories from the sandbox root, so a large vendor/
// tree is never duplicated per mutant.
func NewNative(ctx context.Context, root string) (*Backend, error) {
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
	policy := GoSandboxPolicy()
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	sandboxRoot, err := mutcore.InitSandbox(modRoot, policy)
	if err != nil {
		return nil, err
	}
	tokens, err := newScanner(modRoot, abs).scan()
	if err != nil {
		_ = os.RemoveAll(sandboxRoot)
		return nil, err
	}
	tests, err := discoverTests(ctx, cfg, modRoot, abs)
	if err != nil {
		_ = os.RemoveAll(sandboxRoot)
		return nil, err
	}
	module, err := modulePath(ctx, cfg, modRoot)
	if err != nil {
		_ = os.RemoveAll(sandboxRoot)
		return nil, err
	}
	matrix, err := measureCoverage(ctx, cfg, modRoot, module, tests)
	if err != nil {
		_ = os.RemoveAll(sandboxRoot)
		return nil, err
	}
	byNode := make(map[string]Token, len(tokens))
	for _, tok := range tokens {
		byNode[tok.Node()] = tok
	}
	return &Backend{core: &core{
		root:          modRoot,
		cfg:           cfg,
		module:        module,
		tokens:        tokens,
		byNode:        byNode,
		tests:         tests,
		matrix:        matrix,
		cache:         make(map[string]verify.Outcome, len(tokens)*2),
		staged:        make(map[[32]byte]*stagedEntry),
		stageLRU:      list.New(),
		tokenFP:       fingerprintTokens(tokens),
		testFP:        fingerprintTests(tests, matrix),
		sandboxRoot:   sandboxRoot,
		sandboxPolicy: policy,
		compiler:      newNativeCompiler(cfg),
	}}, nil
}
