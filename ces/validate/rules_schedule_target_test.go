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

// applyMixedProduct turns the method into a mixed mutation-chaos concurrent
// product so the dual-target rule applies.
func applyMixedProduct(e *schema.Experiment) {
	e.Spec.Method.Expression = "m-replace ⊗ operator://c-delay"
	e.Spec.Method.Operators = []schema.OperatorDef{{
		ID: "c-delay", Family: schema.FamilyChaos,
		TargetKind: schema.TargetSimulator, Targets: []string{"sim"},
	}}
}

// TestScheduleRulesEachWay covers one positive and one negative example for
// every CES-VALID-SCHEDULE rule.
func TestScheduleRulesEachWay(t *testing.T) {
	t.Parallel()
	runWays(t, []ruleWay{
		{
			name: "schedule-001-exclusive-mode",
			id:   RuleSchedule001,
			positive: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{
					Continuous: &schema.ContinuousSpec{Duration: schema.Second},
				}
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Schedule.Cron = &schema.CronSpec{Expression: "* * * * *"}
			},
			field: "spec.schedule",
		},
		{
			name: "schedule-002-cron-expression",
			id:   RuleSchedule002,
			positive: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{
					Cron: &schema.CronSpec{Expression: "*/5 0-23 * * 1-5", Timezone: "Virtual/Tokyo"},
				}
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{
					Cron: &schema.CronSpec{Expression: "60 * * * *", Timezone: "UTC"},
				}
			},
			field: "spec.schedule.cron.expression",
		},
		{
			name: "schedule-002-virtual-timezone",
			id:   RuleSchedule002,
			positive: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{
					Cron: &schema.CronSpec{Expression: "* * * * *", Timezone: "UTC"},
				}
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{
					Cron: &schema.CronSpec{Expression: "* * * * *", Timezone: "Asia/Shanghai"},
				}
			},
			field: "spec.schedule.cron.timezone",
		},
		{
			name: "schedule-003-continuous-duration",
			id:   RuleSchedule003,
			positive: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{
					Continuous: &schema.ContinuousSpec{Duration: 10 * schema.Second},
				}
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{
					Continuous: &schema.ContinuousSpec{Duration: 0},
				}
			},
			field: "spec.schedule.continuous.duration",
		},
		{
			name: "schedule-004-once-future",
			id:   RuleSchedule004,
			positive: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{Once: &schema.OnceSpec{At: 5 * schema.Second}}
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Schedule = schema.Schedule{Once: &schema.OnceSpec{At: 0}}
			},
			field: "spec.schedule.once.at",
		},
	})
}

// TestScheduleSkipBranches covers the arms where a schedule mode is unset.
func TestScheduleSkipBranches(t *testing.T) {
	t.Parallel()
	once := fresh(t)
	assertNoRule(t, ruleViolations(t, RuleSchedule002, once), RuleSchedule002)
	assertNoRule(t, ruleViolations(t, RuleSchedule003, once), RuleSchedule003)
	assertNoRule(t, ruleViolations(t, RuleSchedule004, once), RuleSchedule004)
	continuous := fresh(t)
	continuous.Spec.Schedule = schema.Schedule{
		Continuous: &schema.ContinuousSpec{Duration: schema.Second},
	}
	assertNoRule(t, ruleViolations(t, RuleSchedule004, continuous), RuleSchedule004)
}

