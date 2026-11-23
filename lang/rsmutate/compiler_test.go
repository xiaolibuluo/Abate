/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package rsmutate

import (
	"os"
	"path/filepath"
	"testing"
)

// rustToolchain returns the production config or skips the test when rustc
// or the llvm coverage tools are absent from the host.
func rustToolchain(t *testing.T) config {
	t.Helper()
	cfg, err := defaultConfig()
	if err != nil {
		t.Skipf("rust toolchain unavailable: %v", err)
	}
	return cfg
}

// writeCrate materializes one single-file crate inside a fresh temporary
// directory and returns the directory path.
func writeCrate(t *testing.T, src string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "lib.rs"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestFastCompileRejectsSyntaxError(t *testing.T) {
	cfg := rustToolchain(t)
	nc := newNativeCompiler(cfg)
	dir := writeCrate(t, "fn broken( { let x = ; }\n")
	if nc.FastCompile(dir, "lib.rs") {
		t.Fatal("rustc --emit=metadata must reject a syntax error")
	}
	if nc.PackageBuild(dir, "lib.rs") {
		t.Fatal("rustc --emit=obj must reject a syntax error")
	}
}

func TestPackageBuildRejectsCodegenErrorThatMetadataAccepts(t *testing.T) {
	cfg := rustToolchain(t)
	nc := newNativeCompiler(cfg)
	// std::arch::asm! with a directive the LLVM assembler cannot parse:
	// the metadata pass validates the macro shape and types, then the
	// object pass fails inside LLVM. Verified against rustc on this host
	// (metadata exit 0, obj exit 1).
	src := "fn main() { unsafe { std::arch::asm!(\".abate_invalid_directive\"); } }\n"
	dir := writeCrate(t, src)
	if !nc.FastCompile(dir, "lib.rs") {
		t.Fatal("rustc --emit=metadata must accept a well-formed asm! macro")
	}
	if nc.PackageBuild(dir, "lib.rs") {
		t.Fatal("rustc --emit=obj must reject an invalid inline-asm directive")
	}
}

func TestBothGatesAcceptValidCrate(t *testing.T) {
	cfg := rustToolchain(t)
	nc := newNativeCompiler(cfg)
	// rustc compiles a bare .rs file as a bin crate, so the source must
	// carry a main function or the metadata pass fails with E0601.
	dir := writeCrate(t, "pub fn add(a: i32, b: i32) -> i32 { a + b }\nfn main() { let _ = add(1, 2); }\n")
	if !nc.FastCompile(dir, "lib.rs") {
		t.Fatal("rustc --emit=metadata must accept a valid crate")
	}
	if !nc.PackageBuild(dir, "lib.rs") {
		t.Fatal("rustc --emit=obj must accept a valid crate")
	}
}
