/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

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
	"rewrite the javascript frontend coverage checklist")

// jsCoverageCorpus 是真实 JS/TS 语料：既有保真夹具，也有整文件样例。
func jsCoverageCorpus() []string {
	out := make([]string, 0, 48)
	for _, tc := range jsFidelityCases {
		out = append(out, tc.src)
	}
	return append(out,
		jsComprehensiveSource(),
		"import a, {b as c} from 'm';\nimport * as ns from 'n';\nexport {a};\nexport default f;\n",
		"export namespace N { export const v = 1; }\n",
		"class A extends B { static #p = 1; get x() { return 1; } set x(v) {} static { init(); } }\n",
		"function* gen() { yield* other(); }\nconst g = function* () {};\n",
		"const arrow = (a, b = 1, ...rest) => a + b;\nconst obj = { m() {}, get g() { return 1; }, [k]: v, ...spread };\n",
		"async function run() { await p(); const v = await q?.r?.(); }\n",
		"for (const k in o) {}\nfor (const v of arr) {}\ndo { x(); } while (cond);\n",
		"switch (v) { case 1: break; default: break; }\n",
		"try { f(); } catch (e) { g(); } finally { h(); }\n",
		"label: for (;;) { break label; continue label; }\n",
		"with (o) { p(); }\n",
		"const re = /ab+c/gi;\nconst s = `a ${b} c ${d.e}`;\n",
		"const t = a as unknown as string;\nconst u = b satisfies C;\n",
		"interface I { readonly id?: string; m(a: number): void; }\ntype T<X> = X[] | null;\nenum E { A = 1, B }\n",
		"namespace N { export const v = 1; }\ndeclare module 'm' { }\n",
		"const el = <div className=\"x\" {...p}>text {value}</div>;\n",
		"const j = <A.B c:d=\"1\" />;\n",
		"let u = 1;\nconst c = 2;\nvar v = 3;\nusing r = res();\n",
		"a?.b?.[0]?.();\nx ??= 1;\ny ||= 2;\nz &&= 3;\n",
		"debugger;\n;\n",
		"#!/usr/bin/env node\nconst x = 1;\n",
		"const big = 123n;\nconst neg = -1;\n",
		"const a = x as T;\nconst b = y satisfies U;\n",
		"class C { m() { return 1; } }\n",
		"const o = { a: 1, b: [1, 2] };\n",
		"function f(): void {}\n",
		"const t: [number, string] = [1, 'a'];\n",
		"const u: string | null = null;\n",
		"function g<T>(v: T): T { return v; }\n",
	)
}

// TestJsConstructCoverage 证明 JS/TS 前端对上游语法的覆盖是可测量的。
//
// 覆盖数字来自对该语言真实语料的真实解析：上游 tree-sitter-javascript 的
// 具体构造全集与 AIR 实际产出的节点类型求差集，清单入库为
// docs/dev/frontend-coverage/javascript.md。
func TestJsConstructCoverage(t *testing.T) {
	g, err := grammar.Load("js")
	if err != nil {
		t.Fatalf("grammar.Load: %v", err)
	}
	mappings := grammar.Map(g, grammar.DefaultRules("js"))
	observed := map[air.NodeKind]int{}
	for i, src := range jsCoverageCorpus() {
		unit, _ := jsParse(t, jsFrontend(t), "f.js", "", src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
		_ = i
	}
	cov := coverage.Compute(g, mappings, observed)
	t.Logf("javascript constructs=%d covered=%d (%.1f%%) exact=%d degraded=%d out-of-scope=%d unmapped=%d",
		cov.Total, cov.Covered, cov.Percent(), cov.Exact, cov.Degraded,
		cov.OutOfScope, cov.Unmapped)
	if cov.Unmapped != 0 {
		t.Fatalf("%d upstream constructs are not mapped", cov.Unmapped)
	}
	if cov.Covered == 0 {
		t.Fatal("the real parse produced no mapped AIR kind")
	}
	report := coverage.Markdown(cov)
	path := filepath.Join("..", "..", "docs", "dev", "frontend-coverage", "javascript.md")
	if err := coverage.Reconcile(path, report, *coverageUpdate); err != nil {
		t.Fatal(err)
	}
}
