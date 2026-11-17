/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
)

// errInvalidBands reports an impossible byte-window configuration.
var errInvalidBands = errors.New("corpusgen: invalid size bands")

// SizeBands mirrors the sampler's equal-byte stratification exactly.
type SizeBands struct {
	MinBytes int
	MaxBytes int
	Strata   int
}

// NewSizeBands validates the frozen byte window and stratum count.
func NewSizeBands(minBytes, maxBytes, strata int) (SizeBands, error) {
	if minBytes < 1 || maxBytes < minBytes || strata < 1 {
		return SizeBands{}, fmt.Errorf("%w: [%d,%d]/%d",
			errInvalidBands, minBytes, maxBytes, strata)
	}
	return SizeBands{MinBytes: minBytes, MaxBytes: maxBytes, Strata: strata}, nil
}

// Step is the sampler's integer band width.
func (b SizeBands) Step() int {
	return (b.MaxBytes - b.MinBytes) / b.Strata
}

// Low is the inclusive lower byte bound of one stratum.
func (b SizeBands) Low(stratum int) int {
	return b.MinBytes + stratum*b.Step()
}

// High is the inclusive upper byte bound of one stratum. The final band is
// clamped to the plan's max, exactly as the sampler clamps its bucket index.
func (b SizeBands) High(stratum int) int {
	if stratum == b.Strata-1 {
		return b.MaxBytes
	}
	return b.Low(stratum+1) - 1
}

// StratumOf returns the sampler bucket of a byte count.
func (b SizeBands) StratumOf(size int) int {
	idx := (size - b.MinBytes) / b.Step()
	if idx < 0 {
		return 0
	}
	if idx >= b.Strata {
		return b.Strata - 1
	}
	return idx
}

// Quota is the file quota of one stratum. The first band holds the short
// programs; every later band receives the design's fixed 40-file quota.
const bigBandQuota = 40

// firstFileLabel is the lexicographically smallest name of every stratum.
const firstFileLabel = "00000"

// languageQuotas splits the frozen total exactly as the design prescribes:
// the first six languages receive one extra file, the last absorbs the rest.
func languageQuotas(count int, languages []string) map[string]int {
	out := make(map[string]int, len(languages))
	base := count / len(languages)
	for i, language := range languages {
		if i < len(languages)-1 {
			out[language] = base + 1
		} else {
			out[language] = count - (len(languages)-1)*(base+1)
		}
	}
	return out
}

// Quota returns the file quota of one stratum. The first band holds the short
// programs; every later band receives the design's fixed 40-file quota.
func (b SizeBands) Quota(languageTotal, stratum int) int {
	if stratum == 0 {
		return languageTotal - (b.Strata-1)*bigBandQuota
	}
	return bigBandQuota
}

// Target is the convergence target for one file. It sits safely inside the
// band so that the final semantic block cannot overshoot the upper bound.
func (b SizeBands) Target(stratum, index int) int {
	if stratum == 0 {
		return b.MinBytes + 128 + (index*37)%1800
	}
	return b.Low(stratum) + 256 + (index*13)%512
}

// Names returns deterministic relative paths for one language and stratum.
//
// The sampler orders each stratum by sha256(RelPath) and always draws index 0.
// For stratum 0 we therefore skip every candidate name whose digest is
// lexicographically smaller than the digest of "00000", making the clean
// first-unit file both hash-smallest (drawn) and path-smallest (measured).
func (SizeBands) Names(language, ext string, stratum, count int) []string {
	prefix := fmt.Sprintf("%s/s%d/%s_s%d_", language, stratum, language, stratum)
	first := prefix + firstFileLabel + "." + ext
	firstHash := sha256.Sum256([]byte(first))
	out := make([]string, 0, count)
	for i := 0; len(out) < count; i++ {
		name := prefix + fmt.Sprintf("%05d", i) + "." + ext
		if i > 0 && stratum == 0 {
			candidate := sha256.Sum256([]byte(name))
			if bytes.Compare(candidate[:], firstHash[:]) < 0 {
				continue
			}
		}
		out = append(out, name)
	}
	return out
}
