/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"testing"
)

// TestNodeKindsExercise constructs, annotates, and reads every allocated kind.
func TestNodeKindsExercise(t *testing.T) {
	t.Parallel()
	categories := map[string]int{}
	for kind := KindInvalid; kind <= kindLast; kind++ {
		n := NewNode(kind, NewPosition(1, 1, 0, 1), "x")
		if n == nil || n.Kind() != kind || AttributeOf(n) != "x" {
			t.Fatalf("kind %d construction failed", kind)
		}
		if !Annotate(n, SemPurity, "pure") || n.Sem(SemPurity) != "pure" {
			t.Fatalf("kind %d annotation failed", kind)
		}
		if n.Sem(SemScopeID) != nil || SemanticsOf(n).Len() != 1 {
			t.Fatalf("kind %d semantics leak", kind)
		}
		if !IDOf(n).Valid() || len(n.Children()) != 0 {
			t.Fatalf("kind %d id or children failed", kind)
		}
		if kind.Valid() != (kind != KindInvalid) || kind.String() == "" {
			t.Fatalf("kind %d validity or name failed", kind)
		}
		categories[kind.CategoryName()]++
	}
	for _, name := range []string{"declaration", "statement", "expression", "type", "other", "invalid"} {
		if categories[name] == 0 {
			t.Fatalf("category %q has no kinds: %v", name, categories)
		}
	}
	if NodeKind(9999).String() != "Invalid" {
		t.Fatal("out-of-range kind must not produce a name")
	}
}

// TestNodeIDStability proves IDs ignore positions, include attributes, and
// preserve the order of children.
func TestNodeIDStability(t *testing.T) {
	t.Parallel()
	first := NewNode(KindIdent, NewPosition(1, 1, 0, 3), "abc")
	same := NewNode(KindIdent, NewPosition(9, 9, 40, 43), "abc")
	changed := NewNode(KindIdent, NewPosition(1, 1, 0, 3), "abd")
	if IDOf(first) != IDOf(same) {
		t.Fatal("structurally equal nodes must share an ID")
	}
	if IDOf(first) == IDOf(changed) {
		t.Fatal("a one-byte attribute change must change the ID")
	}
	left := NewNode(KindIdent, NewPosition(1, 1, 0, 1), "a")
	right := NewNode(KindIdent, NewPosition(1, 2, 1, 2), "b")
	if IDOf(NewNode(KindCall, NewPosition(1, 1, 0, 2), "", left, right)) ==
		IDOf(NewNode(KindCall, NewPosition(1, 1, 0, 2), "", right, left)) {
		t.Fatal("child order must affect the parent ID")
	}
	if IDOf(nil) != InvalidNodeID || InvalidNodeID.Valid() {
		t.Fatal("invalid ID contract broken")
	}
	if NewNode(KindCall, NewPosition(1, 1, 0, 1), "", nil).Children() != nil {
		t.Fatal("nil children must be filtered")
	}
}

// TestRoundTrip checks a full encode/decode cycle through both entry points.
func TestRoundTrip(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	archive := mustEncode(t, NewEncoder(), u)
	back := mustDecode(t, NewDecoder(), archive)
	if err := CompareUnits(u, back); err != nil {
		t.Fatalf("round trip: %v", err)
	}
	if back.Fingerprint() != u.Fingerprint() {
		t.Fatal("fingerprint changed across the round trip")
	}
	loaded, err := NewDecoder().Load(archive)
	if err != nil || CompareUnits(u, loaded) != nil {
		t.Fatalf("Load alias: %v", err)
	}
	body := archive[:len(archive)-checksumSize]
	sum := sha256.Sum256(body)
	if !bytes.Equal(sum[:], archive[len(archive)-checksumSize:]) {
		t.Fatal("archive checksum does not cover the whole body")
	}
}

// TestStringTableDedupSorted inspects the on-wire string table.
func TestStringTableDedupSorted(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoder(), mustUnit(t))
	payload := sectionPayload(t, archive, sectionStrs)
	if len(payload) < 4 {
		t.Fatal("missing string table")
	}
	cur := &cursor{data: payload}
	count := cur.u32()
	if count < 4 {
		t.Fatalf("string table too small: %d", count)
	}
	seen := false
	previous := ""
	for i := uint32(0); i < count; i++ {
		size := cur.u32()
		raw := cur.take(int(size))
		if cur.err != nil {
			t.Fatalf("string table read: %v", cur.err)
		}
		value := string(raw)
		if seen && value <= previous {
			t.Fatalf("string table not strictly sorted at %q after %q", value, previous)
		}
		previous, seen = value, true
	}
}

