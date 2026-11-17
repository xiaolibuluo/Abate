/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

// Fault is a scheduler fault-injection specification.
//
// Target selects the task id (FaultOOM, FaultSuspend) or cgroup name
// (FaultThrottle). Value is the new share weight for FaultThrottle.
type Fault struct {
	// Kind selects the fault primitive.
	Kind FaultKind

	// Target is the task id or cgroup name affected by the fault.
	Target string

	// Value is the fault parameter interpreted per Kind.
	Value int64
}
