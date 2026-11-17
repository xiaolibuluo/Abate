/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Tolerance type identifiers for steady-state comparisons.
const (
	ToleranceAbsolute = "absolute"
	ToleranceRelative = "relative"
	TolerancePercent  = "percent"
)

// Tolerance bounds the admissible deviation from an Expectation.
//
// Absolute is a fixed delta, relative is a fraction of the expected value and
// percent is a percentage of the expected value. Value is non-negative; the
// validator rejects negative tolerances.
type Tolerance struct {
	Type  string  `json:"type" yaml:"type"`
	Value float64 `json:"value" yaml:"value"`
}
