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

// newSched builds a scheduler with two cgroups: a default and a high-share one.
func newSched(t *testing.T) (*Scheduler, *clock.VirtualClock) {
	t.Helper()
	clk := clock.NewVirtualClock()
	s := NewScheduler(clk)
	if _, err := s.NewCgroup("sys", 1, 1<<30); err != nil {
		t.Fatalf("cgroup sys: %v", err)
	}
	if _, err := s.NewCgroup("hi", 4, 1<<30); err != nil {
		t.Fatalf("cgroup hi: %v", err)
	}
	return s, clk
}

func TestBasicSchedulingAndFairness(t *testing.T) {
	s, _ := newSched(t)
	for i := 0; i < 50; i++ {
		if _, err := s.NewTask(tid("t", i), "sys"); err != nil {
			t.Fatalf("newtask: %v", err)
		}
		_ = s.Enqueue(tid("t", i))
	}
	s.Run(2000)
	// Every task must have been scheduled at least once (fairness).
	for i := 0; i < 50; i++ {
		tk, _ := s.Task(tid("t", i))
		if tk.VRuntime() == 0 {
			t.Fatalf("task %s never scheduled", tid("t", i))
		}
	}
	// Virtual time advanced by steps*quantum.
	if s.Steps() != 2000 {
		t.Fatalf("steps = %d", s.Steps())
	}
}

func TestWeightedFairness(t *testing.T) {
	s, _ := newSched(t)
	if _, err := s.NewTask("hi-1", "hi"); err != nil {
		t.Fatalf("hi task: %v", err)
	}
	if _, err := s.NewTask("lo-1", "sys"); err != nil {
		t.Fatalf("lo task: %v", err)
	}
	_ = s.Enqueue("hi-1")
	_ = s.Enqueue("lo-1")
	s.Run(200)
	hi, _ := s.Task("hi-1")
	lo, _ := s.Task("lo-1")
	if hi.VRuntime() <= lo.VRuntime() {
		t.Fatalf("high-share task should accumulate less vruntime: hi=%d lo=%d", hi.VRuntime(), lo.VRuntime())
	}
	// Both must have run (no permanent starvation).
	if hi.VRuntime() == 0 || lo.VRuntime() == 0 {
		t.Fatalf("both tasks must run: hi=%d lo=%d", hi.VRuntime(), lo.VRuntime())
	}
}

func TestStateMachineTransitions(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	// Created -> Ready.
	if err := s.Enqueue("a"); err != nil {
		t.Fatalf("enqueue: %v", err)
	}
	tk, _ := s.Task("a")
	if tk.State() != stateReady {
		t.Fatalf("state = %d, want ready", tk.State())
	}
	// Ready -> Suspended -> Ready.
	if err := s.Suspend("a"); err != nil {
		t.Fatalf("suspend: %v", err)
	}
	if tk.State() != stateSuspended {
		t.Fatalf("state = %d, want suspended", tk.State())
	}
	if err := s.Resume("a"); err != nil {
		t.Fatalf("resume: %v", err)
	}
	if tk.State() != stateReady {
		t.Fatalf("state = %d, want ready", tk.State())
	}
	// Kill.
	if err := s.Kill("a"); err != nil {
		t.Fatalf("kill: %v", err)
	}
	if tk.State() != stateTerminated {
		t.Fatalf("state = %d, want terminated", tk.State())
	}
}

func TestInvalidTransitions(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	// Suspend a Created task is invalid.
	if err := s.Suspend("a"); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("suspend created want ErrInvalidState, got %v", err)
	}
	// Resume a Created task is invalid.
	if err := s.Resume("a"); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("resume created want ErrInvalidState, got %v", err)
	}
	_ = s.Enqueue("a")
	_ = s.Kill("a")
	// Enqueue a Terminated task is invalid.
	if err := s.Enqueue("a"); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("enqueue terminated want ErrInvalidState, got %v", err)
	}
	// Kill a Terminated task is invalid.
	if err := s.Kill("a"); !errors.Is(err, ErrInvalidState) {
		t.Fatalf("kill terminated want ErrInvalidState, got %v", err)
	}
}

func TestOOMKillsTask(t *testing.T) {
	s, _ := newSched(t)
	if _, err := s.NewCgroup("tight", 1, 100); err != nil {
		t.Fatalf("cgroup: %v", err)
	}
	_, _ = s.NewTask("a", "tight")
	_ = s.Enqueue("a")
	if err := s.Allocate("a", 60); err != nil {
		t.Fatalf("alloc 60: %v", err)
	}
	cg, _ := s.Cgroup("tight")
	if cg.MemUsed() != 60 {
		t.Fatalf("memUsed = %d", cg.MemUsed())
	}
	// Exceeding the limit triggers OOM.
	if err := s.Allocate("a", 50); !errors.Is(err, ErrOOM) {
		t.Fatalf("alloc 50 want ErrOOM, got %v", err)
	}
	tk, _ := s.Task("a")
	if tk.State() != stateTerminated || !tk.OOM() {
		t.Fatalf("task should be OOM-terminated")
	}
	if cg.MemUsed() != 0 {
		t.Fatalf("memUsed after OOM = %d", cg.MemUsed())
	}
}

func TestKillReleasesMemory(t *testing.T) {
	s, _ := newSched(t)
	_, _ = s.NewTask("a", "sys")
	_ = s.Allocate("a", 40)
	cg, _ := s.Cgroup("sys")
	if cg.MemUsed() != 40 {
		t.Fatalf("memUsed = %d", cg.MemUsed())
	}
	if err := s.Kill("a"); err != nil {
		t.Fatalf("kill: %v", err)
	}
	if cg.MemUsed() != 0 {
		t.Fatalf("memUsed after kill = %d", cg.MemUsed())
	}
}

func TestDuplicateAndMissing(t *testing.T) {
	s, _ := newSched(t)
	if _, err := s.NewTask("a", "sys"); err != nil {
		t.Fatalf("newtask: %v", err)
	}
	if _, err := s.NewTask("a", "sys"); !errors.Is(err, ErrDuplicate) {
		t.Fatalf("dup task want ErrDuplicate, got %v", err)
	}
	if _, err := s.NewTask("b", "ghost"); !errors.Is(err, ErrInvalidCgroup) {
		t.Fatalf("bad cgroup want ErrInvalidCgroup, got %v", err)
	}
	if _, err := s.NewCgroup("sys", 1, 1); !errors.Is(err, ErrDuplicate) {
		t.Fatalf("dup cgroup want ErrDuplicate, got %v", err)
	}
	if err := s.Enqueue("ghost"); !errors.Is(err, ErrNoTask) {
		t.Fatalf("enqueue missing want ErrNoTask, got %v", err)
	}
}

// tid builds a deterministic task id from a prefix and index.
func tid(p string, i int) string {
	return p + string(rune('a'+i%26)) + itoa(i)
}

// itoa is a tiny decimal formatter avoiding a strconv import in helpers.
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	var b []byte
	for i > 0 {
		b = append([]byte{byte('0' + i%10)}, b...)
		i /= 10
	}
	if neg {
		b = append([]byte{'-'}, b...)
	}
	return string(b)
}
