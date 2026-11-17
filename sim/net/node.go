/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

// Node is a virtual network endpoint that buffers delivered packets in its inbox.
type Node struct {
	id    string
	inbox []*Packet
}

// recv pops the oldest delivered packet from the node inbox.
func (nd *Node) recv() (*Packet, bool) {
	if len(nd.inbox) == 0 {
		return nil, false
	}
	p := nd.inbox[0]
	nd.inbox = nd.inbox[1:]
	return p, true
}
