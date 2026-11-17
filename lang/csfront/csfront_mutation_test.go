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

// csDropString 返回移除了指定条目的有序表副本。
func csDropString(table []string, drop string) []string {
	out := make([]string, 0, len(table))
	for i := range table {
		if table[i] != drop {
			out = append(out, table[i])
		}
	}
	return out
}

// csDropBinding 返回移除了指定操作符的绑定表副本。
func csDropBinding(table []csBinding, drop string) []csBinding {
	out := make([]csBinding, 0, len(table))
	for i := range table {
		if table[i].text != drop {
			out = append(out, table[i])
		}
	}
	return out
}

// csReduceWith 用指定配置直接归约 token 流并返回 AIR 树哈希。
func csReduceWith(t *testing.T, tokens []frontkit.Token, spec *csSpec) [32]byte {
	t.Helper()
	root, _, err := spec.ParseTokens(tokens, "mut.cs")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if root == nil {
		t.Fatal("ParseTokens returned a nil root")
	}
	return determinism.HashState(root)
}

// TestCsMutationLinqKeywordTable 证明删除 select 后 LINQ 归约行为确实改变。
func TestCsMutationLinqKeywordTable(t *testing.T) {
	src := "var q = from x in xs select x;\n"
	tokens := csTokenize(t, src)
	base := csReduceWith(t, tokens, &csSpec{})
	mutated := &csSpec{linqKeywords: csDropString(csLinqKeywords, "select")}
	changed := csReduceWith(t, tokens, mutated)
	if base == changed {
		t.Fatal("removing the select LINQ keyword left the AIR unchanged")
	}
	if _, root, _ := csParse(t, src); !csHasNote(root, "linq") {
		t.Fatal("baseline query lost its linq note")
	}
}

// TestCsMutationLinqClauseTable 证明子句表驱动 parseLinq 的切分。
func TestCsMutationLinqClauseTable(t *testing.T) {
	src := "var q = from x in xs where x > 0 select x;\n"
	tokens := csTokenize(t, src)
	base := csReduceWith(t, tokens, &csSpec{})
	mutated := &csSpec{linqClauses: csDropString(csLinqClauseKeywords, "where")}
	changed := csReduceWith(t, tokens, mutated)
	if base == changed {
		t.Fatal("removing the where clause keyword left the AIR unchanged")
	}
	root, _, err := mutated.ParseTokens(tokens, "mut.cs")
	if err != nil {
		t.Fatal(err)
	}
	if csHasNote(root, "linq-clause:where") {
		t.Fatal("mutated clause table still produced a where clause node")
	}
}

// TestCsMutationNoteTable 证明语义标注来自可注入表而非硬编码常量。
func TestCsMutationNoteTable(t *testing.T) {
	src := "var q = from x in xs select x;\n"
	tokens := csTokenize(t, src)
	mutated := make([]csNote, len(csNotes))
	copy(mutated, csNotes)
	for i := range mutated {
		if mutated[i].name == "linq" {
			mutated[i].value = "MUTATED-LINQ-NOTE"
		}
	}
	root, _, err := (&csSpec{notes: mutated}).ParseTokens(tokens, "mut.cs")
	if err != nil {
		t.Fatal(err)
	}
	if !csHasNote(root, "MUTATED-LINQ-NOTE") {
		t.Fatal("mutated note value is absent from AIR")
	}
	if csHasNote(root, "linq") {
		t.Fatal("original note value survived the mutation")
	}
	base := csReduceWith(t, tokens, &csSpec{})
	if base == determinism.HashState(root) {
		t.Fatal("note mutation left the AIR hash unchanged")
	}
}

