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
	"sync"
	"testing"

	"abate/kernel/algebra"
	"abate/kernel/verify"
)

// TestNewRegistryStillSevenClasses proves the additive extension leaves the
// M2 public constructor byte-for-byte unchanged.
func TestNewRegistryStillSevenClasses(t *testing.T) {
	t.Parallel()
	base := NewRegistry()
	withEmpty := newRegistryWithNil(t)
	want := []string{"arith-replace", "boundary-value", "constant-perturb",
		"control-flow", "return-value", "negate-condition", "remove-call"}
	if !slices.Equal(base.Classes(), want) || !slices.Equal(withEmpty.Classes(), want) {
		t.Fatal("the seven operator classes changed")
	}
	if len(base.All()) != 7 || len(withEmpty.All()) != 7 {
		t.Fatal("a registry without derived operators must expose exactly 7 operators")
	}
	if !slices.Equal(idsOf(base.All()), idsOf(withEmpty.All())) {
		t.Fatal("NewRegistryWith(nil) changed the base operator set")
	}
	for _, token := range []string{"+", "<", "1", "break", "true", "==", "call"} {
		if !slices.Equal(idsOf(base.For(token)), idsOf(withEmpty.For(token))) {
			t.Fatalf("For(%q) changed without derived operators", token)
		}
	}
}

// TestDerivedRegistryRegistration locks discovery, validation and ordering.
func TestDerivedRegistryRegistration(t *testing.T) {
	t.Parallel()
	ops := []algebra.Operator{
		derivedOperator("b", "boundary-value", "<", "<="),
		derivedOperator("a", "negate-condition", "==", "!="),
		derivedOperator("c", "remove-call", "call", tokenDeleted),
	}
	reg, err := NewRegistryWith(ops)
	if err != nil {
		t.Fatalf("NewRegistryWith: %v", err)
	}
	if len(reg.All()) != 10 {
		t.Fatalf("All() = %d operators, want 7 + 3", len(reg.All()))
	}
	if !slices.IsSortedFunc(reg.All(), byOperatorID) {
		t.Fatal("All() is not sorted by identifier")
	}
	for _, op := range ops {
		if !reg.Valid(op) {
			t.Fatalf("registry rejected its own derived operator %s", op.ID())
		}
		token := op.Params()["from"]
		found := false
		for _, candidate := range reg.For(token) {
			if candidate.ID() == op.ID() {
				found = true
			}
		}
		if !found {
			t.Fatalf("For(%q) did not discover %s", token, op.ID())
		}
	}
	if reg.For("") != nil || len(reg.For("identifier")) != 0 {
		t.Fatal("derived operators leaked onto an unmutatable token")
	}
	if len(reg.For("<deleted>")) != 0 {
		t.Fatal("the deleted token must stay outside the derived index")
	}
	first := reg.For("<")
	first[0] = nil
	if reg.For("<")[0] == nil {
		t.Fatal("For() leaked its internal frozen slice")
	}
}

// TestDerivedRegistryRejections exercises every admission failure branch.
func TestDerivedRegistryRejections(t *testing.T) {
	t.Parallel()
	good := derivedOperator("ok", "boundary-value", "<", "<=")
	badClass := derivedOperator("bad", "not-a-class", "<", "<=")
	badFamily := algebra.NewBasic(
		algebra.FormatOperatorID(algebra.FamilyChaos, "derived-x", "1"),
		algebra.FamilyChaos,
		algebra.Params{"class": "boundary-value", "from": "<", "to": "<=", "derived-from": "x"},
		[]algebra.TargetKey{"<"}, true)
	emptyID := algebra.NewBasic("", algebra.FamilyMutation,
		good.Params(), []algebra.TargetKey{"<"}, true)
	for _, c := range []struct {
		name string
		op   algebra.Operator
	}{
		{"nil", nil},
		{"empty-id", emptyID},
		{"wrong-family", badFamily},
		{"unknown-class", badClass},
		{"missing-from", algebra.NewBasic("m-derived-x-v1", algebra.FamilyMutation,
			algebra.Params{"class": "boundary-value", "to": "<=", "derived-from": "x"}, nil, true)},
		{"missing-to", algebra.NewBasic("m-derived-x-v1", algebra.FamilyMutation,
			algebra.Params{"class": "boundary-value", "from": "<", "derived-from": "x"}, nil, true)},
		{"missing-derived-from", algebra.NewBasic("m-derived-x-v1", algebra.FamilyMutation,
			algebra.Params{"class": "boundary-value", "from": "<", "to": "<="}, nil, true)},
	} {
		t.Run(c.name, func(t *testing.T) {
			if _, err := NewRegistryWith([]algebra.Operator{c.op}); !errors.Is(err, ErrInvalidDerived) {
				t.Fatalf("error = %v, want ErrInvalidDerived", err)
			}
		})
	}
	if _, err := NewRegistryWith([]algebra.Operator{good, good}); !errors.Is(err, ErrInvalidDerived) {
		t.Fatalf("duplicate error = %v, want ErrInvalidDerived", err)
	}
}

