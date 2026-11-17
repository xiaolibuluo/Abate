/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sample

import (
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"testing"

	"abate/kernel/algebra"
	"abate/kernel/mutation"
	"abate/lang/rule"
	"abate/plugin/capability"
	"abate/plugin/host"
	"abate/plugin/report"
	"abate/plugin/signature"
	"abate/sim/clock"
	"abate/sim/rand"
)

// sampleStore creates a trust store with one registered release anchor.
func sampleStore(t *testing.T) (*signature.TrustStore, ed25519.PrivateKey) {
	t.Helper()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	store := signature.NewTrustStore()
	if err := store.Add("sample-anchor", pub); err != nil {
		t.Fatalf("store.Add: %v", err)
	}
	return store, priv
}

// sampleManifest builds a signed manifest for the sample tests.
func sampleManifest(t *testing.T, store *signature.TrustStore, priv ed25519.PrivateKey,
	id string, typ signature.PluginType, decl string) host.Manifest {
	t.Helper()
	m := signature.Manifest{
		ID:               signature.PluginID(id),
		Version:          signature.SemVer{Major: 1, Minor: 0, Patch: 0},
		Type:             typ,
		CapabilitiesDecl: decl,
		TrustAnchor:      "sample-anchor",
		BuildRef:         "build-sample",
	}
	if err := signature.Sign(store, priv, &m); err != nil {
		t.Fatalf("Sign: %v", err)
	}
	return m
}

// sampleCaps grants every known capability.
func sampleCaps() host.Capabilities {
	var c host.Capabilities
	for _, bit := range []capability.Capability{
		capability.CapClock, capability.CapRand, capability.CapObserve,
		capability.CapFS, capability.CapNet,
	} {
		c.Grant(bit)
	}
	return c
}

// sampleFindings returns a deterministic pair of findings.
func sampleFindings() []rule.Finding {
	return []rule.Finding{
		{
			RuleID: "GO-EG-001", Path: "a/b.go", Line: 12, Column: 4,
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Message: "error unchecked",
		},
		{
			RuleID: "GO-VET-003", Path: "c/d.go", Line: 30, Column: 2,
			Severity: rule.SeverityWarning, Category: rule.CategoryConcurrency,
			Message: "possible nil deref",
		},
	}
}

// sampleBackend builds a fresh deterministic host backend for one call.
func sampleBackend(seed uint64) *host.HostBackend {
	return &host.HostBackend{Clock: clock.NewVirtualClock(), RNG: rand.New(seed)}
}

func shaOf(b []byte) [32]byte {
	return sha256.Sum256(b)
}

func TestFactoriesTypeMatrix(t *testing.T) {
	store, priv := sampleStore(t)
	factories := []struct {
		typ     signature.PluginType
		factory host.Factory
	}{
		{signature.PluginTypeOperator, OperatorFactory},
		{signature.PluginTypeReporter, ReporterFactory},
		{signature.PluginTypeAdapter, AdapterFactory},
	}
	for _, f := range factories {
		for _, m := range factories {
			manifest := sampleManifest(t, store, priv, "sample/matrix", m.typ, "clock,rand,observe")
			plugin, err := f.factory(manifest)
			if m.typ == f.typ {
				if err != nil || plugin == nil {
					t.Fatalf("factory %q rejected its own type %q: %v", f.typ, m.typ, err)
				}
				if plugin.Type() != m.typ {
					t.Fatalf("factory %q produced type %q", f.typ, plugin.Type())
				}
				if err := plugin.Close(); err != nil {
					t.Fatalf("Close: %v", err)
				}
			} else if err == nil {
				t.Fatalf("factory %q accepted a %q manifest", f.typ, m.typ)
			}
		}
	}
}

