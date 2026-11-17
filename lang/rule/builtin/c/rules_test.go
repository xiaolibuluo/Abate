/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cbuiltin

import (
	"fmt"
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/cfront"
	"abate/lang/rule"
)

// airKind resolves a fixture kind name without any language frontend.
func airKind(name string) air.NodeKind {
	for k := air.KindInvalid; k <= air.KindRegion; k++ {
		if k.String() == name {
			return k
		}
	}
	return air.KindInvalid
}

// mkNode builds one synthetic AIR leaf node with no semantic annotation.
func mkNode(kind, text string) air.Node {
	pos := air.NewPosition(1, 1, 0, len(text))
	return air.NewNode(airKind(kind), pos, text)
}

// mkUnit wraps synthetic children under a deterministic Region root.
func mkUnit(children ...air.Node) *air.Unit {
	root := air.NewNode(air.KindRegion, air.NewPosition(1, 1, 0, 0), "fixture", children...)
	return &air.Unit{
		Roots: []air.Node{root},
		Files: []air.FileRef{air.NewFileRef("fixture.air", []byte("fixture"))},
	}
}

// fixtureCase pairs a rule identifier with a synthetic positive and negative
// AIR node that the rule must detect and must not detect.
type fixtureCase struct {
	id  string
	pos air.Node
	neg air.Node
}

// localFixtures returns the synthetic positive/negative nodes for every local
// C and C++ rule, modelled directly on the shared AIR node text produced by the
// cfront frontend.
func localFixtures() []fixtureCase {
	return []fixtureCase{
		{"C-CERT-001", mkNode("Directive", "define"), mkNode("Directive", "include")},
		{"C-CERT-002", mkNode("NumLit", "42"), mkNode("Ident", "x")},
		{"C-CERT-003", mkNode("Ident", "malloc"), mkNode("Ident", "memcpy")},
		{"C-CERT-004", mkNode("Star", "*"), mkNode("Ident", "x")},
		{"C-CERT-005", mkNode("NilLit", "NULL"), mkNode("Ident", "x")},
		{"C-CERT-006", mkNode("Ident", "strcpy"), mkNode("Ident", "memmove")},
		{"C-CERT-007", mkNode("Ident", "printf"), mkNode("Ident", "fprintf")},
		{"C-CERT-008", mkNode("Cast", "sizeof"), mkNode("Cast", "const_cast")},
		{"C-CERT-009", mkNode("IncDec", "++"), mkNode("Ident", "x")},
		{"C-CPP-001", mkNode("Unary", "new"), mkNode("Ident", "x")},
		{"C-CPP-002", mkNode("Throw", "e"), mkNode("Ident", "x")},
		{"C-CPP-003", mkNode("Cast", "reinterpret_cast"), mkNode("Cast", "static_cast")},
	}
}

// ruleByID indexes the declared local rules by identifier.
func ruleByID() map[string]rule.Spec {
	out := make(map[string]rule.Spec, 16)
	for _, spec := range cRules() {
		out[string(spec.ID)] = spec
	}
	return out
}

// walkMatch reports whether any node in the unit satisfies the compiled pattern.
func walkMatch(matcher rule.Matcher, unit *air.Unit) bool {
	for _, root := range unit.Roots {
		hit := false
		air.Walk(root, func(current air.Node) bool {
			if matcher.Matches(current) {
				hit = true
				return false
			}
			return true
		})
		if hit {
			return true
		}
	}
	return false
}

// TestRuleSetCount proves the declared C rule set clears the M5 threshold.
func TestRuleSetCount(t *testing.T) {
	t.Parallel()
	if got := Count(); got < 60 {
		t.Fatalf("rule count = %d, want at least 60", got)
	}
}

// TestRuleDeclarations proves identifier validity, uniqueness, reference
// coverage and pattern compilability of every declared rule.
func TestRuleDeclarations(t *testing.T) {
	t.Parallel()
	specs := All()
	seen := make(map[string]struct{}, len(specs))
	previous := ""
	for _, spec := range specs {
		id := string(spec.ID)
		if !spec.ID.Valid() {
			t.Fatalf("invalid rule id %q", id)
		}
		if _, ok := seen[id]; ok {
			t.Fatalf("duplicate rule id %q", id)
		}
		seen[id] = struct{}{}
		if strings.Compare(previous, id) > 0 {
			t.Fatalf("All() is not id-ordered: %s after %s", id, previous)
		}
		previous = id
		if strings.TrimSpace(spec.Reference) == "" {
			t.Fatalf("rule %s has no official reference", id)
		}
		if _, err := rule.Compile(spec); err != nil {
			t.Fatalf("rule %s does not compile: %v", id, err)
		}
	}
}

// TestLocalRulesHitPositiveAndMissNegative proves every local rule detects its
// synthetic positive fixture and stays silent on its synthetic negative fixture.
func TestLocalRulesHitPositiveAndMissNegative(t *testing.T) {
	t.Parallel()
	byID := ruleByID()
	for _, item := range localFixtures() {
		spec, ok := byID[item.id]
		if !ok {
			t.Fatalf("no rule declared for fixture %s", item.id)
		}
		compiled, err := rule.CompilePattern(spec.Pattern)
		if err != nil {
			t.Fatalf("compile %s pattern %q: %v", item.id, spec.Pattern, err)
		}
		matcher := rule.NewMatcher(compiled)
		if !walkMatch(matcher, mkUnit(item.pos)) {
			t.Errorf("rule %s (%s) did not detect its positive fixture", item.id, spec.Pattern)
		}
		if walkMatch(matcher, mkUnit(item.neg)) {
			t.Errorf("rule %s (%s) hit its negative fixture", item.id, spec.Pattern)
		}
	}
}

