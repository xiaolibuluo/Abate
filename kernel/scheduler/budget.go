/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

// Default execution budget limits. A zero field means "unbounded".
const (
	DefaultMaxEvents          uint64      = 1_000_000
	DefaultMaxVirtualDuration VirtualTime = 24 * Hour
	DefaultMaxCost            uint64      = 1_000_000
)

// Budget bounds an experiment along three independent axes.
//
// Any limit being crossed terminates the run with BUDGET_EXCEEDED, preventing
// a runaway experiment from consuming unbounded virtual time.
type Budget struct {
	MaxEvents          uint64
	MaxVirtualDuration VirtualTime
	MaxCost            uint64
}

// DefaultBudget returns the standard execution budget.
func DefaultBudget() Budget {
	return Budget{
		MaxEvents:          DefaultMaxEvents,
		MaxVirtualDuration: DefaultMaxVirtualDuration,
		MaxCost:            DefaultMaxCost,
	}
}

// Allows checks the three limits against the current run counters.
//
// A zero limit disables that axis, which is useful for focused tests.
func (b Budget) Allows(events uint64, span VirtualTime, cost uint64) error {
	if b.MaxEvents > 0 && events > b.MaxEvents {
		return ErrBudgetEvents
	}
	if b.MaxVirtualDuration > 0 && span > b.MaxVirtualDuration {
		return ErrBudgetDuration
	}
	if b.MaxCost > 0 && cost > b.MaxCost {
		return ErrBudgetCost
	}
	return nil
}
