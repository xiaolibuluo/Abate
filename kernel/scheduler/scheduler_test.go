/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

import (
	"context"
	"errors"
	"testing"
)

// noopHandler returns no derived events.
func noopHandler(Event) ([]Event, error) {
	return nil, nil
}

// TestVirtualTimeHelpers covers the integer-time helper surface.
func TestVirtualTimeHelpers(t *testing.T) {
	t.Parallel()
	if Second != 1_000_000_000 || Nanosecond != 1 {
		t.Fatalf("base units broken: %d %d", Nanosecond, Second)
	}
	if Microsecond != 1000*Nanosecond || Millisecond != 1000*Microsecond ||
		Minute != 60*Second || Hour != 60*Minute {
		t.Fatalf("unit ladder broken")
	}
	base := VirtualTime(100)
	if base.Add(5) != 105 || base.Sub(VirtualTime(40)) != 60 {
		t.Fatalf("Add/Sub broken")
	}
	if !base.Before(200) || !base.After(50) || base.Before(100) || base.After(100) {
		t.Fatalf("Before/After broken")
	}
	if base.UnixNano() != 100 {
		t.Fatalf("UnixNano = %d", base.UnixNano())
	}
}

// TestQueueTotalOrder verifies strict (At, Seq) ordering under shuffled input.
func TestQueueTotalOrder(t *testing.T) {
	t.Parallel()
	q := NewQueue()
	const n = 2000
	var seq uint64
	for i := 0; i < n; i++ {
		seq++
		e := Event{At: VirtualTime((i * 7919) % 97), Seq: seq, Kind: EventTick}
		if err := q.Push(e); err != nil {
			t.Fatalf("Push %d: %v", i, err)
		}
	}
	if q.Len() != n {
		t.Fatalf("Len = %d, want %d", q.Len(), n)
	}
	var prev Event
	first := true
	count := 0
	for {
		e, ok := q.Pop()
		if !ok {
			break
		}
		if !first && (e.At < prev.At || (e.At == prev.At && e.Seq <= prev.Seq)) {
			t.Fatalf("out of order: %v after %v", e, prev)
		}
		prev, first, count = e, false, count+1
	}
	if count != n || q.Len() != 0 {
		t.Fatalf("drained %d events, remaining %d", count, q.Len())
	}
}

// TestQueueSeqCollision verifies duplicate keys are rejected.
func TestQueueSeqCollision(t *testing.T) {
	t.Parallel()
	q := NewQueue()
	e := Event{At: 5, Seq: 1, Kind: EventTick}
	if err := q.Push(e); err != nil {
		t.Fatalf("first Push err = %v", err)
	}
	if err := q.Push(e); !errors.Is(err, ErrSeqCollision) {
		t.Fatalf("duplicate Push err = %v, want ErrSeqCollision", err)
	}
	if err := q.Push(Event{At: 5, Seq: 2, Kind: EventTick}); err != nil {
		t.Fatalf("distinct Seq err = %v", err)
	}
	if q.Len() != 2 {
		t.Fatalf("Len = %d, want 2", q.Len())
	}
}

// TestQueueZeroValueAndDefensivePaths covers lazy init and typed-heap guards.
func TestQueueZeroValueAndDefensivePaths(t *testing.T) {
	t.Parallel()
	var q Queue
	if err := q.Push(Event{At: 1, Seq: 1, Kind: EventTick}); err != nil {
		t.Fatalf("zero-value Push err = %v", err)
	}
	if _, ok := q.Pop(); !ok {
		t.Fatalf("Pop must return the queued event")
	}
	if _, ok := q.Pop(); ok {
		t.Fatalf("empty Pop must report false")
	}
	var h eventHeap
	h.Push("not-an-event")
	if h.Len() != 0 {
		t.Fatalf("non-event Push must be ignored")
	}
	got, ok := h.Pop().(Event)
	if !ok || got.At != 0 || got.Seq != 0 || got.Kind != "" || got.Payload != nil {
		t.Fatalf("empty heap Pop = %+v, want zero Event", got)
	}
}

// TestBudgetAllows covers all three limits and the unbounded zero budget.
func TestBudgetAllows(t *testing.T) {
	t.Parallel()
	if err := DefaultBudget().Allows(1, Second, 1); err != nil {
		t.Fatalf("in-budget err = %v", err)
	}
	if !errors.Is(DefaultBudget().Allows(DefaultMaxEvents+1, 0, 0), ErrBudgetEvents) {
		t.Fatalf("event limit not enforced")
	}
	if !errors.Is(DefaultBudget().Allows(0, DefaultMaxVirtualDuration+1, 0), ErrBudgetDuration) {
		t.Fatalf("duration limit not enforced")
	}
	if !errors.Is(DefaultBudget().Allows(0, 0, DefaultMaxCost+1), ErrBudgetCost) {
		t.Fatalf("cost limit not enforced")
	}
	if err := (Budget{}).Allows(1<<62, VirtualTime(1<<62), 1<<62); err != nil {
		t.Fatalf("zero budget must be unbounded: %v", err)
	}
}

// TestRunDeterministic proves 100 double runs produce identical digests.
func TestRunDeterministic(t *testing.T) {
	t.Parallel()
	exp := Experiment{Seed: 7, Initial: []Event{
		{At: 30, Kind: EventApply, Payload: []byte("c")},
		{At: 10, Kind: EventTick, Payload: []byte("a")},
		{At: 20, Kind: EventProbe, Payload: []byte("b")},
	}}
	run := func() Result {
		return New(noopHandler, DefaultBudget()).Run(context.Background(), exp)
	}
	first := run()
	if first.Status != StatusCompleted || first.Events != 3 || first.Cost != 3 || first.End != 30 {
		t.Fatalf("baseline result = %+v", first)
	}
	for i := 0; i < 100; i++ {
		got := run()
		if got.Digest != first.Digest || got.Status != first.Status || got.Events != first.Events {
			t.Fatalf("run %d diverged: %+v", i, got)
		}
	}
}

