/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

// FaultKind enumerates the four link-level fault primitives.
type FaultKind string

// Fault kinds supported by the link injector.
const (
	// FaultLatency overrides the link base latency (Value nanoseconds).
	FaultLatency FaultKind = "latency"

	// FaultPartition partitions the link (Value ignored).
	FaultPartition FaultKind = "partition"

	// FaultLoss sets the loss rate to Value parts per million.
	FaultLoss FaultKind = "loss"

	// FaultBandwidth overrides the link bandwidth (Value bytes per second).
	FaultBandwidth FaultKind = "bandwidth"
)
