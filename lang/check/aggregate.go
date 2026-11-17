/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"slices"
	"strings"

	"abate/lang/rule"
)

// aggregateFindings deduplicates and orders a raw finding list.
//
// The deduplication key is (RuleID, Path, Line, Column). Ordering is by file
// path, then line, then column, then rule identifier. Sorting before collapse
// keeps the result independent of the caller's slice order without ever
// emitting a map in traversal order.
func aggregateFindings(findings []rule.Finding) []rule.Finding {
	ordered := slices.Clone(findings)
	slices.SortStableFunc(ordered, compareAggregate)
	out := make([]rule.Finding, 0, len(ordered))
	for _, finding := range ordered {
		if len(out) > 0 && out[len(out)-1].Key() == finding.Key() {
			continue
		}
		out = append(out, finding)
	}
	return out
}

// compareAggregate implements the canonical aggregate ordering.
func compareAggregate(a, b rule.Finding) int {
	if c := strings.Compare(a.Path, b.Path); c != 0 {
		return c
	}
	if a.Line != b.Line {
		return a.Line - b.Line
	}
	if a.Column != b.Column {
		return a.Column - b.Column
	}
	return strings.Compare(string(a.RuleID), string(b.RuleID))
}
