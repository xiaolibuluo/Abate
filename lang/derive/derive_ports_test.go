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
	"testing"

	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/air"
)

// bareUnit builds an AIR unit whose only node carries no mutatable token, so
// the adapter must supply every representative anchor node itself.
func bareUnit(t *testing.T) *air.Unit {
	t.Helper()
	unit := air.NewUnit(air.NewPackageRef("sample", nil))
	root := air.NewNode(air.KindPackage, air.NewPosition(1, 1, 0, 1), "sample")
	if err := unit.AddFile(
		air.NewFileRef("bare.go", []byte("package sample\n")), root,
	); err != nil {
		t.Fatalf("AddFile: %v", err)
	}
	return unit
}

// TestAIRTargetAnchorsFillMissingTokens covers the synthetic anchor path.
func TestAIRTargetAnchorsFillMissingTokens(t *testing.T) {
	t.Parallel()
	specs, err := NewDeriver().Derive(builtinRules(t))
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}
	target, err := NewAIRTarget(bareUnit(t), specs, experimentTests())
	if err != nil {
		t.Fatalf("NewAIRTarget: %v", err)
	}
	nodes := target.Nodes()
	if len(nodes) == 0 || !slices.IsSorted(nodes) {
		t.Fatal("anchor nodes must be present and sorted")
	}
	for _, node := range nodes {
		if token := target.Token(node); token == "" {
			t.Fatalf("anchor node %q has no lexical token", node)
		}
	}
	again, err := NewAIRTarget(bareUnit(t), specs, experimentTests())
	if err != nil {
		t.Fatalf("second NewAIRTarget: %v", err)
	}
	if !slices.Equal(nodes, again.Nodes()) ||
		target.Fingerprint() != again.Fingerprint() {
		t.Fatal("anchor construction is not deterministic")
	}
}

// TestAIRTargetStaticBranches exercises every branch of the static kill
// predicate and the equivalence check, including an unknown node reference.
func TestAIRTargetStaticBranches(t *testing.T) {
	t.Parallel()
	custom := &AIRTarget{refs: []airNodeRef{{
		key: "n", token: "<", kind: air.KindBinary, assertion: "t2",
	}}}
	changed := mutation.MutationPoint{Node: "n", Original: "a", Mutated: "b"}
	unchanged := mutation.MutationPoint{Node: "n", Original: "a", Mutated: "a"}
	unknown := mutation.MutationPoint{Node: "missing", Original: "a", Mutated: "b"}
	if !custom.Equivalent(unchanged) || custom.Equivalent(changed) {
		t.Fatal("static equivalence decision is wrong")
	}
	if custom.Equivalent(unknown) {
		t.Fatal("an unknown node must not be reported as equivalent")
	}
	if custom.Evaluate(changed, "t2") != verify.OutcomeViolated {
		t.Fatal("a covered assertion must violate once the threshold is reached")
	}
	if custom.Evaluate(changed, "t1") != verify.OutcomeHeld {
		t.Fatal("a test below the assertion threshold must hold")
	}
	if custom.Evaluate(unknown, "t2") != verify.OutcomeHeld {
		t.Fatal("an unknown node must hold")
	}
	if custom.Evaluate(unchanged, "t2") != verify.OutcomeHeld {
		t.Fatal("an unchanged mutant must hold")
	}
	if custom.Token("missing") != "" || custom.Token("n") != "<" {
		t.Fatal("token lookup does not use the frozen projection")
	}
}

// TestDerivedPointHelpers cover nil-operator points in the report projections.
func TestDerivedPointHelpers(t *testing.T) {
	t.Parallel()
	onlyNil := mutation.Report{Points: []mutation.MutationPoint{{Operator: nil}}}
	if hasDerivedPoints(onlyNil) || len(DerivedIDs(onlyNil)) != 0 {
		t.Fatal("a nil-operator point must not count as a derived mutant")
	}
	op, err := validSpec().Operator()
	if err != nil {
		t.Fatalf("Operator: %v", err)
	}
	mixed := mutation.Report{Points: []mutation.MutationPoint{
		{Operator: nil}, {Operator: op}, {Operator: op},
	}}
	if !hasDerivedPoints(mixed) {
		t.Fatal("the derived operator was not detected")
	}
	ids := DerivedIDs(mixed)
	if len(ids) != 1 || ids[0] != string(op.ID()) {
		t.Fatalf("DerivedIDs = %v, want one distinct id", ids)
	}
}

// TestRunExperimentErrorBranches covers the early rejection paths of the
// engine assembly helpers.
func TestRunExperimentErrorBranches(t *testing.T) {
	t.Parallel()
	if _, err := RunExperiment(42, nil, nil, nil); !errors.Is(err, ErrEmptyRuleSet) {
		t.Fatalf("RunExperiment(empty) = %v, want ErrEmptyRuleSet", err)
	}
	if _, _, err := Experiment(42, nil, builtinRules(t)); !errors.Is(err, ErrNilUnit) {
		t.Fatalf("Experiment(nil unit) = %v, want ErrNilUnit", err)
	}
}

// TestOperatorsRejectIncompleteDerivedFrom proves an operator rejected by the
// kernel registry surfaces as the derivation sentinel.
func TestOperatorsRejectIncompleteDerivedFrom(t *testing.T) {
	t.Parallel()
	spec := validSpec()
	spec.SourceRuleID = ""
	_, err := Operators([]MutationSpec{spec})
	if !errors.Is(err, ErrUnknownOperator) {
		t.Fatalf("error = %v, want ErrUnknownOperator", err)
	}
	if !errors.Is(err, mutation.ErrInvalidDerived) {
		t.Fatal("the kernel sentinel must remain visible through errors.Is")
	}
}

// TestAnchorKindArithmetic covers the arithmetic representative mapping.
func TestAnchorKindArithmetic(t *testing.T) {
	t.Parallel()
	if anchorKind("+") != air.KindBinary {
		t.Fatal("the arithmetic anchor must map to a binary node")
	}
	if anchorKind("call") != air.KindCall || anchorKind("true") != air.KindBoolLit {
		t.Fatal("table-driven anchor kinds changed")
	}
}
