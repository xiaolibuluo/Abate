/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

import (
	"sync"
	"testing"
)

// deterministicCaps derives a capability set from i using a fixed rule so the
// result is reproducible with no random source.
func deterministicCaps(i int) Capabilities {
	var c Capabilities
	for k := 0; k < len(allCaps); k++ {
		if (i & (1 << k)) != 0 {
			c.Grant(allCaps[k])
		}
	}
	return c
}

func TestBitOpsDeterministic100(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := deterministicCaps(i)
		b := deterministicCaps(i)
		if a != b {
			t.Fatalf("iteration %d: non-deterministic capability construction", i)
		}
		for k := 0; k < len(allCaps); k++ {
			want := (i & (1 << k)) != 0
			if a.Has(allCaps[k]) != want {
				t.Fatalf("iteration %d cap %d: Has mismatch", i, k)
			}
		}
		// Grant then immediate Revoke returns to zero for that bit and is stable.
		var mut Capabilities
		mut.Grant(CapClock)
		mut.Grant(CapClock) // idempotent
		if !mut.Has(CapClock) {
			t.Fatalf("iteration %d: Grant failed", i)
		}
		mut.Revoke(CapClock)
		mut.Revoke(CapClock) // idempotent
		if mut.Has(CapClock) {
			t.Fatalf("iteration %d: Revoke failed", i)
		}
	}
}

func TestParseGrantDeterministic100(t *testing.T) {
	for i := 0; i < 100; i++ {
		decl := "clock,rand"
		if i%2 == 0 {
			decl = "observe,fs,net"
		}
		first, err1 := ParseGrant(decl)
		second, err2 := ParseGrant(decl)
		if err1 != nil || err2 != nil {
			t.Fatalf("iteration %d: unexpected parse error: %v %v", i, err1, err2)
		}
		if first != second {
			t.Fatalf("iteration %d: non-deterministic ParseGrant", i)
		}
	}
}

func TestConcurrentSafeChaos(t *testing.T) {
	// Per-instance bit ops on distinct values are data-race free.
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			c := deterministicCaps(idx)
			_ = c.Has(CapNet)
			_ = Check(c, CapNet)
		}(i)
	}
	wg.Wait()

	// Shared AuditLog supports concurrent appends safely.
	log := NewAuditLog()
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			log.Append(AuditEvent{Kind: "use", PluginID: "p", Capability: CapClock, Result: true})
			_ = idx
		}(i)
	}
	wg.Wait()
	if log.Len() != n {
		t.Fatalf("expected %d audit entries after concurrent appends, got %d", n, log.Len())
	}

	// Concurrent audited checks against a shared log are safe.
	var caps Capabilities
	caps.Grant(CapClock)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			_ = CheckAudit(caps, CapClock, log, "p")
		}()
	}
	wg.Wait()
	if log.Len() != 2*n {
		t.Fatalf("expected %d audit entries, got %d", 2*n, log.Len())
	}
}
