/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"strings"
	"testing"

	"abate/lang/air"
)

// incrementalUnit is the two-file unit shared by the incremental tests.
func incrementalUnit() air.SourceUnit {
	return air.SourceUnit{
		Language: "go", Dialect: "default",
		Files: []air.SourceFile{
			air.NewSourceFile("a.go", []byte("let a = 1\n")),
			air.NewSourceFile("b.go", []byte("let b = 2\n")),
		},
	}
}

// TestIncrementalReusesUnchangedFile proves the file-level reuse works and
// that the reuse is reported honestly in the unit annotations.
func TestIncrementalReusesUnchangedFile(t *testing.T) {
	frontend := mustFrontend(t, newFakeSpec())
	unit := incrementalUnit()
	old, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatalf("cold parse: %v", err)
	}
	source := string(unit.Files[0].Content)
	at := strings.Index(source, "1")
	edits := []air.Edit{{
		Path: "a.go", Offset: at, EndOffset: at + 1, Text: []byte("7"),
	}}
	fresh, diff, err := frontend.ParseIncremental(unit, old, edits)
	if err != nil {
		t.Fatalf("ParseIncremental: %v", err)
	}
	if fresh.Annotations["lang.go.reusedFiles"] != "1" {
		t.Fatalf("unexpected reuse count %q",
			fresh.Annotations["lang.go.reusedFiles"])
	}
	if fresh.Annotations["lang.go.incremental"] != "file-level-reuse" {
		t.Fatalf("unexpected reuse mode %q",
			fresh.Annotations["lang.go.incremental"])
	}
	if diff.FullRebuild || diff.Reused == 0 {
		t.Fatalf("unexpected diff %s", diff.String())
	}
	if fresh.Roots[0] == old.Roots[0] {
		t.Fatal("the edited file was not reparsed")
	}
	if fresh.Roots[1] != old.Roots[1] {
		t.Fatal("the unchanged subtree was not reused verbatim")
	}
}

// TestIncrementalDialectChangeForcesFullRebuild is the regression test for a
// real defect: a dialect switch used to keep the subtrees of the previous
// dialect, so the result was silently stale.
func TestIncrementalDialectChangeForcesFullRebuild(t *testing.T) {
	frontend := mustFrontend(t, newFakeSpec())
	unit := incrementalUnit()
	old, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatalf("cold parse: %v", err)
	}
	other := unit
	other.Dialect = "other"
	fresh, diff, err := frontend.ParseIncremental(other, old, nil)
	if err != nil {
		t.Fatalf("ParseIncremental: %v", err)
	}
	if fresh.Annotations["lang.go.reusedFiles"] != "0" {
		t.Fatalf("a dialect change reused %q files",
			fresh.Annotations["lang.go.reusedFiles"])
	}
	if fresh.Annotations["lang.go.incremental"] != "full-rebuild" {
		t.Fatalf("unexpected reuse mode %q",
			fresh.Annotations["lang.go.incremental"])
	}
	if !diff.FullRebuild {
		t.Fatalf("a dialect change was not reported as a full rebuild: %s",
			diff.String())
	}
	if fresh.Roots[0] == old.Roots[0] {
		t.Fatal("the stale subtree of the previous dialect was reused")
	}
	if fresh.Annotations["lang.go.dialect"] != "other" {
		t.Fatalf("the new dialect was not recorded: %q",
			fresh.Annotations["lang.go.dialect"])
	}
}

// TestIncrementalEditsMatchFullParse proves the incremental result is the
// same tree as a full parse of the edited source, for insertion, deletion and
// replacement of a token.
func TestIncrementalEditsMatchFullParse(t *testing.T) {
	cases := []struct {
		name  string
		edits func(source string) []air.Edit
	}{
		{"insert", func(source string) []air.Edit {
			at := strings.Index(source, "1")
			return []air.Edit{{
				Path: "a.go", Offset: at, EndOffset: at,
				Text: []byte("0 "),
			}}
		}},
		{"delete", func(source string) []air.Edit {
			at := strings.Index(source, "1")
			return []air.Edit{{
				Path: "a.go", Offset: at, EndOffset: at + 1,
			}}
		}},
		{"replace", func(source string) []air.Edit {
			at := strings.Index(source, "1")
			return []air.Edit{{
				Path: "a.go", Offset: at, EndOffset: at + 1,
				Text: []byte("42"),
			}}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			frontend := mustFrontend(t, newFakeSpec())
			unit := incrementalUnit()
			old, _, err := frontend.Parse(unit)
			if err != nil {
				t.Fatalf("cold parse: %v", err)
			}
			edits := tc.edits(string(unit.Files[0].Content))
			fresh, _, err := frontend.ParseIncremental(unit, old, edits)
			if err != nil {
				t.Fatalf("ParseIncremental: %v", err)
			}
			full, _, err := frontend.Parse(unitAfterEdits(unit, edits))
			if err != nil {
				t.Fatalf("full parse: %v", err)
			}
			if err := compareTrees(fresh, full); err != nil {
				t.Fatalf("incremental result differs from the full parse: %v",
					err)
			}
		})
	}
}

// TestIncrementalRejectsBadInput keeps the error contract of the entry point.
func TestIncrementalRejectsBadInput(t *testing.T) {
	frontend := mustFrontend(t, newFakeSpec())
	unit := incrementalUnit()
	old, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatal(err)
	}
	unknown := []air.Edit{{Path: "missing.go", Offset: 0, EndOffset: 1}}
	if _, _, err := frontend.ParseIncremental(unit, old, unknown); !errors.Is(err, errUnknownEditPath) {
		t.Fatalf("unknown edit path: %v", err)
	}
	bad := unit
	bad.Dialect = "!"
	if _, _, err := frontend.ParseIncremental(bad, old, nil); !errors.Is(err, ErrBadDialect) {
		t.Fatalf("bad dialect: %v", err)
	}
}

// unitAfterEdits returns the unit with the edits applied.
func unitAfterEdits(unit air.SourceUnit, edits []air.Edit) air.SourceUnit {
	out := unit
	for _, edit := range edits {
		src, ok := out.Find(edit.Path)
		if !ok {
			continue
		}
		next, err := air.ApplyEdits(src.Content, edits)
		if err != nil {
			return out
		}
		out = out.WithFile(air.NewSourceFile(edit.Path, next))
	}
	return out
}

// compareTrees compares two units while ignoring the incremental-only
// annotations, because those describe the reuse and not the tree.
func compareTrees(left, right *air.Unit) error {
	stripped := *left
	stripped.Annotations = make(map[string]string, len(left.Annotations))
	for key, value := range left.Annotations {
		if strings.HasSuffix(key, ".incremental") ||
			strings.HasSuffix(key, ".reusedFiles") {
			continue
		}
		stripped.Annotations[key] = value
	}
	return air.CompareUnits(&stripped, right)
}
