/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"encoding/json"
	"math"
	"testing"

	"abate/ces/schema"
)

// TestNormalizeDefaults fills namespace, rollback strategy and a derived seed.
func TestNormalizeDefaults(t *testing.T) {
	t.Parallel()
	exp := validExperiment(t)
	exp.Metadata.Namespace = ""
	exp.Spec.Rollback.Strategy = ""
	exp.Spec.Seed = 0
	got, err := Normalize(exp)
	requireOK(t, err)
	if got.Metadata.Namespace != "default" {
		t.Fatalf("namespace = %q, want default", got.Metadata.Namespace)
	}
	if got.Spec.Rollback.Strategy != schema.StrategyAuto {
		t.Fatalf("strategy = %q, want auto", got.Spec.Rollback.Strategy)
	}
	if got.Spec.Seed == 0 {
		t.Fatal("seed was not derived")
	}
	if exp.Spec.Seed != 0 || exp.Metadata.Namespace != "" {
		t.Fatal("Normalize mutated its input")
	}
}

// TestNormalizePreservesExplicitValues keeps non-default inputs untouched.
func TestNormalizePreservesExplicitValues(t *testing.T) {
	t.Parallel()
	exp := validExperiment(t)
	exp.Metadata.Namespace = "prod"
	exp.Spec.Rollback.Strategy = schema.StrategyManual
	got, err := Normalize(exp)
	requireOK(t, err)
	if got.Metadata.Namespace != "prod" || got.Spec.Rollback.Strategy != schema.StrategyManual {
		t.Fatalf("explicit values overwritten: %+v", got)
	}
	if got.Spec.Seed != 12345 {
		t.Fatalf("explicit seed overwritten: %d", got.Spec.Seed)
	}
}

// TestNormalizeNil rejects a nil experiment.
func TestNormalizeNil(t *testing.T) {
	t.Parallel()
	_, err := Normalize(nil)
	requireRule(t, err, RuleEmptyDocument)
}

// TestNormalizeProbeValidation covers empty, duplicate and bad-interval probes.
func TestNormalizeProbeValidation(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		mut  func(*schema.Experiment)
	}{
		{"empty-id", func(e *schema.Experiment) { e.Spec.Probes[0].ID = "" }},
		{"duplicate-id", func(e *schema.Experiment) {
			p := e.Spec.Probes[0]
			e.Spec.Probes = append(e.Spec.Probes, p)
		}},
		{"zero-interval", func(e *schema.Experiment) { e.Spec.Probes[0].Interval = 0 }},
		{"negative-interval", func(e *schema.Experiment) {
			e.Spec.Probes[0].Interval = -schema.Second
		}},
	} {
		exp := validExperiment(t)
		tc.mut(exp)
		_, err := Normalize(exp)
		requireRule(t, err, RuleMalformedDocument)
	}
}

// TestNormalizeDanglingProbeRef rejects steady-state referencing a missing probe.
func TestNormalizeDanglingProbeRef(t *testing.T) {
	t.Parallel()
	exp := validExperiment(t)
	exp.Spec.SteadyState[0].ProbeRef = "ghost"
	_, err := Normalize(exp)
	requireRule(t, err, RuleDanglingProbeRef)
}

// TestNormalizeScheduleExclusive enforces exactly one scheduling mode.
func TestNormalizeScheduleExclusive(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		set  func(*schema.Schedule)
	}{
		{"none", func(s *schema.Schedule) { *s = schema.Schedule{} }},
		{"two", func(s *schema.Schedule) { s.Cron = &schema.CronSpec{Expression: "* * * * *"} }},
		{"three", func(s *schema.Schedule) {
			s.Cron = &schema.CronSpec{Expression: "* * * * *"}
			s.Continuous = &schema.ContinuousSpec{Duration: schema.Second}
		}},
	} {
		exp := validExperiment(t)
		tc.set(&exp.Spec.Schedule)
		_, err := Normalize(exp)
		requireRule(t, err, RuleScheduleExclusive)
	}
}

// TestScheduleModes counts each combination of non-nil modes.
func TestScheduleModes(t *testing.T) {
	t.Parallel()
	cases := []struct {
		s    schema.Schedule
		want int
	}{
		{schema.Schedule{}, 0},
		{schema.Schedule{Once: &schema.OnceSpec{}}, 1},
		{schema.Schedule{Cron: &schema.CronSpec{}}, 1},
		{schema.Schedule{Continuous: &schema.ContinuousSpec{}}, 1},
		{schema.Schedule{Once: &schema.OnceSpec{}, Cron: &schema.CronSpec{}}, 2},
		{schema.Schedule{
			Once: &schema.OnceSpec{}, Cron: &schema.CronSpec{},
			Continuous: &schema.ContinuousSpec{},
		}, 3},
	}
	for _, tc := range cases {
		if got := scheduleModes(tc.s); got != tc.want {
			t.Fatalf("scheduleModes = %d, want %d", got, tc.want)
		}
	}
}

// TestDeriveSeedStable proves hashing is deterministic and input-sensitive.
func TestDeriveSeedStable(t *testing.T) {
	t.Parallel()
	a := validExperiment(t)
	b := validExperiment(t)
	b.Metadata.Name = "other"
	sa, err := deriveSeed(a)
	requireOK(t, err)
	sa2, err := deriveSeed(a)
	requireOK(t, err)
	sb, err := deriveSeed(b)
	requireOK(t, err)
	if sa != sa2 {
		t.Fatal("deriveSeed is not deterministic")
	}
	if sa == sb {
		t.Fatal("deriveSeed ignores document content")
	}
}

// TestNormalizeSeedMarshalFailure covers the NaN -> json.Marshal error branch.
func TestNormalizeSeedMarshalFailure(t *testing.T) {
	t.Parallel()
	exp := validExperiment(t)
	exp.Spec.Seed = 0
	exp.Spec.SteadyState[0].Tolerance.Value = math.NaN()
	if _, err := json.Marshal(exp); err == nil {
		t.Fatal("precondition: NaN must break json.Marshal")
	}
	_, err := Normalize(exp)
	requireRule(t, err, RuleMalformedDocument)
}
