/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package ktmutate

import (
	"time"

	"abate/lang/frontkit"
)

// ktLexSpec is the pure-data lexical table the mutation scanner consumes.
// The ktfront package keeps its own table unexported, so the backend
// declares the same frozen Kotlin lexical facts here under the frontkit
// contract; the compile-time assertion below keeps the declaration honest
// against the interface.
type ktLexSpec struct{}

// compile-time proof that the table satisfies the frozen frontkit contract.
var _ frontkit.LexSpec = ktLexSpec{}

// ktKeywords is the Kotlin hard-keyword and common-modifier table,
// mirroring the frozen ktfront declaration. Backtick identifiers and
// string templates stay out of the table, exactly as ktfront declares.
var ktKeywords = []string{
	"abstract", "actual", "annotation", "as", "break", "by", "catch", "class",
	"companion", "const", "constructor", "continue", "crossinline", "data",
	"delegate", "do", "dynamic", "else", "enum", "expect", "external", "false",
	"field", "file", "final", "finally", "for", "fun", "get", "if", "import",
	"in", "infix", "init", "inline", "inner", "interface", "internal", "is",
	"lateinit", "noinline", "null", "object", "open", "operator", "out",
	"override", "package", "param", "private", "property", "protected", "public",
	"receiver", "reified", "return", "sealed", "set", "setparam", "super",
	"suspend", "tailrec", "this", "throw", "true", "try", "typealias", "typeof",
	"val", "value", "var", "vararg", "when", "where", "while",
}

// ktOperators is the Kotlin operator spelling table, mirroring the frozen
// ktfront declaration; frontkit.NormalizeTables orders longest match
// first, so ?. ?: !! :: -> ..< win over their single-character prefixes.
var ktOperators = []string{
	"..<", "===", "!==", "?.", "?:", "!!", "::", "->", "..", "++", "--",
	"+=", "-=", "*=", "/=", "%=", "==", "!=", "<=", ">=", "&&", "||",
	"+", "-", "*", "/", "%", "=", "<", ">", "!", "&", "|", "^", "~",
	"?", ".", ",", ":", "@", "$",
}

// ktStringQuotes covers the triple-quoted raw string, the ordinary string
// template, the character literal and the backtick identifier, in that
// longest-match-first order, mirroring the frozen ktfront declaration.
var ktStringQuotes = []string{`"""`, `"`, `'`, "`"}

// ktRawStringPrefixes marks the triple-quoted raw string opener,
// mirroring the frozen ktfront declaration.
var ktRawStringPrefixes = []string{`"""`}

// ktNumberSeparators is the Kotlin digit-group separator.
var ktNumberSeparators = []string{"_"}

// Keywords returns the Kotlin keyword table.
func (ktLexSpec) Keywords() []string { return append([]string(nil), ktKeywords...) }

// Operators returns the Kotlin operator spelling table.
func (ktLexSpec) Operators() []string { return append([]string(nil), ktOperators...) }

// LineComment returns the Kotlin line-comment prefix.
func (ktLexSpec) LineComment() string { return "//" }

// BlockComment returns the Kotlin block-comment delimiters.
func (ktLexSpec) BlockComment() (string, string) { return "/*", "*/" }

// StringQuotes returns the Kotlin string, character and backtick
// delimiters.
func (ktLexSpec) StringQuotes() []string { return append([]string(nil), ktStringQuotes...) }

// RawStringPrefixes returns the Kotlin raw-string prefix table.
func (ktLexSpec) RawStringPrefixes() []string { return append([]string(nil), ktRawStringPrefixes...) }

// NumberSeparators returns the Kotlin digit-group separator.
func (ktLexSpec) NumberSeparators() []string { return append([]string(nil), ktNumberSeparators...) }

// IndentSensitive reports false: Kotlin is a brace-delimited language
// whose tokens never depend on indentation depth.
func (ktLexSpec) IndentSensitive() bool { return false }

// timeoutSeconds converts the configured per-test bound into a duration.
func timeoutSeconds(sec int) time.Duration {
	if sec <= 0 {
		sec = 60
	}
	return time.Duration(sec) * time.Second
}
