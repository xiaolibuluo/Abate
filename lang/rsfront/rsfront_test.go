/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
	"abate/lang/frontkit"
)

// rsFrontend returns a fresh Rust adapter or fails the calling test.
func rsFrontend(t *testing.T) air.Frontend {
	t.Helper()
	fe, err := NewRsFrontend()
	if err != nil {
		t.Fatalf("NewRsFrontend: %v", err)
	}
	return fe
}

// rsParse reduces one in-memory Rust source buffer into a unit and its root.
func rsParse(t *testing.T, src string) (*air.Unit, air.Node, []air.Diagnostic) {
	t.Helper()
	unit := air.SourceUnit{
		Language: "rs", Dialect: "2021",
		Files: []air.SourceFile{air.NewSourceFile("main.rs", []byte(src))},
	}
	out, diags, err := rsFrontend(t).Parse(unit)
	if err != nil {
		t.Fatalf("Parse(%q): %v", src, err)
	}
	if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
		t.Fatalf("Parse(%q): missing AIR root", src)
	}
	return out, out.Roots[0], diags
}

// rsTokenize produces the exact token stream that FrontendBase.Parse delegates.
func rsTokenize(t *testing.T, src string) []frontkit.Token {
	t.Helper()
	lexer, err := frontkit.NewLexer(rsLexSpec{}, []byte(src))
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize(%q): %v", src, err)
	}
	return tokens
}

// rsNodes returns a pre-order snapshot of every reachable AIR vertex.
func rsNodes(root air.Node) []air.Node {
	out := make([]air.Node, 0, 16)
	air.Walk(root, func(node air.Node) bool {
		out = append(out, node)
		return true
	})
	return out
}

func rsHasKind(root air.Node, kind air.NodeKind) bool {
	for _, node := range rsNodes(root) {
		if node.Kind() == kind {
			return true
		}
	}
	return false
}

func rsHasNote(root air.Node, note string) bool {
	return rsHasSem(root, air.SemNote, note)
}

func rsHasSem(root air.Node, key air.SemKey, value string) bool {
	for _, node := range rsNodes(root) {
		if stored, ok := node.Sem(key).(string); ok && stored == value {
			return true
		}
	}
	return false
}

// TestRsConstructToAIRMapping proves real Rust constructions reduce to frozen
// AIR kinds and semantic notes instead of a constant empty shell.
func TestRsConstructToAIRMapping(t *testing.T) {
	if len(rsMappingCases) < 20 {
		t.Fatalf("only %d mapping cases; at least 20 required", len(rsMappingCases))
	}
	for _, tc := range rsMappingCases {
		t.Run(tc.name, func(t *testing.T) {
			out, root, _ := rsParse(t, tc.src)
			if out.NodeCount() < 2 {
				t.Fatalf("empty AIR shell for %q", tc.src)
			}
			if !rsHasKind(root, tc.kind) {
				t.Fatalf("%s: missing %s in tree for %q", tc.name, tc.kind, tc.src)
			}
			if tc.note != "" && !rsHasNote(root, tc.note) {
				t.Fatalf("%s: missing SemNote %q for %q", tc.name, tc.note, tc.src)
			}
		})
	}
}

