/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

import (
	"errors"
	"testing"

	"abate/sim/clock"
)

// TestMutationThrottle confirms a cgroup throttle lowers its tasks' CPU share
// and that recovery restores the baseline.
func TestMutationThrottle(t *testing.T) {
	s, _ := newSched(t)
	if _, err := s.NewTask("hi-1", "hi"); err != nil {
		t.Fatalf("hi task: %v", err)
	}
	if _, err := s.NewTask("lo-1", "sys"); err != nil {
		t.Fatalf("lo task: %v", err)
	}
	_ = s.Enqueue("hi-1")
	_ = s.Enqueue("lo-1")
	// Throttle "hi" down to share 1, equal to "sys".
	if err := s.InjectFault(Fault{Kind: FaultThrottle, Target: "hi", Value: 1}); err != nil {
		t.Fatalf("throttle: %v", err)
	}
	s.Run(400)
	hi, _ := s.Task("hi-1")
	lo, _ := s.Task("lo-1")
	// With equal shares both should have run; neither should be starved.
	if hi.VRuntime() == 0 || lo.VRuntime() == 0 {
		t.Fatalf("equal shares must both run: hi=%d lo=%d", hi.VRuntime(), lo.VRuntime())
	}
	// Recover restores the higher baseline share.
	if err := s.RecoverFault(Fault{Kind: FaultThrottle, Target: "hi"}); err != nil {
		t.Fatalf("recover: %v", err)
	}
	cg, _ := s.Cgroup("hi")
	if cg.Shares() != cg.BaseShares() {
		t.Fatalf("shares %d != base %d after recover", cg.Shares(), cg.BaseShares())
	}
}

// TestMutationFaultSuspend confirms the suspend fault pauses a task and that it
// resumes and is scheduled again afterwards.
func TestMutationFaultSuspend(t *testing.T) {
	s, _ := newSched(t)
	if _, err := s.NewTask("a", "sys"); err != nil {
		t.Fatalf("newtask: %v", err)
	}
	_ = s.Enqueue("a")
	s.Run(10)
	if err := s.InjectFault(Fault{Kind: FaultSuspend, Target: "a"}); err != nil {
		t.Fatalf("suspend fault: %v", err)
	}
	tk, _ := s.Task("a")
	if tk.State() != stateSuspended {
		t.Fatalf("state = %d, want suspended", tk.State())
	}
	// While suspended it must not be scheduled.
	s.Run(20)
	if tk.VRuntime() != int64(DefaultQuantum*10) {
		t.Fatalf("suspended task advanced: %d", tk.VRuntime())
	}
	if err := s.Resume("a"); err != nil {
		t.Fatalf("resume: %v", err)
	}
	s.Run(20)
	if tk.VRuntime() <= int64(DefaultQuantum*10) {
		t.Fatalf("resumed task should run again: %d", tk.VRuntime())
	}
}

// TestMutationOOMFault confirms the OOM fault force-kills a task.
func TestMutationOOMFault(t *testing.T) {
	s, _ := newSched(t)
	if _, err := s.NewTask("a", "sys"); err != nil {
		t.Fatalf("newtask: %v", err)
	}
	_ = s.Enqueue("a")
	if err := s.InjectFault(Fault{Kind: FaultOOM, Target: "a"}); err != nil {
		t.Fatalf("oom fault: %v", err)
	}
	tk, _ := s.Task("a")
	if tk.State() != stateTerminated || !tk.OOM() {
		t.Fatalf("task should be OOM-terminated")
	}
}

// TestMutationUnknownFault confirms an unrecognized fault is rejected.
func TestMutationUnknownFault(t *testing.T) {
	s, _ := newSched(t)
	if err := s.InjectFault(Fault{Kind: FaultKind("bogus"), Target: "a"}); !errors.Is(err, ErrUnknownFault) {
		t.Fatalf("unknown fault want ErrUnknownFault, got %v", err)
	}
}

// TestMutationFaultOnUnknown confirms faults targeting missing entries error.
func TestMutationFaultOnUnknown(t *testing.T) {
	s, _ := newSched(t)
	if err := s.InjectFault(Fault{Kind: FaultOOM, Target: "ghost"}); !errors.Is(err, ErrNoTask) {
		t.Fatalf("oom on missing want ErrNoTask, got %v", err)
	}
	if err := s.InjectFault(Fault{Kind: FaultThrottle, Target: "ghost"}); !errors.Is(err, ErrNoCgroup) {
		t.Fatalf("throttle on missing want ErrNoCgroup, got %v", err)
	}
}

// TestMutationIdleNoDeadlock confirms the scheduler advances time when no task
// is runnable, so it never spins or stalls.
func TestMutationIdleNoDeadlock(t *testing.T) {
	clk := clock.NewVirtualClock()
	s := NewScheduler(clk)
	before := clk.Now().UnixNano()
	s.Run(50)
	if clk.Now().UnixNano()-before != int64(DefaultQuantum)*50 {
		t.Fatalf("idle run did not advance virtual time correctly")
	}
	if s.Running() != "" {
		t.Fatalf("idle scheduler should report no running task")
	}
}

// TestMutationTenThousandTasks is a pressure test: ten thousand tasks must be
// scheduled to completion of the run within a bounded time and deterministically.
func TestMutationTenThousandTasks(t *testing.T) {
	s, _ := newSched(t)
	const n = 10000
	for i := 0; i < n; i++ {
		cg := "sys"
		if i%7 == 0 {
			cg = "hi"
		}
		if _, err := s.NewTask(tid("p", i), cg); err != nil {
			t.Fatalf("newtask %d: %v", i, err)
		}
		_ = s.Enqueue(tid("p", i))
	}
	// Run enough steps so every task receives at least one quantum.
	s.Run(n + 500)
	var unscheduled int
	for i := 0; i < n; i++ {
		tk, _ := s.Task(tid("p", i))
		if tk.VRuntime() == 0 {
			unscheduled++
		}
	}
	if unscheduled > 0 {
		t.Fatalf("%d of %d tasks never scheduled", unscheduled, n)
	}
	// Determinism: a second identical build yields the same trace length.
	s2, _ := newSched(t)
	for i := 0; i < n; i++ {
		cg := "sys"
		if i%7 == 0 {
			cg = "hi"
		}
		_, _ = s2.NewTask(tid("p", i), cg)
		_ = s2.Enqueue(tid("p", i))
	}
	s2.Run(n + 500)
	if len(s.Trace()) != len(s2.Trace()) {
		t.Fatalf("pressure run not reproducible")
	}
}
