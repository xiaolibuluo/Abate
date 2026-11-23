/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsmutate

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"abate/lang/mutcore"
)

// nativeCompiler is the JavaScript implementation of mutcore.NativeCompiler.
type nativeCompiler struct {
	cfg config
	drv driver
}

var _ mutcore.NativeCompiler = (*nativeCompiler)(nil)

func newNativeCompiler(cfg config) mutcore.NativeCompiler {
	return &nativeCompiler{cfg: cfg, drv: driver{cfg: cfg}}
}

// FastCompile reports whether the mutated file still parses using node --check.
// This is the cheapest possible gate: it checks syntax only, no execution.
func (c *nativeCompiler) FastCompile(dir, file string) bool {
	cmd := exec.Command(c.cfg.NodeBin, "--check",
		filepath.Join(dir, filepath.FromSlash(file)))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// PackageBuild reports whether the package holding the mutated file still
// compiles. JS has no separate build step: syntax check is the complete build
// verdict, identical to the legacy Compile path.
func (c *nativeCompiler) PackageBuild(dir, file string) bool {
	return c.FastCompile(dir, file)
}

// PackageTest runs exactly one test against the mutated workspace.
func (c *nativeCompiler) PackageTest(dir, testDir, testName string) (bool, error) {
	return c.drv.RunTest(nil, dir, mutcore.TestRef{
		ID:   testName,
		Dir:  filepath.Join(dir, filepath.FromSlash(testDir)),
		Name: testName,
	})
}
