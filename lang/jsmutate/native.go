/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsmutate

import (
	"context"
	"fmt"
	"os/exec"

	"abate/lang/mutcore"
)

// JsSandboxPolicy returns the CopyPolicy for a native JavaScript/TypeScript
// project. Source files and the package manifest with its lock file are
// copied per mutant; node_modules is linked from the sandbox root so the
// dependency tree is shared across mutants. Generated build outputs are
// excluded.
// NewNative builds a native-sandbox mutation session over a real
// JavaScript/Node project. It validates the toolchain, creates a sandbox
// root with node_modules linked once, and then follows the same
// scan-discover-measure-assemble flow as New.
func NewNative(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg := DefaultConfig()
	if _, err := exec.LookPath(cfg.NodeBin); err != nil {
		return nil, fmt.Errorf("%w: node", mutcore.ErrToolchainMissing)
	}
	policy := JsSandboxPolicy()
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	return mutcore.NewNative(ctx, driver{cfg: cfg}, root, policy)
}

func JsSandboxPolicy() mutcore.CopyPolicy {
	return mutcore.CopyPolicy{
		Always:  []string{"*.js", "*.mjs", "*.cjs", "*.ts", "*.mts", "*.cts", "package.json", "package-lock.json", "yarn.lock", "pnpm-lock.yaml", "tsconfig.json"},
		Once:    []string{"node_modules"},
		Exclude: []string{"dist", "build", "out", ".next", ".nuxt", "coverage", ".nyc_output"},
	}
}
