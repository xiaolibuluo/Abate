/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import (
	"slices"
	"strconv"
	"strings"

	"abate/lang/air"
)

// Mapping is the resolved lowering of one upstream construct onto AIR.
//
// The resolution never falls back to a generic node: a construct no rule
// selected keeps air.KindInvalid and the unmapped status, so the gap stays
// enumerable instead of hiding behind a catch-all.
type Mapping struct {
	// Node is the upstream node type name.
	Node string
	// Kind is the AIR node kind the construct is lowered to.
	Kind air.NodeKind
	// Status classifies the fidelity of the lowering.
	Status string
	// Reason documents the lowering decision.
	Reason string
	// Rule is the name of the rule that selected the construct.
	Rule string
}

// Map resolves every concrete upstream construct of the grammar onto AIR.
//
// The result follows the ascending construct order of the grammar, so the
// mapping of a grammar is byte-for-byte reproducible.
func Map(g *Grammar, rules []MappingRule) []Mapping {
	if g == nil {
		return nil
	}
	out := make([]Mapping, 0, len(g.Nodes))
	for _, node := range g.Constructs() {
		out = append(out, resolve(node.Name, rules))
	}
	return out
}

// resolve selects the first rule that matches one upstream node name.
func resolve(name string, rules []MappingRule) Mapping {
	for _, rule := range rules {
		if !rule.Matches(name) {
			continue
		}
		return Mapping{
			Node:   name,
			Kind:   rule.Kind,
			Status: rule.Status,
			Reason: rule.Reason,
			Rule:   rule.Name,
		}
	}
	return Mapping{
		Node:   name,
		Kind:   air.KindInvalid,
		Status: StatusUnmapped,
		Reason: "映射表未覆盖该上游构造",
	}
}

// Unmapped returns the upstream constructs that no rule selected.
//
// The names are returned in ascending order so a report lists the gaps
// deterministically.
func Unmapped(mappings []Mapping) []string {
	out := make([]string, 0)
	for _, mapping := range mappings {
		if mapping.Status == StatusUnmapped {
			out = append(out, mapping.Node)
		}
	}
	slices.Sort(out)
	return out
}

// ByStatus returns the mappings carrying the given status, in mapping order.
func ByStatus(mappings []Mapping, status string) []Mapping {
	out := make([]Mapping, 0)
	for _, mapping := range mappings {
		if mapping.Status == status {
			out = append(out, mapping)
		}
	}
	return out
}

// Kinds returns the distinct AIR kinds used by the mappings, in ascending
// order, so a report can state the AIR surface a language relies on.
func Kinds(mappings []Mapping) []air.NodeKind {
	seen := make(map[air.NodeKind]struct{})
	out := make([]air.NodeKind, 0)
	for _, mapping := range mappings {
		if _, ok := seen[mapping.Kind]; ok {
			continue
		}
		seen[mapping.Kind] = struct{}{}
		out = append(out, mapping.Kind)
	}
	slices.Sort(out)
	return out
}

// Summary renders the mapping table as one stable line, so two runs can be
// compared without diffing the whole table.
func Summary(mappings []Mapping) string {
	counts := map[string]int{}
	for _, mapping := range mappings {
		counts[mapping.Status]++
	}
	parts := make([]string, 0, len(counts))
	for _, status := range []string{StatusExact, StatusDegraded,
		StatusOutOfScope, StatusUnmapped} {
		parts = append(parts, status+"="+strconv.Itoa(counts[status]))
	}
	return strings.Join(parts, " ")
}
