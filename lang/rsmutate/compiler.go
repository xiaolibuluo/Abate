/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsmutate

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"abate/lang/mutcore"
)

// nativeCompiler is the Rust implementation of mutcore.NativeCompiler.
type nativeCompiler struct {
	cfg config
	drv driver
}

var _ mutcore.NativeCompiler = (*nativeCompiler)(nil)

func newNativeCompiler(cfg config) mutcore.NativeCompiler {
	return &nativeCompiler{cfg: cfg, drv: driver{cfg: cfg}}
}

// FastCompile reports whether the mutated file still parses and type-checks
// using rustc --emit=metadata. This validates syntax and type correctness
// without producing object code, which is the cheapest honest Rust gate.
func (c *nativeCompiler) FastCompile(dir, file string) bool {
	cmd := exec.Command(c.cfg.RustcBin, "--emit=metadata", "-o", os.DevNull,
		filepath.Join(dir, filepath.FromSlash(file)))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// PackageBuild reports whether the crate holding the mutated file still
// compiles to object code. Unlike FastCompile which stops after type
// checking, this gate invokes the full code-generation pipeline,
// catching LLVM-level errors that survive the metadata pass. The output
// object file is discarded on return.
func (c *nativeCompiler) PackageBuild(dir, file string) bool {
	outDir, err := os.MkdirTemp("", "abate-rsmutate-obj-")
	if err != nil {
		return false
	}
	defer func() { _ = os.RemoveAll(outDir) }()
	cmd := exec.Command(c.cfg.RustcBin, "--emit=obj",
		"-o", filepath.Join(outDir, "_.o"),
		filepath.Join(dir, filepath.FromSlash(file)))
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
