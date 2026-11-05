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
	"sync"
	"sync/atomic"
)

// WasmRuntime 是 WASM 运行时引擎接口，可由 wazero 等实现。
// WasmRuntime is the WASM runtime engine interface.
type WasmRuntime interface {
	// Instantiate 编译并实例化 WASM 模块，返回实例句柄。
	Instantiate(ctx context.Context, wasm []byte) (WasmInstance, error)
}

// WasmInstance 是一个已实例化的 WASM 模块。
type WasmInstance interface {
	// Invoke 调用导出函数，返回结果。
	Invoke(ctx context.Context, funcName string, args ...uint64) ([]uint64, error)
	// Close 释放实例资源。
	Close(ctx context.Context) error
}

// SandboxConfig 是沙箱资源限制配置。
type SandboxConfig struct {
	MaxInstructions uint64
	MaxMemoryBytes   uint64
	TimeoutMs        uint64
}

// DefaultSandboxConfig 返回默认沙箱配置。
func DefaultSandboxConfig() SandboxConfig {
	return SandboxConfig{
		MaxInstructions: 1_000_000,
		MaxMemoryBytes: 32 * 1024 * 1024,
		TimeoutMs:      5000,
	}
}

// Sandbox 是 WASM 插件隔离执行沙箱（CES-7.1.4）。
// Sandbox is the WASM plugin isolation sandbox.
type Sandbox struct {
	rt      WasmRuntime
	config  SandboxConfig
	mu      sync.Mutex
	instr   atomic.Uint64
	active  map[string]WasmInstance
}

// NewSandbox 构造沙箱。
func NewSandbox(rt WasmRuntime, config SandboxConfig) (*Sandbox, error) {
	if rt == nil {
		return nil, fmt.Errorf("sandbox: runtime must not be nil")
	}
	return &Sandbox{
		rt:     rt,
		config: config,
		active: make(map[string]WasmInstance),
	}, nil
}

// Instantiate 实例化 WASM 模块并记录活跃实例。
func (s *Sandbox) Instantiate(ctx context.Context, name string, wasm []byte) error {
	inst, err := s.rt.Instantiate(ctx, wasm)
	if err != nil {
		return fmt.Errorf("sandbox: instantiate %s: %w", name, err)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.active[name] = inst
	return nil
}

// Invoke 调用插件导出函数，检查指令计数限制。
func (s *Sandbox) Invoke(ctx context.Context, name, funcName string, args ...uint64) ([]uint64, error) {
	s.mu.Lock()
	inst, ok := s.active[name]
	s.mu.Unlock()
	if !ok {
		return nil, fmt.Errorf("sandbox: plugin %q not instantiated", name)
	}

	result, err := inst.Invoke(ctx, funcName, args...)
	if err != nil {
		return nil, fmt.Errorf("sandbox: invoke %s.%s: %w", name, funcName, err)
	}

	after := s.instr.Load()
	if s.config.MaxInstructions > 0 && after > s.config.MaxInstructions {
	s.mu.Lock()
		delete(s.active, name)
		s.mu.Unlock()
		_ = inst.Close(ctx)
		return nil, fmt.Errorf("sandbox: instruction limit %d exceeded (CES-7.1.4)", s.config.MaxInstructions)
	}
	return result, nil
}

// Close 关闭指定插件实例。
func (s *Sandbox) Close(ctx context.Context, name string) error {
	s.mu.Lock()
	inst, ok := s.active[name]
	if ok {
		delete(s.active, name)
	}
	s.mu.Unlock()
	if !ok {
		return nil
	}
	return inst.Close(ctx)
}

// InstructionCount 返回累计指令计数。
func (s *Sandbox) InstructionCount() uint64 {
	return s.instr.Load()
}

// AddInstructions 增加指令计数（供运行时回调）。
func (s *Sandbox) AddInstructions(n uint64) {
	s.instr.Add(n)
}

// IsActive 报告插件是否已实例化。
func (s *Sandbox) IsActive(name string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.active[name]
	return ok
}
