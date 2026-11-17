/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Expectation type identifiers used by CES-VALID type-matrix checks.
const (
	ExpectValue     = "value"
	ExpectRange     = "range"
	ExpectPredicate = "predicate"
)

// Expectation describes the value a steady-state probe must observe.
//
// Type selects which fields are meaningful: value uses Value, range uses Min
// and Max, and predicate uses Predicate. Values are stored as canonical
// strings so that numeric, boolean and textual observations share one model.
type Expectation struct {
	Type      string `json:"type" yaml:"type"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
	Min       string `json:"min,omitempty" yaml:"min,omitempty"`
	Max       string `json:"max,omitempty" yaml:"max,omitempty"`
	Predicate string `json:"predicate,omitempty" yaml:"predicate,omitempty"`
}
