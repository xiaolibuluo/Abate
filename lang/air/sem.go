/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"slices"
	"strconv"
	"strings"
)

// SemInfo carries the semantic annotations of a single node.
//
// Values are strings so that type references, scope identifiers, control-flow
// edges, mutability, and purity all share one deterministic wire encoding.
// A resolved type reference stores the hexadecimal NodeID of its declaration;
// an unresolved one stores the source-level type name.
type SemInfo map[SemKey]string

// NewSemInfo returns an empty annotation set ready for additive construction.
func NewSemInfo() SemInfo {
	return make(SemInfo)
}

// Get returns the annotation stored under key and whether it was present.
func (s SemInfo) Get(key SemKey) (string, bool) {
	value, ok := s[key]
	return value, ok
}

// Keys returns the annotation keys in ascending numeric order.
func (s SemInfo) Keys() []SemKey {
	keys := make([]SemKey, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

// With returns a copy of the annotation set with key bound to value.
func (s SemInfo) With(key SemKey, value string) SemInfo {
	out := s.Clone()
	out[key] = value
	return out
}

// Clone returns an independent copy of the annotation set.
func (s SemInfo) Clone() SemInfo {
	out := make(SemInfo, len(s))
	for key, value := range s {
		out[key] = value
	}
	return out
}

// Canonical renders the annotation set in a byte-stable form.
//
// Keys are sorted before serialization so map traversal order can never leak
// into a fingerprint or an archive.
func (s SemInfo) Canonical() string {
	var b strings.Builder
	for _, key := range s.Keys() {
		_, _ = b.WriteString(strconv.FormatUint(uint64(key), 10))
		_ = b.WriteByte('=')
		_, _ = b.WriteString(s[key])
		_ = b.WriteByte('\x00')
	}
	return b.String()
}

// Len returns the number of annotations in the set.
func (s SemInfo) Len() int {
	return len(s)
}
