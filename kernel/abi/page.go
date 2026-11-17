/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

// Masks often used when working with alignment in constant expressions.
//
// Adapted from gvisor/pkg/hostarch/sizes_util.go, Apache-2.0.
const (
	// CacheLineMask is CacheLineSize - 1.
	CacheLineMask = CacheLineSize - 1

	// JumboPageMask is the mask for the 1GB jumbo page size.
	JumboPageMask = ^uintptr(JumboPageSize - 1)
)

// bytecount constrains the integer types accepted by the page-alignment
// helpers.
type bytecount interface {
	~uint | ~uint32 | ~uint64 | ~uintptr
}

// hugebytecount constrains the integer types accepted by the huge-page
// helpers.
type hugebytecount interface {
	~uint | ~uint32 | ~uint64 | ~uintptr
}

// PageRoundDown returns x rounded down to the nearest multiple of PageSize.
func PageRoundDown[T bytecount](x T) T {
	return x &^ T(PageMask)
}

// PageRoundUp returns x rounded up to the nearest multiple of PageSize. The ok
// flag is false when rounding up overflows the range of T.
func PageRoundUp[T bytecount](x T) (val T, ok bool) {
	val = PageRoundDown(x + T(PageMask))
	ok = val >= x
	return val, ok
}

// MustPageRoundUp is equivalent to PageRoundUp, but panics if rounding up
// overflows.
func MustPageRoundUp[T bytecount](x T) T {
	val, ok := PageRoundUp(x)
	if !ok {
		panic("abi.PageRoundUp overflows")
	}
	return val
}

// PageOffset returns the offset of x into its containing page.
func PageOffset[T bytecount](x T) T {
	return x & T(PageMask)
}

// IsPageAligned returns true if x is a multiple of PageSize.
func IsPageAligned[T bytecount](x T) bool {
	return PageOffset(x) == 0
}

// ToPagesRoundUp returns the number of pages equal to x bytes rounded up. The
// ok flag is false when rounding x up to a multiple of PageSize overflows T.
func ToPagesRoundUp[T bytecount](x T) (T, bool) {
	y := x + T(PageMask)
	if y < x {
		return x, false
	}
	return y >> PageShift, true
}

// HugePageRoundDown returns x rounded down to the nearest multiple of
// HugePageSize.
func HugePageRoundDown[T hugebytecount](x T) T {
	return x &^ T(HugePageMask)
}

// HugePageRoundUp returns x rounded up to the nearest multiple of
// HugePageSize. The ok flag is false when rounding up overflows T.
func HugePageRoundUp[T hugebytecount](x T) (val T, ok bool) {
	val = HugePageRoundDown(x + T(HugePageMask))
	ok = val >= x
	return val, ok
}

// MustHugePageRoundUp is equivalent to HugePageRoundUp, but panics if rounding
// up overflows.
func MustHugePageRoundUp[T hugebytecount](x T) T {
	val, ok := HugePageRoundUp(x)
	if !ok {
		panic("abi.HugePageRoundUp overflows")
	}
	return val
}

// HugePageOffset returns the offset of x into its containing huge page.
func HugePageOffset[T hugebytecount](x T) T {
	return x & T(HugePageMask)
}

// IsHugePageAligned returns true if x is a multiple of HugePageSize.
func IsHugePageAligned[T hugebytecount](x T) bool {
	return HugePageOffset(x) == 0
}

// CacheLineRoundDown returns x rounded down to the nearest multiple of
// CacheLineSize.
func CacheLineRoundDown[T bytecount](x T) T {
	return x &^ T(CacheLineMask)
}

// CacheLineRoundUp returns x rounded up to the nearest multiple of
// CacheLineSize. The ok flag is false when rounding up overflows T.
func CacheLineRoundUp[T bytecount](x T) (val T, ok bool) {
	val = CacheLineRoundDown(x + T(CacheLineMask))
	ok = val >= x
	return val, ok
}

// MustCacheLineRoundUp is equivalent to CacheLineRoundUp, but panics if
// rounding up overflows.
func MustCacheLineRoundUp[T bytecount](x T) T {
	val, ok := CacheLineRoundUp(x)
	if !ok {
		panic("abi.CacheLineRoundUp overflows")
	}
	return val
}
