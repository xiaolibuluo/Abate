/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"context"
	"sync"
	"testing"

	"abate/sim/rand"
)

// TestSameSeedReportSHAUnique100 certifies that 100 runs of one seed produce
// exactly one SHA-256 digest over the canonical report bytes.
func TestSameSeedReportSHAUnique100(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	engine := NewEngine(20261109)
	unique := make(map[string]int)
	for i := 0; i < 100; i++ {
		rep, err := engine.Run(context.Background(), target, suite)
		if err != nil {
			t.Fatalf("run %d failed: %v", i, err)
		}
		sum := reportSHA(rep)
		unique[shaHex(sum[:])]++
	}
	if len(unique) != 1 {
		t.Fatalf("same-seed runs produced %d distinct SHA-256 digests: %v", len(unique), unique)
	}
	for digest, count := range unique {
		t.Logf("MUT report SHA-256=%s count=%d", digest, count)
	}
}

func TestConcurrentRunSafe(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	engine := NewEngine(99)
	const workers = 32
	results := make(chan [32]byte, workers)
	errs := make(chan error, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rep, err := engine.Run(context.Background(), target, suite)
			if err != nil {
				errs <- err
				return
			}
			results <- reportSHA(rep)
		}()
	}
	wg.Wait()
	close(results)
	close(errs)
	for err := range errs {
		t.Fatalf("concurrent Run failed: %v", err)
	}
	var first [32]byte
	set := false
	for sum := range results {
		if !set {
			first = sum
			set = true
			continue
		}
		if sum != first {
			t.Fatal("concurrent Run digests disagree")
		}
	}
	if !set {
		t.Fatal("no concurrent Run produced a report")
	}
}

func TestDifferentSeedDiverges(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	left, err := NewEngine(1).Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("seed 1: %v", err)
	}
	right, err := NewEngine(2).Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("seed 2: %v", err)
	}
	if reportSHA(left) == reportSHA(right) {
		t.Fatal("different seeds produced identical reports; SIM-RNG is not driving selection")
	}
}

func TestConcurrentJudgeSafe(t *testing.T) {
	t.Parallel()
	target, _, tests := newFixture()
	reg := NewRegistry()
	op := reg.All()[0]
	point := NewMutationPoint("n", op.Params()["from"], op.Params()["to"], op, 1, 1)
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, _, err := Judge(reg, target, tests, point); err != nil {
				t.Errorf("concurrent Judge failed: %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestDiscoverIndependentOfNodeOrder(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	forward, err := Discover(target, suite, NewRegistry(), rand.New(5))
	if err != nil {
		t.Fatalf("forward discover: %v", err)
	}
	reversed := &syntheticTarget{
		nodes:      reverseStrings(target.Nodes()),
		tokens:     target.tokens,
		killFrom:   target.killFrom,
		equivalent: target.equivalent,
		timeout:    target.timeout,
	}
	backward, err := Discover(reversed, suite, NewRegistry(), rand.New(5))
	if err != nil {
		t.Fatalf("reversed discover: %v", err)
	}
	first := Schedule(forward)
	second := Schedule(backward)
	if len(first) != len(second) {
		t.Fatalf("node order changed point count: %d vs %d", len(first), len(second))
	}
	for i := range first {
		if first[i].Fingerprint != second[i].Fingerprint {
			t.Fatalf("node order leaked into scheduled point %d", i)
		}
	}
}

func reverseStrings(in []string) []string {
	out := make([]string, 0, len(in))
	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}
	return out
}
