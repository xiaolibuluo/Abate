/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"fmt"
	"sort"
	"sync"
)

// PluginManager is the plugin port exposed by the control plane. The production
// host adapter and the static fallback both satisfy it, so the real plugin
// store can replace the fallback without altering the API surface.
type PluginManager interface {
	List() []PluginInfo
	Install(upload PluginUpload) (PluginInfo, error)
	Remove(id string) error
	Upgrade(upload PluginUpload) (PluginInfo, error)
}

// staticPluginManager is a thread-safe, in-memory plugin registry used when no
// host adapter is configured. It is deterministic: listings are sorted by id.
type staticPluginManager struct {
	mu      sync.Mutex
	plugins map[string]PluginInfo
}

// newStaticPluginManager returns a manager preloaded with a reference plugin so
// the listing endpoint is never empty in a default deployment.
func newStaticPluginManager() *staticPluginManager {
	return &staticPluginManager{
		plugins: map[string]PluginInfo{
			"reporter-core": {
				ID: "reporter-core", Type: "reporter", Version: "1.0.0",
				Capabilities: []string{"sarif", "json", "junit", "html"},
				Signature:    "verified",
			},
		},
	}
}

// List returns the installed plugins ordered by identifier.
func (m *staticPluginManager) List() []PluginInfo {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]PluginInfo, 0, len(m.plugins))
	for _, info := range m.plugins {
		out = append(out, info)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Install registers a new plugin, rejecting an empty identifier.
func (m *staticPluginManager) Install(upload PluginUpload) (PluginInfo, error) {
	if upload.ID == "" {
		return PluginInfo{}, fmt.Errorf("%w: empty plugin id", errPluginInvalid)
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.plugins[upload.ID]; exists {
		return PluginInfo{}, fmt.Errorf("%w: %s", errPluginExists, upload.ID)
	}
	info := PluginInfo(upload)
	m.plugins[upload.ID] = info
	return info, nil
}

// Remove deletes a plugin by identifier.
func (m *staticPluginManager) Remove(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.plugins[id]; !exists {
		return fmt.Errorf("%w: %s", errPluginMissing, id)
	}
	delete(m.plugins, id)
	return nil
}

// Upgrade replaces a plugin's version and metadata, requiring prior presence.
func (m *staticPluginManager) Upgrade(upload PluginUpload) (PluginInfo, error) {
	if upload.ID == "" {
		return PluginInfo{}, fmt.Errorf("%w: empty plugin id", errPluginInvalid)
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.plugins[upload.ID]; !exists {
		return PluginInfo{}, fmt.Errorf("%w: %s", errPluginMissing, upload.ID)
	}
	info := PluginInfo(upload)
	m.plugins[upload.ID] = info
	return info, nil
}
