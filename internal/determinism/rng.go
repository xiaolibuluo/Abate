/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import (
	"crypto/sha256"
	"encoding/binary"
)

// RNG is a SplitMix64 generator with no global state and no floating point.
// Its initial state is derived from a seed, a manifest digest and a label.
type RNG struct {
	state uint64
}

// NewRNG returns a generator whose state is derived from seed, manifestSHA256
// and label through one domain-separated SHA-256 digest.
func NewRNG(seed uint64, manifestSHA256, label string) RNG {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], seed)
	payload := make([]byte, 0, len(rngDomain)+len(buf)+len(manifestSHA256)+len(label))
	payload = append(payload, rngDomain...)
	payload = append(payload, buf[:]...)
	payload = append(payload, manifestSHA256...)
	payload = append(payload, label...)
	sum := sha256.Sum256(payload)
	return RNG{state: binary.BigEndian.Uint64(sum[0:8])}
}

// Next advances the state by the golden-ratio increment and returns the next
// mixed 64-bit value.
func (r *RNG) Next() uint64 {
	r.state += splitMixGamma
	return mix(r.state)
}

// Intn returns a deterministic value in [0, n), or zero when n is not
// positive. The modulus is taken on the full 64-bit output without floats.
func (r *RNG) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	value := r.Next() % uint64(n)
	if value > uint64(maxInt) {
		return 0
	}
	return int(value)
}

// mix applies the two SplitMix64 multiply-xorshift rounds to one state word.
func mix(z uint64) uint64 {
	z = (z ^ (z >> 30)) * splitMixMulA
	z = (z ^ (z >> 27)) * splitMixMulB
	return z ^ (z >> 31)
}
