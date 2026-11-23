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
	"strings"
	"testing"
)

func TestCopyWorkspaceMaterializesOnlyRequiredFiles(t *testing.T) {
	src := t.TempDir()
	os.WriteFile(filepath.Join(src, "main.go"), []byte("package main\n"), 0o644)
	os.WriteFile(filepath.Join(src, "README.md"), []byte("# doc\n"), 0o644)
	os.MkdirAll(filepath.Join(src, ".git"), 0o755)
	os.WriteFile(filepath.Join(src, ".git", "HEAD"), []byte("ref: refs/heads/main\n"), 0o644)

	dst := t.TempDir()
	err := CopyWorkspace(src, dst, func(name string) bool {
		return strings.HasSuffix(name, ".go")
	})
	if err != nil {
		t.Fatalf("CopyWorkspace: %v", err)
	}
	// main.go must exist.
	if _, err := os.Stat(filepath.Join(dst, "main.go")); err != nil {
		t.Fatal("main.go must be materialized")
	}
	// README.md must NOT exist.
	if _, err := os.Stat(filepath.Join(dst, "README.md")); !os.IsNotExist(err) {
		t.Fatal("README.md must not be materialized")
	}
	// .git must NOT exist (skippedDirs).
	if _, err := os.Stat(filepath.Join(dst, ".git")); !os.IsNotExist(err) {
		t.Fatal(".git must be skipped")
	}
}

func TestRewriteTokenExactMatch(t *testing.T) {
	content := []byte("a + b")
	tok := Token{Offset: 2, End: 3, Text: "+", Path: "test.rs", Line: 1, Column: 3}
	result, err := RewriteToken(content, tok, "-")
	if err != nil {
		t.Fatalf("RewriteToken: %v", err)
	}
	if string(result) != "a - b" {
		t.Fatalf("expected 'a - b', got %q", string(result))
	}
}

func TestRewriteTokenRejectsTokenDrift(t *testing.T) {
	content := []byte("a + b")
	tok := Token{Offset: 2, End: 3, Text: "*", Path: "test.rs", Line: 1, Column: 3}
	if _, err := RewriteToken(content, tok, "-"); err == nil {
		t.Fatal("token drift must be rejected")
	}
}

func TestRewriteTokenRejectsOutOfBounds(t *testing.T) {
	content := []byte("a + b")
	tok := Token{Offset: 10, End: 11, Text: "+", Path: "test.rs"}
	if _, err := RewriteToken(content, tok, "-"); err == nil {
		t.Fatal("out-of-bounds token must be rejected")
	}
}

func TestRewriteTokenRejectsReversedInterval(t *testing.T) {
	content := []byte("a + b")
	tok := Token{Offset: 3, End: 2, Text: "", Path: "test.rs"}
	if _, err := RewriteToken(content, tok, "-"); err == nil {
		t.Fatal("reversed interval must be rejected")
	}
}

func TestRewriteTokenNegativeOffset(t *testing.T) {
	content := []byte("a + b")
	tok := Token{Offset: -1, End: 1, Text: "a", Path: "test.rs"}
	if _, err := RewriteToken(content, tok, "-"); err == nil {
		t.Fatal("negative offset must be rejected")
	}
}

func TestApplyMutationRewritesStagedFile(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "lib.rs"), []byte("fn add() { a + b }\n"), 0o644)
	// "fn add() { a + b }\n": f=0 n=1 ' '=2 a=3 d=4 d=5 (=6 )=7 ' '=8 {=9 ' '=10 a=11 ' '=12 +=13 ' '=14 b=15 ' '=16 }=17 \n=18
	tok := Token{
		Path: "lib.rs", Offset: 13, End: 14,
		Text: "+", Class: ClassOperator, Line: 1, Column: 14,
	}
	if err := ApplyMutation(dir, tok, "-"); err != nil {
		t.Fatalf("ApplyMutation: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib.rs"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "fn add() { a - b }\n" {
		t.Fatalf("unexpected content: %q", string(data))
	}
}

func TestApplyMutationRejectsMissingFile(t *testing.T) {
	dir := t.TempDir()
	tok := Token{Path: "nonexistent.rs", Offset: 0, End: 1, Text: "x"}
	if err := ApplyMutation(dir, tok, "y"); err == nil {
		t.Fatal("missing file must be rejected")
	}
}

func TestApplyMutationRejectsTokenDriftInStagedTree(t *testing.T) {
	dir := t.TempDir()
	// Write file with '+' at offset 14, but tell ApplyMutation the token is '*'.
	os.WriteFile(filepath.Join(dir, "lib.rs"), []byte("fn add() { a + b }\n"), 0o644)
	tok := Token{
		Path: "lib.rs", Offset: 14, End: 15,
		Text: "*", Class: ClassOperator, Line: 1, Column: 15,
	}
	if err := ApplyMutation(dir, tok, "-"); err == nil {
		t.Fatal("token drift must be rejected")
	}
	// File must remain unchanged.
	data, _ := os.ReadFile(filepath.Join(dir, "lib.rs"))
	if string(data) != "fn add() { a + b }\n" {
		t.Fatal("file must not be modified on drift detection")
	}
}
