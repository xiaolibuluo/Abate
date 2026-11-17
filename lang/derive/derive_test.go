/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/kernel/algebra"
	"abate/kernel/mutation"
	"abate/lang/air"
	"abate/lang/rule"
)

// TestDerivationTableFiveClasses locks the single five-row truth source.
func TestDerivationTableFiveClasses(t *testing.T) {
	t.Parallel()
	want := []struct {
		transform  Transform
		class      string
		from       string
		to         string
		kind       air.NodeKind
		invertible bool
	}{
		{TransformBoundary, "boundary-value", "<", "<=", air.KindBinary, true},
		{TransformConstant, "constant-perturb", "1", "2", air.KindNumLit, true},
		{TransformCondition, "negate-condition", "==", "!=", air.KindBinary, true},
		{TransformType, "return-value", "true", "false", air.KindBoolLit, true},
		{TransformCall, "remove-call", "call", derivedDeleted, air.KindCall, false},
	}
	rows := derivationTable()
	if len(rows) != len(want) {
		t.Fatalf("derivation table has %d rows, want %d", len(rows), len(want))
	}
	for i, expected := range want {
		row := rows[i]
		if row.transform != expected.transform || row.class != expected.class ||
			row.from != expected.from || row.to != expected.to ||
			row.targetKind != expected.kind || row.invertible != expected.invertible {
			t.Fatalf("row %d = %+v, want %+v", i, row, expected)
		}
		if !expected.transform.Valid() {
			t.Fatalf("transform %s is not valid", expected.transform)
		}
		if expected.transform.String() != string(expected.transform) {
			t.Fatalf("transform %s has unstable wire text", expected.transform)
		}
		got, ok := transformEntryFor(expected.transform)
		if !ok || got.class != expected.class {
			t.Fatalf("lookup(%s) = %+v, %t", expected.transform, got, ok)
		}
	}
	if Transform("bogus").Valid() {
		t.Fatal("an unknown transform must not validate")
	}
	if _, ok := transformEntryFor(Transform("bogus")); ok {
		t.Fatal("an unknown transform must not resolve to a row")
	}
}

// TestDeriveFiveTransforms proves every class derives and output is sorted.
func TestDeriveFiveTransforms(t *testing.T) {
	t.Parallel()
	decls := []rule.Spec{
		specFor("GO-EG-001", rule.CategoryNaming, "boundary-shift"),
		specFor("GO-EG-002", rule.CategoryNaming, "constant-perturb"),
		specFor("GO-EG-003", rule.CategoryNaming, "condition-negate"),
		specFor("GO-EG-004", rule.CategoryNaming, "type-assert"),
		specFor("GO-EG-005", rule.CategoryNaming, "call-remove"),
	}
	rules := rulesOf(decls[3], decls[0], decls[4], decls[1], decls[2])
	got, err := NewDeriver().Derive(rules)
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}
	if len(got) != 5 {
		t.Fatalf("Derive returned %d specs, want 5", len(got))
	}
	want := []Transform{TransformBoundary, TransformConstant,
		TransformCondition, TransformType, TransformCall}
	for i, expected := range want {
		if got[i].Transform != expected {
			t.Fatalf("spec %d transform = %s, want %s", i, got[i].Transform, expected)
		}
		if !strings.HasPrefix(string(got[i].OperatorID), derivedPrefix) {
			t.Fatalf("spec %d id %q lacks the derived prefix", i, got[i].OperatorID)
		}
		row, ok := transformEntryFor(expected)
		if !ok {
			t.Fatalf("missing row for %s", expected)
		}
		if got[i].TargetKind != row.targetKind.String() {
			t.Fatalf("spec %d target kind = %q, want %q", i,
				got[i].TargetKind, row.targetKind.String())
		}
		if got[i].SourceRuleID != decls[i].ID {
			t.Fatalf("spec %d is not sorted by source rule id", i)
		}
	}
	for i := 1; i < len(got); i++ {
		if got[i-1].Canonical() == got[i].Canonical() {
			t.Fatal("two distinct rules produced identical specifications")
		}
	}
}

// TestDeriveDeterminismAndTwentyIDs proves a double derive hashes identically
// and yields far more than 20 distinct operator identifiers.
func TestDeriveDeterminismAndTwentyIDs(t *testing.T) {
	t.Parallel()
	rules := builtinRules(t)
	if len(rules) != 110 {
		t.Fatalf("built-in rule count = %d, want 110", len(rules))
	}
	first, err := NewDeriver().Derive(rules)
	if err != nil {
		t.Fatalf("first Derive: %v", err)
	}
	second, err := NewDeriver().Derive(rules)
	if err != nil {
		t.Fatalf("second Derive: %v", err)
	}
	if determinism.HashState(first) != determinism.HashState(second) {
		t.Fatal("two derives of one rule set produced different HashState values")
	}
	if len(first) < 20 {
		t.Fatalf("only %d specs derived, want at least 20", len(first))
	}
	ids := make(map[algebra.OperatorID]bool, len(first))
	for i, spec := range first {
		if !strings.HasPrefix(string(spec.OperatorID), derivedPrefix) {
			t.Fatalf("spec %d id %q lacks the derived prefix", i, spec.OperatorID)
		}
		if ids[spec.OperatorID] {
			t.Fatalf("duplicate operator id %s", spec.OperatorID)
		}
		ids[spec.OperatorID] = true
		if i > 0 && first[i-1].SourceRuleID > spec.SourceRuleID {
			t.Fatal("Derive output is not sorted by SourceRuleID")
		}
	}
	if len(ids) < 20 {
		t.Fatalf("only %d distinct derived ids", len(ids))
	}
	if err := Validate(first); err != nil {
		t.Fatalf("Validate: %v", err)
	}
	ops, err := Operators(first)
	if err != nil || len(ops) != len(first) {
		t.Fatalf("Operators = %d, %v", len(ops), err)
	}
}

