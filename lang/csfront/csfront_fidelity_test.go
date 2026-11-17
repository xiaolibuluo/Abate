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
	"abate/lang/frontkit"
)

// csFidelityCorpus 覆盖 C# 各构造形态的保真夹具。
func csFidelityCorpus() []string {
	return []string{
		"using System;\nusing System.Linq;\n",
		"namespace Demo { public class C { } }\n",
		"public class C { private int n = 42; public int M() { return n; } }\n",
		"var q = from x in xs where x > 0 select x;\n",
		"class C { string s = $\"hi {name}\"; string p = @\"C:\\tmp\"; }\n",
		"// line comment\n/* block */\n/// <summary>doc</summary>\n",
		"[Obsolete] public class C { public string? Name { get; set; } }\n",
		"class C { object M(object a) { return a?.X ?? fallback; } }\n",
		"record R(int X, string Y);\n",
		"class C { void M() { if (ready) { Run(); } else { Wait(); } } }\n",
		"class C { void M(int[] a) { foreach (var x in a) { total += x; } } }\n",
		"#region Helpers\nclass C { }\n#endregion\n",
		"class C { char c = 'x'; bool b = true; object o = null; }\n",
	}
}

// TestCsTokenFidelity 证明每个标识符、字面量与注释都字节精确投影进 AIR，
// 且 AIR 不发明源文本中不存在的词法节点。
func TestCsTokenFidelity(t *testing.T) {
	for _, src := range csFidelityCorpus() {
		out, _, _ := csParse(t, src)
		tokens := csTokenize(t, src)
		if err := frontkit.VerifyTokenFidelity(tokens, out); err != nil {
			t.Errorf("token fidelity for %q: %v", src, err)
		}
	}
}

