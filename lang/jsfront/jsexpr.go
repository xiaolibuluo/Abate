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

// jsExprParser 是 token 区间上的 Pratt 表达式归约器.
type jsExprParser struct {
	p     *jsParser
	i     int
	end   int
	depth int
	gen   int
}

// exprParser 返回绑定同一 token 流的新表达式归约器.
func (p *jsParser) exprParser() *jsExprParser {
	return &jsExprParser{p: p, end: len(p.toks)}
}

// parseRange 归约 [from,to) 区间并兜底投影所有未覆盖的 lexical token.
func (e *jsExprParser) parseRange(from, to int) []air.Node {
	e.i, e.end = from, to
	if e.i < 0 {
		e.i = 0
	}
	if e.end > len(e.p.toks) {
		e.end = len(e.p.toks)
	}
	if e.end < e.i {
		e.end = e.i
	}
	out := make([]air.Node, 0, 4)
	limit := e.end - e.i + 1
	for steps := 0; e.i < e.end && steps < limit; steps++ {
		before := e.i
		if node := e.parse(0); node != nil {
			out = append(out, node)
		}
		if e.i <= before {
			e.i = before + 1
		}
	}
	e.p.mapper.cover(from, e.end, &out)
	return out
}

// peek 返回当前 token 而不移动游标.
func (e *jsExprParser) peek() (frontkit.Token, bool) {
	if e.i < 0 || e.i >= e.end || e.i >= len(e.p.toks) {
		return frontkit.Token{}, false
	}
	return e.p.toks[e.i], true
}

// peekAt 返回指定索引的 token 而不移动游标.
func (e *jsExprParser) peekAt(j int) (frontkit.Token, bool) {
	if j < 0 || j >= e.end || j >= len(e.p.toks) {
		return frontkit.Token{}, false
	}
	return e.p.toks[j], true
}

// parse 以最小绑定力归约一个表达式，是 Pratt 优先级攀升主循环.
func (e *jsExprParser) parse(minPower int) air.Node {
	if e.i >= e.end || e.depth > 48 || e.gen > 32 {
		return e.p.mapper.empty(e.i)
	}
	e.depth++
	defer func() { e.depth-- }()
	left := e.prefix()
	return e.parseLoop(left, minPower)
}

// prefix 归约一个前缀 token：字面量、一元运算、分组、字面量与专有构造.
func (e *jsExprParser) prefix() air.Node {
	e.skipNewlines()
	if e.i >= e.end || e.gen > 32 {
		return e.p.mapper.empty(e.i)
	}
	idx := e.i
	tok := e.p.toks[idx]
	e.i++
	return e.prefixDispatch(idx, tok)
}

// parseMember 归约成员访问右侧的属性名 token.
func (e *jsExprParser) parseMember() air.Node {
	if e.i >= e.end {
		return e.p.mapper.empty(e.i)
	}
	idx := e.i
	tok := e.p.toks[idx]
	e.i++
	if jsIsLexicalToken(tok.Kind) {
		return e.p.mapper.leaf(idx, tok)
	}
	return e.p.mapper.opaque(idx, tok)
}

// findClosing 在表达式区间内寻找与 open 匹配的 close，损坏输入收敛到区间末.
func (e *jsExprParser) findClosing(idx int, open, closing string) int {
	depth := 0
	for j := idx; j < e.end; j++ {
		switch e.p.toks[j].Text {
		case open:
			depth++
		case closing:
			depth--
			if depth == 0 {
				return j
			}
		}
	}
	return e.end
}

// parseCallLike 归约调用或下标，把括号内参数 token 全量投影为子节点.
func (e *jsExprParser) parseCallLike(left air.Node, openIdx int,
	kind air.NodeKind, note string) air.Node {
	open := "("
	closing := ")"
	if kind == air.KindIndex {
		open, closing = "[", "]"
	}
	closeIdx := e.findClosing(openIdx, open, closing)
	if closeIdx > e.end {
		closeIdx = e.end
	}
	// 递归使用受限代际的 Pratt 解析器解析参数/下标，使内部的比较与算术运算符
	// 也能绑定成真实的二元节点（与 groupPrefix 一致），从而被单运算符变异捕获。
	children := []air.Node{left}
	// A real call reduces its argument list: anchor a dedicated arguments
	// marker on the opening parenthesis so the construct is modelled by AIR
	// rather than lowered onto the shared Region kind.
	if kind == air.KindCall && openIdx < len(e.p.toks) {
		children = append(children, e.p.mapper.dedicated(
			air.KindJsArguments, tokenPos(e.p.toks[openIdx]), "(", "arguments"))
	}
	sub := &jsExprParser{p: e.p, end: closeIdx, gen: e.gen + 1}
	children = append(children, sub.parseRange(openIdx+1, closeIdx)...)
	e.i = closeIdx
	if closeIdx < e.end && e.p.toks[closeIdx].Text == closing {
		e.i = closeIdx + 1
	}
	n := e.p.mapper.wrapChildren(kind, left, e.i-1, "", children...)
	n = e.p.mapper.annotate(n, air.SemNote, note)
	if note == "optional-call" {
		return e.p.mapper.annotate(n, air.SemOperator, "optional-chain")
	}
	return n
}

// parseTernary 归约条件表达式，真实 KindTernary 加 SemOperator.
func (e *jsExprParser) parseTernary(left air.Node, openIdx int) air.Node {
	colon := -1
	for j := openIdx + 1; j < e.end; j++ {
		if e.p.toks[j].Text == ":" {
			colon = j
			break
		}
	}
	children := []air.Node{left}
	if colon < 0 {
		// 无冒号时把 '?' 之后的全部 token 当作一个表达式序列递归归约。
		sub := &jsExprParser{p: e.p, end: e.end, gen: e.gen + 1}
		children = append(children, sub.parseRange(openIdx+1, e.end)...)
	} else {
		// 两个分支各自递归归约，分支内部的运算符绑定成真实二元节点。
		subT := &jsExprParser{p: e.p, end: colon, gen: e.gen + 1}
		children = append(children, subT.parseRange(openIdx+1, colon)...)
		subF := &jsExprParser{p: e.p, end: e.end, gen: e.gen + 1}
		children = append(children, subF.parseRange(colon+1, e.end)...)
	}
	e.i = e.end
	n := e.p.mapper.wrapChildren(air.KindTernary, left, e.end-1, "?:", children...)
	return e.p.mapper.annotate(n, air.SemOperator, "conditional")
}
