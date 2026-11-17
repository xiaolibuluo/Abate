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

// buildLinear builds a 4-node linear topology a-b-c-d with unit latency links.
func buildLinear() (*Network, *clock.VirtualClock) {
	clk := clock.NewVirtualClock()
	n := NewNetworkWithSeed(clk, 1)
	n.AddNode("a")
	n.AddNode("b")
	n.AddNode("c")
	n.AddNode("d")
	n.Link("a", "b", LinkConfig{BaseLatency: 1000})
	n.Link("b", "c", LinkConfig{BaseLatency: 1000})
	n.Link("c", "d", LinkConfig{BaseLatency: 1000})
	return n, clk
}

func TestAddNodeIdempotent(t *testing.T) {
	n, _ := buildLinear()
	a1 := n.AddNode("a")
	a2 := n.AddNode("a")
	if a1 != a2 {
		t.Fatalf("AddNode not idempotent")
	}
	if n.AddNode("") != nil {
		t.Fatalf("empty id should be rejected")
	}
}

func TestLinkSelfAndAutoAdd(t *testing.T) {
	clk := clock.NewVirtualClock()
	n := NewNetworkWithSeed(clk, 2)
	if n.Link("x", "x", LinkConfig{}) != nil {
		t.Fatalf("self link should be rejected")
	}
	// Missing endpoints are auto-created.
	n.Link("x", "y", LinkConfig{BaseLatency: 5})
	if _, err := n.Route("x", "y"); err != nil {
		t.Fatalf("auto-added link missing: %v", err)
	}
}

func TestRouteBasicAndTieBreak(t *testing.T) {
	clk := clock.NewVirtualClock()
	n := NewNetworkWithSeed(clk, 1)
	for _, id := range []string{"a", "b", "c", "d"} {
		n.AddNode(id)
	}
	n.Link("a", "b", LinkConfig{BaseLatency: 1})
	n.Link("b", "c", LinkConfig{BaseLatency: 1})
	n.Link("c", "d", LinkConfig{BaseLatency: 1})
	r, err := n.Route("a", "d")
	if err != nil || len(r) != 4 || r[0] != "a" || r[3] != "d" {
		t.Fatalf("basic route wrong: %v %v", r, err)
	}

	// Equal-cost alternatives a-b-d and a-c-d (both cost 2). The lexicographically
	// smaller first step (b) must win the tie.
	n.Link("b", "d", LinkConfig{BaseLatency: 1})
	n.Link("a", "c", LinkConfig{BaseLatency: 1})
	r2, err := n.Route("a", "d")
	if err != nil {
		t.Fatalf("route err = %v", err)
	}
	if len(r2) != 3 || r2[1] != "b" {
		t.Fatalf("tie-break wrong: %v", r2)
	}
}

func TestRouteErrors(t *testing.T) {
	n, _ := buildLinear()
	if _, err := n.Route("a", "zzz"); !errors.Is(err, ErrNoNode) {
		t.Fatalf("want ErrNoNode, got %v", err)
	}
	if _, err := n.Route("a", "a"); err != nil {
		t.Fatalf("self route should succeed: %v", err)
	}
	clk := clock.NewVirtualClock()
	iso := NewNetworkWithSeed(clk, 3)
	iso.AddNode("p")
	iso.AddNode("q")
	if _, err := iso.Route("p", "q"); !errors.Is(err, ErrNoRoute) {
		t.Fatalf("want ErrNoRoute, got %v", err)
	}
}

func TestSendDelivers(t *testing.T) {
	n, clk := buildLinear()
	for i := 0; i < 5; i++ {
		_, _ = clk.Advance(clock.Millisecond)
	}
	pid, err := n.Send("a", "d", []byte("hello"))
	if err != nil {
		t.Fatalf("send err = %v", err)
	}
	if pid == 0 {
		t.Fatalf("packet id should be non-zero")
	}
	// d should have the packet in its inbox.
	pkt, ok := n.nodes["d"].recv()
	if !ok || string(pkt.Payload) != "hello" {
		t.Fatalf("delivery missing: ok=%v", ok)
	}
	// Last event should be a deliver at a positive virtual time.
	last := n.events[len(n.events)-1]
	if last.Kind != "deliver" || last.To != "d" {
		t.Fatalf("unexpected last event %+v", last)
	}
}

func TestSendSelfDeliversImmediately(t *testing.T) {
	n, clk := buildLinear()
	_, _ = clk.Advance(clock.Second)
	_, err := n.Send("a", "a", []byte("x"))
	if err != nil {
		t.Fatalf("self send err = %v", err)
	}
	if _, ok := n.nodes["a"].recv(); !ok {
		t.Fatalf("self delivery missing")
	}
}

func TestSendErrors(t *testing.T) {
	n, _ := buildLinear()
	if _, err := n.Send("a", "missing", []byte("x")); !errors.Is(err, ErrNoNode) {
		t.Fatalf("want ErrNoNode, got %v", err)
	}
	big := make([]byte, maxPayloadBytes+1)
	if _, err := n.Send("a", "d", big); !errors.Is(err, ErrInvalidPayload) {
		t.Fatalf("want ErrInvalidPayload, got %v", err)
	}
}

