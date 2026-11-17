/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// TestRsMapperBlockLowersBracedExpression proves block lowers braced children
// to KindBlock with an empty attribute, a byte-exact span and ordered children.
func TestRsMapperBlockLowersBracedExpression(t *testing.T) {
	t.Parallel()
	tokens := rsTokenize(t, "{ a b c }\n")
	open := tokens[0]
	if open.Text != "{" {
		t.Fatalf("first token = %q, want an opening brace", open.Text)
	}
	var closeTok frontkit.Token
	for _, tok := range tokens {
		if tok.Text == "}" {
			closeTok = tok
		}
	}
	if closeTok.Text != "}" {
		t.Fatal("token stream is missing the closing brace")
	}
	view := frontkit.NewView(frontkit.SourceFromTokens(tokens), "block.rs")
	mapper := newRsMapper(view, defaultRsConfig())
	children := make([]air.Node, 0, 3)
	for _, tok := range tokens[1:] {
		if tok.Kind == frontkit.TokenIdent {
			children = append(children, mapper.lexical(tok))
		}
	}
	if len(children) != 3 {
		t.Fatalf("built %d identifier children, want 3", len(children))
	}
	cases := []struct {
		name string
		kids []air.Node
	}{
		{"forward-order", children},
		{"reversed-order", []air.Node{children[2], children[1], children[0]}},
		{"no-children", nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := mapper.block(open, closeTok, tc.kids...)
			if got.Kind() != air.KindBlock {
				t.Fatalf("block kind = %v, want %v", got.Kind(), air.KindBlock)
			}
			if attr := air.AttributeOf(got); attr != "" {
				t.Fatalf("block attribute = %q, want empty", attr)
			}
			start, end := air.Span(got)
			if start != open.Offset || end != closeTok.EndOffset {
				t.Fatalf("block span = [%d,%d), want [%d,%d)",
					start, end, open.Offset, closeTok.EndOffset)
			}
			if got.Pos().Line != open.Line || got.Pos().Column != open.Column {
				t.Fatalf("block position = %+v, want line %d column %d",
					got.Pos(), open.Line, open.Column)
			}
			ordered := got.Children()
			if len(ordered) != len(tc.kids) {
				t.Fatalf("block kept %d children, want %d", len(ordered), len(tc.kids))
			}
			for i := range ordered {
				if air.AttributeOf(ordered[i]) != air.AttributeOf(children[i]) {
					t.Fatalf("child %d = %q, want %q", i,
						air.AttributeOf(ordered[i]), air.AttributeOf(children[i]))
				}
			}
		})
	}
}
