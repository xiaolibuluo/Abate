/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package incremental

import (
	"bytes"
	"errors"
	"testing"

	"abate/lang/air"
	"abate/lang/gofront"
)

const probeSource = "package sample\n\nfunc Add(a int, b int) int { return a + b }\n"

// fakeFrontend is a deterministic test double for the parse contract.
type fakeFrontend struct {
	fresh *air.Unit
	diff  air.Diff
	err   error
}

// Parse returns the configured unit.
func (f fakeFrontend) Parse(air.SourceUnit) (*air.Unit, []air.Diagnostic, error) {
	return f.fresh, nil, f.err
}

// ParseIncremental returns the configured unit and diff.
func (f fakeFrontend) ParseIncremental(air.SourceUnit, *air.Unit,
	[]air.Edit) (*air.Unit, air.Diff, error) {
	return f.fresh, f.diff, f.err
}

// probeUnit builds a two-file source unit with one edited byte range.
func probeUnit() (air.SourceUnit, []air.Edit) {
	src := []byte(probeSource)
	other := []byte("package sample\n\nvar X = 1\n")
	unit := air.SourceUnit{
		Language: "go",
		Files: []air.SourceFile{
			air.NewSourceFile("main.go", src),
			air.NewSourceFile("other.go", other),
		},
	}
	at := bytes.Index(src, []byte(" + "))
	edits := []air.Edit{
		{Path: "main.go", Offset: at + 1, EndOffset: at + 2, Text: []byte("-")},
	}
	return unit, edits
}

// TestProbeCheckConsistentAfterRealEdit proves ParseIncremental keeps the
// unedited file's AIR root identifier and reports one changed range.
func TestProbeCheckConsistentAfterRealEdit(t *testing.T) {
	t.Parallel()
	unit, edits := probeUnit()
	frontend := gofront.NewGoFrontend()
	old, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatalf("initial parse failed: %v", err)
	}
	result, err := NewProbe(frontend).Check(unit, old, edits)
	if err != nil {
		t.Fatalf("probe failed: %v", err)
	}
	if !result.Consistent || result.ChangedRanges != 1 {
		t.Fatalf("probe result wrong: %+v", result)
	}
}

// TestProbeCountRangesMergesDuplicates proves distinct edited intervals are
// counted without map iteration affecting the value.
func TestProbeCountRangesMergesDuplicates(t *testing.T) {
	t.Parallel()
	edits := []air.Edit{
		{Path: "a.go", Offset: 1, EndOffset: 2},
		{Path: "a.go", Offset: 1, EndOffset: 2},
		{Path: "b.go", Offset: 3, EndOffset: 5},
	}
	if got := countRanges(edits); got != 2 {
		t.Fatalf("expected 2 ranges, got %d", got)
	}
	if got := countRanges(nil); got != 0 {
		t.Fatalf("expected 0 ranges, got %d", got)
	}
}

// TestProbeRejectsInvalidInputs proves every contract error is surfaced.
func TestProbeRejectsInvalidInputs(t *testing.T) {
	t.Parallel()
	prior := air.NewUnit(air.NewPackageRef("sample", nil))
	if _, err := NewProbe(nil).Check(air.SourceUnit{}, prior, nil); !errors.Is(err,
		errNilFrontend) {
		t.Fatalf("nil frontend must fail, got %v", err)
	}
	probe := NewProbe(gofront.NewGoFrontend())
	if _, err := probe.Check(air.SourceUnit{}, nil, nil); !errors.Is(err,
		errNilUnit) {
		t.Fatalf("nil prior unit must fail, got %v", err)
	}
	duplicate := air.SourceUnit{Language: "go",
		Files: []air.SourceFile{
			air.NewSourceFile("a.go", []byte("package a")),
			air.NewSourceFile("a.go", []byte("package a")),
		}}
	if _, err := probe.Check(duplicate, prior, nil); err == nil {
		t.Fatal("duplicate paths must fail")
	}
	failing := fakeFrontend{err: errors.New("fixture parse failure")}
	if _, err := NewProbe(failing).Check(air.SourceUnit{}, prior, nil); err == nil {
		t.Fatal("a frontend error must be surfaced")
	}
}

// TestProbeDetectsInconsistentFreshUnit proves a missing unit, an empty diff
// and a changed unedited root all produce an explicit inconsistency.
func TestProbeDetectsInconsistentFreshUnit(t *testing.T) {
	t.Parallel()
	unit, edits := probeUnit()
	prior := air.NewUnit(air.NewPackageRef("sample", nil))
	missing := fakeFrontend{diff: air.Diff{Added: []air.NodeID{air.NodeID(1)}}}
	result, err := NewProbe(missing).Check(unit, prior, edits)
	if err != nil || result.Consistent {
		t.Fatalf("a missing fresh unit must be inconsistent: %+v %v", result, err)
	}
	emptyDiff := fakeFrontend{fresh: prior}
	result, err = NewProbe(emptyDiff).Check(unit, prior, edits)
	if err != nil || result.Consistent {
		t.Fatalf("an empty diff must be inconsistent: %+v %v", result, err)
	}
	changed := air.NewUnit(air.NewPackageRef("sample", nil))
	root := air.NewNode(air.KindIdent, air.NewPosition(1, 1, 0, 1), "different")
	if err = changed.AddFile(air.NewFileRef("other.go",
		[]byte("package sample")), root); err != nil {
		t.Fatalf("fixture failed: %v", err)
	}
	stale := fakeFrontend{fresh: changed,
		diff: air.Diff{Added: []air.NodeID{air.NodeID(1)}}}
	result, err = NewProbe(stale).Check(unit, prior, edits)
	if err != nil || result.Consistent {
		t.Fatalf("a changed unedited root must be inconsistent: %+v %v",
			result, err)
	}
}
