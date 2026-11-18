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

// This file holds the construct markers rsParser emits while it reduces an
// item, block or pattern interval. Each marker is a dedicated AIR node
// produced at the moment the parser recognizes the construct in the unit it
// is reducing, anchored on that construct's real token position. Zero-width
// anchors keep every lexical projection byte-exact (see rsMapper.dedicated).

// rsBlockLeadKinds maps the block-variant lead keywords onto their dedicated
// kinds. The parser emits the marker when the lead keyword directly opens the
// unit's braced body.
var rsBlockLeadKinds = []struct {
	text string
	kind air.NodeKind
	note string
}{
	{"async", air.KindRsAsyncBlock, "rs-async-block"},
	{"const", air.KindRsConstBlock, "rs-const-block"},
	{"gen", air.KindRsGenBlock, "rs-gen-block"},
	{"try", air.KindRsTryBlock, "rs-try-block"},
	{"unsafe", air.KindRsUnsafeBlock, "rs-unsafe-block"},
}

// safeAnchor picks the first candidate token whose offset no already reduced
// sibling strictly contains, mirroring csfront.safeAnchor.
func (*rsParser) safeAnchor(children []air.Node,
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
func (p *rsParser) dedicatedMarkers(mapping rsMapping, i, segEnd, end,
	bodyOpen, _ int, ctx rsContext, children []air.Node) []air.Node {
	tokens := p.tokens
	leadIdx := rsEffectiveLeadIndex(tokens, i, segEnd)
	if leadIdx < 0 {
		return nil
	}
	lead := tokens[leadIdx].Text
	out := make([]air.Node, 0, 2)
	push := func(kind air.NodeKind, anchor frontkit.Token, note string) {
		tok := p.safeAnchor(children, anchor, tokens[i])
		out = append(out, p.mapper.dedicated(kind, tok, tok.Text, note))
	}
	switch {
	case i+1 == end && tokens[i].Text == ";":
		push(air.KindRsEmptyStmt, tokens[i], "rs-empty-stmt")
	case lead == "macro_rules":
		push(air.KindRsMacroDef, tokens[leadIdx], "rs-macro-def")
	case lead == "extern" || lead == "crate":
		if crateAt := rsIndexOfText(tokens, i, segEnd, "crate"); crateAt >= 0 {
			push(air.KindRsExternCrate, tokens[crateAt], "rs-extern-crate")
		}
	case lead == "use":
		// use 的子形态（::、{}、*、as）可能位于花括号体内，扫描整个单元。
		out = append(out, p.useMarkers(i, end, children)...)
	case lead == "fn" && ctx == rsCtxTrait && bodyOpen < 0:
		// A trait member reduced to a signature without a body.
		push(air.KindRsFuncSig, tokens[leadIdx], "rs-func-signature")
	case lead == "fn":
		if ltAt := rsGenericFnAt(tokens, leadIdx, segEnd); ltAt >= 0 {
			push(air.KindRsGenericFunc, tokens[ltAt], "rs-generic-func")
		}
	case lead == "loop":
		// Rust's `loop {}` is an unconditional infinite loop, distinct from
		// `while`/`for`. Anchor a dedicated node on the real `loop` token.
		push(air.KindRsLoop, tokens[leadIdx], "rs-loop")
	case lead == "impl":
		// An `impl` item is Rust's inherent/trait implementation block.
		push(air.KindRsImpl, tokens[leadIdx], "rs-impl")
	}
	for _, blockKind := range rsBlockLeadKinds {
		if lead == blockKind.text && bodyOpen == leadIdx+1 {
			push(blockKind.kind, tokens[leadIdx], blockKind.note)
			break
		}
	}
	if kind, note, ok := p.blockVariantAtBody(bodyOpen, mapping); ok {
		push(kind, tokens[bodyOpen-1], note)
	}
	if ctx == rsCtxStructLit {
		out = append(out, p.fieldInitMarkers(i, segEnd, children)...)
	}
	if ctx == rsCtxMatchArm {
		// 模式花括号可能就是体块边界，扫描整个单元区间。
		out = append(out, p.patternShorthandMarkers(i, end, children)...)
	}
	return out
}

// useMarkers emits the dedicated nodes of the use-declaration sub-forms the
// parser recognizes while reducing one use item.
func (p *rsParser) useMarkers(i, segEnd int, children []air.Node) []air.Node {
	tokens := p.tokens
	out := make([]air.Node, 0, 2)
	push := func(kind air.NodeKind, at int, note string) {
		tok := p.safeAnchor(children, tokens[at], tokens[i])
		out = append(out, p.mapper.dedicated(kind, tok, tok.Text, note))
	}
	if at := rsIndexOfText(tokens, i, segEnd, "::"); at >= 0 {
		push(air.KindRsScopedUse, at, "rs-scoped-use")
	}
	if at := rsIndexOfText(tokens, i, segEnd, "{"); at >= 0 {
		push(air.KindRsUseList, at, "rs-use-list")
	}
	if at := rsIndexOfText(tokens, i, segEnd, "*"); at >= 0 {
		push(air.KindRsUseWildcard, at, "rs-use-wildcard")
	}
	if at := rsIndexOfKeyword(tokens, i, segEnd, "as"); at >= 0 {
		push(air.KindRsUseAs, at, "rs-use-as")
	}
	return out
}

// fieldInitMarkers emits the field-initializer nodes of a struct literal body:
// each entry reduced is either `name: value`, `..base` or a shorthand field.
func (p *rsParser) fieldInitMarkers(i, segEnd int,
	children []air.Node) []air.Node {
	tokens := p.tokens
	out := make([]air.Node, 0, 2)
	push := func(kind air.NodeKind, at int, note string) {
		tok := p.safeAnchor(children, tokens[at], tokens[i])
		out = append(out, p.mapper.dedicated(kind, tok, tok.Text, note))
	}
	if at := rsIndexOfText(tokens, i, segEnd, ".."); at >= 0 {
		push(air.KindRsBaseField, at, "rs-base-field")
		return out
	}
	for j := i; j < segEnd && j+1 < len(tokens); j++ {
		if tokens[j].Kind != frontkit.TokenIdent {
			continue
		}
		if tokens[j+1].Text == ":" {
			push(air.KindRsFieldIdent, j, "rs-field-identifier")
			push(air.KindRsFieldInit, j+1, "rs-field-initializer")
			return out
		}
	}
	for j := i; j < segEnd && j < len(tokens); j++ {
		if tokens[j].Kind == frontkit.TokenIdent {
			push(air.KindRsShorthandField, j, "rs-shorthand-field")
			return out
		}
	}
	return out
}

// patternShorthandMarkers emits one node per shorthand field the parser
// recognizes inside a struct pattern of a match arm (`S { x, y }`).
func (p *rsParser) patternShorthandMarkers(i, segEnd int,
	children []air.Node) []air.Node {
	tokens := p.tokens
	out := make([]air.Node, 0, 2)
	push := func(kind air.NodeKind, at int, note string) {
		tok := p.safeAnchor(children, tokens[at], tokens[i])
		out = append(out, p.mapper.dedicated(kind, tok, tok.Text, note))
	}
	arrowAt := rsIndexOfText(tokens, i, segEnd, "=>")
	if arrowAt < 0 {
		arrowAt = segEnd
	}
	braceOpen := -1
	depth := 0
	for j := i; j < arrowAt && j < len(tokens); j++ {
		switch tokens[j].Text {
		case "{":
			if depth == 0 {
				braceOpen = j
			}
			depth++
		case "}":
			depth--
			if depth == 0 && braceOpen >= 0 {
				for k := braceOpen + 1; k < j && k < len(tokens); k++ {
					if tokens[k].Kind == frontkit.TokenIdent &&
						(k+1 >= j || tokens[k+1].Text != ":") {
						push(air.KindRsShorthandIdent, k, "rs-shorthand-ident")
					}
				}
				braceOpen = -1
			}
		}
	}
	return out
}

// rsGenericFnAt returns the index of the generic argument list that directly
// follows a function item's name, or -1.
func rsGenericFnAt(tokens []frontkit.Token, leadIdx, hi int) int {
	for j := leadIdx + 1; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == "(" {
			return -1
		}
		if tokens[j].Text == "<" {
			return j
		}
	}
	return -1
}

