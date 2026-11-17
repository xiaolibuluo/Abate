/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"errors"
	"fmt"

	"abate/kernel/algebra"
	"abate/sim/fs"
	"abate/sim/vsched"
)

// applyResource injects disk, CPU or memory pressure through SIM-FS/SIM-VSCHED.
func applyResource(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	switch params.class {
	case "disk-io-latency":
		return w.filesystem.InjectFault(fs.Fault{Kind: fs.FaultLatency, Value: params.value})
	case "disk-full":
		w.quotaBytes = 0
		return w.filesystem.InjectFault(fs.Fault{
			Kind: fs.FaultBadBlock, Path: filePath, Offset: 0, Value: params.value,
		})
	case "cpu-pressure":
		return w.scheduler.InjectFault(vsched.Fault{
			Kind: vsched.FaultThrottle, Target: cgroupName, Value: params.value,
		})
	case "memory-pressure":
		if err := w.scheduler.Allocate(params.target, params.value); err != nil &&
			!errors.Is(err, vsched.ErrOOM) {
			return fmt.Errorf("injection: %w", err)
		}
		return nil
	default:
		return ErrUnknownOperator
	}
}

// recoverResource restores disk, CPU or memory pressure through the dual call.
func recoverResource(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	switch params.class {
	case "disk-io-latency":
		return w.filesystem.InjectFault(fs.Fault{
			Kind: fs.FaultLatency, Value: int64(fs.DefaultIOLatency),
		})
	case "disk-full":
		w.quotaBytes = fs.MaxTotalBytes
		return w.filesystem.RecoverFault(filePath)
	case "cpu-pressure":
		return w.scheduler.RecoverFault(vsched.Fault{
			Kind: vsched.FaultThrottle, Target: cgroupName,
		})
	case "memory-pressure":
		w.rebuildScheduler()
		return nil
	default:
		return ErrUnknownOperator
	}
}
