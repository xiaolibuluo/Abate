/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

import (
	"errors"
	"math"
	"testing"
)

func TestInstantArithmetic(t *testing.T) {
	t.Parallel()
	var a Instant = 1000
	if got := a.Add(500); got != 1500 {
		t.Fatalf("Add = %d, want 1500", got)
	}
	if got := a.Sub(300); got != 700 {
		t.Fatalf("Sub = %d, want 700", got)
	}
	if !a.Before(2000) {
		t.Fatalf("Before(2000) = false, want true")
	}
	if !a.After(500) {
		t.Fatalf("After(500) = false, want true")
	}
	if a.IsZero() {
		t.Fatalf("IsZero = true, want false")
	}
	if !Instant(0).IsZero() {
		t.Fatalf("Instant(0).IsZero = false, want true")
	}
	if a.UnixNano() != 1000 {
		t.Fatalf("UnixNano = %d, want 1000", a.UnixNano())
	}
	if MaxInstant <= MinInstant {
		t.Fatalf("MaxInstant must exceed MinInstant")
	}
}

func TestDurationAbs(t *testing.T) {
	t.Parallel()
	d := Duration(-42)
	if d.Abs() != 42 {
		t.Fatalf("Abs = %d, want 42", d.Abs())
	}
	if Duration(7).Abs() != 7 {
		t.Fatalf("Abs = %d, want 7", Duration(7).Abs())
	}
	if Second != 1_000_000_000 {
		t.Fatalf("Second = %d, want 1e9", Second)
	}
	if Hour != 60*Minute {
		t.Fatalf("Hour != 60*Minute")
	}
}

// TestDurationNanoseconds verifies that Nanoseconds returns the raw int64
// nanosecond count for zero, positive, negative, and int64-boundary durations.
func TestDurationNanoseconds(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		d    Duration
		want int64
	}{
		{name: "zero", d: 0, want: 0},
		{name: "single nanosecond", d: Nanosecond, want: 1},
		{name: "one second", d: Second, want: 1_000_000_000},
		{name: "one hour", d: Hour, want: 3_600_000_000_000},
		{name: "negative microsecond", d: -Microsecond, want: -1_000},
		{name: "negative span", d: -42 * Millisecond, want: -42_000_000},
		{name: "max int64", d: Duration(math.MaxInt64), want: math.MaxInt64},
		{name: "min int64", d: Duration(math.MinInt64), want: math.MinInt64},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.d.Nanoseconds(); got != tc.want {
				t.Fatalf("Duration(%d).Nanoseconds() = %d, want %d", tc.d, got, tc.want)
			}
		})
	}
}

func TestVirtualClockAdvance(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	if c.Now() != 0 {
		t.Fatalf("Now = %d, want 0", c.Now())
	}
	got, err := c.Advance(Second)
	if err != nil || got != Instant(Second) {
		t.Fatalf("Advance = (%d, %v), want (1e9, nil)", got, err)
	}
	if c.Now() != Instant(Second) {
		t.Fatalf("Now = %d, want 1e9", c.Now())
	}
}

func TestVirtualClockNegativeAndZero(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	if _, err := c.Advance(0); !errors.Is(err, ErrNegativeDuration) {
		t.Fatalf("Advance(0) err = %v, want ErrNegativeDuration", err)
	}
	if _, err := c.Advance(-1); !errors.Is(err, ErrNegativeDuration) {
		t.Fatalf("Advance(-1) err = %v, want ErrNegativeDuration", err)
	}
}

func TestVirtualClockFreeze(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	if err := c.Freeze(); err != nil {
		t.Fatalf("Freeze err = %v", err)
	}
	if _, err := c.Advance(Second); !errors.Is(err, ErrFrozen) {
		t.Fatalf("Advance while frozen err = %v, want ErrFrozen", err)
	}
	if c.State() != StateFreeze {
		t.Fatalf("State = %s, want FREEZE", c.State())
	}
	if err := c.Unfreeze(); err != nil {
		t.Fatalf("Unfreeze err = %v", err)
	}
	if _, err := c.Advance(Second); err != nil {
		t.Fatalf("Advance after unfreeze err = %v", err)
	}
}

func TestVirtualClockOffset(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	if _, err := c.Advance(Second); err != nil {
		t.Fatalf("Advance err = %v", err)
	}
	c.Offset(500)
	if c.Now() != Instant(int64(Second)+500) {
		t.Fatalf("Now with offset = %d, want %d", c.Now(), Instant(int64(Second)+500))
	}
	// Offset does not disturb the underlying timeline used for ordering.
	if c.now != Instant(Second) {
		t.Fatalf("underlying now = %d, want 1e9", c.now)
	}
	if c.State() != StateOffset {
		t.Fatalf("State = %s, want OFFSET", c.State())
	}
}

func TestVirtualClockRate(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	if err := c.Rate(2, 1); err != nil {
		t.Fatalf("Rate err = %v", err)
	}
	got, err := c.Advance(Second)
	if err != nil || got != Instant(2*int64(Second)) {
		t.Fatalf("Advance at 2x = (%d, %v), want (2e9, nil)", got, err)
	}
	if c.State() != StateFast {
		t.Fatalf("State = %s, want FAST", c.State())
	}
	if err := c.Rate(1, 0); !errors.Is(err, ErrZeroRate) {
		t.Fatalf("Rate(1,0) err = %v, want ErrZeroRate", err)
	}
}

func TestVirtualClockOverflow(t *testing.T) {
	t.Parallel()
	c := NewVirtualClock()
	c.now = MaxInstant - 10
	if _, err := c.Advance(Second); !errors.Is(err, ErrClockOverflow) {
		t.Fatalf("Advance overflow err = %v, want ErrClockOverflow", err)
	}
}

func TestStateRunning(t *testing.T) {
	t.Parallel()
	if NewVirtualClock().State() != StateRunning {
		t.Fatalf("default State = %s, want RUNNING", NewVirtualClock().State())
	}
	if StateRunning.String() != "RUNNING" || StateUnknownCheck() != "UNKNOWN" {
		t.Fatalf("state string rendering incorrect")
	}
}

// StateUnknownCheck exercises the default branch of State.String.
func StateUnknownCheck() string {
	return State(99).String()
}
