/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

import "testing"

// TestMergeSyscallMapsMergesFragments checks disjoint fragments are concatenated.
func TestMergeSyscallMapsMergesFragments(t *testing.T) {
	first := map[uintptr]string{0: "read", 1: "write"}
	second := map[uintptr]string{2: "open"}
	got := mergeSyscallMaps(first, second, map[uintptr]string{})
	if len(got) != 3 {
		t.Fatalf("mergeSyscallMaps returned %d entries, want 3", len(got))
	}
	cases := []struct {
		no   uintptr
		name string
	}{
		{0, "read"},
		{1, "write"},
		{2, "open"},
	}
	for _, c := range cases {
		if got[c.no] != c.name {
			t.Errorf("merged[%d] = %q, want %q", c.no, got[c.no], c.name)
		}
	}
}

// TestMergeSyscallMapsRejectsDuplicate checks a duplicate number panics.
func TestMergeSyscallMapsRejectsDuplicate(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected a panic for a duplicate syscall number")
		}
	}()
	_ = mergeSyscallMaps(
		map[uintptr]string{7: "poll"},
		map[uintptr]string{7: "dup"},
	)
}

// TestPortedTablesAreComplete checks the assembled tables span their fragments.
func TestPortedTablesAreComplete(t *testing.T) {
	amd := mergeSyscallMaps(amd64NamesLow, amd64NamesHigh)
	arm := mergeSyscallMaps(arm64NamesLow, arm64NamesHigh)
	if len(amd) != len(AMD64Names) || len(arm) != len(ARM64Names) {
		t.Errorf("assembled tables differ: amd=%d/%d arm=%d/%d",
			len(amd), len(AMD64Names), len(arm), len(ARM64Names))
	}
	for no, name := range amd {
		if AMD64Names[no] != name {
			t.Errorf("AMD64Names[%d] = %q, want %q", no, AMD64Names[no], name)
		}
	}
	for no, name := range arm {
		if ARM64Names[no] != name {
			t.Errorf("ARM64Names[%d] = %q, want %q", no, ARM64Names[no], name)
		}
	}
}
