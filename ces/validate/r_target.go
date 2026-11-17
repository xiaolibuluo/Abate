/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"slices"
	"strings"

	"abate/ces/schema"
)

// targetRules returns the four target rules in registration order.
func targetRules() []Rule {
	return []Rule{
		newRule(RuleTarget001, checkTarget001),
		newRule(RuleTarget002, checkTarget002),
		newRule(RuleTarget003, checkTarget003),
		newRule(RuleTarget004, checkTarget004),
	}
}

// checkTarget001 enforces a legal target kind; external is explicitly rejected.
func checkTarget001(exp *schema.Experiment) []Violation {
	kind := exp.Spec.Target.Kind
	switch kind {
	case schema.TargetSimulator, schema.TargetSource:
		return nil
	case schema.TargetExternal:
		return []Violation{errorAt(RuleTarget001, "spec.target.kind",
			"external targets are not supported in CES v1")}
	default:
		return []Violation{errorAt(RuleTarget001, "spec.target.kind",
			"target kind must be simulator or source")}
	}
}

// expectedDomain maps an operator family to its required target domain.
func expectedDomain(family string) (string, bool) {
	switch family {
	case schema.FamilyMutation:
		return schema.TargetSource, true
	case schema.FamilyChaos:
		return schema.TargetSimulator, true
	default:
		return "", false
	}
}

// checkTarget002 enforces the operator-family / target-domain matching matrix.
func checkTarget002(exp *schema.Experiment) []Violation {
	var out []Violation
	for _, def := range exp.Spec.Method.Operators {
		want, ok := expectedDomain(def.Family)
		if !ok {
			out = append(out, errorAt(RuleTarget002, "spec.method.operators["+def.ID+"]",
				"operator family is not mutation or chaos"))
			continue
		}
		if def.TargetKind == "" {
			out = append(out, errorAt(RuleTarget002,
				"spec.method.operators["+def.ID+"].targetKind",
				"operator must declare its target domain"))
			continue
		}
		if def.TargetKind != want {
			out = append(out, errorAt(RuleTarget002,
				"spec.method.operators["+def.ID+"].targetKind",
				def.Family+" operators must target "+want))
		}
	}
	return out
}

// checkTarget003 requires a non-empty selector with well-formed keys and values.
func checkTarget003(exp *schema.Experiment) []Violation {
	sel := exp.Spec.Target.Selector
	if len(sel) == 0 {
		return []Violation{errorAt(RuleTarget003, "spec.target.selector",
			"selector must not be empty")}
	}
	keys := make([]string, 0, len(sel))
	for k := range sel {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	var out []Violation
	for _, k := range keys {
		if !selectorKeyValid(k) || !selectorKeyValid(sel[k]) {
			out = append(out, errorAt(RuleTarget003, "spec.target.selector["+k+"]",
				"selector key or value is malformed"))
		}
	}
	return out
}

// mixedProduct reports whether an expression combines a mutation and a chaos
// operator through a concurrent product, which needs dual-target declaration.
func mixedProduct(exp *schema.Experiment) bool {
	expr := exp.Spec.Method.Expression
	if !expressionHasProduct(expr) {
		return false
	}
	byID := make(map[string]schema.OperatorDef, len(exp.Spec.Method.Operators))
	for _, def := range exp.Spec.Method.Operators {
		byID[def.ID] = def
	}
	hasMutation, hasChaos := false, false
	for _, id := range referencedIDs(expr) {
		if def, ok := byID[id]; ok {
			if def.Family == schema.FamilyMutation {
				hasMutation = true
			}
			if def.Family == schema.FamilyChaos {
				hasChaos = true
			}
		}
	}
	// Inline identifiers encode their family in the m-/c- prefix.
	for _, tok := range strings.FieldsFunc(expr, func(r rune) bool {
		return !idRune(r)
	}) {
		if strings.HasPrefix(tok, "m-") {
			hasMutation = true
		}
	}
	return hasMutation && hasChaos
}

// checkTarget004 requires a mixed M-product-C expression to declare both a
// source and a simulator selector (dual-target declaration).
func checkTarget004(exp *schema.Experiment) []Violation {
	if !mixedProduct(exp) {
		return nil
	}
	sel := exp.Spec.Target.Selector
	_, hasSource := sel[schema.TargetSource]
	_, hasSimulator := sel[schema.TargetSimulator]
	if hasSource && hasSimulator {
		return nil
	}
	return []Violation{errorAt(RuleTarget004, "spec.target.selector",
		"mixed M-product-C expression requires source and simulator selectors")}
}
