/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

// Duration is a signed span of virtual time measured in nanoseconds.
//
// It mirrors the integer-nanosecond contract of Instant so that all simulator
// time arithmetic stays exact and deterministic.
type Duration int64

const (
	// Nanosecond is one nanosecond.
	Nanosecond Duration = 1

	// Microsecond is one thousand nanoseconds.
	Microsecond = 1000 * Nanosecond

	// Millisecond is one thousand microseconds.
	Millisecond = 1000 * Microsecond

	// Second is one thousand milliseconds.
	Second = 1000 * Millisecond

	// Minute is sixty seconds.
	Minute = 60 * Second

	// Hour is sixty minutes.
	Hour = 60 * Minute
)

// Abs returns the absolute value of d.
func (d Duration) Abs() Duration {
	if d < 0 {
		return -d
	}
	return d
}

// Nanoseconds returns the duration as a raw int64 nanosecond count.
func (d Duration) Nanoseconds() int64 { return int64(d) }
