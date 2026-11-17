/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"context"
	"sync"

	"abate/kernel/scheduler"
	"abate/sim/clock"
)

// Recorder captures a deterministic experiment run into a replayable archive by
// subscribing to read-only sources and wrapping the scheduler handler, so that no
// recorded frame ever changes the event stream it observes.
type Recorder struct {
	mu      sync.Mutex
	fp      [32]byte
	sources []Source
	frames  []Frame
	maxAt   clock.Instant
}

// NewRecorder builds a recorder that brands its archives with the given experiment
// fingerprint, so two runs of the same experiment yield byte-identical archives.
func NewRecorder(fp [32]byte) *Recorder {
	return &Recorder{fp: fp}
}

// Subscribe attaches read-only sources whose frames are pulled once at Stop.
func (r *Recorder) Subscribe(sources ...Source) {
	r.mu.Lock()
	r.sources = append(r.sources, sources...)
	r.mu.Unlock()
}

// Start is an intentional, contractually-inert no-op. Recording is pull-based:
// frames enter only through Record and WrapHandler, subscribed sources are
// collected once at Stop, and no lifecycle state exists to activate. It exists
// solely for interface symmetry with Stop so orchestrators may call Start on
// every run without branching. Start must ignore the context and must never
// mutate state or add frames; TestRecorderStartIsInert and
// TestRecorderStartConcurrentWithRecord pin that contract.
func (*Recorder) Start(_ context.Context) {}

// Record pushes a single frame, used by orchestrators to log faults, clock moves
// and other injected events, and assigns it the next total-order sequence.
func (r *Recorder) Record(f Frame) {
	r.push(f)
}

// WrapHandler decorates a scheduler handler with zero-intrusion recording: each
// dispatched event is logged as a SCHED frame before the underlying handler runs,
// and the returned event stream is left untouched.
func (r *Recorder) WrapHandler(h scheduler.Handler) scheduler.Handler {
	return func(e scheduler.Event) ([]scheduler.Event, error) {
		r.push(Frame{At: clock.Instant(e.At), Kind: KindSched, Payload: encodeEvent(e)})
		return h(e)
	}
}

// push appends a frame under the recorder lock, updating the monotonic time bound.
func (r *Recorder) push(f Frame) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if f.At > r.maxAt {
		r.maxAt = f.At
	}
	r.frames = append(r.frames, f)
}

// Stop finalizes the recording: it pulls every subscribed source, stamps any
// unstamped frame with the monotonic time bound, assigns the total-order sequence
// and returns the sealed archive.
func (r *Recorder) Stop() (*Archive, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	all := make([]Frame, 0, len(r.frames)+len(r.sources))
	all = append(all, r.frames...)
	for _, src := range r.sources {
		for _, f := range src.Collect() {
			if f.At > r.maxAt {
				r.maxAt = f.At
			}
			all = append(all, f)
		}
	}
	for i := range all {
		if all[i].At == 0 {
			all[i].At = r.maxAt
		}
		all[i].Seq = uint64(i) + 1 //nolint:gosec // frame count is bounded by the experiment size.
	}
	return newArchive(r.fp, all)
}
