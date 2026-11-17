/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package metrics

import (
	"reflect"
	"strings"
	"testing"
)

// recordOptions is the configuration of one synthetic drilldown record.
type recordOptions struct {
	language   string
	path       string
	digest     string
	operator   string
	class      string
	outcome    string
	captured   bool
	equivalent bool
}

// record builds one drilldown record for the aggregation tests.
func record(o recordOptions) DrilldownRecord {
	return DrilldownRecord{
		Language: o.language, FileSHA256: o.digest, RelPath: o.path,
		OperatorID: o.operator, Class: o.class, Offset: 4, EndOffset: 6,
		BeforeText: "a+b", AfterText: "a-b",
		AIRIDBefore: "before-" + o.digest, AIRIDAfter: "after-" + o.digest,
		Captured: o.captured, Equivalent: o.equivalent,
		EquivalentProofID: "EQ-LITERAL-RADIX-001", OutcomeKind: o.outcome,
	}
}

// TestAggregateFourColumnsRecomputesExactly proves every aggregate value is
// derived from the drilldown records and that the four columns stay separate.
func TestAggregateFourColumnsRecomputesExactly(t *testing.T) {
	t.Parallel()
	options := []recordOptions{
		{language: "go", path: "a.go", digest: "aa", operator: "cmp", class: ColumnText, outcome: OutcomeOK, captured: true},
		{language: "go", path: "a.go", digest: "aa", operator: "cmp", class: ColumnText, outcome: OutcomeOK},
		{language: "py", path: "b.py", digest: "bb", operator: "op", class: ColumnStructure, outcome: OutcomeOK, captured: true},
		{language: "c", path: "c.c", digest: "cc", operator: "digit", class: ColumnBroken, outcome: OutcomeTimeout},
		{language: "rs", path: "d.rs", digest: "dd", operator: "literal", class: ColumnEquivalent, outcome: OutcomeOK, captured: true, equivalent: true},
		{language: "kt", path: "e.kt", digest: "ee", operator: "cond", class: ColumnText, outcome: OutcomeOK, captured: true, equivalent: true},
		{language: "js", path: "f.js", digest: "ff", operator: "call", class: "unknown", outcome: OutcomePanic},
		{language: "cs", path: "g.cs", digest: "gg", operator: "block", class: ColumnStructure, outcome: OutcomeMemExceeded},
		{language: "go", path: "h.go", digest: "hh", operator: "return", class: ColumnBroken, outcome: OutcomeDeterminismFailure},
		{language: "py", path: "i.py", digest: "ii", operator: "cmp", class: ColumnText, outcome: OutcomeUnviable},
	}
	recs := make([]DrilldownRecord, 0, len(options))
	for _, option := range options {
		recs = append(recs, record(option))
	}
	sum := Aggregate(recs)
	if sum.Equivalence.Gross != len(recs) || sum.Equivalence.EquivalentCount != 2 ||
		sum.Equivalence.Effective != len(recs)-2 || sum.Equivalence.SilentDrops != 0 {
		t.Fatalf("equivalence summary wrong: %+v", sum.Equivalence)
	}
	if sum.Columns.Text.Gross != 3 || sum.Columns.Structure.Gross != 2 ||
		sum.Columns.Broken.Gross != 3 || sum.Columns.Equivalent.Gross != 2 {
		t.Fatalf("four columns wrong: %+v", sum.Columns)
	}
	if sum.Columns.Equivalent.Effective != 0 || sum.Columns.Text.Effective != 3 {
		t.Fatalf("effective columns wrong: %+v", sum.Columns)
	}
	if !reflect.DeepEqual(sum, Aggregate(slicesClone(recs))) {
		t.Fatal("aggregation is not reproducible from the same records")
	}
	for _, row := range sum.Mutation {
		if err := row.Validate(); err != nil {
			t.Fatalf("row %s|%s|%s: %v", row.Language, row.OperatorID,
				row.Class, err)
		}
		if row.SilentDrops != 0 {
			t.Fatal("a mutation was silently dropped")
		}
	}
	if !sum.Robustness.Complete() || sum.Robustness.ProcessedFiles != 9 {
		t.Fatalf("robustness wrong: %+v", sum.Robustness)
	}
	if sum.Robustness.TimeoutCount != 1 || sum.Robustness.OOMCount != 1 ||
		sum.Robustness.DeterminismFailureCount != 1 {
		t.Fatalf("robustness causes wrong: %+v", sum.Robustness)
	}
	if !strings.Contains(sum.Robustness.PanicLocalizer, "file=ff") ||
		!strings.Contains(sum.Robustness.PanicLocalizer, "airBefore=before-ff") {
		t.Fatalf("panic localizer wrong: %q", sum.Robustness.PanicLocalizer)
	}
}

