/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

import (
	"testing"
)

// TestSyncAlignsUnderlyingTimeline verifies Sync sets the underlying instant
// while Now keeps adding the configured offset.
func TestSyncAlignsUnderlyingTimeline(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	c.Offset(5)
	c.Sync(100)
	if got := c.Now(); got != 105 {
		t.Fatalf("Now = %d, want 105", got)
	}
	if c.State() != StateOffset {
		t.Fatalf("state = %v, want OFFSET", c.State())
	}
}

// TestSyncDoesNotDisturbOffsetOrdering checks that Sync keeps offset semantics
// intact across a subsequent advance.
func TestSyncDoesNotDisturbOffsetOrdering(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	c.Offset(3)
	c.Sync(0)
	before := c.Now()
	if _, err := c.Advance(Second); err != nil {
		t.Fatalf("Advance err = %v", err)
	}
	after := c.Now()
	if after.Sub(before) != Second {
		t.Fatalf("offset disturbed advance: delta=%d want %d", after.Sub(before), Second)
	}
}

// TestSyncDeterminism asserts repeated Sync/Advance sequences are reproducible.
func TestSyncDeterminism(t *testing.T) {
	t.Parallel()
	run := func() int64 {
		c := NewVirtualClock()
		c.Sync(10)
		_, _ = c.Advance(Second)
		c.Sync(50)
		_, _ = c.Advance(Millisecond)
		return c.Now().UnixNano()
	}
	first := run()
	for i := 0; i < 100; i++ {
		if run() != first {
			t.Fatalf("Sync/Advance sequence non-deterministic at run %d", i)
		}
	}
}
