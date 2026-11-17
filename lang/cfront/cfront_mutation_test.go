/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"strings"
	"testing"

	"abate/lang/air"
)

// cDropKeyword 返回移除了指定条目的有序表副本，用于确定性变异。
func cDropKeyword(table []string, drop string) []string {
	out := make([]string, 0, len(table))
	for _, item := range table {
		if item != drop {
			out = append(out, item)
		}
	}
	return out
}

// cSpecParse 用指定配置的 cSpec 直接归约一段源码的 token 流。
func cSpecParse(t *testing.T, spec *cSpec,
	source string) (air.Node, []air.Diagnostic) {
	t.Helper()
	tokens := cLex(t, source)
	root, diagnostics, err := spec.ParseTokens(tokens, "mutation.cpp")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if root == nil {
		t.Fatal("ParseTokens returned a nil root")
	}
	return root, diagnostics
}

// cTreeHasSem 判断一棵节点树里是否存在指定语义标注。
func cTreeHasSem(root air.Node, key air.SemKey, want string) bool {
	found := false
	air.Walk(root, func(node air.Node) bool {
		if value, ok := cSemValue(node, key); ok && value == want {
			found = true
		}
		return true
	})
	return found
}

// cDepthWarnings 统计模板深度相关的 warning 诊断。
func cDepthWarnings(diagnostics []air.Diagnostic) int {
	count := 0
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == air.SeverityWarning &&
			strings.Contains(diagnostic.Message, "template nesting depth") {
			count++
		}
	}
	return count
}

// TestMutatePreprocTableChangesBehavior 变异预处理表后，#include 从
// KindDirective 降级为 KindIdent，节点树 ID 随之改变，证明归约非恒等空壳。
func TestMutatePreprocTableChangesBehavior(t *testing.T) {
	source := "#include \"helper.h\"\n#define X 1\n#pragma once\nint v = X;\n"
	normal := &cSpec{
		preprocKeywords:  append([]string(nil), cPreprocKeywords...),
		maxTemplateDepth: cDefaultMaxTemplateDepth,
	}
	mutated := &cSpec{
		preprocKeywords:  cDropKeyword(cPreprocKeywords, "include"),
		maxTemplateDepth: cDefaultMaxTemplateDepth,
	}
	rootA, _ := cSpecParse(t, normal, source)
	rootB, _ := cSpecParse(t, mutated, source)
	countA := cCountKind(rootA, air.KindDirective)
	countB := cCountKind(rootB, air.KindDirective)
	if countA != 3 {
		t.Fatalf("baseline produced %d directives, want 3", countA)
	}
	if countB != 2 {
		t.Fatalf("mutated table produced %d directives, want 2", countB)
	}
	if air.IDOf(rootA) == air.IDOf(rootB) {
		t.Fatal("mutating the preproc table did not change the node tree")
	}
	if !cTreeHasSem(rootA, air.SemDirective, "include") ||
		cTreeHasSem(rootB, air.SemDirective, "include") {
		t.Fatal("SemDirective=include did not follow the table mutation")
	}
	// 全表移除时所有指令都降级。
	empty := &cSpec{preprocKeywords: []string{},
		maxTemplateDepth: cDefaultMaxTemplateDepth}
	rootC, _ := cSpecParse(t, empty, source)
	if cCountKind(rootC, air.KindDirective) != 0 {
		t.Fatal("an empty preproc table must not recognize any directive")
	}
}

// TestMutateTemplateDepthChangesBehavior 变异模板深度上限后，同一输入
// 从「无 warning」变为「有 warning + SemPartial 降级」，证明上限真实生效。
func TestMutateTemplateDepthChangesBehavior(t *testing.T) {
	limit16 := &cSpec{
		preprocKeywords:  append([]string(nil), cPreprocKeywords...),
		maxTemplateDepth: 16,
	}
	limit4 := &cSpec{
		preprocKeywords:  append([]string(nil), cPreprocKeywords...),
		maxTemplateDepth: 4,
	}
	mid := cDeepTemplate(8)
	rootA, diagnosticsA := cSpecParse(t, limit16, mid)
	rootB, diagnosticsB := cSpecParse(t, limit4, mid)
	if cDepthWarnings(diagnosticsA) != 0 {
		t.Fatalf("depth 8 with limit 16 produced warnings: %v", diagnosticsA)
	}
	if cDepthWarnings(diagnosticsB) == 0 {
		t.Fatalf("depth 8 with limit 4 produced no warnings: %v", diagnosticsB)
	}
	if cTreeHasSem(rootA, air.SemPartial, "template-depth") ||
		!cTreeHasSem(rootB, air.SemPartial, "template-depth") {
		t.Fatal("partial reduction did not follow the depth limit mutation")
	}
	// 超深输入下，warning 文本携带各自的上限值，行为确实不同。
	deep := cDeepTemplate(24)
	_, deepA := cSpecParse(t, limit16, deep)
	_, deepB := cSpecParse(t, limit4, deep)
	messageA, messageB := "", ""
	for _, diagnostic := range deepA {
		if strings.Contains(diagnostic.Message, "template nesting depth") {
			messageA = diagnostic.Message
		}
	}
	for _, diagnostic := range deepB {
		if strings.Contains(diagnostic.Message, "template nesting depth") {
			messageB = diagnostic.Message
		}
	}
	if !strings.Contains(messageA, "limit 16") ||
		!strings.Contains(messageB, "limit 4") || messageA == messageB {
		t.Fatalf("limit mutation did not change the warning: %q vs %q",
			messageA, messageB)
	}
}

