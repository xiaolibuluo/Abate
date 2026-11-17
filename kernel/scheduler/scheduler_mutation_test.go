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

// TestMutationEventStreamChangesDigest proves stream tampering is detected.
func TestMutationEventStreamChangesDigest(t *testing.T) {
	t.Parallel()
	base := []Event{
		{At: 1, Kind: EventTick, Payload: []byte("alpha")},
		{At: 2, Kind: EventApply, Payload: []byte("beta")},
		{At: 3, Kind: EventProbe, Payload: []byte("gamma")},
	}
	run := func(initial []Event) Result {
		return New(noopHandler, DefaultBudget()).Run(
			context.Background(), Experiment{Seed: 1, Initial: initial})
	}
	baseline := run(base).Digest
	copyOf := func() []Event {
		return append([]Event(nil), base...)
	}
	payloadMut := copyOf()
	payloadMut[1].Payload = []byte("BETA")
	timeMut := copyOf()
	timeMut[0].At = 4
	kindMut := copyOf()
	kindMut[2].Kind = EventRollback
	for i, m := range [][]Event{payloadMut, timeMut, kindMut} {
		if run(m).Digest == baseline {
			t.Fatalf("mutation %d was not detected", i)
		}
	}
	// Caller-supplied sequence numbers are overwritten by the scheduler.
	forged := copyOf()
	for i := range forged {
		forged[i].Seq = 999
	}
	if run(forged).Digest != baseline {
		t.Fatalf("caller-supplied Seq must not affect the digest")
	}
}

// TestMutationSeqCollisionRejected verifies a forged duplicate key is refused.
func TestMutationSeqCollisionRejected(t *testing.T) {
	t.Parallel()
	q := NewQueue()
	first := Event{At: 7, Seq: 3, Kind: EventTick}
	if err := q.Push(first); err != nil {
		t.Fatalf("first Push err = %v", err)
	}
	forged := Event{At: 7, Seq: 3, Kind: EventProbe}
	if err := q.Push(forged); !errors.Is(err, ErrSeqCollision) {
		t.Fatalf("forged duplicate err = %v, want ErrSeqCollision", err)
	}
}

// TestMutationBudgetBoundary checks the ±1 event-budget boundary.
func TestMutationBudgetBoundary(t *testing.T) {
	t.Parallel()
	initial := []Event{
		{At: 1, Kind: EventTick},
		{At: 2, Kind: EventTick},
		{At: 3, Kind: EventTick},
	}
	for _, max := range []uint64{2, 3} {
		res := New(noopHandler, Budget{MaxEvents: max}).Run(
			context.Background(), Experiment{Initial: initial})
		want := StatusBudgetExceeded
		if max == 3 {
			want = StatusCompleted
		}
		if res.Status != want {
			t.Fatalf("MaxEvents=%d status = %s, want %s", max, res.Status, want)
		}
	}
}
