/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// ktLeafPart is one lexical token specification for the leafNodes fixtures.
type ktLeafPart struct {
	kind frontkit.TokenKind
	text string
}

// ktLeafTokens builds a token slice whose byte offsets follow a fixed stride,
// so every expected leaf span can be derived from its token index.
func ktLeafTokens(parts []ktLeafPart) []frontkit.Token {
	out := make([]frontkit.Token, 0, len(parts))
	for i, p := range parts {
		out = append(out, frontkit.Token{
			Kind: p.kind, Text: p.text,
			Offset: 2 * i, EndOffset: 2*i + 1,
			Line: 1, Column: 2*i + 1,
		})
	}
	return out
}

// TestKtLeafNodesProjection proves leafNodes skips exactly the four trivia
// kinds, projects every other token to a byte-exact AIR leaf, and honours the
// half-open [lo,hi) interval.
func TestKtLeafNodesProjection(t *testing.T) {
	t.Parallel()
	type want struct {
		text string
		kind air.NodeKind
		idx  int
	}
	cases := []struct {
		name  string
		parts []ktLeafPart
		lo    int
		hi    int
		wants []want
	}{
		{
			name: "mixed-trivia",
			parts: []ktLeafPart{
				{frontkit.TokenIdent, "alpha"},
				{frontkit.TokenNewline, "\n"},
				{frontkit.TokenNum, "42"},
				{frontkit.TokenEOF, ""},
				{frontkit.TokenIndent, "  "},
				{frontkit.TokenKeyword, "class"},
				{frontkit.TokenDedent, ""},
				{frontkit.TokenComment, "// note"},
				{frontkit.TokenStr, "\"s\""},
				{frontkit.TokenPunct, ";"},
				{frontkit.TokenOp, "!!"},
				{frontkit.TokenIllegal, "\x01"},
			},
			lo: 0, hi: 12,
			wants: []want{
				{"alpha", air.KindIdent, 0},
				{"42", air.KindNumLit, 2},
				{"class", air.KindIdent, 5},
				{"// note", air.KindComment, 7},
				{"\"s\"", air.KindStrLit, 8},
				// 标点不再贴 Delimiter 伪标签，标记为零碎恢复节点。
				{";", air.KindPartialNode, 9},
				{"!!", air.KindUnary, 10},
				{"\x01", air.KindPartialNode, 11},
			},
		},
		{
			name: "all-trivia",
			parts: []ktLeafPart{
				{frontkit.TokenNewline, "\n"},
				{frontkit.TokenEOF, ""},
				{frontkit.TokenIndent, " "},
				{frontkit.TokenDedent, ""},
			},
			lo: 0, hi: 4,
			wants: nil,
		},
		{
			name: "keyword-kinds",
			parts: []ktLeafPart{
				{frontkit.TokenKeyword, "true"},
				{frontkit.TokenKeyword, "null"},
				{frontkit.TokenKeyword, "class"},
				{frontkit.TokenChar, "'x'"},
				{frontkit.TokenIdent, "v"},
				{frontkit.TokenPunct, ";"},
			},
			lo: 0, hi: 6,
			wants: []want{
				{"true", air.KindBoolLit, 0},
				{"null", air.KindNilLit, 1},
				{"class", air.KindIdent, 2},
				{"'x'", air.KindCharLit, 3},
				{"v", air.KindIdent, 4},
				{";", air.KindPartialNode, 5},
			},
		},
		{
			name: "string-flavors",
			parts: []ktLeafPart{
				{frontkit.TokenStr, "`id`"},
				{frontkit.TokenStr, `"""raw"""`},
				{frontkit.TokenStr, "\"plain\""},
			},
			lo: 0, hi: 3,
			wants: []want{
				{"`id`", air.KindIdent, 0},
				{`"""raw"""`, air.KindStrLit, 1},
				{"\"plain\"", air.KindStrLit, 2},
			},
		},
		{
			name: "half-open-subrange",
			parts: []ktLeafPart{
				{frontkit.TokenIdent, "a"},
				{frontkit.TokenNewline, "\n"},
				{frontkit.TokenNum, "2"},
				{frontkit.TokenDedent, ""},
				{frontkit.TokenStr, "\"z\""},
				{frontkit.TokenIdent, "b"},
			},
			lo: 1, hi: 5,
			wants: []want{
				{"2", air.KindNumLit, 2},
				{"\"z\"", air.KindStrLit, 4},
			},
		},
		{
			name: "empty-interval",
			parts: []ktLeafPart{
				{frontkit.TokenIdent, "a"},
				{frontkit.TokenNum, "1"},
			},
			lo: 1, hi: 1,
			wants: nil,
		},
	}
	expr := newKtExpr(nil, newKtMap(nil, nil))
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tokens := ktLeafTokens(tc.parts)
			got := expr.leafNodes(tokens, tc.lo, tc.hi)
			if len(got) != len(tc.wants) {
				t.Fatalf("leafNodes returned %d leaves, want %d",
					len(got), len(tc.wants))
			}
			for i, w := range tc.wants {
				node := got[i]
				if node == nil {
					t.Fatalf("leaf %d is nil", i)
				}
				if text := air.AttributeOf(node); text != w.text {
					t.Fatalf("leaf %d text = %q, want %q", i, text, w.text)
				}
				if kind := node.Kind(); kind != w.kind {
					t.Fatalf("leaf %d kind = %v, want %v", i, kind, w.kind)
				}
				start, end := air.Span(node)
				if start != 2*w.idx || end != 2*w.idx+1 {
					t.Fatalf("leaf %d span = [%d,%d), want [%d,%d)",
						i, start, end, 2*w.idx, 2*w.idx+1)
				}
			}
		})
	}
}