// rsIndexOfText returns the first index of a token with the exact text.
func rsIndexOfText(tokens []frontkit.Token, lo, hi int, text string) int {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == text {
			return j
		}
	}
	return -1
}

// rsIndexOfKeyword returns the first index of a keyword token with the text.
func rsIndexOfKeyword(tokens []frontkit.Token, lo, hi int, text string) int {
	for j := lo; j < hi && j < len(tokens); j++ {
		if tokens[j].Text == text && tokens[j].Kind == frontkit.TokenKeyword {
			return j
		}
	}
	return -1
}

// rsEffectiveLeadIndex returns the index of a unit's first non-trivia token
// after skipping visibility modifiers.
func rsEffectiveLeadIndex(tokens []frontkit.Token, lo, hi int) int {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	for j := lo; j < hi; j++ {
		tok := tokens[j]
		if tok.Text == "" || rsIsTrivia(tok.Kind) ||
			tok.Kind == frontkit.TokenComment {
			continue
		}
		if tok.Kind == frontkit.TokenKeyword || tok.Kind == frontkit.TokenIdent {
			if rsContains(rsVisibilityModifiers, tok.Text) {
				if tok.Text == "pub" && j+1 < hi && tokens[j+1].Text == "(" {
					for j+1 < hi && tokens[j+1].Text != ")" {
						j++
					}
					if j+1 < hi && tokens[j+1].Text == ")" {
						j++
					}
				}
				continue
			}
			return j
		}
		if tok.Kind == frontkit.TokenPunct &&
			(tok.Text == "#" || tok.Text == "(" || tok.Text == ")" ||
				tok.Text == "[") {
			continue
		}
		return j
	}
	return -1
}
