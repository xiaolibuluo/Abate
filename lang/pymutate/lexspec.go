/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package pymutate

import (
	"time"

	"abate/lang/frontkit"
)

// pyLexSpec is the pure-data lexical table the mutation scanner consumes.
// The pyfront package keeps its own table unexported, so the backend
// declares the same frozen Python lexical facts here under the frontkit
// contract; the compile-time assertion below keeps the declaration honest
// against the interface, and the lexical facts themselves (keywords,
// operators, comment and string delimiters) are fixed by the Python
// language, not by any single consumer.
type pyLexSpec struct{}

// compile-time proof that the table satisfies the frozen frontkit contract.
var _ frontkit.LexSpec = pyLexSpec{}

// pyKeywords is the Python 3 reserved-word table.
var pyKeywords = []string{
	"False", "None", "True", "and", "as", "assert", "async", "await",
	"break", "class", "continue", "def", "del", "elif", "else", "except",
	"finally", "for", "from", "global", "if", "import", "in", "is",
	"lambda", "nonlocal", "not", "or", "pass", "raise", "return", "try",
	"while", "with", "yield", "match", "case", "type",
}

// pyOperators is the Python operator spelling table, ordered so longest
// match wins during scanning.
var pyOperators = []string{
	"**=", "//=", ">>=", "<<=", "...", "**", "//", "<<", ">>", "<=", ">=",
	"==", "!=", "+=", "-=", "*=", "/=", "%=", "@=", "&=", "|=", "^=",
	"->", ":=", "+", "-", "*", "/", "%", "@", "&", "|", "^", "~",
	"<", ">", "=", "(", ")", "[", "]", "{", "}", ",", ":", ";",
	".",
}

// Keywords returns the Python reserved-word table.
func (pyLexSpec) Keywords() []string { return append([]string(nil), pyKeywords...) }

// Operators returns the Python operator spelling table.
func (pyLexSpec) Operators() []string { return append([]string(nil), pyOperators...) }

// LineComment returns the Python line-comment prefix.
func (pyLexSpec) LineComment() string { return "#" }

// BlockComment returns empty delimiters: Python has no block comment.
func (pyLexSpec) BlockComment() (string, string) { return "", "" }

// StringQuotes returns the Python string literal delimiters.
func (pyLexSpec) StringQuotes() []string {
	return []string{"'''", "\"\"\"", "'", "\""}
}

// RawStringPrefixes returns the prefixes that disable escaping or switch
// literal flavor in Python.
func (pyLexSpec) RawStringPrefixes() []string {
	return []string{"r", "b", "f", "u", "rb", "br", "rf", "fr"}
}

// NumberSeparators returns the digit-group separators Python accepts.
func (pyLexSpec) NumberSeparators() []string { return []string{"_"} }

// IndentSensitive reports that Python synthesizes INDENT/DEDENT tokens.
func (pyLexSpec) IndentSensitive() bool { return true }

// timeoutSeconds converts the configured per-test bound into a duration.
func timeoutSeconds(sec int) time.Duration {
	if sec <= 0 {
		sec = 60
	}
	return time.Duration(sec) * time.Second
}
