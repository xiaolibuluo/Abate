/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csmutate

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"abate/lang/mutcore"
)

// nativeCompiler is the C# implementation of mutcore.NativeCompiler.
type nativeCompiler struct {
	cfg config
	drv driver
}

var _ mutcore.NativeCompiler = (*nativeCompiler)(nil)

func newNativeCompiler(cfg config) mutcore.NativeCompiler {
	return &nativeCompiler{cfg: cfg, drv: driver{cfg: cfg}}
}

// syntaxCheck performs the same csc check as the legacy Compile method.
func (c *nativeCompiler) syntaxCheck(dir, file string) bool {
	outDir, err := os.MkdirTemp("", "abate-csmutate-csc-")
	if err != nil {
		return false
	}
	defer func() { _ = os.RemoveAll(outDir) }()
	args := []string{
		c.cfg.CscDLL, "/nologo", "/t:exe",
		"/out:" + filepath.Join(outDir, harnessAssembly+".dll"),
	}
	// The standalone Roslyn executor (dotnet csc.dll) does not carry the
	// implicit framework references that the MSBuild C# target injects;
	// every invocation must list the BCL assemblies the mutated source
	// may reference. The RefDir is resolved once at construction from
	// the dotnet packs tree.
	if c.cfg.RefDir != "" {
		args = append(args,
			"-r:"+filepath.Join(c.cfg.RefDir, "System.Runtime.dll"),
			"-r:"+filepath.Join(c.cfg.RefDir, "System.Console.dll"),
			"-r:"+filepath.Join(c.cfg.RefDir, "System.Linq.dll"),
			"-r:"+filepath.Join(c.cfg.RefDir, "System.Collections.dll"),
		)
	}
	args = append(args, filepath.Join(dir, filepath.FromSlash(file)))
	cmd := exec.Command(c.cfg.DotnetBin, args...)
	cmd.Env = driver{cfg: c.cfg}.env()
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	return cmd.Run() == nil
}

// FastCompile reports whether the mutated file still parses using the C#
// compiler, identical to the legacy Compile path.
func (c *nativeCompiler) FastCompile(dir, file string) bool {
	return c.syntaxCheck(dir, file)
}

// PackageBuild reports whether the package holding the mutated file still
// compiles. For C#, a single file compilation with the compiler is the
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
