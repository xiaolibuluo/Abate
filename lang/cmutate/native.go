/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cmutate

import (
	"context"
	"fmt"
	"os/exec"

	"abate/lang/mutcore"
)

// CSandboxPolicy returns the CopyPolicy for a native C project. Source
// and header files plus the common build manifests are copied per mutant.
// C projects typically have no package-manager dependency directory, so
// Once is empty. Object files and build outputs are excluded at the
// session layer (skippedDirs already drops bin/ and obj/).
// NewNative builds a native-sandbox mutation session over a real C
// project. It validates the toolchain, creates a sandbox root (C projects
// have no dependency directory by default, so Once is empty), and follows
// the same scan-discover-measure-assemble flow as New.
func NewNative(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg := DefaultConfig()
	for _, bin := range []string{cfg.CcBin, cfg.GcovBin} {
		if _, err := exec.LookPath(bin); err != nil {
			return nil, fmt.Errorf("%w: %s", mutcore.ErrToolchainMissing, bin)
		}
	}
	policy := CSandboxPolicy()
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	return mutcore.NewNative(ctx, driver{cfg: cfg}, root, policy)
}

func CSandboxPolicy() mutcore.CopyPolicy {
	return mutcore.CopyPolicy{
		Always:  []string{"*.c", "*.h", "Makefile", "makefile", "GNUmakefile", "CMakeLists.txt", "meson.build", "configure", "*.mk"},
		Once:    nil,
		Exclude: []string{"*.o", "*.a", "*.so", "build", "out"},
	}
}
