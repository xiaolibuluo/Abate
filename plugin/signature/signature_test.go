/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"errors"
	"testing"
)

// demoBuiltinSeed is the FIXED seed of the built-in demo anchor. It lives ONLY in
// this test file so production code never embeds a private key. The matching
// public key is hardcoded in truststore.go as demoBuiltinPublicKey.
var demoBuiltinSeed = []byte{
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
	0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
}

func demoPrivateKey(t testing.TB) ed25519.PrivateKey {
	t.Helper()
	priv := ed25519.NewKeyFromSeed(demoBuiltinSeed)
	if !bytes.Equal(priv.Public().(ed25519.PublicKey), demoBuiltinPublicKey) {
		t.Fatalf("demo seed no longer derives the hardcoded built-in public key")
	}
	return priv
}

func baseManifest() Manifest {
	return Manifest{
		ID:               "abate/network-delay",
		Version:          SemVer{Major: 1, Minor: 2, Patch: 3},
		Type:             PluginTypeOperator,
		CapabilitiesDecl: "clock,rand",
		TrustAnchor:      demoAnchorID,
		BuildRef:         "build-abc123",
	}
}

func mustSign(t testing.TB, store *TrustStore, priv ed25519.PrivateKey, m Manifest) Manifest {
	t.Helper()
	if err := Sign(store, priv, &m); err != nil {
		t.Fatalf("Sign failed: %v", err)
	}
	return m
}

func TestDemoAnchorConsistency(t *testing.T) {
	priv := demoPrivateKey(t)
	pub, ok := priv.Public().(ed25519.PublicKey)
	if !ok {
		t.Fatalf("private key did not yield an ed25519 public key")
	}
	if !bytes.Equal(pub, demoBuiltinPublicKey) {
		t.Fatalf("derived pub mismatch")
	}
	store := NewTrustStore()
	builtin := store.Builtin()
	if len(builtin) != 1 || builtin[demoAnchorID] == nil {
		t.Fatalf("expected exactly the demo anchor in Builtin(), got %d keys", len(builtin))
	}
	if !bytes.Equal(builtin[demoAnchorID], demoBuiltinPublicKey) {
		t.Fatalf("Builtin() returned wrong demo public key")
	}
	fp := store.Fingerprint(demoBuiltinPublicKey)
	if len(fp) != 16 {
		t.Fatalf("fingerprint should be 16 hex chars, got %q", fp)
	}
}

func TestVerifyPositive(t *testing.T) {
	store := NewTrustStore()
	demoPriv := demoPrivateKey(t)

	// User-supplied anchor.
	userPub, userPriv, err := GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	if err := store.Add("vendor/alpha", userPub); err != nil {
		t.Fatalf("Add: %v", err)
	}

	cases := []Manifest{
		mustSign(t, store, demoPriv, baseManifest()),
		mustSign(t, store, userPriv, func() Manifest {
			m := baseManifest()
			m.ID = "vendor/reporter-x"
			m.Type = PluginTypeReporter
			m.CapabilitiesDecl = "observe,fs,net"
			m.TrustAnchor = "vendor/alpha"
			m.Version = SemVer{Major: 0, Minor: 9, Patch: 1}
			return m
		}()),
		mustSign(t, store, demoPriv, func() Manifest {
			m := baseManifest()
			m.ID = "abate/adapter-y"
			m.Type = PluginTypeAdapter
			m.CapabilitiesDecl = "clock,rand,observe,fs,net"
			return m
		}()),
		mustSign(t, store, demoPriv, func() Manifest {
			m := baseManifest()
			m.CapabilitiesDecl = ""
			return m
		}()),
	}
	if len(cases) < 4 {
		t.Fatalf("need at least 4 positive cases, got %d", len(cases))
	}
	for i, m := range cases {
		if err := Verify(store, m); err != nil {
			t.Fatalf("positive case %d rejected: %v", i, err)
		}
	}
}

