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

// csLexSpec 是 C# 的纯数据词法表，实现 frontkit.LexSpec。
//
// 逐字字符串 @"..."、插值字符串 $"..." 与 $@"..." 由 RawStringPrefixes 拼成
// 单个字面量 token；可空后缀 ?、空条件 ?.、null 合并 ?? 由操作符表识别；
// #region / #endregion 等指令由井号 punct 与后续标识符在归约期识别。
type csLexSpec struct {
	keywords []string
}

// 编译期证明 C# 词法表满足冻结契约。
var _ frontkit.LexSpec = csLexSpec{}

// Keywords 返回 C# 保留字表的独立副本。
func (s csLexSpec) Keywords() []string {
	if len(s.keywords) == 0 {
		return append([]string(nil), csKeywords...)
	}
	return append([]string(nil), s.keywords...)
}

// Operators 返回 C# 操作符表的独立副本。
func (csLexSpec) Operators() []string {
	return append([]string(nil), csOperators...)
}

// LineComment 返回 C# 行注释前缀。
func (csLexSpec) LineComment() string { return csLineComment }

// BlockComment 返回 C# 块注释定界符。
func (csLexSpec) BlockComment() (opening, closing string) {
	return csBlockCommentOpen, csBlockCommentClose
}

// StringQuotes 返回字符串与字符字面量定界符。
func (csLexSpec) StringQuotes() []string {
	return append([]string(nil), csStringQuotes...)
}

// RawStringPrefixes 返回逐字与插值字符串前缀，交给词法器按最长优先归一化。
func (csLexSpec) RawStringPrefixes() []string {
	return append([]string(nil), csRawStringPrefixes...)
}

// NumberSeparators 返回 C# 数字分组分隔符。
func (csLexSpec) NumberSeparators() []string {
	return append([]string(nil), csNumberSeparators...)
}

// IndentSensitive 报告 C# 不合成缩进 token。
func (csLexSpec) IndentSensitive() bool { return false }

// csStringFlavor 报告一个字符串字面量 token 的专有形态，供降级标注使用。
func csStringFlavor(text string) string {
	switch {
	case strings.HasPrefix(text, "$@") || strings.HasPrefix(text, "@$"):
		return "verbatim-interpolated"
	case strings.HasPrefix(text, "$"):
		return "interpolated"
	case strings.HasPrefix(text, "@"):
		return "verbatim"
	default:
		return "plain"
	}
}

// csDirectiveName 从井号指令行的 token 文本序列里取出指令名，确定性。
func csDirectiveName(texts []string) string {
	for i := range texts {
		if texts[i] == "" || texts[i] == "#" {
			continue
		}
		return texts[i]
	}
	return ""
}

// csIsDocComment 报告一个注释 token 是否为 XML 文档注释。
func csIsDocComment(text string) bool { return strings.HasPrefix(text, "///") }

// 以下函数是有序词法表与映射表的确定性查找入口，全部线性扫描，绝不依赖
// map 迭代序，也不读取时钟或随机状态。

// csContains 报告 text 是否为一张有序表里的条目。
func csContains(table []string, text string) bool {
	for i := range table {
		if table[i] == text {
			return true
		}
	}
	return false
}

// csNoteFor 从可注入标注表解析一个命名值。
func csNoteFor(table []csNote, name string) string {
	for i := range table {
		if table[i].name == name {
			return table[i].value
		}
	}
	return ""
}

// csInfixPower 解析一个操作符的绑定力。
func csInfixPower(table []csBinding, text string) int {
	for i := range table {
		if table[i].text == text {
			return table[i].power
		}
	}
	return 0
}

// csInfixRight 报告一个操作符是否右结合。
func csInfixRight(table []csBinding, text string) bool {
	for i := range table {
		if table[i].text == text {
			return table[i].right
		}
	}
	return false
}

// csDeclRuleFor 解析一个声明首关键字的映射。
func csDeclRuleFor(lead string) (csDeclRule, bool) {
	for i := range csDeclRules {
		if csDeclRules[i].lead == lead {
			return csDeclRules[i], true
		}
	}
	return csDeclRule{}, false
}

// csDialectFor 解析一个方言别名的规范拼写。
func csDialectFor(key string) (string, bool) {
	for i := range csDialectAliases {
		if csDialectAliases[i].alias == key {
			return csDialectAliases[i].canonical, true
		}
	}
	return "", false
}

// csHasText 报告区间内是否存在指定拼写的 token。
func csHasText(tokens []frontkit.Token, lo, hi int, text string) bool {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == text {
			return true
		}
	}
	return false
}

// csLeadIndex 跳过前导修饰符，返回声明首关键字的索引。
func csLeadIndex(tokens []frontkit.Token, lo, hi int) int {
	for j := lo; j < hi && j < len(tokens); j++ {
		tok := tokens[j]
		if tok.Text == "" || tok.Kind == frontkit.TokenComment ||
			tok.Kind == frontkit.TokenNewline {
			continue
		}
		if !csContains(csModifierWords, tok.Text) {
			return j
		}
	}
	return hi
}

// csNameOf 挑出构造名，跳过 LINQ 上下文关键字与内建类型拼写。
func csNameOf(tokens []frontkit.Token, lo, hi int) string {
	for j := lo; j < hi && j < len(tokens); j++ {
		tok := tokens[j]
		if tok.Kind != frontkit.TokenIdent {
			continue
		}
		if csContains(csLinqKeywords, tok.Text) || csContains(csTypeWords, tok.Text) {
			continue
		}
		return tok.Text
	}
	return ""
}

// csHasLinqTail 报告区间内是否存在查询终结子句，使用可注入的关键字表。
func csHasLinqTail(tokens []frontkit.Token, lo, hi int, linq []string) bool {
	for j := lo; j < hi && j < len(tokens); j++ {
		text := tokens[j].Text
		if (text == "select" || text == "group" || text == "where") &&
			csContains(linq, text) {
			return true
		}
	}
	return false
}

// csSpanTokens 把一个 token 区间投影为 AIR 起止位置。
func csSpanTokens(tokens []frontkit.Token, lo, hi int) (startPos, endPos air.Position) {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	if lo >= hi {
		zero := air.NewPosition(1, 1, 0, 0)
		return zero, zero
	}
	first := tokens[lo]
	last := tokens[hi-1]
	span := air.NewPosition(first.Line, first.Column, first.Offset, last.EndOffset)
	end := air.NewPosition(last.Line, last.Column, last.EndOffset, last.EndOffset)
	return span, end
}
