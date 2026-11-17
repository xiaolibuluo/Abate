/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import "testing"

func TestSemVerParseAndCompare(t *testing.T) {
	valid := []string{"0.0.0", "1.2.3", "10.20.30", "999.0.1"}
	for _, s := range valid {
		if _, err := ParseSemVer(s); err != nil {
			t.Fatalf("ParseSemVer(%q) unexpected error: %v", s, err)
		}
	}
	invalid := []string{"1.2", "a.b.c", "1.2.3.4", "1.x.3", ""}
	for _, s := range invalid {
		if _, err := ParseSemVer(s); err == nil {
			t.Fatalf("ParseSemVer(%q) expected error", s)
		}
	}
	cmp := func(a, b SemVer) int { return Compare(a, b) }
	a := SemVer{1, 0, 0}
	b := SemVer{0, 9, 9}
	if cmp(a, b) <= 0 {
		t.Fatalf("1.0.0 should be greater than 0.9.9")
	}
	if cmp(SemVer{1, 2, 3}, SemVer{1, 2, 3}) != 0 {
		t.Fatalf("identical versions should compare equal")
	}
	if cmp(SemVer{1, 2, 4}, SemVer{1, 2, 3}) <= 0 {
		t.Fatalf("1.2.4 should be greater than 1.2.3")
	}
	if cmp(SemVer{2, 0, 0}, SemVer{1, 9, 9}) <= 0 {
		t.Fatalf("2.0.0 should be greater than 1.9.9")
	}
}

func TestTrustStoreAddRemove(t *testing.T) {
	store := NewTrustStore()
	pub, _, err := GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	if err := store.Add("vendor/beta", pub); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if got := store.Resolve("vendor/beta"); got == nil {
		t.Fatalf("Resolve failed for added anchor")
	}
	// Builtin remains immutable.
	if _, ok := store.Builtin()["vendor/beta"]; ok {
		t.Fatalf("user anchor must not leak into Builtin()")
	}
	// Dedup by fingerprint under a different id.
	if err := store.Add("vendor/gamma", pub); err == nil {
		t.Fatalf("expected fingerprint collision error")
	}
	// Remove builtin is forbidden.
	if err := store.Remove(demoAnchorID); err == nil {
		t.Fatalf("expected error removing built-in anchor")
	}
	if err := store.Remove("vendor/beta"); err != nil {
		t.Fatalf("Remove user anchor: %v", err)
	}
	if got := store.Resolve("vendor/beta"); got != nil {
		t.Fatalf("Resolve should be nil after removal")
	}
	if err := store.Remove("vendor/missing"); err == nil {
		t.Fatalf("expected error removing unknown anchor")
	}
}
