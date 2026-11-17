/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import (
	"errors"
	"math"

	"abate/sim/clock"
	"abate/sim/rand"
)

// Sentinel errors of the virtual network.
var (
	// ErrNoNode reports an operation referencing an unknown node.
	ErrNoNode = errors.New("net: unknown node")

	// ErrNoRoute reports that no path exists between two nodes.
	ErrNoRoute = errors.New("net: no route between nodes")

	// ErrNoLink reports a fault reference to a non-existent link.
	ErrNoLink = errors.New("net: unknown link")

	// ErrUnknownFault reports an unrecognized fault kind.
	ErrUnknownFault = errors.New("net: unknown fault kind")

	// ErrPartitioned reports that a send traversed a partitioned link.
	ErrPartitioned = errors.New("net: link is partitioned")

	// ErrDropped reports that a packet was dropped by the loss model.
	ErrDropped = errors.New("net: packet dropped by loss")

	// ErrInvalidPayload reports a payload larger than the 1MiB cap.
	ErrInvalidPayload = errors.New("net: payload exceeds 1MiB")

	// ErrConnClosed reports I/O on a closed virtual connection.
	ErrConnClosed = errors.New("net: connection closed")

	// ErrNoMessage reports that no packet is pending on a node inbox.
	ErrNoMessage = errors.New("net: no pending message")
)

// maxPayloadBytes is the virtual per-packet payload cap.
const maxPayloadBytes = 1 << 20

// netEvent is one totally-ordered record of the simulated network stream.
type netEvent struct {
	Seq      uint64
	At       int64
	Kind     string
	PacketID PacketID
	From     string
	To       string
	Hop      string
	Latency  int64
}

// Network is a deterministic in-memory model of a virtual network.
//
// Every propagation, queueing, loss and latency decision is resolved through the
// injected clock and random stream, so the event log it records is reproducible.
type Network struct {
	clk        *clock.VirtualClock
	rng        *rand.Stream
	nodes      map[string]*Node
	links      map[string]*Link
	eventSeq   uint64
	packetSeq  PacketID
	events     []netEvent
	routeCache map[string][]string
}

// NewNetwork builds an empty network driven by the given clock and stream.
func NewNetwork(clk *clock.VirtualClock, rng *rand.Stream) *Network {
	return &Network{
		clk:        clk,
		rng:        rng,
		nodes:      make(map[string]*Node),
		links:      make(map[string]*Link),
		routeCache: make(map[string][]string),
	}
}

// NewNetworkWithSeed builds an empty network with a seeded random stream.
func NewNetworkWithSeed(clk *clock.VirtualClock, seed uint64) *Network {
	return NewNetwork(clk, rand.New(seed))
}

// AddNode registers a node, returning the existing node on duplicate ids.
func (n *Network) AddNode(id string) *Node {
	if id == "" {
		return nil
	}
	if nd, ok := n.nodes[id]; ok {
		return nd
	}
	nd := &Node{id: id}
	n.nodes[id] = nd
	n.routeCache = make(map[string][]string)
	return nd
}

// Link connects two nodes with the given configuration.
func (n *Network) Link(a, b string, cfg LinkConfig) *Link {
	if a == "" || b == "" || a == b {
		return nil
	}
	n.AddNode(a)
	n.AddNode(b)
	lk := &Link{a: a, b: b, cfg: cfg, base: cfg}
	n.links[linkKey(a, b)] = lk
	n.routeCache = make(map[string][]string)
	return lk
}

// Route returns the lexicographically smallest shortest path between two nodes.
//
// The shortest path is by link base latency; ties are broken by preferring the
// smaller neighbor id at every step, which makes the result deterministic.
func (n *Network) Route(from, to string) ([]string, error) {
	if _, ok := n.nodes[from]; !ok {
		return nil, ErrNoNode
	}
	if _, ok := n.nodes[to]; !ok {
		return nil, ErrNoNode
	}
	key := from + "\x00" + to
	if r, ok := n.routeCache[key]; ok {
		return append([]string(nil), r...), nil
	}
	if from == to {
		n.routeCache[key] = []string{from}
		return []string{from}, nil
	}
	dist := n.dijkstra(from)
	if d, ok := dist[to]; !ok || d == math.MaxInt64 {
		return nil, ErrNoRoute
	}
	path := n.lexReconstruct(dist, from, to)
	if path == nil {
		return nil, ErrNoRoute
	}
	n.routeCache[key] = append([]string(nil), path...)
	return append([]string(nil), path...), nil
}

// Send transmits a payload from one node to another, recording every send,
// deliver and drop in the event log. It returns the packet id or an error when
// the packet cannot be delivered.
func (n *Network) Send(from, to string, payload []byte) (PacketID, error) {
	if len(payload) > maxPayloadBytes {
		return 0, ErrInvalidPayload
	}
	if _, ok := n.nodes[from]; !ok {
		return 0, ErrNoNode
	}
	if _, ok := n.nodes[to]; !ok {
		return 0, ErrNoNode
	}
	n.packetSeq++
	pid := n.packetSeq
	sentAt := n.clk.Now().UnixNano()
	n.recordEvent("send", pid, from, to, "", 0, sentAt)
	if from == to {
		n.deliver(pid, from, to, payload, sentAt)
		return pid, nil
	}
	route, err := n.Route(from, to)
	if err != nil {
		return pid, ErrNoRoute
	}
	var total int64
	for i := 0; i < len(route)-1; i++ {
		a, b := route[i], route[i+1]
		lk := n.linkBetween(a, b)
		if lk == nil {
			return pid, ErrNoRoute
		}
		if lk.partitioned {
			return pid, ErrPartitioned
		}
		if n.rng.Float64() < lk.cfg.LossRate {
			n.recordEvent("drop", pid, a, b, b, int64(lk.cfg.BaseLatency), sentAt+total)
			return pid, ErrDropped
		}
		hop := int64(lk.cfg.BaseLatency)
		if lk.cfg.Jitter > 0 {
			hop += int64(n.rng.IntN(int(lk.cfg.Jitter))) //nolint:gosec // jitter is a non-negative duration.
		}
		if lk.cfg.BandwidthBps > 0 {
			arrival := sentAt + total + hop
			_, finish := lk.bucket.reserve(arrival, len(payload), pid, lk.cfg.BandwidthBps)
			hop += finish - arrival
		}
		total += hop
	}
	n.deliver(pid, from, to, payload, sentAt+total)
	return pid, nil
}

// Transport returns a virtual transport bound to the given local node.
func (n *Network) Transport(local string) (Transport, error) {
	if _, ok := n.nodes[local]; !ok {
		return nil, ErrNoNode
	}
	return &transport{net: n, local: local}, nil
}
