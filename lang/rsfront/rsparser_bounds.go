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

// This file holds the construct-interval scanning and braced-body context
// selection of rsParser: unit bounds balancing, block-variant detection at a
// braced body, and the per-body reduction context rules that drive the
// dedicated field-initializer and pattern reductions.

// rsBalanceState 承载扫描一个构造区间时的括号深度与花括号体边界。
type rsBalanceState struct {
	paren, bracket, brace int
	bodyOpen, bodyClose   int
}

// unitBounds scans one logical construct and its braced body. Depth-zero
// semicolons, top-level commas and a balanced closing brace end the unit.
func (p *rsParser) unitBounds(i, hi int) (end, bodyOpen, bodyClose int) {
	st := &rsBalanceState{bodyOpen: -1, bodyClose: -1}
	for j := i; j < hi; j++ {
		tok := p.tokens[j]
		if tok.Kind == frontkit.TokenEOF {
			return j, st.bodyOpen, st.bodyClose
		}
		if term, next := p.rsBalanceOrTerminate(tok, j, i, hi, st); term {
			return next, st.bodyOpen, st.bodyClose
		}
	}
	return hi, st.bodyOpen, st.bodyClose
}

// rsBalanceOrTerminate updates bracket depths and body bounds. It returns true
// when tok terminates the construct (a closing brace, semicolon or comma at
// zero depth) together with the offset to stop at.
func (p *rsParser) rsBalanceOrTerminate(tok frontkit.Token, j, i, hi int,
	st *rsBalanceState) (bool, int) {
	switch tok.Text {
	case "(":
		st.paren++
	case ")":
		if st.paren > 0 {
			st.paren--
		}
	case "[":
		st.bracket++
	case "]":
		if st.bracket > 0 {
			st.bracket--
		}
	case "{":
		if st.brace == 0 && st.bodyOpen < 0 {
			st.bodyOpen = j
		}
		st.brace++
	case "}":
		if st.brace > 0 {
			st.brace--
		}
		if st.brace == 0 && st.bodyOpen >= 0 {
			st.bodyClose = j
			next := j + 1
			if next < hi && p.tokens[next].Text == ";" {
				next++
			}
			return true, next
		}
	case ";":
		if st.paren == 0 && st.bracket == 0 && st.brace == 0 {
			return true, j + 1
		}
	case ",":
		if st.paren == 0 && st.bracket == 0 && st.brace == 0 && j > i {
			return true, j + 1
		}
	}
	return false, 0
}

// blockVariantAtBody reports the dedicated block-variant kind when the token
// directly before a braced body is the variant keyword the parser just
// reduced in a statement context.
func (p *rsParser) blockVariantAtBody(bodyOpen int,
	mapping rsMapping) (air.NodeKind, string, bool) {
	if mapping.kind != air.KindVarDecl && mapping.kind != air.KindExprStmt &&
		mapping.kind != air.KindConstDecl {
		return air.KindInvalid, "", false
	}
	if bodyOpen <= 0 || bodyOpen >= len(p.tokens) {
		return air.KindInvalid, "", false
	}
	tok := p.tokens[bodyOpen-1]
	if tok.Kind != frontkit.TokenKeyword && tok.Text != "gen" {
		return air.KindInvalid, "", false
	}
	for _, bk := range rsBlockLeadKinds {
		if tok.Text == bk.text {
			return bk.kind, bk.note, true
		}
	}
	return air.KindInvalid, "", false
}

// bodyCtx selects the parse context of one braced body. A braced literal
// whose lead token is an identifier and whose body declares fields reduces as
// a struct literal body in statement contexts, so its fields get the
// dedicated initializer reductions.
func (p *rsParser) bodyCtx(mapping rsMapping, bodyOpen,
	bodyClose int) rsContext {
	if mapping.kind == air.KindVarDecl || mapping.kind == air.KindExprStmt ||
		mapping.kind == air.KindConstDecl {
		if bodyOpen > 0 &&
			p.tokens[bodyOpen-1].Kind == frontkit.TokenIdent &&
			bodyClose > bodyOpen &&
			(rsHasText(p.tokens, bodyOpen+1, bodyClose, ":") ||
				rsHasDepth0Comma(p.tokens, bodyOpen+1, bodyClose)) {
			return rsCtxStructLit
		}
	}
	return p.bodyContext(mapping)
}

// rsHasDepth0Comma reports whether the token interval holds a comma outside
// every bracket pair. A braced body of a statement whose lead is an identifier
// never separates plain statements by top-level commas, so such a comma
// identifies a struct-literal body with shorthand fields.
func rsHasDepth0Comma(tokens []frontkit.Token, lo, hi int) bool {
	depth := 0
	for j := lo; j < hi && j < len(tokens); j++ {
		switch tokens[j].Text {
		case "(", "[", "{":
			depth++
		case ")", "]", "}":
			if depth > 0 {
				depth--
			}
		case ",":
			if depth == 0 {
				return true
			}
		}
	}
	return false
}

// bodyContext selects the rule set used inside one braced body.
func (*rsParser) bodyContext(mapping rsMapping) rsContext {
	switch mapping.kind {
	case air.KindMatch:
		return rsCtxMatchArm
	case air.KindTraitDecl:
		return rsCtxTrait
	case air.KindCompositeLit:
		if mapping.noteName == "struct-literal" {
			return rsCtxStructLit
		}
	case air.KindBlock:
		if mapping.noteName == "impl" {
			return rsCtxImpl
		}
	}
	return rsCtxBlock
}
