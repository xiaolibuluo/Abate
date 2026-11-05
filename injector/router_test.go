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

type mockInjector struct {
	cap     core.Capability
	injectErr error
	recoverErr error
	injectCount int
	recoverCount int
}

func (m *mockInjector) Capability() core.Capability { return m.cap }

func (m *mockInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	m.injectCount++
	if m.injectErr != nil {
		return core.Injection{}, m.injectErr
	}
	return core.Injection{ID: e.ID, ExperimentID: e.ID}, nil
}

func (m *mockInjector) Recover(_ context.Context, in core.Injection) error {
	m.recoverCount++
	if m.recoverErr != nil {
		return m.recoverErr
	}
	return nil
}

func TestRouter_Register_Success(t *testing.T) {
	r := NewRouter()
	mi := &mockInjector{cap: core.Capability{
		Name:         "mock",
		Version:      "1.0",
		Protocol:     core.ProtocolVersion,
		FaultClasses: []core.FaultClass{core.F1Latency, core.F2Crash},
	}}
	if err := r.Register(mi); err != nil {
		t.Fatalf("Register: %v", err)
	}
	cap := r.Capability()
	if cap.Name != "router" {
		t.Errorf("name = %s", cap.Name)
	}
	if len(cap.FaultClasses) != 2 {
		t.Errorf("fault classes = %d, want 2", len(cap.FaultClasses))
	}
}

func TestRouter_Register_EmptyClasses(t *testing.T) {
	r := NewRouter()
	mi := &mockInjector{cap: core.Capability{
		Name:     "empty",
		Version:  "1.0",
		Protocol: core.ProtocolVersion,
	}}
	if err := r.Register(mi); err == nil {
		t.Fatal("expected error for empty fault classes")
	}
}

func TestRouter_Register_DuplicateClass(t *testing.T) {
	r := NewRouter()
	mi1 := &mockInjector{cap: core.Capability{
		Name:         "mock1",
		Version:      "1.0",
		Protocol:     core.ProtocolVersion,
		FaultClasses: []core.FaultClass{core.F1Latency},
	}}
	mi2 := &mockInjector{cap: core.Capability{
		Name:         "mock2",
		Version:      "1.0",
		Protocol:     core.ProtocolVersion,
		FaultClasses: []core.FaultClass{core.F1Latency},
	}}
	_ = r.Register(mi1)
	if err := r.Register(mi2); err == nil {
		t.Fatal("expected error for duplicate fault class")
	}
}

func TestRouter_Inject_Success(t *testing.T) {
	r := NewRouter()
	mi := &mockInjector{cap: core.Capability{
		Name:         "mock",
		Version:      "1.0",
		Protocol:     core.ProtocolVersion,
		FaultClasses: []core.FaultClass{core.F1Latency},
	}}
	_ = r.Register(mi)

	exp := core.Experiment{ID: "r1", Fault: core.Fault{Class: core.F1Latency, Params: map[string]string{"delay": "50ms"}}}
	inj, err := r.Inject(context.Background(), exp)
	if err != nil {
		t.Fatalf("Inject: %v", err)
	}
	if inj.ID != "r1" {
		t.Errorf("inj.ID = %s", inj.ID)
	}
	if !r.IsActive() {
		t.Error("IsActive should be true")
	}
}

func TestRouter_Inject_NoInjector(t *testing.T) {
	r := NewRouter()
	exp := core.Experiment{ID: "r2", Fault: core.Fault{Class: core.F5Corruption, Params: map[string]string{}}}
	_, err := r.Inject(context.Background(), exp)
	if err == nil {
		t.Fatal("expected error for no injector")
	}
}

func TestRouter_Inject_InjectorError(t *testing.T) {
	r := NewRouter()
	mi := &mockInjector{
		cap: core.Capability{
			Name:         "mock",
			Version:      "1.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F3Partition},
		},
		injectErr: errors.New("inject failed"),
	}
	_ = r.Register(mi)

	_, err := r.Inject(context.Background(), core.Experiment{ID: "r3", Fault: core.Fault{Class: core.F3Partition, Params: map[string]string{}}})
	if err == nil {
		t.Fatal("expected injector error")
	}
}

