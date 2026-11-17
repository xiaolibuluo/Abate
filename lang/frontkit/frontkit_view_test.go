/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"reflect"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

func leafNode(text string, start, end int) air.Node {
	return air.NewNodeAt(air.KindIdent,
		air.NewPosition(1, start+1, start, end),
		air.NewPosition(1, end+1, end, end), text)
}

func viewTokens(start, end int) (first, second Token) {
	return Token{Kind: TokenIdent, Offset: start, EndOffset: end,
			Line: 1, Column: start + 1},
		Token{Kind: TokenEOF, Offset: end, EndOffset: end,
			Line: 1, Column: end + 1}
}

func TestViewPositionSpanAndBounds(t *testing.T) {
	t.Parallel()
	view := NewView([]byte("let a\r\nlet b\r\n"), "v.fk")
	tok := Token{Kind: TokenIdent, Text: "a", Offset: 4, EndOffset: 5,
		Line: 1, Column: 5}
	if got := view.Position(tok); got.Canonical() != "1/5/4/5" {
		t.Fatalf("Position = %s", got.Canonical())
	}
	start, end := viewTokens(0, 14)
	span := view.Span(start, end)
	if span.Offset != 0 || span.EndOffset != 14 {
		t.Fatalf("Span = %s", span.Canonical())
	}
	clamped := view.Span(Token{Offset: 99, EndOffset: 999, Line: -5, Column: -7},
		Token{EndOffset: 999})
	if clamped.Offset != 14 || clamped.EndOffset != 14 ||
		clamped.Line != 0 || clamped.Column != 0 {
		t.Fatalf("clamped Span = %s", clamped.Canonical())
	}
	var nilView *View
	if nilView.Position(tok).Canonical() != "0/0/0/0" ||
		nilView.Span(start, end).Canonical() != "0/0/0/0" ||
		nilView.Path() != "" || nilView.Size() != 0 ||
		len(nilView.OrderChildren([]air.Node{leafNode("x", 0, 1)})) != 0 {
		t.Fatal("nil view guards are inconsistent")
	}
}

func TestViewNewNodeAtIntervalCoverage(t *testing.T) {
	t.Parallel()
	view := NewView(make([]byte, 12), "v.fk")
	children := []air.Node{
		leafNode("d", 9, 12), nil, leafNode("b", 3, 6),
		leafNode("a", 0, 3), leafNode("c", 6, 9),
	}
	ordered := view.OrderChildren(children)
	if len(ordered) != 4 {
		t.Fatalf("OrderChildren kept %d children, want 4", len(ordered))
	}
	previous := 0
	for i, child := range ordered {
		pos := child.Pos()
		if pos.Offset != previous {
			t.Fatalf("child %d starts at %d, want a gapless boundary at %d",
				i, pos.Offset, previous)
		}
		if pos.Offset < 0 || pos.EndOffset > 12 || pos.EndOffset <= pos.Offset {
			t.Fatalf("child %d interval %s escapes the source", i, pos.String())
		}
		previous = pos.EndOffset
	}
	if previous != 12 {
		t.Fatalf("ordered children end at %d, want 12", previous)
	}
	start, end := viewTokens(0, 12)
	root := view.NewNodeAt(air.KindBlock, start, end, "", children...)
	if err := air.CheckIntervals(root, 12); err != nil {
		t.Fatalf("AIR interval check: %v", err)
	}
	if len(root.Children()) != 4 {
		t.Fatalf("root has %d children, want 4", len(root.Children()))
	}
	rootStart, rootEnd := air.Span(root)
	if rootStart != 0 || rootEnd != 12 {
		t.Fatalf("root interval [%d,%d) does not cover the file", rootStart, rootEnd)
	}
}

func TestViewOrderChildrenDropsInvalidAndOverlapping(t *testing.T) {
	t.Parallel()
	view := NewView(make([]byte, 12), "v.fk")
	input := []air.Node{
		leafNode("a", 0, 3), leafNode("b", 3, 6), leafNode("c", 6, 9),
		leafNode("outside", 10, 20), leafNode("overlap", 2, 5),
		air.NewNodeAt(air.KindRegion,
			air.NewPosition(-1, 0, 0, 2), air.NewPosition(0, 0, 2, 2), "bad"),
	}
	ordered := view.OrderChildren(input)
	if len(ordered) != 3 {
		t.Fatalf("OrderChildren kept %d children, want 3", len(ordered))
	}
	want := []string{"a", "b", "c"}
	for i, child := range ordered {
		if air.AttributeOf(child) != want[i] {
			t.Fatalf("child %d = %q, want %q", i, air.AttributeOf(child), want[i])
		}
	}
	if len(view.OrderChildren(nil)) != 0 {
		t.Fatal("nil children must produce an empty slice")
	}
}

