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

// rsChaosInputs covers damaged files, unclosed macros, nested generics,
// unmatched delimiters, illegal UTF-8, excessive nesting and empty buffers.
var rsChaosInputs = []string{
	"",
	"\n",
	"   \n\n",
	"\t\n",
	"fn f(:\n    ???\n",
	"macro_rules! broken { ( => { ;\n",
	"println!(\"unclosed\n",
	"let v = vec![1, 2,\n",
	"fn g<T<<<<<U>>>>>(x: T) {}\n",
	"impl Trait for Vec<Vec<Vec<Vec<i32>>>> {}\n",
	"match v { Some(x) =>\n",
	"#[derive(Clone\nstruct S;\n",
	"#![allow(\n",
	"unsafe { let p = ; }\n",
	"let s = r#\"unterminated\n",
	"let c = 'x\n",
	"fn h<'a>(x: &'a str) -> &'a str { x }\n\xff\n",
	"async fn z() { fut.await\n",
	"for x in .. { }\n",
	"x = ((((((\n",
	"a::b::c::d::\n",
	"|x| |y| |z| x + y + z,\n",
	"? ! @ $ ~ \\\n",
	strings.Repeat("(", 200) + strings.Repeat(")", 200),
	strings.Repeat("fn f() { ", 120) + "}" + strings.Repeat("}", 120),
	strings.Repeat("let x = 1;\n", 400),
	strings.Repeat("[", 300) + "1" + strings.Repeat("]", 300) + "\n",
	"\xff\xfe\x00\x01\n",
}

// TestRsChaosNeverPanics feeds hostile input through the whole adapter. The
// tested code runs directly: a panic would bubble up and fail the test through
// the Go testing framework, so no explicit recover wrapper is used. Every
// accepted parse must still yield a non-nil unit, a positive node count and a
// stable fingerprint.
func TestRsChaosNeverPanics(t *testing.T) {
	fe := rsFrontend(t)
	for i, src := range rsChaosInputs {
		unit := air.SourceUnit{
			Language: "rs", Dialect: "2021",
			Files: []air.SourceFile{air.NewSourceFile("main.rs", []byte(src))},
		}
		out, diags, err := fe.Parse(unit)
		if err != nil {
			t.Logf("chaos[%d] classified error: %v", i, err)
			continue
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("chaos[%d] %q: missing root",
				i, src[:rsMinInt(len(src), 40)])
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("chaos[%d] %q: invalid unit: %v", i,
				src[:rsMinInt(len(src), 40)], err)
		}
		if out.NodeCount() < 1 {
			t.Fatalf("chaos[%d]: empty tree", i)
		}
		if out.Roots[0].Kind() != air.KindPackage {
			t.Fatalf("chaos[%d]: root is not KindPackage", i)
		}
		first := out.Fingerprint()
		second := out.Fingerprint()
		if first != second {
			t.Fatalf("chaos[%d]: fingerprint unstable", i)
		}
		if err := air.CheckIntervals(out.Roots[0], len(src)); err != nil {
			t.Fatalf("chaos[%d] %q: %v", i,
				src[:rsMinInt(len(src), 40)], err)
		}
		_ = diags
	}
}

// TestRsChaosDeterministicAcrossRuns re-parses every damaged buffer and checks
// the resulting unit hash and node count are identical between runs.
func TestRsChaosDeterministicAcrossRuns(t *testing.T) {
	fe := rsFrontend(t)
	for i, src := range rsChaosInputs {
		unit := air.SourceUnit{
			Language: "rs", Dialect: "2021",
			Files: []air.SourceFile{air.NewSourceFile("main.rs", []byte(src))},
		}
		a, _, errA := fe.Parse(unit)
		b, _, errB := fe.Parse(unit)
		if (errA == nil) != (errB == nil) {
			t.Fatalf("chaos[%d]: inconsistent error classification", i)
		}
		if a == nil || b == nil {
			continue
		}
		if a.Fingerprint() != b.Fingerprint() {
			t.Fatalf("chaos[%d]: non-deterministic damaged parse", i)
		}
		if a.NodeCount() != b.NodeCount() {
			t.Fatalf("chaos[%d]: non-deterministic node count", i)
		}
	}
}

// TestRsChaosDeepNestingStaysBounded proves the depth ceiling caps recursion
// on adversarial nesting without losing the package root.
func TestRsChaosDeepNestingStaysBounded(t *testing.T) {
	src := strings.Repeat("(", 300) + "1" + strings.Repeat(")", 300) + "\n"
	out, root, _ := rsParse(t, src)
	if root.Kind() != air.KindPackage {
		t.Fatal("deep nesting lost the root")
	}
	if out.NodeCount() < 2 || out.NodeCount() > 10000 {
		t.Fatalf("unbounded node count %d", out.NodeCount())
	}
	if err := out.Validate(); err != nil {
		t.Fatalf("deep nesting produced invalid unit: %v", err)
	}
	if !rsHasSem(root, air.SemPartial, "true") {
		t.Fatal("depth-limited reduction is not marked SemPartial")
	}
}

// TestRsChaosIllegalByteDiagnostic checks the deterministic recovery warning.
func TestRsChaosIllegalByteDiagnostic(t *testing.T) {
	src := "let x = 1 \xff\n"
	_, _, diags := rsParse(t, src)
	found := false
	for _, d := range diags {
		if d.Severity == air.SeverityWarning {
			found = true
			if d.Path != "main.rs" || d.Code == "" {
				t.Fatalf("malformed diagnostic: %+v", d)
			}
		}
	}
	if !found {
		t.Fatal("illegal byte did not produce a severity warning")
	}
}

// rsMinInt returns the smaller of two integers without importing math.
func rsMinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
