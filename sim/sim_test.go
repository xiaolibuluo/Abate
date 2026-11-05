// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package sim

import (
	"context"
	"testing"

	"github.com/qingdujian/abate/core"
)

func testExperiment() core.Experiment {
	return core.Experiment{
		ID:          "sim-test",
		Seed:        42,
		Criticality: 3,
		Target: core.Target{
			Kind: core.TargetService,
			Name: "test-svc",
		},
		Fault: core.Fault{
			Class:  core.F1Latency,
			Params: map[string]string{"delay": "100"},
		},
		Blast: core.BlastRadius{
			MaxInstances:  1,
			MaxDurationMs: 60000,
		},
		Steady: []core.SLIAssert{
			{SLI: "latency", Assert: "< 500"},
		},
	}
}

func TestModeValid(t *testing.T) {
	tests := []struct {
		mode Mode
		want bool
	}{
		{ModeSimulation, true},
		{ModeStaging, true},
		{ModeProduction, true},
		{Mode("bogus"), false},
	}
	for _, tt := range tests {
		if got := tt.mode.Valid(); got != tt.want {
			t.Errorf("mode %q: valid = %v, want %v", tt.mode, got, tt.want)
		}
	}
}

func TestSimulatorRun(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	r, err := s.Run(testExperiment())
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if r.Exp.ID != "sim-test" {
		t.Errorf("exp ID = %q, want sim-test", r.Exp.ID)
	}
	if r.WindowMs != 60000 {
		t.Errorf("window = %d, want 60000", r.WindowMs)
		return
	}
	if r.RecoveryMs != 60000 {
		t.Errorf("recovery = %d, want 60000", r.RecoveryMs)
		return
	}
}

func TestSimulatorRunPreflightFail(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	bad := testExperiment()
	bad.ID = ""
	_, err := s.Run(bad)
	if err == nil {
		t.Error("expected error for empty ID")
	}
}

func TestSimulatorDeterministicReplay(t *testing.T) {
	e := testExperiment()
	s1 := NewSimulator(ModeSimulation, 42)
	r1, err := s1.Run(e)
	if err != nil {
		t.Fatalf("run1: %v", err)
	}
	s2 := NewSimulator(ModeSimulation, 42)
	r2, err := s2.Run(e)
	if err != nil {
		t.Fatalf("run2: %v", err)
	}
	if r1.HoldMs != r2.HoldMs {
		t.Errorf("holdMs: %d vs %d", r1.HoldMs, r2.HoldMs)
	}
}

func TestVirtualInjectorIdempotent(t *testing.T) {
	v := &VirtualInjector{rng: core.NewRNG(1)}
	e := testExperiment()
	i1, err := v.Inject(context.Background(), e)
	if err != nil {
		t.Fatalf("inject1: %v", err)
	}
	i2, err := v.Inject(context.Background(), e)
	if err != nil {
		t.Fatalf("inject2: %v", err)
	}
	if i1.ID != i2.ID {
		t.Errorf("idempotent: %q vs %q", i1.ID, i2.ID)
		return
	}
	if err := v.Recover(context.Background(), i1); err != nil {
		t.Errorf("recover: %v", err)
		return
	}
	if v.IsActive() {
		t.Error("should be inactive after recover")
	}
}

func TestVirtualInjectorCapability(t *testing.T) {
	v := &VirtualInjector{rng: core.NewRNG(1)}
	c := v.Capability()
	if c.Name != "sim-injector" {
		t.Errorf("cap name = %q", c.Name)
		return
	}
	if err := c.CheckProtocol(); err != nil {
		t.Errorf("protocol check: %v", err)
	}
}

func TestVirtualProbeObserve(t *testing.T) {
	p := &VirtualProbe{rng: core.NewRNG(1)}
	tgt := core.Target{Kind: core.TargetService, Name: "svc"}
	s1, err := p.Observe(context.Background(), tgt)
	if err != nil {
		t.Fatalf("observe1: %v", err)
	}
	if s1.SLI != "svc" {
		t.Errorf("sli = %q", s1.SLI)
		return
	}
	p2 := &VirtualProbe{rng: core.NewRNG(1)}
	s2, _ := p2.Observe(context.Background(), tgt)
	if s1.Value != s2.Value {
		t.Errorf("deterministic: %d vs %d", s1.Value, s2.Value)
	}
}

func TestVirtualProbeCapability(t *testing.T) {
	p := &VirtualProbe{rng: core.NewRNG(1)}
	c := p.Capability()
	if c.Name != "sim-probe" {
		t.Errorf("cap name = %q", c.Name)
	}
}

func TestVirtualValidatorPreflight(t *testing.T) {
	v := &VirtualValidator{}
	issues, err := v.Preflight(context.Background(), testExperiment())
	if err != nil {
		t.Fatalf("preflight: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("issues = %v", issues)
		return
	}
	bad := testExperiment()
	bad.ID = ""
	issues, _ = v.Preflight(context.Background(), bad)
	if len(issues) == 0 {
		t.Error("expected issues for empty ID")
		return
	}
}

func TestVirtualValidatorPostflight(t *testing.T) {
	v := &VirtualValidator{}
	issues, err := v.Postflight(context.Background(), testExperiment())
	if err != nil {
		t.Fatalf("postflight: %v", err)
	}
	if len(issues) != 0 {
		t.Errorf("issues = %v", issues)
		return
	}
	v.SetResidue(true)
	issues, _ = v.Postflight(context.Background(), testExperiment())
	if len(issues) == 0 {
		t.Error("expected issues when residue set")
		return
	}
}

func TestVirtualValidatorCapability(t *testing.T) {
	v := &VirtualValidator{}
	c := v.Capability()
	if c.Name != "sim-validator" {
		t.Errorf("cap name = %q", c.Name)
	}
}
