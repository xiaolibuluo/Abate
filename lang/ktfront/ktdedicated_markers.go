/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the construct markers ktParser emits while it reduces a
// declaration or statement unit. Each marker is a dedicated AIR node produced
// at the moment the parser recognizes the construct in the unit it is
// reducing, anchored on that construct's real token position. Zero-width
// anchors keep every lexical projection byte-exact (see ktMap.dedicated).

// safeAnchor picks the first candidate token whose offset no already reduced
// sibling strictly contains, mirroring csfront.safeAnchor.
func (*ktParser) safeAnchor(children []air.Node,
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

// dedicatedMarkers returns the construct markers for one reduced unit. The
// result may be empty; callers append it to the unit children.
func (p *ktParser) dedicatedMarkers(mapping ktMapping, i, segEnd, end,
	bodyOpen, _ int, children []air.Node) []air.Node {
	tokens := p.tokens
	at := ktLeadIndex(tokens, i, segEnd)
	lead := ""
	if at < segEnd && at < len(tokens) {
		lead = tokens[at].Text
	}
	out := make([]air.Node, 0, 2)
	push := func(kind air.NodeKind, anchor frontkit.Token, note string) {
		tok := p.safeAnchor(children, anchor, tokens[i])
		out = append(out, p.mapper.dedicated(kind, tok, tok.Text, note))
	}
	switch {
	case mapping.note == "val-binding" || mapping.note == "var-binding":
		push(air.KindKtProperty, tokens[at], "kt-property")
	case mapping.note == "constructor":
		push(air.KindKtSecondaryCtor, tokens[at], "kt-secondary-ctor")
		if at := ktCtorDelegationAt(tokens, i, segEnd); at >= 0 {
			push(air.KindKtCtorDelegation, tokens[at], "kt-ctor-delegation")
		}
	case mapping.note == "getter" || mapping.note == "setter":
		push(air.KindKtAccessor, tokens[at], "kt-accessor")
	case mapping.note == "companion-object":
		if at := ktIndexOfText(tokens, i, segEnd, "companion"); at >= 0 {
			push(air.KindKtCompanion, tokens[at], "kt-companion")
		}
	case lead == "import":
		if at := ktIndexOfText(tokens, i, segEnd, "as"); at >= 0 {
			push(air.KindKtImportAlias, tokens[at], "kt-import-alias")
		} else if at := ktIndexOfText(tokens, i, segEnd, "."); at >= 0 {
			push(air.KindKtImportList, tokens[at], "kt-import-list")
		}
	}
	switch mapping.kind {
	case air.KindIf, air.KindWhile, air.KindDoWhile, air.KindRangeFor:
		// The parser reduced a control structure and its braced body: emit
		// the dedicated control_structure_body node anchored on the brace.
		if bodyOpen >= i && bodyOpen < end {
			push(air.KindKtControlBody, tokens[bodyOpen], "kt-control-body")
		}
	}
	// Accessors written with parenthesized form inside a property unit: the
	// parser recognizes get(/set( adjacency while reducing the property.
	for j := i; j < segEnd && j < len(tokens); j++ {
		tok := tokens[j]
		if (tok.Text == "get" || tok.Text == "set") &&
			j+1 < segEnd && tokens[j+1].Text == "(" {
			push(air.KindKtAccessor, tok, "kt-accessor")
		}
	}
	return out
}

// ktCtorDelegationAt returns the index of the this/super token of a
// constructor delegation call, or -1 when the unit has none.
func ktCtorDelegationAt(tokens []frontkit.Token, lo, hi int) int {
	for j := lo; j < hi-1 && j < len(tokens); j++ {
		if tokens[j].Text == ":" &&
			(tokens[j+1].Text == "this" || tokens[j+1].Text == "super") {
			return j + 1
		}
	}
	return -1
}

// ktIndexOfText returns the first index of a token with the exact text.
func ktIndexOfText(tokens []frontkit.Token, lo, hi int, text string) int {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == text {
			return j
		}
	}
	return -1
}
