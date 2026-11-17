/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"path/filepath"
	"strings"
	"testing"

	"abate/lang/rule"
)

// TestSeverityMappingCoversEveryBranch pins the SARIF level/kind mapping,
// including the conservative fallback for an unknown severity.
func TestSeverityMappingCoversEveryBranch(t *testing.T) {
	cases := []struct {
		severity rule.Severity
		level    string
		kind     string
	}{
		{rule.SeverityError, "error", "fail"},
		{rule.SeverityWarning, "warning", "review"},
		{rule.SeverityInfo, "note", "informational"},
		{rule.Severity("unclassified"), "warning", "review"},
	}
	for _, c := range cases {
		if got := sarifLevel(c.severity); got != c.level {
			t.Fatalf("sarifLevel(%q) = %q, want %q", c.severity, got, c.level)
		}
		if got := sarifKind(c.severity); got != c.kind {
			t.Fatalf("sarifKind(%q) = %q, want %q", c.severity, got, c.kind)
		}
	}
}

// TestFindingLessIsATotalOrder walks every comparison branch of the canonical
// finding order so the SARIF encoder cannot regress into a partial order.
func TestFindingLessIsATotalOrder(t *testing.T) {
	base := rule.Finding{
		RuleID: "RULE-A", Path: "a.go", Line: 1, Column: 1, Message: "m",
	}
	variants := []rule.Finding{
		{RuleID: "RULE-B", Path: "a.go", Line: 1, Column: 1, Message: "m"},
		{RuleID: "RULE-A", Path: "b.go", Line: 1, Column: 1, Message: "m"},
		{RuleID: "RULE-A", Path: "a.go", Line: 2, Column: 1, Message: "m"},
		{RuleID: "RULE-A", Path: "a.go", Line: 1, Column: 2, Message: "m"},
		{RuleID: "RULE-A", Path: "a.go", Line: 1, Column: 1, Message: "n"},
	}
	for i, v := range variants {
		if !findingLess(base, v) {
			t.Fatalf("variant %d: base should sort before the variant", i)
		}
		if findingLess(v, base) {
			t.Fatalf("variant %d: comparison is not antisymmetric", i)
		}
	}
	if findingLess(base, base) {
		t.Fatal("findingLess must be irreflexive")
	}
}

// TestCanonicalOrderIsStableAcrossPermutations proves sorting by the canonical
// key yields one byte-stable document for every input permutation.
func TestCanonicalOrderIsStableAcrossPermutations(t *testing.T) {
	findings := sampleFindings()
	baseline, err := BuildSARIF(findings)
	if err != nil {
		t.Fatalf("BuildSARIF: %v", err)
	}
	for shift := 0; shift < len(findings); shift++ {
		rotated := make([]rule.Finding, 0, len(findings))
		for i := range findings {
			rotated = append(rotated, findings[(i+shift)%len(findings)])
		}
		got, err := BuildSARIF(rotated)
		if err != nil {
			t.Fatalf("BuildSARIF rotation %d: %v", shift, err)
		}
		if string(got) != string(baseline) {
			t.Fatalf("rotation %d changed the SARIF bytes", shift)
		}
	}
}

// TestValidateAgainstSchemaIOBoundaries covers the error paths of the loader
// boundary: unreadable schema, malformed document and schema violation.
func TestValidateAgainstSchemaIOBoundaries(t *testing.T) {
	valid, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF: %v", err)
	}
	missing := filepath.Join(t.TempDir(), "absent-schema.json")
	if err := ValidateSARIFAgainst(valid, missing); err == nil {
		t.Fatal("ValidateSARIFAgainst accepted a missing schema file")
	}
	if err := ValidateSARIFAgainst([]byte("{not json"), missing); err == nil {
		t.Fatal("ValidateSARIFAgainst accepted malformed input with a missing schema")
	}
	path, ok := findOASISSchema()
	if !ok {
		t.Fatal("bundled OASIS schema is not reachable from the report package")
	}
	if err := ValidateSARIFAgainst([]byte("{not json"), path); err == nil {
		t.Fatal("ValidateSARIFAgainst accepted malformed JSON")
	}
	if err := ValidateSARIFAgainst([]byte(`{"version":"2.1.0"}`), path); err == nil {
		t.Fatal("ValidateSARIFAgainst accepted a document missing required runs")
	}
	if err := ValidateSARIFAgainst(valid, path); err != nil {
		t.Fatalf("valid SARIF rejected: %v", err)
	}
}

// TestValidateAgainstOASISMissingBundledSchema proves the auto-discovery path
// fails closed when the repository reference is not reachable.
func TestValidateAgainstOASISMissingBundledSchema(t *testing.T) {
	t.Chdir(t.TempDir())
	valid, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF: %v", err)
	}
	err = ValidateSARIFAgainstOASIS(valid)
	if err == nil || !strings.Contains(err.Error(), "cannot locate OASIS schema") {
		t.Fatalf("expected a missing-schema error, got %v", err)
	}
}
