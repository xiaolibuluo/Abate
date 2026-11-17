/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktbuiltin

import (
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/ktfront"
	"abate/lang/rule"
)

// fixtureRootText is the attribute text of the synthetic AIR root. It is
// deliberately unlike any rule fixture text so no rule can hit the root.
const fixtureRootText = "fixture"

// airKind resolves a fixture kind name without any language frontend.
func airKind(name string) air.NodeKind {
	for k := air.KindInvalid; k <= air.KindRegion; k++ {
		if k.String() == name {
			return k
		}
	}
	return air.KindInvalid
}

// buildKTUnit builds a single-child synthetic AIR unit from one node spec. The
// optional note/async strings attach the matching semantic annotation.
func buildKTUnit(kind air.NodeKind, text, note, async string) *air.Unit {
	pos := air.NewPosition(1, 1, 0, len(text))
	child := air.NewNode(kind, pos, text)
	if note != "" {
		_ = air.Annotate(child, air.SemNote, note)
	}
	if async != "" {
		_ = air.Annotate(child, air.SemAsync, async)
	}
	root := air.NewNode(air.KindRegion, air.NewPosition(1, 1, 0, 0),
		fixtureRootText, child)
	return &air.Unit{
		Roots: []air.Node{root},
		Files: []air.FileRef{air.NewFileRef("fixture.air", []byte("fixture"))},
	}
}

// matchAny reports whether the compiled pattern of spec matches any node of unit.
func matchAny(spec rule.Spec, unit *air.Unit) bool {
	pattern, err := rule.CompilePattern(spec.Pattern)
	if err != nil {
		return false
	}
	matcher := rule.NewMatcher(pattern)
	found := false
	for _, root := range unit.Roots {
		air.Walk(root, func(n air.Node) bool {
			if len(matcher.Match(n)) > 0 {
				found = true
				return false
			}
			return true
		})
		if found {
			break
		}
	}
	return found
}

// findSpec returns the local specification carrying id, or the zero spec.
func findSpec(id rule.RuleID) rule.Spec {
	for _, spec := range All() {
		if spec.ID == id {
			return spec
		}
	}
	return rule.Spec{}
}

// ktLocalFixtures pairs each local rule id with a positive and a negative
// synthetic node spec. Every fixture is flat, so the shared builder applies.
var ktLocalFixtures = []struct {
	id  rule.RuleID
	pos []string
	neg []string
}{
	{"KT-DGN-001", []string{`Ident`, "my_var", "", ""}, []string{`Ident`, "myVar", "", ""}},
	{"KT-DGN-002", []string{`Ident`, "inline", "", ""}, []string{`Ident`, "inlineX", "", ""}},
	{"KT-DGN-003", []string{`Ident`, "suspend", "", ""}, []string{`Ident`, "suspendX", "", ""}},
	{"KT-DGN-004", []string{`Ident`, "when", "", ""}, []string{`Ident`, "whenX", "", ""}},
	{"KT-DGN-005", []string{`Unary`, "!!", "", ""}, []string{`Unary`, "x", "", ""}},
	{"KT-DGN-006", []string{`Optional`, "x", "", ""}, []string{`Ident`, "x", "", ""}},
	{"KT-DGN-007", []string{`NumLit`, "1000", "", ""}, []string{`NumLit`, "42", "", ""}},
	{"KT-DGN-008", []string{`Ident`, "var", "", ""}, []string{`Ident`, "varX", "", ""}},
	{"KT-DGN-009", []string{`Ident`, "lateinit", "", ""}, []string{`Ident`, "lateinitX", "", ""}},
	{"KT-DGN-010", []string{`Ident`, "data", "", ""}, []string{`Ident`, "dataX", "", ""}},
	{"KT-DGN-011", []string{`Ident`, "sealed", "", ""}, []string{`Ident`, "sealedX", "", ""}},
	{"KT-DGN-012", []string{`Ident`, "companion", "", ""}, []string{`Ident`, "companionX", "", ""}},
}

// TestCountAtLeastSixty proves the local set plus the common templates reach
// the required floor of sixty specifications.
func TestCountAtLeastSixty(t *testing.T) {
	t.Parallel()
	if Count() < 60 {
		t.Fatalf("Count() = %d, want >= 60", Count())
	}
}

