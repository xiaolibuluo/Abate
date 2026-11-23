/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutation

import (
	"context"
	"errors"
	"testing"
)

// TestEngineParallelMatchesSequential is the core determinism guarantee of
// the worker-pool engine: for one seed, target and suite, judging every
// point concurrently must produce the byte-identical report of the
// sequential engine, because results are written back by scheduled index.
func TestEngineParallelMatchesSequential(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	want, err := NewEngine(42).Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("sequential Run failed: %v", err)
	}
	for _, workers := range []int{2, 3, 8, 64} {
		got, err := NewEngineParallel(42, workers).Run(context.Background(), target, suite)
		if err != nil {
			t.Fatalf("parallel Run (workers=%d) failed: %v", workers, err)
		}
		if got.Fingerprint != want.Fingerprint {
			t.Fatalf("workers=%d: report fingerprint differs from sequential", workers)
		}
		if reportSHA(got) != reportSHA(want) {
			t.Fatalf("workers=%d: canonical bytes differ from sequential", workers)
		}
		if got.Score != want.Score {
			t.Fatalf("workers=%d: score %+v, want %+v", workers, got.Score, want.Score)
		}
		if len(got.Verdicts) != len(want.Verdicts) || len(got.KillingTests) != len(want.KillingTests) {
			t.Fatalf("workers=%d: report lengths differ from sequential", workers)
		}
		for i := range want.Verdicts {
			if got.Verdicts[i] != want.Verdicts[i] || got.KillingTests[i] != want.KillingTests[i] {
				t.Fatalf("workers=%d: verdict %d differs: got (%s, %q), want (%s, %q)",
					workers, i, got.Verdicts[i], got.KillingTests[i],
					want.Verdicts[i], want.KillingTests[i])
			}
		}
	}
}

// TestEngineParallelRunTwice proves the parallel engine is byte-stable under
// replay: two runs of one seed must agree, which is the property the Report
// fingerprint exists to certify.
func TestEngineParallelRunTwice(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	first, err := NewEngineParallel(7, 4).Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("first Run failed: %v", err)
	}
	second, err := NewEngineParallel(7, 4).Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("second Run failed: %v", err)
	}
	if reportSHA(first) != reportSHA(second) {
		t.Fatal("parallel engine is not byte-stable across two runs of one seed")
	}
}

// TestEngineParallelWorkerFallback proves workers below two keeps the
// sequential semantics and accessors, so callers may pass 0 or a negative
// value to mean "no concurrency".
func TestEngineParallelWorkerFallback(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	want, err := NewEngine(42).Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("sequential Run failed: %v", err)
	}
	for _, workers := range []int{0, -1, 1} {
		engine := NewEngineParallel(42, workers)
		if engine.Seed() != 42 || engine.Registry() == nil {
			t.Fatalf("workers=%d: accessors disagree with construction", workers)
		}
		got, err := engine.Run(context.Background(), target, suite)
		if err != nil {
			t.Fatalf("workers=%d: Run failed: %v", workers, err)
		}
		if reportSHA(got) != reportSHA(want) {
			t.Fatalf("workers=%d: fallback report differs from sequential", workers)
		}
	}
}

// TestEngineParallelContextErrors proves the concurrent path honors context
// cancellation exactly like the sequential loop: a canceled context aborts
// the run with the context error, never a fabricated report.
func TestEngineParallelContextErrors(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	engine := NewEngineParallel(7, 4)
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := engine.Run(canceled, target, suite); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled context error = %v, want context.Canceled", err)
	}
	var missing context.Context
	if _, err := engine.Run(missing, target, suite); !errors.Is(err, ErrCanceled) {
		t.Fatalf("nil context error = %v, want ErrCanceled", err)
	}
}