// TestEngineReportDerivedVerdicts proves the real Engine.Run report carries
// every derived operator with an allocated verdict and stable bytes.
func TestEngineReportDerivedVerdicts(t *testing.T) {
	t.Parallel()
	unit := sampleUnit(t)
	rules := builtinRules(t)
	rep, specs, err := Experiment(42, unit, rules)
	if err != nil {
		t.Fatalf("Experiment: %v", err)
	}
	ids := DerivedIDs(rep)
	if len(ids) < 20 {
		t.Fatalf("Report.Points carries only %d distinct derived ids", len(ids))
	}
	if len(ids) != len(specs) {
		t.Fatalf("Report has %d derived ids but %d specs were derived", len(ids), len(specs))
	}
	if !slices.IsSorted(ids) {
		t.Fatal("DerivedIDs is not sorted")
	}
	for i, point := range rep.Points {
		if point.Operator == nil {
			t.Fatalf("point %d has no operator", i)
		}
	}
	if len(rep.Points) != len(rep.Verdicts) ||
		len(rep.Points) != len(rep.KillingTests) {
		t.Fatal("report columns are not aligned")
	}
	for i, verdict := range rep.Verdicts {
		if !allocatedVerdict(verdict) {
			t.Fatalf("verdict %d = %q is not an allocated state", i, verdict)
		}
	}
	if rep.Score.Total != uint64(len(rep.Points)) {
		t.Fatal("mutation score total disagrees with the point count")
	}
	if mutation.ReportFingerprint(rep) != rep.Fingerprint {
		t.Fatal("report fingerprint disagrees with its canonical body")
	}
	determinism.MustMatch100(t, 42, func() []byte { return rep.Bytes() })
}

// TestDeriveEmptyAndNilRule covers the empty and nil rule boundaries.
func TestDeriveEmptyAndNilRule(t *testing.T) {
	t.Parallel()
	if _, err := NewDeriver().Derive(nil); !errors.Is(err, ErrEmptyRuleSet) {
		t.Fatalf("empty error = %v, want ErrEmptyRuleSet", err)
	}
	if _, err := NewDeriver().Derive([]rule.Rule{nil}); !errors.Is(err, ErrInvalidSpec) {
		t.Fatalf("nil rule error = %v, want ErrInvalidSpec", err)
	}
	if _, err := NewDeriver().Derive(rulesOf(
		specFor("GO-EG-001", rule.Category("unknown"), ""),
	)); !errors.Is(err, ErrUnknownTransform) {
		t.Fatalf("unknown category error = %v, want ErrUnknownTransform", err)
	}
}

// TestMutationSpecOperatorAndCanonical checks the built operator and its
// byte-stable canonical rendering.
func TestMutationSpecOperatorAndCanonical(t *testing.T) {
	t.Parallel()
	spec := MutationSpec{
		SourceRuleID: "GO-EG-001",
		OperatorID:   algebra.FormatOperatorID(algebra.FamilyMutation, "derived-go-eg-001", "1"),
		TargetKind:   air.KindBinary.String(),
		Transform:    TransformBoundary,
	}
	op, err := spec.Operator()
	if err != nil {
		t.Fatalf("Operator: %v", err)
	}
	if op.ID() != spec.OperatorID || op.Family() != algebra.FamilyMutation {
		t.Fatal("built operator does not match the specification")
	}
	params := op.Params()
	for _, key := range []string{"class", "from", "to", "derived-from"} {
		if params[key] == "" {
			t.Fatalf("operator is missing required param %q", key)
		}
	}
	if params["class"] != "boundary-value" || params["from"] != "<" ||
		params["to"] != "<=" || params["derived-from"] != "GO-EG-001" {
		t.Fatalf("unexpected operator params: %v", params)
	}
	first := spec.Canonical()
	if first != spec.Canonical() {
		t.Fatal("Canonical is unstable")
	}
	changed := spec
	changed.SourceRuleID = "GO-EG-002"
	if changed.Canonical() == first {
		t.Fatal("Canonical did not detect a SourceRuleID change")
	}
	if _, err := (MutationSpec{Transform: Transform("bogus")}).Operator(); !errors.Is(err, ErrUnknownTransform) {
		t.Fatalf("bad transform error = %v, want ErrUnknownTransform", err)
	}
}
