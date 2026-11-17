/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tests

import (
	"context"
	"errors"
	"testing"

	"abate/kernel/algebra"
	"abate/plugin/capability"
	"abate/plugin/host"
	"abate/plugin/signature"
)

// The trust anchor used by the signed plugin manifests in this attack suite.
const attackAnchor = "abate/security-attack"

// AttackOperator is a real host.OperatorPlugin for the T2/T3 attack cases. Its
// Apply returns a fixed payload so the test controls the result size exactly.
type attackOperator struct {
	manifest host.Manifest
	payload  []byte
}

// Manifest returns the signed metadata of the attack plugin.
func (p *attackOperator) Manifest() host.Manifest { return p.manifest }

// Type returns the operator plugin kind.
func (*attackOperator) Type() host.PluginType { return host.PluginTypeOperator }

// Close releases no resources for the in-memory attack plugin.
func (*attackOperator) Close() error { return nil }

// Describe returns a stable algebra operator descriptor.
func (*attackOperator) Describe() algebra.Operator {
	return algebra.NewBasic(
		algebra.FormatOperatorID(algebra.FamilyMutation, "security-attack", "1"),
		algebra.FamilyMutation,
		algebra.Params{"algo": "fixed-payload"},
		[]algebra.TargetKey{"security/tests"},
		false,
	)
}

// Apply returns the fixed payload and ignores the injected deterministic
// backend, so the result size is fully controlled by the attack case.
func (p *attackOperator) Apply(_ context.Context, _ []byte) ([]byte, error) {
	return p.payload, nil
}

// AttackFactory builds a host.Factory yielding an attackOperator with payload.
// It fails closed when handed a manifest of the wrong plugin kind.
func attackFactory(payload []byte) host.Factory {
	return func(m host.Manifest) (host.Plugin, error) {
		if m.Type != host.PluginTypeOperator {
			return nil, errors.New("security: attack factory requires an operator manifest")
		}
		return &attackOperator{manifest: m, payload: payload}, nil
	}
}

// SignAttackManifest builds and signs an operator manifest declaring decl. The
// signing chain mirrors TestThreatT1SignatureRejection so the plugin is a real,
// verified artifact before any gate is exercised.
func signAttackManifest(t *testing.T, id string, decl string) (*signature.TrustStore, signature.Manifest) {
	t.Helper()
	store := signature.NewTrustStore()
	pub, priv, err := signature.GenerateKey()
	if err != nil {
		t.Fatalf("generate signing key: %v", err)
	}
	if err := store.Add(attackAnchor, pub); err != nil {
		t.Fatalf("add trust anchor: %v", err)
	}
	m := signature.Manifest{
		ID:               signature.PluginID(id),
		Version:          signature.SemVer{Major: 1, Minor: 0, Patch: 0},
		Type:             signature.PluginTypeOperator,
		CapabilitiesDecl: decl,
		TrustAnchor:      attackAnchor,
		BuildRef:         "security-attack-fixture",
	}
	if err := signature.Sign(store, priv, &m); err != nil {
		t.Fatalf("sign manifest: %v", err)
	}
	return store, m
}

// ClockAndRand grants only the two capabilities a benign operator declares, so
// the effective grant is strictly narrower than the escalation attempts.
func clockAndRand() capability.Capabilities {
	var c capability.Capabilities
	c.Grant(capability.CapClock)
	c.Grant(capability.CapRand)
	return c
}

