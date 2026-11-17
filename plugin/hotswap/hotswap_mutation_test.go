/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

import (
	"context"
	"errors"
	"testing"

	"abate/plugin/host"
	"abate/plugin/signature"
	"abate/sim/clock"
	"abate/sim/rand"
)

// errCloseBoom is the sentinel returned by the mutation-test plugin on Close.
var errCloseBoom = errors.New("hotswap test: close failed")

// closeErrorPlugin is a minimal Plugin whose Close fails, used to prove UNLOAD
// propagates plugin shutdown errors instead of silently discarding them.
type closeErrorPlugin struct {
	manifest host.Manifest
}

// Manifest returns the signed metadata of the plugin.
func (p *closeErrorPlugin) Manifest() host.Manifest { return p.manifest }

// Type returns the plugin kind.
func (p *closeErrorPlugin) Type() host.PluginType { return p.manifest.Type }

// Close always fails with errCloseBoom.
func (*closeErrorPlugin) Close() error { return errCloseBoom }

// rawInstance builds an instance without going through Host.Load, so mutation
// tests can inject incomplete or hostile metadata (missing registry, wrong type,
// failing Close) that a verified Load would never produce.
func rawInstance(m signature.Manifest, p host.Plugin, reg *host.Registry) *host.Instance {
	return &host.Instance{
		Plugin:   p,
		Manifest: m,
		Caps:     allCaps(),
		Clock:    clock.NewVirtualClock(),
		RNG:      rand.New(1),
		Entries:  []string{host.EntryApply},
		Registry: reg,
	}
}

// basicManifest returns a minimal manifest for raw instances.
func basicManifest(v signature.SemVer, typ signature.PluginType) signature.Manifest {
	return signature.Manifest{
		ID:      signature.PluginID("sample/mutation"),
		Version: v,
		Type:    typ,
	}
}

func TestBeginRejectsNilAndIdenticalInstances(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	cases := []struct {
		name string
		old  *host.Instance
		next *host.Instance
	}{
		{"nil-old", nil, f.next},
		{"nil-next", f.old, nil},
		{"both-nil", nil, nil},
		{"identical", f.old, f.old},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sw, err := Begin(context.Background(), c.old, c.next)
			if !errors.Is(err, ErrSwapInvalidInstance) {
				t.Fatalf("Begin = %v, want ErrSwapInvalidInstance", err)
			}
			if sw.State() != PREPARE {
				t.Fatalf("state = %s, want PREPARE", sw.State())
			}
		})
	}
}

func TestBeginRejectsMissingRegistry(t *testing.T) {
	mOld := basicManifest(ver(1, 0, 0), host.PluginTypeOperator)
	mNext := basicManifest(ver(1, 0, 1), host.PluginTypeOperator)
	oldInst := rawInstance(mOld, &closeErrorPlugin{manifest: mOld}, nil)
	nextInst := rawInstance(mNext, &closeErrorPlugin{manifest: mNext}, nil)
	sw, err := Begin(context.Background(), oldInst, nextInst)
	if !errors.Is(err, ErrSwapNoRegistry) {
		t.Fatalf("Begin = %v, want ErrSwapNoRegistry", err)
	}
	if sw.State() != PREPARE {
		t.Fatalf("state = %s, want PREPARE", sw.State())
	}
	if oldInst.IsSuperseded() {
		t.Fatal("PREPARE failure must not mark the old instance superseded")
	}
}

func TestBeginRejectsIdentityMismatch(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	t.Run("different-id", func(t *testing.T) {
		changed := f.next
		changed.Manifest.ID = signature.PluginID("sample/other")
		sw, err := Begin(context.Background(), f.old, changed)
		if !errors.Is(err, ErrSwapIdentityMismatch) || sw.State() != PREPARE {
			t.Fatalf("Begin = (%s,%v), want (PREPARE,ErrSwapIdentityMismatch)", sw.State(), err)
		}
	})
	t.Run("different-type", func(t *testing.T) {
		changed := f.next
		changed.Manifest.Type = host.PluginTypeReporter
		if _, err := Begin(context.Background(), f.old, changed); !errors.Is(err, ErrSwapIdentityMismatch) {
			t.Fatalf("Begin = %v, want ErrSwapIdentityMismatch", err)
		}
	})
}

