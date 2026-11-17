/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package sample provides three reference plugin implementations (operator,
// reporter, adapter) together with their host factories. They form the
// acceptance trio for the plugin ecosystem: each is a concrete, loadable plugin
// that the control plane can sign, load, call, and hot-swap.
package sample

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"abate/kernel/algebra"
	"abate/plugin/host"
	"abate/sim/rand"
)

// OperatorPlugin is a deterministic perturbation operator plugin. Its Apply is a
// pure function of the input bytes and the injected deterministic RNG/clock
// backends, so repeated calls with identical backend state are reproducible.
type OperatorPlugin struct {
	manifest host.Manifest
}

// NewOperatorPlugin constructs a sample operator plugin from its manifest.
func NewOperatorPlugin(m host.Manifest) *OperatorPlugin {
	return &OperatorPlugin{manifest: m}
}

// Manifest returns the signed metadata of the plugin.
func (p *OperatorPlugin) Manifest() host.Manifest { return p.manifest }

// Type returns the plugin kind.
func (*OperatorPlugin) Type() host.PluginType { return host.PluginTypeOperator }

// Close releases no resources for the in-memory sample operator.
func (*OperatorPlugin) Close() error { return nil }

// Describe returns the algebra operator this plugin represents. It is a simple
// XOR keystream perturbation used to demonstrate a real, deterministic operator.
func (*OperatorPlugin) Describe() algebra.Operator {
	return algebra.NewBasic(
		algebra.FormatOperatorID(algebra.FamilyMutation, "sample-perturb", "1"),
		algebra.FamilyMutation,
		algebra.Params{"algo": "xor-keystream", "family": "m"},
		[]algebra.TargetKey{"plugin/sample"},
		false,
	)
}

// Apply performs a deterministic XOR keystream perturbation of the input. It draws
// a keystream from the injected backend RNG when present; otherwise it derives a
// seed from a length-prefixed hash of the input. The backend clock reading is mixed
// into a trailing salt byte. The RNG stream is snapshot and restored so repeated
// calls on the same instance remain reproducible.
func (*OperatorPlugin) Apply(ctx context.Context, input []byte) ([]byte, error) {
	be := host.Backend(ctx)
	var rng *rand.Stream
	if be != nil && be.RNG != nil {
		rng = be.RNG
	} else {
		rng = rand.New(deriveSeedFromInput(input))
	}
	seed0, draws0 := rng.State()
	ks := rng.Bytes(len(input) + 8)
	_ = rng.Restore(seed0, draws0)

	var clockSalt byte
	if be != nil && be.Clock != nil {
		clockSalt = byte(int64(be.Clock.Now()) & 0xff)
	}
	out := make([]byte, len(input)+1)
	for i := range input {
		out[i] = input[i] ^ ks[i]
	}
	out[len(input)] = clockSalt ^ ks[len(input)]
	return out, nil
}

// OperatorFactory instantiates the sample operator plugin from a verified
// manifest. It fails closed if the manifest is not an operator.
func OperatorFactory(m host.Manifest) (host.Plugin, error) {
	if m.Type != host.PluginTypeOperator {
		return nil, fmt.Errorf("sample: manifest %q is not an operator (type %q)", m.ID, m.Type)
	}
	return NewOperatorPlugin(m), nil
}

// deriveSeedFromInput produces a stable fallback RNG seed from a length-prefixed
// hash of the input, used only when no backend RNG is injected.
func deriveSeedFromInput(input []byte) uint64 {
	h := sha256.Sum256(input)
	var tmp [8]byte
	copy(tmp[:], h[:8])
	return binary.LittleEndian.Uint64(tmp[:])
}
