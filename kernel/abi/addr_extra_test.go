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

// TestAddrMustRoundUp covers successful rounding and the overflow panic.
func TestAddrMustRoundUp(t *testing.T) {
	cases := []struct {
		in   abi.Addr
		want abi.Addr
	}{
		{0, 0},
		{1, 0x1000},
		{0x1234, 0x2000},
		{0x1000, 0x1000},
		{0x1001, 0x2000},
		{0x200000, 0x200000},
	}
	for _, c := range cases {
		if got := c.in.MustRoundUp(); got != c.want {
			t.Errorf("Addr(%#x).MustRoundUp() = %#x, want %#x", c.in, got, c.want)
		}
	}
	assertPanics(t, func() { _ = abi.Addr(^uintptr(0) - 10).MustRoundUp() })
}

// TestAddrPageOffset checks the byte offset of an address within its page.
func TestAddrPageOffset(t *testing.T) {
	cases := []struct {
		in   abi.Addr
		want uint64
	}{
		{0, 0},
		{1, 1},
		{0xfff, 4095},
		{0x1000, 0},
		{0x1234, 0x234},
		{0x200001, 1},
	}
	for _, c := range cases {
		if got := c.in.PageOffset(); got != c.want {
			t.Errorf("Addr(%#x).PageOffset() = %d, want %d", c.in, got, c.want)
		}
		if got := abi.IsPageAligned(c.in); got != (c.want == 0) {
			t.Errorf("Addr(%#x).IsPageAligned() = %v, want %v", c.in, got, c.want == 0)
		}
	}
}

// TestAddrMustToRange covers range construction and the wrap-around panic.
func TestAddrMustToRange(t *testing.T) {
	cases := []struct {
		start  abi.Addr
		length uint64
		want   abi.Addr
	}{
		{0, 0, 0},
		{0, 1, 1},
		{0x1000, 0x1000, 0x2000},
		{0x1000, 0x234, 0x1234},
		{0x200000, 0x200000, 0x400000},
	}
	for _, c := range cases {
		got := c.start.MustToRange(c.length)
		if got.Start != c.start || got.End != c.want {
			t.Errorf("Addr(%#x).MustToRange(%d) = [%#x,%#x), want [%#x,%#x)",
				c.start, c.length, got.Start, got.End, c.start, c.want)
		}
	}
	assertPanics(t, func() { _ = abi.Addr(^uintptr(0)).MustToRange(1) })
	assertPanics(t, func() { _ = abi.Addr(0x1000).MustToRange(^uint64(0)) })
}

// TestAddrRangePageAlignment covers both endpoints of IsPageAligned.
func TestAddrRangePageAlignment(t *testing.T) {
	cases := []struct {
		start abi.Addr
		end   abi.Addr
		want  bool
	}{
		{0x1000, 0x3000, true},
		{0x1001, 0x3000, false},
		{0x1000, 0x3001, false},
		{0x1001, 0x3001, false},
		{0, 0, true},
		{0, 0x1000, true},
	}
	for _, c := range cases {
		ar := abi.AddrRange{Start: c.start, End: c.end}
		if got := ar.IsPageAligned(); got != c.want {
			t.Errorf("AddrRange[%#x,%#x).IsPageAligned() = %v, want %v",
				c.start, c.end, got, c.want)
		}
	}
}

// TestAddrRangeHugePageAlignment covers both endpoints of IsHugePageAligned.
func TestAddrRangeHugePageAlignment(t *testing.T) {
	cases := []struct {
		start abi.Addr
		end   abi.Addr
		want  bool
	}{
		{0x200000, 0x400000, true},
		{0x200001, 0x400000, false},
		{0x200000, 0x400001, false},
		{0x200001, 0x400001, false},
		{0, 0, true},
		{0, 0x200000, true},
	}
	for _, c := range cases {
		ar := abi.AddrRange{Start: c.start, End: c.end}
		if got := ar.IsHugePageAligned(); got != c.want {
			t.Errorf("AddrRange[%#x,%#x).IsHugePageAligned() = %v, want %v",
				c.start, c.end, got, c.want)
		}
	}
}

// TestAddrRangeStringExact checks the exact half-open range rendering.
func TestAddrRangeStringExact(t *testing.T) {
	ar := abi.AddrRange{Start: 0x1000, End: 0x2000}
	if got, want := ar.String(), "[0x1000, 0x2000)"; got != want {
		t.Errorf("AddrRange.String() = %q, want %q", got, want)
	}
}
