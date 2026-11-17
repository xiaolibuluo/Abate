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
)

// dedicatedCase mirrors the QA cprobe method: the smallest source that
// triggers one Python construct must really produce that construct's
// dedicated AIR kind. These cases are the evidence that a rule promoted to
// exact is backed by a genuine syntax-node reduction.
type dedicatedCase struct {
	construct string
	src       string
	kind      air.NodeKind
}

// pyDedicatedCases is the full set of G4 construct reductions pyfront owns.
var pyDedicatedCases = []dedicatedCase{
	{"assert", "assert x\n", air.KindPyAssert},
	{"pass", "pass\n", air.KindPyPass},
	{"global", "global g\n", air.KindPyScopeDecl},
	{"nonlocal", "nonlocal n\n", air.KindPyScopeDecl},
	{"delete", "del xs[0]\n", air.KindPyDelete},
	{"yield", "def f():\n    yield 1\n", air.KindPyYield},
	{"not", "x = not y\n", air.KindPyNotOp},
	{"boolean-op", "b = x and y\n", air.KindPyBoolOp},
	{"compare", "b = x < y\n", air.KindPyCompare},
	{"augmented-assignment", "x += 1\n", air.KindPyAugAssign},
	{"named-expression", "if (n := f()):\n    pass\n", air.KindPyNamedExpr},
	{"dict-pair", "d = {\"a\": 1}\n", air.KindPyDictPair},
	{"keyword-argument", "f(x=1)\n", air.KindPyKeywordArg},
	{"comprehension-clause", "xs = [x for x in ys if x]\n", air.KindPyCompClause},
	{"generator", "g = (x for x in xs)\n", air.KindPyGenerator},
	{"string-part", "s = \"abc\"\n", air.KindPyStringFragment},
	{"type-conversion", "s = f\"{a!r}\"\n", air.KindPyTypeConversion},
	{"import-prefix", "from . import rel\n", air.KindPyImportPrefix},
	{"with-item", "with open('f') as fh:\n    pass\n", air.KindPyWithItem},
	{"dotted-name", "import a.b.c\n", air.KindPyQualifiedName},
	{"member-type", "x: os.PathLike = None\n", air.KindPyQualifiedName},
}

// TestPyDedicatedKindsAreReallyProduced proves each dedicated construct kind
// appears in the AIR tree of a real parse of its minimal trigger source.
func TestPyDedicatedKindsAreReallyProduced(t *testing.T) {
	for _, tc := range pyDedicatedCases {
		t.Run(tc.construct, func(t *testing.T) {
			unit, _, _ := pyParse(t, tc.src)
			found := false
			for _, node := range air.Collect(unit) {
				if node.Kind() == tc.kind {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("%s: minimal source %q did not produce %s",
					tc.construct, tc.src, tc.kind)
			}
		})
	}
}
