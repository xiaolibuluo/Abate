/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"crypto/ed25519"
	"errors"
	"testing"
)

func TestSemVerString(t *testing.T) {
	cases := []struct {
		v    SemVer
		want string
	}{
		{SemVer{0, 0, 0}, "0.0.0"},
		{SemVer{1, 2, 3}, "1.2.3"},
		{SemVer{10, 20, 30}, "10.20.30"},
		{SemVer{999, 1, 7}, "999.1.7"},
	}
	for _, tc := range cases {
		if got := tc.v.String(); got != tc.want {
			t.Fatalf("SemVer%v.String() = %q, want %q", tc.v, got, tc.want)
		}
	}
}

func TestSemVerCompareAllBranches(t *testing.T) {
	// Major less, equal, greater.
	if Compare(SemVer{0, 9, 9}, SemVer{1, 0, 0}) >= 0 {
		t.Fatalf("0.9.9 should be less than 1.0.0")
	}
	if Compare(SemVer{2, 0, 0}, SemVer{1, 9, 9}) <= 0 {
		t.Fatalf("2.0.0 should be greater than 1.9.9")
	}
	// Minor decides when major equal.
	if Compare(SemVer{1, 1, 0}, SemVer{1, 2, 0}) >= 0 {
		t.Fatalf("1.1.0 should be less than 1.2.0")
	}
	// Patch decides when major and minor equal.
	if Compare(SemVer{1, 2, 2}, SemVer{1, 2, 3}) >= 0 {
		t.Fatalf("1.2.2 should be less than 1.2.3")
	}
	if Compare(SemVer{1, 2, 3}, SemVer{1, 2, 3}) != 0 {
		t.Fatalf("identical versions compare equal")
	}
}

func TestSignUnresolvedAnchor(t *testing.T) {
	store := NewTrustStore()
	priv := demoPrivateKey(t)
	m := baseManifest()
	m.TrustAnchor = "ghost-anchor"
	if err := Sign(store, priv, &m); !errors.Is(err, ErrTrustAnchorUnknown) {
		t.Fatalf("Sign with unresolved anchor: want ErrTrustAnchorUnknown, got %v", err)
	}
}

func TestAddValidation(t *testing.T) {
	store := NewTrustStore()
	pub, _, err := GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	// Empty id rejected.
	if err := store.Add("", pub); err == nil {
		t.Fatalf("expected error adding empty anchor id")
	}
	// Wrong-size key rejected.
	if err := store.Add("bad/size", ed25519.PublicKey(make([]byte, 10))); err == nil {
		t.Fatalf("expected error adding wrong-size public key")
	}
}
