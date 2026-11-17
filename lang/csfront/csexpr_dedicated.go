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

// This file holds the Pratt rule builders of csExpr that emit the dedicated
// C# construct nodes while the parser consumes a prefix keyword or an infix
// operator (object-creation, typeof, sizeof, conditional-access, pattern,
// alias-qualified).

// keywordPrefix 把 await/new/typeof/sizeof/nameof 归约为专有 AIR 种类，并在
// Pratt 刚消费前缀关键字的时刻补发该构造的 dedicated 节点。
func (e *csExpr) keywordPrefix() func(*frontkit.PrattParser,
	frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		operand, err := parser.ParsePrefixOperand(tok)
		if err != nil || operand == nil {
			return e.mapper.literal(tok), nil
		}
		kind := air.KindUnary
		switch tok.Text {
		case "await":
			kind = air.KindAwait
		case "typeof":
			kind = air.KindTypeAssert
		case "nameof":
			kind = air.KindCall
		}
		// 前缀构造的 dedicated 节点锚定在刚消费的前缀关键字位置。
		var marks []air.Node
		switch tok.Text {
		case "new":
			marks = append(marks, e.mapper.dedicated(air.KindCsObjectCreation,
				tok, tok.Text, "cs-object-creation"))
		case "typeof":
			marks = append(marks, e.mapper.dedicated(air.KindCsTypeof,
				tok, tok.Text, "cs-typeof"))
		case "sizeof":
			marks = append(marks, e.mapper.dedicated(air.KindCsSizeof,
				tok, tok.Text, "cs-sizeof"))
		}
		var node air.Node
		if len(marks) > 0 {
			node = e.mapper.unaryNodeMarked(kind, tok, operand, marks...)
		} else {
			node = e.mapper.unaryNode(kind, tok, operand)
		}
		if tok.Text == "await" {
			e.mapper.annotate(node, air.SemAsync, "true")
		}
		if tok.Text == "typeof" || tok.Text == "nameof" {
			e.mapper.annotate(node, air.SemNote, e.mapper.note(tok.Text))
		}
		return node, nil
	}
}

// infixRule 生成二元/赋值/成员/调用/可空/三元/lambda 节点并写入语义标注；
// 空条件、模式判定与别名限定在 Pratt 刚消费运算符时补发 dedicated 节点。
func (e *csExpr) infixRule() func(*frontkit.PrattParser, air.Node,
	frontkit.Token, air.Node) (air.Node, error) {
	return func(_ *frontkit.PrattParser, left air.Node, tok frontkit.Token,
		right air.Node) (air.Node, error) {
		kind := air.KindBinary
		switch tok.Text {
		case "=", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=",
			"<<=", ">>=", "??=":
			kind = air.KindAssign
		case ".", "?.":
			kind = air.KindField
		case "(":
			kind = air.KindCall
		case "[":
			kind = air.KindIndex
		case "=>":
			kind = air.KindLambda
		case "?":
			kind = air.KindOptional
		case ":":
			kind = air.KindTernary
		case "is", "as":
			kind = air.KindTypeAssert
		}
		var marks []air.Node
		switch tok.Text {
		case "?.":
			marks = append(marks, e.mapper.dedicated(
				air.KindCsConditionalAccess, tok, tok.Text,
				"cs-conditional-access"))
		case "is":
			marks = append(marks, e.mapper.dedicated(
				air.KindCsPattern, tok, tok.Text, "cs-pattern"))
		case "::":
			marks = append(marks, e.mapper.dedicated(
				air.KindCsAliasName, tok, tok.Text, "cs-alias-name"))
		}
		var node air.Node
		if len(marks) > 0 {
			node = e.mapper.infixNodeMarked(kind, left, tok, right, marks...)
		} else {
			node = e.mapper.infixNode(kind, left, tok, right)
		}
		switch tok.Text {
		case "?.":
			e.mapper.annotate(node, air.SemNote, e.mapper.note("null-conditional"))
		case "??":
			e.mapper.annotate(node, air.SemNote, e.mapper.note("null-coalescing"))
		case "??=":
			e.mapper.annotate(node, air.SemNote,
				e.mapper.note("null-coalescing-assign"))
		case "=>":
			e.mapper.annotate(node, air.SemNote, e.mapper.note("lambda"))
		case "?":
			e.mapper.annotate(node, air.SemNote, e.mapper.note("nullable"))
		}
		return node, nil
	}
}
