/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

import (
	"context"
	"crypto/sha256"
	"fmt"
	"hash"
)

// Scheduler is a deterministic discrete-event driver.
//
// A single instance is not safe for concurrent use: concurrency is confined to
// the simulator boundary, and every concurrent producer must hand its events
// to the scheduler, which alone assigns the total-order sequence.
type Scheduler struct {
	queue   *Queue
	budget  Budget
	handler Handler
	now     VirtualTime
	start   VirtualTime
	seq     uint64
	events  uint64
	cost    uint64
	digest  hash.Hash
}

// New constructs a scheduler with the given handler and budget.
func New(handler Handler, budget Budget) *Scheduler {
	return &Scheduler{handler: handler, budget: budget}
}

// Now returns the current virtual time of the scheduler.
func (s *Scheduler) Now() VirtualTime {
	return s.now
}

// Emit enqueues a derived event, assigning the next total-order sequence.
//
// Any Seq supplied by the caller is overwritten, so handlers cannot forge an
// ordering. A negative time or an empty kind is rejected.
func (s *Scheduler) Emit(e Event) error {
	if e.At < 0 || e.Kind == "" {
		return ErrInvalidEvent
	}
	if s.queue == nil {
		s.queue = NewQueue()
	}
	e.Seq = s.nextSeq()
	return s.queue.Push(e)
}

// nextSeq returns the next monotonically increasing sequence number.
func (s *Scheduler) nextSeq() uint64 {
	s.seq++
	return s.seq
}

// Run executes exp to completion and returns a deterministic Result.
//
// Initial events are enqueued in declaration order, then the loop pops the
// earliest event, advances the virtual clock, folds it into the digest and
// dispatches the handler. The loop stops on an empty queue, a budget breach,
// context cancellation or a handler error.
func (s *Scheduler) Run(ctx context.Context, exp Experiment) Result {
	s.reset()
	if s.handler == nil {
		return s.finish(StatusFailed, ErrNilHandler)
	}
	for _, e := range exp.Initial {
		if err := s.Emit(e); err != nil {
			return s.finish(StatusFailed, err)
		}
	}
	for {
		if err := ctx.Err(); err != nil {
			return s.finish(StatusCanceled, err)
		}
		e, ok := s.queue.Pop()
		if !ok {
			return s.finish(StatusCompleted, nil)
		}
		if err := s.budget.Allows(s.events+1, e.At.Sub(s.start), s.cost); err != nil {
			return s.finish(StatusBudgetExceeded, err)
		}
		s.now = e.At
		s.events++
		s.record(e)
		children, err := s.handler(e)
		if err != nil {
			return s.finish(StatusFailed, fmt.Errorf("%w: %w", ErrHandlerFailed, err))
		}
		for _, child := range children {
			if err := s.Emit(child); err != nil {
				return s.finish(StatusFailed, err)
			}
		}
	}
}

// reset clears all run state and starts a fresh digest.
func (s *Scheduler) reset() {
	s.queue = NewQueue()
	s.now = 0
	s.start = 0
	s.seq = 0
	s.events = 0
	s.cost = 0
	s.digest = sha256.New()
}

// record folds one event into the run digest and charges one unit of cost.
func (s *Scheduler) record(e Event) {
	_, _ = fmt.Fprintf(s.digest, "%d/%d/%s/%x", e.At, e.Seq, e.Kind, e.Payload)
	s.cost++
}

// finish materializes the run result.
func (s *Scheduler) finish(status string, err error) Result {
	var sum [32]byte
	if s.digest != nil {
		copy(sum[:], s.digest.Sum(nil))
	}
	return Result{
		Status: status,
		Events: s.events,
		Cost:   s.cost,
		Start:  s.start,
		End:    s.now,
		Digest: sum,
		Err:    err,
	}
}
