/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

// Event is one fully ordered item of the discrete-event stream.
//
// At is the virtual time at which the event becomes due, Seq is assigned by
// the scheduler (never by a caller), Kind selects the dispatch branch and
// Payload is an opaque deterministic byte string.
type Event struct {
	At      VirtualTime
	Seq     uint64
	Kind    EventKind
	Payload []byte
}
