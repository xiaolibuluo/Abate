/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package hotswap implements the deterministic four-phase plugin hot-swap state
// machine: PREPARE, SWITCH, DRAIN and UNLOAD.
//
// A swap keeps an experiment running: the new instance is registered atomically
// before the old one is drained, so subsequent calls are served by the new
// instance while calls already in flight on the old instance are allowed to
// finish. New calls on the superseded instance are rejected with
// host.ErrInstanceStale. No plugin state is rolled back and no experiment is
// interrupted.
//
// Determinism note: all state transitions, capability gates and version checks
// are pure functions of the two instances. The single wall-clock boundary is
// DrainTimeout (see timeout.go), which bounds how long DRAIN may wait for real
// goroutines; the runtime itself never reads time.Now, never uses math/rand and
// never leaks map iteration order.
package hotswap
