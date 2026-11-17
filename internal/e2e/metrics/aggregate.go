/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package metrics

import "slices"

type (
	// Equivalence is the strict equivalence summary of one run. Gross is every
	// mutation, EquivalentCount is the number of equivalent ones, Effective is
	// Gross minus EquivalentCount, and SilentDrops is always zero. Equivalent
	// mutations are tagged, never deleted from the denominator.
	Equivalence struct {
		Gross           int
		Effective       int
		EquivalentCount int
		SilentDrops     int
	}

	// Summary is the complete output of the pure aggregation pass. It carries
	// the four separate columns, the deterministic mutation rows, the
	// independent robustness counters and the equivalence summary.
	Summary struct {
		Columns     Columns
		Mutation    []MutationRow
		Robustness  Robustness
		Equivalence Equivalence
	}
)

// rowAcc accumulates the counters of one mutation row.
type rowAcc struct {
	language   string
	operatorID string
	class      string
	gross      int
	captured   int
	timeout    int
	unviable   int
	equivalent int
}

// Aggregate recomputes every reported counter from drilldown records. It is a
// pure function: no input/output, no wall clock, no map iteration in the output
// order, and no weighted, offsetting or combined score of the four columns.
// Every record enters Gross exactly once; no record is skipped or excluded.
func Aggregate(recs []DrilldownRecord) Summary {
	ordered := orderedRecords(recs)
	out := Summary{}
	rows := make(map[string]*rowAcc)
	files := make(map[string]struct{})
	equivalentCount := 0
	for _, rec := range ordered {
		column := out.Columns.column(rec.ColumnKey())
		column.Gross++
		if rec.Captured {
			column.Captured++
		}
		if rec.Equivalent || rec.Class == ColumnEquivalent {
			equivalentCount++
		}
		accumulateRow(rows, rec)
		accumulateRobustness(&out.Robustness, files, rec)
	}
	out.Columns = out.Columns.finalize(equivalentCount)
	out.Mutation = finalizeRows(rows)
	out.Robustness = finalizeRobustness(out.Robustness, files)
	out.Equivalence = Equivalence{
		Gross:           len(ordered),
		EquivalentCount: equivalentCount,
		Effective:       len(ordered) - equivalentCount,
		SilentDrops:     0,
	}
	return out
}

// orderedRecords returns records sorted by their deterministic grouping key.
func orderedRecords(recs []DrilldownRecord) []DrilldownRecord {
	out := slices.Clone(recs)
	slices.SortFunc(out, func(a, b DrilldownRecord) int {
		left, right := a.Key(), b.Key()
		if left < right {
			return -1
		}
		if left > right {
			return 1
		}
		return 0
	})
	return out
}

// accumulateRow folds one record into its (language, operator, class) row.
func accumulateRow(rows map[string]*rowAcc, rec DrilldownRecord) {
	key := rec.GroupKey()
	acc, found := rows[key]
	if !found {
		acc = &rowAcc{language: rec.Language,
			operatorID: rec.OperatorID, class: rec.Class}
		rows[key] = acc
	}
	acc.gross++
	switch {
	case rec.Captured:
		acc.captured++
	case rec.OutcomeKind == OutcomeTimeout:
		acc.timeout++
	case rec.OutcomeKind == OutcomeUnviable:
		acc.unviable++
	}
	if rec.Equivalent || rec.Class == ColumnEquivalent {
		acc.equivalent++
	}
}

// finalizeRows renders every row in ascending group-key order.
func finalizeRows(rows map[string]*rowAcc) []MutationRow {
	keys := make([]string, 0, len(rows))
	for key := range rows {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	out := make([]MutationRow, 0, len(keys))
	for _, key := range keys {
		acc := rows[key]
		viable := acc.gross - acc.unviable
		row := MutationRow{
			Language:    acc.language,
			OperatorID:  acc.operatorID,
			Class:       acc.class,
			Gross:       acc.gross,
			Effective:   acc.gross - acc.equivalent,
			Equivalent:  acc.equivalent,
			Captured:    acc.captured,
			Missed:      acc.gross - acc.captured - acc.timeout - acc.unviable,
			Timeout:     acc.timeout,
			Unviable:    acc.unviable,
			SilentDrops: 0,
		}
		if acc.gross > 0 {
			row.CapturePct = float64(acc.captured) * 100 / float64(acc.gross)
		}
		if viable > 0 {
			row.ValidOnlyPct = float64(acc.captured) * 100 / float64(viable)
		}
		out = append(out, row)
	}
	return out
}

// accumulateRobustness folds one record into the robustness counters.
func accumulateRobustness(r *Robustness, files map[string]struct{},
	rec DrilldownRecord) {
	files[rec.FileSHA256] = struct{}{}
	switch rec.OutcomeKind {
	case OutcomeTimeout:
		r.TimeoutCount++
	case OutcomeMemExceeded:
		r.OOMCount++
	case OutcomeDeterminismFailure:
		r.DeterminismFailureCount++
	case OutcomePanic:
		if r.PanicLocalizer == "" {
			r.PanicLocalizer = rec.Localizer()
		}
	}
}

// finalizeRobustness sets the file counters from the drilldown set. Every
// record names a registered file, so processed and registered are identical
// unless the runner supplies an external manifest count to compare against.
func finalizeRobustness(r Robustness, files map[string]struct{}) Robustness {
	r.ProcessedFiles = len(files)
	r.RegisteredFiles = len(files)
	return r
}
