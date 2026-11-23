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

// This file holds the comprehension clause reductions: the for/if clause loop,
// the comma-separated target list of a for clause, the async for form and the
// clause node that keeps every clause part inside one clause interval.

// clause lowers one comprehension for/if clause.
func (m *pyMap) clause(start frontkit.Token, body air.Node) air.Node {
	kind := air.KindIf
	if start.Text == "for" {
		kind = air.KindRangeFor
	}
	return m.clauseParts(start, kind, body)
}

// clauseParts lowers one comprehension clause over ordered body parts. The
// clause spans from the clause keyword to the end of its last part, so every
// part stays inside the clause interval.
func (m *pyMap) clauseParts(start frontkit.Token, kind air.NodeKind,
	parts ...air.Node) air.Node {
	end := start
	for _, part := range parts {
		if part == nil {
			continue
		}
		token := spanToken(part, end)
		if token.EndOffset > end.EndOffset {
			end = token
		}
	}
	all := append(append([]air.Node{}, parts...),
		m.dedicated(air.KindPyCompClause, start, end,
			start.Text, "py-comp-clause", parts...))
	node := m.build(kind, start, end, start.Text, all...)
	return m.annotate(node, air.SemNote, m.note("comprehension-clause"))
}

// parseClauses reduces the for/if clauses of a comprehension.
func (e *pyExpr) parseClauses() []air.Node {
	out := make([]air.Node, 0, 2)
	for {
		e.skipTrivia()
		tok, ok := e.stream.Peek()
		if !ok || tok.Kind == frontkit.TokenEOF ||
			!pyContains(pyClauseKeywords, tok.Text) {
			return out
		}
		e.stream.Advance()
		switch tok.Text {
		case "for":
			out = append(out, e.parseForClause(tok))
		case "async":
			out = append(out, e.parseAsyncClause(tok))
		default:
			out = append(out, e.mapper.clause(tok, e.parse(pyClausePower)))
		}
	}
}

// parseAsyncClause reduces an async for clause of an async comprehension. When
// the async keyword is not followed by a for clause it falls back to a plain
// clause without the async marker.
func (e *pyExpr) parseAsyncClause(asyncTok frontkit.Token) air.Node {
	e.skipTrivia()
	forTok, ok := e.stream.Peek()
	if !ok || forTok.Text != "for" {
		return e.mapper.clause(asyncTok, e.parse(pyClausePower))
	}
	e.stream.Advance()
	return e.mapper.annotate(e.parseForClause(forTok), air.SemAsync, "true")
}

// parseForClause reduces one for clause: the target list, the in keyword and
// the iterated expression. The clause spans from the for keyword to the end of
// the iterated expression, so every target stays inside the clause interval
// instead of escaping into the enclosing statement.
func (e *pyExpr) parseForClause(start frontkit.Token) air.Node {
	targets := e.parseTargetList()
	e.skipTrivia()
	tok, ok := e.stream.Peek()
	if !ok || tok.Text != "in" {
		return e.mapper.clauseParts(start, air.KindRangeFor, e.pyTargetNode(targets))
	}
	e.stream.Advance()
	iterated := e.parse(pyClausePower)
	return e.mapper.clauseParts(start, air.KindRangeFor, e.pyTargetNode(targets), iterated)
}

// parseTargetList reduces the comma-separated target list of a for clause.
// The target binding power excludes the in operator, so the list stops before
// the iterated expression and each target keeps its own byte-exact span.
func (e *pyExpr) parseTargetList() []air.Node {
	targets := make([]air.Node, 0, 2)
	for {
		e.skipTrivia()
		before := e.stream.Remaining()
		targets = append(targets, e.parse(pyTargetPower))
		if e.stream.Remaining() >= before {
			return targets
		}
		e.skipTrivia()
		tok, ok := e.stream.Peek()
		if !ok || tok.Text != "," {
			return targets
		}
		e.stream.Advance()
	}
}

// pyTargetNode lowers a for-clause target list to one node: the single target
// itself, or an expression-list node that spans a comma-separated list.
func (e *pyExpr) pyTargetNode(targets []air.Node) air.Node {
	switch len(targets) {
	case 0:
		return nil
	case 1:
		return targets[0]
	default:
		first := spanToken(targets[0], frontkit.Token{})
		last := spanToken(targets[len(targets)-1], first)
		return e.mapper.exprListNode(first, last, targets...)
	}
}
