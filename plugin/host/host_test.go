/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"testing"

	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/rule"
	"abate/plugin/capability"
	"abate/plugin/report"
	"abate/plugin/signature"
)

// testAnchor is the trust anchor id used by the signed-manifest test helpers.
const testAnchor = "test-anchor"

// allCapabilities grants every known capability so a loaded plugin is callable.
func allCapabilities() capability.Capabilities {
	var c capability.Capabilities
	for _, bit := range []capability.Capability{
		capability.CapClock,
		capability.CapRand,
		capability.CapObserve,
		capability.CapFS,
		capability.CapNet,
	} {
		c.Grant(bit)
	}
	return c
}

// operatorCaps returns the capabilities an operator entry requires.
func operatorCaps() []capability.Capability {
	return []capability.Capability{capability.CapRand}
}

// signManifest builds and signs a manifest under the given anchor in store.
func signManifest(t *testing.T, store *signature.TrustStore, anchor, id string, typ signature.PluginType, decl string) signature.Manifest {
	t.Helper()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	if err := store.Add(anchor, pub); err != nil {
		t.Fatalf("store.Add: %v", err)
	}
	m := signature.Manifest{
		ID:               signature.PluginID(id),
		Version:          signature.SemVer{Major: 1, Minor: 0, Patch: 0},
		Type:             typ,
		CapabilitiesDecl: decl,
		TrustAnchor:      anchor,
		BuildRef:         "build-test",
	}
	if err := signature.Sign(store, priv, &m); err != nil {
		t.Fatalf("Sign: %v", err)
	}
	return m
}

func sampleFindings() []rule.Finding {
	return []rule.Finding{
		{
			RuleID:   rule.RuleID("GO-EG-001"),
			Path:     "a/b.go",
			Line:     12,
			Column:   4,
			Severity: rule.SeverityError,
			Category: rule.CategoryErrorProne,
			Message:  "error unchecked",
		},
		{
			RuleID:   rule.RuleID("GO-VET-003"),
			Path:     "c/d.go",
			Line:     30,
			Column:   2,
			Severity: rule.SeverityWarning,
			Category: rule.CategoryConcurrency,
			Message:  "possible nil deref",
		},
	}
}

func sampleReport() mutation.Report {
	return mutation.Report{
		Seed:  0xABCD,
		Score: verify.Score{Total: 4, Killed: 3, Survived: 1},
	}
}

func shaOf(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}

func TestLoadAndCallOperatorDeterministic(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/perturb", signature.PluginTypeOperator, "clock,rand")

	h1 := NewHost(store)
	inst1, err := h1.Load(m, testOperatorFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load 1: %v", err)
	}
	out1, err := h1.Call(inst1, EntryApply, []byte("deterministic-input"), operatorCaps())
	if err != nil {
		t.Fatalf("Call 1: %v", err)
	}

	h2 := NewHost(store)
	inst2, err := h2.Load(m, testOperatorFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load 2: %v", err)
	}
	out2, err := h2.Call(inst2, EntryApply, []byte("deterministic-input"), operatorCaps())
	if err != nil {
		t.Fatalf("Call 2: %v", err)
	}

	if shaOf(out1) != shaOf(out2) {
		t.Fatal("operator output is not deterministic across identical hosts")
	}
	if len(out1) != len("deterministic-input")+1 {
		t.Fatalf("operator output length = %d, want %d", len(out1), len("deterministic-input")+1)
	}
}

func TestLoadUnsignedManifestRejected(t *testing.T) {
	store := signature.NewTrustStore()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	if err := store.Add(testAnchor, pub); err != nil {
		t.Fatalf("store.Add: %v", err)
	}
	_ = priv
	m := signature.Manifest{
		ID:               signature.PluginID("sample/unsigned"),
		Version:          signature.SemVer{Major: 1, Minor: 0, Patch: 0},
		Type:             signature.PluginTypeOperator,
		CapabilitiesDecl: "clock,rand",
		TrustAnchor:      testAnchor,
		BuildRef:         "build-test",
	}
	h := NewHost(store)
	if _, err := h.Load(m, testOperatorFactory, allCapabilities()); err == nil {
		t.Fatal("Load accepted an unsigned manifest")
	} else if !errors.Is(err, signature.ErrSignatureInvalid) && !errors.Is(err, signature.ErrManifestHashMismatch) {
		t.Fatalf("Load returned %v, want ErrSignatureInvalid/ErrManifestHashMismatch", err)
	}
}

