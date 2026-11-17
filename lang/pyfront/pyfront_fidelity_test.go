/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// TestPyTokenFidelity proves every identifier, literal and comment of the real
// token stream is projected byte-exactly into the AIR tree, and no AIR lexical
// node invents source text that is absent from the stream.
func TestPyTokenFidelity(t *testing.T) {
	sources := []string{
		"import os\nfrom sys import argv\n",
		"def f(x: int, y=0) -> bool:\n    return x > y\n",
		"class C(Base):\n    pass\n",
		"xs = [x * 2 for x in items if x > 0]\n",
		"msg = f\"hi {name}, {count}\"\n",
		"# type: ignore\n# plain comment\n",
		"\"\"\"module doc\"\"\"\n",
		"t = (1, 2, 3)\nd = {\"a\": 1, \"b\": 2}\n",
		"x = obj.field.method(a, *args, **kwargs)\n",
		"v = a if c else b\np = a[1:5]\n",
		"t: Optional[int] = None\nalias: list[str] = []\n",
		"try:\n    f()\nexcept E as e:\n    raise e\n",
	}
	for _, src := range sources {
		out, _, _ := pyParse(t, src)
		tokens := pyTokenize(t, src)
		if err := frontkit.VerifyTokenFidelity(tokens, out); err != nil {
			t.Errorf("token fidelity for %q: %v", src, err)
		}
	}
}

// TestPyNodeIntervals checks that every AIR vertex stays inside the source
// extent and carries a valid, non-negative half-open byte interval.
func TestPyNodeIntervals(t *testing.T) {
	src := "import os\n" +
		"def f(a: list[str], b=1) -> int:\n" +
		"    xs = [x for x in a if x > 0]\n" +
		"    return sum(xs) + b\n"
	_, root, _ := pyParse(t, src)
	nodes := pyNodes(root)
	if len(nodes) < 20 {
		t.Fatalf("only %d nodes in layered tree", len(nodes))
	}
	for _, n := range nodes {
		pos := n.Pos()
		if !pos.Valid() || pos.Offset < 0 || pos.EndOffset > len(src) {
			t.Fatalf("%s has out-of-range interval %s", n.Kind(), pos)
		}
		if n.End().Offset < pos.Offset {
			t.Fatalf("%s end precedes start", n.Kind())
		}
	}
}

// TestPyPartialRecoveryMarksDamage proves damaged input is reduced in one pass
// with SemPartial on the recovered vertices rather than aborting the file.
func TestPyPartialRecoveryMarksDamage(t *testing.T) {
	src := "def f(x:\n    y = ?\n    return x\n"
	out, root, _ := pyParse(t, src)
	if out.NodeCount() < 3 {
		t.Fatalf("damaged file produced %d nodes", out.NodeCount())
	}
	if !pyHasSem(root, air.SemPartial, "true") {
		t.Fatal("damaged reduction is not marked SemPartial")
	}
	if out.Roots[0].Kind() != air.KindPackage {
		t.Fatal("damaged file lost its KindPackage root")
	}
	if err := out.Validate(); err != nil {
		t.Fatalf("damaged unit is invalid: %v", err)
	}
}

// TestPyIllegalByteDiagnostic checks the deterministic recovery warning.
func TestPyIllegalByteDiagnostic(t *testing.T) {
	src := "x = 1 \xff\n"
	_, _, diags := pyParse(t, src)
	found := false
	for _, d := range diags {
		if d.Severity == air.SeverityWarning {
			found = true
			if d.Path != "main.py" || d.Code == "" {
				t.Fatalf("malformed diagnostic: %+v", d)
			}
		}
	}
	if !found {
		t.Fatal("illegal byte did not produce a severity warning")
	}
}

