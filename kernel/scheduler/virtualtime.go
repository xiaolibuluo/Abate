/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

// VirtualTime is an integer nanosecond instant on the simulation timeline.
//
// No floating point value ever participates in event ordering, so the same
// input always produces the same total order.
type VirtualTime int64

// Virtual-time unit constants, all expressed in integer nanoseconds.
const (
	Nanosecond  VirtualTime = 1
	Microsecond             = 1000 * Nanosecond
	Millisecond             = 1000 * Microsecond
	Second                  = 1000 * Millisecond
	Minute                  = 60 * Second
	Hour                    = 60 * Minute
)

// UnixNano returns the raw nanosecond count.
func (t VirtualTime) UnixNano() int64 {
	return int64(t)
}

// Before reports whether t is strictly earlier than other.
func (t VirtualTime) Before(other VirtualTime) bool {
	return t < other
}

// After reports whether t is strictly later than other.
func (t VirtualTime) After(other VirtualTime) bool {
	return t > other
}

// Add returns the instant t shifted by d.
func (t VirtualTime) Add(d VirtualTime) VirtualTime {
	return t + d
}

// Sub returns the signed distance from other to t.
func (t VirtualTime) Sub(other VirtualTime) VirtualTime {
	return t - other
}
