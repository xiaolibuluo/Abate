/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "encoding/json"

// Node is the read-only view of one AIR tree vertex.
//
// Frontends build trees with NewNode and NewNodeAt; downstream passes read them
// through this interface. The concrete implementation is unexported so no
// language-specific syntax type can ever leak through the AIR public API.
type Node interface {
	// Kind returns the syntactic class of the node.
	Kind() NodeKind
	// Pos returns the source interval of the node.
	Pos() Position
	// End returns the end line and column of the node.
	End() Position
	// Children returns a defensive copy of the ordered child list.
	Children() []Node
	// Sem returns the annotation stored under key, or nil when absent.
	Sem(key SemKey) any
}

// node is the canonical, self-contained AIR vertex implementation.
type node struct {
	id       NodeID
	kind     NodeKind
	pos      Position
	end      Position
	text     string
	children []Node
	sem      SemInfo
}

// nodeJSON is the canonical JSON projection used by determinism.HashState.
type nodeJSON struct {
	ID       string            `json:"id"`
	Kind     string            `json:"kind"`
	Position string            `json:"position"`
	End      string            `json:"end"`
	Text     string            `json:"text"`
	Sem      map[string]string `json:"sem"`
	Children []json.RawMessage `json:"children"`
}

// NewNode builds a node whose end position is derived from pos.EndOffset.
func NewNode(kind NodeKind, pos Position, text string, children ...Node) Node {
	end := Position{
		Line:      pos.Line,
		Column:    pos.Column,
		Offset:    pos.EndOffset,
		EndOffset: pos.EndOffset,
	}
	return NewNodeAt(kind, pos, end, text, children...)
}

// NewNodeAt builds a node with explicit start and end positions.
func NewNodeAt(kind NodeKind, pos, end Position, text string, children ...Node) Node {
	kept := make([]Node, 0, len(children))
	ids := make([]NodeID, 0, len(children))
	for _, child := range children {
		if child == nil {
			continue
		}
		kept = append(kept, child)
		ids = append(ids, IDOf(child))
	}
	return &node{
		id:       ComputeNodeID(kind, text, ids),
		kind:     kind,
		pos:      pos,
		end:      end,
		text:     text,
		children: kept,
		sem:      NewSemInfo(),
	}
}

// Annotate attaches a semantic annotation to a node in place.
//
// It returns false when the node was not built by this package. Annotations do
// not change the NodeID, so semantic enrichment never invalidates subtree reuse.
func Annotate(target Node, key SemKey, value string) bool {
	inner, ok := target.(*node)
	if !ok {
		return false
	}
	if inner.sem == nil {
		inner.sem = NewSemInfo()
	}
	inner.sem[key] = value
	return true
}

// AttributeOf returns the node's own attribute text, or an empty string for a
// foreign Node implementation.
func AttributeOf(target Node) string {
	inner, ok := target.(*node)
	if !ok {
		return ""
	}
	return inner.text
}

// IDOf returns the content hash of a node, recomputing it for foreign
// implementations from their observable structure.
func IDOf(target Node) NodeID {
	if target == nil {
		return InvalidNodeID
	}
	if inner, ok := target.(*node); ok {
		return inner.id
	}
	children := target.Children()
	ids := make([]NodeID, 0, len(children))
	for _, child := range children {
		ids = append(ids, IDOf(child))
	}
	return ComputeNodeID(target.Kind(), AttributeOf(target), ids)
}

// SemanticsOf returns an independent snapshot of a node's annotations.
func SemanticsOf(target Node) SemInfo {
	inner, ok := target.(*node)
	if !ok {
		return NewSemInfo()
	}
	out := inner.sem.Clone()
	return out
}

// Kind returns the syntactic class of the node.
func (n *node) Kind() NodeKind { return n.kind }

// Pos returns the source interval of the node.
func (n *node) Pos() Position { return n.pos }

// End returns the end line and column of the node.
func (n *node) End() Position { return n.end }

// Children returns a defensive copy of the ordered child list.
func (n *node) Children() []Node {
	return append([]Node(nil), n.children...)
}

// Sem returns the annotation stored under key, or nil when absent.
func (n *node) Sem(key SemKey) any {
	value, ok := n.sem[key]
	if !ok {
		return nil
	}
	return value
}

// MarshalJSON projects the whole subtree into a deterministic JSON value.
//
// This makes determinism.HashState(unit) cover every node instead of seeing an
// opaque empty object for the unexported concrete type.
func (n *node) MarshalJSON() ([]byte, error) {
	view := nodeJSON{
		ID:       n.id.String(),
		Kind:     n.kind.String(),
		Position: n.pos.Canonical(),
		End:      n.end.Canonical(),
		Text:     n.text,
		Sem:      make(map[string]string, len(n.sem)),
		Children: make([]json.RawMessage, 0, len(n.children)),
	}
	for _, key := range n.sem.Keys() {
		view.Sem[key.String()] = n.sem[key]
	}
	for _, child := range n.children {
		encoded, err := json.Marshal(child)
		if err != nil {
			return nil, err
		}
		view.Children = append(view.Children, encoded)
	}
	return json.Marshal(view)
}
