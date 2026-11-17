/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

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
	"rewrite the c frontend coverage checklist")

// cCoverageCorpus 是真实 C/C++ 语料：既有整文件样例，也有构造片段。
func cCoverageCorpus() []string {
	out := make([]string, 0, 32)
	out = append(out, cSampleSource())
	return append(out,
		"#include <stdio.h>\n#define N 4\n#ifdef X\nint y;\n#else\nint z;\n#endif\n",
		"namespace n { template <typename T> struct S { T v; }; }\n",
		"enum class E : int { A = 1, B };\nunion U { int i; float f; };\n",
		"typedef unsigned long W;\nusing A = int;\n",
		"int f(int a, int b) { return a ? b : a; }\n",
		"int g(void) { int *p = 0; int a[4]; return *p + a[0]; }\n",
		"void h(void) { for (int i = 0; i < 3; ++i) { continue; } while (1) { break; } do { } while (0); }\n",
		"void k(void) { switch (x) { case 1: break; default: break; } goto end; end: ; }\n",
		"void m(void) { int v = (int)3.5; unsigned u = sizeof(int); }\n",
		"struct S { int a : 3; const char *s; };\n",
		"static inline int q(void) { return (1 + 2) * 3; }\n",
		"extern int gv;\nconst int cv = 1;\nstatic int sv;\n",
		"void t(void) { char c = 'x'; const char *s = \"str\\n\"; double d = 1.5e3; }\n",
		"int arr[3] = {1, 2, 3};\nstruct P p = {.x = 1, .y = 2};\n",
		"void u(void) { f(1, 2); }\n",
		"bool b = true;\nbool c = false;\n",
		"int *p;\nint a[10];\n",
		"[[nodiscard]] int f(void);\n",
		"void m(void) { int x = (1, 2); }\n",
		"struct S { int a : 3; };\n",
		"int arr[2] = {1, 2};\nstruct P p = {1, 2};\n",
		"void n(void) { s.f = 1; }\n",
		"int v = 1 ? 2 : 3;\n",
		"int y = a[0];\n",
		"void o(void) { __asm__(\"nop\"); }\n",
		"void p2(void) { switch (x) { case 1: break; } }\n",
	)
}

// TestCConstructCoverage 证明 C/C++ 前端对上游语法的覆盖是可测量的。
//
// 覆盖数字来自对该语言真实语料的真实解析：上游 tree-sitter-c 的具体构造
// 全集与 AIR 实际产出的节点类型求差集，清单入库为
// docs/dev/frontend-coverage/c.md。
func TestCConstructCoverage(t *testing.T) {
	g, err := grammar.Load("c")
	if err != nil {
		t.Fatalf("grammar.Load: %v", err)
	}
	mappings := grammar.Map(g, grammar.DefaultRules("c"))
	observed := map[air.NodeKind]int{}
	for _, src := range cCoverageCorpus() {
		unit, _ := cParseOne(t, src, "c++20")
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	cov := coverage.Compute(g, mappings, observed)
	t.Logf("c constructs=%d covered=%d (%.1f%%) exact=%d degraded=%d out-of-scope=%d unmapped=%d",
		cov.Total, cov.Covered, cov.Percent(), cov.Exact, cov.Degraded,
		cov.OutOfScope, cov.Unmapped)
	if cov.Unmapped != 0 {
		t.Fatalf("%d upstream constructs are not mapped", cov.Unmapped)
	}
	if cov.Covered == 0 {
		t.Fatal("the real parse produced no mapped AIR kind")
	}
	report := coverage.Markdown(cov)
	path := filepath.Join("..", "..", "docs", "dev", "frontend-coverage", "c.md")
	if err := coverage.Reconcile(path, report, *coverageUpdate); err != nil {
		t.Fatal(err)
	}
}
