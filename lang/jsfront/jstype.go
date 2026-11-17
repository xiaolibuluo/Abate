/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// jsTypeParser 负责 TS 类型、泛型、JSX 与装饰器到真实 AIR 类型类的降级归约.
type jsTypeParser struct {
	p *jsParser
}

// typeParser 返回绑定同一 token 流的类型归约器.
func (p *jsParser) typeParser() *jsTypeParser { return &jsTypeParser{p: p} }

// simpleNode 构造零文本、单 token 区间的降级节点，避免污染保真子集.
func (t *jsTypeParser) simpleNode(i int, kind air.NodeKind, note string) air.Node {
	tok := t.p.toks[i]
	start, end := tok.Span()
	n := air.NewNodeAt(kind, start, end, "")
	_ = air.Annotate(n, air.SemNote, note)
	return n
}

// typeNode 把 ":" 类型注解降级为 KindNamed 并以 SemTypeRef 全量保留类型名.
func (t *jsTypeParser) typeNode(i, end int) air.Node {
	name := ""
	for j := i + 1; j < end; j++ {
		tok := t.p.toks[j]
		if tok.Kind == frontkit.TokenIdent || jsIsTypeKeyword(tok.Text) {
			name = tok.Text
			break
		}
		if tok.Text == "," || tok.Text == ";" || tok.Text == ")" ||
			tok.Text == "=" || tok.Text == "{" {
			break
		}
	}
	n := t.simpleNode(i, air.KindNamed, noteTypeAnnot)
	_ = air.Annotate(n, air.SemTypeRef, name)
	return n
}

// genericNodes 把泛型参数列表降级为 KindGenericInst/KindGenericParam 并写 SemTypeParam.
func (t *jsTypeParser) genericNodes(i, end int) []air.Node {
	out := make([]air.Node, 0, 4)
	names := make([]string, 0, 4)
	for j := i + 1; j < end; j++ {
		tok := t.p.toks[j]
		if tok.Text == ">" {
			break
		}
		if tok.Kind == frontkit.TokenIdent || jsIsTypeKeyword(tok.Text) {
			names = append(names, tok.Text)
			n := t.simpleNode(j, air.KindGenericParam, noteGenericParam)
			_ = air.Annotate(n, air.SemTypeParam, tok.Text)
			out = append(out, n)
		}
	}
	n := t.simpleNode(i, air.KindGenericInst, noteGeneric)
	_ = air.Annotate(n, air.SemTypeParam, strings.Join(names, ","))
	return append(out, n)
}

// decorate 扫描语句区间，产出 TS 类型、泛型、JSX、装饰器、指令与块内语句节点.
func (t *jsTypeParser) decorate(start, end int, _ string) []air.Node {
	p := t.p
	out := make([]air.Node, 0, 8)
	if end > len(p.toks) {
		end = len(p.toks)
	}
	for i := start; i < end; i++ {
		out = t.decorateToken(out, i, start, end, p.toks[i])
	}
	return out
}

// keywordPrefix 归约出现在表达式位置的 JS/TS 关键字构造.
func (e *jsExprParser) keywordPrefix(idx int, tok frontkit.Token) air.Node {
	m := e.p.mapper
	switch tok.Text {
	case "function", "class":
		kind := air.KindFuncLit
		if tok.Text == "class" {
			kind = air.KindClassDecl
		}
		base := m.opaque(idx, tok)
		return air.NewNodeAt(kind, base.Pos(), base.End(), "", base)
	case "new":
		n := m.wrap(air.KindCall, m.opaque(idx, tok), e.parse(16), "new", 0, "")
		return m.annotate(n, air.SemNote, noteNew)
	case "await":
		n := m.wrap(air.KindAwait, m.opaque(idx, tok), e.parse(14), "await", 0, "")
		n = m.annotate(n, air.SemAsync, "true")
		return m.annotate(n, air.SemNote, "await")
	case "yield":
		return m.annotate(m.unary(tok, e.parse(14)), air.SemNote, "yield")
	case "import":
		n := m.wrap(air.KindCall, m.opaque(idx, tok), e.parse(16), "import", 0, "")
		return m.annotate(n, air.SemNote, noteDynamicImport)
	case "async":
		return e.parse(0)
	case "typeof", "void", "delete":
		return m.unary(tok, e.parse(14))
	case "this", "super":
		start, end := tok.Span()
		return air.NewNodeAt(air.KindIdent, start, end, tok.Text)
	default:
		return m.opaque(idx, tok)
	}
}

