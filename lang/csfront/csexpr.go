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

// csExpr 用 frontkit.PrattParser 归约 C# 表达式区间。
//
// 前缀规则覆盖 LINQ 子句关键字、await、new、typeof 与 nameof；中缀规则覆盖
// 空条件 ?.、null 合并 ??、lambda =>、可空后缀 ? 与三元冒号，绑定力全部是
// 有序整数表。任何片段归约失败都降级为词法叶子序列，绝不向上抛出错误。
type csExpr struct {
	bindings    []csBinding
	linqClauses []string
	mapper      *csMap
}

// newCsExpr 绑定可注入的绑定表、LINQ 子句表与节点映射器。
func newCsExpr(bindings []csBinding, linqClauses []string, mapper *csMap) *csExpr {
	if len(bindings) == 0 {
		bindings = csInfixBindings
	}
	if len(linqClauses) == 0 {
		linqClauses = csLinqClauseKeywords
	}
	return &csExpr{bindings: bindings, linqClauses: linqClauses, mapper: mapper}
}

// reduce 把 [lo,hi) 切成互不重叠的片段并逐段归约，保证每个词法 token 都
// 落在某个 AIR 叶子中，从而满足 token 级保真与位置区间覆盖。
func (e *csExpr) reduce(tokens []frontkit.Token, lo, hi int) []air.Node {
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
				if leaf := e.mapper.literal(tok); leaf != nil {
					out = append(out, leaf)
				}
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
func (e *csExpr) reduceSegment(tokens []frontkit.Token, lo, hi int) []air.Node {
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
			if leaf := e.mapper.literal(tok); leaf != nil {
				out = append(out, leaf)
			}
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
// 词法叶子。Pratt 片段以操作符结尾时解析器会回退并丢弃局部节点，这一步保证
// 该类片段仍满足 token 级保真，且补叶与既有节点按字节区间互不重叠。
func (e *csExpr) coverLeaves(sub []frontkit.Token, out []air.Node) []air.Node {
	limit := len(sub) - 1 // 去掉末尾的 EOF 哨兵
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
			if leaf := e.mapper.literal(sub[k]); leaf != nil {
				out = append(out, leaf)
			}
		}
	}
	return out
}

// register 绑定前缀/中缀规则；注册顺序确定，不使用 map 迭代。
func (e *csExpr) register(parser *frontkit.PrattParser) {
	literalPrefix := func(_ *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		node := e.mapper.literal(tok)
		if node == nil {
			// 标点/运算符到达前缀位置且无结构归约时让表达式失败，由片段
			// recovery 前进跳过，拼写仍由 BindTokenSpellings 兜底。
			return nil, fmt.Errorf("%w: no reduction for %s %q",
				frontkit.ErrUnknownToken, tok.Kind.String(), tok.Text)
		}
		return node, nil
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
		// "(" 与 "[" 的前缀/后缀归约由 parenPrefix 与 bracketPostfix 负责
		// （见下方注册），此处跳过，避免早注册的裸 token 投影遮蔽真实归约。
		if binding.text == "(" || binding.text == "[" {
			continue
		}
		// 每条中缀规则同时保留前缀投影，使定界符作为片段首 token 时仍能被
		// 归约成叶子而不报错。
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
	linqPrefix := e.clausePrefix()
	for _, keyword := range e.linqClauses {
		rule := frontkit.ParseRule{PrefixPower: csLinqPrefixPower, Prefix: linqPrefix}
		_ = parser.RegisterText(frontkit.TokenIdent, keyword, rule)
		_ = parser.RegisterText(frontkit.TokenKeyword, keyword, rule)
	}
	keywordPrefix := e.keywordPrefix()
	for _, keyword := range []string{"await", "new", "typeof", "sizeof",
		"nameof"} {
		rule := frontkit.ParseRule{PrefixPower: 70, Prefix: keywordPrefix}
		_ = parser.RegisterText(frontkit.TokenKeyword, keyword, rule)
	}
	// 前缀运算符在 Pratt 刚消费运算符时归约为真实的一元节点
	// （prefix_unary_expression → cs.unary exact），不再产出裸运算符叶子。
	// 与中缀绑定共存的运算符（-、+）保留原绑定力与中缀归约。
	opPrefix := e.opPrefix()
	for _, op := range csPrefixOps {
		rule := frontkit.ParseRule{Prefix: opPrefix}
		for _, binding := range e.bindings {
			if binding.text == op.text {
				rule.InfixPower = binding.power
				rule.RightAssoc = binding.right
				rule.Infix = infix
				break
			}
		}
		_ = parser.RegisterText(frontkit.TokenOp, op.text, rule)
	}
	// "(" 与 "["：跟在表达式之后时按调用/索引后缀真实归约——Postfix 不
	// 经过通用中缀的右操作数解析，空实参调用与闭括号都能诚实消费；"(" 在
	// 前缀位置（typeof(int)、(int)x、(a, b)）按括号组真实归约，不再产出
	// ParameterList 伪叶子。
	parenPrefix := e.parenPrefix()
	for _, text := range []string{"(", "["} {
		kind := air.KindCall
		if text == "[" {
			kind = air.KindIndex
		}
		rule := frontkit.ParseRule{Postfix: e.bracketPostfix(kind)}
		if text == "(" {
			rule.Prefix = parenPrefix
		}
		_ = parser.RegisterText(frontkit.TokenPunct, text, rule)
	}
}

// clausePrefix 把一条 LINQ 子句归约为 KindCall 带 linq-clause 标注。
func (e *csExpr) clausePrefix() func(*frontkit.PrattParser,
	frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		operand, err := parser.ParsePrefixOperand(tok)
		if err != nil || operand == nil {
			return e.mapper.literal(tok), nil
		}
		node := e.mapper.unaryNode(air.KindCall, tok, operand)
		return e.mapper.annotate(node, air.SemNote,
			e.mapper.note("linq-clause")+":"+tok.Text), nil
	}
}

// leafNodes 把区间内每个非 trivia 的词法 token 映射为冻结叶子；标点与
// 运算符没有独立 AIR 种类，literal 返回 nil，跳过。
func (e *csExpr) leafNodes(tokens []frontkit.Token, lo, hi int) []air.Node {
	out := make([]air.Node, 0, hi-lo)
	for j := lo; j < hi; j++ {
		switch tokens[j].Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF,
			frontkit.TokenIndent, frontkit.TokenDedent:
			continue
		default:
			if leaf := e.mapper.literal(tokens[j]); leaf != nil {
				out = append(out, leaf)
			}
		}
	}
	return out
}
