/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package net

import "abate/sim/clock"

// Fault is a link-level fault-injection specification.
//
// A and B select the link endpoints. Value is the fault parameter interpreted
// per Kind (nanoseconds for latency, bytes per second for bandwidth, parts per
// million for loss; ignored for partition).
type Fault struct {
	// Kind selects the fault primitive.
	Kind FaultKind

	// A and B are the link endpoints.
	A string
	B string

	// Value is the fault parameter interpreted per Kind.
	Value int64
}

// InjectFault applies a link-level fault to the named link.
func (n *Network) InjectFault(f Fault) error {
	lk := n.linkBetween(f.A, f.B)
	if lk == nil {
		return ErrNoLink
	}
	switch f.Kind {
	case FaultLatency:
		lk.SetLatency(clock.Duration(f.Value))
	case FaultBandwidth:
		lk.SetBandwidth(uint64(f.Value)) //nolint:gosec // Value is a non-negative capacity.
	case FaultLoss:
		lk.SetLossRate(float64(f.Value) / 1e6)
	case FaultPartition:
		_ = lk.Partition()
	default:
		return ErrUnknownFault
	}
	return nil
}

// RecoverFault restores the named link to its baseline configuration.
func (n *Network) RecoverFault(f Fault) error {
	lk := n.linkBetween(f.A, f.B)
	if lk == nil {
		return ErrNoLink
	}
	lk.Recover()
	return nil
}
