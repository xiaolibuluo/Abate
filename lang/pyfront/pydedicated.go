/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file carries the genuine G4 reductions of pyfront. Every helper here is
// called from the parser while it lowers a concrete Python construct: an
// assert statement, a comprehension clause, an f-string conversion and so on.
// The emitted node is anchored on that construct's real token span, so the
// kind is produced by syntax-node semantics rather than by switching on a
// detached token's spelling.

// dedicated builds one semantic node for a construct the parser just reduced.
// The node keeps the construct's exact source span and a stable note, which
// makes it a real part of the AIR tree rather than a bare tag.
func (m *pyMap) dedicated(kind air.NodeKind, start, end frontkit.Token,
	text, note string, children ...air.Node) air.Node {
	node := m.build(kind, start, end, text, children...)
	return m.annotate(node, air.SemNote, note)
}

// qualifiedName lowers a parsed dotted access chain into a dedicated qualified
// name node. It is called only from import and type-annotation reductions,
// where tree-sitter-python reports a dotted_name or member_type construct.
func (m *pyMap) qualifiedName(expr air.Node) air.Node {
	return m.qualifiedNameAs(expr, ".")
}

// qualifiedNameAs anchors the dedicated node on the dot between qualifier and
// name. That span is a real token of the construct and never overlaps the
// enclosing projections, so it survives View.OrderChildren.
func (m *pyMap) qualifiedNameAs(expr air.Node, text string) air.Node {
	if expr == nil || expr.Kind() != air.KindField {
		return nil
	}
	parts := expr.Children()
	if len(parts) < 2 {
		return nil
	}
	base := parts[len(parts)-2].Pos()
	name := parts[len(parts)-1].Pos()
	if name.Offset <= base.EndOffset {
		return nil
	}
	start := frontkit.Token{
		Line: name.Line, Column: name.Column,
		Offset: base.EndOffset, EndOffset: base.EndOffset,
	}
	end := frontkit.Token{
		Line: name.Line, Column: name.Column,
		Offset: name.Offset, EndOffset: name.Offset,
	}
	return m.dedicated(air.KindPyQualifiedName, start, end,
		text, "py-qualified-name")
}

// operatorKind classifies an infix operator the Pratt loop has actually
// consumed as part of a binary expression. The order mirrors pyInfixBindings:
// boolean connectives, comparisons, then ordinary arithmetic.
func pyOperatorKind(text string) air.NodeKind {
	switch text {
	case "and", "or":
		return air.KindPyBoolOp
	case "==", "!=", "<", "<=", ">", ">=", "in", "is":
		return air.KindPyCompare
	default:
		return air.KindInvalid
	}
}

// pyQuoteWidth measures the quote width of a literal token at quoteAt: one
// for plain strings, three for triple-quoted docstrings.
func pyQuoteWidth(text string, quoteAt int) int {
	quote := text[quoteAt]
	if quoteAt+2 < len(text) && text[quoteAt+1] == quote &&
		text[quoteAt+2] == quote {
		return 3
	}
	if quoteAt+1 < len(text) && text[quoteAt+1] == quote {
		return 2
	}
	return 1
}

// stringFragments models the parts of a string literal the parser just
// lowering: its opening delimiter, its content, escape sequences and its
// closing delimiter. Each part is a real span inside the literal token.
func (m *pyMap) stringFragments(tok frontkit.Token) []air.Node {
	text := tok.Text
	quoteAt := strings.IndexAny(text, "\"'")
	if quoteAt < 0 {
		return nil
	}
	closeStart := len(text) - pyQuoteWidth(text, quoteAt)
	if closeStart < quoteAt {
		return nil
	}
	parts := make([]air.Node, 0, 3)
	appendPart := func(from, to int, note string) {
		if to <= from {
			return
		}
		part := frontkit.Token{
			Line: tok.Line, Column: tok.Column + from,
			Offset: tok.Offset + from, EndOffset: tok.Offset + to,
		}
		parts = append(parts, m.dedicated(air.KindPyStringFragment,
			part, part, text[from:to], note))
	}
	appendPart(0, closeStart, "py-string-content")
	appendPart(closeStart, len(text), "py-string-delimiter")
	return parts
}

// stringFragment lowers one literal text segment inside an f-string as a
// dedicated string-fragment node. It is called while the f-string literal is
// reduced and anchored on the segment's real span inside the token.
func (m *pyMap) stringFragment(tok frontkit.Token, from, to int) air.Node {
	part := frontkit.Token{
		Line: tok.Line, Column: tok.Column + from,
		Offset: tok.Offset + from, EndOffset: tok.Offset + to,
	}
	return m.dedicated(air.KindPyStringFragment, part, part,
		tok.Text[from:to], "py-string-content")
}

// typeConversion detects the !r/!s/!a conversion marker of one f-string
// interpolation. It is called with the exact interpolation text, not with a
// free-floating token, so the marker is read in its real syntactic context.
func (m *pyMap) typeConversion(tok frontkit.Token, at, end int,
	inner string) air.Node {
	mark := strings.IndexByte(inner, '!')
	if mark < 0 || mark+1 >= len(inner) {
		return nil
	}
	conv := inner[mark+1]
	if conv != 'r' && conv != 's' && conv != 'a' {
		return nil
	}
	startAt := at + 1 + mark
	stopAt := startAt + 2
	part := frontkit.Token{
		Line: tok.Line, Column: tok.Column + startAt,
		Offset: tok.Offset + startAt, EndOffset: tok.Offset + stopAt,
	}
	if part.EndOffset > tok.Offset+end {
		part.EndOffset = tok.Offset + end
	}
	return m.dedicated(air.KindPyTypeConversion, part, part,
		inner[mark:mark+2], "py-type-conversion")
}

// importPrefix models the leading relative-import dots of a from-import as a
// dedicated node. It is invoked only while a from statement is reduced.
func (m *pyMap) importPrefix(tok frontkit.Token, count int) air.Node {
	if count <= 0 {
		return nil
	}
	part := tok
	part.Column = tok.Column + count
	part.Offset = tok.Offset + count
	part.EndOffset = tok.Offset + count
	part.Text = ""
	return m.dedicated(air.KindPyImportPrefix, tok, part,
		strings.Repeat(".", count), "py-import-prefix")
}

// withItem wraps one context manager expression of a with statement.
func (m *pyMap) withItem(item air.Node, start frontkit.Token) air.Node {
	pos := item.Pos()
	end := frontkit.Token{
		Line: pos.Line, Column: pos.Column,
		Offset: pos.Offset, EndOffset: pos.EndOffset,
	}
	return m.dedicated(air.KindPyWithItem, start, end,
		"", "py-with-item", item)
}

// scopeStatement models a global/nonlocal declaration with a dedicated kind.
func (m *pyMap) scopeStatement(tok frontkit.Token, end frontkit.Token,
	children ...air.Node) air.Node {
	return m.dedicated(air.KindPyScopeDecl, tok, end, tok.Text,
		"py-scope-decl", children...)
}
