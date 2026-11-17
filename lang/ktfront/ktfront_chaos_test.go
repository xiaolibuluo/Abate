/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// ktChaosInputs 覆盖损坏文件、未闭合字符串模板、嵌套 when、超深括号与畸形字节。
var ktChaosInputs = []string{
	"",
	"\n",
	"   \n\n",
	"class C {\n",
	"fun f( {\n val x = ;\n return x;\n}\n",
	"val s = \"unclosed ${name\n",
	"val s = \"\"\"unterminated\n",
	"val r = when (x) { 1 -> when (y) { else ->\n",
	"val n = a?. ?: !! ?\n",
	"((((????::;;;@@@###\n",
	"x = 1 \xff\n",
	"<<<<>>>>{}();\n\n\n",
	"fun f(\n",
	"class A { class B { class C {\n",
	strings.Repeat("List<", 200) + "Int" + strings.Repeat(">", 200) + "\n",
	strings.Repeat("(", 200) + strings.Repeat(")", 200),
	strings.Repeat("when (x) { ", 100) + "1 -> 1\n",
	strings.Repeat("val x = 1;\n", 400),
	"${}${x}\"\"\"``:::\n",
	"/* unterminated block comment\n",
}

// ktChaosUnit 构造一份单文件畸形输入的内存编译请求。
func ktChaosUnit(src string) air.SourceUnit {
	return air.SourceUnit{
		Language: "kt", Dialect: "kt2.1",
		Files: []air.SourceFile{air.NewSourceFile("main.kt", []byte(src))},
	}
}

// TestKtChaosNeverPanics 把畸形输入送过整个适配器，要求合法 unit、稳定指纹。
// panic 由 Go 测试框架自然冒泡并判 FAIL，这里只做正向合法性断言。
func TestKtChaosNeverPanics(t *testing.T) {
	fe := ktFrontend(t)
	for i, src := range ktChaosInputs {
		out, diags, err := fe.Parse(ktChaosUnit(src))
		if err != nil {
			t.Logf("chaos[%d] classified error: %v", i, err)
			continue
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("chaos[%d] %q: missing root", i, src[:ktMinInt(len(src), 40)])
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

// TestKtChaosDeterministicAcrossRuns 断言每份畸形缓冲两次独立解析逐字节一致。
func TestKtChaosDeterministicAcrossRuns(t *testing.T) {
	fe := ktFrontend(t)
	for i, src := range ktChaosInputs {
		a, _, errA := fe.Parse(ktChaosUnit(src))
		b, _, errB := fe.Parse(ktChaosUnit(src))
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

// TestKtChaosDeepNestingStaysBounded 证明深度上限约束递归且保留包根。
func TestKtChaosDeepNestingStaysBounded(t *testing.T) {
	src := strings.Repeat("(", 300) + "1" + strings.Repeat(")", 300) + "\n"
	out, root, _ := ktParse(t, src)
	if root.Kind() != air.KindPackage {
		t.Fatal("deep nesting lost the root")
	}
	if out.NodeCount() < 2 || out.NodeCount() > 20000 {
		t.Fatalf("unbounded node count %d", out.NodeCount())
	}
	if err := air.CheckIntervals(root, len(src)); err != nil {
		t.Fatalf("deep nesting produced invalid intervals: %v", err)
	}
}

// TestKtChaosFingerprint100 对混合畸形语料执行 100 次解析哈希一致性认证。
func TestKtChaosFingerprint100(t *testing.T) {
	fe := ktFrontend(t)
	request := air.SourceUnit{
		Language: "kt", Dialect: "kt2.1",
		Files: []air.SourceFile{
			air.NewSourceFile("a.kt", []byte("val s = \"unclosed ${x\n")),
			air.NewSourceFile("b.kt", []byte(strings.Repeat("when (x) { ", 40)+"1 -> 1\n")),
		},
	}
	determinism.MustMatch100(t, 0x4B540001, func() []byte {
		out, _, err := fe.Parse(request)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		sum := out.Fingerprint()
		return sum[:]
	})
}

// TestKtChaosDialectIndependence 断言多方言下的容错路径同样稳定。
func TestKtChaosDialectIndependence(t *testing.T) {
	for _, dialect := range []string{"kt1.6", "kt1.9", "kt2.1"} {
		fe := ktFrontend(t)
		request := ktChaosUnit("class C { fun M( { val x = \xff; }\n")
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