// TestMutateDialectChangesUnitAnnotation 变异方言后 unit 注解确实改变。
func TestMutateDialectChangesUnitAnnotation(t *testing.T) {
	frontend, err := NewCFrontend()
	if err != nil {
		t.Fatal(err)
	}
	source := "int v = 1;\n"
	makeUnit := func(dialect string) *air.Unit {
		unit, _, parseErr := frontend.Parse(air.SourceUnit{
			Language: "c", Dialect: dialect,
			Files: []air.SourceFile{
				air.NewSourceFile("d.cpp", []byte(source)),
			},
		})
		if parseErr != nil {
			t.Fatal(parseErr)
		}
		return unit
	}
	c11 := makeUnit("c11")
	cpp20 := makeUnit("c++20")
	if c11.Annotations["lang.c.dialect"] != "c11" ||
		cpp20.Annotations["lang.c.dialect"] != "c++20" {
		t.Fatal("dialect normalization did not reach the unit annotations")
	}
	if c11.Fingerprint() == cpp20.Fingerprint() {
		t.Fatal("dialect mutation did not change the fingerprint")
	}
}

// TestMutateMappingTableIsAuthoritative 逐一改动构造样例的期望种类，
// 确认映射表是断言的真实依据而非恒通过常量。
func TestMutateMappingTableIsAuthoritative(t *testing.T) {
	cases := cConstructCases()
	mapper := newCMapper()
	kinds := make(map[air.NodeKind]struct{})
	for _, item := range cases {
		tokens := cSyntheticTokens(item.Tokens)
		got := mapper.match(tokens, 0, len(tokens), item.Top)
		kinds[got.kind] = struct{}{}
	}
	if len(kinds) < 12 {
		t.Fatalf("mapping table only exercises %d distinct NodeKinds", len(kinds))
	}
	// 把样例的 token 序列变异后，映射结果至少有一处发生变化。
	changed := false
	for _, item := range cases {
		original := cSyntheticTokens(item.Tokens)
		mutated := cSyntheticTokens(append([]string{"template"}, item.Tokens...))
		if mapper.match(original, 0, len(original), item.Top).kind !=
			mapper.match(mutated, 0, len(mutated), item.Top).kind {
			changed = true
			break
		}
	}
	if !changed {
		t.Fatal("mutating token streams never changed a construct mapping")
	}
}

// TestCFrontendTagAndLexSpec 校验短标签与 C/C++ 词法表内容。
func TestCFrontendTagAndLexSpec(t *testing.T) {
	spec := &cSpec{preprocKeywords: cPreprocKeywords,
		maxTemplateDepth: cDefaultMaxTemplateDepth}
	if spec.Tag() != "c" {
		t.Fatalf("Tag = %q, want c", spec.Tag())
	}
	lex := spec.LexSpec()
	for _, keyword := range []string{"namespace", "template", "sizeof", "new", "class"} {
		found := false
		for _, got := range lex.Keywords() {
			if got == keyword {
				found = true
			}
		}
		if !found {
			t.Fatalf("keyword %q missing from LexSpec", keyword)
		}
	}
	if lex.LineComment() != "//" {
		t.Fatal("wrong line comment")
	}
	opening, closing := lex.BlockComment()
	if opening != "/*" || closing != "*/" {
		t.Fatal("wrong block comment")
	}
	if lex.IndentSensitive() {
		t.Fatal("C/C++ must not be indent sensitive")
	}
	if !strings.Contains(strings.Join(lex.RawStringPrefixes(), "|"), "R") {
		t.Fatal("raw string prefix R missing")
	}
	if len(lex.StringQuotes()) != 2 || len(lex.NumberSeparators()) == 0 {
		t.Fatal("string quotes / number separators missing")
	}
	if len(lex.Operators()) < 20 {
		t.Fatal("operator table is too small")
	}
	// 每个方法都返回独立副本，调用方改写不污染词法表。
	keywords := lex.Keywords()
	keywords[0] = "mutated"
	if lex.Keywords()[0] == "mutated" {
		t.Fatal("Keywords() leaked a shared slice")
	}
}

// TestNormalizeDialect 校验三方言归一、别名与未知方言报错。
func TestNormalizeDialect(t *testing.T) {
	spec := &cSpec{}
	cases := []struct {
		in, want string
	}{
		{"", "c11"}, {"C11", "c11"}, {"c99", "c11"}, {"gnu11", "c11"},
		{"C++17", "c++17"}, {"cpp17", "c++17"}, {"c++1z", "c++17"},
		{"C++20", "c++20"}, {"gnu++20", "c++20"}, {"c++2a", "c++20"},
	}
	for _, item := range cases {
		got, err := spec.NormalizeDialect(item.in)
		if err != nil || got != item.want {
			t.Fatalf("NormalizeDialect(%q) = %q, %v; want %q",
				item.in, got, err, item.want)
		}
	}
	if _, err := spec.NormalizeDialect("rust2021"); err == nil {
		t.Fatal("unknown dialect must be rejected")
	}
	for _, canonical := range cDialectCanonical {
		got, err := spec.NormalizeDialect(canonical)
		if err != nil || got != canonical {
			t.Fatalf("canonical dialect %q did not round-trip", canonical)
		}
	}
}
