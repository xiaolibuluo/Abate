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

// ktExpr 用 frontkit.PrattParser 归约 Kotlin 表达式区间。
//
// 前缀规则覆盖 when 表达式与字符串模板入口；中缀规则覆盖空安全 ?. 、
// Elvis ?: 、lambda -> 、成员调用与区间 .. 等，绑定力全部是有序整数表。
// 后缀 !! 与可空类型 ? 由词法叶子的可注入空安全映射表投影。任何片段归约
// 失败都降级为词法叶子序列，绝不向上抛出错误。
type ktExpr struct {
	bindings []ktBinding
	mapper   *ktMap
}

// newKtExpr 绑定可注入的绑定表与节点映射器。
func newKtExpr(bindings []ktBinding, mapper *ktMap) *ktExpr {
	if len(bindings) == 0 {
		bindings = ktInfixBindings
	}
	return &ktExpr{bindings: bindings, mapper: mapper}
}

// reduce 把 [lo,hi) 切成互不重叠的片段并逐段归约，保证每个词法 token 都
// 落在某个 AIR 叶子中，从而满足 token 级保真与位置区间覆盖。
func (e *ktExpr) reduce(tokens []frontkit.Token, lo, hi int) []air.Node {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	out := make([]air.Node, 0, 8)
	start := lo
	for i := lo; i < hi; i++ {
		tok := tokens[i]
		switch tok.Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF,
			frontkit.TokenIndent, frontkit.TokenDedent:
			if start < i {
				out = append(out, e.reduceSegment(tokens, start, i)...)
			}
			start = i + 1
		case frontkit.TokenComment:
			if start < i {
				out = append(out, e.reduceSegment(tokens, start, i)...)
			}
			out = append(out, e.mapper.comment(tok))
			start = i + 1
		default:
			if tok.Text == ";" || tok.Text == "{" || tok.Text == "}" {
				if start < i {
					out = append(out, e.reduceSegment(tokens, start, i)...)
				}
				out = append(out, e.mapper.literal(tok))
				start = i + 1
			}
		}
	}
	if start < hi {
		out = append(out, e.reduceSegment(tokens, start, hi)...)
	}
	return out
}

// reduceSegment 用 Pratt 解析器归约一段连续 token；失败时降级为词法叶子。
func (e *ktExpr) reduceSegment(tokens []frontkit.Token, lo, hi int) []air.Node {
	sub := make([]frontkit.Token, 0, hi-lo+1)
	for j := lo; j < hi; j++ {
		switch tokens[j].Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF,
			frontkit.TokenIndent, frontkit.TokenDedent, frontkit.TokenComment:
			continue
		default:
			sub = append(sub, tokens[j])
		}
	}
	if len(sub) == 0 {
		return nil
	}
	last := sub[len(sub)-1]
	eof := frontkit.Token{
		Kind: frontkit.TokenEOF, Offset: last.EndOffset, EndOffset: last.EndOffset,
		Line: last.Line, Column: last.Column,
	}
	sub = append(sub, eof)
	stream, err := frontkit.NewTokenStream(sub)
	if err != nil {
		return e.leafNodes(tokens, lo, hi)
	}
	parser, err := frontkit.NewPrattParser(stream, frontkit.NewView(nil, ""))
	if err != nil {
		return e.leafNodes(tokens, lo, hi)
	}
	e.register(parser)
	out := make([]air.Node, 0, 4)
	for {
		peek, ok := stream.Peek()
		if !ok || peek.Kind == frontkit.TokenEOF {
			break
		}
		before := stream.Remaining()
		node, parseErr := parser.ParseExpression(0)
		if parseErr != nil || node == nil {
			tok, advanced := stream.Advance()
			if !advanced {
				break
			}
			out = append(out, e.mapper.literal(tok))
			continue
		}
		out = append(out, node)
		if stream.Remaining() >= before {
			if _, advanced := stream.Advance(); !advanced {
				break
			}
		}
	}
	out = e.coverLeaves(sub, out)
	if e.mapper.view != nil {
		return e.mapper.view.OrderChildren(out)
	}
	return out
}

// coverLeaves 为每个起始偏移尚未落在任何已产出节点区间内的 token 补一个
// 词法叶子，保证以操作符结尾的片段仍满足 token 级保真且不重叠。
func (e *ktExpr) coverLeaves(sub []frontkit.Token, out []air.Node) []air.Node {
	limit := len(sub) - 1
	for k := 0; k < limit; k++ {
		covered := false
		for i := range out {
			start, end := air.Span(out[i])
			if sub[k].Offset >= start && sub[k].Offset < end {
				covered = true
				break
			}
		}
		if !covered {
			out = append(out, e.mapper.literal(sub[k]))
		}
	}
	return out
}

