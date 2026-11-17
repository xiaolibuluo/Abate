/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

// Experiment is the input to a deterministic scheduler run.
//
// Seed is recorded for provenance; the scheduler itself is fully driven by
// Initial events, so any randomness must already have been resolved by the
// caller through SIM-RNG.
type Experiment struct {
	Seed    uint64
	Initial []Event
}
