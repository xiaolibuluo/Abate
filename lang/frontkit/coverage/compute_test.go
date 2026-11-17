/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"maps"
	"slices"
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// SampleMappings returns a deterministic table that exercises every mapping status.
func sampleMappings() []grammar.Mapping {
	return []grammar.Mapping{
		{Node: "function_item", Kind: air.KindFuncDecl, Status: grammar.StatusExact, Reason: "dedicated node", Rule: "rule.fn"},
		{Node: "call_expression", Kind: air.KindCall, Status: grammar.StatusExact, Reason: "shared node", Rule: "rule.call"},
		{Node: "closure_expression", Kind: air.KindFuncLit, Status: grammar.StatusDegraded, Reason: "lowered to func lit", Rule: "rule.lit"},
		{Node: "let_declaration", Kind: air.KindVarDecl, Status: grammar.StatusDegraded, Reason: "lowered to var", Rule: "rule.var"},
		{Node: "attribute_item", Kind: air.KindAttribute, Status: grammar.StatusOutOfScope, Reason: "metadata", Rule: "rule.attr"},
		{Node: "macro_invocation", Kind: air.KindInvalid, Status: grammar.StatusUnmapped, Reason: "no rule", Rule: ""},
		{Node: "use_declaration", Kind: air.KindImport, Status: "custom", Reason: "custom status", Rule: "rule.use"},
		{Node: "method_item", Kind: air.KindFuncDecl, Status: grammar.StatusExact, Reason: "shares func decl", Rule: "rule.method"},
	}
}

// SampleObservations returns the tallies joined with the sample mapping table.
func sampleObservations() map[air.NodeKind]int {
	return map[air.NodeKind]int{
		air.KindFuncDecl:  2,
		air.KindCall:      0,
		air.KindFuncLit:   3,
		air.KindVarDecl:   0,
		air.KindAttribute: 5,
		air.KindInvalid:   0,
		air.KindImport:    4,
	}
}

// ExpectedEntries returns the rows the sample join must produce.
func expectedEntries() []Entry {
	return []Entry{
		{Node: "function_item", Kind: air.KindFuncDecl, Status: grammar.StatusExact, Reason: "dedicated node", Rule: "rule.fn", Observed: true, Count: 2},
		{Node: "call_expression", Kind: air.KindCall, Status: grammar.StatusExact, Reason: "shared node", Rule: "rule.call", Observed: false, Count: 0},
		{Node: "closure_expression", Kind: air.KindFuncLit, Status: grammar.StatusDegraded, Reason: "lowered to func lit", Rule: "rule.lit", Observed: true, Count: 3},
		{Node: "let_declaration", Kind: air.KindVarDecl, Status: grammar.StatusDegraded, Reason: "lowered to var", Rule: "rule.var", Observed: false, Count: 0},
		{Node: "attribute_item", Kind: air.KindAttribute, Status: grammar.StatusOutOfScope, Reason: "metadata", Rule: "rule.attr", Observed: true, Count: 5},
		{Node: "macro_invocation", Kind: air.KindInvalid, Status: grammar.StatusUnmapped, Reason: "no rule", Rule: "", Observed: false, Count: 0},
		{Node: "use_declaration", Kind: air.KindImport, Status: "custom", Reason: "custom status", Rule: "rule.use", Observed: true, Count: 4},
		{Node: "method_item", Kind: air.KindFuncDecl, Status: grammar.StatusExact, Reason: "shares func decl", Rule: "rule.method", Observed: true, Count: 2},
	}
}

// TestComputeJoinsMappingsWithObservedKinds proves every counter and row of the join.
func TestComputeJoinsMappingsWithObservedKinds(t *testing.T) {
	t.Parallel()
	g := &grammar.Grammar{
		Language: "rs",
		Version:  "0.21.0",
		Source:   "tree-sitter-rust",
	}
	mappings := sampleMappings()
	got := Compute(g, mappings, sampleObservations())
	if got.Language != "rs" || got.Version != "0.21.0" || got.Source != "tree-sitter-rust" {
		t.Fatalf("provenance was not carried over: %+v", got)
	}
	if got.Total != len(mappings) {
		t.Fatalf("Total = %d, want %d", got.Total, len(mappings))
	}
	if got.Exact != 2 || got.Degraded != 1 || got.Covered != 3 {
		t.Fatalf("coverage counters = exact %d degraded %d covered %d, want 2 1 3",
			got.Exact, got.Degraded, got.Covered)
	}
	if got.OutOfScope != 1 || got.Unmapped != 1 {
		t.Fatalf("scope counters = out-of-scope %d unmapped %d, want 1 1",
			got.OutOfScope, got.Unmapped)
	}
	if got.ObservedKinds != 4 {
		t.Fatalf("ObservedKinds = %d, want 4 distinct kinds", got.ObservedKinds)
	}
	if len(got.Entries) != len(mappings) {
		t.Fatalf("got %d rows, want %d", len(got.Entries), len(mappings))
	}
	if !slices.Equal(got.Entries, expectedEntries()) {
		t.Fatalf("rows = %+v, want %+v", got.Entries, expectedEntries())
	}
	assertPercentMatches(t, "Percent", got.Percent(), 37.5)
}

