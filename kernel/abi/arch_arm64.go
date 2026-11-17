//go:build arm64

/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

import "encoding/binary"

// Page-size and cache-line parameters for arm64 targets.
//
// Adapted from gvisor/pkg/hostarch/hostarch_arm64.go and
// hostarch_arm64_4k.go, Apache-2.0. arm64 supports 4K/16K/64K page sizes;
// Abate targets the 4K configuration by default.
const (
	// PageShift is the binary log of the system page size (4K).
	PageShift = 12

	// HugePageShift is the binary log of the system huge page size (2MB).
	HugePageShift = 21

	// JumboPageShift is the binary log of the 1GB jumbo page size.
	JumboPageShift = 30

	// CacheLineShift is the binary log of the cache line size.
	CacheLineShift = 6

	// PageSize is the system page size.
	PageSize = 1 << PageShift

	// PageMask is PageSize - 1.
	PageMask = PageSize - 1

	// HugePageSize is the system huge page size.
	HugePageSize = 1 << HugePageShift

	// HugePageMask is HugePageSize - 1.
	HugePageMask = HugePageSize - 1

	// JumboPageSize is the 1GB jumbo page size.
	JumboPageSize = 1 << JumboPageShift

	// CacheLineSize is the size of the cache line.
	CacheLineSize = 1 << CacheLineShift

	// MinPageShift is the binary log of the smallest supported page size (4K).
	MinPageShift = 12

	// MinPageSize is the smallest supported page size.
	MinPageSize = 1 << MinPageShift

	// MaxPageShift is the binary log of the largest supported page size (16K).
	MaxPageShift = 16

	// MaxPageSize is the largest supported page size.
	MaxPageSize = 1 << MaxPageShift
)

// ByteOrder is the native byte order (little endian) on arm64.
var ByteOrder = binary.LittleEndian

// IsSupportedPageSize returns true if Abate supports running on hosts with the
// given page size.
func IsSupportedPageSize(size int) bool {
	return size == 1<<12 || size == 1<<16
}

// UntaggedUserAddr clears the top byte from the address pointer. arm64 enables
// Top-Byte-Ignore (TBI0), so bits [63:56] of user addresses are ignored.
func UntaggedUserAddr(addr Addr) Addr {
	return Addr(int64(addr<<8) >> 8)
}
