/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import "slices"

// NodeType is one upstream node type taken from node-types.json.
//
// The upstream file is the authoritative node inventory of a grammar: it lists
// every node type, the fields each node exposes and, for abstract
// supertypes, the concrete subtypes that fill the slot.
type NodeType struct {
	// Name is the upstream node type name, for example "function_definition".
	Name string
	// Named reports whether the node type carries a name in the syntax tree.
	Named bool
	// Supertype reports whether the node type is an abstract supertype that
	// only groups its subtypes instead of naming a construct of its own.
	Supertype bool
	// Fields lists the field names the node type exposes, in ascending order.
	Fields []string
	// Subtypes lists the concrete subtypes of a supertype, in ascending order.
	Subtypes []string
}

// HasField reports whether the node type exposes the named field.
func (n NodeType) HasField(field string) bool {
	return slices.Contains(n.Fields, field)
}

// Concrete reports whether the node type names a construct of its own.
//
// A concrete node type is either a plain named node or an anonymous node such
// as a literal; a supertype is abstract and only groups its subtypes.
func (n NodeType) Concrete() bool {
	return n.Named && !n.Supertype
}
