/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsmutate

import (
	"context"

	"abate/lang/mutcore"
)

// RsSandboxPolicy returns the CopyPolicy for a native Rust project. Source
// files and the Cargo manifests are copied per mutant; the target
// directory (containing debug/release build artifacts and dependency
// checkouts) is linked from the sandbox root so cargo -p builds can
// reuse the pre-built dependency graph. Build outputs inside target/ are
// excluded at the session layer through skippedDirs.
// NewNative builds a native-sandbox mutation session over a real Rust
// project. It validates the toolchain, creates a sandbox root with the
// target directory linked once, and follows the same scan-discover-
// measure-assemble flow as New.
func NewNative(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	policy := RsSandboxPolicy()
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	return mutcore.NewNative(ctx, driver{cfg: cfg}, root, policy)
}

func RsSandboxPolicy() mutcore.CopyPolicy {
	return mutcore.CopyPolicy{
		Always:  []string{"*.rs", "Cargo.toml", "Cargo.lock", "rust-toolchain.toml", "build.rs"},
		Once:    []string{"target"},
		Exclude: []string{"*.rs.bk", "*.rlib", "*.rmeta"},
	}
}
