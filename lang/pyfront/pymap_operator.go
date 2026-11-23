/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the operator-anchored and expression-form lowering of
// pyMap: the shared operator span helpers, infix binary lowering (with the
// dedicated boolean/comparison construct nodes), assignment forms and every
// postfix form the Pratt loop or the suffix parser produces.

// pyOperatorSpan converts one node interval into explicit AIR start/end positions.
func pyOperatorSpan(start, end frontkit.Token) (pos, endPos air.Position) {
	finish := end
	if finish.EndOffset < start.Offset {
		finish = start
	}
	pos = air.NewPosition(start.Line, start.Column, start.Offset, finish.EndOffset)
	endPos = air.NewPosition(start.Line, start.Column, finish.EndOffset, finish.EndOffset)
	return pos, endPos
}

// boundOperator routes Python operator nodes through frontkit.BindOperator.
func (m *pyMap) boundOperator(kind air.NodeKind, start, end frontkit.Token,
	tok frontkit.Token, children ...air.Node) air.Node {
	pos, endPos := pyOperatorSpan(start, end)
	return frontkit.BindOperator(kind, pos, endPos, tok,
		m.view.OrderChildren(children)...)
}

// binary lowers an infix operator expression to KindBinary. When the Pratt
// loop has just consumed a boolean or comparison operator, the reduction also
// emits that operator's dedicated construct node.
func (m *pyMap) binary(left air.Node, op frontkit.Token, right air.Node) air.Node {
	start, end := spanToken(left, op), spanToken(right, op)
	if op.Text == ":=" {
		return m.assign(start, end, op, left, right)
	}
	if op.Text == "|" && (left.Kind() == air.KindNilLit || right.Kind() == air.KindNilLit) {
		node := m.boundOperator(air.KindOptional, start, end, op, left, right)
		return m.annotate(node, air.SemNote, m.note("optional"))
	}
	children := []air.Node{left, right}
	if kind := pyOperatorKind(op.Text); kind != air.KindInvalid {
		// Anchor on the operator token itself: that span sits between the two
		// operands and survives the non-overlap filter in OrderChildren.
		children = append(children, m.dedicated(
			kind, op, op, op.Text, "py-operator"))
	}
	return m.boundOperator(air.KindBinary, start, end, op, children...)
}

// assign lowers an assignment or walrus expression to KindAssign.
func (m *pyMap) assign(start, end frontkit.Token, op frontkit.Token,
	left, right air.Node) air.Node {
	children := []air.Node{left, right}
	if op.Text == ":=" {
		// A walrus expression is a distinct construct; anchor its dedicated
		// node on the := token so OrderChildren keeps it.
		children = append(children, m.dedicated(
			air.KindPyNamedExpr, op, op, op.Text, "py-named-expr"))
	}
	node := m.boundOperator(air.KindAssign, start, end, op, children...)
	if op.Text == ":=" {
		m.annotate(node, air.SemNote, m.note("walrus"))
	}
	return node
}

// ternary lowers a conditional expression to KindTernary.
func (m *pyMap) ternary(thenNode, condition, alternative air.Node,
	elseTok frontkit.Token) air.Node {
	return m.build(air.KindTernary, spanToken(thenNode, elseTok),
		spanToken(alternative, elseTok), "", thenNode, condition, alternative)
}

// unary lowers a prefix operator to the requested AIR kind.
func (m *pyMap) unary(op frontkit.Token, operand air.Node,
	kind air.NodeKind, note string) air.Node {
	node := m.boundOperator(kind, op, spanToken(operand, op), op, operand)
	if note != "" {
		m.annotate(node, air.SemNote, m.note(note))
	}
	return node
}

// awaitNode lowers an await expression to KindAwait with SemAsync.
func (m *pyMap) awaitNode(op frontkit.Token, operand air.Node) air.Node {
	node := m.build(air.KindAwait, op, spanToken(operand, op), op.Text, operand)
	return m.annotate(node, air.SemAsync, "true")
}

// lambdaNode lowers a lambda expression to KindLambda.
func (m *pyMap) lambdaNode(start frontkit.Token, params []air.Node,
	body air.Node) air.Node {
	children := append(append([]air.Node{}, params...), body)
	return m.build(air.KindLambda, start, spanToken(body, start), "", children...)
}

// call lowers a call expression to KindCall.
func (m *pyMap) call(callee air.Node, open, endTok frontkit.Token,
	args ...air.Node) air.Node {
	children := append([]air.Node{callee}, args...)
	return m.build(air.KindCall, spanToken(callee, open), endTok, "", children...)
}

// field lowers an attribute access to KindField.
func (m *pyMap) field(base air.Node, name frontkit.Token) air.Node {
	return m.build(air.KindField, spanToken(base, name), name, "",
		base, m.lexical(name))
}

// index lowers a subscript to KindIndex.
func (m *pyMap) index(base air.Node, open, endTok frontkit.Token,
	parts ...air.Node) air.Node {
	children := append([]air.Node{base}, parts...)
	return m.build(air.KindIndex, spanToken(base, open), endTok, "", children...)
}

// slice lowers a slice subscript to KindSlice.
func (m *pyMap) slice(base air.Node, open, endTok frontkit.Token,
	parts ...air.Node) air.Node {
	children := append([]air.Node{base}, parts...)
	node := m.build(air.KindSlice, spanToken(base, open), endTok, "", children...)
	return m.annotate(node, air.SemNote, m.note("slice"))
}

// tupleNode lowers a parenthesized tuple to KindTuple.
func (m *pyMap) tupleNode(open, endTok frontkit.Token, items ...air.Node) air.Node {
	return m.build(air.KindTuple, open, endTok, "", items...)
}

// exprListNode lowers a comma-separated expression list to KindPyExprList.
func (m *pyMap) exprListNode(open, endTok frontkit.Token, items ...air.Node) air.Node {
	return m.build(air.KindPyExprList, open, endTok, "", items...)
}

// spread lowers *args and **kwargs to KindSpread or KindPyStarList.
func (m *pyMap) spread(op frontkit.Token, operand air.Node) air.Node {
	kind := air.KindPyStarList
	note := "star-list"
	if op.Text == "**" {
		kind = air.KindSpread
		note = "kwargs"
	}
	node := m.boundOperator(kind, op, spanToken(operand, op), op, operand)
	return m.annotate(node, air.SemNote, m.note(note))
}
