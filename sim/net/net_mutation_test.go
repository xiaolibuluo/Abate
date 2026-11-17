/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import (
	"errors"
	"testing"
)

// TestMutationBandwidthZero confirms a zero bandwidth imposes no queueing delay.
func TestMutationBandwidthZero(t *testing.T) {
	n, _ := buildLinear()
	mid := n.linkBetween("b", "c")
	mid.SetBandwidth(0)
	mid.SetLatency(0)
	base := n.clk.Now().UnixNano()
	_, err := n.Send("a", "d", make([]byte, 1024))
	if err != nil {
		t.Fatalf("send err = %v", err)
	}
	deliver := n.events[len(n.events)-1]
	if deliver.At-base != 2000 {
		t.Fatalf("zero-bandwidth added delay: %d", deliver.At-base)
	}
}

// TestMutationLossBoundary exercises loss rates at the 0 and 1 boundaries.
func TestMutationLossBoundary(t *testing.T) {
	n, _ := buildLinear()
	mid := n.linkBetween("b", "c")
	mid.SetLossRate(0)
	if _, err := n.Send("a", "d", []byte("x")); err != nil {
		t.Fatalf("loss 0 should deliver: %v", err)
	}
	mid.SetLossRate(1)
	if _, err := n.Send("a", "d", []byte("x")); !errors.Is(err, ErrDropped) {
		t.Fatalf("loss 1 should drop, got %v", err)
	}
}

// TestMutationLatencyBoundary confirms a large latency is reflected in delivery.
func TestMutationLatencyBoundary(t *testing.T) {
	n, _ := buildLinear()
	mid := n.linkBetween("b", "c")
	mid.SetLatency(1 << 30)
	base := n.clk.Now().UnixNano()
	_, _ = n.Send("a", "d", nil)
	deliver := n.events[len(n.events)-1]
	if deliver.At-base < int64(1<<30) {
		t.Fatalf("latency not reflected: delta=%d", deliver.At-base)
	}
}

// TestMutationPartitionSendError confirms Send errors while partitioned and that
// Heal restores delivery.
func TestMutationPartitionSendError(t *testing.T) {
	n, _ := buildLinear()
	mid := n.linkBetween("b", "c")
	_ = mid.Partition()
	if !mid.Partitioned() {
		t.Fatalf("link should report partitioned")
	}
	if _, err := n.Send("a", "d", []byte("x")); !errors.Is(err, ErrPartitioned) {
		t.Fatalf("want ErrPartitioned, got %v", err)
	}
	_ = mid.Heal()
	if mid.Partitioned() {
		t.Fatalf("link should report healed")
	}
	if _, err := n.Send("a", "d", []byte("x")); err != nil {
		t.Fatalf("heal should restore delivery: %v", err)
	}
}

// TestMutationFieldVariation verifies each config field independently changes the
// observable outcome (delivery time, drop, or error), which is the core of the
// field-level mutation check.
func TestMutationFieldVariation(t *testing.T) {
	probe := func(mut func(l *Link)) (int64, error) {
		n, _ := buildLinear()
		mid := n.linkBetween("b", "c")
		mut(mid)
		base := n.clk.Now().UnixNano()
		_, err := n.Send("a", "d", make([]byte, 1))
		if err != nil {
			return 0, err
		}
		return n.events[len(n.events)-1].At - base, nil
	}
	baseDelta, err := probe(func(_ *Link) {})
	if err != nil {
		t.Fatalf("base probe err = %v", err)
	}
	latDelta, err := probe(func(l *Link) { l.SetLatency(5000) })
	if err != nil {
		t.Fatalf("latency probe err = %v", err)
	}
	if latDelta <= baseDelta {
		t.Fatalf("latency mutation had no effect: %d vs %d", latDelta, baseDelta)
	}
	_, dropErr := probe(func(l *Link) { l.SetLossRate(1) })
	if !errors.Is(dropErr, ErrDropped) {
		t.Fatalf("loss mutation did not drop: %v", dropErr)
	}
}
