/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

// Transform names one of the five rule-to-operator derivation strategies.
type Transform string

// Transform values enumerate the five derivation classes.
const (
	// TransformBoundary derives a boundary-value operator from boundary rules.
	TransformBoundary Transform = "boundary"
	// TransformConstant derives a constant-perturb operator from constant rules.
	TransformConstant Transform = "constant"
	// TransformCondition derives a negate-condition operator from condition rules.
	TransformCondition Transform = "condition"
	// TransformType derives a return-value operator from type rules.
	TransformType Transform = "type"
	// TransformCall derives a remove-call operator from call rules.
	TransformCall Transform = "call"
)

// Valid reports whether the transform is one of the five allocated values.
func (t Transform) Valid() bool {
	_, ok := transformEntryFor(t)
	return ok
}

// String returns the stable wire text of the transform.
func (t Transform) String() string {
	return string(t)
}
