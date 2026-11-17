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

// ktDropString 返回移除了指定条目的有序表副本。
func ktDropString(table []string, drop string) []string {
	out := make([]string, 0, len(table))
	for i := range table {
		if table[i] != drop {
			out = append(out, table[i])
		}
	}
	return out
}

// ktDropBinding 返回移除了指定操作符的绑定表副本。
func ktDropBinding(table []ktBinding, drop string) []ktBinding {
	out := make([]ktBinding, 0, len(table))
	for i := range table {
		if table[i].text != drop {
			out = append(out, table[i])
		}
	}
	return out
}

// ktDropNullSafe 返回移除了指定操作符的空安全映射表副本。
func ktDropNullSafe(table []ktNullSafeRule, drop string) []ktNullSafeRule {
	out := make([]ktNullSafeRule, 0, len(table))
	for i := range table {
		if table[i].text != drop {
			out = append(out, table[i])
		}
	}
	return out
}

// ktReduceWith 用指定配置直接归约 token 流并返回 AIR 树哈希。
func ktReduceWith(t *testing.T, tokens []frontkit.Token, spec *ktSpec) [32]byte {
	t.Helper()
	root, _, err := spec.ParseTokens(tokens, "mut.kt")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if root == nil {
		t.Fatal("ParseTokens returned a nil root")
	}
	return determinism.HashState(root)
}

// TestKtMutationNullSafeTable 证明删除 ?. 后空安全映射行为确实改变。
func TestKtMutationNullSafeTable(t *testing.T) {
	src := "val n = a?.len ?: 0\n"
	tokens := ktTokenize(t, src)
	base := ktReduceWith(t, tokens, &ktSpec{})
	mutated := &ktSpec{nullSafe: ktDropNullSafe(ktNullSafeRules, "?.")}
	changed := ktReduceWith(t, tokens, mutated)
	if base == changed {
		t.Fatal("removing the ?. null-safety rule left the AIR unchanged")
	}
	if _, root, _ := ktParse(t, src); !ktHasNote(root, "null-safe") {
		t.Fatal("baseline expression lost its null-safe note")
	}
	mutRoot, _, err := mutated.ParseTokens(tokens, "mut.kt")
	if err != nil {
		t.Fatal(err)
	}
	if ktHasNote(mutRoot, "null-safe") {
		t.Fatal("mutated null-safety table still produced a null-safe note")
	}
}

// TestKtMutationNullSafeElvis 证明 Elvis 映射条目同样表驱动。
func TestKtMutationNullSafeElvis(t *testing.T) {
	src := "val n = a ?: b\n"
	tokens := ktTokenize(t, src)
	base := ktReduceWith(t, tokens, &ktSpec{})
	mutated := &ktSpec{nullSafe: ktDropNullSafe(ktNullSafeRules, "?:")}
	if base == ktReduceWith(t, tokens, mutated) {
		t.Fatal("removing the ?: elvis rule left the AIR unchanged")
	}
}

// TestKtMutationNullSafeKind 证明把 !!. 映射种类改写会改变 AIR 结构。
func TestKtMutationNullSafeKind(t *testing.T) {
	src := "val n = a!!\n"
	tokens := ktTokenize(t, src)
	base := ktReduceWith(t, tokens, &ktSpec{})
	shifted := make([]ktNullSafeRule, len(ktNullSafeRules))
	copy(shifted, ktNullSafeRules)
	for i := range shifted {
		if shifted[i].text == "!!" {
			shifted[i].kind = air.KindNumLit
		}
	}
	if base == ktReduceWith(t, tokens, &ktSpec{nullSafe: shifted}) {
		t.Fatal("changing the !!. mapping kind left the AIR unchanged")
	}
}

// TestKtMutationNoteTable 证明语义标注来自可注入表而非硬编码常量。
func TestKtMutationNoteTable(t *testing.T) {
	src := "val n = a?.len\n"
	tokens := ktTokenize(t, src)
	mutated := make([]ktNote, len(ktNotes))
	copy(mutated, ktNotes)
	for i := range mutated {
		if mutated[i].name == "null-safe" {
			mutated[i].value = "MUTATED-NULL-SAFE"
		}
	}
	root, _, err := (&ktSpec{notes: mutated}).ParseTokens(tokens, "mut.kt")
	if err != nil {
		t.Fatal(err)
	}
	if !ktHasNote(root, "MUTATED-NULL-SAFE") {
		t.Fatal("mutated note value is absent from AIR")
	}
	if ktHasNote(root, "null-safe") {
		t.Fatal("original note value survived the mutation")
	}
	base := ktReduceWith(t, tokens, &ktSpec{})
	if base == determinism.HashState(root) {
		t.Fatal("note mutation left the AIR hash unchanged")
	}
}

