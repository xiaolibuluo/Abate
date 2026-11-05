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
	"errors"
	"testing"
)

type mockRuntime struct {
	inst    WasmInstance
	instErr error
}

func (m *mockRuntime) Instantiate(_ context.Context, _ []byte) (WasmInstance, error) {
	if m.instErr != nil {
		return nil, m.instErr
	}
	if m.inst == nil {
		return &mockInstance{}, nil
	}
	return m.inst, nil
}

type mockInstance struct {
	invokeErr error
	closeErr  error
	invoked   bool
	closed    bool
}

func (m *mockInstance) Invoke(_ context.Context, _ string, _ ...uint64) ([]uint64, error) {
	m.invoked = true
	if m.invokeErr != nil {
		return nil, m.invokeErr
	}
	return []uint64{42}, nil
}

func (m *mockInstance) Close(_ context.Context) error {
	m.closed = true
	return m.closeErr
}

func TestNewSandbox_Success(t *testing.T) {
	rt := &mockRuntime{}
	s, err := NewSandbox(rt, DefaultSandboxConfig())
	if err != nil {
		t.Fatalf("NewSandbox: %v", err)
	}
	if s.InstructionCount() != 0 {
		t.Error("InstructionCount should be 0")
	}
	if s.IsActive("x") {
		t.Error("IsActive should be false")
	}
}

func TestNewSandbox_NilRuntime(t *testing.T) {
	_, err := NewSandbox(nil, DefaultSandboxConfig())
	if err == nil {
		t.Fatal("expected error for nil runtime")
	}
}

func TestSandbox_Instantiate_Success(t *testing.T) {
	rt := &mockRuntime{}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	if err := s.Instantiate(context.Background(), "p1", []byte{0x00, 0x61, 0x73, 0x6d}); err != nil {
		t.Fatalf("Instantiate: %v", err)
	}
	if !s.IsActive("p1") {
		t.Error("IsActive should be true")
	}
}

func TestSandbox_Instantiate_Error(t *testing.T) {
	rt := &mockRuntime{instErr: errors.New("compile failed")}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	if err := s.Instantiate(context.Background(), "p1", []byte{}); err == nil {
		t.Fatal("expected error for instantiate failure")
	}
}

func TestSandbox_Invoke_Success(t *testing.T) {
	rt := &mockRuntime{}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	_ = s.Instantiate(context.Background(), "p1", []byte{0x00})

	result, err := s.Invoke(context.Background(), "p1", "run", 1, 2)
	if err != nil {
		t.Fatalf("Invoke: %v", err)
	}
	if len(result) != 1 || result[0] != 42 {
		t.Errorf("result = %v", result)
	}
}

func TestSandbox_Invoke_NotInstantiated(t *testing.T) {
	rt := &mockRuntime{}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	_, err := s.Invoke(context.Background(), "unknown", "run")
	if err == nil {
		t.Fatal("expected error for not instantiated")
	}
}

func TestSandbox_Invoke_InvokeError(t *testing.T) {
	mi := &mockInstance{invokeErr: errors.New("invoke failed")}
	rt := &mockRuntime{inst: mi}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	_ = s.Instantiate(context.Background(), "p1", []byte{0x00})

	_, err := s.Invoke(context.Background(), "p1", "run")
	if err == nil {
		t.Fatal("expected error for invoke failure")
	}
}

func TestSandbox_Invoke_InstructionLimit(t *testing.T) {
	rt := &mockRuntime{}
	s, _ := NewSandbox(rt, SandboxConfig{MaxInstructions: 100, MaxMemoryBytes: 1024, TimeoutMs: 1000})
	_ = s.Instantiate(context.Background(), "p1", []byte{0x00})

	s.AddInstructions(200)
	_, err := s.Invoke(context.Background(), "p1", "run")
	if err == nil {
		t.Fatal("expected error for instruction limit exceeded")
	}
	if s.IsActive("p1") {
		t.Error("IsActive should be false after limit exceeded")
	}
}

func TestSandbox_Invoke_NoLimit(t *testing.T) {
	rt := &mockRuntime{}
	s, _ := NewSandbox(rt, SandboxConfig{MaxInstructions: 0, MaxMemoryBytes: 0, TimeoutMs: 0})
	_ = s.Instantiate(context.Background(), "p1", []byte{0x00})

	s.AddInstructions(999999)
	_, err := s.Invoke(context.Background(), "p1", "run")
	if err != nil {
		t.Fatalf("Invoke with no limit: %v", err)
	}
}

func TestSandbox_Close_Success(t *testing.T) {
	mi := &mockInstance{}
	rt := &mockRuntime{inst: mi}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	_ = s.Instantiate(context.Background(), "p1", []byte{0x00})

	if err := s.Close(context.Background(), "p1"); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if s.IsActive("p1") {
		t.Error("IsActive should be false after close")
	}
	if !mi.closed {
		t.Error("instance.Close was not called")
	}
}

func TestSandbox_Close_NotActive(t *testing.T) {
	rt := &mockRuntime{}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	if err := s.Close(context.Background(), "unknown"); err != nil {
		t.Errorf("Close with unknown: %v", err)
	}
}

func TestSandbox_Close_CloseError(t *testing.T) {
	mi := &mockInstance{closeErr: errors.New("close failed")}
	rt := &mockRuntime{inst: mi}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	_ = s.Instantiate(context.Background(), "p1", []byte{0x00})

	if err := s.Close(context.Background(), "p1"); err == nil {
		t.Fatal("expected error for close failure")
	}
}

func TestSandbox_AddInstructions(t *testing.T) {
	rt := &mockRuntime{}
	s, _ := NewSandbox(rt, DefaultSandboxConfig())
	s.AddInstructions(100)
	s.AddInstructions(50)
	if s.InstructionCount() != 150 {
		t.Errorf("InstructionCount = %d, want 150", s.InstructionCount())
	}
}

func TestDefaultSandboxConfig(t *testing.T) {
	cfg := DefaultSandboxConfig()
	if cfg.MaxInstructions != 1_000_000 {
		t.Errorf("MaxInstructions = %d", cfg.MaxInstructions)
	}
	if cfg.MaxMemoryBytes != 32*1024*1024 {
		t.Errorf("MaxMemoryBytes = %d", cfg.MaxMemoryBytes)
	}
	if cfg.TimeoutMs != 5000 {
		t.Errorf("TimeoutMs = %d", cfg.TimeoutMs)
	}
}