// TestDerivedRegistryConstructionIsOrderIndependent proves the derived table is
// frozen by identifier, not by caller slice order.
func TestDerivedRegistryConstructionIsOrderIndependent(t *testing.T) {
	t.Parallel()
	ops := []algebra.Operator{
		derivedOperator("a", "boundary-value", "<", "<="),
		derivedOperator("b", "negate-condition", "==", "!="),
		derivedOperator("c", "return-value", "true", "false"),
		derivedOperator("d", "remove-call", "call", tokenDeleted),
	}
	forward, err := NewRegistryWith(ops)
	if err != nil {
		t.Fatalf("forward: %v", err)
	}
	reversed := append([]algebra.Operator(nil), ops...)
	slices.Reverse(reversed)
	backward, err := NewRegistryWith(reversed)
	if err != nil {
		t.Fatalf("backward: %v", err)
	}
	if !slices.Equal(idsOf(forward.All()), idsOf(backward.All())) {
		t.Fatal("derived operator ordering depends on caller slice order")
	}
	for _, token := range []string{"<", "==", "true", "call"} {
		if !slices.Equal(idsOf(forward.For(token)), idsOf(backward.For(token))) {
			t.Fatalf("For(%q) depends on caller slice order", token)
		}
	}
}

// TestDerivedEngineRun proves the real Engine.Run emits derived points with
// allocated verdicts and a byte-stable report.
func TestDerivedEngineRun(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	ops := derivedBatch(25, "boundary-value", "<", "<=")
	reg, err := NewRegistryWith(ops)
	if err != nil {
		t.Fatalf("registry: %v", err)
	}
	engine, err := NewEngineWithRegistry(42, reg)
	if err != nil {
		t.Fatalf("engine: %v", err)
	}
	if engine.Seed() != 42 || engine.Registry() != reg {
		t.Fatal("engine did not keep the supplied seed and registry")
	}
	rep, err := engine.Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	seen := make(map[string]bool, 32)
	for _, point := range rep.Points {
		if point.Operator == nil {
			continue
		}
		id := string(point.Operator.ID())
		if strings.HasPrefix(id, "m-derived-") {
			seen[id] = true
		}
	}
	if len(seen) < 20 {
		t.Fatalf("Report.Points carries only %d distinct derived operators", len(seen))
	}
	for i, verdict := range rep.Verdicts {
		switch verdict {
		case verify.MutantKilled, verify.MutantSurvived,
			verify.MutantEquivalent, verify.MutantTimeout:
		default:
			t.Fatalf("verdict %d = %q is not allocated", i, verdict)
		}
	}
	if ReportFingerprint(rep) != rep.Fingerprint {
		t.Fatal("report fingerprint disagrees with its canonical body")
	}
	again, err := engine.Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("second Run: %v", err)
	}
	if string(rep.Bytes()) != string(again.Bytes()) {
		t.Fatal("same-seed runs produced different report bytes")
	}
}

// TestNewEngineWithRegistryNil rejects a missing registry with a sentinel.
func TestNewEngineWithRegistryNil(t *testing.T) {
	t.Parallel()
	if _, err := NewEngineWithRegistry(1, nil); !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("error = %v, want ErrNilRegistry", err)
	}
}

// TestDerivedRegistryConcurrentReads proves the frozen table is race free.
func TestDerivedRegistryConcurrentReads(t *testing.T) {
	t.Parallel()
	reg, err := NewRegistryWith(derivedBatch(24, "boundary-value", "<", "<="))
	if err != nil {
		t.Fatalf("registry: %v", err)
	}
	engine, err := NewEngineWithRegistry(7, reg)
	if err != nil {
		t.Fatalf("engine: %v", err)
	}
	target, suite, _ := newFixture()
	const workers = 16
	results := make(chan string, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			reg.All()
			reg.For("<")
			reg.Valid(reg.All()[0])
			rep, runErr := engine.Run(context.Background(), target, suite)
			if runErr != nil {
				results <- "error: " + runErr.Error()
				return
			}
			results <- shaHex(rep.Bytes())
		}()
	}
	wg.Wait()
	close(results)
	unique := make(map[string]bool, workers)
	for digest := range results {
		unique[digest] = true
	}
	if len(unique) != 1 {
		t.Fatalf("concurrent runs produced %d distinct report digests", len(unique))
	}
}
