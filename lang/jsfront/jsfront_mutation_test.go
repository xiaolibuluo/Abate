/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// jsFingerprint 返回一次解析的单元指纹。
func jsFingerprint(t *testing.T, fe air.Frontend, path, src string) [32]byte {
	t.Helper()
	out, _ := jsParse(t, fe, path, "", src)
	return out.Fingerprint()
}

// jsRuleIndex 在声明映射表中查找指定文本的下标。
func jsRuleIndex(text string) int {
	for i := range jsDeclRules {
		if jsDeclRules[i].text == text {
			return i
		}
	}
	return -1
}

// jsInfixIndex 在中缀绑定表中查找指定运算符的下标。
func jsInfixIndex(text string) int {
	for i := range jsInfixRules {
		if jsInfixRules[i].text == text {
			return i
		}
	}
	return -1
}

// jsWithoutString 返回删除指定字符串后的确定性切片副本。
func jsWithoutString(in []string, drop string) []string {
	out := make([]string, 0, len(in))
	for _, s := range in {
		if s != drop {
			out = append(out, s)
		}
	}
	return out
}

// TestMutationContinuationTable 变异 ASI 续接表后输出确实改变。
func TestMutationContinuationTable(t *testing.T) {
	fe := jsFrontend(t)
	src, path := "a\n.b", "asi.js"
	before := jsFingerprint(t, fe, path, src)
	original := jsContinuationTokens
	mutated := jsWithoutString(jsWithoutString(original, "."), "?.")
	jsContinuationTokens = mutated
	defer func() { jsContinuationTokens = original }()
	after := jsFingerprint(t, fe, path, src)
	if after == before {
		t.Fatal("ASI continuation table mutation did not change the tree")
	}
}

// TestMutationRestrictedTable 变异 ASI 受限产生式表后输出确实改变。
func TestMutationRestrictedTable(t *testing.T) {
	fe := jsFrontend(t)
	src, path := "return\n+1", "asi.js"
	before := jsFingerprint(t, fe, path, src)
	original := jsRestrictedWords
	jsRestrictedWords = jsWithoutString(original, "return")
	defer func() { jsRestrictedWords = original }()
	after := jsFingerprint(t, fe, path, src)
	if after == before {
		t.Fatal("restricted production mutation did not change the tree")
	}
}

// TestMutationInfixBindingTable 变异 "+" 绑定力后表达式树结构改变。
func TestMutationInfixBindingTable(t *testing.T) {
	fe := jsFrontend(t)
	src, path := "a + b * c", "bind.js"
	before := jsFingerprint(t, fe, path, src)
	idx := jsInfixIndex("+")
	if idx < 0 {
		t.Fatal("missing + binding rule")
	}
	original := jsInfixRules[idx].power
	jsInfixRules[idx].power = 14
	defer func() { jsInfixRules[idx].power = original }()
	after := jsFingerprint(t, fe, path, src)
	if after == before {
		t.Fatal("binding power mutation did not change the tree")
	}
}

// TestMutationDeclMappingTable 变异 const 声明映射后 AIR 节点类改变。
func TestMutationDeclMappingTable(t *testing.T) {
	fe := jsFrontend(t)
	src, path := "const x = 1;", "decl.ts"
	before := factsOf(mustParseUnit(t, fe, path, src))
	if !before["kind:ConstDecl"] {
		t.Fatal("baseline lacks KindConstDecl")
	}
	idx := jsRuleIndex("const")
	if idx < 0 {
		t.Fatal("missing const declaration rule")
	}
	original := jsDeclRules[idx].kind
	jsDeclRules[idx].kind = air.KindVarDecl
	defer func() { jsDeclRules[idx].kind = original }()
	after := factsOf(mustParseUnit(t, fe, path, src))
	if after["kind:ConstDecl"] || !after["kind:VarDecl"] {
		t.Fatalf("declaration kind did not change: %v", sortedFactKeys(after))
	}
}

// TestMutationNoteMappingTable 变异 import 的 SemNote 后注记确实改变。
func TestMutationNoteMappingTable(t *testing.T) {
	fe := jsFrontend(t)
	src, path := `import x from "y";`, "note.js"
	before := factsOf(mustParseUnit(t, fe, path, src))
	if !before["note:esm-import"] {
		t.Fatal("baseline lacks esm-import note")
	}
	idx := jsRuleIndex("import")
	original := jsDeclRules[idx].note
	jsDeclRules[idx].note = "mutated-import"
	defer func() { jsDeclRules[idx].note = original }()
	after := factsOf(mustParseUnit(t, fe, path, src))
	if after["note:esm-import"] || !after["note:mutated-import"] {
		t.Fatalf("SemNote mutation did not take effect: %v", sortedFactKeys(after))
	}
}

// TestMutationKeywordTable 变异保留字表后真实 token 流类别改变。
func TestMutationKeywordTable(t *testing.T) {
	src := "const x = 1;"
	before := jsLexTokens(t, src)
	original := jsKeywords
	jsKeywords = jsWithoutString(original, "const")
	defer func() { jsKeywords = original }()
	after := jsLexTokens(t, src)
	if len(before) != len(after) {
		t.Fatalf("token count changed: %d vs %d", len(before), len(after))
	}
	changed := false
	for i := range before {
		if before[i].Kind != after[i].Kind {
			changed = true
		}
	}
	if !changed {
		t.Fatal("keyword table mutation did not change token kinds")
	}
	if before[0].Kind != frontkit.TokenKeyword ||
		after[0].Kind != frontkit.TokenIdent {
		t.Fatalf("unexpected token kinds: %s vs %s", before[0].Kind, after[0].Kind)
	}
}

// TestMutationOperatorTable 变异运算符表后最长匹配结果改变。
func TestMutationOperatorTable(t *testing.T) {
	src := "a?.b"
	before := jsLexTokens(t, src)
	original := jsOperators
	jsOperators = jsWithoutString(original, "?.")
	defer func() { jsOperators = original }()
	after := jsLexTokens(t, src)
	if len(before) == len(after) {
		t.Fatal("operator table mutation did not change tokenization")
	}
}

// TestDifferentSourcesDifferentFingerprints 证明不同输入产生不同指纹，非空恒通过。
func TestDifferentSourcesDifferentFingerprints(t *testing.T) {
	fe := jsFrontend(t)
	sources := []string{
		"const a = 1;",
		"const a = 2;",
		"let a = 1;",
		"function a() { return 1; }",
		"class A {}",
		"interface A { x: number }",
		"const el = <div/>;",
		"// only a comment",
	}
	seen := make(map[[32]byte]bool)
	for i, src := range sources {
		path := "case.js"
		if i == 5 || i == 6 {
			path = "case.tsx"
		}
		seen[jsFingerprint(t, fe, path, src)] = true
	}
	if len(seen) != len(sources) {
		t.Fatalf("expected %d distinct fingerprints, got %d",
			len(sources), len(seen))
	}
}

// mustParseUnit 解析并返回单元，失败即终止。
func mustParseUnit(t *testing.T, fe air.Frontend, path, src string) *air.Unit {
	t.Helper()
	out, _ := jsParse(t, fe, path, "", src)
	return out
}
