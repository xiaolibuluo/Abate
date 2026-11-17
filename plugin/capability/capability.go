/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

// Capability is a single bit in the capability bitmask.
type Capability uint32

// The five known capabilities. Each is a distinct bit so they compose via the
// bitmask. Order is fixed by iota and must not change (it is part of the
// on-disk/over-the-wire contract).
const (
	// CapClock grants virtual-clock reads.
	CapClock Capability = 1 << iota

	// CapRand grants access to the deterministic random stream.
	CapRand

	// CapObserve grants observation/report emission.
	CapObserve

	// CapFS grants access to the plugin's sandboxed data directory.
	CapFS

	// CapNet grants access to the declared network allow-list.
	CapNet
)

// String returns the canonical lower-case name of the capability, or "cap(?)" for
// any value outside the known set.
func (c Capability) String() string {
	switch c {
	case CapClock:
		return "clock"
	case CapRand:
		return "rand"
	case CapObserve:
		return "observe"
	case CapFS:
		return "fs"
	case CapNet:
		return "net"
	default:
		return "cap(?)"
	}
}
