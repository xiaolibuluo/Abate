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

func checkPageInvariants(t *testing.T, x uint64) {
	t.Helper()
	down := abi.PageRoundDown(x)
	offset := abi.PageOffset(x)
	if down+offset != x {
		t.Errorf("page decomposition failed for %d: %d + %d", x, down, offset)
	}
	if offset >= 4096 || down%4096 != 0 {
		t.Errorf("page invariants violated for %d: down=%d offset=%d", x, down, offset)
	}
	if abi.IsPageAligned(x) != (offset == 0) {
		t.Errorf("IsPageAligned disagrees with PageOffset for %d", x)
	}
	up, ok := abi.PageRoundUp(x)
	if !ok {
		return
	}
	if up < x || up%4096 != 0 || up-x >= 4096 {
		t.Errorf("PageRoundUp invariant violated for %d: up=%d", x, up)
	}
	if abi.MustPageRoundUp(x) != up {
		t.Errorf("MustPageRoundUp disagrees with PageRoundUp for %d", x)
	}
	pages, pagesOK := abi.ToPagesRoundUp(x)
	if !pagesOK || pages<<abi.PageShift != up {
		t.Errorf("ToPagesRoundUp disagrees with PageRoundUp for %d: %d", x, pages)
	}
}

func checkHugePageInvariants(t *testing.T, x uint64) {
	t.Helper()
	down := abi.HugePageRoundDown(x)
	offset := abi.HugePageOffset(x)
	if down+offset != x {
		t.Errorf("huge page decomposition failed for %d: %d + %d", x, down, offset)
	}
	if offset >= 2097152 || down%2097152 != 0 {
		t.Errorf("huge page invariants violated for %d: down=%d offset=%d", x, down, offset)
	}
	if abi.IsHugePageAligned(x) != (offset == 0) {
		t.Errorf("IsHugePageAligned disagrees with HugePageOffset for %d", x)
	}
	up, ok := abi.HugePageRoundUp(x)
	if !ok {
		return
	}
	if up < x || up%2097152 != 0 || up-x >= 2097152 {
		t.Errorf("HugePageRoundUp invariant violated for %d: up=%d", x, up)
	}
	if abi.MustHugePageRoundUp(x) != up {
		t.Errorf("MustHugePageRoundUp disagrees with HugePageRoundUp for %d", x)
	}
}

func checkCacheLineInvariants(t *testing.T, x uint64) {
	t.Helper()
	down := abi.CacheLineRoundDown(x)
	offset := x & 63
	if down+offset != x {
		t.Errorf("cache line decomposition failed for %d: %d + %d", x, down, offset)
	}
	if down%64 != 0 {
		t.Errorf("CacheLineRoundDown(%d) = %d is not aligned", x, down)
	}
	up, ok := abi.CacheLineRoundUp(x)
	if !ok {
		return
	}
	if up < x || up%64 != 0 || up-x >= 64 {
		t.Errorf("CacheLineRoundUp invariant violated for %d: up=%d", x, up)
	}
	if abi.MustCacheLineRoundUp(x) != up {
		t.Errorf("MustCacheLineRoundUp disagrees with CacheLineRoundUp for %d", x)
	}
}

func checkRoundTrip(t *testing.T, x uint64) {
	t.Helper()
	if abi.PageRoundDown(abi.PageRoundDown(x)) != abi.PageRoundDown(x) {
		t.Errorf("PageRoundDown is not idempotent for %d", x)
	}
	if abi.HugePageRoundDown(abi.HugePageRoundDown(x)) != abi.HugePageRoundDown(x) {
		t.Errorf("HugePageRoundDown is not idempotent for %d", x)
	}
	if abi.CacheLineRoundDown(abi.CacheLineRoundDown(x)) != abi.CacheLineRoundDown(x) {
		t.Errorf("CacheLineRoundDown is not idempotent for %d", x)
	}
	if abi.PageRoundDown(x) != abi.HugePageRoundDown(x)+(abi.HugePageOffset(x)&^4095) {
		t.Errorf("page and huge page decompositions disagree for %d", x)
	}
}

