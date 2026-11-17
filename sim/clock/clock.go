/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package clock provides the deterministic virtual clock that is the single
// time source of the Abate simulator.
//
// All simulator time is an integer nanosecond count. Floating-point time and any
// real wall-clock access (time.Now) are forbidden so that event ordering and
// hashing are reproducible across runs.
package clock

import "errors"

// Sentinel errors of the virtual clock.
var (
	// ErrClockOverflow reports an advance that would exceed MaxInstant.
	ErrClockOverflow = errors.New("clock: advance would overflow virtual time")

	// ErrFrozen reports an advance attempted while the clock is frozen.
	ErrFrozen = errors.New("clock: clock is frozen")

	// ErrNegativeDuration reports a non-positive advance argument.
	ErrNegativeDuration = errors.New("clock: advance duration must be positive")

	// ErrZeroRate reports a Rate divisor of zero.
	ErrZeroRate = errors.New("clock: clock rate divisor must be non-zero")
)

// VirtualClock is the lock-free, integer-only implementation of Clock.
//
// The clock is driven by a single virtual thread (the scheduler), so no locking
// is required. All state fields are integers and arithmetic is overflow-checked.
type VirtualClock struct {
	now     Instant
	frozen  bool
	offset  Duration
	rateNum int64
	rateDen int64
}

// NewVirtualClock returns a clock at instant 0 with a nominal rate of 1.
func NewVirtualClock() *VirtualClock {
	return &VirtualClock{rateNum: 1, rateDen: 1}
}

// Now returns the reading, which adds the offset to the underlying timeline.
func (c *VirtualClock) Now() Instant {
	return Instant(int64(c.now) + int64(c.offset))
}

// Advance moves the timeline forward, applying the configured rate.
func (c *VirtualClock) Advance(d Duration) (Instant, error) {
	if d <= 0 {
		return c.now, ErrNegativeDuration
	}
	if c.frozen {
		return c.now, ErrFrozen
	}
	scaled := d * Duration(c.rateNum) / Duration(c.rateDen)
	if scaled > 0 && c.now > MaxInstant-Instant(scaled) {
		return c.now, ErrClockOverflow
	}
	c.now += Instant(scaled)
	return c.now, nil
}

// Freeze marks the clock frozen.
func (c *VirtualClock) Freeze() error {
	c.frozen = true
	return nil
}

// Unfreeze clears the frozen flag.
func (c *VirtualClock) Unfreeze() error {
	c.frozen = false
	return nil
}

// Offset adds d to the reading offset.
func (c *VirtualClock) Offset(d Duration) {
	c.offset += d
}

// Rate sets the flow multiplier.
func (c *VirtualClock) Rate(num int64, divisor int64) error {
	if divisor == 0 {
		return ErrZeroRate
	}
	c.rateNum = num
	c.rateDen = divisor
	return nil
}

// Sync aligns the underlying timeline to the given instant without disturbing
// the configured offset, so Now continues to report t plus the offset. SIM
// modules call Sync at scheduler dispatch boundaries to keep the single virtual
// clock honest with the scheduler's own virtual time.
func (c *VirtualClock) Sync(t Instant) {
	c.now = t
}

// State derives the four-state mode from the clock fields.
func (c *VirtualClock) State() State {
	if c.frozen {
		return StateFreeze
	}
	if c.offset != 0 {
		return StateOffset
	}
	if c.rateNum != c.rateDen {
		return StateFast
	}
	return StateRunning
}
