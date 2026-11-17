/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import "abate/ces/schema"

// methodRules returns the four method rules in registration order.
func methodRules() []Rule {
	return []Rule{
		newRule(RuleMethod001, checkMethod001),
		newRule(RuleMethod002, checkMethod002),
		newRule(RuleMethod003, checkMethod003),
		newRule(RuleMethod004, checkMethod004),
	}
}

// checkMethod001 requires a non-empty, compilable expression.
func checkMethod001(exp *schema.Experiment) []Violation {
	const at = "spec.method.expression"
	expr := exp.Spec.Method.Expression
	if expr == "" {
		return []Violation{errorAt(RuleMethod001, at, "expression is empty")}
	}
	_, err := compileMethod(exp)
	if err != nil {
		return []Violation{errorAt(RuleMethod001, at,
			"expression does not compile: "+err.Error())}
	}
	return nil
}

// checkMethod002 requires every referenced non-invertible operator to declare
// an explicit compensation strategy.
func checkMethod002(exp *schema.Experiment) []Violation {
	refs := referencedIDs(exp.Spec.Method.Expression)
	if len(refs) == 0 {
		return nil
	}
	byID := make(map[string]schema.OperatorDef, len(exp.Spec.Method.Operators))
	for _, def := range exp.Spec.Method.Operators {
		byID[def.ID] = def
	}
	var out []Violation
	seen := make(map[string]bool, len(refs))
	for _, id := range refs {
		if seen[id] {
			continue
		}
		seen[id] = true
		def, ok := byID[id]
		if ok && !def.Invertible && def.Compensation == "" {
			out = append(out, errorAt(RuleMethod002,
				"spec.method.operators["+id+"].compensation",
				"non-invertible operator must declare a compensation strategy"))
		}
	}
	return out
}

// checkMethod003 enforces the operator-tree depth bound of 64 via the compiler.
func checkMethod003(exp *schema.Experiment) []Violation {
	_, err := compileMethod(exp)
	depth, _, _ := compileFailure(err)
	if depth {
		return []Violation{errorAt(RuleMethod003, "spec.method.expression",
			"operator tree depth exceeds 64")}
	}
	return nil
}

// checkMethod004 rejects a concurrent product whose operands share a target.
func checkMethod004(exp *schema.Experiment) []Violation {
	if !expressionHasProduct(exp.Spec.Method.Expression) {
		return nil
	}
	_, err := compileMethod(exp)
	_, conflict, _ := compileFailure(err)
	if conflict {
		return []Violation{errorAt(RuleMethod004, "spec.method.expression",
			"concurrent product operands share a target key")}
	}
	return nil
}
