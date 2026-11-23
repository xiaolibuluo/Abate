/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"abate/kernel/algebra"
	"abate/kernel/mutation"
	"abate/kernel/verify"
	"context"
	"os"
	"path/filepath"
	"testing"
)

// moduleSource is a tiny but real Go module with one arithmetic function,
// one comparison and a test asserting their behavior. Mutating "+" to "-"
// or "<" to "<=" genuinely breaks the assertions, so the real toolchain can
// kill those mutants; mutating nothing (identity) survives.
const moduleSource = `package calc

func Add(a, b int) int {
	return a + b
}

func Less(a, b int) bool {
	return a < b
}

const one = 1
`

const moduleTestSource = `package calc

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Fatal("Add(2,3) != 5")
	}
}

func TestLess(t *testing.T) {
	if !Less(2, 3) {
		t.Fatal("Less(2,3) should be true")
	}
}
`

// writeModule materializes the fixture module inside t.TempDir and returns
// its root. The go.mod is pinned to the toolchain's own major version line
// so `go test` works fully offline with zero module downloads.
func writeModule(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	files := map[string]string{
		"go.mod":       "module example.com/calc\n\ngo 1.21\n",
		"calc.go":      moduleSource,
		"calc_test.go": moduleTestSource,
	}
	for name, content := range files {
		path := filepath.Join(root, name)
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return root
}

// TestNewBackendDiscoveries verifies the three real discovery steps: token
// scan finds the mutable operator tokens, test discovery finds both TestXxx
// functions, and coverage measurement reports a positive count for the
// covered lines of calc.go.
func TestNewBackendDiscoveries(t *testing.T) {
	root := writeModule(t)
	backend, err := New(context.Background(), root)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	defer func() { _ = backend.Close() }()
	if backend.TokenCount() == 0 {
		t.Fatal("no tokens discovered")
	}
	if backend.TestCount() != 2 {
		t.Fatalf("tests = %d, want 2", backend.TestCount())
	}
	target := backend.Target()
	suite := backend.Suite()
	if target.Fingerprint() == suite.Fingerprint() {
		t.Fatal("target and suite fingerprints collide")
	}
	nodes := target.Nodes()
	if len(nodes) != backend.TokenCount() {
		t.Fatalf("nodes = %d, want %d", len(nodes), backend.TokenCount())
	}
	tests := suite.Tests()
	if len(tests) != 2 {
		t.Fatalf("suite tests = %d, want 2", len(tests))
	}
}

// TestRealKilling is the end-to-end proof of the backend: a mutant that
// changes the arithmetic operator of Add is killed by the real test suite,
// while the identity mutant survives. Both judgments come from actual
// `go test` executions against materialized workspaces, never from any
// static approximation.
func TestRealKilling(t *testing.T) {
	root := writeModule(t)
	backend, err := New(context.Background(), root)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	defer func() { _ = backend.Close() }()
	target := backend.Target()
	var plusNode, identNode string
	for _, node := range target.Nodes() {
		if target.Token(node) == "+" && plusNode == "" {
			plusNode = node
		}
	}
	if plusNode == "" {
		t.Fatal("no + token discovered in fixture module")
	}
	identNode = plusNode
	reg := mutation.NewRegistry()
	op := reg.For("+")
	if len(op) == 0 {
		t.Fatal("registry has no operator for +")
	}
	mutated := op[0].Params()["to"]
	killedPoint := mutation.NewMutationPoint(plusNode, "+", mutated, op[0], 1, 1)
	identPoint := mutation.NewMutationPoint(identNode, "+", "+", op[0], 1, 1)
	suite := backend.Suite()
	killed := false
	held := false
	for _, test := range suite.Tests() {
		switch target.Evaluate(killedPoint, test) {
		case verify.OutcomeViolated:
			killed = true
		}
		if target.Evaluate(identPoint, test) == verify.OutcomeHeld {
			held = true
		}
	}
	if !killed {
		t.Fatal("+ -> mutated mutant was not killed by the real test suite")
	}
	if !held {
		t.Fatal("identity mutant did not survive the real test suite")
	}
}

// TestEvaluateCachingAndEquivalence verifies the two pure fast paths of the
// Target contract: equivalence for identity mutations is decided without
// the toolchain, and repeated evaluation of one (point, test) pair returns
// the cached outcome deterministically.
func TestEvaluateCachingAndEquivalence(t *testing.T) {
	root := writeModule(t)
	backend, err := New(context.Background(), root)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	defer func() { _ = backend.Close() }()
	target := backend.Target()
	nodes := target.Nodes()
	if len(nodes) == 0 {
		t.Skip("no tokens in fixture")
	}
	reg := mutation.NewRegistry()
	tok := target.Token(nodes[0])
	ops := reg.For(tok)
	var op mutationPointOperator
	if len(ops) > 0 {
		op = ops[0]
	}
	point := mutation.NewMutationPoint(nodes[0], tok, tok, op, 0, 0)
	if !target.Equivalent(point) {
		t.Fatal("identity point not reported equivalent")
	}
	tests := backend.Suite().Tests()
	if len(tests) == 0 {
		t.Skip("no tests in fixture")
	}
	first := target.Evaluate(point, tests[0])
	second := target.Evaluate(point, tests[0])
	if first != second {
		t.Fatalf("cached outcome changed: %s then %s", first, second)
	}
}

// mutationPointOperator aliases the algebra operator type used in tests to
// keep the helper signatures short while remaining the exact type the
// kernel passes through.
type mutationPointOperator = algebra.Operator
