/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"errors"
	"testing"

	"abate/ces/schema"
)

// baseYAML is a complete, executable CES document used as the valid corpus seed.
const baseYAML = `apiVersion: ces.tencent.com/v1
kind: Experiment
metadata:
  name: sample
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

// baseJSON is the JSON twin of baseYAML; durations are text literals.
const baseJSON = `{
  "apiVersion": "ces.tencent.com/v1",
  "kind": "Experiment",
  "metadata": {"name": "sample", "namespace": "default"},
  "spec": {
    "target": {"kind": "simulator", "selector": {"node": "n1"}},
    "steadyState": [{"probeRef": "p1", "expectation": {"type": "value", "value": "42"},
      "tolerance": {"type": "absolute", "value": 1}, "holdFor": "1s"}],
    "method": {"expression": "ε"},
    "probes": [{"id": "p1", "kind": "http",
      "spec": {"url": "http://sim/health", "method": "GET", "expectStatus": 200},
      "interval": "1s"}],
    "rollback": {"strategy": "auto", "verifySteadyState": true},
    "schedule": {"once": {"at": "5s"}},
    "seed": 12345
  }
}
`

// detRNG is a deterministic local LCG used only inside chaos and mutation tests.
type detRNG struct{ state uint64 }

// next advances the generator and returns the next unsigned word.
func (g *detRNG) next() uint64 {
	g.state = g.state*6364136223846793005 + 1442695040888963407
	return g.state
}

// ruleOf extracts the deterministic ParseError rule id, or an empty string.
func ruleOf(err error) string {
	var pe *ParseError
	if errors.As(err, &pe) {
		return pe.RuleID
	}
	return ""
}

// requireRule asserts that err is a ParseError carrying the expected rule id.
func requireRule(t *testing.T, err error, want string) {
	t.Helper()
	if got := ruleOf(err); got != want {
		t.Fatalf("rule = %q, want %q (err = %v)", got, want, err)
	}
}

// requireOK fails the test when a fallible parse step returns an error.
func requireOK(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// validExperiment decodes the base YAML document into an experiment value.
func validExperiment(t *testing.T) *schema.Experiment {
	t.Helper()
	exp, err := Parse([]byte(baseYAML), FormatYAML)
	requireOK(t, err)
	return exp
}

// runChain executes the full parse pipeline on a raw document.
func runChain(data []byte, format Format) (*schema.Experiment, error) {
	exp, err := Parse(data, format)
	if err != nil {
		return nil, err
	}
	normalized, err := Normalize(exp)
	if err != nil {
		return nil, err
	}
	if _, err = CompileExpr(normalized.Spec.Method.Expression,
		normalized.Spec.Method.Operators); err != nil {
		return nil, err
	}
	return normalized, nil
}
