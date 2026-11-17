/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

import (
	"errors"
	"strings"
	"testing"
)

// Mutation testing for the neutral comparison engine: each "mutant" is a driver
// whose observed outcome diverges from a case's expectation in exactly one way.
// A correct runner must kill every mutant (report Failed with the right detail).
// This exercises the mismatch branches that the all-passing embedded library can
// never reach, and proves the TCK actually detects non-compliant implementations.

func runOne(t *testing.T, d Driver, c Case) (CaseResult, Report) {
	t.Helper()
	rep, err := (Runner{Driver: d, Name: "mutant"}).Run([]Case{c})
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if len(rep.Details) != 1 {
		t.Fatalf("expected 1 detail, got %d", len(rep.Details))
	}
	return rep.Details[0], rep
}

func mustKill(t *testing.T, name, detailWant string, d Driver, c Case) {
	t.Helper()
	res, rep := runOne(t, d, c)
	if res.Passed || rep.Failed != 1 || rep.Compliant {
		t.Errorf("mutant %q survived: passed=%t failed=%d compliant=%t detail=%q",
			name, res.Passed, rep.Failed, rep.Compliant, res.Detail)
	}
	if detailWant != "" && !strings.Contains(res.Detail, detailWant) {
		t.Errorf("mutant %q detail %q missing %q", name, res.Detail, detailWant)
	}
}

