/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import "abate/ces/schema"

// steadyRules returns the five steady-state rules in registration order.
func steadyRules() []Rule {
	return []Rule{
		newRule(RuleSteady001, checkSteady001),
		newRule(RuleSteady002, checkSteady002),
		newRule(RuleSteady003, checkSteady003),
		newRule(RuleSteady004, checkSteady004),
		newRule(RuleSteady005, checkSteady005),
	}
}

// checkSteady001 enforces at least one steady-state hypothesis.
func checkSteady001(exp *schema.Experiment) []Violation {
	if len(exp.Spec.SteadyState) == 0 {
		return []Violation{errorAt(RuleSteady001, "spec.steadyState",
			"at least one steady state is required")}
	}
	return nil
}

// checkSteady002 rejects a steady state referencing an undeclared probe.
func checkSteady002(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, ss := range exp.Spec.SteadyState {
		if probeIndex(exp, ss.ProbeRef) < 0 {
			out = append(out, errorAt(RuleSteady002,
				path("spec.steadyState", i)+".probeRef",
				"steady state references an undeclared probe"))
		}
	}
	return out
}

// checkSteady003 requires HoldFor > 0 and no shorter than the probe interval.
func checkSteady003(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, ss := range exp.Spec.SteadyState {
		at := path("spec.steadyState", i) + ".holdFor"
		if ss.HoldFor <= 0 {
			out = append(out, errorAt(RuleSteady003, at,
				"holdFor must be positive"))
			continue
		}
		idx := probeIndex(exp, ss.ProbeRef)
		if idx >= 0 && ss.HoldFor < exp.Spec.Probes[idx].Interval {
			out = append(out, errorAt(RuleSteady003, at,
				"holdFor must be at least the probe interval"))
		}
	}
	return out
}

// checkSteady004 enforces the expectation/tolerance type matrix.
func checkSteady004(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, ss := range exp.Spec.SteadyState {
		at := path("spec.steadyState", i) + ".tolerance"
		if !knownExpectation(ss.Expectation.Type) {
			out = append(out, errorAt(RuleSteady004, at,
				"unknown expectation type"))
			continue
		}
		if ss.Tolerance.Type != "" && !knownTolerance(ss.Tolerance.Type) {
			out = append(out, errorAt(RuleSteady004, at,
				"unknown tolerance type"))
			continue
		}
		if !toleranceAllowed(ss.Expectation.Type, ss.Tolerance.Type) {
			out = append(out, errorAt(RuleSteady004, at,
				"expectation and tolerance types are incompatible"))
		}
	}
	return out
}

// checkSteady005 validates predicate syntax for predicate expectations.
func checkSteady005(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, ss := range exp.Spec.SteadyState {
		if ss.Expectation.Type != schema.ExpectPredicate {
			continue
		}
		if !validPredicate(ss.Expectation.Predicate) {
			out = append(out, errorAt(RuleSteady005,
				path("spec.steadyState", i)+".expectation.predicate",
				"predicate expression is not compilable"))
		}
	}
	return out
}
