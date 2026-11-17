/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"context"

	"abate/sim/clock"
	"abate/sim/rand"
)

// HostBackend carries the deterministic clock and random stream made available to
// a plugin during a Call. Host.Call injects it into the call context so operator
// plugins read virtual time and draw deterministic randomness without touching
// global state, time.Now, or math/rand. This is the WAS-HOST host-function
// backend, realized with SIM-CLK/SIM-RNG instead of a WASM runtime.
//
//nolint:revive // HostBackend is a frozen M4 host-function contract name.
type HostBackend struct {
	// Clock is the per-instance deterministic virtual clock.
	Clock *clock.VirtualClock

	// RNG is the per-instance deterministic random stream.
	RNG *rand.Stream
}

type backendCtxKey struct{}

// WithBackend returns a context carrying the host backend.
func WithBackend(ctx context.Context, b *HostBackend) context.Context {
	return context.WithValue(ctx, backendCtxKey{}, b)
}

// Backend extracts the host backend from ctx, or nil if absent.
func Backend(ctx context.Context) *HostBackend {
	if b, ok := ctx.Value(backendCtxKey{}).(*HostBackend); ok {
		return b
	}
	return nil
}
