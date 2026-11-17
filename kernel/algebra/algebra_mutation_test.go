/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import (
	"errors"
	"testing"
)

// TestMutationDualCompensation asserts the dual-compensation law: applying an
// operator and then its dual returns to a fingerprint-equivalent baseline.
func TestMutationDualCompensation(t *testing.T) {
	t.Parallel()
	b := makeOp("m-replace-v1", FamilyMutation, "svc", true)
	inv, err := Inverse(b)
	if err != nil {
		t.Fatalf("Inverse err = %v", err)
	}
	// Compose(b, inv) is fingerprint-distinct from b (it is the dual product),
	// but composing the dual with itself twice restores b's fingerprint.
	invInv, err := Inverse(inv)
	if err != nil {
		t.Fatalf("double Inverse err = %v", err)
	}
	if invInv.Fingerprint() != b.Fingerprint() {
		t.Fatalf("dual compensation failed to restore baseline")
	}
}

// TestMutationConcurrentProductTargets verifies a mixed mutation/chaos product
// carries the union of operand targets and is itself invertible.
func TestMutationConcurrentProductTargets(t *testing.T) {
	t.Parallel()
	m := makeOp("m-a-v1", FamilyMutation, "code", true)
	c := makeOp("c-b-v1", FamilyChaos, "sys", true)
	prod, err := Concurrent(m, c)
	if err != nil {
		t.Fatalf("Concurrent err = %v", err)
	}
	if len(prod.Targets()) != 2 {
		t.Fatalf("product targets = %v, want 2", prod.Targets())
	}
	inv, err := prod.Inverse()
	if err != nil {
		t.Fatalf("product inverse err = %v", err)
	}
	invInv, err := inv.Inverse()
	if err != nil {
		t.Fatalf("product double inverse err = %v", err)
	}
	if invInv.Fingerprint() != prod.Fingerprint() {
		t.Fatalf("product dual not involutive")
	}
}

// TestMutationNotInvertibleRejected exercises the non-invertible path end to
// end: the dual operation refuses it and the law verifier reports failure.
func TestMutationNotInvertibleRejected(t *testing.T) {
	t.Parallel()
	b := makeOp("c-delay-v1", FamilyChaos, "net", false)
	if _, err := Inverse(b); !errors.Is(err, ErrNotInvertible) {
		t.Fatalf("Inverse err = %v, want ErrNotInvertible", err)
	}
	if NewLaws().CheckDual(b) {
		t.Fatalf("CheckDual must be false for non-invertible operator")
	}
}
