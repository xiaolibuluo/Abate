/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package abi

import "fmt"

// AddrRange is a range of addresses delimited by a start (inclusive) and an
// end (exclusive).
//
// Adapted from gvisor/pkg/hostarch/addr.go, Apache-2.0.
type AddrRange struct {
	// Start is the inclusive start of the range.
	Start Addr

	// End is the exclusive end of the range.
	End Addr
}

// IsPageAligned returns true if both Start and End are page-aligned.
func (ar AddrRange) IsPageAligned() bool {
	return ar.Start.IsPageAligned() && ar.End.IsPageAligned()
}

// IsHugePageAligned returns true if both Start and End are huge-page-aligned.
func (ar AddrRange) IsHugePageAligned() bool {
	return IsHugePageAligned(ar.Start) && IsHugePageAligned(ar.End)
}

// String implements fmt.Stringer.
func (ar AddrRange) String() string {
	return fmt.Sprintf("[%#x, %#x)", ar.Start, ar.End)
}
