/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package rand provides the deterministic pseudo-random stream of the Abate
// simulator.
//
// Every random value in an experiment flows through a Stream seeded from a single
// 64-bit seed. The same seed always yields the same sequence, and State/Restore
// lets a replay jump to any draw ordinal without re-simulating the prefix.
package rand

import "errors"

// ErrInvalidRestore reports a Restore call with a seed that disagrees with the
// stream's origin, which would make subsequent values non-reproducible.
var ErrInvalidRestore = errors.New("rand: restore origin disagrees with stream")

// Stream is a deterministic SplitMix64 random stream with replay accounting.
//
// It is the only sanctioned source of randomness inside the simulator. Using
// math/rand or crypto/rand would break the determinism and air-gap guarantees.
type Stream struct {
	origin uint64
	sm     *splitMix64
	draws  uint64
	acct   *Accounting
}

// New constructs a stream deterministically from a 64-bit seed.
func New(seed uint64) *Stream {
	return &Stream{
		origin: seed,
		sm:     newSplitMix64(seed),
		draws:  0,
		acct:   newAccounting(),
	}
}

// drawNoArgs advances the stream and records a draw that carries no arguments.
//
// Keeping the argument-free case on a non-variadic path guarantees that no
// slice is ever built for the majority of draws, so the ledger stores a nil
// Args value exactly as the original variadic call did when given no arguments.
func (s *Stream) drawNoArgs(kind uint8) uint64 {
	v := s.sm.next()
	s.draws++
	s.acct.record(kind, v)
	return v
}

// drawArg advances the stream and records a draw carrying one integer argument.
//
// The argument is passed as a value instead of a variadic slice, so the hot path
// allocates nothing; Snapshot rebuilds the one-element Args slice with the same
// value and encoding as the original one-argument variadic call.
func (s *Stream) drawArg(kind uint8, arg uint64) uint64 {
	v := s.sm.next()
	s.draws++
	s.acct.recordArg(kind, arg, v)
	return v
}

// Uint64 returns the next 64-bit value of the stream.
func (s *Stream) Uint64() uint64 {
	return s.drawNoArgs(drawKindUint64)
}

// Int63 returns a non-negative 63-bit integer in [0, 2^63).
func (s *Stream) Int63() int64 {
	return int64(s.drawNoArgs(drawKindInt63) & (1<<63 - 1)) //nolint:gosec // value is bounded by a bitmask; conversion is safe.
}

// Float64 returns a pseudo-random float64 in the half-open interval [0, 1).
//
// The scale is 1/2^53: a float64 has a 52-bit mantissa, so drawing 53 bits
// and dividing by 2^53 yields a uniform value in [0, 1). The exponent is built
// from a shift rather than a bare decimal literal.
func (s *Stream) Float64() float64 {
	return float64(s.drawNoArgs(drawKindFloat64)>>11) * (1.0 / float64(uint64(1)<<53))
}

// IntN returns a uniform integer in [0, n) using unbiased rejection sampling.
//
// For n <= 1 it degenerates to 0, which is the only value in an empty range.
func (s *Stream) IntN(n int) int {
	if n <= 1 {
		return 0
	}
	limit := ^uint64(0) - (^uint64(0) % uint64(n))
	for {
		v := s.drawArg(drawKindIntN, uint64(n))
		if v < limit {
			return int(v % uint64(n)) //nolint:gosec // value is bounded by a modulo; conversion is safe.
		}
	}
}

// Bool returns a uniformly distributed boolean.
func (s *Stream) Bool() bool {
	return s.drawNoArgs(drawKindBool)&1 == 1
}

// Bytes returns n bytes of pseudo-random data.
func (s *Stream) Bytes(n int) []byte {
	out := make([]byte, n)
	for i := 0; i < n; i += 8 {
		v := s.drawArg(drawKindBytes, uint64(n)) //nolint:gosec // n is a non-negative byte count.
		for j := 0; j < 8 && i+j < n; j++ {
			out[i+j] = byte(v >> (8 * uint(j)))
		}
	}
	return out
}

// Fork derives an independent child stream from the next draw of the parent.
//
// After Fork the two streams evolve independently, which lets parallel
// experiments consume disjoint, reproducible sub-streams.
func (s *Stream) Fork() *Stream {
	return New(s.drawNoArgs(drawKindFork))
}

// State returns the stream origin and the number of draws consumed so far.
func (s *Stream) State() (seed, draws uint64) {
	return s.origin, s.draws
}

// Restore jumps the stream to a draw ordinal without replaying the prefix.
//
// The seed must equal the stream origin; otherwise the replay would no longer
// be reproducible and ErrInvalidRestore is returned.
func (s *Stream) Restore(seed uint64, draws uint64) error {
	if seed != s.origin {
		return ErrInvalidRestore
	}
	s.sm.state = seed + draws*splitMixGamma
	s.draws = draws
	s.acct = newAccounting()
	return nil
}

// Snapshot returns a defensive copy of the random draw ledger for replay
// certification. It exposes the internal accounting so SIM-REC can prove the
// exact stream an experiment consumed from its seed.
func (s *Stream) Snapshot() []DrawRecord {
	return s.acct.Snapshot()
}
