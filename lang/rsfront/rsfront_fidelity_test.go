/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package rsfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
	"testing"
)

type rsMapCase struct {
	name string
	src  string
	kind air.NodeKind
	note string
}

var rsMappingCases = []rsMapCase{
	{"use-import", "use std::io;\n", air.KindImport, ""},
	{"module", "mod inner { }\n", air.KindModuleDecl, ""},
	{"function", "fn add(a: i32, b: i32) -> i32 { a + b }\n", air.KindFuncDecl, ""},
	{"struct", "struct Point { x: i32 }\n", air.KindStructDecl, ""},
	{"enum", "enum Color { Red, Green }\n", air.KindEnumDecl, ""},
	{"trait", "trait Draw { fn draw(&self); }\n", air.KindTraitDecl, ""},
	{"impl", "impl Point { fn origin() -> Self { Point { x: 0 } } }\n", air.KindBlock, "impl"},
	{"method", "impl Point { fn origin() -> Self { Point { x: 0 } } }\n", air.KindMethodDecl, ""},
	{"macro-rules", "macro_rules! twice { ($x:expr) => { $x + $x }; }\n", air.KindCall, "macro"},
	{"macro-invocation", "println!(\"hi {}\", x);\n", air.KindCall, "macro"},
	{"attribute", "#[derive(Clone)]\nstruct S;\n", air.KindAttribute, ""},
	{"unsafe-block", "unsafe { let p = 1; }\n", air.KindBlock, "unsafe"},
	{"match", "match v { Some(x) => x, None => 0 }\n", air.KindMatch, ""},
	{"match-arm", "match v { Some(x) => x, None => 0 }\n", air.KindCase, "match-arm"},
	{"if-let", "if let Some(v) = opt { v } else { 0 }\n", air.KindIf, "if-let"},
	{"for-range", "for x in start..end { }\n", air.KindRangeFor, ""},
	{"range", "let r = start..end;\n", air.KindBinary, "range"},
	{"while", "while n < 3 { n += 1; }\n", air.KindWhile, ""},
	{"loop", "loop { break; }\n", air.KindWhile, "loop"},
	{"let-mut", "let mut count = 0;\n", air.KindVarDecl, ""},
	{"const", "const MAX: usize = 10;\n", air.KindConstDecl, ""},
	{"static", "static FLAG: bool = true;\n", air.KindVarDecl, "static"},
	{"type-alias", "type Alias = Vec<i32>;\n", air.KindTypeDecl, ""},
	{"generic-inst", "type Alias = Vec<i32>;\n", air.KindGenericInst, ""},
	{"lifetime", "fn link<'a>(s: &'a str) -> &'a str { s }\n", air.KindGenericParam, "lifetime"},
	{"closure", "let f = |x: i32| x * 2;\n", air.KindLambda, "closure"},
	{"move-closure", "let g = move |x| x;\n", air.KindLambda, "move"},
	{"try-op", "let v = foo()?;\n", air.KindUnary, "try-op"},
	{"reference", "let r = &value;\n", air.KindRef, "ref"},
	{"deref", "let v = *ptr;\n", air.KindStar, ""},
	{"cast", "let y = x as u64;\n", air.KindCast, ""},
	{"selector", "let p = std::mem::size_of::<i32>();\n", air.KindSelector, ""},
	{"await", "let z = fut.await;\n", air.KindAwait, ""},
	{"return", "fn f() -> i32 { return 1; }\n", air.KindReturn, ""},
	{"break", "loop { break; }\n", air.KindBranch, "break"},
	{"array", "let a = [1, 2, 3];\n", air.KindCompositeLit, "array"},
	{"tuple", "let t = (1, 2);\n", air.KindTuple, ""},
	{"struct-literal", "Point { x: 1, y: 2 };\n", air.KindCompositeLit, "struct-literal"},
	{"char-literal", "let c = 'x';\n", air.KindCharLit, "char-literal"},
	{"bool-lit", "let b = true;\n", air.KindBoolLit, ""},
	{"str-lit", "let s = \"hello\";\n", air.KindStrLit, ""},
	{"num-lit", "let n = 42;\n", air.KindNumLit, ""},
	{"doc-comment", "/// doc\nfn f() {}\n", air.KindDoc, ""},
	{"directive", "// clippy: allow\nfn f() {}\n", air.KindDirective, ""},
	{"comment", "// plain\nfn f() {}\n", air.KindComment, ""},
	{"call", "let x = foo(1);\n", air.KindCall, ""},
	{"index", "let x = arr[0];\n", air.KindIndex, ""},
	{"field", "let x = obj.field;\n", air.KindField, ""},
	{"binary", "let x = 1 + 2 * 3;\n", air.KindBinary, ""},
	{"assign", "let x = 1;\n", air.KindAssign, ""},
	{"block", "fn f() { }\n", air.KindBlock, ""},
	{"if", "if flag { work(); }\n", air.KindIf, ""},
	{"async-fn", "async fn fetch() -> i32 { 1 }\n", air.KindFuncDecl, ""},
	{"union", "union U { a: i32, b: u32 }\n", air.KindStructDecl, "union"},
	{"extern-crate", "extern crate core;\n", air.KindImport, "extern"},
	{"where", "where T: Clone,\n", air.KindGenericParam, "where"},
}

