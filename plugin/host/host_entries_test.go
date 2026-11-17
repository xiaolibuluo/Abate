/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"abate/lang/rule"
	"abate/plugin/capability"
	"abate/plugin/report"
	"abate/plugin/signature"
)

// onlyClock grants a capability set that deliberately lacks CapObserve, so the
// observation gate on reporter and adapter entries is observable.
func onlyClock() Capabilities {
	var c Capabilities
	c.Grant(capability.CapClock)
	return c
}

func TestCallReporterRequiresCapObserve(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/rep-gate", signature.PluginTypeReporter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testReporterFactory, onlyClock())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if _, err := h.Call(inst, EntryBuild, nil, nil); !errors.Is(err, capability.ErrCapabilityDenied) {
		t.Fatalf("reporter call without CapObserve = %v", err)
	}
}

func TestCallAdapterRequiresCapObserve(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/adp-gate", signature.PluginTypeAdapter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testAdapterFactory, onlyClock())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if _, err := h.Call(inst, EntryDetect, nil, nil); !errors.Is(err, capability.ErrCapabilityDenied) {
		t.Fatalf("adapter detect without CapObserve = %v", err)
	}
	if _, err := h.Call(inst, EntryEmit, []byte("payload"), nil); !errors.Is(err, capability.ErrCapabilityDenied) {
		t.Fatalf("adapter emit without CapObserve = %v", err)
	}
}

func TestCallBuildMalformedArgsRejected(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/rep-badargs", signature.PluginTypeReporter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testReporterFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if _, err := h.Call(inst, EntryBuild, []byte("{broken"), []Capability{capability.CapObserve}); err == nil {
		t.Fatal("EntryBuild accepted malformed BuildArgs JSON")
	}
}

func TestCallBuildEmptyArgsProducesValidSARIF(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/rep-empty", signature.PluginTypeReporter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testReporterFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	out, err := h.Call(inst, EntryBuild, nil, []Capability{capability.CapObserve})
	if err != nil {
		t.Fatalf("Call build: %v", err)
	}
	if err := report.ValidateSARIFAgainstOASIS(out); err != nil {
		t.Fatalf("empty-args SARIF invalid: %v", err)
	}
}

func TestCallBuildResultTooLarge(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/rep-big", signature.PluginTypeReporter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testReporterFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	args, err := EncodeBuildArgs(BuildArgs{Findings: []rule.Finding{{
		RuleID:  "GO-BIG-001",
		Path:    "huge.go",
		Message: strings.Repeat("x", MaxResultBytes+64),
	}}})
	if err != nil {
		t.Fatalf("EncodeBuildArgs: %v", err)
	}
	if _, err := h.Call(inst, EntryBuild, args, []Capability{capability.CapObserve}); !errors.Is(err, ErrResultTooLarge) {
		t.Fatalf("oversize reporter result = %v, want ErrResultTooLarge", err)
	}
}

func TestCallDetectMalformedArgsRejected(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/adp-badargs", signature.PluginTypeAdapter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testAdapterFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if _, err := h.Call(inst, EntryDetect, []byte("{broken"), []Capability{capability.CapObserve}); err == nil {
		t.Fatal("EntryDetect accepted malformed environment JSON")
	}
}

func TestCallDetectEmptyEnvAndEmitPayload(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/adp-empty", signature.PluginTypeAdapter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testAdapterFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	wants := []Capability{capability.CapObserve}
	out, err := h.Call(inst, EntryDetect, nil, wants)
	if err != nil {
		t.Fatalf("Call detect: %v", err)
	}
	var res AdapterDetectResult
	if err := json.Unmarshal(out, &res); err != nil {
		t.Fatalf("decode detect result: %v", err)
	}
	if res.OK || res.ID != "" {
		t.Fatalf("empty environment must not match, got %+v", res)
	}
	if _, err := h.Call(inst, EntryEmit, []byte{0xff}, wants); err == nil {
		t.Fatal("EntryEmit accepted invalid UTF-8 payload")
	}
	emitOut, err := h.Call(inst, EntryEmit, []byte("ok"), wants)
	if err != nil {
		t.Fatalf("Call emit: %v", err)
	}
	if len(emitOut) != 0 {
		t.Fatalf("emit result length = %d, want 0", len(emitOut))
	}
}

func TestCallEntryTypeMismatchForbidden(t *testing.T) {
	entries := []string{EntryApply, EntryBuild, EntryDetect, EntryEmit}
	wants := []Capability{capability.CapRand, capability.CapObserve}
	for _, entry := range entries {
		inst := &Instance{
			Plugin:   &plainPlugin{},
			Manifest: Manifest{ID: PluginID("sample/plain")},
			Caps:     allCapabilities(),
			Entries:  []string{entry},
		}
		if _, err := NewHost(signature.NewTrustStore()).Call(inst, entry, nil, wants); !errors.Is(err, ErrEntryForbidden) {
			t.Fatalf("entry %q with a non-conforming plugin = %v", entry, err)
		}
	}
}
