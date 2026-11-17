/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import "abate/sim/clock"

// CLKSource records the final state of a virtual clock as a single CLK frame.
type CLKSource struct {
	clk *clock.VirtualClock
}

// NewCLKSource builds a source that snapshots the clock when the recorder stops.
func NewCLKSource(clk *clock.VirtualClock) *CLKSource {
	return &CLKSource{clk: clk}
}

// Collect returns one frame carrying the clock instant and four-state mode.
func (s *CLKSource) Collect() []Frame {
	now := s.clk.Now()
	return []Frame{{
		At:      now,
		Kind:    KindClk,
		Payload: encodeClock(now, s.clk.State()),
	}}
}
