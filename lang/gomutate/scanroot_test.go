/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// scanRootTargetSource is the judgeable subtree: one arithmetic token and
// one real test that kills mutations of it. It lives below target/ so the
// module also carries an unrelated sibling package outside the scan root.
const scanRootTargetSource = `package target

func Add(a, b int) int {
	return a + b
}
`

const scanRootTargetTestSource = `package target

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Fatal("Add(2,3) != 5")
	}
}
`

// scanRootOutsideSource belongs to the same module but outside the selected
// subtree. Its test is intentionally red: if discovery ever regresses to the
// module root, baseline coverage executes it and New fails before any mutant
// can be judged.
const scanRootOutsideSource = `package outside

func Sub(a, b int) int {
	return a - b
}
`

const scanRootOutsideTestSource = `package outside

import "testing"

func TestOutsideFails(t *testing.T) {
	if Sub(3, 2) != 1 {
		t.Fatal("unreachable")
	}
	t.Fatal("outside scanRoot test must never run")
}
`

// writeScanRootModule materializes one module whose judgeable target sits in
// a subdirectory while a sibling package outside that subdirectory carries a
// failing test. go.mod is pinned to a released toolchain line so discovery
// and coverage run fully offline.
func writeScanRootModule(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	files := map[string]string{
		"go.mod":                  "module example.com/scanroot\n\ngo 1.21\n",
		"target/target.go":        scanRootTargetSource,
		"target/target_test.go":   scanRootTargetTestSource,
		"outside/outside.go":      scanRootOutsideSource,
		"outside/outside_test.go": scanRootOutsideTestSource,
	}
	for name, content := range files {
		path := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return root
}

// TestNewBackendScanRootExcludesOutsideTests is the regression proof for the
// discovery-range fix: a mutate session anchored at a module subdirectory
// discovers and executes only the tests below that subdirectory. The sibling
// package is red on purpose; a module-root discovery would abort New while
// measuring the pristine baseline.
func TestNewBackendScanRootExcludesOutsideTests(t *testing.T) {
	root := writeScanRootModule(t)
	scanRoot := filepath.Join(root, "target")
	backend, err := New(context.Background(), scanRoot)
	if err != nil {
		t.Fatalf("New over scanRoot must ignore outside red tests: %v", err)
	}
	defer func() { _ = backend.Close() }()
	if backend.TestCount() != 1 {
		t.Fatalf("tests = %d, want exactly the in-scope test", backend.TestCount())
	}
	tests := backend.Suite().Tests()
	if len(tests) != 1 {
		t.Fatalf("suite tests = %d, want 1", len(tests))
	}
	joined := fmt.Sprint(tests)
	if strings.Contains(joined, "outside") {
		t.Fatalf("suite leaked outside-scanRoot tests: %s", joined)
	}
	if !strings.Contains(joined, "target.TestAdd") {
		t.Fatalf("suite lost the in-scope test: %s", joined)
	}
}

// TestWithinDir pins the lexical boundary rule used to drop packages outside
// scanRoot: equality and descendants pass, while a sibling sharing the same
// name prefix must not be mistaken for a descendant.
func TestWithinDir(t *testing.T) {
	root := string(filepath.Separator) + filepath.Join("mod", "target")
	cases := []struct {
		name string
		dir  string
		want bool
	}{
		{name: "equal", dir: root, want: true},
		{name: "child", dir: filepath.Join(root, "pkg"), want: true},
		{name: "sibling prefix", dir: root + "2", want: false},
		{name: "parent", dir: filepath.Dir(root), want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := withinDir(root, tc.dir); got != tc.want {
				t.Fatalf("withinDir(%q, %q) = %v, want %v", root, tc.dir, got, tc.want)
			}
		})
	}
}
