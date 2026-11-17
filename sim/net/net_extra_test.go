/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import (
	"errors"
	"testing"

	"abate/sim/clock"
)

// TestConnReadClosedAndEmpty exercises the virtual connection read path: an
// empty inbox, a successful read, and the closed-state errors.
func TestConnReadClosedAndEmpty(t *testing.T) {
	n, _ := buildLinear()
	tr, err := n.Transport("a")
	if err != nil {
		t.Fatalf("transport err %v", err)
	}
	c, err := tr.Dial("d")
	if err != nil {
		t.Fatalf("dial err %v", err)
	}

	// Reading before any packet is delivered yields ErrNoMessage.
	buf := make([]byte, 16)
	if _, rerr := c.Read(buf); !errors.Is(rerr, ErrNoMessage) {
		t.Fatalf("empty read want ErrNoMessage, got %v", rerr)
	}

	// Deliver a packet from a to d through the writing side.
	if _, werr := c.Write([]byte("ping")); werr != nil {
		t.Fatalf("write err %v", werr)
	}
	trd, err := n.Transport("d")
	if err != nil {
		t.Fatalf("transport d err %v", err)
	}
	cd, err := trd.Dial("d")
	if err != nil {
		t.Fatalf("dial d err %v", err)
	}
	nn, rerr := cd.Read(buf)
	if rerr != nil || string(buf[:nn]) != "ping" {
		t.Fatalf("read err=%v data=%q", rerr, buf[:nn])
	}

	// Both endpoints refuse I/O after close.
	if cerr := c.Close(); cerr != nil {
		t.Fatalf("close err %v", cerr)
	}
	if _, rerr := c.Read(buf); !errors.Is(rerr, ErrConnClosed) {
		t.Fatalf("read after close want ErrConnClosed, got %v", rerr)
	}
	if _, werr := c.Write([]byte("x")); !errors.Is(werr, ErrConnClosed) {
		t.Fatalf("write after close want ErrConnClosed, got %v", werr)
	}
}

// TestTransportUnknownNode covers the error branches of Transport and Dial.
func TestTransportUnknownNode(t *testing.T) {
	n, _ := buildLinear()
	if _, err := n.Transport("missing"); !errors.Is(err, ErrNoNode) {
		t.Fatalf("transport missing want ErrNoNode, got %v", err)
	}
	tr, err := n.Transport("a")
	if err != nil {
		t.Fatalf("transport a err %v", err)
	}
	if _, err := tr.Dial("missing"); !errors.Is(err, ErrNoNode) {
		t.Fatalf("dial missing want ErrNoNode, got %v", err)
	}
}

// TestLinkEndpoints covers the A and B endpoint accessors.
func TestLinkEndpoints(t *testing.T) {
	n, _ := buildLinear()
	lk := n.linkBetween("b", "c")
	if lk.A() != "b" || lk.B() != "c" {
		t.Fatalf("endpoints wrong: %q %q", lk.A(), lk.B())
	}
}

// TestRouteCacheHit confirms a repeated route query hits the cache path.
func TestRouteCacheHit(t *testing.T) {
	n, _ := buildLinear()
	r1, err := n.Route("a", "d")
	if err != nil {
		t.Fatalf("route err %v", err)
	}
	r2, err := n.Route("a", "d")
	if err != nil {
		t.Fatalf("route cached err %v", err)
	}
	if len(r1) != len(r2) {
		t.Fatalf("cache mismatch: %v vs %v", r1, r2)
	}
}

// TestInjectBandwidthFault covers the remaining InjectFault branch and the
// RecoverFault error branch.
func TestInjectBandwidthFault(t *testing.T) {
	n, _ := buildLinear()
	if err := n.InjectFault(Fault{Kind: FaultBandwidth, A: "b", B: "c", Value: 1 << 30}); err != nil {
		t.Fatalf("inject bandwidth err %v", err)
	}
	if got := n.linkBetween("b", "c").Config().BandwidthBps; got != 1<<30 {
		t.Fatalf("bandwidth not injected: %d", got)
	}
	if err := n.RecoverFault(Fault{A: "x", B: "y"}); !errors.Is(err, ErrNoLink) {
		t.Fatalf("recover unknown want ErrNoLink, got %v", err)
	}
}

// TestLinkKeyEdge covers both orderings of linkKey and both splitKey branches.
func TestLinkKeyEdge(t *testing.T) {
	if linkKey("y", "x") != linkKey("x", "y") {
		t.Fatalf("linkKey not order independent")
	}
	a, b := splitKey(linkKey("x", "y"))
	if a != "x" || b != "y" {
		t.Fatalf("splitKey wrong: %q %q", a, b)
	}
	only, other := splitKey("lonely")
	if only != "lonely" || other != "" {
		t.Fatalf("splitKey no-separator wrong: %q %q", only, other)
	}
}

// TestNodeRecvEmpty covers the empty-inbox branch of Node.recv.
func TestNodeRecvEmpty(t *testing.T) {
	n, _ := buildLinear()
	if _, ok := n.nodes["a"].recv(); ok {
		t.Fatalf("empty node recv should report false")
	}
}

// TestSendNoRoute confirms Send returns ErrNoRoute when endpoints are valid but
// unreachable, exercising the route-error branch of Send.
func TestSendNoRoute(t *testing.T) {
	clk := clock.NewVirtualClock()
	n := NewNetworkWithSeed(clk, 7)
	n.AddNode("p")
	n.AddNode("q")
	if _, err := n.Send("p", "q", []byte("x")); !errors.Is(err, ErrNoRoute) {
		t.Fatalf("want ErrNoRoute, got %v", err)
	}
}