// TestCsMutationInfixBindings 证明删除一条绑定力会改变表达式树。
func TestCsMutationInfixBindings(t *testing.T) {
	src := "class C { object M(object a, int b, int c, int d) " +
		"{ return a ?? b + c * d; } }\n"
	tokens := csTokenize(t, src)
	base := csReduceWith(t, tokens, &csSpec{})
	mutated := csDropBinding(csInfixBindings, "??")
	if base == csReduceWith(t, tokens, &csSpec{infix: mutated}) {
		t.Fatal("removing the ?? binding left the AIR unchanged")
	}
	shifted := make([]csBinding, 0, len(csInfixBindings))
	for _, entry := range csInfixBindings {
		entry.power = 1
		shifted = append(shifted, entry)
	}
	if base == csReduceWith(t, tokens, &csSpec{infix: shifted}) {
		t.Fatal("changing every binding power left the AIR unchanged")
	}
}

// TestCsMutationKeywordTable 证明关键字表改变 token 种类并传导到 AIR。
func TestCsMutationKeywordTable(t *testing.T) {
	src := "class C { public int X { get; set; } }\n"
	mutSpec := &csSpec{keywords: csDropString(csKeywords, "class")}
	baseTokens := csTokenize(t, src)
	mutTokens := csTokenizeWith(t, mutSpec.LexSpec(), src)
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
	base := csReduceWith(t, baseTokens, &csSpec{})
	if base == csReduceWith(t, mutTokens, &csSpec{}) {
		t.Fatal("keyword table mutation left the AIR unchanged")
	}
}

// TestCsMutationDialectTable 证明方言接受集合是表驱动的。
func TestCsMutationDialectTable(t *testing.T) {
	if !csContains(csDialects, "cs8") || !csContains(csDialects, "cs12") {
		t.Fatal("canonical dialects absent from table")
	}
	if csContains(csDialects, "cs7") {
		t.Fatal("unsupported dialect accepted by table")
	}
	mutated := []string{"cs8"}
	if !csContains(mutated, "cs8") || csContains(mutated, "cs12") {
		t.Fatal("mutated dialect table behaves incorrectly")
	}
	if _, err := (&csSpec{}).NormalizeDialect("net7"); err != nil {
		t.Fatalf("net7 must be accepted: %v", err)
	}
	if _, err := (&csSpec{}).NormalizeDialect("cs7"); err == nil {
		t.Fatal("non-canonical dialect accepted")
	}
}

// TestCsTableLookups 覆盖所有有序表查找的命中与未命中分支。
func TestCsTableLookups(t *testing.T) {
	if !csContains(csKeywords, "class") || csContains(csKeywords, "nope") {
		t.Fatal("csContains broken")
	}
	if csNoteFor(csNotes, "linq") != "linq" || csNoteFor(csNotes, "nope") != "" {
		t.Fatal("csNoteFor broken")
	}
	if csInfixPower(csInfixBindings, "??") != 15 ||
		csInfixPower(csInfixBindings, ",") != 3 ||
		csInfixPower(csInfixBindings, "nope") != 0 {
		t.Fatal("csInfixPower broken")
	}
	if !csInfixRight(csInfixBindings, "=") ||
		csInfixRight(csInfixBindings, "+") || csInfixRight(csInfixBindings, "nope") {
		t.Fatal("csInfixRight broken")
	}
	if rule, ok := csDeclRuleFor("namespace"); !ok || rule.kind != air.KindModuleDecl {
		t.Fatal("csDeclRuleFor broken")
	}
	if _, ok := csDeclRuleFor("nope"); ok {
		t.Fatal("csDeclRuleFor accepted unknown keyword")
	}
	if got, ok := csDialectFor("net8"); !ok || got != "cs12" {
		t.Fatal("csDialectFor broken")
	}
	if _, ok := csDialectFor("nope"); ok {
		t.Fatal("csDialectFor accepted unknown dialect")
	}
}

