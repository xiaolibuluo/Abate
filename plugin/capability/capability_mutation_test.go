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

func TestUnknownCapabilityRejected(t *testing.T) {
	if _, err := ParseGrant("unknown"); !errors.Is(err, ErrUnknownCapability) {
		t.Fatalf("ParseGrant(unknown): want ErrUnknownCapability, got %v", err)
	}
	if _, err := ParseGrant("clock,mystery"); !errors.Is(err, ErrUnknownCapability) {
		t.Fatalf("ParseGrant(clock,mystery): want ErrUnknownCapability, got %v", err)
	}
}

func TestDeclFormatBoundaries(t *testing.T) {
	cases := []struct {
		decl string
		want Capabilities
	}{
		{"clock,", func() (c Capabilities) { c.Grant(CapClock); return c }()},
		{",clock", func() (c Capabilities) { c.Grant(CapClock); return c }()},
		{" clock ", func() (c Capabilities) { c.Grant(CapClock); return c }()},
		{"clock,,rand", func() (c Capabilities) { c.Grant(CapClock); c.Grant(CapRand); return c }()},
		{"\tobsERVE\t", func() (c Capabilities) { return c }()}, // case-sensitive -> no match
	}
	for _, tc := range cases {
		got, err := ParseGrant(tc.decl)
		if tc.decl == "\tobsERVE\t" {
			// "obsERVE" is not a known name -> ErrUnknownCapability.
			if !errors.Is(err, ErrUnknownCapability) {
				t.Fatalf("ParseGrant(%q): want ErrUnknownCapability, got %v", tc.decl, err)
			}
			continue
		}
		if err != nil {
			t.Fatalf("ParseGrant(%q) unexpected error: %v", tc.decl, err)
		}
		if got != tc.want {
			t.Fatalf("ParseGrant(%q) = %d, want %d", tc.decl, got, tc.want)
		}
	}
}

func TestGrantRevokeIdempotent(t *testing.T) {
	var c Capabilities
	c.Grant(CapFS)
	c.Grant(CapFS) // repeat: no change
	if !c.Has(CapFS) {
		t.Fatalf("Grant idempotence broken")
	}
	c.Revoke(CapFS)
	c.Revoke(CapFS) // repeat: no change
	if c.Has(CapFS) {
		t.Fatalf("Revoke idempotence broken")
	}

	// Granting then revoking a single cap leaves others intact.
	var multi Capabilities
	for _, x := range allCaps {
		multi.Grant(x)
	}
	multi.Revoke(CapRand)
	if multi.Has(CapRand) {
		t.Fatalf("Revoke did not remove CapRand")
	}
	for _, x := range []Capability{CapClock, CapObserve, CapFS, CapNet} {
		if !multi.Has(x) {
			t.Fatalf("Revoke of CapRand wrongly affected %s", x)
		}
	}
}

func TestAuditAppendDeterministic(t *testing.T) {
	log := NewAuditLog()
	const n = 50
	for i := 0; i < n; i++ {
		log.Append(AuditEvent{Kind: "use", PluginID: "p", Capability: CapClock, Result: true})
	}
	entries := log.Entries()
	if len(entries) != n {
		t.Fatalf("expected %d entries, got %d", n, len(entries))
	}
	// At is the virtual monotonic counter, assigned 1..n deterministically.
	for i, e := range entries {
		if e.At != int64(i+1) {
			t.Fatalf("entry %d: At=%d, want %d", i, e.At, i+1)
		}
	}
	// Entries() returns a copy; mutating it must not affect the log.
	entries[0].Kind = "tampered"
	if log.Entries()[0].Kind == "tampered" {
		t.Fatalf("Entries() did not return a copy")
	}
}
