/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// Coverage is the construct coverage of one language frontend.
//
// Every counter is derived from the same real parse, so the summary and the
// per construct rows can never disagree.
type Coverage struct {
	// Language is the Abate short language tag of the language.
	Language string
	// Version is the upstream grammar package version.
	Version string
	// Source is the upstream repository the grammar facts came from.
	Source string
	// Total is the number of concrete upstream constructs.
	Total int
	// Covered is the number of constructs whose AIR kind a real parse produced.
	Covered int
	// Exact is the number of exactly mapped constructs a real parse produced.
	Exact int
	// Degraded is the number of degraded constructs a real parse produced.
	Degraded int
	// OutOfScope is the number of constructs outside the AIR capability domain.
	OutOfScope int
	// Unmapped is the number of constructs no mapping rule selected.
	Unmapped int
	// ObservedKinds is the number of distinct AIR kinds a real parse produced.
	ObservedKinds int
	// Entries lists one row per upstream construct, in upstream order.
	Entries []Entry
}

// Percent returns the share of constructs a real parse covered, in percent.
//
// Constructs outside the AIR capability domain are not counted as covered,
// because Abate cannot analyse them at all.
func (c Coverage) Percent() float64 {
	if c.Total == 0 {
		return 0
	}
	return float64(c.Covered) * 100 / float64(c.Total)
}

// AcknowledgedPercent returns the share of constructs that are either covered
// or explicitly outside the AIR capability domain, in percent.
func (c Coverage) AcknowledgedPercent() float64 {
	if c.Total == 0 {
		return 0
	}
	return float64(c.Covered+c.OutOfScope) * 100 / float64(c.Total)
}

// MappingPercent returns the share of constructs the mapping table resolves,
// in percent. A complete table resolves every construct.
func (c Coverage) MappingPercent() float64 {
	if c.Total == 0 {
		return 0
	}
	return float64(c.Total-c.Unmapped) * 100 / float64(c.Total)
}

// Uncovered returns the rows a real parse did not produce, in upstream order.
func (c Coverage) Uncovered() []Entry {
	out := make([]Entry, 0)
	for _, entry := range c.Entries {
		if entry.Status == grammar.StatusOutOfScope ||
			entry.Status == grammar.StatusUnmapped {
			continue
		}
		if !entry.Observed {
			out = append(out, entry)
		}
	}
	return out
}

// ObservedKindsOf returns the AIR kinds a real parse produced, in ascending
// kind order, so a report can state the AIR surface the frontend reaches.
func (c Coverage) ObservedKindsOf() []air.NodeKind {
	seen := make(map[air.NodeKind]struct{})
	out := make([]air.NodeKind, 0)
	for _, entry := range c.Entries {
		if !entry.Observed {
			continue
		}
		if _, ok := seen[entry.Kind]; ok {
			continue
		}
		seen[entry.Kind] = struct{}{}
		out = append(out, entry.Kind)
	}
	return out
}
