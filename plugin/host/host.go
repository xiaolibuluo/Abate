/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"

	"abate/plugin/capability"
	"abate/plugin/signature"
	"abate/sim/clock"
	"abate/sim/rand"
)

// Host is the native plugin runtime. It verifies signatures, grants
// capabilities, instantiates plugins through a per-type Factory, registers them,
// and dispatches capability-gated calls. The host holds no global mutable state
// and never calls time.Now or math/rand; per-instance Clock/RNG are derived
// deterministically from Seed and the manifest digest (see deriveSeed).
type Host struct {
	store *signature.TrustStore
	reg   *Registry

	// Seed is the deterministic base seed for derived per-instance RNG streams.
	Seed uint64
}

// NewHost constructs a Host bound to the given trust store.
func NewHost(store *signature.TrustStore) *Host {
	return &Host{store: store, reg: newRegistry()}
}

// Load verifies, grants, instantiates, and registers a plugin.
//
// Load chain:
//  1. signature.Verify(store, m): an unsigned or tampered manifest is rejected
//     outright (ErrSignatureInvalid / ErrTrustAnchorUnknown / ErrManifestHashMismatch).
//     There is no fall-through to an "unsigned" state.
//  2. capability.ParseGrant(m.CapabilitiesDecl): an unknown capability name
//     rejects the whole load with ErrUnknownCapability (fail-closed).
//  3. The factory instantiates the Plugin from the verified manifest; it never
//     sees an unverified manifest.
//  4. The effective grant is the intersection of the declared capabilities and
//     the caller-supplied caps (default-deny: a plugin gets at most both).
//  5. A deterministic Clock and RNG are derived and injected, then the instance
//     is registered atomically by manifest id.
func (h *Host) Load(m Manifest, factory Factory, caps Capabilities) (*Instance, error) {
	if err := signature.Verify(h.store, m); err != nil {
		return nil, err
	}
	declared, err := capability.ParseGrant(m.CapabilitiesDecl)
	if err != nil {
		return nil, err
	}
	p, err := factory(m)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errNilPlugin
	}
	granted := declared & caps
	seed := deriveSeed(h.Seed, m)
	inst := &Instance{
		Plugin:   p,
		Manifest: m,
		Caps:     granted,
		Clock:    clock.NewVirtualClock(),
		RNG:      rand.New(seed),
		Entries:  entriesFor(p),
		Registry: h.reg,
	}
	h.reg.put(m.ID, inst)
	return inst, nil
}

// Lookup returns the live instance registered under id.
func (h *Host) Lookup(id PluginID) (*Instance, bool) {
	return h.reg.Get(id)
}

// Registry returns the host registry.
func (h *Host) Registry() *Registry {
	return h.reg
}

// Call dispatches a capability-gated call to an instance entry.
//
// Call chain: entry-whitelist check (ErrEntryForbidden) -> capability gate via
// capability.CheckMany (ErrCapabilityDenied on missing authority) -> dispatch.
// Reporter and adapter entries additionally require CapObserve regardless of the
// caller-supplied want list, so an empty want list can never bypass observation
// authority. Dispatch targets OperatorPlugin for EntryApply, ReporterPlugin for
// EntryBuild, and AdapterPlugin for EntryDetect and EntryEmit; every result is
// capped at MaxResultBytes (ErrResultTooLarge). Dispatch runs against the
// instance's injected Clock/RNG backends so output is deterministic. The entry-whitelist and the BeginCall superseded-guard are
// applied before any dispatch, so a stale instance is never invoked.
func (*Host) Call(inst *Instance, entry string, args []byte, want []Capability) ([]byte, error) {
	if inst == nil {
		return nil, ErrEntryForbidden
	}
	if !contains(inst.Entries, entry) {
		return nil, ErrEntryForbidden
	}
	// Reporter and adapter entries are always observation-bound: even when the
	// caller supplies an empty want list, CapObserve is enforced fail-closed.
	wants := append([]Capability(nil), want...)
	switch entry {
	case EntryBuild, EntryDetect, EntryEmit:
		if !hasCapability(wants, capability.CapObserve) {
			wants = append(wants, capability.CapObserve)
		}
	}
	if err := capability.CheckMany(inst.Caps, wants); err != nil {
		return nil, err
	}
	release, ok := inst.BeginCall()
	if !ok {
		return nil, ErrInstanceStale
	}
	defer release()

	switch entry {
	case EntryApply:
		return hostCallApply(inst, args)
	case EntryBuild:
		return hostCallBuild(inst, args)
	case EntryDetect:
		return hostCallDetect(inst, args)
	case EntryEmit:
		return hostCallEmit(inst, args)
	}
	return nil, ErrEntryForbidden
}

