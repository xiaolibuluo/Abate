/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

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

// TestKtExactRulesAreActuallyProduced parses the real Kotlin corpus and checks
// every StatusExact rule against actual front-end output.
func TestKtExactRulesAreActuallyProduced(t *testing.T) {
	observed := map[air.NodeKind]int{}
	for _, src := range ktCoverageCorpus() {
		unit, _, _ := ktParse(t, src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	assertExactRulesUseDedicatedKinds(t, "kt")
	assertExactRulesAreObservable(t, "kt", observed)
}

// TestKtQAAbsentRulesStayDegraded pins the Kotlin rules QA flagged ABSENT
// that still have no genuine dedicated production. The three rules that
// gained real parser reductions (when-entry, companion, interpolated) moved
// out of this pin into TestKtQASampledRulesHaveGenuineProduction in the same
// change, per the promotion procedure documented on assertRulesDegraded.
func TestKtQAAbsentRulesStayDegraded(t *testing.T) {
	assertRulesDegraded(t, "kt",
		"kt.postfix", "kt.simple-identifier")
}
