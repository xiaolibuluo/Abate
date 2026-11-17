/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

const goInventoryFile = "../../../ci/constructs_go.json"

// parsedUnit builds a real AIR unit whose roots contain the supplied kinds.
func parsedUnit(t *testing.T, kinds ...air.NodeKind) *air.Unit {
	t.Helper()
	unit := air.NewUnit(air.NewPackageRef("sample", nil))
	children := make([]air.Node, 0, len(kinds))
	for _, kind := range kinds {
		children = append(children,
			air.NewNode(kind, air.NewPosition(1, 1, 0, 1), "x"))
	}
	root := air.NewNode(air.KindFuncDecl,
		air.NewPosition(1, 1, 0, 10), "func", children...)
	if err := unit.AddFile(air.NewFileRef("main.go",
		[]byte("package sample")), root); err != nil {
		t.Fatalf("unit fixture failed: %v", err)
	}
	return unit
}

// TestNewCoverageIndexLoadsSixMeasuredLanguages proves the six embedded
// grammars load and the Go row stays honestly NOT_MEASURED.
func TestNewCoverageIndexLoadsSixMeasuredLanguages(t *testing.T) {
	t.Parallel()
	index, err := NewCoverageIndex(goInventoryFile)
	if err != nil {
		t.Fatalf("coverage index failed: %v", err)
	}
	if !slices.Equal(index.Languages(), indexedLanguages) {
		t.Fatalf("language order wrong: %v", index.Languages())
	}
	for _, tag := range []string{"c", "cs", "js", "kt", "py", "rs"} {
		if !index.IsMeasured(tag) || index.Status(tag) != "MEASURED" ||
			len(index.Expected(tag)) == 0 {
			t.Fatalf("embedded language %s is not measured", tag)
		}
	}
	if index.IsMeasured("go") || index.Status("go") != statusNotMeasured ||
		len(index.Expected("go")) < 30 {
		t.Fatal("go inventory must be explicit and NOT_MEASURED")
	}
	goNodes := index.Expected("go")
	if !slices.Contains(goNodes, "FuncDecl") ||
		!slices.Contains(goNodes, "Comment") {
		t.Fatalf("go inventory is incomplete: %v", goNodes)
	}
	if index.Status("zz") != statusNotMeasured || len(index.Expected("zz")) != 0 {
		t.Fatal("unsupported language must not claim measurement")
	}
}

// TestCoverageComputesFromRealParse proves coverage.Compute and Observe are
// driven by a real AIR unit for all six embedded languages.
func TestCoverageComputesFromRealParse(t *testing.T) {
	t.Parallel()
	index, err := NewCoverageIndex(goInventoryFile)
	if err != nil {
		t.Fatalf("coverage index failed: %v", err)
	}
	unit := parsedUnit(t, air.KindBinary, air.KindNumLit, air.KindIdent,
		air.KindComment, air.KindCall)
	for _, tag := range []string{"c", "cs", "js", "kt", "py", "rs"} {
		report, coverageErr := index.Coverage(tag, unit)
		if coverageErr != nil {
			t.Fatalf("coverage %s failed: %v", tag, coverageErr)
		}
		if report.Language != tag || report.Total == 0 ||
			len(report.Entries) != report.Total || report.Percent() < 0 {
			t.Fatalf("coverage %s report wrong: %+v", tag, report)
		}
		if len(report.ObservedKindsOf()) == 0 || len(report.Uncovered()) == 0 {
			t.Fatalf("coverage %s lacks real observation rows", tag)
		}
		if index.Degradation(tag, unit) < 0 {
			t.Fatalf("degradation %s is negative", tag)
		}
	}
}

