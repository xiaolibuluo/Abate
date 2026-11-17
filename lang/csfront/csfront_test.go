/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
	"abate/lang/frontkit"
)

// csFrontend 返回一个新的 C# 适配器，失败即终止调用测试。
func csFrontend(t *testing.T) air.Frontend {
	t.Helper()
	fe, err := NewCsFrontend()
	if err != nil {
		t.Fatalf("NewCsFrontend: %v", err)
	}
	return fe
}

// csParse 把一段单文件源码归约为 unit、根节点与诊断。
func csParse(t *testing.T, src string) (*air.Unit, air.Node, []air.Diagnostic) {
	t.Helper()
	unit := air.SourceUnit{
		Language: "cs", Dialect: "cs12",
		Files: []air.SourceFile{air.NewSourceFile("main.cs", []byte(src))},
	}
	out, diags, err := csFrontend(t).Parse(unit)
	if err != nil {
		t.Fatalf("Parse(%q): %v", src, err)
	}
	if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
		t.Fatalf("Parse(%q): missing AIR root", src)
	}
	return out, out.Roots[0], diags
}

// csTokenize 用 C# 词法表产生与 FrontendBase.Parse 相同的 token 流。
func csTokenize(t *testing.T, src string) []frontkit.Token {
	t.Helper()
	return csTokenizeWith(t, (&csSpec{}).LexSpec(), src)
}

// csTokenizeWith 用指定的词法表产生 token 流，供变异测试使用。
func csTokenizeWith(t *testing.T, spec frontkit.LexSpec, src string) []frontkit.Token {
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

// csNodes 返回一棵树的先序快照。
func csNodes(root air.Node) []air.Node {
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

func csHasKind(root air.Node, kind air.NodeKind) bool {
	for _, node := range csNodes(root) {
		if node.Kind() == kind {
			return true
		}
	}
	return false
}

func csHasNote(root air.Node, note string) bool {
	return csHasSem(root, air.SemNote, note)
}

func csHasSem(root air.Node, key air.SemKey, value string) bool {
	for _, node := range csNodes(root) {
		if stored, ok := node.Sem(key).(string); ok && stored == value {
			return true
		}
	}
	return false
}

// 构造到 AIR 的映射表放在 csfront_fidelity_test.go，保证每个文件不超过 300 行。

// TestCsConstructToAIRMapping 证明 ≥20 个真实 C# 构造归约为冻结 AIR 种类与标注。
func TestCsConstructToAIRMapping(t *testing.T) {
	cases := csConstructCases()
	if len(cases) < 20 {
		t.Fatalf("only %d mapping cases; at least 20 required", len(cases))
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, root, _ := csParse(t, tc.src)
			if out.NodeCount() < 2 {
				t.Fatalf("empty AIR shell for %q", tc.src)
			}
			if !csHasKind(root, tc.kind) {
				t.Fatalf("%s: missing %s in tree for %q", tc.name, tc.kind, tc.src)
			}
			if tc.note != "" && !csHasNote(root, tc.note) {
				t.Fatalf("%s: missing SemNote %q for %q", tc.name, tc.note, tc.src)
			}
		})
	}
}

