/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "fmt"

// Position locates a node in its source file.
//
// Line and Column are one-based, Offset is a zero-based byte offset, and
// EndOffset is the exclusive end of the half-open interval [Offset, EndOffset).
// A zero Line denotes an unknown source line, which is legal for synthetic
// nodes produced by downstream passes.
type Position struct {
	// Line is the one-based start line of the node.
	Line int
	// Column is the one-based start column of the node in bytes.
	Column int
	// Offset is the zero-based byte offset of the first byte of the node.
	Offset int
	// EndOffset is the exclusive byte offset just past the node.
	EndOffset int
}

// NewPosition builds a position from its four canonical components.
func NewPosition(line, column, offset, endOffset int) Position {
	return Position{
		Line:      line,
		Column:    column,
		Offset:    offset,
		EndOffset: endOffset,
	}
}

// Valid reports whether the position describes a non-negative, ordered span.
func (p Position) Valid() bool {
	return p.Line >= 0 && p.Column >= 0 && p.Offset >= 0 && p.EndOffset >= p.Offset
}

// Contains reports whether the byte interval of other lies inside this one.
func (p Position) Contains(other Position) bool {
	return p.Offset <= other.Offset && other.EndOffset <= p.EndOffset &&
		other.Offset <= other.EndOffset
}

// Overlaps reports whether two byte intervals share at least one byte.
func (p Position) Overlaps(other Position) bool {
	return p.Offset < other.EndOffset && other.Offset < p.EndOffset
}

// String renders the position in a stable, human-readable form.
func (p Position) String() string {
	return fmt.Sprintf("%d:%d[%d,%d)", p.Line, p.Column, p.Offset, p.EndOffset)
}

// Canonical returns the position as a deterministic comparison key.
func (p Position) Canonical() string {
	return fmt.Sprintf("%d/%d/%d/%d", p.Line, p.Column, p.Offset, p.EndOffset)
}
