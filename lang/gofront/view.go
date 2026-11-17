/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"go/ast"
	"go/token"
	"go/types"
	"slices"
	"strings"

	"abate/lang/air"
)

// GoView is an opaque view of one parsed Go file.
//
// It holds the upstream syntax tree, its file set, and the type information
// gathered by the in-memory checker. All fields and all methods that mention
// go/ast or go/token are unexported, so upstream AST types cannot leak into
// the AIR-facing public API.
type GoView struct {
	file     *ast.File
	fset     *token.FileSet
	info     *types.Info
	src      []byte
	bindings map[ast.Node]air.Node
}

// newGoView creates a view for one parsed file.
func newGoView(file *ast.File, fset *token.FileSet, src []byte) *GoView {
	return &GoView{
		file:     file,
		fset:     fset,
		src:      src,
		bindings: make(map[ast.Node]air.Node),
	}
}

// PackageName returns the declared package name of the viewed file.
func (v *GoView) PackageName() string {
	if v == nil || v.file == nil || v.file.Name == nil {
		return ""
	}
	return v.file.Name.Name
}

// NodeCount returns the number of syntax nodes bound to AIR nodes so far.
func (v *GoView) NodeCount() int {
	if v == nil {
		return 0
	}
	return len(v.bindings)
}

// Imports returns the sorted, deduplicated import paths of the viewed file.
func (v *GoView) Imports() []string {
	if v == nil || v.file == nil || len(v.file.Imports) == 0 {
		return nil
	}
	paths := make([]string, 0, len(v.file.Imports))
	for _, spec := range v.file.Imports {
		if spec.Path != nil {
			paths = append(paths, strings.Trim(spec.Path.Value, "`\""))
		}
	}
	slices.Sort(paths)
	return slices.Compact(paths)
}

// position converts a token position into an AIR start position.
func (v *GoView) position(pos token.Pos) air.Position {
	if v == nil || v.fset == nil || !pos.IsValid() {
		return air.NewPosition(0, 0, 0, 0)
	}
	at := v.fset.Position(pos)
	return air.NewPosition(at.Line, at.Column, at.Offset, at.Offset)
}

// span converts a token interval into an AIR byte interval.
func (v *GoView) span(pos, end token.Pos) air.Position {
	start := v.position(pos)
	if v == nil || v.fset == nil || !end.IsValid() {
		start.EndOffset = start.Offset
		return start
	}
	finish := v.fset.Position(end)
	if finish.Offset >= start.Offset {
		start.EndOffset = finish.Offset
	}
	return start
}

// spanAt converts one token interval into explicit AIR start and end positions.
func (v *GoView) spanAt(pos, end token.Pos) (start, endPos air.Position) {
	start = v.span(pos, end)
	endPos = air.NewPosition(start.Line, start.Column, start.EndOffset, start.EndOffset)
	if v != nil && v.fset != nil && end.IsValid() {
		finish := v.fset.Position(end)
		endPos = air.NewPosition(finish.Line, finish.Column,
			start.EndOffset, start.EndOffset)
	}
	return start, endPos
}

// newNode builds an AIR node with exact start and end positions.
func (v *GoView) newNode(kind air.NodeKind, pos, end token.Pos,
	text string, children ...air.Node) air.Node {
	start, endPos := v.spanAt(pos, end)
	return air.NewNodeAt(kind, start, endPos, text, children...)
}

// synthetic builds a zero-width AIR node used when a syntax node is absent.
func (v *GoView) synthetic(kind air.NodeKind, text string) air.Node {
	return v.newNode(kind, token.NoPos, token.NoPos, text)
}

// bind records the AIR node produced for one upstream syntax node.
func (v *GoView) bind(key ast.Node, value air.Node) {
	if v == nil || isNilASTNode(key) || value == nil {
		return
	}
	v.bindings[key] = value
}

// lookup finds the AIR node bound to an upstream syntax node.
func (v *GoView) lookup(key ast.Node) (air.Node, bool) {
	if v == nil || isNilASTNode(key) {
		return nil, false
	}
	value, ok := v.bindings[key]
	return value, ok
}

// appendNode appends a non-nil node to a child list.
func appendNode(children []air.Node, value air.Node) []air.Node {
	if value == nil {
		return children
	}
	return append(children, value)
}

// orderChildren sorts siblings by their byte offset, keeping equal offsets in
// their original order so comment and declaration nodes never overlap.
func orderChildren(children []air.Node) []air.Node {
	slices.SortStableFunc(children, func(a, b air.Node) int {
		left, right := a.Pos().Offset, b.Pos().Offset
		switch {
		case left < right:
			return -1
		case left > right:
			return 1
		default:
			return 0
		}
	})
	return children
}
