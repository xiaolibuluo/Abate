/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"fmt"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// csPrefixOp 是一条可作前缀使用的 C# 运算符拼写。
type csPrefixOp struct {
	text string
}

// csPrefixOps 列出 C# 的前缀运算符（unary_expression 的运算符集合）。
var csPrefixOps = []csPrefixOp{
	{"-"}, {"+"}, {"!"}, {"~"}, {"++"}, {"--"},
	{".."},
}

// opPrefix 把刚消费的前缀运算符连同其操作数归约为 KindUnary 节点。
func (e *csExpr) opPrefix() func(*frontkit.PrattParser,
	frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		operand, err := parser.ParsePrefixOperand(tok)
		if err != nil {
			return nil, err
		}
		if operand == nil {
			return nil, fmt.Errorf("%w: missing operand after %q",
				frontkit.ErrUnknownToken, tok.Text)
		}
		if tok.Text == ".." {
			mark := e.mapper.dedicated(air.KindSpread, tok, tok.Text, "cs-spread")
			return e.mapper.unaryNodeMarked(air.KindUnary, tok, operand, mark), nil
		}
		return e.mapper.unaryNode(air.KindUnary, tok, operand), nil
	}
}

// parenPrefix 归约到达前缀位置的括号组：消费到配对闭括号为止，对内部区间
// 做真实子归约——单表达式直接返回其节点（括号是分组语法糖，拼写由
// BindTokenSpellings 兜底），逗号分隔的多元素归约为 KindTuple
// （tuple_expression → cs.tuple exact）。
func (e *csExpr) parenPrefix() func(*frontkit.PrattParser,
	frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		stream := parser.Stream()
		depth := 0
		inner := make([]frontkit.Token, 0, 8)
		var closeTok frontkit.Token
		for {
			next, ok := stream.Peek()
			if !ok || next.Kind == frontkit.TokenEOF {
				return nil, fmt.Errorf("%w: unterminated %q group",
					frontkit.ErrUnknownToken, tok.Text)
			}
			stream.Advance()
			if next.Text == tok.Text {
				depth++
			} else if next.Text == csCloseOf(tok.Text) {
				if depth == 0 {
					closeTok = next
					break
				}
				depth--
			}
			inner = append(inner, next)
		}
		nodes := e.reduceSegment(inner, 0, len(inner))
		if len(nodes) == 0 {
			// 空组（如 lambda 的空参数表 "()"）：诚实标注 partial 的占位
			// 节点，跨度覆盖整组，拼写由 BindTokenSpellings 兜底。
			return e.mapper.emptyGroupNode(tok, closeTok), nil
		}
		if len(nodes) == 1 {
			// 转型检测：若括号内容为单个类型节点（Named），且后续
			// token 可作为表达式操作数（标识符/字面量/关键字/括号/前缀
			// 运算符），则按 (Type)expr 归约为 KindCast。
			k := nodes[0].Kind()
			if k == air.KindNamed {
				next, ok := stream.Peek()
				if ok && isCastOperandStart(next) {
					operand, err := parser.ParsePrefixOperand(next)
					if err == nil && operand != nil {
						cast := e.mapper.infixNode(
							air.KindCast,
							nodes[0], closeTok, operand)
						e.mapper.annotate(cast, air.SemNote,
							e.mapper.note("cast"))
						return cast, nil
					}
				}
			}
			return nodes[0], nil
		}
		// 多元素：真实 KindTuple 归约（cs.tuple → tuple_expression exact）。
		return e.mapper.tupleNode(tok, closeTok, nodes), nil
	}
}

// csCloseOf 返回一个开括号的配对闭括号拼写。
func csCloseOf(open string) string {
	if open == "(" {
		return ")"
	}
	return "]"
}

// isCastOperandStart 报告一个 token 是否可作为转型表达式右操作数的起点。
func isCastOperandStart(tok frontkit.Token) bool {
	if tok.Kind == frontkit.TokenEOF ||
		tok.Kind == frontkit.TokenNewline ||
		tok.Kind == frontkit.TokenComment ||
		tok.Text == "" {
		return false
	}
	// 排除语句定界符和后缀运算符
	switch tok.Text {
	case ")", "]", "}", ";", ",", ".", ":", "=>":
		return false
	}
	// 排除中缀/赋值运算符
	switch tok.Text {
	case "=", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=",
		"<<=", ">>=", "??=", "??", "||", "&&", "==", "!=",
		"<", ">", "<=", ">=", "|", "^", "&", "<<", ">>",
		"+", "-", "*", "/", "%", "is", "as", "in":
		return false
	}
	return true
}

// bracketPostfix 把跟在表达式之后的括号对归约为调用/索引节点：紧跟闭括号
// 时按空实参调用处理（Pratt 通用中缀路径会在闭括号处解析失败，因此调用与
// 索引必须走 Postfix 自行消费实参区间），否则对实参区间做一次真实子归约。
func (e *csExpr) bracketPostfix(kind air.NodeKind) func(*frontkit.PrattParser,
	air.Node, frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, left air.Node,
		tok frontkit.Token) (air.Node, error) {
		if left == nil {
			return nil, fmt.Errorf("%w: postfix %q without an operand",
				frontkit.ErrUnknownToken, tok.Text)
		}
		stream := parser.Stream()
		closer := csCloseOf(tok.Text)
		depth := 0
		inner := make([]frontkit.Token, 0, 8)
		var closeTok frontkit.Token
		for {
			next, ok := stream.Peek()
			if !ok || next.Kind == frontkit.TokenEOF {
				return nil, fmt.Errorf("%w: unterminated %q group",
					frontkit.ErrUnknownToken, tok.Text)
			}
			stream.Advance()
			switch next.Text {
			case tok.Text:
				depth++
			case closer:
				if depth == 0 {
					closeTok = next
					return e.mapper.postfixNode(kind, left, tok, closeTok,
						e.reduceSegment(inner, 0, len(inner))), nil
				}
				depth--
			}
			inner = append(inner, next)
		}
	}
}
