/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

// FaultKind enumerates the fault primitives understood by the scheduler.
type FaultKind string

// Fault kinds supported by the scheduler injector.
const (
	// FaultOOM force-kills a task as if it had breached its memory limit.
	FaultOOM FaultKind = "oom"

	// FaultSuspend suspends a task.
	FaultSuspend FaultKind = "suspend"

	// FaultThrottle lowers a cgroup's CPU share weight to Value.
	FaultThrottle FaultKind = "throttle"
)
