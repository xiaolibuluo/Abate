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

// pyIncrementalSource 是增量校验用的真实 Python 样例。
const pyIncrementalSource = "import os\n\n\ndef f(x):\n    return x + 1\n"

// TestPyIncrementalMatchesFullParse proves the incremental result of a real
// Python frontend equals the result of a full parse of the edited source.
//
// Three real editing scenarios are exercised: inserting a token, deleting a
// token and replacing a token. A frontend that kept a stale subtree would
// produce a different tree here.
func TestPyIncrementalMatchesFullParse(t *testing.T) {
	cases := []struct {
		name  string
		edits func(source string) []air.Edit
	}{
		{"insert", func(source string) []air.Edit {
			at := strings.Index(source, "return x")
			return []air.Edit{{
				Path: "main.py", Offset: at, EndOffset: at,
				Text: []byte("y = 1\n    "),
			}}
		}},
		{"delete", func(source string) []air.Edit {
			at := strings.Index(source, " + 1")
			return []air.Edit{{
				Path: "main.py", Offset: at, EndOffset: at + 4,
			}}
		}},
		{"replace", func(source string) []air.Edit {
			at := strings.Index(source, "x + 1")
			return []air.Edit{{
				Path: "main.py", Offset: at, EndOffset: at + 5,
				Text: []byte("x * 2"),
			}}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			frontend := pyFrontend(t)
			unit := air.SourceUnit{
				Language: "py", Dialect: "3.12",
				Files: []air.SourceFile{
					air.NewSourceFile("main.py", []byte(pyIncrementalSource)),
				},
			}
			old, _, err := frontend.Parse(unit)
			if err != nil {
				t.Fatalf("cold parse: %v", err)
			}
			edits := tc.edits(pyIncrementalSource)
			fresh, _, err := frontend.ParseIncremental(unit, old, edits)
			if err != nil {
				t.Fatalf("ParseIncremental: %v", err)
			}
			next, err := air.ApplyEdits([]byte(pyIncrementalSource), edits)
			if err != nil {
				t.Fatalf("ApplyEdits: %v", err)
			}
			full, _, err := frontend.Parse(air.SourceUnit{
				Language: "py", Dialect: "3.12",
				Files: []air.SourceFile{air.NewSourceFile("main.py", next)},
			})
			if err != nil {
				t.Fatalf("full parse: %v", err)
			}
			if err := pyCompareTrees(fresh, full); err != nil {
				t.Fatalf("incremental result differs from the full parse: %v",
					err)
			}
		})
	}
}

// TestPyIncrementalDialectChangeForcesRebuild proves the Python frontend does
// not reuse a subtree that was parsed under another dialect.
func TestPyIncrementalDialectChangeForcesRebuild(t *testing.T) {
	frontend := pyFrontend(t)
	unit := air.SourceUnit{
		Language: "py", Dialect: "3.12",
		Files: []air.SourceFile{
			air.NewSourceFile("main.py", []byte(pyIncrementalSource)),
		},
	}
	old, _, err := frontend.Parse(unit)
	if err != nil {
		t.Fatalf("cold parse: %v", err)
	}
	other := unit
	other.Dialect = "3.10"
	fresh, _, err := frontend.ParseIncremental(other, old, nil)
	if err != nil {
		t.Fatalf("ParseIncremental: %v", err)
	}
	if fresh.Annotations["lang.py.reusedFiles"] != "0" {
		t.Fatalf("a dialect change reused %q files",
			fresh.Annotations["lang.py.reusedFiles"])
	}
	if fresh.Roots[0] == old.Roots[0] {
		t.Fatal("the stale subtree of the previous dialect was reused")
	}
}

// pyCompareTrees compares two units while ignoring the incremental-only
// annotations, because those describe the reuse and not the tree.
func pyCompareTrees(left, right *air.Unit) error {
	stripped := *left
	stripped.Annotations = make(map[string]string, len(left.Annotations))
	for key, value := range left.Annotations {
		if strings.HasSuffix(key, ".incremental") ||
			strings.HasSuffix(key, ".reusedFiles") {
			continue
		}
		stripped.Annotations[key] = value
	}
	return air.CompareUnits(&stripped, right)
}
