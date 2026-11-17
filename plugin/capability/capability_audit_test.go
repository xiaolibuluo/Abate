/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

import (
	"errors"
	"testing"
)

func TestCapabilityStringAllBranches(t *testing.T) {
	want := map[Capability]string{
		CapClock:   "clock",
		CapRand:    "rand",
		CapObserve: "observe",
		CapFS:      "fs",
		CapNet:     "net",
	}
	for c, name := range want {
		if c.String() != name {
			t.Fatalf("Capability(%d).String() = %q, want %q", c, c.String(), name)
		}
	}
	// Unknown bit must fall back to the placeholder form.
	if Capability(99).String() != "cap(?)" {
		t.Fatalf("unknown capability string = %q, want cap(?)", Capability(99).String())
	}
}

func TestKnownCapabilitiesSorted(t *testing.T) {
	got := KnownCapabilities()
	want := []string{"clock", "fs", "net", "observe", "rand"}
	if len(got) != len(want) {
		t.Fatalf("KnownCapabilities len = %d, want %d (%v)", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("KnownCapabilities[%d] = %q, want %q (must be sorted)", i, got[i], want[i])
		}
	}
}

func TestCheckManyAndManyAudit(t *testing.T) {
	var all Capabilities
	for _, c := range allCaps {
		all.Grant(c)
	}
	// All granted: no error.
	if err := CheckMany(all, allCaps); err != nil {
		t.Fatalf("CheckMany all-granted: %v", err)
	}

	var partial Capabilities
	partial.Grant(CapClock)
	// Missing one capability short-circuits with denial.
	err := CheckMany(partial, allCaps)
	if !errors.Is(err, ErrCapabilityDenied) {
		t.Fatalf("CheckMany partial: want ErrCapabilityDenied, got %v", err)
	}

	// Audited variant records use events for the granted subset.
	log := NewAuditLog()
	if err := CheckManyAudit(partial, []Capability{CapClock}, log, "p"); err != nil {
		t.Fatalf("CheckManyAudit granted subset: %v", err)
	}
	// Audited variant records a deny for the missing capability.
	if err := CheckManyAudit(partial, allCaps, log, "p"); !errors.Is(err, ErrCapabilityDenied) {
		t.Fatalf("CheckManyAudit partial: want ErrCapabilityDenied, got %v", err)
	}
	entries := log.Entries()
	// Call 1 appends one use (CapClock); call 2 appends a use (CapClock) then a
	// deny (CapRand, where it short-circuits) before returning.
	if len(entries) != 3 {
		t.Fatalf("expected 3 audit entries, got %d", len(entries))
	}
	if !entries[0].IsUse() || !entries[0].Result {
		t.Fatalf("entry 0 should be a successful use event, got %+v", entries[0])
	}
	if !entries[1].IsUse() || !entries[1].Result {
		t.Fatalf("entry 1 should be a successful use event, got %+v", entries[1])
	}
	if !entries[2].IsDeny() || entries[2].Result {
		t.Fatalf("entry 2 should be a deny event, got %+v", entries[2])
	}
}

func TestAuditEventPredicates(t *testing.T) {
	grant := AuditEvent{Kind: "grant", PluginID: "p", Capability: CapClock, Result: true}
	deny := AuditEvent{Kind: "deny", PluginID: "p", Capability: CapRand, Result: false}
	use := AuditEvent{Kind: "use", PluginID: "p", Capability: CapFS, Result: true}

	if !grant.IsGrant() || grant.IsDeny() || grant.IsUse() {
		t.Fatalf("grant predicates wrong: %+v", grant)
	}
	if deny.IsGrant() || !deny.IsDeny() || deny.IsUse() {
		t.Fatalf("deny predicates wrong: %+v", deny)
	}
	if use.IsGrant() || use.IsDeny() || !use.IsUse() {
		t.Fatalf("use predicates wrong: %+v", use)
	}
}