func TestViewAnnotatePreservesNodeID(t *testing.T) {
	t.Parallel()
	view := NewView(make([]byte, 4), "v.fk")
	node := leafNode("abc", 0, 3)
	before := air.IDOf(node)
	if !view.Annotate(node, air.SemNote, "hello") {
		t.Fatal("Annotate returned false on a package-built node")
	}
	if air.IDOf(node) != before {
		t.Fatal("annotation changed the NodeID")
	}
	if value := node.Sem(air.SemNote); value != "hello" {
		t.Fatalf("SemNote = %v, want hello", value)
	}
	if view.Annotate(nil, air.SemNote, "x") ||
		view.Annotate(node, air.SemInvalid, "x") {
		t.Fatal("Annotate accepted a nil node or invalid key")
	}
}

func TestViewSyntheticPartialAndBindings(t *testing.T) {
	t.Parallel()
	view := NewView(make([]byte, 4), "v.fk")
	partial := view.Synthetic(air.KindRegion, "missing")
	if partial.Kind() != air.KindRegion || air.AttributeOf(partial) != "missing" {
		t.Fatal("synthetic node lost its kind or text")
	}
	if partial.Sem(air.SemPartial) != "true" {
		t.Fatal("synthetic node is not marked partial")
	}
	if partial.Pos().Offset != 0 || partial.Pos().EndOffset != partial.Pos().Offset {
		t.Fatal("synthetic node is not zero-width")
	}
	node := leafNode("a", 0, 1)
	view.Bind(3, node)
	view.Bind(4, nil)
	if got, ok := view.Lookup(3); !ok || got != node {
		t.Fatal("Bind/Lookup round trip failed")
	}
	if _, ok := view.Lookup(99); ok {
		t.Fatal("Lookup invented a binding")
	}
	if view.NodeCount() != 1 {
		t.Fatalf("NodeCount = %d, want 1", view.NodeCount())
	}
}

func TestViewPackageImportsAndNilGuards(t *testing.T) {
	t.Parallel()
	view := NewView(make([]byte, 4), "v.fk")
	view.AddImport("z")
	view.AddImport("a")
	view.AddImport("a")
	view.AddImport("")
	imports := view.Imports()
	if !reflect.DeepEqual(imports, []string{"a", "z"}) {
		t.Fatalf("Imports = %v, want [a z]", imports)
	}
	imports[0] = "mutated"
	if view.Imports()[0] != "a" {
		t.Fatal("Imports returned a shared slice")
	}
	view.SetPackageName("pkg")
	if view.PackageName() != "pkg" || view.PackageRef().Name != "pkg" ||
		len(view.PackageRef().Imports) != 2 {
		t.Fatal("PackageRef does not reflect the view")
	}
	var nilView *View
	if nilView.PackageName() != "" || len(nilView.Imports()) != 0 ||
		nilView.PackageRef().Name != "" || nilView.NodeCount() != 0 ||
		nilView.Annotate(nil, air.SemNote, "x") {
		t.Fatal("nil view guards are inconsistent")
	}
	if node := nilView.NewNodeAt(air.KindBlock, Token{}, Token{}, "x"); node == nil {
		t.Fatal("nil view must still build a node")
	}
	if node := nilView.Synthetic(air.KindRegion, "x"); node == nil {
		t.Fatal("nil view must still build a synthetic node")
	}
	if _, ok := nilView.Lookup(0); ok {
		t.Fatal("nil view invented a binding")
	}
}

func TestSourceFromTokensExtent(t *testing.T) {
	t.Parallel()
	tokens := tokenList([]TokenKind{TokenIdent, TokenOp, TokenIdent},
		[]string{"ab", "+", "cd"})
	src := SourceFromTokens(tokens)
	if len(src) != tokens[len(tokens)-1].EndOffset {
		t.Fatalf("extent = %d, want %d", len(src), tokens[len(tokens)-1].EndOffset)
	}
	if string(src[0:2]) != "ab" || string(src[3:4]) != "+" ||
		string(src[5:7]) != "cd" || src[2] != 0 {
		t.Fatalf("reconstructed source = %q, gaps not preserved", src)
	}
	malformed := append([]Token(nil), tokens...)
	malformed = append(malformed, Token{Offset: -1, EndOffset: 999})
	if len(SourceFromTokens(malformed)) != len(src) {
		t.Fatal("malformed token changed the reconstructed extent")
	}
	if len(SourceFromTokens(nil)) != 0 {
		t.Fatal("nil tokens must produce an empty buffer")
	}
}

func TestViewDeterminism100(t *testing.T) {
	t.Parallel()
	determinism.MustMatch100(t, 42, func() []byte {
		view := NewView(make([]byte, 12), "v.fk")
		start, end := viewTokens(0, 12)
		root := view.NewNodeAt(air.KindBlock, start, end, "",
			leafNode("d", 9, 12), leafNode("a", 0, 3), leafNode("c", 6, 9))
		digest := determinism.HashState(root)
		return digest[:]
	})
}
