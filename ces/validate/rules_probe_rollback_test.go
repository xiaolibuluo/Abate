/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"testing"

	"abate/ces/schema"
)

// assertViolationShape fails unless one violation carries the exact rule id,
// field path, severity and a non-empty diagnostic message.
func assertViolationShape(t *testing.T, vs []Violation, id, field string, sev Severity) {
	t.Helper()
	for _, v := range vs {
		if v.RuleID != id {
			continue
		}
		if v.FieldPath != field {
			t.Fatalf("%s fieldPath = %q, want %q", id, v.FieldPath, field)
		}
		if v.Severity != sev {
			t.Fatalf("%s severity = %q, want %q", id, v.Severity, sev)
		}
		if v.Message == "" {
			t.Fatalf("%s message must not be empty", id)
		}
		return
	}
	t.Fatalf("expected violation %s, got %+v", id, vs)
}

// ruleWay is one positive and one negative example of a single rule.
type ruleWay struct {
	name     string
	id       string
	positive func(*schema.Experiment)
	negative func(*schema.Experiment)
	field    string
}

// runWays exercises each case's positive and negative example in isolation.
func runWays(t *testing.T, cases []ruleWay) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pos := fresh(t)
			c.positive(pos)
			assertNoRule(t, ruleViolations(t, c.id, pos), c.id)
			neg := fresh(t)
			c.negative(neg)
			assertViolationShape(t, ruleViolations(t, c.id, neg),
				c.id, c.field, SeverityError)
		})
	}
}

// TestProbeRulesEachWay covers one positive and one negative example for
// every CES-VALID-PROBE rule, asserting the exact rule id, field path and
// severity of the negative result.
func TestProbeRulesEachWay(t *testing.T) {
	t.Parallel()
	runWays(t, []ruleWay{
		{
			name: "probe-001-unique-id",
			id:   RuleProbe001,
			positive: func(e *schema.Experiment) {
				e.Spec.Probes = append(e.Spec.Probes, schema.Probe{
					ID: "p2", Kind: schema.ProbeHTTP, Interval: schema.Second,
					Spec: schema.ProbeSpec{URL: "http://sim/ready", ExpectStatus: 200},
				})
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Probes = append(e.Spec.Probes, e.Spec.Probes[0])
			},
			field: "spec.probes[1].id",
		},
		{
			name:     "probe-002-positive-interval",
			id:       RuleProbe002,
			positive: func(e *schema.Experiment) { e.Spec.Probes[0].Interval = 2 * schema.Second },
			negative: func(e *schema.Experiment) { e.Spec.Probes[0].Interval = 0 },
			field:    "spec.probes[0].interval",
		},
		{
			name:     "probe-003-simulator-url",
			id:       RuleProbe003,
			positive: func(e *schema.Experiment) { e.Spec.Probes[0].Spec.URL = "https://localhost:8443/health" },
			negative: func(e *schema.Experiment) { e.Spec.Probes[0].Spec.URL = "https://api.example.com/health" },
			field:    "spec.probes[0].spec.url",
		},
		{
			name: "probe-004-safe-shell",
			id:   RuleProbe004,
			positive: func(e *schema.Experiment) {
				e.Spec.Probes[0].Kind = schema.ProbeShell
				e.Spec.Probes[0].Spec.Command = "echo ready"
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Probes[0].Kind = schema.ProbeShell
				e.Spec.Probes[0].Spec.Command = "echo ready; curl http://outside"
			},
			field: "spec.probes[0].spec.command",
		},
		{
			name:     "probe-005-http-status",
			id:       RuleProbe005,
			positive: func(e *schema.Experiment) { e.Spec.Probes[0].Spec.ExpectStatus = 204 },
			negative: func(e *schema.Experiment) { e.Spec.Probes[0].Spec.ExpectStatus = 99 },
			field:    "spec.probes[0].spec.expectStatus",
		},
	})
}

