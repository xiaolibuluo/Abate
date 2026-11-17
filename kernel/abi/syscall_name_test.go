/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi_test

import (
	"testing"

	"abate/kernel/abi"
)

// TestAMD64NameMatchesTable checks every ported number resolves to its name.
func TestAMD64NameMatchesTable(t *testing.T) {
	for no, want := range abi.AMD64Names {
		got, ok := abi.AMD64Name(abi.Syscall(no))
		if !ok || got != want {
			t.Errorf("AMD64Name(%d) = (%q, %v), want (%q, true)", no, got, ok, want)
		}
	}
}

// TestARM64NameMatchesTable checks every ported number resolves to its name.
func TestARM64NameMatchesTable(t *testing.T) {
	for no, want := range abi.ARM64Names {
		got, ok := abi.ARM64Name(abi.Syscall(no))
		if !ok || got != want {
			t.Errorf("ARM64Name(%d) = (%q, %v), want (%q, true)", no, got, ok, want)
		}
	}
}

// TestSyscallNamePresence checks present and absent numbers agree with the maps.
func TestSyscallNamePresence(t *testing.T) {
	for no := uintptr(0); no < 512; no++ {
		_, amdPresent := abi.AMD64Names[no]
		_, armPresent := abi.ARM64Names[no]
		amdName, amdOK := abi.AMD64Name(abi.Syscall(no))
		armName, armOK := abi.ARM64Name(abi.Syscall(no))
		if amdOK != amdPresent || armOK != armPresent {
			t.Errorf("syscall %d presence: got (%v, %v), want (%v, %v)",
				no, amdOK, armOK, amdPresent, armPresent)
		}
		if !amdOK && amdName != "" {
			t.Errorf("AMD64Name(%d) absent but returned %q", no, amdName)
		}
		if !armOK && armName != "" {
			t.Errorf("ARM64Name(%d) absent but returned %q", no, armName)
		}
		if amdOK && amdName != abi.AMD64Names[no] {
			t.Errorf("AMD64Name(%d) = %q, map has %q", no, amdName, abi.AMD64Names[no])
		}
		if armOK && armName != abi.ARM64Names[no] {
			t.Errorf("ARM64Name(%d) = %q, map has %q", no, armName, abi.ARM64Names[no])
		}
	}
}

// TestSyscallNameOutOfRange checks very large numbers are absent.
func TestSyscallNameOutOfRange(t *testing.T) {
	cases := []abi.Syscall{512, 1000, 65535, abi.Syscall(^uintptr(0))}
	for _, no := range cases {
		if name, ok := abi.AMD64Name(no); ok || name != "" {
			t.Errorf("AMD64Name(%d) = (%q, %v), want (\"\", false)", no, name, ok)
		}
		if name, ok := abi.ARM64Name(no); ok || name != "" {
			t.Errorf("ARM64Name(%d) = (%q, %v), want (\"\", false)", no, name, ok)
		}
	}
}
