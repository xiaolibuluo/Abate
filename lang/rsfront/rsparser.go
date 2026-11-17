/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

// This file holds the whole-file reduction: item, impl, trait, match, pattern
// arm and attribute productions. Expression reduction is delegated to rsExpr.

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// rsConfig is the complete injectable data set for one Rust reducer.
type rsConfig struct {
	infix      []rsBinding
	notes      []rsNote
	itemRules  []rsItemRule
	macroNames []string
	maxDepth   int
}

// defaultRsConfig returns independent copies of the frozen Rust tables.
func defaultRsConfig() rsConfig {
	return rsConfig{
		infix:      append([]rsBinding(nil), rsInfixBindings...),
		notes:      append([]rsNote(nil), rsNotes...),
		itemRules:  append([]rsItemRule(nil), rsItemRules...),
		macroNames: append([]string(nil), rsMacroNames...),
		maxDepth:   rsMaxDepth,
	}
}

// rsParser reduces one file's token stream into a layered Rust AIR tree.
type rsParser struct {
	tokens []frontkit.Token
	path   string
	view   *frontkit.View
	mapper *rsMapper
	expr   *rsExpr
	sink   *frontkit.DiagnosticSink
	cfg    rsConfig
}

// newRsParser builds the reducer for one file with an injected configuration.
func newRsParser(tokens []frontkit.Token, path string, cfg rsConfig) *rsParser {
	view := frontkit.NewView(frontkit.SourceFromTokens(tokens), path)
	mapper := newRsMapper(view, cfg)
	return &rsParser{
		tokens: tokens, path: path, view: view, mapper: mapper,
		expr: newRsExpr(mapper, cfg),
		sink: frontkit.NewDiagnosticSink(path, "rs"), cfg: cfg,
	}
}

// parse reduces the whole file and reports diagnostics without ever panicking.
func (p *rsParser) parse() (air.Node, []air.Diagnostic, error) {
	for i := range p.tokens {
		if p.tokens[i].Kind == frontkit.TokenIllegal {
			_ = p.sink.Warning(2, p.tokens[i].Offset, p.tokens[i].Line,
				p.tokens[i].Column,
				"illegal byte sequence skipped during recovery")
		}
	}
	return p.parseFile(), p.sink.Diagnostics(), nil
}

// parseFile builds the KindPackage root covering [0, fileSize).
func (p *rsParser) parseFile() air.Node {
	size := p.view.Size()
	children := p.parseSeq(0, len(p.tokens), rsCtxTop)
	start := air.NewPosition(1, 1, 0, size)
	end := air.NewPosition(1, 1, size, size)
	root := air.NewNodeAt(air.KindPackage, start, end, p.path,
		p.view.OrderChildren(children)...)
	_ = air.Annotate(root, air.SemScopeID, "crate")
	return root
}

// parseSeq reduces one token interval into a sequence of sibling nodes.
func (p *rsParser) parseSeq(lo, hi int, ctx rsContext) []air.Node {
	if lo < 0 {
		lo = 0
	}
	if hi > len(p.tokens) {
		hi = len(p.tokens)
	}
	out := make([]air.Node, 0, 8)
	for i := lo; i < hi; {
		tok := p.tokens[i]
		if tok.Kind == frontkit.TokenEOF {
			break
		}
		if rsIsTrivia(tok.Kind) {
			i++
			continue
		}
		if tok.Kind == frontkit.TokenComment {
			out = append(out, p.mapper.comment(tok))
			i++
			continue
		}
		if p.isAttributeStart(i, hi) {
			end := p.attributeEnd(i, hi)
			out = append(out, p.mapper.attribute(p.tokens, i, end))
			i = end
			continue
		}
		end, bodyOpen, bodyClose := p.unitBounds(i, hi)
		if end <= i {
			end = i + 1
		}
		out = append(out, p.parseUnit(i, end, bodyOpen, bodyClose, ctx))
		i = end
	}
	return p.view.OrderChildren(out)
}

// isAttributeStart reports whether one token opens an outer or inner attribute.
func (p *rsParser) isAttributeStart(i, hi int) bool {
	if p.tokens[i].Text != "#" {
		return false
	}
	if i+1 < hi && p.tokens[i+1].Text == "[" {
		return true
	}
	return i+2 < hi && p.tokens[i+1].Text == "!" && p.tokens[i+2].Text == "["
}

// attributeEnd finds the exclusive end of one #[...] attribute by balancing
// all three bracket families; an unclosed attribute consumes the interval.
func (p *rsParser) attributeEnd(i, hi int) int {
	paren, bracket, brace := 0, 0, 0
	for j := i; j < hi; j++ {
		switch p.tokens[j].Text {
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
			if bracket == 0 && paren == 0 && brace == 0 {
				return j + 1
			}
		case "{":
			brace++
		case "}":
			if brace > 0 {
				brace--
			}
		}
	}
	return hi
}

// parseUnit reduces one construct interval into a layered AIR node whose span
// covers every consumed token, so interval coverage always holds.
func (p *rsParser) parseUnit(i, end, bodyOpen, bodyClose int,
	ctx rsContext) air.Node {
	if end > len(p.tokens) {
		end = len(p.tokens)
	}
	if end <= i {
		end = i + 1
	}
	mapping := p.mapper.match(p.tokens, i, end, bodyOpen, bodyClose, ctx)
	segEnd := end
	if bodyOpen > i && bodyOpen < end {
		segEnd = bodyOpen
	}
	children := p.expr.reduce(p.tokens, i, segEnd)
	// 体块先于标记归约，让 safeAnchor 能看见体块子节点的 span：自然锚点
	// 落在体块复合节点内部时回退到单元首 token，避免被 OrderChildren 丢弃。
	var bodyChildren []air.Node
	hasBody := bodyOpen >= i && bodyClose >= bodyOpen && bodyClose < end
	if hasBody {
		bodyChildren = p.parseSeq(bodyOpen+1, bodyClose,
			p.bodyCtx(mapping, bodyOpen, bodyClose))
	}
	anchors := children
	if len(bodyChildren) > 0 {
		anchors = make([]air.Node, 0, len(children)+len(bodyChildren))
		anchors = append(anchors, children...)
		anchors = append(anchors, bodyChildren...)
	}
	marks := p.dedicatedMarkers(mapping, i, segEnd, end,
		bodyOpen, bodyClose, ctx, anchors)
	if hasBody {
		// 落在花括号体内区间上的标记挂到体块节点，避免被体块这一同级
		// 复合节点覆盖而丢弃；体外的标记保持为单元子节点。边界比较使用
		// 体块首尾 token 的字节偏移，与标记锚点的字节偏移同口径。
		bodyFrom := p.tokens[bodyOpen].Offset
		bodyTo := p.tokens[bodyClose].Offset
		for _, mk := range marks {
			at := mk.Pos().Offset
			if at > bodyFrom && at < bodyTo {
				bodyChildren = append(bodyChildren, mk)
			} else {
				children = append(children, mk)
			}
		}
		block := p.mapper.build(air.KindBlock, p.tokens[bodyOpen],
			p.tokens[bodyClose], "", bodyChildren...)
		children = append(children, block)
	} else {
		children = append(children, marks...)
	}
	node := p.mapper.build(mapping.kind, p.tokens[i], p.tokens[end-1],
		mapping.name, children...)
	p.mapper.annotateNode(node, mapping)
	return node
}
