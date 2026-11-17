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

func TestErrnoValues(t *testing.T) {
	want := map[abi.Errno]uint32{
		abi.NOERRNO: 0,
		abi.EPERM:   1,
		abi.ENOENT:  2,
		abi.EAGAIN:  11,
		abi.EDEADLK: 35,
		abi.ELOOP:   40,
	}
	for c, v := range want {
		if uint32(c) != v {
			t.Errorf("errno %v = %d, want %d", c, uint32(c), v)
		}
	}
	if abi.EWOULDBLOCK != abi.EAGAIN {
		t.Errorf("EWOULDBLOCK = %d, want %d", abi.EWOULDBLOCK, abi.EAGAIN)
	}
	if abi.EDEADLOCK != abi.EDEADLK {
		t.Errorf("EDEADLOCK = %d, want %d", abi.EDEADLOCK, abi.EDEADLK)
	}
	if abi.ERESTARTSYS != 512 || abi.ERestartBlock != 516 {
		t.Errorf("restart errnos unexpected: %d %d",
			abi.ERESTARTSYS, abi.ERestartBlock)
	}
}

func TestAddrRoundAndAlign(t *testing.T) {
	if got := abi.Addr(0x1234).RoundDown(); got != 0x1000 {
		t.Errorf("RoundDown = %#x, want %#x", got, uintptr(0x1000))
	}
	if got, ok := abi.Addr(0x1234).RoundUp(); !ok || got != 0x2000 {
		t.Errorf("RoundUp = %#x ok=%v, want %#x", got, ok, uintptr(0x2000))
	}
	if !abi.Addr(0x2000).IsPageAligned() {
		t.Errorf("0x2000 should be page aligned")
	}
	if abi.Addr(0x1234).IsPageAligned() {
		t.Errorf("0x1234 should not be page aligned")
	}
	end, ok := abi.Addr(0x1000).AddLength(0x1000)
	if !ok || end != 0x2000 {
		t.Errorf("AddLength = %#x ok=%v, want %#x", end, ok, uintptr(0x2000))
	}
}

func TestAddrRangeToRange(t *testing.T) {
	ar, ok := abi.Addr(0x1000).ToRange(0x1000)
	if !ok {
		t.Fatalf("ToRange overflowed unexpectedly")
	}
	if ar.Start != 0x1000 || ar.End != 0x2000 {
		t.Errorf("ToRange = [%#x,%#x), want [%#x,%#x)",
			ar.Start, ar.End, uintptr(0x1000), uintptr(0x2000))
	}
	if ar.String() == "" {
		t.Errorf("AddrRange.String returned empty")
	}
}

func TestPageHelpers(t *testing.T) {
	if got := abi.PageRoundDown(uintptr(0x1234)); got != 0x1000 {
		t.Errorf("PageRoundDown = %#x, want %#x", got, uintptr(0x1000))
	}
	val, ok := abi.PageRoundUp(uint64(0x1234))
	if !ok || val != 0x2000 {
		t.Errorf("PageRoundUp = %#x ok=%v, want %#x", val, ok, uintptr(0x2000))
	}
	if !abi.IsPageAligned(uintptr(0x2000)) {
		t.Errorf("0x2000 should be page aligned")
	}
}
