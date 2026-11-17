/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import (
	"math"
	"testing"

	"abate/sim/clock"
)

// TestSliceTransmitNSBoundaries locks the integer-only transmit-time model.
func TestSliceTransmitNSBoundaries(t *testing.T) {
	t.Parallel()
	if got := sliceTransmitNS(0, 1000); got != 0 {
		t.Fatalf("empty payload duration = %d", got)
	}
	if got := sliceTransmitNS(10, 0); got != 0 {
		t.Fatalf("disabled bandwidth duration = %d", got)
	}
	if got := sliceTransmitNS(1, 1_000_000_000); got != 1 {
		t.Fatalf("one byte at 1 GB/s = %d ns", got)
	}
	if got := sliceTransmitNS(1000, 1000); got != 1_000_000_000 {
		t.Fatalf("1000 bytes at 1000 B/s = %d ns", got)
	}
	// A payload spanning more than one slice must equal the sum of the slices.
	want := int64(1<<16)*1_000_000_000/1000 + 1_000_000_000/1000
	if got := sliceTransmitNS(1<<16+1, 1000); got != want {
		t.Fatalf("sliced duration = %d, want %d", got, want)
	}
	// An unsigned capacity beyond the signed range must never yield a negative delay.
	if got := sliceTransmitNS(1, ^uint64(0)); got < 0 {
		t.Fatalf("oversized capacity produced a negative duration: %d", got)
	}
	if got := sliceTransmitNS(10_000_000_000, 1); got != math.MaxInt64 {
		t.Fatalf("overflow was not clamped: %d", got)
	}
}

// TestTokenBucketReservations covers FIFO queueing, idle refill, tie-break and reset.
func TestTokenBucketReservations(t *testing.T) {
	t.Parallel()
	tb := &tokenBucket{}
	firstStart, firstFinish := tb.reserve(100, 10, 7, 10)
	if firstStart != 100 || firstFinish != 100+1_000_000_000 {
		t.Fatalf("first reservation = (%d,%d)", firstStart, firstFinish)
	}
	// A second packet at the same instant queues behind the first in FIFO order.
	secondStart, secondFinish := tb.reserve(100, 10, 8, 10)
	if secondStart != firstFinish || secondFinish <= secondStart {
		t.Fatalf("queued reservation = (%d,%d)", secondStart, secondFinish)
	}
	// Advancing virtual time past the debt refills the byte tokens.
	idleStart, _ := tb.reserve(secondFinish+50, 10, 9, 10)
	if idleStart != secondFinish+50 {
		t.Fatalf("idle reservation start = %d", idleStart)
	}
	// A smaller packet id arriving at an identical slot cannot jump the queue.
	_, _ = tb.reserve(idleStart, 0, 20, 10)
	tieStart, _ := tb.reserve(idleStart, 0, 3, 10)
	if tieStart < idleStart {
		t.Fatalf("tie-break moved a packet backwards: %d", tieStart)
	}
	tb.reset()
	resetStart, _ := tb.reserve(500, 10, 1, 10)
	if resetStart != 500 {
		t.Fatalf("reset did not clear accumulated credit: %d", resetStart)
	}
}

// TestSendBandwidthQueueing proves cross-packet queueing is observable and
// respects the integer token-bucket transmission time.
func TestSendBandwidthQueueing(t *testing.T) {
	clk := clock.NewVirtualClock()
	n := NewNetworkWithSeed(clk, 1)
	n.Link("a", "b", LinkConfig{BandwidthBps: 1000})
	deliveries := make([]int64, 0, 4)
	for i := 0; i < 4; i++ {
		if _, err := n.Send("a", "b", make([]byte, 100)); err != nil {
			t.Fatalf("send %d: %v", i, err)
		}
		deliveries = append(deliveries, n.events[len(n.events)-1].At)
	}
	for i := 1; i < len(deliveries); i++ {
		if deliveries[i]-deliveries[i-1] != 100_000_000 {
			t.Fatalf("queue spacing = %v", deliveries)
		}
	}
	// Recovering the link clears the queued debt as dual compensation.
	lk := n.linkBetween("a", "b")
	lk.Recover()
	if lk.bucket.active {
		t.Fatal("link recovery did not reset the bandwidth bucket")
	}
}