// TestChecksumDetectsEveryByteFlip flips every byte of an archive.
func TestChecksumDetectsEveryByteFlip(t *testing.T) {
	t.Parallel()
	archive := mustEncode(t, NewEncoder(), mustUnit(t))
	for i := range archive {
		mutated := append([]byte(nil), archive...)
		mutated[i] ^= 0x01
		unit, err := NewDecoder().Decode(mutated)
		if err == nil || unit != nil {
			t.Fatalf("flipping byte %d was accepted", i)
		}
		if !isAIRSentinel(err) {
			t.Fatalf("byte %d: unclassified error %v", i, err)
		}
		if !errors.Is(err, ErrCorruptChecksum) {
			t.Fatalf("byte %d: expected checksum sentinel, got %v", i, err)
		}
	}
}

// TestWalkCollectSpanIntervals exercises the traversal and interval helpers.
func TestWalkCollectSpanIntervals(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	nodes := Collect(u)
	if len(nodes) != u.NodeCount() || len(nodes) == 0 {
		t.Fatalf("Collect/NodeCount disagree: %d vs %d", len(nodes), u.NodeCount())
	}
	visited := 0
	Walk(u.Roots[0], func(_ Node) bool {
		visited++
		return false
	})
	if visited != 1 {
		t.Fatalf("pruning failed: visited %d", visited)
	}
	start, end := Span(u.Roots[0])
	if start != 0 || end != 3 {
		t.Fatalf("Span = [%d,%d)", start, end)
	}
	if err := CheckIntervals(u.Roots[0], 9); err != nil {
		t.Fatalf("valid intervals rejected: %v", err)
	}
	bad := NewNode(KindBlock, NewPosition(1, 1, 0, 5), "",
		NewNode(KindIdent, NewPosition(1, 1, 10, 12), "x"))
	if err := CheckIntervals(bad, 20); !errors.Is(err, ErrBadPosition) {
		t.Fatalf("escaping child not rejected: %v", err)
	}
	if err := CheckIntervals(nil, 10); err != nil {
		t.Fatalf("nil root: %v", err)
	}
}

// TestCompareUnitsDetectsChanges mutates one field at a time.
func TestCompareUnitsDetectsChanges(t *testing.T) {
	t.Parallel()
	base := mustUnit(t)
	if err := CompareUnits(base, base); err != nil {
		t.Fatalf("self comparison: %v", err)
	}
	if err := CompareUnits(nil, nil); err != nil {
		t.Fatalf("nil comparison: %v", err)
	}
	if err := CompareUnits(nil, base); !errors.Is(err, ErrUnitMismatch) {
		t.Fatalf("nil mismatch: %v", err)
	}
	copyUnit := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), base))
	Annotate(copyUnit.Roots[0], SemNote, "changed")
	if err := CompareUnits(base, copyUnit); !errors.Is(err, ErrUnitMismatch) {
		t.Fatal("semantic drift was not detected")
	}
	copyUnit = mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), base))
	copyUnit.Annotations["alpha"] = "2"
	if err := CompareUnits(base, copyUnit); !errors.Is(err, ErrUnitMismatch) {
		t.Fatal("annotation drift was not detected")
	}
	copyUnit = mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), base))
	copyUnit.Roots = append(copyUnit.Roots, nil)
	if err := CompareUnits(base, copyUnit); !errors.Is(err, ErrUnitMismatch) {
		t.Fatal("root count drift was not detected")
	}
}

// TestUnitFingerprintAndValidation covers empty units and invalid metadata.
func TestUnitFingerprintAndValidation(t *testing.T) {
	t.Parallel()
	var nilUnit *Unit
	if nilUnit.Fingerprint() != [32]byte{} || nilUnit.NodeCount() != 0 {
		t.Fatal("nil unit fingerprint or count")
	}
	if err := nilUnit.Validate(); !errors.Is(err, ErrNilUnit) {
		t.Fatalf("nil validate: %v", err)
	}
	empty := NewUnit(NewPackageRef("empty", nil))
	back := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), empty))
	if back.NodeCount() != 0 || len(back.Files) != 0 {
		t.Fatal("empty unit round trip changed shape")
	}
	root := NewNode(KindPackage, NewPosition(1, 1, 0, 1), "x")
	if err := empty.AddFile(NewFileRef("a.go", []byte("x")), root); err != nil {
		t.Fatalf("AddFile: %v", err)
	}
	if err := empty.AddFile(NewFileRef("a.go", []byte("y")), root); !errors.Is(err, ErrDuplicatePath) {
		t.Fatalf("duplicate path: %v", err)
	}
	if err := nilUnit.AddFile(NewFileRef("a.go", nil), root); !errors.Is(err, ErrNilUnit) {
		t.Fatalf("nil AddFile: %v", err)
	}
	if err := empty.AddFile(FileRef{}, root); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("empty FileRef: %v", err)
	}
	if err := empty.AddFile(NewFileRef("b.go", nil), nil); !errors.Is(err, ErrMalformedNode) {
		t.Fatalf("nil root: %v", err)
	}
	if _, err := NewEncoder().Encode(nil); !errors.Is(err, ErrNilUnit) {
		t.Fatalf("nil Encode: %v", err)
	}
}
