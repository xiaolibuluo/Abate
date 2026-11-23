/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"fmt"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// cParser 是 C/C++ 前端的一次性归约器，持有单文件 token 流与全部局部状态。
//
// 归约点直接产出 AIR 节点：顶层切分出预处理指令、声明、函数、命名空间、
// 模板与记录体，函数体与命名空间体递归归约，因此结果是有层级的真实节点树。
type cParser struct {
	tokens           []frontkit.Token
	path             string
	view             *frontkit.View
	sink             *frontkit.DiagnosticSink
	preproc          *cPreproc
	mapper           *cMapper
	expr             *cExprReducer
	maxTemplateDepth int
}

// newCParser 为一文件构造归约器；所有配置来自 cSpec，无跨文件共享状态。
func newCParser(tokens []frontkit.Token, path string,
	keywords []string, maxDepth int) *cParser {
	if maxDepth <= 0 {
		maxDepth = cDefaultMaxTemplateDepth
	}
	return &cParser{
		tokens:           tokens,
		path:             path,
		sink:             frontkit.NewDiagnosticSink(path, "c"),
		preproc:          newCPreproc(keywords),
		mapper:           newCMapper(),
		expr:             newCExprReducer(),
		maxTemplateDepth: maxDepth,
	}
}

// parse 归约整文件：根 KindPackage 覆盖 [0,size)，子节点经 View 排序裁剪。
func (p *cParser) parse() (air.Node, []air.Diagnostic, error) {
	src := frontkit.SourceFromTokens(p.tokens)
	p.view = frontkit.NewView(src, p.path)
	if _, err := frontkit.NewTokenStream(p.tokens); err != nil {
		_ = p.sink.Warning(1, 0, 1, 1, "token stream validation failed; flat reduction")
	}
	children := p.parseSeq(0, len(p.tokens), true, cBodyStatement)
	size := len(src)
	start := air.NewPosition(1, 1, 0, size)
	end := air.NewPosition(1, 1, size, size)
	root := air.NewNodeAt(air.KindPackage, start, end, p.path,
		p.view.OrderChildren(children)...)
	_ = air.Annotate(root, air.SemScopeID, "translation-unit")
	return root, p.sink.Diagnostics(), nil
}

// cBodyKind marks whether a sequence holds statements or record members.
type cBodyKind int

const (
	cBodyStatement cBodyKind = iota
	cBodyRecord
	cBodyEnum
)

// parseSeq 顺序归约一个 token 区间；top 区分翻译单元与函数体/命名空间体，
// record 标记该区间是否是记录体（字段成员所在的大括号）。
func (p *cParser) parseSeq(lo, hi int, top bool, record cBodyKind) []air.Node {
	if lo < 0 {
		lo = 0
	}
	if hi > len(p.tokens) {
		hi = len(p.tokens)
	}
	out := make([]air.Node, 0, 16)
	for i := lo; i < hi; {
		token := p.tokens[i]
		switch token.Kind {
		case frontkit.TokenEOF:
			return p.view.OrderChildren(out)
		case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent:
			i++
			continue
		case frontkit.TokenComment:
			out = append(out, cLiteralNode(token))
			i++
			continue
		}
		if token.Text == "#" {
			node, next := p.preproc.parseLine(p.tokens, i, hi)
			out = append(out, node)
			out = append(out, p.definedMarkers(i, next)...)
			if next <= i {
				i++
			} else {
				i = next
			}
			continue
		}
		end, bodyOpen, bodyClose := p.unitBounds(i, hi)
		if end <= i {
			end = i + 1
		}
		out = append(out, p.parseUnit(i, end, bodyOpen, bodyClose, top, record))
		i = end
	}
	return p.view.OrderChildren(out)
}

