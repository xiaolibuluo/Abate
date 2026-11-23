/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import (
	"os"
	"path/filepath"
	"testing"
)

// requiredAll returns true for every file, simulating the simplest driver.
func requiredAll(name string) bool { return true }

func TestCopyPolicyValidateRejectsEmptyAlways(t *testing.T) {
	p := CopyPolicy{}
	if err := p.Validate(); err == nil {
		t.Fatal("empty Always must be rejected")
	}
}

func TestCopyPolicyValidateAllowsValidPolicy(t *testing.T) {
	p := CopyPolicy{
		Always:  []string{"*.rs"},
		Once:    []string{"target"},
		Exclude: []string{"*.tmp"},
	}
	if err := p.Validate(); err != nil {
		t.Fatalf("valid policy must not fail: %v", err)
	}
}

func TestCopyPolicyValidateRejectsOverlappingOnceExclude(t *testing.T) {
	p := CopyPolicy{
		Always:  []string{"*.rs"},
		Once:    []string{"target"},
		Exclude: []string{"target"},
	}
	if err := p.Validate(); err == nil {
		t.Fatal("overlapping Once/Exclude must be rejected")
	}
}

func TestInitSandboxCreatesSandboxRoot(t *testing.T) {
	// Setup a minimal project with one Once directory.
	proj := t.TempDir()
	onceDir := filepath.Join(proj, "vendor")
	if err := os.MkdirAll(onceDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(onceDir, "lib.go"), []byte("package lib\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	policy := CopyPolicy{
		Always: []string{"*.go"},
		Once:   []string{"vendor"},
	}
	sandRoot, err := InitSandbox(proj, policy)
	if err != nil {
		t.Fatalf("InitSandbox: %v", err)
	}
	defer os.RemoveAll(sandRoot)
	// The vendor directory must exist inside the sandbox root.
	if info, err := os.Stat(filepath.Join(sandRoot, "vendor")); err != nil || !info.IsDir() {
		t.Fatal("Once directory not materialized inside sandbox root")
	}
	// The content must match.
	data, err := os.ReadFile(filepath.Join(sandRoot, "vendor", "lib.go"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "package lib\n" {
		t.Fatalf("vendor/lib.go content mismatch: %q", string(data))
	}
}

func TestInitSandboxSkipsMissingOnceDirectory(t *testing.T) {
	proj := t.TempDir()
	policy := CopyPolicy{
		Always: []string{"*.go"},
		Once:   []string{"nonexistent"},
	}
	sandRoot, err := InitSandbox(proj, policy)
	if err != nil {
		t.Fatalf("InitSandbox must not fail on missing Once: %v", err)
	}
	defer os.RemoveAll(sandRoot)
}

func TestCopyWorkspaceNativeMaterializesAlwaysAndLinksOnce(t *testing.T) {
	proj := t.TempDir()
	// Always source file — written in the project root so copyDir reproduces
	// the real InitSandbox snapshot that includes both sources and Once dirs.
	if err := os.WriteFile(filepath.Join(proj, "main.rs"), []byte("fn main() {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Once directory (in the real project root).
	onceDir := filepath.Join(proj, "target")
	if err := os.MkdirAll(onceDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(onceDir, "dep.o"), []byte("object\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Simulate sandbox root prepared by InitSandbox: the Once directory
	// is materialized once, and the source files (Always) also live there
	// because the real InitSandbox copies the project tree.
	sandRoot := t.TempDir()
	if err := copyDir(proj, sandRoot); err != nil {
		t.Fatal(err)
	}

	policy := CopyPolicy{
		Always: []string{"*.rs"},
		Once:   []string{"target"},
	}
	dst := t.TempDir()
	if err := CopyWorkspaceNative(sandRoot, dst, policy, requiredAll); err != nil {
		t.Fatalf("CopyWorkspaceNative: %v", err)
	}

	// Always file must be physically copied.
	data, err := os.ReadFile(filepath.Join(dst, "main.rs"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "fn main() {}\n" {
		t.Fatalf("main.rs content mismatch: %q", string(data))
	}

	// Once directory must be a symlink.
	link, err := os.Readlink(filepath.Join(dst, "target"))
	if err != nil {
		t.Fatalf("target must be a symlink: %v", err)
	}
	if link != filepath.Join(sandRoot, "target") {
		t.Fatalf("symlink target mismatch: %q", link)
	}
}

func TestCopyWorkspaceNativeSkipsSkippedDirs(t *testing.T) {
	proj := t.TempDir()
	// .git and __pycache__ must never be materialized.
	for _, d := range []string{".git", "__pycache__"} {
		os.MkdirAll(filepath.Join(proj, d), 0o755)
		os.WriteFile(filepath.Join(proj, d, "meta.json"), []byte("{}"), 0o644)
	}
	if err := os.WriteFile(filepath.Join(proj, "main.go"), []byte("package main\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	sandRoot := t.TempDir()
	if err := copyDir(filepath.Join(proj, ".git"), filepath.Join(sandRoot, ".git")); err != nil {
		t.Fatal(err)
	}
	policy := CopyPolicy{Always: []string{"*.go"}}
	dst := t.TempDir()
	if err := CopyWorkspaceNative(sandRoot, dst, policy, requiredAll); err != nil {
		t.Fatalf("CopyWorkspaceNative: %v", err)
	}
	for _, d := range []string{".git", "__pycache__"} {
		if _, err := os.Stat(filepath.Join(dst, d)); !os.IsNotExist(err) {
			t.Fatalf("skipped dir %s must not exist in staged tree", d)
		}
	}
}
