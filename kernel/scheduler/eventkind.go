/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

// EventKind classifies one scheduled event for dispatch and reporting.
type EventKind string

// Event kind constants of the discrete-event loop.
const (
	EventTick     EventKind = "tick"
	EventProbe    EventKind = "probe"
	EventApply    EventKind = "apply"
	EventRollback EventKind = "rollback"
	EventTimeout  EventKind = "timeout"
)
