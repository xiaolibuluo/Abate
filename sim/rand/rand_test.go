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

func TestUint64RangeAndDeterminism(t *testing.T) {
	t.Parallel()
	a := New(42)
	b := New(42)
	for i := 0; i < 1000; i++ {
		if a.Uint64() != b.Uint64() {
			t.Fatalf("seed 42 divergent at draw %d", i)
		}
	}
}

func TestInt63NonNegative(t *testing.T) {
	t.Parallel()
	s := New(7)
	for i := 0; i < 1000; i++ {
		v := s.Int63()
		if v < 0 {
			t.Fatalf("Int63 returned negative: %d", v)
		}
	}
}

func TestFloat64Bounds(t *testing.T) {
	t.Parallel()
	s := New(11)
	for i := 0; i < 1000; i++ {
		v := s.Float64()
		if v < 0 || v >= 1 {
			t.Fatalf("Float64 out of [0,1): %v", v)
		}
	}
}

func TestIntNBounds(t *testing.T) {
	t.Parallel()
	s := New(3)
	if s.IntN(0) != 0 || s.IntN(1) != 0 {
		t.Fatalf("IntN degenerate cases failed")
	}
	for i := 0; i < 1000; i++ {
		v := s.IntN(10)
		if v < 0 || v >= 10 {
			t.Fatalf("IntN(10) out of range: %d", v)
		}
	}
}

func TestBoolDistribution(t *testing.T) {
	t.Parallel()
	s := New(99)
	var ones int
	const n = 20000
	for i := 0; i < n; i++ {
		if s.Bool() {
			ones++
		}
	}
	ratio := float64(ones) / float64(n)
	if math.Abs(ratio-0.5) > 0.05 {
		t.Fatalf("Bool ratio %.4f too far from 0.5", ratio)
	}
}

func TestBytesLength(t *testing.T) {
	t.Parallel()
	s := New(5)
	b := s.Bytes(37)
	if len(b) != 37 {
		t.Fatalf("Bytes length = %d, want 37", len(b))
	}
	if len(s.Bytes(0)) != 0 {
		t.Fatalf("Bytes(0) should be empty")
	}
}

func TestForkIndependence(t *testing.T) {
	t.Parallel()
	parent := New(123)
	child := parent.Fork()
	// Parent and child must diverge immediately.
	if parent.Uint64() == child.Uint64() {
		t.Fatalf("fork produced identical streams")
	}
	// Two independent parents forked at the same seed yield identical children.
	p1 := New(123)
	p2 := New(123)
	c1 := p1.Fork()
	c2 := p2.Fork()
	for i := 0; i < 100; i++ {
		if c1.Uint64() != c2.Uint64() {
			t.Fatalf("child streams diverge for same fork seed at %d", i)
		}
	}
}

func TestAccounting(t *testing.T) {
	t.Parallel()
	s := New(1)
	for i := 0; i < 25; i++ {
		s.Uint64()
	}
	if s.acct.Len() != 25 {
		t.Fatalf("accounting len = %d, want 25", s.acct.Len())
	}
	snap := s.acct.Snapshot()
	if len(snap) != 25 || snap[0].Index != 0 {
		t.Fatalf("snapshot malformed")
	}
	// Snapshot is a copy: mutating it must not affect the ledger.
	snap[0].Result = 0
	if s.acct.Snapshot()[0].Result == 0 {
		t.Fatalf("snapshot is not a defensive copy")
	}
}
