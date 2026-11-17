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

// allCaps is the ordered list of the five known capabilities.
var allCaps = []Capability{CapClock, CapRand, CapObserve, CapFS, CapNet}

func TestCapabilityMatrix5x2(t *testing.T) {
	// 5 capabilities x 2 states (granted / denied) must hold for every bit.
	for _, c := range allCaps {
		var granted Capabilities
		granted.Grant(c)
		if !granted.Has(c) {
			t.Fatalf("%s: granted capability missing from Has", c)
		}
		var zero Capabilities
		if zero.Has(c) {
			t.Fatalf("%s: zero capability set must deny", c)
		}
	}
	// All-granted set reports every capability present.
	var all Capabilities
	for _, c := range allCaps {
		all.Grant(c)
	}
	for _, c := range allCaps {
		if !all.Has(c) {
			t.Fatalf("%s: present in all-granted set", c)
		}
	}
	if all.IsZero() {
		t.Fatalf("all-granted set must not be zero")
	}
}

func TestParseGrantLegalIllegalUnknown(t *testing.T) {
	legal := []struct {
		decl string
		want Capabilities
	}{
		{"clock", func() (c Capabilities) { c.Grant(CapClock); return c }()},
		{"clock,rand", func() (c Capabilities) { c.Grant(CapClock); c.Grant(CapRand); return c }()},
		{"observe,fs,net", func() (c Capabilities) { c.Grant(CapObserve); c.Grant(CapFS); c.Grant(CapNet); return c }()},
		{"clock,rand,observe,fs,net", func() (c Capabilities) {
			for _, x := range allCaps {
				c.Grant(x)
			}
			return c
		}()},
		{" clock , rand ", func() (c Capabilities) { c.Grant(CapClock); c.Grant(CapRand); return c }()},
		{"clock,", func() (c Capabilities) { c.Grant(CapClock); return c }()},
		{",clock", func() (c Capabilities) { c.Grant(CapClock); return c }()},
	}
	for _, tc := range legal {
		got, err := ParseGrant(tc.decl)
		if err != nil {
			t.Fatalf("ParseGrant(%q) unexpected error: %v", tc.decl, err)
		}
		if got != tc.want {
			t.Fatalf("ParseGrant(%q) = %d, want %d", tc.decl, got, tc.want)
		}
	}

	// Empty / whitespace-only declaration -> zero mask, no error (default deny).
	for _, empty := range []string{"", "   ", ","} {
		got, err := ParseGrant(empty)
		if err != nil {
			t.Fatalf("ParseGrant(%q) unexpected error: %v", empty, err)
		}
		if !got.IsZero() {
			t.Fatalf("ParseGrant(%q) should yield zero mask", empty)
		}
	}

	// Unknown capability names reject the whole declaration.
	unknown := []string{"foo", "clock,foo", "CLOCK"}
	for _, decl := range unknown {
		if _, err := ParseGrant(decl); !errors.Is(err, ErrUnknownCapability) {
			t.Fatalf("ParseGrant(%q): want ErrUnknownCapability, got %v", decl, err)
		}
	}
}

func TestOverprivilegePerClassIntercepted(t *testing.T) {
	log := NewAuditLog()
	for _, c := range allCaps {
		var zero Capabilities
		err := CheckAudit(zero, c, log, "p/"+c.String())
		if !errors.Is(err, ErrCapabilityDenied) {
			t.Fatalf("%s: unprivileged call should be denied, got %v", c, err)
		}
	}
	entries := log.Entries()
	if len(entries) != len(allCaps) {
		t.Fatalf("expected %d deny entries, got %d", len(allCaps), len(entries))
	}
	for _, e := range entries {
		if !e.IsDeny() || e.Result {
			t.Fatalf("expected deny event with Result=false, got %+v", e)
		}
	}
}

func TestAuditThreeKindsCovered(t *testing.T) {
	log := NewAuditLog()

	var canUse Capabilities
	canUse.Grant(CapClock)
	if err := CheckAudit(canUse, CapClock, log, "p1"); err != nil {
		t.Fatalf("use check: %v", err)
	}

	var zero Capabilities
	if err := CheckAudit(zero, CapRand, log, "p1"); !errors.Is(err, ErrCapabilityDenied) {
		t.Fatalf("deny check: %v", err)
	}

	log.Append(AuditEvent{Kind: "grant", PluginID: "p1", Capability: CapFS, Result: true})

	entries := log.Entries()
	if len(entries) != 3 {
		t.Fatalf("expected 3 events, got %d", len(entries))
	}
	kinds := map[string]int{}
	for _, e := range entries {
		kinds[e.Kind]++
	}
	if kinds["use"] != 1 || kinds["deny"] != 1 || kinds["grant"] != 1 {
		t.Fatalf("expected one of each kind, got %v", kinds)
	}
	// Virtual monotonic At must be strictly increasing.
	for i := 1; i < len(entries); i++ {
		if entries[i].At <= entries[i-1].At {
			t.Fatalf("audit At not monotonic increasing: %d then %d", entries[i-1].At, entries[i].At)
		}
	}
}
