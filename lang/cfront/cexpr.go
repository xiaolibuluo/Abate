/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"slices"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// cOpBinding 是一条 Pratt 中缀绑定力，全部为整数常量，绝不引入浮点或时钟。
type cOpBinding struct {
	text  string
	power int
	right bool
}

// cInfixOperators 是确定序的中缀绑定力表。
var cInfixOperators = []cOpBinding{
	{"=", 10, true}, {"+=", 10, true}, {"-=", 10, true}, {"*=", 10, true},
	{"/=", 10, true}, {"%=", 10, true}, {"&=", 10, true}, {"|=", 10, true},
	{"^=", 10, true}, {"<<=", 10, true}, {">>=", 10, true},
	{"||", 20, false}, {"&&", 25, false},
	{"==", 30, false}, {"!=", 30, false}, {"<", 30, false}, {">", 30, false},
	{"<=", 30, false}, {">=", 30, false}, {"<=>", 30, false},
	{"|", 35, false}, {"^", 38, false}, {"&", 40, false},
	{"<<", 45, false}, {">>", 45, false},
	{"+", 50, false}, {"-", 50, false}, {"*", 55, false}, {"/", 55, false},
	{"%", 55, false}, {".", 60, false}, {"->", 60, false}, {"::", 65, false},
	{".*", 60, false}, {"->*", 60, false},
}

// cPrefixOperators 是既可作前缀也可作中缀的一元操作符。
var cPrefixOperators = []string{"*", "&", "+", "-"}

// cExtraPrefixOperators 是仅作前缀的操作符。
var cExtraPrefixOperators = []string{"!", "~", "++", "--"}

// cExprReducer 用 frontkit.PrattParser 归约表达式片段。
//
// 指针/引用由 * 与 & 前缀规则映射为 KindStar/KindAddr，cast 由 sizeof 与
// C++ 显式转换前缀映射为 KindCast，new/delete 映射为 KindUnary，二元与赋值
// 运算映射为 KindBinary/KindAssign 并携带 SemOperator。任何片段归约失败都
// 降级为词法叶子序列，绝不向上抛出错误或 panic。
type cExprReducer struct{}

func newCExprReducer() *cExprReducer { return &cExprReducer{} }

// cIsPrefixOp 报告 spelling 是否拥有一元前缀规则。
func cIsPrefixOp(text string) bool { return slices.Contains(cPrefixOperators, text) }

// reduce 把 [lo,hi) 切成互不重叠的片段并逐段归约，保证每个非 trivia token
// 都落在某个节点的字节区间内。
func (r *cExprReducer) reduce(tokens []frontkit.Token, lo, hi int) []air.Node {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	out := make([]air.Node, 0, hi-lo)
	start := lo
	for i := lo; i < hi; i++ {
		token := tokens[i]
		if token.Kind == frontkit.TokenNewline || token.Kind == frontkit.TokenEOF {
			if start < i {
				out = append(out, r.reduceSegment(tokens, start, i)...)
			}
			start = i + 1
			continue
		}
		if token.Kind == frontkit.TokenComment || token.Text == "," || token.Text == ";" {
			if start < i {
				out = append(out, r.reduceSegment(tokens, start, i)...)
			}
			out = append(out, cLiteralNode(token))
			start = i + 1
		}
	}
	if start < hi {
		out = append(out, r.reduceSegment(tokens, start, hi)...)
	}
	return out
}

// reduceSegment 用 Pratt 解析器归约一段连续 token；失败时降级为词法叶子。
func (r *cExprReducer) reduceSegment(tokens []frontkit.Token, lo, hi int) []air.Node {
	if ternary := r.reduceTernary(tokens, lo, hi); ternary != nil {
		return ternary
	}
	sub := make([]frontkit.Token, 0, hi-lo+1)
	for j := lo; j < hi; j++ {
		token := tokens[j]
		switch token.Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF, frontkit.TokenComment,
			frontkit.TokenIndent, frontkit.TokenDedent:
			continue
		default:
			sub = append(sub, token)
		}
	}
	if len(sub) == 0 {
		return cLeafNodes(tokens, lo, hi)
	}
	last := sub[len(sub)-1]
	eof := frontkit.Token{
		Kind: frontkit.TokenEOF, Offset: last.EndOffset, EndOffset: last.EndOffset,
		Line: last.Line, Column: last.Column,
	}
	sub = append(sub, eof)
	stream, err := frontkit.NewTokenStream(sub)
	if err != nil {
		return cLeafNodes(tokens, lo, hi)
	}
	parser, err := frontkit.NewPrattParser(stream, frontkit.NewView(nil, ""))
	if err != nil {
		return cLeafNodes(tokens, lo, hi)
	}
	r.register(parser)
	node, err := parser.Parse()
	if err != nil || node == nil {
		return cLeafNodes(tokens, lo, hi)
	}
	return []air.Node{node}
}

