/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"strings"
	"testing"
)

// TestLexTokenKinds covers every lexical category plus whitespace handling.
func TestLexTokenKinds(t *testing.T) {
	t.Parallel()
	in := "(ε ∘ c-delay-v1 ⊗ operator://m-repl-v1)⁻¹"
	want := []tokenKind{
		tokLParen, tokIdentity, tokCompose, tokBare, tokProduct,
		tokRef, tokRParen, tokInverse, tokEOF,
	}
	tokens, err := lex(in)
	requireOK(t, err)
	if len(tokens) != len(want) {
		t.Fatalf("got %d tokens, want %d: %+v", len(tokens), len(want), tokens)
	}
	for i := range want {
		if tokens[i].kind != want[i] {
			t.Fatalf("token %d = %v, want %v", i, tokens[i].kind, want[i])
		}
	}
	if tokens[5].text != "m-repl-v1" {
		t.Fatalf("ref text = %q, want %q", tokens[5].text, "m-repl-v1")
	}
	if tokens[3].text != "c-delay-v1" {
		t.Fatalf("bare text = %q", tokens[3].text)
	}
}

// TestLexErrors covers unexpected runes and empty references.
func TestLexErrors(t *testing.T) {
	t.Parallel()
	for _, in := range []string{"a @ b", "operator://∘", "a / b", "a⁻"} {
		if _, err := lex(in); err == nil {
			t.Fatalf("lex(%q) must fail", in)
		}
	}
	// A lone superscript minus is not a valid inverse token.
	if _, err := lex("⁻"); err == nil {
		t.Fatal("lex of a lone inverse rune must fail")
	}
}

// TestLexAllRuneClasses exercises isIDRune across letters, digits and separators.
func TestLexAllRuneClasses(t *testing.T) {
	t.Parallel()
	tokens, err := lex("abc-1_2.3")
	requireOK(t, err)
	if len(tokens) != 2 || tokens[0].text != "abc-1_2.3" {
		t.Fatalf("identifier lexing drift: %+v", tokens)
	}
}

// TestHasPrefixAt covers boundary, match and mismatch branches.
func TestHasPrefixAt(t *testing.T) {
	t.Parallel()
	r := []rune("operator://x")
	if !hasPrefixAt(r, 0, "operator://") {
		t.Fatal("prefix must match at offset 0")
	}
	if hasPrefixAt(r, 0, "operator://longer-than-input") {
		t.Fatal("overlong prefix must not match")
	}
	if hasPrefixAt(r, 1, "operator://") {
		t.Fatal("prefix must not match at a shifted offset")
	}
	if hasPrefixAt(r, 0, "zzzzzzzzzzz") {
		t.Fatal("different prefix must not match")
	}
}

// TestIsIDRune covers all accepted and rejected rune classes.
func TestIsIDRune(t *testing.T) {
	t.Parallel()
	for _, r := range []rune{'a', 'Z', '0', '9', '-', '_', '.', '中', 'é'} {
		if !isIDRune(r) {
			t.Fatalf("isIDRune(%q) must be true", r)
		}
	}
	for _, r := range []rune{' ', '(', ')', '∘', '⊗', '/', ':', '@', '\n'} {
		if isIDRune(r) {
			t.Fatalf("isIDRune(%q) must be false", r)
		}
	}
}

// TestParserGrammar covers the full EBNF: composition, product, dual,
// identity, parentheses, references and inline operands.
func TestParserGrammar(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		in   string
		kind nodeKind
	}{
		{"ε", nodeIdentity},
		{"c-delay-v1", nodeOperand},
		{"operator://m-repl-v1", nodeOperand},
		{"a ∘ b", nodeCompose},
		{"a ⊗ b", nodeProduct},
		{"a⁻¹", nodeInverse},
		{"(a ∘ b) ⊗ c⁻¹", nodeProduct},
		{"a ∘ b ⊗ c", nodeCompose}, // composition has the lowest precedence.
		{"a⁻¹⁻¹", nodeInverse},     // postfix dual is left-associative.
	} {
		tokens, err := lex(tc.in)
		requireOK(t, err)
		p := &exprParser{tokens: tokens}
		n, err := p.parse()
		requireOK(t, err)
		if n.kind != tc.kind {
			t.Fatalf("parse(%q) root = %v, want %v", tc.in, n.kind, tc.kind)
		}
	}
}

// TestParserPrecedence checks that ⊗ binds tighter than ∘.
func TestParserPrecedence(t *testing.T) {
	t.Parallel()
	tokens, err := lex("a ∘ b ⊗ c")
	requireOK(t, err)
	n, err := (&exprParser{tokens: tokens}).parse()
	requireOK(t, err)
	if n.kind != nodeCompose || n.left.kind != nodeOperand || n.right.kind != nodeProduct {
		t.Fatalf("precedence drift: %+v", n)
	}
}

// TestParserReferenceFlagging distinguishes operator:// refs from inline ids.
func TestParserReferenceFlagging(t *testing.T) {
	t.Parallel()
	refTokens, err := lex("operator://foo")
	requireOK(t, err)
	refNode, err := (&exprParser{tokens: refTokens}).parse()
	requireOK(t, err)
	if !refNode.ref || refNode.text != "foo" {
		t.Fatalf("reference flag drift: %+v", refNode)
	}
	bareTokens, err := lex("foo")
	requireOK(t, err)
	bareNode, err := (&exprParser{tokens: bareTokens}).parse()
	requireOK(t, err)
	if bareNode.ref || bareNode.text != "foo" {
		t.Fatalf("inline operand flag drift: %+v", bareNode)
	}
}

// TestParserErrors covers empty input, trailing tokens, missing operands,
// missing parenthesis and the nesting-depth guard.
func TestParserErrors(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		in   string
		want string
	}{
		{"", RuleExprSyntax},
		{"a b", RuleExprSyntax},
		{"∘ a", RuleExprSyntax},
		{"a ∘", RuleExprSyntax},
		{"(a", RuleExprSyntax},
		{"a ⊗ ⊗ b", RuleExprSyntax},
	} {
		tokens, err := lex(tc.in)
		requireOK(t, err)
		_, err = (&exprParser{tokens: tokens}).parse()
		requireRule(t, err, tc.want)
	}
	deep := strings.Repeat("(", MaxExprDepth+1) + "a" + strings.Repeat(")", MaxExprDepth+1)
	tokens, err := lex(deep)
	requireOK(t, err)
	_, err = (&exprParser{tokens: tokens}).parse()
	requireRule(t, err, RuleExprDepth)
}

// TestParserNextAtEOF verifies next never walks past the final token.
func TestParserNextAtEOF(t *testing.T) {
	t.Parallel()
	p := &exprParser{tokens: []token{{kind: tokEOF, pos: 0}}}
	for i := 0; i < 3; i++ {
		if tk := p.next(); tk.kind != tokEOF {
			t.Fatalf("next at EOF returned %v", tk.kind)
		}
	}
	if p.peek().kind != tokEOF {
		t.Fatal("peek at EOF drifted")
	}
}
