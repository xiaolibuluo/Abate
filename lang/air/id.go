/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

// NodeID is the deterministic content hash of a node and its ordered children.
//
// The hash covers the kind, the length-prefixed attribute text, the child
// count, and every child identifier in order. Source positions and semantic
// annotations are deliberately excluded, so structurally equal subtrees at
// different offsets share an identifier. This is what makes subtree reuse
// across incremental parses possible.
type NodeID uint64

// InvalidNodeID is reserved for an unset identifier and is never produced by
// ComputeNodeID.
const InvalidNodeID NodeID = 0

// ComputeNodeID derives the FNV-1a content hash of a node description.
//
// All length prefixes and integers are big-endian, and the child order is the
// caller's responsibility. The encoder always walks children in source order.
func ComputeNodeID(kind NodeKind, attribute string, children []NodeID) NodeID {
	payload := make([]byte, 0, 2+4+len(attribute)+4+8*len(children))
	payload = binary.BigEndian.AppendUint16(payload, uint16(kind))
	payload = binary.BigEndian.AppendUint32(payload, toU32(len(attribute)))
	payload = append(payload, attribute...)
	payload = binary.BigEndian.AppendUint32(payload, toU32(len(children)))
	for _, child := range children {
		payload = binary.BigEndian.AppendUint64(payload, uint64(child))
	}

	h := fnv.New64a()
	_, _ = h.Write(payload)
	sum := h.Sum64()
	if sum == 0 {
		sum = 1
	}
	return NodeID(sum)
}

// String returns the fixed-width lowercase hexadecimal form of the identifier.
func (id NodeID) String() string {
	return fmt.Sprintf("%016x", uint64(id))
}

// Valid reports whether the identifier can denote a real node.
func (id NodeID) Valid() bool {
	return id != InvalidNodeID
}
