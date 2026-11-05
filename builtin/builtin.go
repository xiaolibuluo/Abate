// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package builtin 将全部内置角色一键注册到插件注册表。
//
// Package builtin registers all built-in roles into the plugin registry.
package builtin

import (
	"fmt"

	"github.com/qingdujian/abate/adapter"
	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/injector"
	"github.com/qingdujian/abate/plugin"
	"github.com/qingdujian/abate/probe"
	"github.com/qingdujian/abate/reporter"
	"github.com/qingdujian/abate/validator"
)

// reporterNew 可在测试中替换以触发报告器创建失败。
var reporterNew = reporter.New

// Register 将全部内置角色注册到注册表。dir 为文件系统注入器和报告器的工作目录。
// Register registers all built-in roles into the registry.
func Register(r *plugin.Registry, dir string) error {
	inj, err := injector.New(dir)
	if err != nil {
		return fmt.Errorf("builtin: file injector: %w", err)
	}
	if err := r.Register(plugin.Plugin{Injector: inj}); err != nil {
		return fmt.Errorf("builtin: register injector: %w", err)
	}

	p := probe.New(42)
	if err := r.Register(plugin.Plugin{Probe: p}); err != nil {
		return fmt.Errorf("builtin: register probe: %w", err)
	}

	v := validator.New()
	if err := r.Register(plugin.Plugin{Validator: v}); err != nil {
		return fmt.Errorf("builtin: register validator: %w", err)
	}

	rep, err := reporterNew(dir)
	if err != nil {
		return fmt.Errorf("builtin: reporter: %w", err)
	}
	if err := r.Register(plugin.Plugin{Reporter: rep}); err != nil {
		return fmt.Errorf("builtin: register reporter: %w", err)
	}

	return nil
}

// RegisterAdapter 注册一个第三方适配器到注册表。
// RegisterAdapter registers a third-party adapter into the registry.
func RegisterAdapter(r *plugin.Registry, name, version string, backend adapter.Backend, table *adapter.MappingTable) error {
	a, err := adapter.New(name, version, backend, table)
	if err != nil {
		return fmt.Errorf("builtin: adapter %s: %w", name, err)
	}
	return r.Register(plugin.Plugin{Injector: a})
}

// EnsureCapability 确保注册表中存在指定角色的可用实现。
func EnsureCapability(r *plugin.Registry, role plugin.Role) error {
	switch role {
	case plugin.RoleInjector:
		_, err := r.FindInjector(core.F1Latency, core.TargetService)
		return err
	case plugin.RoleProbe:
		_, err := r.FindProbe(core.TargetService)
		return err
	case plugin.RoleValidator:
		_, err := r.FindValidator()
		return err
	case plugin.RoleReporter:
		_, err := r.FindReporter()
		return err
	}
	return fmt.Errorf("builtin: unknown role %q", role)
}