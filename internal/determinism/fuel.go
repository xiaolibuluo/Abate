/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import "math"

// FuelMeter is a deterministic, cooperative instruction-fuel meter.
//
// The meter models the SandboxFuel specification: fuel is a non-negative
// counter that starts at an initial budget, can never go negative, can never be
// consumed beyond the initial budget, and once exhausted leaves the meter
// permanently halted. Consumption is cooperative: it is charged by the host at
// its call boundary. The meter is exact at the host-call granularity and
// over-approximates the instruction count of pure guest loops.
type FuelMeter struct {
	initial   uint64
	remaining uint64
	used      uint64
	halted    bool
}

// NewFuelMeter constructs a meter with the given initial budget.
//
// A budget of zero disables metering: Consume never halts and reports a maximal
// remaining budget, allowing legacy guests to run unbounded.
func NewFuelMeter(initial uint64) *FuelMeter {
	return &FuelMeter{initial: initial, remaining: initial}
}

// Initial returns the starting fuel budget.
func (m *FuelMeter) Initial() uint64 { return m.initial }

// Remaining returns the fuel left before exhaustion.
func (m *FuelMeter) Remaining() uint64 { return m.remaining }

// Used returns the fuel consumed so far.
func (m *FuelMeter) Used() uint64 { return m.used }

// Halted reports whether the meter has been exhausted.
func (m *FuelMeter) Halted() bool { return m.halted }

// Consume charges n units of fuel.
//
// It returns the remaining budget and whether the meter is now halted. When n
// exceeds the remaining budget the meter clamps to zero, marks itself halted
// and reports remaining zero, preserving the never-negative invariant.
func (m *FuelMeter) Consume(n uint64) (uint64, bool) {
	if m.initial == 0 {
		return math.MaxUint64, false
	}
	if m.halted {
		return 0, true
	}
	if n == 0 {
		return m.remaining, false
	}
	if n >= m.remaining {
		m.remaining = 0
		m.used = m.initial
		m.halted = true
		return 0, true
	}
	m.remaining -= n
	m.used += n
	return m.remaining, false
}
