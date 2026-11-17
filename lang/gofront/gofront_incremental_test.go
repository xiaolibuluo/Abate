/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"errors"
	"strings"
	"testing"

	"abate/lang/air"
)

// incrementalUnit is a two-file unit used by all incremental tests.
func incrementalUnit() air.SourceUnit {
	return air.SourceUnit{
		Language: "go",
		Files: []air.SourceFile{
			air.NewSourceFile("a.go", []byte("package p\n\nfunc Fa() int { return 1 }\n")),
			air.NewSourceFile("b.go", []byte("package p\n\nfunc Fb() int { return 9 }\n")),
		},
	}
}

// TestIncrementalReusesUnchangedSubtrees proves real file-level reuse.
func TestIncrementalReusesUnchangedSubtrees(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	unit := incrementalUnit()
	old, diagnostics, err := frontend.Parse(unit)
	if err != nil || len(diagnostics) != 0 {
		t.Fatalf("cold parse: %v %v", err, diagnostics)
	}
	source := string(unit.Files[0].Content)
	at := strings.Index(source, "return 1") + len("return ")
	edits := []air.Edit{{
		Path: "a.go", Offset: at, EndOffset: at + 1, Text: []byte("2"),
	}}
	fresh, diff, err := frontend.ParseIncremental(unit, old, edits)
	if err != nil {
		t.Fatalf("ParseIncremental: %v", err)
	}
	if len(fresh.Files) != 2 || len(fresh.Roots) != 2 {
		t.Fatalf("incremental unit lost files: %d", len(fresh.Files))
	}
	if fresh.Roots[0] == old.Roots[0] {
		t.Fatal("the edited file was not reparsed")
	}
	if fresh.Roots[1] != old.Roots[1] {
		t.Fatal("the unchanged subtree was not reused verbatim")
	}
	if diff.FullRebuild || diff.Reused == 0 || diff.Empty() {
		t.Fatalf("unexpected diff: %s", diff.String())
	}
	if fresh.Files[0].ContentHash == old.Files[0].ContentHash {
		t.Fatal("edited file metadata was not refreshed")
	}
	if fresh.Files[1].Canonical() != old.Files[1].Canonical() {
		t.Fatal("unchanged file metadata drifted")
	}
}

// TestIncrementalNoEditsAndColdStart covers the reuse boundaries.
func TestIncrementalNoEditsAndColdStart(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	unit := incrementalUnit()
	old, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatal(err)
	}
	unchanged, diff, err := frontend.ParseIncremental(unit, old, nil)
	if err != nil || !diff.Empty() || diff.FullRebuild {
		t.Fatalf("no-op incremental: %v %s", err, diff.String())
	}
	if err := air.CompareUnits(old, unchanged); err != nil {
		t.Fatalf("no-op incremental changed the unit: %v", err)
	}
	cold, diff, err := frontend.ParseIncremental(unit, nil, nil)
	if err != nil || !diff.FullRebuild || cold.NodeCount() != old.NodeCount() {
		t.Fatalf("cold start: %v %s", err, diff.String())
	}
}

// TestIncrementalEditValidation checks malformed edits and diagnostics.
func TestIncrementalEditValidation(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	unit := incrementalUnit()
	old, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatal(err)
	}
	unknown := []air.Edit{{Path: "missing.go", Offset: 0, EndOffset: 1}}
	if _, _, err := frontend.ParseIncremental(unit, old, unknown); !errors.Is(err, ErrUnknownEditPath) {
		t.Fatalf("unknown edit path: %v", err)
	}
	outside := []air.Edit{{Path: "a.go", Offset: 0, EndOffset: 10000}}
	if _, _, err := frontend.ParseIncremental(unit, old, outside); !errors.Is(err, air.ErrInvalidEdit) {
		t.Fatalf("out-of-range edit: %v", err)
	}
	overlap := []air.Edit{
		{Path: "a.go", Offset: 0, EndOffset: 5},
		{Path: "a.go", Offset: 2, EndOffset: 7},
	}
	if _, _, err := frontend.ParseIncremental(unit, old, overlap); !errors.Is(err, air.ErrInvalidEdit) {
		t.Fatalf("overlapping edits: %v", err)
	}
	badUnit := air.SourceUnit{Language: "python"}
	if _, _, err := frontend.ParseIncremental(badUnit, old, nil); !errors.Is(err, ErrUnsupportedLanguage) {
		t.Fatalf("unsupported language: %v", err)
	}
}

// TestIncrementalSyntaxDiagnostics checks that parse findings are reported
// through unit annotations because the F1 contract has no diagnostic slot.
func TestIncrementalSyntaxDiagnostics(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	bad := air.SourceUnit{Language: "go", Files: []air.SourceFile{
		air.NewSourceFile("a.go", []byte("package p\n\nfunc F() { return 1 }\n")),
		air.NewSourceFile("b.go", []byte("func (")),
	}}
	unit, diff, err := frontend.ParseIncremental(bad, nil, nil)
	if err != nil || unit == nil {
		t.Fatalf("broken incremental parse: %v", err)
	}
	if !diff.FullRebuild || unit.Annotations["lang.go.diagnostics"] == "" {
		t.Fatalf("syntax findings were not surfaced: %s %v", diff.String(), unit.Annotations)
	}
}

// extraGo covers remaining literal, constraint, and alias families.
const extraGo = `package extra

import "unsafe"

type Alias = int
type Number interface {
	~int | ~string
}
type Pair[K comparable, V any] struct {
	Key K
	Val V
}

func Generic[K comparable](k K) K { return k }

func Extras() {
	flag := true
	var ch rune = 'x'
	_ = unsafe.Sizeof(flag)
	_ = Generic[int](1)
	_ = Pair[string, int]{Key: "a", Val: 2}
	var f func(...int) (int, error)
	_ = f
	var p *int
	_ = p
	_ = map[int][]string{}
	_ = ch
}
`

// TestExtraMappingFixtures exercises families absent from the main fixture.
func TestExtraMappingFixtures(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, extraGo)
	for _, kind := range []air.NodeKind{
		air.KindBoolLit, air.KindCharLit, air.KindUnion, air.KindRef,
		air.KindGenericInst, air.KindSpread, air.KindPointer, air.KindMap,
		air.KindSliceType, air.KindGenericParam,
	} {
		if !hasKind(unit, kind) {
			t.Fatalf("extra fixture is missing %s", kind)
		}
	}
	if got := annotationOf(unit, air.KindTypeDecl, air.SemNote); got != "type-alias" {
		t.Fatalf("type alias annotation = %q", got)
	}
	if err := VerifyTokenFidelity([]byte(extraGo), unit); err != nil {
		t.Fatalf("extra fixture token fidelity: %v", err)
	}
}
