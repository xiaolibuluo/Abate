/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"slices"
	"testing"

	"abate/ces/schema"
)

// chaosGen is a deterministic SplitMix64 generator used only by chaos tests.
type chaosGen struct{ state uint64 }

// next advances the generator and returns the next unsigned word.
func (g *chaosGen) next() uint64 {
	g.state += 0x9E3779B97F4A7C15
	z := g.state
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

// violationLess implements the documented total order for assertions.
func violationLess(a, b Violation) bool {
	if a.RuleID != b.RuleID {
		return a.RuleID < b.RuleID
	}
	if a.FieldPath != b.FieldPath {
		return a.FieldPath < b.FieldPath
	}
	if a.Message != b.Message {
		return a.Message < b.Message
	}
	return string(a.Severity) < string(b.Severity)
}

// TestChaosRandomPerturbationsNeverPanics fuzzes document fields with a
// deterministic local generator; any panic fails the test.
func TestChaosRandomPerturbationsNeverPanics(t *testing.T) {
	t.Parallel()
	g := chaosGen{state: 0xD1B54A32D192ED03}
	base := fresh(t)
	values := []string{"", "ε", "@", "operator://x", "⊗", "http://sim/health",
		"../../etc", "a ⊗ b", "999", "-1", "Virtual/Tokyo", "Asia/Shanghai"}
	intervals := []schema.VirtualDuration{-1, 0, schema.Second}
	statuses := []int{0, 99, 200, 600}
	for i := 0; i < 2000; i++ {
		exp := *base
		v := values[g.next()%uint64(len(values))]
		switch g.next() % 8 {
		case 0:
			exp.Spec.Method.Expression = v
		case 1:
			exp.Spec.Target.Kind = v
		case 2:
			exp.Spec.Rollback.Strategy = schema.Strategy(v)
		case 3:
			exp.Spec.Probes = []schema.Probe{{
				ID: v, Kind: schema.ProbeKind(v),
				Interval: intervals[g.next()%uint64(len(intervals))],
				Spec: schema.ProbeSpec{
					URL: v, Command: v,
					ExpectStatus: statuses[g.next()%uint64(len(statuses))],
				},
			}}
		case 4:
			exp.Spec.Schedule = schema.Schedule{
				Cron: &schema.CronSpec{Expression: v, Timezone: v},
			}
		case 5:
			exp.Spec.SteadyState = []schema.SteadyState{{
				ProbeRef:    v,
				Expectation: schema.Expectation{Type: v, Predicate: v},
				Tolerance:   schema.Tolerance{Type: v},
				HoldFor:     intervals[g.next()%uint64(len(intervals))],
			}}
		case 6:
			exp.Spec.Target.Selector = map[string]string{v: v}
		case 7:
			exp.Metadata.Namespace = v
			exp.Spec.Seed = g.next()
		}
		_, _ = Validate(&exp)
	}
}

// TestChaosSortViolationsTotalOrder proves the comparator is a stable total
// order independent of the input permutation.
func TestChaosSortViolationsTotalOrder(t *testing.T) {
	t.Parallel()
	base := []Violation{
		{RuleID: "B", FieldPath: "x", Message: "m", Severity: SeverityError},
		{RuleID: "A", FieldPath: "y", Message: "n", Severity: SeverityWarning},
		{RuleID: "A", FieldPath: "x", Message: "z", Severity: SeverityError},
		{RuleID: "A", FieldPath: "x", Message: "a", Severity: SeverityWarning},
		{RuleID: "A", FieldPath: "x", Message: "a", Severity: SeverityError},
		{RuleID: "A", FieldPath: "x", Message: "a", Severity: SeverityError},
	}
	want := append([]Violation(nil), base...)
	sortViolations(want)
	for i := 1; i < len(want); i++ {
		if violationLess(want[i], want[i-1]) {
			t.Fatalf("sort is not ascending at %d: %+v", i, want)
		}
	}
	for shift := 0; shift < len(base); shift++ {
		perm := make([]Violation, 0, len(base))
		for j := 0; j < len(base); j++ {
			perm = append(perm, base[(j+shift)%len(base)])
		}
		sortViolations(perm)
		if !slices.Equal(perm, want) {
			t.Fatalf("sort result drifted for rotation %d", shift)
		}
	}
}
