/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package metrics

import "fmt"

// Outcome labels recorded on one drilldown row.
const (
	// OutcomeOK marks a completed parse.
	OutcomeOK = "OK"
	// OutcomeTimeout marks a parse that exceeded the time budget.
	OutcomeTimeout = "Timeout"
	// OutcomeMemExceeded marks a parse that exceeded the memory budget.
	OutcomeMemExceeded = "MemExceeded"
	// OutcomeContractError marks a frontend contract failure.
	OutcomeContractError = "ContractError"
	// OutcomeUnviable marks a mutation that could not be parsed.
	OutcomeUnviable = "Unviable"
	// OutcomePanic marks a run that stopped before the report was complete.
	OutcomePanic = "Panic"
	// OutcomeDeterminismFailure marks a mismatched double run.
	OutcomeDeterminismFailure = "DeterminismFailure"
)

// DrilldownRecord is the full recomputation input for exactly one mutation.
// FileSHA256, OperatorID, AIRIDBefore and AIRIDAfter make every aggregate value
// recomputable from the records alone; SilentDrops is always zero.
type DrilldownRecord struct {
	Language          string
	FileSHA256        string
	RelPath           string
	OperatorID        string
	Class             string
	Offset            int
	EndOffset         int
	BeforeText        string
	AfterText         string
	AIRIDBefore       string
	AIRIDAfter        string
	Captured          bool
	Equivalent        bool
	EquivalentProofID string
	OutcomeKind       string
	AllocDeltaBytes   uint64
	SilentDrops       int
}

// Key returns the deterministic aggregation key of the record. The key never
// contains a wall-clock value, an absolute path or an environment-derived value.
func (r DrilldownRecord) Key() string {
	return fmt.Sprintf("%s|%s|%s|%d", r.Language, r.RelPath,
		r.OperatorID, r.Offset)
}

// Localizer renders the crash or zero-tolerance locator of the record. It
// contains the file digest, the operator and both AIR identifiers.
func (r DrilldownRecord) Localizer() string {
	return fmt.Sprintf("file=%s|operator=%s|airBefore=%s|airAfter=%s",
		r.FileSHA256, r.OperatorID, r.AIRIDBefore, r.AIRIDAfter)
}

// ColumnKey returns the four-way column the record belongs to. An equivalent
// record is always the fourth column, independent of its class label.
func (r DrilldownRecord) ColumnKey() string {
	if r.Equivalent || r.Class == ColumnEquivalent {
		return ColumnEquivalent
	}
	switch r.Class {
	case ColumnText:
		return ColumnText
	case ColumnStructure:
		return ColumnStructure
	default:
		return ColumnBroken
	}
}

// GroupKey returns the mutation-row grouping key of the record.
func (r DrilldownRecord) GroupKey() string {
	return fmt.Sprintf("%s|%s|%s", r.Language, r.OperatorID, r.Class)
}
