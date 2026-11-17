/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import (
	"errors"
	"testing"
)

// TestVirtualClockEpochIsStable verifies Epoch returns the fixed origin regardless of advances.
func TestVirtualClockEpochIsStable(t *testing.T) {
	c := NewVirtualClock(7_000)
	if got := c.Epoch(); got != 7_000 {
		t.Fatalf("Epoch() = %d, want 7000", got)
	}
	c.Advance(500)
	if got := c.Epoch(); got != 7_000 {
		t.Fatalf("Epoch() after Advance = %d, want 7000", got)
	}
	c.Reset()
	if got := c.Epoch(); got != 7_000 {
		t.Fatalf("Epoch() after Reset = %d, want 7000", got)
	}
}

// TestVirtualClockAdvanceMovesForward verifies monotonic forward-only steps.
func TestVirtualClockAdvanceMovesForward(t *testing.T) {
	cases := []struct {
		name  string
		steps []uint64
		want  uint64
	}{
		{name: "zero step", steps: []uint64{0}, want: 100},
		{name: "single step", steps: []uint64{25}, want: 125},
		{name: "multiple steps", steps: []uint64{1, 2, 3, 4}, want: 110},
		{name: "large step", steps: []uint64{1 << 40}, want: 100 + (1 << 40)},
	}
	for _, tc := range cases {
		c := NewVirtualClock(100)
		for _, step := range tc.steps {
			c.Advance(step)
		}
		if got := c.Now(); got != tc.want {
			t.Errorf("%s: Now() = %d, want %d", tc.name, got, tc.want)
		}
	}
}

// TestVirtualClockResetRestoresEpoch verifies Reset discards accumulated advances.
func TestVirtualClockResetRestoresEpoch(t *testing.T) {
	c := NewVirtualClock(500)
	c.Advance(13)
	c.Advance(29)
	if got := c.Now(); got != 542 {
		t.Fatalf("Now() before Reset = %d, want 542", got)
	}
	c.Reset()
	if got := c.Now(); got != 500 {
		t.Fatalf("Now() after Reset = %d, want 500", got)
	}
	c.Advance(1)
	if got := c.Now(); got != 501 {
		t.Fatalf("Now() after Reset+Advance = %d, want 501", got)
	}
}

// TestVirtualClockDeterminism verifies identical traces yield identical timestamps.
func TestVirtualClockDeterminism(t *testing.T) {
	trace := []uint64{0, 7, 11, 1 << 32, 3}
	left := NewVirtualClock(99)
	right := NewVirtualClock(99)
	for _, step := range trace {
		left.Advance(step)
		right.Advance(step)
	}
	if left.Now() != right.Now() || left.Epoch() != right.Epoch() {
		t.Fatalf("clocks diverged: %d/%d vs %d/%d", left.Epoch(), left.Now(), right.Epoch(), right.Now())
	}
}

// TestDefaultConfigValues verifies the documented deterministic defaults.
func TestDefaultConfigValues(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.FuelLimit != 1<<20 {
		t.Errorf("FuelLimit = %d, want %d", cfg.FuelLimit, 1<<20)
	}
	if cfg.MemoryMinPages != 1 || cfg.MemoryMaxPages != 256 {
		t.Errorf("memory pages = %d/%d, want 1/256", cfg.MemoryMinPages, cfg.MemoryMaxPages)
	}
	if !cfg.ClockEnabled || cfg.ClockEpoch != 0 {
		t.Errorf("clock = %v/%d, want true/0", cfg.ClockEnabled, cfg.ClockEpoch)
	}
}

// TestConfigValidateAcceptsConsistentLimits verifies valid configurations pass.
func TestConfigValidateAcceptsConsistentLimits(t *testing.T) {
	cases := []struct {
		name string
		cfg  Config
	}{
		{name: "defaults", cfg: DefaultConfig()},
		{name: "equal min and max", cfg: Config{MemoryMinPages: 4, MemoryMaxPages: 4}},
		{name: "min below max", cfg: Config{MemoryMinPages: 1, MemoryMaxPages: 9}},
	}
	for _, tc := range cases {
		if err := tc.cfg.Validate(); err != nil {
			t.Errorf("%s: Validate() = %v, want nil", tc.name, err)
		}
	}
}

// TestConfigValidateRejectsInconsistentLimits verifies both failure branches.
func TestConfigValidateRejectsInconsistentLimits(t *testing.T) {
	cases := []struct {
		name string
		cfg  Config
		want error
	}{
		{name: "zero max", cfg: Config{MemoryMaxPages: 0}, want: errConfigMemoryMax},
		{name: "min exceeds max", cfg: Config{MemoryMaxPages: 2, MemoryMinPages: 3}, want: errConfigMemoryMin},
	}
	for _, tc := range cases {
		if err := tc.cfg.Validate(); !errors.Is(err, tc.want) {
			t.Errorf("%s: Validate() = %v, want %v", tc.name, err, tc.want)
		}
	}
}