// TestRsFrontendContract proves the returned air.Frontend produces a layered,
// valid unit with a KindPackage root and many reachable vertices.
func TestRsFrontendContract(t *testing.T) {
	fe := rsFrontend(t)
	if fe == nil {
		t.Fatal("NewRsFrontend returned a nil air.Frontend")
	}
	src := "use std::io;\n" +
		"struct Point { x: i32 }\n" +
		"trait Draw { fn draw(&self); }\n" +
		"impl Point { fn origin() -> Self { Point { x: 0 } } }\n" +
		"fn main() { let mut n = 1; println!(\"{}\", n); }\n"
	unit := air.SourceUnit{
		Language: "rs", Dialect: "rust2021",
		Files: []air.SourceFile{air.NewSourceFile("main.rs", []byte(src))},
	}
	out, diags, err := fe.Parse(unit)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if err := out.Validate(); err != nil {
		t.Fatalf("Unit.Validate: %v", err)
	}
	if out.NodeCount() < 20 {
		t.Fatalf("NodeCount=%d, want a layered tree", out.NodeCount())
	}
	if out.Roots[0].Kind() != air.KindPackage {
		t.Fatalf("root kind=%s, want KindPackage", out.Roots[0].Kind())
	}
	if len(out.Roots[0].Children()) < 4 {
		t.Fatal("root has no item children")
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got := out.Annotations["lang.rs.dialect"]; got != "2021" {
		t.Fatalf("dialect annotation=%q, want 2021", got)
	}
}

// TestRsDeterminismFingerprint certifies Unit.Fingerprint is byte-stable.
func TestRsDeterminismFingerprint(t *testing.T) {
	fe := rsFrontend(t)
	src := "use std::collections::HashMap;\n" +
		"fn count<'a>(xs: &'a [i32]) -> usize {\n" +
		"    let mut n = 0;\n" +
		"    for x in xs { n += *x as usize; }\n" +
		"    println!(\"{}\", n);\n" +
		"    n\n" +
		"}\n"
	unit := air.SourceUnit{
		Language: "rs", Dialect: "2021",
		Files: []air.SourceFile{air.NewSourceFile("main.rs", []byte(src))},
	}
	determinism.MustMatch100(t, 0x52530001, func() []byte {
		out, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		sum := out.Fingerprint()
		return sum[:]
	})
}

// TestRsDialectNormalization checks the canonical 2015-2024 edition mapping.
func TestRsDialectNormalization(t *testing.T) {
	good := map[string]string{
		"": "2021", "2015": "2015", "2018": "2018", "2021": "2021",
		"2024": "2024", "rust2021": "2021", "edition2018": "2018",
		"RUST 2015": "2015", "  RS2024 ": "2024", "stable": "2021",
		"nightly": "2021", "beta": "2021",
	}
	for in, want := range good {
		got, err := defaultRsSpec().NormalizeDialect(in)
		if err != nil || got != want {
			t.Fatalf("NormalizeDialect(%q)=%q,%v want %q", in, got, err, want)
		}
	}
	for _, in := range []string{"4.0", "ruby3", "2022", "go1.26", "python3.12"} {
		if got, err := defaultRsSpec().NormalizeDialect(in); err == nil {
			t.Fatalf("NormalizeDialect(%q)=%q, want error", in, got)
		}
	}
}

// TestRsLexSpecContract checks the frozen lexical table contract.
func TestRsLexSpecContract(t *testing.T) {
	spec := defaultRsSpec()
	if spec.Tag() != "rs" {
		t.Fatalf("Tag=%q, want rs", spec.Tag())
	}
	lex := spec.LexSpec()
	if lex == nil || lex.IndentSensitive() {
		t.Fatal("rust lexer must not be indentation sensitive")
	}
	if lex.LineComment() != "//" {
		t.Fatalf("LineComment=%q, want //", lex.LineComment())
	}
	if open, closeTok := lex.BlockComment(); open != "/*" || closeTok != "*/" {
		t.Fatalf("BlockComment=%q,%q", open, closeTok)
	}
	for _, table := range []struct {
		name   string
		values []string
		want   string
	}{
		{"keywords", lex.Keywords(), "macro_rules"},
		{"operators", lex.Operators(), "::"},
		{"operators", lex.Operators(), "'"},
		{"quotes", lex.StringQuotes(), `"`},
		{"prefixes", lex.RawStringPrefixes(), "r#"},
		{"separators", lex.NumberSeparators(), "_"},
	} {
		if !rsContains(table.values, table.want) {
			t.Fatalf("%s table missing %q", table.name, table.want)
		}
	}
	if rsContains(lex.StringQuotes(), "'") {
		t.Fatal("single quote must not be a string delimiter")
	}
}

// TestRsSemanticKeys checks non-note AIR annotations on proprietary forms.
func TestRsSemanticKeys(t *testing.T) {
	cases := []struct {
		name, src string
		key       air.SemKey
		value     string
	}{
		{"async-fn", "async fn f() {}\n", air.SemAsync, "true"},
		{"mutability", "let mut count = 0;\n", air.SemMutability, "mutable"},
		{"binary-op", "let x = 1 + 2;\n", air.SemOperator, "+"},
		{"assign-op", "let x = 1;\n", air.SemOperator, "="},
		{"type-param", "type A = Vec<i32>;\n", air.SemTypeParam, "1"},
		{"directive", "// clippy: allow\nfn f() {}\n", air.SemDirective,
			"// clippy: allow"},
		{"tag", "#[derive(Clone)]\nstruct S;\n", air.SemTag, "derive"},
		{"known-macro", "vec![1];\n", air.SemTag, "known-macro"},
		{"scope", "fn f() {}\n", air.SemScopeID, "crate"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, root, _ := rsParse(t, tc.src)
			if !rsHasSem(root, tc.key, tc.value) {
				t.Fatalf("missing %s=%q for %q", tc.key, tc.value, tc.src)
			}
		})
	}
}

// TestRsParseTokensDirect proves the LanguageSpec reduction is a real parser
// when called without the FrontendBase skeleton.
func TestRsParseTokensDirect(t *testing.T) {
	src := "fn f(x: i32) -> i32 {\n    match x { 0 => 1, _ => x + 1 }\n}\n"
	tokens := rsTokenize(t, src)
	root, diags, err := defaultRsSpec().ParseTokens(tokens, "direct.rs")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if root == nil || root.Kind() != air.KindPackage {
		t.Fatal("ParseTokens did not produce a KindPackage root")
	}
	if !rsHasKind(root, air.KindFuncDecl) || !rsHasKind(root, air.KindMatch) ||
		!rsHasKind(root, air.KindBlock) {
		t.Fatal("ParseTokens produced an empty shell")
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if err := air.CheckIntervals(root, len(src)); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
}
