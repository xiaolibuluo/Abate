/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// rsDedicatedCases pairs a minimal Rust source with the dedicated kind the
// parser must produce while genuinely reducing that construct.
var rsDedicatedCases = []struct {
	construct string
	src       string
	kind      air.NodeKind
}{
	{"async-block", "fn main() { let b = async { 1 }; }\n", air.KindRsAsyncBlock},
	{"const-block", "fn main() { let c = const { 2 }; }\n", air.KindRsConstBlock},
	{"gen-block", "fn main() { let e = gen { yield 1 }; }\n", air.KindRsGenBlock},
	{"try-block", "fn main() { let t = try { 5 }; }\n", air.KindRsTryBlock},
	{"unsafe-block", "fn main() { let b = unsafe { 1 }; }\n", air.KindRsUnsafeBlock},
	{"compound-assign", "fn main() { z += 2; }\n", air.KindRsCompoundAssign},
	{"field-identifier",
		"fn main() { let s = S { x: 1, ..d }; }\n", air.KindRsFieldIdent},
	{"field-initializer",
		"fn main() { let s = S { x: 1, ..d }; }\n", air.KindRsFieldInit},
	{"base-field",
		"fn main() { let s = S { x: 1, ..Default::default() }; }\n",
		air.KindRsBaseField},
	{"shorthand-field", "fn main() { let s = P { q, r }; }\n",
		air.KindRsShorthandField},
	{"shorthand-ident",
		"fn main() { match v { S { a, b } => 1, _ => 0 } }\n",
		air.KindRsShorthandIdent},
	{"metavariable", "macro_rules! m { ($x:expr) => { $x } }\n",
		air.KindRsMetavar},
	{"self-expr", "fn main() { let y = self.x; }\n", air.KindRsSelf},
	{"crate-path", "extern crate serde;\n", air.KindRsCrate},
	{"super-expr", "fn main() { let y = super.y; }\n", air.KindRsSuper},
	{"empty-stmt", "fn main() { ; }\n", air.KindRsEmptyStmt},
	{"loop", "fn main() { loop { break; } }\n", air.KindRsLoop},
	{"impl", "struct S;\nimpl S { fn new() -> S { S } }\n", air.KindRsImpl},
	{"string-part", "fn main() { let s = \"a\\nb\"; }\n", air.KindRsStringFrag},
	{"extern-crate", "extern crate serde;\n", air.KindRsExternCrate},
	{"scoped-use", "use std::collections::{HashMap as M, *};\n",
		air.KindRsScopedUse},
	{"use-list", "use std::collections::{HashMap as M, *};\n",
		air.KindRsUseList},
	{"use-wildcard", "use std::collections::{HashMap as M, *};\n",
		air.KindRsUseWildcard},
	{"use-as", "use std::collections::{HashMap as M, *};\n", air.KindRsUseAs},
	{"negative-literal", "fn main() { let n = -42; }\n", air.KindRsNegativeLit},
	{"yield-expr", "fn main() { let e = gen { yield 1 }; }\n", air.KindRsYield},
	{"macro-call", "fn main() { let _ = m!(1); }\n", air.KindRsMacroCall},
	{"macro-def", "macro_rules! m { ($x:expr) => { $x } }\n", air.KindRsMacroDef},
	{"func-signature", "trait T { fn sig(&self) -> i32; }\n", air.KindRsFuncSig},
	{"generic-func", "fn g<T>(x: T) -> T { x }\n", air.KindRsGenericFunc},
	{"turbofish", "fn main() { let f = Vec::<i32>::new(); }\n",
		air.KindRsTurbofish},
	{"type-args", "fn main() { let f = Vec::<i32>::new(); }\n",
		air.KindRsTypeArgs},
	{"enum-variant", "enum E { A, B(u32), C { x: i32 } }\n",
		air.KindRsEnumVariant},
	{"visibility-modifier", "pub fn open() {}\n", air.KindRsVisibility},
	{"mutable-specifier", "fn main() { let mut n = 0; n += 1; }\n",
		air.KindRsMutable},
}

// TestRsDedicatedKindsAreReallyProduced parses each minimal source and asserts
// the parser genuinely emitted the dedicated kind for that construct.
func TestRsDedicatedKindsAreReallyProduced(t *testing.T) {
	for _, tc := range rsDedicatedCases {
		unit, _, _ := rsParse(t, tc.src)
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

// TestRsQASampledRulesHaveGenuineProduction is the strengthened form of the
// earlier QA-absent pin (see pyfront for the pattern): every promoted rule
// must be exact, must name a genuine dedicated kind (>= 128, outside the
// fabricated 87-108 bucket), and a minimal source must really produce it.
func TestRsQASampledRulesHaveGenuineProduction(t *testing.T) {
	wanted := map[string]string{
		// async 块在全语料零产出，已回退 degraded；改用真实产出且仍为
		// exact 的 unsafe 块做 dedicated 生产锚点。
		"rs.unsafe-block":     "unsafe-block",
		"rs.yield":            "yield-expr",
		"rs.macro-def":        "macro-def",
		"rs.negative-literal": "negative-literal",
	}
	byName := make(map[string]grammar.MappingRule)
	for _, rule := range grammar.DefaultRules("rs") {
		byName[rule.Name] = rule
	}
	produced := make(map[string]bool)
	for _, tc := range rsDedicatedCases {
		unit, _, _ := rsParse(t, tc.src)
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
			t.Errorf("rule %q not found in rs rule table", name)
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
