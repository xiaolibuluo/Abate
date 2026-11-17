/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csbuiltin

import (
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/csfront"
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

// buildCSUnit builds a single-child synthetic AIR unit from one node spec. The
// optional note/async strings attach the matching semantic annotation.
func buildCSUnit(kind air.NodeKind, text, note, async string) *air.Unit {
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

// csLocalFixtures pairs each local rule id with a positive and a negative
// synthetic node spec. Every fixture is flat, so the shared builder applies.
var csLocalFixtures = []struct {
	id  rule.RuleID
	pos []string
	neg []string
}{
	{"CS-DGN-001", []string{`Ident`, "my_var", "", ""}, []string{`Ident`, "MyVar", "", ""}},
	{"CS-DGN-002", []string{`Ident`, "MAX_COUNT", "", ""}, []string{`Ident`, "MaxCount", "", ""}},
	{"CS-DGN-003", []string{`Ident`, "async", "", ""}, []string{`Ident`, "asyncX", "", ""}},
	{"CS-DGN-004", []string{`Ident`, "from", "", ""}, []string{`Ident`, "fromX", "", ""}},
	{"CS-DGN-005", []string{`Optional`, "", "", ""}, []string{`Ident`, "x", "", ""}},
	{"CS-DGN-006", []string{`Ident`, "record", "", ""}, []string{`Ident`, "recordX", "", ""}},
	{"CS-DGN-007", []string{`Ident`, "Exception", "", ""}, []string{`Ident`, "ArgumentException", "", ""}},
	{"CS-DGN-008", []string{`Ident`, "using", "", ""}, []string{`Ident`, "usingX", "", ""}},
	{"CS-DGN-009", []string{`NumLit`, "1000", "", ""}, []string{`NumLit`, "42", "", ""}},
	{"CS-DGN-010", []string{`Ident`, "ToString", "", ""}, []string{`Ident`, "ToStringX", "", ""}},
	{"CS-DGN-011", []string{`Ident`, "get", "", ""}, []string{`Ident`, "getX", "", ""}},
	{"CS-DGN-012", []string{`Ident`, "sealed", "", ""}, []string{`Ident`, "sealedX", "", ""}},
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
	for _, fx := range csLocalFixtures {
		spec := findSpec(fx.id)
		if spec.ID == "" {
			t.Fatalf("missing local spec %s", fx.id)
		}
		pos := buildCSUnit(airKind(fx.pos[0]), fx.pos[1], fx.pos[2], fx.pos[3])
		neg := buildCSUnit(airKind(fx.neg[0]), fx.neg[1], fx.neg[2], fx.neg[3])
		if !matchAny(spec, pos) {
			t.Errorf("rule %s did not detect positive fixture %v", spec.ID, fx.pos)
		}
		if matchAny(spec, neg) {
			t.Errorf("rule %s hit negative fixture %v", spec.ID, fx.neg)
		}
	}
}

// TestEndToEndHits proves representative C# rules really fire against the
// production csfront frontend on genuine source constructs.
func TestEndToEndHits(t *testing.T) {
	t.Parallel()
	fe, err := csfront.NewCsFrontend()
	if err != nil {
		t.Fatalf("NewCsFrontend: %v", err)
	}
	violation := "using System;\n" +
		"public async Task<int> Fetch() { await Task.Delay(1); return 1; }\n" +
		"record Widget(int Id);\n" +
		"string? name;\n" +
		"int item_count = 5;\n" +
		"const int MAX_COUNT = 1000;\n" +
		"var q = from a in items where a > 0 select a;\n" +
		"class C { public int X { get; set; } }\n" +
		"sealed class D {}\n" +
		"try {} catch (Exception e) {}\n" +
		"public void Dump() { var s = ToString(); }\n"
	unit, _, perr := fe.Parse(air.SourceUnit{
		Language: "cs",
		Files:    []air.SourceFile{air.NewSourceFile("sample.cs", []byte(violation))},
	})
	if perr != nil {
		t.Fatalf("Parse violation: %v", perr)
	}
	hits := []rule.RuleID{
		"CS-DGN-001", "CS-DGN-002", "CS-DGN-003", "CS-DGN-004", "CS-DGN-005",
		"CS-DGN-006", "CS-DGN-007", "CS-DGN-008", "CS-DGN-009", "CS-DGN-010",
		"CS-DGN-011", "CS-DGN-012",
	}
	for _, id := range hits {
		if !matchAny(findSpec(id), unit) {
			t.Errorf("rule %s did not hit the real-frontend violation unit", id)
		}
	}
}

// TestEndToEndCompliantMiss proves the naming, exception, magic-number,
// property and sealing rules stay silent on a clean C# fragment.
func TestEndToEndCompliantMiss(t *testing.T) {
	t.Parallel()
	fe, err := csfront.NewCsFrontend()
	if err != nil {
		t.Fatalf("NewCsFrontend: %v", err)
	}
	compliant := "namespace App {\n" +
		"  public class Service {\n" +
		"    private int count;\n" +
		"    public void Run() { var x = 1; }\n" +
		"  }\n" +
		"}\n"
	unit, _, perr := fe.Parse(air.SourceUnit{
		Language: "cs",
		Files:    []air.SourceFile{air.NewSourceFile("good.cs", []byte(compliant))},
	})
	if perr != nil {
		t.Fatalf("Parse compliant: %v", perr)
	}
	misses := []rule.RuleID{
		"CS-DGN-001", "CS-DGN-002", "CS-DGN-007", "CS-DGN-009", "CS-DGN-011", "CS-DGN-012",
	}
	for _, id := range misses {
		if matchAny(findSpec(id), unit) {
			t.Errorf("rule %s wrongly hit the compliant unit", id)
		}
	}
}
