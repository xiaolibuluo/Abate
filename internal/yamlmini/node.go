/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import "strconv"

// Node is one node of the ordered parse tree. Mappings keep source order in
// Keys/Vals; sequences use Items; scalars and nulls use Value and Sk.
type Node struct {
	Kind  Kind
	Value string
	Sk    scalarKind
	Keys  []string
	Vals  []*Node
	Items []*Node
	Line  int
}

func (n *Node) keyIndex(key string) int {
	if n == nil || n.Kind != Mapping {
		return -1
	}
	for i, k := range n.Keys {
		if k == key {
			return i
		}
	}
	return -1
}

func (n *Node) lookup(key string) (*Node, bool) {
	i := n.keyIndex(key)
	if i < 0 {
		return nil, false
	}
	return n.Vals[i], true
}

func nullNode(line int) *Node { return &Node{Kind: Null, Sk: sNull, Line: line} }

func scalarNode(value string, sk scalarKind, line int) *Node {
	return &Node{Kind: Scalar, Value: value, Sk: sk, Line: line}
}

func emptySequence(line int) *Node { return &Node{Kind: Sequence, Line: line} }
func emptyMapping(line int) *Node  { return &Node{Kind: Mapping, Line: line} }

// nodeToAny converts a node into the generic shape used by map[string]any
// destinations: map[string]any, []any, string, int, uint64, float64, bool or
// nil. This matches the type set the TCK reference driver type-asserts on.
func nodeToAny(n *Node) any {
	if n == nil {
		return nil
	}
	switch n.Kind {
	case Null:
		return nil
	case Mapping:
		m := make(map[string]any, len(n.Keys))
		for i, k := range n.Keys {
			m[k] = nodeToAny(n.Vals[i])
		}
		return m
	case Sequence:
		s := make([]any, 0, len(n.Items))
		for _, item := range n.Items {
			s = append(s, nodeToAny(item))
		}
		return s
	default:
		switch n.Sk {
		case sInt:
			if v, err := strconv.ParseInt(n.Value, 10, 64); err == nil {
				return int(v)
			}
			if v, err := strconv.ParseUint(n.Value, 10, 64); err == nil {
				return v
			}
		case sFloat:
			if v, err := strconv.ParseFloat(n.Value, 64); err == nil {
				return v
			}
		case sBool:
			return n.Value == "true"
		}
		return n.Value
	}
}
