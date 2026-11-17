/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

// Tolerance type identifiers for steady-state comparisons.
const (
	ToleranceAbsolute = "absolute"
	ToleranceRelative = "relative"
	TolerancePercent  = "percent"
)

// Tolerance bounds the admissible deviation from an expectation.
//
// Absolute is a fixed delta, relative is a fraction of the expected value and
// percent is a percentage of the expected value. A deviation exactly on the
// boundary counts as satisfied.
type Tolerance struct {
	Type  string
	Value float64
}

// margin converts the tolerance into an absolute delta around expected.
func (t Tolerance) margin(expected float64) float64 {
	switch t.Type {
	case ToleranceAbsolute:
		return t.Value
	case ToleranceRelative:
		return t.Value * abs(expected)
	case TolerancePercent:
		return t.Value / 100 * abs(expected)
	default:
		return 0
	}
}

// abs returns the magnitude of x without importing math for one operation.
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
