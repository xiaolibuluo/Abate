/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/rule"
)

// sampleFindings returns a small, deterministic set of findings used across tests.
func sampleFindings() []rule.Finding {
	return []rule.Finding{
		{
			RuleID:    rule.RuleID("GO-EG-001"),
			Path:      "internal/server/handler.go",
			Line:      42,
			Column:    7,
			Severity:  rule.SeverityError,
			Category:  rule.CategoryErrorProne,
			Message:   "error not checked",
			Reference: "EG-001",
		},
		{
			RuleID:    rule.RuleID("GO-VET-003"),
			Path:      "cmd/root.go",
			Line:      10,
			Column:    3,
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryConcurrency,
			Message:   "possible nil dereference",
			Reference: "VET-003",
		},
		{
			RuleID:    rule.RuleID("GO-CRC-014"),
			Path:      "pkg/crc/crc.go",
			Line:      88,
			Column:    1,
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Message:   "odd length crc table",
			Reference: "CRC-014",
		},
	}
}

func sampleReport() mutation.Report {
	var tf, sf, fp [32]byte
	copy(tf[:], []byte("target-fingerprint-padding......"))
	copy(sf[:], []byte("suite-fingerprint-padding......."))
	copy(fp[:], []byte("report-fingerprint-padding....."))
	rep := mutation.Report{
		Seed:              0xC0FFEE,
		TargetFingerprint: tf,
		SuiteFingerprint:  sf,
		Fingerprint:       fp,
		Score:             verify.Score{Total: 10, Killed: 7, Survived: 2, Equivalent: 1, Timeout: 0},
	}
	return rep
}

func TestBuildSARIFValidatesAgainstOASIS(t *testing.T) {
	data, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF: %v", err)
	}
	if err := ValidateSARIFAgainstOASIS(data); err != nil {
		t.Fatalf("generated SARIF failed OASIS validation: %v", err)
	}
}

func TestBuildSARIFRequiredFields(t *testing.T) {
	data, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF: %v", err)
	}
	var doc map[string]any
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if doc["version"] != "2.1.0" {
		t.Fatalf("version = %v, want 2.1.0", doc["version"])
	}
	if _, ok := doc["$schema"]; !ok {
		t.Fatal("missing $schema field")
	}
	runs, ok := doc["runs"].([]any)
	if !ok || len(runs) < 1 {
		t.Fatalf("expected at least one run, got %v", doc["runs"])
	}
	run, ok := runs[0].(map[string]any)
	if !ok {
		t.Fatalf("runs[0] is not a map: %T", runs[0])
	}
	tool, ok := run["tool"].(map[string]any)
	if !ok {
		t.Fatalf("run[tool] is not a map: %T", run["tool"])
	}
	driver, ok := tool["driver"].(map[string]any)
	if !ok {
		t.Fatalf("tool[driver] is not a map: %T", tool["driver"])
	}
	if driver["name"] != "Abate RPT" {
		t.Fatalf("driver name = %v", driver["name"])
	}
	results, ok := run["results"].([]any)
	if !ok {
		t.Fatalf("run[results] is not a slice: %T", run["results"])
	}
	if len(results) != len(sampleFindings()) {
		t.Fatalf("results count = %d, want %d", len(results), len(sampleFindings()))
	}
	// The schema URI must be an absolute URI as required by the format keyword.
	if !strings.HasPrefix(doc["$schema"].(string), "https://") {
		t.Fatalf("schema uri not absolute: %v", doc["$schema"])
	}
}

func TestBuildSARIFDeterministic(t *testing.T) {
	a, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF a: %v", err)
	}
	b, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF b: %v", err)
	}
	if hex.EncodeToString(a) != hex.EncodeToString(b) {
		t.Fatal("SARIF output is not byte-deterministic for identical inputs")
	}
	// A different finding order must still yield identical bytes (canonical sort).
	rev := append([]rule.Finding(nil), sampleFindings()...)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	c, err := BuildSARIF(rev)
	if err != nil {
		t.Fatalf("BuildSARIF c: %v", err)
	}
	if hex.EncodeToString(a) != hex.EncodeToString(c) {
		t.Fatal("SARIF output changes with input ordering (must be canonicalized)")
	}
}

func TestBuildJSONDeterministic(t *testing.T) {
	a, err := BuildJSON(sampleFindings(), sampleReport())
	if err != nil {
		t.Fatalf("BuildJSON a: %v", err)
	}
	b, err := BuildJSON(sampleFindings(), sampleReport())
	if err != nil {
		t.Fatalf("BuildJSON b: %v", err)
	}
	sumA := sha256.Sum256(a)
	sumB := sha256.Sum256(b)
	if sumA != sumB {
		t.Fatal("JSON report is not byte-deterministic for identical inputs")
	}
	var doc jsonDocument
	if err := json.Unmarshal(a, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if doc.Report.Seed != 0xC0FFEE {
		t.Fatalf("report seed = %d, want %d", doc.Report.Seed, 0xC0FFEE)
	}
	if doc.Report.Score.Killed != 7 {
		t.Fatalf("killed = %d, want 7", doc.Report.Score.Killed)
	}
	if len(doc.Findings) != len(sampleFindings()) {
		t.Fatalf("findings = %d, want %d", len(doc.Findings), len(sampleFindings()))
	}
}