// TestProbe005KindBranches covers the remaining PROBE-005 switch arms.
func TestProbe005KindBranches(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name  string
		mut   func(*schema.Experiment)
		field string
	}{
		{"http-status-high", func(e *schema.Experiment) {
			e.Spec.Probes[0].Spec.ExpectStatus = 600
		}, "spec.probes[0].spec.expectStatus"},
		{"shell-exit-high", func(e *schema.Experiment) {
			e.Spec.Probes[0].Kind = schema.ProbeShell
			e.Spec.Probes[0].Spec.ExpectExitCode = 256
		}, "spec.probes[0].spec.expectExitCode"},
		{"shell-exit-negative", func(e *schema.Experiment) {
			e.Spec.Probes[0].Kind = schema.ProbeShell
			e.Spec.Probes[0].Spec.ExpectExitCode = -1
		}, "spec.probes[0].spec.expectExitCode"},
		{"metric-empty-expect", func(e *schema.Experiment) {
			e.Spec.Probes[0].Kind = schema.ProbeMetric
			e.Spec.Probes[0].Spec.Expect = ""
		}, "spec.probes[0].spec.expect"},
		{"prom-empty-expect", func(e *schema.Experiment) {
			e.Spec.Probes[0].Kind = schema.ProbeProm
			e.Spec.Probes[0].Spec.Expect = ""
		}, "spec.probes[0].spec.expect"},
		{"unknown-kind", func(e *schema.Experiment) {
			e.Spec.Probes[0].Kind = "quantum"
		}, "spec.probes[0].kind"},
	} {
		t.Run(c.name, func(t *testing.T) {
			exp := fresh(t)
			c.mut(exp)
			assertViolationShape(t, ruleViolations(t, RuleProbe005, exp),
				RuleProbe005, c.field, SeverityError)
		})
	}
}

// TestProbeKindSkipBranches covers the non-matching probe-kind arms of the
// simulator URL and shell subset rules.
func TestProbeKindSkipBranches(t *testing.T) {
	t.Parallel()
	shell := fresh(t)
	shell.Spec.Probes[0].Kind = schema.ProbeShell
	shell.Spec.Probes[0].Spec.Command = "echo ok"
	assertNoRule(t, ruleViolations(t, RuleProbe003, shell), RuleProbe003)
	assertNoRule(t, ruleViolations(t, RuleProbe004, shell), RuleProbe004)
	httpOnly := fresh(t)
	assertNoRule(t, ruleViolations(t, RuleProbe004, httpOnly), RuleProbe004)
}

// TestRollbackRulesEachWay covers one positive and one negative example for
// every CES-VALID-ROLLBACK rule.
func TestRollbackRulesEachWay(t *testing.T) {
	t.Parallel()
	runWays(t, []ruleWay{
		{
			name:     "rollback-001-strategy-enum",
			id:       RuleRollback001,
			positive: func(e *schema.Experiment) { e.Spec.Rollback.Strategy = schema.StrategyManual },
			negative: func(e *schema.Experiment) { e.Spec.Rollback.Strategy = "quantum" },
			field:    "spec.rollback.strategy",
		},
		{
			name: "rollback-002-timeout-positive",
			id:   RuleRollback002,
			positive: func(e *schema.Experiment) {
				e.Spec.Rollback.Strategy = schema.StrategyTimeout
				e.Spec.Rollback.Timeout = schema.Second
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Rollback.Strategy = schema.StrategyTimeout
				e.Spec.Rollback.Timeout = 0
			},
			field: "spec.rollback.timeout",
		},
		{
			name:     "rollback-003-steady-verification",
			id:       RuleRollback003,
			positive: func(e *schema.Experiment) { e.Spec.Rollback.VerifySteadyState = true },
			negative: func(e *schema.Experiment) { e.Spec.SteadyState = nil },
			field:    "spec.rollback.verifySteadyState",
		},
	})
}

// TestRollbackBranches covers the remaining strategy and verification arms.
func TestRollbackBranches(t *testing.T) {
	t.Parallel()
	for _, strategy := range []schema.Strategy{
		schema.StrategyAuto, schema.StrategyManual, schema.StrategyTimeout,
	} {
		exp := fresh(t)
		exp.Spec.Rollback.Strategy = strategy
		exp.Spec.Rollback.Timeout = schema.Second
		assertNoRule(t, ruleViolations(t, RuleRollback001, exp), RuleRollback001)
		assertNoRule(t, ruleViolations(t, RuleRollback002, exp), RuleRollback002)
	}
	disabled := fresh(t)
	disabled.Spec.Rollback.VerifySteadyState = false
	disabled.Spec.SteadyState = nil
	assertNoRule(t, ruleViolations(t, RuleRollback003, disabled), RuleRollback003)
}

// TestRollback003ViaAggregate proves the aggregate validator reports both
// the steady-state failure and the dependent rollback failure.
func TestRollback003ViaAggregate(t *testing.T) {
	t.Parallel()
	exp := fresh(t)
	exp.Spec.SteadyState = nil
	vs, err := Validate(exp)
	ok(t, err)
	assertRule(t, vs, RuleSteady001)
	assertRule(t, vs, RuleRollback003)
}
