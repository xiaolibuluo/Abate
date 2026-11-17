/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"testing"

	"abate/lang/air"
)

// TestKtWhenStatementBranches 证明 when 语句被切成 KindCase 分支层。
func TestKtWhenStatementBranches(t *testing.T) {
	src := "fun classify(x: Int): String {\n" +
		"  when (x) {\n" +
		"    1 -> println(\"one\")\n" +
		"    in 2..10 -> println(\"few\")\n" +
		"    is Int -> println(\"int\")\n" +
		"    else -> println(\"other\")\n" +
		"  }\n" +
		"  return \"done\"\n" +
		"}\n"
	out, root, diags := ktParse(t, src)
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if !ktHasNote(root, "when") {
		t.Fatal("when statement is missing the when SemNote")
	}
	if got := ktCountKind(root, air.KindCase); got < 4 {
		t.Fatalf("only %d KindCase nodes for a four-branch when", got)
	}
	if !ktHasNote(root, "when-branch") {
		t.Fatal("when branches are missing their SemNote")
	}
	if out.NodeCount() < 20 {
		t.Fatalf("when reduction produced only %d nodes", out.NodeCount())
	}
}

// TestKtWhenExpression 断言作为表达式的 when 归约为 KindMatchExpr。
func TestKtWhenExpression(t *testing.T) {
	src := "fun tone(x: Int): String {\n" +
		"  val r = when (x) { 1 -> \"a\" 2 -> \"b\" else -> \"z\" }\n" +
		"  return r\n" +
		"}\n"
	_, root, _ := ktParse(t, src)
	if !ktHasKind(root, air.KindMatchExpr) {
		t.Fatal("expression when did not lower to KindMatchExpr")
	}
	if !ktHasNote(root, "when-expr") {
		t.Fatal("expression when is missing the when-expr SemNote")
	}
}

// TestKtNestedWhenStaysStructured 断言嵌套 when 仍逐层归约且不 panic。
func TestKtNestedWhenStaysStructured(t *testing.T) {
	src := "fun f(a: Int, b: Int): Int {\n" +
		"  return when (a) {\n" +
		"    1 -> when (b) { 1 -> 10 else -> 20 }\n" +
		"    else -> 0\n" +
		"  }\n" +
		"}\n"
	out, root, _ := ktParse(t, src)
	if !ktHasKind(root, air.KindMatchExpr) {
		t.Fatal("nested when lost its match expression node")
	}
	if err := air.CheckIntervals(root, len(src)); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
	if out.NodeCount() < 12 {
		t.Fatalf("nested when produced only %d nodes", out.NodeCount())
	}
}

// TestKtSuspendFunction 专项验证 suspend 协程函数标注 SemAsync。
func TestKtSuspendFunction(t *testing.T) {
	src := "suspend fun fetch(url: String): Int {\n" +
		"  delay(100)\n" +
		"  return url?.length ?: 0\n" +
		"}\n"
	_, root, _ := ktParse(t, src)
	if !ktHasKind(root, air.KindFuncDecl) {
		t.Fatal("suspend function must lower to KindFuncDecl")
	}
	if !ktHasSem(root, air.SemAsync, "true") {
		t.Fatal("suspend function is missing SemAsync")
	}
	if !ktHasNote(root, "suspend") {
		t.Fatal("suspend function is missing its SemNote")
	}
}

// TestKtExtensionFunction 专项验证扩展函数归约与 extension 标注。
func TestKtExtensionFunction(t *testing.T) {
	src := "fun String.shout(): String {\n" +
		"  return uppercase() + \"!\"\n" +
		"}\n" +
		"fun List<Int>.sum2(): Int = size\n"
	_, root, _ := ktParse(t, src)
	if !ktHasNote(root, "extension") {
		t.Fatal("extension function is missing the extension SemNote")
	}
	if got := ktCountKind(root, air.KindFuncDecl); got < 2 {
		t.Fatalf("only %d KindFuncDecl nodes for two extension functions", got)
	}
}

// TestKtNullSafetyOperators 专项验证 ?. !! ?: 与可空类型 ? 的降级映射。
func TestKtNullSafetyOperators(t *testing.T) {
	src := "fun probe(a: String?, b: String): Int {\n" +
		"  val x = a?.length\n" +
		"  val y = a ?: b\n" +
		"  val z = a!!.length\n" +
		"  return x ?: z\n" +
		"}\n"
	_, root, _ := ktParse(t, src)
	for _, note := range []string{
		"null-safe", "elvis", "not-null-assert", "nullable",
	} {
		if !ktHasNote(root, note) {
			t.Fatalf("missing null-safety note %q", note)
		}
	}
	if ktCountKind(root, air.KindField) < 2 ||
		!ktHasSem(root, air.SemOperator, "?.") {
		t.Fatal("null-conditional member access is missing")
	}
	if !ktHasSem(root, air.SemOperator, "?:") {
		t.Fatal("Elvis operator is missing its SemOperator")
	}
}

// TestKtLambdaAndReceiver 专项验证 lambda 与带接收者 lambda。
func TestKtLambdaAndReceiver(t *testing.T) {
	src := "val doubled = list.map { x -> x * 2 }\n" +
		"val built = apply { x -> this.run(x) }\n"
	_, root, _ := ktParse(t, src)
	if !ktHasNote(root, "lambda") {
		t.Fatal("plain lambda is missing its SemNote")
	}
	if !ktHasNote(root, "lambda-with-receiver") {
		t.Fatal("lambda with receiver is missing its SemNote")
	}
	if ktCountKind(root, air.KindLambda) < 2 {
		t.Fatal("both lambdas must lower to KindLambda")
	}
}

// TestKtStringTemplateParts 专项验证 $name 与 ${expr} 模板片段。
func TestKtStringTemplateParts(t *testing.T) {
	src := "fun label(user: String, n: Int): String {\n" +
		"  return \"hi $user, count ${user.length + n}!\"\n" +
		"}\n"
	_, root, _ := ktParse(t, src)
	if !ktHasNote(root, "string-template") {
		t.Fatal("string template is missing its SemNote")
	}
	if !ktHasNote(root, "template-entry") {
		t.Fatal("template entries are missing their SemNote")
	}
	if got := ktCountKind(root, air.KindKtStringFrag); got < 2 {
		t.Fatalf("only %d template entry fragments", got)
	}
	if !ktHasSem(root, air.SemTag, "string-template") {
		t.Fatal("string template is missing its SemTag")
	}
}

// TestKtCompanionAndSealed 专项验证伴生对象与 sealed 类层级。
func TestKtCompanionAndSealed(t *testing.T) {
	src := "sealed class Result {\n" +
		"  data class Ok(val n: Int) : Result()\n" +
		"  companion object { const val ZERO = 0 }\n" +
		"}\n"
	_, root, _ := ktParse(t, src)
	for _, note := range []string{
		"sealed-class", "data-class", "companion-object", "const-val",
	} {
		if !ktHasNote(root, note) {
			t.Fatalf("missing class modifier note %q", note)
		}
	}
	if !ktHasKind(root, air.KindClassDecl) || !ktHasKind(root, air.KindConstDecl) {
		t.Fatal("class hierarchy lost its declaration kinds")
	}
}
