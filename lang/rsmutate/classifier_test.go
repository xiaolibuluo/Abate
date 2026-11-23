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

	"abate/lang/frontkit"
	"abate/lang/mutcore"
)

// tok builds a frontkit token of one kind and text; offsets are irrelevant
// to the classifier, which only inspects Kind and Text.
func tok(kind frontkit.TokenKind, text string) frontkit.Token {
	return frontkit.Token{Kind: kind, Text: text, Line: 1, Column: 1}
}

// classifyAll feeds the token stream through one classifier and returns the
// class/keep verdict for every token, so tests can assert the full per-token
// outcome instead of only the final state.
func classifyAll(c *rsClassifier, toks []frontkit.Token) []bool {
	kept := make([]bool, len(toks))
	for i, lt := range toks {
		_, kept[i] = c.classify("", lt)
	}
	return kept
}

func TestClassifierMutableTokensOutsideTestFn(t *testing.T) {
	// Production code with a mutable operator: the classifier must keep it.
	c := &rsClassifier{}
	stream := []frontkit.Token{
		tok(frontkit.TokenKeyword, "fn"),
		tok(frontkit.TokenIdent, "add"),
		tok(frontkit.TokenPunct, "("),
		tok(frontkit.TokenPunct, ")"),
		tok(frontkit.TokenPunct, "{"),
		tok(frontkit.TokenIdent, "a"),
		tok(frontkit.TokenOp, "+"),
		tok(frontkit.TokenIdent, "b"),
		tok(frontkit.TokenPunct, "}"),
		tok(frontkit.TokenEOF, ""),
	}
	kept := classifyAll(c, stream)
	// Only the '+' operator is mutable; everything else is skipped or unknown.
	for i, lt := range stream {
		want := lt.Text == "+"
		if kept[i] != want {
			t.Fatalf("token %d (%q): kept=%v, want %v", i, lt.Text, kept[i], want)
		}
	}
}

func TestClassifierKillsEveryTokenInsideTestFn(t *testing.T) {
	// The regression this machine fixes: a single-file crate whose test fn
	// sits next to production code must never expose test-body tokens.
	c := &rsClassifier{}
	stream := []frontkit.Token{
		tok(frontkit.TokenKeyword, "fn"),
		tok(frontkit.TokenIdent, "test_add"),
		tok(frontkit.TokenPunct, "("),
		tok(frontkit.TokenPunct, ")"),
		tok(frontkit.TokenPunct, "{"),
		tok(frontkit.TokenIdent, "assert_eq"),
		tok(frontkit.TokenOp, "+"),
		tok(frontkit.TokenNum, "42"),
		tok(frontkit.TokenPunct, "}"),
		tok(frontkit.TokenEOF, ""),
	}
	kept := classifyAll(c, stream)
	for i, lt := range stream {
		if kept[i] {
			t.Fatalf("token %d (%q) inside fn test_* must not be mutable", i, lt.Text)
		}
	}
}

func TestClassifierNestedBraceDepth(t *testing.T) {
	// Nested blocks inside a test fn must keep the machine in rsSkipping
	// until the outermost closing brace returns depth to zero.
	c := &rsClassifier{}
	stream := []frontkit.Token{
		tok(frontkit.TokenKeyword, "fn"),
		tok(frontkit.TokenIdent, "test_nested"),
		tok(frontkit.TokenPunct, "{"), // depth 1
		tok(frontkit.TokenKeyword, "if"),
		tok(frontkit.TokenPunct, "{"), // depth 2
		tok(frontkit.TokenOp, "+"),
		tok(frontkit.TokenPunct, "}"), // depth 1
		tok(frontkit.TokenOp, "-"),    // still inside test body
		tok(frontkit.TokenPunct, "}"), // depth 0 -> rsNormal
		tok(frontkit.TokenKeyword, "fn"),
		tok(frontkit.TokenIdent, "prod"),
		tok(frontkit.TokenPunct, "{"),
		tok(frontkit.TokenOp, "*"),
		tok(frontkit.TokenPunct, "}"),
		tok(frontkit.TokenEOF, ""),
	}
	kept := classifyAll(c, stream)
	for i, lt := range stream {
		// Only the '*' in the production fn body is mutable.
		want := lt.Text == "*"
		if kept[i] != want {
			t.Fatalf("token %d (%q): kept=%v, want %v", i, lt.Text, kept[i], want)
		}
	}
}

