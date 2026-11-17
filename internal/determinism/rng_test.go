/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import "testing"

// TestRNGMixKnownVectors pins the SplitMix64 mixing rounds to their published
// first four outputs for a state seeded at zero.
func TestRNGMixKnownVectors(t *testing.T) {
	t.Parallel()
	want := [4]uint64{
		0xe220a8397b1dcdaf,
		0x6e789e6aa1b965f4,
		0x06c45d188009454f,
		0xf88bb8a8724c81ec,
	}
	state := uint64(0)
	for i, expect := range want {
		state += splitMixGamma
		if got := mix(state); got != expect {
			t.Fatalf("mix vector %d: got 0x%016x want 0x%016x", i, got, expect)
		}
	}
}

// TestRNGSameSeedDoubleRun asserts that one seed, manifest digest and label
// reproduce a byte-identical stream across one hundred regenerations.
func TestRNGSameSeedDoubleRun(t *testing.T) {
	t.Parallel()
	build := func() []byte {
		r := NewRNG(4242, "manifest-digest", "mutation")
		out := make([]byte, 64)
		for i := range out {
			out[i] = byte(r.Next() >> 56)
		}
		return out
	}
	MustMatch100(t, 4242, build)
	MustMatch100(t, 4242, build)
}

// TestRNGLabelSeparatesStreams asserts that a different label derives a
// different stream even when the seed and manifest digest are unchanged.
func TestRNGLabelSeparatesStreams(t *testing.T) {
	t.Parallel()
	left := NewRNG(7, "digest", "mutation")
	right := NewRNG(7, "digest", "chaos")
	for i := 0; i < 16; i++ {
		if left.Next() == right.Next() {
			t.Fatalf("streams collided at draw %d", i)
		}
	}
}

// TestRNGIntnBounds asserts Intn stays inside its range and handles a
// non-positive bound without dividing by zero.
func TestRNGIntnBounds(t *testing.T) {
	t.Parallel()
	r := NewRNG(99, "digest", "points")
	for i := 0; i < 256; i++ {
		got := r.Intn(13)
		if got < 0 || got >= 13 {
			t.Fatalf("Intn(13) returned %d", got)
		}
	}
	if r.Intn(0) != 0 {
		t.Fatalf("Intn(0) must return zero")
	}
}

// TestStrideEndpoints asserts Stride is ascending, includes both endpoints
// and returns exactly the requested number of indices.
func TestStrideEndpoints(t *testing.T) {
	t.Parallel()
	got := Stride(1000, 7)
	if len(got) != 7 {
		t.Fatalf("Stride(1000, 7) returned %d indices", len(got))
	}
	if got[0] != 0 || got[len(got)-1] != 999 {
		t.Fatalf("Stride endpoints are %d and %d", got[0], got[len(got)-1])
	}
	for i := 1; i < len(got); i++ {
		if got[i] <= got[i-1] {
			t.Fatalf("Stride is not strictly ascending at %d", i)
		}
	}
}

// TestStrideDegenerateCases asserts empty and saturated inputs are handled
// without panics or duplicate indices.
func TestStrideDegenerateCases(t *testing.T) {
	t.Parallel()
	if len(Stride(0, 5)) != 0 || len(Stride(5, 0)) != 0 {
		t.Fatalf("non-positive inputs must yield an empty selection")
	}
	if got := Stride(3, 10); len(got) != 3 || got[2] != 2 {
		t.Fatalf("saturated Stride returned %v", got)
	}
	if got := Stride(10, 1); len(got) != 1 || got[0] != 0 {
		t.Fatalf("single Stride returned %v", got)
	}
}

// TestStrideDeterministic asserts identical inputs yield identical indices.
func TestStrideDeterministic(t *testing.T) {
	t.Parallel()
	left := Stride(6262, 200)
	right := Stride(6262, 200)
	if len(left) != len(right) {
		t.Fatalf("Stride lengths differ")
	}
	for i := range left {
		if left[i] != right[i] {
			t.Fatalf("Stride differs at %d", i)
		}
	}
}
