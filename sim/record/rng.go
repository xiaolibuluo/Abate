/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"abate/sim/clock"
	"abate/sim/rand"
)

// RNGSource exposes the deterministic draw accounting of a random stream as a
// sequence of RNG frames for the recorder to append at Stop.
type RNGSource struct {
	at clock.Instant
	s  *rand.Stream
}

// NewRNGSource builds a source that records every draw of s at the given instant.
func NewRNGSource(at clock.Instant, s *rand.Stream) *RNGSource {
	return &RNGSource{at: at, s: s}
}

// Collect returns one frame per recorded draw, carrying its index, method kind,
// arguments and raw result so a replay can prove identical random consumption.
func (s *RNGSource) Collect() []Frame {
	records := s.s.Snapshot()
	out := make([]Frame, 0, len(records))
	for _, d := range records {
		out = append(out, Frame{At: s.at, Kind: KindRNG, Payload: encodeDraw(d)})
	}
	return out
}
