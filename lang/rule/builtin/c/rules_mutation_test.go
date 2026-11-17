/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cbuiltin

import (
	"errors"
	"testing"

	"abate/lang/rule"
)

// TestBadIDRejected proves a malformed three-part identifier or a duplicate
// identifier is rejected by the registry at load time.
func TestBadIDRejected(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	if err := LoadBuiltins(registry); err != nil {
		t.Fatalf("baseline load: %v", err)
	}
	badSeq := rule.Spec{
		ID: "C-CERT-XYZ", Reference: "docs/reference/c-cert.md — CERT C / EXP",
		Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
		Pattern: "Ident", Message: "bad sequence segment",
	}
	if err := registry.Register(badSeq); !errors.Is(err, rule.ErrInvalidRuleID) {
		t.Fatalf("bad id %s: error = %v, want ErrInvalidRuleID", badSeq.ID, err)
	}
	dup := cRules()[2]
	if err := registry.Register(dup); !errors.Is(err, rule.ErrDuplicateID) {
		t.Fatalf("duplicate id %s: error = %v, want ErrDuplicateID", dup.ID, err)
	}
}

// TestEmptyReferenceRejected proves a rule without an official reference is
// rejected at load time.
func TestEmptyReferenceRejected(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	bad := rule.Spec{
		ID: "C-CERT-900", Reference: "",
		Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
		Pattern: "Ident", Message: "missing reference",
	}
	if err := registry.Register(bad); !errors.Is(err, rule.ErrEmptyReference) {
		t.Fatalf("empty reference: error = %v, want ErrEmptyReference", err)
	}
}

// TestPatternMutationChangesHits proves that changing a rule's structural
// pattern flips which synthetic fixtures are detected.
func TestPatternMutationChangesHits(t *testing.T) {
	t.Parallel()
	original := cRules()[2] // C-CERT-003: Ident~alloc
	mutated := original
	mutated.Pattern = `Ident=free`

	compileOriginal, err := rule.CompilePattern(original.Pattern)
	if err != nil {
		t.Fatalf("compile original: %v", err)
	}
	compileMutated, err := rule.CompilePattern(mutated.Pattern)
	if err != nil {
		t.Fatalf("compile mutated: %v", err)
	}
	matcherOriginal := rule.NewMatcher(compileOriginal)
	matcherMutated := rule.NewMatcher(compileMutated)

	malloc := mkNode("Ident", "malloc")
	free := mkNode("Ident", "free")

	if !walkMatch(matcherOriginal, mkUnit(malloc)) {
		t.Error("original pattern should detect malloc")
	}
	if walkMatch(matcherOriginal, mkUnit(free)) {
		t.Error("original pattern should not detect free")
	}
	if walkMatch(matcherMutated, mkUnit(malloc)) {
		t.Error("mutated pattern should not detect malloc")
	}
	if !walkMatch(matcherMutated, mkUnit(free)) {
		t.Error("mutated pattern should detect free")
	}
}

// TestCanonicalStable proves two identical specifications render the same
// canonical bytes, so the fingerprint is reproducible.
func TestCanonicalStable(t *testing.T) {
	t.Parallel()
	a := cRules()[0]
	b := cRules()[0]
	if a.Canonical() != b.Canonical() {
		t.Fatal("identical specs render different canonical bytes")
	}
}
