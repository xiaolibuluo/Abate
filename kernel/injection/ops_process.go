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
	"abate/sim/vsched"
)

// applyProcess freezes and terminates a virtual task through SIM-VSCHED.
func applyProcess(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	if params.class != "process-kill" {
		return ErrUnknownOperator
	}
	if err := w.scheduler.Suspend(params.target); err != nil {
		return fmt.Errorf("injection: %w", err)
	}
	if op.Params()["mode"] == "suspend" {
		return nil
	}
	if err := w.scheduler.Kill(params.target); err != nil {
		return fmt.Errorf("injection: %w", err)
	}
	return nil
}

// recoverProcess restores the task through Resume or a baseline rebuild.
func recoverProcess(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	if params.class != "process-kill" {
		return ErrUnknownOperator
	}
	if op.Params()["mode"] == "suspend" {
		if err := w.scheduler.Resume(params.target); err != nil {
			return fmt.Errorf("injection: %w", err)
		}
		return nil
	}
	// A terminated task cannot be resumed; the invalid-state result is the
	// explicit signal that the baseline snapshot must rebuild the scheduler.
	if err := w.scheduler.Resume(params.target); err != nil &&
		!errors.Is(err, vsched.ErrInvalidState) {
		return fmt.Errorf("injection: %w", err)
	}
	w.rebuildScheduler()
	return nil
}
