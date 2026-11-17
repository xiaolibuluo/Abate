/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

import (
	"context"
	"crypto/ed25519"
	"errors"
	"runtime"
	"testing"
	"time"

	"abate/plugin/capability"
	"abate/plugin/host"
	"abate/plugin/sample"
	"abate/plugin/signature"
)

// swapFixture owns a signed pair of instances registered on one host. The new
// instance is loaded first (to prove the factory accepts it) and the registry is
// then rolled back to the old instance, reproducing the state immediately before
// a hot-swap.
type swapFixture struct {
	store *signature.TrustStore
	priv  ed25519.PrivateKey
	h     *host.Host
	old   *host.Instance
	next  *host.Instance
	id    signature.PluginID
}

// allCaps grants every known capability to a loaded instance.
func allCaps() capability.Capabilities {
	var c capability.Capabilities
	for _, bit := range []capability.Capability{
		capability.CapClock, capability.CapRand, capability.CapObserve,
		capability.CapFS, capability.CapNet,
	} {
		c.Grant(bit)
	}
	return c
}

// factoryFor maps a plugin type to its reference factory.
func factoryFor(t host.PluginType) host.Factory {
	switch t {
	case host.PluginTypeReporter:
		return sample.ReporterFactory
	case host.PluginTypeAdapter:
		return sample.AdapterFactory
	default:
		return sample.OperatorFactory
	}
}

// newSignedManifest builds and signs a manifest under the fixture anchor.
func newSignedManifest(t *testing.T, f *swapFixture, v signature.SemVer, typ signature.PluginType) signature.Manifest {
	t.Helper()
	m := signature.Manifest{
		ID:               f.id,
		Version:          v,
		Type:             typ,
		CapabilitiesDecl: "clock,rand,observe",
		TrustAnchor:      "hotswap-anchor",
		BuildRef:         "build-hotswap",
	}
	if err := signature.Sign(f.store, f.priv, &m); err != nil {
		t.Fatalf("Sign: %v", err)
	}
	return m
}

// newSwapFixture constructs a signed old/new pair with the requested versions.
func newSwapFixture(t *testing.T, from, to signature.SemVer, typ signature.PluginType) *swapFixture {
	t.Helper()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	store := signature.NewTrustStore()
	if err := store.Add("hotswap-anchor", pub); err != nil {
		t.Fatalf("store.Add: %v", err)
	}
	f := &swapFixture{
		store: store,
		priv:  priv,
		h:     host.NewHost(store),
		id:    signature.PluginID("sample/hotswap"),
	}
	mOld := newSignedManifest(t, f, from, typ)
	mNext := newSignedManifest(t, f, to, typ)
	f.old, err = f.h.Load(mOld, factoryFor(typ), allCaps())
	if err != nil {
		t.Fatalf("Load old: %v", err)
	}
	f.next, err = f.h.Load(mNext, factoryFor(typ), allCaps())
	if err != nil {
		t.Fatalf("Load next: %v", err)
	}
	f.h.Registry().Replace(f.id, f.old)
	return f
}

// waitUntil polls a condition with a bounded number of scheduler yields. It
// avoids wall-clock sleeps while still letting a swapping goroutine progress.
func waitUntil(t *testing.T, cond func() bool) {
	t.Helper()
	for i := 0; i < 2000000; i++ {
		if cond() {
			return
		}
		runtime.Gosched()
	}
	t.Fatal("condition never became true")
}

func ver(major, minor, patch uint) signature.SemVer {
	return signature.SemVer{Major: major, Minor: minor, Patch: patch}
}

func TestSwapStateNames(t *testing.T) {
	cases := []struct {
		state SwapState
		want  string
	}{
		{PREPARE, "PREPARE"},
		{SWITCH, "SWITCH"},
		{DRAIN, "DRAIN"},
		{UNLOAD, "UNLOAD"},
		{SwapState(99), "UNKNOWN"},
	}
	for _, c := range cases {
		if got := c.state.String(); got != c.want {
			t.Fatalf("state %d = %q, want %q", c.state, got, c.want)
		}
	}
	if PREPARE != 0 || SWITCH != 1 || DRAIN != 2 || UNLOAD != 3 {
		t.Fatal("frozen phase ordering changed")
	}
}

