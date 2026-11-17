/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the construct markers csParser emits while it reduces a
// declaration or statement unit. Each marker is a dedicated AIR node produced
// at the moment the parser recognizes the construct in the unit it is
// reducing, anchored on that construct's real token position. Zero-width
// anchors keep every lexical projection byte-exact (see csMap.dedicated).

// csAccessorWords are the accessor keywords a property/indexer body declares.
var csAccessorWords = []string{"get", "set", "init", "add", "remove"}

// safeAnchor picks the first candidate token whose offset no already reduced
// sibling strictly contains. A zero-width marker at such a position always
// sorts before (or alongside) the sibling at the same offset, so
// View.OrderChildren keeps both the marker and the sibling projection.
func (*csParser) safeAnchor(children []air.Node,
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
func (p *csParser) dedicatedMarkers(mapping csMapping, i, segEnd,
	bodyOpen, linqAt int, scope csScope, typeName string,
	children []air.Node) []air.Node {
	out := p.declarationMarkers(mapping, i, segEnd, bodyOpen, scope,
		typeName, children)
	if linqAt >= i && linqAt < segEnd {
		// The parser recognized a query_expression while reducing this unit.
		tok := p.safeAnchor(children, p.tokens[linqAt], p.tokens[i])
		out = append(out, p.mapper.dedicated(air.KindCsQuery, tok,
			tok.Text, "cs-query"))
	}
	out = append(out, p.memberMarkers(mapping, i, segEnd, children)...)
	out = append(out, p.modifierMarkers(i, segEnd, children)...)
	return out
}

// markerSink accumulates the markers of one unit and anchors each new marker
// through safeAnchor so it can never be dropped by sibling containment.
type markerSink struct {
	p        *csParser
	i        int
	children []air.Node
	out      []air.Node
}

// push anchors one dedicated marker at anchor (falling back to the unit's
// lead token) and appends it to the sink.
func (s *markerSink) push(kind air.NodeKind, anchor frontkit.Token,
	note string) {
	tok := s.p.safeAnchor(s.children, anchor, s.p.tokens[s.i])
	s.out = append(s.out, s.p.mapper.dedicated(kind, tok, tok.Text, note))
}

// declarationMarkers emits the declaration-form markers of one unit: yield,
// record, event, extern-alias, destructor, operator, constructor and local
// function, each recognized from the mapping the classifier just returned.
func (p *csParser) declarationMarkers(mapping csMapping, i, segEnd,
	bodyOpen int, scope csScope, typeName string,
	children []air.Node) []air.Node {
	tokens := p.tokens
	at := csLeadIndex(tokens, i, segEnd)
	lead := ""
	if at < segEnd && at < len(tokens) {
		lead = tokens[at].Text
	}
	s := &markerSink{p: p, i: i, children: children,
		out: make([]air.Node, 0, 2)}
	switch {
	case mapping.note == "yield":
		s.push(air.KindCsYield, tokens[at], "cs-yield")
	case mapping.note == "record":
		s.push(air.KindCsRecord, tokens[at], "cs-record")
	case mapping.note == "event":
		s.push(air.KindCsEventDecl, tokens[at], "cs-event")
	case csHasText(tokens, i, segEnd, "extern") &&
		csHasText(tokens, i, segEnd, "alias"):
		if extAt := csIndexOfText(tokens, i, segEnd, "extern"); extAt >= 0 {
			s.push(air.KindCsExternAlias, tokens[extAt], "cs-extern-alias")
		}
	case mapping.kind == air.KindMethodDecl && lead == "~":
		s.push(air.KindCsDestructor, tokens[at], "cs-destructor")
	case (mapping.kind == air.KindMethodDecl || mapping.kind == air.KindCall) &&
		csHasText(tokens, i, segEnd, "operator"):
		// operator 与 conversion operator 声明都带 operator 关键字，Pratt 把
		// 它们归约为调用或方法区间，此处按声明形态补发 dedicated 节点。
		if opAt := csIndexOfText(tokens, i, segEnd, "operator"); opAt >= 0 {
			s.push(air.KindCsOperatorDecl, tokens[opAt], "cs-operator")
		}
	case mapping.kind == air.KindMethodDecl && scope == csScopeType &&
		bodyOpen >= 0 && lead != "" && lead == typeName:
		// The parser has reduced a member whose lead identifier equals the
		// enclosing type and that carries a parameter list and body: a
		// genuine constructor_declaration.
		if open := csIndexOfText(tokens, i, segEnd, "("); open >= 0 {
			s.push(air.KindCsConstructor, tokens[open], "cs-constructor")
		}
	case mapping.kind == air.KindMethodDecl && scope == csScopeBody:
		if open := csIndexOfText(tokens, i, segEnd, "("); open >= 0 {
			s.push(air.KindCsLocalFunc, tokens[open], "cs-local-func")
		}
	}
	return s.out
}

// memberMarkers emits the member-form markers of one unit: the property name
// of a property declaration and the this keyword of an indexer declaration.
func (p *csParser) memberMarkers(mapping csMapping, i, segEnd int,
	children []air.Node) []air.Node {
	tokens := p.tokens
	s := &markerSink{p: p, i: i, children: children,
		out: make([]air.Node, 0, 2)}
	switch {
	case mapping.note == "property" && mapping.kind == air.KindFuncDecl:
		if nameAt := csNameIndex(tokens, i, segEnd); nameAt >= 0 {
			s.push(air.KindCsProperty, tokens[nameAt], "cs-property")
		}
	case mapping.note == "indexer" && mapping.kind == air.KindIndex:
		if thisAt := csIndexOfText(tokens, i, segEnd, "this"); thisAt >= 0 {
			s.push(air.KindCsIndexer, tokens[thisAt], "cs-indexer")
		}
	}
	return s.out
}

// modifierMarkers emits one marker per visibility or declaration modifier
// keyword the unit carries.
func (p *csParser) modifierMarkers(i, segEnd int,
	children []air.Node) []air.Node {
	tokens := p.tokens
	s := &markerSink{p: p, i: i, children: children,
		out: make([]air.Node, 0, 2)}
	for j := i; j < segEnd && j < len(tokens); j++ {
		if csContains(csModifierWords, tokens[j].Text) &&
			tokens[j].Kind == frontkit.TokenKeyword {
			s.push(air.KindCsModifier, tokens[j], "cs-modifier")
		}
	}
	return s.out
}

// accessorMarkers emits one marker per accessor keyword the parser recognizes
// inside a property or indexer body it is reducing. The markers attach to the
// body block, whose span is the only sibling family containing them.
func (p *csParser) accessorMarkers(bodyOpen, bodyHi int) []air.Node {
	if bodyOpen < 0 || bodyHi <= bodyOpen || bodyHi > len(p.tokens) {
		return nil
	}
	out := make([]air.Node, 0, 2)
	for j := bodyOpen + 1; j < bodyHi; j++ {
		tok := p.tokens[j]
		if csContains(csAccessorWords, tok.Text) &&
			(tok.Kind == frontkit.TokenKeyword || tok.Kind == frontkit.TokenIdent) {
			out = append(out, p.mapper.dedicated(air.KindCsAccessor, tok,
				tok.Text, "cs-accessor"))
		}
	}
	return out
}

// csIndexOfText returns the first index of a token with the exact text.
func csIndexOfText(tokens []frontkit.Token, lo, hi int, text string) int {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == text {
			return j
		}
	}
	return -1
}

// csNameIndex returns the index of the construct name token (see csNameOf).
func csNameIndex(tokens []frontkit.Token, lo, hi int) int {
	for j := lo; j < hi && j < len(tokens); j++ {
		tok := tokens[j]
		if tok.Kind != frontkit.TokenIdent {
			continue
		}
		if csContains(csLinqKeywords, tok.Text) || csContains(csTypeWords, tok.Text) {
			continue
		}
		return j
	}
	return -1
}
