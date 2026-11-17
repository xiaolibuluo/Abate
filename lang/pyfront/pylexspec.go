/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the Python lexical table, its ordered-table lookups, and the
// pyExpr reductions that consume those tables: string prefixes, interpolation,
// container forms, lambdas and conditional expressions.

// pyLexSpec is the pure-data Python lexical table consumed by frontkit.Lexer.
type pyLexSpec struct{}

// Compile-time proof that the Python tables satisfy the shared contract.
var _ frontkit.LexSpec = pyLexSpec{}

// Keywords returns the frozen Python keyword table.
func (pyLexSpec) Keywords() []string { return pyKeywords }

// Operators returns the frozen Python operator table.
func (pyLexSpec) Operators() []string { return pyOperators }

// LineComment returns the hash line-comment prefix.
func (pyLexSpec) LineComment() string { return "#" }

// BlockComment reports that Python has no block-comment delimiter; triple
// quoted strings are handled by StringQuotes instead.
func (pyLexSpec) BlockComment() (opening, closing string) { return "", "" }

// StringQuotes returns single, double and triple quoted string delimiters.
func (pyLexSpec) StringQuotes() []string { return pyStringQuotes }

// RawStringPrefixes returns the r/f/b/u literal prefixes in longest-first order.
func (pyLexSpec) RawStringPrefixes() []string { return pyRawPrefixes }

// NumberSeparators returns the underscore digit-group separator.
func (pyLexSpec) NumberSeparators() []string { return pyNumberSeparators }

// IndentSensitive enables INDENT and DEDENT synthesis for Python suites.
func (pyLexSpec) IndentSensitive() bool { return true }

// pyContains reports whether text is one entry of an ordered table.
func pyContains(table []string, text string) bool {
	for i := range table {
		if table[i] == text {
			return true
		}
	}
	return false
}

// pyHasPrefix reports whether text starts with one entry of an ordered table.
func pyHasPrefix(text string, prefixes []string) bool {
	for i := range prefixes {
		if strings.HasPrefix(text, prefixes[i]) {
			return true
		}
	}
	return false
}

// pyTableAt indexes an ordered table, clamping out-of-range modes to the tail.
func pyTableAt(table []string, at int) string {
	if len(table) == 0 {
		return ""
	}
	if at < 0 || at >= len(table) {
		return table[len(table)-1]
	}
	return table[at]
}

func pyCloseOf(mode int) string           { return pyTableAt(pyCloseTexts, mode) }
func pyCollectionNote(mode int) string    { return pyTableAt(pyCollectionNotes, mode) }
func pyComprehensionNote(mode int) string { return pyTableAt(pyComprehensionNotes, mode) }

// parseDelimited reduces a parenthesized, list, dict, set or comprehension form.
func (e *pyExpr) parseDelimited(open frontkit.Token, mode int) air.Node {
	e.bracket++
	defer func() { e.bracket-- }()
	e.skipTrivia()
	if tok, ok := e.stream.Peek(); ok && tok.Text == pyCloseOf(mode) {
		e.stream.Advance()
		return e.mapper.collection(open, tok, pyCollectionNote(mode))
	}
	head := e.parse(0)
	e.skipTrivia()
	if mode == pyModeDict {
		if tok, ok := e.stream.Peek(); ok && tok.Text == ":" {
			e.stream.Advance()
			head = e.mapper.pair(head, e.parse(0), tok)
		} else {
			mode = pyModeSet
		}
	}
	e.skipTrivia()
	if tok, ok := e.stream.Peek(); ok && pyContains(pyClauseKeywords, tok.Text) {
		return e.finishComprehension(open, mode, head)
	}
	items := []air.Node{head}
	for {
		e.skipTrivia()
		tok, ok := e.stream.Peek()
		if !ok || tok.Kind == frontkit.TokenEOF {
			break
		}
		if tok.Text == pyCloseOf(mode) {
			e.stream.Advance()
			return e.finishDelimited(open, tok, mode, items)
		}
		if tok.Text == "," || tok.Text == ":" || tok.Text == ";" {
			e.stream.Advance()
			continue
		}
		if pyContains(pyClauseKeywords, tok.Text) || tok.Kind == frontkit.TokenNewline ||
			tok.Kind == frontkit.TokenDedent {
			break
		}
		before := e.stream.Remaining()
		items = append(items, e.parse(0))
		if e.stream.Remaining() >= before {
			e.stream.Advance()
		}
	}
	e.skipTrivia()
	closeTok, _ := e.stream.Peek()
	if closeTok.Text == pyCloseOf(mode) {
		e.stream.Advance()
	}
	return e.finishDelimited(open, closeTok, mode, items)
}

// finishComprehension reduces the clause tail of one comprehension form.
func (e *pyExpr) finishComprehension(open frontkit.Token, mode int,
	head air.Node) air.Node {
	clauses := e.parseClauses()
	e.skipTrivia()
	closeTok, _ := e.stream.Peek()
	if closeTok.Text == pyCloseOf(mode) {
		e.stream.Advance()
	}
	children := append([]air.Node{head}, clauses...)
	if mode == pyModeParen {
		children = append(children, e.mapper.dedicated(
			air.KindPyGenerator, open, closeTok, "",
			"py-generator", children...))
	}
	return e.mapper.comprehension(open, closeTok, pyComprehensionNote(mode), children...)
}

