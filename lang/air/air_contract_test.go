/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

// TestPositionHelpers checks interval ordering and rendering.
func TestPositionHelpers(t *testing.T) {
	t.Parallel()
	outer := NewPosition(1, 1, 0, 10)
	inner := NewPosition(2, 3, 4, 8)
	if !outer.Valid() || !outer.Contains(inner) || !outer.Overlaps(inner) {
		t.Fatal("valid outer interval rejected")
	}
	if inner.Contains(outer) || NewPosition(1, 1, 12, 13).Overlaps(outer) {
		t.Fatal("interval containment inverted")
	}
	if NewPosition(-1, 0, 0, 1).Valid() {
		t.Fatal("negative line must be invalid")
	}
	if NewPosition(1, 1, 5, 4).Valid() {
		t.Fatal("reversed interval must be invalid")
	}
	if outer.String() == "" || outer.Canonical() != "1/1/0/10" {
		t.Fatalf("position rendering: %s %s", outer.String(), outer.Canonical())
	}
}

// TestSemInfoCanonicalAndKeys proves map order never leaks into the encoding.
func TestSemInfoCanonicalAndKeys(t *testing.T) {
	t.Parallel()
	first := NewSemInfo().With(SemPurity, "pure").With(SemTypeRef, "int")
	second := NewSemInfo().With(SemTypeRef, "int").With(SemPurity, "pure")
	if first.Canonical() != second.Canonical() || first.Len() != 2 {
		t.Fatal("SemInfo canonical form depends on insertion order")
	}
	if value, ok := first.Get(SemPurity); !ok || value != "pure" {
		t.Fatal("SemInfo.Get lost a value")
	}
	if _, ok := first.Get(SemScopeID); ok {
		t.Fatal("SemInfo.Get invented a value")
	}
	keys := first.Keys()
	if len(keys) != 2 || keys[0] != SemTypeRef || keys[1] != SemPurity {
		t.Fatalf("unsorted keys: %v", keys)
	}
	clone := first.Clone()
	clone[SemPurity] = "side-effect"
	if first.Canonical() == clone.Canonical() {
		t.Fatal("Clone shares backing storage")
	}
	var nilSem SemInfo
	if nilSem.Canonical() != "" || nilSem.Len() != 0 {
		t.Fatal("nil SemInfo must canonicalize to the empty string")
	}
}

// TestSemKeyNames covers every annotation slot and its string form.
func TestSemKeyNames(t *testing.T) {
	t.Parallel()
	for key := SemInvalid; key <= SemKeyLast; key++ {
		if key.String() == "" {
			t.Fatalf("key %d has no name", key)
		}
		if key.Valid() != (key != SemInvalid) {
			t.Fatalf("key %d validity", key)
		}
	}
	if SemKey(9999).String() != "Invalid" || !SemKeyLast.Valid() {
		t.Fatal("semantic key boundary broken")
	}
}

// TestFileRefAndPackageRef checks metadata derivation and normalization.
func TestFileRefAndPackageRef(t *testing.T) {
	t.Parallel()
	lf := NewFileRef("a.go", []byte("x\ny\n"))
	crlf := NewFileRef("b.go", []byte("x\r\ny\r\n"))
	bom := NewFileRef("c.go", []byte("\xef\xbb\xbfx"))
	if lf.LineEnding != "\n" || crlf.LineEnding != "\r\n" || bom.Encoding != "utf-8-bom" {
		t.Fatalf("line ending or encoding detection: %q %q %q",
			lf.LineEnding, crlf.LineEnding, bom.Encoding)
	}
	if lf.Digest() == crlf.Digest() {
		t.Fatal("different content must produce different digests")
	}
	if lf.Size != 4 || len(lf.Digest()) != 64 {
		t.Fatalf("file metadata: %d %q", lf.Size, lf.Digest())
	}
	pkg := NewPackageRef("p", []string{"z", "a", "a"})
	if len(pkg.Imports) != 2 || pkg.Imports[0] != "a" || pkg.Imports[1] != "z" {
		t.Fatalf("import normalization: %v", pkg.Imports)
	}
	copyPkg := pkg.Clone()
	copyPkg.Imports[0] = "mutated"
	if pkg.Imports[0] != "a" || pkg.Canonical() != "p\x00a\x00z" {
		t.Fatal("PackageRef.Clone shares backing storage")
	}
	if err := NewPackageRef("p", []string{""}).Validate(); !errors.Is(err, ErrMalformedUnit) {
		t.Fatalf("empty import: %v", err)
	}
}

// TestSourceUnitContract checks in-memory source handling and ordering.
func TestSourceUnitContract(t *testing.T) {
	t.Parallel()
	unit := SourceUnit{
		Language: "go",
		Dialect:  "go1.26",
		Files: []SourceFile{
			NewSourceFile("z.go", []byte("z")),
			NewSourceFile("a.go", []byte("a")),
		},
	}
	sorted := unit.SortedFiles()
	if sorted[0].Path != "a.go" || unit.Files[0].Path != "z.go" {
		t.Fatal("SortedFiles mutated or misordered the source unit")
	}
	if found, ok := unit.Find("a.go"); !ok || found.Size() != 1 {
		t.Fatal("Find failed")
	}
	if _, ok := unit.Find("missing.go"); ok {
		t.Fatal("Find invented a file")
	}
	changed := unit.WithFile(NewSourceFile("a.go", []byte("longer")))
	if bytes.Equal(unit.Files[1].Content, changed.Files[1].Content) {
		t.Fatal("WithFile mutated the original unit")
	}
	if len(changed.Files) != 2 {
		t.Fatal("WithFile appended instead of replacing")
	}
	added := unit.WithFile(NewSourceFile("m.go", nil))
	if len(added.Files) != 3 || added.Language != "go" {
		t.Fatal("WithFile add path failed")
	}
	bad := SourceUnit{Files: []SourceFile{NewSourceFile("a.go", nil), NewSourceFile("a.go", nil)}}
	if err := bad.Validate(); !errors.Is(err, ErrDuplicatePath) {
		t.Fatalf("duplicate path: %v", err)
	}
	if err := (SourceUnit{Files: []SourceFile{{Path: ""}}}).Validate(); !errors.Is(err, ErrMalformedUnit) {
		t.Fatal("empty path accepted")
	}
}

