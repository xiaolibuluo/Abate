/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

import (
	"context"
	"testing"
)

// chaosSrc is a local deterministic integer source used only by chaos tests.
type chaosSrc struct{ state uint64 }

// next advances the generator with integer arithmetic.
func (g *chaosSrc) next() uint64 {
	g.state = g.state*6364136223846793005 + 1442695040888963407
	return g.state
}

// bounded returns a deterministic value in [0, n).
func (g *chaosSrc) bounded(n uint64) uint64 {
	return g.next() % n
}

// nanos returns a deterministic virtual time in [1, n].
func (g *chaosSrc) nanos(n int64) VirtualTime {
	return VirtualTime(g.next()%uint64(n)) + 1 //nolint:gosec // the modulo result is below n and fits int64.
}

// TestChaosQueueTotalOrder storms the queue with 10000 seeded events.
func TestChaosQueueTotalOrder(t *testing.T) {
	t.Parallel()
	g := chaosSrc{state: 0x9E3779B97F4A7C15}
	q := NewQueue()
	const n = 10000
	var seq uint64
	for i := 0; i < n; i++ {
		seq++
		e := Event{At: g.nanos(500), Seq: seq, Kind: EventTick}
		if err := q.Push(e); err != nil {
			t.Fatalf("Push %d: %v", i, err)
		}
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
			t.Fatalf("total order violated at event %d", count)
		}
		prev, first, count = e, false, count+1
	}
	if count != n {
		t.Fatalf("drained %d events, want %d", count, n)
	}
}

// TestChaosRunDigestStable replays a seeded event storm 100 times.
func TestChaosRunDigestStable(t *testing.T) {
	t.Parallel()
	g := chaosSrc{state: 0xD1B54A32D192ED03}
	kinds := []EventKind{EventTick, EventProbe, EventApply, EventRollback, EventTimeout}
	alphabet := []byte("abcdefghijklmnop")
	initial := make([]Event, 0, 200)
	for i := 0; i < 200; i++ {
		initial = append(initial, Event{
			At:      g.nanos(1000),
			Kind:    kinds[g.bounded(uint64(len(kinds)))],
			Payload: []byte{alphabet[g.bounded(uint64(len(alphabet)))]},
		})
	}
	handler := func(e Event) ([]Event, error) {
		if e.Kind == EventApply && e.At < 50 {
			return []Event{{At: e.At + 1, Kind: EventProbe, Payload: e.Payload}}, nil
		}
		return nil, nil
	}
	exp := Experiment{Seed: 99, Initial: initial}
	run := func() Result {
		return New(handler, DefaultBudget()).Run(context.Background(), exp)
	}
	baseline := run()
	if baseline.Status != StatusCompleted || baseline.Events < 200 {
		t.Fatalf("baseline = %+v", baseline)
	}
	for i := 0; i < 100; i++ {
		if got := run(); got.Digest != baseline.Digest || got.Events != baseline.Events {
			t.Fatalf("chaos run %d diverged", i)
		}
	}
}

// TestChaosBudgetNeverPanics runs seeded budgets without panics.
func TestChaosBudgetNeverPanics(t *testing.T) {
	t.Parallel()
	g := chaosSrc{state: 0x2545F4914F6CDD1D}
	initial := []Event{
		{At: 1, Kind: EventTick},
		{At: 2, Kind: EventTick},
		{At: 3, Kind: EventTick},
		{At: 4, Kind: EventTick},
	}
	for i := 0; i < 200; i++ {
		b := Budget{
			MaxEvents:          g.bounded(6),
			MaxVirtualDuration: g.nanos(10),
			MaxCost:            g.bounded(6),
		}
		res := New(noopHandler, b).Run(context.Background(), Experiment{Initial: initial})
		if res.Status == "" {
			t.Fatalf("iteration %d produced an empty status", i)
		}
	}
}
