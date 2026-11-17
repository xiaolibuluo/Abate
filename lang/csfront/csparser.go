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

// csParser 是 C# 单文件 token 流的一次性归约器。
//
// 顶层切分出井号指令、特性标记、命名空间/类型声明与成员；花括号体递归归约，
// LINQ 查询按子句切分，因此产出的是有层级的真实 AIR 节点树而非空壳。
type csParser struct {
	tokens       []frontkit.Token
	path         string
	view         *frontkit.View
	mapper       *csMap
	expr         *csExpr
	sink         *frontkit.DiagnosticSink
	linqKeywords []string
	linqClauses  []string
	maxDepth     int
	typeNames    []string
}

// newCsParser 为一文件构造归约器；空表回退到冻结默认值，便于变异测试注入。
func newCsParser(tokens []frontkit.Token, path string, spec *csSpec) *csParser {
	linq := spec.linqKeywords
	if len(linq) == 0 {
		linq = csLinqKeywords
	}
	clauses := spec.linqClauses
	if len(clauses) == 0 {
		clauses = csLinqClauseKeywords
	}
	bindings := spec.infix
	if len(bindings) == 0 {
		bindings = csInfixBindings
	}
	notes := spec.notes
	if len(notes) == 0 {
		notes = csNotes
	}
	p := &csParser{
		tokens: tokens, path: path, linqKeywords: linq, linqClauses: clauses,
		sink: frontkit.NewDiagnosticSink(path, "cs"), maxDepth: csMaxDepth,
	}
	p.mapper = newCsMap(notes, linq)
	p.expr = newCsExpr(bindings, clauses, p.mapper)
	return p
}

// parse 归约整文件为覆盖 [0,size) 的 KindPackage 根。
func (p *csParser) parse() (air.Node, []air.Diagnostic, error) {
	src := frontkit.SourceFromTokens(p.tokens)
	p.view = frontkit.NewView(src, p.path)
	p.mapper.bindView(p.view)
	p.scanIllegal()
	children := p.parseSequence(0, len(p.tokens), csScopeTop)
	size := len(src)
	start := air.NewPosition(1, 1, 0, size)
	end := air.NewPosition(1, 1, size, size)
	root := air.NewNodeAt(air.KindPackage, start, end, p.path,
		p.view.OrderChildren(children)...)
	_ = air.Annotate(root, air.SemScopeID, "compilation-unit")
	return root, p.sink.Diagnostics(), nil
}

// parseSequence 顺序归约一个 token 区间。
func (p *csParser) parseSequence(lo, hi int, scope csScope) []air.Node {
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
		if tok.Text == "#" {
			node, next := p.parseDirective(i, hi)
			out = append(out, node)
			i = next
			continue
		}
		if tok.Text == "[" {
			node, next := p.parseAttribute(i, hi)
			out = append(out, node)
			i = next
			continue
		}
		end, bodyOpen, bodyClose := p.unitBounds(i, hi)
		if end <= i {
			end = i + 1
		}
		out = append(out, p.parseUnit(i, end, bodyOpen, bodyClose, scope))
		i = end
	}
	return p.view.OrderChildren(out)
}

// unitBounds 扫描一个逻辑构造的结束索引与花括号体区间。
func (p *csParser) unitBounds(i, hi int) (end, bodyOpen, bodyClose int) {
	bodyOpen, bodyClose = -1, -1
	paren, bracket, brace := 0, 0, 0
	for j := i; j < hi; j++ {
		tok := p.tokens[j]
		if tok.Kind == frontkit.TokenEOF {
			return j, bodyOpen, bodyClose
		}
		switch tok.Text {
		case "(":
			paren++
		case ")":
			if paren > 0 {
				paren--
			}
		case "[":
			bracket++
		case "]":
			if bracket > 0 {
				bracket--
			}
		case "{":
			if brace == 0 && bodyOpen < 0 {
				bodyOpen = j
			}
			brace++
		case "}":
			if brace > 0 {
				brace--
			}
			if brace == 0 && bodyOpen >= 0 {
				bodyClose = j
				next := j + 1
				if next < hi && p.tokens[next].Text == ";" {
					next++
				}
				return next, bodyOpen, bodyClose
			}
		case ";":
			if paren == 0 && bracket == 0 && brace == 0 {
				return j + 1, bodyOpen, bodyClose
			}
		}
	}
	return hi, bodyOpen, bodyClose
}

// currentTypeName returns the innermost enclosing type name the reducer has
// pushed while recursing into a class/struct/interface body.
func (p *csParser) currentTypeName() string {
	if len(p.typeNames) == 0 {
		return ""
	}
	return p.typeNames[len(p.typeNames)-1]
}

// parseUnit 把一个构造区间归约成带语义标注的 AIR 节点。
func (p *csParser) parseUnit(i, end, bodyOpen, bodyClose int, scope csScope) air.Node {
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
	children := make([]air.Node, 0, 8)
	linqAt := p.linqStart(i, segEnd)
	switch {
	case linqAt == i:
		children = append(children, p.parseLinq(i, segEnd))
	case linqAt > i:
		children = append(children, p.expr.reduce(p.tokens, i, linqAt)...)
		children = append(children, p.parseLinq(linqAt, segEnd))
	default:
		children = append(children, p.expr.reduce(p.tokens, i, segEnd)...)
	}
	// 归约器在进入花括号体前记录外围类型名；构造函数的判定依赖它。
	typeName := p.currentTypeName()
	if mapping.kind == air.KindClassDecl || mapping.kind == air.KindStructDecl ||
		mapping.kind == air.KindInterfaceDecl {
		if name := csNameOf(p.tokens, i, segEnd); name != "" {
			p.typeNames = append(p.typeNames, name)
			defer func() { p.typeNames = p.typeNames[:len(p.typeNames)-1] }()
		}
	}
	bodyScope := csScopeBody
	switch mapping.kind {
	case air.KindClassDecl, air.KindStructDecl, air.KindInterfaceDecl,
		air.KindEnumDecl, air.KindModuleDecl:
		bodyScope = csScopeType
	}
	if bodyOpen >= i && bodyOpen < end {
		// 花括号未闭合时把剩余 token 全部收进一个部分块，既不丢词法投影，
		// 也用 SemPartial 明确标记降级，畸形输入因此保持可恢复。
		bodyHi := bodyClose
		unclosed := bodyClose < bodyOpen || bodyClose >= end
		if unclosed {
			bodyHi = end - 1
		}
		if bodyHi >= bodyOpen {
			bodyChildren := p.parseSequence(bodyOpen+1, bodyHi, bodyScope)
			if mapping.note == "property" || mapping.note == "indexer" {
				bodyChildren = append(bodyChildren,
					p.accessorMarkers(bodyOpen, bodyHi)...)
			}
			blockStart, blockEnd := csSpanTokens(p.tokens, bodyOpen, bodyHi+1)
			block := air.NewNodeAt(air.KindBlock, blockStart, blockEnd, "",
				p.view.OrderChildren(bodyChildren)...)
			if mapping.note == "property" || mapping.note == "indexer" {
				p.mapper.annotate(block, air.SemNote, p.mapper.note("property-accessors"))
			}
			if unclosed {
				p.mapper.annotate(block, air.SemPartial, "true")
			}
			children = append(children, block)
		}
	}
	start, last := csSpanTokens(p.tokens, i, end)
	children = append(children, p.dedicatedMarkers(mapping, i, segEnd,
		bodyOpen, linqAt, scope, typeName, children)...)
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
