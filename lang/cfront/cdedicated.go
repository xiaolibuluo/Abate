/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"slices"

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
		case tok.Text == "->" || tok.Text == ".":
			if at, ok := cFieldIdentAt(tokens, j, i, segEnd); ok {
				push(air.KindCFieldIdent, at, "c-field-ident")
			}
		}
	}
	for j := i; j < end && j < len(tokens); j++ {
		if tokens[j].Text == "." && cIsInitPair(tokens, j, i, end) {
			push(air.KindCInitPair, j, "c-init-pair")
		}
	}
	return out
}

// cSizedTypeWords lists the lead modifiers of a sized_type_specifier.
var cSizedTypeWords = []string{"unsigned", "signed", "long", "short"}

// cIsSizedTypeWord reports whether the spelling leads a sized type specifier.
func cIsSizedTypeWord(s string) bool { return slices.Contains(cSizedTypeWords, s) }

// cFuncDeclaratorName locates the name token of a function declarator: the
// identifier immediately preceding the first '(' of the unit head, which is
// the declarator anchor of a tree-sitter function_declarator. A leading
// parenthesised declarator (function pointer) has no such identifier and is
// left unmarked.
func cFuncDeclaratorName(tokens []frontkit.Token, lo, hi int) (int, bool) {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text != "(" {
			continue
		}
		prev, idx := cSigBackIdx(tokens, j-1, lo, j)
		if prev == nil || prev.Kind != frontkit.TokenIdent {
			return -1, false
		}
		return idx, true
	}
	return -1, false
}

// cInitDeclaratorEq reports whether the head of a declaration unit holds a
// top-level '=' of an init_declarator and returns its index. '=' tokens
// nested inside parentheses, subscripts or braces belong to child constructs
// (for-initializers, array initializers) and are excluded by depth tracking.
func cInitDeclaratorEq(tokens []frontkit.Token, lo, hi int) (int, bool) {
	paren, bracket, brace := 0, 0, 0
	for j := lo; j < hi && j < len(tokens); j++ {
		switch tokens[j].Text {
		case "(":
			paren++
		case ")":
			if paren > 0 {
				paren--
			}
		case "[":
			bracket++
		case "]":
			if bracket > 0 {
				bracket--
			}
		case "{":
			brace++
		case "}":
			if brace > 0 {
				brace--
			}
		case "=":
			if paren == 0 && bracket == 0 && brace == 0 {
				return j, true
			}
		}
	}
	return -1, false
}

// cMacroTypeAt reports whether the user identifier at leadIdx opens a
// macro_type_specifier: it is directly followed by a balanced parenthesis
// group, and declarator evidence (an identifier or '*') remains past the
// group — LIST(int) items. Without a trailing declarator the group is a
// parameter list and the construct is an ordinary function declarator.
func cMacroTypeAt(tokens []frontkit.Token, leadIdx, hi int) bool {
	parenIdx := -1
	for j := leadIdx + 1; j < hi && j < len(tokens); j++ {
		t := tokens[j]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			continue
		}
		if t.Text == "(" {
			parenIdx = j
		}
		break
	}
	if parenIdx < 0 {
		return false
	}
	closeIdx := cBalancedEnd(tokens, parenIdx, hi, "(", ")")
	if closeIdx <= parenIdx {
		return false
	}
	after := cSig(tokens, closeIdx, closeIdx, hi)
	if after == nil {
		return false
	}
	return after.Kind == frontkit.TokenIdent || after.Text == "*"
}

// cDeclTypeLead returns the position of the type lead word of a declaration
// unit together with the dedicated kind it deserves: a sizing modifier
// (unsigned/signed/long/short) maps to sized_type_specifier; a user
// identifier maps to macro_type_specifier when it directly opens a
// parenthesised type argument, and to type_identifier otherwise. The lead
// must be followed by declarator evidence (an identifier or '*'), otherwise
// the unit is an expression that merely fell onto a declaration mapping and
// carries no type position at all.
func cDeclTypeLead(tokens []frontkit.Token, lo, hi int) (int, air.NodeKind, bool) {
	start := cAttributeExtent(tokens, lo, hi)
	leadIdx := -1
	for j := start; j < hi && j < len(tokens); j++ {
		text := tokens[j].Text
		if text == "" {
			continue
		}
		ch := text[0]
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' {
			leadIdx = j
			break
		}
	}
	if leadIdx < 0 {
		return -1, 0, false
	}
	lead := tokens[leadIdx]
	if cIsSizedTypeWord(lead.Text) {
		return leadIdx, air.KindCSizedType, true
	}
	if lead.Kind != frontkit.TokenIdent || cIsTypeWord(lead.Text) ||
		cIsSpecifier(lead.Text) {
		return -1, 0, false
	}
	if cMacroTypeAt(tokens, leadIdx, hi) {
		return leadIdx, air.KindCMacroType, true
	}
	next := cSig(tokens, leadIdx+1, leadIdx+1, hi)
	if next == nil {
		return -1, 0, false
	}
	if next.Kind == frontkit.TokenIdent || next.Text == "*" {
		return leadIdx, air.KindCTypeIdent, true
	}
	return -1, 0, false
}

// cFieldIdentAt reports whether the member-access token at j ('->' or '.')
// is followed by a member identifier and returns its index. '->' always
// leads a field_identifier; '.' counts only in member-access position, i.e.
// its previous significant token closes an operand (identifier, ')' or ']'),
// which excludes designated initializers like .x = 1.
func cFieldIdentAt(tokens []frontkit.Token, j, lo, hi int) (int, bool) {
	switch tokens[j].Text {
	case "->":
	case ".":
		prev, _ := cSigBackIdx(tokens, j-1, lo, j)
		if prev == nil {
			return -1, false
		}
		if prev.Kind != frontkit.TokenIdent && prev.Text != ")" && prev.Text != "]" {
			return -1, false
		}
	default:
		return -1, false
	}
	for k := j + 1; k < hi && k < len(tokens); k++ {
		t := tokens[k]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			continue
		}
		if t.Kind == frontkit.TokenIdent {
			return k, true
		}
		return -1, false
	}
	return -1, false
}

// cDeclaratorMarkers returns the dedicated markers that depend on the whole
// unit mapping: function_declarator anchors the function name,
// init_declarator anchors the top-level '=', and sized_type_specifier /
// macro_type_specifier / type_identifier anchor the type lead word. All
// markers are zero-width nodes pinned on real head tokens; they never rewrite
// an existing projection, so lexical fidelity and AIR-IDs stay byte-exact.
func (p *cParser) cDeclaratorMarkers(mapping cMapping, i, end, segEnd int,
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
	pushType := func() air.NodeKind {
		at, kind, ok := cDeclTypeLead(tokens, i, segEnd)
		if !ok {
			return 0
		}
		note := "c-sized-type"
		switch kind {
		case air.KindCMacroType:
			note = "c-macro-type"
		case air.KindCTypeIdent:
			note = "c-type-ident"
		}
		push(kind, at, note)
		return kind
	}
	switch mapping.kind {
	case air.KindFuncDecl, air.KindMethodDecl:
		// A leading macro type specifier owns the parenthesis group, so the
		// unit carries no function declarator at all.
		if pushType() != air.KindCMacroType {
			if at, ok := cFuncDeclaratorName(tokens, i, segEnd); ok {
				push(air.KindCFuncDeclarator, at, "c-func-declarator")
			}
		}
	case air.KindVarDecl, air.KindConstDecl:
		if at, ok := cInitDeclaratorEq(tokens, i, segEnd); ok {
			push(air.KindCInitDeclarator, at, "c-init-declarator")
		}
		pushType()
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