// finishDelimited chooses tuple or composite-lit lowering for one container.
func (e *pyExpr) finishDelimited(open, endTok frontkit.Token, mode int,
	items []air.Node) air.Node {
	if mode == pyModeParen {
		if len(items) == 1 {
			return items[0]
		}
		if len(items) > 1 {
			return e.mapper.tupleNode(open, endTok, items...)
		}
	}
	return e.mapper.collection(open, endTok, pyCollectionNote(mode), items...)
}

// parseString lowers plain, formatted and docstring literals.
func (e *pyExpr) parseString(tok frontkit.Token) air.Node {
	at := strings.IndexAny(tok.Text, "\"'")
	if at > 0 && pyContains(pyFStringPrefixes, tok.Text[:at]) {
		return e.mapper.fstring(tok, e.fstringParts(tok)...)
	}
	if strings.HasPrefix(tok.Text, `"""`) || strings.HasPrefix(tok.Text, `'''`) {
		return e.mapper.annotate(e.mapper.leaf(tok, air.KindDoc),
			air.SemNote, e.mapper.note("docstring"))
	}
	// The string construct has been reduced: project its content and delimiter
	// parts as dedicated string-fragment nodes.
	parts := e.mapper.stringFragments(tok)
	return e.mapper.build(air.KindStrLit, tok, tok, tok.Text, parts...)
}

// fstringParts lowers an f-string's real sub-constructs while the literal is
// reduced: each literal text segment becomes a dedicated string-fragment node
// and each !r/!s/!a conversion keeps its dedicated type-conversion node. The
// {expression} grouping itself has no dedicated AIR kind and its rule is
// locked degraded, so no fabricated grouping node is emitted here; its bytes
// stay covered by the enclosing string-literal node.
func (e *pyExpr) fstringParts(tok frontkit.Token) []air.Node {
	text := tok.Text
	quoteAt := strings.IndexAny(text, "\"'")
	if quoteAt < 0 {
		return nil
	}
	width := pyQuoteWidth(text, quoteAt)
	open := quoteAt + width
	end := len(text) - width
	if end < open {
		return nil
	}
	parts := make([]air.Node, 0, 3)
	seg := open
	for i := open; i < end; i++ {
		if text[i] != '{' || (i+1 < end && text[i+1] == '{') {
			continue
		}
		rel := strings.IndexByte(text[i+1:end], '}')
		if rel < 0 {
			continue
		}
		closeAt := i + 1 + rel
		if i > seg {
			parts = append(parts, e.mapper.stringFragment(tok, seg, i))
		}
		if conv := e.mapper.typeConversion(tok, i, closeAt+1,
			text[i+1:closeAt]); conv != nil {
			parts = append(parts, conv)
		}
		seg = closeAt + 1
		i = closeAt
	}
	if end > seg {
		parts = append(parts, e.mapper.stringFragment(tok, seg, end))
	}
	return parts
}

// pySeparator reports one list separator accepted by the current delimiter mode.
func pySeparator(text string, allowSlice bool) bool {
	return text == "," || text == ";" || (allowSlice && text == ":")
}

// pyAtListEnd reports one token that closes the current list by recovery.
func pyAtListEnd(tok frontkit.Token) bool {
	return pyContains(pyClauseKeywords, tok.Text) ||
		tok.Kind == frontkit.TokenNewline || tok.Kind == frontkit.TokenDedent
}

// parseTernary reduces the conditional expression form.
func (e *pyExpr) parseTernary(left air.Node, ifTok frontkit.Token,
	power int) air.Node {
	condition := e.parse(power + 1)
	e.skipTrivia()
	elseTok, ok := e.stream.Peek()
	if !ok || elseTok.Text != "else" {
		return e.mapper.partial(e.mapper.binary(left, ifTok, condition))
	}
	e.stream.Advance()
	return e.mapper.ternary(left, condition, e.parse(power), elseTok)
}

// parseLambda reduces a lambda expression and its parameter list.
func (e *pyExpr) parseLambda(tok frontkit.Token) air.Node {
	params := make([]air.Node, 0, 2)
	for {
		e.skipTrivia()
		cur, ok := e.stream.Peek()
		if !ok || cur.Kind == frontkit.TokenEOF || cur.Text == ":" {
			break
		}
		if cur.Text == "," || cur.Text == "=" {
			e.stream.Advance()
			continue
		}
		if cur.Kind == frontkit.TokenNewline || cur.Kind == frontkit.TokenDedent {
			break
		}
		before := e.stream.Remaining()
		params = append(params, e.parse(0))
		if e.stream.Remaining() >= before {
			e.stream.Advance()
		}
	}
	e.skipTrivia()
	if cur, ok := e.stream.Peek(); ok && cur.Text == ":" {
		e.stream.Advance()
	}
	return e.mapper.lambdaNode(tok, params, e.parse(0))
}
