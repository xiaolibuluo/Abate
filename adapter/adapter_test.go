// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package adapter

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/qingdujian/abate/core"
)

// mockBackend 是测试用的 Backend 替身。
type mockBackend struct {
	applyErr error
	cleanErr error
	applied  []MappedFault
	cleaned []string
	nextID   int
}

func (m *mockBackend) Apply(mapped MappedFault) (string, error) {
	if m.applyErr != nil {
		return "", m.applyErr
	}
	m.applied = append(m.applied, mapped)
	m.nextID++
	return fmt.Sprintf("upstream-%d", m.nextID), nil
}

func (m *mockBackend) Cleanup(upstreamID string) error {
	if m.cleanErr != nil {
		return m.cleanErr
	}
	m.cleaned = append(m.cleaned, upstreamID)
	return nil
}

func validExperiment() core.Experiment {
	return core.Experiment{
		ID:          "exp-test",
		Seed:        42,
		Criticality: 3,
		Target:      core.Target{Kind: core.TargetService, Name: "svc-a"},
		Fault:       core.Fault{Class: core.FaultClass("F1.1"), Params: map[string]string{"delay": "100"}},
		Blast:       core.BlastRadius{MaxInstances: 1, MaxDurationMs: 60000},
		Steady:      []core.SLIAssert{{SLI: "latency", Assert: "< 500ms"}},
	}
}

func TestNew_Success(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, err := New("chaosmesh", "0.1.0", mb, mt)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if a.cap.Name != "chaosmesh" {
		t.Errorf("name = %q, want chaosmesh", a.cap.Name)
	}
	if a.cap.Protocol != core.ProtocolVersion {
		t.Errorf("protocol = %q, want %q", a.cap.Protocol, core.ProtocolVersion)
	}
	if len(a.cap.FaultClasses) == 0 {
		t.Error("FaultClasses should not be empty")
	}
}

func TestNew_NilBackend(t *testing.T) {
	mt := NewMappingTable(DefaultChaosMeshMappings())
	_, err := New("x", "1", nil, mt)
	if err == nil {
		t.Fatal("expected error for nil backend")
	}
}

func TestNew_NilTable(t *testing.T) {
	mb := &mockBackend{}
	_, err := New("x", "1", mb, nil)
	if err == nil {
		t.Fatal("expected error for nil table")
	}
}

func TestCapability(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	got := a.Capability()
	if got.Name != "cm" {
		t.Errorf("name = %q", got.Name)
	}
	if got.Version != "1.0" {
		t.Errorf("version = %q", got.Version)
	}
}

func TestInject_Success(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	inj, err := a.Inject(ctx, validExperiment())
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID != "upstream-1" {
		t.Errorf("inj.ID = %q, want upstream-1", inj.ID)
	}
	if inj.ExperimentID != "exp-test" {
		t.Errorf("inj.ExperimentID = %q", inj.ExperimentID)
	}
	if len(mb.applied) != 1 {
		t.Fatalf("applied count = %d, want 1", len(mb.applied))
	}
	if mb.applied[0].UpstreamType != "IOChaos/delay" {
		t.Errorf("upstreamType = %q, want IOChaos/delay", mb.applied[0].UpstreamType)
	}
	if mb.applied[0].Params["latency"] != "100" {
		t.Errorf("params[latency] = %q, want 100", mb.applied[0].Params["latency"])
	}
}

func TestInject_Idempotent(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	exp := validExperiment()
	inj1, _ := a.Inject(ctx, exp)
	inj2, err := a.Inject(ctx, exp)
	if err != nil {
		t.Fatalf("second Inject: %v", err)
	}
	if inj1.ID != inj2.ID {
		t.Errorf("idempotent: inj1=%q inj2=%q", inj1.ID, inj2.ID)
	}
	if len(mb.applied) != 1 {
		t.Errorf("applied count = %d, want 1 (idempotent)", len(mb.applied))
	}
}

func TestInject_TranslateError(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	exp := validExperiment()
	exp.Fault.Class = core.F5Corruption
	_, err := a.Inject(ctx, exp)
	if err == nil {
		t.Fatal("expected translate error for unmapped F5")
	}
}

func TestInject_BackendError(t *testing.T) {
	mb := &mockBackend{applyErr: errors.New("cluster unreachable")}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	_, err := a.Inject(ctx, validExperiment())
	if err == nil {
		t.Fatal("expected backend error")
	}
}

func TestRecover_Success(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	inj, _ := a.Inject(ctx, validExperiment())
	if err := a.Recover(ctx, inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if len(mb.cleaned) != 1 || mb.cleaned[0] != "upstream-1" {
		t.Errorf("cleaned = %v, want [upstream-1]", mb.cleaned)
	}
}

func TestRecover_Idempotent(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	inj, _ := a.Inject(ctx, validExperiment())
	_ = a.Recover(ctx, inj)
	if err := a.Recover(ctx, inj); err != nil {
		t.Errorf("second Recover: %v", err)
	}
}

func TestRecover_NoActive(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	if err := a.Recover(ctx, core.Injection{ID: "x", ExperimentID: "none"}); err != nil {
		t.Errorf("Recover with no active: %v", err)
	}
}

func TestRecover_BackendError(t *testing.T) {
	mb := &mockBackend{cleanErr: errors.New("cleanup failed")}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	inj, _ := a.Inject(ctx, validExperiment())
	err := a.Recover(ctx, inj)
	if err == nil {
		t.Fatal("expected backend cleanup error")
	}
}

func TestAdapter_Degradations(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable(DefaultChaosMeshMappings())
	a, _ := New("cm", "1.0", mb, mt)
	degs := a.Degradations()
	if len(degs) == 0 {
		t.Fatal("expected degradations from adapter")
	}
}

func TestInject_TranslatePassthroughParam(t *testing.T) {
	mb := &mockBackend{}
	mt := NewMappingTable([]FaultMapping{
		{SubCode: "F1", UpstreamType: "Delay", ParamMap: map[string]string{"delay": "latency"}},
	})
	a, _ := New("cm", "1.0", mb, mt)
	ctx := context.Background()
	exp := validExperiment()
	exp.Fault.Class = core.FaultClass("F1")
	exp.Fault.Params = map[string]string{"delay": "100", "unknown": "val"}
	inj, err := a.Inject(ctx, exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID == "" {
		t.Error("inj.ID should not be empty")
	}
	if len(mb.applied) != 1 {
		t.Fatalf("applied = %d, want 1", len(mb.applied))
	}
	if mb.applied[0].UpstreamType != "Delay" {
		t.Errorf("type = %q", mb.applied[0].UpstreamType)
	}
	if mb.applied[0].Params["latency"] != "100" {
		t.Errorf("latency = %q", mb.applied[0].Params["latency"])
	}
	if mb.applied[0].Params["unknown"] != "val" {
		t.Errorf("unknown passthrough = %q", mb.applied[0].Params["unknown"])
	}
}
