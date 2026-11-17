/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

// FaultKind enumerates the fault primitives understood by the file system.
type FaultKind string

// Fault kinds supported by the file-system injector.
const (
	// FaultLatency overrides the virtual nanoseconds charged per operation.
	FaultLatency FaultKind = "latency"

	// FaultBadBlock marks a byte range of a file as unreadable.
	FaultBadBlock FaultKind = "badblock"
)
