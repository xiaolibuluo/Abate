/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"slices"
	"strings"

	"abate/lang/air"
)

// The two interfaces below are grouped into one declaration block. They are
// pure data contracts: every signature refers only to B01 types, the frozen
// lang/air types and standard-library primitives. No implementation type from
// a later batch (Lexer, View, TokenStream, DiagnosticSink, PrattParser) appears
// here, so the contract layer compiles and passes the gate entirely on its own.
type (
	// LexSpec is the pure-data lexical table contract supplied by one language.
	// A later batch consumes these tables to build a deterministic token stream.
	LexSpec interface {
		// Keywords returns the reserved words of the language.
		Keywords() []string
		// Operators returns the operator spellings of the language.
		Operators() []string
		// LineComment returns the line-comment prefix, or an empty string.
		LineComment() string
		// BlockComment returns the opening and closing block-comment delimiters.
		BlockComment() (opening, closing string)
		// StringQuotes returns the string literal delimiters.
		StringQuotes() []string
		// RawStringPrefixes returns the prefixes that disable escaping.
		RawStringPrefixes() []string
		// NumberSeparators returns the digit-group separators.
		NumberSeparators() []string
		// IndentSensitive reports whether the language synthesizes INDENT/DEDENT.
		IndentSensitive() bool
	}

	// LanguageSpec is the single hook a language frontend must implement.
	LanguageSpec interface {
		// Tag returns the canonical short tag (go/py/js/c/rs/cs/kt).
		Tag() string
		// LexSpec returns the lexical tables consumed by the token scanner.
		LexSpec() LexSpec
		// ParseTokens reduces one file's token stream directly into an air.Node.
		ParseTokens(tokens []Token, path string) (air.Node, []air.Diagnostic, error)
		// NormalizeDialect maps a dialect spelling to one canonical string.
		NormalizeDialect(d string) (string, error)
	}
)

// NormalizeTables returns independent, deterministically ordered copies of the
// keyword and operator tables. Entries sort by descending byte length first and
// by lexicographic order second, so longest-match scanning never depends on the
// caller's slice order or on map iteration order.
func NormalizeTables(keywords, operators []string) (normalizedKeywords, normalizedOperators []string) {
	normalizedKeywords = normalizeTable(keywords)
	normalizedOperators = normalizeTable(operators)
	return normalizedKeywords, normalizedOperators
}

// normalizeTable copies one table and applies the frozen deterministic sort key.
func normalizeTable(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	slices.SortStableFunc(out, func(a, b string) int {
		if len(a) != len(b) {
			if len(a) > len(b) {
				return -1
			}
			return 1
		}
		return strings.Compare(a, b)
	})
	return out
}
