/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package jsmutate

import (
	"time"

	"abate/lang/frontkit"
)

// jsLexSpec is the pure-data lexical table the mutation scanner consumes.
// The jsfront package keeps its own table unexported, so the backend
// declares the same frozen JavaScript lexical facts here under the
// frontkit contract; the compile-time assertion below keeps the
// declaration honest against the interface.
type jsLexSpec struct{}

// compile-time proof that the table satisfies the frozen frontkit contract.
var _ frontkit.LexSpec = jsLexSpec{}

// jsKeywords is the JS/TS reserved-word and contextual-keyword table,
// mirroring the frozen jsfront declaration.
var jsKeywords = []string{
	"abstract", "any", "as", "asserts", "async", "await", "bigint", "boolean",
	"break", "case", "catch", "class", "const", "constructor", "continue",
	"debugger", "declare", "default", "delete", "do", "else", "enum", "export",
	"extends", "finally", "for", "from", "function", "get", "global", "if",
	"implements", "import", "in", "infer", "instanceof", "interface", "is",
	"keyof", "let", "module", "namespace", "never", "new", "number", "object",
	"of", "override", "private", "protected", "public", "readonly", "require",
	"return", "satisfies", "set", "static", "string", "super", "switch",
	"symbol", "this", "throw", "try", "type", "typeof", "unique", "unknown",
	"var", "void", "while", "with", "yield",
}

// jsOperators is the JS/TS operator spelling table, mirroring the frozen
// jsfront declaration; frontkit.NormalizeTables orders longest match first.
var jsOperators = []string{
	">>>=", "===", "!==", "**=", "&&=", "||=", "??=",
	"...", "=>", "?.", "??", ">>>", "<<=", ">>=", "==", "!=", "<=", ">=",
	"&&", "||", "++", "--", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=",
	"<<", ">>", "**", "<", ">", "+", "-", "*", "/", "%", "&", "|", "^",
	"!", "=", "~",
}

// Keywords returns the JS/TS keyword table.
func (jsLexSpec) Keywords() []string { return append([]string(nil), jsKeywords...) }

// Operators returns the JS/TS operator spelling table.
func (jsLexSpec) Operators() []string { return append([]string(nil), jsOperators...) }

// LineComment returns the JavaScript line-comment prefix.
func (jsLexSpec) LineComment() string { return "//" }

// BlockComment returns the JavaScript block-comment delimiters.
func (jsLexSpec) BlockComment() (string, string) { return "/*", "*/" }

// StringQuotes returns the JavaScript string literal delimiters, including
// the backtick template literal.
func (jsLexSpec) StringQuotes() []string { return []string{"\"", "'", "`"} }

// RawStringPrefixes returns an empty table: JavaScript has no raw-string
// prefix syntax.
func (jsLexSpec) RawStringPrefixes() []string { return []string{} }

// NumberSeparators returns the digit-group separator of ECMAScript numeric
// literals.
func (jsLexSpec) NumberSeparators() []string { return []string{"_"} }

// IndentSensitive reports false: automatic semicolon insertion is decided
// by newlines and neighbouring tokens, never by indentation depth.
func (jsLexSpec) IndentSensitive() bool { return false }

// timeoutSeconds converts the configured per-test bound into a duration.
func timeoutSeconds(sec int) time.Duration {
	if sec <= 0 {
		sec = 60
	}
	return time.Duration(sec) * time.Second
}
