/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktmutate

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"abate/lang/mutcore"
)

// nativeCompiler is the Kotlin implementation of mutcore.NativeCompiler.
type nativeCompiler struct {
	cfg config
	drv driver
}

var _ mutcore.NativeCompiler = (*nativeCompiler)(nil)

func newNativeCompiler(cfg config) mutcore.NativeCompiler {
	return &nativeCompiler{cfg: cfg, drv: driver{cfg: cfg}}
}

// syntaxCheck performs the same kotlinc check as the legacy Compile method.
func (c *nativeCompiler) syntaxCheck(dir, file string) bool {
	classDir, err := os.MkdirTemp("", "abate-ktmutate-csc-")
	if err != nil {
		return false
	}
	defer func() { _ = os.RemoveAll(classDir) }()
	cmd := exec.Command(c.cfg.KotlincBin,
		filepath.Join(dir, filepath.FromSlash(file)), "-d", classDir)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// FastCompile reports whether the mutated file still parses using kotlinc,
// identical to the legacy Compile path.
func (c *nativeCompiler) FastCompile(dir, file string) bool {
	return c.syntaxCheck(dir, file)
}

// PackageBuild reports whether the package holding the mutated file still
// compiles. For Kotlin, a single file compilation with kotlinc is the
// complete build verdict.
func (c *nativeCompiler) PackageBuild(dir, file string) bool {
	return c.syntaxCheck(dir, file)
}

// PackageTest runs exactly one test against the mutated workspace.
func (c *nativeCompiler) PackageTest(dir, testDir, testName string) (bool, error) {
	return c.drv.RunTest(nil, dir, mutcore.TestRef{
		ID:   testName,
		Dir:  filepath.Join(dir, filepath.FromSlash(testDir)),
		Name: testName,
	})
}
