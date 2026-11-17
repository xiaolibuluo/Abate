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

var archConstants = []struct {
	name string
	got  uint64
	want uint64
}{
	{"PageShift", abi.PageShift, 12},
	{"PageSize", abi.PageSize, 4096},
	{"PageMask", abi.PageMask, 4095},
	{"HugePageShift", abi.HugePageShift, 21},
	{"HugePageSize", abi.HugePageSize, 2097152},
	{"HugePageMask", abi.HugePageMask, 2097151},
	{"JumboPageShift", abi.JumboPageShift, 30},
	{"JumboPageSize", abi.JumboPageSize, 1073741824},
	{"CacheLineShift", abi.CacheLineShift, 6},
	{"CacheLineSize", abi.CacheLineSize, 64},
	{"CacheLineMask", abi.CacheLineMask, 63},
	{"MinPageShift", abi.MinPageShift, 12},
	{"MinPageSize", abi.MinPageSize, 4096},
	{"MaxPageShift", abi.MaxPageShift, 12},
	{"MaxPageSize", abi.MaxPageSize, 4096},
}

var errnoAnchors = []struct {
	name string
	got  uint64
	want uint64
}{
	{"NOERRNO", abi.NOERRNO, 0},
	{"EPERM", abi.EPERM, 1},
	{"ENOENT", abi.ENOENT, 2},
	{"EAGAIN", abi.EAGAIN, 11},
	{"EDEADLK", abi.EDEADLK, 35},
	{"ELOOP", abi.ELOOP, 40},
	{"ENOSYS", abi.ENOSYS, 38},
	{"EHWPOISON", abi.EHWPOISON, 133},
	{"EWOULDBLOCK", abi.EWOULDBLOCK, 11},
	{"EDEADLOCK", abi.EDEADLOCK, 35},
	{"ERESTARTSYS", abi.ERESTARTSYS, 512},
	{"ERESTARTNOINTR", abi.ERESTARTNOINTR, 513},
	{"ERESTARTNOHAND", abi.ERESTARTNOHAND, 514},
	{"ERestartBlock", abi.ERestartBlock, 516},
}

// TestMutationArchitectureConstants pins the host architecture parameters.
func TestMutationArchitectureConstants(t *testing.T) {
	for _, c := range archConstants {
		if c.got != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.got, c.want)
		}
	}
}

// TestMutationErrnoAnchors pins representative Linux errno numbers.
func TestMutationErrnoAnchors(t *testing.T) {
	for _, c := range errnoAnchors {
		if c.got != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.got, c.want)
		}
	}
}

// TestMutationJumboPageMask pins the 1GB jumbo page mask.
func TestMutationJumboPageMask(t *testing.T) {
	want := ^uintptr(1073741823)
	if abi.JumboPageMask != want {
		t.Errorf("JumboPageMask = %#x, want %#x", abi.JumboPageMask, want)
	}
	got := abi.HugePageRoundDown(uint64(abi.JumboPageSize) + 1)
	if got != abi.JumboPageSize {
		t.Errorf("HugePageRoundDown(JumboPageSize+1) = %d, want %d",
			got, abi.JumboPageSize)
	}
}

// TestMutationPageBoundaryExpectations pins 4KB page rounding results.
func TestMutationPageBoundaryExpectations(t *testing.T) {
	cases := []struct {
		in      uint64
		down    uint64
		up      uint64
		pages   uint64
		offset  uint64
		aligned bool
	}{
		{0, 0, 0, 0, 0, true},
		{1, 0, 4096, 1, 1, false},
		{4095, 0, 4096, 1, 4095, false},
		{4096, 4096, 4096, 1, 0, true},
		{4097, 4096, 8192, 2, 1, false},
		{2097152, 2097152, 2097152, 512, 0, true},
		{2097153, 2097152, 2101248, 513, 1, false},
	}
	for _, c := range cases {
		if got := abi.PageRoundDown(c.in); got != c.down {
			t.Errorf("PageRoundDown(%d) = %d, want %d", c.in, got, c.down)
		}
		val, ok := abi.PageRoundUp(c.in)
		if !ok || val != c.up {
			t.Errorf("PageRoundUp(%d) = (%d, %v), want (%d, true)",
				c.in, val, ok, c.up)
		}
		if got := abi.MustPageRoundUp(c.in); got != c.up {
			t.Errorf("MustPageRoundUp(%d) = %d, want %d", c.in, got, c.up)
		}
		pages, pagesOK := abi.ToPagesRoundUp(c.in)
		if !pagesOK || pages != c.pages {
			t.Errorf("ToPagesRoundUp(%d) = (%d, %v), want (%d, true)",
				c.in, pages, pagesOK, c.pages)
		}
		if got := abi.PageOffset(c.in); got != c.offset {
			t.Errorf("PageOffset(%d) = %d, want %d", c.in, got, c.offset)
		}
		if got := abi.IsPageAligned(c.in); got != c.aligned {
			t.Errorf("IsPageAligned(%d) = %v, want %v", c.in, got, c.aligned)
		}
	}
}

