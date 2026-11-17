/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"errors"
	"testing"

	"abate/plugin/signature"
)

// TestLoadTamperedHashRejected proves a manifest whose SHA256 field was altered after
// signing is rejected, even though the signature still verifies.
func TestLoadTamperedHashRejected(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/tamper", signature.PluginTypeOperator, "clock,rand")
	m.SHA256[0] ^= 0xff
	h := NewHost(store)
	if _, err := h.Load(m, testOperatorFactory, allCapabilities()); err == nil {
		t.Fatal("Load accepted a tampered-hash manifest")
	} else if !errors.Is(err, signature.ErrManifestHashMismatch) {
		t.Fatalf("Load returned %v, want ErrManifestHashMismatch", err)
	}
}

// TestLoadUnknownAnchorRejected proves a manifest referencing an unregistered trust
// anchor is rejected before any signature check.
func TestLoadUnknownAnchorRejected(t *testing.T) {
	store := signature.NewTrustStore()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	_ = store.Add("real-anchor", pub)
	m := signature.Manifest{
		ID:               signature.PluginID("sample/ghost"),
		Version:          signature.SemVer{Major: 1, Minor: 0, Patch: 0},
		Type:             signature.PluginTypeOperator,
		CapabilitiesDecl: "clock,rand",
		TrustAnchor:      "no-such-anchor",
		BuildRef:         "build-test",
	}
	_ = signature.Sign(store, priv, &m)
	h := NewHost(store)
	if _, err := h.Load(m, testOperatorFactory, allCapabilities()); err == nil {
		t.Fatal("Load accepted a manifest with an unknown anchor")
	} else if !errors.Is(err, signature.ErrTrustAnchorUnknown) {
		t.Fatalf("Load returned %v, want ErrTrustAnchorUnknown", err)
	}
}

// TestCallNilInstance proves calling a nil instance is rejected at the whitelist stage.
func TestCallNilInstance(t *testing.T) {
	h := NewHost(signature.NewTrustStore())
	if _, err := h.Call(nil, EntryApply, nil, operatorCaps()); err == nil {
		t.Fatal("Call accepted a nil instance")
	} else if !errors.Is(err, ErrEntryForbidden) {
		t.Fatalf("Call returned %v, want ErrEntryForbidden", err)
	}
}

// TestLoadWrongTypeFactory proves a manifest whose type disagrees with the factory is
// rejected by the factory itself (fail-closed).
func TestLoadWrongTypeFactory(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/wrong", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	if _, err := h.Load(m, testReporterFactory, allCapabilities()); err == nil {
		t.Fatal("Load accepted an operator manifest via the reporter factory")
	}
}
