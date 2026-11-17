/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import "abate/internal/determinism"

// generatorDomain separates the corpus generator's hash stream from every
// other consumer of internal/determinism in the repository.
const generatorDomain = "abate-self-researched-corpus/v1"

// LabeledRNG is one deterministic SplitMix64 stream. Its state depends only on
// the frozen seed, the fixed generator domain and the caller's label.
type LabeledRNG struct {
	stream determinism.RNG
	label  string
}

// NewLabeledRNG derives a stream for one generation scope (for example one
// language or one stratum). Equal inputs always produce equal streams.
func NewLabeledRNG(seed uint64, label string) LabeledRNG {
	return LabeledRNG{
		stream: determinism.NewRNG(seed, generatorDomain, label),
		label:  label,
	}
}

// Intn returns a deterministic integer in [0, n).
func (r *LabeledRNG) Intn(n int) int {
	return r.stream.Intn(n)
}

// Label returns the immutable domain label of this stream.
func (r LabeledRNG) Label() string {
	return r.label
}
