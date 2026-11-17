/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// node 以 [first,last] token 区间构造一个结构节点，children 顺序完全可控.
func (m *jsMapper) node(first, last int, kind air.NodeKind, text string,
	children ...air.Node) air.Node {
	if len(m.toks) == 0 {
		zero := air.NewPosition(1, 1, 0, 0)
		return air.NewNodeAt(kind, zero, zero, text, children...)
	}
	if first < 0 {
		first = 0
	}
	if last >= len(m.toks) {
		last = len(m.toks) - 1
	}
	if last < first {
		last = first
	}
	a, b := m.toks[first], m.toks[last]
	pos := m.view.Span(a, b)
	end := air.NewPosition(b.Line, b.Column, pos.EndOffset, pos.EndOffset)
	return air.NewNodeAt(kind, pos, end, text, jsNestChildren(children)...)
}

// cover 把区间内尚未投影的 lexical token 补成叶节点，保证 token 级保真.
func (m *jsMapper) cover(from, to int, dst *[]air.Node) {
	if dst == nil {
		return
	}
	if from < 0 {
		from = 0
	}
	if to > len(m.toks) {
		to = len(m.toks)
	}
	for i := from; i < to; i++ {
		if m.covered[i] {
			continue
		}
		tok := m.toks[i]
		if tok.Text == "" || !jsIsLexicalToken(tok.Kind) {
			continue
		}
		*dst = append(*dst, m.leaf(i, tok))
	}
}

// annotate 写入一个真实 SemKey，返回原节点便于链式归约.
func (*jsMapper) annotate(n air.Node, key air.SemKey, value string) air.Node {
	if n != nil && key.Valid() {
		_ = air.Annotate(n, key, value)
	}
	return n
}

// jsOperatorSpan converts two ordered child nodes into one wrapped interval.
func jsOperatorSpan(left, right air.Node) (pos, end air.Position) {
	lp, rp := left.Pos(), right.Pos()
	endOff := rp.EndOffset
	if endOff < lp.EndOffset {
		endOff = lp.EndOffset
	}
	pos = air.NewPosition(lp.Line, lp.Column, lp.Offset, endOff)
	end = air.NewPosition(rp.Line, rp.Column, endOff, endOff)
	return pos, end
}

// wrap 用左右子节点的区间构造二元结构节点.
func (m *jsMapper) wrap(kind air.NodeKind, left, right air.Node, text string,
	key air.SemKey, value string) air.Node {
	pos, end := jsOperatorSpan(left, right)
	n := air.NewNodeAt(kind, pos, end, text, jsOrderChildren([]air.Node{left, right})...)
	return m.annotate(n, key, value)
}

// boundWrap routes one binary-shaped operator through frontkit.BindOperator.
func (*jsMapper) boundWrap(kind air.NodeKind, left, right air.Node,
	tok frontkit.Token) air.Node {
	pos, end := jsOperatorSpan(left, right)
	return frontkit.BindOperator(kind, pos, end, tok,
		jsOrderChildren([]air.Node{left, right})...)
}

// wrapChildren 用首子节点起点与指定末 token 构造多子结构节点.
func (m *jsMapper) wrapChildren(kind air.NodeKind, first air.Node, last int,
	text string, children ...air.Node) air.Node {
	lp := first.Pos()
	endOff := lp.EndOffset
	if last >= 0 && last < len(m.toks) && m.toks[last].EndOffset > endOff {
		endOff = m.toks[last].EndOffset
	}
	line, column := lp.Line, lp.Column
	if last >= 0 && last < len(m.toks) {
		line, column = m.toks[last].Line, m.toks[last].Column
	}
	pos := air.NewPosition(lp.Line, lp.Column, lp.Offset, endOff)
	end := air.NewPosition(line, column, endOff, endOff)
	return air.NewNodeAt(kind, pos, end, text, jsOrderChildren(children)...)
}

