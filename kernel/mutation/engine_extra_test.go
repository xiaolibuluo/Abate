/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"errors"
	"testing"

	"abate/kernel/algebra"
	"abate/kernel/verify"
	"abate/sim/rand"
)

func TestDiscoverErrors(t *testing.T) {
	t.Parallel()
	target, suite, _ := newFixture()
	reg := NewRegistry()
	stream := func() *rand.Stream { return rand.New(11) }
	if _, err := Discover(nil, suite, reg, stream()); !errors.Is(err, ErrNilTarget) {
		t.Fatalf("nil target error = %v", err)
	}
	if _, err := Discover(target, nil, reg, stream()); !errors.Is(err, ErrNilSuite) {
		t.Fatalf("nil suite error = %v", err)
	}
	if _, err := Discover(target, suite, nil, stream()); !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("nil registry error = %v", err)
	}
	emptyTarget := &syntheticTarget{nodes: nil, tokens: map[string]string{}}
	if _, err := Discover(emptyTarget, suite, reg, stream()); !errors.Is(err, ErrNoNodes) {
		t.Fatalf("empty nodes error = %v", err)
	}
	emptySuite := &syntheticSuite{tests: nil, cover: map[string]uint64{}}
	if _, err := Discover(target, emptySuite, reg, stream()); !errors.Is(err, ErrNoTests) {
		t.Fatalf("empty tests error = %v", err)
	}
	idle := &syntheticTarget{nodes: []string{"x"}, tokens: map[string]string{"x": "identifier"}}
	if _, err := Discover(idle, suite, reg, stream()); !errors.Is(err, ErrNoMutants) {
		t.Fatalf("no mutants error = %v", err)
	}
}

func TestJudgeRejectsUnknownOperator(t *testing.T) {
	t.Parallel()
	target, _, tests := newFixture()
	bogus := algebra.NewBasic(algebra.FormatOperatorID(algebra.FamilyMutation, "bogus", "1"),
		algebra.FamilyMutation, algebra.Params{"class": "not-registered"}, nil, true)
	point := NewMutationPoint("x", "+", "-", bogus, 1, 1)
	if _, _, err := Judge(NewRegistry(), target, tests, point); !errors.Is(err, ErrUnknownOperator) {
		t.Fatalf("unregistered operator error = %v", err)
	}
	if _, _, err := Judge(nil, target, tests, point); !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("nil registry error = %v", err)
	}
}

func TestReportDefensivePaths(t *testing.T) {
	t.Parallel()
	point := NewMutationPoint("n", "+", "-", nil, 1, 1)
	if point.operatorKey() != "" || point.Fingerprint == [32]byte{} {
		t.Fatal("nil-operator point did not hash defensively")
	}
	rep := Report{
		Points:       []MutationPoint{point, point},
		Verdicts:     []verify.MutationVerdict{verify.MutantKilled},
		KillingTests: []string{},
	}
	if len(rep.Bytes()) == 0 || ReportFingerprint(rep) == [32]byte{} {
		t.Fatal("ragged report did not encode")
	}
	other := point
	other.Node = "z"
	if !point.Less(other) {
		t.Fatal("Less node branch disagrees")
	}
	mutated := point
	mutated.Mutated = "changed"
	if !point.Less(mutated) {
		t.Fatal("Less mutated branch disagrees")
	}
	selected := point
	selected.Selector = 9
	if !point.Less(selected) {
		t.Fatal("Less selector branch disagrees")
	}
	hashed := point
	hashed.Fingerprint[0] = 0xff
	if !point.Less(hashed) {
		t.Fatal("Less fingerprint branch disagrees")
	}
}

func TestConstantVariantsOverflowBounds(t *testing.T) {
	t.Parallel()
	high := constantVariants("9223372036854775807")
	low := constantVariants("-9223372036854775808")
	if len(high) != 2 || len(low) != 2 {
		t.Fatalf("overflow guards failed: high=%v low=%v", high, low)
	}
	if len(constantVariants("not-a-number")) != 0 {
		t.Fatal("non-numeric constant produced variants")
	}
}
