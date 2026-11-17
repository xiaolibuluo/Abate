/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cbuiltin

import (
	"errors"
	"slices"
	"sync"
	"testing"

	"abate/lang/rule"
)

// TestRuleSetChaosHash proves repeated and concurrent builds yield one
// identical rule-set digest.
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

// TestLoadShuffledOrder proves shuffled declaration input still yields the same
// ordered, duplicate-free rule set.
func TestLoadShuffledOrder(t *testing.T) {
	t.Parallel()
	specs := All()
	shuffled := slices.Clone(specs)
	slices.Reverse(shuffled)
	registry := rule.NewRegistry()
	if err := registry.Load(shuffled); err != nil {
		t.Fatalf("shuffled load: %v", err)
	}
	for i, id := range registry.IDs() {
		if id != specs[i].ID {
			t.Fatalf("shuffled input changed order at %d", i)
		}
	}
}

// TestDuplicateIDRejected proves loading a second copy of an identifier is
// rejected with ErrDuplicateID rather than silently overwriting.
func TestDuplicateIDRejected(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	if err := LoadBuiltins(registry); err != nil {
		t.Fatalf("baseline load: %v", err)
	}
	dup := cRules()[0]
	err := registry.Register(dup)
	if !errors.Is(err, rule.ErrDuplicateID) {
		t.Fatalf("duplicate id %s: error = %v, want ErrDuplicateID", dup.ID, err)
	}
}

// TestNilRegistryRejected proves LoadBuiltins guards against a nil registry.
func TestNilRegistryRejected(t *testing.T) {
	t.Parallel()
	err := LoadBuiltins(nil)
	if !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("LoadBuiltins(nil): error = %v, want ErrNilRegistry", err)
	}
}

// TestEmptyLoadDoesNotPanic proves loading an empty batch leaves the registry
// empty without panicking.
func TestEmptyLoadDoesNotPanic(t *testing.T) {
	t.Parallel()
	registry := rule.NewRegistry()
	if err := registry.Load([]rule.Spec{}); err != nil {
		t.Fatalf("empty load: %v", err)
	}
	if registry.Len() != 0 {
		t.Fatalf("empty load left %d rules", registry.Len())
	}
}
