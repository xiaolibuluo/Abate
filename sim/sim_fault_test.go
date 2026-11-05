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
	"fmt"
	"testing"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/engine"
)

var errSimTest = fmt.Errorf("sim test error")

// failingInjector 注入时返回错误，用于覆盖 Run 的错误分支。
type failingInjector struct {
	failAt string
}

func (f *failingInjector) Capability() core.Capability {
	return core.Capability{Name: "fail", Protocol: core.ProtocolVersion}
}

func (f *failingInjector) Inject(_ context.Context, _ core.Experiment) (core.Injection, error) {
	if f.failAt == "inject" {
		return core.Injection{}, errSimTest
	}
	return core.Injection{ID: "ok"}, nil
}

func (f *failingInjector) Recover(_ context.Context, _ core.Injection) error {
	if f.failAt == "recover" {
		return errSimTest
	}
	return nil
}

// failingProbe 采样时返回错误。
type failingProbe struct{}

func (f *failingProbe) Capability() core.Capability {
	return core.Capability{Name: "fail-probe", Protocol: core.ProtocolVersion}
}

func (f *failingProbe) Observe(_ context.Context, _ core.Target) (core.Sample, error) {
	return core.Sample{}, errSimTest
}

// failingValidator 在 Preflight 时返回错误。
type failingValidator struct{}

func (f *failingValidator) Capability() core.Capability {
	return core.Capability{Name: "fail-val", Protocol: core.ProtocolVersion}
}

func (f *failingValidator) Preflight(_ context.Context, _ core.Experiment) ([]string, error) {
	return nil, errSimTest
}

func (f *failingValidator) Postflight(_ context.Context, _ core.Experiment) ([]string, error) {
	return nil, nil
}

func (f *failingValidator) SetResidue(_ bool) {}

// issuesValidator 在 Preflight 时返回非空 issues 但无 error。
type issuesValidator struct{}

func (i *issuesValidator) Capability() core.Capability {
	return core.Capability{Name: "issues-val", Protocol: core.ProtocolVersion}
}

func (i *issuesValidator) Preflight(_ context.Context, _ core.Experiment) ([]string, error) {
	return []string{"issue-1"}, nil
}

func (i *issuesValidator) Postflight(_ context.Context, _ core.Experiment) ([]string, error) {
	return nil, nil
}

func (i *issuesValidator) SetResidue(_ bool) {}

func TestSimulatorRunInjectFail(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	s.inj = &failingInjector{failAt: "inject"}
	_, err := s.Run(testExperiment())
	if err == nil {
		t.Error("expected inject error")
	}
}

func TestSimulatorRunRecoverFail(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	s.inj = &failingInjector{failAt: "recover"}
	_, err := s.Run(testExperiment())
	if err == nil {
		t.Error("expected recover error")
	}
}

func TestSimulatorRunObserveFail(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	s.probe = &failingProbe{}
	_, err := s.Run(testExperiment())
	if err == nil {
		t.Error("expected observe error")
	}
}

func TestSimulatorRunPreflightIssues(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	s.val = &VirtualValidator{residue: false}
	bad := testExperiment()
	bad.Target = core.Target{Kind: "bogus", Name: ""}
	_, err := s.Run(bad)
	if err == nil {
		t.Error("expected preflight error")
	}
}

func TestSimulatorRunPreflightError(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	s.val = &failingValidator{}
	_, err := s.Run(testExperiment())
	if err == nil {
		t.Error("expected preflight error")
	}
}

func TestSimulatorRunPreflightIssuesNonEmpty(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	s.val = &issuesValidator{}
	_, err := s.Run(testExperiment())
	if err == nil {
		t.Error("expected preflight issues error")
	}
}

func TestSimulatorRunDefaultWindow(t *testing.T) {
	s := NewSimulator(ModeSimulation, 42)
	e := testExperiment()
	e.Blast.MaxDurationMs = 0
	e.Blast.MaxTraffic = 1
	r, err := s.Run(e)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if r.WindowMs != 60000 {
		t.Errorf("window = %d, want 60000", r.WindowMs)
	}
}

func TestSimulatorRunTransitFail(t *testing.T) {
	orig := transitFn
	defer func() { transitFn = orig }()
	transitFn = func(_, _ engine.State) (engine.State, error) {
		return "", errSimTest
	}
	s := NewSimulator(ModeSimulation, 42)
	_, err := s.Run(testExperiment())
	if err == nil {
		t.Error("expected transit error")
	}
}

func TestSimulatorRunTransitFailAtN(t *testing.T) {
	orig := transitFn
	defer func() { transitFn = orig }()
	for failAt := 1; failAt <= 5; failAt++ {
		callCount := 0
		transitFn = func(_, _ engine.State) (engine.State, error) {
			callCount++
			if callCount == failAt {
				return "", errSimTest
			}
			return engine.Transition(engine.StatePending, engine.StatePreflight)
		}
		s := NewSimulator(ModeSimulation, 42)
		_, err := s.Run(testExperiment())
		if err == nil {
			t.Errorf("failAt %d: expected transit error", failAt)
		}
	}
}
