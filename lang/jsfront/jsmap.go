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

// 专有构造降级注记：全部写入真实 SemKey，绝不扩展 AIR 枚举.
const (
	noteModule        = "js-module"
	noteJSX           = "jsx"
	noteTemplate      = "template-string"
	noteInterpolation = "template-interpolation"
	noteRegex         = "regex-literal"
	noteOptionalChain = "optional-chain"
	noteNullish       = "nullish"
	noteArrow         = "arrow-function"
	noteDecorator     = "decorator"
	noteSpread        = "spread"
	noteTypeAnnot     = "type-annotation"
	noteGeneric       = "generic-instantiation"
	noteGenericParam  = "generic-parameter"
	noteOptional      = "optional-type"
	noteUnion         = "union-type"
	noteTuple         = "tuple-type"
	noteReadonly      = "readonly"
	noteNew           = "new-expression"
	noteDynamicImport = "dynamic-import"
	noteArrayLiteral  = "array-literal"
	noteObjectLiteral = "object-literal"
	noteJSXAttribute  = "jsx-attribute"
	noteUseStrict     = "use-strict"
)

// jsMapper 是 JS/TS 构造到 AIR 的节点工厂，持有 token 区间视图与投影台账.
type jsMapper struct {
	toks    []frontkit.Token
	view    *frontkit.View
	covered []bool
	d       jsDialect
}

// newJsMapper 用 token 流重建源区间视图，供位置计算与保真投影使用.
func newJsMapper(toks []frontkit.Token, path string, d jsDialect) *jsMapper {
	return &jsMapper{
		toks:    toks,
		view:    frontkit.NewView(frontkit.SourceFromTokens(toks), path),
		covered: make([]bool, len(toks)),
		d:       d,
	}
}

// lexKind 把真实 token 类别映射到 AIR lexical 节点类.
func (*jsMapper) lexKind(tok frontkit.Token) air.NodeKind {
	switch tok.Kind {
	case frontkit.TokenIdent:
		switch tok.Text {
		case "true", "false":
			return air.KindBoolLit
		case "null", "undefined":
			return air.KindNilLit
		default:
			return air.KindIdent
		}
	case frontkit.TokenNum:
		return air.KindNumLit
	case frontkit.TokenStr, frontkit.TokenChar:
		return air.KindStrLit
	case frontkit.TokenComment, frontkit.TokenIllegal:
		return air.KindComment
	default:
		return air.KindRegion
	}
}

// leaf 以 token 的精确半开区间投影一个 AIR lexical 叶节点并登记台账.
func (m *jsMapper) leaf(i int, tok frontkit.Token) air.Node {
	start, end := tok.Span()
	if i >= 0 && i < len(m.covered) {
		m.covered[i] = true
	}
	return air.NewNodeAt(m.lexKind(tok), start, end, tok.Text)
}

// opaque 为关键字、运算符与标点构造结构节点.
//
// 运算符与标点的真实拼写写进 attribute，使 ComputeNodeID 能区分只有该
// token 不同的两棵树；这类 token 不属于 lexical 保真子集，但它们是
// cmp/op/digit 变异的真实目标，零文本会让变异无法被 AIR-ID 捕获。
func (*jsMapper) opaque(_ int, tok frontkit.Token) air.Node {
	start, end := tok.Span()
	return air.NewNodeAt(air.KindRegion, start, end, tok.Text)
}

// empty 构造一个合法零宽节点，用于表达式缺位时的容错归约.
func (m *jsMapper) empty(i int) air.Node {
	pos := air.NewPosition(1, 1, 0, 0)
	if i >= 0 && i < len(m.toks) {
		t := m.toks[i]
		pos = air.NewPosition(t.Line, t.Column, t.Offset, t.Offset)
	}
	return air.NewNodeAt(air.KindRegion, pos, pos, "")
}
