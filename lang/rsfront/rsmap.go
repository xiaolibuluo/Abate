/*
Package rsfront reduces Rust sources into AIR.
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
This package comment ends with a period.
*/
package rsfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
	"strings"
)

type rsMapping struct {
	kind     air.NodeKind
	name     string
	noteName string
	semKey   air.SemKey
	semValue string
}

type rsMapper struct {
	view       *frontkit.View
	notes      []rsNote
	itemRules  []rsItemRule
	macroNames []string
}

func newRsMapper(view *frontkit.View, cfg rsConfig) *rsMapper {
	return &rsMapper{
		view: view, notes: cfg.notes, itemRules: cfg.itemRules,
		macroNames: cfg.macroNames,
	}
}

func (m *rsMapper) note(name string) string { return rsNoteFor(m.notes, name) }

func (*rsMapper) annotate(node air.Node, key air.SemKey, value string) air.Node {
	_ = air.Annotate(node, key, value)
	return node
}

func (m *rsMapper) partial(node air.Node) air.Node {
	return m.annotate(node, air.SemPartial, "true")
}

func (m *rsMapper) build(kind air.NodeKind, start, end frontkit.Token,
	text string, children ...air.Node) air.Node {
	finish := end
	if finish.EndOffset < start.Offset {
		finish = start
	}
	pos := air.NewPosition(start.Line, start.Column, start.Offset, finish.EndOffset)
	return air.NewNode(kind, pos, text, m.view.OrderChildren(children)...)
}

func (m *rsMapper) leaf(tok frontkit.Token, kind air.NodeKind) air.Node {
	return m.build(kind, tok, tok, tok.Text)
}

func (m *rsMapper) lexical(tok frontkit.Token) air.Node {
	return m.leaf(tok, rsTokenNodeKind(tok.Kind))
}

// partialLeaf 诚实投影恢复位置上的 token：词法 token 保持冻结投影，标点、
// 运算符等非词法 token 标记为带 SemPartial 的部分节点，绝不按拼写贴专用
// Kind 标签冒充真实语法归约。
func (m *rsMapper) partialLeaf(tok frontkit.Token) air.Node {
	if rsIsLexicalKind(tok.Kind) {
		return m.lexical(tok)
	}
	return m.partial(m.build(air.KindPartialNode, tok, tok, ""))
}

func (m *rsMapper) annotateNode(node air.Node, mapping rsMapping) {
	if node == nil {
		return
	}
	if mapping.semKey.Valid() && mapping.semValue != "" {
		_ = air.Annotate(node, mapping.semKey, mapping.semValue)
	}
	if mapping.noteName != "" {
		_ = air.Annotate(node, air.SemNote, m.note(mapping.noteName))
	}
}

func spanToken(node air.Node, fallback frontkit.Token) frontkit.Token {
	if node == nil {
		return fallback
	}
	pos := node.Pos()
	return frontkit.Token{
		Line: pos.Line, Column: pos.Column,
		Offset: pos.Offset, EndOffset: pos.EndOffset,
	}
}

func (m *rsMapper) comment(tok frontkit.Token) air.Node {
	kind := air.KindComment
	if rsHasPrefix(tok.Text, rsDocPrefixes) {
		kind = air.KindDoc
	}
	if rsHasPrefix(tok.Text, rsDirectivePrefixes) {
		node := m.leaf(tok, air.KindDirective)
		return m.annotate(node, air.SemDirective, strings.TrimSpace(tok.Text))
	}
	node := m.leaf(tok, kind)
	if kind == air.KindDoc {
		m.annotate(node, air.SemNote, m.note("doc-comment"))
	}
	return node
}

func (m *rsMapper) attribute(tokens []frontkit.Token, lo, hi int) air.Node {
	end := hi - 1
	if end < lo {
		end = lo
	}
	// Bind the exact source spelling of the whole attribute as the node's own
	// attribute text. Without this, the attribute is a region whose payload is
	// independent of its inner tokens, so a single-character mutation inside
	// #[cfg(...)] or #[allow(/* comment */ ...)] leaves the AIR identifier
	// unchanged and the mutation is recorded as uncaptured. Routing the real
	// source text through NewNodeAt makes every inner operator, digit, comment
	// and string contribute to the node identity.
	text := string(frontkit.SourceFromTokens(tokens[lo:hi]))
	node := m.build(air.KindAttribute, tokens[lo], tokens[end], text)
	if name := rsAttributeName(tokens, lo, hi); name != "" {
		m.annotate(node, air.SemTag, name)
	}
	if rsHasText(tokens, lo, hi, "!") {
		m.annotate(node, air.SemNote, m.note("inner-attribute"))
	}
	return node
}

