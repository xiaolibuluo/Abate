/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

import (
	"context"
	"sync"
	"testing"

	"abate/plugin/capability"
	"abate/plugin/host"
)

// TestConcurrentSwapsOnIndependentHosts swaps 32 independent hosts in parallel
// to surface any data race in the phase machine, registry or instance flags.
func TestConcurrentSwapsOnIndependentHosts(t *testing.T) {
	const n = 32
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			patch := uint(idx + 1) //nolint:gosec // idx is a small non-negative loop counter.
			f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, patch), host.PluginTypeOperator)
			sw, err := Begin(context.Background(), f.old, f.next)
			if err != nil {
				t.Errorf("iteration %d Begin: %v", idx, err)
				return
			}
			if sw.State() != UNLOAD {
				t.Errorf("iteration %d state = %s", idx, sw.State())
			}
			if got, ok := f.h.Lookup(f.id); !ok || got != f.next {
				t.Errorf("iteration %d registry did not switch", idx)
			}
		}(i)
	}
	wg.Wait()
}

// TestDrainWithConcurrentInFlightCalls holds many in-flight calls on the old
// instance, switches to the new one, then releases them all and verifies the
// DRAIN phase waited for every single call.
func TestDrainWithConcurrentInFlightCalls(t *testing.T) {
	f := newSwapFixture(t, ver(2, 0, 0), ver(2, 0, 1), host.PluginTypeOperator)
	const n = 24
	allow := make(chan struct{})
	var held sync.WaitGroup
	var done sync.WaitGroup
	held.Add(n)
	done.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			release, ok := f.old.BeginCall()
			if !ok {
				t.Errorf("BeginCall rejected before SWITCH")
				held.Done()
				done.Done()
				return
			}
			held.Done()
			defer done.Done()
			<-allow
			release()
		}()
	}
	held.Wait()

	type result struct {
		sw  *Swap
		err error
	}
	resCh := make(chan result, 1)
	go func() {
		sw, err := Begin(context.Background(), f.old, f.next)
		resCh <- result{sw: sw, err: err}
	}()
	waitUntil(t, func() bool {
		got, found := f.h.Lookup(f.id)
		return found && got == f.next
	})
	wants := []capability.Capability{capability.CapRand}
	if _, err := f.h.Call(f.next, host.EntryApply, []byte("served"), wants); err != nil {
		t.Fatalf("new instance must serve while old drains: %v", err)
	}
	close(allow)
	done.Wait()
	res := <-resCh
	if res.err != nil {
		t.Fatalf("Begin with concurrent drain: %v", res.err)
	}
	if res.sw.State() != UNLOAD {
		t.Fatalf("state = %s, want UNLOAD", res.sw.State())
	}
}

// TestConcurrentStateReads proves State is safe to poll from monitoring
// goroutines after the transaction has completed.
func TestConcurrentStateReads(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	sw, err := Begin(context.Background(), f.old, f.next)
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if sw.State() != UNLOAD {
				t.Error("concurrent State read observed a wrong phase")
			}
		}()
	}
	wg.Wait()
}

// TestSequentialSwapsChainVersions proves a chain of monotonically increasing
// upgrades never interrupts deterministic operator calls.
func TestSequentialSwapsChainVersions(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	current := f.old
	for patch := uint(1); patch <= 8; patch++ {
		nextManifest := newSignedManifest(t, f, ver(1, 0, patch+1), host.PluginTypeOperator)
		next, err := f.h.Load(nextManifest, factoryFor(host.PluginTypeOperator), allCaps())
		if err != nil {
			t.Fatalf("iteration %d Load: %v", patch, err)
		}
		f.h.Registry().Replace(f.id, current)
		sw, err := Begin(context.Background(), current, next)
		if err != nil {
			t.Fatalf("iteration %d Begin: %v", patch, err)
		}
		if sw.State() != UNLOAD {
			t.Fatalf("iteration %d state = %s", patch, sw.State())
		}
		current = next
	}
	if got, ok := f.h.Lookup(f.id); !ok || got != current {
		t.Fatal("final registry does not serve the newest instance")
	}
}
