/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rand

// splitMixGamma is the SplitMix64 golden-ratio step constant.
const splitMixGamma = 0x9E3779B97F4A7C15

// splitMixMulA and splitMixMulB are the two SplitMix64 mixing multipliers
// from the public-domain reference. They are named to avoid bare magic
// constants in the mixing step below.
const (
	splitMixMulA = 0xBF58476D1CE4E5B9
	splitMixMulB = 0x94D049BB133111EB
)

// splitMix64 is the private SplitMix64 generator state machine.
//
// It is public-domain (no license burden) and integer-only, so it satisfies the
// simulator's determinism and air-gap constraints. The exported Stream wraps it.
type splitMix64 struct {
	state uint64
}

// newSplitMix64 seeds a fresh generator.
func newSplitMix64(seed uint64) *splitMix64 {
	return &splitMix64{state: seed}
}

// next advances the state and returns a mixed 64-bit value.
func (s *splitMix64) next() uint64 {
	s.state += splitMixGamma
	z := s.state
	z = (z ^ (z >> 30)) * splitMixMulA
	z = (z ^ (z >> 27)) * splitMixMulB
	return z ^ (z >> 31)
}
