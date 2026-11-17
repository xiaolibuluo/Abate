/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

// csDedicatedCases pairs a minimal C# source with the dedicated kind the
// parser must produce while genuinely reducing that construct.
var csDedicatedCases = []struct {
	construct string
	src       string
	kind      air.NodeKind
}{
	{"constructor", "class C { public C() { } }\n", air.KindCsConstructor},
	{"destructor", "class C { ~C() { } }\n", air.KindCsDestructor},
	{"operator", "class C { public static C operator +(C a, C b) => a; }\n",
		air.KindCsOperatorDecl},
	{"conversion-operator", "class C { public static implicit operator int(C c) => 1; }\n",
		air.KindCsOperatorDecl},
	{"local-func", "class C { void M() { int L(int v) { return v; } } }\n",
		air.KindCsLocalFunc},
	{"accessor", "class C { public int P { get; set; } }\n", air.KindCsAccessor},
	{"property", "class C { public int P { get; set; } }\n", air.KindCsProperty},
	{"indexer", "class C { public int this[int i] { get { return i; } } }\n",
		air.KindCsIndexer},
	{"event", "class C { public event System.EventHandler E; }\n",
		air.KindCsEventDecl},
	{"record", "public record R(int X);\n", air.KindCsRecord},
	{"yield", "class C { System.Collections.Generic.IEnumerable<int> M() { yield return 1; } }\n",
		air.KindCsYield},
	{"conditional-access", "class C { void M(C c) { var x = c?.P; } }\n",
		air.KindCsConditionalAccess},
	{"object-creation", "class C { void M() { var x = new C(); } }\n",
		air.KindCsObjectCreation},
	{"sizeof", "class C { void M() { var x = sizeof(int); } }\n",
		air.KindCsSizeof},
	{"typeof", "class C { void M() { var x = typeof(int); } }\n",
		air.KindCsTypeof},
	{"query", "class C { void M() { var q = from n in ns where n > 0 select n; } }\n",
		air.KindCsQuery},
	{"interpolated-string", "class C { void M() { var s = $\"a {b} c\"; } }\n",
		air.KindCsInterpString},
	{"string-part", "class C { void M() { var s = \"a\\nb\"; } }\n",
		air.KindCsStringFrag},
	{"char-body", "class C { void M() { var c = 'q'; } }\n", air.KindCsCharBody},
	{"modifier", "public class C { }\n", air.KindCsModifier},
	{"extern-alias", "extern alias Lib;\n", air.KindCsExternAlias},
	{"is-pattern", "class C { void M(object o) { var b = o is string s; } }\n",
		air.KindCsPattern},
	{"alias-qualified", "class C { void M() { var f = global::System.IO.File; } }\n",
		air.KindCsAliasName},
}

// TestCsDedicatedKindsAreReallyProduced parses each minimal source and asserts
// the parser genuinely emitted the dedicated kind for that construct.
func TestCsDedicatedKindsAreReallyProduced(t *testing.T) {
	for _, tc := range csDedicatedCases {
		unit, _, _ := csParse(t, tc.src)
		produced := false
		for _, node := range air.Collect(unit) {
			if node.Kind() == tc.kind {
				produced = true
				break
			}
		}
		if !produced {
			t.Errorf("construct %q: minimal source %q produced no %s node",
				tc.construct, tc.src, tc.kind)
		}
	}
}

// TestCsQASampledRulesHaveGenuineProduction is the strengthened form of the
// earlier QA-absent pin (see pyfront for the pattern): every promoted rule
// must be exact, must name a genuine dedicated kind (>= 128, outside the
// fabricated 87-108 bucket), and a minimal source must really produce it.
func TestCsQASampledRulesHaveGenuineProduction(t *testing.T) {
	wanted := map[string]string{
		"cs.constructor":         "constructor",
		"cs.query":               "query",
		"cs.interpolated-string": "interpolated-string",
		"cs.typeof":              "typeof",
	}
	byName := make(map[string]grammar.MappingRule)
	for _, rule := range grammar.DefaultRules("cs") {
		byName[rule.Name] = rule
	}
	produced := make(map[string]bool)
	for _, tc := range csDedicatedCases {
		unit, _, _ := csParse(t, tc.src)
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
			t.Errorf("rule %q not found in cs rule table", name)
			continue
		}
		if rule.Status != grammar.StatusExact {
			t.Errorf("rule %q is not StatusExact", name)
		}
		if rule.Kind < 128 {
			t.Errorf("rule %q uses kind %s outside the genuine dedicated range",
				name, rule.Kind)
		}
		if !produced[construct] {
			t.Errorf("rule %q claims kind %s but %q produced no such node",
				name, rule.Kind, construct)
		}
	}
}

// TestCsExactRulesAreProducedFromRealCorpus re-checks the whole exact table
// against the real coverage corpus in one pass, mirroring pyfront.
func TestCsExactRulesAreProducedFromRealCorpus(t *testing.T) {
	observed := map[air.NodeKind]int{}
	for _, src := range csCoverageCorpus() {
		unit, _, _ := csParse(t, src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	assertExactRulesUseDedicatedKinds(t, "cs")
	assertExactRulesAreObservable(t, "cs", observed)
}