// TestKtLeafNodesNullSafeLeafKeepsNote proves the null-safe operator leaf
// carries both its dedicated AIR kind and the semantic note from the table.
func TestKtLeafNodesNullSafeLeafKeepsNote(t *testing.T) {
	t.Parallel()
	tokens := ktLeafTokens([]ktLeafPart{
		{frontkit.TokenOp, "!!"},
		{frontkit.TokenIdent, "v"},
	})
	expr := newKtExpr(nil, newKtMap(nil, nil))
	got := expr.leafNodes(tokens, 0, len(tokens))
	if len(got) != 2 {
		t.Fatalf("leafNodes returned %d leaves, want 2", len(got))
	}
	if got[0].Kind() != air.KindUnary {
		t.Fatalf("operator leaf kind = %v, want %v", got[0].Kind(), air.KindUnary)
	}
	if air.SemanticsOf(got[0])[air.SemNote] == "" {
		t.Fatal("the null-safe operator leaf lost its semantic note")
	}
	if air.SemanticsOf(got[1])[air.SemNote] != "" {
		t.Fatal("a plain identifier leaf must not carry a semantic note")
	}
}

// TestKtLeafNodesIllegalLeafCarriesPartialMarker proves the fallback leaf for
// an illegal token keeps the SemPartial recovery marker from the mapper.
func TestKtLeafNodesIllegalLeafCarriesPartialMarker(t *testing.T) {
	t.Parallel()
	tokens := ktLeafTokens([]ktLeafPart{
		{frontkit.TokenIllegal, "\x01"},
		{frontkit.TokenIdent, "ok"},
	})
	expr := newKtExpr(nil, newKtMap(nil, nil))
	got := expr.leafNodes(tokens, 0, len(tokens))
	if len(got) != 2 {
		t.Fatalf("leafNodes returned %d leaves, want 2", len(got))
	}
	if air.SemanticsOf(got[0])[air.SemPartial] != "true" {
		t.Fatal("the illegal token leaf lost its SemPartial marker")
	}
	if air.SemanticsOf(got[1])[air.SemPartial] != "" {
		t.Fatal("a plain identifier leaf must not be marked partial")
	}
}
