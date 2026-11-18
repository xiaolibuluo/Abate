/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file carries the token-stream context predicates of the dedicated
// markers: each predicate decides, from the surrounding tokens of one anchor,
// whether a specific JavaScript construct genuinely occurs there. They are
// pure token-interval helpers and emit nothing; the emission sites live in
// jsdedicated.go.

// jsSig walks the token stream past newlines, comments and EOF and returns the
// nearest significant token with its index, mirroring jsParser.significant.
func jsSig(toks []frontkit.Token, j, step int) (*frontkit.Token, int) {
	for j >= 0 && j < len(toks) {
		t := toks[j]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			j += step
			continue
		}
		return &t, j
	}
	return nil, -1
}

// jsMatchForward returns the index of the bracket that closes the one at i,
// or -1 when the stream ends first. Only same-kind brackets are counted.
func jsMatchForward(toks []frontkit.Token, i int, open, closing string) int {
	depth := 0
	for j := i; j < len(toks); j++ {
		t := toks[j]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			continue
		}
		switch t.Text {
		case open:
			depth++
		case closing:
			depth--
			if depth == 0 {
				return j
			}
		}
	}
	return -1
}

// jsMatchBackward returns the index of the bracket that opens the one at i.
func jsMatchBackward(toks []frontkit.Token, i int, open, closing string) int {
	depth := 0
	for j := i; j >= 0; j-- {
		t := toks[j]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			continue
		}
		switch t.Text {
		case closing:
			depth++
		case open:
			if depth == 0 {
				return j
			}
			depth--
		}
	}
	return -1
}

// jsIsClassExprContext reports whether the class keyword at i opens a class
// expression: a class keyword directly preceded by an expression-position
// token. A declaration's class keyword is preceded by a statement boundary.
func jsIsClassExprContext(toks []frontkit.Token, i int) bool {
	prev, _ := jsSig(toks, i-1, -1)
	if prev == nil {
		return false
	}
	switch prev.Text {
	case "=", "(", ",", "=>", "[", ":", "?", "!", "return",
		"await", "new", "yield", "throw":
		return true
	default:
		return false
	}
}

// jsIsComputedName reports whether the '[' at i opens a computed property or
// method name: it must sit at member position (after '{', ',' or ';') and its
// closing ']' must be followed by ':', '=' or '(' (property, field or method).
// A bracket that fails this shape is an array literal or index expression.
func jsIsComputedName(toks []frontkit.Token, i int) bool {
	if i <= 0 || toks[i].Text != "[" {
		return false
	}
	prev, _ := jsSig(toks, i-1, -1)
	if prev == nil || (prev.Text != "{" && prev.Text != "," &&
		prev.Text != ";") {
		return false
	}
	shut := jsMatchForward(toks, i, "[", "]")
	if shut < 0 || shut <= i+1 {
		return false
	}
	next, _ := jsSig(toks, shut+1, 1)
	return next != nil && (next.Text == ":" || next.Text == "=" ||
		next.Text == "(")
}

// jsStatementKeywords lists the keywords that may only follow a label colon,
// never a pair value, so a block-level label is never mistaken for a pair.
func jsStatementKeyword(text string) bool {
	switch text {
	case "for", "while", "do", "if", "switch", "try", "var", "let",
		"const", "return", "break", "continue", "throw", "debugger":
		return true
	default:
		return false
	}
}

// jsIsPairColon reports whether the ':' at i separates an object-literal
// pair or a destructuring pair_pattern. The key must sit at member position
// and the value must not start a statement (labels).
func jsIsPairColon(toks []frontkit.Token, i int) bool {
	if i <= 0 || toks[i].Text != ":" {
		return false
	}
	next, _ := jsSig(toks, i+1, 1)
	if next == nil || jsStatementKeyword(next.Text) {
		return false
	}
	key, keyIdx := jsSig(toks, i-1, -1)
	if key == nil || (key.Kind != frontkit.TokenIdent &&
		key.Kind != frontkit.TokenStr && key.Kind != frontkit.TokenNum) {
		return false
	}
	keyPrev, _ := jsSig(toks, keyIdx-1, -1)
	if keyPrev == nil || (keyPrev.Text != "{" && keyPrev.Text != ",") {
		return false
	}
	return true
}

