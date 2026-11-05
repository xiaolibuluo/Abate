// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package injector

import (
	"context"
	"errors"
	"testing"

	"github.com/qingdujian/abate/core"
)

func validProcessExperiment() core.Experiment {
	return core.Experiment{
		ID:          "exp-proc",
		Seed:        1,
		Criticality: 3,
		Target:      core.Target{Kind: core.TargetInstance, Name: "test-svc"},
		Fault:       core.Fault{Class: core.F2Crash, Params: map[string]string{"signal": "TERM"}},
		Blast:       core.BlastRadius{MaxInstances: 1, MaxDurationMs: 30000},
		Steady:      []core.SLIAssert{{SLI: "availability", Assert: ">= 99%"}},
	}
}

func TestProcessInjector_Capability(t *testing.T) {
	p := NewProcessInjector()
	cap := p.Capability()
	if cap.Name != "process-injector" {
		t.Errorf("name = %q", cap.Name)
	}
	if cap.Protocol != core.ProtocolVersion {
		t.Errorf("protocol = %q", cap.Protocol)
	}
}

func TestProcessInjector_Inject_Success(t *testing.T) {
	old := killFn
	killFn = func(signal, pattern string) error { return nil }
	defer func() { killFn = old }()

	p := NewProcessInjector()
	inj, err := p.Inject(context.Background(), validProcessExperiment())
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID != "exp-proc" {
		t.Errorf("inj.ID = %q", inj.ID)
	}
	if !p.IsActive() {
		t.Error("should be active")
	}
}

func TestProcessInjector_Inject_Idempotent(t *testing.T) {
	calls := 0
	old := killFn
	killFn = func(signal, pattern string) error {
		calls++
		return nil
	}
	defer func() { killFn = old }()

	p := NewProcessInjector()
	exp := validProcessExperiment()
	_, _ = p.Inject(context.Background(), exp)
	_, err := p.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("second Inject: %v", err)
	}
	if calls != 1 {
		t.Errorf("kill called %d times, want 1", calls)
	}
}

func TestProcessInjector_Inject_EmptyTarget(t *testing.T) {
	old := killFn
	killFn = func(signal, pattern string) error { return nil }
	defer func() { killFn = old }()

	p := NewProcessInjector()
	exp := validProcessExperiment()
	exp.Target.Name = ""
	_, err := p.Inject(context.Background(), exp)
	if err == nil {
		t.Fatal("expected error for empty target")
	}
}

func TestProcessInjector_Inject_KillError(t *testing.T) {
	old := killFn
	killFn = func(signal, pattern string) error { return errors.New("permission denied") }
	defer func() { killFn = old }()

	p := NewProcessInjector()
	_, err := p.Inject(context.Background(), validProcessExperiment())
	if err == nil {
		t.Fatal("expected kill error")
	}
}

func TestProcessInjector_Inject_DefaultSignal(t *testing.T) {
	var captured string
	old := killFn
	killFn = func(signal, pattern string) error {
		captured = signal
		return nil
	}
	defer func() { killFn = old }()

	p := NewProcessInjector()
	exp := validProcessExperiment()
	exp.Fault.Params = map[string]string{"x": "y"}
	_, err := p.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if captured != "TERM" {
		t.Errorf("signal = %q, want TERM", captured)
	}
}

func TestProcessInjector_Recover(t *testing.T) {
	old := killFn
	killFn = func(signal, pattern string) error { return nil }
	defer func() { killFn = old }()

	p := NewProcessInjector()
	inj, _ := p.Inject(context.Background(), validProcessExperiment())
	if err := p.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if p.IsActive() {
		t.Error("should not be active after recover")
	}
}

func TestProcessInjector_Recover_Idempotent(t *testing.T) {
	p := NewProcessInjector()
	if err := p.Recover(context.Background(), core.Injection{ID: "none"}); err != nil {
		t.Errorf("Recover with no active: %v", err)
	}
}

func TestDefaultKill(t *testing.T) {
	err := defaultKill("TERM", "nonexistent-process-xyz-12345")
	if err == nil {
		t.Log("pkill returned nil (unexpected but not fatal)")
	}
}