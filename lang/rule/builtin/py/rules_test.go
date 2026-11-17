/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pybuiltin

import (
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/pyfront"
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

// buildBinaryCompareUnit builds a Binary node with the operator text op and a
// right operand of the given kind and text, used for the nested None comparisons.
func buildBinaryCompareUnit(op, rightText string, rightKind air.NodeKind) *air.Unit {
	left := air.NewNode(air.KindIdent, air.NewPosition(1, 1, 0, 1), "x")
	right := air.NewNode(rightKind, air.NewPosition(1, 1, 1, 1+len(rightText)), rightText)
	bin := air.NewNode(air.KindBinary, air.NewPosition(1, 1, 0, 1+len(rightText)), op, left, right)
	root := air.NewNode(air.KindRegion, air.NewPosition(1, 1, 0, 0), fixtureRootText, bin)
	return &air.Unit{
		Roots: []air.Node{root},
		Files: []air.FileRef{air.NewFileRef("fixture.air", []byte("x"+op+rightText))},
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

// pyLocalFixtures pairs each local rule id with a positive and a negative
// synthetic fragment. The two nested rules use empty fragments resolved below.
var pyLocalFixtures = []struct {
	id  rule.RuleID
	pos string
	neg string
}{
	{"PY-PEP8-001", "Ident:badName", "Ident:good_name"},
	{"PY-PEP8-002", "ClassDecl:my_class", "ClassDecl:MyClass"},
	{"PY-PEP8-003", "Ident:l", "Ident:width"},
	{"PY-PEP8-004", "NumLit:1000", "NumLit:42"},
	{"PY-PEP8-005", "Directive:# noqa", "Comment:# plain"},
	{"PY-PEP8-006", "Directive:# type: ignore", "Comment:# plain"},
	{"PY-PEP8-007", "Star:*", "Ident:ok"},
	{"PY-PEP8-008", "Ident:eval", "Ident:my_func"},
	{"PY-PEP8-009", "Ident:exec", "Ident:my_func"},
	{"PY-PEP484-001", "", ""},
	{"PY-PEP484-002", "", ""},
	{"PY-PEP484-003", "Named:Any", "Named:Int"},
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
// AIR matcher, both for flat and nested patterns.
func TestLocalRulesHitPositiveMissNegative(t *testing.T) {
	t.Parallel()
	for _, fx := range pyLocalFixtures {
		spec := findSpec(fx.id)
		if spec.ID == "" {
			t.Fatalf("missing local spec %s", fx.id)
		}
		var posUnit, negUnit *air.Unit
		switch fx.id {
		case "PY-PEP484-001":
			posUnit = buildBinaryCompareUnit("==", "None", air.KindNilLit)
			negUnit = buildBinaryCompareUnit("is", "None", air.KindNilLit)
		case "PY-PEP484-002":
			posUnit = buildBinaryCompareUnit("!=", "None", air.KindNilLit)
			negUnit = buildBinaryCompareUnit("is", "None", air.KindNilLit)
		default:
			posUnit = buildFixtureUnit(fx.pos)
			negUnit = buildFixtureUnit(fx.neg)
		}
		if !matchAny(spec, posUnit) {
			t.Errorf("rule %s did not detect positive fixture", spec.ID)
		}
		if matchAny(spec, negUnit) {
			t.Errorf("rule %s hit negative fixture", spec.ID)
		}
	}
}

// TestPythonEndToEnd proves representative rules fire on source parsed by the
// real pyfront frontend, and that a compliant fragment stays silent.
func TestPythonEndToEnd(t *testing.T) {
	t.Parallel()
	fe, err := pyfront.NewPyFrontend()
	if err != nil {
		t.Fatalf("NewPyFrontend: %v", err)
	}
	violation := "badName = 1000\n" +
		"x == None\n" +
		"x != None\n" +
		"from os import *\n" +
		"result = eval(\"1+1\")\n" +
		"value: Any = None\n" +
		"# noqa\n" +
		"class my_class:\n" +
		"    pass\n"
	compliant := "def compute(items):\n" +
		"    count = 5\n" +
		"    if count is None:\n" +
		"        return 0\n" +
		"    return count\n"

	vUnit, _, err := fe.Parse(air.SourceUnit{
		Language: "py",
		Files:    []air.SourceFile{air.NewSourceFile("sample.py", []byte(violation))},
	})
	if err != nil {
		t.Fatalf("Parse violation: %v", err)
	}
	cUnit, _, err := fe.Parse(air.SourceUnit{
		Language: "py",
		Files:    []air.SourceFile{air.NewSourceFile("good.py", []byte(compliant))},
	})
	if err != nil {
		t.Fatalf("Parse compliant: %v", err)
	}

	hits := []rule.RuleID{
		"PY-PEP8-001", "PY-PEP8-002", "PY-PEP8-004", "PY-PEP8-005",
		"PY-PEP8-007", "PY-PEP8-008", "PY-PEP484-001", "PY-PEP484-002",
		"PY-PEP484-003",
	}
	for _, id := range hits {
		if !matchAny(findSpec(id), vUnit) {
			t.Errorf("rule %s did not hit the real-frontend violation unit", id)
		}
	}
	misses := []rule.RuleID{"PY-PEP8-001", "PY-PEP8-002", "PY-PEP484-001"}
	for _, id := range misses {
		if matchAny(findSpec(id), cUnit) {
			t.Errorf("rule %s wrongly hit the compliant unit", id)
		}
	}
}
