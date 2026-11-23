/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cmutate

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"abate/lang/mutcore"
)

// nativeCompiler is the C implementation of mutcore.NativeCompiler.
type nativeCompiler struct {
	cfg config
	drv driver
}

var _ mutcore.NativeCompiler = (*nativeCompiler)(nil)

func newNativeCompiler(cfg config) mutcore.NativeCompiler {
	return &nativeCompiler{cfg: cfg, drv: driver{cfg: cfg}}
}

// syntaxCheck performs the same gcc -fsyntax-only check as the legacy
// Compile method. It is extracted so both FastCompile and PackageBuild
// can share the identical logic without duplicating the exec path.
func (c *nativeCompiler) syntaxCheck(dir, file string) bool {
	cmd := exec.Command(c.cfg.CcBin, "-fsyntax-only",
		filepath.Join(dir, filepath.FromSlash(file)))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// FastCompile reports whether the mutated file still parses using gcc
// -fsyntax-only. This is the fastest honest check: it validates syntax
// without invoking the assembler or linker.
func (c *nativeCompiler) FastCompile(dir, file string) bool {
	return c.syntaxCheck(dir, file)
}

// PackageBuild reports whether the translation unit holding the mutated
// file still compiles to an object file. Unlike FastCompile which stops
// at syntax, this gate invokes the full frontend and code generator,
// catching type errors and semantic violations that the syntax-only pass
// would miss. It produces no executable, only an object file written to
// a temporary location that is discarded on return.
func (c *nativeCompiler) PackageBuild(dir, file string) bool {
	outDir, err := os.MkdirTemp("", "abate-cmutate-obj-")
	if err != nil {
		return false
	}
	defer func() { _ = os.RemoveAll(outDir) }()
	cmd := exec.Command(c.cfg.CcBin, "-c",
		filepath.Join(dir, filepath.FromSlash(file)),
		"-o", filepath.Join(outDir, "_.o"))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// PackageTest runs exactly one test against the mutated workspace.
func (c *nativeCompiler) PackageTest(dir, testDir, testName string) (bool, error) {
	return c.drv.RunTest(nil, dir, mutcore.TestRef{
		ID:   testName,
		Dir:  filepath.Join(dir, filepath.FromSlash(testDir)),
		Name: testName,
	})
}
