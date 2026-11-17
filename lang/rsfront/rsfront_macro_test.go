/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"strings"
	"testing"

	"abate/lang/air"
)

// TestRsMacroRulesDefinition proves macro_rules lowers to KindCall with a
// macro note and keeps its braced body as a real child block.
func TestRsMacroRulesDefinition(t *testing.T) {
	src := "macro_rules! add {\n    ($a:expr, $b:expr) => { $a + $b };\n}\n"
	out, root, diags := rsParse(t, src)
	if !rsHasKind(root, air.KindCall) || !rsHasNote(root, "macro") {
		t.Fatal("macro_rules! did not lower to KindCall with a macro note")
	}
	if !rsHasKind(root, air.KindBlock) {
		t.Fatal("macro_rules! body was not reduced to a block")
	}
	if out.NodeCount() < 4 {
		t.Fatalf("macro definition produced only %d nodes", out.NodeCount())
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
}

// TestRsMacroInvocations checks built-in and user macros, including SemTag.
func TestRsMacroInvocations(t *testing.T) {
	cases := []struct {
		name, src string
		known     bool
	}{
		{"println", "println!(\"hi {}\", x);\n", true},
		{"vec", "let v = vec![1, 2, 3];\n", true},
		{"format", "let s = format!(\"{}\", 1);\n", true},
		{"user-macro", "my_macro!(a, b);\n", false},
		{"assert-eq", "assert_eq!(a, b);\n", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, root, _ := rsParse(t, tc.src)
			if !rsHasKind(root, air.KindCall) || !rsHasNote(root, "macro") {
				t.Fatalf("missing macro call for %q", tc.src)
			}
			got := rsHasSem(root, air.SemTag, "known-macro")
			if got != tc.known {
				t.Fatalf("known-macro SemTag=%v, want %v for %q",
					got, tc.known, tc.src)
			}
		})
	}
}

// TestRsAttributeForms checks outer, inner and cfg attributes.
func TestRsAttributeForms(t *testing.T) {
	cases := []struct {
		name, src, tag string
		inner          bool
	}{
		{"derive", "#[derive(Clone)]\nstruct S;\n", "derive", false},
		{"cfg", "#[cfg(feature = \"x\")]\nfn f() {}\n", "cfg", false},
		{"inner", "#![allow(dead_code)]\nfn f() {}\n", "allow", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, root, _ := rsParse(t, tc.src)
			if !rsHasKind(root, air.KindAttribute) {
				t.Fatalf("missing KindAttribute for %q", tc.src)
			}
			if !rsHasSem(root, air.SemTag, tc.tag) {
				t.Fatalf("missing SemTag=%q for %q", tc.tag, tc.src)
			}
			got := rsHasNote(root, "inner-attribute")
			if got != tc.inner {
				t.Fatalf("inner-attribute note=%v, want %v", got, tc.inner)
			}
		})
	}
}

// TestRsLifetimeAndReference checks lifetime parameters and reference forms.
func TestRsLifetimeAndReference(t *testing.T) {
	src := "fn take<'a>(x: &'a str, y: &mut i32) -> &'a str { x }\n"
	_, root, _ := rsParse(t, src)
	if !rsHasKind(root, air.KindGenericParam) ||
		!rsHasNote(root, "lifetime") {
		t.Fatal("missing KindGenericParam with a lifetime note")
	}
	if !rsHasKind(root, air.KindRef) || !rsHasNote(root, "ref") {
		t.Fatal("missing KindRef with a ref note")
	}
	if !rsHasSem(root, air.SemMutability, "mutable") {
		t.Fatal("missing mutable reference annotation")
	}
	if !rsHasKind(root, air.KindGenericInst) {
		t.Fatal("missing generic instantiation for the lifetime parameter list")
	}
}

// TestRsRawAndByteStringLiterals checks Rust literal prefixes stay lexical.
func TestRsRawAndByteStringLiterals(t *testing.T) {
	src := "let a = r\"raw\\ntext\";\n" +
		"let b = b\"bytes\";\n" +
		"let c = c\"cstr\";\n" +
		"let d = r#\"hash raw\"#;\n"
	_, root, _ := rsParse(t, src)
	if !rsHasKind(root, air.KindStrLit) {
		t.Fatal("raw and byte strings did not project to KindStrLit")
	}
	count := 0
	for _, node := range rsNodes(root) {
		if node.Kind() == air.KindStrLit {
			count++
		}
	}
	if count < 4 {
		t.Fatalf("only %d string literal nodes for four prefixed literals", count)
	}
}