// unitBounds 扫描出一个逻辑构造的结束索引与函数体/记录体花括号区间。
// 括号深度归零时遇到分号或闭合花括号即结束；畸形输入退化为吞并剩余区间，
// 但绝不 panic 也绝不无限循环。
func (p *cParser) unitBounds(i, hi int) (end, bodyOpen, bodyClose int) {
	bodyOpen, bodyClose = -1, -1
	paren, bracket, brace := 0, 0, 0
	for j := i; j < hi; j++ {
		token := p.tokens[j]
		if token.Kind == frontkit.TokenEOF {
			return j, bodyOpen, bodyClose
		}
		switch token.Text {
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

// parseUnit 把一个构造区间归约成带语义标注的 AIR 节点。
func (p *cParser) parseUnit(i, end, bodyOpen, bodyClose int, top bool,
	record cBodyKind) air.Node {
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
	mapping := p.mapper.match(p.tokens, i, end, cNesting(top))
	if record == cBodyRecord {
		mapping = cMemberMapping(mapping, p.tokens, i, end)
	}
	children := make([]air.Node, 0, 8)
	if mapping.kind == air.KindGenericInst {
		if a, b, ok := cAngleRange(p.tokens, i, segEnd); ok {
			children = append(children, p.expr.reduce(p.tokens, i, a)...)
			genericStart, genericEnd := cSpan(p.tokens, a, b)
			generic := air.NewNodeAt(air.KindGenericInst, genericStart, genericEnd,
				"template", p.view.OrderChildren(cLeafNodes(p.tokens, a+1, b-1))...)
			_ = air.Annotate(generic, air.SemTypeParam, "template")
			children = append(children, generic)
			children = append(children, p.expr.reduce(p.tokens, b, segEnd)...)
		} else {
			children = append(children, p.expr.reduce(p.tokens, i, segEnd)...)
		}
	} else {
		children = append(children, p.attributeChildren(i, segEnd)...)
	}
	if bodyOpen >= i && bodyClose >= bodyOpen && bodyClose < end {
		bodyChildren := p.parseSeq(bodyOpen+1, bodyClose, false,
			cRecordKind(mapping.kind))
		blockStart, blockEnd := cSpan(p.tokens, bodyOpen, bodyClose+1)
		block := air.NewNodeAt(air.KindBlock, blockStart, blockEnd, "",
			p.view.OrderChildren(bodyChildren)...)
		children = append(children, block)
	}
	children = append(children,
		p.cDedicatedMarkers(i, end, segEnd, record, children)...)
	children = append(children,
		p.cDeclaratorMarkers(mapping, i, end, segEnd, children)...)
	start, last := cSpan(p.tokens, i, end)
	node := air.NewNodeAt(mapping.kind, start, last, mapping.name,
		p.view.OrderChildren(children)...)
	p.mapper.annotate(node, mapping)
	depth := p.templateDepth(i, segEnd)
	if depth > p.maxTemplateDepth {
		_ = p.sink.Warning(2, start.Offset, start.Line, start.Column,
			fmt.Sprintf("template nesting depth %d exceeds limit %d; reduction flattened",
				depth, p.maxTemplateDepth))
		_ = air.Annotate(node, air.SemPartial, "template-depth")
	}
	return node
}

// attributeChildren 归约一个声明区间的子节点。
//
// 区间开头若有属性前缀（[[attr]] 或 __attribute__((...))），属性单独成为
// 一个 KindAttribute 子节点，其余部分照常归约，这样属性与被声明的实体不会
// 互相吞掉对方的区间。
func (p *cParser) attributeChildren(lo, hi int) []air.Node {
	extent := cAttributeExtent(p.tokens, lo, hi)
	if extent <= lo {
		return p.expr.reduce(p.tokens, lo, hi)
	}
	start, end := cSpan(p.tokens, lo, extent)
	attribute := air.NewNodeAt(air.KindAttribute, start, end, "")
	children := make([]air.Node, 0, 4)
	children = append(children, attribute)
	return append(children, p.expr.reduce(p.tokens, extent, hi)...)
}

// templateDepth 统计尖括号模板嵌套的最大深度，">>" 记两次闭合。
func (p *cParser) templateDepth(lo, hi int) int {
	depth, maxDepth := 0, 0
	for j := lo; j < hi && j < len(p.tokens); j++ {
		switch p.tokens[j].Text {
		case "<":
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		case ">":
			if depth > 0 {
				depth--
			}
		case ">>":
			if depth > 1 {
				depth -= 2
			} else {
				depth = 0
			}
		}
	}
	return maxDepth
}
