/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import (
	"strings"
	"testing"
)

func TestIdentityAllAccessors(t *testing.T) {
	t.Parallel()
	e := Identity("svc")
	if e.ID() != "e-identity-vv1" {
		t.Fatalf("identity ID = %s", e.ID())
	}
	if e.Family() != FamilyIdentity {
		t.Fatalf("identity family = %s", e.Family())
	}
	if e.Params()["target"] != "svc" {
		t.Fatalf("identity params = %v", e.Params())
	}
	if len(e.Targets()) != 1 || e.Targets()[0] != "svc" {
		t.Fatalf("identity targets = %v", e.Targets())
	}
	if !e.Invertible() {
		t.Fatalf("identity must be invertible")
	}
	inv, err := e.Inverse()
	if err != nil || inv.Fingerprint() != e.Fingerprint() {
		t.Fatalf("identity inverse must equal itself")
	}
	if e.Fingerprint() == ([32]byte{}) {
		t.Fatalf("identity fingerprint must be non-zero")
	}
}

func TestCompositeAccessors(t *testing.T) {
	t.Parallel()
	p1 := makeOp("m-a-v1", FamilyMutation, "a", true)
	p2 := makeOp("m-b-v1", FamilyMutation, "b", true)
	c, err := Compose(p1, p2)
	if err != nil {
		t.Fatalf("Compose err = %v", err)
	}
	if !strings.Contains(string(c.ID()), "compose") {
		t.Fatalf("composite ID = %s", c.ID())
	}
	if c.Params()["outer"] == "" || c.Params()["inner"] == "" {
		t.Fatalf("composite params incomplete: %v", c.Params())
	}
	if c.Fingerprint() == ([32]byte{}) {
		t.Fatalf("composite fingerprint must be non-zero")
	}
}

func TestConcurrentAccessors(t *testing.T) {
	t.Parallel()
	a := makeOp("m-a-v1", FamilyMutation, "a", true)
	b := makeOp("c-b-v1", FamilyChaos, "b", true)
	prod, err := Concurrent(a, b)
	if err != nil {
		t.Fatalf("Concurrent err = %v", err)
	}
	if !strings.Contains(string(prod.ID()), "concurrent") {
		t.Fatalf("concurrent ID = %s", prod.ID())
	}
	if prod.Params()["left"] == "" || prod.Params()["right"] == "" {
		t.Fatalf("concurrent params incomplete: %v", prod.Params())
	}
	if !prod.Invertible() {
		t.Fatalf("concurrent of invertible ops must be invertible")
	}
}

func TestOperatorPackageFunctions(t *testing.T) {
	t.Parallel()
	fp := ComputeFingerprint("m-x-v1", FamilyMutation, Params{"k": "v"})
	if fp == ([32]byte{}) {
		t.Fatalf("ComputeFingerprint must be non-zero")
	}
	id := FormatOperatorID(FamilyChaos, "delay", "2")
	if id != "c-delay-v2" {
		t.Fatalf("FormatOperatorID = %s, want c-delay-v2", id)
	}
	if got := (Params{"b": "2", "a": "1"}.Canonical()); got != "a=1"+separator+"b=2" {
		t.Fatalf("Canonical = %q", got)
	}
}

func TestAlgebraErrorBranches(t *testing.T) {
	t.Parallel()
	l := NewLaws()
	p := makeOp("m-a-v1", FamilyMutation, "a", true)
	// CheckAssociativity returns false when an operand is nil (Compose errors).
	if l.CheckAssociativity(nil, p, p) {
		t.Fatalf("CheckAssociativity must be false for nil operand")
	}
	// Composite inverse fails when a sub-operator is non-invertible.
	ni := makeOp("c-ni-v1", FamilyChaos, "x", false)
	comp, err := Compose(p, ni)
	if err != nil {
		t.Fatalf("compose with non-invertible leaf ok: %v", err)
	}
	if _, err := comp.Inverse(); err == nil {
		t.Fatalf("composite inverse must fail for non-invertible leaf")
	}
	// Concurrent inverse fails when a sub-operator is non-invertible.
	prod, err := Concurrent(p, ni)
	if err != nil {
		t.Fatalf("concurrent with non-invertible leaf ok: %v", err)
	}
	if _, err := prod.Inverse(); err == nil {
		t.Fatalf("concurrent inverse must fail for non-invertible leaf")
	}
}

