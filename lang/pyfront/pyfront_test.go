/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
	"abate/lang/frontkit"
)

// pyFrontend returns a fresh Python adapter or fails the calling test.
func pyFrontend(t *testing.T) air.Frontend {
	t.Helper()
	fe, err := NewPyFrontend()
	if err != nil {
		t.Fatalf("NewPyFrontend: %v", err)
	}
	return fe
}

// pyParse reduces one in-memory Python source buffer into a unit and its root.
func pyParse(t *testing.T, src string) (*air.Unit, air.Node, []air.Diagnostic) {
	t.Helper()
	unit := air.SourceUnit{
		Language: "py", Dialect: "3.12",
		Files: []air.SourceFile{air.NewSourceFile("main.py", []byte(src))},
	}
	out, diags, err := pyFrontend(t).Parse(unit)
	if err != nil {
		t.Fatalf("Parse(%q): %v", src, err)
	}
	if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
		t.Fatalf("Parse(%q): missing AIR root", src)
	}
	return out, out.Roots[0], diags
}

// pyTokenize produces the exact token stream that FrontendBase.Parse delegates.
func pyTokenize(t *testing.T, src string) []frontkit.Token {
	t.Helper()
	lexer, err := frontkit.NewLexer(pyLexSpec{}, []byte(src))
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize(%q): %v", src, err)
	}
	return tokens
}

// pyNodes returns a pre-order snapshot of every reachable AIR vertex.
func pyNodes(root air.Node) []air.Node {
	out := make([]air.Node, 0, 16)
	var walk func(air.Node)
	walk = func(n air.Node) {
		if n == nil {
			return
		}
		out = append(out, n)
		for _, child := range n.Children() {
			walk(child)
		}
	}
	walk(root)
	return out
}

func pyHasKind(root air.Node, kind air.NodeKind) bool {
	for _, n := range pyNodes(root) {
		if n.Kind() == kind {
			return true
		}
	}
	return false
}

func pyHasNote(root air.Node, note string) bool {
	return pyHasSem(root, air.SemNote, note)
}

func pyHasSem(root air.Node, key air.SemKey, value string) bool {
	for _, n := range pyNodes(root) {
		if stored, ok := n.Sem(key).(string); ok && stored == value {
			return true
		}
	}
	return false
}

// The construct-to-AIR mapping table lives in pyfront_fidelity_test.go so
// that every source file stays within the 300-line limit.

// TestPyConstructToAIRMapping proves 70 real Python constructions reduce to
// frozen AIR kinds and semantic notes instead of a constant empty shell.
func TestPyConstructToAIRMapping(t *testing.T) {
	if len(pyMappingCases) < 20 {
		t.Fatalf("only %d mapping cases; at least 20 required", len(pyMappingCases))
	}
	for _, tc := range pyMappingCases {
		t.Run(tc.name, func(t *testing.T) {
			out, root, _ := pyParse(t, tc.src)
			if out.NodeCount() < 2 {
				t.Fatalf("empty AIR shell for %q", tc.src)
			}
			if !pyHasKind(root, tc.kind) {
				t.Fatalf("%s: missing %s in tree for %q", tc.name, tc.kind, tc.src)
			}
			if tc.note != "" && !pyHasNote(root, tc.note) {
				t.Fatalf("%s: missing SemNote %q for %q", tc.name, tc.note, tc.src)
			}
		})
	}
}

// TestPyFrontendContract proves the returned air.Frontend produces a layered,
// valid unit with a KindPackage root and more than one reachable vertex.
func TestPyFrontendContract(t *testing.T) {
	fe := pyFrontend(t)
	if fe == nil {
		t.Fatal("NewPyFrontend returned a nil air.Frontend")
	}
	src := "import os\nclass C:\n    def f(self, x):\n        return x + 1\n"
	unit := air.SourceUnit{
		Language: "py", Dialect: "3.12",
		Files: []air.SourceFile{air.NewSourceFile("main.py", []byte(src))},
	}
	out, diags, err := fe.Parse(unit)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if err := out.Validate(); err != nil {
		t.Fatalf("Unit.Validate: %v", err)
	}
	if out.NodeCount() < 10 {
		t.Fatalf("NodeCount=%d, want a layered tree", out.NodeCount())
	}
	if out.Roots[0].Kind() != air.KindPackage {
		t.Fatalf("root kind=%s, want KindPackage", out.Roots[0].Kind())
	}
	if len(out.Roots[0].Children()) < 2 {
		t.Fatal("root has no statement children")
	}
	if len(diags) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got := out.Annotations["lang.py.dialect"]; got != "3.12" {
		t.Fatalf("dialect annotation=%q, want 3.12", got)
	}
}