// template 保留完整模板字符串文本，并把 ${} 插值降级为带 SemNote 的子节点.
func (m *jsMapper) template(_ air.Node, tok frontkit.Token) air.Node {
	parts := jsSplitTemplate(tok.Text)
	children := make([]air.Node, 0, len(parts))
	for _, part := range parts {
		if !part.expr || part.text == "" {
			continue
		}
		off := tok.Offset + part.offset
		width := len(part.text)
		pos := air.NewPosition(tok.Line, tok.Column, off, off+width)
		end := air.NewPosition(tok.Line, tok.Column, off+width, off+width)
		n := air.NewNodeAt(air.KindRegion, pos, end, "")
		_ = air.Annotate(n, air.SemNote, noteInterpolation)
		_ = air.Annotate(n, air.SemTag, part.text)
		children = append(children, n)
	}
	start, end := tok.Span()
	root := air.NewNodeAt(air.KindStrLit, start, end, tok.Text, children...)
	return m.annotate(root, air.SemNote, noteTemplate)
}

// unary 构造一元运算节点并把真实 token 原文写入 AIR-ID.
func (*jsMapper) unary(tok frontkit.Token, operand air.Node) air.Node {
	start, _ := tok.Span()
	op := operand.Pos()
	endOff := op.EndOffset
	if endOff < start.EndOffset {
		endOff = start.EndOffset
	}
	pos := air.NewPosition(start.Line, start.Column, start.Offset, endOff)
	end := air.NewPosition(op.Line, op.Column, endOff, endOff)
	return frontkit.BindOperator(air.KindUnary, pos, end, tok, operand)
}

// spread 构造展开运算符节点并标注真实降级注记.
func (m *jsMapper) spread(tok frontkit.Token, operand air.Node) air.Node {
	base := m.unary(tok, operand)
	n := frontkit.BindOperator(air.KindSpread, base.Pos(), base.End(), tok, operand)
	return m.annotate(n, air.SemNote, noteSpread)
}

// binary 按绑定表把真实 token 原文写入 AIR-ID，并保留既有语义算子事实.
func (m *jsMapper) binary(rule jsBindRule, left, right air.Node,
	tok frontkit.Token) air.Node {
	n := m.boundWrap(rule.kind, left, right, tok)
	n = m.annotate(n, air.SemOperator, rule.operator)
	if rule.operator == "nullish" {
		return m.annotate(n, air.SemNote, noteNullish)
	}
	return n
}

// chain 把成员访问或可选链归约为 KindField/KindSelector.
func (m *jsMapper) chain(left, right air.Node, kind air.NodeKind,
	operator string, tok frontkit.Token) air.Node {
	n := m.boundWrap(kind, left, right, tok)
	n = m.annotate(n, air.SemOperator, operator)
	if operator == "optional-chain" {
		return m.annotate(n, air.SemNote, noteOptionalChain)
	}
	return n
}

// arrow 把箭头函数归约为真实 KindArrow.
func (m *jsMapper) arrow(left, body air.Node, tok frontkit.Token) air.Node {
	n := m.boundWrap(air.KindArrow, left, body, tok)
	return m.annotate(n, air.SemNote, noteArrow)
}

// typeAssert 把 TS as/satisfies 归约为真实 KindTypeAssert.
func (m *jsMapper) typeAssert(left, right air.Node, tok frontkit.Token) air.Node {
	n := m.wrap(air.KindTypeAssert, left, right, tok.Text, 0, "")
	return m.annotate(n, air.SemNote, "type-assert-"+tok.Text)
}

// decorator 把 @decorator 归约为真实 KindDecorator.
//
// 节点区间从 @ 起、到被装饰表达式止，否则装饰目标会逃出装饰器节点区间.
func (m *jsMapper) decorator(tok frontkit.Token, target air.Node) air.Node {
	start, _ := tok.Span()
	endOff, line, column := start.EndOffset, start.Line, start.Column
	if target != nil {
		tp := target.Pos()
		if tp.EndOffset > endOff {
			endOff = tp.EndOffset
		}
		line, column = tp.Line, tp.Column
	}
	pos := air.NewPosition(start.Line, start.Column, start.Offset, endOff)
	end := air.NewPosition(line, column, endOff, endOff)
	n := air.NewNodeAt(air.KindDecorator, pos, end, "@", target)
	return m.annotate(n, air.SemNote, noteDecorator)
}
