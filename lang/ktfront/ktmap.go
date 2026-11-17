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

// ktMap 构造 AIR 节点并应用 Kotlin 构造到冻结枚举的降级映射。
type ktMap struct {
	view     *frontkit.View
	notes    []ktNote
	nullSafe []ktNullSafeRule
}

// newKtMap 绑定可注入的标注表与空安全映射表。
func newKtMap(notes []ktNote, nullSafe []ktNullSafeRule) *ktMap {
	if len(notes) == 0 {
		notes = ktNotes
	}
	if len(nullSafe) == 0 {
		nullSafe = ktNullSafeRules
	}
	return &ktMap{notes: notes, nullSafe: nullSafe}
}

// bindView 绑定单文件的节点构造视图。
func (m *ktMap) bindView(view *frontkit.View) { m.view = view }

// note 解析一个命名语义标注；表外名字回退到名字本身，保持确定性。
func (m *ktMap) note(name string) string {
	if value := ktNoteValue(m.notes, name); value != "" {
		return value
	}
	return name
}

// annotate 附加一个语义值并返回同一节点。
func (*ktMap) annotate(node air.Node, key air.SemKey, value string) air.Node {
	if node != nil {
		_ = air.Annotate(node, key, value)
	}
	return node
}

// leafKind 把一个 token 投影为指定种类的字节精确叶子。
func (*ktMap) leafKind(tok frontkit.Token, kind air.NodeKind,
	children ...air.Node) air.Node {
	start, end := tok.Span()
	return air.NewNodeAt(kind, start, end, tok.Text, children...)
}

// literal 把一个词法 token 按冻结枚举投影为 AIR 叶子。
//
// 词法 token 与关键字保持冻结投影；标点、运算符等非词法 token 不再按拼写
// 贴专用 Kind 标签——它们由真实语法归约的父节点承载语义，孤立出现时以
// PartialNode + SemPartial 诚实标记为未归约片段。
func (m *ktMap) literal(tok frontkit.Token) air.Node {
	var kind air.NodeKind
	switch tok.Kind {
	case frontkit.TokenIdent:
		kind = air.KindIdent
	case frontkit.TokenKeyword:
		switch tok.Text {
		case "true", "false":
			kind = air.KindBoolLit
		case "null":
			kind = air.KindNilLit
		default:
			kind = air.KindIdent
		}
	case frontkit.TokenNum:
		kind = air.KindNumLit
	case frontkit.TokenStr:
		return m.stringLiteral(tok)
	case frontkit.TokenChar:
		kind = air.KindCharLit
	case frontkit.TokenComment:
		return m.comment(tok)
	case frontkit.TokenIllegal:
		return m.annotate(m.leafKind(tok, air.KindPartialNode), air.SemPartial, "true")
	default:
		if rule, ok := ktNullSafeFor(m.nullSafe, tok.Text); ok {
			node := m.leafKind(tok, rule.kind)
			return m.annotate(node, air.SemNote, m.note(rule.note))
		}
		return m.annotate(m.leafKind(tok, air.KindPartialNode),
			air.SemPartial, "true")
	}
	return m.leafKind(tok, kind)
}

// stringLiteral 归约字符串模板、三引号原始字符串、反引号标识符与普通字符串。
//
// 归约的同时产出字符串片段的 dedicated 子节点：普通/原始字符串补内容与
// 收尾定界符，模板字符串因内容与插值节点交错，只补收尾定界符。
func (m *ktMap) stringLiteral(tok frontkit.Token) air.Node {
	switch ktStringFlavor(tok.Text) {
	case "backtick-ident":
		node := m.leafKind(tok, air.KindIdent)
		return m.annotate(node, air.SemNote, m.note("backtick-ident"))
	case "string-template":
		parts := ktTemplateParts(m, tok)
		var children []air.Node
		if len(parts) > 0 {
			children = append(children, m.dedicated(air.KindKtInterp, tok,
				tok.Text, "kt-interp"))
		}
		children = append(children, parts...)
		children = append(children, m.delimiterFragment(tok)...)
		node := m.leafKind(tok, air.KindStrLit, children...)
		m.annotate(node, air.SemNote, m.note("string-template"))
		return m.annotate(node, air.SemTag, "string-template")
	case "raw-string":
		node := m.leafKind(tok, air.KindStrLit, m.stringFragments(tok)...)
		m.annotate(node, air.SemNote, m.note("raw-string"))
		return m.annotate(node, air.SemTag, "raw-string")
	default:
		return m.leafKind(tok, air.KindStrLit, m.stringFragments(tok)...)
	}
}

// comment 归约普通注释与 KDoc 文档注释。
func (m *ktMap) comment(tok frontkit.Token) air.Node {
	if ktIsDocComment(tok.Text) {
		node := m.leafKind(tok, air.KindDoc)
		return m.annotate(node, air.SemNote, m.note("kdoc"))
	}
	return m.leafKind(tok, air.KindComment)
}

// infixNode 构造一个覆盖 [left.Start, right.End) 的中缀节点。
func (m *ktMap) infixNode(kind air.NodeKind, left air.Node, tok frontkit.Token,
	right air.Node) air.Node {
	if left == nil || right == nil {
		return m.leafKind(tok, kind)
	}
	lp := left.Pos()
	end := right.End()
	pos := air.NewPosition(lp.Line, lp.Column, lp.Offset, end.Offset)
	children := []air.Node{left, right}
	if m.view != nil {
		children = m.view.OrderChildren(children)
	}
	return frontkit.BindOperator(kind, pos, end, tok, children...)
}

