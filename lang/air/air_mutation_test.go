/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"errors"
	"testing"
)

// TestMutationEveryTruncation cuts an archive at every possible prefix length.
func TestMutationEveryTruncation(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoder(), mustUnit(t))
	for size := 0; size < len(archive); size++ {
		unit, err := NewDecoder().Decode(archive[:size])
		if err == nil || unit != nil {
			t.Fatalf("truncation to %d bytes was accepted", size)
		}
		if !isAIRSentinel(err) {
			t.Fatalf("truncation %d: unclassified error %v", size, err)
		}
	}
}

// TestMutationEmptyAndShortArchives covers the smallest malformed inputs.
func TestMutationEmptyAndShortArchives(t *testing.T) {
	t.Parallel()
	cases := [][]byte{
		{},
		[]byte("ABAT"),
		make([]byte, headerSize),
		make([]byte, headerSize+checksumSize),
		[]byte("XXXX\x00\x01\x00\x01"),
	}
	for i, data := range cases {
		unit, err := NewDecoder().Load(data)
		if err == nil || unit != nil {
			t.Fatalf("case %d accepted", i)
		}
		if !isAIRSentinel(err) {
			t.Fatalf("case %d: unclassified error %v", i, err)
		}
	}
}

// TestMutationChecksumFlip flips each byte of the trailing digest.
func TestMutationChecksumFlip(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoder(), mustUnit(t))
	for i := len(archive) - checksumSize; i < len(archive); i++ {
		mutated := append([]byte(nil), archive...)
		mutated[i] ^= 0x80
		if _, err := NewDecoder().Decode(mutated); !errors.Is(err, ErrCorruptChecksum) {
			t.Fatalf("checksum byte %d: %v", i, err)
		}
	}
}

// TestMutationBadMagicAndDuplicateSection checks framing-level mutations.
func TestMutationBadMagicAndDuplicateSection(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoder(), mustUnit(t))
	if _, err := NewDecoder().Decode(mutateByte(t, archive, 0, 'X')); !errors.Is(err, ErrBadMagic) {
		t.Fatalf("bad magic: %v", err)
	}
	body := archiveBody(archive)
	nodePayload := sectionPayload(t, archive, sectionNodes)
	body = append(body, appendSection(nil, sectionNodes, nodePayload)...)
	if _, err := NewDecoder().Decode(reseal(body)); !errors.Is(err, ErrDuplicateSection) {
		t.Fatalf("duplicate section: %v", err)
	}
	oversized := archiveBody(archive)
	for i := headerSize + 1; i <= headerSize+4; i++ {
		oversized[i] = 0xFF
	}
	if _, err := NewDecoder().Decode(reseal(oversized)); !errors.Is(err, ErrTruncatedArchive) {
		t.Fatalf("oversized section length: %v", err)
	}
}

// TestMutationNodeTableFields changes one node field at a time.
func TestMutationNodeTableFields(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	archive := mustEncode(t, NewEncoder(), u)
	nodesAt := sectionOffset(t, archive, sectionNodes)
	if nodesAt < 0 {
		t.Fatal("node section missing")
	}
	// A legal kind change must decode but fail the structural comparison.
	kindLow := nodesAt + 4 + 8 + 1
	changed := mustDecode(t, NewDecoder(), mutateByte(t, archive, kindLow, byte(KindIdent)))
	if err := CompareUnits(u, changed); !errors.Is(err, ErrUnitMismatch) {
		t.Fatalf("kind mutation not detected: %v", err)
	}
	// An illegal kind must be rejected by the loader.
	if _, err := NewDecoder().Decode(mutateByte(t, archive, kindLow-1, 0xFF)); !errors.Is(err, ErrMalformedNode) {
		t.Fatalf("illegal kind: %v", err)
	}
	// A different string index changes the node attribute.
	textLow := nodesAt + 4 + 8 + 2 + 6*4 + 3
	retexted := mustDecode(t, NewDecoder(), mutateByte(t, archive, textLow, 2))
	if err := CompareUnits(u, retexted); !errors.Is(err, ErrUnitMismatch) {
		t.Fatalf("attribute mutation not detected: %v", err)
	}
	// A changed position offset is still valid but must be detected.
	offsetLow := nodesAt + 4 + 8 + 2 + 4 + 4 + 3
	moved := mustDecode(t, NewDecoder(), mutateByte(t, archive, offsetLow, 1))
	if err := CompareUnits(u, moved); !errors.Is(err, ErrUnitMismatch) {
		t.Fatalf("position mutation not detected: %v", err)
	}
	// An enormous child count must be rejected before allocation.
	childHigh := nodesAt + 4 + 8 + 2 + 6*4 + 4
	if _, err := NewDecoder().Decode(mutateByte(t, archive, childHigh, 0xFF)); !errors.Is(err, ErrMalformedNode) {
		t.Fatalf("child count mutation: %v", err)
	}
	// An out-of-range string index must be reported as an index error.
	if _, err := NewDecoder().Decode(mutateByte(t, archive, textLow-3, 0xFF)); !errors.Is(err, ErrIndexOutOfRange) {
		t.Fatalf("string index mutation: %v", err)
	}
}

