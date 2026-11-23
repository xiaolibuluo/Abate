/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package csmutate

import (
	"os"
	"path/filepath"
	"testing"
)

// csToolchain returns the production config or skips the test when dotnet
// or the Roslyn compiler assembly is absent from the host.
func csToolchain(t *testing.T) config {
	t.Helper()
	cfg, err := defaultConfig()
	if err != nil {
		t.Skipf("C# toolchain unavailable: %v", err)
	}
	return cfg
}

// writeCSModule creates a temporary .cs file and returns its directory.
func writeCSModule(t *testing.T, src string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Prog.cs"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestCSFastCompileRejectsSyntaxError(t *testing.T) {
	cfg := csToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writeCSModule(t, "class Broken { void M( { } }")
	if nc.FastCompile(dir, "Prog.cs") {
		t.Fatal("Roslyn csc must reject a syntax error")
	}
	// C#: PackageBuild is the same csc gate.
	if nc.PackageBuild(dir, "Prog.cs") {
		t.Fatal("Roslyn csc must reject a syntax error (PackageBuild)")
	}
}

func TestCSBothGatesAcceptValidSource(t *testing.T) {
	cfg := csToolchain(t)
	nc := newNativeCompiler(cfg)
	src := `using System;
public class Prog {
    public static int Add(int a, int b) { return a + b; }
    public static void Main() { Console.WriteLine(Add(1, 2)); }
}
`
	dir := writeCSModule(t, src)
	if !nc.FastCompile(dir, "Prog.cs") {
		t.Fatal("Roslyn csc must accept a valid source file")
	}
	if !nc.PackageBuild(dir, "Prog.cs") {
		t.Fatal("Roslyn csc must accept a valid source file (PackageBuild)")
	}
}
