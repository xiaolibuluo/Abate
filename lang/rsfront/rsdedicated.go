/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"strconv"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file carries the dedicated() helper and the string fragment reduction
// of rsfront. Every helper here is called from the parser while it lowers a
// concrete Rust construct; the emitted node is anchored on that construct's
// real token position, so the kind is produced by syntax-node semantics
// rather than by switching on a detached token's spelling.

// dedicated builds one semantic node for a construct the parser just reduced.
// The anchor is the real token position of the construct; the span is
// zero-width, which mirrors pyfront.qualifiedNameAs: it can never overlap an
// already projected sibling, so View.OrderChildren keeps it while the
// construct's own lexical projection stays byte-exact.
func (m *rsMapper) dedicated(kind air.NodeKind, at frontkit.Token,
	text, noteName string) air.Node {
	pos := air.NewPosition(at.Line, at.Column, at.Offset, at.Offset)
	node := air.NewNode(kind, pos, text)
	return m.annotate(node, air.SemNote, m.note(noteName))
}

// rsQuoteSpan returns the content bounds of one Rust string token: raw strings
// r#".."# keep their hash suffix inside the delimiters, plain strings open at
// the first double quote.
func rsQuoteSpan(text string) (contentStart, contentEnd int, ok bool) {
	first := strings.IndexByte(text, '"')
	if first < 0 || first+1 >= len(text) {
		return 0, 0, false
	}
	last := strings.LastIndexByte(text, '"')
	if last <= first {
		return 0, 0, false
	}
	return first + 1, last, true
}

// stringFragments models the parts of a string literal the parser just
// lowered: its content, its escape sequences and its closing delimiter. Each
// part is a real span inside the literal token, attached as a child of the
// string leaf so the byte-exact projection is untouched.
func (m *rsMapper) stringFragments(tok frontkit.Token) []air.Node {
	text := tok.Text
	contentStart, contentEnd, ok := rsQuoteSpan(text)
	if !ok || contentEnd < contentStart {
		return nil
	}
	parts := make([]air.Node, 0, 3)
	part := func(from, to int, noteName string) {
		if to <= from {
			return
		}
		node := frontkit.Token{
			Line: tok.Line, Column: tok.Column + from,
			Offset: tok.Offset + from, EndOffset: tok.Offset + to,
		}
		start, end := node.Span()
		child := air.NewNodeAt(air.KindRsStringFrag, start, end, text[from:to])
		parts = append(parts, m.annotate(child, air.SemNote, m.note(noteName)))
	}
	part(contentStart, contentEnd, "rs-string-content")
	for i := contentStart; i+1 < contentEnd; i++ {
		if text[i] != '\\' {
			continue
		}
		stop := i + 2
		if text[i+1] == 'u' && i+1 < contentEnd && i+2 < contentEnd &&
			text[i+2] == '{' {
			if brace := strings.IndexByte(text[i+2:], '}'); brace >= 0 {
				stop = i + 2 + brace + 1
			}
		}
		if stop > contentEnd {
			stop = contentEnd
		}
		part(i, stop, "rs-escape")
		i = stop - 1
	}
	part(contentEnd, len(text), "rs-string-delimiter")
	return parts
}

// stringNode lowers a string token into the byte-exact StrLit leaf with the
// fragment children the parser recognized while reducing it.
func (m *rsMapper) stringNode(tok frontkit.Token) air.Node {
	return m.build(air.KindStrLit, tok, tok, tok.Text,
		m.stringFragments(tok)...)
}

// assign lowers an assignment to KindAssign. A compound assignment spelling
// is a distinct construct: the parser emits its dedicated node anchored on
// the operator token it just consumed.
func (m *rsMapper) assign(left air.Node, op frontkit.Token,
	right air.Node) air.Node {
	var children []air.Node
	if op.Text != "=" {
		children = append(children, m.dedicated(air.KindRsCompoundAssign,
			op, op.Text, "rs-compound-assign"))
	}
	children = append(children, left, right)
	return m.boundOperator(air.KindAssign, spanToken(left, op),
		spanToken(right, op), op, children...)
}

// unary lowers a prefix operator to the requested AIR kind.
func (m *rsMapper) unary(op frontkit.Token, operand air.Node,
	kind air.NodeKind, noteName string) air.Node {
	return m.unaryMarked(op, operand, kind, noteName)
}

// unaryMarked lowers a prefix operator with additional dedicated markers
// anchored at the operator token the parser just consumed.
func (m *rsMapper) unaryMarked(op frontkit.Token, operand air.Node,
	kind air.NodeKind, noteName string, marks ...air.Node) air.Node {
	end := op
	if operand != nil {
		end = spanToken(operand, op)
	}
	children := append([]air.Node{operand}, marks...)
	node := m.boundOperator(kind, op, end, op, children...)
	if noteName != "" {
		m.annotate(node, air.SemNote, m.note(noteName))
	}
	return node
}

// genericInstMarked lowers a generic instantiation with additional dedicated
// markers anchored at the angle brackets the parser just consumed.
func (m *rsMapper) genericInstMarked(base air.Node, end frontkit.Token,
	marks []air.Node, args ...air.Node) air.Node {
	children := append([]air.Node{base}, args...)
	children = append(children, marks...)
	node := m.build(air.KindGenericInst, spanToken(base, end), end, "", children...)
	return m.annotate(node, air.SemTypeParam, strconv.Itoa(len(args)))
}

// MacroCall lowers a macro invocation to KindCall with a macro note and the
// dedicated macro-invocation node anchored on the bang the parser consumed.
func (m *rsMapper) macroCall(left air.Node, bang, end frontkit.Token,
	args ...air.Node) air.Node {
	children := append([]air.Node{left,
		m.dedicated(air.KindRsMacroCall, bang, bang.Text, "rs-macro-call")},
		args...)
	node := m.build(air.KindCall, spanToken(left, bang), end, "", children...)
	m.annotate(node, air.SemNote, m.note("macro"))
	if rsContains(m.macroNames, air.AttributeOf(left)) {
		m.annotate(node, air.SemTag, m.note("known-macro"))
	}
	return node
}
