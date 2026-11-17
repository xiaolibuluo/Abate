/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"encoding/binary"
	"slices"
)

// Wire format constants. Sections are length-prefixed so a decoder can skip
// any section identifier introduced by a newer additive schema version.
const (
	archiveMagic  = "ABAT"
	headerSize    = 8
	checksumSize  = 32
	sectionUnit   = byte(1)
	sectionNodes  = byte(2)
	sectionStrs   = byte(3)
	sectionSem    = byte(4)
	sectionExt    = byte(5)
	maxTreeDepth  = 512
	minNodeBytes  = 8 + 2 + 6*4 + 4 + 4
	unknownTailID = byte(0x7F)
)

// stringTable interns every string used by an archive, then freezes into a
// lexicographically sorted, deduplicated lookup table.
type stringTable struct {
	seen  map[string]struct{}
	order []string
	index map[string]uint32
}

// newStringTable returns an empty interning table.
func newStringTable() *stringTable {
	return &stringTable{
		seen:  make(map[string]struct{}),
		order: []string{},
		index: make(map[string]uint32),
	}
}

// add records a string for later inclusion in the table.
func (t *stringTable) add(value string) {
	t.seen[value] = struct{}{}
}

// freeze sorts the collected strings and assigns their final indices.
func (t *stringTable) freeze() {
	t.order = make([]string, 0, len(t.seen))
	for value := range t.seen {
		t.order = append(t.order, value)
	}
	slices.Sort(t.order)
	for i, value := range t.order {
		t.index[value] = toU32(i)
	}
}

// id returns the frozen index of a previously added string.
func (t *stringTable) id(value string) uint32 {
	if at, ok := t.index[value]; ok {
		return at
	}
	return 0
}

// appendSection wraps a payload in a one-byte identifier and a u32 length.
func appendSection(out []byte, id byte, payload []byte) []byte {
	out = append(out, id)
	out = binary.BigEndian.AppendUint32(out, toU32(len(payload)))
	return append(out, payload...)
}

// toU32 clamps a signed integer into the unsigned wire domain.
func toU32(value int) uint32 {
	if value < 0 {
		return 0
	}
	if uint64(value) > uint64(^uint32(0)) {
		return ^uint32(0)
	}
	return uint32(value) //nolint:gosec // the range check above bounds the conversion.
}

// toU64 clamps a signed integer into the unsigned wire domain.
func toU64(value int) uint64 {
	if value < 0 {
		return 0
	}
	return uint64(value) //nolint:gosec // negative values are clamped to zero above.
}
