/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

import (
	"reflect"
	"strings"
	"testing"
)

const testSeed = 20261114

var expectedOperators = []string{
	"block", "call", "cmp", "cond", "digit", "literal", "op", "return",
}

// TestDefaultRegistryHasEightOperators proves the three P0 and five P1
// operators are registered with the identifiers of ci/e2e_plan.json.
func TestDefaultRegistryHasEightOperators(t *testing.T) {
	t.Parallel()
	registry, err := DefaultRegistry()
	if err != nil {
		t.Fatalf("default registry failed: %v", err)
	}
	if !reflect.DeepEqual(registry.IDs(), expectedOperators) {
		t.Fatalf("operator ids wrong: %v", registry.IDs())
	}
	if len(registry.All()) != len(expectedOperators) {
		t.Fatalf("All returned %d operators", len(registry.All()))
	}
	for _, id := range []string{"cmp", "op", "digit"} {
		if _, found := registry.Lookup(id); !found {
			t.Fatalf("P0 operator %s missing", id)
		}
	}
	if _, found := registry.Lookup("absent"); found {
		t.Fatal("unknown operator must not be found")
	}
}

// TestEveryOperatorHasThirtyPointsPerLanguage proves the registry declares at
// least thirty real mutation points for every operator and every language.
func TestEveryOperatorHasThirtyPointsPerLanguage(t *testing.T) {
	t.Parallel()
	registry, err := DefaultRegistry()
	if err != nil {
		t.Fatalf("default registry failed: %v", err)
	}
	for _, spec := range registry.All() {
		if err = spec.Validate(); err != nil {
			t.Fatalf("spec %s invalid: %v", spec.ID, err)
		}
		if spec.MinValidPoints < minValidPoints ||
			len(spec.Languages()) != len(languageTags) {
			t.Fatalf("spec %s has wrong budget or languages", spec.ID)
		}
		for _, tag := range languageTags {
			if len(spec.PointsFor(tag)) < minValidPoints {
				t.Fatalf("spec %s/%s has only %d points", spec.ID, tag,
					len(spec.PointsFor(tag)))
			}
		}
	}
}

// TestEveryOperatorProducesThirtyEditsPerLanguage proves the point budget is
// real: a fixture with thirty-two occurrences yields at least thirty valid,
// non-identical edits for each language.
func TestEveryOperatorProducesThirtyEditsPerLanguage(t *testing.T) {
	t.Parallel()
	registry, err := DefaultRegistry()
	if err != nil {
		t.Fatalf("default registry failed: %v", err)
	}
	for _, spec := range registry.All() {
		for _, tag := range languageTags {
			tokens := spec.Tokens(tag)
			if len(tokens) == 0 {
				t.Fatalf("spec %s/%s has no tokens", spec.ID, tag)
			}
			content := []byte(strings.Repeat(tokens[0]+" ", 32))
			ctx := OpContext{Content: content, Language: tag,
				ManifestSHA256: "fixture-digest", Seed: testSeed}
			edits, applyErr := spec.Apply(ctx)
			if applyErr != nil {
				t.Fatalf("spec %s/%s failed: %v", spec.ID, tag, applyErr)
			}
			if len(edits) < minValidPoints {
				t.Fatalf("spec %s/%s produced only %d edits", spec.ID, tag,
					len(edits))
			}
			for _, edit := range edits {
				if err = edit.Validate(len(content)); err != nil ||
					!edit.Matches(content) || edit.After == edit.Before ||
					edit.OperatorID != spec.ID {
					t.Fatalf("spec %s/%s produced a bad edit: %+v", spec.ID,
						tag, edit)
				}
			}
			again, againErr := spec.Apply(ctx)
			if againErr != nil || !reflect.DeepEqual(edits, again) {
				t.Fatalf("spec %s/%s is not deterministic", spec.ID, tag)
			}
		}
	}
}

// TestRegistryRefusesFormattingOperators proves comment, whitespace and
// formatting families are explicitly rejected by the registry.
func TestRegistryRefusesFormattingOperators(t *testing.T) {
	t.Parallel()
	points := spreadPoints(crossPoints([]string{"a", "b", "c", "d", "e", "f"},
		[]string{"a", "b", "c", "d", "e", "f"}))
	registry := NewRegistry()
	commentSpec := newSpec(OperatorSpec{ID: "comment.swap",
		Class: ClassText, Points: points})
	if err := registry.Register(commentSpec); err == nil {
		t.Fatal("comment operator was accepted")
	}
	whitespaceSpec := newSpec(OperatorSpec{ID: "format.indent",
		Class: ClassText, Points: points})
	if err := registry.Register(whitespaceSpec); err == nil {
		t.Fatal("formatting operator was accepted")
	}
	blank := newSpec(OperatorSpec{ID: "blank", Class: ClassText,
		Points: points})
	if err := registry.Register(blank); err == nil {
		t.Fatal("blank-line family operator was accepted")
	}
	good := newSpec(OperatorSpec{ID: "cmp", Class: ClassText, Points: points})
	if err := registry.Register(good); err != nil {
		t.Fatalf("valid operator rejected: %v", err)
	}
	if err := registry.Register(good); err == nil {
		t.Fatal("duplicate operator was accepted")
	}
}

// TestNoRegisteredPointIsFormatting proves every declared point is a real
// token rewrite, never a comment, whitespace or formatting-only change.
func TestNoRegisteredPointIsFormatting(t *testing.T) {
	t.Parallel()
	registry, err := DefaultRegistry()
	if err != nil {
		t.Fatalf("default registry failed: %v", err)
	}
	for _, spec := range registry.All() {
		if containsForbidden(spec.ID) {
			t.Fatalf("forbidden family registered: %s", spec.ID)
		}
		for _, tag := range languageTags {
			for _, point := range spec.PointsFor(tag) {
				if strings.TrimSpace(point.Before) == "" ||
					strings.TrimSpace(point.Before) != point.Before ||
					forbiddenToken(point.Before) || point.Before == point.After {
					t.Fatalf("formatting point %s/%s: %+v", spec.ID, tag,
						point)
				}
			}
		}
	}
}