func TestBeginRejectsVersionViolations(t *testing.T) {
	cases := []struct {
		name string
		from signature.SemVer
		to   signature.SemVer
	}{
		{"patch-rollback", ver(1, 0, 2), ver(1, 0, 1)},
		{"minor-rollback", ver(1, 3, 0), ver(1, 2, 9)},
		{"major-upgrade", ver(1, 0, 0), ver(2, 0, 0)},
		{"major-downgrade", ver(2, 0, 0), ver(1, 9, 9)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newSwapFixture(t, c.from, c.to, host.PluginTypeOperator)
			sw, err := Begin(context.Background(), f.old, f.next)
			if !errors.Is(err, ErrSwapNotAllowed) {
				t.Fatalf("Begin = %v, want ErrSwapNotAllowed", err)
			}
			if sw.State() != PREPARE {
				t.Fatalf("state = %s, want PREPARE", sw.State())
			}
			if got, ok := f.h.Lookup(f.id); !ok || got != f.old {
				t.Fatal("rejected swap changed the registry")
			}
			if f.old.IsSuperseded() {
				t.Fatal("rejected swap superseded the old instance")
			}
		})
	}
}

func TestBeginRejectsUnknownPluginType(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	reg := f.h.Registry()
	mOld := basicManifest(ver(1, 0, 0), host.PluginType("sandbox"))
	mNext := basicManifest(ver(1, 0, 1), host.PluginType("sandbox"))
	oldInst := rawInstance(mOld, &closeErrorPlugin{manifest: mOld}, reg)
	nextInst := rawInstance(mNext, &closeErrorPlugin{manifest: mNext}, reg)
	if _, err := Begin(context.Background(), oldInst, nextInst); !errors.Is(err, ErrSwapNotAllowed) {
		t.Fatalf("Begin = %v, want ErrSwapNotAllowed", err)
	}
}

func TestBeginRejectsSupersededReplacement(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	f.next.MarkSuperseded()
	sw, err := Begin(context.Background(), f.old, f.next)
	if !errors.Is(err, ErrSwapNotAllowed) {
		t.Fatalf("Begin = %v, want ErrSwapNotAllowed", err)
	}
	if sw.State() != PREPARE {
		t.Fatalf("state = %s, want PREPARE", sw.State())
	}
}

func TestBeginPropagatesCloseError(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	reg := f.h.Registry()
	mOld := basicManifest(ver(1, 0, 0), host.PluginTypeOperator)
	mNext := basicManifest(ver(1, 0, 1), host.PluginTypeOperator)
	oldInst := rawInstance(mOld, &closeErrorPlugin{manifest: mOld}, reg)
	nextInst := rawInstance(mNext, &closeErrorPlugin{manifest: mNext}, nil)
	sw, err := Begin(context.Background(), oldInst, nextInst)
	if !errors.Is(err, errCloseBoom) {
		t.Fatalf("Begin = %v, want errCloseBoom", err)
	}
	if sw.State() != UNLOAD {
		t.Fatalf("state = %s, want UNLOAD", sw.State())
	}
}

func TestSwitchOverUsesNextRegistry(t *testing.T) {
	f := newSwapFixture(t, ver(1, 0, 0), ver(1, 0, 1), host.PluginTypeOperator)
	reg := f.h.Registry()
	mOld := basicManifest(ver(1, 0, 0), host.PluginTypeOperator)
	mNext := basicManifest(ver(1, 0, 1), host.PluginTypeOperator)
	oldInst := rawInstance(mOld, &closeErrorPlugin{manifest: mOld}, nil)
	nextInst := rawInstance(mNext, &closeErrorPlugin{manifest: mNext}, reg)
	if _, err := Begin(context.Background(), oldInst, nextInst); !errors.Is(err, errCloseBoom) {
		t.Fatalf("Begin = %v, want errCloseBoom from UNLOAD", err)
	}
	if got, ok := reg.Get(mOld.ID); !ok || got != nextInst {
		t.Fatal("SWITCH did not use the replacement instance registry")
	}
}