func checkAddrInvariants(t *testing.T, x uint64) {
	t.Helper()
	addr := abi.Addr(x)
	if addr.PageOffset() != abi.PageOffset(x) {
		t.Errorf("Addr.PageOffset disagrees with PageOffset for %d", x)
	}
	if uint64(addr.RoundDown()) != abi.PageRoundDown(x) {
		t.Errorf("Addr.RoundDown disagrees with PageRoundDown for %d", x)
	}
	if addr.IsPageAligned() != abi.IsPageAligned(x) {
		t.Errorf("Addr.IsPageAligned disagrees with IsPageAligned for %d", x)
	}
	up, ok := addr.RoundUp()
	if ok && addr.MustRoundUp() != up {
		t.Errorf("Addr.MustRoundUp disagrees with RoundUp for %d", x)
	}
	ar, rangeOK := addr.ToRange(64)
	if rangeOK {
		if ar.Start != addr || ar.End-addr != 64 {
			t.Errorf("Addr.ToRange produced a broken range for %d", x)
		}
		if ar != addr.MustToRange(64) {
			t.Errorf("Addr.MustToRange disagrees with ToRange for %d", x)
		}
	}
	untagged := abi.UntaggedUserAddr(addr)
	if abi.UntaggedUserAddr(untagged) != untagged {
		t.Errorf("UntaggedUserAddr is not idempotent for %d", x)
	}
}

func checkRangeInvariants(t *testing.T, x uint64) {
	t.Helper()
	start := abi.Addr(x)
	end := abi.Addr(abi.PageRoundDown(x))
	ar := abi.AddrRange{Start: start, End: end}
	wantPage := start.IsPageAligned() && end.IsPageAligned()
	wantHuge := abi.IsHugePageAligned(start) && abi.IsHugePageAligned(end)
	if ar.IsPageAligned() != wantPage {
		t.Errorf("AddrRange.IsPageAligned disagrees with endpoints for %d", x)
	}
	if ar.IsHugePageAligned() != wantHuge {
		t.Errorf("AddrRange.IsHugePageAligned disagrees with endpoints for %d", x)
	}
	if ar.String() == "" {
		t.Errorf("AddrRange.String returned empty for %d", x)
	}
}

// TestChaosAlignmentInvariants checks alignment invariants on random inputs.
func TestChaosAlignmentInvariants(t *testing.T) {
	values := chaosUint64Values()
	assertNoPanic(t, func() {
		for _, x := range values {
			checkPageInvariants(t, x)
			checkHugePageInvariants(t, x)
			checkCacheLineInvariants(t, x)
			checkRoundTrip(t, x)
		}
	})
}

// TestChaosAddressInvariants checks address and range invariants on random inputs.
func TestChaosAddressInvariants(t *testing.T) {
	values := chaosUint64Values()
	assertNoPanic(t, func() {
		for _, x := range values {
			checkAddrInvariants(t, x)
			checkRangeInvariants(t, x)
		}
	})
}

// TestChaosSyscallLookupPanicFree checks random syscall lookups never panic.
func TestChaosSyscallLookupPanicFree(t *testing.T) {
	gen := newChaosSource(0x5eedcafef00d1234)
	assertNoPanic(t, func() {
		for i := 0; i < 8192; i++ {
			no := abi.Syscall(gen.next())
			amdName, amdOK := abi.AMD64Name(no)
			armName, armOK := abi.ARM64Name(no)
			_, amdPresent := abi.AMD64Names[uintptr(no)]
			_, armPresent := abi.ARM64Names[uintptr(no)]
			if amdOK != amdPresent || armOK != armPresent {
				t.Errorf("presence mismatch for syscall %d", no)
			}
			if amdOK && amdName == "" {
				t.Errorf("AMD64Name(%d) present but empty", no)
			}
			if armOK && armName == "" {
				t.Errorf("ARM64Name(%d) present but empty", no)
			}
		}
	})
}

// TestChaosOverflowInputsPanic checks Must* helpers reject overflowing inputs.
func TestChaosOverflowInputsPanic(t *testing.T) {
	for _, delta := range []uint64{0, 1, 2, 62, 63, 64, 4094} {
		x := ^uint64(0) - delta
		assertPanics(t, func() { _ = abi.MustPageRoundUp(x) })
	}
	for _, delta := range []uint64{0, 1, 63, 4095, 2097150} {
		x := ^uint64(0) - delta
		assertPanics(t, func() { _ = abi.MustHugePageRoundUp(x) })
	}
	for _, delta := range []uint64{0, 1, 2, 61, 62} {
		x := ^uint64(0) - delta
		assertPanics(t, func() { _ = abi.MustCacheLineRoundUp(x) })
	}
}