// TestAggregateEmptyKeepsZeroPercentages proves an empty record set yields
// zeroed columns without a division by zero.
func TestAggregateEmptyKeepsZeroPercentages(t *testing.T) {
	t.Parallel()
	sum := Aggregate(nil)
	if sum.Columns.Text.CapturePct != 0 || sum.Columns.Equivalent.CapturePct != 0 ||
		len(sum.Mutation) != 0 || sum.Equivalence.Gross != 0 ||
		sum.Equivalence.SilentDrops != 0 {
		t.Fatalf("empty aggregate wrong: %+v", sum)
	}
	if sum.Robustness.ProcessedFiles != 0 || !sum.Robustness.Complete() {
		t.Fatalf("empty robustness wrong: %+v", sum.Robustness)
	}
}

// TestPrecisionRowValidate proves precision rows reject out-of-range values.
func TestPrecisionRowValidate(t *testing.T) {
	t.Parallel()
	good := PrecisionRow{Language: "go", ZeroDiagnosticZeroError: 3,
		Diagnostics: 1, ParseErrors: 0, DegradeNodePct: 2.5,
		OutsideDamageRetentionPct: 99.5, ConstructCoveragePct: 87}
	if err := good.Validate(); err != nil {
		t.Fatalf("valid row rejected: %v", err)
	}
	bad := []PrecisionRow{
		{Language: "", DegradeNodePct: 1},
		{Language: "go", Diagnostics: -1},
		{Language: "go", DegradeNodePct: -1},
		{Language: "go", OutsideDamageRetentionPct: 101},
		{Language: "go", ConstructCoveragePct: 100.5},
	}
	for i, row := range bad {
		if err := row.Validate(); err == nil {
			t.Fatalf("invalid row %d accepted: %+v", i, row)
		}
	}
}

// TestMutationRowValidate proves the strict denominator invariants.
func TestMutationRowValidate(t *testing.T) {
	t.Parallel()
	good := MutationRow{Language: "go", OperatorID: "cmp", Class: ColumnText,
		Gross: 5, Effective: 4, Equivalent: 1, Captured: 3, Missed: 1,
		Timeout: 1, CapturePct: 60, ValidOnlyPct: 75}
	if err := good.Validate(); err != nil {
		t.Fatalf("valid row rejected: %v", err)
	}
	bad := []MutationRow{
		{Language: "go", OperatorID: "cmp", Class: ColumnText, Gross: 1},
		{Language: "go", OperatorID: "cmp", Class: ColumnText, Gross: 2,
			Effective: 2, Captured: 1, Missed: 1, CapturePct: 101},
		{Language: "go", OperatorID: "cmp", Class: ColumnText, Gross: 1,
			Effective: 1, SilentDrops: 1},
		{Language: "", OperatorID: "cmp", Class: ColumnText},
	}
	for i, row := range bad {
		if err := row.Validate(); err == nil {
			t.Fatalf("invalid row %d accepted: %+v", i, row)
		}
	}
}

// TestDrilldownKeysAreDeterministic proves keys and localizers carry the
// provenance needed by the recompute and zero-tolerance paths.
func TestDrilldownKeysAreDeterministic(t *testing.T) {
	t.Parallel()
	rec := record(recordOptions{language: "go", path: "x.go",
		digest: "deadbeef", operator: "cmp", class: ColumnText,
		outcome: OutcomeOK, captured: true})
	if rec.Key() != "go|x.go|cmp|4" || rec.GroupKey() != "go|cmp|text" {
		t.Fatalf("keys wrong: %q / %q", rec.Key(), rec.GroupKey())
	}
	if rec.ColumnKey() != ColumnText {
		t.Fatalf("column key wrong: %q", rec.ColumnKey())
	}
	equivalent := rec
	equivalent.Equivalent = true
	if equivalent.ColumnKey() != ColumnEquivalent {
		t.Fatal("equivalent record must be the fourth column")
	}
	equivalent.Class = ColumnEquivalent
	if equivalent.ColumnKey() != ColumnEquivalent {
		t.Fatal("equivalent class must be the fourth column")
	}
	unknown := rec
	unknown.Class = "unknown"
	if unknown.ColumnKey() != ColumnBroken {
		t.Fatal("unknown class must not be dropped")
	}
	if !strings.Contains(rec.Localizer(), "file=deadbeef") ||
		!strings.Contains(rec.Localizer(), "operator=cmp") ||
		!strings.Contains(rec.Localizer(), "airAfter=after-deadbeef") {
		t.Fatalf("localizer wrong: %q", rec.Localizer())
	}
}

// slicesClone returns an independent copy so the reproducibility check uses
// fresh backing storage.
func slicesClone(in []DrilldownRecord) []DrilldownRecord {
	out := make([]DrilldownRecord, len(in))
	copy(out, in)
	return out
}