// TestPyParseTokensDirect proves the LanguageSpec reduction is a real parser
// when called without the FrontendBase skeleton.
func TestPyParseTokensDirect(t *testing.T) {
	src := "def f(x):\n    if x:\n        return x\n    return None\n"
	tokens := pyTokenize(t, src)
	root, diags, err := (pySpec{}).ParseTokens(tokens, "direct.py")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if root == nil || root.Kind() != air.KindPackage {
		t.Fatal("ParseTokens did not produce a KindPackage root")
	}
	if !pyHasKind(root, air.KindFuncDecl) || !pyHasKind(root, air.KindIf) ||
		!pyHasKind(root, air.KindBlock) {
		t.Fatal("ParseTokens produced an empty shell")
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if err := frontkit.VerifyNodeFidelity(tokens, root); err != nil {
		t.Fatalf("node fidelity: %v", err)
	}
}

// pyMapCase is one Python construct and the AIR kind plus note it must produce.
type pyMapCase struct {
	name string
	src  string
	kind air.NodeKind
	note string
}

// pyMappingCases enumerates the Python 3.8-3.12 construct-to-AIR contract.
var pyMappingCases = []pyMapCase{
	{"import", "import os\n", air.KindImport, ""},
	{"from-import", "from sys import argv\n", air.KindImport, "from-import"},
	{"def", "def f(x):\n    return x\n", air.KindFuncDecl, ""},
	{"return", "def f():\n    return 1\n", air.KindReturn, ""},
	{"class", "class C(Base):\n    pass\n", air.KindClassDecl, ""},
	{"async-def", "async def g():\n    await h()\n", air.KindFuncDecl, ""},
	{"await", "async def g():\n    await h()\n", air.KindAwait, ""},
	{"decorator", "@deco\ndef f():\n    pass\n", air.KindDecorator, "decorator"},
	{"if", "if x:\n    pass\n", air.KindIf, ""},
	{"elif", "if x:\n    pass\nelif y:\n    pass\n", air.KindIf, "elif"},
	{"else", "if x:\n    pass\nelse:\n    pass\n", air.KindIf, "else"},
	{"try", "try:\n    f()\nexcept E:\n    raise e\nfinally:\n    g()\n", air.KindTry, ""},
	{"except", "try:\n    f()\nexcept E:\n    raise e\nfinally:\n    g()\n", air.KindCatch, ""},
	{"finally", "try:\n    f()\nexcept E:\n    raise e\nfinally:\n    g()\n", air.KindFinally, ""},
	{"with", "with open(\"f\") as fh:\n    pass\n", air.KindWith, ""},
	{"for", "for i in xs:\n    break\n", air.KindRangeFor, ""},
	{"while", "while n:\n    continue\n", air.KindWhile, ""},
	{"match", "match v:\n    case 1:\n        pass\n", air.KindMatch, ""},
	{"match-expr", "match v:\n    case 1:\n        pass\n", air.KindMatchExpr, ""},
	{"case", "match v:\n    case 1:\n        pass\n", air.KindCase, ""},
	{"list-comprehension", "xs = [x for x in ys]\n", air.KindCompositeLit, "comprehension"},
	{"dict-comprehension", "d = {k: v for k, v in ps}\n", air.KindCompositeLit, "dict-comprehension"},
	{"set-comprehension", "s = {x for x in xs}\n", air.KindCompositeLit, "set-comprehension"},
	{"generator-expression", "g = (x for x in xs)\n", air.KindCompositeLit, "generator-comprehension"},
	{"lambda", "f = lambda a, b: a + b\n", air.KindLambda, ""},
	{"ternary", "v = a if c else b\n", air.KindTernary, ""},
	{"walrus", "if (n := f()) > 0:\n    pass\n", air.KindAssign, "walrus"},
	{"slice", "p = a[1:5:2]\n", air.KindSlice, "slice"},
	{"optional-subscript", "t: Optional[int] = None\n", air.KindOptional, "optional"},
	{"generic-inst", "t: list[str] = []\n", air.KindGenericInst, ""},
	{"union-none", "v = int | None\n", air.KindOptional, "optional"},
	{"fstring", "s = f\"hi {name}\"\n", air.KindStrLit, "fstring"},
	{"fstring-content", "s = f\"hi {name}\"\n", air.KindPyStringFragment, "py-string-content"},
	{"docstring", "\"\"\"doc\"\"\"\n", air.KindDoc, "docstring"},
	{"directive", "# type: ignore\n", air.KindDirective, ""},
	{"star-args", "f(*args, **kw)\n", air.KindStar, "star-args"},
	{"kwargs", "f(*args, **kw)\n", air.KindSpread, "kwargs"},
	{"ellipsis", "...\n", air.KindRegion, "ellipsis"},
	{"assert", "assert x\n", air.KindCall, "assert"},
	{"global", "global g\n", air.KindDeclStmt, "global"},
	{"nonlocal", "nonlocal n\n", air.KindDeclStmt, "nonlocal"},
	{"del", "del xs[0]\n", air.KindUnary, "del"},
	{"tuple", "t = (1, 2)\n", air.KindTuple, ""},
	{"dict-literal", "d = {\"a\": 1}\n", air.KindCompositeLit, "dict-literal"},
	{"dict-pair", "d = {\"a\": 1}\n", air.KindAssign, "dict-pair"},
	{"set-literal", "s = {1, 2}\n", air.KindCompositeLit, "set-literal"},
	{"list-literal", "l = [1, 2]\n", air.KindCompositeLit, "list-literal"},
	{"bool-lit", "b = True\n", air.KindBoolLit, ""},
	{"nil-lit", "n = None\n", air.KindNilLit, ""},
	{"yield", "def f():\n    yield 1\n", air.KindUnary, "yield"},
	{"raise", "raise ValueError()\n", air.KindThrow, ""},
	{"break", "for i in xs:\n    break\n", air.KindBranch, "break"},
	{"continue", "for i in xs:\n    continue\n", air.KindBranch, "continue"},
	{"pass", "pass\n", air.KindDeclStmt, "pass"},
	{"type-annotation", "x: int = 1\n", air.KindDeclStmt, "type-annotation"},
	{"named-type", "x: int = 1\n", air.KindNamed, ""},
	{"call", "f(1)\n", air.KindCall, ""},
	{"field", "a.b\n", air.KindField, ""},
	{"index", "a[0]\n", air.KindIndex, ""},
	{"binary", "x = 1 + 2 * 3\n", air.KindBinary, ""},
	{"unary", "x = -1\n", air.KindUnary, ""},
	{"expr-stmt", "f()\n", air.KindExprStmt, ""},
	{"comment", "# plain\n", air.KindComment, ""},
	{"indent-block", "if x:\n    pass\n", air.KindBlock, "indent-based"},
	{"char-lit", "s = 'x'\n", air.KindCharLit, ""},
	{"str-lit", "s = \"x\"\n", air.KindStrLit, ""},
	{"num-lit", "n = 42\n", air.KindNumLit, ""},
	{"ident", "name\n", air.KindIdent, ""},
	{"comprehension-clause", "xs = [x for x in ys]\n", air.KindRangeFor, "comprehension-clause"},
}

// TestPyOperatorFidelity asserts AIR-ID changes for cmp, op and digit mutations, with identical trailing comments on both variants.
func TestPyOperatorFidelity(t *testing.T) {
	cases := []struct{ name, base, mutated string }{
		{"cmp", "x = 1 == 2\n", "x = 1 != 2\n"},
		{"op", "x = 1 + 2\n", "x = 1 - 2\n"},
		{"digit", "x = 41\n", "x = 42\n"},
	}
	rootID := func(t *testing.T, src string) air.NodeID {
		t.Helper()
		out, _, _ := pyParse(t, src)
		return air.IDOf(out.Roots[0])
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if rootID(t, tc.base) == rootID(t, tc.mutated) {
				t.Fatal("operator or digit mutation did not change AIR-ID")
			}
			baseComment, mutatedComment := tc.base+"# tail\n", tc.mutated+"# tail\n"
			if rootID(t, baseComment) == rootID(t, mutatedComment) {
				t.Fatal("identical comments changed the capture conclusion")
			}
		})
	}
}
