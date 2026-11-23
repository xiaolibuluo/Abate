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

// csMapping 是一个 C# 构造归约后的 AIR 种类与语义标注集合。
type csMapping struct {
	kind       air.NodeKind
	name       string
	note       string
	semKey     air.SemKey
	semValue   string
	async      bool
	partial    bool
	mutability string
}

// csMap 构造 AIR 节点并应用 C# 构造到冻结枚举的降级映射。
type csMap struct {
	view         *frontkit.View
	notes        []csNote
	linqKeywords []string
}

// newCsMap 绑定可注入的标注表与 LINQ 关键字表。
func newCsMap(notes []csNote, linq []string) *csMap {
	if len(notes) == 0 {
		notes = csNotes
	}
	if len(linq) == 0 {
		linq = csLinqKeywords
	}
	return &csMap{notes: notes, linqKeywords: linq}
}

// bindView 绑定单文件的节点构造视图。
func (m *csMap) bindView(view *frontkit.View) { m.view = view }

// note 解析一个命名语义标注；表外名字回退到名字本身，保持确定性。
func (m *csMap) note(name string) string {
	if value := csNoteFor(m.notes, name); value != "" {
		return value
	}
	return name
}

// annotate 附加一个语义值并返回同一节点。
func (*csMap) annotate(node air.Node, key air.SemKey, value string) air.Node {
	if node != nil {
		_ = air.Annotate(node, key, value)
	}
	return node
}

// leafKind 把一个 token 投影为指定种类的字节精确叶子。
func (*csMap) leafKind(tok frontkit.Token, kind air.NodeKind,
	children ...air.Node) air.Node {
	start, end := tok.Span()
	return air.NewNodeAt(kind, start, end, tok.Text, children...)
}

// literal 把一个词法 token 按冻结枚举投影为 AIR 叶子。
//
// 关键字按冻结枚举投影（保留真实拼写），畸形字节诚实标注 partial；标点与
// 运算符不设独立 AIR 种类——它们折叠进结构节点或经 BindTokenSpellings 绑定
// 拼写，因此返回 nil，调用方跳过。
func (m *csMap) literal(tok frontkit.Token) air.Node {
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
			if csContains(csTypeWords, tok.Text) {
				kind = air.KindNamed
			}
		}
	case frontkit.TokenNum:
		kind = air.KindNumLit
	case frontkit.TokenStr:
		return m.stringLiteral(tok)
	case frontkit.TokenChar:
		return m.leafKind(tok, air.KindCharLit, m.charBody(tok)...)
	case frontkit.TokenComment:
		return m.comment(tok)
	case frontkit.TokenIllegal:
		return m.annotate(m.leafKind(tok, air.KindPartialNode), air.SemPartial, "true")
	default:
		// 标点、运算符与 trivia 标记没有独立 AIR 种类，也不属于词法保真
		// 等价子集：不产节点，拼写由 BindTokenSpellings 兜底绑定。
		return nil
	}
	return m.leafKind(tok, kind)
}

// comment 归约普通注释、XML 文档注释与块注释。
func (m *csMap) comment(tok frontkit.Token) air.Node {
	kind := air.KindComment
	if csIsDocComment(tok.Text) {
		kind = air.KindDoc
		node := m.leafKind(tok, kind)
		return m.annotate(node, air.SemNote, m.note("xml-doc"))
	}
	return m.leafKind(tok, kind)
}

// infixNode 构造一个覆盖 [left.Start, right.End) 的中缀节点。
func (m *csMap) infixNode(kind air.NodeKind, left air.Node, tok frontkit.Token,
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
func (m *csMap) unaryNode(kind air.NodeKind, tok frontkit.Token,
	operand air.Node) air.Node {
	if operand == nil {
		return m.leafKind(tok, kind)
	}
	start, _ := tok.Span()
	end := operand.End()
	pos := air.NewPosition(start.Line, start.Column, start.Offset, end.Offset)
	children := make([]air.Node, 0, 2)
	if leaf := m.literal(tok); leaf != nil {
		children = append(children, leaf)
	}
	children = append(children, operand)
	if m.view != nil {
		children = m.view.OrderChildren(children)
	}
	return frontkit.BindOperator(kind, pos, end, tok, children...)
}

// classify 依据首关键字与区间特征，确定性地选择一个构造映射。
func (m *csMap) classify(tokens []frontkit.Token, lo, hi int,
	scope csScope, hasBody bool) csMapping {
	at := csLeadIndex(tokens, lo, hi)
	lead := ""
	if at < hi && at < len(tokens) {
		lead = tokens[at].Text
	}
	result := csMapping{name: csNameOf(tokens, lo, hi)}
	if csHasText(tokens, lo, hi, "async") {
		result.async = true
	}
	if csHasText(tokens, lo, hi, "partial") {
		result.partial = true
	}
	if csHasText(tokens, lo, hi, "readonly") {
		result.mutability = "readonly"
	}
	if rule, ok := csDeclRuleFor(lead); ok {
		result.kind, result.note = rule.kind, rule.note
		result.semKey, result.semValue = rule.semKey, rule.semValue
		if lead == "using" && csHasText(tokens, lo, hi, "(") {
			result.kind, result.note = air.KindUsingDecl, "using-statement"
		}
	}
	if csContains(csLinqStartKeywords, lead) ||
		(csHasText(tokens, lo, hi, "from") &&
			csHasLinqTail(tokens, lo, hi, m.linqKeywords)) {
		result.kind, result.note = air.KindCall, "linq"
	}
	if result.kind == air.KindInvalid {
		// 标签语句：首 token 是标识符，紧随的下一有效 token 是冒号。
		if at < hi-1 && tokens[at].Kind == frontkit.TokenIdent &&
			tokens[at+1].Text == ":" && scope == csScopeBody {
			result.kind, result.note = air.KindLabel, "label"
			return result
		}
		// 表达式语句：在函数体/块作用域中无体无括号时，首标记若非
		// 声明关键字且其后不紧跟另一标识符（排除 "Type name;" 模式），
		// 则为独立的表达式语句。
		if scope == csScopeBody && !hasBody {
			next := at + 1
			for next < hi && next < len(tokens) {
				tk2 := tokens[next]
				if tk2.Kind != frontkit.TokenComment &&
					tk2.Kind != frontkit.TokenNewline &&
					tk2.Text != "" {
					break
				}
				next++
			}
			isDecl := next < hi &&
				tokens[next].Kind == frontkit.TokenIdent &&
				tokens[next].Text != ";" &&
				tokens[next].Text != "="
			if !isDecl {
				result.kind, result.note = air.KindExprStmt,
					"expression-statement"
				return result
			}
		}
		hasParen := csHasText(tokens, lo, hi, "(")
		switch {
		case csHasText(tokens, lo, hi, "this") && csHasText(tokens, lo, hi, "["):
			result.kind, result.note = air.KindIndex, "indexer"
		case hasBody && hasParen:
			result.kind = air.KindMethodDecl
		case hasBody && scope == csScopeType:
			result.kind, result.note = air.KindFuncDecl, "property"
		case hasBody && scope == csScopeBody:
			result.kind = air.KindBlock
		case hasParen:
			result.kind = air.KindCall
		default:
			result.kind = air.KindVarDecl
		}
	}
	return result
}

// apply 把映射的语义标注写到节点上，全部使用真实 SemKey。
func (m *csMap) apply(node air.Node, mapping csMapping) {
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