func TestAllowedPolicy(t *testing.T) {
	cases := []struct {
		name string
		typ  host.PluginType
		from signature.SemVer
		to   signature.SemVer
		want bool
	}{
		{"patch-upgrade", host.PluginTypeOperator, ver(1, 0, 0), ver(1, 0, 1), true},
		{"minor-upgrade", host.PluginTypeOperator, ver(1, 0, 0), ver(1, 2, 0), true},
		{"identical", host.PluginTypeOperator, ver(1, 2, 3), ver(1, 2, 3), true},
		{"reporter", host.PluginTypeReporter, ver(2, 0, 0), ver(2, 0, 1), true},
		{"adapter", host.PluginTypeAdapter, ver(3, 1, 0), ver(3, 2, 0), true},
		{"patch-downgrade", host.PluginTypeOperator, ver(1, 0, 2), ver(1, 0, 1), false},
		{"minor-downgrade", host.PluginTypeOperator, ver(1, 2, 0), ver(1, 1, 9), false},
		{"major-upgrade", host.PluginTypeOperator, ver(1, 0, 0), ver(2, 0, 0), false},
		{"major-downgrade", host.PluginTypeOperator, ver(2, 0, 0), ver(1, 9, 9), false},
		{"unknown-type", host.PluginType("sandbox"), ver(1, 0, 0), ver(1, 0, 1), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Allowed(c.typ, c.from, c.to); got != c.want {
				t.Fatalf("Allowed(%q,%v,%v) = %v, want %v", c.typ, c.from, c.to, got, c.want)
			}
		})
	}
}

func TestBeginFullLifecycle(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	sw, err := Begin(context.Background(), f.old, f.next)
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	if sw.State() != UNLOAD {
		t.Fatalf("state = %s, want UNLOAD", sw.State())
	}
	if !f.old.IsSuperseded() {
		t.Fatal("old instance was not superseded")
	}
	if got, ok := f.h.Lookup(f.id); !ok || got != f.next {
		t.Fatal("registry does not serve the new instance")
	}
	wants := []capability.Capability{capability.CapRand}
	out, err := f.h.Call(f.next, host.EntryApply, []byte("after-swap"), wants)
	if err != nil {
		t.Fatalf("new instance call: %v", err)
	}
	if len(out) != len("after-swap")+1 {
		t.Fatalf("new instance output length = %d", len(out))
	}
	if _, err := f.h.Call(f.old, host.EntryApply, []byte("stale"), wants); !errors.Is(err, host.ErrInstanceStale) {
		t.Fatalf("stale instance call = %v, want ErrInstanceStale", err)
	}
}

func TestInFlightCallDrainsAndNewInstanceServes(t *testing.T) {
	f := newSwapFixture(t, ver(1, 2, 0), ver(1, 2, 1), host.PluginTypeOperator)
	release, ok := f.old.BeginCall()
	if !ok {
		t.Fatal("could not start an in-flight call on the old instance")
	}
	started := make(chan struct{})
	finish := make(chan struct{})
	go func() {
		close(started)
		<-finish
		release()
	}()
	<-started

	type result struct {
		sw  *Swap
		err error
	}
	resCh := make(chan result, 1)
	go func() {
		sw, err := Begin(context.Background(), f.old, f.next)
		resCh <- result{sw: sw, err: err}
	}()

	waitUntil(t, func() bool {
		got, found := f.h.Lookup(f.id)
		return found && got == f.next
	})

	wants := []capability.Capability{capability.CapRand}
	if _, err := f.h.Call(f.next, host.EntryApply, []byte("during-drain"), wants); err != nil {
		t.Fatalf("new instance must serve during DRAIN: %v", err)
	}
	if _, err := f.h.Call(f.old, host.EntryApply, []byte("rejected"), wants); !errors.Is(err, host.ErrInstanceStale) {
		t.Fatalf("old instance must reject new calls: %v", err)
	}
	close(finish)
	res := <-resCh
	if res.err != nil {
		t.Fatalf("Begin while draining: %v", res.err)
	}
	if res.sw.State() != UNLOAD {
		t.Fatalf("state = %s, want UNLOAD", res.sw.State())
	}
}

func TestBeginDrainTimeout(t *testing.T) {
	previous := DrainTimeout
	DrainTimeout = 15 * time.Millisecond
	t.Cleanup(func() { DrainTimeout = previous })
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 1, 0), host.PluginTypeOperator)
	release, ok := f.old.BeginCall()
	if !ok {
		t.Fatal("could not hold an in-flight call")
	}
	defer release()
	sw, err := Begin(context.Background(), f.old, f.next)
	if !errors.Is(err, ErrDrainTimeout) {
		t.Fatalf("Begin = %v, want ErrDrainTimeout", err)
	}
	if sw.State() != DRAIN {
		t.Fatalf("state = %s, want DRAIN", sw.State())
	}
	if got, found := f.h.Lookup(f.id); !found || got != f.next {
		t.Fatal("new instance stopped serving after a drain timeout")
	}
}

func TestBeginContextCancellation(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 2), host.PluginTypeOperator)
	release, ok := f.old.BeginCall()
	if !ok {
		t.Fatal("could not hold an in-flight call")
	}
	defer release()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sw, err := Begin(ctx, f.old, f.next)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Begin = %v, want context.Canceled", err)
	}
	if sw.State() != DRAIN {
		t.Fatalf("state = %s, want DRAIN", sw.State())
	}
}

func TestBeginNilContextUsesBackground(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	var noContext context.Context
	sw, err := Begin(noContext, f.old, f.next)
	if err != nil {
		t.Fatalf("Begin with nil context: %v", err)
	}
	if sw.State() != UNLOAD {
		t.Fatalf("state = %s, want UNLOAD", sw.State())
	}
}
