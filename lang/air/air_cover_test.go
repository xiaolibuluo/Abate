/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"testing"
)

// foreignNode is a Node implementation outside this package's constructor.
type foreignNode struct {
	kind NodeKind
	pos  Position
	end  Position
	kids []Node
}

// Kind returns the declared kind.
func (f foreignNode) Kind() NodeKind { return f.kind }

// Pos returns the declared interval.
func (f foreignNode) Pos() Position { return f.pos }

func (f foreignNode) End() Position    { return f.end }
func (f foreignNode) Children() []Node { return f.kids }
func (foreignNode) Sem(_ SemKey) any   { return nil }

// badJSONNode makes json.Marshal fail inside a parent's MarshalJSON.
type badJSONNode struct {
	foreignNode
	Fn func()
}

// TestCoverForeignNodePaths exercises the fallback paths for custom nodes.
func TestCoverForeignNodePaths(t *testing.T) {
	t.Parallel()
	foreign := foreignNode{kind: KindIdent, pos: NewPosition(1, 1, 0, 1)}
	if Annotate(foreign, SemPurity, "pure") || AttributeOf(foreign) != "" {
		t.Fatal("foreign node accepted package-internal operations")
	}
	if !IDOf(foreign).Valid() || SemanticsOf(foreign).Len() != 0 {
		t.Fatal("foreign node ID or semantics fallback failed")
	}
	parent := foreignNode{kind: KindCall, kids: []Node{foreign, nil}}
	if !IDOf(parent).Valid() {
		t.Fatal("foreign node child hashing failed")
	}
	visited := 0
	Walk(parent, func(_ Node) bool {
		visited++
		return true
	})
	if visited != 2 {
		t.Fatalf("Walk skipped a foreign child: %d", visited)
	}
	bad := NewNode(KindBlock, NewPosition(1, 1, 0, 1), "",
		badJSONNode{foreignNode: foreign, Fn: func() {}})
	if _, err := json.Marshal(bad); err == nil {
		t.Fatal("MarshalJSON must propagate a child encoding failure")
	}
	if start, end := Span(nil); start != 0 || end != 0 {
		t.Fatal("Span(nil) must be zero")
	}
	if err := compareNodes(nil, foreign); !errors.Is(err, ErrUnitMismatch) {
		t.Fatal("nil comparison must fail")
	}
}

// TestCoverUnitValidation rejects each malformed unit invariant.
func TestCoverUnitValidation(t *testing.T) {
	t.Parallel()
	root := NewNode(KindIdent, NewPosition(1, 1, 0, 1), "x")
	cases := []struct {
		name string
		unit *Unit
		want error
	}{
		{"bad package", NewUnit(NewPackageRef("p", []string{""})), ErrMalformedUnit},
		{"file count", &Unit{
			Package: NewPackageRef("p", nil),
			Files:   []FileRef{{Path: "a.go"}}, Roots: []Node{},
		}, ErrMalformedUnit},
		{"bad file", &Unit{
			Package: NewPackageRef("p", nil),
			Files:   []FileRef{{}}, Roots: []Node{root},
		}, ErrMalformedUnit},
		{"nil root", &Unit{
			Package: NewPackageRef("p", nil),
			Files:   []FileRef{{Path: "a.go"}}, Roots: []Node{nil},
		}, ErrMalformedNode},
		{"bad position", &Unit{
			Package: NewPackageRef("p", nil),
			Files:   []FileRef{{Path: "a.go"}},
			Roots:   []Node{NewNode(KindIdent, NewPosition(1, 1, 5, 4), "x")},
		}, ErrMalformedNode},
	}
	for _, tc := range cases {
		if err := tc.unit.Validate(); !errors.Is(err, tc.want) {
			t.Fatalf("%s: got %v", tc.name, err)
		}
		if _, err := NewEncoder().Encode(tc.unit); !errors.Is(err, tc.want) {
			t.Fatalf("%s Encode: %v", tc.name, err)
		}
	}
}

