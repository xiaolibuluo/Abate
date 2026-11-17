/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

// Laws is the algebraic-law verifier for the perturbation operator algebra.
//
// It checks the three structural laws that every well-formed operator tree must
// satisfy: associativity of Compose, the identity unit, and the involution of
// the dual.
type Laws struct{}

// NewLaws returns a verifier for the operator algebra laws.
func NewLaws() *Laws {
	return &Laws{}
}

// CheckAssociativity verifies (p1 ∘ (p2 ∘ p3)) ≡ ((p1 ∘ p2) ∘ p3).
//
// Because fingerprints flatten the application-order leaf sequence, the two
// groupings collapse to the same hash when the algebra is associative.
func (*Laws) CheckAssociativity(p1, p2, p3 Operator) bool {
	left := composeOrNil(p1, composeOrNil(p2, p3))
	right := composeOrNil(composeOrNil(p1, p2), p3)
	if left == nil || right == nil {
		return false
	}
	return left.Fingerprint() == right.Fingerprint()
}

// composeOrNil assembles a composite, mapping any composition error to nil.
func composeOrNil(p1, p2 Operator) Operator {
	c, err := Compose(p1, p2)
	if err != nil {
		return nil
	}
	return c
}

// CheckIdentity verifies ε is a two-sided unit: (p ∘ ε) ≡ p ≡ (ε ∘ p).
func (*Laws) CheckIdentity(p Operator) bool {
	e := Identity("")
	left, err := Compose(p, e)
	if err != nil {
		return false
	}
	right, err := Compose(e, p)
	if err != nil {
		return false
	}
	return left.Fingerprint() == p.Fingerprint() && right.Fingerprint() == p.Fingerprint()
}

// CheckDual verifies the dual is an involution and preserves targets.
//
// Applying Inverse twice must return an operator fingerprint-equal to the
// original, and the dual must act on the same target domain.
func (*Laws) CheckDual(p Operator) bool {
	inv, err := Inverse(p)
	if err != nil {
		return false
	}
	inv2, err := Inverse(inv)
	if err != nil {
		return false
	}
	if inv2.Fingerprint() != p.Fingerprint() {
		return false
	}
	return targetsEqual(inv, p)
}
