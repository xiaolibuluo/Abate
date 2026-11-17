/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"slices"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// cKeywords 是 C11 / C++17 / C++20 的保留字全集，供词法器与构造分派共用。
var cKeywords = []string{
	"alignas", "alignof", "and", "and_eq", "asm", "auto", "bitand", "bitor",
	"bool", "break", "case", "catch", "char", "char8_t", "char16_t", "char32_t",
	"class", "co_await", "co_return", "co_yield", "compl", "concept", "const",
	"consteval", "constexpr", "constinit", "const_cast", "continue", "decltype",
	"default", "delete", "do", "double", "dynamic_cast", "else", "enum",
	"explicit", "export", "extern", "false", "float", "for", "friend", "goto",
	"if", "inline", "int", "long", "mutable", "namespace", "new", "noexcept",
	"not", "not_eq", "nullptr", "operator", "or", "or_eq", "private", "protected",
	"public", "register", "reinterpret_cast", "requires", "restrict", "return",
	"short", "signed", "sizeof", "static", "static_assert", "static_cast",
	"struct", "switch", "template", "this", "thread_local", "throw", "true",
	"try", "typedef", "typeid", "typename", "union", "unsigned", "using",
	"virtual", "void", "volatile", "wchar_t", "while", "xor", "xor_eq",
	"_Alignas", "_Alignof", "_Atomic", "_Bool", "_Generic", "_Noreturn",
	"_Static_assert", "_Thread_local",
}

// cOperators 是 C/C++ 操作符表；frontkit.NormalizeTables 按长度降序归一化。
// "#" 刻意不列入，以便预处理行首始终是一个 TokenPunct 井号。
var cOperators = []string{
	"<=>", ">>=", "<<=", "->*", "...", "::", ".*", "->", "++", "--", "<<",
	">>", "<=", ">=", "==", "!=", "&&", "||", "+=", "-=", "*=", "/=", "%=",
	"&=", "|=", "^=", "##", "+", "-", "*", "/", "%", "&", "|", "^", "~",
	"!", "=", "<", ">", "?", ":", ".",
}

// cStringQuotes 是字符串与字符字面量定界符。
var cStringQuotes = []string{"\"", "'"}

// cRawStringPrefixes 覆盖 C++ 原始字符串与字符/字符串编码前缀。
var cRawStringPrefixes = []string{
	"u8R", "uR", "UR", "LR", "R", "u8", "u", "U", "L",
}

// cNumberSeparators 是 C++14 数字分组分隔符。
var cNumberSeparators = []string{"'"}

const (
	cLineComment       = "//"
	cBlockCommentOpen  = "/*"
	cBlockCommentClose = "*/"
)

// cPreprocKeywords 决定哪些 # 指令行被识别为 KindDirective，供变异测试改写。
var cPreprocKeywords = []string{
	"include", "define", "undef", "if", "ifdef", "ifndef", "elif", "else",
	"endif", "pragma", "error", "warning", "line",
}

// cTypeWords 是内建类型拼写，用于把类型引用映射为 KindNamed。
var cTypeWords = []string{
	"void", "bool", "char", "char8_t", "char16_t", "char32_t", "wchar_t",
	"short", "int", "long", "signed", "unsigned", "float", "double", "auto",
}

// cSpecifierWords 是声明说明符，cNameOf 跳过它们以挑出真正的声明名。
var cSpecifierWords = []string{
	"const", "constexpr", "consteval", "constinit", "static", "extern",
	"inline", "virtual", "explicit", "mutable", "register", "volatile",
	"thread_local", "unsigned", "signed", "long", "short", "auto", "void",
	"bool", "char", "char8_t", "char16_t", "char32_t", "wchar_t", "int",
	"float", "double", "struct", "class", "union", "enum", "typename",
	"template", "using", "typedef", "namespace", "friend", "decltype",
}

// cDialect 把一个方言别名绑定到三方言之一的规范拼写。
type cDialect struct {
	alias     string
	canonical string
}

// cDialectAliases 是有序别名表，绝不依赖 map 迭代序。
var cDialectAliases = []cDialect{
	{"", "c11"}, {"c", "c11"}, {"c11", "c11"}, {"c99", "c11"},
	{"c17", "c11"}, {"c18", "c11"}, {"gnu11", "c11"}, {"gnu99", "c11"},
	{"ansi-c", "c11"}, {"iso-c11", "c11"},
	{"c++", "c++17"}, {"cpp", "c++17"}, {"c++11", "c++17"},
	{"c++14", "c++17"}, {"c++17", "c++17"}, {"c++1z", "c++17"},
	{"gnu++17", "c++17"}, {"cpp17", "c++17"},
	{"c++20", "c++20"}, {"c++2a", "c++20"}, {"gnu++20", "c++20"},
	{"cpp20", "c++20"}, {"c++23", "c++20"},
}

// cDialectCanonical 是三方言规范拼写集合。
var cDialectCanonical = []string{"c11", "c++17", "c++20"}

// cIsKeyword 报告 spelling 是否为 C/C++ 保留字。
func cIsKeyword(s string) bool { return slices.Contains(cKeywords, s) }

// cIsTypeWord 报告 spelling 是否为内建类型拼写。
func cIsTypeWord(s string) bool { return slices.Contains(cTypeWords, s) }

// cIsSpecifier 报告 spelling 是否为声明说明符。
func cIsSpecifier(s string) bool { return slices.Contains(cSpecifierWords, s) }

// cAngleRange 返回模板参数尖括号区间 [a,b)，未闭合时仍归约到 hi。
func cAngleRange(tokens []frontkit.Token, lo, hi int) (start, end int, ok bool) {
	open, depth := -1, 0
	for j := lo; j < hi && j < len(tokens); j++ {
		switch tokens[j].Text {
		case "<":
			if depth == 0 {
				open = j
			}
			depth++
		case ">":
			if depth > 0 {
				depth--
			}
			if depth == 0 && open >= 0 {
				return open, j + 1, true
			}
		case ">>":
			if depth > 1 {
				depth -= 2
			} else {
				depth = 0
			}
			if depth == 0 && open >= 0 {
				return open, j + 1, true
			}
		}
	}
	if open >= 0 {
		return open, hi, true
	}
	return 0, 0, false
}

// cSpan 把一个 token 区间投影为 AIR 起止位置。
func cSpan(tokens []frontkit.Token, lo, hi int) (start, end air.Position) {
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
	first, _ := tokens[lo].Span()
	last := tokens[hi-1]
	_, lastEnd := last.Span()
	// air.NewNodeAt 以 pos.EndOffset 作为节点区间终点（与 frontkit.View 一致），
	// 因此这里把区间终点设为最后一个 token 的 EndOffset，保证父节点覆盖全部子节点。
	span := air.NewPosition(first.Line, first.Column, first.Offset, last.EndOffset)
	return span, lastEnd
}

// cLeafNodes 把区间内每个非 trivia token 映射为叶子节点。
func cLeafNodes(tokens []frontkit.Token, lo, hi int) []air.Node {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	out := make([]air.Node, 0, hi-lo)
	for j := lo; j < hi; j++ {
		switch tokens[j].Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF,
			frontkit.TokenIndent, frontkit.TokenDedent:
			continue
		default:
			out = append(out, cLiteralNode(tokens[j]))
		}
	}
	return out
}
