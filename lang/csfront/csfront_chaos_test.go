/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// csChaosInputs 覆盖损坏文件、未闭合插值字符串、超深泛型、畸形字节与空缓冲。
var csChaosInputs = []string{
	"",
	"\n",
	"   \n\n",
	"class C {\n",
	"class C { void M( {\n int x = ;\n return x;\n}\n",
	"string s = $\"unclosed {name\n",
	"string s = @\"unterminated\n",
	"var q = from x in xs where select group\n",
	"#region\n#pragma\n#\n",
	"[[[[Obsolete]]]]\n",
	"x = 1 \xff\n",
	"<<<<>>>>{}();\n\n\n",
	"int f(\n",
	"namespace N { class C { class D { class E {\n",
	strings.Repeat("List<", 200) + "int" + strings.Repeat(">", 200) + ";\n",
	strings.Repeat("(", 200) + strings.Repeat(")", 200),
	strings.Repeat("if (x) { ", 100) + "return 1;\n",
	strings.Repeat("x = 1;\n", 400),
	":::.???;;;@@@###\n",
	"/* unterminated block comment\n",
}

// csChaosUnit 构造一份单文件畸形输入的内存编译请求。
func csChaosUnit(src string) air.SourceUnit {
	return air.SourceUnit{
		Language: "cs", Dialect: "cs12",
		Files: []air.SourceFile{air.NewSourceFile("main.cs", []byte(src))},
	}
}

// TestCsChaosNeverPanics 把畸形输入送过整个适配器，要求合法 unit、稳定指纹。
// panic 由 Go 测试框架自然冒泡并判 FAIL，这里只做正向合法性断言。
func TestCsChaosNeverPanics(t *testing.T) {
	fe := csFrontend(t)
	for i, src := range csChaosInputs {
		out, diags, err := fe.Parse(csChaosUnit(src))
		if err != nil {
			t.Logf("chaos[%d] classified error: %v", i, err)
			continue
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("chaos[%d] %q: missing root", i, src[:csMinInt(len(src), 40)])
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("chaos[%d]: invalid unit: %v", i, err)
		}
		if out.NodeCount() < 1 {
			t.Fatalf("chaos[%d]: empty tree", i)
		}
		if first, second := out.Fingerprint(), out.Fingerprint(); first != second {
			t.Fatalf("chaos[%d]: fingerprint unstable", i)
		}
		_ = diags
	}
}

// TestCsChaosDeterministicAcrossRuns 断言每份畸形缓冲两次独立解析逐字节一致。
func TestCsChaosDeterministicAcrossRuns(t *testing.T) {
	fe := csFrontend(t)
	for i, src := range csChaosInputs {
		a, _, errA := fe.Parse(csChaosUnit(src))
		b, _, errB := fe.Parse(csChaosUnit(src))
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

// TestCsChaosDeepNestingStaysBounded 证明深度上限约束递归且保留包根。
func TestCsChaosDeepNestingStaysBounded(t *testing.T) {
	src := strings.Repeat("[", 300) + "1" + strings.Repeat("]", 300) + ";\n"
	out, root, _ := csParse(t, src)
	if root.Kind() != air.KindPackage {
		t.Fatal("deep nesting lost the root")
	}
	if out.NodeCount() < 2 || out.NodeCount() > 10000 {
		t.Fatalf("unbounded node count %d", out.NodeCount())
	}
	if err := air.CheckIntervals(root, len(src)); err != nil {
		t.Fatalf("deep nesting produced invalid intervals: %v", err)
	}
}

// TestCsChaosFingerprint100 对混合畸形语料执行 100 次解析哈希一致性认证。
func TestCsChaosFingerprint100(t *testing.T) {
	fe := csFrontend(t)
	request := air.SourceUnit{
		Language: "cs", Dialect: "cs12",
		Files: []air.SourceFile{
			air.NewSourceFile("a.cs", []byte("string s = $\"unclosed {x\n")),
			air.NewSourceFile("b.cs", []byte(strings.Repeat("List<", 40)+"int>\n")),
		},
	}
	determinism.MustMatch100(t, 0x43530001, func() []byte {
		out, _, err := fe.Parse(request)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		sum := out.Fingerprint()
		return sum[:]
	})
}

// TestCsChaosDialectIndependence 断言三种方言下的容错路径同样稳定。
func TestCsChaosDialectIndependence(t *testing.T) {
	for _, dialect := range []string{"cs8", "cs10", "cs12"} {
		fe := csFrontend(t)
		request := csChaosUnit("class C { void M( { int x = \xff; }\n")
		request.Dialect = dialect
		out, _, err := fe.Parse(request)
		if err != nil {
			t.Fatalf("Parse(%s): %v", dialect, err)
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("invalid unit for %s: %v", dialect, err)
		}
		if out.NodeCount() < 2 {
			t.Fatalf("empty tree for %s", dialect)
		}
	}
}
