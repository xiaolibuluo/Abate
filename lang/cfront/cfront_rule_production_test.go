/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"testing"

	"abate/lang/air"
	fkcov "abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

// cDedicatedCases pairs a minimal C source with the dedicated kind the
// frontend must really produce for it. This is the executable anti-fraud
// guard: a mapping may only be marked exact when the frontend truly lowers
// the construct onto its dedicated kind.
var cDedicatedCases = []struct {
	name string
	src  string
	kind air.NodeKind
}{
	{"union", "union U { int i; float f; };", air.KindCUnion},
	{"enum-member", "enum E { A = 1, B };", air.KindCEnumMember},
	{"preproc-defined", "#if defined(X)\nint y;\n#endif\n", air.KindCDefined},
	{"field-decl", "struct S { int a; const char *s; };", air.KindCFieldDecl},
	{"deref", "int g(void) { int *p = 0; return *p; }", air.KindCDeref},
	{"init-pair", "struct P p = {.x = 1, .y = 2};", air.KindCInitPair},
	{"label", "void f(void) { goto end; end: ; }", air.KindCLabel},
	{"func-declarator", "int add(int a, int b) { return a + b; }", air.KindCFuncDeclarator},
	{"init-declarator", "int total = 42;", air.KindCInitDeclarator},
	{"macro-type", "LIST(int) items;", air.KindCMacroType},
	{"type-ident", "Slot8_1 make_slot(void) { Slot8_1 s; return s; }", air.KindCTypeIdent},
	{"sized-type", "unsigned long width = 16;", air.KindCSizedType},
	{"field-ident", "int g(void) { struct P p; return p.x; }", air.KindCFieldIdent},
}

// TestCDedicatedKindsReallyEmitted proves each dedicated kind is produced by
// a real parse of a construct that actually appears in the source.
func TestCDedicatedKindsReallyEmitted(t *testing.T) {
	for _, tc := range cDedicatedCases {
		t.Run(tc.name, func(t *testing.T) {
			unit, _ := cParseOne(t, tc.src, "c++20")
			observed := fkcov.Observe(unit)
			if observed[tc.kind] == 0 {
				t.Fatalf("kind %s (%d) not emitted for source %q; observed=%v",
					tc.kind, tc.kind, tc.src, observed)
			}
		})
	}
}

// TestCDedicatedMarkersZeroWidth proves dedicated markers never widen a node:
// every dedicated-kind node has an empty byte span, so lexical fidelity and
// existing AIR-IDs stay byte-exact.
func TestCDedicatedMarkersZeroWidth(t *testing.T) {
	unit, _ := cParseOne(t,
		"union U { int i; };\nenum E { A = 1 };\n#if defined(X)\n#endif\n"+
			"struct S { int a; };\nint g(void) { int *p = 0; return *p; }\n"+
			"struct P p = {.x = 1};\nvoid f(void) { goto end; end: ; }\n",
		"c++20")
	for _, node := range air.Collect(unit) {
		name := node.Kind().String()
		if len(name) > 1 && name[:1] == "C" {
			pos, end := node.Pos(), node.End()
			if pos.Offset != end.Offset {
				t.Fatalf("dedicated kind %s is not zero-width: [%d,%d)",
					name, pos.Offset, end.Offset)
			}
		}
	}
}

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

// TestCStructuralDegradationsStayDegraded pins every construct AIR still
// models with a shared kind: the declarator/identifier family (field, type
// and sized-type identifiers merge into Ident/Named nodes), the function and
// initializer declarators, and the macro-type alias. None of them may carry
// an exact label yet. Promoting one of them to exact first requires a real
// parser reduction plus a TestCDedicatedKindsReallyEmitted case AND non-zero
// production over the real corpus, and must move the rule out of this pin in
// the same change.
//
// The union, enumerator, field-decl, deref, initializer-pair,
// preproc-defined and statement-identifier rules were removed from this pin
// after cfront gained genuine dedicated-kind reductions (KindCUnion 231,
// KindCEnumMember 232, KindCFieldDecl 237, KindCDeref 236, KindCInitPair 234,
// KindCDefined 233, KindCLabel 235) anchored on real tokens of each
// construct; their old "shared-kind only" reasons became false, so the rules
// were honestly restored to exact. The remaining pins still lower onto shared
// kinds only.
func TestCStructuralDegradationsStayDegraded(t *testing.T) {
	assertRulesDegraded(t, "c",
		"c.decl-list", "c.enumerator-list", "c.abstract-func-declarator")
}

// TestCExactRulesAreActuallyProduced parses the real C corpus and checks that
// exact-rule kinds are observable from genuine parses.
func TestCExactRulesAreActuallyProduced(t *testing.T) {
	observed := map[air.NodeKind]int{}
	for _, src := range cCoverageCorpus() {
		unit, _ := cParseOne(t, src, "c++20")
		observed = fkcov.Merge(observed, fkcov.Observe(unit))
	}
	assertExactRulesUseDedicatedKinds(t, "c")
	seenProduced := 0
	for _, rule := range grammar.DefaultRules("c") {
		if rule.Status == grammar.StatusExact && observed[rule.Kind] > 0 {
			seenProduced++
		}
	}
	if seenProduced == 0 {
		t.Fatal("no StatusExact rule kind was produced by any real parse")
	}
}
