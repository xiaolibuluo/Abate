/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"strings"
	"testing"

	"abate/lang/air"
)

// pyChaosInputs covers damaged files, unmatched delimiters, mixed tab/space
// indentation, illegal UTF-8, excessive nesting and empty buffers.
var pyChaosInputs = []string{
	"",
	"\n",
	"   \n\n",
	"\t\n",
	"def f(:\n    ???\n",
	"if x:\n\tpass\n        pass\n",
	"x = ((((((",
	"x = [1, 2,",
	"s = \"unterminated\n",
	"x = 1 \xff\n",
	"@\n",
	"async\n",
	"match:\n  case:\n",
	"def f():\n\treturn 1\n    return 2\n",
	strings.Repeat("(", 200) + strings.Repeat(")", 200),
	strings.Repeat("if x:\n    ", 100) + "pass\n",
	strings.Repeat("x = 1\n", 400),
	"f'{unclosed\n",
	"\\",
	":::???;;;@@@\n",
}

// TestPyChaosNeverPanics feeds hostile input through the whole adapter and
// requires a valid unit, a deterministic fingerprint and no panic.
func TestPyChaosNeverPanics(t *testing.T) {
	fe := pyFrontend(t)
	for i, src := range pyChaosInputs {
		unit := air.SourceUnit{
			Language: "py", Dialect: "3.12",
			Files: []air.SourceFile{air.NewSourceFile("main.py", []byte(src))},
		}
		out, diags, err := fe.Parse(unit)
		if err != nil {
			t.Logf("chaos[%d] classified error: %v", i, err)
			continue
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("chaos[%d] %q: missing root", i, src[:minInt(len(src), 40)])
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("chaos[%d] %q: invalid unit: %v", i,
				src[:minInt(len(src), 40)], err)
		}
		if out.NodeCount() < 1 {
			t.Fatalf("chaos[%d]: empty tree", i)
		}
		first := out.Fingerprint()
		second := out.Fingerprint()
		if first != second {
			t.Fatalf("chaos[%d]: fingerprint unstable", i)
		}
		_ = diags
	}
}

// TestPyChaosDeterministicAcrossRuns re-parses every damaged buffer and checks
// the resulting unit hash is byte-identical between independent runs.
func TestPyChaosDeterministicAcrossRuns(t *testing.T) {
	fe := pyFrontend(t)
	for i, src := range pyChaosInputs {
		unit := air.SourceUnit{
			Language: "py", Dialect: "3.12",
			Files: []air.SourceFile{air.NewSourceFile("main.py", []byte(src))},
		}
		a, _, errA := fe.Parse(unit)
		b, _, errB := fe.Parse(unit)
		if (errA == nil) != (errB == nil) {
			t.Fatalf("chaos[%d]: inconsistent error classification", i)
		}
		if a == nil || b == nil {
			continue
		}
		if a.Fingerprint() != b.Fingerprint() {
			t.Fatalf("chaos[%d]: non-deterministic damaged parse", i)
		}
		if a.NodeCount() != b.NodeCount() {
			t.Fatalf("chaos[%d]: non-deterministic node count", i)
		}
	}
}

// TestPyChaosDeepNestingStaysBounded proves the depth ceiling caps recursion
// on adversarial nesting without losing the package root.
func TestPyChaosDeepNestingStaysBounded(t *testing.T) {
	src := strings.Repeat("[", 300) + "1" + strings.Repeat("]", 300) + "\n"
	out, root, _ := pyParse(t, src)
	if root.Kind() != air.KindPackage {
		t.Fatal("deep nesting lost the root")
	}
	if out.NodeCount() < 2 || out.NodeCount() > 10000 {
		t.Fatalf("unbounded node count %d", out.NodeCount())
	}
	if err := out.Validate(); err != nil {
		t.Fatalf("deep nesting produced invalid unit: %v", err)
	}
}

// minInt returns the smaller of two integers without importing math.
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
