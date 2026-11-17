/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"testing"

	"abate/lang/air"
)

// csCountKind 统计一棵树里指定种类的节点数量。
func csCountKind(root air.Node, kind air.NodeKind) int {
	count := 0
	air.Walk(root, func(node air.Node) bool {
		if node.Kind() == kind {
			count++
		}
		return true
	})
	return count
}

// TestCsLinqQueryClauses 证明 LINQ 查询被切成 from/where/orderby/select 子句层。
func TestCsLinqQueryClauses(t *testing.T) {
	src := "var q = from x in xs where x > 0 orderby x select x;\n"
	out, root, diags := csParse(t, src)
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if !csHasNote(root, "linq") {
		t.Fatal("query is missing the linq SemNote")
	}
	for _, note := range []string{
		"linq-clause:from", "linq-clause:where",
		"linq-clause:orderby", "linq-clause:select",
	} {
		if !csHasNote(root, note) {
			t.Fatalf("query is missing clause %q", note)
		}
	}
	if got := csCountKind(root, air.KindCall); got < 4 {
		t.Fatalf("only %d KindCall nodes for a four-clause query", got)
	}
	if out.NodeCount() < 10 {
		t.Fatalf("linq reduction produced only %d nodes", out.NodeCount())
	}
}

// TestCsLinqGroupJoinLet 覆盖 group/join/let/into 子句的降级归约。
func TestCsLinqGroupJoinLet(t *testing.T) {
	src := "var q = from a in xs " +
		"join b in ys on a equals b " +
		"let z = a + b " +
		"group z by z into g " +
		"select g;\n"
	_, root, _ := csParse(t, src)
	for _, note := range []string{
		"linq-clause:from", "linq-clause:join", "linq-clause:let",
		"linq-clause:group", "linq-clause:into", "linq-clause:select",
	} {
		if !csHasNote(root, note) {
			t.Fatalf("query is missing clause %q", note)
		}
	}
	if !csHasKind(root, air.KindCall) {
		t.Fatal("LINQ clauses must lower to KindCall")
	}
}

// TestCsLinqNestedExpression 断言 LINQ 里的表达式仍由 Pratt 结构化归约。
func TestCsLinqNestedExpression(t *testing.T) {
	src := "var q = from x in xs where x > 0 && x < 10 select x + 1;\n"
	out, root, _ := csParse(t, src)
	if !csHasKind(root, air.KindBinary) {
		t.Fatal("query predicates were not reduced to KindBinary")
	}
	if !csHasSem(root, air.SemOperator, ">") {
		t.Fatal("missing SemOperator for the LINQ predicate")
	}
	if err := air.CheckIntervals(root, len(src)); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
	if out.NodeCount() < 12 {
		t.Fatalf("only %d nodes in the layered LINQ tree", out.NodeCount())
	}
}

// TestCsAttributeAndIndexer 专项验证特性标记与索引器声明。
func TestCsAttributeAndIndexer(t *testing.T) {
	src := "[Obsolete(\"old\")]\n" +
		"public class C {\n" +
		"  [IndexerName(\"Item\")]\n" +
		"  public int this[int i] { get { return i; } }\n" +
		"}\n"
	out, root, _ := csParse(t, src)
	if !csHasNote(root, "attribute") || csCountKind(root, air.KindAttribute) < 2 {
		t.Fatal("both attribute markers must lower to KindAttribute")
	}
	if !csHasNote(root, "indexer") || !csHasKind(root, air.KindIndex) {
		t.Fatal("indexer declaration must lower to KindIndex with a note")
	}
	if !csHasNote(root, "property-accessors") {
		t.Fatal("indexer accessor block is missing its note")
	}
	if out.NodeCount() < 15 {
		t.Fatalf("only %d nodes in the attribute/indexer tree", out.NodeCount())
	}
}

// TestCsPropertyAccessors 专项验证自动属性与表达式体属性。
func TestCsPropertyAccessors(t *testing.T) {
	src := "public class C {\n" +
		"  public string? Name { get; set; }\n" +
		"  public int Size => items.Count;\n" +
		"}\n"
	_, root, _ := csParse(t, src)
	if !csHasNote(root, "property") {
		t.Fatal("auto-property must lower to KindFuncDecl with a property note")
	}
	if !csHasKind(root, air.KindFuncDecl) {
		t.Fatal("missing KindFuncDecl for the property")
	}
	if !csHasKind(root, air.KindLambda) {
		t.Fatal("expression-bodied property must produce a KindLambda")
	}
	if !csHasSem(root, air.SemNote, "nullable") {
		t.Fatal("nullable reference property must be annotated")
	}
}

// TestCsNullConditionalAndCoalescing 专项验证空条件与 null 合并操作符。
func TestCsNullConditionalAndCoalescing(t *testing.T) {
	src := "class C {\n" +
		"  object? M(A? a, object b) {\n" +
		"    a?.Reset();\n" +
		"    var v = a?.Name ?? b ?? fallback;\n" +
		"    return v;\n" +
		"  }\n" +
		"}\n"
	_, root, _ := csParse(t, src)
	if csCountKind(root, air.KindField) < 2 ||
		!csHasNote(root, "null-conditional") {
		t.Fatal("null-conditional member access is missing")
	}
	if csCountKind(root, air.KindOptional) < 2 ||
		!csHasNote(root, "nullable") {
		t.Fatal("nullable reference annotations are missing")
	}
	if !csHasNote(root, "null-coalescing") ||
		!csHasSem(root, air.SemOperator, "??") {
		t.Fatal("null-coalescing operator is missing")
	}
}

// TestCsVerbatimAndInterpolatedStrings 专项验证 C# 字符串字面量形态。
func TestCsVerbatimAndInterpolatedStrings(t *testing.T) {
	src := "class C {\n" +
		"  string a = @\"C:\\temp\\f.txt\";\n" +
		"  string b = $\"hello {name}, {count}\";\n" +
		"  string c = $@\"raw {name}\";\n" +
		"}\n"
	_, root, _ := csParse(t, src)
	for _, note := range []string{
		"verbatim", "interpolated", "verbatim-interpolated", "interpolation",
	} {
		if !csHasNote(root, note) {
			t.Fatalf("missing string flavor note %q", note)
		}
	}
	if csCountKind(root, air.KindStrLit) < 3 {
		t.Fatal("all three string flavors must project to KindStrLit")
	}
	if !csHasSem(root, air.SemTag, "verbatim") {
		t.Fatal("verbatim string is missing its SemTag")
	}
}
