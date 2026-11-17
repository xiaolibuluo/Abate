/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"crypto/sha256"
	"sync"
	"testing"

	"abate/plugin/capability"
	"abate/plugin/signature"
)

// TestDeterministicAcrossIterations proves operator output is reproducible across 100
// independent load+call sequences with no hidden state.
func TestDeterministicAcrossIterations(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/iter", signature.PluginTypeOperator, "clock,rand")
	var baseline string
	for i := 0; i < 100; i++ {
		h := NewHost(store)
		inst, err := h.Load(m, testOperatorFactory, allCapabilities())
		if err != nil {
			t.Fatalf("iteration %d Load: %v", i, err)
		}
		out, err := h.Call(inst, EntryApply, []byte("repeatable"), operatorCaps())
		if err != nil {
			t.Fatalf("iteration %d Call: %v", i, err)
		}
		got := sha256.Sum256(out)
		if i == 0 {
			baseline = string(got[:])
			continue
		}
		if string(got[:]) != baseline {
			t.Fatalf("iteration %d produced divergent operator output", i)
		}
	}
}

// TestConcurrentLoadCallChaos exercises independent hosts under concurrent load and
// call to surface any data race in the host runtime.
func TestConcurrentLoadCallChaos(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/cx", signature.PluginTypeOperator, "clock,rand")
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			h := NewHost(store)
			inst, err := h.Load(m, testOperatorFactory, allCapabilities())
			if err != nil {
				t.Errorf("Load: %v", err)
				return
			}
			if _, err := h.Call(inst, EntryApply, []byte("concurrent"), operatorCaps()); err != nil {
				t.Errorf("Call: %v", err)
			}
		}()
	}
	wg.Wait()
}

// TestConcurrentRegistryChaos stresses the registry with concurrent lookups.
func TestConcurrentRegistryChaos(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/regcx", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	inst, err := h.Load(m, testOperatorFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	const n = 64
	var wg sync.WaitGroup
	wg.Add(2 * n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if _, ok := h.Lookup(inst.Manifest.ID); !ok {
				t.Errorf("Lookup missed a live instance")
			}
			_ = h.Registry().Snapshot()
		}()
		go func() {
			defer wg.Done()
			_ = capability.CapClock
		}()
	}
	wg.Wait()
}