// TestMutationParseBranches kills parse mutants and covers the accept-mismatch
// (with and without an error), the rule-id mismatch and the accept-pass paths.
func TestMutationParseBranches(t *testing.T) {
	// Expect accept, driver rejects with an error -> detail must be err text.
	mustKill(t, "parse reject with err", "scripted rejection",
		scriptedDriver{parseOK: false, parseRule: "CES-PARSE-006", parseErr: errors.New("scripted rejection")},
		Case{ID: "M1", Spec: "s", Kind: KindParse, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	// Expect accept, driver rejects without an error -> generic mismatch detail.
	mustKill(t, "parse reject no err", "parse accept mismatch",
		scriptedDriver{parseOK: false, parseRule: "CES-PARSE-006"},
		Case{ID: "M2", Spec: "s", Kind: KindParse, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	// Expect reject with rule R, driver rejects with a different rule -> mismatch.
	mustKill(t, "parse wrong rule", "parse rule id mismatch",
		scriptedDriver{parseOK: false, parseRule: "CES-PARSE-999", parseErr: errors.New("x")},
		Case{ID: "M3", Spec: "s", Kind: KindParse, Input: "a: 1",
			Expect: CaseExpect{Accept: false, RuleIDs: []string{"CES-PARSE-006"}}})
	// Expect accept, driver accepts -> survivor (this is the compliant path).
	res, rep := runOne(t, scriptedDriver{parseOK: true},
		Case{ID: "M4", Spec: "s", Kind: KindParse, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	if !res.Passed || rep.Passed != 1 || !rep.Compliant {
		t.Fatalf("compliant parse case was wrongly killed: %+v", rep)
	}
}

// TestMutationValidateBranches kills validate mutants and covers the returned
// error, accept-mismatch, rule-id-set mismatch and the compliant paths.
func TestMutationValidateBranches(t *testing.T) {
	// Driver returns an error -> "validate returned error".
	mustKill(t, "validate error", "validate returned error",
		scriptedDriver{valErr: errors.New("boom")},
		Case{ID: "V1", Spec: "s", Kind: KindValidate, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	// Expect accept, driver produces a violation -> accept mismatch.
	mustKill(t, "unexpected violation", "validate accept mismatch",
		scriptedDriver{viols: []Violation{{RuleID: "CES-VALID-X"}}},
		Case{ID: "V2", Spec: "s", Kind: KindValidate, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	// Expect a specific rule set, driver produces another -> set mismatch.
	mustKill(t, "wrong violation set", "validate rule id set mismatch",
		scriptedDriver{viols: []Violation{{RuleID: "CES-VALID-A"}, {RuleID: "CES-VALID-B"}}},
		Case{ID: "V3", Spec: "s", Kind: KindValidate, Input: "a: 1",
			Expect: CaseExpect{Accept: false, RuleIDs: []string{"CES-VALID-A"}}})
	// Matching violation set (order independent) -> compliant path.
	res, rep := runOne(t, scriptedDriver{viols: []Violation{{RuleID: "A"}, {RuleID: "B"}}},
		Case{ID: "V4", Spec: "s", Kind: KindValidate, Input: "a: 1",
			Expect: CaseExpect{Accept: false, RuleIDs: []string{"B", "A"}}})
	if !res.Passed || !rep.Compliant {
		t.Fatalf("order-independent violation set wrongly killed: %+v", rep)
	}
}

// flakyParseDriver accepts the first Parse call and rejects the second, so the
// round-trip check's "second parse diverged" branch becomes reachable.
type flakyParseDriver struct{ calls int }

func (d *flakyParseDriver) Parse(_ []byte, _ string) (ParseResult, error) {
	d.calls++
	if d.calls == 1 {
		return ParseResult{OK: true}, nil
	}
	return ParseResult{OK: false, RuleID: "CES-PARSE-006"}, errors.New("diverged")
}

func (*flakyParseDriver) Validate(_ []byte) ([]Violation, error) { return nil, nil }

// TestMutationRoundtripBranches covers document-fails-to-parse, second-parse
// divergence, expected-reject-but-parsed, and the compliant round-trip path.
func TestMutationRoundtripBranches(t *testing.T) {
	// First parse fails.
	mustKill(t, "roundtrip first parse fails", "roundtrip document failed to parse",
		scriptedDriver{parseOK: false, parseRule: "CES-PARSE-006", parseErr: errors.New("bad")},
		Case{ID: "R1", Spec: "s", Kind: KindRoundtrip, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	// Second parse diverges.
	mustKill(t, "roundtrip second parse diverges", "roundtrip second parse diverged",
		&flakyParseDriver{},
		Case{ID: "R2", Spec: "s", Kind: KindRoundtrip, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	// Expect reject but the document actually parses.
	mustKill(t, "roundtrip expected reject", "roundtrip expected reject",
		scriptedDriver{parseOK: true},
		Case{ID: "R3", Spec: "s", Kind: KindRoundtrip, Input: "a: 1", Expect: CaseExpect{Accept: false}})
	// Compliant round-trip.
	res, rep := runOne(t, scriptedDriver{parseOK: true},
		Case{ID: "R4", Spec: "s", Kind: KindRoundtrip, Input: "a: 1", Expect: CaseExpect{Accept: true}})
	if !res.Passed || !rep.Compliant {
		t.Fatalf("compliant roundtrip wrongly killed: %+v", rep)
	}
}

// TestMutationAbateDriverErrorPaths drives the real bundled driver through
// malformed inputs to cover its parse/normalize error branches and the neutral
// violation mapping, proving ruleIDOf is applied on real structured errors.
func TestMutationAbateDriverErrorPaths(t *testing.T) {
	d := AbateDriver{}
	// Structurally invalid YAML must be rejected by Parse with a non-empty rule id.
	res, err := d.Parse([]byte("this: [is: not: yaml"), "yaml")
	if err == nil || res.OK {
		t.Fatalf("malformed document was accepted: ok=%t err=%v", res.OK, err)
	}
	if res.RuleID == "" {
		t.Error("rejected document must carry a CES-PARSE rule id")
	}
	// Validate on malformed input must return the parse error.
	if _, verr := d.Validate([]byte("\tbad: : yaml")); verr == nil {
		t.Error("Validate must propagate parse errors on malformed input")
	}
	// A syntactically valid but non-CES document is rejected by the pipeline,
	// exercising the version/kind rule branches through both methods.
	nonCES := []byte("apiVersion: wrong/v9\nkind: Nope\nmetadata:\n  name: x\n")
	if r, perr := d.Parse(nonCES, "auto"); perr == nil || r.OK {
		t.Errorf("non-CES document accepted by Parse: ok=%t err=%v", r.OK, perr)
	}
	if _, verr := d.Validate(nonCES); verr == nil {
		t.Error("non-CES document accepted by Validate")
	}
	// JSON hint path is honored and still rejects a non-CES JSON document.
	jsonDoc := []byte(`{"apiVersion":"wrong/v9","kind":"Nope"}`)
	if r, perr := d.Parse(jsonDoc, "json"); perr == nil || r.OK {
		t.Errorf("non-CES JSON accepted by Parse: ok=%t err=%v", r.OK, perr)
	}
}
