/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
	"abate/lang/frontkit"
)

// ktFrontend 返回一个新的 Kotlin 适配器，失败即终止调用测试。
func ktFrontend(t *testing.T) air.Frontend {
	t.Helper()
	fe, err := NewKtFrontend()
	if err != nil {
		t.Fatalf("NewKtFrontend: %v", err)
	}
	return fe
}

// ktParse 把一段单文件源码归约为 unit、根节点与诊断。
func ktParse(t *testing.T, src string) (*air.Unit, air.Node, []air.Diagnostic) {
	t.Helper()
	unit := air.SourceUnit{
		Language: "kt", Dialect: "kt2.1",
		Files: []air.SourceFile{air.NewSourceFile("main.kt", []byte(src))},
	}
	out, diags, err := ktFrontend(t).Parse(unit)
	if err != nil {
		t.Fatalf("Parse(%q): %v", src, err)
	}
	if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
		t.Fatalf("Parse(%q): missing AIR root", src)
	}
	return out, out.Roots[0], diags
}

// ktTokenize 用 Kotlin 词法表产生与 FrontendBase.Parse 相同的 token 流。
func ktTokenize(t *testing.T, src string) []frontkit.Token {
	t.Helper()
	return ktTokenizeWith(t, (&ktSpec{}).LexSpec(), src)
}

// ktTokenizeWith 用指定的词法表产生 token 流，供变异测试使用。
func ktTokenizeWith(t *testing.T, spec frontkit.LexSpec, src string) []frontkit.Token {
	t.Helper()
	lexer, err := frontkit.NewLexer(spec, []byte(src))
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize: %v", err)
	}
	return tokens
}

// ktNodes 返回一棵树的先序快照。
func ktNodes(root air.Node) []air.Node {
	out := make([]air.Node, 0, 16)
	var walk func(air.Node)
	walk = func(node air.Node) {
		if node == nil {
			return
		}
		out = append(out, node)
		for _, child := range node.Children() {
			walk(child)
		}
	}
	walk(root)
	return out
}

func ktHasKind(root air.Node, kind air.NodeKind) bool {
	for _, node := range ktNodes(root) {
		if node.Kind() == kind {
			return true
		}
	}
	return false
}

func ktHasNote(root air.Node, note string) bool {
	return ktHasSem(root, air.SemNote, note)
}

func ktHasSem(root air.Node, key air.SemKey, value string) bool {
	for _, node := range ktNodes(root) {
		if stored, ok := node.Sem(key).(string); ok && stored == value {
			return true
		}
	}
	return false
}

// ktCountKind 统计一棵树里指定种类的节点数量。
func ktCountKind(root air.Node, kind air.NodeKind) int {
	count := 0
	for _, node := range ktNodes(root) {
		if node.Kind() == kind {
			count++
		}
	}
	return count
}

// ktMinInt 返回两个整数中较小的一个，避免引入额外依赖。
func ktMinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 构造到 AIR 的映射表放在 ktfront_fidelity_test.go，保证每个文件不超过 300 行。

// TestKtConstructToAIRMapping 证明 ≥20 个真实 Kotlin 构造归约为冻结 AIR 种类与标注。
func TestKtConstructToAIRMapping(t *testing.T) {
	cases := ktConstructCases()
	if len(cases) < 20 {
		t.Fatalf("only %d mapping cases; at least 20 required", len(cases))
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, root, _ := ktParse(t, tc.src)
			if out.NodeCount() < 2 {
				t.Fatalf("empty AIR shell for %q", tc.src)
			}
			if !ktHasKind(root, tc.kind) {
				t.Fatalf("%s: missing %s in tree for %q", tc.name, tc.kind, tc.src)
			}
			if tc.note != "" && !ktHasNote(root, tc.note) {
				t.Fatalf("%s: missing SemNote %q for %q", tc.name, tc.note, tc.src)
			}
		})
	}
}

