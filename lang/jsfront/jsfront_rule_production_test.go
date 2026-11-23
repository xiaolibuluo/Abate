/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"testing"

	"abate/lang/air"
	fkcov "abate/lang/frontkit/coverage"
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

// TestJsStructuralDegradationsStayDegraded pins every construct AIR still
// models with a shared kind: the object/array/assignment destructuring
// patterns have no dedicated AIR node yet, the optional chain and field
// definition lower onto the shared Field kind, the empty statement lowers
// onto Region and the labeled statement onto Label. None of them may carry an
// exact label yet. Promoting one of them to exact first requires a real parser
// reduction plus a TestDedicatedKindsReallyEmitted case AND non-zero
// production over the real corpus, and must move the rule out of this pin in
// the same change.
//
// The jsx, regex, assign-augmented, using-decl, static-block, class-expr,
// debugger, meta-property, computed-name, pair, pair-pattern,
// object-assign-pattern and rest-pattern rules were removed from this pin
// after jsfront gained genuine dedicated-kind reductions (KindJsJsx 218,
// KindJsRegex 215, KindJsAugAssign 205, KindJsUsingDecl 203,
// KindJsStaticBlock 199, KindJsClassExpr 197, KindJsDebugger 226,
// KindJsMetaProperty 227, KindJsComputedName 228, KindJsPair 211,
// KindJsPairPattern 229, KindJsObjectAssignPat 230, KindJsRestPattern 212)
// anchored on real tokens of each construct; their old "shared-kind only"
// reasons became false, so the rules were honestly restored to exact. The
// remaining pins still lower onto shared kinds only.
func TestJsStructuralDegradationsStayDegraded(t *testing.T) {
	assertRulesDegraded(t, "js",
		"js.optional-chain", "js.field-def", "js.label", "js.empty-stmt")
}

// TestJsExactRulesAreActuallyProduced parses the real JavaScript corpus and
// checks every StatusExact rule against actual front-end output.
func TestJsExactRulesAreActuallyProduced(t *testing.T) {
	observed := map[air.NodeKind]int{}
	fe := jsFrontend(t)
	for _, src := range jsCoverageCorpus() {
		unit, _ := jsParse(t, fe, "f.js", "", src)
		observed = fkcov.Merge(observed, fkcov.Observe(unit))
	}
	assertExactRulesUseDedicatedKinds(t, "js")
	assertExactRulesAreObservable(t, "js", observed)
}
