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

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

var newRuntimeFn = func(ctx context.Context) wazero.Runtime {
	return wazero.NewRuntime(ctx)
}

// WazeroRuntime 是基于 wazero 的 WASM 运行时实现。
// WazeroRuntime is the wazero-based WASM runtime.
type WazeroRuntime struct {
	runtime wazero.Runtime
}

// NewWazeroRuntime 构造 wazero 运行时。
func NewWazeroRuntime(ctx context.Context) (*WazeroRuntime, error) {
	rt := newRuntimeFn(ctx)
	if rt == nil {
		return nil, fmt.Errorf("wazero: failed to create runtime")
	}
	return &WazeroRuntime{runtime: rt}, nil
}

var instantiateModuleFn = func(ctx context.Context, rt wazero.Runtime, compiled wazero.CompiledModule) (api.Module, error) {
	return rt.InstantiateModule(ctx, compiled, wazero.NewModuleConfig())
}

// Instantiate 编译并实例化 WASM 模块。
func (w *WazeroRuntime) Instantiate(ctx context.Context, wasm []byte) (WasmInstance, error) {
	compiled, err := w.runtime.CompileModule(ctx, wasm)
	if err != nil {
		return nil, fmt.Errorf("wazero: compile: %w", err)
	}
	mod, err := instantiateModuleFn(ctx, w.runtime, compiled)
	if err != nil {
		return nil, fmt.Errorf("wazero: instantiate: %w", err)
	}
	if mod == nil {
		return nil, fmt.Errorf("wazero: instantiate returned nil module")
	}
	return &wazeroInstance{mod: mod}, nil
}

// Close 关闭运行时，释放所有资源。
func (w *WazeroRuntime) Close(ctx context.Context) error {
	return w.runtime.Close(ctx)
}

// wazeroInstance 是 wazero 模块实例。
type wazeroInstance struct {
	mod api.Module
}

// Invoke 调用导出函数。
func (wi *wazeroInstance) Invoke(ctx context.Context, funcName string, args ...uint64) ([]uint64, error) {
	fn := wi.mod.ExportedFunction(funcName)
	if fn == nil {
		return nil, fmt.Errorf("wazero: function %q not exported", funcName)
	}
	results, err := fn.Call(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("wazero: call %q: %w", funcName, err)
	}
	return results, nil
}

// Close 关闭模块实例。
func (wi *wazeroInstance) Close(ctx context.Context) error {
	return wi.mod.Close(ctx)
}
