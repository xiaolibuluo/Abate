/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"errors"
	"testing"

	"abate/lang/rule"
)

// TestMutationBadPattern proves an uncompilable pattern is reported instead of
// being quietly dropped from the generated batch.
func TestMutationBadPattern(t *testing.T) {
	t.Parallel()
	bad := chaosTemplate(3)
	bad.Pattern = "("
	if _, err := GenerateChecked("PY", []Template{bad}); err == nil {
		t.Fatal("a broken pattern was accepted by GenerateChecked")
	}
	if got := Generate("PY", []Template{bad}); got != nil {
		t.Fatalf("Generate returned %d specs for a broken batch, want nil", len(got))
	}
	if _, err := rule.Compile(rule.Spec{
		ID: "PY-AIR-901", Reference: "PEP 8", Severity: rule.SeverityInfo,
		Category: rule.CategoryStyle, Pattern: "(", Message: "x",
	}); err == nil {
		t.Fatal("the rule engine accepted the mutated pattern")
	}
}

// TestMutationBadWhere proves an unknown predicate op changes load behaviour.
func TestMutationBadWhere(t *testing.T) {
	t.Parallel()
	bad := chaosTemplate(4)
	bad.Where = []string{"unknown-op:argument"}
	if _, err := GenerateChecked("PY", []Template{bad}); err == nil {
		t.Fatal("an unknown predicate op was accepted by GenerateChecked")
	}
	if got := Generate("PY", []Template{bad}); got != nil {
		t.Fatalf("Generate returned %d specs for a bad predicate batch", len(got))
	}
	registry := rule.NewRegistry()
	if err := registry.Register(rule.Spec{
		ID: "PY-AIR-902", Reference: "PEP 8", Severity: rule.SeverityInfo,
		Category: rule.CategoryStyle, Pattern: "Ident",
		Where: []string{"unknown-op:argument"}, Message: "x",
	}); !errors.Is(err, rule.ErrInvalidPredicate) {
		t.Fatalf("registry error = %v, want ErrInvalidPredicate", err)
	}
}

// TestMutationEmptyReference proves a blank citation changes load behaviour.
func TestMutationEmptyReference(t *testing.T) {
	t.Parallel()
	bad := chaosTemplate(5)
	bad.Reference = "   "
	if _, err := GenerateChecked("PY", []Template{bad}); !errors.Is(err, rule.ErrEmptyReference) {
		t.Fatalf("blank reference error = %v, want ErrEmptyReference", err)
	}
	if got := Generate("PY", []Template{bad}); got != nil {
		t.Fatalf("Generate returned %d specs for a blank reference batch", len(got))
	}
	registry := rule.NewRegistry()
	if err := registry.Register(rule.Spec{
		ID: "PY-AIR-903", Reference: "   ", Severity: rule.SeverityInfo,
		Category: rule.CategoryStyle, Pattern: "Ident", Message: "x",
	}); !errors.Is(err, rule.ErrEmptyReference) {
		t.Fatalf("registry error = %v, want ErrEmptyReference", err)
	}
}

// TestMutationTagChangesIdentity proves the language tag is part of the output.
func TestMutationTagChangesIdentity(t *testing.T) {
	t.Parallel()
	py := Generate("PY", nil)
	rs := Generate("RS", nil)
	if len(py) != len(rs) {
		t.Fatal("different tags produced different batch sizes")
	}
	changed := 0
	for i := range py {
		if py[i].ID != rs[i].ID {
			changed++
		}
		if py[i].Canonical() == rs[i].Canonical() {
			t.Fatalf("spec %d did not change with the tag", i)
		}
	}
	if changed != len(py) {
		t.Fatalf("only %d/%d identifiers changed with the tag", changed, len(py))
	}
}

// TestMutationFixtureHitChanges proves a mutated pattern changes real matches.
func TestMutationFixtureHitChanges(t *testing.T) {
	t.Parallel()
	item := builtinTemplates()[0]
	original, err := rule.Compile(All()[0])
	if err != nil {
		t.Fatalf("compile original: %v", err)
	}
	mutatedItem := item
	mutatedItem.Pattern = `Ident/^NEVER_MATCH$/`
	spec := mutatedItem.toSpec("MUT", 1, "docs/reference/c-cert.md — General")
	mutated, err := rule.Compile(spec)
	if err != nil {
		t.Fatalf("compile mutated: %v", err)
	}
	positive := buildFixtureUnit(item.Positive)
	if len(original.Check(positive)) == 0 {
		t.Fatal("original rule did not hit the positive fixture")
	}
	if len(mutated.Check(positive)) != 0 {
		t.Fatal("mutated rule still hit the positive fixture")
	}
}

// TestMutationLocalTemplateChangesBatch proves a local template is observable.
func TestMutationLocalTemplateChangesBatch(t *testing.T) {
	t.Parallel()
	base := Generate("PY", nil)
	extended := Generate("PY", []Template{chaosTemplate(6)})
	if len(extended) != len(base)+1 {
		t.Fatalf("extended batch size = %d, want %d", len(extended), len(base)+1)
	}
	if extended[len(extended)-1].ID == base[len(base)-1].ID {
		t.Fatal("the local template did not receive a fresh identifier")
	}
}