// TestRsUnsafeAndAsyncForms checks unsafe and async lowering variants.
func TestRsUnsafeAndAsyncForms(t *testing.T) {
	cases := []struct {
		name, src string
		kind      air.NodeKind
		note      string
	}{
		{"unsafe-block", "unsafe { let p = 1; }\n", air.KindBlock, "unsafe"},
		{"unsafe-fn", "unsafe fn danger() {}\n", air.KindFuncDecl, "unsafe"},
		{"async-fn", "async fn load() {}\n", air.KindFuncDecl, ""},
		{"extern-fn", "extern \"C\" fn abi() {}\n", air.KindFuncDecl, "extern-abi"},
		{"const-fn", "const fn fixed() -> i32 { 1 }\n",
			air.KindFuncDecl, "const-fn"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, root, _ := rsParse(t, tc.src)
			if !rsHasKind(root, tc.kind) {
				t.Fatalf("missing %s for %q", tc.kind, tc.src)
			}
			if tc.note != "" && !rsHasNote(root, tc.note) {
				t.Fatalf("missing SemNote %q for %q", tc.note, tc.src)
			}
		})
	}
}

// TestRsDocAndDirectiveComments checks comment classification.
func TestRsDocAndDirectiveComments(t *testing.T) {
	src := "/// public doc\n// clippy: allow\n// plain\n/* block */\nfn f() {}\n"
	_, root, _ := rsParse(t, src)
	for _, kind := range []air.NodeKind{
		air.KindDoc, air.KindDirective, air.KindComment,
	} {
		if !rsHasKind(root, kind) {
			t.Fatalf("missing %s node", kind)
		}
	}
	if !rsHasNote(root, "doc-comment") {
		t.Fatal("missing doc-comment note")
	}
	if !rsHasSem(root, air.SemDirective, "// clippy: allow") {
		t.Fatal("missing directive semantic annotation")
	}
}

// TestRsMatchPatternArms checks pattern arms, wildcards and guards.
func TestRsMatchPatternArms(t *testing.T) {
	src := "match value {\n    Some(x) if x > 0 => x,\n    None => 0,\n    _ => -1,\n}\n"
	_, root, _ := rsParse(t, src)
	if !rsHasKind(root, air.KindMatch) {
		t.Fatal("missing KindMatch")
	}
	if !rsHasNote(root, "match-arm") {
		t.Fatal("missing match-arm note")
	}
	count := 0
	for _, node := range rsNodes(root) {
		if node.Kind() == air.KindCase {
			count++
		}
	}
	if count < 2 {
		t.Fatalf("only %d match arms reduced to KindCase", count)
	}
}

// rsBranchSources drives broad coverage of mapper and expression branches.
func rsBranchSources() []string {
	return []string{
		"pub(crate) fn visible() {}\n",
		"extern crate alloc;\n",
		"default fn d() {}\n",
		"dyn Trait,\n",
		"where T: Clone,\n",
		"box value,\n",
		"yield 1,\n",
		"try { work() },\n",
		"union U { a: i32 }\n",
		"let a = !flag;\n",
		"let b = -1;\n",
		"let c = &&x;\n",
		"let d = ..end;\n",
		"let e = start..=end;\n",
		"let f = || 1;\n",
		"let g = (x);\n",
		"let h = { x + 1 };\n",
		"let i = arr[1..3];\n",
		"let j = a.b.c;\n",
		"let k = a::b::c;\n",
		"let l = x += 2;\n",
		"let m = &raw const x;\n",
		"let n = '\\n';\n",
		"if flag { } else { }\n",
		"while let Some(v) = next() { }\n",
		"for i in 0..n { continue; }\n",
		"loop { break; }\n",
		"pub struct S;\n",
		"type Pair = (i32, i32);\n",
		"static N: usize = 1_000;\n",
		"const C: bool = false;\n",
		"let t = (1, 2, 3);\n",
		"let v = [1, 2];\n",
		"let w = HashMap<String, Vec<i32>>::new();\n",
	}
}

// TestRsBroadConstructCorpus parses diverse constructs to exercise every
// deterministic branch of the mapper, parser and expression reducer.
func TestRsBroadConstructCorpus(t *testing.T) {
	for i, src := range rsBranchSources() {
		out, root, _ := rsParse(t, src)
		if err := out.Validate(); err != nil {
			t.Fatalf("branch[%d] %q invalid unit: %v", i, src, err)
		}
		if out.NodeCount() < 2 {
			t.Fatalf("branch[%d] %q produced an empty shell", i, src)
		}
		if err := air.CheckIntervals(root, len(src)); err != nil {
			t.Fatalf("branch[%d] %q: %v", i, src, err)
		}
		if strings.TrimSpace(src) == "" {
			t.Fatal("empty branch source")
		}
	}
}
