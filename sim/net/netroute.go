/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import (
	"container/heap"
	"math"
	"sort"
)

// edge is an adjacency entry used by the shortest-path search.
type edge struct {
	to     string
	weight int64
}

// distItem is a Dijkstra priority-queue entry ordered by (distance, node id).
type distItem struct {
	id   string
	dist int64
}

// distHeap is a min-heap of distance items with deterministic tie-breaking.
type distHeap []distItem

func (h distHeap) Len() int { return len(h) }

func (h distHeap) Less(i, j int) bool {
	if h[i].dist != h[j].dist {
		return h[i].dist < h[j].dist
	}
	return h[i].id < h[j].id
}

func (h distHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *distHeap) Push(x any) { *h = append(*h, x.(distItem)) }

func (h *distHeap) Pop() any {
	old := *h
	n := len(old)
	it := old[n-1]
	*h = old[:n-1]
	return it
}

// dijkstra computes the shortest distance from src to every node.
func (n *Network) dijkstra(src string) map[string]int64 {
	dist := make(map[string]int64, len(n.nodes))
	for id := range n.nodes {
		dist[id] = math.MaxInt64
	}
	dist[src] = 0
	h := &distHeap{{id: src, dist: 0}}
	heap.Init(h)
	for h.Len() > 0 {
		cur, ok := heap.Pop(h).(distItem)
		if !ok {
			continue
		}
		if cur.dist > dist[cur.id] {
			continue
		}
		for _, e := range n.adj(cur.id) {
			nd := cur.dist + e.weight
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(h, distItem{id: e.to, dist: nd})
			}
		}
	}
	return dist
}

// lexReconstruct walks a deterministic lexicographically smallest shortest path.
//
// A visited set guards against zero-weight cycles: a distance-monotonic edge
// over a zero-weight link can otherwise let the walk revisit nodes forever.
func (n *Network) lexReconstruct(dist map[string]int64, from, to string) []string {
	path := []string{from}
	visited := make(map[string]bool, len(dist))
	visited[from] = true
	cur := from
	for cur != to {
		best := ""
		for _, e := range n.adj(cur) {
			if visited[e.to] {
				continue
			}
			if dist[e.to] == dist[cur]+e.weight {
				if best == "" || e.to < best {
					best = e.to
				}
			}
		}
		if best == "" {
			return nil
		}
		path = append(path, best)
		visited[best] = true
		cur = best
	}
	return path
}

// adj returns the sorted adjacency list of a node (deterministic by neighbor id).
func (n *Network) adj(id string) []edge {
	out := make([]edge, 0, 4)
	for key, lk := range n.links {
		a, b := splitKey(key)
		var other string
		switch {
		case a == id:
			other = b
		case b == id:
			other = a
		default:
			continue
		}
		out = append(out, edge{to: other, weight: int64(lk.cfg.BaseLatency)})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].to < out[j].to })
	return out
}

// linkBetween returns the link connecting a and b, regardless of direction.
func (n *Network) linkBetween(a, b string) *Link {
	return n.links[linkKey(a, b)]
}

// deliver records a delivery event and enqueues the packet on the destination.
func (n *Network) deliver(pid PacketID, from, to string, payload []byte, at int64) {
	n.recordEvent("deliver", pid, from, to, to, 0, at)
	node := n.nodes[to]
	node.inbox = append(node.inbox, &Packet{
		ID:        pid,
		From:      from,
		To:        to,
		Payload:   append([]byte(nil), payload...),
		SentAt:    n.clk.Now().UnixNano(),
		DeliverAt: at,
	})
}

// recordEvent appends a totally-ordered event to the network log.
func (n *Network) recordEvent(kind string, pid PacketID, from, to, hop string, lat, at int64) {
	n.eventSeq++
	n.events = append(n.events, netEvent{
		Seq:      n.eventSeq,
		At:       at,
		Kind:     kind,
		PacketID: pid,
		From:     from,
		To:       to,
		Hop:      hop,
		Latency:  lat,
	})
}

// linkKey builds an order-independent key for an undirected link.
func linkKey(a, b string) string {
	if a < b {
		return a + "\x1f" + b
	}
	return b + "\x1f" + a
}

// splitKey splits a link key into its two endpoint ids.
func splitKey(k string) (first, second string) {
	for i := 0; i < len(k); i++ {
		if k[i] == '\x1f' {
			return k[:i], k[i+1:]
		}
	}
	return k, ""
}
