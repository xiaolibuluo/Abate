/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package algebra defines the perturbation operator algebra that unifies the
// mutation and chaos kernels.
//
// Algebraic operations:
//
//   - Compose: p1 ∘ p2, apply p2 first and then p1.
//   - Concurrent: p1 ⊗ p2, apply both operators at once.
//   - Dual: p⁻¹, undo the operator.
//   - Identity: ε, the no-op operator.
//
// Laws: composition is associative, ε is the unit, and p ∘ p⁻¹ restores the
// baseline within tolerance.
package algebra
