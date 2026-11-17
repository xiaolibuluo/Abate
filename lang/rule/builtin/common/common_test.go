/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"errors"
	"strings"
	"testing"

	"abate/lang/rule"
)

// TestTemplateCount proves the frozen sixty-template family is intact.
func TestTemplateCount(t *testing.T) {
	t.Parallel()
	if got := Count(); got != templateCount {
		t.Fatalf("Count() = %d, want %d", got, templateCount)
	}
	if got := len(All()); got != templateCount {
		t.Fatalf("len(All()) = %d, want %d", got, templateCount)
	}
	if got := len(builtinTemplates()); got != templateCount {
		t.Fatalf("len(builtinTemplates()) = %d, want %d", got, templateCount)
	}
	if got := len(families()); got != familyCount {
		t.Fatalf("len(families()) = %d, want %d", got, familyCount)
	}
	groups := variantGroups()
	if len(groups) != familyCount {
		t.Fatalf("variant group count = %d, want %d", len(groups), familyCount)
	}
	for i, group := range groups {
		if len(group) != 5 {
			t.Fatalf("group %d has %d variants, want 5", i, len(group))
		}
	}
}

// TestAllSpecsCompile proves every generated spec passes schema, pattern,
// predicate and identifier validation, and that all ids are unique and ordered.
func TestAllSpecsCompile(t *testing.T) {
	t.Parallel()
	specs := All()
	seen := make(map[string]struct{}, len(specs))
	previous := ""
	references := 0
	for _, spec := range specs {
		id := string(spec.ID)
		if !spec.ID.Valid() {
			t.Fatalf("invalid rule id %q", id)
		}
		if spec.ID.Source() != sourceAIR {
			t.Fatalf("rule %s has source %q, want %q", id, spec.ID.Source(), sourceAIR)
		}
		if _, dup := seen[id]; dup {
			t.Fatalf("duplicate rule id %q", id)
		}
		seen[id] = struct{}{}
		if strings.Compare(previous, id) > 0 {
			t.Fatalf("All() is not id-ordered: %s after %s", id, previous)
		}
		previous = id
		if err := spec.Validate(); err != nil {
			t.Fatalf("spec %s fails Validate: %v", id, err)
		}
		if _, err := rule.CompilePattern(spec.Pattern); err != nil {
			t.Fatalf("spec %s fails CompilePattern: %v", id, err)
		}
		for _, clause := range spec.Where {
			if _, err := rule.ParsePredicate(clause); err != nil {
				t.Fatalf("spec %s has invalid clause %q: %v", id, clause, err)
			}
		}
		if _, err := rule.Compile(spec); err != nil {
			t.Fatalf("spec %s fails Compile: %v", id, err)
		}
		if !strings.Contains(spec.Reference, "docs/reference/") {
			t.Fatalf("spec %s has a non-traceable reference %q", id, spec.Reference)
		}
		references++
	}
	if references != len(specs) || len(specs) != templateCount {
		t.Fatalf("reference coverage = %d/%d", references, len(specs))
	}
}

// TestRegistryLoadsAllRules proves the four load-time checks accept the set.
func TestRegistryLoadsAllRules(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	if err := LoadBuiltins(registry); err != nil {
		t.Fatalf("LoadBuiltins: %v", err)
	}
	if registry.Len() != Count() {
		t.Fatalf("registry size = %d, want %d", registry.Len(), Count())
	}
	loaded := registry.For("com")
	if len(loaded) != Count() {
		t.Fatalf("For(com) returned %d rules, want %d", len(loaded), Count())
	}
	for i := 1; i < len(loaded); i++ {
		if loaded[i-1].Spec().ID > loaded[i].Spec().ID {
			t.Fatalf("For() is not id-ordered at %d", i)
		}
	}
}

// TestLoadBuiltinsNilRegistry proves the nil guard reports its own sentinel.
func TestLoadBuiltinsNilRegistry(t *testing.T) {
	t.Parallel()
	if err := LoadBuiltins(nil); !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("LoadBuiltins(nil) = %v, want ErrNilRegistry", err)
	}
}

// TestGenerateTagMatrix proves every short tag names its own namespace.
func TestGenerateTagMatrix(t *testing.T) {
	t.Parallel()
	tags := []string{"go", "PY", "Js", "c", "RS", "cs", "kt", "", "unknown"}
	for _, tag := range tags {
		specs := Generate(tag, nil)
		if len(specs) != templateCount {
			t.Fatalf("Generate(%q) returned %d specs, want %d", tag, len(specs), templateCount)
		}
		wantPrefix := strings.ToUpper(strings.TrimSpace(tag))
		if wantPrefix == "" {
			wantPrefix = DefaultTag
		}
		for _, spec := range specs {
			if !strings.HasPrefix(string(spec.ID), wantPrefix+"-"+sourceAIR+"-") {
				t.Fatalf("Generate(%q) produced id %s", tag, spec.ID)
			}
			if spec.Language() != strings.ToLower(wantPrefix) {
				t.Fatalf("Generate(%q) produced language %q", tag, spec.Language())
			}
		}
	}
}

// TestReferenceTableIsTraceable proves the citation table is complete and stable.
func TestReferenceTableIsTraceable(t *testing.T) {
	t.Parallel()
	table := ReferenceTable()
	if len(table) != len(languageChapters)*len(familySections) {
		t.Fatalf("reference table size = %d, want %d", len(table),
			len(languageChapters)*len(familySections))
	}
	first := ReferenceTable()
	for i, ref := range table {
		if ref.Section == "" || ref.Cite() == "" || ref.Tag == "" || ref.Family == "" {
			t.Fatalf("reference %d is empty: %+v", i, ref)
		}
		if !strings.Contains(ref.Section, "docs/reference/") {
			t.Fatalf("reference %d is not traceable: %q", i, ref.Section)
		}
		if i >= len(first) || ref != first[i] {
			t.Fatalf("reference table is not stable at %d", i)
		}
	}
	known, ok := ReferenceFor("py", "naming")
	if !ok || !strings.Contains(known.Section, "python-pep8.md") {
		t.Fatalf("ReferenceFor(py,naming) = %+v, %v", known, ok)
	}
	fallback, ok := ReferenceFor("unknown-tag", "unknown-family")
	if !ok || !strings.Contains(fallback.Section, "General") {
		t.Fatalf("fallback reference = %+v, %v", fallback, ok)
	}
}

// TestFingerprintIsStable proves one hundred rebuilds produce one digest.
func TestFingerprintIsStable(t *testing.T) {
	t.Parallel()
	base := Fingerprint()
	for i := 0; i < 100; i++ {
		if got := Fingerprint(); got != base {
			t.Fatalf("fingerprint drifted at rebuild %d", i)
		}
	}
	if base == [32]byte{} {
		t.Fatal("fingerprint of a full rule set must not be zero")
	}
}
