/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import "abate/ces/schema"

// rollbackRules returns the three rollback rules in registration order.
func rollbackRules() []Rule {
	return []Rule{
		newRule(RuleRollback001, checkRollback001),
		newRule(RuleRollback002, checkRollback002),
		newRule(RuleRollback003, checkRollback003),
	}
}

// checkRollback001 requires a legal rollback strategy enum value.
func checkRollback001(exp *schema.Experiment) []Violation {
	s := exp.Spec.Rollback.Strategy
	switch s {
	case schema.StrategyAuto, schema.StrategyManual, schema.StrategyTimeout:
		return nil
	default:
		return []Violation{errorAt(RuleRollback001, "spec.rollback.strategy",
			"rollback strategy must be auto, manual or timeout")}
	}
}

// checkRollback002 requires a positive timeout when strategy is timeout.
func checkRollback002(exp *schema.Experiment) []Violation {
	if exp.Spec.Rollback.Strategy != schema.StrategyTimeout {
		return nil
	}
	if exp.Spec.Rollback.Timeout <= 0 {
		return []Violation{errorAt(RuleRollback002, "spec.rollback.timeout",
			"timeout strategy requires a positive timeout")}
	}
	return nil
}

// checkRollback003 requires all steady-state rules to pass when the rollback
// re-verifies the steady state (it depends on a legal steady-state definition).
func checkRollback003(exp *schema.Experiment) []Violation {
	if !exp.Spec.Rollback.VerifySteadyState {
		return nil
	}
	if len(steadyViolations(exp)) > 0 {
		return []Violation{errorAt(RuleRollback003, "spec.rollback.verifySteadyState",
			"steady-state rules must all pass when verification is enabled")}
	}
	return nil
}

// steadyViolations re-runs the steady-state rule group for ROLLBACK-003.
func steadyViolations(exp *schema.Experiment) []Violation {
	var out []Violation
	for _, rule := range steadyRules() {
		out = append(out, rule.Check(exp)...)
	}
	return out
}