// register 绑定前缀/中缀规则；注册顺序确定，不使用 map。
func (r *cExprReducer) register(parser *frontkit.PrattParser) {
	literalPrefix := func(_ *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		return cLiteralNode(tok), nil
	}
	for _, kind := range []frontkit.TokenKind{
		frontkit.TokenIdent, frontkit.TokenKeyword, frontkit.TokenNum,
		frontkit.TokenStr, frontkit.TokenChar, frontkit.TokenComment,
	} {
		_ = parser.Register(kind, frontkit.ParseRule{Prefix: literalPrefix})
	}
	unaryPrefix := r.unaryPrefix()
	infix := r.infix()
	for _, op := range cInfixOperators {
		rule := frontkit.ParseRule{
			InfixPower: op.power, RightAssoc: op.right, Infix: infix,
		}
		if cIsPrefixOp(op.text) {
			rule.PrefixPower, rule.Prefix = 70, unaryPrefix
		}
		_ = parser.RegisterText(frontkit.TokenOp, op.text, rule)
	}
	for _, op := range cExtraPrefixOperators {
		_ = parser.RegisterText(frontkit.TokenOp, op,
			frontkit.ParseRule{PrefixPower: 70, Prefix: unaryPrefix})
	}
	keywordPrefix := r.keywordPrefix()
	for _, keyword := range []string{"sizeof", "new", "delete"} {
		_ = parser.RegisterText(frontkit.TokenKeyword, keyword,
			frontkit.ParseRule{PrefixPower: 70, Prefix: keywordPrefix})
	}
	_ = parser.RegisterText(frontkit.TokenPunct, "[",
		frontkit.ParseRule{Postfix: r.postfix()})
}

// unaryPrefix 生成 * / & / ++ / -- 等一元节点。
func (*cExprReducer) unaryPrefix() func(*frontkit.PrattParser, frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		operand, err := parser.ParsePrefixOperand(tok)
		if err != nil || operand == nil {
			return cLiteralNode(tok), nil
		}
		kind := air.KindUnary
		switch tok.Text {
		case "*":
			kind = air.KindStar
		case "&":
			kind = air.KindAddr
		case "++", "--":
			kind = air.KindIncDec
		}
		node := cUnaryNode(kind, tok, operand)
		_ = air.Annotate(node, air.SemOperator, tok.Text)
		return node, nil
	}
}

// keywordPrefix 把 sizeof/new/delete 归约为带 SemNote 的 cast/unary 节点。
func (*cExprReducer) keywordPrefix() func(*frontkit.PrattParser, frontkit.Token) (air.Node, error) {
	return func(parser *frontkit.PrattParser, tok frontkit.Token) (air.Node, error) {
		operand, err := parser.ParsePrefixOperand(tok)
		if err != nil || operand == nil {
			return cLiteralNode(tok), nil
		}
		kind := air.KindUnary
		if tok.Text == "sizeof" {
			kind = air.KindCast
		}
		node := cUnaryNode(kind, tok, operand)
		_ = air.Annotate(node, air.SemNote, tok.Text)
		return node, nil
	}
}

// infix 生成二元/赋值/成员选择节点并写入 SemOperator。
func (*cExprReducer) infix() func(*frontkit.PrattParser, air.Node, frontkit.Token, air.Node) (air.Node, error) {
	return func(_ *frontkit.PrattParser, left air.Node, tok frontkit.Token, right air.Node) (air.Node, error) {
		kind := air.KindBinary
		switch tok.Text {
		case "=", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "<<=", ">>=":
			kind = air.KindAssign
		case ".", "->", "::", ".*", "->*":
			kind = air.KindSelector
		}
		node := cInfixNode(kind, left, tok, right)
		_ = air.Annotate(node, air.SemOperator, tok.Text)
		return node, nil
	}
}

// cLiteralNode 把一个词法 token 映射为 AIR 叶子节点。
func cLiteralNode(tok frontkit.Token) air.Node {
	start, end := tok.Span()
	kind := air.KindIdent
	note := ""
	switch tok.Kind {
	case frontkit.TokenNum:
		kind = air.KindNumLit
	case frontkit.TokenStr:
		kind = air.KindStrLit
	case frontkit.TokenChar:
		kind = air.KindCharLit
	case frontkit.TokenComment:
		kind = air.KindComment
	case frontkit.TokenOp:
		switch tok.Text {
		case "*":
			kind = air.KindStar
		case "&":
			kind = air.KindAddr
		case "=":
			kind = air.KindAssign
		default:
			kind = air.KindBinary
		}
	default:
		switch tok.Text {
		case "true", "false":
			kind = air.KindBoolLit
		case "nullptr", "NULL":
			kind = air.KindNilLit
		case "sizeof":
			kind, note = air.KindCast, "sizeof"
		case "new", "delete":
			kind, note = air.KindUnary, tok.Text
		case "static_cast", "dynamic_cast", "const_cast", "reinterpret_cast":
			kind, note = air.KindCast, "cast"
		default:
			if cIsTypeWord(tok.Text) || cIsKeyword(tok.Text) && cIsSpecifier(tok.Text) {
				kind = air.KindNamed
			}
		}
	}
	if tok.Kind == frontkit.TokenOp {
		node := cBoundOperator(kind, start, end, tok)
		if note != "" {
			_ = air.Annotate(node, air.SemNote, note)
		}
		return node
	}
	node := air.NewNode(kind, start, tok.Text)
	if note != "" {
		_ = air.Annotate(node, air.SemNote, note)
	}
	return node
}

// cUnaryNode 构造覆盖 [tok.Start, operand.End) 的一元节点。
func cUnaryNode(kind air.NodeKind, tok frontkit.Token, operand air.Node) air.Node {
	end := operand.End()
	start := air.NewPosition(tok.Line, tok.Column, tok.Offset, end.Offset)
	return cBoundOperator(kind, start, end, tok, operand)
}

// cInfixNode 构造覆盖 [left.Start, right.End) 的二元节点。
func cInfixNode(kind air.NodeKind, left air.Node, tok frontkit.Token, right air.Node) air.Node {
	startPos := left.Pos()
	end := right.End()
	start := air.NewPosition(startPos.Line, startPos.Column, startPos.Offset, end.Offset)
	return cBoundOperator(kind, start, end, tok, left, right)
}