// hostCallApply dispatches the OperatorPlugin.Apply entry.
func hostCallApply(inst *Instance, args []byte) ([]byte, error) {
	op, ok := inst.Plugin.(OperatorPlugin)
	if !ok {
		return nil, ErrEntryForbidden
	}
	backend := &HostBackend{Clock: inst.Clock, RNG: inst.RNG}
	out, err := op.Apply(WithBackend(context.Background(), backend), args)
	if err != nil {
		return nil, err
	}
	if len(out) > MaxResultBytes {
		return nil, ErrResultTooLarge
	}
	return out, nil
}

// hostCallBuild dispatches the ReporterPlugin.Build entry.
func hostCallBuild(inst *Instance, args []byte) ([]byte, error) {
	rp, ok := inst.Plugin.(ReporterPlugin)
	if !ok {
		return nil, ErrEntryForbidden
	}
	var ba BuildArgs
	if len(args) > 0 {
		if err := DecodeBuildArgs(args, &ba); err != nil {
			return nil, err
		}
	}
	out, err := rp.Build(ba.Findings, ba.Report)
	if err != nil {
		return nil, err
	}
	if len(out) > MaxResultBytes {
		return nil, ErrResultTooLarge
	}
	return out, nil
}

// hostCallDetect dispatches the AdapterPlugin.Detect entry.
func hostCallDetect(inst *Instance, args []byte) ([]byte, error) {
	ap, ok := inst.Plugin.(AdapterPlugin)
	if !ok {
		return nil, ErrEntryForbidden
	}
	var env map[string]string
	if len(args) > 0 {
		if err := json.Unmarshal(args, &env); err != nil {
			return nil, err
		}
	}
	id, detected := ap.Detect(env)
	res, err := json.Marshal(AdapterDetectResult{ID: id, OK: detected})
	if err != nil {
		return nil, err
	}
	if len(res) > MaxResultBytes {
		return nil, ErrResultTooLarge
	}
	return res, nil
}

// hostCallEmit dispatches the AdapterPlugin.Emit entry.
func hostCallEmit(inst *Instance, args []byte) ([]byte, error) {
	ap, ok := inst.Plugin.(AdapterPlugin)
	if !ok {
		return nil, ErrEntryForbidden
	}
	if err := ap.Emit(args); err != nil {
		return nil, err
	}
	return nil, nil
}

// Close drains and releases an instance, removing it from the registry.
func (h *Host) Close(inst *Instance) error {
	if inst == nil {
		return nil
	}
	h.reg.remove(inst.Manifest.ID)
	return inst.Close()
}

// deriveSeed blends the host base seed with the manifest canonical digest into a
// deterministic 64-bit stream seed. It is a pure function of its inputs, free of
// time.Now and global state, so identical (Seed, manifest) pairs always yield the
// same per-instance RNG.
func deriveSeed(base uint64, m Manifest) uint64 {
	sum := sha256.Sum256(m.CanonicalBytes())
	return base ^ binary.LittleEndian.Uint64(sum[:8])
}

// entriesFor returns the entry whitelist for a concrete plugin type.
func entriesFor(p Plugin) []string {
	switch p.(type) {
	case OperatorPlugin:
		return []string{EntryApply}
	case ReporterPlugin:
		return []string{EntryBuild}
	case AdapterPlugin:
		return []string{EntryDetect, EntryEmit}
	default:
		return nil
	}
}

// contains reports whether s is present in ss.
func contains(ss []string, s string) bool {
	for i := range ss {
		if ss[i] == s {
			return true
		}
	}
	return false
}

// hasCapability reports whether c is present in cs.
func hasCapability(cs []Capability, c Capability) bool {
	for i := range cs {
		if cs[i] == c {
			return true
		}
	}
	return false
}
