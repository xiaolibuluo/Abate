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
	"abate/lang/frontkit"
)

// ktFidelityCorpus 覆盖 Kotlin 各构造形态的保真夹具。
func ktFidelityCorpus() []string {
	return []string{
		"package demo\nimport kotlin.math.abs\n",
		"data class User(val name: String, val age: Int)\n",
		"interface Repo { fun load(id: Int): String }\n",
		"suspend fun fetch(url: String): Int { return url.length }\n",
		"fun String.shout() = uppercase()\n",
		"sealed class Result { data class Ok(val n: Int) : Result() }\n",
		"val s = \"hi $name and ${name.len}\"\n",
		"val raw = \"\"\"raw $notemplate\"\"\"\n",
		"val `weird name` = 42\n",
		"// line comment\n/* block */\n/** KDoc */\n",
		"val n = a?.len ?: 0\n",
		"val r = when (x) { 1 -> a() else -> b() }\n",
		"enum class Dir { NORTH, SOUTH }\n",
		"companion object Factory { const val MAX = 10 }\n",
		"fun box(): Int { val c = 'x'; val b = true; val z = null; return 1 }\n",
		"for (x in 1..10) { println(x) }\n",
	}
}

// TestKtTokenFidelity 证明每个标识符、字面量与注释都字节精确投影进 AIR，
// 且 AIR 不发明源文本中不存在的词法节点。
func TestKtTokenFidelity(t *testing.T) {
	for _, src := range ktFidelityCorpus() {
		out, _, _ := ktParse(t, src)
		tokens := ktTokenize(t, src)
		if err := frontkit.VerifyTokenFidelity(tokens, out); err != nil {
			t.Errorf("token fidelity for %q: %v", src, err)
		}
	}
}

// TestKtNodeIntervals 断言每个节点区间合法、落在源范围内且兄弟不重叠。
func TestKtNodeIntervals(t *testing.T) {
	src := "package demo\n" +
		"data class User(val name: String)\n" +
		"suspend fun fetch(url: String?): Int {\n" +
		"  val label = \"hi $url\"\n" +
		"  val n = url?.length ?: 0\n" +
		"  val r = when (n) { 1 -> a() else -> b() }\n" +
		"  return r!!\n" +
		"}\n"
	_, root, _ := ktParse(t, src)
	nodes := ktNodes(root)
	if len(nodes) < 30 {
		t.Fatalf("only %d nodes in layered tree", len(nodes))
	}
	for _, node := range nodes {
		pos := node.Pos()
		if !pos.Valid() || pos.Offset < 0 || pos.EndOffset > len(src) {
			t.Fatalf("%s has out-of-range interval %s", node.Kind(), pos)
		}
		if node.End().Offset < pos.Offset {
			t.Fatalf("%s end precedes start", node.Kind())
		}
	}
	if err := air.CheckIntervals(root, len(src)); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
}

// TestKtRootCoversWholeFile 断言根节点覆盖 [0,size) 的完整源区间。
func TestKtRootCoversWholeFile(t *testing.T) {
	for _, src := range ktFidelityCorpus() {
		_, root, _ := ktParse(t, src)
		start, end := air.Span(root)
		if start != 0 || end != len(src) {
			t.Fatalf("root span [%d,%d) does not cover [0,%d) for %q",
				start, end, len(src), src[:ktMinInt(len(src), 24)])
		}
	}
}

// TestKtStructureSpansMultipleTokens 证明归约是有层级的结构树而非 token 等价投影。
func TestKtStructureSpansMultipleTokens(t *testing.T) {
	src := "package demo\ndata class User(val name: String)\n" +
		"suspend fun fetch(url: String): Int { return url?.length ?: 0 }\n"
	_, root, _ := ktParse(t, src)
	tokens := ktTokenize(t, src)
	structured := 0
	for _, node := range ktNodes(root) {
		start, end := air.Span(node)
		contained := 0
		for _, tok := range tokens {
			switch tok.Kind {
			case frontkit.TokenNewline, frontkit.TokenEOF,
				frontkit.TokenIndent, frontkit.TokenDedent:
				continue
			}
			if start <= tok.Offset && tok.EndOffset <= end {
				contained++
			}
		}
		if contained >= 6 {
			structured++
		}
	}
	if structured < 3 {
		t.Fatalf("only %d structural nodes span multiple tokens; reduction looks flat",
			structured)
	}
}

