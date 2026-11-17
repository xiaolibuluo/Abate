/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"math"
	"slices"
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// AssertPercentMatches fails when a computed percentage differs from the expected one.
func assertPercentMatches(t *testing.T, label string, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 1e-9 {
		t.Fatalf("%s = %.6f, want %.6f", label, got, want)
	}
}

// EntryNodeNames returns the upstream construct names of a row slice in order.
func entryNodeNames(entries []Entry) []string {
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Node)
	}
	return names
}

// TestPercentHelpersHandleZeroTotal proves all three ratios are zero for an empty inventory.
func TestPercentHelpersHandleZeroTotal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		data Coverage
	}{
		{"plain-zero", Coverage{}},
		{"zero-with-nonzero-numerators", Coverage{Covered: 4, OutOfScope: 3, Unmapped: 2}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertPercentMatches(t, "Percent", tc.data.Percent(), 0)
			assertPercentMatches(t, "AcknowledgedPercent", tc.data.AcknowledgedPercent(), 0)
			assertPercentMatches(t, "MappingPercent", tc.data.MappingPercent(), 0)
		})
	}
}

// TestPercentHelpersComputeRatios proves every numerator is divided by the total.
func TestPercentHelpersComputeRatios(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		data        Coverage
		wantPercent float64
		wantAck     float64
		wantMapping float64
	}{
		{"mixed", Coverage{Total: 8, Covered: 3, OutOfScope: 2, Unmapped: 2}, 37.5, 62.5, 75},
		{"one-covered", Coverage{Total: 4, Covered: 1}, 25, 25, 100},
		{"fully-covered", Coverage{Total: 4, Covered: 4}, 100, 100, 100},
		{"fully-acknowledged", Coverage{Total: 4, OutOfScope: 4}, 0, 100, 100},
		{"fully-unmapped", Coverage{Total: 4, Unmapped: 4}, 0, 0, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertPercentMatches(t, "Percent", tc.data.Percent(), tc.wantPercent)
			assertPercentMatches(t, "AcknowledgedPercent", tc.data.AcknowledgedPercent(), tc.wantAck)
			assertPercentMatches(t, "MappingPercent", tc.data.MappingPercent(), tc.wantMapping)
		})
	}
}

// TestUncoveredFiltersExcludedStatusesKeepsOrder proves only unobserved mapped rows survive.
func TestUncoveredFiltersExcludedStatusesKeepsOrder(t *testing.T) {
	t.Parallel()
	entries := []Entry{
		{Node: "exact-gap", Status: grammar.StatusExact, Observed: false},
		{Node: "exact-hit", Status: grammar.StatusExact, Observed: true},
		{Node: "scope-gap", Status: grammar.StatusOutOfScope, Observed: false},
		{Node: "unmapped-gap", Status: grammar.StatusUnmapped, Observed: false},
		{Node: "degraded-gap", Status: grammar.StatusDegraded, Observed: false},
		{Node: "custom-gap", Status: "custom", Observed: false},
		{Node: "degraded-hit", Status: grammar.StatusDegraded, Observed: true},
	}
	want := []string{"exact-gap", "degraded-gap", "custom-gap"}
	got := Coverage{Entries: entries}.Uncovered()
	if !slices.Equal(entryNodeNames(got), want) {
		t.Fatalf("Uncovered names = %v, want %v", entryNodeNames(got), want)
	}
}

// TestUncoveredOnEmptyInventory proves an empty inventory yields an empty non-nil slice.
func TestUncoveredOnEmptyInventory(t *testing.T) {
	t.Parallel()
	got := Coverage{}.Uncovered()
	if got == nil {
		t.Fatal("Uncovered returned a nil slice")
	}
	if len(got) != 0 {
		t.Fatalf("Uncovered returned %d rows, want none", len(got))
	}
}

// TestObservedKindsOfDeduplicatesAscendingKinds proves deduplication and ascending order for upstream ordered rows.
func TestObservedKindsOfDeduplicatesAscendingKinds(t *testing.T) {
	t.Parallel()
	entries := []Entry{
		{Node: "import", Kind: air.KindImport, Observed: true},
		{Node: "function", Kind: air.KindFuncDecl, Observed: true},
		{Node: "function-again", Kind: air.KindFuncDecl, Observed: true},
		{Node: "block", Kind: air.KindBlock, Observed: false},
		{Node: "call", Kind: air.KindCall, Observed: true},
		{Node: "ident", Kind: air.KindIdent, Observed: false},
	}
	want := []air.NodeKind{air.KindImport, air.KindFuncDecl, air.KindCall}
	got := Coverage{Entries: entries}.ObservedKindsOf()
	if !slices.Equal(got, want) {
		t.Fatalf("ObservedKindsOf = %v, want %v", got, want)
	}
	if !slices.IsSorted(got) {
		t.Fatalf("ObservedKindsOf is not ascending: %v", got)
	}
}

// TestObservedKindsOfFollowsEntryOrder proves the helper keeps first-seen order instead of sorting.
func TestObservedKindsOfFollowsEntryOrder(t *testing.T) {
	t.Parallel()
	entries := []Entry{
		{Node: "ident", Kind: air.KindIdent, Observed: true},
		{Node: "import", Kind: air.KindImport, Observed: true},
		{Node: "ident-again", Kind: air.KindIdent, Observed: true},
		{Node: "unobserved", Kind: air.KindReturn, Observed: false},
	}
	want := []air.NodeKind{air.KindIdent, air.KindImport}
	got := Coverage{Entries: entries}.ObservedKindsOf()
	if !slices.Equal(got, want) {
		t.Fatalf("ObservedKindsOf = %v, want %v", got, want)
	}
}

// TestObservedKindsOfOnEmptyInventory proves an empty inventory yields an empty non-nil slice.
func TestObservedKindsOfOnEmptyInventory(t *testing.T) {
	t.Parallel()
	got := Coverage{}.ObservedKindsOf()
	if got == nil {
		t.Fatal("ObservedKindsOf returned a nil slice")
	}
	if len(got) != 0 {
		t.Fatalf("ObservedKindsOf returned %d kinds, want none", len(got))
	}
}
