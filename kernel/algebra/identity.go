/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

// identityOp is the neutral operator ε of the algebra.
//
// It is the unit for Compose: composing any operator with ε yields that
// operator unchanged, and ε is its own dual.
type identityOp struct {
	target TargetKey
}

// Identity returns the neutral operator acting on the given target.
func Identity(target TargetKey) Operator {
	return &identityOp{target: target}
}

// ID returns the identity operator identifier.
func (*identityOp) ID() OperatorID {
	return FormatOperatorID(FamilyIdentity, "identity", "v1")
}

// Family reports the identity family.
func (*identityOp) Family() Family { return FamilyIdentity }

// Params returns the identity parameters.
func (i *identityOp) Params() Params {
	return Params{"target": string(i.target)}
}

// Targets returns the identity target.
func (i *identityOp) Targets() []TargetKey {
	return []TargetKey{i.target}
}

// Invertible reports that the identity is always invertible.
func (*identityOp) Invertible() bool {
	return true
}

// Inverse returns the identity itself.
func (i *identityOp) Inverse() (Operator, error) {
	return i, nil
}

// Fingerprint returns the deterministic identity fingerprint.
func (i *identityOp) Fingerprint() [32]byte {
	return ComputeFingerprint(i.ID(), FamilyIdentity, i.Params())
}
