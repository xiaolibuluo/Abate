/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"fmt"

	"abate/ces/schema"
)

func validateNegative() []scenario {
	var out []scenario
	add := func(spec string, mutate func(*schema.Experiment)) {
		e := clone(baseExp())
		mutate(&e)
		out = append(out, scenario{spec: spec, e: e})
	}
	// STEADY-001: no steady state.
	add("03-ces-spec.md §4 STEADY-001", func(e *schema.Experiment) { e.Spec.SteadyState = nil })
	// STEADY-003: holdFor not positive.
	add("03-ces-spec.md §4 STEADY-003", func(e *schema.Experiment) { e.Spec.SteadyState[0].HoldFor = 0 })
	// STEADY-003: holdFor shorter than probe interval.
	add("03-ces-spec.md §4 STEADY-003", func(e *schema.Experiment) { e.Spec.SteadyState[0].HoldFor = base / 2 })
	// STEADY-004: unknown expectation type.
	add("03-ces-spec.md §4 STEADY-004", func(e *schema.Experiment) { e.Spec.SteadyState[0].Expectation.Type = "bogus" })
	// STEADY-004: predicate with numeric tolerance (incompatible).
	add("03-ces-spec.md §4 STEADY-004", func(e *schema.Experiment) {
		e.Spec.SteadyState[0].Expectation.Type = schema.ExpectPredicate
		e.Spec.SteadyState[0].Expectation.Predicate = ">5"
		e.Spec.SteadyState[0].Tolerance.Type = schema.ToleranceAbsolute
		e.Spec.SteadyState[0].Tolerance.Value = 1
	})
	// STEADY-005: malformed predicate.
	add("03-ces-spec.md §4 STEADY-005", func(e *schema.Experiment) {
		e.Spec.SteadyState[0].Expectation.Type = schema.ExpectPredicate
		e.Spec.SteadyState[0].Expectation.Predicate = "not-a-predicate"
	})
	// METHOD-001: empty expression.
	add("03-ces-spec.md §4 METHOD-001", func(e *schema.Experiment) { e.Spec.Method.Expression = "" })
	// METHOD-002: non-invertible operator without compensation.
	add("03-ces-spec.md §4 METHOD-002", func(e *schema.Experiment) {
		e.Spec.Method.Expression = "operator://x"
		e.Spec.Method.Operators = []schema.OperatorDef{{
			ID: "x", Family: schema.FamilyMutation, Invertible: false, TargetKind: schema.TargetSource,
		}}
	})
	// METHOD-003: expression depth exceeds 64.
	add("03-ces-spec.md §4 METHOD-003", func(e *schema.Experiment) {
		expr := "ε"
		for i := 0; i < 100; i++ {
			expr = expr + " ∘ ε"
		}
		e.Spec.Method.Expression = expr
	})
	// METHOD-004: concurrent product sharing a target key.
	add("03-ces-spec.md §4 METHOD-004", func(e *schema.Experiment) {
		e.Spec.Method.Expression = "operator://a ⊗ operator://b"
		e.Spec.Method.Operators = []schema.OperatorDef{
			{ID: "a", Family: schema.FamilyChaos, Invertible: true, TargetKind: schema.TargetSimulator, Targets: []string{"node"}},
			{ID: "b", Family: schema.FamilyChaos, Invertible: true, TargetKind: schema.TargetSimulator, Targets: []string{"node"}},
		}
	})
	// PROBE-003: http probe with external URL.
	add("03-ces-spec.md §4 PROBE-003", func(e *schema.Experiment) {
		e.Spec.Probes = []schema.Probe{{
			ID: "p1", Kind: schema.ProbeHTTP, Spec: schema.ProbeSpec{URL: "https://example.com/health", Method: "GET", ExpectStatus: 200}, Interval: base,
		}}
	})
	// PROBE-004: shell probe with unsafe command.
	add("03-ces-spec.md §4 PROBE-004", func(e *schema.Experiment) {
		e.Spec.Probes = []schema.Probe{{
			ID: "p1", Kind: schema.ProbeShell, Spec: schema.ProbeSpec{Command: "rm -rf /", ExpectExitCode: 0}, Interval: base,
		}}
	})
	// PROBE-005: http expectStatus out of range.
	add("03-ces-spec.md §4 PROBE-005", func(e *schema.Experiment) {
		e.Spec.Probes = []schema.Probe{{
			ID: "p1", Kind: schema.ProbeHTTP, Spec: schema.ProbeSpec{URL: "sim://health", Method: "GET", ExpectStatus: 99}, Interval: base,
		}}
	})
	// PROBE-005: metric probe with empty expect.
	add("03-ces-spec.md §4 PROBE-005", func(e *schema.Experiment) {
		e.Spec.Probes = []schema.Probe{{
			ID: "p1", Kind: schema.ProbeMetric, Spec: schema.ProbeSpec{Query: "up"}, Interval: base,
		}}
	})
	// ROLLBACK-001: unknown strategy.
	add("03-ces-spec.md §4 ROLLBACK-001", func(e *schema.Experiment) { e.Spec.Rollback.Strategy = "bogus" })
	// ROLLBACK-002: timeout strategy without positive timeout.
	add("03-ces-spec.md §4 ROLLBACK-002", func(e *schema.Experiment) { e.Spec.Rollback.Strategy = schema.StrategyTimeout })
	// ROLLBACK-003: verifySteadyState with broken steady state.
	add("03-ces-spec.md §4 ROLLBACK-003", func(e *schema.Experiment) {
		e.Spec.Rollback.VerifySteadyState = true
		e.Spec.SteadyState[0].HoldFor = 0
	})
	// SCHEDULE-002: malformed cron expression.
	add("03-ces-spec.md §4 SCHEDULE-002", func(e *schema.Experiment) {
		e.Spec.Schedule = schema.Schedule{Cron: &schema.CronSpec{Expression: "* * *", Timezone: "UTC"}}
	})
	// SCHEDULE-003: continuous duration not positive.
	add("03-ces-spec.md §4 SCHEDULE-003", func(e *schema.Experiment) {
		e.Spec.Schedule = schema.Schedule{Continuous: &schema.ContinuousSpec{Duration: 0}}
	})
	// SCHEDULE-004: once.at not future.
	add("03-ces-spec.md §4 SCHEDULE-004", func(e *schema.Experiment) {
		e.Spec.Schedule = schema.Schedule{Once: &schema.OnceSpec{At: 0}}
	})
	// TARGET-001: external target rejected in v1.
	add("03-ces-spec.md §4 TARGET-001", func(e *schema.Experiment) { e.Spec.Target.Kind = schema.TargetExternal })
	// TARGET-002: mutation operator must target source.
	add("03-ces-spec.md §4 TARGET-002", func(e *schema.Experiment) {
		e.Spec.Method.Expression = "operator://x"
		e.Spec.Method.Operators = []schema.OperatorDef{{
			ID: "x", Family: schema.FamilyMutation, Invertible: true, TargetKind: schema.TargetSimulator,
		}}
	})
	// TARGET-003: empty selector.
	add("03-ces-spec.md §4 TARGET-003", func(e *schema.Experiment) { e.Spec.Target.Selector = map[string]string{} })
	// TARGET-004: mixed M⊗C needs dual selectors.
	add("03-ces-spec.md §4 TARGET-004", func(e *schema.Experiment) {
		e.Spec.Method.Expression = "m-op ⊗ c-op"
		e.Spec.Method.Operators = []schema.OperatorDef{
			{ID: "m-op", Family: schema.FamilyMutation, Invertible: true, TargetKind: schema.TargetSource},
			{ID: "c-op", Family: schema.FamilyChaos, Invertible: true, TargetKind: schema.TargetSimulator},
		}
	})
	// Variety: extra ROLLBACK-001 negatives to pad the validate group.
	for i := 0; i < 14; i++ {
		e := clone(baseExp())
		e.Spec.Rollback.Strategy = "unknown-strategy"
		e.Metadata.Name = fmt.Sprintf("tck-validate-extra-%d", i)
		out = append(out, scenario{spec: "03-ces-spec.md §4 ROLLBACK-001", e: e})
	}
	return out
}

