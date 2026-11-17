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

// The keyword, operator and punctuation prefix reductions below consume
// the lexical tables directly, mirroring the pyfront file split.

// parseKeywordPrefix lowers literals, closures and keyword operands.
func (e *rsExpr) parseKeywordPrefix(tok frontkit.Token) air.Node {
	switch tok.Text {
	case "true", "false":
		return e.mapper.leaf(tok, air.KindBoolLit)
	case "move":
		if e.at < e.hi &&
			(e.tokens[e.at].Text == "|" || e.tokens[e.at].Text == "||") {
			open := e.tokens[e.at]
			e.at++
			return e.finishClosure(tok, open, "move")
		}
		return e.mapper.leaf(tok, air.KindIdent)
	case "unsafe":
		return e.parseUnsafePrefix(tok)
	case "self":
		node := e.mapper.leaf(tok, air.KindRsSelf)
		return e.mapper.annotate(node, air.SemNote, e.mapper.note("rs-self"))
	case "super":
		node := e.mapper.leaf(tok, air.KindRsSuper)
		return e.mapper.annotate(node, air.SemNote, e.mapper.note("rs-super"))
	case "crate":
		node := e.mapper.leaf(tok, air.KindRsCrate)
		return e.mapper.annotate(node, air.SemNote, e.mapper.note("rs-crate"))
	case "yield":
		node := e.mapper.leaf(tok, air.KindRsYield)
		return e.mapper.annotate(node, air.SemNote, e.mapper.note("yield"))
	case "dyn":
		return e.mapper.annotate(e.mapper.leaf(tok, air.KindNamed),
			air.SemNote, e.mapper.note("dyn"))
	case "box":
		return e.mapper.unary(tok, e.parseExpression(rsPrefixPower),
			air.KindUnary, "box")
	case "async", "const", "gen", "try":
		// async/const/gen/try followed by "{" reduce as expression-level
		// block variants; the dedicated node anchors on the keyword the
		// parser just consumed.
		if kind, note, isBlock := rsBlockMarkerKind(tok.Text); isBlock &&
			e.at < e.hi && e.tokens[e.at].Text == "{" {
			// 真正识别到 "{...}" 块时才产出 dedicated 标记；不再用关键
			// 字拼写给 BlockVariant 等伪 kind 贴标签。
			return e.mapper.dedicated(kind, tok, tok.Text, note)
		}
		return e.mapper.leaf(tok, air.KindIdent)
	default:
		// 其它关键字保持冻结的标识符投影；专用 Kind 只由真实语法归约
		// （本函数上述分支及 item/statement 解析）产出。
		return e.mapper.leaf(tok, air.KindIdent)
	}
}

// parseUnsafePrefix lowers an unsafe block nested inside an expression. The
// dedicated unsafe-block node is a child of the block the parser just built.
func (e *rsExpr) parseUnsafePrefix(tok frontkit.Token) air.Node {
	if e.at < e.hi && e.tokens[e.at].Text == "{" {
		open := e.tokens[e.at]
		e.at++
		children, closeTok := e.parseDelimited(open, "}")
		children = append(children, e.mapper.dedicated(
			air.KindRsUnsafeBlock, tok, tok.Text, "rs-unsafe-block"))
		node := e.mapper.build(air.KindBlock, tok, closeTok, "", children...)
		return e.mapper.annotate(node, air.SemNote, e.mapper.note("unsafe"))
	}
	return e.mapper.leaf(tok, air.KindIdent)
}

// rsBlockMarkerKind maps an expression-level block-variant keyword onto its
// dedicated kind; ok reports whether the keyword opens a block here.
func rsBlockMarkerKind(text string) (air.NodeKind, string, bool) {
	switch text {
	case "async":
		return air.KindRsAsyncBlock, "rs-async-block", true
	case "const":
		return air.KindRsConstBlock, "rs-const-block", true
	case "gen":
		return air.KindRsGenBlock, "rs-gen-block", true
	case "try":
		return air.KindRsTryBlock, "rs-try-block", true
	}
	return air.KindInvalid, "", false
}

// parseOpPrefix lowers prefix operators, lifetimes, closures and ranges.
func (e *rsExpr) parseOpPrefix(tok frontkit.Token) air.Node {
	switch tok.Text {
	case "&":
		if e.at < e.hi && e.tokens[e.at].Text == "mut" {
			e.at++
			return e.mapper.mutRefNode(tok, e.parseExpression(rsPrefixPower))
		}
		return e.mapper.refNode(tok, e.parseExpression(rsPrefixPower))
	case "&&":
		return e.mapper.addrNode(tok, e.parseExpression(rsPrefixPower))
	case "*":
		return e.mapper.starNode(tok, e.parseExpression(rsPrefixPower))
	case "-":
		operand := e.parseExpression(rsPrefixPower)
		var marks []air.Node
		if operand != nil && operand.Kind() == air.KindNumLit {
			// The parser reduced a unary minus over a numeric literal: a
			// genuine negative literal.
			marks = append(marks, e.mapper.dedicated(air.KindRsNegativeLit,
				tok, tok.Text, "rs-negative-literal"))
		}
		return e.mapper.unaryMarked(tok, operand, air.KindUnary, "", marks...)
	case "!":
		return e.mapper.unary(tok, e.parseExpression(rsPrefixPower),
			air.KindUnary, "")
	case "'":
		return e.parseQuotePrefix(tok)
	case "|":
		return e.finishClosure(tok, tok, "closure")
	case "||":
		return e.finishClosure(tok, frontkit.Token{}, "closure")
	case "..", "..=":
		return e.mapper.rangeNode(nil, tok, e.parseExpression(rsRangePower))
	case "::":
		return e.mapper.leaf(tok, air.KindIdent)
	default:
		return nil
	}
}

// parsePunctPrefix lowers parenthesized, array, braced and metavariable forms.
func (e *rsExpr) parsePunctPrefix(tok frontkit.Token) air.Node {
	switch tok.Text {
	case "(":
		children, closeTok := e.parseDelimited(tok, ")")
		return e.mapper.tuple(tok, closeTok, children...)
	case "[":
		children, closeTok := e.parseDelimited(tok, "]")
		return e.mapper.array(tok, closeTok, children...)
	case "{":
		children, closeTok := e.parseDelimited(tok, "}")
		return e.mapper.block(tok, closeTok, children...)
	case "$":
		// A metavariable of a macro_rules body: $name reduces as a unit, and
		// the dedicated node rides as a child of the name leaf so it never
		// overlaps an enclosing composite projection.
		if e.at < e.hi && e.tokens[e.at].Kind == frontkit.TokenIdent {
			name := e.tokens[e.at]
			e.at++
			marker := e.mapper.dedicated(air.KindRsMetavar, name, name.Text,
				"rs-metavariable")
			return e.mapper.build(air.KindIdent, name, name, name.Text, marker)
		}
		// 无后续标识符的 "$" 只是未归约片段，不能贴 SyntaxDetail 伪标签。
		return e.mapper.partial(e.mapper.leaf(tok, air.KindPartialNode))
	default:
		return nil
	}
}