// TestCoverWireHelpers covers clamping, interning, and accessor boundaries.
func TestCoverWireHelpers(t *testing.T) {
	t.Parallel()
	if toU32(-1) != 0 || toU64(-1) != 0 || toU64(42) != 42 {
		t.Fatal("numeric clamping failed")
	}
	if toU32(int(^uint(0)>>1)) != ^uint32(0) {
		t.Fatal("large int was not clamped to uint32 maximum")
	}
	table := newStringTable()
	table.add("b")
	table.add("a")
	table.add("a")
	table.freeze()
	if table.id("a") != 0 || table.id("b") != 1 || table.id("missing") != 0 {
		t.Fatalf("string table ordering: %v", table.order)
	}
	if NewEncoder().Version() != NewVersion() || NewDecoder().Version() != NewVersion() {
		t.Fatal("codec version accessors drifted")
	}
	if NewEncoderAt(NewVersion()).Version().Current() != SchemaVersion {
		t.Fatal("NewEncoderAt ignored its argument")
	}
	cr := NewFileRef("x.go", []byte("a\rb"))
	if cr.LineEnding != "\r" || cr.Encoding != "utf-8" {
		t.Fatalf("carriage-return detection: %+v", cr)
	}
	file := NewSourceFile("x.go", []byte("abc"))
	if file.Digest() != NewFileRef("x.go", []byte("abc")).ContentHash {
		t.Fatal("SourceFile.Digest disagrees with FileRef")
	}
	tie := SourceUnit{Files: []SourceFile{
		NewSourceFile("a.go", []byte("x")), NewSourceFile("a.go", []byte("y")),
	}}
	if len(tie.SortedFiles()) != 2 {
		t.Fatal("equal paths must remain stable under sorting")
	}
}

// TestCoverCursorFailures checks the bounds-checked reader after a failure.
func TestCoverCursorFailures(t *testing.T) {
	t.Parallel()
	c := &cursor{data: []byte{1, 2, 3}}
	if c.u8() != 1 || c.u16() != (2<<8|3) {
		t.Fatal("cursor short reads failed")
	}
	_ = c.u32()
	if c.remaining() != 0 || c.u8() != 0 || c.u16() != 0 ||
		c.u32() != 0 || c.u64() != 0 {
		t.Fatal("cursor must stay inert after a failed read")
	}
	if _, err := readSections([]byte{sectionUnit, 0xFF, 0xFF, 0xFF, 0xFF}); !errors.Is(err, ErrTruncatedArchive) {
		t.Fatalf("framing overflow: %v", err)
	}
}

// TestCoverSectionDecoders calls every section decoder with malformed inputs.
func TestCoverSectionDecoders(t *testing.T) {
	t.Parallel()
	if _, err := decodeStrings(nil); !errors.Is(err, ErrMalformedStrings) {
		t.Fatalf("short strings: %v", err)
	}
	if _, err := decodeStrings([]byte{0xFF, 0xFF, 0xFF, 0xFF}); !errors.Is(err, ErrMalformedStrings) {
		t.Fatalf("string count: %v", err)
	}
	if _, err := decodeStrings([]byte{0, 0, 0, 1, 0, 0, 0, 9, 'x'}); !errors.Is(err, ErrTruncatedArchive) {
		t.Fatalf("string length: %v", err)
	}
	if _, err := decodeNodes(nil, []string{""}); !errors.Is(err, ErrMalformedNode) {
		t.Fatalf("short nodes: %v", err)
	}
	if _, err := decodeNodes([]byte{0xFF, 0xFF, 0xFF, 0xFF}, []string{""}); !errors.Is(err, ErrMalformedNode) {
		t.Fatalf("node count: %v", err)
	}
	if _, err := decodeNodes([]byte{0, 0, 0, 1, 0}, []string{""}); !errors.Is(err, ErrTruncatedArchive) {
		t.Fatalf("truncated node record: %v", err)
	}
	if err := decodeSemantics(nil, []string{""}, nil); !errors.Is(err, ErrMalformedSemantics) {
		t.Fatalf("short semantics: %v", err)
	}
	if err := decodeSemantics([]byte{0xFF, 0, 0, 0}, []string{""}, nil); !errors.Is(err, ErrMalformedSemantics) {
		t.Fatalf("semantic count: %v", err)
	}
	if err := decodeSemantics([]byte{0, 0, 0, 1, 0}, []string{""}, nil); !errors.Is(err, ErrMalformedSemantics) {
		t.Fatalf("truncated semantic record: %v", err)
	}
	if _, err := decodeExtension(nil, []string{""}); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("short extension: %v", err)
	}
	if _, err := decodeExtension([]byte{0xFF, 0, 0, 0}, []string{""}); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("extension count: %v", err)
	}
	if _, err := decodeExtension([]byte{0, 0, 0, 1}, []string{""}); !errors.Is(err, ErrTruncatedArchive) {
		t.Fatalf("truncated extension record: %v", err)
	}
}

