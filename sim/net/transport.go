/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

// Transport is the virtual channel used by in-sim HTTP and shell probes.
//
// It is a deliberate subset of net.Conn semantics that never performs any real
// socket I/O; every Dial, Read and Write is resolved through the in-memory
// network so the whole stack stays air-gapped and deterministic.
type Transport interface {
	// Dial opens a virtual connection to the named node.
	Dial(addr string) (Conn, error)
}
