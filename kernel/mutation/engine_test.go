/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"context"
	"errors"
	"slices"
	"strings"
	"testing"

	"abate/kernel/algebra"
)

func TestRegistrySevenClasses(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	want := []string{"arith-replace", "boundary-value", "constant-perturb",
		"control-flow", "return-value", "negate-condition", "remove-call"}
	if !slices.Equal(reg.Classes(), want) {
		t.Fatalf("classes = %v, want %v", reg.Classes(), want)
	}
	all := reg.All()
	if len(all) != 7 {
		t.Fatalf("All() returned %d operators, want 7", len(all))
	}
	seen := make(map[algebra.OperatorID]bool)
	for _, op := range all {
		if op.Family() != algebra.FamilyMutation {
			t.Fatalf("operator %s has family %s", op.ID(), op.Family())
		}
		if seen[op.ID()] {
			t.Fatalf("duplicate operator id %s", op.ID())
		}
		seen[op.ID()] = true
		first := op.Fingerprint()
		if op.Fingerprint() != first {
			t.Fatalf("operator %s fingerprint is unstable", op.ID())
		}
		if !reg.Valid(op) {
			t.Fatalf("registry rejected its own operator %s", op.ID())
		}
	}
	if !slices.IsSortedFunc(all, byOperatorID) {
		t.Fatal("All() is not sorted by operator id")
	}
}

func TestOperatorPositiveNegativeCases(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	cases := []struct {
		class string
		token string
		to    string
	}{
		{"arith-replace", "+", "-"},
		{"boundary-value", "<", "<="},
		{"constant-perturb", "1", "2"},
		{"control-flow", "break", tokenDeleted},
		{"return-value", "true", "false"},
		{"negate-condition", "==", "!="},
		{"remove-call", "call", tokenDeleted},
	}
	for _, tc := range cases {
		ops := reg.For(tc.token)
		if len(ops) == 0 {
			t.Fatalf("class %s has no operator for token %q", tc.class, tc.token)
		}
		found := false
		for _, op := range ops {
			params := op.Params()
			if params["class"] == tc.class && params["to"] == tc.to && params["from"] == tc.token {
				found = true
			}
			params["from"] = "mutated-copy"
			if op.Params()["from"] == "mutated-copy" {
				t.Fatal("Params() leaked its internal map")
			}
		}
		if !found {
			t.Fatalf("class %s did not produce %q -> %q: %v", tc.class, tc.token, tc.to, ops)
		}
	}
	if reg.For("") != nil || len(reg.For("identifier")) != 0 {
		t.Fatal("non-mutatable token produced operators")
	}
	tokens := []string{"+", "-", "*", "/", "&&", "||", "<", "<=", ">", ">=",
		"1", "0", "break", "continue", "true", "false", "nil", "nonnil",
		"==", "!=", "call"}
	for _, token := range tokens {
		ops := reg.For(token)
		if len(ops) == 0 {
			t.Fatalf("token %q has no applicable operator", token)
		}
		changed := false
		for _, op := range ops {
			if op.Params()["to"] != token {
				changed = true
			}
		}
		if !changed {
			t.Fatalf("token %q produced only identical mutants", token)
		}
	}
}

func TestOperatorInverseBranches(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	laws := algebra.NewLaws()
	for _, op := range reg.All() {
		class := op.Params()["class"]
		dual, err := op.Inverse()
		if class == "control-flow" || class == "remove-call" {
			if !errors.Is(err, algebra.ErrNotInvertible) {
				t.Fatalf("class %s should be irreversible, got %v", class, err)
			}
			if op.Invertible() {
				t.Fatalf("class %s declared invertible", class)
			}
			continue
		}
		if err != nil || dual == nil {
			t.Fatalf("class %s inverse failed: %v", class, err)
		}
		if !laws.CheckDual(op) {
			t.Fatalf("class %s violates the dual law", class)
		}
	}
}