func rsCorpus() []struct {
	name, dialect, source string
} {
	return []struct {
		name, dialect, source string
	}{
		{"edition2015", "2015", "fn main() { let x = 1; }\n"},
		{"edition2018", "2018", "async fn f() { let y = g().await; }\n"},
		{"edition2021", "2021", rsSampleSource()},
		{"edition2024", "2024", "use a::b;\nmod m { pub fn h() {} }\n"},
		{"attributes", "2021", "#![allow(dead_code)]\n#[derive(Clone)]\nstruct S;\n"},
		{"macros", "2021", "macro_rules! m { () => {}; }\nprintln!(\"x\");\nvec![1, 2];\n"},
		{"literals", "2021", "let s = r\"raw\";\nlet b = b\"bytes\";\nlet c = 'x';\nlet n = 1_000;\n"},
		{"comments", "2021", "// plain\n/// doc\n/* block */\nfn f() {}\n"},
		{"generics", "2021", "fn f<'a, T: Clone>(x: &'a Vec<T>) -> T { x[0].clone() }\n"},
		{"empty", "2021", ""},
		{"only-newlines", "2021", "\n\n\n"},
	}
}

func rsSampleSource() string {
	return "use std::collections::HashMap;\n" +
		"/// A point.\n" +
		"#[derive(Clone)]\n" +
		"pub struct Point { pub x: i32, pub y: i32 }\n" +
		"enum Shape { Circle(f64), Square }\n" +
		"trait Area { fn area(&self) -> f64; }\n" +
		"impl Area for Point {\n" +
		"    fn area(&self) -> f64 { (self.x * self.y) as f64 }\n" +
		"}\n" +
		"fn total<'a>(xs: &'a [i32]) -> i32 {\n" +
		"    let mut sum = 0;\n" +
		"    for x in xs {\n" +
		"        sum += *x;\n" +
		"    }\n" +
		"    match sum { 0 => 1, n => n }\n" +
		"}\n" +
		"fn main() {\n" +
		"    let mut m: HashMap<String, Vec<i32>> = HashMap::new();\n" +
		"    m.insert(\"a\".to_string(), vec![1, 2]);\n" +
		"    let p = Point { x: 1, y: 2 };\n" +
		"    let f = |v: i32| v + 1;\n" +
		"    let ok = f(p.x)?;\n" +
		"    unsafe { let raw = &p.x; let _ = *raw; }\n" +
		"    println!(\"{:?} {}\", m, ok);\n" +
		"}\n"
}

func rsAssertIntervalCoverage(t *testing.T, root air.Node,
	tokens []frontkit.Token, size int) {
	t.Helper()
	if root == nil {
		t.Fatal("nil root")
	}
	span := root.Pos()
	if span.Offset != 0 || span.EndOffset != size {
		t.Fatalf("root span [%d,%d) does not cover [0,%d)",
			span.Offset, span.EndOffset, size)
	}
	if err := air.CheckIntervals(root, size); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
	nodes := make([]air.Node, 0, 32)
	air.Walk(root, func(node air.Node) bool {
		if !node.Pos().Valid() || !node.End().Valid() {
			t.Fatalf("node %s has an invalid position", node.Kind())
		}
		nodes = append(nodes, node)
		return true
	})
	inner := nodes[1:]
	for _, token := range tokens {
		switch token.Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF,
			frontkit.TokenIndent, frontkit.TokenDedent:
			continue
		}
		covered := false
		for _, node := range inner {
			start, end := air.Span(node)
			if start <= token.Offset && token.EndOffset <= end {
				covered = true
				break
			}
		}
		if !covered {
			t.Fatalf("token %q %s at [%d,%d) has no covering AIR node",
				token.Text, token.Kind, token.Offset, token.EndOffset)
		}
	}
}