// TestKtParseTokensDirect 证明脱离 FrontendBase 时 LanguageSpec 仍是真解析器。
func TestKtParseTokensDirect(t *testing.T) {
	src := "package demo\ndata class User(val name: String)\n" +
		"fun greet(n: String): String { return \"hi $n\" }\n"
	tokens := ktTokenize(t, src)
	root, diags, err := (&ktSpec{}).ParseTokens(tokens, "direct.kt")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if root == nil || root.Kind() != air.KindPackage {
		t.Fatal("ParseTokens did not produce a KindPackage root")
	}
	for _, kind := range []air.NodeKind{
		air.KindModuleDecl, air.KindClassDecl, air.KindFuncDecl,
		air.KindBlock, air.KindReturn, air.KindStrLit,
	} {
		if !ktHasKind(root, kind) {
			t.Fatalf("ParseTokens produced an empty shell: missing %s", kind)
		}
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if err := frontkit.VerifyNodeFidelity(tokens, root); err != nil {
		t.Fatalf("node fidelity: %v", err)
	}
}

// TestKtPartialRecoveryMarksDamage 断言损坏输入单遍恢复并标注 SemPartial。
func TestKtPartialRecoveryMarksDamage(t *testing.T) {
	src := "class C { fun M( {\n    val x = \xff;\n    return x\n}\n"
	out, root, diags := ktParse(t, src)
	if out.NodeCount() < 4 {
		t.Fatalf("damaged file produced %d nodes", out.NodeCount())
	}
	if !ktHasSem(root, air.SemPartial, "true") {
		t.Fatal("damaged reduction is not marked SemPartial")
	}
	if root.Kind() != air.KindPackage {
		t.Fatal("damaged file lost its KindPackage root")
	}
	if err := out.Validate(); err != nil {
		t.Fatalf("damaged unit is invalid: %v", err)
	}
	warning := false
	for _, diag := range diags {
		if diag.Severity == air.SeverityWarning && diag.Path == "main.kt" {
			warning = true
			if diag.Code == "" {
				t.Fatalf("malformed diagnostic: %+v", diag)
			}
		}
	}
	if !warning {
		t.Fatal("illegal byte did not produce a severity warning")
	}
}

// ktConstructCase 是一个「Kotlin 构造 → AIR 种类与语义标注」契约。
type ktConstructCase struct {
	name string
	src  string
	kind air.NodeKind
	note string
}

// ktConstructCases 枚举 Kotlin 语言构造到冻结 AIR 枚举的映射契约。
func ktConstructCases() []ktConstructCase {
	return []ktConstructCase{
		{"package", "package demo\n", air.KindModuleDecl, "package"},
		{"import", "import kotlin.math.abs\n", air.KindImport, "import"},
		{"class", "class Foo { }\n", air.KindClassDecl, ""},
		{"data-class", "data class User(val n: Int)\n", air.KindClassDecl, "data-class"},
		{"sealed-class", "sealed class Result { }\n", air.KindClassDecl, "sealed-class"},
		{"enum-class", "enum class Dir { A, B }\n", air.KindEnumDecl, "enum-class"},
		{"annotation-class", "annotation class Marker\n", air.KindAnnotationDecl, "annotation-class"},
		{"interface", "interface Repo { }\n", air.KindInterfaceDecl, ""},
		{"object", "object Singleton { }\n", air.KindClassDecl, "object"},
		{"companion-object", "companion object Factory { }\n", air.KindClassDecl, "companion-object"},
		{"fun", "fun greet(name: String) { }\n", air.KindFuncDecl, ""},
		{"extension", "fun String.shout() = uppercase()\n", air.KindFuncDecl, "extension"},
		{"operator-fun", "operator fun plus(x: Int) = x\n", air.KindFuncDecl, "operator"},
		{"when-stmt", "when (x) { 1 -> a() else -> b() }\n", air.KindMatch, "when"},
		{"when-case", "when (x) { 1 -> a() else -> b() }\n", air.KindCase, "when-branch"},
		{"when-expr", "val r = when (x) { 1 -> a() else -> b() }\n", air.KindMatchExpr, "when-expr"},
		{"val", "val x = 1\n", air.KindVarDecl, "val-binding"},
		{"var", "var y = 2\n", air.KindVarDecl, "var-binding"},
		{"const-val", "const val MAX = 10\n", air.KindConstDecl, "const-val"},
		{"typealias", "typealias Handler = (Int) -> Unit\n", air.KindTypeDecl, "typealias"},
		{"init", "init { println(1) }\n", air.KindBlock, "init"},
		{"null-safe", "val n = a?.len\n", air.KindField, "null-safe"},
		{"not-null", "val n = a!!\n", air.KindUnary, "not-null-assert"},
		{"elvis", "val n = a ?: b\n", air.KindBinary, "elvis"},
		{"nullable", "val s: String? = null\n", air.KindOptional, "nullable"},
		{"string-template", "val s = \"hi $name\"\n", air.KindStrLit, "string-template"},
		{"template-entry", "val s = \"hi ${name.len}\"\n", air.KindKtStringFrag, "template-entry"},
		{"raw-string", "val s = \"\"\"raw text\"\"\"\n", air.KindStrLit, "raw-string"},
		{"backtick-ident", "val `weird name` = 1\n", air.KindIdent, "backtick-ident"},
		{"lambda", "list.map { x -> x + 1 }\n", air.KindLambda, "lambda"},
		{"lambda-receiver", "apply { x -> this.run(x) }\n", air.KindLambda, "lambda-with-receiver"},
		{"kdoc", "/** doc */\n", air.KindDoc, "kdoc"},
		{"comment", "// note\n", air.KindComment, ""},
		{"if", "if (ok) { run() }\n", air.KindIf, ""},
		{"for", "for (x in xs) { run(x) }\n", air.KindRangeFor, "for-in"},
		{"while", "while (ok) { run() }\n", air.KindWhile, ""},
		{"return", "fun f(): Int { return 1 }\n", air.KindReturn, ""},
		{"throw", "throw Error()\n", air.KindThrow, ""},
		{"try", "try { run() } catch (e: E) { } finally { }\n", air.KindTry, ""},
		{"call", "foo(1)\n", air.KindCall, ""},
		{"field-access", "val n = obj.field\n", air.KindField, ""},
		{"index", "val n = arr[0]\n", air.KindIndex, ""},
		{"binary", "val n = a + b\n", air.KindBinary, ""},
		{"assign", "x = 1\n", air.KindAssign, ""},
		{"type-assert", "val b = a is String\n", air.KindTypeAssert, ""},
		{"num-lit", "val n = 42\n", air.KindNumLit, ""},
		{"char-lit", "val c = 'x'\n", air.KindCharLit, ""},
		{"bool-lit", "val b = true\n", air.KindBoolLit, ""},
		{"nil-lit", "val n = null\n", air.KindNilLit, ""},
	}
}

// TestKtOperatorFidelity asserts AIR-ID changes for cmp, op and digit mutations, with identical trailing comments on both variants.
func TestKtOperatorFidelity(t *testing.T) {
	cases := []struct{ name, base, mutated string }{
		{"cmp", "fun eq(a: Int, b: Int): Boolean { return a == b }\n", "fun eq(a: Int, b: Int): Boolean { return a != b }\n"},
		{"op", "fun add(a: Int, b: Int): Int { return a + b }\n", "fun add(a: Int, b: Int): Int { return a - b }\n"},
		{"digit", "fun value(): Int { return 41 }\n", "fun value(): Int { return 42 }\n"},
	}
	rootID := func(t *testing.T, src string) air.NodeID {
		t.Helper()
		out, _, _ := ktParse(t, src)
		return air.IDOf(out.Roots[0])
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if rootID(t, tc.base) == rootID(t, tc.mutated) {
				t.Fatal("operator or digit mutation did not change AIR-ID")
			}
			baseComment, mutatedComment := tc.base+"// tail\n", tc.mutated+"// tail\n"
			if rootID(t, baseComment) == rootID(t, mutatedComment) {
				t.Fatal("identical comments changed the capture conclusion")
			}
		})
	}
}
