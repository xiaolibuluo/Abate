/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"sync"
	"testing"
)

// TestCompareIsStableOver100Runs certifies that one hundred comparisons over
// the frozen golden produce exactly one JSON document.
func TestCompareIsStableOver100Runs(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	unique := make(map[string]int)
	for i := 0; i < 100; i++ {
		report, err := Compare(records)
		if err != nil {
			t.Fatalf("run %d failed: %v", i, err)
		}
		payload, err := report.JSON()
		if err != nil {
			t.Fatalf("run %d JSON failed: %v", i, err)
		}
		unique[string(payload)]++
	}
	if len(unique) != 1 {
		t.Fatalf("the hundred runs produced %d distinct reports", len(unique))
	}
	for payload, count := range unique {
		t.Logf("GOLDEN-REGRESSION report bytes=%d count=%d", len(payload), count)
	}
}

// TestConcurrentCompareIsSafe proves the harness holds no shared state.
func TestConcurrentCompareIsSafe(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	const workers = 16
	results := make(chan string, workers)
	errs := make(chan error, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			report, err := Compare(records)
			if err != nil {
				errs <- err
				return
			}
			payload, err := report.JSON()
			if err != nil {
				errs <- err
				return
			}
			results <- string(payload)
		}()
	}
	wg.Wait()
	close(results)
	close(errs)
	for err := range errs {
		t.Fatalf("concurrent compare failed: %v", err)
	}
	seen := ""
	for payload := range results {
		if seen == "" {
			seen = payload
			continue
		}
		if payload != seen {
			t.Fatal("concurrent comparisons disagree")
		}
	}
	if seen == "" {
		t.Fatal("no concurrent comparison produced a report")
	}
}

// TestFixtureNeverLeaksItsState proves the fixture stays immutable.
func TestFixtureNeverLeaksItsState(t *testing.T) {
	t.Parallel()
	for i := 0; i < 50; i++ {
		target, suite := newFixture()
		nodes := target.Nodes()
		if len(nodes) != 10 {
			t.Fatalf("iteration %d: fixture holds %d nodes", i, len(nodes))
		}
		nodes[0] = "tampered"
		if target.Nodes()[0] == "tampered" {
			t.Fatalf("iteration %d: Nodes() leaked its backing array", i)
		}
		tests := suite.Tests()
		tests[0] = "tampered"
		if suite.Tests()[0] == "tampered" {
			t.Fatalf("iteration %d: Tests() leaked its backing array", i)
		}
		digest := target.Fingerprint()
		if digest != target.Fingerprint() {
			t.Fatalf("iteration %d: the target fingerprint drifted", i)
		}
	}
}
