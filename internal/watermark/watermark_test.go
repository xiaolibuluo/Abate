/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package watermark

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

// TestOwnerContainsTribe proves the owner constant carries the watermark.
func TestOwnerContainsTribe(t *testing.T) {
	t.Parallel()
	if !strings.Contains(Owner(), "小黎的部落") {
		t.Fatalf("owner = %q, want 小黎的部落", Owner())
	}
	if Source() != SourceURL {
		t.Fatalf("source = %q, want %q", Source(), SourceURL)
	}
}

// TestFingerprintMatchesCanonical recomputes the digest from the canonical string.
func TestFingerprintMatchesCanonical(t *testing.T) {
	t.Parallel()
	digest := sha256.Sum256([]byte(Canonical))
	want := hex.EncodeToString(digest[:])
	if want != "e92fa6a0a80b33498123d5545e07e5b2c8b1ce30f220e3a933644bd8118b81bc" {
		t.Fatalf("canonical digest = %s, unexpected anchor", want)
	}
	if Fingerprint() != want || FingerprintSHA256 != want {
		t.Fatalf("fingerprint = %q, want %s", Fingerprint(), want)
	}
}

// TestVerifyAcceptsRealAndRejectsDrift covers both outcomes of the verifier.
func TestVerifyAcceptsRealAndRejectsDrift(t *testing.T) {
	t.Parallel()
	if !Verify() {
		t.Fatal("Verify() = false for the committed watermark")
	}
	for _, c := range []struct {
		name        string
		canonical   string
		fingerprint string
		owner       string
		source      string
		want        bool
	}{
		{"real", Canonical, FingerprintSHA256, OwnerName, SourceURL, true},
		{"bad-owner", Canonical, FingerprintSHA256, "别人", SourceURL, false},
		{"bad-source", Canonical, FingerprintSHA256, OwnerName, "https://example.com", false},
		{"bad-canonical", "tampered", FingerprintSHA256, OwnerName, SourceURL, false},
		{"bad-fingerprint", Canonical, strings.Repeat("0", len(FingerprintSHA256)), OwnerName, SourceURL, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if got := verifyCandidate(c.canonical, c.fingerprint, c.owner, c.source); got != c.want {
				t.Fatalf("verify() = %v, want %v", got, c.want)
			}
		})
	}
}

// TestVersionProvenanceLabel proves the version accessor is stable.
func TestVersionProvenanceLabel(t *testing.T) {
	t.Parallel()
	if Version() != "0.1.0" {
		t.Fatalf("version = %q, want 0.1.0", Version())
	}
}
