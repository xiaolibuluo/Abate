/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package fs

// Fault is a file-system fault-injection specification.
//
// For FaultLatency, Value is the per-operation virtual latency in nanoseconds.
// For FaultBadBlock, Path selects the target file, Offset is the start of the
// faulted range and Value is its length in bytes.
type Fault struct {
	// Kind selects the fault primitive.
	Kind FaultKind

	// Path is the target entry (required for FaultBadBlock).
	Path string

	// Offset is the start of a bad-block range in bytes.
	Offset int64

	// Value is the fault parameter interpreted per Kind.
	Value int64
}
