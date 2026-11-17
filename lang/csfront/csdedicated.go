/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file carries the dedicated() helper and the string/char fragment
// reductions of csfront. Every helper here is called from the parser while it
// lowers a concrete C# construct; the emitted node is anchored on that
// construct's real token position, so the kind is produced by syntax-node
// semantics rather than by switching on a detached token's spelling.

// dedicated builds one semantic node for a construct the parser just reduced.
// The anchor is the real token position of the construct; the span is
// zero-width, which mirrors pyfront.qualifiedNameAs: it can never overlap an
// already projected sibling, so View.OrderChildren keeps it while the
// construct's own lexical projection stays byte-exact.
func (m *csMap) dedicated(kind air.NodeKind, at frontkit.Token,
	text, note string) air.Node {
	pos := air.NewPosition(at.Line, at.Column, at.Offset, at.Offset)
	node := air.NewNodeAt(kind, pos, pos, text)
	return m.annotate(node, air.SemNote, note)
}

// quoteWidth returns the width of the opening quote at start of text.
func quoteWidth(text string, at int) int {
	if at >= len(text) || text[at] != '"' {
		return 0
	}
	width := 1
	for at+width < len(text) && width < 3 && text[at+width] == '"' {
		width++
	}
	return width
}

// stringFragments models the parts of a string literal the parser just
// lowered: its content, its escape sequences and its closing delimiter. Each
// part is a real span inside the literal token, attached as a child of the
// string leaf so the byte-exact projection is untouched.
func (m *csMap) stringFragments(tok frontkit.Token) []air.Node {
	text := tok.Text
	verbatim := strings.HasPrefix(text, "@")
	quoteAt := strings.IndexByte(text, '"')
	if quoteAt < 0 {
		return nil
	}
	width := quoteWidth(text, quoteAt)
	contentStart := quoteAt + width
	contentEnd := len(text) - width
	if contentEnd < contentStart {
		return nil
	}
	parts := make([]air.Node, 0, 3)
	part := func(from, to int, kind air.NodeKind, note string) {
		if to <= from {
			return
		}
		node := frontkit.Token{
			Line: tok.Line, Column: tok.Column + from,
			Offset: tok.Offset + from, EndOffset: tok.Offset + to,
		}
		start, end := node.Span()
		child := air.NewNodeAt(kind, start, end, text[from:to])
		parts = append(parts, m.annotate(child, air.SemNote, m.note(note)))
	}
	part(contentStart, contentEnd, air.KindCsStringFrag, "cs-string-content")
	if !verbatim {
		for i := contentStart; i+1 < contentEnd; i++ {
			if text[i] != '\\' {
				continue
			}
			stop := i + 2
			if text[i+1] == 'u' || text[i+1] == 'U' {
				stop = i + 6
			}
			if stop > contentEnd {
				stop = contentEnd
			}
			part(i, stop, air.KindCsStringFrag, "cs-escape")
			i = stop - 1
		}
	}
	part(contentEnd, len(text), air.KindCsStringFrag, "cs-string-delimiter")
	return parts
}

// delimiterFragment models the closing delimiter of a string literal.
func (m *csMap) delimiterFragment(tok frontkit.Token) []air.Node {
	text := tok.Text
	quoteAt := strings.IndexByte(text, '"')
	if quoteAt < 0 {
		return nil
	}
	width := quoteWidth(text, quoteAt)
	contentEnd := len(text) - width
	if contentEnd < quoteAt+width {
		return nil
	}
	part := frontkit.Token{
		Line: tok.Line, Column: tok.Column + contentEnd,
		Offset: tok.Offset + contentEnd, EndOffset: tok.Offset + len(text),
	}
	start, end := part.Span()
	node := air.NewNodeAt(air.KindCsStringFrag, start, end, text[contentEnd:])
	return []air.Node{m.annotate(node, air.SemNote, m.note("cs-string-delimiter"))}
}

// charBody models the content of a character literal the parser just lowered.
func (m *csMap) charBody(tok frontkit.Token) []air.Node {
	text := tok.Text
	if len(text) < 2 || text[0] != '\'' || text[len(text)-1] != '\'' {
		return nil
	}
	if len(text) == 2 {
		return nil
	}
	body := frontkit.Token{
		Line: tok.Line, Column: tok.Column + 1,
		Offset: tok.Offset + 1, EndOffset: tok.Offset + len(text) - 1,
	}
	start, end := body.Span()
	node := air.NewNodeAt(air.KindCsCharBody, start, end, text[1:len(text)-1])
	return []air.Node{m.annotate(node, air.SemNote, m.note("cs-char-body"))}
}

// infixNodeMarked builds an infix node with additional dedicated markers
// anchored between the operands, mirroring pyfront's operator anchoring.
func (m *csMap) infixNodeMarked(kind air.NodeKind, left air.Node,
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

// unaryNodeMarked builds a prefix node with additional dedicated markers
// anchored at the operator token the parser just consumed.
func (m *csMap) unaryNodeMarked(kind air.NodeKind, tok frontkit.Token,
	operand air.Node, marks ...air.Node) air.Node {
	if operand == nil {
		return m.leafKind(tok, kind)
	}
	start, _ := tok.Span()
	end := operand.End()
	pos := air.NewPosition(start.Line, start.Column, start.Offset, end.Offset)
	children := make([]air.Node, 0, len(marks)+2)
	if leaf := m.literal(tok); leaf != nil {
		children = append(children, leaf)
	}
	children = append(children, operand)
	children = append(children, marks...)
	if m.view != nil {
		children = m.view.OrderChildren(children)
	}
	return frontkit.BindOperator(kind, pos, end, tok, children...)
}
