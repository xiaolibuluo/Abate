/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsbuiltin

import (
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/rsfront"
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

// mkNoted builds a synthetic node carrying a SemNote annotation.
func mkNoted(kind, text, note string) air.Node {
	n := mkNode(kind, text)
	_ = air.Annotate(n, air.SemNote, note)
	return n
}

// mkMut builds a synthetic node carrying a SemMutability annotation.
func mkMut(kind, text string) air.Node {
	n := mkNode(kind, text)
	_ = air.Annotate(n, air.SemMutability, "mutable")
	return n
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
// Rust rule, modelled directly on the shared AIR node text produced by the
// rsfront frontend.
func localFixtures() []fixtureCase {
	return []fixtureCase{
		{"RS-API-001", mkNode("Ident", "unwrap"), mkNode("Ident", "wrapped")},
		{"RS-API-002", mkNode("Ident", "expect"), mkNode("Ident", "inspect")},
		{"RS-API-003", mkNode("Ident", "panic"), mkNode("Ident", "println")},
		{"RS-API-004", mkNode("Ident", "todo"), mkNode("Ident", "done")},
		{"RS-API-005", mkNode("Ident", "unimplemented"), mkNode("Ident", "implemented")},
		{"RS-API-006", mkNode("Ident", "clone"), mkNode("Ident", "copy")},
		{"RS-API-007", mkNoted("Block", "", "unsafe"), mkNode("Block", "")},
		{"RS-API-008", mkNoted("GenericParam", "a", "lifetime"), mkNode("Ident", "T")},
		{"RS-API-009", mkMut("VarDecl", "x"), mkNode("VarDecl", "x")},
		{"RS-API-010", mkNode("Cast", "as"), mkNode("Cast", "x")},
		{"RS-API-011", mkNode("Ident", "dbg"), mkNode("Ident", "log")},
		{"RS-API-012", mkNode("Ident", "unreachable"), mkNode("Ident", "panic")},
	}
}

// ruleByID indexes the declared local rules by identifier.
func ruleByID() map[string]rule.Spec {
	out := make(map[string]rule.Spec, 16)
	for _, spec := range rsRules() {
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

// TestRuleSetCount proves the declared Rust rule set clears the M5 threshold.
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

// parseRS parses one in-memory Rust source buffer into an AIR unit.
func parseRS(t *testing.T, source string) *air.Unit {
	t.Helper()
	unit := air.SourceUnit{
		Language: "rs",
		Files:    []air.SourceFile{air.NewSourceFile("sample.rs", []byte(source))},
	}
	frontend, ferr := rsfront.NewRsFrontend()
	if ferr != nil {
		t.Fatalf("construct rsfront: %v", ferr)
	}
	built, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatalf("parse RS sample: %v", err)
	}
	if built == nil {
		t.Fatal("parse RS sample returned a nil unit")
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
	if !walkMatch(matcher, parseRS(t, source)) {
		t.Errorf("rule %s (%s) failed to hit real RS frontend output", id, spec.Pattern)
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
	if walkMatch(matcher, parseRS(t, source)) {
		t.Errorf("rule %s (%s) unexpectedly hit a compliant RS snippet", id, spec.Pattern)
	}
}

// TestEndToEndHits proves every Rust rule really fires against the production
// rsfront frontend on genuine source constructs.
func TestEndToEndHits(t *testing.T) {
	t.Parallel()
	assertHit(t, "RS-API-001", "fn f(x: Option<i32>) -> i32 { x.unwrap() }\n")
	assertHit(t, "RS-API-002", "fn f(x: Option<i32>) -> i32 { x.expect(\"e\") }\n")
	assertHit(t, "RS-API-003", "fn f() { panic!(\"x\"); }\n")
	assertHit(t, "RS-API-004", "fn f() -> i32 { todo!() }\n")
	assertHit(t, "RS-API-005", "fn f() -> i32 { unimplemented!() }\n")
	assertHit(t, "RS-API-006", "fn f(v: Vec<i32>) { let x = v.clone(); }\n")
	assertHit(t, "RS-API-007", "fn f() { unsafe { let p = 1; } }\n")
	assertHit(t, "RS-API-008", "fn link<'a>(s: &'a str) -> &'a str { s }\n")
	assertHit(t, "RS-API-009", "fn f() { let mut x = 0; }\n")
	assertHit(t, "RS-API-010", "fn f(x: i64) { let y = x as u32; }\n")
	assertHit(t, "RS-API-011", "fn f(x: i32) { dbg!(x); }\n")
	assertHit(t, "RS-API-012", "fn f() -> i32 { unreachable!() }\n")
}

// TestEndToEndCompliantMiss proves the same rules stay silent on clean Rust.
func TestEndToEndCompliantMiss(t *testing.T) {
	t.Parallel()
	assertMiss(t, "RS-API-001", "fn f(x: i32) -> i32 { x + 1 }\n")
	assertMiss(t, "RS-API-002", "fn f(x: i32) -> i32 { x + 1 }\n")
	assertMiss(t, "RS-API-003", "fn f() { println!(\"x\"); }\n")
	assertMiss(t, "RS-API-004", "fn f() -> i32 { 1 }\n")
	assertMiss(t, "RS-API-005", "fn f() -> i32 { 1 }\n")
	assertMiss(t, "RS-API-006", "fn f(v: Vec<i32>) { let x = v.copy(); }\n")
	assertMiss(t, "RS-API-007", "fn f() { let p = 1; }\n")
	assertMiss(t, "RS-API-008", "fn f(x: &str) -> &str { x }\n")
	assertMiss(t, "RS-API-009", "fn f() { let x = 0; }\n")
	assertMiss(t, "RS-API-010", "fn f(x: i64) { let y = x; }\n")
	assertMiss(t, "RS-API-011", "fn f(x: i32) { println!(\"{:?}\", x); }\n")
	assertMiss(t, "RS-API-012", "fn f() -> i32 { panic!(\"x\"); }\n")
}

// TestEndToEndUnwrapAndUnsafeHits proves the representative unwrap and unsafe
// rules fire against the real frontend and report source positions.
func TestEndToEndUnwrapAndUnsafeHits(t *testing.T) {
	t.Parallel()
	for _, id := range []string{"RS-API-001", "RS-API-007"} {
		spec := ruleByID()[id]
		compiled, err := rule.CompilePattern(spec.Pattern)
		if err != nil {
			t.Fatalf("compile %s: %v", id, err)
		}
		matcher := rule.NewMatcher(compiled)
		source := "// safety review\nfn f(x: Option<i32>) -> i32 {\n    let v = x.unwrap();\n    let _ = unsafe { v };\n    v\n}\n"
		unit := parseRS(t, source)
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
			t.Fatalf("rule %s did not capture any node in the real unit", id)
		}
		if captured.Line < 2 {
			t.Fatalf("rule %s captured node on line %d, want >= 2", id, captured.Line)
		}
	}
}
