/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cmutate

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// cToolchain returns the production config or skips the test when gcc or
// gcov is absent from the host.
func cToolchain(t *testing.T) config {
	t.Helper()
	cfg := DefaultConfig()
	for _, bin := range []string{cfg.CcBin, cfg.GcovBin} {
		if _, err := exec.LookPath(bin); err != nil {
			t.Skipf("C toolchain unavailable: %v", err)
		}
	}
	return cfg
}

// writeTU materializes one single translation unit inside a fresh temporary
// directory and returns the directory path.
func writeTU(t *testing.T, src string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "a.c"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestFastCompileRejectsSyntaxError(t *testing.T) {
	cfg := cToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writeTU(t, "int main(void) { return ;;; }\n")
	if nc.FastCompile(dir, "a.c") {
		t.Fatal("gcc -fsyntax-only must reject a syntax error")
	}
	if nc.PackageBuild(dir, "a.c") {
		t.Fatal("gcc -c must reject a syntax error")
	}
}

func TestPackageBuildRejectsAssembleErrorThatSyntaxAccepts(t *testing.T) {
	cfg := cToolchain(t)
	nc := newNativeCompiler(cfg)
	// GNU inline asm with a directive the assembler cannot parse: the
	// syntax-only pass accepts the statement, then gcc -c fails inside as.
	// Verified against gcc on this host (syntax exit 0, compile exit 1).
	src := "int main(void) { __asm__ volatile (\".abate_invalid_directive\"); return 0; }\n"
	dir := writeTU(t, src)
	if !nc.FastCompile(dir, "a.c") {
		t.Fatal("gcc -fsyntax-only must accept well-formed GNU inline asm")
	}
	if nc.PackageBuild(dir, "a.c") {
		t.Fatal("gcc -c must reject an invalid assembler directive")
	}
}

func TestBothGatesAcceptValidTranslationUnit(t *testing.T) {
	cfg := cToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writeTU(t, "int add(int a, int b) { return a + b; }\n")
	if !nc.FastCompile(dir, "a.c") {
		t.Fatal("gcc -fsyntax-only must accept a valid translation unit")
	}
	if !nc.PackageBuild(dir, "a.c") {
		t.Fatal("gcc -c must accept a valid translation unit")
	}
}