// TestGoCoverageRefusesFabricatedRatio proves the Go row and unknown languages
// expose explicit errors instead of a fake coverage percentage.
func TestGoCoverageRefusesFabricatedRatio(t *testing.T) {
	t.Parallel()
	index, err := NewCoverageIndex(goInventoryFile)
	if err != nil {
		t.Fatalf("coverage index failed: %v", err)
	}
	unit := parsedUnit(t, air.KindBinary)
	if _, err = index.Coverage("go", unit); !errors.Is(err, ErrNotMeasured) {
		t.Fatalf("go coverage must be NOT_MEASURED, got %v", err)
	}
	if index.Degradation("go", unit) != 0 {
		t.Fatal("go degradation must stay zero")
	}
	if _, err = index.Coverage("zz", unit); !errors.Is(err,
		ErrUnsupportedLanguage) {
		t.Fatalf("unknown language must fail, got %v", err)
	}
	if index.Degradation("zz", unit) != 0 {
		t.Fatal("unknown language degradation must stay zero")
	}
	if _, err = NewCoverageIndex("../../../ci/absent.json"); err == nil {
		t.Fatal("a missing inventory must be refused")
	}
}

// TestGoInventoryValidation proves the inventory reader and its honest label.
func TestGoInventoryValidation(t *testing.T) {
	t.Parallel()
	inventory, err := LoadGoInventory(goInventoryFile)
	if err != nil {
		t.Fatalf("inventory load failed: %v", err)
	}
	if inventory.IsMeasured() || inventory.Status != statusNotMeasured ||
		inventory.Language != "go" || len(inventory.Nodes()) < 30 {
		t.Fatalf("inventory label wrong: %+v", inventory)
	}
	if !slices.IsSorted(inventory.Nodes()) {
		t.Fatal("inventory nodes are not sorted")
	}
	if inventory.CategoryOf("FuncDecl") != "declaration" ||
		inventory.CategoryOf("absent") != "" {
		t.Fatal("inventory categories wrong")
	}
}

// TestGoInventoryRejectsBadDocuments proves a document claiming measurement is
// refused rather than being silently accepted.
func TestGoInventoryRejectsBadDocuments(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	documents := []string{
		`{"schemaVersion":1,"language":"py","status":"NOT_MEASURED","measured":false,"constructs":[{"nodeType":"X","category":"t"}]}`,
		`{"schemaVersion":1,"language":"go","status":"NOT_MEASURED","measured":true,"constructs":[{"nodeType":"X","category":"t"}]}`,
		`{"schemaVersion":1,"language":"go","status":"MEASURED","measured":false,"constructs":[{"nodeType":"X","category":"t"}]}`,
		`{"schemaVersion":1,"language":"go","status":"NOT_MEASURED","measured":false,"constructs":[]}`,
		`{"schemaVersion":1,"language":"go","status":"NOT_MEASURED","measured":false,"constructs":[{"nodeType":"X","category":"t"},{"nodeType":"X","category":"t"}]}`,
		`{"schemaVersion":0,"language":"go","status":"NOT_MEASURED","measured":false,"constructs":[{"nodeType":"X","category":"t"}]}`,
		`{`,
	}
	for i, document := range documents {
		path := filepath.Join(dir, "doc.json")
		if err := os.WriteFile(path, []byte(document), 0o600); err != nil {
			t.Fatalf("fixture %d write failed: %v", i, err)
		}
		if _, err := LoadGoInventory(path); err == nil {
			t.Fatalf("bad document %d was accepted", i)
		}
	}
	if _, err := LoadGoInventory(filepath.Join(dir, "absent.json")); err == nil {
		t.Fatal("missing document was accepted")
	}
}

// TestDegradationKeepsStatusDegraded proves degraded mappings are counted with
// their original status and out-of-scope rows do not dilute the ratio.
func TestDegradationKeepsStatusDegraded(t *testing.T) {
	t.Parallel()
	mappings := []grammar.Mapping{
		{Node: "a", Kind: air.KindBinary, Status: grammar.StatusExact},
		{Node: "b", Kind: air.KindNumLit, Status: grammar.StatusDegraded},
		{Node: "c", Kind: air.KindCall, Status: grammar.StatusDegraded},
		{Node: "d", Kind: air.KindInvalid, Status: grammar.StatusOutOfScope},
		{Node: "e", Kind: air.KindInvalid, Status: grammar.StatusUnmapped},
	}
	observed := map[air.NodeKind]int{air.KindNumLit: 3}
	got := Degradation(mappings, observed)
	if got <= 33 || got >= 34 {
		t.Fatalf("degradation ratio wrong: %v", got)
	}
	if Degradation(nil, nil) != 0 ||
		Degradation(mappings, nil) != 0 {
		t.Fatal("a zero denominator must yield zero")
	}
}