// TestComputeWithNilGrammarAndNoMappings proves the empty join is safe and empty.
func TestComputeWithNilGrammarAndNoMappings(t *testing.T) {
	t.Parallel()
	got := Compute(nil, nil, nil)
	if got.Language != "" || got.Version != "" || got.Source != "" {
		t.Fatalf("a nil grammar must not invent provenance: %+v", got)
	}
	if got.Total != 0 || got.Covered != 0 || got.ObservedKinds != 0 {
		t.Fatalf("an empty join must stay zero: %+v", got)
	}
	if got.Entries == nil || len(got.Entries) != 0 {
		t.Fatalf("an empty join must keep an empty non-nil row slice: %+v", got.Entries)
	}
}

// TestComputeWithoutAnyObservation proves every row stays unobserved with a nil tally.
func TestComputeWithoutAnyObservation(t *testing.T) {
	t.Parallel()
	mappings := sampleMappings()
	got := Compute(nil, mappings, nil)
	if got.Total != len(mappings) {
		t.Fatalf("Total = %d, want %d", got.Total, len(mappings))
	}
	if got.Covered != 0 || got.Exact != 0 || got.Degraded != 0 || got.ObservedKinds != 0 {
		t.Fatalf("a nil tally must cover nothing: %+v", got)
	}
	if got.OutOfScope != 1 || got.Unmapped != 1 {
		t.Fatalf("status counters must not depend on observation: %+v", got)
	}
	for _, entry := range got.Entries {
		if entry.Observed || entry.Count != 0 {
			t.Fatalf("row %q claims an observation without a tally", entry.Node)
		}
	}
}

// TestObserveCountsEveryKindInAUnit proves the tally walks two real AIR trees.
func TestObserveCountsEveryKindInAUnit(t *testing.T) {
	t.Parallel()
	unit := air.NewUnit(air.NewPackageRef("demo", nil))
	pos := air.NewPosition(1, 1, 0, 12)
	ident := air.NewNode(air.KindIdent, pos, "x")
	ret := air.NewNode(air.KindReturn, pos, "return")
	block := air.NewNode(air.KindBlock, pos, "body", ret)
	fn := air.NewNode(air.KindFuncDecl, pos, "fn", ident, block)
	imp := air.NewNode(air.KindImport, pos, "import")
	firstRoot := air.NewNode(air.KindPackage, pos, "pkg", fn, imp)
	call := air.NewNode(air.KindCall, pos, "call", air.NewNode(air.KindIdent, pos, "f"))
	secondRoot := air.NewNode(air.KindModuleDecl, pos, "mod", call)
	if err := unit.AddFile(air.NewFileRef("a.rs", []byte("fn")), firstRoot); err != nil {
		t.Fatalf("add a.rs: %v", err)
	}
	if err := unit.AddFile(air.NewFileRef("b.rs", []byte("call")), secondRoot); err != nil {
		t.Fatalf("add b.rs: %v", err)
	}
	want := map[air.NodeKind]int{
		air.KindPackage:    1,
		air.KindFuncDecl:   1,
		air.KindIdent:      2,
		air.KindBlock:      1,
		air.KindReturn:     1,
		air.KindImport:     1,
		air.KindModuleDecl: 1,
		air.KindCall:       1,
	}
	got := Observe(unit)
	if !maps.Equal(got, want) {
		t.Fatalf("Observe = %v, want %v", got, want)
	}
	if len(air.Collect(unit)) != 9 {
		t.Fatalf("the fixture must contain nine nodes, got %d", len(air.Collect(unit)))
	}
}

// TestObserveHandlesNilAndEmptyUnits proves missing trees yield an empty non-nil tally.
func TestObserveHandlesNilAndEmptyUnits(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		unit *air.Unit
	}{
		{"nil-unit", nil},
		{"zero-unit", &air.Unit{}},
		{"new-unit", air.NewUnit(air.NewPackageRef("demo", nil))},
		{"nil-root", &air.Unit{Roots: []air.Node{nil}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Observe(tc.unit)
			if got == nil {
				t.Fatal("Observe returned a nil map")
			}
			if len(got) != 0 {
				t.Fatalf("Observe returned %d kinds, want none", len(got))
			}
		})
	}
}

// TestMergeBuildsTallyFromNilAndAccumulates proves a nil destination is created and counts add up.
func TestMergeBuildsTallyFromNilAndAccumulates(t *testing.T) {
	t.Parallel()
	from := map[air.NodeKind]int{air.KindFuncDecl: 2, air.KindCall: 3}
	first := Merge(nil, from)
	if !maps.Equal(first, from) {
		t.Fatalf("Merge into nil = %v, want %v", first, from)
	}
	second := Merge(first, map[air.NodeKind]int{air.KindFuncDecl: 4, air.KindIdent: 1})
	want := map[air.NodeKind]int{air.KindFuncDecl: 6, air.KindCall: 3, air.KindIdent: 1}
	if !maps.Equal(second, want) {
		t.Fatalf("merged tally = %v, want %v", second, want)
	}
	if first[air.KindFuncDecl] != 6 {
		t.Fatal("Merge must accumulate into the supplied destination map")
	}
}

// TestMergeKeepsTallyWhenSourceIsEmpty proves an empty source does not replace the destination.
func TestMergeKeepsTallyWhenSourceIsEmpty(t *testing.T) {
	t.Parallel()
	into := map[air.NodeKind]int{air.KindImport: 7}
	got := Merge(into, nil)
	if !maps.Equal(got, into) {
		t.Fatalf("Merge with nil source = %v, want %v", got, into)
	}
	got[air.KindImport] = 9
	if into[air.KindImport] != 9 {
		t.Fatal("Merge must return the same destination map instance")
	}
	again := Merge(into, map[air.NodeKind]int{})
	if !maps.Equal(again, into) {
		t.Fatalf("Merge with empty source = %v, want %v", again, into)
	}
}