// TestKtMutationInfixBindings 证明删除一条绑定力会改变表达式树。
func TestKtMutationInfixBindings(t *testing.T) {
	src := "fun f(a: Int, b: Int, c: Int): Int { return a ?: b + c * 2 }\n"
	tokens := ktTokenize(t, src)
	base := ktReduceWith(t, tokens, &ktSpec{})
	mutated := ktDropBinding(ktInfixBindings, "?:")
	if base == ktReduceWith(t, tokens, &ktSpec{infix: mutated}) {
		t.Fatal("removing the ?: binding left the AIR unchanged")
	}
	shifted := make([]ktBinding, 0, len(ktInfixBindings))
	for _, entry := range ktInfixBindings {
		entry.power = 1
		shifted = append(shifted, entry)
	}
	if base == ktReduceWith(t, tokens, &ktSpec{infix: shifted}) {
		t.Fatal("changing every binding power left the AIR unchanged")
	}
}

// TestKtMutationKeywordTable 证明关键字表改变 token 种类并传导到 AIR。
func TestKtMutationKeywordTable(t *testing.T) {
	src := "class C { fun f() { } }\n"
	mutSpec := &ktSpec{keywords: ktDropString(ktKeywords, "class")}
	baseTokens := ktTokenize(t, src)
	mutTokens := ktTokenizeWith(t, mutSpec.LexSpec(), src)
	baseKinds, mutIdents := 0, 0
	for i := range baseTokens {
		if baseTokens[i].Text == "class" && baseTokens[i].Kind == frontkit.TokenKeyword {
			baseKinds++
		}
		if i < len(mutTokens) && mutTokens[i].Text == "class" &&
			mutTokens[i].Kind == frontkit.TokenIdent {
			mutIdents++
		}
	}
	if baseKinds != 1 || mutIdents != 1 {
		t.Fatalf("keyword mutation did not change token kinds: %d/%d",
			baseKinds, mutIdents)
	}
	base := ktReduceWith(t, baseTokens, &ktSpec{})
	if base == ktReduceWith(t, mutTokens, &ktSpec{}) {
		t.Fatal("keyword table mutation left the AIR unchanged")
	}
}

// TestKtMutationWhenWords 证明 when 分支切分由可注入表驱动。
func TestKtMutationWhenWords(t *testing.T) {
	src := "when (x) { 1 -> a() 2 -> b() else -> c() }\n"
	tokens := ktTokenize(t, src)
	base := ktReduceWith(t, tokens, &ktSpec{})
	mutated := &ktSpec{whenWords: ktDropString(ktWhenBranchWords, "->")}
	if base == ktReduceWith(t, tokens, mutated) {
		t.Fatal("removing the -> branch marker left the AIR unchanged")
	}
	mutRoot, _, err := mutated.ParseTokens(tokens, "mut.kt")
	if err != nil {
		t.Fatal(err)
	}
	if ktCountKind(mutRoot, air.KindCase) >= ktCountKind(
		mustParseRoot(t, tokens), air.KindCase) {
		t.Fatal("mutated when table did not reduce the branch count")
	}
}

// mustParseRoot 用默认配置归约 token 流并返回根节点。
func mustParseRoot(t *testing.T, tokens []frontkit.Token) air.Node {
	t.Helper()
	root, _, err := (&ktSpec{}).ParseTokens(tokens, "base.kt")
	if err != nil {
		t.Fatal(err)
	}
	return root
}

// TestKtMutationDialectTable 证明方言接受集合是表驱动的。
func TestKtMutationDialectTable(t *testing.T) {
	if !ktContains(ktDialects, "kt1.6") || !ktContains(ktDialects, "kt2.1") {
		t.Fatal("canonical dialects absent from table")
	}
	if ktContains(ktDialects, "kt3.0") {
		t.Fatal("unsupported dialect accepted by table")
	}
	mutated := []string{"kt1.6"}
	if !ktContains(mutated, "kt1.6") || ktContains(mutated, "kt2.1") {
		t.Fatal("mutated dialect table behaves incorrectly")
	}
	if _, err := (&ktSpec{}).NormalizeDialect("2.0"); err != nil {
		t.Fatalf("2.0 must be accepted: %v", err)
	}
	if _, err := (&ktSpec{}).NormalizeDialect("kt3.0"); err == nil {
		t.Fatal("non-canonical dialect accepted")
	}
}

// TestKtDialectNormalization 检查别名到规范方言的确定性归一。
func TestKtDialectNormalization(t *testing.T) {
	good := map[string]string{
		"": "kt2.1", "kt": "kt2.1", "kotlin": "kt2.1", "KOTLIN": "kt2.1",
		"1.9": "kt1.9", "2.0": "kt2.0", "kt1.6": "kt1.6", "latest": "kt2.1",
		"  2.1  ": "kt2.1",
	}
	for in, want := range good {
		got, err := (&ktSpec{}).NormalizeDialect(in)
		if err != nil || got != want {
			t.Fatalf("NormalizeDialect(%q)=%q,%v want %q", in, got, err, want)
		}
	}
	for _, in := range []string{"java", "kt3.0", "go1.26", "9.9"} {
		if got, err := (&ktSpec{}).NormalizeDialect(in); err == nil {
			t.Fatalf("NormalizeDialect(%q)=%q, want error", in, got)
		}
	}
}
