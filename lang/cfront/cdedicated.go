/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the genuine dedicated-kind reductions of cfront. Every
// marker below is emitted for a construct the parser actually recognized, and
// is anchored on a real token (or the unit lead) of that construct. A marker
// is a zero-width node: it never rewrites an existing projection, so lexical
// fidelity and every existing AIR-ID stay byte-exact. This mirrors the
// dedicated-marker pattern already proven in jsfront/csfront/rsfront/ktfront.

// dedicated anchors a zero-width dedicated-kind marker and tags it with the
// tree-sitter construct name that produced it, so the emission is auditable.
func (*cMapper) dedicated(kind air.NodeKind, tok frontkit.Token,
	note string) air.Node {
	pos := air.NewPosition(tok.Line, tok.Column, tok.Offset, tok.Offset)
	node := air.NewNode(kind, pos, tok.Text)
	_ = air.Annotate(node, air.SemNote, note)
	return node
}

// cSig returns the nearest significant token at or after j within [lo, hi),
// skipping comments, newlines and EOF. It returns nil when the interval ends
// first, so a scan can never leak across a unit boundary.
func cSig(tokens []frontkit.Token, j, lo, hi int) *frontkit.Token {
	for j >= lo && j < hi && j < len(tokens) {
		t := tokens[j]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			j++
			continue
		}
		return &t
	}
	return nil
}

// cSigBackIdx returns the nearest significant token at or before j within
// [lo, hi) together with its index, mirroring cSig in the backward direction.
func cSigBackIdx(tokens []frontkit.Token, j, lo, hi int) (*frontkit.Token, int) {
	for j >= lo && j < hi && j >= 0 && j < len(tokens) {
		t := tokens[j]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			j--
			continue
		}
		return &t, j
	}
	return nil, -1
}

// cDerefOperandExcluded lists the token texts after which a '*' is an infix
// operator or a declarator star, never a dereference.
func cDerefOperandExcluded(prev *frontkit.Token) bool {
	if prev == nil {
		return false
	}
	switch prev.Kind {
	case frontkit.TokenIdent, frontkit.TokenNum, frontkit.TokenStr,
		frontkit.TokenChar:
		return true
	}
	switch prev.Text {
	case "]", ")", "*", "/", "++", "--":
		return true
	default:
		return false
	}
}

// cIsLabel reports whether the identifier at j starts a labeled statement:
// it is directly followed by ':' and preceded by a statement boundary
// ('{', '}' or ';') or sits at the very start of the unit. A bitfield width
// (ident ':' after a type name) and a ternary branch (after '?') never match.
func cIsLabel(tokens []frontkit.Token, colonIdx, lo int) bool {
	ident, idx := cSigBackIdx(tokens, colonIdx-1, lo, colonIdx)
	if ident == nil || ident.Kind != frontkit.TokenIdent {
		return false
	}
	prev, _ := cSigBackIdx(tokens, idx-1, lo, idx)
	if prev == nil {
		return true
	}
	return prev.Text == "{" || prev.Text == "}" || prev.Text == ";"
}

// cIsDeref reports whether the '*' at j dereferences its operand: the
// previous significant token must not end an operand (which would make this
// an infix multiply or a declarator star) and the next significant token must
// start one.
func cIsDeref(tokens []frontkit.Token, j, lo, hi int) bool {
	if tokens[j].Text != "*" {
		return false
	}
	prev, _ := cSigBackIdx(tokens, j-1, lo, j)
	if cDerefOperandExcluded(prev) {
		return false
	}
	next := cSig(tokens, j+1, j+1, hi)
	if next == nil {
		return false
	}
	switch next.Kind {
	case frontkit.TokenIdent, frontkit.TokenKeyword:
		return true
	}
	return next.Text == "("
}

// cInitBraceOpen reports whether the '{' at j opens an initializer list: the
// nearest significant token before it is '='.
func cInitBraceOpen(tokens []frontkit.Token, j, lo int) bool {
	if tokens[j].Text != "{" {
		return false
	}
	prev, _ := cSigBackIdx(tokens, j-1, lo, j)
	return prev != nil && prev.Text == "="
}

