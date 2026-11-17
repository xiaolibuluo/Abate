/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "encoding/binary"

// cursor is a bounds-checked reader that remembers its first failure.
//
// Every read is a no-op after an error, so malformed archives can never panic
// regardless of the byte patterns a fuzzer supplies.
type cursor struct {
	data []byte
	off  int
	err  error
}

// remaining returns the number of unread bytes, or zero after a failure.
func (c *cursor) remaining() int {
	if c.err != nil {
		return 0
	}
	return len(c.data) - c.off
}

// take consumes exactly n bytes, failing on any out-of-range request.
func (c *cursor) take(n int) []byte {
	if c.err != nil {
		return nil
	}
	if n < 0 || n > c.remaining() {
		c.err = ErrTruncatedArchive
		return nil
	}
	out := c.data[c.off : c.off+n]
	c.off += n
	return out
}

// u8 reads one byte.
func (c *cursor) u8() byte {
	b := c.take(1)
	if b == nil {
		return 0
	}
	return b[0]
}

// u16 reads a big-endian uint16.
func (c *cursor) u16() uint16 {
	b := c.take(2)
	if b == nil {
		return 0
	}
	return binary.BigEndian.Uint16(b)
}

// u32 reads a big-endian uint32.
func (c *cursor) u32() uint32 {
	b := c.take(4)
	if b == nil {
		return 0
	}
	return binary.BigEndian.Uint32(b)
}

// u64 reads a big-endian uint64.
func (c *cursor) u64() uint64 {
	b := c.take(8)
	if b == nil {
		return 0
	}
	return binary.BigEndian.Uint64(b)
}

// lookup resolves a string table index and records an invalid index.
func lookup(c *cursor, strings []string, index uint32) string {
	if int(index) >= len(strings) {
		c.err = ErrIndexOutOfRange
		return ""
	}
	return strings[index]
}

// decodeStrings reads the sorted string table.
func decodeStrings(payload []byte) ([]string, error) {
	if len(payload) < 4 {
		return nil, ErrMalformedStrings
	}
	c := &cursor{data: payload}
	count := c.u32()
	if uint64(count) > toU64(len(payload)/4)+1 {
		return nil, ErrMalformedStrings
	}
	out := make([]string, 0, count)
	for i := uint32(0); i < count; i++ {
		size := c.u32()
		raw := c.take(int(size))
		if c.err != nil {
			return nil, c.err
		}
		out = append(out, string(raw))
	}
	return out, nil
}

// decodeNodes reads the flat node table, links children, and bounds the depth.
func decodeNodes(payload []byte, strings []string) ([]*node, error) {
	if len(payload) < 4 {
		return nil, ErrMalformedNode
	}
	c := &cursor{data: payload}
	count := c.u32()
	if uint64(count) > toU64(len(payload)/minNodeBytes)+1 {
		return nil, ErrMalformedNode
	}
	nodes := make([]*node, 0, count)
	children := make([][]uint32, 0, count)
	for i := uint32(0); i < count; i++ {
		id, kind := c.u64(), NodeKind(c.u16())
		line, column := c.u32(), c.u32()
		offset, endOffset := c.u32(), c.u32()
		endLine, endColumn := c.u32(), c.u32()
		text := lookup(c, strings, c.u32())
		childCount := c.u32()
		if c.err != nil {
			return nil, c.err
		}
		if !kind.Valid() {
			return nil, ErrMalformedNode
		}
		if uint64(childCount) > toU64(c.remaining()/4) {
			return nil, ErrMalformedNode
		}
		kids := make([]uint32, 0, childCount)
		for j := uint32(0); j < childCount; j++ {
			child := c.u32()
			if child >= count || child <= i {
				return nil, ErrMalformedNode
			}
			kids = append(kids, child)
		}
		nodes = append(nodes, &node{
			id:   NodeID(id),
			kind: kind,
			pos:  NewPosition(int(line), int(column), int(offset), int(endOffset)),
			end:  NewPosition(int(endLine), int(endColumn), int(endOffset), int(endOffset)),
			text: text,
			sem:  NewSemInfo(),
		})
		children = append(children, kids)
	}
	for i, kids := range children {
		for _, child := range kids {
			nodes[i].children = append(nodes[i].children, nodes[child])
		}
	}
	if err := checkDepth(children); err != nil {
		return nil, err
	}
	return nodes, nil
}

