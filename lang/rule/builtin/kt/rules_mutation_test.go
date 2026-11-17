/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktbuiltin

import (
	"testing"

	"abate/lang/rule"
)

// TestMutationInvalidIDRejected proves a malformed three-part identifier is
// rejected at compile time before it can enter a registry.
func TestMutationInvalidIDRejected(t *testing.T) {
	t.Parallel()
	spec := findSpec("KT-DGN-001")
	spec.ID = "KT-BAD"
	if _, err := rule.Compile(spec); err == nil {
		t.Fatal("Compile accepted an invalid three-part identifier")
	}
}

// TestMutationDuplicateIDRejected proves cloning a local identifier onto another
// local rule makes the whole batch fail registration with ErrDuplicateID.
func TestMutationDuplicateIDRejected(t *testing.T) {
	t.Parallel()
	specs := All()
	mutated := make([]rule.Spec, len(specs))
	copy(mutated, specs)
	for i := range mutated {
		if mutated[i].ID == "KT-DGN-002" {
			mutated[i].ID = "KT-DGN-001"
			break
		}
	}
	registry := rule.NewRegistry()
	if err := registry.Load(mutated); err == nil {
		t.Fatal("Load accepted a duplicate identifier")
	}
}

// TestMutationPatternNoMatch proves rewriting a pattern to one that matches no
// node changes the hit count from a positive fixture to zero.
func TestMutationPatternNoMatch(t *testing.T) {
	t.Parallel()
	original := findSpec("KT-DGN-008")
	unit := buildKTUnit(airKind("Ident"), "var", "", "")
	if !matchAny(original, unit) {
		t.Fatal("original var pattern should hit the fixture")
	}
	mutated := original
	mutated.Pattern = `Ident=__zzz_no_match__`
	if matchAny(mutated, unit) {
		t.Fatal("mutated pattern should no longer match the fixture")
	}
}

// TestMutationUnsortedIDsStillRejected proves identifier validation is order
// independent: a duplicate introduced anywhere in the batch is caught.
func TestMutationUnsortedIDsStillRejected(t *testing.T) {
	t.Parallel()
	specs := All()
	for i := range specs {
		if specs[i].ID == "KT-DGN-012" {
			specs[i].ID = "KT-DGN-007"
			break
		}
	}
	registry := rule.NewRegistry()
	if err := registry.Load(specs); err == nil {
		t.Fatal("Load accepted an unsorted duplicate identifier")
	}
}