func TestVerifyTamperedRejected(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	signed := mustSign(t, store, priv, baseManifest())

	tamper := func(mut func(m *Manifest)) Manifest {
		m := signed
		mut(&m)
		return m
	}
	cases := []struct {
		name string
		m    Manifest
	}{
		{"id", tamper(func(m *Manifest) { m.ID = "abate/tampered" })},
		{"version", tamper(func(m *Manifest) { m.Version.Patch = 99 })},
		{"type", tamper(func(m *Manifest) { m.Type = PluginTypeAdapter })},
		{"capdecl", tamper(func(m *Manifest) { m.CapabilitiesDecl = "fs,net" })},
		{"buildref", tamper(func(m *Manifest) { m.BuildRef = "build-evil" })},
	}
	if len(cases) < 4 {
		t.Fatalf("need at least 4 tamper cases, got %d", len(cases))
	}
	for _, tc := range cases {
		err := Verify(store, tc.m)
		if !errors.Is(err, ErrSignatureInvalid) {
			t.Fatalf("tamper %q: want ErrSignatureInvalid, got %v", tc.name, err)
		}
	}
}

func TestVerifyWrongAnchorRejected(t *testing.T) {
	store := NewTrustStore()
	demoPriv := demoPrivateKey(t)
	userPub, _, err := GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	if err := store.Add("vendor/alpha", userPub); err != nil {
		t.Fatalf("Add: %v", err)
	}

	// Case A: unknown anchor id.
	mUnknown := mustSign(t, store, demoPriv, baseManifest())
	mUnknown.TrustAnchor = "no-such-anchor"
	if err := Verify(store, mUnknown); !errors.Is(err, ErrTrustAnchorUnknown) {
		t.Fatalf("unknown anchor: want ErrTrustAnchorUnknown, got %v", err)
	}

	// Case B: empty anchor id.
	mEmpty := mustSign(t, store, demoPriv, baseManifest())
	mEmpty.TrustAnchor = ""
	if err := Verify(store, mEmpty); !errors.Is(err, ErrTrustAnchorUnknown) {
		t.Fatalf("empty anchor: want ErrTrustAnchorUnknown, got %v", err)
	}

	// Case C: signature made by demo key but declared anchor is a different
	// registered key -> signature does not verify.
	mWrong := mustSign(t, store, demoPriv, baseManifest())
	mWrong.TrustAnchor = "vendor/alpha"
	if err := Verify(store, mWrong); !errors.Is(err, ErrSignatureInvalid) {
		t.Fatalf("wrong anchor (different key): want ErrSignatureInvalid, got %v", err)
	}

	// Case D: signature made by an unregistered key, anchor points to a real
	// registered key -> signature does not verify.
	roguePub, roguePriv, err := GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	_ = roguePub
	mRogue := mustSign(t, store, roguePriv, baseManifest())
	mRogue.TrustAnchor = demoAnchorID
	if err := Verify(store, mRogue); !errors.Is(err, ErrSignatureInvalid) {
		t.Fatalf("rogue key under real anchor: want ErrSignatureInvalid, got %v", err)
	}
}

func TestVerifyHashMismatchRejected(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	signed := mustSign(t, store, priv, baseManifest())

	mutateHash := func(mut func(h *[32]byte)) Manifest {
		m := signed
		mut(&m.SHA256)
		return m
	}
	cases := []struct {
		name string
		m    Manifest
	}{
		{"zero", mutateHash(func(h *[32]byte) { *h = [32]byte{} })},
		{"flipped-byte", mutateHash(func(h *[32]byte) { h[0] ^= 0xff })},
		{"fixed-other", mutateHash(func(h *[32]byte) { *h = sha256.Sum256([]byte("other")) })},
		{"all-ones", mutateHash(func(h *[32]byte) { *h = [32]byte{0xff} })},
	}
	if len(cases) < 4 {
		t.Fatalf("need at least 4 hash-mismatch cases, got %d", len(cases))
	}
	for _, tc := range cases {
		err := Verify(store, tc.m)
		if !errors.Is(err, ErrManifestHashMismatch) {
			t.Fatalf("hash mismatch %q: want ErrManifestHashMismatch, got %v", tc.name, err)
		}
	}
}

func TestSignatureFieldIrrelevantToDigest(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	m := mustSign(t, store, priv, baseManifest())

	d1 := m.SignableDigest()
	m.Signature = []byte{0xde, 0xad, 0xbe, 0xef}
	d2 := m.SignableDigest()
	if d1 != d2 {
		t.Fatalf("Signature field must not affect SignableDigest")
	}
	// With the broken signature, verification now fails as invalid.
	if err := Verify(store, m); !errors.Is(err, ErrSignatureInvalid) {
		t.Fatalf("broken signature: want ErrSignatureInvalid, got %v", err)
	}
}
