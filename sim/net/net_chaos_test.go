/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import (
	"testing"

	"abate/internal/determinism"
	"abate/sim/clock"
)

// runNetScenario executes a fixed traffic pattern against a fixed seeded network
// and returns the SHA-256 of the totally-ordered event stream.
func runNetScenario(seed uint64) [32]byte {
	clk := clock.NewVirtualClock()
	n := NewNetworkWithSeed(clk, seed)
	for _, id := range []string{"a", "b", "c", "d", "e"} {
		n.AddNode(id)
	}
	n.Link("a", "b", LinkConfig{BaseLatency: 1000, Jitter: 500, BandwidthBps: 1000})
	n.Link("b", "c", LinkConfig{BaseLatency: 1000, LossRate: 0.1})
	n.Link("c", "d", LinkConfig{BaseLatency: 1000})
	n.Link("a", "d", LinkConfig{BaseLatency: 8000})
	n.Link("d", "e", LinkConfig{BaseLatency: 1000})

	for i := 0; i < 200; i++ {
		if i == 100 {
			// Deterministic fault window: partition b-c, then heal at 150.
			_ = n.linkBetween("b", "c").Partition()
		}
		if i == 150 {
			_ = n.linkBetween("b", "c").Heal()
		}
		_, _ = n.Send("a", "e", []byte("p"))
		_, _ = n.Send("a", "d", []byte("q"))
		_, _ = clk.Advance(clock.Millisecond)
	}
	return determinism.HashState(n.events)
}

// TestChaosEventStreamDeterminism asserts the packet event stream is byte-identical
// across 100 independent runs of the same seed and traffic.
func TestChaosEventStreamDeterminism(t *testing.T) {
	first := runNetScenario(12345)
	seen := make(map[[32]byte]int)
	seen[first]++
	for i := 0; i < 100; i++ {
		h := runNetScenario(12345)
		seen[h]++
		if h != first {
			t.Fatalf("event stream divergent on run %d", i)
		}
	}
	if len(seen) != 1 {
		t.Fatalf("expected exactly 1 unique SHA-256, got %d", len(seen))
	}
	t.Logf("NET unique SHA-256 over 100 runs = %d (hash=%x)", len(seen), first)
}

// TestChaosRouteDeterminism asserts routing is identical across 100 runs.
func TestChaosRouteDeterminism(t *testing.T) {
	var first []string
	seen := make(map[[32]byte]int)
	for i := 0; i < 100; i++ {
		clk := clock.NewVirtualClock()
		n := NewNetworkWithSeed(clk, 7)
		for _, id := range []string{"a", "b", "c", "d", "e"} {
			n.AddNode(id)
		}
		n.Link("a", "b", LinkConfig{BaseLatency: 1})
		n.Link("a", "c", LinkConfig{BaseLatency: 1})
		n.Link("b", "d", LinkConfig{BaseLatency: 1})
		n.Link("c", "d", LinkConfig{BaseLatency: 1})
		n.Link("d", "e", LinkConfig{BaseLatency: 1})
		r, err := n.Route("a", "e")
		if err != nil {
			t.Fatalf("route err = %v", err)
		}
		if i == 0 {
			first = r
		}
		h := determinism.HashState(r)
		seen[h]++
		if determinism.HashState(first) != h {
			t.Fatalf("route divergent on run %d: %v vs %v", i, first, r)
		}
	}
	if len(seen) != 1 {
		t.Fatalf("expected exactly 1 unique route hash, got %d", len(seen))
	}
}
