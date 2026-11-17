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

// TestIsSupportedPageSize pins the accepted host page sizes.
func TestIsSupportedPageSize(t *testing.T) {
	cases := []struct {
		size int
		want bool
	}{
		{4096, true},
		{abi.MaxPageSize, true},
		{0, false},
		{1, false},
		{4095, false},
		{4097, false},
		{8192, false},
		{16384, false},
		{-1, false},
	}
	for _, c := range cases {
		if got := abi.IsSupportedPageSize(c.size); got != c.want {
			t.Errorf("IsSupportedPageSize(%d) = %v, want %v", c.size, got, c.want)
		}
	}
}

// TestUntaggedUserAddr checks that untagged user addresses are preserved.
func TestUntaggedUserAddr(t *testing.T) {
	cases := []abi.Addr{
		0,
		1,
		0x1234,
		0x0000123456789abc,
		0x0000ffffffffffff,
		0x00007fffffffffff,
	}
	for _, addr := range cases {
		if got := abi.UntaggedUserAddr(addr); got != addr {
			t.Errorf("UntaggedUserAddr(%#x) = %#x, want %#x", addr, got, addr)
		}
	}
}

// TestUntaggedUserAddrKeepsLowerBits checks the low address bits survive.
func TestUntaggedUserAddrKeepsLowerBits(t *testing.T) {
	const lowBitsMask = uint64(0x00ffffffffffffff)
	cases := []uint64{
		0xab00000000001234,
		0xff123456789abcde,
		^uint64(0),
		0x0000800000000000,
	}
	for _, raw := range cases {
		got := uint64(abi.UntaggedUserAddr(abi.Addr(raw)))
		if got&lowBitsMask != raw&lowBitsMask {
			t.Errorf("UntaggedUserAddr(%#x) = %#x, low 56 bits differ", raw, got)
		}
		if again := abi.UntaggedUserAddr(abi.Addr(got)); again != abi.Addr(got) {
			t.Errorf("UntaggedUserAddr is not idempotent for %#x", raw)
		}
	}
}

// TestByteOrderDecodesLittleEndian checks the native byte order decoder.
func TestByteOrderDecodesLittleEndian(t *testing.T) {
	if got := abi.ByteOrder.Uint32([]byte{0x01, 0x02, 0x03, 0x04}); got != 0x04030201 {
		t.Errorf("ByteOrder.Uint32 = %#x, want 0x04030201", got)
	}
	if got := abi.ByteOrder.Uint16([]byte{0x34, 0x12}); got != 0x1234 {
		t.Errorf("ByteOrder.Uint16 = %#x, want 0x1234", got)
	}
	if got := abi.ByteOrder.Uint64([]byte{8, 7, 6, 5, 4, 3, 2, 1}); got != 0x0102030405060708 {
		t.Errorf("ByteOrder.Uint64 = %#x, want 0x0102030405060708", got)
	}
}
