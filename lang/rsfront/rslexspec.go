/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the Rust lexical table and the literal, delimiter, closure
// and generic reductions that consume it, mirroring the split used by pyfront.

// RsLexSpec is the pure-data Rust lexical table consumed by the frontkit lexer.
//
// Lifetimes ('a) survive because the single quote is an operator, not a quote.
// Attributes with brackets stay two punctuation tokens.
// Macro calls keep the bang as an operator.
// Raw strings use r# and br prefixes; byte or C strings use b and c prefixes.
type rsLexSpec struct {
	keywords []string
}

// Compile-time proof that the Rust tables satisfy the shared contract.
var _ frontkit.LexSpec = rsLexSpec{}

// Keywords returns the frozen Rust keyword table.
func (s rsLexSpec) Keywords() []string {
	if len(s.keywords) == 0 {
		return rsKeywords
	}
	return s.keywords
}

// Operators returns the frozen Rust operator table.
func (rsLexSpec) Operators() []string { return rsOperators }

// LineComment returns the Rust line-comment prefix.
func (rsLexSpec) LineComment() string { return "//" }

// BlockComment returns the Rust block-comment delimiters.
func (rsLexSpec) BlockComment() (opening, closing string) { return "/*", "*/" }

// StringQuotes returns the double-quote string delimiter only.
func (rsLexSpec) StringQuotes() []string { return rsStringQuotes }

// RawStringPrefixes returns raw, byte and C string prefixes in longest order.
func (rsLexSpec) RawStringPrefixes() []string { return rsRawPrefixes }

// NumberSeparators returns the underscore digit-group separator.
func (rsLexSpec) NumberSeparators() []string { return rsNumberSeparators }

// IndentSensitive reports that Rust does not synthesize INDENT and DEDENT.
func (rsLexSpec) IndentSensitive() bool { return false }

// parseDelimited reduces the tokens after an opening delimiter up to its
// matching close. An unmatched delimiter consumes the remaining interval and
// uses the last token as the end, so every produced node stays in bounds.
func (e *rsExpr) parseDelimited(openTok frontkit.Token,
	closeText string) ([]air.Node, frontkit.Token) {
	start := e.at
	closeIdx := -1
	depth := 1
	for j := start; j < e.hi; j++ {
		text := e.tokens[j].Text
		if text == openTok.Text {
			depth++
		}
		if text == closeText {
			depth--
			if depth <= 0 {
				closeIdx = j
				break
			}
		}
	}
	if closeIdx < 0 {
		inner := e.reduceSub(start, e.hi)
		last := openTok
		if e.hi-1 >= start && e.hi-1 < len(e.tokens) {
			last = e.tokens[e.hi-1]
		}
		e.at = e.hi
		return inner, last
	}
	inner := e.reduceSub(start, closeIdx)
	e.at = closeIdx + 1
	return inner, e.tokens[closeIdx]
}

// reduceSub reduces one sub-interval with a fresh cursor and an inherited
// depth, so nested delimiters never share mutable state and never overflow.
func (e *rsExpr) reduceSub(lo, hi int) []air.Node {
	if lo >= hi {
		return nil
	}
	if e.depth >= e.maxDepth {
		if hi-1 < len(e.tokens) {
			node := e.mapper.build(air.KindPartialNode, e.tokens[lo], e.tokens[hi-1], "")
			return []air.Node{e.mapper.partial(node)}
		}
		return nil
	}
	sub := rsExpr{
		tokens: e.tokens, mapper: e.mapper, infix: e.infix,
		at: lo, hi: hi, depth: e.depth + 1, maxDepth: e.maxDepth,
		extra: make([]air.Node, 0),
	}
	return sub.reduceAll()
}

// findAngleClose locates the matching greater-than of a generic argument list,
// bailing out at statement boundaries so comparisons are not misread as types.
func (e *rsExpr) findAngleClose(at int) (int, bool) {
	depth := 0
	for j := at; j < e.hi; j++ {
		switch e.tokens[j].Text {
		case "<":
			depth++
		case ">":
			depth--
			if depth == 0 {
				return j, true
			}
		case ">>":
			depth -= 2
			if depth <= 0 {
				return j, true
			}
		case ";", "{", "=", "=>":
			return 0, false
		}
	}
	return 0, false
}

// parseQuotePrefix recombines a char literal or lowers a lifetime parameter.
func (e *rsExpr) parseQuotePrefix(tok frontkit.Token) air.Node {
	limit := e.at + 3
	if limit > e.hi {
		limit = e.hi
	}
	for j := e.at; j < limit; j++ {
		if e.tokens[j].Text == "'" {
			end := e.tokens[j]
			inner := e.collectParamLeaves(e.at, j)
			e.at = j + 1
			node := e.mapper.build(air.KindCharLit, tok, end, "", inner...)
			return e.mapper.annotate(node, air.SemNote,
				e.mapper.note("char-literal"))
		}
	}
	if e.at < e.hi {
		name := e.tokens[e.at]
		e.at++
		return e.mapper.lifetime(tok, name)
	}
	return e.mapper.leaf(tok, air.KindIdent)
}

// finishClosure reduces a |params| body or a || body closure to KindLambda.
func (e *rsExpr) finishClosure(start, open frontkit.Token, noteName string) air.Node {
	params := make([]air.Node, 0, 2)
	if open.Text != "" {
		closeIdx := -1
		for j := e.at; j < e.hi; j++ {
			if e.tokens[j].Text == "|" {
				closeIdx = j
				break
			}
		}
		if closeIdx >= 0 {
			params = e.collectParamLeaves(e.at, closeIdx)
			e.at = closeIdx + 1
		} else {
			params = e.collectParamLeaves(e.at, e.hi)
			e.at = e.hi
		}
	}
	body := e.parseExpression(0)
	end := start
	if body != nil {
		end = spanToken(body, start)
	}
	node := e.mapper.lambda(start, end, params, body)
	if noteName != "" {
		e.mapper.annotate(node, air.SemNote, e.mapper.note(noteName))
	}
	return node
}

// collectParamLeaves projects identifier tokens of a parameter list directly,
// without recursing, so damaged or empty lists cannot inflate the parse depth.
func (e *rsExpr) collectParamLeaves(lo, hi int) []air.Node {
	out := make([]air.Node, 0, 4)
	for j := lo; j < hi && j < len(e.tokens); j++ {
		if e.tokens[j].Kind == frontkit.TokenIdent ||
			e.tokens[j].Kind == frontkit.TokenNum {
			out = append(out, e.mapper.lexical(e.tokens[j]))
		}
	}
	return out
}
