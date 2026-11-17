/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csbuiltin

import (
	"errors"
	"slices"
	"testing"

	"abate/lang/rule"
	"abate/lang/rule/builtin/common"
)

// TestLoadBuiltinsNilRegistry proves a nil registry is rejected up front.
func TestLoadBuiltinsNilRegistry(t *testing.T) {
	t.Parallel()
	err := LoadBuiltins(nil)
	if !errors.Is(err, common.ErrNilRegistry) {
		t.Fatalf("LoadBuiltins(nil) = %v, want ErrNilRegistry", err)
	}
}

// TestLoadBuiltinsRejectsDuplicateID proves loading the same batch a second
// time into the same registry is rejected with ErrDuplicateID.
func TestLoadBuiltinsRejectsDuplicateID(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	if err := LoadBuiltins(registry); err != nil {
		t.Fatalf("first LoadBuiltins: %v", err)
	}
	err := LoadBuiltins(registry)
	if !errors.Is(err, rule.ErrDuplicateID) {
		t.Fatalf("second LoadBuiltins = %v, want ErrDuplicateID", err)
	}
}

// TestLoadEmptySetNoPanic proves an empty batch loads without side effects.
func TestLoadEmptySetNoPanic(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	if err := registry.Load(nil); err != nil {
		t.Fatalf("Load(nil) = %v", err)
	}
	if registry.Len() != 0 {
		t.Fatalf("registry.Len() = %d, want 0", registry.Len())
	}
}

// TestAllDeterministic proves repeated All calls yield identical identifiers.
func TestAllDeterministic(t *testing.T) {
	t.Parallel()
	first := All()
	second := All()
	if len(first) != len(second) {
		t.Fatalf("All length drift: %d vs %d", len(first), len(second))
	}
	for i := range first {
		if first[i].ID != second[i].ID {
			t.Fatalf("identifier drift at %d: %s vs %s", i, first[i].ID, second[i].ID)
		}
	}
}

// TestFingerprintStable proves the gate key is identical across calls.
func TestFingerprintStable(t *testing.T) {
	t.Parallel()
	first := Fingerprint()
	second := Fingerprint()
	if first != second {
		t.Fatal("Fingerprint differs between calls")
	}
}

// TestShuffledLoadOrderIsDeterministic proves load order never changes the
// registered identifier set, which is what makes the gate reproducible.
func TestShuffledLoadOrderIsDeterministic(t *testing.T) {
	t.Parallel()
	specs := All()
	reversed := make([]rule.Spec, len(specs))
	for i, spec := range specs {
		reversed[len(specs)-1-i] = spec
	}
	registry := rule.NewRegistry()
	if err := registry.Load(reversed); err != nil {
		t.Fatalf("Load(reversed) = %v", err)
	}
	if registry.Len() != len(specs) {
		t.Fatalf("registry.Len() = %d, want %d", registry.Len(), len(specs))
	}
	got := registry.IDs()
	want := registry.IDs()
	slices.Sort(got)
	slices.Sort(want)
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("identifier set drift at %d: %s vs %s", i, got[i], want[i])
		}
	}
}
