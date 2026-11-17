/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"abate/kernel/abi"
)

const (
	amd64Golden = "testdata/amd64.golden"
	arm64Golden = "testdata/arm64.golden"
)

// loadFrozenGolden reads a self-developed frozen fact-data golden. The golden
// is a number-to-name snapshot of OS ABI facts, not upstream source code.
// Tests reading it never touch any removed upstream source path.
func loadFrozenGolden(t *testing.T, path string) map[uintptr]string {
	t.Helper()
	// The path is supplied by this test file's own golden fixtures.
	raw, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		t.Fatalf("read frozen golden %s: %v", path, err)
	}
	out := make(map[uintptr]string)
	for _, line := range strings.Split(string(raw), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		fields := strings.Fields(trimmed)
		if len(fields) != 2 {
			t.Fatalf("malformed golden line in %s: %q", path, line)
		}
		no, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			t.Fatalf("malformed syscall number in %s: %q: %v", path, fields[0], err)
		}
		out[uintptr(no)] = fields[1]
	}
	if len(out) == 0 {
		t.Fatalf("frozen golden %s holds no record", path)
	}
	return out
}

// TestAMD64SyscallFrozenGolden verifies the amd64 table has not drifted from
// its frozen snapshot. This is a regression check, not an upstream comparison.
func TestAMD64SyscallFrozenGolden(t *testing.T) {
	golden := loadFrozenGolden(t, amd64Golden)
	if len(golden) != len(abi.AMD64Names) {
		t.Fatalf("amd64 syscall count mismatch: golden=%d table=%d",
			len(golden), len(abi.AMD64Names))
	}
	for no, name := range golden {
		if got, ok := abi.AMD64Names[no]; !ok || got != name {
			t.Errorf("amd64 syscall %d: golden=%q table=%q (present=%v)",
				no, name, got, ok)
		}
	}
}

// TestARM64SyscallFrozenGolden verifies the arm64 table has not drifted from
// its frozen snapshot. This is a regression check, not an upstream comparison.
func TestARM64SyscallFrozenGolden(t *testing.T) {
	golden := loadFrozenGolden(t, arm64Golden)
	if len(golden) != len(abi.ARM64Names) {
		t.Fatalf("arm64 syscall count mismatch: golden=%d table=%d",
			len(golden), len(abi.ARM64Names))
	}
	for no, name := range golden {
		if got, ok := abi.ARM64Names[no]; !ok || got != name {
			t.Errorf("arm64 syscall %d: golden=%q table=%q (present=%v)",
				no, name, got, ok)
		}
	}
}

// TestKnownSyscallNames verifies a representative set of well-known amd64
// syscalls. The arm64 port uses a different numbering scheme and is covered
// exhaustively by TestARM64SyscallFrozenGolden.
func TestKnownSyscallNames(t *testing.T) {
	cases := []struct {
		no   uintptr
		name string
	}{
		{0, "read"},
		{1, "write"},
		{2, "open"},
		{59, "execve"},
		{60, "exit"},
		{322, "execveat"},
		{441, "epoll_pwait2"},
	}
	for _, c := range cases {
		if name, ok := abi.AMD64Names[c.no]; !ok || name != c.name {
			t.Errorf("amd64 %d: want %q got %q (present=%v)", c.no, c.name, name, ok)
		}
	}
}