// unaryNode 构造一个覆盖 [tok.Start, operand.End) 的一元节点。
func (m *ktMap) unaryNode(kind air.NodeKind, tok frontkit.Token,
	operand air.Node) air.Node {
	if operand == nil {
		return m.leafKind(tok, kind)
	}
	start, _ := tok.Span()
	end := operand.End()
	pos := air.NewPosition(start.Line, start.Column, start.Offset, end.Offset)
	children := []air.Node{m.literal(tok), operand}
	if m.view != nil {
		children = m.view.OrderChildren(children)
	}
	return frontkit.BindOperator(kind, pos, end, tok, children...)
}

// classify 依据主关键字、修饰符与区间特征，确定性地选择一个构造映射。
func (m *ktMap) classify(tokens []frontkit.Token, lo, hi int,
	scope ktScope, hasBody bool) ktMapping {
	at := ktLeadIndex(tokens, lo, hi)
	lead := ""
	if at < hi && at < len(tokens) {
		lead = tokens[at].Text
	}
	result := ktMapping{name: ktNameOf(tokens, lo, hi), kind: air.KindVarDecl}
	matched := false
	if rule, ok := ktDeclRuleFor(lead); ok {
		result.kind, result.note = rule.kind, rule.note
		result.semKey, result.semValue = rule.semKey, rule.semValue
		matched = true
	}
	m.applyModifiers(tokens, lo, hi, lead, &result)
	if !matched {
		m.applyFallback(tokens, lo, hi, scope, hasBody, &result)
	}
	return result
}

// applyModifiers 把 Kotlin 修饰符翻译成冻结 AIR 种类与语义标注。
func (*ktMap) applyModifiers(tokens []frontkit.Token, lo, hi int,
	lead string, result *ktMapping) {
	if ktHasText(tokens, lo, hi, "suspend") {
		result.async = true
		if result.note == "" {
			result.note = "suspend"
		}
	}
	if ktHasText(tokens, lo, hi, "const") && lead == "val" {
		result.kind, result.note = air.KindConstDecl, "const-val"
	}
	if lead == "class" {
		switch {
		case ktHasText(tokens, lo, hi, "enum"):
			result.kind, result.note = air.KindEnumDecl, "enum-class"
		case ktHasText(tokens, lo, hi, "annotation"):
			result.kind, result.note = air.KindAnnotationDecl, "annotation-class"
		case ktHasText(tokens, lo, hi, "data"):
			result.note = "data-class"
		case ktHasText(tokens, lo, hi, "sealed"):
			result.note = "sealed-class"
		case ktHasText(tokens, lo, hi, "inner"):
			result.note = "inner-class"
		}
	}
	if lead == "object" && ktHasText(tokens, lo, hi, "companion") {
		result.note = "companion-object"
	}
	if lead == "fun" {
		switch {
		case ktHasExtensionDot(tokens, lo, hi):
			result.note = "extension"
		case ktHasText(tokens, lo, hi, "operator"):
			result.note = "operator"
		case ktHasText(tokens, lo, hi, "infix"):
			result.note = "infix"
		}
	}
	if ktHasText(tokens, lo, hi, "lateinit") {
		result.mutability = "lateinit"
	}
}

// applyFallback 对未匹配声明规则的区间按结构特征归类。
func (*ktMap) applyFallback(tokens []frontkit.Token, lo, hi int,
	scope ktScope, hasBody bool, result *ktMapping) {
	hasParen := ktHasText(tokens, lo, hi, "(")
	switch {
	case ktHasText(tokens, lo, hi, "->"):
		result.kind = air.KindLambda
		if ktHasReceiverHint(tokens, lo, hi) {
			result.note = "lambda-with-receiver"
		} else {
			result.note = "lambda"
		}
	case hasBody && hasParen:
		result.kind = air.KindMethodDecl
	case hasBody && scope == ktScopeType:
		result.kind = air.KindFuncDecl
	case hasBody && scope == ktScopeBody:
		result.kind = air.KindBlock
	case hasParen:
		result.kind = air.KindCall
	default:
		result.kind = air.KindVarDecl
	}
}

// ktHasReceiverHint 报告区间是否含 lambda 接收者提示（this/it 或 Type.()）。
func ktHasReceiverHint(tokens []frontkit.Token, lo, hi int) bool {
	for j := lo; j < hi && j < len(tokens); j++ {
		if ktContains(ktLambdaReceiverWords, tokens[j].Text) {
			return true
		}
		if tokens[j].Text == "." && j+2 < hi &&
			tokens[j+1].Text == "(" && tokens[j+2].Text == ")" {
			return true
		}
	}
	return false
}

// apply 把映射的语义标注写到节点上，全部使用真实 SemKey。
func (m *ktMap) apply(node air.Node, mapping ktMapping) {
	if node == nil {
		return
	}
	if mapping.note != "" {
		m.annotate(node, air.SemNote, m.note(mapping.note))
	}
	if mapping.semKey.Valid() && mapping.semValue != "" {
		m.annotate(node, mapping.semKey, mapping.semValue)
	}
	if mapping.async {
		m.annotate(node, air.SemAsync, "true")
	}
	if mapping.partial {
		m.annotate(node, air.SemPartial, "true")
	}
	if mapping.mutability != "" {
		m.annotate(node, air.SemMutability, mapping.mutability)
	}
}
