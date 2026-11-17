/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"errors"
	"testing"
)

// TestMutationPerFieldDetection ensures a single-field mutation of a signed
// manifest is always detected (no silent acceptance).
func TestMutationPerFieldDetection(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	good := mustSign(t, store, priv, baseManifest())

	cases := []struct {
		name string
		mut  func(m *Manifest)
		want error
	}{
		{"id", func(m *Manifest) { m.ID = "abate/other" }, ErrSignatureInvalid},
		{"version", func(m *Manifest) { m.Version.Minor = 42 }, ErrSignatureInvalid},
		{"type", func(m *Manifest) { m.Type = PluginTypeReporter }, ErrSignatureInvalid},
		{"capdecl", func(m *Manifest) { m.CapabilitiesDecl = "fs" }, ErrSignatureInvalid},
		{"buildref", func(m *Manifest) { m.BuildRef = "x" }, ErrSignatureInvalid},
		{"sha256", func(m *Manifest) { m.SHA256 = [32]byte{0x01} }, ErrManifestHashMismatch},
		{"trustanchor", func(m *Manifest) { m.TrustAnchor = "unknown" }, ErrTrustAnchorUnknown},
	}
	for _, tc := range cases {
		m := good
		tc.mut(&m)
		err := Verify(store, m)
		if !errors.Is(err, tc.want) {
			t.Fatalf("mutation %q: want %v, got %v", tc.name, tc.want, err)
		}
	}
}

func TestEmptyManifestRejected(t *testing.T) {
	store := NewTrustStore()
	m := Manifest{}
	if err := Verify(store, m); err == nil {
		t.Fatalf("empty manifest must be rejected")
	}
}

func TestSignatureLengthAnomaly(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	m := mustSign(t, store, priv, baseManifest())

	// Wrong-length signature must fail verification, not panic.
	m.Signature = []byte{0x01, 0x02, 0x03}
	if err := Verify(store, m); !errors.Is(err, ErrSignatureInvalid) {
		t.Fatalf("wrong-length signature: want ErrSignatureInvalid, got %v", err)
	}

	// Zero-length signature must also be rejected cleanly.
	m.Signature = nil
	if err := Verify(store, m); !errors.Is(err, ErrSignatureInvalid) {
		t.Fatalf("nil signature: want ErrSignatureInvalid, got %v", err)
	}
}

func TestCrossCheckBoundary(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	m := mustSign(t, store, priv, baseManifest())

	// A valid signature with a zero SHA256 field must be flagged as a hash
	// mismatch, never accepted.
	m.SHA256 = [32]byte{}
	if err := Verify(store, m); !errors.Is(err, ErrManifestHashMismatch) {
		t.Fatalf("zero hash with valid signature: want ErrManifestHashMismatch, got %v", err)
	}
}
