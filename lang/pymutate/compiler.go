/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pymutate

import (
	"path/filepath"

	"abate/lang/mutcore"
)

// nativeCompiler is the Python implementation of mutcore.NativeCompiler.
// Python has no separate compile step: the honest FastCompile verdict is a
// py_compile syntax check, and PackageBuild reuses the same check because
// a Python file that passes py_compile is ready to import—there is no
// additional package-level build gate.
type nativeCompiler struct {
	cfg config
	drv driver
}

var _ mutcore.NativeCompiler = (*nativeCompiler)(nil)

func newNativeCompiler(cfg config) mutcore.NativeCompiler {
	return &nativeCompiler{cfg: cfg, drv: driver{cfg: cfg}}
}

// FastCompile reports whether the mutated file still parses using py_compile.
func (c *nativeCompiler) FastCompile(dir, file string) bool {
	return compileFile(nil, c.cfg, filepath.Join(dir, filepath.FromSlash(file)))
}

// PackageBuild reports whether the package holding the mutated file still
// compiles. Python has no package-level build: a file passing py_compile is
// the complete build verdict, identical to the legacy Compile path.
func (c *nativeCompiler) PackageBuild(dir, file string) bool {
	return compileFile(nil, c.cfg, filepath.Join(dir, filepath.FromSlash(file)))
}

// PackageTest runs exactly one test against the mutated workspace.
func (c *nativeCompiler) PackageTest(dir, testDir, testName string) (bool, error) {
	// Python tests are node-id based: testDir is workspace-relative, testName is node id.
	// Delegate to RunTest via driver helper.
	return c.drv.RunTest(nil, dir, mutcore.TestRef{
		ID:   testName,
		Dir:  filepath.Join(dir, filepath.FromSlash(testDir)),
		Name: testName,
	})
}