// TestPyDeterminismFingerprint certifies Unit.Fingerprint is byte-stable.
func TestPyDeterminismFingerprint(t *testing.T) {
	fe := pyFrontend(t)
	src := "import os\nfrom typing import Optional\n" +
		"def f(a: Optional[int] = None):\n" +
		"    xs = [x * 2 for x in range(10) if x > 1]\n" +
		"    return f\"{xs}\"\n"
	unit := air.SourceUnit{
		Language: "py", Dialect: "3.12",
		Files: []air.SourceFile{air.NewSourceFile("main.py", []byte(src))},
	}
	determinism.MustMatch100(t, 0x5059, func() []byte {
		out, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		sum := out.Fingerprint()
		return sum[:]
	})
}

// TestPyDialectNormalization checks the canonical 3.8-3.13 dialect mapping.
func TestPyDialectNormalization(t *testing.T) {
	good := map[string]string{
		"": "3.12", "3.8": "3.8", "3.12": "3.12", "3.13": "3.13",
		"python3.11": "3.11", "py3.10": "3.10", "cpython3.9": "3.9",
		"  PY3.8 ": "3.8",
	}
	for in, want := range good {
		got, err := (pySpec{}).NormalizeDialect(in)
		if err != nil || got != want {
			t.Fatalf("NormalizeDialect(%q)=%q,%v want %q", in, got, err, want)
		}
	}
	for _, in := range []string{"4.0", "ruby3", "3.7", "go1.26"} {
		if got, err := (pySpec{}).NormalizeDialect(in); err == nil {
			t.Fatalf("NormalizeDialect(%q)=%q, want error", in, got)
		}
	}
}

// TestPyLexSpecContract checks the frozen lexical table contract.
func TestPyLexSpecContract(t *testing.T) {
	spec := pySpec{}
	if spec.Tag() != "py" {
		t.Fatalf("Tag=%q, want py", spec.Tag())
	}
	lex := spec.LexSpec()
	if lex == nil || !lex.IndentSensitive() {
		t.Fatal("python lexer must be indentation sensitive")
	}
	if opening, closing := lex.BlockComment(); opening != "" || closing != "" {
		t.Fatal("python has no block comments")
	}
	if lex.LineComment() != "#" {
		t.Fatalf("LineComment=%q, want #", lex.LineComment())
	}
	for _, table := range []struct {
		name   string
		values []string
		want   string
	}{
		{"keywords", lex.Keywords(), "def"},
		{"operators", lex.Operators(), ":="},
		{"quotes", lex.StringQuotes(), `"""`},
		{"prefixes", lex.RawStringPrefixes(), "f"},
		{"separators", lex.NumberSeparators(), "_"},
	} {
		if !pyContains(table.values, table.want) {
			t.Fatalf("%s table missing %q", table.name, table.want)
		}
	}
}

// TestPySemanticKeys checks non-note AIR annotations on proprietary forms.
func TestPySemanticKeys(t *testing.T) {
	cases := []struct {
		name, src string
		key       air.SemKey
		value     string
	}{
		{"async-def", "async def g():\n    await h()\n", air.SemAsync, "true"},
		{"binary-op", "x = 1 + 2\n", air.SemOperator, "+"},
		{"assign-op", "x = 1\n", air.SemOperator, "="},
		{"annotated-assign", "x: int = 1\n", air.SemOperator, "="},
		{"type-param", "t: list[str] = []\n", air.SemTypeParam, "1"},
		{"directive", "# type: ignore\n", air.SemDirective, "# type: ignore"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, root, _ := pyParse(t, tc.src)
			if !pyHasSem(root, tc.key, tc.value) {
				t.Fatalf("missing %s=%q for %q", tc.key, tc.value, tc.src)
			}
		})
	}
}

// Ordered-table lookup coverage lives in pyfront_mutation_test.go so that
// every source file stays within the 300-line limit.
