/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

import "container/heap"

// eventKey is the total-order identity of an event.
type eventKey struct {
	at  VirtualTime
	seq uint64
}

// eventHeap is the backing min-heap ordered by (At, Seq).
type eventHeap []Event

// Len reports the heap size.
func (h eventHeap) Len() int {
	return len(h)
}

// Less orders by virtual time first and sequence second.
func (h eventHeap) Less(i, j int) bool {
	if h[i].At != h[j].At {
		return h[i].At < h[j].At
	}
	return h[i].Seq < h[j].Seq
}

// Swap exchanges two heap entries.
func (h eventHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push appends an event through the heap interface.
func (h *eventHeap) Push(x any) {
	e, ok := x.(Event)
	if ok {
		*h = append(*h, e)
	}
}

// Pop removes the last heap entry through the heap interface.
func (h *eventHeap) Pop() any {
	old := *h
	n := len(old)
	if n == 0 {
		return Event{}
	}
	e := old[n-1]
	*h = old[:n-1]
	return e
}

// Queue is a total-order priority queue keyed by (At, Seq).
//
// The zero value is ready to use. A duplicate key is rejected rather than
// silently reordered, which keeps replay unambiguous.
type Queue struct {
	items eventHeap
	seen  map[eventKey]struct{}
}

// NewQueue returns an empty queue.
func NewQueue() *Queue {
	return &Queue{seen: make(map[eventKey]struct{})}
}

// Push inserts an event, rejecting a duplicate (At, Seq) key.
func (q *Queue) Push(e Event) error {
	if q.seen == nil {
		q.seen = make(map[eventKey]struct{})
	}
	k := eventKey{at: e.At, seq: e.Seq}
	if _, exists := q.seen[k]; exists {
		return ErrSeqCollision
	}
	q.seen[k] = struct{}{}
	heap.Push(&q.items, e)
	return nil
}

// Pop removes and returns the earliest event in (At, Seq) order.
func (q *Queue) Pop() (Event, bool) {
	if len(q.items) == 0 {
		return Event{}, false
	}
	e, ok := heap.Pop(&q.items).(Event)
	if !ok {
		return Event{}, false
	}
	return e, true
}

// Len returns the number of queued events.
func (q *Queue) Len() int {
	return len(q.items)
}
