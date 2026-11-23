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

// ktParser 是 Kotlin 单文件 token 流的一次性归约器。
//
// 顶层切分出包/导入声明、类、函数、when、伴生对象与扩展函数；花括号体递归
// 归约，表达式片段交给 Pratt 解析器，因此产出的是有层级的真实 AIR 节点树。
type ktParser struct {
	tokens    []frontkit.Token
	path      string
	view      *frontkit.View
	mapper    *ktMap
	expr      *ktExpr
	sink      *frontkit.DiagnosticSink
	whenWords []string
	maxDepth  int
}

// newKtParser 为一文件构造归约器；空表回退到冻结默认值，便于变异测试注入。
func newKtParser(tokens []frontkit.Token, path string, spec *ktSpec) *ktParser {
	whenWords := spec.whenWords
	if len(whenWords) == 0 {
		whenWords = ktWhenBranchWords
	}
	bindings := spec.infix
	if len(bindings) == 0 {
		bindings = ktInfixBindings
	}
	notes := spec.notes
	if len(notes) == 0 {
		notes = ktNotes
	}
	nullSafe := spec.nullSafe
	if len(nullSafe) == 0 {
		nullSafe = ktNullSafeRules
	}
	p := &ktParser{
		tokens: tokens, path: path, whenWords: whenWords,
		sink: frontkit.NewDiagnosticSink(path, "kt"), maxDepth: ktMaxDepth,
	}
	p.mapper = newKtMap(notes, nullSafe)
	p.expr = newKtExpr(bindings, p.mapper)
	return p
}

// parse 归约整文件为覆盖 [0,size) 的 KindPackage 根。
func (p *ktParser) parse() (air.Node, []air.Diagnostic, error) {
	src := frontkit.SourceFromTokens(p.tokens)
	p.view = frontkit.NewView(src, p.path)
	p.mapper.bindView(p.view)
	p.scanIllegal()
	children := p.parseSequence(0, len(p.tokens), ktScopeTop, 0)
	size := len(src)
	start := air.NewPosition(1, 1, 0, size)
	end := air.NewPosition(1, 1, size, size)
	root := air.NewNodeAt(air.KindPackage, start, end, p.path,
		p.view.OrderChildren(children)...)
	_ = air.Annotate(root, air.SemScopeID, "compilation-unit")
	return root, p.sink.Diagnostics(), nil
}

// scanIllegal 对无法分类的字节序列发出确定性的恢复警告。
func (p *ktParser) scanIllegal() {
	for i := range p.tokens {
		tok := p.tokens[i]
		if tok.Kind == frontkit.TokenIllegal {
			_ = p.sink.Warning(2, tok.Offset, tok.Line, tok.Column,
				"illegal byte sequence skipped during reduction")
		}
	}
}

// nestingDepth 统计一个区间的最大括号嵌套深度。
func (p *ktParser) nestingDepth(lo, hi int) int {
	depth, maximum := 0, 0
	for j := lo; j < hi && j < len(p.tokens); j++ {
		switch p.tokens[j].Text {
		case "(", "[", "{", "<":
			depth++
		case ")", "]", "}", ">":
			if depth > 0 {
				depth--
			}
		}
		if depth > maximum {
			maximum = depth
		}
	}
	return maximum
}

// parseSequence 顺序归约一个 token 区间，深度受 ktMaxDepth 约束。
func (p *ktParser) parseSequence(lo, hi int, scope ktScope,
	depth int) []air.Node {
	if lo < 0 {
		lo = 0
	}
	if hi > len(p.tokens) {
		hi = len(p.tokens)
	}
	out := make([]air.Node, 0, 16)
	for i := lo; i < hi; {
		tok := p.tokens[i]
		switch tok.Kind {
		case frontkit.TokenEOF:
			return p.view.OrderChildren(out)
		case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent:
			i++
			continue
		case frontkit.TokenComment:
			out = append(out, p.mapper.comment(tok))
			i++
			continue
		}
		end, bodyOpen, bodyClose := p.unitBounds(i, hi)
		if end <= i {
			end = i + 1
		}
		out = append(out, p.parseUnit(i, end, bodyOpen, bodyClose, scope, depth))
		i = end
	}
	return p.view.OrderChildren(out)
}

// unitBounds 扫描一个逻辑构造的结束索引与花括号体区间；Kotlin 语句以换行
// 结束，因此在括号深度归零时把换行也作为单元边界。
func (p *ktParser) unitBounds(i, hi int) (end, bodyOpen, bodyClose int) {
	bodyOpen, bodyClose = -1, -1
	paren, bracket, brace := 0, 0, 0
	for j := i; j < hi; j++ {
		tok := p.tokens[j]
		if boundary, at := ktUnitBoundary(tok, j, paren, bracket, brace); at {
			return boundary, bodyOpen, bodyClose
		}
		if term, next := p.ktBalanceOrTerminate(tok, j, hi,
			&paren, &bracket, &brace, &bodyOpen, &bodyClose); term {
			return next, bodyOpen, bodyClose
		}
	}
	return hi, bodyOpen, bodyClose
}