func TestSchedulePriorityDeterminism(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	ops := reg.All()
	points := make([]MutationPoint, 0, 12)
	for i, op := range ops {
		priority := uint64(i % 3) //nolint:gosec // bounded loop index.
		selector := uint64(i)     //nolint:gosec // bounded loop index.
		points = append(points, NewMutationPoint("n-"+string(rune('a'+i)), op.Params()["from"],
			op.Params()["to"], op, priority, selector))
	}
	points = append(points, points[0])
	first := Schedule(points)
	second := Schedule(points)
	if len(first) != len(ops) {
		t.Fatalf("Schedule kept %d points, want %d after dedup", len(first), len(ops))
	}
	for i := range first {
		if first[i].Fingerprint != second[i].Fingerprint || first[i].Priority != second[i].Priority {
			t.Fatalf("Schedule is not deterministic at index %d", i)
		}
		if i > 0 && first[i-1].Priority < first[i].Priority {
			t.Fatalf("coverage priority not descending at index %d", i)
		}
	}
}

func TestMinimizeEndToEnd(t *testing.T) {
	t.Parallel()
	point := NewMutationPoint("n", "+", "-", NewRegistry().All()[0], 1, 1)
	tests := []string{"t1", "t2", "t3", "t4", "t5"}
	index, err := Minimize(point, tests, func(_ MutationPoint, test string) bool {
		return test >= "t3"
	})
	if err != nil || index != 2 {
		t.Fatalf("Minimize = (%d, %v), want (2, nil)", index, err)
	}
	index, err = Minimize(point, tests, func(_ MutationPoint, _ string) bool { return false })
	if err != nil || index != -1 {
		t.Fatalf("no-kill Minimize = (%d, %v), want (-1, nil)", index, err)
	}
	index, err = Minimize(point, tests, func(_ MutationPoint, _ string) bool { return true })
	if err != nil || index != 0 {
		t.Fatalf("all-kill Minimize = (%d, %v), want (0, nil)", index, err)
	}
	index, err = Minimize(point, nil, func(_ MutationPoint, _ string) bool { return true })
	if err != nil || index != -1 {
		t.Fatalf("empty Minimize = (%d, %v), want (-1, nil)", index, err)
	}
	if _, err := Minimize(point, tests, nil); !errors.Is(err, ErrNilKill) {
		t.Fatalf("nil predicate error = %v, want ErrNilKill", err)
	}
}

func TestEngineRunBasic(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	engine := NewEngine(42)
	rep, err := engine.Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}
	if engine.Seed() != 42 || engine.Registry() == nil {
		t.Fatal("engine accessors disagree with construction")
	}
	if rep.Score.Total != uint64(len(rep.Points)) || len(rep.Verdicts) != len(rep.Points) ||
		len(rep.KillingTests) != len(rep.Points) {
		t.Fatalf("report lengths disagree: points=%d verdicts=%d killing=%d score=%+v",
			len(rep.Points), len(rep.Verdicts), len(rep.KillingTests), rep.Score)
	}
	if rep.Fingerprint != ReportFingerprint(rep) {
		t.Fatal("report fingerprint disagrees with canonical body")
	}
	if !strings.Contains(rep.Describe(), "seed=42") {
		t.Fatalf("Describe lost the seed: %s", rep.Describe())
	}
	if len(rep.Bytes()) < 64 {
		t.Fatalf("canonical encoding too short: %d", len(rep.Bytes()))
	}
	for _, point := range rep.Points {
		if point.Operator == nil || !engine.Registry().Valid(point.Operator) {
			t.Fatalf("point has an invalid operator: %+v", point)
		}
	}
}

func TestEngineContextErrors(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	engine := NewEngine(7)
	var missing context.Context
	if _, err := engine.Run(missing, target, suite); !errors.Is(err, ErrCanceled) {
		t.Fatalf("nil context error = %v, want ErrCanceled", err)
	}
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := engine.Run(canceled, target, suite); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled context error = %v, want context.Canceled", err)
	}
	if _, err := engine.Run(&flipContext{}, target, suite); !errors.Is(err, context.Canceled) {
		t.Fatalf("loop guard error = %v, want context.Canceled", err)
	}
}