func TestLoadUnknownCapabilityRejected(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/badcap", signature.PluginTypeOperator, "bogus")
	h := NewHost(store)
	if _, err := h.Load(m, testOperatorFactory, allCapabilities()); err == nil {
		t.Fatal("Load accepted an unknown capability declaration")
	} else if !errors.Is(err, capability.ErrUnknownCapability) {
		t.Fatalf("Load returned %v, want ErrUnknownCapability", err)
	}
}

func TestCallMissingCapabilityDenied(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/nocap", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	var none capability.Capabilities
	inst, err := h.Load(m, testOperatorFactory, none)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if _, err := h.Call(inst, EntryApply, []byte("x"), operatorCaps()); err == nil {
		t.Fatal("Call succeeded without the required capability")
	} else if !errors.Is(err, capability.ErrCapabilityDenied) {
		t.Fatalf("Call returned %v, want ErrCapabilityDenied", err)
	}
}

func TestCallForbiddenEntry(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/forbidden", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	inst, err := h.Load(m, testOperatorFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if _, err := h.Call(inst, EntryBuild, nil, operatorCaps()); err == nil {
		t.Fatal("Call accepted a non-whitelisted entry")
	} else if !errors.Is(err, ErrEntryForbidden) {
		t.Fatalf("Call returned %v, want ErrEntryForbidden", err)
	}
}

func TestCallResultTooLarge(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/big", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	inst, err := h.Load(m, testOperatorFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	big := make([]byte, MaxResultBytes+5)
	if _, err := h.Call(inst, EntryApply, big, operatorCaps()); err == nil {
		t.Fatal("Call accepted an over-size result")
	} else if !errors.Is(err, ErrResultTooLarge) {
		t.Fatalf("Call returned %v, want ErrResultTooLarge", err)
	}
}

func TestRegistryAtomicity(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/reg", signature.PluginTypeOperator, "clock,rand")
	h := NewHost(store)
	inst, err := h.Load(m, testOperatorFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got, ok := h.Lookup(inst.Manifest.ID); !ok || got != inst {
		t.Fatal("Lookup did not return the loaded instance")
	}
	if snap := h.Registry().Snapshot(); len(snap) != 1 {
		t.Fatalf("snapshot size = %d, want 1", len(snap))
	}
	if err := h.Close(inst); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if _, ok := h.Lookup(inst.Manifest.ID); ok {
		t.Fatal("Lookup still returned instance after Close")
	}
}

func TestCallReporterBuild(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/reporter", signature.PluginTypeReporter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testReporterFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	args, err := EncodeBuildArgs(BuildArgs{Findings: sampleFindings(), Report: sampleReport()})
	if err != nil {
		t.Fatalf("EncodeBuildArgs: %v", err)
	}
	out, err := h.Call(inst, EntryBuild, args, []capability.Capability{capability.CapObserve})
	if err != nil {
		t.Fatalf("Call build: %v", err)
	}
	if len(out) == 0 {
		t.Fatal("reporter produced empty SARIF")
	}
	if err := report.ValidateSARIFAgainstOASIS(out); err != nil {
		t.Fatalf("reporter SARIF failed OASIS validation: %v", err)
	}
}

func TestCallAdapterDetectAndEmit(t *testing.T) {
	store := signature.NewTrustStore()
	m := signManifest(t, store, testAnchor, "sample/adapter", signature.PluginTypeAdapter, "observe")
	h := NewHost(store)
	inst, err := h.Load(m, testAdapterFactory, allCapabilities())
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	env, _ := json.Marshal(map[string]string{"CI": "true"})
	detectOut, err := h.Call(inst, EntryDetect, env, []capability.Capability{capability.CapObserve})
	if err != nil {
		t.Fatalf("Call detect: %v", err)
	}
	var dr AdapterDetectResult
	if err := json.Unmarshal(detectOut, &dr); err != nil {
		t.Fatalf("decode detect result: %v", err)
	}
	if !dr.OK || dr.ID != "adp-ci" {
		t.Fatalf("detect result = %+v, want {adp-ci,true}", dr)
	}

	if _, err := h.Call(inst, EntryEmit, []byte("valid-payload"), []capability.Capability{capability.CapObserve}); err != nil {
		t.Fatalf("Call emit valid: %v", err)
	}
	if _, err := h.Call(inst, EntryEmit, nil, []capability.Capability{capability.CapObserve}); err == nil {
		t.Fatal("Call emit accepted an empty payload")
	}
}
