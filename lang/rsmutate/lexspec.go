/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package rsmutate

import (
	"time"

	"abate/lang/frontkit"
)

// rsLexSpec is the pure-data lexical table the mutation scanner consumes.
// The rsfront package keeps its own table unexported, so the backend
// declares the same frozen Rust lexical facts here under the frontkit
// contract; the compile-time assertion below keeps the declaration honest
// against the interface.
type rsLexSpec struct{}

// compile-time proof that the table satisfies the frozen frontkit contract.
var _ frontkit.LexSpec = rsLexSpec{}

// rsKeywords lists Rust 2015-2024 reserved words, including macro_rules,
// mirroring the frozen rsfront declaration.
var rsKeywords = []string{
	"as", "async", "await", "break", "const", "continue", "crate", "dyn", "else", "enum",
	"extern", "false", "fn", "for", "if", "impl", "in", "let", "loop", "macro_rules",
	"match", "mod", "move", "mut", "pub", "ref", "return", "self", "Self", "static",
	"struct", "super", "trait", "true", "type", "unsafe", "use", "where", "while",
	"abstract", "become", "box", "do", "final", "macro", "override", "priv", "typeof",
	"unsized", "virtual", "yield", "try", "union", "default",
}

// rsOperators holds Rust operator spellings, mirroring the frozen rsfront
// declaration. The single quote is an operator, not a string delimiter,
// so a lifetime such as 'a stays a two-token sequence; char literals are
// recombined by the reducer. NormalizeTables sorts longest first, so the
// compound spellings win over their single-character prefixes.
var rsOperators = []string{
	">>=", "<<=", "..=", "...", "::", "->", "=>", "==", "!=", "<=", ">=", "&&", "||",
	"<<", ">>", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "..", "'", "&", "|",
	"^", "+", "-", "*", "/", "%", "<", ">", "=", "!", "?", "@", "~",
}

// rsStringQuotes holds the ordinary double-quoted string delimiter,
// mirroring the frozen rsfront declaration. The single quote is
// deliberately absent so lifetimes are never swallowed by the lexer.
var rsStringQuotes = []string{`"`}

// rsRawPrefixes covers raw strings r"..", r#".."#, br".." and c".." plus
// their hash variants, mirroring the frozen rsfront declaration.
var rsRawPrefixes = []string{
	"br##", "cr##", "r###", "br#", "cr#", "r##", "br", "cr", "r#",
	"r", "b", "c",
}

// rsNumberSeparators is the Rust digit-group separator.
var rsNumberSeparators = []string{"_"}

// Keywords returns the Rust keyword table.
func (rsLexSpec) Keywords() []string { return append([]string(nil), rsKeywords...) }

// Operators returns the Rust operator spelling table.
func (rsLexSpec) Operators() []string { return append([]string(nil), rsOperators...) }

// LineComment returns the Rust line-comment prefix.
func (rsLexSpec) LineComment() string { return "//" }

// BlockComment returns the Rust block-comment delimiters.
func (rsLexSpec) BlockComment() (string, string) { return "/*", "*/" }

// StringQuotes returns the Rust string literal delimiters.
func (rsLexSpec) StringQuotes() []string { return append([]string(nil), rsStringQuotes...) }

// RawStringPrefixes returns the Rust raw-string prefix table.
func (rsLexSpec) RawStringPrefixes() []string { return append([]string(nil), rsRawPrefixes...) }

// NumberSeparators returns the Rust digit-group separator.
func (rsLexSpec) NumberSeparators() []string { return append([]string(nil), rsNumberSeparators...) }

// IndentSensitive reports false: Rust is a brace-delimited language whose
// tokens never depend on indentation depth.
func (rsLexSpec) IndentSensitive() bool { return false }

// timeoutSeconds converts the configured per-test bound into a duration.
func timeoutSeconds(sec int) time.Duration {
	if sec <= 0 {
		sec = 60
	}
	return time.Duration(sec) * time.Second
}
