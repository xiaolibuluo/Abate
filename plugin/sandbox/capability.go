/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sandbox

import (
	"slices"
	"strings"
)

// Capability is a bitmask of sandbox capabilities.
//
// The four capabilities mirror the four resource families a plugin can reach:
// the filesystem, the network, the virtual clock and the environment. A zero
// value grants nothing, which keeps the default deny posture of the sandbox.
type Capability uint32

// The known sandbox capabilities. The bit order is part of the serialized
// contract and must not change.
const (
	// CapFS grants access to the declared filesystem allow list.
	CapFS Capability = 1 << iota

	// CapNet grants access to the declared network allow list.
	CapNet

	// CapClock grants access to the virtual clock.
	CapClock

	// CapEnv grants access to the declared environment variables.
	CapEnv
)

// capabilityNames holds the canonical name of every known capability.
var capabilityNames = []struct {
	bit  Capability
	name string
}{
	{CapFS, "fs"},
	{CapNet, "net"},
	{CapClock, "clock"},
	{CapEnv, "env"},
}

// Has reports whether one capability bit is granted.
func (c Capability) Has(bit Capability) bool {
	return c&bit != 0
}

// Grant adds one capability bit. It is idempotent.
func (c *Capability) Grant(bit Capability) {
	*c |= bit
}

// Revoke removes one capability bit. It is idempotent.
func (c *Capability) Revoke(bit Capability) {
	*c &^= bit
}

// IsZero reports whether no capability is granted.
func (c Capability) IsZero() bool {
	return c == 0
}

// Names returns the names of the granted capabilities in canonical order.
func (c Capability) Names() []string {
	out := make([]string, 0, len(capabilityNames))
	for _, entry := range capabilityNames {
		if c.Has(entry.bit) {
			out = append(out, entry.name)
		}
	}
	return out
}

// Missing returns the capabilities that are requested but not granted.
func (c Capability) Missing(granted Capability) []string {
	out := make([]string, 0, len(capabilityNames))
	for _, entry := range capabilityNames {
		if c.Has(entry.bit) && !granted.Has(entry.bit) {
			out = append(out, entry.name)
		}
	}
	slices.Sort(out)
	return out
}

// String renders the capability mask as a comma separated name list.
func (c Capability) String() string {
	return strings.Join(c.Names(), ",")
}
