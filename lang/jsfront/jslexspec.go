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

// jsLexSpec 是 JS/TS 词法契约的纯数据实现，供 frontkit.Lexer 消费.
type jsLexSpec struct{}

// 编译期断言 jsLexSpec 满足共享词法契约.
var _ frontkit.LexSpec = jsLexSpec{}

// Keywords 返回 JS/TS 保留字与上下文关键字表.
func (jsLexSpec) Keywords() []string {
	return append([]string(nil), jsKeywords...)
}

// Operators 返回运算符表，最长匹配由 frontkit 归一.
func (jsLexSpec) Operators() []string {
	return append([]string(nil), jsOperators...)
}

// LineComment 返回行注释前缀.
func (jsLexSpec) LineComment() string { return "//" }

// BlockComment 返回块注释的起始与结束定界符.
func (jsLexSpec) BlockComment() (opening, closing string) { return "/*", "*/" }

// StringQuotes 返回字符串与模板字符串定界符.
func (jsLexSpec) StringQuotes() []string {
	return append([]string(nil), jsStringQuotes...)
}

// RawStringPrefixes 返回空表：JS/TS 无 raw 字符串前缀，标签模板由前缀标识符归约.
func (jsLexSpec) RawStringPrefixes() []string { return []string{} }

// NumberSeparators 返回数字分组分隔符.
func (jsLexSpec) NumberSeparators() []string {
	return append([]string(nil), jsNumberSeparators...)
}

// IndentSensitive 返回 false：ASI 由换行与前后 token 判定，不依赖缩进层级.
func (jsLexSpec) IndentSensitive() bool { return false }

// jsTemplatePart 是模板字符串的一段：静态文本或 ${} 插值表达式.
type jsTemplatePart struct {
	text   string
	expr   bool
	offset int
}

// jsIsTemplate 判定 token 是否为反引号模板字符串.
func jsIsTemplate(tok frontkit.Token) bool {
	return tok.Kind == frontkit.TokenStr && strings.HasPrefix(tok.Text, "`")
}

// jsSplitTemplate 把模板字符串按 ${...} 拆成静态段与插值段，损坏输入也收敛.
func jsSplitTemplate(text string) []jsTemplatePart {
	parts := make([]jsTemplatePart, 0, 4)
	inner, base := text, 0
	if len(inner) > 0 && inner[0] == '`' {
		inner, base = inner[1:], 1
	}
	if len(inner) > 0 && inner[len(inner)-1] == '`' {
		inner = inner[:len(inner)-1]
	}
	for i := 0; i < len(inner); {
		rel := strings.Index(inner[i:], "${")
		if rel < 0 {
			parts = append(parts, jsTemplatePart{text: inner[i:], offset: base + i})
			break
		}
		at := i + rel
		if at > i {
			parts = append(parts, jsTemplatePart{text: inner[i:at], offset: base + i})
		}
		j, depth := at+2, 1
		for j < len(inner) && depth > 0 {
			switch inner[j] {
			case '{':
				depth++
			case '}':
				depth--
			case '\\':
				if j+1 < len(inner) {
					j++
				}
			}
			j++
		}
		if j > len(inner) {
			j = len(inner)
		}
		parts = append(parts, jsTemplatePart{
			text: inner[at:j], expr: true, offset: base + at,
		})
		i = j
	}
	if len(inner) == 0 {
		parts = append(parts, jsTemplatePart{offset: base})
	}
	return parts
}

// jsCanEndExpression 判定 token 能否结束一个表达式，是 ASI 的必要条件.
func jsCanEndExpression(prev *frontkit.Token) bool {
	if prev == nil {
		return false
	}
	switch prev.Kind {
	case frontkit.TokenIdent, frontkit.TokenNum, frontkit.TokenStr,
		frontkit.TokenChar, frontkit.TokenKeyword:
		return true
	}
	switch prev.Text {
	case ")", "]", "}", "++", "--", "`":
		return true
	}
	return false
}

// jsCanInsertSemicolonAfterNewline 判定换行处是否应执行自动分号插入（ASI）.
func jsCanInsertSemicolonAfterNewline(prev, next *frontkit.Token) bool {
	if prev == nil || !jsCanEndExpression(prev) {
		return false
	}
	if jsIsRestricted(prev.Text) {
		return true
	}
	if next == nil {
		return true
	}
	return !jsIsContinuation(next.Text)
}

// jsRegexAllowed 判定 "/" 是否应作为正则字面量起始而非除法运算符.
func jsRegexAllowed(prev *frontkit.Token) bool {
	if prev == nil {
		return true
	}
	if jsIsRegexOpener(prev.Text) {
		return true
	}
	switch prev.Kind {
	case frontkit.TokenIdent, frontkit.TokenNum, frontkit.TokenStr, frontkit.TokenChar:
		return false
	}
	switch prev.Text {
	case ")", "]", "}", "++", "--":
		return false
	}
	return true
}