// checkDepth rejects a forest deeper than the recursive consumers allow.
func checkDepth(children [][]uint32) error {
	depth := make([]uint32, len(children))
	for i, kids := range children {
		for _, child := range kids {
			candidate := depth[i] + 1
			if candidate > depth[child] {
				depth[child] = candidate
			}
			if depth[child] > maxTreeDepth {
				return ErrTreeTooDeep
			}
		}
	}
	return nil
}

// decodedUnit groups the metadata recovered from the unit section.
type decodedUnit struct {
	pkg   PackageRef
	files []FileRef
	roots []Node
}

// decodeUnit reads package metadata and aligns files with their root nodes.
func decodeUnit(payload []byte, strings []string, nodes []*node) (decodedUnit, error) {
	if len(payload) < 12 {
		return decodedUnit{}, ErrMalformedUnit
	}
	c := &cursor{data: payload}
	name := lookup(c, strings, c.u32())
	importCount := c.u32()
	if uint64(importCount) > toU64(len(payload)/4)+1 {
		return decodedUnit{}, ErrMalformedUnit
	}
	imports := make([]string, 0, importCount)
	for i := uint32(0); i < importCount; i++ {
		imports = append(imports, lookup(c, strings, c.u32()))
	}
	fileCount := c.u32()
	if uint64(fileCount) > toU64(len(payload)/52)+1 {
		return decodedUnit{}, ErrMalformedUnit
	}
	files := make([]FileRef, 0, fileCount)
	roots := make([]Node, 0, fileCount)
	for i := uint32(0); i < fileCount; i++ {
		path := lookup(c, strings, c.u32())
		hash := c.take(32)
		ending := lookup(c, strings, c.u32())
		encoding := lookup(c, strings, c.u32())
		size, rootIndex := c.u32(), c.u32()
		if c.err != nil {
			return decodedUnit{}, c.err
		}
		if int(rootIndex) >= len(nodes) {
			return decodedUnit{}, ErrIndexOutOfRange
		}
		file := FileRef{Path: path, LineEnding: ending, Encoding: encoding, Size: int(size)}
		copy(file.ContentHash[:], hash)
		files = append(files, file)
		roots = append(roots, nodes[rootIndex])
	}
	if c.err != nil {
		return decodedUnit{}, c.err
	}
	return decodedUnit{
		pkg:   PackageRef{Name: name, Imports: imports},
		files: files,
		roots: roots,
	}, nil
}

// decodeSemantics attaches annotation triples, skipping unknown future keys.
func decodeSemantics(payload []byte, strings []string, nodes []*node) error {
	if len(payload) < 4 {
		return ErrMalformedSemantics
	}
	c := &cursor{data: payload}
	count := c.u32()
	if uint64(count) > toU64(len(payload)/10)+1 {
		return ErrMalformedSemantics
	}
	for i := uint32(0); i < count; i++ {
		nodeIndex, key := c.u32(), SemKey(c.u16())
		value := lookup(c, strings, c.u32())
		if c.err != nil || int(nodeIndex) >= len(nodes) {
			return ErrMalformedSemantics
		}
		if !key.Valid() {
			continue
		}
		nodes[nodeIndex].sem[key] = value
	}
	return nil
}

// decodeExtension reads the version 2 additive unit annotations.
func decodeExtension(payload []byte, strings []string) (map[string]string, error) {
	if len(payload) < 4 {
		return nil, ErrMalformedUnit
	}
	c := &cursor{data: payload}
	count := c.u32()
	if uint64(count) > toU64(len(payload)/8)+1 {
		return nil, ErrMalformedUnit
	}
	out := make(map[string]string, count)
	for i := uint32(0); i < count; i++ {
		key := lookup(c, strings, c.u32())
		value := lookup(c, strings, c.u32())
		if c.err != nil {
			return nil, c.err
		}
		out[key] = value
	}
	return out, nil
}
