/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package pymutate

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// pyToolchain returns the production config or skips the test when python3
// is absent from the host.
func pyToolchain(t *testing.T) config {
	t.Helper()
	cfg := DefaultConfig()
	if _, err := exec.LookPath(cfg.PythonBin); err != nil {
		t.Skipf("Python toolchain unavailable: %v", err)
	}
	return cfg
}

// writePyModule creates a temporary .py file and returns its directory.
func writePyModule(t *testing.T, src string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "mod.py"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestPythonFastCompileRejectsSyntaxError(t *testing.T) {
	cfg := pyToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writePyModule(t, "def broken(:")
	if nc.FastCompile(dir, "mod.py") {
		t.Fatal("py_compile must reject a syntax error")
	}
	// Python: PackageBuild is the same py_compile gate.
	if nc.PackageBuild(dir, "mod.py") {
		t.Fatal("py_compile must reject a syntax error (PackageBuild)")
	}
}

func TestPythonBothGatesAcceptValidModule(t *testing.T) {
	cfg := pyToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writePyModule(t, "def add(a: int, b: int) -> int:\n    return a + b\n")
	if !nc.FastCompile(dir, "mod.py") {
		t.Fatal("py_compile must accept a valid module")
	}
	if !nc.PackageBuild(dir, "mod.py") {
		t.Fatal("py_compile must accept a valid module (PackageBuild)")
	}
}
