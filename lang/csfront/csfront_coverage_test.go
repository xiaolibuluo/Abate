/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

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
	"rewrite the csharp frontend coverage checklist")

// csCoverageCorpus 是真实 C# 语料：既有保真夹具，也有构造片段。
func csCoverageCorpus() []string {
	out := make([]string, 0, 40)
	out = append(out, csFidelityCorpus()...)
	return append(out,
		"using System;\nusing static System.Math;\nusing Alias = System.Collections.Generic.List<int>;\n",
		"namespace Demo.Sub { public class C { } }\n",
		"public interface I { void M(); int P { get; set; } }\n",
		"public enum E { A, B = 2 }\n",
		"public struct S { public int X; }\npublic record R(int X, string Y);\n",
		"public delegate int D(int x);\n",
		"public class C { public event EventHandler E; public int this[int i] => i; }\n",
		"public static int operator +(C a, C b) => 0;\npublic C() { } ~C() { }\n",
		"public class C { public C(int x) : this() { } }\n",
		"class C { void M() { var q = from x in xs where x > 0 select x; } }\n",
		"class C { void M() { var q = from x in xs group x by x into g select g; } }\n",
		"class C { void M() { var q = from x in xs join y in ys on x equals y select x; } }\n",
		"class C { void M() { int? n = null; n ??= 1; var v = n!.ToString(); } }\n",
		"class C { void M() { var a = new[] { 1, 2 }; var b = new List<int> { 1 }; } }\n",
		"class C { void M() { var (a, b) = (1, 2); var c = a switch { 1 => 2, _ => 3 }; } }\n",
		"class C { void M() { try { f(); } catch (Exception e) when (e != null) { g(); } finally { h(); } } }\n",
		"class C { void M() { checked { int x = 1; } unchecked { int y = 2; } } }\n",
		"class C { unsafe void M() { fixed (int* p = &x) { } } }\n",
		"class C { void M() { lock (o) { } using (s) { } foreach (var x in xs) { } } }\n",
		"class C { void M() { for (int i = 0; i < 3; i++) { continue; } while (true) { break; } do { } while (false); } }\n",
		"class C { void M() { goto end; end: return; } }\n",
		"class C { async Task M() { await Task.Delay(1); } }\n",
		"class C { void M() { var t = (1, 2); var u = typeof(int); var s = sizeof(int); } }\n",
		"class C { void M() { string s = $\"a {b} c\"; string v = @\"C:\\x\"; } }\n",
		"[Obsolete, Serializable]\npublic class C { }\n",
		"class C { const int X = 1; static readonly int Y = 2; public int Z { get; init; } }\n",
		"class C { void M(object o) { if (o is string s) { } var t = o as string; } }\n",
		"class C { void M() { Action a = () => { }; Func<int,int> f = x => x + 1; } }\n",
		"class C { void M() { f(); } }\n",
		"class C { void M() { switch (x) { case 1: break; default: break; } } }\n",
		"class C { void M() { throw new Exception(); } }\n",
		"class C { void M() { var t = (1, 2); } }\n",
		"class C { void M() { int[] a = new int[1]; } }\n",
		"class C { void M() { ref int r = ref x; } }\n",
		"class C { void M() { int[] a = [1, 2]; } }\n",
		"class C { int M() => 1; }\n",
		"class C { void M() { var c = x switch { 1 => 2, _ => 3 }; } }\n",
		"class C { void M() { int? n = null; } }\n",
		"#pragma warning disable\nclass C { }\n",
	)
}

// TestCsConstructCoverage 证明 C# 前端对上游语法的覆盖是可测量的。
//
// 覆盖数字来自对该语言真实语料的真实解析：上游 tree-sitter-c-sharp 的具体
// 构造全集与 AIR 实际产出的节点类型求差集，清单入库为
// docs/dev/frontend-coverage/csharp.md。
func TestCsConstructCoverage(t *testing.T) {
	g, err := grammar.Load("cs")
	if err != nil {
		t.Fatalf("grammar.Load: %v", err)
	}
	mappings := grammar.Map(g, grammar.DefaultRules("cs"))
	observed := map[air.NodeKind]int{}
	for _, src := range csCoverageCorpus() {
		unit, _, _ := csParse(t, src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	cov := coverage.Compute(g, mappings, observed)
	t.Logf("csharp constructs=%d covered=%d (%.1f%%) exact=%d degraded=%d out-of-scope=%d unmapped=%d",
		cov.Total, cov.Covered, cov.Percent(), cov.Exact, cov.Degraded,
		cov.OutOfScope, cov.Unmapped)
	if cov.Unmapped != 0 {
		t.Fatalf("%d upstream constructs are not mapped", cov.Unmapped)
	}
	if cov.Covered == 0 {
		t.Fatal("the real parse produced no mapped AIR kind")
	}
	report := coverage.Markdown(cov)
	path := filepath.Join("..", "..", "docs", "dev", "frontend-coverage", "csharp.md")
	if err := coverage.Reconcile(path, report, *coverageUpdate); err != nil {
		t.Fatal(err)
	}
}
