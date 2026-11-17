/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rand

import (
	"math"
	"testing"
)

// TestChaosChiSquareIntN runs a chi-square goodness-of-fit test on IntN over a
// moderate range to certify that the rejection sampler is unbiased.
func TestChaosChiSquareIntN(t *testing.T) {
	t.Parallel()
	const (
		buckets = 16
		trials  = 160000
	)
	s := New(2026)
	counts := make([]int, buckets)
	for i := 0; i < trials; i++ {
		counts[s.IntN(buckets)]++
	}
	expected := float64(trials) / float64(buckets)
	chi := 0.0
	for _, c := range counts {
		d := float64(c) - expected
		chi += d * d / expected
	}
	// Critical value for df=15 at alpha=0.001 is ~37.7; stay well under it.
	if chi > 37.7 {
		t.Fatalf("chi-square = %.2f exceeds 0.001 critical value", chi)
	}
}

// TestChaosChiSquareFloat64Bins checks Float64 uniformity across 10 bins.
func TestChaosChiSquareFloat64Bins(t *testing.T) {
	t.Parallel()
	const (
		buckets = 10
		trials  = 100000
	)
	s := New(555)
	counts := make([]int, buckets)
	for i := 0; i < trials; i++ {
		counts[int(s.Float64()*float64(buckets))]++
	}
	expected := float64(trials) / float64(buckets)
	chi := 0.0
	for _, c := range counts {
		d := float64(c) - expected
		chi += d * d / expected
	}
	if chi > 29.6 {
		t.Fatalf("chi-square = %.2f exceeds 0.001 critical value for df=9", chi)
	}
}

// TestChaosNoPanicUnderAdversarialSeeds hammers the generator with edge-case
// seeds and range arguments to ensure no panic under perturbation.
func TestChaosNoPanicUnderAdversarialSeeds(t *testing.T) {
	t.Parallel()
	seeds := []uint64{0, 1, ^uint64(0), 1 << 63, 0xDEADBEEF}
	for _, sd := range seeds {
		s := New(sd)
		for i := 0; i < 1000; i++ {
			_ = s.Uint64()
			_ = s.IntN(2)
			_ = s.IntN(1000)
			_ = s.Bool()
			_ = s.Bytes(3)
			_ = s.Float64()
		}
	}
}

// TestChaosDeterministicHash asserts the full 1000-draw hash is stable across
// 100 independent constructions from the same seed.
func TestChaosDeterministicHash(t *testing.T) {
	t.Parallel()
	first := chaosHash(777)
	for i := 0; i < 100; i++ {
		if got := chaosHash(777); got != first {
			t.Fatalf("chaos hash divergent: %d != %d", got, first)
		}
	}
}

func chaosHash(seed uint64) uint64 {
	s := New(seed)
	var h uint64
	for i := 0; i < 1000; i++ {
		h ^= s.Uint64()
		h = h<<21 | h>>43
	}
	return h
}

// ensure math import is exercised for clarity of statistical bounds.
var _ = math.Abs
