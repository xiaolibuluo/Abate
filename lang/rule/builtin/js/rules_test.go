/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsbuiltin

import (
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/jsfront"
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

// buildFixtureUnit turns a compact "Kind:text|Kind:text" fragment into an AIR
// unit. It is a pure in-memory constructor: no frontend, no filesystem, no
// environment and no clock is involved.
func buildFixtureUnit(fragment string) *air.Unit {
	children := make([]air.Node, 0, 2)
	for _, token := range strings.Split(fragment, "|") {
		kindName, text, ok := strings.Cut(token, ":")
		if !ok {
			continue
		}
		kind := airKind(kindName)
		if kind == air.KindInvalid {
			continue
		}
		pos := air.NewPosition(1, 1, 0, len(text))
		children = append(children, air.NewNode(kind, pos, text))
	}
	root := air.NewNode(air.KindRegion, air.NewPosition(1, 1, 0, 0),
		fixtureRootText, children...)
	return &air.Unit{
		Roots: []air.Node{root},
		Files: []air.FileRef{air.NewFileRef("fixture.air", []byte(fragment))},
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

// jsLocalFixtures pairs each local rule id with a positive and a negative
// synthetic fragment. Every fixture is flat, so the shared builder applies.
var jsLocalFixtures = []struct {
	id  rule.RuleID
	pos string
	neg string
}{
	{"JS-ES-001", "Binary:==", "Binary:==="},
	{"JS-ES-002", "Binary:!=", "Binary:!=="},
	{"JS-ES-003", "VarDecl:", "ConstDecl:"},
	{"JS-ES-004", "With:", "If:"},
	{"JS-ES-005", "Ident:eval", "Ident:myFunc"},
	{"JS-ES-006", "NumLit:1000", "NumLit:42"},
	{"JS-ES-007", "Comment:// eslint", "Comment:// ok"},
	{"JS-ES-008", "Binary:instanceof", "Binary:=="},
	{"JS-ES-009", "Ident:alert", "Ident:log"},
	{"JS-TS-001", "Unary:void", "Unary:typeof"},
	{"JS-TS-002", "TypeAssert:", "Arrow:"},
	{"JS-TS-003", "EnumDecl:", "InterfaceDecl:"},
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
	for _, fx := range jsLocalFixtures {
		spec := findSpec(fx.id)
		if spec.ID == "" {
			t.Fatalf("missing local spec %s", fx.id)
		}
		if !matchAny(spec, buildFixtureUnit(fx.pos)) {
			t.Errorf("rule %s did not detect positive fixture %q", spec.ID, fx.pos)
		}
		if matchAny(spec, buildFixtureUnit(fx.neg)) {
			t.Errorf("rule %s hit negative fixture %q", spec.ID, fx.neg)
		}
	}
}

// TestJavaScriptEndToEnd proves representative rules fire on source parsed by
// the real jsfront frontend, and that a compliant fragment stays silent.
func TestJavaScriptEndToEnd(t *testing.T) {
	t.Parallel()
	fe, err := jsfront.NewJsFrontend()
	if err != nil {
		t.Fatalf("NewJsFrontend: %v", err)
	}
	violation := "var total = 1000;\n" +
		"const helper = eval(\"1+1\");\n" +
		"function run(x) {\n" +
		"  if (x == null) { return; }\n" +
		"  alert(\"x\");\n" +
		"}\n" +
		"var eq = y == null;\n" +
		"with (obj) { y = 1; }\n" +
		"enum Color { Red }\n" +
		"const z = w as Number;\n" +
		"/* eslint-disable */\n" +
		"void 0;\n" +
		"x instanceof Foo;\n"
	compliant := "const total = 5;\n" +
		"function run(items) {\n" +
		"  if (items === null) { return items; }\n" +
		"  return items.length;\n" +
		"}\n"

	vUnit, _, err := fe.Parse(air.SourceUnit{
		Language: "ts",
		Files:    []air.SourceFile{air.NewSourceFile("sample.ts", []byte(violation))},
	})
	if err != nil {
		t.Fatalf("Parse violation: %v", err)
	}
	cUnit, _, err := fe.Parse(air.SourceUnit{
		Language: "ts",
		Files:    []air.SourceFile{air.NewSourceFile("good.ts", []byte(compliant))},
	})
	if err != nil {
		t.Fatalf("Parse compliant: %v", err)
	}

	hits := []rule.RuleID{
		"JS-ES-001", "JS-ES-003", "JS-ES-005", "JS-ES-006", "JS-ES-007",
		"JS-ES-008", "JS-ES-009", "JS-TS-001", "JS-TS-002", "JS-TS-003", "JS-ES-004",
	}
	for _, id := range hits {
		if !matchAny(findSpec(id), vUnit) {
			t.Errorf("rule %s did not hit the real-frontend violation unit", id)
		}
	}
	misses := []rule.RuleID{"JS-ES-001", "JS-ES-003"}
	for _, id := range misses {
		if matchAny(findSpec(id), cUnit) {
			t.Errorf("rule %s wrongly hit the compliant unit", id)
		}
	}
}
