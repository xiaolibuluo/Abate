// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package plugin 实现插件注册、握手验证与路由（CES 第 7 章）。
//
// Package plugin implements plugin registration, handshake validation,
// and routing per CES Chapter 7.
package plugin

import (
	"fmt"
	"sort"
	"sync"

	"github.com/qingdujian/abate/core"
)

// Role 是插件角色标识（CES-7.1.3）。
type Role string

const (
	RoleInjector  Role = "injector"
	RoleProbe     Role = "probe"
	RoleValidator Role = "validator"
	RoleReporter  Role = "reporter"
)

// entry 是注册表中一个插件条目。
type entry struct {
	cap   core.Capability
	inj   core.Injector
	probe core.Probe
	val   core.Validator
	rep   core.Reporter
}

// roles 返回该条目声明的角色集合。
func (e *entry) roles() []Role {
	var rs []Role
	if e.inj != nil {
		rs = append(rs, RoleInjector)
	}
	if e.probe != nil {
		rs = append(rs, RoleProbe)
	}
	if e.val != nil {
		rs = append(rs, RoleValidator)
	}
	if e.rep != nil {
		rs = append(rs, RoleReporter)
	}
	return rs
}

// Registry 是线程安全的插件注册表（CES-7.1）。
type Registry struct {
	mu      sync.RWMutex
	entries map[string]*entry
}

// NewRegistry 构造空注册表。
func NewRegistry() *Registry {
	return &Registry{entries: make(map[string]*entry)}
}

// Plugin 声明一个待注册插件，可携带多个角色实现。
// 至少一个角色字段非 nil 才可注册。
type Plugin struct {
	Injector  core.Injector
	Probe     core.Probe
	Validator core.Validator
	Reporter  core.Reporter
}

// capability 从第一个非 nil 角色提取能力声明。
func (p Plugin) capability() (core.Capability, error) {
	if p.Injector != nil {
		return p.Injector.Capability(), nil
	}
	if p.Probe != nil {
		return p.Probe.Capability(), nil
	}
	if p.Validator != nil {
		return p.Validator.Capability(), nil
	}
	if p.Reporter != nil {
		return p.Reporter.Capability(), nil
	}
	return core.Capability{}, fmt.Errorf("plugin: no role implemented (CES-7.1.3)")
}

// Register 注册插件，执行握手验证后存入注册表。
// 插件名由 Capability.Name 决定，重复注册返回错误。
func (r *Registry) Register(p Plugin) error {
	cap, err := p.capability()
	if err != nil {
		return err
	}
	if err := Handshake(cap); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.entries[cap.Name]; exists {
		return fmt.Errorf("plugin %q: already registered", cap.Name)
	}
	r.entries[cap.Name] = &entry{
		cap: cap, inj: p.Injector, probe: p.Probe, val: p.Validator, rep: p.Reporter,
	}
	return nil
}

// Unregister 移除已注册插件。
func (r *Registry) Unregister(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.entries[name]; !exists {
		return fmt.Errorf("plugin %q: not registered", name)
	}
	delete(r.entries, name)
	return nil
}

// Lookup 按名称查找插件条目。
func (r *Registry) Lookup(name string) (*entry, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.entries[name]
	return e, ok
}

// List 返回全部已注册插件名，按字典序排列。
func (r *Registry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.entries))
	for n := range r.entries {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// FindInjector 按故障编码与目标类型查找 Injector（CES-7.1.2）。
func (r *Registry) FindInjector(fc core.FaultClass, tk core.TargetKind) (core.Injector, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, e := range r.entries {
		if e.inj == nil {
			continue
		}
		if !containsFault(e.cap.FaultClasses, fc) {
			continue
		}
		if !containsTarget(e.cap.TargetKinds, tk) {
			continue
		}
		return e.inj, nil
	}
	return nil, fmt.Errorf("no injector for fault %q target %q (CES-7.1.2)", fc, tk)
}

// FindProbe 按目标类型查找 Probe。
func (r *Registry) FindProbe(tk core.TargetKind) (core.Probe, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, e := range r.entries {
		if e.probe == nil {
			continue
		}
		if !containsTarget(e.cap.TargetKinds, tk) {
			continue
		}
		return e.probe, nil
	}
	return nil, fmt.Errorf("no probe for target %q (CES-7.1.2)", tk)
}

// FindValidator 查找已注册 Validator；未注册时返回错误（核心兜底，CES-7.2.3）。
func (r *Registry) FindValidator() (core.Validator, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, e := range r.entries {
		if e.val != nil {
			return e.val, nil
		}
	}
	return nil, fmt.Errorf("no validator registered (CES-7.2.3)")
}

// FindReporter 查找已注册 Reporter；未注册时返回错误（核心兜底，CES-7.2.3）。
func (r *Registry) FindReporter() (core.Reporter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, e := range r.entries {
		if e.rep != nil {
			return e.rep, nil
		}
	}
	return nil, fmt.Errorf("no reporter registered (CES-7.2.3)")
}

func containsFault(haystack []core.FaultClass, needle core.FaultClass) bool {
	for _, f := range haystack {
		if f == needle {
			return true
		}
	}
	return false
}

func containsTarget(haystack []core.TargetKind, needle core.TargetKind) bool {
	for _, t := range haystack {
		if t == needle {
			return true
		}
	}
	return false
}
