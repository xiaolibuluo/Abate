/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

// Conn is a virtual connection between two nodes of the simulated network.
//
// It mirrors the read/write/close subset of net.Conn but resolves entirely
// through the in-memory network, so probes can run with zero real socket I/O.
type Conn interface {
	// Read copies the next delivered packet payload into p.
	Read(p []byte) (int, error)

	// Write sends p as a packet from the local node to the remote node.
	Write(p []byte) (int, error)

	// Close marks the connection closed.
	Close() error
}

// transport is the in-memory Transport implementation.
type transport struct {
	net   *Network
	local string
}

// Dial opens a virtual connection to the named remote node.
func (t *transport) Dial(addr string) (Conn, error) {
	if _, ok := t.net.nodes[addr]; !ok {
		return nil, ErrNoNode
	}
	return &conn{net: t.net, local: t.local, remote: addr}, nil
}

// conn is the in-memory Conn implementation.
type conn struct {
	net    *Network
	local  string
	remote string
	closed bool
}

// Write sends the payload from the local node to the remote node.
func (c *conn) Write(p []byte) (int, error) {
	if c.closed {
		return 0, ErrConnClosed
	}
	if _, err := c.net.Send(c.local, c.remote, p); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Read copies the next delivered packet destined to the local node.
func (c *conn) Read(p []byte) (int, error) {
	if c.closed {
		return 0, ErrConnClosed
	}
	pkt, ok := c.net.nodes[c.local].recv()
	if !ok {
		return 0, ErrNoMessage
	}
	return copy(p, pkt.Payload), nil
}

// Close marks the connection closed.
func (c *conn) Close() error {
	c.closed = true
	return nil
}
