/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"testing"
)

// The fixture must stay isomorphic with the kernel fixture declared in
// kernel/mutation/helpers_test.go, otherwise the Abate verdicts of this
// harness would not be the verdicts of the kernel.
func TestFixtureMirrorsTheKernelFixture(t *testing.T) {
	t.Parallel()
	target, suite := newFixture()
	nodes := target.Nodes()
	want := []string{
		"a-arith", "b-bound", "c-const", "d-control", "e-return",
		"f-negate", "g-call", "h-equivalent", "i-timeout", "j-survived",
	}
	if len(nodes) != len(want) {
		t.Fatalf("fixture holds %d nodes, want %d", len(nodes), len(want))
	}
	for i, node := range want {
		if nodes[i] != node {
			t.Fatalf("node %d = %q, want %q", i, nodes[i], node)
		}
	}
	tokens := map[string]string{
		"a-arith": "+", "b-bound": "<", "c-const": "1", "d-control": "break",
		"e-return": "true", "f-negate": "==", "g-call": "call",
		"h-equivalent": "+", "i-timeout": "-", "j-survived": "*",
	}
	for node, token := range tokens {
		if target.Token(node) != token {
			t.Fatalf("token of %q = %q, want %q", node, target.Token(node), token)
		}
	}
	if suite.Coverage("j-survived") != 0 {
		t.Fatal("the no coverage node must carry a zero coverage count")
	}
	if suite.Coverage("a-arith") != 3 {
		t.Fatal("the arithmetic node must carry three covering tests")
	}
	if len(suite.Tests()) != 3 {
		t.Fatalf("suite holds %d tests, want 3", len(suite.Tests()))
	}
}

// TestFixtureIsPureAndDefensive covers the fixture purity contract.
func TestFixtureIsPureAndDefensive(t *testing.T) {
	t.Parallel()
	target, suite := newFixture()
	first := target.Fingerprint()
	if first != target.Fingerprint() {
		t.Fatal("the target fingerprint drifted")
	}
	nodes := target.Nodes()
	nodes[0] = "tampered"
	if target.Nodes()[0] == "tampered" {
		t.Fatal("Nodes() leaked its backing array")
	}
	tests := suite.Tests()
	tests[0] = "tampered"
	if suite.Tests()[0] == "tampered" {
		t.Fatal("Tests() leaked its backing array")
	}
}

// TestBaselinePortMatchesTheJavaSuite proves the port reproduces the Java
// assertions, which is what makes the witness independent of the PIT status.
func TestBaselinePortMatchesTheJavaSuite(t *testing.T) {
	t.Parallel()
	for _, entry := range witnessTable {
		if len(entry.inputs) == 0 {
			continue
		}
		for i, input := range entry.inputs {
			if got := entry.original(input); got != entry.expected[i] {
				t.Fatalf("%s:%d input %d observes %s, want %s",
					entry.method, entry.line, input, got, entry.expected[i])
			}
		}
	}
}
