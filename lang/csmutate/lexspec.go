/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package csmutate

import (
	"time"

	"abate/lang/frontkit"
)

// csLexSpec is the pure-data lexical table the mutation scanner consumes.
// The csfront package keeps its own table unexported, so the backend
// declares the same frozen C# lexical facts here under the frontkit
// contract; the compile-time assertion below keeps the declaration honest
// against the interface.
type csLexSpec struct{}

// compile-time proof that the table satisfies the frozen frontkit contract.
var _ frontkit.LexSpec = csLexSpec{}

// csKeywords is the C# reserved-word table, mirroring the frozen csfront
// declaration. LINQ contextual keywords stay out of the table so they
// remain identifiers for the lexer, exactly as csfront declares.
var csKeywords = []string{
	"abstract", "add", "as", "async", "await", "base", "bool", "break", "byte",
	"case", "catch", "char", "checked", "class", "const", "continue", "decimal",
	"default", "delegate", "do", "double", "dynamic", "else", "enum", "event",
	"explicit", "extern", "false", "finally", "fixed", "float", "for", "foreach",
	"get", "goto", "if", "implicit", "in", "int", "interface", "internal", "is",
	"lock", "long", "nameof", "namespace", "new", "null", "object", "operator",
	"out", "override", "params", "partial", "private", "protected", "public",
	"readonly", "record", "ref", "remove", "return", "sbyte", "sealed", "set",
	"short", "sizeof", "stackalloc", "static", "string", "struct", "switch",
	"this", "throw", "true", "try", "typeof", "uint", "ulong", "unchecked",
	"unsafe", "ushort", "using", "value", "var", "virtual", "void", "volatile",
	"while", "yield",
}

// csOperators is the C# operator spelling table, mirroring the frozen
// csfront declaration; frontkit.NormalizeTables orders longest match
// first. The # and @ prefixes stay out of the table so directive lines
// and verbatim identifiers tokenize identically to the csfront lexer.
var csOperators = []string{
	"??=", "<<=", ">>=", ">>>=", "=>", "?.", "??", "++", "--", "->", "::",
	"==", "!=", "<=", ">=", "&&", "||", "+=", "-=", "*=", "/=", "%=", "&=",
	"|=", "^=", "<<", ">>", ">>>", "+", "-", "*", "/", "%", "&", "|", "^",
	"!", "~", "=", "<", ">", "?",
}

// csStringQuotes is the C# string and character literal delimiter table,
// mirroring the frozen csfront declaration.
var csStringQuotes = []string{`"`, `'`}

// csRawStringPrefixes covers verbatim and interpolated string prefixes,
// mirroring the frozen csfront declaration.
var csRawStringPrefixes = []string{"$@", "@$", "$", "@"}

// csNumberSeparators is the C# digit-group separator.
var csNumberSeparators = []string{"_"}

// Keywords returns the C# keyword table.
func (csLexSpec) Keywords() []string { return append([]string(nil), csKeywords...) }

// Operators returns the C# operator spelling table.
func (csLexSpec) Operators() []string { return append([]string(nil), csOperators...) }

// LineComment returns the C# line-comment prefix.
func (csLexSpec) LineComment() string { return "//" }

// BlockComment returns the C# block-comment delimiters.
func (csLexSpec) BlockComment() (string, string) { return "/*", "*/" }

// StringQuotes returns the C# string and character literal delimiters.
func (csLexSpec) StringQuotes() []string { return append([]string(nil), csStringQuotes...) }

// RawStringPrefixes returns the C# verbatim and interpolated prefix table.
func (csLexSpec) RawStringPrefixes() []string { return append([]string(nil), csRawStringPrefixes...) }

// NumberSeparators returns the C# digit-group separator.
func (csLexSpec) NumberSeparators() []string { return append([]string(nil), csNumberSeparators...) }

// IndentSensitive reports false: C# is a brace-delimited language whose
// tokens never depend on indentation depth.
func (csLexSpec) IndentSensitive() bool { return false }

// timeoutSeconds converts the configured per-test bound into a duration.
func timeoutSeconds(sec int) time.Duration {
	if sec <= 0 {
		sec = 60
	}
	return time.Duration(sec) * time.Second
}
