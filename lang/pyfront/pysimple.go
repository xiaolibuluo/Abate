/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the simple-statement productions of pyParser: the
// keyword-led forms (assert, pass, del, global, nonlocal, yield, import),
// the dedicated construct markers they emit while being reduced, and the
// expression-statement, assignment and annotation forms.

// parseSimpleKeyword reduces a keyword-led simple statement.
func (p *pyParser) parseSimpleKeyword(tok frontkit.Token, rule pyStmtRule) air.Node {
	p.stream.Advance()
	// A from-import may start with relative dots; they form the import_prefix
	// construct and are read here, while the statement is being reduced.
	var prefix air.Node
	if rule.keyword == "from" {
		dots := 0
		for {
			cur, ok := p.stream.Peek()
			if !ok || cur.Text != "." {
				break
			}
			p.stream.Advance()
			dots++
		}
		if dots > 0 {
			prefix = p.mapper.importPrefix(tok, dots)
		}
	}
	var operand air.Node
	if cur, ok := p.stream.Peek(); ok && cur.Kind != frontkit.TokenNewline &&
		cur.Kind != frontkit.TokenDedent && cur.Kind != frontkit.TokenEOF &&
		cur.Text != ";" && cur.Text != ":" {
		operand = p.expr.parse(0)
	}
	if (rule.keyword == "import" || rule.keyword == "from") && operand != nil {
		if qualified := p.mapper.qualifiedName(operand); qualified != nil {
			operand = qualified
		}
		p.view.AddImport(air.AttributeOf(operand))
	}
	p.synchronize()
	end := tok
	children := make([]air.Node, 0, 3)
	if prefix != nil {
		children = append(children, prefix)
	}
	if operand != nil {
		children = append(children, operand)
		end = spanToken(operand, tok)
	}
	if marker := p.statementMarker(rule.keyword, tok, end, operand); marker != nil {
		children = append(children, marker)
	}
	node := p.finish(rule.kind, tok, end, tok.Text, children...)
	if rule.note != "" {
		p.mapper.annotate(node, air.SemNote, p.mapper.note(rule.note))
	}
	return node
}

// statementMarker emits the dedicated construct node for one simple
// statement the parser just reduced. Keeping the original outer kind intact
// preserves the byte-exact projection while making the construct observable.
func (p *pyParser) statementMarker(keyword string, start, end frontkit.Token,
	operand air.Node) air.Node {
	children := make([]air.Node, 0, 1)
	if operand != nil {
		children = append(children, operand)
	}
	switch keyword {
	case "assert":
		return p.mapper.dedicated(air.KindPyAssert, start, end,
			keyword, "py-assert", children...)
	case "pass":
		return p.mapper.dedicated(air.KindPyPass, start, end,
			keyword, "py-pass")
	case "del":
		return p.mapper.dedicated(air.KindPyDelete, start, end,
			keyword, "py-delete", children...)
	case "global", "nonlocal":
		return p.mapper.scopeStatement(start, end, children...)
	case "yield":
		return p.mapper.dedicated(air.KindPyYield, start, end,
			keyword, "py-yield", children...)
	default:
		return nil
	}
}

// parseSimple reduces an expression statement, an assignment or an annotation.
func (p *pyParser) parseSimple(tok frontkit.Token) air.Node {
	left := p.expr.parse(0)
	p.expr.skipTrivia()
	next, ok := p.stream.Peek()
	if ok && pyContains(pyAssignOps, next.Text) {
		p.stream.Advance()
		right := p.expr.parse(0)
		p.synchronize()
		end := spanToken(right, tok)
		// The parser has now reduced a concrete assignment form: emit the
		// dedicated node for augmented assignment or the walrus expression.
		children := []air.Node{left, right}
		switch {
		case next.Text == ":=":
			children = append(children, p.mapper.dedicated(
				air.KindPyNamedExpr, next, next, next.Text, "py-named-expr"))
		case next.Text != "=":
			children = append(children, p.mapper.dedicated(
				air.KindPyAugAssign, next, next, next.Text, "py-aug-assign"))
		}
		node := p.finish(air.KindAssign, tok, end, next.Text, children...)
		p.mapper.annotate(node, air.SemOperator, next.Text)
		if next.Text == ":=" {
			p.mapper.annotate(node, air.SemNote, p.mapper.note("walrus"))
		}
		return node
	}
	if ok && next.Text == ":" {
		p.stream.Advance()
		annotation := p.expr.parse(0)
		p.expr.skipTrivia()
		children := []air.Node{left}
		end := spanToken(annotation, tok)
		operator := ":"
		if annotation != nil {
			children = append(children, p.mapper.typeNode(annotation))
		}
		if rhs, found := p.stream.Peek(); found && pyContains(pyAssignOps, rhs.Text) {
			p.stream.Advance()
			value := p.expr.parse(0)
			children = append(children, value)
			end = spanToken(value, tok)
			operator = rhs.Text
		}
		p.synchronize()
		node := p.finish(air.KindDeclStmt, tok, end, ":", children...)
		if operator != ":" {
			p.mapper.annotate(node, air.SemOperator, operator)
		}
		return p.mapper.annotate(node, air.SemNote, p.mapper.note("type-annotation"))
	}
	p.synchronize()
	return p.finish(air.KindExprStmt, tok, spanToken(left, tok), "", left)
}
