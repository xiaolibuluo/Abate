/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

import (
	"crypto/sha256"
	"slices"
)

// Pipeline normalizes probe samples and fans them out to attached sinks.
//
// The zero value is ready to use. Probe IDs are kept sorted so that collection
// order is independent of registration order.
type Pipeline struct {
	probes map[string]Probe
	order  []string
	sinks  []Sink
}

// NewPipeline returns an empty pipeline.
func NewPipeline() *Pipeline {
	return &Pipeline{probes: make(map[string]Probe)}
}

// Register adds a probe, rejecting a nil probe, an empty ID or a duplicate.
func (p *Pipeline) Register(probe Probe) error {
	if probe == nil {
		return ErrNilProbe
	}
	if probe.ID() == "" {
		return ErrInvalidProbeID
	}
	if p.probes == nil {
		p.probes = make(map[string]Probe)
	}
	if _, exists := p.probes[probe.ID()]; exists {
		return ErrDuplicateProbe
	}
	p.probes[probe.ID()] = probe
	p.order = append(p.order, probe.ID())
	slices.Sort(p.order)
	return nil
}

// Attach adds a sink to the fan-out list.
func (p *Pipeline) Attach(sink Sink) error {
	if sink == nil {
		return ErrNilSink
	}
	p.sinks = append(p.sinks, sink)
	return nil
}

// Collect samples every probe at one virtual time and returns the ordered
// stream, starting sequence numbers at seq.
//
// Sink failures are isolated: they append a meta observation with failed
// status and never abort the remaining sinks.
func (p *Pipeline) Collect(at int64, seq uint64) []Observation {
	obs := make([]Observation, 0, len(p.order))
	next := seq
	for _, id := range p.order {
		probe := p.probes[id]
		value, status := probe.Collect(at)
		o := Observation{At: at, Seq: next, ProbeID: id, Value: value, Status: status}
		obs = append(obs, o)
		next++
		for _, sink := range p.sinks {
			if err := sink.Observe(o); err != nil {
				meta := Observation{
					At:      at,
					Seq:     next,
					ProbeID: "meta/sink-error",
					Value:   NewText(err.Error()),
					Status:  StatusFailed,
				}
				obs = append(obs, meta)
				next++
			}
		}
	}
	return obs
}

// StreamDigest computes the SHA-256 digest of an observation batch.
//
// A record whose value cannot be encoded (for example a NaN number) is skipped
// so that one malformed sample cannot abort the whole stream hash.
func StreamDigest(obs []Observation) [32]byte {
	h := sha256.New()
	for _, o := range obs {
		data, err := o.Encode()
		if err != nil {
			continue
		}
		_, _ = h.Write(data)
	}
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out
}
