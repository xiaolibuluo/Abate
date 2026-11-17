/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import (
	"sort"

	"abate/internal/determinism"
)

// SamplePlan is the option struct for deterministic stratified sampling.
// Files outside the byte window remain fully registered in the manifest and
// are processed by the runner; this is the libFuzzer max_len layering rule,
// not a provenance-based exclusion.
type SamplePlan struct {
	FilesPerLanguage int
	StrataCount      int
	MinFileBytes     int64
	MaxFileBytes     int64
}

type sampleKey struct {
	entry Entry
	key   string
}

// Sample draws a deterministic, byte-window stratified, equally spaced file
// set. The ordering key is sha256(RelPath); no random generator participates in
// choosing which file is drawn, so the manifest and plan fix the set exactly.
func Sample(m Manifest, p SamplePlan) []Entry {
	p = normalizePlan(p)
	groups := groupByLanguage(m.Sorted(), p)
	out := make([]Entry, 0)
	for _, lang := range languageOrder {
		out = append(out, stratifiedPick(groups[lang], p)...)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].RelPath < out[j].RelPath
	})
	return out
}

// StratifiedSample is the name used by the tenth gate for the deterministic
// draw; it shares the single implementation above rather than duplicating it.
func StratifiedSample(m Manifest, p SamplePlan) []Entry {
	return Sample(m, p)
}

// groupByLanguage partitions the byte-window candidates by language.
func groupByLanguage(entries []Entry, p SamplePlan) map[string][]Entry {
	groups := make(map[string][]Entry, len(languageOrder))
	for _, entry := range entries {
		if entry.Bytes >= p.MinFileBytes && entry.Bytes <= p.MaxFileBytes {
			groups[entry.Language] = append(groups[entry.Language], entry)
		}
	}
	return groups
}

// normalizePlan fills unset option fields from the named package constants.
func normalizePlan(p SamplePlan) SamplePlan {
	out := p
	if out.MinFileBytes < 1 {
		out.MinFileBytes = minFileBytes
	}
	if out.MaxFileBytes < out.MinFileBytes {
		out.MaxFileBytes = maxFileBytes
	}
	if out.StrataCount < 1 {
		out.StrataCount = strataCount
	}
	if out.FilesPerLanguage < 1 {
		out.FilesPerLanguage = defaultFilesPerLanguage
	}
	return out
}

// stratifiedPick splits one language into equal byte strata and draws an equal
// number of files from each stratum by a deterministic stride. Quota wasted by
// under-filled strata (a stratum holding fewer files than its share) is
// recovered deterministically in a second pass, so the draw reaches the target
// `want` whenever the language holds at least `want` candidates. No random
// generator participates; the ordering key remains sha256(RelPath).
func stratifiedPick(entries []Entry, p SamplePlan) []Entry {
	layers := strataOf(entries, p)
	want := p.FilesPerLanguage
	if want < 1 {
		want = defaultFilesPerLanguage
	}
	if len(layers) == 0 {
		return nil
	}
	per := want / len(layers)
	remainder := want % len(layers)
	out := make([]Entry, 0, want)
	selected := make([][]bool, len(layers))
	// First pass: equal stride per stratum (original behaviour, unchanged).
	for i, layer := range layers {
		quota := per
		if i < remainder {
			quota++
		}
		sel := make([]bool, len(layer))
		if quota > len(layer) {
			quota = len(layer)
		}
		for _, idx := range determinism.Stride(len(layer), quota) {
			sel[idx] = true
			out = append(out, layer[idx].entry)
		}
		selected[i] = sel
	}
	// Second pass: recover quota wasted by under-filled strata. Scan each
	// stratum for its next undrawn entry (by the deterministic key order) and
	// append it, round-robin, until `want` is met or every stratum is empty.
	leftover := want - len(out)
	for leftover > 0 {
		progressed := false
		for i := 0; i < len(layers) && leftover > 0; i++ {
			sel := selected[i]
			for k := 0; k < len(sel); k++ {
				if !sel[k] {
					sel[k] = true
					out = append(out, layers[i][k].entry)
					leftover--
					progressed = true
					break
				}
			}
		}
		if !progressed {
			break
		}
	}
	return out
}

// strataOf buckets entries by equal byte boundaries and sorts every stratum by
// the sha256(RelPath) key so the draw is fully reproducible.
func strataOf(entries []Entry, p SamplePlan) [][]sampleKey {
	count := p.StrataCount
	if count < 1 {
		count = strataCount
	}
	layers := make([][]sampleKey, count)
	span := p.MaxFileBytes - p.MinFileBytes
	step := span / int64(count)
	if step < 1 {
		step = 1
	}
	for _, entry := range entries {
		idx := int((entry.Bytes - p.MinFileBytes) / step)
		if idx < 0 {
			idx = 0
		}
		if idx >= count {
			idx = count - 1
		}
		layers[idx] = append(layers[idx], sampleKey{
			entry: entry,
			key:   hashBytes([]byte(entry.RelPath)),
		})
	}
	for i := range layers {
		sort.Slice(layers[i], func(a, b int) bool {
			return layers[i][a].key < layers[i][b].key
		})
	}
	return layers
}
