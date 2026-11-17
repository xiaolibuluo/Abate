/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"abate/kernel/observe"
	"abate/sim/clock"
)

// OBSSource exposes a collected observation stream as OBS frames.
type OBSSource struct {
	at  clock.Instant
	obs []observe.Observation
}

// NewOBSSource builds a source that records the given observations at the instant.
func NewOBSSource(at clock.Instant, obs []observe.Observation) *OBSSource {
	return &OBSSource{at: at, obs: obs}
}

// Collect returns one frame per observation, encoding its canonical bytes.
func (s *OBSSource) Collect() []Frame {
	out := make([]Frame, 0, len(s.obs))
	for _, o := range s.obs {
		out = append(out, Frame{At: s.at, Kind: KindObs, Payload: encodeObservation(o)})
	}
	return out
}