func TestAlgebraHelperBranches(t *testing.T) {
	// depthOf: basic returns 1, composite returns 1+max(depth).
	b := makeOp("m-a-v1", FamilyMutation, "a", true)
	if depthOf(b) != 1 {
		t.Fatalf("depthOf basic = %d, want 1", depthOf(b))
	}
	c, _ := Compose(b, b)
	if depthOf(c) != 2 {
		t.Fatalf("depthOf composite = %d, want 2", depthOf(c))
	}
	// deriveFamily: identity-only, single, mixed.
	if deriveFamily([]Family{FamilyIdentity, FamilyIdentity}) != FamilyIdentity {
		t.Fatalf("deriveFamily identity-only failed")
	}
	if deriveFamily([]Family{FamilyMutation, FamilyMutation}) != FamilyMutation {
		t.Fatalf("deriveFamily single family failed")
	}
	if deriveFamily([]Family{FamilyMutation, FamilyChaos}) != FamilyMixed {
		t.Fatalf("deriveFamily mixed failed")
	}
	// targetsEqual and intersect paths.
	a1 := makeOp("m-a-v1", FamilyMutation, "a", true)
	a2 := makeOp("m-a-v1", FamilyMutation, "a", true)
	if !targetsEqual(a1, a2) {
		t.Fatalf("targetsEqual true case failed")
	}
	b1 := makeOp("m-b-v1", FamilyMutation, "b", true)
	if targetsEqual(a1, b1) {
		t.Fatalf("targetsEqual false case failed")
	}
	if !intersect(targetsOf(a1), []TargetKey{"a", "z"}) {
		t.Fatalf("intersect true case failed")
	}
	if intersect(targetsOf(a1), []TargetKey{"z"}) {
		t.Fatalf("intersect false case failed")
	}
}

// TestAlgebraBranchCoverage exercises structural branches that the property
// tests do not reach.
func TestAlgebraBranchCoverage(t *testing.T) {
	t.Parallel()
	a := makeOp("m-a-v1", FamilyMutation, "a", true)
	b := makeOp("m-b-v1", FamilyMutation, "b", true)
	z := makeOp("m-z-v1", FamilyMutation, "z", true)
	ni := makeOp("c-ni-v1", FamilyChaos, "x", false)

	prod, err := Concurrent(a, b)
	if err != nil {
		t.Fatalf("Concurrent err = %v", err)
	}
	// A concurrent node nested in a composite reaches the *concurrentOp
	// branches of flattenTargets and depthOf.
	nested, err := Compose(prod, z)
	if err != nil {
		t.Fatalf("Compose nested err = %v", err)
	}
	if len(nested.Targets()) != 3 {
		t.Fatalf("nested targets = %v, want 3", nested.Targets())
	}
	if depthOf(nested) != 3 {
		t.Fatalf("depthOf nested = %d, want 3", depthOf(nested))
	}

	// max(a, b) with a > b: the inner operand is the deeper subtree.
	deep, err := Compose(a, a)
	if err != nil {
		t.Fatalf("Compose deep err = %v", err)
	}
	innerDeeper, err := Compose(z, deep)
	if err != nil {
		t.Fatalf("Compose innerDeeper err = %v", err)
	}
	if depthOf(innerDeeper) != 3 {
		t.Fatalf("depthOf innerDeeper = %d, want 3", depthOf(innerDeeper))
	}

	// Composite inverse: the outer dual fails while the inner dual succeeds.
	outerNI, err := Compose(ni, a)
	if err != nil {
		t.Fatalf("Compose outerNI err = %v", err)
	}
	if _, err := outerNI.Inverse(); err == nil {
		t.Fatalf("composite inverse must fail on a non-invertible outer")
	}

	// Concurrent inverse: the first dual fails.
	prodNI, err := Concurrent(ni, b)
	if err != nil {
		t.Fatalf("Concurrent prodNI err = %v", err)
	}
	if _, err := prodNI.Inverse(); err == nil {
		t.Fatalf("concurrent inverse must fail on a non-invertible first operand")
	}

	// targetsEqual with different cardinalities.
	if targetsEqual(prod, a) {
		t.Fatalf("targetsEqual must be false for different cardinalities")
	}

	// Nil-operand and target-conflict error branches of Concurrent.
	if _, err := Concurrent(nil, a); err == nil {
		t.Fatalf("Concurrent(nil, a) must fail")
	}
	if _, err := Concurrent(a, a); err == nil {
		t.Fatalf("Concurrent(a, a) must fail with a target conflict")
	}

	// Law verifiers reject invalid inputs.
	l := NewLaws()
	if l.CheckIdentity(nil) {
		t.Fatalf("CheckIdentity(nil) must be false")
	}
	if l.CheckDual(ni) {
		t.Fatalf("CheckDual(non-invertible) must be false")
	}
}
