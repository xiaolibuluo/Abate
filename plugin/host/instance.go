/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"sync"
	"sync/atomic"

	"abate/plugin/capability"
	"abate/plugin/signature"
	"abate/sim/clock"
	"abate/sim/rand"
)

// Instance is a single loaded plugin together with its granted capabilities, the
// injected deterministic backends, and the entry whitelist. It is the unit of
// registration, capability gating, and hot-swap atomicity.
type Instance struct {
	// Plugin is the underlying plugin implementation (operator/reporter/adapter).
	Plugin

	// Manifest is the signed metadata, cached for O(1) access.
	Manifest signature.Manifest

	// Caps is the granted capability mask (default-deny; zero denies everything).
	Caps capability.Capabilities

	// Clock is the per-instance deterministic virtual clock.
	Clock *clock.VirtualClock

	// RNG is the per-instance deterministic random stream.
	RNG *rand.Stream

	// Entries is the entry whitelist this instance may be called through.
	Entries []string

	// Registry is the host registry that owns this instance. It is a back-reference
	// set by Host.Load so plugin/hotswap can atomically replace a superseded
	// instance without the caller threading the registry through every swap.
	Registry *Registry

	inflight   sync.WaitGroup
	superseded atomic.Bool
}

// MarkSuperseded marks the instance as replaced by a hot-swap. After this call,
// Host.Call rejects new invocations with ErrInstanceStale while already in-flight
// calls are allowed to finish. It is invoked by plugin/hotswap during SWITCH.
func (inst *Instance) MarkSuperseded() {
	inst.superseded.Store(true)
}

// IsSuperseded reports whether the instance has been replaced by a hot-swap.
func (inst *Instance) IsSuperseded() bool {
	return inst.superseded.Load()
}

// BeginCall atomically checks the superseded flag and, when the instance is still
// live, increments the in-flight call counter, returning a release function to be
// deferred by the caller. The ok result is false when the instance is stale, in which case
// no counter is taken and release is nil. The atomic check-then-add closes the
// race between a reader obtaining the old instance and a concurrent swap.
func (inst *Instance) BeginCall() (release func(), ok bool) {
	if inst.superseded.Load() {
		return nil, false
	}
	inst.inflight.Add(1)
	return inst.inflight.Done, true
}

// WaitCalls blocks until every in-flight call has released. It is used by
// plugin/hotswap during the DRAIN phase and must be bounded by the real-time
// DrainTimeout by the caller.
func (inst *Instance) WaitCalls() {
	inst.inflight.Wait()
}