// TestMutationHugePageBoundaryExpectations pins 2MB huge page rounding results.
func TestMutationHugePageBoundaryExpectations(t *testing.T) {
	cases := []struct {
		in      uint64
		down    uint64
		up      uint64
		offset  uint64
		aligned bool
	}{
		{0, 0, 0, 0, true},
		{1, 0, 2097152, 1, false},
		{2097151, 0, 2097152, 2097151, false},
		{2097152, 2097152, 2097152, 0, true},
		{2097153, 2097152, 4194304, 1, false},
		{4194304, 4194304, 4194304, 0, true},
	}
	for _, c := range cases {
		if got := abi.HugePageRoundDown(c.in); got != c.down {
			t.Errorf("HugePageRoundDown(%d) = %d, want %d", c.in, got, c.down)
		}
		val, ok := abi.HugePageRoundUp(c.in)
		if !ok || val != c.up {
			t.Errorf("HugePageRoundUp(%d) = (%d, %v), want (%d, true)",
				c.in, val, ok, c.up)
		}
		if got := abi.MustHugePageRoundUp(c.in); got != c.up {
			t.Errorf("MustHugePageRoundUp(%d) = %d, want %d", c.in, got, c.up)
		}
		if got := abi.HugePageOffset(c.in); got != c.offset {
			t.Errorf("HugePageOffset(%d) = %d, want %d", c.in, got, c.offset)
		}
		if got := abi.IsHugePageAligned(c.in); got != c.aligned {
			t.Errorf("IsHugePageAligned(%d) = %v, want %v", c.in, got, c.aligned)
		}
	}
}

// TestMutationCacheLineBoundaryExpectations pins 64 byte cache line results.
func TestMutationCacheLineBoundaryExpectations(t *testing.T) {
	cases := []struct {
		in     uint64
		down   uint64
		up     uint64
		offset uint64
	}{
		{0, 0, 0, 0},
		{1, 0, 64, 1},
		{63, 0, 64, 63},
		{64, 64, 64, 0},
		{65, 64, 128, 1},
		{4096, 4096, 4096, 0},
		{4097, 4096, 4160, 1},
	}
	for _, c := range cases {
		if got := abi.CacheLineRoundDown(c.in); got != c.down {
			t.Errorf("CacheLineRoundDown(%d) = %d, want %d", c.in, got, c.down)
		}
		val, ok := abi.CacheLineRoundUp(c.in)
		if !ok || val != c.up {
			t.Errorf("CacheLineRoundUp(%d) = (%d, %v), want (%d, true)",
				c.in, val, ok, c.up)
		}
		if got := abi.MustCacheLineRoundUp(c.in); got != c.up {
			t.Errorf("MustCacheLineRoundUp(%d) = %d, want %d", c.in, got, c.up)
		}
		if got := c.in & 63; got != c.offset {
			t.Errorf("cache line offset of %d = %d, want %d", c.in, got, c.offset)
		}
	}
}

// TestMutationSyscallAnchors pins canonical syscall number to name pairs.
func TestMutationSyscallAnchors(t *testing.T) {
	amdCases := []struct {
		no   abi.Syscall
		name string
	}{
		{0, "read"},
		{1, "write"},
		{2, "open"},
		{59, "execve"},
		{60, "exit"},
		{221, "fadvise64"},
		{322, "execveat"},
		{334, "rseq"},
		{424, "pidfd_send_signal"},
		{441, "epoll_pwait2"},
	}
	for _, c := range amdCases {
		got, ok := abi.AMD64Name(c.no)
		if !ok || got != c.name {
			t.Errorf("AMD64Name(%d) = (%q, %v), want (%q, true)",
				c.no, got, ok, c.name)
		}
	}
	armCases := []struct {
		no   abi.Syscall
		name string
	}{
		{0, "io_setup"},
		{63, "read"},
		{64, "write"},
		{221, "execve"},
		{222, "mmap"},
		{281, "execveat"},
		{293, "rseq"},
		{441, "epoll_pwait2"},
	}
	for _, c := range armCases {
		got, ok := abi.ARM64Name(c.no)
		if !ok || got != c.name {
			t.Errorf("ARM64Name(%d) = (%q, %v), want (%q, true)",
				c.no, got, ok, c.name)
		}
	}
}

// TestMutationAddrAnchors pins address range construction results.
func TestMutationAddrAnchors(t *testing.T) {
	ar := abi.Addr(0x1000).MustToRange(0x2000)
	if ar.Start != 0x1000 || ar.End != 0x3000 {
		t.Errorf("MustToRange = [%#x,%#x), want [0x1000,0x3000)", ar.Start, ar.End)
	}
	if !ar.IsPageAligned() || ar.IsHugePageAligned() {
		t.Errorf("MustToRange alignment flags are wrong for %v", ar)
	}
	if got := abi.Addr(0x1234).PageOffset(); got != 0x234 {
		t.Errorf("Addr(0x1234).PageOffset() = %d, want 0x234", got)
	}
	if got := abi.Addr(0x1234).MustRoundUp(); got != 0x2000 {
		t.Errorf("Addr(0x1234).MustRoundUp() = %#x, want 0x2000", got)
	}
	big := abi.Addr(0x200000).MustToRange(0x200000)
	if !big.IsHugePageAligned() || !big.IsPageAligned() {
		t.Errorf("huge page range is not aligned: %v", big)
	}
	if got := big.String(); got != "[0x200000, 0x400000)" {
		t.Errorf("huge page range String() = %q, want [0x200000, 0x400000)", got)
	}
}
