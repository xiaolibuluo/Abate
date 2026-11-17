/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"sync"
	"testing"

	"abate/internal/determinism"
)

// TestSameSeed100UniqueReportHash runs the whole pipeline 100 times under one
// seed and proves the report hash stays unique.
func TestSameSeed100UniqueReportHash(t *testing.T) {
	t.Parallel()
	unit := sampleUnit(t)
	rules := builtinRules(t)
	seen := make(map[[32]byte]bool, 1)
	for i := 0; i < 100; i++ {
		rep, _, err := Experiment(42, unit, rules)
		if err != nil {
			t.Fatalf("run %d: %v", i, err)
		}
		seen[determinism.HashState(rep.Bytes())] = true
	}
	if len(seen) != 1 {
		t.Fatalf("100 same-seed runs produced %d distinct report hashes", len(seen))
	}
}

// TestConcurrentDeriveAndRun proves a shared Deriver, unit and rule set are
// safe under concurrent derivations and engine runs.
func TestConcurrentDeriveAndRun(t *testing.T) {
	t.Parallel()
	unit := sampleUnit(t)
	rules := builtinRules(t)
	const workers = 24
	var wg sync.WaitGroup
	reportHashes := make(chan [32]byte, workers)
	specHashes := make(chan [32]byte, workers)
	failures := make(chan error, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			specs, err := NewDeriver().Derive(rules)
			if err != nil {
				failures <- err
				return
			}
			specHashes <- determinism.HashState(specs)
			rep, _, err := Experiment(42, unit, rules)
			if err != nil {
				failures <- err
				return
			}
			reportHashes <- determinism.HashState(rep.Bytes())
		}()
	}
	wg.Wait()
	close(failures)
	close(reportHashes)
	close(specHashes)
	for err := range failures {
		t.Fatalf("concurrent run: %v", err)
	}
	uniqueReports := make(map[[32]byte]bool, workers)
	for hash := range reportHashes {
		uniqueReports[hash] = true
	}
	if len(uniqueReports) != 1 {
		t.Fatalf("concurrent runs produced %d distinct report hashes", len(uniqueReports))
	}
	uniqueSpecs := make(map[[32]byte]bool, workers)
	for hash := range specHashes {
		uniqueSpecs[hash] = true
	}
	if len(uniqueSpecs) != 1 {
		t.Fatalf("concurrent derives produced %d distinct spec hashes", len(uniqueSpecs))
	}
}

// TestDifferentSeedsDiverge proves the seed genuinely drives point selection
// while each individual seed remains byte-stable.
func TestDifferentSeedsDiverge(t *testing.T) {
	t.Parallel()
	unit := sampleUnit(t)
	rules := builtinRules(t)
	first, _, err := Experiment(1, unit, rules)
	if err != nil {
		t.Fatalf("seed 1: %v", err)
	}
	second, _, err := Experiment(2, unit, rules)
	if err != nil {
		t.Fatalf("seed 2: %v", err)
	}
	if string(first.Bytes()) == string(second.Bytes()) {
		t.Fatal("two different seeds produced identical report bytes")
	}
	determinism.MustMatch100(t, 1, func() []byte { return first.Bytes() })
	determinism.MustMatch100(t, 2, func() []byte { return second.Bytes() })
}
