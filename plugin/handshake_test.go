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
	"context"
	"fmt"
	"testing"

	"github.com/qingdujian/abate/core"
)

var errTest = fmt.Errorf("test error")

// mockInjector 实现 core.Injector 用于测试。
type mockInjector struct {
	cap       core.Capability
	injectErr error
	recoverErr error
}

func (m *mockInjector) Capability() core.Capability { return m.cap }
func (m *mockInjector) Inject(_ context.Context, e core.Experiment) (core.Injection, error) {
	if m.injectErr != nil {
		return core.Injection{}, m.injectErr
	}
	return core.Injection{ID: "inj-1", ExperimentID: e.ID}, nil
}
func (m *mockInjector) Recover(_ context.Context, _ core.Injection) error {
	return m.recoverErr
}

// mockProbe 实现 core.Probe 用于测试。
type mockProbe struct {
	cap        core.Capability
	observeErr error
}

func (m *mockProbe) Capability() core.Capability { return m.cap }
func (m *mockProbe) Observe(_ context.Context, t core.Target) (core.Sample, error) {
	if m.observeErr != nil {
		return core.Sample{}, m.observeErr
	}
	return core.Sample{SLI: t.Name, Value: 9900, TimestampMs: 1000}, nil
}

// mockValidator 实现 core.Validator 用于测试。
type mockValidator struct {
	cap           core.Capability
	preflightErr  error
	postflightErr error
	issues        []string
}

func (m *mockValidator) Capability() core.Capability { return m.cap }
func (m *mockValidator) Preflight(_ context.Context, _ core.Experiment) ([]string, error) {
	return m.issues, m.preflightErr
}
func (m *mockValidator) Postflight(_ context.Context, _ core.Experiment) ([]string, error) {
	return nil, m.postflightErr
}
func (m *mockValidator) SetResidue(_ bool) {}

// mockReporter 实现 core.Reporter 用于测试。
type mockReporter struct {
	cap      core.Capability
	emitErr  error
	emitLoc  string
}

func (m *mockReporter) Capability() core.Capability { return m.cap }
func (m *mockReporter) Emit(_ context.Context, _ core.Score) (string, error) {
	if m.emitErr != nil {
		return "", m.emitErr
	}
	return m.emitLoc, nil
}

func validCap(name string) core.Capability {
	return core.Capability{
		Name:    name,
		Version: "1.0.0",
		Protocol: core.ProtocolVersion,
		FaultClasses: []core.FaultClass{core.F1Latency},
		TargetKinds:  []core.TargetKind{core.TargetService},
	}
}

func TestHandshakeOK(t *testing.T) {
	if err := Handshake(validCap("ok")); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestHandshakeEmptyName(t *testing.T) {
	cap := validCap("")
	if err := Handshake(cap); err == nil {
		t.Error("expected error for empty name")
	}
}

func TestHandshakeEmptyProtocol(t *testing.T) {
	cap := validCap("no-proto")
	cap.Protocol = ""
	if err := Handshake(cap); err == nil {
		t.Error("expected error for empty protocol")
	}
}

func TestHandshakeProtocolMismatch(t *testing.T) {
	cap := validCap("bad-proto")
	cap.Protocol = "v0"
	if err := Handshake(cap); err == nil {
		t.Error("expected error for protocol mismatch")
	}
}

func TestHandshakeInjectorOK(t *testing.T) {
	if err := HandshakeInjector(validCap("inj")); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestHandshakeInjectorEmptyFaultClasses(t *testing.T) {
	cap := validCap("no-faults")
	cap.FaultClasses = nil
	if err := HandshakeInjector(cap); err == nil {
		t.Error("expected error for empty faultClasses")
	}
}
