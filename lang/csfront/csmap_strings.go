/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// stringLiteral 归约逐字字符串、插值字符串与普通字符串。
//
// 归约的同时按 pyfront.stringFragments 的方式产出字符串片段的 dedicated
// 子节点：内容、转义与收尾定界符都是字符串字面量真实区间的一部分。
func (m *csMap) stringLiteral(tok frontkit.Token) air.Node {
	flavor := csStringFlavor(tok.Text)
	var children []air.Node
	if flavor == "interpolated" || flavor == "verbatim-interpolated" {
		// 插值字符串构造的 dedicated 节点锚定在字面量起始位置，作为字符串
		// 叶子的子节点产出；插值片段占据内容子区间，因此内容片段只补
		// 收尾定界符，避免片段与插值节点兄弟重叠。
		marker := m.dedicated(air.KindCsInterpString, tok, tok.Text,
			"interpolation")
		children = append(children, marker)
		children = append(children, m.interpolationSegments(tok)...)
		children = append(children, m.delimiterFragment(tok)...)
	} else {
		children = m.stringFragments(tok)
	}
	node := m.leafKind(tok, air.KindStrLit, children...)
	if flavor != "plain" {
		m.annotate(node, air.SemNote, m.note(flavor))
		m.annotate(node, air.SemTag, flavor)
	}
	return node
}

// interpolationSegments 抽取一个插值字符串里的字面文本段，产出到位的
// KindCsStringFrag 内容片段；{表达式} 分组本身没有独立 AIR 种类且其规则锁定
// 降级，不产伪造分组节点，其字节由外层字符串字面量节点覆盖。双大括号是
// 字面转义，按普通文本处理。
func (m *csMap) interpolationSegments(tok frontkit.Token) []air.Node {
	text := tok.Text
	quoteAt := strings.IndexByte(text, '"')
	if quoteAt < 0 {
		return nil
	}
	contentStart := quoteAt + quoteWidth(text, quoteAt)
	contentEnd := len(text) - quoteWidth(text, quoteAt)
	if contentEnd < contentStart {
		return nil
	}
	out := make([]air.Node, 0, 2)
	seg := contentStart
	part := func(from, to int) {
		if to <= from {
			return
		}
		node := frontkit.Token{
			Line: tok.Line, Column: tok.Column + from,
			Offset: tok.Offset + from, EndOffset: tok.Offset + to,
		}
		start, end := node.Span()
		child := air.NewNodeAt(air.KindCsStringFrag, start, end, text[from:to])
		out = append(out, m.annotate(child, air.SemNote,
			m.note("cs-string-content")))
	}
	for i := contentStart; i < contentEnd; i++ {
		if text[i] != '{' || (i+1 < contentEnd && text[i+1] == '{') {
			continue
		}
		rel := strings.IndexByte(text[i+1:contentEnd], '}')
		if rel < 0 {
			continue
		}
		closeAt := i + 1 + rel
		part(seg, i)
		seg = closeAt + 1
		i = closeAt
	}
	part(seg, contentEnd)
	return out
}