func TestRouter_Recover_Success(t *testing.T) {
	r := NewRouter()
	mi := &mockInjector{cap: core.Capability{
		Name:         "mock",
		Version:      "1.0",
		Protocol:     core.ProtocolVersion,
		FaultClasses: []core.FaultClass{core.F4Exhaustion},
	}}
	_ = r.Register(mi)

	exp := core.Experiment{ID: "r4", Fault: core.Fault{Class: core.F4Exhaustion, Params: map[string]string{"size": "1M"}}}
	inj, _ := r.Inject(context.Background(), exp)

	if err := r.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover: %v", err)
	}
	if r.IsActive() {
		t.Error("IsActive should be false after recover")
	}
}

func TestRouter_Recover_Idempotent(t *testing.T) {
	r := NewRouter()
	mi := &mockInjector{cap: core.Capability{
		Name:         "mock",
		Version:      "1.0",
		Protocol:     core.ProtocolVersion,
		FaultClasses: []core.FaultClass{core.F6ClockSkew},
	}}
	_ = r.Register(mi)

	exp := core.Experiment{ID: "r5", Fault: core.Fault{Class: core.F6ClockSkew, Params: map[string]string{"offset": "1000"}}}
	inj, _ := r.Inject(context.Background(), exp)

	_ = r.Recover(context.Background(), inj)
	if err := r.Recover(context.Background(), inj); err != nil {
		t.Fatalf("Recover again: %v", err)
	}
}

func TestRouter_Recover_NoActive(t *testing.T) {
	r := NewRouter()
	if err := r.Recover(context.Background(), core.Injection{ID: "x", ExperimentID: "none"}); err != nil {
		t.Errorf("Recover with no active: %v", err)
	}
}

func TestRouter_Recover_InjectorError(t *testing.T) {
	r := NewRouter()
	mi := &mockInjector{
		cap: core.Capability{
			Name:         "mock",
			Version:      "1.0",
			Protocol:     core.ProtocolVersion,
			FaultClasses: []core.FaultClass{core.F2Crash},
		},
		recoverErr: errors.New("recover failed"),
	}
	_ = r.Register(mi)

	exp := core.Experiment{ID: "r6", Fault: core.Fault{Class: core.F2Crash, Params: map[string]string{"signal": "SIGTERM"}}}
	inj, _ := r.Inject(context.Background(), exp)

	if err := r.Recover(context.Background(), inj); err == nil {
		t.Fatal("expected recover error")
	}
}

func TestRouter_IsActive_Empty(t *testing.T) {
	r := NewRouter()
	if r.IsActive() {
		t.Error("IsActive should be false on new router")
	}
}

func TestParseSubCode_Defaults(t *testing.T) {
	e := &core.Experiment{Fault: core.Fault{Class: core.F1Latency}}
	parseSubCode(e)
	if e.Fault.Params["delay"] != "100ms" {
		t.Errorf("delay = %q, want 100ms", e.Fault.Params["delay"])
	}
}

func TestParseSubCode_PreservesExisting(t *testing.T) {
	e := &core.Experiment{Fault: core.Fault{
		Class:  core.F1Latency,
		Params: map[string]string{"delay": "200ms"},
	}}
	parseSubCode(e)
	if e.Fault.Params["delay"] != "200ms" {
		t.Errorf("delay = %q, want 200ms", e.Fault.Params["delay"])
	}
}

func TestParseSubCode_ExactMatch(t *testing.T) {
	e := &core.Experiment{Fault: core.Fault{Class: core.FaultClass("F2.1")}}
	parseSubCode(e)
	if e.Fault.Params["signal"] != "SIGTERM" {
		t.Errorf("signal = %q, want SIGTERM", e.Fault.Params["signal"])
	}
}

func TestParseSubCode_NoMatch(t *testing.T) {
	e := &core.Experiment{Fault: core.Fault{Class: core.FaultClass("F9")}}
	parseSubCode(e)
	if len(e.Fault.Params) != 0 {
		t.Errorf("params = %v, want empty", e.Fault.Params)
	}
}
