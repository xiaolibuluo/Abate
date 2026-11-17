/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

// Inverse returns the dual (rollback) operator of op.
//
// It is the p⁻¹ algebraic operation. Non-invertible operators return
// ErrNotInvertible, which the mutation and chaos kernels must handle by
// declaring an explicit compensation strategy in the CES document.
func Inverse(op Operator) (Operator, error) {
	return op.Inverse()
}
