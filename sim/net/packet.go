/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

// Packet is a virtual packet delivered through the simulated network.
type Packet struct {
	// ID is the global packet sequence number assigned at send time.
	ID PacketID

	// From is the source node id.
	From string

	// To is the destination node id.
	To string

	// Payload is the application payload (capped at 1MiB).
	Payload []byte

	// SentAt is the virtual send time in nanoseconds.
	SentAt int64

	// DeliverAt is the virtual delivery time in nanoseconds.
	DeliverAt int64
}
