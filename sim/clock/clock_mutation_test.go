/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

import "testing"

// TestClockOffsetMutationCompensation asserts the dual-compensation law for a
// chaos clock-offset mutation: applying offset d and then its negation restores
// the baseline reading exactly.
func TestClockOffsetMutationCompensation(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	if _, err := c.Advance(10 * Second); err != nil {
		t.Fatalf("advance err = %v", err)
	}
	baseline := c.Now()

	const shift Duration = 1234567
	c.Offset(shift)
	if c.Now() != Instant(int64(baseline)+int64(shift)) {
		t.Fatalf("offset mutation reading = %d, want %d", c.Now(), Instant(int64(baseline)+int64(shift)))
	}
	// Rollback: the dual of +shift is -shift.
	c.Offset(-shift)
	if c.Now() != baseline {
		t.Fatalf("compensated reading = %d, want baseline %d", c.Now(), baseline)
	}
}

// TestClockRateMutationCompensation verifies a rate mutation followed by a
// reset to nominal rate leaves the timeline deterministic and monotonic.
func TestClockRateMutationCompensation(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	baseline := c.Now()
	if err := c.Rate(4, 1); err != nil {
		t.Fatalf("rate err = %v", err)
	}
	got, err := c.Advance(Second)
	if err != nil {
		t.Fatalf("advance err = %v", err)
	}
	if got != Instant(int64(baseline)+4*int64(Second)) {
		t.Fatalf("fast advance = %d, want %d", got, Instant(int64(baseline)+4*int64(Second)))
	}
	// Rollback to nominal rate.
	if err := c.Rate(1, 1); err != nil {
		t.Fatalf("rate reset err = %v", err)
	}
	got2, err := c.Advance(Second)
	if err != nil || got2 != Instant(int64(got)+int64(Second)) {
		t.Fatalf("post-rollback advance = (%d, %v), want (%d, nil)", got2, err, Instant(int64(got)+int64(Second)))
	}
}

// TestClockMutationDeterminism asserts that the same mutation sequence applied
// from the same baseline yields byte-identical readings across 100 runs.
func TestClockMutationDeterminism(t *testing.T) {
	t.Parallel()
	const rounds = 100
	first := mutateClock(7)
	for i := 0; i < rounds; i++ {
		if got := mutateClock(7); got != first {
			t.Fatalf("round %d reading = %d, want %d", i, got, first)
		}
	}
}

func mutateClock(seed uint64) int64 {
	c := NewVirtualClock()
	offsets := []Duration{10, -5, 20, 0, -20, 5}
	idx := seed % uint64(len(offsets))
	for i := 0; i < 6; i++ {
		c.Offset(offsets[(int(idx)+i)%len(offsets)]) //nolint:gosec // value is bounded by a modulo or bitmask; conversion is safe.
		_, _ = c.Advance(Millisecond)
	}
	return c.Now().UnixNano()
}