// register 绑定前缀/中缀规则；注册顺序确定，不使用 map 迭代。
func (e *ktExpr) register(parser *frontkit.PrattParser) {
	literalPrefix := func(_ *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		return e.mapper.literal(tok), nil
	}
	for _, kind := range []frontkit.TokenKind{
		frontkit.TokenIdent, frontkit.TokenKeyword, frontkit.TokenNum,
		frontkit.TokenStr, frontkit.TokenChar, frontkit.TokenComment,
		frontkit.TokenOp, frontkit.TokenPunct, frontkit.TokenIllegal,
	} {
		_ = parser.Register(kind, frontkit.ParseRule{Prefix: literalPrefix})
	}
	infix := e.infixRule()
	for _, binding := range e.bindings {
		// 每条中缀规则同时保留前缀投影，使被注册为操作符文本的定界符
		// 作为片段首 token 时仍能归约成叶子而不报错。
		rule := frontkit.ParseRule{
			Prefix: literalPrefix, InfixPower: binding.power,
			RightAssoc: binding.right, Infix: infix,
		}
		for _, kind := range []frontkit.TokenKind{
			frontkit.TokenOp, frontkit.TokenKeyword,
			frontkit.TokenIdent, frontkit.TokenPunct,
		} {
			_ = parser.RegisterText(kind, binding.text, rule)
		}
	}
	whenPrefix := e.whenPrefixRule()
	for _, kind := range []frontkit.TokenKind{
		frontkit.TokenKeyword, frontkit.TokenIdent,
	} {
		_ = parser.RegisterText(kind, "when",
			frontkit.ParseRule{PrefixPower: ktWhenPrefixPower, Prefix: whenPrefix})
	}
	// this/super 表达式由 Pratt 前缀归约产出专有 AIR 节点。
	for _, special := range []struct {
		text string
		kind air.NodeKind
		note string
	}{
		{"this", air.KindKtThis, "kt-this"},
		{"super", air.KindKtSuper, "kt-super"},
	} {
		text, kind, note := special.text, special.kind, special.note
		_ = parser.RegisterText(frontkit.TokenKeyword, text,
			frontkit.ParseRule{Prefix: func(_ *frontkit.PrattParser,
				tok frontkit.Token) (air.Node, error) {
				node := e.mapper.leafKind(tok, kind)
				return e.mapper.annotate(node, air.SemNote,
					e.mapper.note(note)), nil
			}})
	}
}

// whenPrefixRule 把 when 表达式归约为 KindMatchExpr 带 when-expr 标注。
func (e *ktExpr) whenPrefixRule() func(*frontkit.PrattParser,
	frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		operand, err := parser.ParsePrefixOperand(tok)
		if err != nil || operand == nil {
			node := e.mapper.leafKind(tok, air.KindMatchExpr)
			return e.mapper.annotate(node, air.SemNote,
				e.mapper.note("when-expr")), nil
		}
		node := e.mapper.unaryNode(air.KindMatchExpr, tok, operand)
		return e.mapper.annotate(node, air.SemNote,
			e.mapper.note("when-expr")), nil
	}
}

// infixRule 生成二元/赋值/成员/调用/lambda/空安全节点并写入语义标注；
// 空安全相关标注来自可注入映射表，删去表项后行为随之改变。
func (e *ktExpr) infixRule() func(*frontkit.PrattParser, air.Node,
	frontkit.Token, air.Node) (air.Node, error) {
	return func(_ *frontkit.PrattParser, left air.Node, tok frontkit.Token,
		right air.Node) (air.Node, error) {
		kind := air.KindBinary
		switch tok.Text {
		case "=", "+=", "-=", "*=", "/=", "%=":
			kind = air.KindAssign
		case ".", "?.":
			kind = air.KindField
		case "(":
			kind = air.KindCall
		case "[":
			kind = air.KindIndex
		case "->":
			kind = air.KindLambda
		case "::":
			kind = air.KindSelector
		case "is", "as":
			kind = air.KindTypeAssert
		}
		var node air.Node
		if tok.Text == "::" && left != nil && right != nil {
			// Pratt 刚消费 "::"：可调用引用构造的 dedicated 节点锚定在其记号上。
			node = e.mapper.infixNodeMarked(kind, left, tok, right,
				e.mapper.dedicated(air.KindKtCallableRef, tok, tok.Text,
					"kt-callable-ref"))
		} else {
			node = e.mapper.infixNode(kind, left, tok, right)
		}
		if rule, ok := ktNullSafeFor(e.mapper.nullSafe, tok.Text); ok {
			e.mapper.annotate(node, air.SemNote, e.mapper.note(rule.note))
		}
		if tok.Text == "->" {
			e.mapper.annotate(node, air.SemNote, e.mapper.note("lambda"))
		}
		return node, nil
	}
}

// leafNodes 把区间内每个非 trivia token 映射为词法叶子。
func (e *ktExpr) leafNodes(tokens []frontkit.Token, lo, hi int) []air.Node {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	if hi <= lo {
		return nil
	}
	out := make([]air.Node, 0, hi-lo)
	for j := lo; j < hi; j++ {
		switch tokens[j].Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF,
			frontkit.TokenIndent, frontkit.TokenDedent:
			continue
		default:
			out = append(out, e.mapper.literal(tokens[j]))
		}
	}
	return out
}