// parseUnit 把一个构造区间归约成带语义标注的 AIR 节点。
func (p *ktParser) parseUnit(i, end, bodyOpen, bodyClose int, scope ktScope,
	depth int) air.Node {
	if end <= i {
		end = i + 1
	}
	if end > len(p.tokens) {
		end = len(p.tokens)
	}
	segEnd := end
	if bodyOpen > i && bodyOpen < end {
		segEnd = bodyOpen
	}
	mapping := p.mapper.classify(p.tokens, i, segEnd, scope, bodyOpen >= 0)
	children := p.expr.reduce(p.tokens, i, segEnd)
	bodyScope := ktScopeBody
	switch mapping.kind {
	case air.KindClassDecl, air.KindStructDecl, air.KindInterfaceDecl,
		air.KindEnumDecl, air.KindModuleDecl, air.KindAnnotationDecl:
		bodyScope = ktScopeType
	}
	if bodyOpen >= i && bodyOpen < end {
		bodyHi := bodyClose
		unclosed := bodyClose < bodyOpen || bodyClose >= end
		if unclosed {
			bodyHi = end - 1
		}
		if bodyHi >= bodyOpen {
			bodyChildren := p.parseBody(bodyOpen+1, bodyHi, bodyScope,
				mapping.kind, depth)
			blockStart, blockEnd := ktSpanTokens(p.tokens, bodyOpen, bodyHi+1)
			block := air.NewNodeAt(air.KindBlock, blockStart, blockEnd, "",
				p.view.OrderChildren(bodyChildren)...)
			if unclosed {
				p.mapper.annotate(block, air.SemPartial, "true")
			}
			if mapping.kind == air.KindMatch {
				p.mapper.annotate(block, air.SemNote, p.mapper.note("when"))
			}
			children = append(children, block)
		}
	}
	start, last := ktSpanTokens(p.tokens, i, end)
	children = append(children, p.dedicatedMarkers(mapping, i, segEnd, end,
		bodyOpen, bodyClose, children)...)
	node := air.NewNodeAt(mapping.kind, start, last, mapping.name,
		p.view.OrderChildren(children)...)
	p.mapper.apply(node, mapping)
	if p.nestingDepth(i, segEnd) > p.maxDepth {
		p.mapper.annotate(node, air.SemPartial, "depth")
		_ = p.sink.Warning(3, start.Offset, start.Line, start.Column,
			"nesting depth exceeds the deterministic limit; reduction flattened")
	}
	return node
}

// parseBody 递归归约花括号体；when 语句体按分支切分，深度超限时平铺。
func (p *ktParser) parseBody(lo, hi int, scope ktScope, kind air.NodeKind,
	depth int) []air.Node {
	if depth >= p.maxDepth {
		return p.expr.reduce(p.tokens, lo, hi)
	}
	if kind == air.KindMatch {
		return p.parseWhenBranches(lo, hi)
	}
	return p.parseSequence(lo, hi, scope, depth+1)
}

// parseWhenBranches 把 when 的花括号体按 -> 切成 KindCase 层级，分支内仍由
// Pratt 归约；可注入的 whenWords 表决定切分标记，变异后行为随之改变。
func (p *ktParser) parseWhenBranches(lo, hi int) []air.Node {
	out := make([]air.Node, 0, 8)
	seg := lo
	arrowEnabled := ktContains(p.whenWords, "->")
	for j := lo; j < hi; j++ {
		if arrowEnabled && p.tokens[j].Text == "->" {
			stop := j + 1
			branch := p.caseNode(seg, stop)
			out = append(out, branch)
			seg = stop
		}
	}
	if seg < hi {
		out = append(out, p.expr.reduce(p.tokens, seg, hi)...)
	}
	return p.view.OrderChildren(out)
}

// caseNode 把一个 when 分支投影为带 when-branch 标注的 KindCase 节点；分支
// 归约的同时在 "->" 记号上补发 when_entry 的 dedicated 节点。
func (p *ktParser) caseNode(lo, hi int) air.Node {
	inner := p.expr.reduce(p.tokens, lo, hi)
	start, end := ktSpanTokens(p.tokens, lo, hi)
	arrow := p.tokens[hi-1]
	// 箭头 token 已被分支内归约出的 Lambda 等子节点覆盖，零宽标记锚在
	// 箭头结束偏移处，作为最后子节点不与前序兄弟区间重叠。
	anchor := frontkit.Token{
		Line: arrow.Line, Column: arrow.Column + len(arrow.Text),
		Offset: arrow.EndOffset, EndOffset: arrow.EndOffset,
	}
	marker := p.mapper.dedicated(air.KindKtWhenEntry, anchor,
		arrow.Text, "kt-when-entry")
	node := air.NewNodeAt(air.KindCase, start, end, "",
		append(p.view.OrderChildren(inner), marker)...)
	return p.mapper.annotate(node, air.SemNote, p.mapper.note("when-branch"))
}
