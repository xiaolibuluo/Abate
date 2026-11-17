/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import (
	"errors"
	"slices"
	"strings"
	"sync"
	"testing"

	"abate/lang/rule"
)

// TestRuleSetCount proves the declared Go rule set clears the M3 threshold.
func TestRuleSetCount(t *testing.T) {
	t.Parallel()
	if got := Count(); got < 100 {
		t.Fatalf("rule count = %d, want at least 100", got)
	}
}

// TestRuleDeclarations proves identifier uniqueness, validity, reference
// coverage and compilability of every declared rule.
func TestRuleDeclarations(t *testing.T) {
	t.Parallel()
	specs := All()
	seen := make(map[string]struct{}, len(specs))
	previous := ""
	for _, spec := range specs {
		id := string(spec.ID)
		if !spec.ID.Valid() {
			t.Fatalf("invalid rule id %q", id)
		}
		if _, ok := seen[id]; ok {
			t.Fatalf("duplicate rule id %q", id)
		}
		seen[id] = struct{}{}
		if strings.Compare(previous, id) > 0 {
			t.Fatalf("All() is not id-ordered: %s after %s", id, previous)
		}
		previous = id
		if strings.TrimSpace(spec.Reference) == "" {
			t.Fatalf("rule %s has no official reference", id)
		}
		if strings.TrimSpace(spec.Pattern) == "" {
			t.Fatalf("rule %s has an empty pattern", id)
		}
		if _, err := rule.Compile(spec); err != nil {
			t.Fatalf("rule %s does not compile: %v", id, err)
		}
	}
}

// TestFixtureCoverage proves every rule has one positive and one negative case.
func TestFixtureCoverage(t *testing.T) {
	t.Parallel()
	fixtures := allFixtures()
	for _, spec := range All() {
		item, ok := fixtures[string(spec.ID)]
		if !ok {
			t.Errorf("rule %s has no fixture", spec.ID)
			continue
		}
		if item.Pos == "" || item.Neg == "" {
			t.Errorf("rule %s has an empty fixture side", spec.ID)
		}
	}
}

// TestRuleFixtures proves every rule detects its positive case and stays
// silent on its negative case through the real AIR frontend and matcher.
func TestRuleFixtures(t *testing.T) {
	t.Parallel()
	byID := make(map[string]rule.Spec, 128)
	for _, spec := range All() {
		byID[string(spec.ID)] = spec
	}
	fixtures := allFixtures()
	ids := make([]string, 0, len(byID))
	for id := range byID {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	for _, id := range ids {
		item, ok := fixtures[id]
		if !ok {
			continue
		}
		t.Run(id, func(t *testing.T) {
			compiled, err := rule.Compile(byID[id])
			if err != nil {
				t.Fatalf("compile %s: %v", id, err)
			}
			positive, err := ParseUnit("positive.go", fixtureSource(item.Pos))
			if err != nil {
				t.Fatalf("parse positive %s: %v", id, err)
			}
			negative, err := ParseUnit("negative.go", fixtureSource(item.Neg))
			if err != nil {
				t.Fatalf("parse negative %s: %v", id, err)
			}
			if hits := compiled.Check(positive); len(hits) == 0 {
				t.Errorf("rule %s did not detect its positive fixture", id)
			}
			if hits := compiled.Check(negative); len(hits) != 0 {
				t.Errorf("rule %s hit its negative fixture %d times: %s",
					id, len(hits), hits[0].String())
			}
		})
	}
}

// TestRegistryLoadsAllRules proves the four load-time validations accept the
// full set and that For returns it in identifier order.
func TestRegistryLoadsAllRules(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	if err := LoadBuiltins(registry); err != nil {
		t.Fatalf("LoadBuiltins: %v", err)
	}
	if registry.Len() != Count() {
		t.Fatalf("registry size = %d, want %d", registry.Len(), Count())
	}
	loaded := registry.For("go")
	for i := 1; i < len(loaded); i++ {
		if loaded[i-1].Spec().ID > loaded[i].Spec().ID {
			t.Fatalf("For() is not id-ordered at %d", i)
		}
	}
}

// TestRuleSetChaosHash proves 100 repeated builds and concurrent read-only
// builds produce one identical rule-set digest.
func TestRuleSetChaosHash(t *testing.T) {
	t.Parallel()
	base := Fingerprint()
	for i := 0; i < 100; i++ {
		if got := Fingerprint(); got != base {
			t.Fatalf("rule-set digest changed at build %d", i)
		}
	}
	var group sync.WaitGroup
	digests := make(chan [32]byte, 32)
	for i := 0; i < 32; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			digests <- Fingerprint()
		}()
	}
	group.Wait()
	close(digests)
	for digest := range digests {
		if digest != base {
			t.Fatal("concurrent rule-set build produced a different digest")
		}
	}
}

