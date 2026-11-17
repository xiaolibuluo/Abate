/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

// VirtualClock is a deterministic virtual time source for guest plugins.
//
// It replaces wall-clock access: guests read time only through host functions,
// never through time.Now. The clock advances only when Advance is called, so
// identical execution traces produce identical timestamps and replays are exact.
type VirtualClock struct {
	epoch uint64
	now   uint64
}

// NewVirtualClock creates a clock anchored at the given epoch in nanoseconds.
func NewVirtualClock(epoch uint64) *VirtualClock {
	return &VirtualClock{epoch: epoch, now: epoch}
}

// Epoch returns the fixed virtual time origin in nanoseconds.
func (c *VirtualClock) Epoch() uint64 { return c.epoch }

// Now returns the current virtual time in nanoseconds.
func (c *VirtualClock) Now() uint64 { return c.now }

// Advance moves the virtual clock forward by ns nanoseconds.
//
// The delta is always applied as a forward step, preserving the monotonic
// determinism contract: time only moves forward and never reads the wall clock.
func (c *VirtualClock) Advance(ns uint64) {
	c.now += ns
}

// Reset restores the clock to its epoch.
func (c *VirtualClock) Reset() {
	c.now = c.epoch
}
