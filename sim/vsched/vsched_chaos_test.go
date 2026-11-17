/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

import (
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/sim/clock"
)

// vschedScenario builds a fresh scheduler, enqueues a fixed workload and runs a
// fixed number of steps. Because the scheduler uses no randomness or wall-clock,
// the produced trace must be identical across runs.
func vschedScenario() []byte {
	clk := clock.NewVirtualClock()
	s := NewScheduler(clk)
	_, _ = s.NewCgroup("sys", 1, 1<<30)
	_, _ = s.NewCgroup("hi", 3, 1<<30)
	for i := 0; i < 120; i++ {
		cg := "sys"
		if i%4 == 0 {
			cg = "hi"
		}
		if _, err := s.NewTask(tid("t", i), cg); err != nil {
			panic(err)
		}
		_ = s.Enqueue(tid("t", i))
	}
	s.Run(800)
	return []byte(strings.Join(s.Trace(), "|"))
}

// TestVschedChaosDeterminism certifies the scheduler trace is reproducible over
// one hundred independent runs.
func TestVschedChaosDeterminism(t *testing.T) {
	determinism.MustMatch100(t, 1, vschedScenario)
}

// TestVschedTraceDeterminism is an independent weaker check that the trace is
// stable for a different seed and workload shape.
func TestVschedTraceDeterminism(t *testing.T) {
	determinism.MustMatch100(t, 2, func() []byte {
		clk := clock.NewVirtualClock()
		s := NewScheduler(clk)
		_, _ = s.NewCgroup("a", 2, 1<<30)
		for i := 0; i < 64; i++ {
			if _, err := s.NewTask(tid("w", i), "a"); err != nil {
				panic(err)
			}
			_ = s.Enqueue(tid("w", i))
		}
		s.Run(400)
		return []byte(strings.Join(s.Trace(), ","))
	})
}