func rsParseFiles(t *testing.T, files []air.SourceFile,
	dialect string) (*air.Unit, []air.Diagnostic) {
	t.Helper()
	out, diags, err := rsFrontend(t).Parse(air.SourceUnit{
		Language: "rs", Dialect: dialect, Files: files,
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if out == nil {
		t.Fatal("Parse returned a nil unit")
	}
	return out, diags
}

func TestRsIntervalCoverageCorpus(t *testing.T) {
	for _, item := range rsCorpus() {
		t.Run(item.name, func(t *testing.T) {
			tokens := rsTokenize(t, item.source)
			unit, _ := rsParseFiles(t, []air.SourceFile{
				air.NewSourceFile("sample.rs", []byte(item.source)),
			}, item.dialect)
			if len(unit.Roots) != 1 {
				t.Fatalf("expected one root, got %d", len(unit.Roots))
			}
			rsAssertIntervalCoverage(t, unit.Roots[0], tokens, len(item.source))
		})
	}
}

func TestRsStructureSpansMultipleTokens(t *testing.T) {
	source := rsSampleSource()
	tokens := rsTokenize(t, source)
	unit, _ := rsParseFiles(t, []air.SourceFile{
		air.NewSourceFile("sample.rs", []byte(source)),
	}, "2021")
	multi := false
	for _, node := range air.Collect(unit) {
		start, end := air.Span(node)
		contained := 0
		for _, token := range tokens {
			if rsIsTrivia(token.Kind) || token.Kind == frontkit.TokenEOF {
				continue
			}
			if start <= token.Offset && token.EndOffset <= end {
				contained++
			}
		}
		if contained >= 6 {
			multi = true
			break
		}
	}
	if !multi {
		t.Fatal("no structural node spans multiple tokens; reduction looks flat")
	}
}

func TestRsFingerprintStableAcrossParses(t *testing.T) {
	source := rsSampleSource()
	first := rsParseOne(t, source)
	second := rsParseOne(t, source)
	if first.Fingerprint() != second.Fingerprint() {
		t.Fatal("identical input produced different fingerprints")
	}
	if err := air.CompareUnits(first, second); err != nil {
		t.Fatalf("CompareUnits: %v", err)
	}
	if err := air.CheckIntervals(second.Roots[0], len(source)); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
}

func rsParseOne(t *testing.T, source string) *air.Unit {
	t.Helper()
	unit, _ := rsParseFiles(t, []air.SourceFile{
		air.NewSourceFile("sample.rs", []byte(source)),
	}, "2021")
	return unit
}

func TestRsNodeIntervals(t *testing.T) {
	source := rsSampleSource()
	unit := rsParseOne(t, source)
	for _, node := range air.Collect(unit) {
		pos := node.Pos()
		if !pos.Valid() || pos.Offset < 0 || pos.EndOffset > len(source) {
			t.Fatalf("%s has out-of-range interval %s", node.Kind(), pos)
		}
		if node.End().Offset < pos.Offset {
			t.Fatalf("%s end precedes start", node.Kind())
		}
	}
}

func TestRsFileOrderIndependence(t *testing.T) {
	files := []air.SourceFile{
		air.NewSourceFile("z.rs", []byte("fn z() -> i32 { 1 }\n")), air.NewSourceFile("a.rs", []byte("fn a() -> i32 { 2 }\n")),
	}
	first, _ := rsParseFiles(t, files, "2021")
	second, _ := rsParseFiles(t, []air.SourceFile{files[1], files[0]}, "2021")
	if first.Fingerprint() != second.Fingerprint() {
		t.Fatal("file order leaked into the fingerprint")
	}
}
func TestRsOperatorFidelity(t *testing.T) {
	cases := []struct{ name, base, mutated string }{
		{"cmp", "fn f(a: i32, b: i32) -> bool { a == b }\n", "fn f(a: i32, b: i32) -> bool { a != b }\n"},
		{"op", "fn f(a: i32, b: i32) -> i32 { a + b }\n", "fn f(a: i32, b: i32) -> i32 { a - b }\n"},
		{"digit", "fn f() -> i32 { 41 }\n", "fn f() -> i32 { 42 }\n"},
	}
	rootID := func(src string) air.NodeID {
		unit := rsParseOne(t, src)
		return air.IDOf(unit.Roots[0])
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if rootID(tc.base) == rootID(tc.mutated) {
				t.Fatal("operator or digit mutation did not change AIR-ID")
			}
			baseComment, mutatedComment := tc.base+"// tail\n", tc.mutated+"// tail\n"
			if rootID(baseComment) == rootID(mutatedComment) {
				t.Fatal("identical comments changed the capture conclusion")
			}
		})
	}
}