// TestCoverSemanticAndUnitTables drives the remaining table-level branches.
func TestCoverSemanticAndUnitTables(t *testing.T) {
	t.Parallel()
	payload := make([]byte, 0, 14)
	payload = binary.BigEndian.AppendUint32(payload, 1)
	payload = binary.BigEndian.AppendUint32(payload, 0)
	payload = binary.BigEndian.AppendUint16(payload, 0xFFFF)
	payload = binary.BigEndian.AppendUint32(payload, 0)
	n := &node{sem: NewSemInfo()}
	if err := decodeSemantics(payload, []string{"v"}, []*node{n}); err != nil || n.sem.Len() != 0 {
		t.Fatalf("unknown semantic key must be skipped: %v", err)
	}
	badNodeIndex := binary.BigEndian.AppendUint32(nil, 9)
	badNodeIndex = binary.BigEndian.AppendUint32(badNodeIndex, 1)
	badNodeIndex = binary.BigEndian.AppendUint16(badNodeIndex, uint16(SemPurity))
	badNodeIndex = binary.BigEndian.AppendUint32(badNodeIndex, 0)
	if err := decodeSemantics(badNodeIndex, []string{"v"}, []*node{n}); !errors.Is(err, ErrMalformedSemantics) {
		t.Fatalf("out-of-range semantic node: %v", err)
	}
	if _, err := decodeUnit(nil, []string{""}, nil); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("short unit: %v", err)
	}
	hugeImports := make([]byte, 12)
	binary.BigEndian.PutUint32(hugeImports[4:8], 0xFF000000)
	if _, err := decodeUnit(hugeImports, []string{""}, nil); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("huge import count: %v", err)
	}
	hugeFiles := make([]byte, 12)
	binary.BigEndian.PutUint32(hugeFiles[8:12], 0xFF000000)
	if _, err := decodeUnit(hugeFiles, []string{""}, nil); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("huge file count: %v", err)
	}
	truncated := make([]byte, 12)
	if _, err := decodeUnit(truncated, nil, nil); !errors.Is(err, ErrIndexOutOfRange) && !errors.Is(err, ErrTruncatedArchive) {
		t.Fatalf("truncated unit record: %v", err)
	}
}

// TestCoverNodeAndIntervalBranches covers child links and overlap detection.
func TestCoverNodeAndIntervalBranches(t *testing.T) {
	t.Parallel()
	payload := make([]byte, 0, 50)
	payload = binary.BigEndian.AppendUint32(payload, 1)
	payload = binary.BigEndian.AppendUint64(payload, 7)
	payload = binary.BigEndian.AppendUint16(payload, uint16(KindIdent))
	for i := 0; i < 6; i++ {
		payload = binary.BigEndian.AppendUint32(payload, 0)
	}
	payload = binary.BigEndian.AppendUint32(payload, 0)
	payload = binary.BigEndian.AppendUint32(payload, 1)
	payload = binary.BigEndian.AppendUint32(payload, 0)
	if _, err := decodeNodes(payload, []string{"x"}); !errors.Is(err, ErrMalformedNode) {
		t.Fatalf("backward child link: %v", err)
	}
	overlap := NewNode(KindBlock, NewPosition(1, 1, 0, 10), "",
		NewNode(KindIdent, NewPosition(1, 1, 0, 6), "a"),
		NewNode(KindIdent, NewPosition(1, 1, 2, 8), "b"))
	if err := CheckIntervals(overlap, 10); !errors.Is(err, ErrBadPosition) {
		t.Fatalf("sibling overlap: %v", err)
	}
}

// TestCoverCompareUnitsBranches mutates each compared field in isolation.
func TestCoverCompareUnitsBranches(t *testing.T) {
	t.Parallel()
	base := mustUnit(t)
	variants := []func(*Unit){
		func(u *Unit) { u.Package.Name = "other" },
		func(u *Unit) { u.Files[0].Size++ },
		func(u *Unit) { u.Files = u.Files[:1] },
		func(u *Unit) { u.Roots = u.Roots[:1] },
		func(u *Unit) { u.Annotations["extra"] = "v" },
		func(u *Unit) { u.Annotations = map[string]string{} },
		func(u *Unit) {
			child := NewNode(KindIdent, NewPosition(1, 1, 0, 1), "z")
			u.Roots[0] = NewNode(KindBlock, NewPosition(1, 1, 0, 2), "", child)
		},
		func(u *Unit) { u.Roots[0] = nil },
	}
	for i, mutate := range variants {
		copyUnit := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), base))
		mutate(copyUnit)
		if err := CompareUnits(base, copyUnit); !errors.Is(err, ErrUnitMismatch) {
			t.Fatalf("variant %d was not detected", i)
		}
	}
}
