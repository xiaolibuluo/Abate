// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package plugin

import (
	"testing"

	"github.com/qingdujian/abate/core"
)

func TestRegisterAndLookup(t *testing.T) {
	r := NewRegistry()
	inj := &mockInjector{cap: validCap("svc-inj")}
	if err := r.Register(Plugin{Injector: inj}); err != nil {
		t.Fatalf("register: %v", err)
	}
	e, ok := r.Lookup("svc-inj")
	if !ok {
		t.Fatal("lookup failed")
	}
	if e.cap.Name != "svc-inj" {
		t.Errorf("name = %q", e.cap.Name)
	}
	if e.inj == nil {
		t.Error("injector is nil")
	}
}

func TestRegisterDuplicate(t *testing.T) {
	r := NewRegistry()
	inj := &mockInjector{cap: validCap("dup")}
	if err := r.Register(Plugin{Injector: inj}); err != nil {
		t.Fatalf("register1: %v", err)
	}
	if err := r.Register(Plugin{Injector: inj}); err == nil {
		t.Error("expected duplicate error")
	}
}

func TestRegisterNoRole(t *testing.T) {
	r := NewRegistry()
	if err := r.Register(Plugin{}); err == nil {
		t.Error("expected error for no role")
	}
}

func TestRegisterBadProtocol(t *testing.T) {
	r := NewRegistry()
	cap := validCap("bad")
	cap.Protocol = "v0"
	inj := &mockInjector{cap: cap}
	if err := r.Register(Plugin{Injector: inj}); err == nil {
		t.Error("expected protocol mismatch error")
	}
}

func TestUnregister(t *testing.T) {
	r := NewRegistry()
	inj := &mockInjector{cap: validCap("rm")}
	_ = r.Register(Plugin{Injector: inj})
	if err := r.Unregister("rm"); err != nil {
		t.Fatalf("unregister: %v", err)
	}
	if _, ok := r.Lookup("rm"); ok {
		t.Error("should be removed")
	}
}

func TestUnregisterNotFound(t *testing.T) {
	r := NewRegistry()
	if err := r.Unregister("ghost"); err == nil {
		t.Error("expected not found error")
	}
}

func TestList(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Injector: &mockInjector{cap: validCap("zeta")}})
	_ = r.Register(Plugin{Probe: &mockProbe{cap: validCap("alpha")}})
	names := r.List()
	if len(names) != 2 {
		t.Fatalf("list len = %d, want 2", len(names))
	}
	if names[0] != "alpha" || names[1] != "zeta" {
		t.Errorf("order = %v", names)
	}
}

func TestFindInjector(t *testing.T) {
	r := NewRegistry()
	cap := validCap("f1-inj")
	cap.FaultClasses = []core.FaultClass{core.F1Latency, core.F2Crash}
	cap.TargetKinds = []core.TargetKind{core.TargetService}
	_ = r.Register(Plugin{Injector: &mockInjector{cap: cap}})
	got, err := r.FindInjector(core.F1Latency, core.TargetService)
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if got == nil {
		t.Error("injector is nil")
	}
}

func TestFindInjectorByFaultClass(t *testing.T) {
	r := NewRegistry()
	cap := validCap("f2-inj")
	cap.FaultClasses = []core.FaultClass{core.F2Crash}
	cap.TargetKinds = []core.TargetKind{core.TargetService}
	_ = r.Register(Plugin{Injector: &mockInjector{cap: cap}})
	if _, err := r.FindInjector(core.F1Latency, core.TargetService); err == nil {
		t.Error("expected no injector for F1")
	}
}

func TestFindInjectorByTargetKind(t *testing.T) {
	r := NewRegistry()
	cap := validCap("svc-only")
	cap.FaultClasses = []core.FaultClass{core.F1Latency}
	cap.TargetKinds = []core.TargetKind{core.TargetService}
	_ = r.Register(Plugin{Injector: &mockInjector{cap: cap}})
	if _, err := r.FindInjector(core.F1Latency, core.TargetInstance); err == nil {
		t.Error("expected no injector for instance target")
	}
}

func TestFindInjectorNone(t *testing.T) {
	r := NewRegistry()
	if _, err := r.FindInjector(core.F1Latency, core.TargetService); err == nil {
		t.Error("expected no injector error")
	}
}

func TestFindProbe(t *testing.T) {
	r := NewRegistry()
	cap := validCap("probe-1")
	cap.TargetKinds = []core.TargetKind{core.TargetService}
	_ = r.Register(Plugin{Probe: &mockProbe{cap: cap}})
	got, err := r.FindProbe(core.TargetService)
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if got == nil {
		t.Error("probe is nil")
	}
}

