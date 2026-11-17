/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"context"
	"errors"
	"testing"

	"abate/plugin/capability"
	"abate/plugin/signature"
)

// plainPlugin implements only Plugin, so entriesFor has no entry to whitelist
// and every dispatch branch must fall through to ErrEntryForbidden.
type plainPlugin struct {
	manifest Manifest
}

// Manifest returns the signed metadata of the plugin.
func (p *plainPlugin) Manifest() Manifest { return p.manifest }

// Type returns the plugin kind.
func (p *plainPlugin) Type() PluginType { return p.manifest.Type }

// Close releases no resources.
func (*plainPlugin) Close() error { return nil }

func TestEntriesForUnknownPlugin(t *testing.T) {
	if got := entriesFor(&plainPlugin{}); len(got) != 0 {
		t.Fatalf("entriesFor(plain) = %v, want none", got)
	}
}

func TestRegistryRegisterAndReplace(t *testing.T) {
	r := newRegistry()
	factory := func(m Manifest) (Plugin, error) { return &plainPlugin{manifest: m}, nil }
	if err := r.Register(factory, PluginTypeOperator); err != nil {
		t.Fatalf("Register: %v", err)
	}
	if err := r.Register(factory, PluginTypeOperator); !errors.Is(err, errFactoryExists) {
		t.Fatalf("duplicate Register = %v, want errFactoryExists", err)
	}
	id := PluginID("sample/replace")
	first := &Instance{Manifest: Manifest{ID: id}}
	second := &Instance{Manifest: Manifest{ID: id}}
	r.Replace(id, first)
	if got, ok := r.Get(id); !ok || got != first {
		t.Fatal("Replace did not register the first instance")
	}
	r.Replace(id, second)
	if got, ok := r.Get(id); !ok || got != second {
		t.Fatal("Replace is not atomic for readers")
	}
	if snap := r.Snapshot(); len(snap) != 1 || snap[id] != second {
		t.Fatalf("Snapshot after Replace = %v", snap)
	}
}

func TestInstanceSupersededLifecycle(t *testing.T) {
	inst := &Instance{}
	if inst.IsSuperseded() {
		t.Fatal("a new instance must not be superseded")
	}
	release, ok := inst.BeginCall()
	if !ok {
		t.Fatal("BeginCall rejected a live instance")
	}
	inst.MarkSuperseded()
	if !inst.IsSuperseded() {
		t.Fatal("MarkSuperseded did not set the flag")
	}
	if _, ok := inst.BeginCall(); ok {
		t.Fatal("BeginCall accepted a superseded instance")
	}
	done := make(chan struct{})
	go func() {
		inst.WaitCalls()
		close(done)
	}()
	select {
	case <-done:
		t.Fatal("WaitCalls returned while a call was still in flight")
	default:
	}
	release()
	<-done
}

func TestBackendAbsentAndPresent(t *testing.T) {
	if got := Backend(context.Background()); got != nil {
		t.Fatalf("Backend without injection = %v, want nil", got)
	}
	want := &HostBackend{}
	ctx := WithBackend(context.Background(), want)
	if got := Backend(ctx); got != want {
		t.Fatal("Backend did not return the injected value")
	}
}

func TestLoadNilPluginRejected(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/nilplugin", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	_, err := h.Load(m, func(Manifest) (Plugin, error) { return nil, nil }, allCapabilities())
	if !errors.Is(err, errNilPlugin) {
		t.Fatalf("Load with a nil factory result = %v, want errNilPlugin", err)
	}
}

func TestLoadPlainPluginHasNoEntries(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/plain", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	inst, err := h.Load(m, func(mm Manifest) (Plugin, error) { return &plainPlugin{manifest: mm}, nil }, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(inst.Entries) != 0 {
		t.Fatalf("plain plugin entries = %v", inst.Entries)
	}
	if _, err := h.Call(inst, EntryApply, nil, nil); !errors.Is(err, ErrEntryForbidden) {
		t.Fatalf("plain plugin call = %v, want ErrEntryForbidden", err)
	}
}

func TestCloseNilInstanceAndStaleCall(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/stale", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	if err := h.Close(nil); err != nil {
		t.Fatalf("Close(nil): %v", err)
	}
	inst, err := h.Load(m, testOperatorFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	inst.MarkSuperseded()
	_, err = h.Call(inst, EntryApply, []byte("x"), []Capability{capability.CapRand})
	if !errors.Is(err, ErrInstanceStale) {
		t.Fatalf("stale call = %v, want ErrInstanceStale", err)
	}
}
