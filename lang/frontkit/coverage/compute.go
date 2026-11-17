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

// Compute joins the upstream construct inventory with a real parse result.
//
// The observed map holds the AIR kinds a real parse of the language produced
// together with how often each kind appeared. A construct counts as covered
// only when its mapped kind is present in that map, which keeps the coverage
// number tied to an executed frontend.
func Compute(g *grammar.Grammar, mappings []grammar.Mapping,
	observed map[air.NodeKind]int) Coverage {
	out := Coverage{}
	if g != nil {
		out.Language = g.Language
		out.Version = g.Version
		out.Source = g.Source
	}
	out.Entries = make([]Entry, 0, len(mappings))
	kinds := make(map[air.NodeKind]struct{})
	for _, mapping := range mappings {
		count := observed[mapping.Kind]
		entry := Entry{
			Node:     mapping.Node,
			Kind:     mapping.Kind,
			Status:   mapping.Status,
			Reason:   mapping.Reason,
			Rule:     mapping.Rule,
			Observed: count > 0,
			Count:    count,
		}
		out.Entries = append(out.Entries, entry)
		out.Total++
		switch mapping.Status {
		case grammar.StatusExact:
			if entry.Observed {
				out.Exact++
			}
		case grammar.StatusDegraded:
			if entry.Observed {
				out.Degraded++
			}
		case grammar.StatusOutOfScope:
			out.OutOfScope++
		case grammar.StatusUnmapped:
			out.Unmapped++
		}
		if entry.Observed {
			kinds[mapping.Kind] = struct{}{}
		}
	}
	out.Covered = out.Exact + out.Degraded
	out.ObservedKinds = len(kinds)
	return out
}

// Observe tallies the AIR kinds of one parsed unit.
//
// The tally walks the tree in pre-order, so the counts do not depend on any
// traversal order of a map.
func Observe(unit *air.Unit) map[air.NodeKind]int {
	out := make(map[air.NodeKind]int)
	for _, node := range air.Collect(unit) {
		out[node.Kind()]++
	}
	return out
}

// Merge adds the kinds of one observation into an existing tally.
func Merge(into, from map[air.NodeKind]int) map[air.NodeKind]int {
	if into == nil {
		into = make(map[air.NodeKind]int)
	}
	for kind, count := range from {
		into[kind] += count
	}
	return into
}