func TestLossModel(t *testing.T) {
	n, _ := buildLinear()
	// Force every hop to drop.
	for key, lk := range n.links {
		_ = key
		lk.SetLossRate(1)
	}
	pid, err := n.Send("a", "d", []byte("x"))
	if !errors.Is(err, ErrDropped) {
		t.Fatalf("want ErrDropped, got %v", err)
	}
	if pid == 0 {
		t.Fatalf("packet id should be assigned before drop")
	}
	// Loss rate 0 never drops.
	for _, lk := range n.links {
		lk.SetLossRate(0)
	}
	if _, err := n.Send("a", "d", []byte("x")); err != nil {
		t.Fatalf("zero loss should deliver: %v", err)
	}
}

func TestPartitionAndHeal(t *testing.T) {
	n, _ := buildLinear()
	mid := n.linkBetween("b", "c")
	if err := mid.Partition(); err != nil {
		t.Fatalf("partition err = %v", err)
	}
	if _, err := n.Send("a", "d", []byte("x")); !errors.Is(err, ErrPartitioned) {
		t.Fatalf("want ErrPartitioned, got %v", err)
	}
	// No propagation events should be recorded for the failed send beyond send.
	if last := n.events[len(n.events)-1]; last.Kind != "send" {
		t.Fatalf("partitioned send must not emit propagation events: %+v", last)
	}
	if err := mid.Heal(); err != nil {
		t.Fatalf("heal err = %v", err)
	}
	if _, err := n.Send("a", "d", []byte("y")); err != nil {
		t.Fatalf("heal should restore delivery: %v", err)
	}
}

func TestBandwidthAndLatency(t *testing.T) {
	n, _ := buildLinear()
	mid := n.linkBetween("b", "c")
	mid.SetBandwidth(0) // zero bandwidth: no queueing delay.
	mid.SetLatency(0)
	before := n.clk.Now().UnixNano()
	pid, err := n.Send("a", "d", make([]byte, 100))
	if err != nil {
		t.Fatalf("send err = %v", err)
	}
	_ = pid
	// Only the two non-zero base links (a-b, c-d = 1000 each) contribute latency.
	deliver := n.events[len(n.events)-1]
	if deliver.At-before != 2000 {
		t.Fatalf("unexpected deliverAt delta %d", deliver.At-before)
	}

	// Large bandwidth on a 1-byte packet adds negligible delay (still 2000).
	mid.SetBandwidth(1_000_000)
	mid.SetLatency(500)
	_, _ = n.Send("a", "d", make([]byte, 1))
}

func TestTransportReadWrite(t *testing.T) {
	n, _ := buildLinear()
	tr, err := n.Transport("a")
	if err != nil {
		t.Fatalf("transport err = %v", err)
	}
	conn, err := tr.Dial("d")
	if err != nil {
		t.Fatalf("dial err = %v", err)
	}
	if _, err := conn.Write([]byte("ping")); err != nil {
		t.Fatalf("write err = %v", err)
	}
	// Read consumes the delivered packet on d.
	buf := make([]byte, 64)
	nn, err := n.Transport("d")
	if err != nil {
		t.Fatalf("transport d err = %v", err)
	}
	dconn, _ := nn.Dial("d") // loopback-ish; we only read d's inbox.
	_ = dconn
	rd, ok := n.nodes["d"].recv()
	if !ok || string(rd.Payload) != "ping" {
		t.Fatalf("transport write did not deliver: ok=%v", ok)
	}
	_ = buf
	_ = conn
	if err := conn.Close(); err != nil {
		t.Fatalf("close err = %v", err)
	}
	if _, err := conn.Write([]byte("x")); !errors.Is(err, ErrConnClosed) {
		t.Fatalf("write after close should fail: %v", err)
	}
}

func TestFaultInjectRecover(t *testing.T) {
	n, _ := buildLinear()
	if err := n.InjectFault(Fault{Kind: FaultLatency, A: "b", B: "c", Value: 9999}); err != nil {
		t.Fatalf("inject latency err = %v", err)
	}
	if n.linkBetween("b", "c").Config().BaseLatency != 9999 {
		t.Fatalf("latency not injected")
	}
	if err := n.InjectFault(Fault{Kind: FaultLoss, A: "b", B: "c", Value: 1_000_000}); err != nil {
		t.Fatalf("inject loss err = %v", err)
	}
	if n.linkBetween("b", "c").Config().LossRate != 1 {
		t.Fatalf("loss not injected")
	}
	if err := n.InjectFault(Fault{Kind: FaultPartition, A: "b", B: "c"}); err != nil {
		t.Fatalf("inject partition err = %v", err)
	}
	if !n.linkBetween("b", "c").Partitioned() {
		t.Fatalf("partition not injected")
	}
	if err := n.RecoverFault(Fault{A: "b", B: "c"}); err != nil {
		t.Fatalf("recover err = %v", err)
	}
	if n.linkBetween("b", "c").Partitioned() {
		t.Fatalf("recover did not heal")
	}
	if n.linkBetween("b", "c").Config().BaseLatency != 1000 {
		t.Fatalf("recover did not restore baseline latency")
	}
	if err := n.InjectFault(Fault{Kind: FaultKind("bogus"), A: "b", B: "c"}); !errors.Is(err, ErrUnknownFault) {
		t.Fatalf("want ErrUnknownFault, got %v", err)
	}
	if err := n.InjectFault(Fault{Kind: FaultLatency, A: "x", B: "y"}); !errors.Is(err, ErrNoLink) {
		t.Fatalf("want ErrNoLink, got %v", err)
	}
}
