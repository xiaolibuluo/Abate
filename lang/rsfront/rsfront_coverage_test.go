/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"flag"
	"path/filepath"
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

// coverageUpdate 报告是否重写覆盖清单，用 go test -update 打开。
var coverageUpdate = flag.Bool("update", false,
	"rewrite the rust frontend coverage checklist")

// rsCoverageCorpus 是真实 Rust 语料：既有整文件样例，也有构造片段。
func rsCoverageCorpus() []string {
	out := make([]string, 0, 32)
	out = append(out, rsSampleSource())
	return append(out,
		"use std::fmt::{self, Display};\nuse crate::module::*;\npub use a::b;\n",
		"pub struct S { pub a: i32, b: String }\nstruct T(i32, String);\nunion U { a: i32, b: f32 }\n",
		"enum E { A, B(i32), C { x: i32 } }\n",
		"trait Tr { fn m(&self) -> i32; type Out; }\nimpl Tr for S { fn m(&self) -> i32 { 0 } type Out = i32; }\n",
		"fn f<const N: usize>(x: [i32; N]) -> i32 { x[0] }\n",
		"fn g<T: Clone + Send>(v: Vec<T>) -> Option<T> { v.first().cloned() }\n",
		"fn h() { let v: Vec<i32> = vec![1, 2]; let s = &v[0..1]; let r = 1..2; }\n",
		"fn i() { if let Some(x) = opt { } else if let Err(e) = res { } else { } }\n",
		"fn j() { match v { 1 => a, 2..=3 => b, _ if c => d, _ => e } }\n",
		"fn k() { loop { break; } while let Some(x) = it.next() { continue; } for x in 0..3 {} }\n",
		"fn l() { let c = |x: i32| -> i32 { x + 1 }; let d = move || 2; }\n",
		"fn m() { unsafe { let r = &raw; } let a = *p; let b = &mut q; }\n",
		"fn n() -> Result<i32, String> { Err(\"x\".to_string())?; Ok(1) }\n",
		"async fn o() { let v = fut.await; }\n",
		"fn p() { let t = (1, 2); let s = S { a: 1, b: \"x\".into() }; let u = s.a; }\n",
		"fn q() { let v: Box<dyn Tr> = Box::new(S {}); let w = v as i32; }\n",
		"fn r() { println!(\"{}\", 1); let v = vec![1]; assert_eq!(v.len(), 1); }\n",
		"fn s() { let a = 1i64; let b = 1.5f32; let c = 'x'; let d = true; let e = \"s\"; }\n",
		"#[derive(Debug, Clone)]\n#[allow(dead_code)]\nstruct D;\n",
		"//! inner doc\n/// outer doc\nfn t() {}\n",
		"fn f() -> i32 { 1 }\n",
		"fn g(x: &str) {}\n",
		"fn h(x: *const u8) {}\n",
		"fn i(x: [u8; 4]) {}\n",
		"fn j(f: fn(i32) -> i32) {}\n",
		"fn k() { let v: Vec<i32> = Vec::new(); }\n",
	)
}

// TestRsConstructCoverage 证明 Rust 前端对上游语法的覆盖是可测量的。
//
// 覆盖数字来自对该语言真实语料的真实解析：上游 tree-sitter-rust 的具体
// 构造全集与 AIR 实际产出的节点类型求差集，清单入库为
// docs/dev/frontend-coverage/rust.md。
func TestRsConstructCoverage(t *testing.T) {
	g, err := grammar.Load("rs")
	if err != nil {
		t.Fatalf("grammar.Load: %v", err)
	}
	mappings := grammar.Map(g, grammar.DefaultRules("rs"))
	observed := map[air.NodeKind]int{}
	for _, src := range rsCoverageCorpus() {
		unit := rsParseOne(t, src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	cov := coverage.Compute(g, mappings, observed)
	t.Logf("rust constructs=%d covered=%d (%.1f%%) exact=%d degraded=%d out-of-scope=%d unmapped=%d",
		cov.Total, cov.Covered, cov.Percent(), cov.Exact, cov.Degraded,
		cov.OutOfScope, cov.Unmapped)
	if cov.Unmapped != 0 {
		t.Fatalf("%d upstream constructs are not mapped", cov.Unmapped)
	}
	if cov.Covered == 0 {
		t.Fatal("the real parse produced no mapped AIR kind")
	}
	report := coverage.Markdown(cov)
	path := filepath.Join("..", "..", "docs", "dev", "frontend-coverage", "rust.md")
	if err := coverage.Reconcile(path, report, *coverageUpdate); err != nil {
		t.Fatal(err)
	}
}