// TestRuleSetMutationOrder proves shuffled declaration input still yields the
// same ordered rule set, and that bad rules are rejected at load time.
func TestRuleSetMutationOrder(t *testing.T) {
	t.Parallel()
	specs := All()
	shuffled := slices.Clone(specs)
	slices.Reverse(shuffled)
	half := len(shuffled) / 2
	rotated := append(slices.Clone(shuffled[half:]), shuffled[:half]...)
	registry := rule.NewRegistry()
	if err := registry.Load(rotated); err != nil {
		t.Fatalf("shuffled load: %v", err)
	}
	for i, id := range registry.IDs() {
		if id != specs[i].ID {
			t.Fatalf("shuffled input changed order at %d", i)
		}
	}
	assertBadRule(t, rule.Spec{
		ID: "GO-EG-901", Reference: "", Severity: rule.SeverityWarning,
		Category: rule.CategoryStyle, Pattern: "Ident", Message: "x",
	}, rule.ErrEmptyReference)
	assertBadRule(t, rule.Spec{
		ID: "GO-EG-902", Reference: "Effective Go", Severity: rule.SeverityWarning,
		Category: rule.CategoryStyle, Pattern: "(", Message: "x",
	}, rule.ErrBadPattern)
	assertBadRule(t, rule.Spec{
		ID: "GO-EG-001", Reference: "Effective Go", Severity: rule.SeverityWarning,
		Category: rule.CategoryStyle, Pattern: "Ident", Message: "x",
	}, rule.ErrDuplicateID)
	assertBadRule(t, rule.Spec{
		ID: "GO-EG-903", Reference: "Effective Go", Severity: "bogus",
		Category: rule.CategoryStyle, Pattern: "Ident", Message: "x",
	}, rule.ErrInvalidSpec)
	assertBadRule(t, rule.Spec{
		ID: "GO-EG-904", Reference: "Effective Go", Severity: rule.SeverityWarning,
		Category: rule.CategoryStyle, Pattern: "Ident",
		Where: []string{"unknown-op:x"}, Message: "x",
	}, rule.ErrInvalidPredicate)
}

// assertBadRule injects one bad specification into a fresh registry.
func assertBadRule(t *testing.T, spec rule.Spec, want error) {
	t.Helper()
	registry := rule.NewRegistry()
	if err := LoadBuiltins(registry); err != nil {
		t.Fatalf("baseline load: %v", err)
	}
	err := registry.Register(spec)
	if !errors.Is(err, want) {
		t.Fatalf("bad rule %s: error = %v, want %v", spec.ID, err, want)
	}
}

// TestExemptions proves every registered exemption names a real rule and a
// non-empty justification.
func TestExemptions(t *testing.T) {
	t.Parallel()
	valid := make(map[string]struct{}, 128)
	for _, spec := range All() {
		valid[string(spec.ID)] = struct{}{}
	}
	table := Exemptions()
	for id, reason := range table {
		if _, ok := valid[id]; !ok {
			t.Errorf("exemption %s does not name a declared rule", id)
		}
		if strings.TrimSpace(reason) == "" {
			t.Errorf("exemption %s has no justification", id)
		}
		if _, ok := ExemptionFor(id); !ok {
			t.Errorf("ExemptionFor(%s) missed a table entry", id)
		}
	}
	if _, ok := ExemptionFor("GO-EG-001"); ok {
		t.Fatal("GO-EG-001 must not be exempted")
	}
}