// TestAllIDsUniqueAndReferencesNonEmpty proves every emitted identifier is
// unique and that no specification ships an empty official reference.
func TestAllIDsUniqueAndReferencesNonEmpty(t *testing.T) {
	t.Parallel()
	specs := All()
	if len(specs) != Count() {
		t.Fatalf("All() returned %d but Count() is %d", len(specs), Count())
	}
	seen := make(map[rule.RuleID]int, len(specs))
	for _, spec := range specs {
		if strings.TrimSpace(spec.Reference) == "" {
			t.Errorf("spec %s has an empty reference", spec.ID)
		}
		seen[spec.ID]++
	}
	for id, count := range seen {
		if count > 1 {
			t.Errorf("identifier %s appears %d times", id, count)
		}
	}
}

// TestLocalRulesHitPositiveMissNegative proves each local rule detects its
// positive fixture and stays silent on its negative fixture through the real
// AIR matcher.
func TestLocalRulesHitPositiveMissNegative(t *testing.T) {
	t.Parallel()
	for _, fx := range ktLocalFixtures {
		spec := findSpec(fx.id)
		if spec.ID == "" {
			t.Fatalf("missing local spec %s", fx.id)
		}
		pos := buildKTUnit(airKind(fx.pos[0]), fx.pos[1], fx.pos[2], fx.pos[3])
		neg := buildKTUnit(airKind(fx.neg[0]), fx.neg[1], fx.neg[2], fx.neg[3])
		if !matchAny(spec, pos) {
			t.Errorf("rule %s did not detect positive fixture %v", spec.ID, fx.pos)
		}
		if matchAny(spec, neg) {
			t.Errorf("rule %s hit negative fixture %v", spec.ID, fx.neg)
		}
	}
}

// TestEndToEndHits proves representative Kotlin rules really fire against the
// production ktfront frontend on genuine source constructs.
func TestEndToEndHits(t *testing.T) {
	t.Parallel()
	fe, err := ktfront.NewKtFrontend()
	if err != nil {
		t.Fatalf("NewKtFrontend: %v", err)
	}
	violation := "package demo\n" +
		"data class Point(val x: Int, var y: Int)\n" +
		"sealed class Shape\n" +
		"class Outer { companion object { const val MAX = 1000 } }\n" +
		"suspend fun fetch(): String { return \"\" }\n" +
		"val risky: String? = null\n" +
		"fun use(s: String?) { val v = s!! }\n" +
		"lateinit var name: String\n" +
		"inline fun id(x: Int) = x\n" +
		"fun compute() { val total = a + 12345; when(val k = 1) { 1 -> println(k) } }\n" +
		"fun bad_name() {}\n"
	unit, _, perr := fe.Parse(air.SourceUnit{
		Language: "kt",
		Files:    []air.SourceFile{air.NewSourceFile("sample.kt", []byte(violation))},
	})
	if perr != nil {
		t.Fatalf("Parse violation: %v", perr)
	}
	hits := []rule.RuleID{
		"KT-DGN-001", "KT-DGN-002", "KT-DGN-003", "KT-DGN-004", "KT-DGN-005",
		"KT-DGN-006", "KT-DGN-007", "KT-DGN-008", "KT-DGN-009", "KT-DGN-010",
		"KT-DGN-011", "KT-DGN-012",
	}
	for _, id := range hits {
		if !matchAny(findSpec(id), unit) {
			t.Errorf("rule %s did not hit the real-frontend violation unit", id)
		}
	}
}

// TestEndToEndCompliantMiss proves the naming, suspend, non-null, magic-number,
// mutable, lateinit, data, sealed, companion and inline rules stay silent on a
// clean Kotlin fragment.
func TestEndToEndCompliantMiss(t *testing.T) {
	t.Parallel()
	fe, err := ktfront.NewKtFrontend()
	if err != nil {
		t.Fatalf("NewKtFrontend: %v", err)
	}
	compliant := "package demo\n" +
		"class Service {\n" +
		"    val count = 1\n" +
		"    fun run() { val x = 1 }\n" +
		"}\n"
	unit, _, perr := fe.Parse(air.SourceUnit{
		Language: "kt",
		Files:    []air.SourceFile{air.NewSourceFile("good.kt", []byte(compliant))},
	})
	if perr != nil {
		t.Fatalf("Parse compliant: %v", perr)
	}
	misses := []rule.RuleID{
		"KT-DGN-001", "KT-DGN-002", "KT-DGN-003", "KT-DGN-004", "KT-DGN-005",
		"KT-DGN-006", "KT-DGN-007", "KT-DGN-008", "KT-DGN-009", "KT-DGN-010",
		"KT-DGN-011", "KT-DGN-012",
	}
	for _, id := range misses {
		if matchAny(findSpec(id), unit) {
			t.Errorf("rule %s wrongly hit the compliant unit", id)
		}
	}
}
