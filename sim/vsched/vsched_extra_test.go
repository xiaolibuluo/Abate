/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo/163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

import (
	"errors"
	"testing"
)

// TestCgroupAccessors covers cgroup getters and the share clamp.
func TestCgroupAccessors(t *testing.T) {
	s, _ := newSched(t)
	cg, ok := s.Cgroup("sys")
	if !ok || cg == nil {
		t.Fatalf("cgroup sys missing")
	}
	if cg.Name() != "sys" {
		t.Fatalf("name = %q", cg.Name())
	}
	if cg.Shares() != 1 || cg.BaseShares() != 1 {
		t.Fatalf("shares = %d base = %d", cg.Shares(), cg.BaseShares())
	}
	if cg.MemLimit() != 1<<30 {
		t.Fatalf("memLimit = %d", cg.MemLimit())
	}
	if cg.TaskCount() != 0 {
		t.Fatalf("taskCount = %d", cg.TaskCount())
	}
	if cg.MemUsed() != 0 {
		t.Fatalf("memUsed = %d", cg.MemUsed())
	}
	// Share clamping to a minimum of one.
	cg.SetShares(0)
	if cg.Shares() != 1 {
		t.Fatalf("SetShares(0) should clamp to 1, got %d", cg.Shares())
	}
}

// TestNewCgroupValidation covers name, duplicate, share and limit clamp branches.
func TestNewCgroupValidation(t *testing.T) {
	s, _ := newSched(t)
	if _, err := s.NewCgroup("", 1, 1); !errors.Is(err, ErrInvalidCgroup) {
		t.Fatalf("empty name want ErrInvalidCgroup, got %v", err)
	}
	cg, err := s.NewCgroup("clamped", 0, -5)
	if err != nil {
		t.Fatalf("clamped cgroup: %v", err)
	}
	if cg.Shares() != 1 {
		t.Fatalf("shares should clamp to 1, got %d", cg.Shares())
	}
	if cg.MemLimit() != 0 {
		t.Fatalf("memLimit should clamp to 0, got %d", cg.MemLimit())
	}
}

// TestTaskAccessors covers the Task getter surface.
func TestTaskAccessors(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	_ = s.Allocate("a", 7)
	tk, _ := s.Task("a")
	if tk.ID() != "a" {
		t.Fatalf("id = %q", tk.ID())
	}
	if tk.State() != stateCreated {
		t.Fatalf("state = %d", tk.State())
	}
	if tk.Cgroup() != "sys" {
		t.Fatalf("cgroup = %q", tk.Cgroup())
	}
	if tk.MemBytes() != 7 {
		t.Fatalf("mem = %d", tk.MemBytes())
	}
	if tk.VRuntime() != 0 {
		t.Fatalf("vruntime = %d", tk.VRuntime())
	}
	if tk.OOM() {
		t.Fatalf("oom should be false")
	}
}

// TestTaskIDsSorted covers the public sorted-id accessor.
func TestTaskIDsSorted(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("c", "sys")
	_, _ = s.NewTask("a", "sys")
	_, _ = s.NewTask("b", "sys")
	ids := s.TaskIDs()
	if len(ids) != 3 || ids[0] != "a" || ids[1] != "b" || ids[2] != "c" {
		t.Fatalf("ids not sorted: %v", ids)
	}
}

// TestRunningReportsTask covers the non-idle branch of Running.
func TestRunningReportsTask(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	_ = s.Enqueue("a")
	s.Run(1)
	if s.Running() == "" {
		t.Fatalf("running should report a task id")
	}
}

// TestSuspendTerminated covers the invalid-transition branch of Suspend.
func TestSuspendTerminated(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	_ = s.Kill("a")
	if err := s.Suspend("a"); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("suspend terminated want ErrInvalidState, got %v", err)
	}
}

// TestResumeNonSuspended covers the invalid-transition branch of Resume.
func TestResumeNonSuspended(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	_ = s.Enqueue("a")
	if err := s.Resume("a"); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("resume ready want ErrInvalidState, got %v", err)
	}
}

// TestKillMissing covers the lookup-error branch of Kill.
func TestKillMissing(t *testing.T) {
	s, _ := newSched(t)
	if err := s.Kill("ghost"); !errors.Is(err, ErrNoTask) {
		t.Fatalf("kill missing want ErrNoTask, got %v", err)
	}
}

// TestAllocateOnTerminated covers the terminated-guard branch of Allocate.
func TestAllocateOnTerminated(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	_ = s.Kill("a")
	if err := s.Allocate("a", 1); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("allocate terminated want ErrInvalidState, got %v", err)
	}
}

// TestRecoverUnknownFault covers the default branch of RecoverFault.
func TestRecoverUnknownFault(t *testing.T) {
	s, _ := newSched(t)
	if err := s.RecoverFault(Fault{Kind: FaultKind("bogus"), Target: "x"}); !errors.Is(err, ErrUnknownFault) {
		t.Fatalf("recover unknown want ErrUnknownFault, got %v", err)
	}
	if err := s.RecoverFault(Fault{Kind: FaultThrottle, Target: "ghost"}); !errors.Is(err, ErrNoCgroup) {
		t.Fatalf("recover missing cgroup want ErrNoCgroup, got %v", err)
	}
}

// TestTerminateOrphanCgroup covers the missing-cgroup branch of terminate and
// Allocate by registering a task against a cgroup that is not in the scheduler.
func TestTerminateOrphanCgroup(t *testing.T) {
	s, _ := newSched(t)
	// A task whose cgroup is absent from the scheduler maps.
	orphan := &Task{id: "orphan", state: stateReady, cgroup: "ghost"}
	s.tasks["orphan"] = orphan
	s.terminate(orphan)
	if orphan.State() != stateTerminated {
		t.Fatalf("orphan should terminate cleanly")
	}
	// A non-terminated orphan exercises the cgroup lookup branch of Allocate.
	orphan2 := &Task{id: "orphan2", state: stateCreated, cgroup: "ghost"}
	s.tasks["orphan2"] = orphan2
	if err := s.Allocate("orphan2", 1); !errors.Is(err, ErrInvalidCgroup) {
		t.Fatalf("allocate orphan cgroup want ErrInvalidCgroup, got %v", err)
	}
}
