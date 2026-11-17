/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import (
	"abate/kernel/algebra"
	"abate/sim/clock"
)

// applyClock shifts the single virtual clock through the SIM-CLOCK Offset API.
func applyClock(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	delta := clock.Duration(params.value)
	w.clk.Offset(delta)
	w.offsetNS += delta
	return nil
}

// recoverClock applies the exact dual offset, so Now and State return to baseline.
func recoverClock(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	delta := clock.Duration(params.value)
	w.clk.Offset(-delta)
	w.offsetNS -= delta
	return nil
}
