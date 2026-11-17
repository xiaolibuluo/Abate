/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// pyAsyncExprAfterKeyword builds a real pyExpr whose stream is positioned just
// past the leading async token, mirroring the parseClauses call site.
func pyAsyncExprAfterKeyword(t *testing.T, src string) (*pyExpr, frontkit.Token) {
	t.Helper()
	tokens := pyTokenize(t, src)
	parser, err := newPyParser(tokens, "async.py")
	if err != nil {
		t.Fatalf("newPyParser: %v", err)
	}
	asyncTok, ok := parser.stream.Peek()
	if !ok || asyncTok.Text != "async" {
		t.Fatalf("leading token = %+v (ok=%v), want async", asyncTok, ok)
	}
	if _, more := parser.stream.Advance(); !more {
		t.Fatal("stream ended before the async token was consumed")
	}
	return parser.expr, asyncTok
}

// TestPyParseAsyncClauseBranches proves parseAsyncClause reduces a real async
// for clause with the SemAsync marker, and degrades to a plain clause anchored
// on the async token when the next token is not for.
func TestPyParseAsyncClauseBranches(t *testing.T) {
	t.Parallel()
	hasText := func(root air.Node, text string) bool {
		for _, node := range pyNodes(root) {
			if air.AttributeOf(node) == text {
				return true
			}
		}
		return false
	}
	cases := []struct {
		name         string
		src          string
		wantKind     air.NodeKind
		wantAttr     string
		wantAsync    string
		wantBodyText string
	}{
		{"for-with-in", "async for x in xs\n", air.KindRangeFor, "for", "true", "xs"},
		{"for-without-in", "async for x\n", air.KindRangeFor, "for", "true", "x"},
		{"if-fallback", "async if cond\n", air.KindIf, "async", "", ""},
		{"newline-fallback", "async\n", air.KindIf, "async", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			expr, asyncTok := pyAsyncExprAfterKeyword(t, tc.src)
			node := expr.parseAsyncClause(asyncTok)
			if node == nil {
				t.Fatal("parseAsyncClause returned nil")
			}
			if kind := node.Kind(); kind != tc.wantKind {
				t.Fatalf("clause kind = %v, want %v", kind, tc.wantKind)
			}
			if attr := air.AttributeOf(node); attr != tc.wantAttr {
				t.Fatalf("clause attribute = %q, want %q", attr, tc.wantAttr)
			}
			sem := air.SemanticsOf(node)
			if sem[air.SemAsync] != tc.wantAsync {
				t.Fatalf("SemAsync = %q, want %q", sem[air.SemAsync], tc.wantAsync)
			}
			if sem[air.SemNote] != "comprehension-clause" {
				t.Fatalf("SemNote = %q, want comprehension-clause", sem[air.SemNote])
			}
			if tc.wantBodyText != "" && !hasText(node, tc.wantBodyText) {
				t.Fatalf("clause lost the body token %q", tc.wantBodyText)
			}
		})
	}
}

// TestPyParseAsyncClauseAtStreamEnd proves the exhausted-stream Peek branch
// degrades to a clause anchored on the async token instead of panicking.
func TestPyParseAsyncClauseAtStreamEnd(t *testing.T) {
	t.Parallel()
	tokens := pyTokenize(t, "async\n")
	parser, err := newPyParser(tokens, "async.py")
	if err != nil {
		t.Fatalf("newPyParser: %v", err)
	}
	asyncTok, ok := parser.stream.Peek()
	if !ok || asyncTok.Text != "async" {
		t.Fatalf("leading token = %+v (ok=%v), want async", asyncTok, ok)
	}
	for {
		if _, more := parser.stream.Advance(); !more {
			break
		}
	}
	if next, more := parser.stream.Peek(); more {
		t.Fatalf("stream still has a token at exhaustion: %+v", next)
	}
	node := parser.expr.parseAsyncClause(asyncTok)
	if node == nil {
		t.Fatal("parseAsyncClause returned nil on an exhausted stream")
	}
	if node.Kind() != air.KindIf {
		t.Fatalf("clause kind = %v, want %v", node.Kind(), air.KindIf)
	}
	if air.AttributeOf(node) != "async" {
		t.Fatalf("clause attribute = %q, want async", air.AttributeOf(node))
	}
	sem := air.SemanticsOf(node)
	if sem[air.SemAsync] != "" {
		t.Fatalf("exhausted clause must not be async, got %q", sem[air.SemAsync])
	}
	if sem[air.SemNote] != "comprehension-clause" {
		t.Fatalf("SemNote = %q, want comprehension-clause", sem[air.SemNote])
	}
}
