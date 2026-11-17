/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"context"
	"slices"
	"testing"

	"abate/internal/determinism"
	"abate/kernel/mutation"
	"abate/kernel/verify"
)

// TestFullPipelineAIRToReport walks the whole M3 chain explicitly: AIR parse,
// rule loading, derivation, registry injection, engine run and report replay.
func TestFullPipelineAIRToReport(t *testing.T) {
	t.Parallel()
	unit := sampleUnit(t)
	if unit.NodeCount() == 0 {
		t.Fatal("the AIR sample parsed into an empty tree")
	}
	rules := builtinRules(t)
	specs, err := NewDeriver().Derive(rules)
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}
	ops, err := Operators(specs)
	if err != nil {
		t.Fatalf("Operators: %v", err)
	}
	reg, err := mutation.NewRegistryWith(ops)
	if err != nil {
		t.Fatalf("NewRegistryWith: %v", err)
	}
	if len(reg.All()) != 7+len(ops) {
		t.Fatalf("registry exposes %d operators, want 7 + %d", len(reg.All()), len(ops))
	}
	engine, err := mutation.NewEngineWithRegistry(42, reg)
	if err != nil {
		t.Fatalf("NewEngineWithRegistry: %v", err)
	}
	tests := experimentTests()
	target, err := NewAIRTarget(unit, specs, tests)
	if err != nil {
		t.Fatalf("NewAIRTarget: %v", err)
	}
	suite, err := NewAIRSuite(target.Nodes(), tests)
	if err != nil {
		t.Fatalf("NewAIRSuite: %v", err)
	}
	rep, err := engine.Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("Engine.Run: %v", err)
	}
	ids := DerivedIDs(rep)
	if len(ids) < 20 {
		t.Fatalf("end-to-end report has only %d derived operators", len(ids))
	}
	if len(ids) != len(specs) {
		t.Fatalf("end-to-end report has %d derived ids, want %d", len(ids), len(specs))
	}
	for i, verdict := range rep.Verdicts {
		if !allocatedVerdict(verdict) {
			t.Fatalf("verdict %d = %q is not an allocated state", i, verdict)
		}
	}
	if rep.Describe() == "" {
		t.Fatal("report summary is empty")
	}
	firstDescribe := rep.Describe()
	if rep.Describe() != firstDescribe {
		t.Fatal("report summary is unstable")
	}
	if rep.TargetFingerprint != target.Fingerprint() ||
		rep.SuiteFingerprint != suite.Fingerprint() {
		t.Fatal("report did not bind the target and suite fingerprints")
	}
	determinism.MustMatch100(t, 42, func() []byte { return rep.Bytes() })
}

// TestAIRTargetContract locks the language-level target port: its node set,
// representative tokens, fingerprint stability, and the kill predicate.
func TestAIRTargetContract(t *testing.T) {
	t.Parallel()
	unit := sampleUnit(t)
	specs, err := NewDeriver().Derive(builtinRules(t))
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}
	tests := experimentTests()
	target, err := NewAIRTarget(unit, specs, tests)
	if err != nil {
		t.Fatalf("NewAIRTarget: %v", err)
	}
	nodes := target.Nodes()
	if len(nodes) == 0 || !slices.IsSorted(nodes) {
		t.Fatal("AIRTarget.Nodes must be non-empty and sorted")
	}
	tokens := make(map[string]bool, len(nodes))
	for _, node := range nodes {
		token := target.Token(node)
		if token != "" {
			tokens[token] = true
		}
	}
	for _, want := range []string{"<", "1", "true", "break", "call", "==", "+"} {
		if !tokens[want] {
			t.Fatalf("AIRTarget is missing the representative token %q", want)
		}
	}
	if target.Token("no-such-node") != "" {
		t.Fatal("an unknown node must not produce a token")
	}
	firstDigest := target.Fingerprint()
	secondDigest := target.Fingerprint()
	if firstDigest != secondDigest {
		t.Fatal("AIRTarget.Fingerprint is unstable")
	}
	firstNode := nodes[0]
	if !target.Equivalent(mutation.MutationPoint{
		Node: firstNode, Original: "x", Mutated: "x",
	}) {
		t.Fatal("an unchanged mutant must be reported as equivalent")
	}
	if target.Equivalent(mutation.MutationPoint{
		Node: firstNode, Original: "x", Mutated: "y",
	}) {
		t.Fatal("a changed mutant must not be reported as equivalent")
	}
	violated := false
	for _, node := range nodes {
		point := mutation.MutationPoint{Node: node, Original: "a", Mutated: "b"}
		outcome := target.Evaluate(point, "t1")
		if outcome == verify.OutcomeViolated {
			violated = true
		}
		if outcome != verify.OutcomeViolated && outcome != verify.OutcomeHeld {
			t.Fatalf("Evaluate returned an unallocated outcome %q", outcome)
		}
	}
	if !violated {
		t.Fatal("the static AIR kill predicate never produced a violation")
	}
	if target.Evaluate(mutation.MutationPoint{
		Node: firstNode, Original: "a", Mutated: "a",
	}, "t1") != verify.OutcomeHeld {
		t.Fatal("an unchanged mutant must not be killed")
	}
}

// TestAIRSuiteContract locks the language-level suite port: completeness,
// coverage, and determinism for identical inputs.
func TestAIRSuiteContract(t *testing.T) {
	t.Parallel()
	unit := sampleUnit(t)
	specs, err := NewDeriver().Derive(builtinRules(t))
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}
	tests := experimentTests()
	target, err := NewAIRTarget(unit, specs, tests)
	if err != nil {
		t.Fatalf("NewAIRTarget: %v", err)
	}
	nodes := target.Nodes()
	suite, err := NewAIRSuite(nodes, tests)
	if err != nil {
		t.Fatalf("NewAIRSuite: %v", err)
	}
	if !slices.IsSorted(suite.Tests()) || len(suite.Tests()) != len(tests) {
		t.Fatal("AIRSuite.Tests must be sorted and complete")
	}
	for _, node := range nodes {
		if suite.Coverage(node) == 0 {
			t.Fatalf("node %q is not covered by the static suite", node)
		}
	}
	againTarget, err := NewAIRTarget(unit, specs, tests)
	if err != nil {
		t.Fatalf("second NewAIRTarget: %v", err)
	}
	againSuite, err := NewAIRSuite(againTarget.Nodes(), tests)
	if err != nil {
		t.Fatalf("second NewAIRSuite: %v", err)
	}
	if target.Fingerprint() != againTarget.Fingerprint() ||
		suite.Fingerprint() != againSuite.Fingerprint() {
		t.Fatal("the AIR ports are not deterministic for identical inputs")
	}
}
