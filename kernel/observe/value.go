/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

import (
	"encoding/json"
	"reflect"
)

// Value type identifiers for the four observation kinds.
const (
	ValueNumber    = "number"
	ValueBoolean   = "boolean"
	ValueText      = "text"
	ValueHistogram = "histogram"
)

// Value is a tagged union carrying one of the four probe value kinds.
//
// The JSON tags fix the field order, so encoding is byte-stable even when a
// value is round-tripped through JSON and YAML.
type Value struct {
	Kind    string   `json:"kind"`
	Number  float64  `json:"number"`
	Text    string   `json:"text"`
	Flag    bool     `json:"flag"`
	Buckets []uint64 `json:"buckets,omitempty"`
}

// NewNumber builds a numeric observation value.
func NewNumber(n float64) Value {
	return Value{Kind: ValueNumber, Number: n}
}

// NewBoolean builds a boolean observation value.
func NewBoolean(b bool) Value {
	return Value{Kind: ValueBoolean, Flag: b}
}

// NewText builds a textual observation value.
func NewText(s string) Value {
	return Value{Kind: ValueText, Text: s}
}

// NewHistogram builds a histogram value from a copied bucket slice.
func NewHistogram(buckets []uint64) Value {
	cp := append([]uint64(nil), buckets...)
	return Value{Kind: ValueHistogram, Buckets: cp}
}

// Encode returns the canonical JSON encoding of the value.
func (v Value) Encode() ([]byte, error) {
	return json.Marshal(v)
}

// DecodeValue parses a canonical JSON value.
func DecodeValue(data []byte) (Value, error) {
	var v Value
	if err := json.Unmarshal(data, &v); err != nil {
		return Value{}, err
	}
	return v, nil
}

// Equal reports whether two values carry identical tagged payloads.
func (v Value) Equal(other Value) bool {
	return reflect.DeepEqual(v, other)
}
