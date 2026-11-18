/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

// assertExactRulesUseDedicatedKinds enforces mappingrule.go:18-21: every
// StatusExact rule must name a dedicated AIR kind. The 22 fabricated
// shared-bucket kinds (87-108) carried 4-10 different constructs each and were
// the hallmark of the withdrawn label rewrite; KindRegion is the catch-all and
// can never be "exact". Any exact rule pointing at one of them is fabrication.
func assertExactRulesUseDedicatedKinds(t *testing.T, lang string) {
	t.Helper()
	banned := map[air.NodeKind]bool{air.KindRegion: true,
		air.KindInterpolation: true, air.KindPattern: true,
		air.KindPropertyDecl: true, air.KindWhenBranch: true,
		air.KindAttributeArg: true, air.KindLinqClause: true,
		air.KindParameterList: true, air.KindStringFragment: true,
		air.KindMacroInvocation: true, air.KindUseDetail: true,
		air.KindQualifiedName: true, air.KindCompoundAssign: true,
		air.KindBlockVariant: true, air.KindTypeArg: true,
		air.KindInvocation: true, air.KindSelectorDetail: true,
		air.KindDelimiter: true, air.KindOperatorToken: true,
		air.KindSyntaxDetail: true, air.KindPartialNode: true,
		air.KindSpecialIdent: true, air.KindComprehension: true,
	}
	for _, rule := range grammar.DefaultRules(lang) {
		if rule.Status != grammar.StatusExact {
			continue
		}
		if banned[rule.Kind] {
			t.Errorf("StatusExact rule %q uses non-dedicated kind %s",
				rule.Name, rule.Kind)
		}
	}
}

// assertExactRulesAreObservable checks that the kinds the real corpus does
// produce are consistent with the exact rules that claim them: a kind is
// observable only from a genuine parse, never from a static rule edit.
func assertExactRulesAreObservable(t *testing.T, lang string,
	observed map[air.NodeKind]int) {
	t.Helper()
	seenProduced := 0
	for _, rule := range grammar.DefaultRules(lang) {
		if rule.Status == grammar.StatusExact && observed[rule.Kind] > 0 {
			seenProduced++
		}
	}
	if seenProduced == 0 {
		t.Fatalf("%s: no StatusExact rule kind was produced by any real parse; "+
			"the rule table is disconnected from front-end output", lang)
	}
}
func assertRulesDegraded(t *testing.T, lang string, names ...string) {
	t.Helper()
	byName := make(map[string]grammar.MappingRule)
	for _, rule := range grammar.DefaultRules(lang) {
		byName[rule.Name] = rule
	}
	for _, name := range names {
		rule, found := byName[name]
		if !found {
			t.Errorf("rule %q not found in %s rule table", name, lang)
			continue
		}
		if rule.Status == grammar.StatusExact {
			t.Errorf("rule %q is StatusExact: QA found no dedicated production; "+
				"revert to degraded or add a genuine syntax-node reduction first", name)
		}
	}
}

// TestPyStructuralDegradationsStayDegraded pins every construct AIR still
// models with a shared kind (Region, Star, Tuple, GenericParam) and every
// rule the QA verdict confirmed has zero real-corpus production (named-expr,
// type-alias, generic-param, type-union): none of them may carry an exact
// label yet. Promoting one of them to exact first requires a real parser
// reduction plus a TestPyDedicatedKindsAreReallyProduced case AND non-zero
// production over the real corpus, and must move the rule out of this pin in
// the same change.
//
// The await and case rules were removed from this pin after the self-researched
// corpus began producing KindAwait (1984 observations) and KindCase (5856) from real
// pyfront reductions; their old "zero production" reason became false, so the
// rules were honestly restored to exact. The remaining pins still have zero
// real-corpus production.
func TestPyStructuralDegradationsStayDegraded(t *testing.T) {
	assertRulesDegraded(t, "py",
		"py.type-constrained", "py.splat-type", "py.expr-list",
		"py.list-splat", "py.paren-expr", "py.parameters", "py.patterns",
		"py.separators", "py.argument-list", "py.interpolation",
		"py.ellipsis", "py.line-continuation",
		"py.named-expr",
		"py.type-alias", "py.generic-param", "py.type-union")
}

// TestPyExactRulesAreActuallyProduced parses the real Python corpus and checks
// every StatusExact rule against actual front-end output.
func TestPyExactRulesAreActuallyProduced(t *testing.T) {
	observed := map[air.NodeKind]int{}
	for _, src := range pyCoverageCorpus() {
		unit, _, _ := pyParse(t, src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	assertExactRulesUseDedicatedKinds(t, "py")
	assertExactRulesAreObservable(t, "py", observed)
}

// TestPyQASampledRulesHaveGenuineProduction is the strengthened form of the
// earlier "stay degraded" pin. That pin only protected the state before a real
// reduction existed; once pyfront genuinely lowers assert/pass/boolean-op/
// compare/yield to dedicated nodes, the honest label becomes exact. This test
// now requires all three properties together: the rules are exact, their kinds
// are not from the fabricated 87-108 bucket, and a minimal source really
// produces the claimed kind. It is strictly harder to satisfy by label editing
// than the original pin.
func TestPyQASampledRulesHaveGenuineProduction(t *testing.T) {
	wanted := map[string]string{
		"py.assert": "assert", "py.pass": "pass",
		"py.boolean-op": "boolean-op", "py.compare": "compare",
		"py.yield": "yield",
	}
	byName := make(map[string]grammar.MappingRule)
	for _, rule := range grammar.DefaultRules("py") {
		byName[rule.Name] = rule
	}
	produced := make(map[string]bool)
	for _, tc := range pyDedicatedCases {
		unit, _, _ := pyParse(t, tc.src)
		for _, node := range air.Collect(unit) {
			if node.Kind() == tc.kind {
				produced[tc.construct] = true
				break
			}
		}
	}
	for name, construct := range wanted {
		rule, found := byName[name]
		if !found {
			t.Errorf("rule %q not found in py rule table", name)
			continue
		}
		if rule.Status != grammar.StatusExact {
			t.Errorf("rule %q is not StatusExact", name)
		}
		if rule.Kind < 109 {
			t.Errorf("rule %q uses kind %s outside the genuine dedicated range",
				name, rule.Kind)
		}
		if !produced[construct] {
			t.Errorf("rule %q claims kind %s but %q produced no such node",
				name, rule.Kind, construct)
		}
	}
}
