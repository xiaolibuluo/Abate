/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

import (
	"context"
	"sync/atomic"
	"time"

	"abate/plugin/host"
)

// Swap is one hot-swap transaction. It retains both instances so a caller can
// audit which plugin was drained and which plugin now serves the registry. The
// phase is stored atomically because State may be read by a monitoring
// goroutine while Begin is still advancing the machine.
type Swap struct {
	state atomic.Int32
	old   *host.Instance
	next  *host.Instance
}

// Begin runs the complete four-phase swap and returns the transaction together
// with the first error encountered. On error the returned Swap is non-nil and
// its State() names the phase in which the failure happened; this lets the
// control plane report a precise, deterministic failure phase rather than a
// bare error. A failed swap may have already switched the registry (the new
// instance keeps serving), which is the intended "never interrupt the
// experiment" behaviour.
//
// Phases:
//  1. PREPARE: validate instances, registry, identity and Allowed version step.
//  2. SWITCH:  old.MarkSuperseded() then Registry.Replace(id, next), a single
//     atomic pointer assignment; new calls are rejected on old and served by next.
//  3. DRAIN:   wait for old.WaitCalls(), bounded by DrainTimeout and ctx.
//  4. UNLOAD:  old.Close() releases plugin resources without touching the
//     registry (Host.Close would remove the newly registered instance).
func Begin(ctx context.Context, old, next *host.Instance) (*Swap, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	s := &Swap{old: old, next: next}
	s.setState(PREPARE)
	if err := s.prepare(); err != nil {
		return s, err
	}
	s.setState(SWITCH)
	if err := s.switchOver(); err != nil {
		return s, err
	}
	s.setState(DRAIN)
	if err := s.drain(ctx); err != nil {
		return s, err
	}
	s.setState(UNLOAD)
	if err := old.Close(); err != nil {
		return s, err
	}
	return s, nil
}

// State returns the current phase of the transaction.
func (s *Swap) State() SwapState {
	return SwapState(s.state.Load())
}

// setState advances the phase atomically.
func (s *Swap) setState(v SwapState) {
	s.state.Store(int32(v)) //nolint:gosec // SwapState is a small bounded four-value enum.
}

// prepare validates every precondition before any mutation of the registry.
func (s *Swap) prepare() error {
	if s.old == nil || s.next == nil || s.old == s.next {
		return ErrSwapInvalidInstance
	}
	if s.old.Registry == nil && s.next.Registry == nil {
		return ErrSwapNoRegistry
	}
	if s.old.Manifest.ID != s.next.Manifest.ID ||
		s.old.Manifest.Type != s.next.Manifest.Type {
		return ErrSwapIdentityMismatch
	}
	if s.next.IsSuperseded() {
		return ErrSwapNotAllowed
	}
	if !Allowed(s.old.Manifest.Type, s.old.Manifest.Version, s.next.Manifest.Version) {
		return ErrSwapNotAllowed
	}
	return nil
}

// switchOver performs the atomic registry replacement. The superseded flag is
// set first so no new call can start on old from the moment the registry is
// about to change; calls already counted by BeginCall finish under DRAIN.
func (s *Swap) switchOver() error {
	reg := s.old.Registry
	if reg == nil {
		reg = s.next.Registry
	}
	if reg == nil {
		return ErrSwapNoRegistry
	}
	s.old.MarkSuperseded()
	reg.Replace(s.old.Manifest.ID, s.next)
	return nil
}

// drain waits for in-flight calls on the old instance. The wait runs in a
// dedicated goroutine so the wall-clock DrainTimeout and ctx cancellation can
// bound it. On timeout the old instance is left unclosed: closing a plugin
// while a call is still executing would corrupt the experiment.
func (s *Swap) drain(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		s.old.WaitCalls()
		close(done)
	}()
	timer := time.NewTimer(DrainTimeout)
	defer timer.Stop()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return ErrDrainTimeout
	}
}
