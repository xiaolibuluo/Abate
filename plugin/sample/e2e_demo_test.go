/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sample

import (
	"encoding/json"
	"errors"
	"testing"

	"abate/kernel/mutation"
	"abate/plugin/capability"
	"abate/plugin/host"
	"abate/plugin/report"
	"abate/plugin/signature"
)

// TestEndToEndThreePluginDemo loads the operator, reporter and adapter trinity
// into one Host, calls every entry and validates the reporter artifact against
// the bundled OASIS SARIF 2.1.0 schema. It is the executable evidence for the
// M4 plugin ecosystem acceptance criteria.
func TestEndToEndThreePluginDemo(t *testing.T) {
	store, priv := sampleStore(t)
	h := host.NewHost(store)
	mOp := sampleManifest(t, store, priv, "sample/e2e-operator", signature.PluginTypeOperator, "clock,rand")
	mRep := sampleManifest(t, store, priv, "sample/e2e-reporter", signature.PluginTypeReporter, "observe")
	mAdp := sampleManifest(t, store, priv, "sample/e2e-adapter", signature.PluginTypeAdapter, "observe")

	instOp, err := h.Load(mOp, OperatorFactory, sampleCaps())
	if err != nil {
		t.Fatalf("Load operator: %v", err)
	}
	instRep, err := h.Load(mRep, ReporterFactory, sampleCaps())
	if err != nil {
		t.Fatalf("Load reporter: %v", err)
	}
	instAdp, err := h.Load(mAdp, AdapterFactory, sampleCaps())
	if err != nil {
		t.Fatalf("Load adapter: %v", err)
	}
	if got := len(h.Registry().Snapshot()); got != 3 {
		t.Fatalf("registry size = %d, want 3", got)
	}

	applyWants := []capability.Capability{capability.CapRand}
	first, err := h.Call(instOp, host.EntryApply, []byte("e2e-payload"), applyWants)
	if err != nil {
		t.Fatalf("Call apply: %v", err)
	}
	second, err := h.Call(instOp, host.EntryApply, []byte("e2e-payload"), applyWants)
	if err != nil {
		t.Fatalf("Call apply again: %v", err)
	}
	if shaOf(first) != shaOf(second) {
		t.Fatal("operator output is not deterministic across repeated calls")
	}

	args, err := host.EncodeBuildArgs(host.BuildArgs{
		Findings: sampleFindings(),
		Report:   mutation.Report{Seed: 7},
	})
	if err != nil {
		t.Fatalf("EncodeBuildArgs: %v", err)
	}
	observeWants := []capability.Capability{capability.CapObserve}
	sarif, err := h.Call(instRep, host.EntryBuild, args, observeWants)
	if err != nil {
		t.Fatalf("Call build: %v", err)
	}
	if err := report.ValidateSARIFAgainstOASIS(sarif); err != nil {
		t.Fatalf("demo SARIF failed OASIS validation: %v", err)
	}

	env, err := json.Marshal(map[string]string{"CI": "true"})
	if err != nil {
		t.Fatalf("encode environment: %v", err)
	}
	detected, err := h.Call(instAdp, host.EntryDetect, env, observeWants)
	if err != nil {
		t.Fatalf("Call detect: %v", err)
	}
	var res host.AdapterDetectResult
	if err := json.Unmarshal(detected, &res); err != nil {
		t.Fatalf("decode detect result: %v", err)
	}
	if !res.OK || res.ID != "adp-ci" {
		t.Fatalf("detect result = %+v", res)
	}
	if _, err := h.Call(instAdp, host.EntryEmit, []byte("e2e-report"), observeWants); err != nil {
		t.Fatalf("Call emit: %v", err)
	}
	if _, err := h.Call(instAdp, host.EntryEmit, []byte{0xff}, observeWants); err == nil {
		t.Fatal("demo accepted an invalid adapter payload")
	}

	unsigned := signature.Manifest{
		ID:               "sample/e2e-unsigned",
		Version:          signature.SemVer{Major: 1, Minor: 0, Patch: 0},
		Type:             signature.PluginTypeOperator,
		CapabilitiesDecl: "clock,rand",
		TrustAnchor:      "sample-anchor",
	}
	if _, err := h.Load(unsigned, OperatorFactory, sampleCaps()); !errors.Is(err, signature.ErrSignatureInvalid) {
		t.Fatalf("unsigned manifest = %v, want ErrSignatureInvalid", err)
	}
	if got := len(h.Registry().Snapshot()); got != 3 {
		t.Fatalf("rejected load changed the registry size to %d", got)
	}
}
