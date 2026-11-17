/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

// Addr represents an address in an unspecified address space.
//
// Adapted from gvisor/pkg/hostarch/addr.go, Apache-2.0.
type Addr uintptr

// AddLength adds the given length to v and returns the resulting end address.
// The ok flag is false when the addition overflows the range of Addr.
func (v Addr) AddLength(length uint64) (end Addr, ok bool) {
	end = v + Addr(length)
	ok = end >= v
	return end, ok
}

// RoundDown returns v rounded down to the nearest page boundary.
func (v Addr) RoundDown() Addr {
	return PageRoundDown(v)
}

// RoundUp returns v rounded up to the nearest page boundary. The ok flag is
// false when rounding up overflows the range of Addr.
func (v Addr) RoundUp() (Addr, bool) {
	return PageRoundUp(v)
}

// MustRoundUp is equivalent to RoundUp, but panics if rounding up overflows.
func (v Addr) MustRoundUp() Addr {
	return MustPageRoundUp(v)
}

// PageOffset returns the offset of v into its containing page as a uint64.
func (v Addr) PageOffset() uint64 {
	return uint64(PageOffset(v))
}

// IsPageAligned returns true if v is a multiple of PageSize.
func (v Addr) IsPageAligned() bool {
	return IsPageAligned(v)
}

// ToRange returns the half-open address range [v, v+length). The ok flag is
// false when the end of the range wraps around the address space.
func (v Addr) ToRange(length uint64) (AddrRange, bool) {
	end, ok := v.AddLength(length)
	return AddrRange{Start: v, End: end}, ok
}

// MustToRange is equivalent to ToRange, but panics if the end wraps around.
func (v Addr) MustToRange(length uint64) AddrRange {
	ar, ok := v.ToRange(length)
	if !ok {
		panic("abi.Addr.ToRange wraps around the address space")
	}
	return ar
}