func TestClassifierFnWithNonTestNameFallsBack(t *testing.T) {
	// fn followed by a non-test_* identifier returns to rsNormal, so a
	// mutable token in that ordinary function is still classified.
	c := &rsClassifier{}
	stream := []frontkit.Token{
		tok(frontkit.TokenKeyword, "fn"),
		tok(frontkit.TokenIdent, "compute"),
		tok(frontkit.TokenPunct, "{"),
		tok(frontkit.TokenNum, "1"),
		tok(frontkit.TokenOp, "=="),
		tok(frontkit.TokenNum, "2"),
		tok(frontkit.TokenPunct, "}"),
		tok(frontkit.TokenEOF, ""),
	}
	kept := classifyAll(c, stream)
	if !kept[4] {
		t.Fatal("'==' inside an ordinary fn must stay mutable")
	}
}

func TestClassifierEOFResetsState(t *testing.T) {
	// EOF in any state resets to rsNormal with depth zero: a truncated
	// stream must not leak the skipping state into the next file.
	c := &rsClassifier{}
	open := []frontkit.Token{
		tok(frontkit.TokenKeyword, "fn"),
		tok(frontkit.TokenIdent, "test_trunc"),
		tok(frontkit.TokenPunct, "{"),
		tok(frontkit.TokenEOF, ""),
	}
	_ = classifyAll(c, open)
	if c.state != rsNormal || c.depth != 0 {
		t.Fatalf("EOF must reset classifier, got state=%v depth=%d", c.state, c.depth)
	}
	// A fresh stream after the reset classifies normally.
	class, ok := c.classify("", tok(frontkit.TokenOp, "+"))
	if !ok || class != mutcore.ClassOperator {
		t.Fatalf("after EOF reset '+' must be ClassOperator, got %v/%v", class, ok)
	}
}

func TestScanTokensHybridCrateKeepsProductionTokens(t *testing.T) {
	// End-to-end through the real lexer and filesystem: one .rs file
	// holding both production code and a test fn must expose exactly the
	// production-side mutable tokens.
	root := t.TempDir()
	src := "fn add(a: i32, b: i32) -> i32 { a + b }\n" +
		"fn test_add() { assert_eq!(add(1, 2), 3); }\n"
	if err := os.WriteFile(filepath.Join(root, "lib.rs"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	d := driver{}
	toks, err := d.ScanTokens(root)
	if err != nil {
		t.Fatalf("ScanTokens: %v", err)
	}
	if len(toks) != 1 {
		t.Fatalf("expected exactly 1 mutable token (the '+'), got %d: %+v", len(toks), toks)
	}
	if toks[0].Text != "+" || toks[0].Class != mutcore.ClassOperator {
		t.Fatalf("unexpected surviving token: %+v", toks[0])
	}
}

func TestScanTokensTestOnlyCrateYieldsNoSourceFiles(t *testing.T) {
	// A crate whose only function is a test fn produces zero mutable
	// tokens, which LexicalScan surfaces as ErrNoSourceFiles.
	root := t.TempDir()
	src := "fn test_only() { assert_eq!(1 + 1, 2); }\n"
	if err := os.WriteFile(filepath.Join(root, "lib.rs"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	d := driver{}
	if _, err := d.ScanTokens(root); err == nil {
		t.Fatal("test-only crate must fail with ErrNoSourceFiles, got nil")
	}
}
