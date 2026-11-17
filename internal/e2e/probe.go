/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"context"
	"errors"
	"runtime"
	"time"

	"abate/internal/e2e/metrics"
	"abate/lang/air"
)

// ResourceProbe measures one synchronous parse against a timeout and an
// allocation budget. It never starts a goroutine: the parse runs in the
// caller's goroutine and ctx.Err() is inspected afterwards, which keeps the
// single-goroutine execution order intact and avoids scheduling randomness.
type ResourceProbe struct {
	budget ProbeBudget
}

// NewResourceProbe builds a probe from one immutable budget.
func NewResourceProbe(budget ProbeBudget) ResourceProbe {
	return ResourceProbe{budget: budget}
}

// Measure runs parse, classifies its outcome and returns the real allocation
// increment. The deadline covers only the synchronous parse: the two
// stop-the-world MemStats reads stay outside that window, so measurement
// overhead can never be mistaken for a parse timeout. MemLimitBytes bounds
// the TotalAlloc increment, which the design declares as the platform-neutral
// allocation proxy. That verdict takes precedence over the deadline verdict:
// the allocation increment is a deterministic count, while a wall-clock
// deadline that fires on the very same heavy parse would otherwise split one
// out-of-budget parse between two columns at random. Timeout therefore counts
// only deadline overruns of parses that stayed inside the allocation ceiling.
// The two counters remain independent and never offset each other.
func (p ResourceProbe) Measure(parse func() (*air.Unit, []air.Diagnostic, error)) Outcome {
	deadline := time.Duration(p.budget.TimeoutMillis) * time.Millisecond
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	unit, diags, parseErr := parse()
	expired := errors.Is(ctx.Err(), context.DeadlineExceeded)
	cancel()
	runtime.ReadMemStats(&after)
	out := Outcome{
		AllocDeltaBytes: after.TotalAlloc - before.TotalAlloc,
		Diagnostics:     len(diags),
	}
	if parseErr != nil {
		out.Errors = 1
	}
	switch {
	case out.AllocDeltaBytes > p.budget.MemLimitBytes:
		out.Kind = metrics.OutcomeMemExceeded
	case expired:
		out.Kind = metrics.OutcomeTimeout
	case parseErr != nil || unit == nil:
		out.Kind = metrics.OutcomeContractError
	default:
		out.Kind = metrics.OutcomeOK
	}
	return out
}
