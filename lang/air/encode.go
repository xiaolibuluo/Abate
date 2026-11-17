/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"crypto/sha256"
	"encoding/binary"
	"slices"
	"strings"
)

// Encoder serializes a Unit into a versioned, checksummed AIR archive.
//
// An archive is the ABAT magic, a u16 writer version, a u16 writer minimum, a
// sequence of length-prefixed sections, and a SHA-256 checksum of every
// preceding byte. The encoder is stateless and safe for concurrent use.
type Encoder struct {
	version Version
}

// NewEncoder returns an encoder that writes the current schema version.
func NewEncoder() Encoder {
	return Encoder{version: NewVersion()}
}

// NewEncoderAt returns an encoder that writes an explicit version range.
func NewEncoderAt(version Version) Encoder {
	return Encoder{version: version}
}

// Version returns the version range used by this encoder.
func (e Encoder) Version() Version {
	return e.version
}

// Encode validates the unit and produces a complete archive.
func (e Encoder) Encode(u *Unit) ([]byte, error) {
	if u == nil {
		return nil, ErrNilUnit
	}
	if err := u.Validate(); err != nil {
		return nil, err
	}
	body := encodeSections(u, e.version)
	sum := sha256.Sum256(body)
	out := make([]byte, 0, len(body)+checksumSize)
	out = append(out, body...)
	return append(out, sum[:]...), nil
}

// encodeSections builds every byte before the trailing checksum.
//
// It is intentionally infallible: Unit.Validate rejects malformed inputs before
// encoding, and the node walk is iterative so no tree shape can panic here.
func encodeSections(u *Unit, version Version) []byte {
	table := newStringTable()
	layout := buildLayout(u, table, version)
	table.freeze()

	out := make([]byte, 0, 128+len(layout.nodes)*48)
	out = append(out, archiveMagic...)
	out = binary.BigEndian.AppendUint16(out, version.current)
	out = binary.BigEndian.AppendUint16(out, version.min)
	out = appendSection(out, sectionUnit, buildUnitSection(u, layout, table))
	out = appendSection(out, sectionNodes, buildNodeSection(layout, table))
	out = appendSection(out, sectionStrs, buildStringSection(table))
	out = appendSection(out, sectionSem, buildSemSection(layout, table))
	if version.current >= 2 {
		out = appendSection(out, sectionExt, buildExtensionSection(u, table))
	}
	return out
}

// encodeLayout is the deterministic flattening of a unit's node forest.
type encodeLayout struct {
	nodes    []Node
	children [][]uint32
	roots    []uint32
	order    []int
}

// buildLayout flattens roots in path order and interns every string used.
func buildLayout(u *Unit, table *stringTable, version Version) encodeLayout {
	layout := encodeLayout{order: make([]int, len(u.Files))}
	for i := range layout.order {
		layout.order[i] = i
	}
	slices.SortFunc(layout.order, func(a, b int) int {
		return strings.Compare(u.Files[a].Path, u.Files[b].Path)
	})

	table.add(u.Package.Name)
	for _, path := range u.Package.Imports {
		table.add(path)
	}
	for _, file := range u.Files {
		table.add(file.Path)
		table.add(file.LineEnding)
		table.add(file.Encoding)
	}
	if version.current >= 2 {
		for key, value := range u.Annotations {
			table.add(key)
			table.add(value)
		}
	}

	for _, at := range layout.order {
		layout.roots = append(layout.roots,
			appendTree(u.Roots[at], &layout.nodes, &layout.children))
	}
	for _, current := range layout.nodes {
		table.add(AttributeOf(current))
		sem := SemanticsOf(current)
		for _, key := range sem.Keys() {
			table.add(sem[key])
		}
	}
	return layout
}

// appendTree assigns pre-order indices to a subtree without recursion.
func appendTree(root Node, nodes *[]Node, children *[][]uint32) uint32 {
	rootIndex := toU32(len(*nodes))
	*nodes = append(*nodes, root)
	*children = append(*children, nil)

	type frame struct {
		parent   uint32
		children []Node
		next     int
	}
	stack := []frame{{parent: rootIndex, children: root.Children(), next: 0}}
	for len(stack) > 0 {
		top := &stack[len(stack)-1]
		if top.next >= len(top.children) {
			stack = stack[:len(stack)-1]
			continue
		}
		child := top.children[top.next]
		top.next++
		if child == nil {
			continue
		}
		childIndex := toU32(len(*nodes))
		*nodes = append(*nodes, child)
		*children = append(*children, nil)
		(*children)[top.parent] = append((*children)[top.parent], childIndex)
		stack = append(stack, frame{
			parent:   childIndex,
			children: child.Children(),
			next:     0,
		})
	}
	return rootIndex
}

