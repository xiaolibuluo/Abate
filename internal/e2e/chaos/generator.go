/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package chaos

import "abate/internal/determinism"

// Generator produces deterministic damaged variants of one clean in-memory
// source buffer. The seed and manifest digest only choose damage offsets.
type Generator struct {
	Seed              uint64
	ManifestSHA256    string
	BrokenPerLanguage int
}

// NewGenerator builds a generator from a seed, a manifest digest and the
// per-language damage budget.
func NewGenerator(seed uint64, manifestSHA256 string,
	perLanguage int) Generator {
	return Generator{
		Seed:              seed,
		ManifestSHA256:    manifestSHA256,
		BrokenPerLanguage: perLanguage,
	}
}

// Generate returns the merged, deduplicated damaged variants of content. The
// truncated body is always present and the remaining variants are byte splices
// at offsets selected by the SplitMix64 generator.
func (g Generator) Generate(content []byte, label string) [][]byte {
	if g.BrokenPerLanguage <= 0 {
		return [][]byte{}
	}
	rng := determinism.NewRNG(g.Seed, g.ManifestSHA256, label)
	candidates := make([][]byte, 0, g.BrokenPerLanguage+1)
	candidates = append(candidates, TruncateHalf(content))
	for i := 0; i < g.BrokenPerLanguage; i++ {
		at := rng.Intn(len(content) + 1)
		candidates = append(candidates, Splice(content, at, spliceChunk))
	}
	return Merge(candidates)
}
