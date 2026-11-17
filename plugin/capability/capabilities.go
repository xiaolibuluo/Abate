/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

// Capabilities is a bitmask of granted capabilities. A zero value denies
// everything, enforcing the default-deny security posture.
type Capabilities uint32

// Has reports whether bit is granted.
func (c Capabilities) Has(bit Capability) bool {
	return c&Capabilities(bit) != 0
}

// Grant adds bit to the set. It is idempotent.
func (c *Capabilities) Grant(bit Capability) {
	*c |= Capabilities(bit)
}

// Revoke removes bit from the set. It is idempotent.
func (c *Capabilities) Revoke(bit Capability) {
	*c &^= Capabilities(bit)
}

// IsZero reports whether no capability is granted.
func (c Capabilities) IsZero() bool {
	return c == 0
}