// buildUnitSection writes package metadata and the path-ordered file table.
func buildUnitSection(u *Unit, layout encodeLayout, table *stringTable) []byte {
	out := binary.BigEndian.AppendUint32(nil, table.id(u.Package.Name))
	imports := slices.Clone(u.Package.Imports)
	slices.Sort(imports)
	out = binary.BigEndian.AppendUint32(out, toU32(len(imports)))
	for _, path := range imports {
		out = binary.BigEndian.AppendUint32(out, table.id(path))
	}
	out = binary.BigEndian.AppendUint32(out, toU32(len(layout.order)))
	for at, fileIndex := range layout.order {
		file := u.Files[fileIndex]
		out = binary.BigEndian.AppendUint32(out, table.id(file.Path))
		out = append(out, file.ContentHash[:]...)
		out = binary.BigEndian.AppendUint32(out, table.id(file.LineEnding))
		out = binary.BigEndian.AppendUint32(out, table.id(file.Encoding))
		out = binary.BigEndian.AppendUint32(out, toU32(file.Size))
		out = binary.BigEndian.AppendUint32(out, layout.roots[at])
	}
	return out
}

// buildNodeSection writes the flat node table in pre-order.
func buildNodeSection(layout encodeLayout, table *stringTable) []byte {
	out := binary.BigEndian.AppendUint32(nil, toU32(len(layout.nodes)))
	for i, current := range layout.nodes {
		pos, end := current.Pos(), current.End()
		out = binary.BigEndian.AppendUint64(out, uint64(IDOf(current)))
		out = binary.BigEndian.AppendUint16(out, uint16(current.Kind()))
		out = binary.BigEndian.AppendUint32(out, toU32(pos.Line))
		out = binary.BigEndian.AppendUint32(out, toU32(pos.Column))
		out = binary.BigEndian.AppendUint32(out, toU32(pos.Offset))
		out = binary.BigEndian.AppendUint32(out, toU32(pos.EndOffset))
		out = binary.BigEndian.AppendUint32(out, toU32(end.Line))
		out = binary.BigEndian.AppendUint32(out, toU32(end.Column))
		out = binary.BigEndian.AppendUint32(out, table.id(AttributeOf(current)))
		out = binary.BigEndian.AppendUint32(out, toU32(len(layout.children[i])))
		for _, child := range layout.children[i] {
			out = binary.BigEndian.AppendUint32(out, child)
		}
	}
	return out
}

// buildStringSection writes the sorted, deduplicated string table.
func buildStringSection(table *stringTable) []byte {
	out := binary.BigEndian.AppendUint32(nil, toU32(len(table.order)))
	for _, value := range table.order {
		out = binary.BigEndian.AppendUint32(out, toU32(len(value)))
		out = append(out, value...)
	}
	return out
}

// buildSemSection writes semantic annotations as sorted global triples.
func buildSemSection(layout encodeLayout, table *stringTable) []byte {
	records := make([]byte, 0, len(layout.nodes)*8)
	count := uint32(0)
	for i, current := range layout.nodes {
		sem := SemanticsOf(current)
		for _, key := range sem.Keys() {
			records = binary.BigEndian.AppendUint32(records, toU32(i))
			records = binary.BigEndian.AppendUint16(records, uint16(key))
			records = binary.BigEndian.AppendUint32(records, table.id(sem[key]))
			count++
		}
	}
	out := binary.BigEndian.AppendUint32(nil, count)
	return append(out, records...)
}

// buildExtensionSection writes the version 2 additive unit annotations.
func buildExtensionSection(u *Unit, table *stringTable) []byte {
	out := binary.BigEndian.AppendUint32(nil, toU32(len(u.Annotations)))
	for _, key := range u.AnnotationKeys() {
		out = binary.BigEndian.AppendUint32(out, table.id(key))
		out = binary.BigEndian.AppendUint32(out, table.id(u.Annotations[key]))
	}
	return out
}
