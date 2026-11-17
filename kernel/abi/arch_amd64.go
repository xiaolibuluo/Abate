//go:build amd64 || 386

/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

import "encoding/binary"

// Page-size and cache-line parameters for x86 targets.
//
// Adapted from gvisor/pkg/hostarch/hostarch_x86.go, Apache-2.0.
const (
	// PageShift is the binary log of the system page size.
	PageShift = 12

	// HugePageShift is the binary log of the system huge page size.
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

	// MinPageShift is the binary log of the smallest supported page size.
	MinPageShift = PageShift

	// MinPageSize is the smallest supported page size.
	MinPageSize = 1 << MinPageShift

	// MaxPageShift is the binary log of the largest supported page size.
	MaxPageShift = PageShift

	// MaxPageSize is the largest supported page size.
	MaxPageSize = 1 << MaxPageShift
)

// ByteOrder is the native byte order (little endian) on x86.
var ByteOrder = binary.LittleEndian

// IsSupportedPageSize returns true if Abate supports running on hosts with the
// given page size.
func IsSupportedPageSize(size int) bool {
	return size == PageSize
}

// UntaggedUserAddr is a no-op on x86, which does not use address tagging.
func UntaggedUserAddr(addr Addr) Addr {
	return addr
}
