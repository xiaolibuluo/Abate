/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import "fmt"

// Diff is the structured, deterministic result of comparing two archives.
//
// It follows the deep-recursive comparison model of google/go-cmp but is
// handwritten for the frame domain without reflection: every header and frame
// field is visited in a fixed order, the earliest divergence is captured with
// its complete path, and a statistical summary aggregates equality counts.
// Map-derived output is emitted in sorted key order, so constructing the same
// archive through different source orders still yields byte-identical reports.
type Diff struct {
	// Equal reports whether every compared field of both archives agreed.
	Equal bool

	// First is the earliest divergence, or nil when Equal is true.
	First *DiffPoint

	// Summary aggregates the comparison statistics.
	Summary DiffSummary
}

// String renders the comparison as deterministic, human-readable lines.
func (d Diff) String() string {
	out := fmt.Sprintf("equal: %t\n", d.Equal)
	if d.First != nil {
		out += fmt.Sprintf("firstDiff: %s\n", d.First.String())
	}
	out += fmt.Sprintf("summary: %s", d.Summary.String())
	return out
}