// jsInPatternBraces reports whether the ':' or '=' at i sits inside a brace
// pair that is a destructuring pattern: the brace's closing '}' is directly
// followed by '=' (declaration or assignment destructuring), or the brace is
// nested inside a formal-parameter list.
func jsInPatternBraces(toks []frontkit.Token, i int) bool {
	brace := jsMatchBackward(toks, i, "{", "}")
	if brace < 0 {
		return false
	}
	end := jsMatchForward(toks, brace, "{", "}")
	if end < 0 {
		return false
	}
	if after, _ := jsSig(toks, end+1, 1); after != nil && after.Text == "=" {
		return true
	}
	paren := jsMatchBackward(toks, brace-1, "(", ")")
	if paren < 0 {
		return false
	}
	closeParen := jsMatchForward(toks, paren, "(", ")")
	return closeParen > brace && jsIsFormalParamsOpen(toks, paren)
}

// jsIsObjectAssignDefault reports whether the '=' at i carries an object
// destructuring default: its left operand is a bare identifier at member
// position inside a pattern brace. A plain assignment's left operand is never
// preceded by '{' or ',' inside a pattern, and a var-declarator chain's
// subsequent '=' sits outside any pattern brace.
func jsIsObjectAssignDefault(toks []frontkit.Token, i int) bool {
	if i <= 0 || toks[i].Text != "=" {
		return false
	}
	prev, prevIdx := jsSig(toks, i-1, -1)
	if prev == nil || prev.Kind != frontkit.TokenIdent {
		return false
	}
	prevPrev, _ := jsSig(toks, prevIdx-1, -1)
	if prevPrev == nil || (prevPrev.Text != "{" && prevPrev.Text != ",") {
		return false
	}
	if next, _ := jsSig(toks, i+1, 1); next == nil || next.Text == "=" {
		return false
	}
	return jsInPatternBraces(toks, i)
}

// jsIsRestInParams reports whether the '...' at i opens a rest parameter:
// it must follow a parameter boundary inside a formal-parameter list, either
// a named function's parameter list or an arrow function's parameter list.
// Spread elements inside calls and array/object literals are excluded.
func jsIsRestInParams(toks []frontkit.Token, i int) bool {
	if i <= 0 || toks[i].Text != "..." {
		return false
	}
	prev, _ := jsSig(toks, i-1, -1)
	if prev == nil || (prev.Text != "(" && prev.Text != ",") {
		return false
	}
	paren := jsMatchBackward(toks, i-1, "(", ")")
	if paren < 0 {
		return false
	}
	return jsIsFormalParamsOpen(toks, paren) || jsIsArrowParamsOpen(toks, paren)
}

// jsIsArrowParamsOpen reports whether the '(' at i opens an arrow function's
// parameter list: its matching ')' is directly followed by '=>'.
func jsIsArrowParamsOpen(toks []frontkit.Token, i int) bool {
	if i <= 0 || toks[i].Text != "(" {
		return false
	}
	shut := jsMatchForward(toks, i, "(", ")")
	if shut < 0 {
		return false
	}
	next, _ := jsSig(toks, shut+1, 1)
	return next != nil && next.Text == "=>"
}

// noteMarkers emits dedicated markers for constructs the parser tagged on the
// reduced tree (JSX elements and parenthesized expressions).
func (p *jsParser) noteMarkers(out *[]air.Node) {
	for _, root := range p.nodes {
		air.Walk(root, func(n air.Node) bool {
			note, _ := air.SemanticsOf(n).Get(air.SemNote)
			switch note {
			case noteJSX:
				*out = append(*out, p.mapper.dedicated(
					air.KindJsJsx, n.Pos(), "", "jsx_element"))
			case "grouping":
				*out = append(*out, p.mapper.dedicated(
					air.KindJsParenExpr, n.Pos(), "", "parenthesized_expression"))
			}
			return true
		})
	}
}

// dedicatedMarkers walks the reduced tree and the token stream of one parse
// and emits one zero-width marker per construct the frontend recognized. Each
// marker is anchored on a real token of that construct, never synthesized from
// a detached spelling, so the kind is produced by syntax-node semantics.
func (p *jsParser) dedicatedMarkers() []air.Node {
	out := make([]air.Node, 0, 48)
	for i, tok := range p.toks {
		if tok.Kind == frontkit.TokenNewline || tok.Kind == frontkit.TokenEOF ||
			tok.Kind == frontkit.TokenComment {
			continue
		}
		if !p.keywordMarker(&out, i, tok) {
			p.symbolMarker(&out, i, tok)
		}
	}
	p.noteMarkers(&out)
	return out
}
