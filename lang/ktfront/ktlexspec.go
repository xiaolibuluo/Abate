/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// ktLexSpec 是 Kotlin 的纯数据词法表，实现 frontkit.LexSpec。
//
// 三引号原始字符串 """"、字符串模板 "..."、字符字面量与反引号标识符由
// StringQuotes 拼成单个字面量 token；?. ?: !! :: -> 等复合操作符由操作符表
// 按最长优先识别；KDoc 注释由归约期依据前缀识别。
type ktLexSpec struct {
	keywords []string
}

// 编译期证明 Kotlin 词法表满足冻结契约。
var _ frontkit.LexSpec = ktLexSpec{}

// Keywords 返回 Kotlin 保留字表的独立副本。
func (s ktLexSpec) Keywords() []string {
	if len(s.keywords) == 0 {
		return append([]string(nil), ktKeywords...)
	}
	return append([]string(nil), s.keywords...)
}

// Operators 返回 Kotlin 操作符表的独立副本。
func (ktLexSpec) Operators() []string {
	return append([]string(nil), ktOperators...)
}

// LineComment 返回 Kotlin 行注释前缀。
func (ktLexSpec) LineComment() string { return ktLineComment }

// BlockComment 返回 Kotlin 块注释定界符。
func (ktLexSpec) BlockComment() (opening, closing string) {
	return ktBlockCommentOpen, ktBlockCommentClose
}

// StringQuotes 返回三引号、普通、字符与反引号字面量定界符。
func (ktLexSpec) StringQuotes() []string {
	return append([]string(nil), ktStringQuotes...)
}

// RawStringPrefixes 返回三引号原始字符串前缀，交给词法器按最长优先归一化。
func (ktLexSpec) RawStringPrefixes() []string {
	return append([]string(nil), ktRawStringPrefixes...)
}

// NumberSeparators 返回 Kotlin 数字分组分隔符。
func (ktLexSpec) NumberSeparators() []string {
	return append([]string(nil), ktNumberSeparators...)
}

// IndentSensitive 报告 Kotlin 不合成缩进 token。
func (ktLexSpec) IndentSensitive() bool { return false }

// ktStringFlavor 报告一个字符串字面量 token 的专有形态，供降级标注使用。
func ktStringFlavor(text string) string {
	switch {
	case strings.HasPrefix(text, `"""`) && len(text) >= 6:
		return "raw-string"
	case strings.HasPrefix(text, "`") && len(text) >= 2:
		return "backtick-ident"
	case strings.Contains(text, "$"):
		return "string-template"
	default:
		return "plain"
	}
}

// ktIsDocComment 报告一个注释 token 是否为 KDoc 文档注释。
func ktIsDocComment(text string) bool { return strings.HasPrefix(text, "/**") }

// 以下函数是有序词法表与映射表的确定性查找入口，全部线性扫描，绝不依赖
// map 迭代序，也不读取时钟或随机状态。

// ktContains 报告 text 是否为一张有序表里的条目。
func ktContains(table []string, text string) bool {
	for i := range table {
		if table[i] == text {
			return true
		}
	}
	return false
}

// ktNoteValue 从可注入标注表解析一个命名值，表外名字回退到名字本身。
func ktNoteValue(table []ktNote, name string) string {
	for i := range table {
		if table[i].name == name {
			return table[i].value
		}
	}
	return name
}

// ktInfixPower 解析一个操作符的绑定力。
func ktInfixPower(table []ktBinding, text string) int {
	for i := range table {
		if table[i].text == text {
			return table[i].power
		}
	}
	return 0
}

// ktInfixRight 报告一个操作符是否右结合。
func ktInfixRight(table []ktBinding, text string) bool {
	for i := range table {
		if table[i].text == text {
			return table[i].right
		}
	}
	return false
}

// ktDeclRuleFor 解析一个声明主关键字的映射。
func ktDeclRuleFor(lead string) (ktDeclRule, bool) {
	for i := range ktDeclRules {
		if ktDeclRules[i].lead == lead {
			return ktDeclRules[i], true
		}
	}
	return ktDeclRule{}, false
}

// ktDialectFor 解析一个方言别名的规范拼写。
func ktDialectFor(key string) (string, bool) {
	for i := range ktDialectAliases {
		if ktDialectAliases[i].alias == key {
			return ktDialectAliases[i].canonical, true
		}
	}
	return "", false
}

