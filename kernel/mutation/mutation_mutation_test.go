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
	"testing"

	"abate/kernel/algebra"
	"abate/kernel/verify"
)

var allTokens = []string{
	"+", "-", "*", "/", "&&", "||", "<", "<=", ">", ">=",
	"1", "0", "-1", "42", "break", "continue", "true", "false",
	"nil", "nonnil", "==", "!=", "call", "<deleted>", "identifier", "",
}

func TestEveryOperatorMutantIsDetected(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	target := &syntheticTarget{
		nodes:      []string{"n"},
		tokens:     map[string]string{"n": "+"},
		killFrom:   map[string]string{"n": "t2"},
		equivalent: map[string]bool{},
		timeout:    map[string]bool{},
	}
	tests := []string{"t1", "t2", "t3"}
	checked := 0
	for _, token := range allTokens {
		for _, op := range reg.For(token) {
			mutated := op.Params()["to"]
			point := NewMutationPoint("n", token, mutated, op, 1, 1)
			verdict, killer, err := Judge(reg, target, tests, point)
			if err != nil {
				t.Fatalf("token %q operator %s: %v", token, op.ID(), err)
			}
			checked++
			if mutated == token {
				if verdict != verify.MutantEquivalent || killer != "" {
					t.Fatalf("identity mutant %q -> %q got %s/%s", token, mutated, verdict, killer)
				}
				continue
			}
			if verdict != verify.MutantKilled || killer != "t2" {
				t.Fatalf("mutant %q -> %q got %s/%s, want KILLED/t2", token, mutated, verdict, killer)
			}
		}
	}
	if checked < 30 {
		t.Fatalf("only %d operator variants were checked", checked)
	}
	t.Logf("checked %d operator variants across all seven classes", checked)
}

func TestOperatorFingerprintsAreUnique(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	seen := make(map[[32]byte]string)
	for _, token := range allTokens {
		for _, op := range reg.For(token) {
			fp := op.Fingerprint()
			if prior, ok := seen[fp]; ok {
				t.Fatalf("fingerprint collision between %s and %s", prior, op.ID())
			}
			seen[fp] = string(op.ID())
		}
	}
	if len(seen) < 30 {
		t.Fatalf("only %d distinct operator fingerprints", len(seen))
	}
}

func TestBoundaryAndIllegalInputs(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	if reg.For("") != nil {
		t.Fatal("empty token produced operators")
	}
	if len(reg.For("<deleted>")) != 0 {
		t.Fatal("deleted sentinel produced operators")
	}
	long := "x"
	for i := 0; i < 64; i++ {
		long += "x"
	}
	if len(reg.For(long)) != 0 {
		t.Fatal("oversized token produced operators")
	}
	if reg.Valid(nil) {
		t.Fatal("Valid(nil) should be false, not panic")
	}
	wrongFamily := algebra.NewBasic(algebra.FormatOperatorID(algebra.FamilyChaos, "x", "1"),
		algebra.FamilyChaos, algebra.Params{"class": "arith-replace"}, nil, true)
	if reg.Valid(wrongFamily) {
		t.Fatal("chaos-family operator accepted by mutation registry")
	}
	missingClass := algebra.NewBasic(algebra.FormatOperatorID(algebra.FamilyMutation, "y", "1"),
		algebra.FamilyMutation, algebra.Params{}, nil, true)
	if reg.Valid(missingClass) {
		t.Fatal("operator without a class accepted")
	}
}

func TestInvalidTargetAndSuiteThroughEngine(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	engine := NewEngine(3)
	ctx := context.Background()
	cases := []struct {
		name   string
		target Target
		suite  Suite
		want   error
	}{
		{"nil-target", nil, suite, ErrNilTarget},
		{"nil-suite", target, nil, ErrNilSuite},
		{"no-nodes", &syntheticTarget{nodes: nil, tokens: map[string]string{}}, suite, ErrNoNodes},
		{"no-tests", target, &syntheticSuite{tests: nil, cover: map[string]uint64{}}, ErrNoTests},
		{"no-mutants", &syntheticTarget{
			nodes: []string{"x"}, tokens: map[string]string{"x": "identifier"},
			killFrom: map[string]string{}, equivalent: map[string]bool{}, timeout: map[string]bool{},
		}, suite, ErrNoMutants},
	}
	for _, tc := range cases {
		if _, err := engine.Run(ctx, tc.target, tc.suite); !errors.Is(err, tc.want) {
			t.Fatalf("%s: error = %v, want %v", tc.name, err, tc.want)
		}
	}
}

func TestIrreversibleOperatorsUseSentinel(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	for _, class := range []string{"control-flow", "remove-call"} {
		var found algebra.Operator
		for _, op := range reg.All() {
			if op.Params()["class"] == class {
				found = op
				break
			}
		}
		if found == nil {
			t.Fatalf("class %s missing from registry", class)
		}
		if _, err := algebra.Inverse(found); !errors.Is(err, algebra.ErrNotInvertible) {
			t.Fatalf("algebra.Inverse(%s) = %v, want ErrNotInvertible", class, err)
		}
		if _, err := found.Inverse(); !errors.Is(err, algebra.ErrNotInvertible) {
			t.Fatalf("Basic.Inverse(%s) = %v, want ErrNotInvertible", class, err)
		}
	}
}

func TestTargetPortIsPureAndDefensive(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	first := target.Fingerprint()
	if first != target.Fingerprint() {
		t.Fatal("target fingerprint drifted")
	}
	nodes := target.Nodes()
	nodes[0] = "tampered"
	if target.Nodes()[0] == "tampered" {
		t.Fatal("Nodes() leaked its backing array")
	}
	tests := suite.Tests()
	tests[0] = "tampered"
	if suite.Tests()[0] == "tampered" {
		t.Fatal("Tests() leaked its backing array")
	}
	point := NewMutationPoint("a-arith", "+", "-", NewRegistry().All()[0], 1, 1)
	for i := 0; i < 8; i++ {
		if target.Evaluate(point, "t2") != verify.OutcomeViolated {
			t.Fatal("Evaluate is not a pure function")
		}
	}
}
