/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

import "sync"

// AuditLog is an append-only, concurrency-safe record of capability events.
//
// The At field of every appended event is assigned from a virtual monotonic
// counter (never time.Now), so the log is fully deterministic: the same sequence
// of appends always yields the same entries in the same order.
type AuditLog struct {
	mu     sync.Mutex
	seq    int64
	events []AuditEvent
}

// NewAuditLog returns an empty audit log.
func NewAuditLog() *AuditLog {
	return &AuditLog{}
}

// Append records event, overriding its At field with the next virtual sequence
// number. It is safe for concurrent use.
func (l *AuditLog) Append(event AuditEvent) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.seq++
	event.At = l.seq
	l.events = append(l.events, event)
}

// Entries returns a copy of all recorded events in append order. The returned
// slice is owned by the caller and may be mutated freely.
func (l *AuditLog) Entries() []AuditEvent {
	l.mu.Lock()
	defer l.mu.Unlock()
	out := make([]AuditEvent, len(l.events))
	copy(out, l.events)
	return out
}

// Len returns the number of recorded events.
func (l *AuditLog) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.events)
}
