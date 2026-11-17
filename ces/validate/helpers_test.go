/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"testing"

	"abate/ces/parse"
	"abate/ces/schema"
)

// cleanYAML is a fully legal, normalization-safe CES document that passes
// every one of the 25 listed CES-VALID rules.
const cleanYAML = `apiVersion: ces.tencent.com/v1
kind: Experiment
metadata:
  name: validate-baseline
  namespace: default
spec:
  target:
    kind: simulator
    selector:
      node: n1
  steadyState:
    - probeRef: p1
      expectation:
        type: value
        value: "42"
      tolerance:
        type: absolute
        value: 1
      holdFor: 1s
  method:
    expression: "ε"
  probes:
    - id: p1
      kind: http
      spec:
        url: http://sim/health
        method: GET
        expectStatus: 200
      interval: 1s
  rollback:
    strategy: auto
    verifySteadyState: true
  schedule:
    once:
      at: 5s
  seed: 12345
`

// fresh decodes and normalizes a brand-new baseline experiment, so every test
// case works on independent maps and slices with zero shared mutable state.
func fresh(t *testing.T) *schema.Experiment {
	t.Helper()
	exp, err := parse.Parse([]byte(cleanYAML), parse.FormatYAML)
	if err != nil {
		t.Fatalf("parse baseline: %v", err)
	}
	normalized, err := parse.Normalize(exp)
	if err != nil {
		t.Fatalf("normalize baseline: %v", err)
	}
	return normalized
}

// ok fails the test on an unexpected error.
func ok(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// hasRule reports whether a violation list contains the given rule id.
func hasRule(vs []Violation, id string) bool {
	for _, v := range vs {
		if v.RuleID == id {
			return true
		}
	}
	return false
}

// assertRule fails unless the exact rule id is present.
func assertRule(t *testing.T, vs []Violation, id string) {
	t.Helper()
	if !hasRule(vs, id) {
		t.Fatalf("expected violation %s, got %+v", id, vs)
	}
}

// assertNoRule fails if the rule id is present.
func assertNoRule(t *testing.T, vs []Violation, id string) {
	t.Helper()
	if hasRule(vs, id) {
		t.Fatalf("unexpected violation %s in %+v", id, vs)
	}
}

// ruleViolations runs one registered rule against exp, isolating its result.
func ruleViolations(t *testing.T, id string, exp *schema.Experiment) []Violation {
	t.Helper()
	reg, err := builtinRegistry()
	ok(t, err)
	rule, exists := reg.rules[id]
	if !exists {
		t.Fatalf("rule %s is not registered", id)
	}
	return rule.Check(exp)
}
