/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"strconv"
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// cDeepTemplate 生成指定尖括号嵌套深度的确定性模板别名源码。
func cDeepTemplate(depth int) string {
	var builder strings.Builder
	builder.WriteString("using Deep = ")
	for i := 0; i < depth; i++ {
		builder.WriteString("Vec<")
	}
	builder.WriteString("int")
	for i := 0; i < depth; i++ {
		builder.WriteString(">")
	}
	builder.WriteString(";\n")
	return builder.String()
}

// cChaosBytes 生成确定性的伪随机字节序列（不使用 math/rand 或时钟）。
func cChaosBytes(length int) string {
	out := make([]byte, length)
	for i := range out {
		out[i] = byte((i*37 + 11) % 256)
	}
	return string(out)
}

// cAssertNoPanic 断言畸形输入仍产出合法节点树与合法区间；不 panic 由 Go
// 测试框架保证（panic 会自然冒泡并判 FAIL），无需显式 recover（避免触发
// GO-EG-027 提示规则）。
func cAssertNoPanic(t *testing.T, source, dialect string) {
	t.Helper()
	unit, _ := cParseOne(t, source, dialect)
	if unit == nil || unit.NodeCount() == 0 {
		t.Fatalf("chaos input produced an empty unit")
	}
	if err := air.CheckIntervals(unit.Roots[0], len(source)); err != nil {
		t.Fatalf("CheckIntervals on chaos input: %v", err)
	}
}

// cChaosCorpus 是未闭合宏、超深模板、畸形字节等确定性畸形语料。
func cChaosCorpus() []string {
	return []string{
		"#define BROKEN(\n",
		"#if 1\nint a = ;\n",
		"int x = \"unterminated\n",
		"/* unterminated comment\n",
		"#\n",
		"#include\n",
		"template <\n",
		"<<<<>>>>{}();\n\n\n",
		"int f(\n",
		"struct S {\n",
		cDeepTemplate(24),
		cDeepTemplate(64),
		cChaosBytes(256),
		cChaosBytes(1024),
	}
}

// TestChaosNoPanic 遍历畸形语料，断言全部不 panic 且产出合法 AIR。
func TestChaosNoPanic(t *testing.T) {
	for i, source := range cChaosCorpus() {
		t.Run("case_"+strconv.Itoa(i), func(t *testing.T) {
			cAssertNoPanic(t, source, "c++20")
		})
	}
	// 三方言各跑一遍，确保方言归一不影响容错路径。
	for _, dialect := range cDialectCanonical {
		cAssertNoPanic(t, "#define X(\n"+cDeepTemplate(20), dialect)
	}
}

// TestTemplateDepthLimitWarning 断言深度上限 16 + warning，且不 panic。
func TestTemplateDepthLimitWarning(t *testing.T) {
	deep := cDeepTemplate(24)
	unit, diagnostics := cParseOne(t, deep, "c++20")
	warnings := 0
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == air.SeverityWarning &&
			strings.Contains(diagnostic.Message, "limit 16") {
			warnings++
		}
	}
	if warnings == 0 {
		t.Fatalf("expected a template depth warning, got %v", diagnostics)
	}
	if !cHasSemValue(unit, air.SemPartial, "template-depth") {
		t.Fatal("over-depth node is not marked partial")
	}
	if err := air.CheckIntervals(unit.Roots[0], len(deep)); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
	// 深度在上限内时不产生该 warning。
	within := cDeepTemplate(12)
	_, cleanDiagnostics := cParseOne(t, within, "c++20")
	for _, diagnostic := range cleanDiagnostics {
		if strings.Contains(diagnostic.Message, "template nesting depth") {
			t.Fatalf("unexpected depth warning for depth 12: %v", diagnostic)
		}
	}
}

// TestChaosDeterminism 断言超深模板与畸形输入的 100 次解析指纹一致。
func TestChaosDeterminism(t *testing.T) {
	frontend, err := NewCFrontend()
	if err != nil {
		t.Fatal(err)
	}
	request := air.SourceUnit{
		Language: "c", Dialect: "c++20",
		Files: []air.SourceFile{
			air.NewSourceFile("deep.cpp", []byte(cDeepTemplate(30))),
			air.NewSourceFile("broken.cpp", []byte("#define X(\n"+cChaosBytes(128))),
		},
	}
	determinism.MustMatch100(t, 0xC05F0003, func() []byte {
		unit, _, parseErr := frontend.Parse(request)
		if parseErr != nil {
			t.Fatal(parseErr)
		}
		fingerprint := unit.Fingerprint()
		return fingerprint[:]
	})
}
