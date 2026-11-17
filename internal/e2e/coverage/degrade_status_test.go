/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// fabricatedKinds are the 22 node kinds appended for the withdrawn G4 label
// rewrite. They were emitted by token-text switches (never by genuine syntax
// node reductions) and were used to relabel 78 degraded mappings as exact.
// A rule table that references any of them is the exact fraud signature this
// regression test must keep out.
var fabricatedKinds = []air.NodeKind{
	air.KindInterpolation, air.KindPattern, air.KindPropertyDecl,
	air.KindWhenBranch, air.KindAttributeArg, air.KindLinqClause,
	air.KindParameterList, air.KindStringFragment, air.KindMacroInvocation,
	air.KindUseDetail, air.KindQualifiedName, air.KindCompoundAssign,
	air.KindBlockVariant, air.KindTypeArg, air.KindInvocation,
	air.KindSelectorDetail, air.KindDelimiter, air.KindOperatorToken,
	air.KindSyntaxDetail, air.KindPartialNode, air.KindSpecialIdent,
	air.KindComprehension,
}

// TestDegradationNeverReclassifiesDegradedByKind enforces the promise at
// degrade.go:16-18: "a degraded mapping is never reclassified as exact to
// lower the ratio". Two mappings share one observed kind — one is exact, one
// is degraded. Degradation must still count the degraded one; it must not
// quietly promote it just because an exact rule uses the same kind.
func TestDegradationNeverReclassifiesDegradedByKind(t *testing.T) {
	mappings := []grammar.Mapping{
		{Node: "shared-exact", Kind: air.KindMethodDecl, Status: grammar.StatusExact},
		{Node: "shared-degraded", Kind: air.KindMethodDecl, Status: grammar.StatusDegraded},
	}
	observed := map[air.NodeKind]int{air.KindMethodDecl: 7}
	got := Degradation(mappings, observed)
	if got != 50 {
		t.Fatalf("Degradation() = %v, want 50: a degraded mapping sharing an "+
			"exact kind was reclassified to lower the ratio", got)
	}
}

// TestDegradationReclassificationWouldLowerRatio documents why the promise is
// load-bearing: promoting one degraded mapping to exact lowers the reported
// percentage. The function itself must keep reporting the declared statuses,
// so the test asserts both ratios and the strict decrease.
func TestDegradationReclassificationWouldLowerRatio(t *testing.T) {
	honest := []grammar.Mapping{
		{Node: "a", Kind: air.KindMethodDecl, Status: grammar.StatusExact},
		{Node: "b", Kind: air.KindCall, Status: grammar.StatusDegraded},
		{Node: "c", Kind: air.KindField, Status: grammar.StatusDegraded},
	}
	observed := map[air.NodeKind]int{
		air.KindMethodDecl: 1, air.KindCall: 1, air.KindField: 1,
	}
	before := Degradation(honest, observed)
	if before != 200.0/3.0 {
		t.Fatalf("honest Degradation() = %v, want %v", before, 200.0/3.0)
	}

	reclassified := []grammar.Mapping{
		{Node: "a", Kind: air.KindMethodDecl, Status: grammar.StatusExact},
		{Node: "b", Kind: air.KindCall, Status: grammar.StatusExact},
		{Node: "c", Kind: air.KindField, Status: grammar.StatusDegraded},
	}
	after := Degradation(reclassified, observed)
	if after >= before {
		t.Fatalf("reclassified Degradation() = %v, want < %v", after, before)
	}
	if after != 100.0/3.0 {
		t.Fatalf("reclassified Degradation() = %v, want %v", after, 100.0/3.0)
	}
}

// TestNoRuleTableReferencesFabricatedKinds is the direct regression guard for
// the 78-label rewrite: none of the four affected rule tables may reference a
// fabricated shared-bucket kind. Genuine dedicated kinds added later with real
// syntax-node reductions would require updating this list deliberately.
func TestNoRuleTableReferencesFabricatedKinds(t *testing.T) {
	banned := make(map[air.NodeKind]bool, len(fabricatedKinds))
	for _, kind := range fabricatedKinds {
		banned[kind] = true
	}
	for _, lang := range []string{"cs", "kt", "py", "rs"} {
		for _, rule := range grammar.DefaultRules(lang) {
			if banned[rule.Kind] {
				t.Errorf("%s rule %q references fabricated kind %s",
					lang, rule.Name, rule.Kind)
			}
		}
	}
}
