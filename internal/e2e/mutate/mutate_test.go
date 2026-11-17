/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"abate/internal/e2e/metrics"
)

const equivalentFixturePath = "../../../ci/equivalent_mutations.json"

// TestApplyRewritesInMemory proves Apply produces a fresh buffer with both
// replacements and never mutates the input.
func TestApplyRewritesInMemory(t *testing.T) {
	t.Parallel()
	content := []byte("a == b && c != d")
	first := bytes.Index(content, []byte("=="))
	second := bytes.Index(content, []byte("!="))
	edits := []Edit{
		{OperatorID: "cmp", Before: "==", After: "!=", Offset: first, EndOffset: first + 2},
		{OperatorID: "cmp", Before: "!=", After: "==", Offset: second, EndOffset: second + 2},
	}
	result, err := Apply(edits, content)
	if err != nil {
		t.Fatalf("apply failed: %v", err)
	}
	if string(result) != "a != b && c == d" {
		t.Fatalf("apply produced %q", result)
	}
	if string(content) != "a == b && c != d" {
		t.Fatal("apply mutated the input buffer")
	}
	if _, err = Apply(nil, content); err != nil {
		t.Fatalf("empty edit set failed: %v", err)
	}
}

// TestApplyRejectsBadEdits proves every error branch of the in-memory rewriter.
func TestApplyRejectsBadEdits(t *testing.T) {
	t.Parallel()
	content := []byte("a == b")
	bad := [][]Edit{
		{{OperatorID: "cmp", Before: "", After: "!=", Offset: 2, EndOffset: 4}},
		{{OperatorID: "cmp", Before: "==", After: "==", Offset: 2, EndOffset: 4}},
		{{OperatorID: "cmp", Before: "zz", After: "!=", Offset: 2, EndOffset: 4}},
		{{OperatorID: "cmp", Before: "==", After: "!=", Offset: 2, EndOffset: 99}},
		{{OperatorID: "cmp", Before: "==", After: "!=", Offset: -1, EndOffset: 4}},
	}
	for i, edits := range bad {
		if _, err := Apply(edits, content); err == nil {
			t.Fatalf("bad edit set %d accepted", i)
		}
	}
	overlapping := []byte("aaaa")
	overlap := []Edit{
		{OperatorID: "x", Before: "aa", After: "b", Offset: 0, EndOffset: 2},
		{OperatorID: "x", Before: "aa", After: "c", Offset: 1, EndOffset: 3},
	}
	if _, err := Apply(overlap, overlapping); err == nil {
		t.Fatal("overlapping edits were accepted")
	}
	stale := Edit{OperatorID: "cmp", Before: "zz", After: "!=", Offset: 2, EndOffset: 4}
	if stale.Matches(content) {
		t.Fatal("stale edit matched the source buffer")
	}
}

// TestEquivalentSetLoadsExplicitDefinitions proves the on-disk set is the
// single source of equivalent mutations and that classification is explicit.
func TestEquivalentSetLoadsExplicitDefinitions(t *testing.T) {
	t.Parallel()
	set, err := LoadEquivalentSet(equivalentFixturePath)
	if err != nil {
		t.Fatalf("load equivalent set failed: %v", err)
	}
	defs := set.All()
	if len(defs) != 4 {
		t.Fatalf("expected four explicit definitions, got %d", len(defs))
	}
	for _, def := range defs {
		if err = def.Validate(); err != nil {
			t.Fatalf("definition %s invalid: %v", def.ProofID, err)
		}
	}
	proof, found := set.Classify(metrics.DrilldownRecord{
		EquivalentProofID: "EQ-LITERAL-RADIX-001",
	})
	if !found || proof != "EQ-LITERAL-RADIX-001" {
		t.Fatalf("proof classification failed: %q %t", proof, found)
	}
	proof, found = set.Classify(metrics.DrilldownRecord{
		OperatorID: "expr.commutative.swap", Equivalent: true,
	})
	if !found || proof != "EQ-COMMUTATIVE-001" {
		t.Fatalf("operator classification failed: %q %t", proof, found)
	}
	if _, found = set.Classify(metrics.DrilldownRecord{OperatorID: "cmp"}); found {
		t.Fatal("an undeclared mutation was classified as equivalent")
	}
}