func validatePositive() []scenario {
	var out []scenario
	for i := 0; i < 8; i++ {
		e := clone(baseExp())
		e.Metadata.Name = fmt.Sprintf("tck-validate-ok-%d", i)
		out = append(out, scenario{spec: "03-ces-spec.md §4 (valid document passes)", e: e, cat: "validate-ok"})
	}
	return out
}

func roundtrip() []scenario {
	var out []scenario
	kinds := []schema.ProbeKind{schema.ProbeMetric, schema.ProbeHTTP, schema.ProbeShell, schema.ProbeProm}
	for i := 0; i < 50; i++ {
		e := clone(baseExp())
		e.Metadata.Name = fmt.Sprintf("tck-roundtrip-%d", i)
		e.Spec.SteadyState[0].Expectation.Value = fmt.Sprintf("v%d", i%7)
		pk := kinds[i%len(kinds)]
		switch pk {
		case schema.ProbeMetric:
			e.Spec.Probes = []schema.Probe{{ID: "p1", Kind: pk, Spec: schema.ProbeSpec{Query: "up", Expect: "1"}, Interval: base}}
		case schema.ProbeHTTP:
			e.Spec.Probes = []schema.Probe{{ID: "p1", Kind: pk, Spec: schema.ProbeSpec{URL: "sim://health", Method: "GET", ExpectStatus: 200}, Interval: base}}
		case schema.ProbeShell:
			e.Spec.Probes = []schema.Probe{{ID: "p1", Kind: pk, Spec: schema.ProbeSpec{Command: "echo ok", ExpectExitCode: 0}, Interval: base}}
		case schema.ProbeProm:
			e.Spec.Probes = []schema.Probe{{ID: "p1", Kind: pk, Spec: schema.ProbeSpec{PromQL: "up", Range: base, Expect: "1"}, Interval: base}}
		}
		switch i % 3 {
		case 0:
			e.Spec.Schedule = schema.Schedule{Continuous: &schema.ContinuousSpec{Duration: base}}
		case 1:
			e.Spec.Schedule = schema.Schedule{Cron: &schema.CronSpec{Expression: "*/5 * * * *", Timezone: "UTC"}}
		case 2:
			e.Spec.Schedule = schema.Schedule{Once: &schema.OnceSpec{At: 5 * base}}
		}
		out = append(out, scenario{spec: "03-ces-spec.md §3 (YAML/JSON round-trip lossless)", e: e, cat: "roundtrip"})
	}
	return out
}