func TestOperatorMetadataAndDescribe(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/op", signature.PluginTypeOperator, "clock,rand")
	p := NewOperatorPlugin(m)
	if p.Manifest().ID != m.ID || p.Type() != signature.PluginTypeOperator {
		t.Fatal("operator metadata does not mirror the manifest")
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	op := p.Describe()
	if op == nil {
		t.Fatal("Describe returned nil")
	}
	if op.Family() != algebra.FamilyMutation {
		t.Fatalf("family = %q", op.Family())
	}
	if op.ID() == "" || len(op.Targets()) != 1 {
		t.Fatalf("operator descriptor incomplete: %q %v", op.ID(), op.Targets())
	}
	if op.Params()["algo"] != "xor-keystream" {
		t.Fatalf("params = %v", op.Params())
	}
	if op.Invertible() {
		t.Fatal("sample perturbation must be non-invertible")
	}
	if _, err := op.Inverse(); err == nil {
		t.Fatal("Inverse succeeded for a non-invertible operator")
	}
	if op.Fingerprint() == [32]byte{} {
		t.Fatal("operator fingerprint is zero")
	}
}

func TestOperatorApplyDeterministicWithBackend(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/apply", signature.PluginTypeOperator, "clock,rand")
	p := NewOperatorPlugin(m)
	backend := sampleBackend(0xCAFE)
	ctx := host.WithBackend(context.Background(), backend)
	first, err := p.Apply(ctx, []byte("deterministic-input"))
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	second, err := p.Apply(ctx, []byte("deterministic-input"))
	if err != nil {
		t.Fatalf("Apply again: %v", err)
	}
	if shaOf(first) != shaOf(second) {
		t.Fatal("Apply is not reproducible on one backend")
	}
	if len(first) != len("deterministic-input")+1 {
		t.Fatalf("output length = %d", len(first))
	}
	fresh, err := p.Apply(host.WithBackend(context.Background(), sampleBackend(0xCAFE)), []byte("deterministic-input"))
	if err != nil {
		t.Fatalf("Apply on fresh backend: %v", err)
	}
	if shaOf(first) != shaOf(fresh) {
		t.Fatal("Apply is not reproducible across identical backends")
	}
}

func TestOperatorApplyFallbackWithoutBackend(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/fallback", signature.PluginTypeOperator, "clock,rand")
	p := NewOperatorPlugin(m)
	first, err := p.Apply(context.Background(), []byte("abc"))
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	second, err := p.Apply(context.Background(), []byte("abc"))
	if err != nil {
		t.Fatalf("Apply again: %v", err)
	}
	if shaOf(first) != shaOf(second) {
		t.Fatal("fallback seed is not deterministic")
	}
	changed, err := p.Apply(context.Background(), []byte("abd"))
	if err != nil {
		t.Fatalf("Apply changed input: %v", err)
	}
	if shaOf(first) == shaOf(changed) {
		t.Fatal("fallback seed does not depend on the input")
	}
	if len(first) != len("abc")+1 {
		t.Fatalf("output length = %d", len(first))
	}
}

func TestOperatorApplyPartialBackends(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/partial", signature.PluginTypeOperator, "clock,rand")
	p := NewOperatorPlugin(m)
	clockOnly := host.WithBackend(context.Background(), &host.HostBackend{Clock: clock.NewVirtualClock()})
	rngOnly := host.WithBackend(context.Background(), &host.HostBackend{RNG: rand.New(7)})
	a, err := p.Apply(clockOnly, []byte("partial"))
	if err != nil {
		t.Fatalf("clock-only Apply: %v", err)
	}
	b, err := p.Apply(rngOnly, []byte("partial"))
	if err != nil {
		t.Fatalf("rng-only Apply: %v", err)
	}
	if len(a) != len("partial")+1 || len(b) != len("partial")+1 {
		t.Fatalf("partial backend output lengths = %d,%d", len(a), len(b))
	}
	aAgain, err := p.Apply(clockOnly, []byte("partial"))
	if err != nil {
		t.Fatalf("clock-only Apply again: %v", err)
	}
	bAgain, err := p.Apply(rngOnly, []byte("partial"))
	if err != nil {
		t.Fatalf("rng-only Apply again: %v", err)
	}
	if shaOf(a) != shaOf(aAgain) || shaOf(b) != shaOf(bAgain) {
		t.Fatal("partial backends are not reproducible")
	}
}

func TestReporterMetadataAndBuild(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/rep", signature.PluginTypeReporter, "observe")
	p := NewReporterPlugin(m)
	if p.Manifest().ID != m.ID || p.Type() != signature.PluginTypeReporter {
		t.Fatal("reporter metadata does not mirror the manifest")
	}
	if err := p.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	first, err := p.Build(sampleFindings(), mutation.Report{Seed: 42})
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	if err := report.ValidateSARIFAgainstOASIS(first); err != nil {
		t.Fatalf("reporter SARIF failed OASIS validation: %v", err)
	}
	second, err := p.Build(sampleFindings(), mutation.Report{Seed: 42})
	if err != nil {
		t.Fatalf("Build again: %v", err)
	}
	if shaOf(first) != shaOf(second) {
		t.Fatal("reporter output is not deterministic")
	}
}

func TestAdapterFactoryIsAirGapped(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/adp", signature.PluginTypeAdapter, "observe")
	plugin, err := AdapterFactory(m)
	if err != nil {
		t.Fatalf("AdapterFactory: %v", err)
	}
	ap, ok := plugin.(host.AdapterPlugin)
	if !ok {
		t.Fatal("adapter factory did not produce an AdapterPlugin")
	}
	if id, ok := ap.Detect(map[string]string{"CI": "true"}); !ok || id != "adp-ci" {
		t.Fatalf("Detect = (%q,%v)", id, ok)
	}
	if err := ap.Emit([]byte("air-gapped")); err != nil {
		t.Fatalf("Emit: %v", err)
	}
	if err := ap.Emit(nil); err == nil {
		t.Fatal("Emit accepted an empty payload")
	}
}