func (m *rsMapper) match(tokens []frontkit.Token, lo, hi, bodyOpen,
	bodyClose int, ctx rsContext) rsMapping {
	lead := rsEffectiveLead(tokens, lo, hi)
	rule, ok := rsItemRuleFor(m.itemRules, lead)
	out := rsMapping{kind: air.KindExprStmt, name: rsNameOf(tokens, lo, hi)}
	if ok {
		out.kind = rule.kind
		out.noteName = rule.noteName
		out.semKey = rule.semKey
		out.semValue = rule.semValue
	}
	if ctx == rsCtxMatchArm && rsHasText(tokens, lo, hi, "=>") {
		out.kind, out.noteName = air.KindCase, "match-arm"
	}
	if (ctx == rsCtxImpl || ctx == rsCtxTrait) && lead == "fn" {
		out.kind, out.noteName = air.KindMethodDecl, "method"
	}
	rsMatchFnPrefix(lead, tokens, lo, hi, &out)
	if rsHasText(tokens, lo, hi, "mut") && out.kind == air.KindVarDecl {
		out.semKey, out.semValue = air.SemMutability, "mutable"
	}
	if lead == "if" && rsHasText(tokens, lo, hi, "let") {
		out.noteName = "if-let"
	}
	if lead == "while" && rsHasText(tokens, lo, hi, "let") {
		out.noteName = "while-let"
	}
	if !ok && out.kind == air.KindExprStmt {
		rsMatchComposite(tokens, lo, bodyOpen, bodyClose, &out)
	}
	if !out.kind.Valid() {
		out.kind = air.KindExprStmt
	}
	return out
}

func rsMatchFnPrefix(lead string, tokens []frontkit.Token, lo, hi int, out *rsMapping) {
	if rsHasText(tokens, lo, hi, "fn") &&
		(lead == "unsafe" || lead == "extern" || lead == "const") {
		out.kind = air.KindFuncDecl
		switch lead {
		case "extern":
			out.noteName = "extern-abi"
		case "const":
			out.noteName = "const-fn"
		default:
			out.noteName = "unsafe"
		}
	}
	if rsHasText(tokens, lo, hi, "unsafe") && out.kind == air.KindFuncDecl {
		out.noteName = "unsafe"
	}
}

func rsMatchComposite(tokens []frontkit.Token, lo, bodyOpen,
	bodyClose int, out *rsMapping) {
	if bodyOpen <= lo || bodyClose <= bodyOpen {
		return
	}
	if bodyOpen-1 >= len(tokens) {
		return
	}
	if tokens[bodyOpen-1].Kind != frontkit.TokenIdent {
		return
	}
	if !rsHasText(tokens, bodyOpen+1, bodyClose, ":") {
		return
	}
	out.kind, out.noteName = air.KindCompositeLit, "struct-literal"
}

func rsEffectiveLead(tokens []frontkit.Token, lo, hi int) string {
	at := rsEffectiveLeadIndex(tokens, lo, hi)
	if at < 0 {
		return ""
	}
	return tokens[at].Text
}

func rsNameOf(tokens []frontkit.Token, lo, hi int) string {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	for j := lo; j < hi; j++ {
		if tokens[j].Kind == frontkit.TokenIdent {
			return tokens[j].Text
		}
	}
	return ""
}

func rsHasText(tokens []frontkit.Token, lo, hi int, text string) bool {
	if lo < 0 {
		lo = 0
	}
	if hi > len(tokens) {
		hi = len(tokens)
	}
	for j := lo; j < hi; j++ {
		if tokens[j].Text == text {
			return true
		}
	}
	return false
}

func rsAttributeName(tokens []frontkit.Token, lo, hi int) string {
	return rsNameOf(tokens, lo, hi)
}

func rsOperatorSpan(start, end frontkit.Token) (pos, endPos air.Position) {
	finish := end
	if finish.EndOffset < start.Offset {
		finish = start
	}
	pos = air.NewPosition(start.Line, start.Column, start.Offset, finish.EndOffset)
	endPos = air.NewPosition(start.Line, start.Column, finish.EndOffset, finish.EndOffset)
	return pos, endPos
}

func (m *rsMapper) boundOperator(kind air.NodeKind, start, end frontkit.Token,
	tok frontkit.Token, children ...air.Node) air.Node {
	pos, endPos := rsOperatorSpan(start, end)
	return frontkit.BindOperator(kind, pos, endPos, tok,
		m.view.OrderChildren(children)...)
}
