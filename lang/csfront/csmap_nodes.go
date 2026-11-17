/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// tupleNode 把逗号分隔的括号组归约为 KindTuple
// （tuple_expression → cs.tuple exact），跨度覆盖 [开括号, 闭括号)。
func (m *csMap) tupleNode(open, closing frontkit.Token,
	children []air.Node) air.Node {
	start, _ := open.Span()
	_, closeEnd := closing.Span()
	pos := air.NewPosition(start.Line, start.Column, start.Offset,
		closeEnd.Offset)
	if m.view != nil {
		children = m.view.OrderChildren(children)
	}
	return air.NewNodeAt(air.KindTuple, pos, closeEnd, "", children...)
}

// postfixNode 构造覆盖 [left.Start, 闭括号.End) 的调用/索引节点，实参按序
// 作子节点；空实参调用允许 args 为空，括号拼写由 BindOperator 写入标识。
func (m *csMap) postfixNode(kind air.NodeKind, left air.Node, open,
	closing frontkit.Token, args []air.Node) air.Node {
	lp := left.Pos()
	_, closeEnd := closing.Span()
	pos := air.NewPosition(lp.Line, lp.Column, lp.Offset, closeEnd.Offset)
	children := make([]air.Node, 0, len(args)+1)
	children = append(children, left)
	children = append(children, args...)
	if m.view != nil {
		children = m.view.OrderChildren(children)
	}
	return frontkit.BindOperator(kind, pos, closeEnd, open, children...)
}

// emptyGroupNode 构造覆盖 [开括号, 闭括号) 的诚实 partial 占位节点（如
// lambda 的空参数表 "()"），拼写由 BindTokenSpellings 兜底。
func (m *csMap) emptyGroupNode(open, closing frontkit.Token) air.Node {
	start, _ := open.Span()
	_, closeEnd := closing.Span()
	pos := air.NewPosition(start.Line, start.Column, start.Offset,
		closeEnd.Offset)
	node := air.NewNodeAt(air.KindPartialNode, pos, closeEnd, "")
	return m.annotate(node, air.SemPartial, "true")
}
