/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

// Run status values reported by Scheduler.Run.
const (
	StatusCompleted      = "COMPLETED"
	StatusBudgetExceeded = "BUDGET_EXCEEDED"
	StatusCanceled       = "CANCELED"
	StatusFailed         = "FAILED"
)

// Result is the deterministic outcome of one experiment run.
//
// Digest is a SHA-256 over the ordered event stream and is the primary
// evidence used by the double-run determinism check.
type Result struct {
	Status string
	Events uint64
	Cost   uint64
	Start  VirtualTime
	End    VirtualTime
	Digest [32]byte
	Err    error
}
