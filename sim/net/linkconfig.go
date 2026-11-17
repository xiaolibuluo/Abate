/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import "abate/sim/clock"

// LinkConfig configures a virtual link.
type LinkConfig struct {
	// BaseLatency is the per-hop propagation delay in virtual nanoseconds.
	BaseLatency clock.Duration

	// Jitter is the maximum extra latency drawn per hop via the SIM random stream.
	Jitter clock.Duration

	// BandwidthBps is the link capacity in bytes per virtual second.
	BandwidthBps uint64

	// LossRate is the per-hop drop probability in [0, 1] resolved via the stream.
	LossRate float64
}
