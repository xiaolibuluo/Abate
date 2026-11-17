/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"crypto/sha256"
	"errors"
	"strings"
	"testing"

	"abate/kernel/algebra"
	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/air"
	"abate/lang/rule"
)

// nativeOnlyTarget exposes one arithmetic node and never a derived token.
type nativeOnlyTarget struct{}

// Fingerprint returns a fixed digest for the native-only target.
func (nativeOnlyTarget) Fingerprint() [32]byte {
	return sha256.Sum256([]byte("native-only"))
}

// Nodes returns the single arithmetic node.
func (nativeOnlyTarget) Nodes() []string { return []string{"n"} }

// Token returns an arithmetic token that no derived operator indexes.
func (nativeOnlyTarget) Token(_ string) string { return "+" }

// Equivalent never reports an equivalent mutant.
func (nativeOnlyTarget) Equivalent(_ mutation.MutationPoint) bool { return false }

// Evaluate always holds so the static model yields a survived native mutant.
func (nativeOnlyTarget) Evaluate(_ mutation.MutationPoint, _ string) verify.Outcome {
	return verify.OutcomeHeld
}

// localSuite is a minimal deterministic suite for the error-path tests.
type localSuite struct{}

// Fingerprint returns a fixed digest for the local suite.
func (localSuite) Fingerprint() [32]byte { return sha256.Sum256([]byte("local")) }

// Tests returns the fixed static assertions.
func (localSuite) Tests() []string { return []string{"t1", "t2"} }

// Coverage reports two covering tests for every node.
func (localSuite) Coverage(_ string) uint64 { return 2 }

// validSpec builds one fully valid derived specification.
func validSpec() MutationSpec {
	return MutationSpec{
		SourceRuleID: "GO-EG-001",
		OperatorID:   algebra.FormatOperatorID(algebra.FamilyMutation, "derived-go-eg-001", "1"),
		TargetKind:   air.KindBinary.String(),
		Transform:    TransformBoundary,
	}
}

// TestFieldMutationDetection proves every specification field is observed.
func TestFieldMutationDetection(t *testing.T) {
	t.Parallel()
	base := validSpec()
	mutants := []struct {
		name   string
		change func(*MutationSpec)
	}{
		{"source-rule-id", func(s *MutationSpec) { s.SourceRuleID = "GO-EG-002" }},
		{"operator-id", func(s *MutationSpec) { s.OperatorID = "m-derived-other-v1" }},
		{"target-kind", func(s *MutationSpec) { s.TargetKind = air.KindCall.String() }},
		{"transform", func(s *MutationSpec) { s.Transform = TransformCall }},
	}
	for _, m := range mutants {
		t.Run(m.name, func(t *testing.T) {
			changed := base
			m.change(&changed)
			if changed.Canonical() == base.Canonical() {
				t.Fatal("Canonical did not detect the field mutation")
			}
			if err := Validate([]MutationSpec{changed}); err != nil {
				t.Fatalf("mutated spec should still be usable: %v", err)
			}
		})
	}
}

// TestUnknownOperatorAndIllegalClass locks the sentinel rejections.
func TestUnknownOperatorAndIllegalClass(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name string
		id   string
		want error
	}{
		{"foreign-prefix", "m-native-x-v1", ErrUnknownOperator},
		{"empty-id", "", ErrUnknownOperator},
	} {
		t.Run(c.name, func(t *testing.T) {
			spec := validSpec()
			spec.OperatorID = algebra.OperatorID(c.id)
			if _, err := Operators([]MutationSpec{spec}); !errors.Is(err, c.want) {
				t.Fatalf("error = %v, want %v", err, c.want)
			}
		})
	}
	badTransform := validSpec()
	badTransform.Transform = Transform("bogus")
	if _, err := Operators([]MutationSpec{badTransform}); !errors.Is(err, ErrUnknownTransform) {
		t.Fatalf("unknown transform error = %v, want ErrUnknownTransform", err)
	}
	duplicate := validSpec()
	if _, err := Operators([]MutationSpec{duplicate, duplicate}); !errors.Is(err, ErrInvalidSpec) {
		t.Fatalf("duplicate error = %v, want ErrInvalidSpec", err)
	}
	illegal := algebra.NewBasic(
		algebra.FormatOperatorID(algebra.FamilyMutation, "derived-illegal", "1"),
		algebra.FamilyMutation,
		algebra.Params{
			"class": "not-a-class", "from": "<", "to": "<=", "derived-from": "x",
		},
		[]algebra.TargetKey{"<"}, true)
	if err := checkAvailability([]algebra.Operator{illegal}); !errors.Is(err, ErrUnknownOperator) {
		t.Fatalf("illegal class error = %v, want ErrUnknownOperator", err)
	}
	if !errors.Is(checkAvailability([]algebra.Operator{illegal}), mutation.ErrInvalidDerived) {
		t.Fatal("the kernel sentinel must remain visible through errors.Is")
	}
}