// TestEquivalentSetRejectsBadDeclarations proves invalid definitions and files
// are refused instead of silently dropped.
func TestEquivalentSetRejectsBadDeclarations(t *testing.T) {
	t.Parallel()
	complete := EquivalentDef{OperatorID: "op", ContextPattern: "ctx",
		ProofID: "EQ-1", Rationale: "reason"}
	if _, err := NewEquivalentSet([]EquivalentDef{complete, complete}); err == nil {
		t.Fatal("duplicate proof was accepted")
	}
	incomplete := complete
	incomplete.Rationale = ""
	if _, err := NewEquivalentSet([]EquivalentDef{incomplete}); err == nil {
		t.Fatal("incomplete definition was accepted")
	}
	if _, err := LoadEquivalentSet("../../../ci/absent.json"); err == nil {
		t.Fatal("missing file was accepted")
	}
	dir := t.TempDir()
	badPath := filepath.Join(dir, "bad.json")
	if err := os.WriteFile(badPath, []byte("{"), 0o600); err != nil {
		t.Fatalf("fixture write failed: %v", err)
	}
	if _, err := LoadEquivalentSet(badPath); err == nil {
		t.Fatal("malformed JSON was accepted")
	}
	emptyPath := filepath.Join(dir, "empty.json")
	if err := os.WriteFile(emptyPath,
		[]byte(`{"schemaVersion":1,"definitions":[]}`), 0o600); err != nil {
		t.Fatalf("fixture write failed: %v", err)
	}
	if _, err := LoadEquivalentSet(emptyPath); err == nil {
		t.Fatal("empty definition set was accepted")
	}
}

// TestClassContextAndEditValidation proves the small value contracts reject
// incomplete configurations and stale edits.
func TestClassContextAndEditValidation(t *testing.T) {
	t.Parallel()
	if !ClassText.Valid() || parseClassOK(ClassText.Column()) != ClassText {
		t.Fatal("class round trip failed")
	}
	if _, ok := ParseClass("unknown"); ok {
		t.Fatal("unknown class was accepted")
	}
	goodCtx := OpContext{Content: []byte("a"), Language: "go"}
	if err := goodCtx.Validate(); err != nil {
		t.Fatalf("valid context rejected: %v", err)
	}
	badCtx := []OpContext{
		{Language: "go"},
		{Content: []byte("a")},
		{Content: []byte("a"), Language: "go", Offset: -1},
		{Content: []byte("a"), Language: "go", Offset: 2},
	}
	for i, ctx := range badCtx {
		if err := ctx.Validate(); err == nil {
			t.Fatalf("invalid context %d accepted", i)
		}
	}
	goodEdit := Edit{OperatorID: "cmp", Before: "==", After: "!=", Offset: 2, EndOffset: 4}
	if err := goodEdit.Validate(6); err != nil || !goodEdit.Matches([]byte("a == b")) {
		t.Fatalf("valid edit rejected: %v", err)
	}
	badEdits := []Edit{
		{OperatorID: "", Before: "==", After: "!=", Offset: 2, EndOffset: 4},
		{OperatorID: "cmp", Before: "==", After: "==", Offset: 2, EndOffset: 4},
		{OperatorID: "cmp", Before: "==", After: "!=", Offset: -1, EndOffset: 4},
		{OperatorID: "cmp", Before: "==", After: "!=", Offset: 4, EndOffset: 4},
	}
	for i, edit := range badEdits {
		if err := edit.Validate(6); err == nil {
			t.Fatalf("invalid edit %d accepted", i)
		}
	}
	for _, edit := range badEdits[2:] {
		if edit.Matches([]byte("a == b")) {
			t.Fatalf("malformed edit matched: %+v", edit)
		}
	}
}

