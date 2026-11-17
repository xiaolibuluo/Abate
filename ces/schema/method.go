/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Method declares the perturbation algebra applied to the target.
//
// Expression is written in the CES operator calculus: ∘ for composition, ⊗
// for the concurrent product, a ⁻¹ suffix for the dual, ε for identity and
// parentheses for grouping. Operands are operator:// references resolved
// against Operators, or inline operator definitions.
type Method struct {
	Expression string        `json:"expression" yaml:"expression"`
	Operators  []OperatorDef `json:"operators,omitempty" yaml:"operators,omitempty"`
}