// TestEmptyBoundaries covers every empty-input rejection of the pipeline.
func TestEmptyBoundaries(t *testing.T) {
	t.Parallel()
	if _, err := Operators(nil); !errors.Is(err, ErrEmptyRuleSet) {
		t.Fatalf("Operators(nil) = %v", err)
	}
	if err := Validate(nil); !errors.Is(err, ErrEmptyRuleSet) {
		t.Fatalf("Validate(nil) = %v", err)
	}
	if _, err := NewAIRSuite(nil, nil); !errors.Is(err, ErrEmptySuite) {
		t.Fatalf("NewAIRSuite(nil, nil) = %v", err)
	}
	if _, err := NewAIRTarget(nil, nil, nil); !errors.Is(err, ErrNilUnit) {
		t.Fatalf("NewAIRTarget(nil) = %v", err)
	}
	unit := sampleUnit(t)
	if _, err := NewAIRTarget(unit, nil, nil); !errors.Is(err, ErrEmptySuite) {
		t.Fatalf("NewAIRTarget(empty tests) = %v", err)
	}
	if _, _, err := Experiment(42, unit, nil); !errors.Is(err, ErrEmptyRuleSet) {
		t.Fatalf("Experiment(empty rules) = %v", err)
	}
	specs, err := NewDeriver().Derive(builtinRules(t))
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}
	if _, err := RunExperiment(42, nil, localSuite{}, specs); !errors.Is(err, mutation.ErrNilTarget) {
		t.Fatalf("RunExperiment(nil target) = %v", err)
	}
	if _, err := RunExperiment(42, nativeOnlyTarget{}, nil, specs); !errors.Is(err, mutation.ErrNilSuite) {
		t.Fatalf("RunExperiment(nil suite) = %v", err)
	}
	if _, err := RunExperiment(42, nativeOnlyTarget{}, localSuite{}, specs); !errors.Is(err, ErrNoDerivedPoints) {
		t.Fatalf("RunExperiment(native only) = %v, want ErrNoDerivedPoints", err)
	}
}

// TestTokenMappingAndAnchors exercises the AIR token projection helpers.
func TestTokenMappingAndAnchors(t *testing.T) {
	t.Parallel()
	pos := air.NewPosition(1, 1, 0, 1)
	cases := []struct {
		kind  air.NodeKind
		text  string
		token string
	}{
		{air.KindBinary, "<", "<"},
		{air.KindBinary, "==", "=="},
		{air.KindBinary, "++", ""},
		{air.KindNumLit, "1", "1"},
		{air.KindNumLit, "not-a-number", ""},
		{air.KindBoolLit, "true", "true"},
		{air.KindBranch, "break", "break"},
		{air.KindCall, "", "call"},
		{air.KindIdent, "x", ""},
	}
	for _, c := range cases {
		node := air.NewNode(c.kind, pos, c.text)
		if got := tokenOf(node); got != c.token {
			t.Fatalf("tokenOf(%s,%q) = %q, want %q", c.kind, c.text, got, c.token)
		}
	}
	if tokenOf(nil) != "" {
		t.Fatal("a nil node must not produce a token")
	}
	if !isLexicalOperator(">=") || isLexicalOperator("===") {
		t.Fatal("lexical operator classification is wrong")
	}
	anchors := anchorTokens([]MutationSpec{validSpec()})
	for index := 1; index < len(anchors); index++ {
		if anchors[index-1] > anchors[index] {
			t.Fatal("anchor tokens are not sorted")
		}
	}
	if !contains(anchors, "+") || !contains(anchors, "break") || !contains(anchors, "<") {
		t.Fatalf("anchor tokens missing representatives: %v", anchors)
	}
	if anchorKind("1") != air.KindNumLit || anchorKind("break") != air.KindBranch ||
		anchorKind("unknown") != air.KindIdent {
		t.Fatal("anchor kind mapping is wrong")
	}
	if assertionFor(air.KindBinary, "t1") != "t1" ||
		assertionFor(air.KindNumLit, "t1") != "" {
		t.Fatal("static assertion mapping changed")
	}
}

// contains reports whether a sorted token list holds one value.
func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

// TestTransformForCategoryFallback proves category classification and its
// unknown-category rejection without any language-specific identifier.
func TestTransformForCategoryFallback(t *testing.T) {
	t.Parallel()
	cases := []struct {
		category rule.Category
		want     Transform
	}{
		{rule.CategoryBoundary, TransformBoundary},
		{rule.CategoryPerformance, TransformConstant},
		{rule.CategoryErrorProne, TransformCondition},
		{rule.CategoryConcurrency, TransformCall},
		{rule.CategoryStyle, TransformType},
		{rule.CategoryNaming, TransformType},
	}
	for _, c := range cases {
		got, ok := transformFor(specFor("GO-EG-001", c.category, "rename"))
		if !ok || got != c.want {
			t.Fatalf("category %s -> %s, want %s", c.category, got, c.want)
		}
	}
	if _, ok := transformFor(specFor("GO-EG-001", rule.Category("bogus"), "")); ok {
		t.Fatal("an unknown category must not classify")
	}
	got, ok := transformFor(specFor("GO-EG-001", rule.CategoryNaming, "BoundaryShift"))
	if !ok || got != TransformBoundary {
		t.Fatalf("explicit kind did not win: %s, %t", got, ok)
	}
	if !strings.HasPrefix(string(validSpec().OperatorID), derivedPrefix) {
		t.Fatal("a derived operator id must carry the frozen prefix")
	}
}
