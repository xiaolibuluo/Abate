/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import "crypto/sha256"

// Operator is the algebraic contract every perturbation operator implements.
type Operator interface {
	// ID returns the globally unique operator identifier.
	ID() OperatorID

	// Family reports the operator family: mutation, chaos, identity, or mixed.
	Family() Family

	// Params returns the operator parameters.
	Params() Params

	// Targets returns the target keys the operator acts on.
	//
	// The target set drives the domain compatibility check in Compose and
	// the conflict detection in Concurrent. It must be stable across calls and
	// return a fresh slice on every invocation.
	Targets() []TargetKey

	// Invertible reports whether the operator has a dual.
	//
	// It is a structural query used by Compose to detect incompatible pairs of
	// non-invertible operators without invoking Inverse, which would recurse.
	Invertible() bool

	// Inverse returns the dual (rollback) operator.
	//
	// The dual preserves the family and the target set of the original
	// operator. Non-invertible operators must return ErrNotInvertible and
	// declare a compensation strategy in the CES document.
	Inverse() (Operator, error)

	// Fingerprint returns the deterministic fingerprint of the operator.
	//
	// The fingerprint covers the identifier, the family, and the canonical
	// parameter serialization, and is stable across runs.
	Fingerprint() [32]byte
}

// ComputeFingerprint computes the deterministic fingerprint of an operator.
//
// The input is the canonical serialization of identifier, family, and params.
func ComputeFingerprint(id OperatorID, family Family, params Params) [32]byte {
	canonical := string(id) + separator + string(family) + separator + params.Canonical()

	return sha256.Sum256([]byte(canonical))
}

// FormatOperatorID builds an operator identifier from family, name, and version.
func FormatOperatorID(family Family, name, version string) OperatorID {
	return OperatorID(string(family) + "-" + name + "-v" + version)
}