// TestKtFrontendSatisfiesAirFrontend 证明返回值满足 air.Frontend 接口。
func TestKtFrontendSatisfiesAirFrontend(t *testing.T) {
	fe := ktFrontend(t)
	if fe == nil {
		t.Fatal("NewKtFrontend returned a nil air.Frontend")
	}
	asFrontend := fe
	unit, _, err := asFrontend.Parse(air.SourceUnit{
		Language: "kt", Dialect: "kt2.1",
		Files: []air.SourceFile{
			air.NewSourceFile("main.kt", []byte("package demo\nclass C { }\n")),
		},
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if unit == nil || unit.NodeCount() < 4 {
		t.Fatalf("expected a real node tree, got %v", unit)
	}
	if err := unit.Validate(); err != nil {
		t.Fatalf("unit.Validate: %v", err)
	}
}

// TestKtFrontendContract 断言有层级、结构合法且带确定性注释的 unit。
func TestKtFrontendContract(t *testing.T) {
	src := "package demo\n" +
		"import kotlin.math.abs\n" +
		"data class User(val name: String)\n" +
		"suspend fun fetch(url: String): Int {\n" +
		"  val label = \"hi $url\"\n" +
		"  val n = url?.length ?: 0\n" +
		"  return n!!\n" +
		"}\n"
	out, diags, err := ktFrontend(t).Parse(air.SourceUnit{
		Language: "kt", Dialect: "kt2.1",
		Files: []air.SourceFile{air.NewSourceFile("main.kt", []byte(src))},
	})
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
	if len(out.Roots[0].Children()) < 3 {
		t.Fatal("root has no statement children")
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got := out.Annotations["lang.kt.dialect"]; got != "kt2.1" {
		t.Fatalf("dialect annotation=%q, want kt2.1", got)
	}
}

// TestKtDeterminismFingerprint 用 MustMatch100 认证指纹逐字节稳定。
func TestKtDeterminismFingerprint(t *testing.T) {
	fe := ktFrontend(t)
	src := "package demo\n" +
		"data class User(val name: String)\n" +
		"fun String.shout() = uppercase()\n" +
		"suspend fun fetch(url: String?): Int {\n" +
		"  val label = \"hi $url ${url.length}\"\n" +
		"  val n = when (url?.length) { 1 -> 1 else -> 0 }\n" +
		"  return url?.length ?: n\n" +
		"}\n"
	unit := air.SourceUnit{
		Language: "kt", Dialect: "kt2.1",
		Files: []air.SourceFile{air.NewSourceFile("main.kt", []byte(src))},
	}
	determinism.MustMatch100(t, 0x4B54, func() []byte {
		out, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		sum := out.Fingerprint()
		return sum[:]
	})
}

// TestKtLexSpecContract 检查冻结词法表契约。
func TestKtLexSpecContract(t *testing.T) {
	spec := &ktSpec{}
	if spec.Tag() != "kt" {
		t.Fatalf("Tag=%q, want kt", spec.Tag())
	}
	lex := spec.LexSpec()
	if lex == nil || lex.IndentSensitive() {
		t.Fatal("Kotlin lexer must not be indentation sensitive")
	}
	if lex.LineComment() != "//" {
		t.Fatalf("LineComment=%q", lex.LineComment())
	}
	if opening, closing := lex.BlockComment(); opening != "/*" || closing != "*/" {
		t.Fatalf("BlockComment=%q,%q", opening, closing)
	}
	for _, table := range []struct {
		name   string
		values []string
		want   string
	}{
		{"keywords", lex.Keywords(), "suspend"},
		{"operators", lex.Operators(), "?."},
		{"quotes", lex.StringQuotes(), `"""`},
		{"prefixes", lex.RawStringPrefixes(), `"""`},
		{"separators", lex.NumberSeparators(), "_"},
	} {
		if !ktContains(table.values, table.want) {
			t.Fatalf("%s table missing %q", table.name, table.want)
		}
	}
}

// TestKtSemanticKeys 断言非 note 的 AIR 语义标注落在真实 SemKey 上。
func TestKtSemanticKeys(t *testing.T) {
	cases := []struct {
		name, src string
		key       air.SemKey
		value     string
	}{
		{"scope", "class C { }\n", air.SemScopeID, "compilation-unit"},
		{"async", "suspend fun fetch() { }\n", air.SemAsync, "true"},
		{"binary-op", "fun f(a: Int): Int { return a + 1 }\n", air.SemOperator, "+"},
		{"assign-op", "x = 1\n", air.SemOperator, "="},
		{"mutability", "lateinit var s: String\n", air.SemMutability, "lateinit"},
		{"partial", "class C {\n", air.SemPartial, "true"},
		{"string-tag", "val s = \"hi $n\"\n", air.SemTag, "string-template"},
		{"raw-tag", "val s = \"\"\"raw\"\"\"\n", air.SemTag, "raw-string"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, root, _ := ktParse(t, tc.src)
			if !ktHasSem(root, tc.key, tc.value) {
				t.Fatalf("missing %s=%q for %q", tc.key, tc.value, tc.src)
			}
		})
	}
}