// jsIsJSXOpen 在 JSX 方言下判定 "<" 后是否跟随元素名或 fragment 闭合尖括号.
func jsIsJSXOpen(next frontkit.Token, d jsDialect) bool {
	if !d.jsx {
		return false
	}
	if next.Text == ">" || next.Text == "/" {
		return true
	}
	return next.Kind == frontkit.TokenIdent || next.Kind == frontkit.TokenKeyword
}

// isUseStrictToken 判定字符串字面量是否为指令序言.
func isUseStrictToken(tok frontkit.Token) bool {
	return tok.Kind == frontkit.TokenStr &&
		(tok.Text == "\"use strict\"" || tok.Text == "'use strict'")
}

// typeFollows 判定 ":" 后是否跟随类型表达式.
func typeFollows(toks []frontkit.Token, i, end int) bool {
	if i+1 >= end || i+1 >= len(toks) {
		return false
	}
	tok := toks[i+1]
	return tok.Kind == frontkit.TokenIdent || jsIsTypeKeyword(tok.Text) ||
		tok.Text == "[" || tok.Text == "{" || tok.Text == "(" || tok.Text == "<"
}

// looksGeneric 判定 "<" 是否为泛型参数列表起始.
func looksGeneric(toks []frontkit.Token, i, end int) bool {
	for j := i + 1; j < end && j < len(toks); j++ {
		switch toks[j].Text {
		case ">":
			return j > i+1
		case ";", "{", "}", "=>", "=", "+", "(", ")", "?", ":":
			return false
		}
	}
	return false
}

// genericPrefix 判定 "<" 前置 token 是类型名或函数名，排除 JSX 起始标签.
func genericPrefix(toks []frontkit.Token, i int) bool {
	if i <= 0 || i >= len(toks) {
		return false
	}
	prev := toks[i-1]
	return prev.Kind == frontkit.TokenIdent || prev.Kind == frontkit.TokenKeyword
}

// jsxFollows 判定 "<" 后是否形成 JSX 起始标签.
func jsxFollows(toks []frontkit.Token, i, end int, d jsDialect) bool {
	if i+1 >= end || i+1 >= len(toks) || !jsIsJSXOpen(toks[i+1], d) {
		return false
	}
	for j := i + 2; j < end && j < len(toks); j++ {
		switch toks[j].Text {
		case ">":
			return true
		case ";", "{":
			return false
		}
	}
	return false
}

// optionalFollows 排除三元运算符后判定 TS 可选标记.
func optionalFollows(toks []frontkit.Token, i, end int) bool {
	if i+1 >= end || i+1 >= len(toks) {
		return false
	}
	next := toks[i+1].Text
	return next != ":" && next != "."
}

// tupleFollows 判定 "[" 是否位于 TS 类型位置.
func tupleFollows(toks []frontkit.Token, i int) bool {
	if i == 0 || i >= len(toks) {
		return false
	}
	prev := toks[i-1]
	return prev.Text == ":" || jsIsTypeKeyword(prev.Text) ||
		prev.Kind == frontkit.TokenIdent
}

// jsxAttrFollows 判定 "=" 是否位于 JSX 起始标签的属性区间内.
func jsxAttrFollows(toks []frontkit.Token, i, start int) bool {
	for j := i - 1; j >= start && j >= 0 && j < len(toks); j-- {
		switch toks[j].Text {
		case "<":
			return true
		case ">", ";":
			return false
		}
	}
	return false
}

// regexOrDivision 依据前置 token 判定 "/" 是正则字面量起始还是容错节点.
//
// 正则字面量内部的运算符 token（例如 /<h1[^>]*>/ 里的 "<"）不属于 lexical
// 保真子集，不会被 cover 投影成叶节点。把整段真实 token 拼写写入外层正则
// 节点的 attribute，才能让这类单字符变异真正改变 AIR-ID。
func (e *jsExprParser) regexOrDivision(idx int, tok frontkit.Token) air.Node {
	if !jsRegexAllowed(e.p.significant(idx-1, -1)) {
		return e.p.mapper.opaque(idx, tok)
	}
	closeIdx := idx
	for j := idx + 1; j < e.end; j++ {
		if e.p.toks[j].Text == "/" {
			closeIdx = j
			break
		}
		if e.p.toks[j].Kind == frontkit.TokenNewline {
			break
		}
	}
	children := make([]air.Node, 0, 4)
	e.p.mapper.cover(idx, closeIdx+1, &children)
	e.i = closeIdx + 1
	flags := ""
	if e.i < e.end && e.p.toks[e.i].Kind == frontkit.TokenIdent {
		flags = e.p.toks[e.i].Text
		children = append(children, e.p.mapper.leaf(e.i, e.p.toks[e.i]))
		e.i++
	}
	n := e.p.mapper.node(idx, e.i-1, air.KindRegion,
		jsTokenSpelling(e.p.toks, idx, e.i), children...)
	n = e.p.mapper.annotate(n, air.SemNote, noteRegex)
	n = e.p.mapper.annotate(n, air.SemOperator, "regex")
	return e.p.mapper.annotate(n, air.SemTag, flags)
}
