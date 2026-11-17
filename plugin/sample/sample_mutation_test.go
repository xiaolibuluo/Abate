/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sample

import (
	"context"
	"errors"
	"testing"

	"abate/kernel/mutation"
	"abate/plugin/host"
	"abate/plugin/report"
	"abate/plugin/signature"
)

// TestOperatorApplyIsSensitiveToSeedAndInput proves the operator is not a
// constant function: changing the injected RNG seed or the input changes bytes.
func TestOperatorApplyIsSensitiveToSeedAndInput(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/op-mut", signature.PluginTypeOperator, "clock,rand")
	p := NewOperatorPlugin(m)
	apply := func(seed uint64, input string) [32]byte {
		ctx := host.WithBackend(context.Background(), sampleBackend(seed))
		out, err := p.Apply(ctx, []byte(input))
		if err != nil {
			t.Fatalf("Apply: %v", err)
		}
		return shaOf(out)
	}
	base := apply(1, "payload")
	if base == apply(2, "payload") {
		t.Fatal("changing the RNG seed did not change the output")
	}
	if base == apply(1, "payloaD") {
		t.Fatal("changing the input did not change the output")
	}
	if base != apply(1, "payload") {
		t.Fatal("identical inputs diverged")
	}
}

// TestReporterBuildIsSensitiveToFindings proves a single finding mutation is
// visible in the emitted SARIF artifact and keeps it schema-valid.
func TestReporterBuildIsSensitiveToFindings(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/rep-mut", signature.PluginTypeReporter, "observe")
	p := NewReporterPlugin(m)
	base, err := p.Build(sampleFindings(), mutation.Report{Seed: 1})
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	mutated := sampleFindings()
	mutated[0].Message = "a different diagnostic message"
	changed, err := p.Build(mutated, mutation.Report{Seed: 1})
	if err != nil {
		t.Fatalf("Build mutated: %v", err)
	}
	if shaOf(base) == shaOf(changed) {
		t.Fatal("changing a finding did not change the SARIF output")
	}
	if err := report.ValidateSARIFAgainstOASIS(changed); err != nil {
		t.Fatalf("mutated SARIF is invalid: %v", err)
	}
}

// TestFactoryRejectsTamperedType proves the factories fail closed when a
// manifest type is changed after signing.
func TestFactoryRejectsTamperedType(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/tampered", signature.PluginTypeOperator, "clock,rand")
	if err := signature.Verify(store, m); err != nil {
		t.Fatalf("signed manifest should verify: %v", err)
	}
	m.Type = signature.PluginTypeReporter
	if err := signature.Verify(store, m); err == nil {
		t.Fatal("tampering with the type should invalidate the signature")
	}
	h := host.NewHost(store)
	if _, err := h.Load(m, ReporterFactory, sampleCaps()); !errors.Is(err, signature.ErrSignatureInvalid) {
		t.Fatalf("Host.Load = %v, want ErrSignatureInvalid", err)
	}
}
