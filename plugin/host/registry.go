/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"sync"

	"abate/plugin/signature"
)

// Registry maps plugin types to factories and plugin ids to live instances. All
// access is guarded by a single RWMutex so concurrent Load/Lookup/Replace are
// safe and Replace is an atomic single-pointer assignment (zero intermediate
// state observed by readers).
type Registry struct {
	mu        sync.RWMutex
	factories map[signature.PluginType]Factory
	instances map[signature.PluginID]*Instance
}

// newRegistry builds an empty registry.
func newRegistry() *Registry {
	return &Registry{
		factories: make(map[signature.PluginType]Factory),
		instances: make(map[signature.PluginID]*Instance),
	}
}

// Register binds a factory to a plugin type. A duplicate registration returns
// errFactoryExists.
func (r *Registry) Register(factory Factory, t signature.PluginType) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, dup := r.factories[t]; dup {
		return errFactoryExists
	}
	r.factories[t] = factory
	return nil
}

// Get returns the live instance registered under id.
func (r *Registry) Get(id signature.PluginID) (*Instance, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	inst, ok := r.instances[id]
	return inst, ok
}

// Replace atomically swaps the instance registered under id. It is a single
// assignment under the write lock, so no reader ever observes a torn pointer.
func (r *Registry) Replace(id signature.PluginID, inst *Instance) {
	r.mu.Lock()
	r.instances[id] = inst
	r.mu.Unlock()
}

// Snapshot returns a shallow copy of the id->instance map. The returned map is a
// snapshot; callers must not mutate the instances. Because Go map iteration order
// is non-deterministic, sort the keys before hashing or comparing snapshots.
func (r *Registry) Snapshot() map[signature.PluginID]*Instance {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make(map[signature.PluginID]*Instance, len(r.instances))
	for k, v := range r.instances {
		out[k] = v
	}
	return out
}

func (r *Registry) put(id signature.PluginID, inst *Instance) {
	r.mu.Lock()
	r.instances[id] = inst
	r.mu.Unlock()
}

func (r *Registry) remove(id signature.PluginID) {
	r.mu.Lock()
	delete(r.instances, id)
	r.mu.Unlock()
}
