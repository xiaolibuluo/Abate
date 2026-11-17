/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

// 本文件承载对外的 air.Frontend 接线，以及 csParser 的指令、特性、LINQ 子句、
// 深度限制与恢复归约；声明与语句主产生式留在 csparser.go，与 pyfront 的分工一致。

import (
	"fmt"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// CsFrontend 是 C# 前端句柄，是共享 frontkit 骨架绑定到 csSpec 的别名。
type CsFrontend = frontkit.FrontendBase

// 编译期证明适配器满足冻结契约，且返回值满足 air.Frontend。
var (
	_ frontkit.LanguageSpec = (*csSpec)(nil)
	_ air.Frontend          = (*CsFrontend)(nil)
)

// NewCsFrontend 构造 C# 前端，返回满足 air.Frontend 的共享解析骨架。
//
// 骨架负责文件排序、方言归一、词法化、诊断聚合与确定性注释；本包只负责
// 词法表与 ParseTokens 的真归约。零真实文件系统读取、零子进程、零 CGO。
func NewCsFrontend() (air.Frontend, error) {
	spec := &csSpec{
		keywords:     append([]string(nil), csKeywords...),
		linqKeywords: append([]string(nil), csLinqKeywords...),
		linqClauses:  append([]string(nil), csLinqClauseKeywords...),
		infix:        append([]csBinding(nil), csInfixBindings...),
		notes:        append([]csNote(nil), csNotes...),
	}
	return frontkit.NewFrontendBase(spec)
}

// csSpec 实现 frontkit.LanguageSpec，是注入共享解析骨架的 C# 适配器。
//
// 字段可注入以便变异测试改写词法表、绑定表与标注表；零值字段回退到冻结默认值。
type csSpec struct {
	keywords     []string
	linqKeywords []string
	linqClauses  []string
	infix        []csBinding
	notes        []csNote
}

// Tag 返回 C# 的规范短标签。
func (*csSpec) Tag() string { return "cs" }

// LexSpec 返回 C# 词法表。
func (s *csSpec) LexSpec() frontkit.LexSpec {
	return csLexSpec{keywords: s.specKeywords()}
}

// specKeywords 返回关键字表的独立副本，空表回退默认值。
func (s *csSpec) specKeywords() []string {
	if len(s.keywords) == 0 {
		return append([]string(nil), csKeywords...)
	}
	return append([]string(nil), s.keywords...)
}

// NormalizeDialect 把别名归一到 cs8-cs12 五种规范方言之一，确定性。
func (*csSpec) NormalizeDialect(dialect string) (string, error) {
	key := strings.ToLower(strings.TrimSpace(dialect))
	if canonical, ok := csDialectFor(key); ok {
		return canonical, nil
	}
	return "", fmt.Errorf("csfront: unsupported dialect %q (want cs8..cs12)", dialect)
}

// ParseTokens 把一文件的 token 流真归约成有层级的 AIR 节点树。
func (s *csSpec) ParseTokens(tokens []frontkit.Token,
	path string) (air.Node, []air.Diagnostic, error) {
	return newCsParser(tokens, path, s).parse()
}

// scanIllegal 对无法分类的字节序列发出确定性的恢复警告。
func (p *csParser) scanIllegal() {
	for i := range p.tokens {
		tok := p.tokens[i]
		if tok.Kind == frontkit.TokenIllegal {
			_ = p.sink.Warning(2, tok.Offset, tok.Line, tok.Column,
				"illegal byte sequence skipped during reduction")
		}
	}
}

// nestingDepth 统计一个区间的最大括号与尖括号嵌套深度。
func (p *csParser) nestingDepth(lo, hi int) int {
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

// linqStart 找出区间内一个 LINQ 查询表达式的起始索引。
func (p *csParser) linqStart(lo, hi int) int {
	if !csContains(p.linqKeywords, "from") {
		return -1
	}
	for j := lo; j < hi; j++ {
		if p.tokens[j].Text == "from" &&
			csHasLinqTail(p.tokens, j+1, hi, p.linqKeywords) {
			return j
		}
	}
	return -1
}

// parseLinq 把查询区间按子句切分为 KindCall 层级的真实子树。
func (p *csParser) parseLinq(lo, hi int) air.Node {
	starts := make([]int, 0, 4)
	for j := lo; j < hi; j++ {
		if csContains(p.linqClauses, p.tokens[j].Text) {
			starts = append(starts, j)
		}
	}
	clauses := make([]air.Node, 0, len(starts))
	for k, at := range starts {
		stop := hi
		if k+1 < len(starts) {
			stop = starts[k+1]
		}
		clauseLeaves := make([]air.Node, 0, 2)
		if leaf := p.mapper.literal(p.tokens[at]); leaf != nil {
			clauseLeaves = append(clauseLeaves, leaf)
		}
		inner := append(clauseLeaves,
			p.expr.reduce(p.tokens, at+1, stop)...)
		clauseStart, clauseEnd := csSpanTokens(p.tokens, at, stop)
		clause := air.NewNodeAt(air.KindCall, clauseStart, clauseEnd, "",
			p.view.OrderChildren(inner)...)
		p.mapper.annotate(clause, air.SemNote,
			p.mapper.note("linq-clause")+":"+p.tokens[at].Text)
		clauses = append(clauses, clause)
	}
	queryStart, queryEnd := csSpanTokens(p.tokens, lo, hi)
	query := air.NewNodeAt(air.KindCall, queryStart, queryEnd, "",
		p.view.OrderChildren(clauses)...)
	return p.mapper.annotate(query, air.SemNote, p.mapper.note("linq"))
}

// parseDirective 归约一行井号指令；#region 归约为专用区域细节节点。
func (p *csParser) parseDirective(i, hi int) (air.Node, int) {
	j := i + 1
	for j < hi {
		if p.tokens[j].Kind == frontkit.TokenNewline ||
			p.tokens[j].Kind == frontkit.TokenEOF {
			break
		}
		j++
	}
	texts := make([]string, 0, j-i)
	leaves := make([]air.Node, 0, j-i)
	for k := i; k < j; k++ {
		texts = append(texts, p.tokens[k].Text)
		if leaf := p.mapper.literal(p.tokens[k]); leaf != nil {
			leaves = append(leaves, leaf)
		}
	}
	name := csDirectiveName(texts)
	// #region/#endregion 属 preproc_region 构造（cs.preproc exact），与其他
	// 预处理指令一样归约为 KindDirective；不再落到 87-108 伪 kind。
	kind, note := air.KindDirective, p.mapper.note("directive")
	if name == "region" || name == "endregion" {
		note = p.mapper.note(name)
	}
	start, end := csSpanTokens(p.tokens, i, j)
	node := air.NewNodeAt(kind, start, end, "", p.view.OrderChildren(leaves)...)
	p.mapper.annotate(node, air.SemDirective, name)
	p.mapper.annotate(node, air.SemNote, note)
	next := j
	if next <= i {
		next = i + 1
	}
	return node, next
}

// parseAttribute 归约一对方括号里的特性标记。
func (p *csParser) parseAttribute(i, hi int) (air.Node, int) {
	depth, j := 0, i
	for ; j < hi; j++ {
		tok := p.tokens[j]
		if tok.Kind == frontkit.TokenEOF || tok.Kind == frontkit.TokenNewline {
			break
		}
		if tok.Text == "[" {
			depth++
		}
		if tok.Text == "]" {
			depth--
			if depth <= 0 {
				j++
				break
			}
		}
	}
	if j <= i {
		j = i + 1
	}
	if j > hi {
		j = hi
	}
	children := p.expr.reduce(p.tokens, i, j)
	start, end := csSpanTokens(p.tokens, i, j)
	node := air.NewNodeAt(air.KindAttribute, start, end, "",
		p.view.OrderChildren(children)...)
	p.mapper.annotate(node, air.SemNote, p.mapper.note("attribute"))
	next := j
	if next <= i {
		next = i + 1
	}
	return node, next
}
