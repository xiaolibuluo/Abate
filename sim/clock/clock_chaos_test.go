/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

import (
	"hash/fnv"
	"testing"
)

// TestClockChaosOperationStorm fires pseudo-random operation sequences drawn
// from a deterministic stream at the clock and asserts that the resulting
// reading sequence is reproducible and that no illegal transition panics.
func TestClockChaosOperationStorm(t *testing.T) {
	t.Parallel()
	const runs = 200
	for seed := uint64(0); seed < runs; seed++ {
		got := runChaosSequence(seed)
		want := runChaosSequence(seed)
		if got != want {
			t.Fatalf("seed %d produced divergent readings: %d != %d", seed, got, want)
		}
	}
}

// runChaosSequence applies 64 random operations and returns the final reading.
func runChaosSequence(seed uint64) int64 {
	h := fnv.New64a()
	c := NewVirtualClock()
	state := seed
	for i := 0; i < 64; i++ {
		state = state*2862933555777941757 + 1
		op := state % 6
		switch op {
		case 0:
			_, _ = c.Advance(Second)
		case 1:
			_ = c.Freeze()
		case 2:
			_ = c.Unfreeze()
		case 3:
			c.Offset(Duration(int64(state % 1000))) //nolint:gosec // value is bounded by a modulo or bitmask; conversion is safe.
		case 4:
			_ = c.Rate(int64(1+state%4), 1) //nolint:gosec // value is bounded by a modulo or bitmask; conversion is safe.
		case 5:
			_ = c.Rate(1, 1)
		}
		_, _ = h.Write([]byte(c.State().String()))
	}
	_, _ = h.Write([]byte{})
	return c.Now().UnixNano()
}

// TestClockChaosOffsetStorm verifies that injecting many offset perturbations
// never disturbs the underlying event-ordering timeline (the Now reading may
// shift, but c.now stays monotonic).
func TestClockChaosOffsetStorm(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	lastUnderlying := c.now
	for i := 0; i < 1000; i++ {
		c.Offset(Duration((i % 2) * -10))
		if _, err := c.Advance(Millisecond); err != nil {
			t.Fatalf("advance %d err = %v", i, err)
		}
		if c.now < lastUnderlying {
			t.Fatalf("underlying timeline regressed at step %d", i)
		}
		lastUnderlying = c.now
	}
}
