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

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// minimalWASM 是一个最小 WASM 模块，导出 add(i32, i32) → i32。
var minimalWASM = []byte{
	0x00, 0x61, 0x73, 0x6d, // magic
	0x01, 0x00, 0x00, 0x00, // version
	// Type section
	0x01, 0x07, 0x01, 0x60, 0x02, 0x7f, 0x7f, 0x01, 0x7f,
	// Function section
	0x03, 0x02, 0x01, 0x00,
	// Export section
	0x07, 0x07, 0x01, 0x03, 0x61, 0x64, 0x64, 0x00, 0x00,
	// Code section
	0x0a, 0x09, 0x01, 0x07, 0x00, 0x20, 0x00, 0x20, 0x01, 0x6a, 0x0b,
}

func TestNewWazeroRuntime_NilRuntime(t *testing.T) {
	orig := newRuntimeFn
	defer func() { newRuntimeFn = orig }()
	newRuntimeFn = func(_ context.Context) wazero.Runtime { return nil }
	_, err := NewWazeroRuntime(context.Background())
	if err == nil {
		t.Fatal("expected error for nil runtime")
	}
}

func TestWazeroRuntime_InstantiateModuleError(t *testing.T) {
	ctx := context.Background()
	orig := instantiateModuleFn
	defer func() { instantiateModuleFn = orig }()
	instantiateModuleFn = func(_ context.Context, _ wazero.Runtime, _ wazero.CompiledModule) (api.Module, error) {
		return nil, errors.New("instantiate failed")
	}
	rt, _ := NewWazeroRuntime(ctx)
	defer rt.Close(ctx)
	_, err := rt.Instantiate(ctx, minimalWASM)
	if err == nil {
		t.Fatal("expected instantiate error")
	}
}

func TestWazeroRuntime_InstantiateNilModule(t *testing.T) {
	ctx := context.Background()
	orig := instantiateModuleFn
	defer func() { instantiateModuleFn = orig }()
	instantiateModuleFn = func(_ context.Context, _ wazero.Runtime, _ wazero.CompiledModule) (api.Module, error) {
		return nil, nil
	}
	rt, _ := NewWazeroRuntime(ctx)
	defer rt.Close(ctx)
	_, err := rt.Instantiate(ctx, minimalWASM)
	if err == nil {
		t.Fatal("expected nil module error")
	}
}

func TestWazeroRuntime_InstantiateAndInvoke(t *testing.T) {
	ctx := context.Background()
	rt, err := NewWazeroRuntime(ctx)
	if err != nil {
		t.Fatalf("NewWazeroRuntime: %v", err)
	}
	defer rt.Close(ctx)

	inst, err := rt.Instantiate(ctx, minimalWASM)
	if err != nil {
		t.Fatalf("Instantiate: %v", err)
	}
	defer inst.Close(ctx)

	result, err := inst.Invoke(ctx, "add", 3, 7)
	if err != nil {
		t.Fatalf("Invoke: %v", err)
	}
	if len(result) != 1 || result[0] != 10 {
		t.Errorf("add(3,7) = %v, want [10]", result)
	}
}

func TestWazeroRuntime_InstantiateInvalidWASM(t *testing.T) {
	ctx := context.Background()
	rt, _ := NewWazeroRuntime(ctx)
	defer rt.Close(ctx)

	_, err := rt.Instantiate(ctx, []byte{0x00, 0x00})
	if err == nil {
		t.Fatal("expected error for invalid WASM")
	}
}

func TestWazeroInstance_InvokeNotFound(t *testing.T) {
	ctx := context.Background()
	rt, _ := NewWazeroRuntime(ctx)
	defer rt.Close(ctx)

	inst, err := rt.Instantiate(ctx, minimalWASM)
	if err != nil {
		t.Fatalf("Instantiate: %v", err)
	}
	defer inst.Close(ctx)

	_, err = inst.Invoke(ctx, "nonexistent", 1)
	if err == nil {
		t.Fatal("expected error for nonexistent function")
	}
}

func TestWazeroRuntime_InstantiateCompileError(t *testing.T) {
	ctx := context.Background()
	rt, _ := NewWazeroRuntime(ctx)
	defer rt.Close(ctx)

	// Valid magic but invalid module body triggers compile error
	_, err := rt.Instantiate(ctx, []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00, 0xFF})
	if err == nil {
		t.Fatal("expected compile error")
	}
}

func TestWazeroInstance_InvokeCallError(t *testing.T) {
	ctx := context.Background()
	rt, _ := NewWazeroRuntime(ctx)
	defer rt.Close(ctx)

	inst, _ := rt.Instantiate(ctx, minimalWASM)
	defer inst.Close(ctx)

	// Call with wrong number of args triggers call error
	_, err := inst.Invoke(ctx, "add", 1)
	if err == nil {
		t.Fatal("expected call error for wrong arg count")
	}
}

func TestWazeroInstance_Close(t *testing.T) {
	ctx := context.Background()
	rt, _ := NewWazeroRuntime(ctx)
	defer rt.Close(ctx)

	inst, _ := rt.Instantiate(ctx, minimalWASM)
	if err := inst.Close(ctx); err != nil {
		t.Fatalf("Close: %v", err)
	}
}
