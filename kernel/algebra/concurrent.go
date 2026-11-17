/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import "fmt"

// concurrentOp is the wrapper produced by Concurrent.
//
// It applies both operands at the same virtual time slice. The scheduler
// total-orders the resulting events, so the product is deterministic.
type concurrentOp struct {
	a Operator
	b Operator
}

// Concurrent assembles p1 ⊗ p2, applying both operators together.
//
// Concurrent operators must have disjoint target sets; sharing a target is
// rejected with ErrTargetConflict. The product depth is bounded by MaxDepth.
func Concurrent(p1, p2 Operator) (Operator, error) {
	if p1 == nil || p2 == nil {
		return nil, ErrNilOperator
	}
	if depthOf(p1)+depthOf(p2)+1 > MaxDepth {
		return nil, ErrTooDeep
	}
	if intersect(targetsOf(p1), targetsOf(p2)) {
		return nil, ErrTargetConflict
	}
	return &concurrentOp{a: p1, b: p2}, nil
}

// ID returns the concurrent operator identifier.
func (*concurrentOp) ID() OperatorID {
	return FormatOperatorID(FamilyMixed, "concurrent", "v1")
}

// Family derives the concurrent family from its operands.
func (o *concurrentOp) Family() Family {
	return deriveFamily(flattenFamilies(o))
}

// Params returns the canonical parameters of the concurrent product.
func (o *concurrentOp) Params() Params {
	return Params{
		"left":  fmt.Sprintf("%x", o.a.Fingerprint()),
		"right": fmt.Sprintf("%x", o.b.Fingerprint()),
	}
}

// Targets returns the union of the operand target keys.
func (o *concurrentOp) Targets() []TargetKey {
	return targetsUnion([]Operator{o.a, o.b})
}

// Invertible reports whether both operands are invertible.
func (o *concurrentOp) Invertible() bool {
	return o.a.Invertible() && o.b.Invertible()
}

// Inverse returns the concurrent product of the duals.
func (o *concurrentOp) Inverse() (Operator, error) {
	ia, err := Inverse(o.a)
	if err != nil {
		return nil, err
	}
	ib, err := Inverse(o.b)
	if err != nil {
		return nil, err
	}
	return Concurrent(ia, ib)
}

// Fingerprint returns the flattened fingerprint of the product.
func (o *concurrentOp) Fingerprint() [32]byte {
	return seqFingerprint(flattenLeaves(o))
}
