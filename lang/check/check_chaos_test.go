/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"sync"
	"testing"

	"abate/lang/gofront"
	"abate/lang/rule"
)

// TestChaosCheckFingerprint proves 100 sequential and concurrent checks of the
// same sources produce one result fingerprint and identical report bytes.
func TestChaosCheckFingerprint(t *testing.T) {
	t.Parallel()
	path := sourceFile(t, "chaos.go", "package p\nvar foo_bar int\n")
	current := unitRule(t, "GO-EG-001", "Ident~foo_bar", rule.SeverityWarning)
	checker := NewChecker(gofront.NewGoFrontend().WithTypeCheck(false),
		[]rule.Rule{current})
	options := DefaultOptions()
	options.Format = FormatSARIF
	base, err := checker.Check([]string{path}, options)
	if err != nil {
		t.Fatalf("baseline check: %v", err)
	}
	report := base.Render(options.Format)
	for i := 0; i < 100; i++ {
		got, checkErr := checker.Check([]string{path}, options)
		if checkErr != nil {
			t.Fatalf("iteration %d: %v", i, checkErr)
		}
		if got.Fingerprint != base.Fingerprint ||
			string(got.Render(options.Format)) != string(report) {
			t.Fatalf("check output changed at iteration %d", i)
		}
	}
	var group sync.WaitGroup
	fingerprints := make(chan [32]byte, 24)
	for i := 0; i < 24; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			got, checkErr := checker.Check([]string{path}, options)
			if checkErr != nil {
				fingerprints <- [32]byte{}
				return
			}
			fingerprints <- got.Fingerprint
		}()
	}
	group.Wait()
	close(fingerprints)
	for value := range fingerprints {
		if value != base.Fingerprint {
			t.Fatal("concurrent check produced a different fingerprint")
		}
	}
}
