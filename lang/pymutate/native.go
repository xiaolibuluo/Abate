/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pymutate

import (
	"context"
	"fmt"
	"os/exec"

	"abate/lang/mutcore"
)

// PySandboxPolicy returns the CopyPolicy for a native Python project.
// Source files and the standard project manifests are copied per mutant;
// a virtual environment directory is linked from the sandbox root when
// present so the installed packages are shared across mutants. Test
// caches and compiled bytecode are excluded at the session layer.
// NewNative builds a native-sandbox mutation session over a real Python
// project. It validates the toolchain, creates a sandbox root with the
// virtual environment directory linked once, and then follows the same
// scan-discover-measure-assemble flow as New. Per-mutant workspaces
// symlink the Once directories from the sandbox root.
func NewNative(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg := DefaultConfig()
	if _, err := exec.LookPath(cfg.PythonBin); err != nil {
		return nil, fmt.Errorf("%w: python3", mutcore.ErrToolchainMissing)
	}
	if _, err := exec.LookPath(cfg.PytestBin); err != nil {
		return nil, fmt.Errorf("%w: pytest", mutcore.ErrToolchainMissing)
	}
	policy := PySandboxPolicy()
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	return mutcore.NewNative(ctx, driver{cfg: cfg}, root, policy)
}

func PySandboxPolicy() mutcore.CopyPolicy {
	return mutcore.CopyPolicy{
		Always:  []string{"*.py", "requirements.txt", "pyproject.toml", "setup.cfg", "setup.py", "conftest.py", "pytest.ini", "tox.ini"},
		Once:    []string{"venv", ".venv", "env"},
		Exclude: []string{"__pycache__", ".pytest_cache", "*.pyc", ".tox", "dist", "build", "*.egg-info"},
	}
}
