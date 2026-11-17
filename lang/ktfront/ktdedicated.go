/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file carries the dedicated() helper and the string fragment
// reductions of ktfront. Every helper here is called from the parser while it
// lowers a concrete Kotlin construct; the emitted node is anchored on that
// construct's real token position, so the kind is produced by syntax-node
// semantics rather than by switching on a detached token's spelling.

// dedicated builds one semantic node for a construct the parser just reduced.
// The anchor is the real token position of the construct; the span is
// zero-width, which mirrors pyfront.qualifiedNameAs: it can never overlap an
// already projected sibling, so View.OrderChildren keeps it while the
// construct's own lexical projection stays byte-exact.
func (m *ktMap) dedicated(kind air.NodeKind, at frontkit.Token,
	text, note string) air.Node {
	pos := air.NewPosition(at.Line, at.Column, at.Offset, at.Offset)
	node := air.NewNodeAt(kind, pos, pos, text)
	return m.annotate(node, air.SemNote, note)
}

// quoteSpan returns the content bounds of one Kotlin string token. Raw strings
// open with a triple quote; plain and template strings with a single quote.
// The delimiter width always equals contentStart, so callers derive it from
// there instead of a fourth result.
func ktQuoteSpan(text string) (contentStart, contentEnd int, ok bool) {
	if len(text) < 2 {
		return 0, 0, false
	}
	if len(text) >= 3 && text[:3] == `"""` && len(text) >= 6 {
		return 3, len(text) - 3, true
	}
	if text[0] != '"' {
		return 0, 0, false
	}
	return 1, len(text) - 1, true
}

// stringFragments models the parts of a string literal the parser just
// lowered: its content and its closing delimiter. Each part is a real span
// inside the literal token, attached as a child of the string leaf so the
// byte-exact projection is untouched. Template strings carry interpolation
// nodes inside their content, so callers skip the content fragment there and
// use delimiterFragment instead.
func (m *ktMap) stringFragments(tok frontkit.Token) []air.Node {
	text := tok.Text
	contentStart, contentEnd, ok := ktQuoteSpan(text)
	if !ok || contentEnd < contentStart {
		return nil
	}
	width := contentStart
	parts := make([]air.Node, 0, 2)
	part := func(from, to int) {
		if to <= from {
			return
		}
		node := frontkit.Token{
			Line: tok.Line, Column: tok.Column + from,
			Offset: tok.Offset + from, EndOffset: tok.Offset + to,
		}
		start, end := node.Span()
		child := air.NewNodeAt(air.KindKtStringFrag, start, end, text[from:to])
		parts = append(parts, m.annotate(child, air.SemNote,
			m.note("kt-string-content")))
	}
	part(contentStart, contentEnd)
	part(len(text)-width, len(text))
	return parts
}

// delimiterFragment models only the closing delimiter of a string token; used
// for template strings whose content is interleaved with interpolation nodes.
func (m *ktMap) delimiterFragment(tok frontkit.Token) []air.Node {
	text := tok.Text
	width, contentEnd, ok := ktQuoteSpan(text)
	if !ok || contentEnd < width {
		return nil
	}
	part := frontkit.Token{
		Line: tok.Line, Column: tok.Column + contentEnd,
		Offset: tok.Offset + contentEnd, EndOffset: tok.Offset + len(text),
	}
	start, end := part.Span()
	node := air.NewNodeAt(air.KindKtStringFrag, start, end, text[contentEnd:])
	return []air.Node{m.annotate(node, air.SemNote,
		m.note("kt-string-delimiter"))}
}

// infixNodeMarked builds an infix node with additional dedicated markers
// anchored between the operands, mirroring pyfront's operator anchoring.
func (m *ktMap) infixNodeMarked(kind air.NodeKind, left air.Node,
	tok frontkit.Token, right air.Node, marks ...air.Node) air.Node {
	if left == nil || right == nil {
		return m.leafKind(tok, kind)
	}
	lp := left.Pos()
	end := right.End()
	pos := air.NewPosition(lp.Line, lp.Column, lp.Offset, end.Offset)
	children := append([]air.Node{left, right}, marks...)
	if m.view != nil {
		children = m.view.OrderChildren(children)
	}
	return frontkit.BindOperator(kind, pos, end, tok, children...)
}
