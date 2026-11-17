/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

// Clock is the deterministic time source contract used across the simulator.
//
// Now returns the current virtual instant. Advance advances the underlying
// timeline. Freeze and Unfreeze toggle a frozen mode in which advances are
// rejected. Offset shifts the reading without disturbing event ordering. Rate
// changes the virtual-time flow multiplier.
type Clock interface {
	// Now returns the current virtual instant including any offset shift.
	Now() Instant

	// Advance moves the underlying timeline forward by d nanoseconds.
	Advance(d Duration) (Instant, error)

	// Freeze rejects further advances until Unfreeze is called.
	Freeze() error

	// Unfreeze re-enables advances.
	Unfreeze() error

	// Offset shifts the Now reading by d without changing event ordering.
	Offset(d Duration)

	// Rate sets the virtual-time flow multiplier num/divisor.
	Rate(num int64, divisor int64) error

	// State reports the derived four-state mode of the clock.
	State() State
}
