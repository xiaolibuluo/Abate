/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"fmt"
	"sort"
	"strings"
)

// DiffSummary aggregates the statistics of a structured comparison.
//
// KindCounts is derived from a map but is never iterated to reach a decision;
// String sorts its keys first, so two summaries built with different insertion
// orders render identically and no map traversal order leaks into the report.
type DiffSummary struct {
	// LeftFrames is the frame count of the expected archive.
	LeftFrames int

	// RightFrames is the frame count of the actual archive.
	RightFrames int

	// FramesCompared counts frame pairs visited.
	FramesCompared int

	// FieldsCompared counts every visited field and payload byte.
	FieldsCompared int

	// FieldsEqual counts agreed fields.
	FieldsEqual int

	// FieldsDifferent counts divergent fields.
	FieldsDifferent int

	// FirstDivergentFrame is the earliest divergent frame index, or -1.
	FirstDivergentFrame int

	// KindCounts tallies the expected frames by category.
	KindCounts map[FrameKind]int
}

// String renders the summary in fixed field order with sorted kind keys.
func (s DiffSummary) String() string {
	keys := make([]string, 0, len(s.KindCounts))
	for k := range s.KindCounts {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	counts := make([]string, 0, len(keys))
	for _, k := range keys {
		counts = append(counts, fmt.Sprintf("%s=%d", k, s.KindCounts[FrameKind(k)]))
	}
	return fmt.Sprintf(
		"leftFrames=%d rightFrames=%d framesCompared=%d fieldsCompared=%d "+
			"fieldsEqual=%d fieldsDifferent=%d firstDivergentFrame=%d kindCounts=[%s]",
		s.LeftFrames, s.RightFrames, s.FramesCompared, s.FieldsCompared,
		s.FieldsEqual, s.FieldsDifferent, s.FirstDivergentFrame,
		strings.Join(counts, " "))
}