// TestOperatorSpecValidateRejectsBadSpecs proves the operator contract rejects
// formatting families, missing languages and short point budgets.
func TestOperatorSpecValidateRejectsBadSpecs(t *testing.T) {
	t.Parallel()
	points := spreadPoints(crossPoints([]string{"a", "b", "c", "d", "e", "f"},
		[]string{"a", "b", "c", "d", "e", "f"}))
	valid := newSpec(OperatorSpec{ID: "cmp", Class: ClassText, Points: points})
	if err := valid.Validate(); err != nil {
		t.Fatalf("valid spec rejected: %v", err)
	}
	short := newSpec(OperatorSpec{ID: "cmp", Class: ClassText, Points: points})
	short.MinValidPoints = 1
	short.Points = map[string][]MutationPoint{"go": points["go"][:2]}
	if err := short.Validate(); err == nil {
		t.Fatal("short point budget was accepted")
	}
	noApply := OperatorSpec{ID: "cmp", Class: ClassText,
		MinValidPoints: minValidPoints, Points: points}
	if err := noApply.Validate(); err == nil {
		t.Fatal("spec without an Apply entry was accepted")
	}
	badClass := newSpec(OperatorSpec{ID: "cmp", Class: "other", Points: points})
	if err := badClass.Validate(); err == nil {
		t.Fatal("spec with an unknown class was accepted")
	}
	blank := map[string][]MutationPoint{}
	for tag, list := range points {
		blank[tag] = append(slicesClonePoints(list),
			MutationPoint{Before: " ", After: "="})
	}
	blankSpec := newSpec(OperatorSpec{ID: "cmp", Class: ClassText, Points: blank})
	if err := blankSpec.Validate(); err == nil {
		t.Fatal("whitespace point was accepted")
	}
}

// TestScanEditsRejectsInvalidContext proves the scanner keeps its deterministic
// error surface for an empty buffer and an unregistered language.
func TestScanEditsRejectsInvalidContext(t *testing.T) {
	t.Parallel()
	registry, err := DefaultRegistry()
	if err != nil {
		t.Fatalf("default registry failed: %v", err)
	}
	spec, _ := registry.Lookup("cmp")
	if _, err = spec.Apply(OpContext{Language: "go"}); err == nil {
		t.Fatal("empty buffer was accepted")
	}
	if _, err = spec.Apply(OpContext{Content: []byte("a == b"),
		Language: "zz"}); err == nil {
		t.Fatal("unregistered language was accepted")
	}
	if !isWordByte('_') || !isWordByte('Z') || !isWordByte('7') ||
		isWordByte(' ') {
		t.Fatal("word-byte classification failed")
	}
}

// TestZeroValueRegistryAcceptsRegistration proves a Registry declared without
// NewRegistry still initialises its map on first registration.
func TestZeroValueRegistryAcceptsRegistration(t *testing.T) {
	t.Parallel()
	var registry Registry
	points := spreadPoints(crossPoints([]string{"a", "b", "c", "d", "e", "f"},
		[]string{"a", "b", "c", "d", "e", "f"}))
	if err := registry.Register(newSpec(OperatorSpec{ID: "cmp",
		Class: ClassText, Points: points})); err != nil {
		t.Fatalf("zero-value registry failed: %v", err)
	}
	if len(registry.IDs()) != 1 {
		t.Fatalf("zero-value registry has %d operators", len(registry.IDs()))
	}
}

// parseClassOK resolves a key and reports the zero class for an unknown key.
func parseClassOK(key string) MutationClass {
	class, _ := ParseClass(key)
	return class
}

// slicesClonePoints copies one point list without sharing backing storage.
func slicesClonePoints(in []MutationPoint) []MutationPoint {
	out := make([]MutationPoint, len(in))
	copy(out, in)
	return out
}