// TestThreatT2CapabilityEscalation proves T2: a signed plugin declaring only
// clock/rand cannot escalate to net/fs through the real host call gate, and the
// capability sandbox emits an auditable deny event.
func TestThreatT2CapabilityEscalation(t *testing.T) {
	store, m := signAttackManifest(t, "abate/t2-escalate", "clock,rand")
	h := host.NewHost(store)
	inst, err := h.Load(m, attackFactory([]byte("benign")), clockAndRand())
	if err != nil {
		t.Fatalf("load signed operator: %v", err)
	}
	if inst.Caps.Has(capability.CapNet) || inst.Caps.Has(capability.CapFS) {
		t.Fatal("elevated capability was granted to a clock/rand plugin")
	}

	// First escalation vector: request CapNet through host.Call. The entry
	// whitelist and capability gate run before any plugin code executes.
	out, callErr := h.Call(inst, host.EntryApply, []byte("escalate"), []host.Capability{capability.CapNet})
	if !errors.Is(callErr, capability.ErrCapabilityDenied) {
		t.Fatalf("net escalation was not denied by host.Call: err=%v out=%q", callErr, out)
	}
	if len(out) != 0 {
		t.Fatalf("denied call returned a payload: %q", out)
	}

	// Second escalation vector: CapFS must be refused by the same gate.
	if _, fsErr := h.Call(inst, host.EntryApply, []byte("escalate"), []host.Capability{capability.CapFS}); !errors.Is(fsErr, capability.ErrCapabilityDenied) {
		t.Fatalf("filesystem escalation was not denied: %v", fsErr)
	}

	// The audited variant of the same gate records a deny event for forensics.
	declared, parseErr := capability.ParseGrant("clock,rand")
	if parseErr != nil {
		t.Fatalf("parse grant: %v", parseErr)
	}
	auditLog := capability.NewAuditLog()
	if auditErr := capability.CheckAudit(declared, capability.CapNet, auditLog, "abate/t2-escalate"); !errors.Is(auditErr, capability.ErrCapabilityDenied) {
		t.Fatalf("audited escalation was not denied: %v", auditErr)
	}
	entries := auditLog.Entries()
	if len(entries) != 1 {
		t.Fatalf("audit log has %d events, want 1", len(entries))
	}
	event := entries[0]
	if event.Kind != "deny" || !event.IsDeny() {
		t.Fatalf("audit event kind = %q, want deny", event.Kind)
	}
	if event.PluginID != "abate/t2-escalate" {
		t.Fatalf("audit plugin id = %q, want abate/t2-escalate", event.PluginID)
	}
	if event.Capability != capability.CapNet {
		t.Fatalf("audit capability = %v, want net", event.Capability)
	}
	if event.Result {
		t.Fatal("deny event incorrectly reported success")
	}
	if event.At != 1 {
		t.Fatalf("audit virtual sequence = %d, want 1", event.At)
	}
}

// TestThreatT3ResourceExhaustion proves T3: a plugin returning more than
// host.MaxResultBytes is refused by the real host call path with
// host.ErrResultTooLarge, while a result exactly at the cap is accepted.
func TestThreatT3ResourceExhaustion(t *testing.T) {
	if host.MaxResultBytes != 1<<20 {
		t.Fatalf("result cap = %d, want 1048576", host.MaxResultBytes)
	}
	store, m := signAttackManifest(t, "abate/t3-exhaust", "clock,rand")
	h := host.NewHost(store)
	payload := make([]byte, host.MaxResultBytes+1)
	inst, err := h.Load(m, attackFactory(payload), clockAndRand())
	if err != nil {
		t.Fatalf("load signed operator: %v", err)
	}
	// CapRand is granted so the capability gate passes; otherwise T2 would
	// mask T3 and the assertion would not exercise the result-size ceiling.
	out, callErr := h.Call(inst, host.EntryApply, []byte("input"), []host.Capability{capability.CapRand})
	if !errors.Is(callErr, host.ErrResultTooLarge) {
		t.Fatalf("oversize result was not rejected: err=%v len(out)=%d", callErr, len(out))
	}
	if len(out) != 0 {
		t.Fatalf("oversize call returned %d bytes, want none", len(out))
	}

	// Boundary: a payload of exactly MaxResultBytes is still accepted. A fresh
	// host avoids reusing the same registry id for a distinct instance.
	boundaryStore, boundaryManifest := signAttackManifest(t, "abate/t3-boundary", "clock,rand")
	boundaryHost := host.NewHost(boundaryStore)
	exactInst, loadErr := boundaryHost.Load(boundaryManifest,
		attackFactory(make([]byte, host.MaxResultBytes)), clockAndRand())
	if loadErr != nil {
		t.Fatalf("load boundary operator: %v", loadErr)
	}
	exactOut, exactErr := boundaryHost.Call(exactInst, host.EntryApply, nil,
		[]host.Capability{capability.CapRand})
	if exactErr != nil {
		t.Fatalf("result at cap was rejected: %v", exactErr)
	}
	if len(exactOut) != host.MaxResultBytes {
		t.Fatalf("result at cap = %d bytes, want %d", len(exactOut), host.MaxResultBytes)
	}
}
