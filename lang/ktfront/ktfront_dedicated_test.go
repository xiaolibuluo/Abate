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

// ktDedicatedCases pairs a minimal Kotlin source with the dedicated kind the
// parser must produce while genuinely reducing that construct.
var ktDedicatedCases = []struct {
	construct string
	src       string
	kind      air.NodeKind
}{
	{"companion", "class C { companion object { } }\n", air.KindKtCompanion},
	{"secondary-ctor", "class C { constructor() { } }\n",
		air.KindKtSecondaryCtor},
	{"ctor-delegation", "class C { constructor(x: Int) : this(0) { } }\n",
		air.KindKtCtorDelegation},
	{"accessor", "class C { val p: Int get() = 1\n set(v) { } }\n",
		air.KindKtAccessor},
	{"import-alias", "import java.io.File as JFile\n", air.KindKtImportAlias},
	{"import-list", "import kotlin.collections.List\n", air.KindKtImportList},
	{"interpolated", "fun f(n: String) = \"hi $n\"\n", air.KindKtInterp},
	{"string-part", "val s = \"abc\"\n", air.KindKtStringFrag},
	{"property", "val x = 1\n", air.KindKtProperty},
	{"this-expr", "class C { fun f(): Int = this.hashCode()\n}\n",
		air.KindKtThis},
	{"super-expr", "open class B\n\nclass C : B() { fun f() = super.hashCode() }\n",
		air.KindKtSuper},
	{"when-entry", "fun f(x: Int): Int {\n\twhen (x) {\n\t\t1 -> return 2\n\t\telse -> return 3\n\t}\n\treturn 0\n}\n",
		air.KindKtWhenEntry},
	{"callable-ref", "val r = String::length\n", air.KindKtCallableRef},
	{"control-body", "fun f(x: Int) { if (x > 0) { return } }\n",
		air.KindKtControlBody},
	{"for-in", "fun f(xs: List<Int>) { for (x in xs) { println(x) } }\n",
		air.KindKtForIn},
}

// TestKtDedicatedKindsAreReallyProduced parses each minimal source and asserts
// the parser genuinely emitted the dedicated kind for that construct.
func TestKtDedicatedKindsAreReallyProduced(t *testing.T) {
	for _, tc := range ktDedicatedCases {
		unit, _, _ := ktParse(t, tc.src)
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

// TestKtQASampledRulesHaveGenuineProduction is the strengthened form of the
// earlier QA-absent pin (see pyfront for the pattern): every promoted rule
// must be exact, must name a genuine dedicated kind (>= 128, outside the
// fabricated 87-108 bucket), and a minimal source must really produce it.
func TestKtQASampledRulesHaveGenuineProduction(t *testing.T) {
	wanted := map[string]string{
		"kt.when-entry":   "when-entry",
		"kt.companion":    "companion",
		"kt.interpolated": "interpolated",
	}
	byName := make(map[string]grammar.MappingRule)
	for _, rule := range grammar.DefaultRules("kt") {
		byName[rule.Name] = rule
	}
	produced := make(map[string]bool)
	for _, tc := range ktDedicatedCases {
		unit, _, _ := ktParse(t, tc.src)
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
			t.Errorf("rule %q not found in kt rule table", name)
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

// TestKtExactRulesAreProducedFromRealCorpus re-checks the whole exact table
// against the real coverage corpus in one pass, mirroring pyfront.
func TestKtExactRulesAreProducedFromRealCorpus(t *testing.T) {
	observed := map[air.NodeKind]int{}
	for _, src := range ktCoverageCorpus() {
		unit, _, _ := ktParse(t, src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	assertExactRulesUseDedicatedKinds(t, "kt")
	assertExactRulesAreObservable(t, "kt", observed)
}