func TestFindProbeNone(t *testing.T) {
	r := NewRegistry()
	if _, err := r.FindProbe(core.TargetService); err == nil {
		t.Error("expected no probe error")
	}
}

func TestFindValidator(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Validator: &mockValidator{cap: validCap("val-1")}})
	got, err := r.FindValidator()
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if got == nil {
		t.Error("validator is nil")
	}
}

func TestFindValidatorNone(t *testing.T) {
	r := NewRegistry()
	if _, err := r.FindValidator(); err == nil {
		t.Error("expected no validator error")
	}
}

func TestFindReporter(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Reporter: &mockReporter{cap: validCap("rep-1"), emitLoc: "/tmp/r.json"}})
	got, err := r.FindReporter()
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if got == nil {
		t.Error("reporter is nil")
	}
}

func TestFindReporterNone(t *testing.T) {
	r := NewRegistry()
	if _, err := r.FindReporter(); err == nil {
		t.Error("expected no reporter error")
	}
}

func TestMultiRolePlugin(t *testing.T) {
	r := NewRegistry()
	cap := validCap("multi")
	inj := &mockInjector{cap: cap}
	val := &mockValidator{cap: cap}
	if err := r.Register(Plugin{Injector: inj, Validator: val}); err != nil {
		t.Fatalf("register: %v", err)
	}
	e, ok := r.Lookup("multi")
	if !ok {
		t.Fatal("lookup failed")
	}
	roles := e.roles()
	if len(roles) != 2 {
		t.Errorf("roles = %v, want 2", roles)
	}
	if roles[0] != RoleInjector || roles[1] != RoleValidator {
		t.Errorf("roles = %v", roles)
	}
}

func TestRolesProbeOnly(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Probe: &mockProbe{cap: validCap("probe-only")}})
	e, _ := r.Lookup("probe-only")
	roles := e.roles()
	if len(roles) != 1 || roles[0] != RoleProbe {
		t.Errorf("roles = %v, want [probe]", roles)
	}
}

func TestRolesReporterOnly(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Reporter: &mockReporter{cap: validCap("rep-only"), emitLoc: "/tmp/r.json"}})
	e, _ := r.Lookup("rep-only")
	roles := e.roles()
	if len(roles) != 1 || roles[0] != RoleReporter {
		t.Errorf("roles = %v, want [reporter]", roles)
	}
}

func TestFindInjectorSkipsNonInjector(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Probe: &mockProbe{cap: validCap("non-inj")}})
	_ = r.Register(Plugin{Injector: &mockInjector{cap: validCap("real-inj")}})
	got, err := r.FindInjector(core.F1Latency, core.TargetService)
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if got == nil {
		t.Error("injector is nil")
	}
}

func TestFindProbeSkipsNonProbe(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Injector: &mockInjector{cap: validCap("non-probe")}})
	_ = r.Register(Plugin{Probe: &mockProbe{cap: validCap("real-probe")}})
	got, err := r.FindProbe(core.TargetService)
	if err != nil {
		t.Fatalf("find: %v", err)
	}
	if got == nil {
		t.Error("probe is nil")
	}
}

func TestFindProbeSkipsWrongTarget(t *testing.T) {
	r := NewRegistry()
	cap := validCap("inst-probe")
	cap.TargetKinds = []core.TargetKind{core.TargetInstance}
	_ = r.Register(Plugin{Probe: &mockProbe{cap: cap}})
	if _, err := r.FindProbe(core.TargetService); err == nil {
		t.Error("expected no probe for service target")
	}
}

func TestHandshakeInjectorBadProtocol(t *testing.T) {
	cap := validCap("bad-proto-inj")
	cap.Protocol = "v0"
	if err := HandshakeInjector(cap); err == nil {
		t.Error("expected handshake error")
	}
}

func TestFindProbeOnlyNonProbe(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Injector: &mockInjector{cap: validCap("inj-no-probe")}})
	if _, err := r.FindProbe(core.TargetService); err == nil {
		t.Error("expected no probe error")
	}
}

func TestFindInjectorOnlyNonInjector(t *testing.T) {
	r := NewRegistry()
	_ = r.Register(Plugin{Probe: &mockProbe{cap: validCap("probe-no-inj")}})
	if _, err := r.FindInjector(core.F1Latency, core.TargetService); err == nil {
		t.Error("expected no injector error")
	}
}