// TestCsTokenLookups covers the token inspection helpers.
func TestCsTokenLookups(t *testing.T) {
	tokens := csTokenize(t, "class C { }\n")
	if !csHasText(tokens, 0, len(tokens), "class") ||
		csHasText(tokens, 0, len(tokens), "nope") {
		t.Fatal("csHasText broken")
	}
	if csLeadIndex(tokens, 0, 1) != 0 {
		t.Fatal("csLeadIndex broken")
	}
	if csNameOf(tokens, 0, len(tokens)) != "C" {
		t.Fatal("csNameOf broken")
	}
	if !csHasLinqTail(csTokenize(t, "select x;\n"), 0, 2, csLinqKeywords) ||
		csHasLinqTail(tokens, 0, len(tokens), csLinqKeywords) {
		t.Fatal("csHasLinqTail broken")
	}
	if start, end := csSpanTokens(tokens, 5, 2); start.Offset != 0 || end.Offset != 0 {
		t.Fatal("csSpanTokens empty interval broken")
	}
}

// TestCsStringAndDirectiveLookups covers string flavors, directives and doc
// comments.
func TestCsStringAndDirectiveLookups(t *testing.T) {
	for _, tc := range []struct{ text, want string }{
		{`@"a"`, "verbatim"}, {`$"a"`, "interpolated"},
		{`$@"a"`, "verbatim-interpolated"}, {`@$"a"`, "verbatim-interpolated"},
		{`"a"`, "plain"},
	} {
		if got := csStringFlavor(tc.text); got != tc.want {
			t.Fatalf("csStringFlavor(%q)=%q, want %q", tc.text, got, tc.want)
		}
	}
	if csDirectiveName([]string{"#", "region", "X"}) != "region" ||
		csDirectiveName([]string{"#", ""}) != "" {
		t.Fatal("csDirectiveName broken")
	}
	if !csIsDocComment("/// doc") || csIsDocComment("// doc") {
		t.Fatal("csIsDocComment broken")
	}
}

// TestCsMapperEdgeCases 覆盖节点构造器的空值与表外标注分支。
func TestCsMapperEdgeCases(t *testing.T) {
	mapper := newCsMap(nil, nil)
	if mapper.note("unknown-note") != "unknown-note" {
		t.Fatal("note fallback broken")
	}
	mapper.apply(nil, csMapping{kind: air.KindVarDecl, note: "linq"})
	tok := frontkit.Token{
		Kind: frontkit.TokenIdent, Text: "x", Offset: 0, EndOffset: 1,
		Line: 1, Column: 1,
	}
	if node := mapper.infixNode(air.KindBinary, nil, tok, mapper.literal(tok)); node.Kind() != air.KindBinary {
		t.Fatal("infixNode nil-left fallback broken")
	}
	if node := mapper.unaryNode(air.KindAwait, tok, nil); node.Kind() != air.KindAwait {
		t.Fatal("unaryNode nil-operand fallback broken")
	}
	// 词法 token 必须产出冻结投影叶子；标点/运算符不再贴伪 Kind 标签，
	// 其拼写由 BindTokenSpellings 在父节点负载中绑定。
	for _, kind := range []frontkit.TokenKind{
		frontkit.TokenIdent, frontkit.TokenKeyword, frontkit.TokenNum,
		frontkit.TokenStr, frontkit.TokenChar, frontkit.TokenComment,
		frontkit.TokenIllegal,
	} {
		probe := tok
		probe.Kind = kind
		if mapper.literal(probe) == nil {
			t.Fatalf("literal returned nil for %s", kind)
		}
	}
	for _, kind := range []frontkit.TokenKind{
		frontkit.TokenOp, frontkit.TokenPunct,
	} {
		probe := tok
		probe.Kind = kind
		if mapper.literal(probe) != nil {
			t.Fatalf("literal fabricated a node for %s", kind)
		}
	}
	if mapper.comment(frontkit.Token{
		Kind: frontkit.TokenComment, Text: "/* block */",
		Offset: 0, EndOffset: 11, Line: 1, Column: 1,
	}).Kind() != air.KindComment {
		t.Fatal("block comment lowering broken")
	}
}