// TestApplyEdits checks deterministic, non-overlapping byte replacements.
func TestApplyEdits(t *testing.T) {
	t.Parallel()
	src := []byte("hello world")
	edits := []Edit{
		{Path: "a.go", Offset: 6, EndOffset: 11, Text: []byte("air")},
		{Path: "a.go", Offset: 0, EndOffset: 5, Text: []byte("hi")},
	}
	out, err := ApplyEdits(src, edits)
	if err != nil || string(out) != "hi air" {
		t.Fatalf("ApplyEdits = %q, %v", out, err)
	}
	reversed := []Edit{edits[1], edits[0]}
	again, err := ApplyEdits(src, reversed)
	if err != nil || !bytes.Equal(out, again) {
		t.Fatal("edit order leaked into the result")
	}
	if string(src) != "hello world" {
		t.Fatal("ApplyEdits mutated the source buffer")
	}
	overlap := []Edit{{Offset: 0, EndOffset: 5}, {Offset: 4, EndOffset: 6}}
	if _, err := ApplyEdits(src, overlap); !errors.Is(err, ErrInvalidEdit) {
		t.Fatalf("overlap: %v", err)
	}
	if _, err := ApplyEdits(src, []Edit{{Offset: 0, EndOffset: 99}}); !errors.Is(err, ErrInvalidEdit) {
		t.Fatalf("out of range: %v", err)
	}
	if err := (Edit{Offset: -1, EndOffset: 2}).Validate(10); !errors.Is(err, ErrInvalidEdit) {
		t.Fatalf("negative offset: %v", err)
	}
	if cloned := NewSourceFile("x.go", src).Clone(); cloned.Size() != 11 {
		t.Fatal("SourceFile.Clone changed the buffer")
	}
}

// TestDiagnosticAndDiff checks rendering and diff bookkeeping.
func TestDiagnosticAndDiff(t *testing.T) {
	t.Parallel()
	d := Diagnostic{
		Path: "a.go", Line: 3, Column: 7, Severity: SeverityError,
		Code: "LANG-GO-E-001", Message: "unexpected token",
	}
	if !d.IsError() || d.String() != "a.go:3:7 error LANG-GO-E-001: unexpected token" {
		t.Fatalf("diagnostic rendering: %s", d.String())
	}
	if (Diagnostic{Severity: SeverityWarning}).IsError() {
		t.Fatal("warning classified as error")
	}
	if MaxDiagnostics != 100 || SeverityInfo != "info" {
		t.Fatal("diagnostic constants drifted")
	}
	diff := Diff{Added: []NodeID{1}, Reused: 2}
	if diff.Empty() || diff.String() == "" {
		t.Fatal("non-empty diff reported empty")
	}
	if !(Diff{}).Empty() {
		t.Fatal("empty diff reported non-empty")
	}
}

// TestFrontendInterfaceShape verifies the seven-language contract without
// defining any placeholder implementation in production code.
func TestFrontendInterfaceShape(t *testing.T) {
	t.Parallel()
	typ := reflect.TypeOf((*Frontend)(nil)).Elem()
	if typ.NumMethod() != 2 {
		t.Fatalf("Frontend has %d methods", typ.NumMethod())
	}
	parse, ok := typ.MethodByName("Parse")
	if !ok || parse.Type.NumIn() != 1 || parse.Type.NumOut() != 3 {
		t.Fatalf("Parse signature drifted: %s", parse.Type.String())
	}
	if parse.Type.In(0) != reflect.TypeOf(SourceUnit{}) {
		t.Fatal("Parse must accept SourceUnit")
	}
	incremental, ok := typ.MethodByName("ParseIncremental")
	if !ok || incremental.Type.NumIn() != 3 || incremental.Type.NumOut() != 3 {
		t.Fatalf("ParseIncremental signature drifted: %s", incremental.Type.String())
	}
	if incremental.Type.In(1) != reflect.TypeOf((*Unit)(nil)) ||
		incremental.Type.In(2) != reflect.TypeOf([]Edit{}) {
		t.Fatal("ParseIncremental must accept the old unit and edits")
	}
	if incremental.Type.Out(1) != reflect.TypeOf(Diff{}) {
		t.Fatal("ParseIncremental must return Diff")
	}
}

// TestComputeDiff verifies identifier multiset arithmetic.
func TestComputeDiff(t *testing.T) {
	t.Parallel()
	old := mustUnit(t)
	fresh := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), old))
	unchanged := ComputeDiff(old, fresh)
	if !unchanged.Empty() || unchanged.Reused != old.NodeCount() || unchanged.FullRebuild {
		t.Fatalf("unchanged diff: %s", unchanged.String())
	}
	cold := ComputeDiff(nil, fresh)
	if !cold.FullRebuild || len(cold.Added) == 0 || cold.Reused != 0 {
		t.Fatalf("cold diff: %s", cold.String())
	}
	removed := ComputeDiff(old, NewUnit(NewPackageRef("x", nil)))
	if len(removed.Removed) != old.NodeCount() || removed.Reused != 0 {
		t.Fatalf("removal diff: %s", removed.String())
	}
}