// TestRunDerivedEvents checks handler-produced events are queued and ordered.
func TestRunDerivedEvents(t *testing.T) {
	t.Parallel()
	handler := func(e Event) ([]Event, error) {
		if e.Kind == EventTick && e.At < 3 {
			return []Event{{At: e.At + 1, Kind: EventTick, Payload: e.Payload}}, nil
		}
		return nil, nil
	}
	s := New(handler, DefaultBudget())
	res := s.Run(context.Background(), Experiment{Initial: []Event{{At: 1, Kind: EventTick}}})
	if res.Status != StatusCompleted || res.Events != 3 || s.Now() != 3 {
		t.Fatalf("derived run = %+v, now = %d", res, s.Now())
	}
}

// TestRunBudgetBreaches covers each budget axis through the run loop.
func TestRunBudgetBreaches(t *testing.T) {
	t.Parallel()
	initial := []Event{
		{At: 1, Kind: EventTick},
		{At: 2, Kind: EventTick},
		{At: 3, Kind: EventTick},
	}
	cases := []struct {
		budget Budget
		want   error
	}{
		{Budget{MaxEvents: 2}, ErrBudgetEvents},
		{Budget{MaxVirtualDuration: 1}, ErrBudgetDuration},
		{Budget{MaxCost: 1}, ErrBudgetCost},
	}
	for _, c := range cases {
		res := New(noopHandler, c.budget).Run(context.Background(), Experiment{Initial: initial})
		if res.Status != StatusBudgetExceeded || !errors.Is(res.Err, c.want) {
			t.Fatalf("budget %+v -> %+v", c.budget, res)
		}
	}
}

// TestRunCancellation checks a canceled context terminates before dispatch.
func TestRunCancellation(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	res := New(noopHandler, DefaultBudget()).Run(ctx, Experiment{
		Initial: []Event{{At: 1, Kind: EventTick}},
	})
	if res.Status != StatusCanceled || !errors.Is(res.Err, context.Canceled) || res.Events != 0 {
		t.Fatalf("canceled run = %+v", res)
	}
}

// TestRunFailureModes covers handler errors, a nil handler and bad events.
func TestRunFailureModes(t *testing.T) {
	t.Parallel()
	boom := errors.New("boom")
	res := New(func(Event) ([]Event, error) { return nil, boom }, DefaultBudget()).Run(
		context.Background(), Experiment{Initial: []Event{{At: 1, Kind: EventTick}}})
	if res.Status != StatusFailed || !errors.Is(res.Err, ErrHandlerFailed) || !errors.Is(res.Err, boom) {
		t.Fatalf("handler error run = %+v", res)
	}
	res = New(nil, DefaultBudget()).Run(context.Background(), Experiment{})
	if res.Status != StatusFailed || !errors.Is(res.Err, ErrNilHandler) {
		t.Fatalf("nil handler run = %+v", res)
	}
	res = New(noopHandler, DefaultBudget()).Run(context.Background(), Experiment{
		Initial: []Event{{At: -1, Kind: EventTick}},
	})
	if res.Status != StatusFailed || !errors.Is(res.Err, ErrInvalidEvent) {
		t.Fatalf("invalid initial run = %+v", res)
	}
	res = New(func(Event) ([]Event, error) {
		return []Event{{At: -1, Kind: EventTick}}, nil
	}, DefaultBudget()).Run(context.Background(), Experiment{
		Initial: []Event{{At: 1, Kind: EventTick}},
	})
	if res.Status != StatusFailed || !errors.Is(res.Err, ErrInvalidEvent) {
		t.Fatalf("invalid child run = %+v", res)
	}
}

// TestEmitValidation checks Emit overwrites caller Seq and rejects bad events.
func TestEmitValidation(t *testing.T) {
	t.Parallel()
	s := New(noopHandler, DefaultBudget())
	if err := s.Emit(Event{At: -1, Kind: EventTick}); !errors.Is(err, ErrInvalidEvent) {
		t.Fatalf("negative time err = %v", err)
	}
	if err := s.Emit(Event{At: 1}); !errors.Is(err, ErrInvalidEvent) {
		t.Fatalf("empty kind err = %v", err)
	}
	if err := s.Emit(Event{At: 1, Seq: 999, Kind: EventTick}); err != nil {
		t.Fatalf("valid Emit err = %v", err)
	}
	if s.queue.Len() != 1 {
		t.Fatalf("queue Len = %d", s.queue.Len())
	}
	e, ok := s.queue.Pop()
	if !ok || e.Seq != 1 {
		t.Fatalf("Emit must assign Seq 1, got (%v, %v)", e, ok)
	}
	if s.Now() != 0 {
		t.Fatalf("Now = %d", s.Now())
	}
}

// TestFinishWithoutRun covers the digest-absent finalization branch.
func TestFinishWithoutRun(t *testing.T) {
	t.Parallel()
	res := New(noopHandler, DefaultBudget()).finish(StatusFailed, nil)
	if res.Digest != ([32]byte{}) || res.Status != StatusFailed {
		t.Fatalf("finish without run = %+v", res)
	}
}
