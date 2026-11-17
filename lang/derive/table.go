/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"strings"

	"abate/lang/air"
	"abate/lang/rule"
)

// derivedDeleted is the lexical token produced by a remove-call mutation.
const derivedDeleted = "<deleted>"

// derivedPrefix is the fixed identifier prefix of every derived operator.
const derivedPrefix = "m-derived-"

// derivationRow binds one transform to its operator class, lexical token pair
// and AIR node kind. The five rows are the single derivation truth source.
type derivationRow struct {
	transform  Transform
	class      string
	from       string
	to         string
	targetKind air.NodeKind
	invertible bool
}

// derivationTable returns the five derivation rows in a fixed declaration
// order, rebuilt per call so no mutable package state is shared.
func derivationTable() []derivationRow {
	return []derivationRow{
		{TransformBoundary, "boundary-value", "<", "<=", air.KindBinary, true},
		{TransformConstant, "constant-perturb", "1", "2", air.KindNumLit, true},
		{TransformCondition, "negate-condition", "==", "!=", air.KindBinary, true},
		{TransformType, "return-value", "true", "false", air.KindBoolLit, true},
		{TransformCall, "remove-call", "call", derivedDeleted, air.KindCall, false},
	}
}

// transformEntryFor looks up one row by its transform.
func transformEntryFor(transform Transform) (derivationRow, bool) {
	for _, row := range derivationTable() {
		if row.transform == transform {
			return row, true
		}
	}
	return derivationRow{}, false
}

// transformFor classifies one rule specification into a derivation strategy.
// An explicit mutation kind wins; otherwise the rule category decides, so the
// mapping never depends on a concrete language identifier.
func transformFor(spec rule.Spec) (Transform, bool) {
	kind := strings.ToLower(strings.TrimSpace(spec.Mutation.Kind))
	for _, row := range derivationTable() {
		if kind != "" && strings.HasPrefix(kind, string(row.transform)) {
			return row.transform, true
		}
	}
	switch spec.Category {
	case rule.CategoryBoundary:
		return TransformBoundary, true
	case rule.CategoryPerformance:
		return TransformConstant, true
	case rule.CategoryErrorProne:
		return TransformCondition, true
	case rule.CategoryConcurrency:
		return TransformCall, true
	case rule.CategoryStyle, rule.CategoryNaming:
		return TransformType, true
	default:
		return "", false
	}
}
