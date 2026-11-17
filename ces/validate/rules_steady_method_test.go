/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"strings"
	"testing"

	"abate/ces/schema"
)

// TestSteadyMethodBaseline proves the clean document passes all nine rules.
func TestSteadyMethodBaseline(t *testing.T) {
	t.Parallel()
	ids := append(allRuleIDs()[:0:0], allRuleIDs()...)
	for _, id := range ids {
		if !strings.HasPrefix(id, "CES-VALID-STEADY") &&
			!strings.HasPrefix(id, "CES-VALID-METHOD") {
			continue
		}
		assertNoRule(t, ruleViolations(t, id, fresh(t)), id)
	}
}

// TestSteadyRulesEachWay covers one negative case per steady-state rule.
func TestSteadyRulesEachWay(t *testing.T) {
	t.Parallel()
	cases := []struct {
		id  string
		mut func(*schema.Experiment)
	}{
		{RuleSteady001, func(e *schema.Experiment) { e.Spec.SteadyState = nil }},
		{RuleSteady002, func(e *schema.Experiment) {
			e.Spec.SteadyState[0].ProbeRef = "ghost"
		}},
		{RuleSteady003, func(e *schema.Experiment) {
			e.Spec.SteadyState[0].HoldFor = 0
		}},
		{RuleSteady003, func(e *schema.Experiment) {
			e.Spec.SteadyState[0].HoldFor = schema.Millisecond
			e.Spec.Probes[0].Interval = schema.Second
		}},
		{RuleSteady004, func(e *schema.Experiment) {
			e.Spec.SteadyState[0].Expectation.Type = schema.ExpectRange
			e.Spec.SteadyState[0].Tolerance.Type = schema.ToleranceRelative
		}},
		{RuleSteady004, func(e *schema.Experiment) {
			e.Spec.SteadyState[0].Tolerance.Type = "quantum"
		}},
		{RuleSteady005, func(e *schema.Experiment) {
			e.Spec.SteadyState[0].Expectation.Type = schema.ExpectPredicate
			e.Spec.SteadyState[0].Expectation.Predicate = ">=abc"
		}},
	}
	for i, c := range cases {
		exp := fresh(t)
		c.mut(exp)
		assertRule(t, ruleViolations(t, c.id, exp), c.id)
		_ = i
	}
}

// TestMethodRulesEachWay covers one negative case per method rule.
func TestMethodRulesEachWay(t *testing.T) {
	t.Parallel()
	cases := []struct {
		id  string
		mut func(*schema.Experiment)
	}{
		{RuleMethod001, func(e *schema.Experiment) {
			e.Spec.Method.Expression = ""
		}},
		{RuleMethod001, func(e *schema.Experiment) {
			e.Spec.Method.Expression = "a @ b"
		}},
		{RuleMethod002, func(e *schema.Experiment) {
			e.Spec.Method.Expression = "operator://m-noninv"
			e.Spec.Method.Operators = []schema.OperatorDef{{
				ID: "m-noninv", Family: schema.FamilyMutation,
				TargetKind: schema.TargetSource, Targets: []string{"src"},
			}}
		}},
		{RuleMethod003, func(e *schema.Experiment) {
			e.Spec.Method.Expression = strings.Repeat("(", 70) + "ε" + strings.Repeat(")", 70)
		}},
		{RuleMethod004, func(e *schema.Experiment) {
			e.Spec.Method.Expression = "a ⊗ a"
		}},
	}
	for _, c := range cases {
		exp := fresh(t)
		c.mut(exp)
		assertRule(t, ruleViolations(t, c.id, exp), c.id)
	}
}

// TestMethod002CompensationPositive accepts a non-invertible operator that
// declares a compensation strategy.
func TestMethod002CompensationPositive(t *testing.T) {
	t.Parallel()
	exp := fresh(t)
	exp.Spec.Method.Expression = "operator://m-noninv"
	exp.Spec.Method.Operators = []schema.OperatorDef{{
		ID: "m-noninv", Family: schema.FamilyMutation,
		TargetKind: schema.TargetSource, Targets: []string{"src"},
		Compensation: "restore-from-snapshot",
	}}
	assertNoRule(t, ruleViolations(t, RuleMethod002, exp), RuleMethod002)
}
