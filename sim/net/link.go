/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import "abate/sim/clock"

// Link is a virtual bidirectional link between two nodes.
//
// It carries the four link-level fault primitives (partition, latency, loss and
// bandwidth) used by the kernel fault injector. Recover restores the baseline.
type Link struct {
	a, b        string
	cfg         LinkConfig
	base        LinkConfig
	partitioned bool
	bucket      tokenBucket
}

// A returns the first endpoint id of the link.
func (l *Link) A() string { return l.a }

// B returns the second endpoint id of the link.
func (l *Link) B() string { return l.b }

// Config returns the active link configuration.
func (l *Link) Config() LinkConfig { return l.cfg }

// Partitioned reports whether the link is currently partitioned.
func (l *Link) Partitioned() bool { return l.partitioned }

// Partition makes the link bidirectionally unreachable.
func (l *Link) Partition() error {
	l.partitioned = true
	return nil
}

// Heal restores reachability after a partition.
func (l *Link) Heal() error {
	l.partitioned = false
	return nil
}

// SetLatency overrides the base latency of the link.
func (l *Link) SetLatency(d clock.Duration) { l.cfg.BaseLatency = d }

// SetLossRate overrides the packet loss rate of the link.
func (l *Link) SetLossRate(r float64) { l.cfg.LossRate = r }

// SetBandwidth overrides the link bandwidth in bytes per second.
func (l *Link) SetBandwidth(bps uint64) { l.cfg.BandwidthBps = bps }

// Recover restores the link to its baseline configuration and heals partitions.
//
// The bandwidth bucket is reset as well, so a recovered link starts with a full
// byte-token credit instead of carrying the queued debt of the faulted state.
func (l *Link) Recover() {
	l.cfg = l.base
	l.partitioned = false
	l.bucket.reset()
}
