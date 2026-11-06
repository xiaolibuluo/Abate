// Abate —— 止息混沌 · 测试插件注册表
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"fmt"
	"sort"
	"sync"
)

// Registry 是线程安全的测试插件注册表。
type Registry struct {
	mu      sync.RWMutex
	entries map[string]TestPlugin
}

// NewRegistry 构造空注册表。
func NewRegistry() *Registry {
	return &Registry{entries: make(map[string]TestPlugin)}
}

// Register 注册测试插件。
func (r *Registry) Register(p TestPlugin) error {
	cap := p.Capability()
	if cap.Name == "" {
		return fmt.Errorf("test: plugin name is empty")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.entries[cap.Name]; exists {
		return fmt.Errorf("test plugin %q: already registered", cap.Name)
	}
	r.entries[cap.Name] = p
	return nil
}

// Unregister 移除已注册插件。
func (r *Registry) Unregister(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.entries[name]; !exists {
		return fmt.Errorf("test plugin %q: not registered", name)
	}
	delete(r.entries, name)
	return nil
}

// Lookup 按名称查找插件。
func (r *Registry) Lookup(name string) (TestPlugin, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.entries[name]
	return p, ok
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

// DefaultRegistry 是包级默认注册表，供插件 init() 自注册。
var DefaultRegistry = NewRegistry()
