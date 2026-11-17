/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"abate/ces/schema"
	"abate/kernel/observe"
	"abate/kernel/scheduler"
	"abate/sim/clock"
	"abate/sim/rand"
	"abate/sim/record"
)

// rngStreamSalt decorrelates the recording sub-stream from the report seed so
// that consuming a draw for fault parameters never changes the report bytes.
const rngStreamSalt uint64 = 0x9e3779b97f4a7c15

// experimentHandler owns the virtual clock, random stream and recorder of one
// experiment run. The scheduler handler is wrapped by a read-only recorder, so
// every dispatched event becomes a SCHED frame without altering the event
// stream that the report digests.
type experimentHandler struct {
	base  scheduler.Handler
	clk   *clock.VirtualClock
	rng   *rand.Stream
	rec   *record.Recorder
	seed  uint64
	mixed func(seed uint64) ([]record.Frame, error)
}

// newExperimentHandler wires a clock, a seeded random sub-stream and a recorder
// subscribed to the real RNG accounting, clock state and observation stream.
func newExperimentHandler(exp *schema.Experiment, times []scheduler.VirtualTime,
	obs []observe.Observation,
) *experimentHandler {
	clk := clock.NewVirtualClock()
	rng := rand.New(exp.Spec.Seed ^ rngStreamSalt)
	rec := record.NewRecorder(experimentFingerprint(exp))
	at := clock.Instant(0)
	if len(times) > 0 {
		at = clock.Instant(times[len(times)-1].UnixNano())
	}
	rec.Subscribe(
		record.NewRNGSource(at, rng),
		record.NewCLKSource(clk),
		record.NewOBSSource(at, obs),
	)
	rec.Start(context.Background())
	return &experimentHandler{
		base:  cascadeHandler(times, probeInterval(exp)),
		clk:   clk,
		rng:   rng,
		rec:   rec,
		seed:  exp.Spec.Seed,
		mixed: mixedFrames,
	}
}

// experimentFingerprint derives a stable archive fingerprint from the CES
// identity and seed. No wall-clock value is ever written into an archive.
func experimentFingerprint(exp *schema.Experiment) [32]byte {
	key := fmt.Sprintf("%s/%s/%d", exp.Metadata.Namespace,
		exp.Metadata.Name, exp.Spec.Seed)
	return sha256.Sum256([]byte(key))
}

// handler returns the read-only recording wrapper around the cascade handler.
func (h *experimentHandler) handler() scheduler.Handler {
	return h.rec.WrapHandler(h.dispatch)
}

// dispatch records the perturbation injection of an apply event as a FAULT
// frame, folds the M⊗C mixed-perturbation frames derived by the L6 mutation
// engine and L7 injection kernel into the same recorder, advances the virtual
// clock to the scheduler instant, and delegates child-event derivation to the
// untouched cascade handler. The event stream returned to the scheduler is
// unchanged, so the report digest stays byte-stable while the archive gains
// the NET, FS and VTASK categories that only real SIM APIs can produce.
func (h *experimentHandler) dispatch(e scheduler.Event) ([]scheduler.Event, error) {
	h.clk.Sync(clock.Instant(e.At.UnixNano()))
	if e.Kind == scheduler.EventApply {
		at := clock.Instant(e.At.UnixNano())
		param := make([]byte, 8)
		binary.BigEndian.PutUint64(param, h.rng.Uint64())
		h.rec.Record(record.Frame{At: at, Kind: record.KindFault, Payload: param})
		frames, err := h.mixed(h.seed)
		if err != nil {
			return nil, fmt.Errorf("experiment: mixed perturbation: %w", err)
		}
		for _, frame := range frames {
			frame.At = at
			h.rec.Record(frame)
		}
	}
	return h.base(e)
}

// stop seals the recorded frames into a versioned archive.
func (h *experimentHandler) stop() (*record.Archive, error) {
	return h.rec.Stop()
}
