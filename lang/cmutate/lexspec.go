/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cmutate

import (
	"time"

	"abate/lang/frontkit"
)

// cLexSpec is the pure-data lexical table the mutation scanner consumes.
// The cfront package keeps its own table unexported, so the backend
// declares the same frozen C/C++ lexical facts here under the frontkit
// contract; the compile-time assertion below keeps the declaration honest
// against the interface.
type cLexSpec struct{}

// compile-time proof that the table satisfies the frozen frontkit contract.
var _ frontkit.LexSpec = cLexSpec{}

// cKeywords is the C11 / C++17 / C++20 reserved-word table, mirroring the
// frozen cfront declaration.
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

// cOperators is the C/C++ operator spelling table, mirroring the frozen
// cfront declaration; frontkit.NormalizeTables orders longest match first.
// The # preprocessing punctuator stays out of the table, exactly as the
// cfront lexer declares it, so directive lines tokenize identically.
var cOperators = []string{
	"<=>", ">>=", "<<=", "->*", "...", "::", ".*", "->", "++", "--", "<<",
	">>", "<=", ">=", "==", "!=", "&&", "||", "+=", "-=", "*=", "/=", "%=",
	"&=", "|=", "^=", "##", "+", "-", "*", "/", "%", "&", "|", "^", "~",
	"!", "=", "<", ">", "?", ":", ".",
}

// Keywords returns the C/C++ keyword table.
func (cLexSpec) Keywords() []string { return append([]string(nil), cKeywords...) }

// Operators returns the C/C++ operator spelling table.
func (cLexSpec) Operators() []string { return append([]string(nil), cOperators...) }

// LineComment returns the C line-comment prefix.
func (cLexSpec) LineComment() string { return "//" }

// BlockComment returns the C block-comment delimiters.
func (cLexSpec) BlockComment() (string, string) { return "/*", "*/" }

// StringQuotes returns the C string and character literal delimiters.
func (cLexSpec) StringQuotes() []string { return []string{"\"", "'"} }

// RawStringPrefixes covers the C++ raw-string and encoding prefixes,
// mirroring the frozen cfront declaration.
func (cLexSpec) RawStringPrefixes() []string {
	return []string{"u8R", "uR", "UR", "LR", "R", "u8", "u", "U", "L"}
}

// NumberSeparators returns the C++14 digit-group separator.
func (cLexSpec) NumberSeparators() []string { return []string{"'"} }

// IndentSensitive reports false: C is a brace-delimited language whose
// tokens never depend on indentation depth.
func (cLexSpec) IndentSensitive() bool { return false }

// timeoutSeconds converts the configured per-test bound into a duration.
func timeoutSeconds(sec int) time.Duration {
	if sec <= 0 {
		sec = 60
	}
	return time.Duration(sec) * time.Second
}
