/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package ktmutate

import (
	"os"
	"path/filepath"
	"testing"
)

// ktToolchain returns the production config or skips the test when kotlinc
// or the Java runtime is absent from the host.
func ktToolchain(t *testing.T) config {
	t.Helper()
	cfg, err := defaultConfig()
	if err != nil {
		t.Skipf("Kotlin toolchain unavailable: %v", err)
	}
	return cfg
}

// writeKTModule creates a temporary .kt file and returns its directory.
func writeKTModule(t *testing.T, src string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Prog.kt"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestKTFastCompileRejectsSyntaxError(t *testing.T) {
	cfg := ktToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writeKTModule(t, "fun broken( {")
	if nc.FastCompile(dir, "Prog.kt") {
		t.Fatal("kotlinc must reject a syntax error")
	}
	// Kotlin: PackageBuild is the same kotlinc gate.
	if nc.PackageBuild(dir, "Prog.kt") {
		t.Fatal("kotlinc must reject a syntax error (PackageBuild)")
	}
}

func TestKTBothGatesAcceptValidSource(t *testing.T) {
	cfg := ktToolchain(t)
	nc := newNativeCompiler(cfg)
	src := "fun add(a: Int, b: Int): Int = a + b\nfun main() { println(add(1, 2)) }\n"
	dir := writeKTModule(t, src)
	if !nc.FastCompile(dir, "Prog.kt") {
		t.Fatal("kotlinc must accept a valid source file")
	}
	if !nc.PackageBuild(dir, "Prog.kt") {
		t.Fatal("kotlinc must accept a valid source file (PackageBuild)")
	}
}