// cIsInitPair reports whether the '.' at j designates an initializer member:
// it follows an initializer '{' within the same unit and is followed by an
// identifier and '='. A member access assignment (a.b = c) has no
// initializer brace and never matches.
func cIsInitPair(tokens []frontkit.Token, j, lo, hi int) bool {
	if tokens[j].Text != "." {
		return false
	}
	key := cSig(tokens, j+1, j+1, hi)
	if key == nil || key.Kind != frontkit.TokenIdent {
		return false
	}
	assign := cSig(tokens, j+2, j+2, hi)
	if assign == nil || assign.Text != "=" {
		return false
	}
	for k := lo; k < j; k++ {
		if cInitBraceOpen(tokens, k, lo) {
			return true
		}
	}
	return false
}

// cDedicatedMarkers returns the dedicated markers of one reduced unit: the
// record/enum-member/union lead markers plus the head-segment scans for
// labels, dereferences and designated initializer pairs. Scans stay inside
// [i, segEnd) so a nested unit's constructs are claimed only by that unit.
func (p *cParser) cDedicatedMarkers(i, end, segEnd int, record cBodyKind,
	children []air.Node) []air.Node {
	tokens := p.tokens
	if segEnd > end {
		segEnd = end
	}
	out := make([]air.Node, 0, 2)
	push := func(kind air.NodeKind, at int, note string) {
		tok := cSafeAnchor(children, tokens[at], tokens[i])
		out = append(out, p.mapper.dedicated(kind, tok, note))
	}
	leadIdx := -1
	for j := i; j < segEnd; j++ {
		if tokens[j].Kind == frontkit.TokenNewline ||
			tokens[j].Kind == frontkit.TokenComment ||
			tokens[j].Kind == frontkit.TokenEOF {
			continue
		}
		leadIdx = j
		break
	}
	if leadIdx >= 0 {
		lead := tokens[leadIdx].Text
		switch {
		case lead == "union" && record == cBodyStatement:
			push(air.KindCUnion, leadIdx, "c-union")
		case record == cBodyEnum:
			if tokens[leadIdx].Kind == frontkit.TokenIdent {
				push(air.KindCEnumMember, leadIdx, "c-enum-member")
			}
		case record == cBodyRecord:
			push(air.KindCFieldDecl, leadIdx, "c-field-decl")
		}
	}
	for j := i; j < segEnd && j < len(tokens); j++ {
		tok := tokens[j]
		switch {
		case tok.Text == "defined":
			push(air.KindCDefined, j, "c-defined")
		case tok.Text == ":" && cIsLabel(tokens, j, i):
			push(air.KindCLabel, j-1, "c-label")
		case tok.Text == "*" && cIsDeref(tokens, j, i, segEnd):
			push(air.KindCDeref, j, "c-deref")
		}
	}
	for j := i; j < end && j < len(tokens); j++ {
		if tokens[j].Text == "." && cIsInitPair(tokens, j, i, end) {
			push(air.KindCInitPair, j, "c-init-pair")
		}
	}
	return out
}

// cSafeAnchor picks the first candidate token whose offset no already reduced
// sibling strictly contains, mirroring rsfront.safeAnchor. A marker anchored
// inside a sibling's span would be pruned by View.OrderChildren's overlap
// rule, so the unit lead token is the safe fallback.
func cSafeAnchor(children []air.Node,
	candidates ...frontkit.Token) frontkit.Token {
	for _, tok := range candidates {
		at := tok.Offset
		contained := false
		for _, child := range children {
			pos := child.Pos()
			if pos.Offset < at && at < pos.EndOffset {
				contained = true
				break
			}
		}
		if !contained {
			return tok
		}
	}
	if len(candidates) == 0 {
		return frontkit.Token{}
	}
	return candidates[len(candidates)-1]
}

// definedMarkers scans one preprocessor directive interval for `defined`
// macro tests and anchors one dedicated marker per occurrence. Directives are
// reduced by cpreproc outside parseUnit, so this hook keeps the emission
// inside the same audit trail.
func (p *cParser) definedMarkers(lo, hi int) []air.Node {
	out := make([]air.Node, 0, 1)
	for j := lo; j < hi && j < len(p.tokens); j++ {
		if p.tokens[j].Text == "defined" {
			out = append(out, p.mapper.dedicated(
				air.KindCDefined, cSafeAnchor(out, p.tokens[lo],
					p.tokens[j]), "c-defined"))
		}
	}
	return out
}