// TestTargetRulesEachWay covers one positive and one negative example for
// every CES-VALID-TARGET rule.
func TestTargetRulesEachWay(t *testing.T) {
	t.Parallel()
	runWays(t, []ruleWay{
		{
			name:     "target-001-kind-enum",
			id:       RuleTarget001,
			positive: func(e *schema.Experiment) { e.Spec.Target.Kind = schema.TargetSource },
			negative: func(e *schema.Experiment) { e.Spec.Target.Kind = schema.TargetExternal },
			field:    "spec.target.kind",
		},
		{
			name:     "target-001-unknown-kind",
			id:       RuleTarget001,
			positive: func(e *schema.Experiment) { e.Spec.Target.Kind = schema.TargetSimulator },
			negative: func(e *schema.Experiment) { e.Spec.Target.Kind = "quantum" },
			field:    "spec.target.kind",
		},
		{
			name: "target-002-family-domain",
			id:   RuleTarget002,
			positive: func(e *schema.Experiment) {
				e.Spec.Method.Operators = []schema.OperatorDef{{
					ID: "c-delay", Family: schema.FamilyChaos,
					TargetKind: schema.TargetSimulator, Targets: []string{"sim"},
				}}
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Method.Operators = []schema.OperatorDef{{
					ID: "m-replace", Family: schema.FamilyMutation,
					TargetKind: schema.TargetSimulator, Targets: []string{"src"},
				}}
			},
			field: "spec.method.operators[m-replace].targetKind",
		},
		{
			name: "target-003-nonempty-selector",
			id:   RuleTarget003,
			positive: func(e *schema.Experiment) {
				e.Spec.Target.Selector = map[string]string{"app": "web", "node": "n1"}
			},
			negative: func(e *schema.Experiment) { e.Spec.Target.Selector = nil },
			field:    "spec.target.selector",
		},
		{
			name: "target-003-malformed-key",
			id:   RuleTarget003,
			positive: func(e *schema.Experiment) {
				e.Spec.Target.Selector = map[string]string{"node": "n1"}
			},
			negative: func(e *schema.Experiment) {
				e.Spec.Target.Selector = map[string]string{"bad key": "n1"}
			},
			field: "spec.target.selector[bad key]",
		},
		{
			name: "target-004-dual-selector",
			id:   RuleTarget004,
			positive: func(e *schema.Experiment) {
				applyMixedProduct(e)
				e.Spec.Target.Selector = map[string]string{
					schema.TargetSource: "src", schema.TargetSimulator: "sim",
				}
			},
			negative: func(e *schema.Experiment) {
				applyMixedProduct(e)
				e.Spec.Target.Selector = map[string]string{"node": "n1"}
			},
			field: "spec.target.selector",
		},
	})
}

// TestTarget002Branches covers the unknown-family and missing-domain arms.
func TestTarget002Branches(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name  string
		def   schema.OperatorDef
		field string
	}{
		{"unknown-family", schema.OperatorDef{
			ID: "ghost", Family: "quantum",
		}, "spec.method.operators[ghost]"},
		{"missing-target-kind", schema.OperatorDef{
			ID: "m-replace", Family: schema.FamilyMutation, Targets: []string{"src"},
		}, "spec.method.operators[m-replace].targetKind"},
	} {
		t.Run(c.name, func(t *testing.T) {
			exp := fresh(t)
			exp.Spec.Method.Operators = []schema.OperatorDef{c.def}
			assertViolationShape(t, ruleViolations(t, RuleTarget002, exp),
				RuleTarget002, c.field, SeverityError)
		})
	}
	legal := fresh(t)
	legal.Spec.Method.Operators = []schema.OperatorDef{
		{ID: "m-replace", Family: schema.FamilyMutation,
			TargetKind: schema.TargetSource, Targets: []string{"src"}},
		{ID: "c-delay", Family: schema.FamilyChaos,
			TargetKind: schema.TargetSimulator, Targets: []string{"sim"}},
	}
	assertNoRule(t, ruleViolations(t, RuleTarget002, legal), RuleTarget002)
}

// TestTarget003SelectorOrdering proves sorted keys yield a stable field path.
func TestTarget003SelectorOrdering(t *testing.T) {
	t.Parallel()
	exp := fresh(t)
	exp.Spec.Target.Selector = map[string]string{
		"z-key": "bad value", "a-key": "ok",
	}
	vs := ruleViolations(t, RuleTarget003, exp)
	if len(vs) != 1 {
		t.Fatalf("selector violations = %d, want 1: %+v", len(vs), vs)
	}
	assertViolationShape(t, vs, RuleTarget003,
		"spec.target.selector[z-key]", SeverityError)
}

// TestTarget004NonMixed covers the arms where the expression is not a mixed
// mutation-chaos concurrent product.
func TestTarget004NonMixed(t *testing.T) {
	t.Parallel()
	plain := fresh(t)
	assertNoRule(t, ruleViolations(t, RuleTarget004, plain), RuleTarget004)
	noChaos := fresh(t)
	noChaos.Spec.Method.Expression = "a ⊗ b"
	assertNoRule(t, ruleViolations(t, RuleTarget004, noChaos), RuleTarget004)
	oneTarget := fresh(t)
	applyMixedProduct(oneTarget)
	oneTarget.Spec.Target.Selector = map[string]string{schema.TargetSource: "src"}
	assertViolationShape(t, ruleViolations(t, RuleTarget004, oneTarget),
		RuleTarget004, "spec.target.selector", SeverityError)
}
