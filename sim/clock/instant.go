/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

// Instant is a monotonically increasing virtual time measured in nanoseconds.
//
// Instant is the only time source inside the deterministic simulator. It is an
// integer so that event ordering and hashing are reproducible across runs; the
// simulator forbids any floating-point time and any real wall-clock access.
type Instant int64

const (
	// MaxInstant is the largest representable virtual instant (≈292 years).
	MaxInstant = Instant(1<<63 - 1)

	// MinInstant is the smallest representable virtual instant.
	MinInstant = Instant(-1 << 63)
)

// Add returns the instant t advanced by the given duration.
func (t Instant) Add(d Duration) Instant { return t + Instant(d) }

// Sub returns the signed duration from u to t.
func (t Instant) Sub(u Instant) Duration { return Duration(t - u) }

// Before reports whether t is strictly before u.
func (t Instant) Before(u Instant) bool { return t < u }

// After reports whether t is strictly after u.
func (t Instant) After(u Instant) bool { return t > u }

// IsZero reports whether t equals the zero instant.
func (t Instant) IsZero() bool { return t == 0 }

// UnixNano returns the virtual time as a raw int64 nanosecond count.
func (t Instant) UnixNano() int64 { return int64(t) }
