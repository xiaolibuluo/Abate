/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import "math"

// tokenBucket is a deterministic integer byte-token bucket bound to one link.
//
// It is a standard-library rewrite of the bandwidth toxic model from
// Shopify/toxiproxy. The improvement is a fully deterministic virtual-time
// formulation: wall-clock sleeps and floating-point rates are replaced by
// integer nanosecond credit, large payloads are sliced into bounded chunks
// before the transmit duration is summed, and reservations at the same virtual
// instant are totally ordered by packet id so no traversal order can leak.
type tokenBucket struct {
	readyAt int64
	lastAt  int64
	lastSeq PacketID
	active  bool
}

// sliceTransmitNS returns the integer virtual nanoseconds needed for n bytes.
//
// Chunking mirrors the slicer toxic from Shopify/toxiproxy but is deterministic:
// every chunk has a fixed byte size, the per-chunk duration is an integer
// division, and the total is clamped instead of overflowing or going negative
// when the configured capacity exceeds the signed integer range.
func sliceTransmitNS(n int, bps uint64) int64 {
	if n <= 0 || bps == 0 {
		return 0
	}
	const nsPerSecond uint64 = 1_000_000_000
	const chunkBytes uint64 = 1 << 16
	remaining := uint64(n) //nolint:gosec // n is a non-negative payload length.
	var total uint64
	for remaining > chunkBytes {
		total += chunkBytes * nsPerSecond / bps
		remaining -= chunkBytes
	}
	total += remaining * nsPerSecond / bps
	if total > math.MaxInt64 {
		return math.MaxInt64
	}
	return int64(total) //nolint:gosec // total was clamped to MaxInt64 above.
}

// reserve allocates link time for one packet and returns its start and finish.
//
// The bucket refills byte tokens as virtual time advances: a request arriving
// before readyAt queues behind the previous packet in FIFO order. When two
// requests share a virtual instant, the smaller packet id owns the earlier
// slot, which is the deterministic tie-break required by the SIM-NET design.
func (tb *tokenBucket) reserve(now int64, n int, pid PacketID, bps uint64) (start int64, finish int64) {
	duration := sliceTransmitNS(n, bps)
	start = now
	if tb.active && tb.readyAt > start {
		start = tb.readyAt
	}
	if tb.active && start == tb.lastAt && pid < tb.lastSeq {
		start = tb.readyAt
	}
	finish = start + duration
	tb.readyAt = finish
	tb.lastAt = start
	tb.lastSeq = pid
	tb.active = true
	return start, finish
}

// reset clears accumulated credit as the dual compensation of a link recovery.
func (tb *tokenBucket) reset() {
	tb.readyAt = 0
	tb.lastAt = 0
	tb.lastSeq = 0
	tb.active = false
}
