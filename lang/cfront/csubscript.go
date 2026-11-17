/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"fmt"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the C/C++ expression forms that are not plain binary
// operators: the subscript postfix and the conditional expression.

// postfix 把 a[i] 归约为真实 KindIndex 节点.
func (*cExprReducer) postfix() func(*frontkit.PrattParser, air.Node,
	frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, left air.Node,
		tok frontkit.Token) (air.Node, error) {
		index, err := parser.ParseExpression(0)
		if err != nil || index == nil {
			return nil, err
		}
		closeTok, ok := parser.Stream().Advance()
		if !ok || closeTok.Text != "]" {
			return nil, fmt.Errorf("%w: subscript is not closed",
				frontkit.ErrUnknownToken)
		}
		start := left.Pos()
		_, closeEnd := closeTok.Span()
		pos := air.NewPosition(start.Line, start.Column, start.Offset,
			closeEnd.EndOffset)
		node := air.NewNodeAt(air.KindIndex, pos, closeEnd, "", left, index)
		_ = air.Annotate(node, air.SemOperator, tok.Text)
		return node, nil
	}
}

// The reduceTernary reducer lowers a ternary expression to a KindTernary node.
//
// Note: a ternary has no binary infix form, so the Pratt parser cannot handle
// it directly. This reducer splits the condition, true branch and false branch
// by bracket depth, then assembles them. When no top-level question mark exists
// in the range, it returns nil.
func (r *cExprReducer) reduceTernary(tokens []frontkit.Token, lo, hi int) []air.Node {
	question := cTopLevelToken(tokens, lo, hi, "?")
	if question < 0 {
		return nil
	}
	colon := cTopLevelToken(tokens, question+1, hi, ":")
	if colon < 0 {
		return nil
	}
	condition := r.reduceSegment(tokens, lo, question)
	consequent := r.reduceSegment(tokens, question+1, colon)
	alternative := r.reduceSegment(tokens, colon+1, hi)
	if len(condition) == 0 || len(consequent) == 0 || len(alternative) == 0 {
		return nil
	}
	children := make([]air.Node, 0, len(condition)+len(consequent)+len(alternative))
	children = append(children, condition...)
	children = append(children, consequent...)
	children = append(children, alternative...)
	start, end := cSpan(tokens, lo, hi)
	return []air.Node{air.NewNodeAt(air.KindTernary, start, end, "?", children...)}
}

// cTopLevelToken 返回区间内第一个括号深度为零的指定拼写 token 下标；没有则
// 返回 -1。嵌套在括号内的同名 token 不算顶层。
func cTopLevelToken(tokens []frontkit.Token, lo, hi int, text string) int {
	depth := 0
	for j := lo; j < hi && j < len(tokens); j++ {
		switch tokens[j].Text {
		case "(", "[", "{":
			depth++
			continue
		case ")", "]", "}":
			if depth > 0 {
				depth--
			}
			continue
		}
		if depth == 0 && tokens[j].Text == text {
			return j
		}
	}
	return -1
}
