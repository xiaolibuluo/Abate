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

// TestMustPageRoundUp covers successful rounding and the overflow panic.
func TestMustPageRoundUp(t *testing.T) {
	cases := []struct {
		in   uint64
		want uint64
	}{
		{0, 0},
		{1, 4096},
		{4095, 4096},
		{4096, 4096},
		{4097, 8192},
		{2097153, 2101248},
	}
	for _, c := range cases {
		got := abi.MustPageRoundUp(c.in)
		if got != c.want {
			t.Errorf("MustPageRoundUp(%d) = %d, want %d", c.in, got, c.want)
		}
	}
	assertPanics(t, func() { _ = abi.MustPageRoundUp(^uint64(0) - 10) })
}

// TestToPagesRoundUp verifies page counting and the overflow branch.
func TestToPagesRoundUp(t *testing.T) {
	cases := []struct {
		in     uint64
		want   uint64
		wantOK bool
	}{
		{0, 0, true},
		{1, 1, true},
		{4095, 1, true},
		{4096, 1, true},
		{4097, 2, true},
		{2097153, 513, true},
	}
	for _, c := range cases {
		got, ok := abi.ToPagesRoundUp(c.in)
		if got != c.want || ok != c.wantOK {
			t.Errorf("ToPagesRoundUp(%d) = (%d, %v), want (%d, %v)",
				c.in, got, ok, c.want, c.wantOK)
		}
	}
	overflow := ^uint64(0) - 10
	got, ok := abi.ToPagesRoundUp(overflow)
	if got != overflow || ok {
		t.Errorf("ToPagesRoundUp(overflow) = (%d, %v), want (%d, false)",
			got, ok, overflow)
	}
	val, valid := abi.PageRoundUp(overflow)
	if valid || val >= overflow {
		t.Errorf("PageRoundUp(overflow) = (%d, %v), want wrapped value and false",
			val, valid)
	}
}

// TestHugePageHelpers covers the whole huge-page rounding family.
func TestHugePageHelpers(t *testing.T) {
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
	overflow := ^uint64(0) - 10
	val, ok := abi.HugePageRoundUp(overflow)
	if ok || val >= overflow {
		t.Errorf("HugePageRoundUp(overflow) = (%d, %v), want wrapped value and false",
			val, ok)
	}
	assertPanics(t, func() { _ = abi.MustHugePageRoundUp(overflow) })
}

// TestCacheLineHelpers covers the whole cache-line rounding family.
func TestCacheLineHelpers(t *testing.T) {
	cases := []struct {
		in   uint64
		down uint64
		up   uint64
	}{
		{0, 0, 0},
		{1, 0, 64},
		{63, 0, 64},
		{64, 64, 64},
		{65, 64, 128},
		{4097, 4096, 4160},
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
	}
	overflow := ^uint64(0) - 10
	val, ok := abi.CacheLineRoundUp(overflow)
	if ok || val >= overflow {
		t.Errorf("CacheLineRoundUp(overflow) = (%d, %v), want wrapped value and false",
			val, ok)
	}
	assertPanics(t, func() { _ = abi.MustCacheLineRoundUp(overflow) })
}

// TestPageHelpersAcrossIntegerTypes exercises every allowed integer type.
func TestPageHelpersAcrossIntegerTypes(t *testing.T) {
	if got := abi.MustPageRoundUp(uint(4097)); got != 8192 {
		t.Errorf("MustPageRoundUp(uint) = %d, want 8192", got)
	}
	if got := abi.MustPageRoundUp(uint32(4097)); got != 8192 {
		t.Errorf("MustPageRoundUp(uint32) = %d, want 8192", got)
	}
	if got := abi.MustPageRoundUp(uintptr(4097)); got != 8192 {
		t.Errorf("MustPageRoundUp(uintptr) = %d, want 8192", got)
	}
	if got := abi.MustHugePageRoundUp(uint32(2097153)); got != 4194304 {
		t.Errorf("MustHugePageRoundUp(uint32) = %d, want 4194304", got)
	}
	if got := abi.MustCacheLineRoundUp(uint32(65)); got != 128 {
		t.Errorf("MustCacheLineRoundUp(uint32) = %d, want 128", got)
	}
	assertPanics(t, func() { _ = abi.MustPageRoundUp(^uint32(0)) })
	assertPanics(t, func() { _ = abi.MustHugePageRoundUp(^uint32(0)) })
	assertPanics(t, func() { _ = abi.MustCacheLineRoundUp(^uint32(0)) })
}
