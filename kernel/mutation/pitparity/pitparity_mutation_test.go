/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"errors"
	"testing"

	"abate/kernel/mutation"
	"abate/kernel/verify"
)

// TestEngineExercisesEveryOperatorClass proves the parity fixture drives the
// whole M family operator set through the real mutation engine.
func TestEngineExercisesEveryOperatorClass(t *testing.T) {
	t.Parallel()
	target, suite := newFixture()
	registry, err := mutation.NewRegistryWith(nil)
	if err != nil {
		t.Fatalf("registry: %v", err)
	}
	report, err := runEngine(registry, target, suite)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if len(report.Points) == 0 {
		t.Fatal("the engine produced no mutation point")
	}
	classes := make(map[string]bool, len(report.Points))
	for i, point := range report.Points {
		if point.Operator == nil {
			t.Fatalf("point %d carries no operator", i)
		}
		classes[point.Operator.Params()["class"]] = true
		switch report.Verdicts[i] {
		case verify.MutantKilled, verify.MutantSurvived,
			verify.MutantEquivalent, verify.MutantTimeout:
		default:
			t.Fatalf("point %d carries an unknown verdict %q", i, report.Verdicts[i])
		}
	}
	for _, class := range []string{
		"arith-replace", "boundary-value", "constant-perturb", "control-flow",
		"return-value", "negate-condition", "remove-call",
	} {
		if !classes[class] {
			t.Fatalf("the fixture does not exercise the %s class", class)
		}
	}
	if report.Score.Killed == 0 {
		t.Fatal("the engine killed no mutant of the parity fixture")
	}
	if report.Score.Equivalent == 0 {
		t.Fatal("the fixture lost its equivalent mutant branch")
	}
	if report.Score.Timeout == 0 {
		t.Fatal("the fixture lost its timeout branch")
	}
}

// TestEngineFailuresAreReported covers the kernel failure branches of the
// parity run, so a broken fixture never turns into a silent pass.
func TestEngineFailuresAreReported(t *testing.T) {
	t.Parallel()
	target, suite := newFixture()
	registry, err := mutation.NewRegistryWith(nil)
	if err != nil {
		t.Fatalf("registry: %v", err)
	}
	if _, err := runEngine(nil, target, suite); !errors.Is(err, mutation.ErrNilRegistry) {
		t.Fatalf("error = %v, want ErrNilRegistry", err)
	}
	if _, err := runEngine(registry, nil, suite); !errors.Is(err, mutation.ErrNilTarget) {
		t.Fatalf("error = %v, want ErrNilTarget", err)
	}
	if _, err := runEngine(registry, target, nil); !errors.Is(err, mutation.ErrNilSuite) {
		t.Fatalf("error = %v, want ErrNilSuite", err)
	}
}

// TestVerdictIndexRejectsBrokenRuns covers the index failure branches.
func TestVerdictIndexRejectsBrokenRuns(t *testing.T) {
	t.Parallel()
	registry, err := mutation.NewRegistryWith(nil)
	if err != nil {
		t.Fatalf("registry: %v", err)
	}
	operators := registry.For("+")
	if len(operators) == 0 {
		t.Fatal("the registry produced no arithmetic operator")
	}
	point := mutation.NewMutationPoint("a-arith", "+", "-", operators[0], 1, 1)
	other := mutation.NewMutationPoint("a-arith", "+", "*", operators[0], 1, 2)
	broken := mutation.Report{Points: []mutation.MutationPoint{point}}
	if _, err := newVerdictIndex(broken); !errors.Is(err, ErrMissingVerdict) {
		t.Fatalf("error = %v, want ErrMissingVerdict", err)
	}
	conflicting := mutation.Report{
		Points:   []mutation.MutationPoint{point, other},
		Verdicts: []verify.MutationVerdict{verify.MutantKilled, verify.MutantSurvived},
	}
	if _, err := newVerdictIndex(conflicting); !errors.Is(err, ErrMissingVerdict) {
		t.Fatalf("error = %v, want ErrMissingVerdict", err)
	}
	otherClass := registry.For("<")
	if len(otherClass) == 0 {
		t.Fatal("the registry produced no boundary operator")
	}
	mixed := mutation.NewMutationPoint("a-arith", "<", "<=", otherClass[0], 1, 3)
	nodeConflict := mutation.Report{
		Points:   []mutation.MutationPoint{point, mixed},
		Verdicts: []verify.MutationVerdict{verify.MutantKilled, verify.MutantSurvived},
	}
	if _, err := newVerdictIndex(nodeConflict); !errors.Is(err, ErrMissingVerdict) {
		t.Fatalf("error = %v, want ErrMissingVerdict", err)
	}
}

// TestUncoveredHelpersAreReachable pins the uncovered branch helpers.
func TestUncoveredHelpersAreReachable(t *testing.T) {
	t.Parallel()
	if got := uncoveredOriginal(3); got != "9" {
		t.Fatalf("uncovered original = %q, want 9", got)
	}
	if got := uncoveredDivided(9); got != "3" {
		t.Fatalf("uncovered divided = %q, want 3", got)
	}
	if got := uncoveredZeroReturn(9); got != "0" {
		t.Fatalf("uncovered zero return = %q, want 0", got)
	}
	if _, ok := witnessFor("nowhere", "NoMutator", 42); ok {
		t.Fatal("witnessFor found a case that was never registered")
	}
	if witnessMatchesStatus("unknown", StatusKilled) {
		t.Fatal("an unknown witness matched a known status")
	}
	if markerInconsistent != "否" {
		t.Fatal("the negative marker is wrong")
	}
	if markerConsistent != "是" {
		t.Fatal("the positive marker is wrong")
	}
}

// TestResolveSiteRejectsUnknownSites covers the site resolution failure.
func TestResolveSiteRejectsUnknownSites(t *testing.T) {
	t.Parallel()
	if _, _, err := resolveSite("nowhere", "NoMutator", 42); !errors.Is(err, ErrUnmappedMutation) {
		t.Fatalf("error = %v, want ErrUnmappedMutation", err)
	}
	target, suite := newFixture()
	index, err := newVerdictIndex(mutationReport(t, target, suite))
	if err != nil {
		t.Fatalf("index: %v", err)
	}
	unknown := Mutation{
		Status: StatusKilled, MutatedMethod: "nowhere",
		Mutator: "NoMutator", LineNumber: 42,
	}
	if _, err := compareRecord(unknown, index, suite); !errors.Is(err, ErrUnmappedMutation) {
		t.Fatalf("error = %v, want ErrUnmappedMutation", err)
	}
}

// TestMutantHelpersCoverEveryBranch pins the helper branches of the witness.
func TestMutantHelpersCoverEveryBranch(t *testing.T) {
	t.Parallel()
	if got := acceptsBoundary(6); got != "true" {
		t.Fatalf("accepts boundary = %q, want true", got)
	}
	if got := acceptsBoundary(4); got != "false" {
		t.Fatalf("accepts boundary = %q, want false", got)
	}
	if got := sumUntilSubtracted(4); got != "-3" {
		t.Fatalf("sum until subtracted = %q, want -3", got)
	}
	if got := sumUntilSubtracted(3); got != "-3" {
		t.Fatalf("sum until subtracted = %q, want -3", got)
	}
}