// ktNullSafeFor 从可注入空安全映射表解析一条规则。
func ktNullSafeFor(table []ktNullSafeRule, text string) (ktNullSafeRule, bool) {
	for i := range table {
		if table[i].text == text {
			return table[i], true
		}
	}
	return ktNullSafeRule{}, false
}

// ktHasText 报告区间内是否存在指定拼写的 token。
func ktHasText(tokens []frontkit.Token, lo, hi int, text string) bool {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == text {
			return true
		}
	}
	return false
}

// ktLeadIndex 跳过前导修饰符，返回声明主关键字的索引。
func ktLeadIndex(tokens []frontkit.Token, lo, hi int) int {
	for j := lo; j < hi && j < len(tokens); j++ {
		tok := tokens[j]
		if tok.Text == "" || tok.Kind == frontkit.TokenComment ||
			tok.Kind == frontkit.TokenNewline {
			continue
		}
		if !ktContains(ktModifierWords, tok.Text) {
			return j
		}
	}
	return hi
}

// ktNameOf 挑出构造名，跳过上下文关键字与内建类型拼写。
func ktNameOf(tokens []frontkit.Token, lo, hi int) string {
	for j := lo; j < hi && j < len(tokens); j++ {
		tok := tokens[j]
		if tok.Kind != frontkit.TokenIdent {
			continue
		}
		if ktContains(ktSoftContextWords, tok.Text) ||
			ktContains(ktTypeWords, tok.Text) {
			continue
		}
		return tok.Text
	}
	return ""
}

// ktHasExtensionDot 报告一个 fun 声明是否在参数括号前带接收者类型点号。
func ktHasExtensionDot(tokens []frontkit.Token, lo, hi int) bool {
	seenFun := false
	for j := lo; j < hi && j < len(tokens); j++ {
		text := tokens[j].Text
		if text == "(" || text == "{" || text == "=" {
			break
		}
		if text == "fun" {
			seenFun = true
			continue
		}
		if seenFun && text == "." {
			return true
		}
	}
	return false
}

// ktSpanTokens 把一个 token 区间投影为 AIR 起止位置。
func ktSpanTokens(tokens []frontkit.Token, lo, hi int) (startPos, endPos air.Position) {
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

// ktMapping 是一个 Kotlin 构造归约后的 AIR 种类与语义标注集合。
type ktMapping struct {
	kind       air.NodeKind
	name       string
	note       string
	semKey     air.SemKey
	semValue   string
	async      bool
	partial    bool
	mutability string
}

// ktTemplateParts 抽取一个字符串模板里的 $name 与 ${expr} 片段，产出带标注的
// KindKtStringFrag 片段节点与 KtInterp dedicated 锚点；未闭合的 ${ 安全跳过。
func ktTemplateParts(m *ktMap, tok frontkit.Token) []air.Node {
	out := make([]air.Node, 0, 2)
	text := tok.Text
	for i := 0; i < len(text); i++ {
		if text[i] != '$' || i+1 >= len(text) {
			continue
		}
		stop := 0
		switch {
		case text[i+1] == '{':
			rel := strings.IndexByte(text[i+2:], '}')
			if rel < 0 {
				continue
			}
			stop = i + 2 + rel + 1
		case ktIdentStartByte(text[i+1]):
			stop = i + 2
			for stop < len(text) && ktIdentContinueByte(text[stop]) {
				stop++
			}
		default:
			continue
		}
		part := frontkit.Token{
			Line: tok.Line, Column: tok.Column + i,
			Offset: tok.Offset + i, EndOffset: tok.Offset + stop,
		}
		start, end := part.Span()
		// 插值片段归约的同时在其 "$" 记号位置补发 dedicated 节点。
		marker := m.dedicated(air.KindKtInterp, part, text[i:stop],
			"kt-interp")
		// 插值段的字节由字符串字面量的真实片段承载；包装层改用专用
		// KtStringFrag 节点，不再借 KindInterpolation 伪标签充 dedicated。
		node := air.NewNodeAt(air.KindKtStringFrag, start, end, text[i:stop])
		m.annotate(node, air.SemNote, m.note("template-entry"))
		out = append(out, marker, node)
		i = stop - 1
	}
	return out
}

// ktIdentStartByte 报告一个 ASCII 字节能否开启标识符。
func ktIdentStartByte(ch byte) bool {
	return ch == '_' || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// ktIdentContinueByte 报告一个字节能否延续标识符。
func ktIdentContinueByte(ch byte) bool {
	return ktIdentStartByte(ch) || (ch >= '0' && ch <= '9') || ch >= 0x80
}
