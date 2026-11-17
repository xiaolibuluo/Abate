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
	"sync"
	"testing"

	"abate/lang/rule"
)

// TestBuildSARIFDeterministic100 proves the encoder is reproducible across many
// iterations with no hidden state or random source.
func TestBuildSARIFDeterministic100(t *testing.T) {
	findings := sampleFindings()
	first, err := BuildSARIF(findings)
	if err != nil {
		t.Fatalf("BuildSARIF: %v", err)
	}
	firstHex := hex.EncodeToString(sha256.New().Sum(append([]byte(nil), first...)))
	for i := 0; i < 100; i++ {
		got, err := BuildSARIF(findings)
		if err != nil {
			t.Fatalf("iteration %d: %v", i, err)
		}
		if hex.EncodeToString(sha256.New().Sum(append([]byte(nil), got...))) != firstHex {
			t.Fatalf("iteration %d: SARIF output drifted", i)
		}
	}
}

// TestBuildConcurrentSafeChaos exercises concurrent SARIF construction and OASIS
// validation to surface any data race in the encoder or validator.
func TestBuildConcurrentSafeChaos(t *testing.T) {
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			data, err := BuildSARIF(sampleFindings())
			if err != nil {
				t.Errorf("BuildSARIF: %v", err)
				return
			}
			if err := ValidateSARIFAgainstOASIS(data); err != nil {
				t.Errorf("concurrent OASIS validation failed: %v", err)
			}
		}()
	}
	wg.Wait()
}

// TestBuildJSONConcurrentSafeChaos proves concurrent JSON report construction is
// free of data races.
func TestBuildJSONConcurrentSafeChaos(t *testing.T) {
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if _, err := BuildJSON(sampleFindings(), sampleReport()); err != nil {
				t.Errorf("BuildJSON: %v", err)
			}
		}()
	}
	wg.Wait()
}

// ensure rule package import is used even if sampleFindings changes.
var _ = rule.RuleID("")
