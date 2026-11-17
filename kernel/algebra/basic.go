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

import (
	"crypto/sha256"
	"fmt"
	"slices"
)

// Basic is a concrete perturbation operator assembled from explicit fields.
//
// It is the building block used by the mutation and chaos kernels to register
// operators; composite and concurrent operators are assembled on top of Basics.
type Basic struct {
	id         OperatorID
	family     Family
	params     Params
	targets    []TargetKey
	invertible bool
	inverted   bool
}

// NewBasic constructs a Basic operator from explicit algebra fields.
//
// The params and targets slices are copied so callers may reuse them freely.
func NewBasic(id OperatorID, family Family, params Params, targets []TargetKey, invertible bool) *Basic {
	cp := make(Params, len(params))
	for k, v := range params {
		cp[k] = v
	}
	tg := make([]TargetKey, len(targets))
	copy(tg, targets)
	return &Basic{id: id, family: family, params: cp, targets: tg, invertible: invertible}
}

// ID returns the operator identifier.
func (b *Basic) ID() OperatorID {
	if b.inverted {
		return b.id + "~inv"
	}
	return b.id
}

// Family reports the operator family.
func (b *Basic) Family() Family { return b.family }

// Params returns a fresh copy of the operator parameters.
func (b *Basic) Params() Params {
	cp := make(Params, len(b.params))
	for k, v := range b.params {
		cp[k] = v
	}
	return cp
}

// Targets returns a fresh copy of the operator target keys.
func (b *Basic) Targets() []TargetKey {
	out := make([]TargetKey, len(b.targets))
	copy(out, b.targets)
	return out
}

// Invertible reports whether the operator declares a dual.
func (b *Basic) Invertible() bool {
	return b.invertible
}

// Inverse returns the dual operator, or ErrNotInvertible when declared so.
//
// The dual is an involution: Inverse(Inverse(b)) is fingerprint-equal to b.
func (b *Basic) Inverse() (Operator, error) {
	if !b.invertible {
		return nil, ErrNotInvertible
	}
	return &Basic{
		id:         b.id,
		family:     b.family,
		params:     b.Params(),
		targets:    b.Targets(),
		invertible: true,
		inverted:   !b.inverted,
	}, nil
}

// Fingerprint returns the deterministic fingerprint of the operator.
func (b *Basic) Fingerprint() [32]byte {
	eff := string(b.id)
	if b.inverted {
		eff += "~inv"
	}
	canon := eff + separator + string(b.family) + separator + b.params.Canonical()
	tks := make([]TargetKey, len(b.targets))
	copy(tks, b.targets)
	slices.Sort(tks)
	for _, tk := range tks {
		canon += separator + string(tk)
	}
	canon += fmt.Sprintf("%s%t", separator, b.inverted)
	return sha256.Sum256([]byte(canon))
}
