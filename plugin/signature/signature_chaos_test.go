/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"crypto/sha256"
	"sync"
	"testing"
)

// deterministicManifest builds a manifest whose content depends only on i, using
// no random source, so the result is fully reproducible across runs.
func deterministicManifest(i uint) Manifest {
	return Manifest{
		ID:               PluginID("abate/plugin-" + itoa(i)),
		Version:          SemVer{Major: i % 7, Minor: i % 13, Patch: i % 29},
		Type:             PluginTypeOperator,
		CapabilitiesDecl: "clock,rand,observe",
		TrustAnchor:      demoAnchorID,
		BuildRef:         "build-" + itoa(i*31),
	}
}

func itoa(i uint) string {
	if i == 0 {
		return "0"
	}
	var buf [20]byte
	pos := len(buf)
	for i > 0 {
		pos--
		buf[pos] = byte('0' + i%10)
		i /= 10
	}
	return string(buf[pos:])
}

func TestCanonicalBytesDeterministicChaos(t *testing.T) {
	// Same manifest -> 100 canonical byte slices must be byte-identical, yielding
	// exactly one unique hash.
	m := deterministicManifest(0)
	seen := make(map[string]bool)
	for i := uint(0); i < 100; i++ {
		h := sha256.Sum256(m.CanonicalBytes())
		seen[string(h[:])] = true
	}
	if len(seen) != 1 {
		t.Fatalf("same manifest produced %d unique canonical hashes, want 1", len(seen))
	}

	// 100 distinct-but-deterministic manifests: each must be internally stable
	// (compute twice equal) and collectively produce 100 distinct hashes.
	distinct := make(map[string]bool)
	for i := uint(0); i < 100; i++ {
		m := deterministicManifest(i)
		a := m.CanonicalBytes()
		b := m.CanonicalBytes()
		if string(a) != string(b) {
			t.Fatalf("manifest %d canonical bytes not stable across calls", i)
		}
		sum := sha256.Sum256(a)
		distinct[string(sum[:])] = true
	}
	if len(distinct) != 100 {
		t.Fatalf("expected 100 distinct canonical hashes, got %d", len(distinct))
	}
}

func TestVerifyDeterministicChaos(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)

	// Sign with deterministic content, then verify 100 times: identical result.
	m := mustSign(t, store, priv, deterministicManifest(7))
	for i := uint(0); i < 100; i++ {
		if err := Verify(store, m); err != nil {
			t.Fatalf("verify iteration %d failed: %v", i, err)
		}
	}

	// Also verify 100 different deterministic manifests, all must pass.
	for i := uint(0); i < 100; i++ {
		m := mustSign(t, store, priv, deterministicManifest(i))
		if err := Verify(store, m); err != nil {
			t.Fatalf("verify manifest %d failed: %v", i, err)
		}
	}
}

func TestVerifyConcurrentChaos(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	m := mustSign(t, store, priv, deterministicManifest(3))

	const n = 64
	errs := make([]error, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			errs[idx] = Verify(store, m)
		}(i)
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Fatalf("concurrent verify %d failed: %v", i, err)
		}
	}
}

func TestUnknownAnchorRejectedDeterministicChaos(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	m := mustSign(t, store, priv, deterministicManifest(5))
	m.TrustAnchor = "ghost-anchor"
	for i := uint(0); i < 100; i++ {
		if err := Verify(store, m); err == nil {
			t.Fatalf("iteration %d: unknown anchor should be rejected", i)
		}
	}
}