// TestMutationEveryKindNodeID requires an attribute mutation to change an ID.
func TestMutationEveryKindNodeID(t *testing.T) {
	t.Parallel()
	pos := NewPosition(1, 1, 0, 1)
	for kind := KindInvalid; kind <= kindLast; kind++ {
		if IDOf(NewNode(kind, pos, "x")) == IDOf(NewNode(kind, pos, "y")) {
			t.Fatalf("kind %d ignored an attribute mutation", kind)
		}
	}
	if got := ComputeNodeID(KindIdent, "", nil); !got.Valid() || got.String() == "" {
		t.Fatal("ComputeNodeID returned an invalid identifier")
	}
}

// TestMutationEverySemKey requires annotation mutations to survive the wire.
func TestMutationEverySemKey(t *testing.T) {
	t.Parallel()
	for key := SemTypeRef; key <= SemKeyLast; key++ {
		first, second := mustUnit(t), mustUnit(t)
		Annotate(first.Roots[0], key, "one")
		Annotate(second.Roots[0], key, "two")
		if err := CompareUnits(first, second); !errors.Is(err, ErrUnitMismatch) {
			t.Fatalf("key %d mutation not detected", key)
		}
		back := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), first))
		if back.Roots[0].Sem(key) != "one" {
			t.Fatalf("key %d lost its annotation", key)
		}
	}
	first := mustUnit(t)
	Annotate(first.Roots[0], SemInvalid, "ignored")
	back := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), first))
	if back.Roots[0].Sem(SemInvalid) != nil {
		t.Fatal("an unknown future annotation key must be skipped, not retained")
	}
}

// TestMutationSectionTables corrupts the string and semantic tables.
func TestMutationSectionTables(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoder(), mustUnit(t))
	strAt := sectionOffset(t, archive, sectionStrs)
	if _, err := NewDecoder().Decode(mutateByte(t, archive, strAt, 0xFF)); !errors.Is(err, ErrMalformedStrings) {
		t.Fatalf("string table count: %v", err)
	}
	badLen := archiveBody(archive)
	for i := strAt + 4; i <= strAt+7; i++ {
		badLen[i] = 0xFF
	}
	if _, err := NewDecoder().Decode(reseal(badLen)); !errors.Is(err, ErrTruncatedArchive) {
		t.Fatalf("string length: %v", err)
	}
	semAt := sectionOffset(t, archive, sectionSem)
	if _, err := NewDecoder().Decode(mutateByte(t, archive, semAt, 0xFF)); !errors.Is(err, ErrMalformedSemantics) {
		t.Fatalf("semantic table count: %v", err)
	}
	extAt := sectionOffset(t, archive, sectionExt)
	if _, err := NewDecoder().Decode(mutateByte(t, archive, extAt, 0xFF)); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("extension count: %v", err)
	}
}
