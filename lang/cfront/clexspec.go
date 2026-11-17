/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import "abate/lang/frontkit"

// cLexSpec 是 C/C++ 的纯数据词法表，实现 frontkit.LexSpec。
//
// 预处理行以井号开头，由 cpreproc 在 token 流上归约，不引入额外词法种类；
// 字符/字符串前缀与原始字符串前缀交给词法器拼成单个字面量 token。
type cLexSpec struct{}

// 编译期证明词法表满足冻结契约。
var _ frontkit.LexSpec = cLexSpec{}

// Keywords 返回保留字表的独立副本。
func (cLexSpec) Keywords() []string { return append([]string(nil), cKeywords...) }

// Operators 返回操作符表的独立副本。
func (cLexSpec) Operators() []string { return append([]string(nil), cOperators...) }

// LineComment 返回 C/C++ 行注释前缀。
func (cLexSpec) LineComment() string { return cLineComment }

// BlockComment 返回 C/C++ 块注释定界符。
func (cLexSpec) BlockComment() (opening, closing string) {
	return cBlockCommentOpen, cBlockCommentClose
}

// StringQuotes 返回字符串与字符字面量定界符。
func (cLexSpec) StringQuotes() []string {
	return append([]string(nil), cStringQuotes...)
}

// RawStringPrefixes 返回字符/字符串编码前缀与原始字符串前缀。
func (cLexSpec) RawStringPrefixes() []string {
	return append([]string(nil), cRawStringPrefixes...)
}

// NumberSeparators 返回数字分组分隔符。
func (cLexSpec) NumberSeparators() []string {
	return append([]string(nil), cNumberSeparators...)
}

// IndentSensitive 报告 C/C++ 不合成缩进 token。
func (cLexSpec) IndentSensitive() bool { return false }
