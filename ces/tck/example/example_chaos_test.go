/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package example

import (
	"errors"
	"strings"
	"testing"

	"abate/ces/parse"
)

// TestChaosToFormatBranches covers every branch of the format-hint mapping,
// including case-insensitivity, surrounding whitespace and the auto fallback.
func TestChaosToFormatBranches(t *testing.T) {
	cases := map[string]parse.Format{
		"json": parse.FormatJSON, "JSON": parse.FormatJSON, " Json ": parse.FormatJSON,
		"yaml": parse.FormatYAML, "YAML": parse.FormatYAML, "\tyaml\n": parse.FormatYAML,
		"auto": parse.FormatAuto, "": parse.FormatAuto, "toml": parse.FormatAuto,
	}
	for hint, want := range cases {
		if got := toFormat(hint); got != want {
			t.Errorf("toFormat(%q) = %q, want %q", hint, got, want)
		}
	}
}

// TestChaosRuleIDOfBranches proves structured parse errors yield their rule id
// while plain/wrapped/unstructured errors are handled correctly.
func TestChaosRuleIDOfBranches(t *testing.T) {
	structured := &parse.ParseError{RuleID: "CES-PARSE-006"}
	if got := ruleIDOf(structured); got != "CES-PARSE-006" {
		t.Errorf("structured = %q, want CES-PARSE-006", got)
	}
	if got := ruleIDOf(errors.New("plain")); got != "" {
		t.Errorf("plain error = %q, want empty", got)
	}
	if got := ruleIDOf(nil); got != "" {
		t.Errorf("nil error = %q, want empty", got)
	}
	wrapped := &wrappedParseError{inner: structured}
	if got := ruleIDOf(wrapped); got != "CES-PARSE-006" {
		t.Errorf("wrapped = %q, want CES-PARSE-006", got)
	}
}

// wrappedParseError is an error that wraps a structured parse error, used to
// verify errors.As unwrapping in ruleIDOf.
type wrappedParseError struct{ inner error }

func (w *wrappedParseError) Error() string { return "wrapped: " + w.inner.Error() }
func (w *wrappedParseError) Unwrap() error { return w.inner }

// TestChaosDriverRejectsMalformed proves the example driver's Parse and
// Validate propagate parse/normalize errors and attach a real CES-PARSE rule id
// on rejection — the branches the all-valid conformance run cannot reach.
func TestChaosDriverRejectsMalformed(t *testing.T) {
	d := Driver{}
	malformed := []byte("spec: [unterminated")
	res, err := d.Parse(malformed, "yaml")
	if err == nil || res.OK {
		t.Fatalf("malformed doc accepted: ok=%t err=%v", res.OK, err)
	}
	if res.RuleID == "" {
		t.Error("rejection must carry a CES-PARSE rule id")
	}
	if _, verr := d.Validate(malformed); verr == nil {
		t.Error("Validate must propagate parse errors on malformed input")
	}

	// A well-formed YAML document that is not a CES experiment is rejected by
	// the version/kind rules through both adapter methods.
	nonCES := []byte("apiVersion: wrong/v9\nkind: Nope\nmetadata:\n  name: x\n")
	if r, perr := d.Parse(nonCES, "auto"); perr == nil || r.OK {
		t.Errorf("non-CES doc accepted by Parse: ok=%t err=%v", r.OK, perr)
	}
	if _, verr := d.Validate(nonCES); verr == nil {
		t.Error("non-CES doc accepted by Validate")
	}

	// The JSON hint path is honored: a non-CES JSON document is still rejected.
	jsonDoc := []byte(`{"apiVersion":"wrong/v9","kind":"Nope"}`)
	if r, perr := d.Parse(jsonDoc, "json"); perr == nil || r.OK {
		t.Errorf("non-CES JSON accepted by Parse: ok=%t err=%v", r.OK, perr)
	}
}

// TestChaosDriverNormalizeErrorBranches builds documents that decode cleanly but
// violate Normalize invariants (dangling probe ref, non-positive interval). This
// reaches the Normalize-error branch of both Parse and Validate, which a fully
// valid conformance document can never exercise.
func TestChaosDriverNormalizeErrorBranches(t *testing.T) {
	d := Driver{}
	// Rename the declared probe id in the marshaled YAML so the steady-state
	// hypothesis references a probe that no longer exists.
	doc, err := baseBuilder().Marshal()
	if err != nil {
		t.Fatalf("marshal base: %v", err)
	}
	dangling := []byte(strings.Replace(string(doc), "id: p1", "id: renamed", 1))
	if res, perr := d.Parse(dangling, "yaml"); perr == nil || res.OK {
		t.Errorf("dangling-probe doc accepted by Parse: ok=%t err=%v", res.OK, perr)
	}
	if _, verr := d.Validate(dangling); verr == nil {
		t.Error("dangling-probe doc accepted by Validate (Normalize error not propagated)")
	}

	// A non-positive probe interval is another Normalize rejection. VirtualDuration
	// marshals as a duration string ("1s"), so lower it to "0s" in the YAML text.
	zeroInterval := []byte(strings.Replace(string(doc), "interval: 1s", "interval: 0s", 1))
	if string(zeroInterval) == string(doc) {
		t.Fatal("interval replacement did not apply; check marshaled field name")
	}
	if _, verr := d.Validate(zeroInterval); verr == nil {
		t.Error("zero-interval doc accepted by Validate (Normalize error not propagated)")
	}
}

// TestChaosBuilderAdversarial exercises the producer across many deterministic
// field combinations, asserting the assembled document always marshals and
// round-trips through the public pipeline without unexpected violations. It is
// the builder's robustness contract under repeated/edge-case inputs.
func TestChaosBuilderAdversarial(t *testing.T) {
	d := Driver{}
	for i := 0; i < 8; i++ {
		b := NewBuilder("exp-chaos", "node-"+string(rune('a'+i))).
			WithSteadyState("p", "ok", int64(i+1)).
			WithMetricProbe("p", "up", "1", int64(i+1)).
			WithMethodExpression("ε")
		doc, err := b.Marshal()
		if err != nil {
			t.Fatalf("iteration %d marshal: %v", i, err)
		}
		if len(doc) == 0 {
			t.Fatalf("iteration %d produced empty document", i)
		}
		// The built document must be a valid CES experiment with the target
		// selector and method expression the builder was given.
		if !strings.Contains(string(doc), "node-"+string(rune('a'+i))) {
			t.Errorf("iteration %d lost the node selector", i)
		}
		viols, verr := d.Validate(doc)
		if verr != nil {
			t.Fatalf("iteration %d validate: %v", i, verr)
		}
		if len(viols) != 0 {
			t.Errorf("iteration %d produced %d violations", i, len(viols))
		}
	}
}
