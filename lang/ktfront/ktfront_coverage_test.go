/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

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
	"rewrite the kotlin frontend coverage checklist")

// ktCoverageCorpus 是真实 Kotlin 语料：既有保真夹具，也有构造片段。
func ktCoverageCorpus() []string {
	out := make([]string, 0, 40)
	out = append(out, ktFidelityCorpus()...)
	return append(out,
		"package demo\nimport kotlin.math.abs\nimport kotlin.collections.* as k\n",
		"data class User(val name: String, val age: Int = 0)\n",
		"sealed class Result { data class Ok(val n: Int) : Result(); data class Err(val m: String) : Result() }\n",
		"open class Base { open fun m() {} }\nclass Derived : Base() { override fun m() {} }\n",
		"interface Repo { fun load(id: Int): String }\n",
		"object Singleton { val v = 1 }\ncompanion object Factory { const val MAX = 10 }\n",
		"enum class Dir(val deg: Int) { NORTH(0), SOUTH(180) }\n",
		"annotation class Marker\n@Marker fun tagged() {}\n",
		"typealias Handler = (Int) -> Unit\n",
		"suspend fun fetch(url: String): Int { return url.length }\n",
		"fun String.shout(): String = this.uppercase()\n",
		"fun f(vararg xs: Int, block: () -> Unit) {}\n",
		"fun g(x: Int = 1, y: Int = 2): Int = x + y\n",
		"val lambda: (Int) -> Int = { it + 1 }\n",
		"val n = a?.length ?: 0\nval m = a!!.length\n",
		"val r = when (x) { 1 -> a(); in 2..3 -> b(); else -> c() }\n",
		"val t: String? = null\nval u: String = t!!\n",
		"val list = listOf(1, 2, 3)\nval map = mapOf(\"a\" to 1)\nval arr = intArrayOf(1)\n",
		"for (i in 0..10) { println(i) }\nwhile (true) { break }\ndo { } while (false)\n",
		"try { f() } catch (e: Exception) { g() } finally { h() }\n",
		"val c = x as? String\nval d = x is String\n",
		"class C { init { } constructor() { } }\n",
		"class C { val p: Int get() = 1; var q: Int = 2 }\n",
		"fun <T> generic(v: T): T = v\n",
		"inline fun <reified T> reified(v: T): String = T::class.simpleName ?: \"\"\n",
		"val s = \"hi $name and ${name.length}\"\nval raw = \"\"\"raw $x\"\"\"\n",
		"val n = 42\nval l = 42L\nval u = 42u\nval h = 0xFF\nval b = 0b1010\nval f = 1.5\nval c = 'x'\nval t = true\n",
		"// line\n/* block */\n/** KDoc */\n",
		"val n = arr[0]\n",
		"val s: String = \"x\"\n",
		"fun <T> g(v: T): T = v\n",
		"val f: (Int) -> Int = { it }\n",
		"val l = listOf(1)\n",
		"val o = object {}\n",
		"val sp = *arr\n",
		"loop@ for (i in 0..1) { break@loop }\n",
		"#!/usr/bin/env kotlin\n",
		"val c = x as String\n",
		"val p: Int? = null\nval q: Int = p!!\n",
		"fun pick(x: Int): String {\n  when (x) {\n    1 -> return \"a\"\n    else -> return \"b\"\n  }\n}\n",
		"open class B { open fun m() {} }\nclass D : B() { override fun m() { super.m() } }\n",
		"class E { constructor(x: Int) : this() { } }\n",
	)
}

// TestKtConstructCoverage 证明 Kotlin 前端对上游语法的覆盖是可测量的。
//
// 覆盖数字来自对该语言真实语料的真实解析：上游 tree-sitter-kotlin 的具体
// 构造全集与 AIR 实际产出的节点类型求差集，清单入库为
// docs/dev/frontend-coverage/kotlin.md。
func TestKtConstructCoverage(t *testing.T) {
	g, err := grammar.Load("kt")
	if err != nil {
		t.Fatalf("grammar.Load: %v", err)
	}
	mappings := grammar.Map(g, grammar.DefaultRules("kt"))
	observed := map[air.NodeKind]int{}
	for _, src := range ktCoverageCorpus() {
		unit, _, _ := ktParse(t, src)
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	cov := coverage.Compute(g, mappings, observed)
	t.Logf("kotlin constructs=%d covered=%d (%.1f%%) exact=%d degraded=%d out-of-scope=%d unmapped=%d",
		cov.Total, cov.Covered, cov.Percent(), cov.Exact, cov.Degraded,
		cov.OutOfScope, cov.Unmapped)
	if cov.Unmapped != 0 {
		t.Fatalf("%d upstream constructs are not mapped", cov.Unmapped)
	}
	if cov.Covered == 0 {
		t.Fatal("the real parse produced no mapped AIR kind")
	}
	report := coverage.Markdown(cov)
	path := filepath.Join("..", "..", "docs", "dev", "frontend-coverage", "kotlin.md")
	if err := coverage.Reconcile(path, report, *coverageUpdate); err != nil {
		t.Fatal(err)
	}
}