// TestCsFrontendSatisfiesAirFrontend 证明返回值满足 air.Frontend 接口。
func TestCsFrontendSatisfiesAirFrontend(t *testing.T) {
	fe := csFrontend(t)
	if fe == nil {
		t.Fatal("NewCsFrontend returned a nil air.Frontend")
	}
	asFrontend := fe
	unit, _, err := asFrontend.Parse(air.SourceUnit{
		Language: "cs", Dialect: "cs12",
		Files: []air.SourceFile{
			air.NewSourceFile("main.cs", []byte("namespace N { class C { } }\n")),
		},
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if unit == nil || unit.NodeCount() < 5 {
		t.Fatalf("expected a real node tree, got %v", unit)
	}
	if err := unit.Validate(); err != nil {
		t.Fatalf("unit.Validate: %v", err)
	}
}

// TestCsFrontendContract 断言有层级、结构合法且带确定性注释的 unit。
func TestCsFrontendContract(t *testing.T) {
	src := "using System;\n" +
		"namespace Demo {\n" +
		"public class C {\n" +
		"  public string? Name { get; set; }\n" +
		"  public int M(int x) { return x + 1; }\n" +
		"}\n}\n"
	out, diags, err := csFrontend(t).Parse(air.SourceUnit{
		Language: "cs", Dialect: "cs12",
		Files: []air.SourceFile{air.NewSourceFile("main.cs", []byte(src))},
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
	if len(out.Roots[0].Children()) < 2 {
		t.Fatal("root has no statement children")
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got := out.Annotations["lang.cs.dialect"]; got != "cs12" {
		t.Fatalf("dialect annotation=%q, want cs12", got)
	}
}

// TestCsDeterminismFingerprint 用 MustMatch100 认证指纹逐字节稳定。
func TestCsDeterminismFingerprint(t *testing.T) {
	fe := csFrontend(t)
	src := "using System.Linq;\n" +
		"namespace Demo {\n" +
		"public class C {\n" +
		"  public string? Name { get; set; }\n" +
		"  public object M(object a) { return a?.Name ?? \"none\"; }\n" +
		"  public void Q() { var q = from x in xs where x > 0 select x; }\n" +
		"}\n}\n"
	unit := air.SourceUnit{
		Language: "cs", Dialect: "cs12",
		Files: []air.SourceFile{air.NewSourceFile("main.cs", []byte(src))},
	}
	determinism.MustMatch100(t, 0x4353, func() []byte {
		out, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		sum := out.Fingerprint()
		return sum[:]
	})
}

// TestCsDialectNormalization 检查别名到 cs8-cs12 的确定性归一。
func TestCsDialectNormalization(t *testing.T) {
	good := map[string]string{
		"": "cs12", "cs12": "cs12", "CS8": "cs8", "9.0": "cs9",
		"C# 12": "cs12", "net8": "cs12", "latest": "cs12",
		"  NET6 ": "cs10", "c#11": "cs11",
	}
	for in, want := range good {
		got, err := (&csSpec{}).NormalizeDialect(in)
		if err != nil || got != want {
			t.Fatalf("NormalizeDialect(%q)=%q,%v want %q", in, got, err, want)
		}
	}
	for _, in := range []string{"java", "cs7", "go1.26", "13.0"} {
		if got, err := (&csSpec{}).NormalizeDialect(in); err == nil {
			t.Fatalf("NormalizeDialect(%q)=%q, want error", in, got)
		}
	}
}

// TestCsLexSpecContract 检查冻结词法表契约。
func TestCsLexSpecContract(t *testing.T) {
	spec := &csSpec{}
	if spec.Tag() != "cs" {
		t.Fatalf("Tag=%q, want cs", spec.Tag())
	}
	lex := spec.LexSpec()
	if lex == nil || lex.IndentSensitive() {
		t.Fatal("C# lexer must not be indentation sensitive")
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
		{"keywords", lex.Keywords(), "record"},
		{"operators", lex.Operators(), "??="},
		{"quotes", lex.StringQuotes(), `"`},
		{"prefixes", lex.RawStringPrefixes(), "@"},
		{"separators", lex.NumberSeparators(), "_"},
	} {
		if !csContains(table.values, table.want) {
			t.Fatalf("%s table missing %q", table.name, table.want)
		}
	}
}

// TestCsSemanticKeys 断言非 note 的 AIR 语义标注落在真实 SemKey 上。
func TestCsSemanticKeys(t *testing.T) {
	cases := []struct {
		name, src string
		key       air.SemKey
		value     string
	}{
		{"scope", "class C { }\n", air.SemScopeID, "compilation-unit"},
		{"async-method", "class C { async Task M() { } }\n", air.SemAsync, "true"},
		{"binary-op", "class C { int M(int a) { return a + 1; } }\n", air.SemOperator, "+"},
		{"assign-op", "class C { void M() { x = 1; } }\n", air.SemOperator, "="},
		{"directive", "#region Helpers\n", air.SemDirective, "region"},
		{"mutability", "class C { readonly int x; }\n", air.SemMutability, "readonly"},
		{"partial", "partial class C { }\n", air.SemPartial, "true"},
		{"string-tag", "class C { string s = $\"hi {n}\"; }\n", air.SemTag, "interpolated"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, root, _ := csParse(t, tc.src)
			if !csHasSem(root, tc.key, tc.value) {
				t.Fatalf("missing %s=%q for %q", tc.key, tc.value, tc.src)
			}
		})
	}
}