// groupPrefix 归约圆括号分组，内部递归使用受限代际的 Pratt 解析器.
func (e *jsExprParser) groupPrefix(idx int, _ frontkit.Token) air.Node {
	closeIdx := e.findClosing(idx, "(", ")")
	if closeIdx > e.end {
		closeIdx = e.end
	}
	sub := &jsExprParser{p: e.p, end: closeIdx, gen: e.gen + 1}
	inner := sub.parseRange(idx+1, closeIdx)
	e.i = closeIdx
	if closeIdx < e.end && e.p.toks[closeIdx].Text == ")" {
		e.i = closeIdx + 1
	}
	children := make([]air.Node, 0, len(inner))
	children = append(children, inner...)
	n := e.p.mapper.node(idx, e.i-1, air.KindRegion, "", children...)
	return e.p.mapper.annotate(n, air.SemNote, "grouping")
}

// bracketPrefix 归约数组或对象字面量（以及函数体、类体等块）为真实
// KindCompositeLit，内部递归使用受限代际的 Pratt 解析器，与 groupPrefix 一致。
//
// 这样字面量或块内部的比较/算术运算符也会经 BindOperator 绑定成真实的二元节点，
// 而不是被 cover 当成非保真 token 直接丢弃；单运算符变异因此改变 AIR-ID（被捕获）。
func (e *jsExprParser) bracketPrefix(idx int, open, closing string,
	kind air.NodeKind, note string) air.Node {
	closeIdx := e.findClosing(idx, open, closing)
	if closeIdx > e.end {
		closeIdx = e.end
	}
	sub := &jsExprParser{p: e.p, end: closeIdx, gen: e.gen + 1}
	inner := sub.parseRange(idx+1, closeIdx)
	e.i = closeIdx
	if closeIdx < e.end && e.p.toks[closeIdx].Text == closing {
		e.i = closeIdx + 1
	}
	children := make([]air.Node, 0, len(inner))
	children = append(children, inner...)
	n := e.p.mapper.node(idx, e.i-1, kind, "", children...)
	return e.p.mapper.annotate(n, air.SemNote, note)
}

// arrayPrefix 归约数组字面量.
func (e *jsExprParser) arrayPrefix(idx int, _ frontkit.Token) air.Node {
	return e.bracketPrefix(idx, "[", "]", air.KindCompositeLit, noteArrayLiteral)
}

// objectPrefix 归约对象字面量.
func (e *jsExprParser) objectPrefix(idx int, _ frontkit.Token) air.Node {
	return e.bracketPrefix(idx, "{", "}", air.KindCompositeLit, noteObjectLiteral)
}

// jsxPrefix 把 JSX 起始标签降级为 KindCompositeLit；损坏标签发警告并标 SemPartial.
func (e *jsExprParser) jsxPrefix(idx int, tok frontkit.Token,
	head air.Node) air.Node {
	closeIdx := idx + 1
	for j := idx + 1; j < e.end; j++ {
		if e.p.toks[j].Text == ">" || e.p.toks[j].Text == ";" {
			closeIdx = j
			break
		}
		closeIdx = j + 1
	}
	children := make([]air.Node, 0, 4)
	if head != nil {
		children = append(children, head)
	}
	e.p.mapper.cover(idx+1, closeIdx, &children)
	unclosed := closeIdx >= e.end ||
		(closeIdx < e.end && e.p.toks[closeIdx].Text != ">")
	e.i = closeIdx
	if !unclosed {
		e.i = closeIdx + 1
	}
	n := e.p.mapper.node(idx, e.i-1, air.KindCompositeLit, "", children...)
	n = e.p.mapper.annotate(n, air.SemNote, noteJSX)
	if unclosed {
		_ = e.p.sink.Warning(2, tok.Offset, tok.Line, tok.Column,
			"unterminated JSX element")
		n = e.p.mapper.annotate(n, air.SemPartial, "true")
	}
	return n
}
