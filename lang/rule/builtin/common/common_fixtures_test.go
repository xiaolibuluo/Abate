/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/rule"
)

// fixtureRootText is the attribute text of the synthetic AIR root. It is
// deliberately unlike any template fixture text so no rule can hit the root.
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

// TestTemplatesCarryBothFixtureSides proves every declaration has one positive
// and one negative fragment, and that both fragments parse into real AIR nodes.
func TestTemplatesCarryBothFixtureSides(t *testing.T) {
	t.Parallel()
	templates := builtinTemplates()
	for i, item := range templates {
		if !item.Fixtures() {
			t.Fatalf("template %d (%s) is missing a fixture side", i, item.Family)
		}
		for _, fragment := range []string{item.Positive, item.Negative} {
			unit := buildFixtureUnit(fragment)
			if len(unit.Roots) != 1 || unit.NodeCount() < 2 {
				t.Fatalf("template %d fragment %q did not build a node", i, fragment)
			}
		}
	}
	keys := make(map[string]struct{}, len(templates))
	for _, item := range templates {
		if _, dup := keys[item.Key()]; dup {
			t.Fatalf("duplicate declaration key in family %s", item.Family)
		}
		keys[item.Key()] = struct{}{}
	}
}

// TestGeneratedRulesHitPositiveAndMissNegative proves every one of the sixty
// generated rules really detects its positive fixture and stays silent on its
// negative fixture through the real AIR matcher and predicate engine.
func TestGeneratedRulesHitPositiveAndMissNegative(t *testing.T) {
	t.Parallel()
	specs := All()
	templates := builtinTemplates()
	if len(specs) != len(templates) {
		t.Fatalf("spec/template pairing mismatch: %d vs %d", len(specs), len(templates))
	}
	for i, spec := range specs {
		item := templates[i]
		compiled, err := rule.Compile(spec)
		if err != nil {
			t.Fatalf("compile %s: %v", spec.ID, err)
		}
		positive := compiled.Check(buildFixtureUnit(item.Positive))
		if len(positive) == 0 {
			t.Errorf("rule %s (%s) did not detect positive fixture %q",
				spec.ID, item.Family, item.Positive)
		}
		negative := compiled.Check(buildFixtureUnit(item.Negative))
		if len(negative) != 0 {
			t.Errorf("rule %s (%s) hit negative fixture %q %d times",
				spec.ID, item.Family, item.Negative, len(negative))
		}
	}
}

// TestFixtureUnitsAreDeterministic proves the AIR constructor is repeatable.
func TestFixtureUnitsAreDeterministic(t *testing.T) {
	t.Parallel()
	first := buildFixtureUnit("Ident:bad_name|Comment:stub")
	second := buildFixtureUnit("Ident:bad_name|Comment:stub")
	if air.IDOf(first.Roots[0]) != air.IDOf(second.Roots[0]) {
		t.Fatal("identical fixture fragments produced different AIR node ids")
	}
	if first.NodeCount() != 3 {
		t.Fatalf("fixture node count = %d, want 3", first.NodeCount())
	}
	if got := airKind("NotAKind"); got != air.KindInvalid {
		t.Fatalf("unknown kind resolved to %v", got)
	}
}
