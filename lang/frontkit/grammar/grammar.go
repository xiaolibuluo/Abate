/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import "slices"

// Grammar is the upstream grammar fact set of one language.
//
// The struct is immutable after Load returns it, so every consumer observes
// the same node and production inventories in the same order.
type Grammar struct {
	// Language is the Abate short language tag of the grammar.
	Language string
	// Version is the upstream grammar package version.
	Version string
	// Source is the upstream repository the facts were taken from.
	Source string
	// Nodes lists every upstream node type in ascending name order.
	Nodes []NodeType
	// Rules lists every upstream production in ascending name order.
	Rules []Rule
	// Supertypes lists the abstract supertype names in ascending order.
	Supertypes []string
}

// Node returns the upstream node type with the given name.
func (g *Grammar) Node(name string) (NodeType, bool) {
	if g == nil {
		return NodeType{}, false
	}
	for _, node := range g.Nodes {
		if node.Name == name {
			return node, true
		}
	}
	return NodeType{}, false
}

// Production returns the upstream production with the given name.
func (g *Grammar) Production(name string) (Rule, bool) {
	if g == nil {
		return Rule{}, false
	}
	for _, rule := range g.Rules {
		if rule.Name == name {
			return rule, true
		}
	}
	return Rule{}, false
}

// Constructs returns the concrete named node types in ascending name order.
//
// Abstract supertypes are excluded because they group their subtypes rather
// than naming a construct of their own, and anonymous nodes are excluded
// because they carry no name in the syntax tree.
func (g *Grammar) Constructs() []NodeType {
	if g == nil {
		return nil
	}
	out := make([]NodeType, 0, len(g.Nodes))
	for _, node := range g.Nodes {
		if node.Concrete() {
			out = append(out, node)
		}
	}
	return out
}

// FieldNames returns the sorted union of every field name the grammar
// exposes, so a report can state the complete field inventory of a language.
func (g *Grammar) FieldNames() []string {
	if g == nil {
		return nil
	}
	seen := make(map[string]struct{})
	for _, node := range g.Nodes {
		for _, field := range node.Fields {
			seen[field] = struct{}{}
		}
	}
	out := make([]string, 0, len(seen))
	for field := range seen {
		out = append(out, field)
	}
	slices.Sort(out)
	return out
}
