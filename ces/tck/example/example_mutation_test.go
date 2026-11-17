/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package example

import (
	"strings"
	"testing"
)

// Mutation testing for the example producer/consumer pairing: each mutant is a
// CES document derived from a valid Builder output but altered in exactly one
// field. A correct pipeline must "kill" the mutant by emitting the expected
// CES-VALID rule or rejecting it at parse time. This proves the example driver
// actually validates semantics rather than accepting everything, and covers the
// violation-mapping branch of Validate.

func baseBuilder() *Builder {
	return NewBuilder("mut", "n1").
		WithSteadyState("p1", "ok", 2).
		WithMetricProbe("p1", "up", "1", 1)
}

// TestMutationValidateDetectsTampering mutates a valid document's steady-state
// and target fields and requires the driver to surface the corresponding rule.
func TestMutationValidateDetectsTampering(t *testing.T) {
	doc, err := baseBuilder().Marshal()
	if err != nil {
		t.Fatalf("marshal base: %v", err)
	}
	d := Driver{}

	// Mutant 1: empty the steady-state hypothesis list by removing probes/holds
	// via a fresh builder that omits WithSteadyState. Validation must flag the
	// missing steady-state rule (CES-VALID-STEADY-001).
	noSteady, err := NewBuilder("mut", "n1").
		WithMetricProbe("p1", "up", "1", 1).Marshal()
	if err != nil {
		t.Fatalf("marshal no-steady mutant: %v", err)
	}
	viols, verr := d.Validate(noSteady)
	if verr != nil {
		t.Fatalf("validate no-steady mutant: %v", verr)
	}
	foundSteady := false
	for _, v := range viols {
		if strings.Contains(v.RuleID, "STEADY") {
			foundSteady = true
		}
	}
	if !foundSteady {
		t.Errorf("no-steady mutant survived: %+v", viols)
	}

	// Mutant 2: an external target kind is forbidden by the CES spec. Replace
	// the simulator target with an external one in the marshaled YAML.
	external := strings.Replace(string(doc), "kind: simulator", "kind: external", 1)
	viols, verr = d.Validate([]byte(external))
	if verr != nil {
		t.Fatalf("validate external-target mutant: %v", verr)
	}
	foundTarget := false
	for _, v := range viols {
		if strings.Contains(v.RuleID, "TARGET") {
			foundTarget = true
		}
	}
	if !foundTarget {
		t.Errorf("external-target mutant survived: %+v", viols)
	}

	// Control: the unmutated document has zero violations (no false positives).
	viols, verr = d.Validate(doc)
	if verr != nil {
		t.Fatalf("validate control: %v", verr)
	}
	if len(viols) != 0 {
		t.Errorf("control document has %d violations: %+v", len(viols), viols)
	}
}

// TestMutationMethodExpressionOverrides proves WithMethodExpression genuinely
// changes the assembled document (otherwise the mutation would be inert) and
// that the resulting expression is preserved through marshaling.
func TestMutationMethodExpressionOverrides(t *testing.T) {
	base, err := baseBuilder().WithMethodExpression("ε").Marshal()
	if err != nil {
		t.Fatalf("marshal base: %v", err)
	}
	mutated, err := baseBuilder().WithMethodExpression("ε ⊗ ε").Marshal()
	if err != nil {
		t.Fatalf("marshal mutated: %v", err)
	}
	if string(base) == string(mutated) {
		t.Fatal("WithMethodExpression did not change the document")
	}
	if !strings.Contains(string(mutated), "ε ⊗ ε") {
		t.Error("overridden expression not present in marshaled document")
	}
	// Both expressions remain parseable CES documents.
	d := Driver{}
	for name, doc := range map[string][]byte{"base": base, "mutated": mutated} {
		if res, perr := d.Parse(doc, "yaml"); perr != nil || !res.OK {
			t.Errorf("%s expression rejected: ok=%t err=%v", name, res.OK, perr)
		}
	}
}

// TestMutationBuilderFieldAssembly proves every With* setter mutates exactly the
// intended section of the assembled experiment (a setter that silently no-ops
// would be a mutation survivor).
func TestMutationBuilderFieldAssembly(t *testing.T) {
	b := baseBuilder()
	exp := b.Build()
	if exp.Metadata.Name != "mut" {
		t.Errorf("name = %q", exp.Metadata.Name)
	}
	if len(exp.Spec.SteadyState) != 1 {
		t.Errorf("steady states = %d, want 1", len(exp.Spec.SteadyState))
	}
	if len(exp.Spec.Probes) != 1 {
		t.Errorf("probes = %d, want 1", len(exp.Spec.Probes))
	}
	// Appending a second hypothesis/probe must grow the slices, not overwrite.
	b.WithSteadyState("p2", "ok", 3).WithMetricProbe("p2", "up", "1", 2)
	exp = b.Build()
	if len(exp.Spec.SteadyState) != 2 || len(exp.Spec.Probes) != 2 {
		t.Errorf("append failed: steady=%d probes=%d", len(exp.Spec.SteadyState), len(exp.Spec.Probes))
	}
	if exp.Spec.Target.Selector["node"] != "n1" {
		t.Errorf("target selector lost: %+v", exp.Spec.Target.Selector)
	}
}