// parseC parses one in-memory C or C++ source buffer into an AIR unit.
func parseC(t *testing.T, source string) *air.Unit {
	t.Helper()
	unit := air.SourceUnit{
		Language: "c",
		Files:    []air.SourceFile{air.NewSourceFile("sample.c", []byte(source))},
	}
	frontend, ferr := cfront.NewCFrontend()
	if ferr != nil {
		t.Fatalf("construct cfront: %v", ferr)
	}
	built, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatalf("parse C sample: %v", err)
	}
	if built == nil {
		t.Fatal("parse C sample returned a nil unit")
	}
	return built
}

// assertHit proves the named rule matches at least one node of the unit.
func assertHit(t *testing.T, id, source string) {
	t.Helper()
	spec, ok := ruleByID()[id]
	if !ok {
		t.Fatalf("no local rule %s", id)
	}
	compiled, err := rule.CompilePattern(spec.Pattern)
	if err != nil {
		t.Fatalf("compile %s: %v", id, err)
	}
	matcher := rule.NewMatcher(compiled)
	unit := parseC(t, source)
	if !walkMatch(matcher, unit) {
		t.Errorf("rule %s (%s) failed to hit real C frontend output", id, spec.Pattern)
	}
}

// assertMiss proves the named rule matches no node of the unit.
func assertMiss(t *testing.T, id, source string) {
	t.Helper()
	spec, ok := ruleByID()[id]
	if !ok {
		t.Fatalf("no local rule %s", id)
	}
	compiled, err := rule.CompilePattern(spec.Pattern)
	if err != nil {
		t.Fatalf("compile %s: %v", id, err)
	}
	matcher := rule.NewMatcher(compiled)
	unit := parseC(t, source)
	if walkMatch(matcher, unit) {
		t.Errorf("rule %s (%s) unexpectedly hit a compliant C snippet", id, spec.Pattern)
	}
}

// TestEndToEndHits proves representative C and C++ rules really fire against
// the production cfront frontend on genuine source constructs.
func TestEndToEndHits(t *testing.T) {
	t.Parallel()
	assertHit(t, "C-CERT-001", "#define MAX 100\nint x = MAX;\n")
	assertHit(t, "C-CERT-002", "int x = 42;\n")
	assertHit(t, "C-CERT-003", "void *p = malloc(10);\n")
	assertHit(t, "C-CERT-004", "int x = *p;\n")
	assertHit(t, "C-CERT-005", "int *p = NULL;\nif (p == NULL) {}\n")
	assertHit(t, "C-CERT-006", "char buf[8];\nstrcpy(buf, \"hi\");\n")
	assertHit(t, "C-CERT-007", "printf(\"hello\");\n")
	assertHit(t, "C-CERT-008", "size_t n = sizeof(int);\n")
	assertHit(t, "C-CERT-009", "void f() { int i = 0; ++i; }\n")
	assertHit(t, "C-CPP-001", "int *p = new int(5);\n")
	assertHit(t, "C-CPP-002", "void f() { throw 1; }\n")
	assertHit(t, "C-CPP-003", "int *q = reinterpret_cast<int *>(p);\n")
}

// TestEndToEndCompliantMiss proves the same rules stay silent on clean C.
func TestEndToEndCompliantMiss(t *testing.T) {
	t.Parallel()
	assertMiss(t, "C-CERT-001", "#include <stdio.h>\nint x = 1;\n")
	assertMiss(t, "C-CERT-002", "int x = y;\n")
	assertMiss(t, "C-CERT-003", "memcpy(buf, src, n);\n")
	assertMiss(t, "C-CERT-004", "int x = y + 1;\n")
	assertMiss(t, "C-CERT-005", "if (p) {}\n")
	assertMiss(t, "C-CERT-006", "memmove(buf, src, n);\n")
	assertMiss(t, "C-CERT-007", "puts(\"hello\");\n")
	assertMiss(t, "C-CERT-008", "int n = 4;\n")
	assertMiss(t, "C-CERT-009", "void f() { int i = 0; i = i + 1; }\n")
	assertMiss(t, "C-CPP-001", "int *p = (int *)malloc(5);\n")
	assertMiss(t, "C-CPP-002", "void f() { return; }\n")
	assertMiss(t, "C-CPP-003", "int *q = (int *)p;\n")
}

// TestEndToEndPositionCoverage proves the matched node carries the source
// position of the offending construct through the real frontend.
func TestEndToEndPositionCoverage(t *testing.T) {
	t.Parallel()
	spec := ruleByID()["C-CERT-001"]
	compiled, err := rule.CompilePattern(spec.Pattern)
	if err != nil {
		t.Fatalf("compile %s: %v", spec.ID, err)
	}
	matcher := rule.NewMatcher(compiled)
	unit := parseC(t, "#define MAX 100\nint x = MAX;\n")
	var captured rule.Capture
	found := false
	for _, root := range unit.Roots {
		air.Walk(root, func(current air.Node) bool {
			caps := matcher.Match(current)
			if len(caps) > 0 {
				captured = caps[0]
				found = true
				return false
			}
			return true
		})
		if found {
			break
		}
	}
	if !found {
		t.Fatal("C-CERT-001 did not capture any node in the real unit")
	}
	if captured.Line != 1 {
		t.Fatalf("expected directive on line 1, got line %d", captured.Line)
	}
	if captured.Kind != "Directive" {
		t.Fatalf("expected captured kind Directive, got %q", captured.Kind)
	}
	_ = fmt.Sprintf("captured %s at %d:%d", captured.Kind, captured.Line, captured.Column)
}
