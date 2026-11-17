/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"slices"
	"strings"

	"abate/lang/air"
)

// View is the language-neutral node-building surface of one source buffer. It
// generalizes the Go frontend view: tokens replace upstream syntax nodes as the
// only positional input, so no language-specific AST type can leak into AIR.
// Bindings are keyed by integer token index rather than by a syntax node type.
type View struct {
	src         []byte
	path        string
	bindings    map[int]air.Node
	packageName string
	imports     []string
}

// NewView returns an independent view over a copy of src.
func NewView(src []byte, path string) *View {
	return &View{
		src: slices.Clone(src), path: path,
		bindings: make(map[int]air.Node), imports: []string{},
	}
}

// SourceFromTokens rebuilds a byte buffer with the same extent as the token
// stream, copying each token's text at its exact offset. Gaps stay zeroed; the
// buffer is used only for interval clamping, never for source regeneration. A
// language adapter that receives tokens but not raw bytes can use it to give
// NewView a correctly sized source extent.
func SourceFromTokens(tokens []Token) []byte {
	size := 0
	for _, tok := range tokens {
		if tok.Offset < 0 || tok.EndOffset < tok.Offset {
			continue
		}
		if tok.EndOffset > size {
			size = tok.EndOffset
		}
	}
	out := make([]byte, size)
	for _, tok := range tokens {
		if tok.Offset < 0 || tok.EndOffset > size || tok.EndOffset < tok.Offset {
			continue
		}
		copy(out[tok.Offset:tok.EndOffset], tok.Text)
	}
	return out
}

// Path returns the virtual source path of the view.
func (v *View) Path() string {
	if v == nil {
		return ""
	}
	return v.path
}

// Size returns the byte length of the viewed source buffer.
func (v *View) Size() int {
	if v == nil {
		return 0
	}
	return len(v.src)
}

// Position converts one token into an air.Position clamped to the source.
func (v *View) Position(tok Token) air.Position {
	if v == nil {
		return air.NewPosition(0, 0, 0, 0)
	}
	line := clampInt(tok.Line, 0, 1<<30)
	column := clampInt(tok.Column, 0, 1<<30)
	start := clampInt(tok.Offset, 0, len(v.src))
	end := clampInt(tok.EndOffset, start, len(v.src))
	return air.NewPosition(line, column, start, end)
}

// Span converts a token interval into one clamped air.Position.
func (v *View) Span(startToken, endToken Token) air.Position {
	span := v.Position(startToken)
	if v == nil {
		return span
	}
	span.EndOffset = clampInt(endToken.EndOffset, span.Offset, len(v.src))
	return span
}

// NewNodeAt builds an AIR node from a token interval and orders its children.
// The end position collapses to the interval's exclusive end offset, mirroring
// air.NewNode so byte-exact half-open spans are preserved.
func (v *View) NewNodeAt(kind air.NodeKind, startToken, endToken Token,
	text string, children ...air.Node) air.Node {
	if v == nil {
		return air.NewNode(kind, air.NewPosition(0, 0, 0, 0), text)
	}
	span := v.Span(startToken, endToken)
	line := clampInt(endToken.Line, 0, 1<<30)
	column := clampInt(endToken.Column, 0, 1<<30)
	end := air.NewPosition(line, column, span.EndOffset, span.EndOffset)
	return air.NewNodeAt(kind, span, end, text, v.OrderChildren(children)...)
}

// Synthetic builds a zero-width node marked as a partial reduction.
func (v *View) Synthetic(kind air.NodeKind, text string) air.Node {
	pos := v.Position(Token{})
	pos.EndOffset = pos.Offset
	node := air.NewNodeAt(kind, pos, pos, text)
	_ = air.Annotate(node, air.SemPartial, "true")
	return node
}

// Annotate attaches a semantic value to a node without changing its NodeID.
func (v *View) Annotate(node air.Node, key air.SemKey, value string) bool {
	if v == nil || node == nil || !key.Valid() {
		return false
	}
	return air.Annotate(node, key, value)
}

// Bind records the AIR node produced for one token index.
func (v *View) Bind(key int, value air.Node) {
	if v == nil || value == nil {
		return
	}
	if v.bindings == nil {
		v.bindings = make(map[int]air.Node)
	}
	v.bindings[key] = value
}

// Lookup finds the AIR node bound to one token index.
func (v *View) Lookup(key int) (air.Node, bool) {
	if v == nil || v.bindings == nil {
		return nil, false
	}
	node, ok := v.bindings[key]
	return node, ok
}

// NodeCount returns how many token indices are bound so far.
func (v *View) NodeCount() int {
	if v == nil {
		return 0
	}
	return len(v.bindings)
}

// OrderChildren sorts siblings by byte offset, drops nil and out-of-bounds
// children, and removes any child overlapping its predecessor. The result has
// no overlap and no interval beyond the source extent. Gaps between siblings
// are preserved because they hold keywords and separators that a neutral IR
// does not model as vertices; when children tile exactly, the ordered list has
// no gaps at all, which is the interval coverage contract exercised by tests.
func (v *View) OrderChildren(children []air.Node) []air.Node {
	bound := 0
	if v != nil {
		bound = len(v.src)
	}
	out := make([]air.Node, 0, len(children))
	for _, child := range children {
		if child == nil {
			continue
		}
		pos := child.Pos()
		if !pos.Valid() || pos.Offset < 0 || pos.EndOffset > bound {
			continue
		}
		out = append(out, child)
	}
	slices.SortStableFunc(out, func(a, b air.Node) int {
		if a.Pos().Offset != b.Pos().Offset {
			if a.Pos().Offset < b.Pos().Offset {
				return -1
			}
			return 1
		}
		if a.Pos().EndOffset != b.Pos().EndOffset {
			if a.Pos().EndOffset < b.Pos().EndOffset {
				return -1
			}
			return 1
		}
		return strings.Compare(air.AttributeOf(a), air.AttributeOf(b))
	})
	kept := make([]air.Node, 0, len(out))
	previous := 0
	for _, child := range out {
		if child.Pos().Offset < previous {
			continue
		}
		kept = append(kept, child)
		previous = child.Pos().EndOffset
	}
	return kept
}

// SetPackageName records the declared package name.
func (v *View) SetPackageName(name string) {
	if v != nil {
		v.packageName = name
	}
}

// PackageName returns the declared package name.
func (v *View) PackageName() string {
	if v == nil {
		return ""
	}
	return v.packageName
}

// AddImport records one import path.
func (v *View) AddImport(importPath string) {
	if v == nil || importPath == "" {
		return
	}
	v.imports = append(v.imports, importPath)
}

// Imports returns sorted, deduplicated import paths as a fresh slice.
func (v *View) Imports() []string {
	if v == nil || len(v.imports) == 0 {
		return []string{}
	}
	out := slices.Clone(v.imports)
	slices.Sort(out)
	return slices.Compact(out)
}

// PackageRef builds the AIR package metadata from the view.
func (v *View) PackageRef() air.PackageRef {
	if v == nil {
		return air.NewPackageRef("", nil)
	}
	return air.NewPackageRef(v.PackageName(), v.Imports())
}

// clampInt keeps one integer inside an inclusive interval.
func clampInt(value, low, high int) int {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}
