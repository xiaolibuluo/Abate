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

func makeOp(id string, fam Family, target string, invertible bool) *Basic {
	return NewBasic(OperatorID(id), fam, Params{"k": "v"}, []TargetKey{TargetKey(target)}, invertible)
}

func TestBasicInterface(t *testing.T) {
	t.Parallel()
	b := makeOp("m-replace-v1", FamilyMutation, "svc", true)
	if b.ID() != "m-replace-v1" {
		t.Fatalf("ID = %s", b.ID())
	}
	if b.Family() != FamilyMutation {
		t.Fatalf("Family = %s", b.Family())
	}
	if len(b.Targets()) != 1 || b.Targets()[0] != "svc" {
		t.Fatalf("Targets = %v", b.Targets())
	}
	inv, err := b.Inverse()
	if err != nil {
		t.Fatalf("Inverse err = %v", err)
	}
	if inv.Family() != FamilyMutation {
		t.Fatalf("inverse family = %s", inv.Family())
	}
	// ID changes to mark the dual, fingerprint does not collide.
	if inv.ID() == b.ID() {
		t.Fatalf("inverse ID must differ from original")
	}
	if b.Fingerprint() == inv.Fingerprint() {
		t.Fatalf("inverse fingerprint must differ from original")
	}
}

func TestBasicNotInvertible(t *testing.T) {
	t.Parallel()
	b := makeOp("c-delay-v1", FamilyChaos, "net", false)
	if _, err := b.Inverse(); !errors.Is(err, ErrNotInvertible) {
		t.Fatalf("Inverse err = %v, want ErrNotInvertible", err)
	}
}

func TestComposeIdentityUnit(t *testing.T) {
	t.Parallel()
	b := makeOp("m-x-v1", FamilyMutation, "a", true)
	e := Identity("a")
	left, err := Compose(b, e)
	if err != nil {
		t.Fatalf("Compose err = %v", err)
	}
	right, err := Compose(e, b)
	if err != nil {
		t.Fatalf("Compose err = %v", err)
	}
	if left.Fingerprint() != b.Fingerprint() || right.Fingerprint() != b.Fingerprint() {
		t.Fatalf("identity unit failed")
	}
}

func TestComposeErrors(t *testing.T) {
	t.Parallel()
	b := makeOp("m-x-v1", FamilyMutation, "a", true)
	if _, err := Compose(nil, b); !errors.Is(err, ErrNilOperator) {
		t.Fatalf("nil compose err = %v", err)
	}
	// Two non-invertible operators sharing a target are incompatible.
	n1 := makeOp("c-a-v1", FamilyChaos, "shared", false)
	n2 := makeOp("c-b-v1", FamilyChaos, "shared", false)
	if _, err := Compose(n1, n2); !errors.Is(err, ErrIncompatible) {
		t.Fatalf("incompatible compose err = %v, want ErrIncompatible", err)
	}
	// Two non-invertible operators on disjoint targets are allowed.
	d1 := makeOp("c-a-v1", FamilyChaos, "x", false)
	d2 := makeOp("c-b-v1", FamilyChaos, "y", false)
	if _, err := Compose(d1, d2); err != nil {
		t.Fatalf("disjoint non-invertible compose err = %v", err)
	}
}

func TestComposeFingerprintFlatten(t *testing.T) {
	t.Parallel()
	p1 := makeOp("m-a-v1", FamilyMutation, "a", true)
	p2 := makeOp("m-b-v1", FamilyMutation, "b", true)
	c, err := Compose(p1, p2)
	if err != nil {
		t.Fatalf("Compose err = %v", err)
	}
	if c.Family() != FamilyMutation {
		t.Fatalf("composite family = %s, want m", c.Family())
	}
	// Reversing operands yields a different composite fingerprint.
	c2, _ := Compose(p2, p1)
	if c.Fingerprint() == c2.Fingerprint() {
		t.Fatalf("compose order must affect fingerprint")
	}
}

func TestConcurrentConflict(t *testing.T) {
	t.Parallel()
	a := makeOp("m-a-v1", FamilyMutation, "shared", true)
	b := makeOp("c-b-v1", FamilyChaos, "shared", true)
	if _, err := Concurrent(a, b); !errors.Is(err, ErrTargetConflict) {
		t.Fatalf("concurrent conflict err = %v, want ErrTargetConflict", err)
	}
	x := makeOp("m-a-v1", FamilyMutation, "x", true)
	y := makeOp("c-b-v1", FamilyChaos, "y", true)
	prod, err := Concurrent(x, y)
	if err != nil {
		t.Fatalf("concurrent err = %v", err)
	}
	if prod.Family() != FamilyMixed {
		t.Fatalf("concurrent family = %s, want x", prod.Family())
	}
}

func TestConcurrentNil(t *testing.T) {
	t.Parallel()
	a := makeOp("m-a-v1", FamilyMutation, "x", true)
	if _, err := Concurrent(nil, a); !errors.Is(err, ErrNilOperator) {
		t.Fatalf("nil concurrent err = %v", err)
	}
}

func TestCompositeInverse(t *testing.T) {
	t.Parallel()
	p1 := makeOp("m-a-v1", FamilyMutation, "a", true)
	p2 := makeOp("m-b-v1", FamilyMutation, "b", true)
	c, err := Compose(p1, p2)
	if err != nil {
		t.Fatalf("Compose err = %v", err)
	}
	inv, err := c.Inverse()
	if err != nil {
		t.Fatalf("composite inverse err = %v", err)
	}
	inv2, err := inv.Inverse()
	if err != nil {
		t.Fatalf("double inverse err = %v", err)
	}
	if inv2.Fingerprint() != c.Fingerprint() {
		t.Fatalf("composite inverse is not an involution")
	}
}

func TestDepthLimit(t *testing.T) {
	t.Parallel()
	// Build a chain deeper than MaxDepth.
	var chain Operator = makeOp("m-a-v1", FamilyMutation, "a", true)
	var err error
	for i := 0; i < MaxDepth+1; i++ {
		chain, err = Compose(chain, makeOp("m-x-v1", FamilyMutation, "x", true))
		if err != nil {
			break
		}
	}
	if !errors.Is(err, ErrTooDeep) {
		t.Fatalf("deep compose err = %v, want ErrTooDeep", err)
	}
}

func TestLawsIdentityAndDual(t *testing.T) {
	t.Parallel()
	l := NewLaws()
	b := makeOp("m-a-v1", FamilyMutation, "a", true)
	if !l.CheckIdentity(b) {
		t.Fatalf("CheckIdentity failed")
	}
	if !l.CheckDual(b) {
		t.Fatalf("CheckDual failed")
	}
}

func TestTargetsUnionAndEqual(t *testing.T) {
	t.Parallel()
	a := makeOp("m-a-v1", FamilyMutation, "a", true)
	b := makeOp("m-b-v1", FamilyMutation, "b", true)
	union := targetsUnion([]Operator{a, b})
	if len(union) != 2 {
		t.Fatalf("union size = %d, want 2", len(union))
	}
	if !targetsEqual(a, a) {
		t.Fatalf("self targetsEqual failed")
	}
	if targetsEqual(a, b) {
		t.Fatalf("distinct targetsEqual should be false")
	}
	if !isIdentity(Identity("z")) {
		t.Fatalf("isIdentity failed for identity")
	}
	if isIdentity(a) {
		t.Fatalf("isIdentity true for non-identity")
	}
}
