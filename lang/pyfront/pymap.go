/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"strconv"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the pyMap core: node construction, lexical projection,
// trivia and the construct-level helpers for containers, types and literals.
// The operator-anchored and expression-form lowering lives in pymap_operator.go.

// pyMap builds AIR nodes and applies the Python construct lowering rules.
type pyMap struct {
	view  *frontkit.View
	notes []pyNote
}

// newPyMap binds one node-building view and the default semantic note table.
func newPyMap(view *frontkit.View) *pyMap {
	return &pyMap{view: view, notes: pyNotes}
}

// note resolves one named semantic note from the injectable table.
func (m *pyMap) note(name string) string { return pyNoteFor(m.notes, name) }

// annotate attaches a semantic value to a node and returns the same node.
func (*pyMap) annotate(node air.Node, key air.SemKey, value string) air.Node {
	_ = air.Annotate(node, key, value)
	return node
}

// partial marks a node as an incomplete, recoverable reduction.
func (m *pyMap) partial(node air.Node) air.Node {
	return m.annotate(node, air.SemPartial, "true")
}

// build creates one node over [start.Offset, end.EndOffset) with ordered,
// non-overlapping children.
func (m *pyMap) build(kind air.NodeKind, start, end frontkit.Token,
	text string, children ...air.Node) air.Node {
	finish := end
	if finish.EndOffset < start.Offset {
		finish = start
	}
	pos := air.NewPosition(start.Line, start.Column, start.Offset, finish.EndOffset)
	return air.NewNode(kind, pos, text, m.view.OrderChildren(children)...)
}

// leaf projects one lexical token onto a byte-exact AIR leaf.
func (m *pyMap) leaf(tok frontkit.Token, kind air.NodeKind) air.Node {
	return m.build(kind, tok, tok, tok.Text)
}

// lexical projects a token using the frozen token-to-AIR kind table.
func (m *pyMap) lexical(tok frontkit.Token) air.Node {
	return m.leaf(tok, pyTokenNodeKind(tok.Kind))
}

// partialLeaf projects an unexpected token at a recovery position. Lexical
// tokens keep their frozen byte-exact projection kind; every keyword, operator
// or punctuation token that reaches a recovery position is an honest partial
// node — never a fabricated structural kind guessed from the token text.
func (m *pyMap) partialLeaf(tok frontkit.Token) air.Node {
	kind := air.KindPartialNode
	if pyIsLexicalKind(tok.Kind) {
		kind = pyTokenNodeKind(tok.Kind)
	}
	return m.partial(m.build(kind, tok, tok, ""))
}

// spanToken projects an AIR node interval back onto a token-shaped span.
func spanToken(node air.Node, fallback frontkit.Token) frontkit.Token {
	if node == nil {
		return fallback
	}
	pos := node.Pos()
	return frontkit.Token{
		Line: pos.Line, Column: pos.Column,
		Offset: pos.Offset, EndOffset: pos.EndOffset,
	}
}

// comment lowers a comment, a directive comment, or a docstring.
func (m *pyMap) comment(tok frontkit.Token) air.Node {
	kind := air.KindComment
	if pyHasPrefix(tok.Text, pyDirectivePrefixes) {
		kind = air.KindDirective
	}
	if strings.HasPrefix(tok.Text, `"""`) || strings.HasPrefix(tok.Text, `'''`) {
		kind = air.KindDoc
	}
	node := m.leaf(tok, kind)
	if kind == air.KindDirective {
		m.annotate(node, air.SemDirective, strings.TrimSpace(tok.Text))
	}
	return node
}

// fstring lowers an f-prefixed literal with interpolation children.
func (m *pyMap) fstring(tok frontkit.Token, parts ...air.Node) air.Node {
	node := m.build(air.KindStrLit, tok, tok, tok.Text, parts...)
	return m.annotate(node, air.SemNote, m.note("fstring"))
}

// comprehension lowers a comprehension expression to KindCompositeLit.
func (m *pyMap) comprehension(start, end frontkit.Token, note string,
	children ...air.Node) air.Node {
	node := m.build(air.KindCompositeLit, start, end, "", children...)
	return m.annotate(node, air.SemNote, m.note(note))
}

// collection lowers a literal container to KindCompositeLit.
func (m *pyMap) collection(open, endTok frontkit.Token, note string,
	items ...air.Node) air.Node {
	node := m.build(air.KindCompositeLit, open, endTok, "", items...)
	return m.annotate(node, air.SemNote, m.note(note))
}

// pair lowers one dictionary key:value entry.
func (m *pyMap) pair(key, value air.Node, colon frontkit.Token) air.Node {
	start, end := spanToken(key, colon), spanToken(value, colon)
	node := m.build(air.KindAssign, start, end, ":", key, value,
		m.dedicated(air.KindPyDictPair, colon, colon, ":", "py-dict-pair"))
	return m.annotate(node, air.SemNote, m.note("dict-pair"))
}

// indentBlock lowers an indentation-driven suite to KindBlock.
func (m *pyMap) indentBlock(start, end frontkit.Token,
	children ...air.Node) air.Node {
	node := m.build(air.KindBlock, start, end, "", children...)
	return m.annotate(node, air.SemNote, m.note("indent"))
}

// decorator lowers a decorator line to KindDecorator.
func (m *pyMap) decorator(start, end frontkit.Token,
	children ...air.Node) air.Node {
	node := m.build(air.KindDecorator, start, end, "", children...)
	return m.annotate(node, air.SemNote, m.note("decorator"))
}

// genericInst lowers a subscripted generic type to KindGenericInst.
func (m *pyMap) genericInst(base air.Node, endTok frontkit.Token,
	args ...air.Node) air.Node {
	children := append([]air.Node{base}, args...)
	node := m.build(air.KindGenericInst, spanToken(base, endTok), endTok, "", children...)
	return m.annotate(node, air.SemTypeParam, strconv.Itoa(len(args)))
}

// optional lowers Optional[T] and T | None to KindOptional.
func (m *pyMap) optional(start, end frontkit.Token,
	children ...air.Node) air.Node {
	node := m.build(air.KindOptional, start, end, "", children...)
	return m.annotate(node, air.SemNote, m.note("optional"))
}

// typeNode lowers a type annotation expression to KindNamed.
func (m *pyMap) typeNode(expr air.Node) air.Node {
	if qualified := m.qualifiedNameAs(expr, air.AttributeOf(expr)); qualified != nil {
		// A dotted type is a qualified name, not a plain named type.
		return qualified
	}
	start := spanToken(expr, frontkit.Token{})
	return m.build(air.KindNamed, start, start, air.AttributeOf(expr), expr)
}

// matchExpr wraps a match subject in KindMatchExpr.
func (m *pyMap) matchExpr(expr air.Node) air.Node {
	start := spanToken(expr, frontkit.Token{})
	return m.build(air.KindMatchExpr, start, start, "", expr)
}
