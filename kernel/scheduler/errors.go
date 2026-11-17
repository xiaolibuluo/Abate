/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

import "errors"

// Sentinel errors of the scheduler package.
var (
	// ErrSeqCollision reports a duplicate (At, Seq) key on enqueue.
	ErrSeqCollision = errors.New("scheduler: duplicate event sequence key")

	// ErrBudgetEvents reports the event-count budget being exceeded.
	ErrBudgetEvents = errors.New("scheduler: event budget exceeded")

	// ErrBudgetDuration reports the virtual-duration budget being exceeded.
	ErrBudgetDuration = errors.New("scheduler: virtual duration budget exceeded")

	// ErrBudgetCost reports the operator-cost budget being exceeded.
	ErrBudgetCost = errors.New("scheduler: operator cost budget exceeded")

	// ErrInvalidEvent reports an event with a negative time or an empty kind.
	ErrInvalidEvent = errors.New("scheduler: event must have non-negative time and a kind")

	// ErrNilHandler reports a run started without an event handler.
	ErrNilHandler = errors.New("scheduler: event handler must not be nil")

	// ErrHandlerFailed reports a handler error, wrapped with its cause.
	ErrHandlerFailed = errors.New("scheduler: event handler failed")
)
