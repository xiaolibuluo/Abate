/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

import (
	"errors"
	"testing"

	"abate/internal/yamlmini"
)

// TestTCKBootstrap runs the full embedded case library against the bundled
// AbateDriver and requires 100% compliance. This is the project's "eat our own
// dog food" gate: a regression in any CES parse or validate rule surfaces here.
func TestTCKBootstrap(t *testing.T) {
	cases, err := LoadCases()
	if err != nil {
		t.Fatalf("load cases: %v", err)
	}
	if len(cases) < 154 {
		t.Fatalf("case library too small: got %d want >=154", len(cases))
	}
	byKind := map[CaseKind]int{}
	for _, c := range cases {
		byKind[c.Kind]++
	}
	if byKind[KindParse] < 60 {
		t.Errorf("parse cases = %d, want >=60", byKind[KindParse])
	}
	if byKind[KindValidate] < 44 {
		t.Errorf("validate cases = %d, want >=44", byKind[KindValidate])
	}
	if byKind[KindRoundtrip] < 50 {
		t.Errorf("roundtrip cases = %d, want >=50", byKind[KindRoundtrip])
	}

	rep, err := Runner{Driver: AbateDriver{}, Name: "abate"}.Run(cases)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if !rep.Compliant {
		t.Errorf("AbateDriver not compliant: %s", rep.summary())
		for _, d := range rep.Details {
			if !d.Passed {
				t.Logf("FAIL %s [%s] expected=%+v actual=%+v :: %s", d.ID, d.Kind, d.Expected, d.Actual, d.Detail)
			}
		}
	}
	if rep.Passed != rep.Total {
		t.Errorf("passed %d of %d", rep.Passed, rep.Total)
	}
}

// TestTCKCaseQuality enforces the data-quality invariants the specification
// mandates: unique case ids, every case cites a spec clause, and every case
// parses as valid YAML.
func TestTCKCaseQuality(t *testing.T) {
	cases, err := LoadCases()
	if err != nil {
		t.Fatalf("load cases: %v", err)
	}
	seen := map[string]bool{}
	for _, c := range cases {
		if seen[c.ID] {
			t.Errorf("duplicate case id %s", c.ID)
		}
		seen[c.ID] = true
		if c.Spec == "" {
			t.Errorf("case %s missing spec citation", c.ID)
		}
		// Parse-reject cases may carry intentionally malformed documents, so the
		// YAML-validity assertion only applies to cases that should parse.
		if c.Kind != KindParse || c.Expect.Accept {
			var doc map[string]any
			if err := yamlmini.Unmarshal([]byte(c.Input), &doc); err != nil {
				t.Errorf("case %s input is not valid YAML: %v", c.ID, err)
			}
		}
	}
}

// TestTCKNeutrality proves the kit does not depend on Abate-specific behavior by
// running it against a minimal independent reference implementation. The reference
// handles a well-defined subset (valid documents plus a few canonical violations)
// and skips cases outside its scope; it must pass every case it claims to support.
func TestTCKNeutrality(t *testing.T) {
	cases, err := LoadCases()
	if err != nil {
		t.Fatalf("load cases: %v", err)
	}
	subset := make([]Case, 0, len(cases))
	for _, c := range cases {
		if referenceSupports(c) {
			subset = append(subset, c)
		}
	}
	rep, err := Runner{Driver: referenceDriver{}, Name: "reference"}.Run(subset)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if rep.Skipped == rep.Total {
		t.Fatal("reference driver skipped every case; neutrality is unproven")
	}
	if rep.Failed != 0 {
		t.Errorf("reference driver failed %d supported case(s): %s", rep.Failed, rep.summary())
		for _, d := range rep.Details {
			if !d.Passed {
				t.Logf("FAIL %s [%s] expected=%+v actual=%+v :: %s", d.ID, d.Kind, d.Expected, d.Actual, d.Detail)
			}
		}
	}
}

// referenceSupports reports whether the minimal reference implementation can
// deterministically decide a case. It covers every valid document (accept) and a
// small, deliberately independent set of canonical rejections, and declines the
// rest so neutrality is demonstrated over a meaningful but bounded subset.
func referenceSupports(c Case) bool {
	switch c.Kind {
	case KindRoundtrip:
		return true // every roundtrip document is a valid, lossless doc.
	case KindParse:
		if c.Expect.Accept {
			return true
		}
		return containsAny(c.Expect.RuleIDs, "CES-PARSE-002", "CES-PARSE-006")
	case KindValidate:
		if c.Expect.Accept {
			return true
		}
		return containsAny(c.Expect.RuleIDs, "CES-VALID-STEADY-001", "CES-VALID-TARGET-001")
	}
	return false
}

func containsAny(set []string, want ...string) bool {
	for _, s := range set {
		for _, w := range want {
			if s == w {
				return true
			}
		}
	}
	return false
}

// referenceDriver is a deliberately small, Abate-independent CES implementation
// used only to prove the TCK is a neutral arbiter. It reuses the shared
// YAML-subset decoder for generic documents but no Abate CES package.
type referenceDriver struct{}

func (referenceDriver) Parse(data []byte, _ string) (ParseResult, error) {
	var doc map[string]any
	if err := yamlmini.Unmarshal(data, &doc); err != nil {
		return ParseResult{OK: false, RuleID: "CES-PARSE-005", Error: err.Error()}, err
	}
	if _, ok := doc["apiVersion"]; !ok {
		return ParseResult{OK: false, RuleID: "CES-PARSE-006"}, errors.New("missing apiVersion")
	}
	if v, ok := doc["apiVersion"].(string); !ok || v != "ces.tencent.com/v1" {
		return ParseResult{OK: false, RuleID: "CES-PARSE-002"}, errors.New("unsupported version")
	}
	if v, ok := doc["kind"].(string); !ok || v != "Experiment" {
		return ParseResult{OK: false, RuleID: "CES-PARSE-006"}, errors.New("kind must be Experiment")
	}
	meta, ok := doc["metadata"].(map[string]any)
	if !ok {
		return ParseResult{OK: false, RuleID: "CES-PARSE-006"}, errors.New("metadata required")
	}
	name, ok := meta["name"].(string)
	if !ok || name == "" {
		return ParseResult{OK: false, RuleID: "CES-PARSE-006"}, errors.New("metadata.name required")
	}
	return ParseResult{OK: true}, nil
}

func (referenceDriver) Validate(data []byte) ([]Violation, error) {
	var doc map[string]any
	if err := yamlmini.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	spec, ok := doc["spec"].(map[string]any)
	if !ok {
		return nil, errors.New("missing spec")
	}
	var out []Violation
	if ss, ok := spec["steadyState"].([]any); !ok || len(ss) == 0 {
		out = append(out, Violation{RuleID: "CES-VALID-STEADY-001", FieldPath: "spec.steadyState", Severity: "error"})
	}
	if tgt, ok := spec["target"].(map[string]any); ok {
		if kind, ok := tgt["kind"].(string); ok && kind == "external" {
			out = append(out, Violation{RuleID: "CES-VALID-TARGET-001", FieldPath: "spec.target.kind", Severity: "error"})
		}
	}
	return out, nil
}

// ensure the reference driver honors the supported-set contract by skipping the
// rest; this test guards the neutrality harness itself.
func TestReferenceSupportsSubset(t *testing.T) {
	cases, err := LoadCases()
	if err != nil {
		t.Fatalf("load cases: %v", err)
	}
	supported := 0
	for _, c := range cases {
		if referenceSupports(c) {
			supported++
		}
	}
	if supported < 60 {
		t.Fatalf("reference supports only %d cases; neutrality subset too small", supported)
	}
}
