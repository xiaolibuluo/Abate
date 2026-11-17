/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import "abate/lang/air"

// Token is one deterministic lexical unit produced from source bytes. It holds
// no time, random or global-state field, so a token stream is a pure value.
type Token struct {
	// Kind is the lexical category of the token.
	Kind TokenKind
	// Text is the exact source slice covered by the token.
	Text string
	// Offset is the zero-based byte offset of the first byte.
	Offset int
	// EndOffset is the exclusive byte offset just past the token.
	EndOffset int
	// Line is the one-based source line of the token.
	Line int
	// Column is the one-based byte column of the token.
	Column int
}

// Span converts the token into the start and end positions used by air.Position.
// Line and Column are one-based, Offset is zero-based, and the byte interval is
// the half-open range [Offset, EndOffset). The end position mirrors
// air.NewNode: its offsets collapse to EndOffset while line and column repeat
// the start, so the projection is byte-exact and independent of the source text.
func (t Token) Span() (start, end air.Position) {
	start = air.NewPosition(t.Line, t.Column, t.Offset, t.EndOffset)
	end = air.NewPosition(t.Line, t.Column, t.EndOffset, t.EndOffset)
	return start, end
}

// Valid reports whether the token has a known kind and an ordered, non-negative
// span with one-based line and column coordinates.
func (t Token) Valid() bool {
	return t.Kind.Valid() && t.Offset >= 0 && t.EndOffset >= t.Offset &&
		t.Line >= 1 && t.Column >= 1
}
