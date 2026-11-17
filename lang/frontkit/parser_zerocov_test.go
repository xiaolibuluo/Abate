/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"testing"

	"abate/lang/air"
)

// TestPrattViewReturnsBoundView proves View exposes the exact live view bound
// to the parser at construction, stable across calls and shared by reference.
func TestPrattViewReturnsBoundView(t *testing.T) {
	t.Parallel()
	tokens := tokenList([]TokenKind{TokenIdent}, []string{"a"})
	stream, err := NewTokenStream(tokens)
	if err != nil {
		t.Fatalf("NewTokenStream: %v", err)
	}
	view := NewView(SourceFromTokens(tokens), "view.fk")
	parser, err := NewPrattParser(stream, view)
	if err != nil {
		t.Fatalf("NewPrattParser: %v", err)
	}
	got := parser.View()
	if got == nil {
		t.Fatal("View returned nil for a parser built with a view")
	}
	if got != view {
		t.Fatal("View must return the exact view bound at construction")
	}
	if got.Path() != "view.fk" {
		t.Fatalf("view path = %q, want %q", got.Path(), "view.fk")
	}
	if parser.View() != got {
		t.Fatal("View must return the same view on every call")
	}
	bound := got.Synthetic(air.KindIdent, "bound")
	got.Bind(0, bound)
	stored, ok := parser.View().Lookup(0)
	if !ok || stored != bound {
		t.Fatalf("binding through View was not observable: %v, %v", stored, ok)
	}
}

// TestPrattViewNilReceiver proves the nil-receiver guard reports a nil view
// instead of dereferencing an absent parser.
func TestPrattViewNilReceiver(t *testing.T) {
	t.Parallel()
	var absent *PrattParser
	if absent.View() != nil {
		t.Fatal("a nil parser must report a nil view")
	}
}
