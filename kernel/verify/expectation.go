/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

// Expectation type identifiers for steady-state comparisons.
const (
	ExpectValue     = "value"
	ExpectRange     = "range"
	ExpectPredicate = "predicate"
)

// Expectation states what an observed value must look like.
//
// Type selects the active fields: value uses Value, range uses Min and Max,
// and predicate uses Predicate with a comparison operator such as ">=5".
type Expectation struct {
	Type      string
	Value     float64
	Min       float64
	Max       float64
	Predicate string
}
