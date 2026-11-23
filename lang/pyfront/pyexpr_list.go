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

// pyListMode selects whether a list allows slice colon separators.
type pyListMode int

const (
	pyListPlain pyListMode = iota
	pyListSlice
)

// pySubscriptKind selects slice or index lowering for a bracketed subscript.
type pySubscriptKind int

const (
	pySubscriptIndex pySubscriptKind = iota
	pySubscriptSlice
)

// parseSuffix reduces a postfix attribute, call, index or generic subscript.
func (e *pyExpr) parseSuffix(left air.Node, tok frontkit.Token) air.Node {
	switch tok.Text {
	case ".":
		e.skipTrivia()
		name, ok := e.stream.Advance()
		if !ok {
			return e.mapper.partial(left)
		}
		return e.mapper.field(left, name)
	case "(":
		e.bracket++
		args, closeTok, _ := e.parseList(")", pyListPlain)
		e.bracket--
		return e.mapper.call(left, tok, closeTok, args...)
	case "[":
		e.bracket++
		parts, closeTok, sliced := e.parseList("]", pyListSlice)
		e.bracket--
		kind := pySubscriptIndex
		if sliced {
			kind = pySubscriptSlice
		}
		return e.finishSubscript(left, tok, closeTok, parts, kind)
	default:
		return e.mapper.partial(left)
	}
}

// finishSubscript chooses generic, optional, slice or index lowering.
func (e *pyExpr) finishSubscript(left air.Node, open, endTok frontkit.Token,
	parts []air.Node, kind pySubscriptKind) air.Node {
	baseName := air.AttributeOf(left)
	sliced := kind == pySubscriptSlice
	if pyContains(pyGenericNames, baseName) && !sliced {
		if baseName == "Optional" {
			children := append([]air.Node{left}, parts...)
			return e.mapper.optional(spanToken(left, open), endTok, children...)
		}
		return e.mapper.genericInst(left, endTok, parts...)
	}
	if sliced {
		return e.mapper.slice(left, open, endTok, parts...)
	}
	return e.mapper.index(left, open, endTok, parts...)
}

// parseList reduces comma-separated parts until one closing delimiter. Colons
// mark a slice only when the list mode is pyListSlice, so the same loop serves
// calls, parameter lists, subscripts and literal containers.
func (e *pyExpr) parseList(closeText string, mode pyListMode) ([]air.Node,
	frontkit.Token, bool) {
	parts := make([]air.Node, 0, 4)
	sliced := false
	allowSlice := mode == pyListSlice
	for {
		e.skipTrivia()
		tok, ok := e.stream.Peek()
		if !ok || tok.Kind == frontkit.TokenEOF {
			return parts, tok, sliced
		}
		if tok.Text == closeText {
			e.stream.Advance()
			return parts, tok, sliced
		}
		if pySeparator(tok.Text, allowSlice) {
			if tok.Text == ":" {
				sliced = true
			}
			e.stream.Advance()
			continue
		}
		if pyAtListEnd(tok) {
			return parts, tok, sliced
		}
		part := e.parse(0)
		e.skipTrivia()
		if eq, found := e.stream.Peek(); found && eq.Text == "=" &&
			part != nil && part.Kind() == air.KindIdent {
			// A name=value entry is a keyword argument, a distinct construct.
			e.stream.Advance()
			value := e.parse(0)
			end := spanToken(value, eq)
			part = e.mapper.dedicated(air.KindPyKeywordArg,
				spanToken(part, eq), end, air.AttributeOf(part),
				"py-keyword-arg", part, value)
		}
		parts = append(parts, part)
		e.skipTrivia()
		next, ok := e.stream.Peek()
		if !ok {
			return parts, next, sliced
		}
		if next.Text == closeText {
			e.stream.Advance()
			return parts, next, sliced
		}
		if pySeparator(next.Text, allowSlice) {
			if next.Text == ":" {
				sliced = true
			}
			e.stream.Advance()
			continue
		}
		if pyAtListEnd(next) {
			return parts, next, sliced
		}
		if pyIsLexicalKind(next.Kind) {
			parts = append(parts, e.mapper.partial(e.mapper.lexical(next)))
		}
		e.stream.Advance()
	}
}

// parseExprList reduces a comma-separated expression list (expression_list
// in the Python grammar) to a single node: a lone expression returns as-is;
// two or more are wrapped in KindPyExprList anchored on real token spans.
// Because comma carries binding power 0 in pyInfixBindings, the Pratt loop
// discards it; this method fills the gap by reading successive expressions
// separated by commas until a statement terminator, the end of stream, or a
// token that cannot start an expression.
func (e *pyExpr) parseExprList() air.Node {
	first := e.parse(0)
	items := []air.Node{first}
	for {
		e.skipTrivia()
		tok, ok := e.stream.Peek()
		if !ok || tok.Text != "," {
			break
		}
		e.stream.Advance()
		items = append(items, e.parse(0))
	}
	if len(items) == 1 {
		return items[0]
	}
	start := spanToken(items[0], frontkit.Token{})
	end := spanToken(items[len(items)-1], start)
	return e.mapper.exprListNode(start, end, items...)
}
