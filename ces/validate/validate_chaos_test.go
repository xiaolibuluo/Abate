/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"slices"
	"sync"
	"testing"

	"abate/ces/schema"
)

// digestViolations returns the SHA-256 of the canonical JSON encoding.
func digestViolations(t *testing.T, vs []Violation) string {
	t.Helper()
	b, err := json.Marshal(vs)
	ok(t, err)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

// chaoticDocument returns a normalized document that trips several rules so
// aggregate ordering and severity mapping are exercised.
func chaoticDocument(t *testing.T) *schema.Experiment {
	t.Helper()
	exp := fresh(t)
	exp.Spec.Probes[0].Spec.ExpectStatus = 99
	exp.Spec.Probes[0].Interval = 0
	exp.Spec.Rollback.Strategy = "quantum"
	exp.Spec.Target.Kind = "quantum"
	exp.Spec.SteadyState[0].HoldFor = 0
	return exp
}

// TestChaosDeterminism100 proves 100 sequential validations of one document
// produce byte-identical JSON and a non-decreasing aggregate rule order.
func TestChaosDeterminism100(t *testing.T) {
	t.Parallel()
	exp := chaoticDocument(t)
	first := ""
	for i := 0; i < 100; i++ {
		vs, err := Validate(exp)
		ok(t, err)
		if len(vs) == 0 {
			t.Fatal("chaotic document produced no violations")
		}
		for j := 1; j < len(vs); j++ {
			if violationLess(vs[j], vs[j-1]) {
				t.Fatalf("aggregate order broke at %d on iteration %d", j, i)
			}
		}
		d := digestViolations(t, vs)
		if i == 0 {
			first = d
			continue
		}
		if d != first {
			t.Fatalf("validate digest drifted on iteration %d", i)
		}
	}
}

// TestChaosConcurrent100 validates 100 independent copies concurrently and
// requires every digest to match.
func TestChaosConcurrent100(t *testing.T) {
	t.Parallel()
	const n = 100
	docs := make([]*schema.Experiment, n)
	for i := range docs {
		docs[i] = chaoticDocument(t)
	}
	digests := make([]string, n)
	errs := make([]error, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			vs, err := Validate(docs[k])
			if err != nil {
				errs[k] = err
				return
			}
			b, err := json.Marshal(vs)
			if err != nil {
				errs[k] = err
				return
			}
			sum := sha256.Sum256(b)
			digests[k] = hex.EncodeToString(sum[:])
		}(i)
	}
	wg.Wait()
	for i := range errs {
		if errs[i] != nil {
			t.Fatalf("concurrent validate failed: %v", errs[i])
		}
		if digests[i] != digests[0] {
			t.Fatalf("concurrent digest drifted at %d", i)
		}
	}
}

// TestChaosConcurrentRegistry100 proves registry construction is reentrant.
func TestChaosConcurrentRegistry100(t *testing.T) {
	t.Parallel()
	const n = 100
	ids := make([][]string, n)
	errs := make([]error, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			got, err := RegisteredIDs()
			if err != nil {
				errs[k] = err
				return
			}
			ids[k] = got
		}(i)
	}
	wg.Wait()
	for k := range errs {
		if errs[k] != nil {
			t.Fatalf("concurrent registry build failed: %v", errs[k])
		}
		if !slices.Equal(ids[k], ids[0]) {
			t.Fatalf("registry id order drifted at %d", k)
		}
	}
}

// TestChaosRegisteredIDsCanonical proves the 25 builtin ids are unique,
// lexicographic and equal to the enumerated list.
func TestChaosRegisteredIDsCanonical(t *testing.T) {
	t.Parallel()
	ids, err := RegisteredIDs()
	ok(t, err)
	if len(ids) != 25 {
		t.Fatalf("registered ids = %d, want 25", len(ids))
	}
	if !slices.IsSorted(ids) {
		t.Fatalf("registered ids not lexicographic: %v", ids)
	}
	seen := make(map[string]bool, len(ids))
	for _, id := range ids {
		if seen[id] {
			t.Fatalf("duplicate rule id %s", id)
		}
		seen[id] = true
	}
	want := allRuleIDs()
	if !slices.IsSorted(want) || !slices.Equal(ids, want) {
		t.Fatalf("registered ids disagree with the enumeration: %v", ids)
	}
	reg, err := builtinRegistry()
	ok(t, err)
	if !slices.Equal(reg.IDs(), ids) {
		t.Fatal("registry ids disagree with RegisteredIDs")
	}
	vs := reg.run(chaoticDocument(t))
	if len(vs) == 0 {
		t.Fatal("registry aggregate produced no violations")
	}
	for i := 1; i < len(vs); i++ {
		if vs[i].RuleID < vs[i-1].RuleID {
			t.Fatalf("aggregate order not deterministic at %d", i)
		}
	}
}

// TestChaosUnnormalizedInputs proves every normalization invariant is
// enforced with the sentinel error before any rule runs.
func TestChaosUnnormalizedInputs(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name string
		mut  func(*schema.Experiment)
	}{
		{"empty-namespace", func(e *schema.Experiment) { e.Metadata.Namespace = "" }},
		{"empty-strategy", func(e *schema.Experiment) { e.Spec.Rollback.Strategy = "" }},
		{"zero-seed", func(e *schema.Experiment) { e.Spec.Seed = 0 }},
		{"no-schedule", func(e *schema.Experiment) { e.Spec.Schedule = schema.Schedule{} }},
		{"two-schedules", func(e *schema.Experiment) {
			e.Spec.Schedule.Cron = &schema.CronSpec{Expression: "* * * * *"}
		}},
	} {
		t.Run(c.name, func(t *testing.T) {
			exp := fresh(t)
			c.mut(exp)
			vs, err := Validate(exp)
			if !errors.Is(err, ErrUnnormalized) {
				t.Fatalf("error = %v, want ErrUnnormalized", err)
			}
			if vs != nil {
				t.Fatalf("unnormalized input produced violations: %+v", vs)
			}
		})
	}
	vs, err := Validate(nil)
	if !errors.Is(err, ErrUnnormalized) || vs != nil {
		t.Fatalf("nil input: vs=%+v err=%v, want ErrUnnormalized", vs, err)
	}
}

// TestDeterminismJSONSHA256 proves the same document serializes to the exact
// same SHA-256 digest across repeated validations.
func TestDeterminismJSONSHA256(t *testing.T) {
	t.Parallel()
	exp := fresh(t)
	first, err := Validate(exp)
	ok(t, err)
	want := digestViolations(t, first)
	for i := 0; i < 100; i++ {
		got, err := Validate(exp)
		ok(t, err)
		if d := digestViolations(t, got); d != want {
			t.Fatalf("JSON SHA-256 drifted on iteration %d", i)
		}
	}
	bad := chaoticDocument(t)
	if digestViolations(t, mustValidate(t, bad)) == want {
		t.Fatal("distinct documents produced the same digest")
	}
}

// mustValidate validates a normalized document and fails on an error.
func mustValidate(t *testing.T, exp *schema.Experiment) []Violation {
	t.Helper()
	vs, err := Validate(exp)
	ok(t, err)
	return vs
}
