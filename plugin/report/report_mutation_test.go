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
	"testing"
)

// TestBuildSARIFFindsMutation shows the encoder is sensitive to its inputs: a
// single changed finding changes the emitted bytes, while an empty input still
// yields a schema-valid document.
func TestBuildSARIFFindsMutation(t *testing.T) {
	base, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF base: %v", err)
	}
	if err := ValidateSARIFAgainstOASIS(base); err != nil {
		t.Fatalf("base SARIF invalid: %v", err)
	}

	// Mutate one finding and confirm the output changes.
	mut := sampleFindings()
	mut[0].Message = "mutated message text"
	changed, err := BuildSARIF(mut)
	if err != nil {
		t.Fatalf("BuildSARIF mutated: %v", err)
	}
	if hex.EncodeToString(base) == hex.EncodeToString(changed) {
		t.Fatal("changing a finding did not change the SARIF output")
	}
	if err := ValidateSARIFAgainstOASIS(changed); err != nil {
		t.Fatalf("mutated SARIF invalid: %v", err)
	}

	// An empty finding set must still produce a valid, stable document.
	empty, err := BuildSARIF(nil)
	if err != nil {
		t.Fatalf("BuildSARIF empty: %v", err)
	}
	if err := ValidateSARIFAgainstOASIS(empty); err != nil {
		t.Fatalf("empty SARIF invalid: %v", err)
	}
}

// TestOASISRejectsInvalidDocument proves the validator is not vacuous: a document
// that violates required fields or additionalProperties is rejected.
func TestOASISRejectsInvalidDocument(t *testing.T) {
	valid, err := BuildSARIF(sampleFindings())
	if err != nil {
		t.Fatalf("BuildSARIF: %v", err)
	}
	if err := ValidateSARIFAgainstOASIS(valid); err != nil {
		t.Fatalf("valid SARIF wrongly rejected: %v", err)
	}

	// Missing required "version" must fail.
	missingVersion := `{"$schema":"https://example.com/sarif.json","runs":[]}`
	if err := ValidateSARIFAgainstOASIS([]byte(missingVersion)); err == nil {
		t.Fatal("validator accepted a document missing required version")
	}

	// additionalProperties:false violation at the root must fail.
	extraProp := `{"version":"2.1.0","runs":[],"bogusExtraProperty":true}`
	if err := ValidateSARIFAgainstOASIS([]byte(extraProp)); err == nil {
		t.Fatal("validator accepted a document with an extra root property")
	}

	// A non-object value cannot satisfy the root object schema.
	if err := ValidateSARIFAgainstOASIS([]byte(`"not-an-object"`)); err == nil {
		t.Fatal("validator accepted a non-object root")
	}
}

// TestBuildJSONSensitiveToInputs proves the JSON reporter responds to input changes
// and that empty inputs remain valid.
func TestBuildJSONSensitiveToInputs(t *testing.T) {
	base, err := BuildJSON(sampleFindings(), sampleReport())
	if err != nil {
		t.Fatalf("BuildJSON base: %v", err)
	}
	mut := sampleFindings()
	mut[1].Line = 999
	changed, err := BuildJSON(mut, sampleReport())
	if err != nil {
		t.Fatalf("BuildJSON mutated: %v", err)
	}
	if sha256.Sum256(base) == sha256.Sum256(changed) {
		t.Fatal("changing a finding did not change the JSON output")
	}
	empty, err := BuildJSON(nil, sampleReport())
	if err != nil {
		t.Fatalf("BuildJSON empty: %v", err)
	}
	if len(empty) == 0 {
		t.Fatal("empty JSON report produced no bytes")
	}
}
