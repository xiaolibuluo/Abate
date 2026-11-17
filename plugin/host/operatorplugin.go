/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"context"

	"abate/kernel/algebra"
)

// OperatorPlugin is a plugin that applies a deterministic perturbation operator
// to an input payload and produces a deterministic output. The operator identity
// is exposed via Describe so the host and scheduler can reason about it.
type OperatorPlugin interface {
	Plugin

	// Describe returns the underlying algebra operator this plugin represents.
	Describe() algebra.Operator

	// Apply runs the operator on input under the host-provided deterministic
	// backend (clock + RNG carried in ctx). Apply must be a pure function of
	// (input, granted capabilities, backend state) so repeated calls with the
	// same backend stream are reproducible.
	Apply(ctx context.Context, input []byte) ([]byte, error)
}
