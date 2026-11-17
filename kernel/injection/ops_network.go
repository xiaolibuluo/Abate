/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"fmt"

	"abate/kernel/algebra"
	"abate/sim/net"
)

// applyNetwork injects one of the four virtual link faults through SIM-NET.
func applyNetwork(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	fault := net.Fault{A: linkA, B: linkB, Value: params.value}
	switch params.class {
	case "network-latency":
		fault.Kind = net.FaultLatency
	case "network-partition":
		fault.Kind = net.FaultPartition
	case "network-loss":
		fault.Kind = net.FaultLoss
	case "network-bandwidth":
		fault.Kind = net.FaultBandwidth
	default:
		return ErrUnknownOperator
	}
	if err := w.network.InjectFault(fault); err != nil {
		return fmt.Errorf("injection: %w", err)
	}
	return nil
}

// recoverNetwork restores the virtual link through the SIM-NET dual call.
func recoverNetwork(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	fault := net.Fault{A: linkA, B: linkB, Value: params.value}
	switch params.class {
	case "network-latency":
		fault.Kind = net.FaultLatency
	case "network-partition":
		fault.Kind = net.FaultPartition
	case "network-loss":
		fault.Kind = net.FaultLoss
	case "network-bandwidth":
		fault.Kind = net.FaultBandwidth
	default:
		return ErrUnknownOperator
	}
	if err := w.network.RecoverFault(fault); err != nil {
		return fmt.Errorf("injection: %w", err)
	}
	return nil
}