// TestCsNodeIntervals 断言每个节点区间合法、落在源范围内且兄弟不重叠。
func TestCsNodeIntervals(t *testing.T) {
	src := "using System.Linq;\n" +
		"namespace Demo {\n" +
		"public class C {\n" +
		"  public string? Name { get; set; }\n" +
		"  public int M(int x) { var q = from y in xs where y > 0 select y; return x + 1; }\n" +
		"}\n}\n"
	_, root, _ := csParse(t, src)
	nodes := csNodes(root)
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

// TestCsRootCoversWholeFile 断言根节点覆盖 [0,size) 的完整源区间。
func TestCsRootCoversWholeFile(t *testing.T) {
	for _, src := range csFidelityCorpus() {
		_, root, _ := csParse(t, src)
		start, end := air.Span(root)
		if start != 0 || end != len(src) {
			t.Fatalf("root span [%d,%d) does not cover [0,%d) for %q",
				start, end, len(src), src[:csMinInt(len(src), 24)])
		}
	}
}

// TestCsStructureSpansMultipleTokens 证明归约是有层级的结构树而非 token 等价投影。
func TestCsStructureSpansMultipleTokens(t *testing.T) {
	src := "namespace Demo { public class C { public int M(int x) { return x + 1; } } }\n"
	_, root, _ := csParse(t, src)
	tokens := csTokenize(t, src)
	structured := 0
	for _, node := range csNodes(root) {
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

// TestCsParseTokensDirect 证明脱离 FrontendBase 时 LanguageSpec 仍是真解析器。
func TestCsParseTokensDirect(t *testing.T) {
	src := "namespace Demo { class C { int M(int x) { return x; } } }\n"
	tokens := csTokenize(t, src)
	root, diags, err := (&csSpec{}).ParseTokens(tokens, "direct.cs")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if root == nil || root.Kind() != air.KindPackage {
		t.Fatal("ParseTokens did not produce a KindPackage root")
	}
	for _, kind := range []air.NodeKind{
		air.KindModuleDecl, air.KindClassDecl, air.KindMethodDecl,
		air.KindBlock, air.KindReturn,
	} {
		if !csHasKind(root, kind) {
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

// TestCsPartialRecoveryMarksDamage 断言损坏输入单遍恢复并标注 SemPartial。
func TestCsPartialRecoveryMarksDamage(t *testing.T) {
	src := "class C { void M( {\n    int x = \xff;\n    return x;\n}\n"
	out, root, diags := csParse(t, src)
	if out.NodeCount() < 4 {
		t.Fatalf("damaged file produced %d nodes", out.NodeCount())
	}
	if !csHasSem(root, air.SemPartial, "true") {
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
		if diag.Severity == air.SeverityWarning && diag.Path == "main.cs" {
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

// csMinInt 返回两个整数中较小的一个，避免引入额外依赖。
func csMinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// csConstructCase 是一个「C# 构造 → AIR 种类与语义标注」契约。
type csConstructCase struct {
	name string
	src  string
	kind air.NodeKind
	note string
}

// csConstructCases 枚举 C# 语言构造到冻结 AIR 枚举的映射契约。
func csConstructCases() []csConstructCase {
	return []csConstructCase{
		{"namespace", "namespace Demo { }\n", air.KindModuleDecl, "namespace"},
		{"using-directive", "using System;\n", air.KindImport, "using-directive"},
		{"using-static", "using static System.Math;\n", air.KindImport, "using-directive"},
		{"using-statement", "using (var f = Open()) { }\n", air.KindUsingDecl, "using-statement"},
		{"class", "class C { }\n", air.KindClassDecl, ""},
		{"struct", "struct S { }\n", air.KindStructDecl, ""},
		{"interface", "interface I { }\n", air.KindInterfaceDecl, ""},
		{"enum", "enum E { A, B }\n", air.KindEnumDecl, ""},
		{"record", "record R(int X);\n", air.KindClassDecl, "record"},
		{"delegate", "delegate int D(int x);\n", air.KindFuncDecl, "delegate"},
		{"method", "class C { public int M(int x) { return x; } }\n", air.KindMethodDecl, ""},
		{"property", "class C { public int X { get; set; } }\n", air.KindFuncDecl, "property"},
		{"indexer", "class C { public int this[int i] { get; } }\n", air.KindIndex, "indexer"},
		{"field", "class C { private int n = 1; }\n", air.KindVarDecl, ""},
		{"const", "class C { const int N = 4; }\n", air.KindConstDecl, ""},
		{"event", "class C { public event Action E; }\n", air.KindVarDecl, "event"},
		{"attribute", "[Obsolete] class C { }\n", air.KindAttribute, "attribute"},
		{"nullable", "class C { string? name; }\n", air.KindOptional, "nullable"},
		{"linq", "var q = from x in xs where x > 0 select x;\n", air.KindCall, "linq"},
		{"linq-clause", "var q = from x in xs select x;\n", air.KindCall, "linq-clause:from"},
		{"verbatim-string", "class C { string p = @\"C:\\tmp\"; }\n", air.KindStrLit, "verbatim"},
		{"interpolated-string", "class C { string s = $\"hi {name}\"; }\n", air.KindStrLit, "interpolated"},
		{"null-conditional", "class C { int? M(A a) { return a?.X; } }\n", air.KindField, "null-conditional"},
		{"null-coalescing", "class C { object M(object a, object b) { return a ?? b; } }\n", air.KindBinary, "null-coalescing"},
		{"lambda", "class C { Func<int,int> f = x => x + 1; }\n", air.KindLambda, "lambda"},
		{"ternary", "class C { int M(bool c) { return c ? 1 : 0; } }\n", air.KindTernary, ""},
		{"if", "class C { void M(bool c) { if (c) { } } }\n", air.KindIf, ""},
		{"else", "class C { void M(bool c) { if (c) { } else { } } }\n", air.KindIf, "else"},
		{"for", "class C { void M() { for (int i = 0; i < 3; i++) { } } }\n", air.KindFor, ""},
		{"foreach", "class C { void M(int[] a) { foreach (var x in a) { } } }\n", air.KindRangeFor, "foreach"},
		{"while", "class C { void M(bool c) { while (c) { } } }\n", air.KindWhile, ""},
		{"do-while", "class C { void M(bool c) { do { } while (c); } }\n", air.KindDoWhile, ""},
		{"switch", "class C { void M(int x) { switch (x) { case 1: break; } } }\n", air.KindSwitch, ""},
		{"case", "class C { void M(int x) { switch (x) { case 1: break; } } }\n", air.KindCase, ""},
		{"default-case", "class C { void M(int x) { switch (x) { default: break; } } }\n", air.KindCase, "default"},
		{"try", "class C { void M() { try { } catch (E e) { } finally { } } }\n", air.KindTry, ""},
		{"catch", "class C { void M() { try { } catch (E e) { } finally { } } }\n", air.KindCatch, ""},
		{"finally", "class C { void M() { try { } catch (E e) { } finally { } } }\n", air.KindFinally, ""},
		{"return", "class C { int M(int x) { return x; } }\n", air.KindReturn, ""},
		{"throw", "class C { void M() { throw new E(); } }\n", air.KindThrow, ""},
		{"break", "class C { void M() { for (;;) { break; } } }\n", air.KindBranch, "break"},
		{"continue", "class C { void M() { for (;;) { continue; } } }\n", air.KindBranch, "continue"},
		{"goto", "class C { void M() { goto done; } }\n", air.KindBranch, "goto"},
		{"yield", "class C { IEnumerable M() { yield return 1; } }\n", air.KindReturn, "yield"},
		{"lock", "class C { void M(object o) { lock (o) { } } }\n", air.KindWith, "lock"},
		{"checked", "class C { void M() { checked { x++; } } }\n", air.KindBlock, "checked"},
		{"unsafe", "class C { void M() { unsafe { } } }\n", air.KindBlock, "unsafe"},
		{"await", "class C { async Task M() { return await F(); } }\n", air.KindAwait, ""},
		{"new-expression", "class C { object o = new Foo(); }\n", air.KindUnary, ""},
		{"typeof", "class C { Type t = typeof(int); }\n", air.KindTypeAssert, "typeof"},
		{"nameof", "class C { string s = nameof(X); }\n", air.KindCall, "nameof"},
		{"region", "#region Helpers\n#endregion\n", air.KindDirective, "region"},
		{"directive", "#pragma warning disable\n", air.KindDirective, "directive"},
		{"comment", "// hello\n", air.KindComment, ""},
		{"xml-doc", "/// <summary>doc</summary>\n", air.KindDoc, "xml-doc"},
		{"num-lit", "class C { int n = 42; }\n", air.KindNumLit, ""},
		{"str-lit", "class C { string s = \"hi\"; }\n", air.KindStrLit, ""},
		{"char-lit", "class C { char c = 'x'; }\n", air.KindCharLit, ""},
		{"bool-lit", "class C { bool b = true; }\n", air.KindBoolLit, ""},
		{"nil-lit", "class C { object o = null; }\n", air.KindNilLit, ""},
		{"call", "class C { void M() { F(1); } }\n", air.KindCall, ""},
		{"field-access", "class C { void M(A a) { var x = a.X; } }\n", air.KindField, ""},
		{"index", "class C { void M(int[] a) { var x = a[0]; } }\n", air.KindIndex, ""},
		{"binary", "class C { int M(int a) { return a + 1; } }\n", air.KindBinary, ""},
		{"assign", "class C { void M() { x = 1; } }\n", air.KindAssign, ""},
		{"type-assert", "class C { void M(object o) { var x = o is string; } }\n", air.KindTypeAssert, ""},
	}
}

// TestCsOperatorFidelity asserts AIR-ID changes for cmp, op and digit mutations, with identical trailing comments on both variants.
func TestCsOperatorFidelity(t *testing.T) {
	cases := []struct{ name, base, mutated string }{
		{"cmp", "class C { bool M(int a, int b) { return a == b; } }\n", "class C { bool M(int a, int b) { return a != b; } }\n"},
		{"op", "class C { int M(int a, int b) { return a + b; } }\n", "class C { int M(int a, int b) { return a - b; } }\n"},
		{"digit", "class C { int M() { return 41; } }\n", "class C { int M() { return 42; } }\n"},
	}
	rootID := func(t *testing.T, src string) air.NodeID {
		t.Helper()
		out, _, _ := csParse(t, src)
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
