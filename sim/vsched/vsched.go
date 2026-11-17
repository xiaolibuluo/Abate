/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package vsched is a deterministic virtual scheduler modeling tasks, cgroups,
// weighted fair CPU sharing and memory-limit OOM kills without a real OS.
//
// Every decision is resolved through the injected virtual clock and explicit
// faults, so a rebuilt scheduler produces a byte-identical run. No goroutine,
// wall-clock read or pseudo-random source is ever used.
package vsched

import (
	"errors"
	"sort"

	"abate/sim/clock"
)

// DefaultQuantum is the virtual nanoseconds charged per scheduling step.
const DefaultQuantum = clock.Duration(1000)

// Sentinel errors of the virtual scheduler.
var (
	// ErrNoTask reports an operation referencing an unknown task.
	ErrNoTask = errors.New("vsched: no such task")

	// ErrNoCgroup reports an operation referencing an unknown cgroup.
	ErrNoCgroup = errors.New("vsched: no such cgroup")

	// ErrInvalidState reports a disallowed state transition.
	ErrInvalidState = errors.New("vsched: invalid state transition")

	// ErrOOM reports an allocation that breached the cgroup memory limit.
	ErrOOM = errors.New("vsched: out of memory")

	// ErrDuplicate reports a reused task or cgroup id.
	ErrDuplicate = errors.New("vsched: duplicate id")

	// ErrInvalidCgroup reports an empty or unknown cgroup reference.
	ErrInvalidCgroup = errors.New("vsched: invalid cgroup")

	// ErrUnknownFault reports an unrecognized fault kind.
	ErrUnknownFault = errors.New("vsched: unknown fault kind")
)

// Scheduler is a deterministic virtual CPU scheduler.
//
// It maintains tasks and cgroups, performs weighted-fair scheduling each Step,
// enforces per-cgroup memory limits through OOM kills, and supports the fault
// injector. Task selection always iterates task ids in sorted order so the
// outcome is independent of map traversal order.
type Scheduler struct {
	clk     *clock.VirtualClock
	quantum clock.Duration
	tasks   map[string]*Task
	cgroups map[string]*Cgroup
	order   []string
	running *Task
	steps   int64
	trace   []string
}

// NewScheduler builds an empty scheduler driven by the given virtual clock.
func NewScheduler(clk *clock.VirtualClock) *Scheduler {
	return &Scheduler{
		clk:     clk,
		quantum: DefaultQuantum,
		tasks:   make(map[string]*Task),
		cgroups: make(map[string]*Cgroup),
	}
}

// NewCgroup registers a cgroup with the given CPU shares and memory limit.
func (s *Scheduler) NewCgroup(name string, shares, memLimit int64) (*Cgroup, error) {
	if name == "" {
		return nil, ErrInvalidCgroup
	}
	if _, ok := s.cgroups[name]; ok {
		return nil, ErrDuplicate
	}
	if shares < 1 {
		shares = 1
	}
	if memLimit < 0 {
		memLimit = 0
	}
	cg := &Cgroup{name: name, shares: shares, baseShares: shares, memLimit: memLimit}
	s.cgroups[name] = cg
	return cg, nil
}

// NewTask creates a task attached to an existing cgroup in the Created state.
func (s *Scheduler) NewTask(id, cgroup string) (*Task, error) {
	if id == "" {
		return nil, ErrNoTask
	}
	if _, ok := s.tasks[id]; ok {
		return nil, ErrDuplicate
	}
	cg, ok := s.cgroups[cgroup]
	if !ok {
		return nil, ErrInvalidCgroup
	}
	t := &Task{id: id, state: stateCreated, cgroup: cgroup}
	s.tasks[id] = t
	s.order = append(s.order, id)
	sort.Strings(s.order)
	cg.tasks++
	return t, nil
}

// TaskIDs returns all task ids in lexicographic order.
func (s *Scheduler) TaskIDs() []string {
	ids := make([]string, len(s.order))
	copy(ids, s.order)
	return ids
}

// Task returns the task with the given id.
func (s *Scheduler) Task(id string) (*Task, bool) {
	t, ok := s.tasks[id]
	return t, ok
}

// Cgroup returns the cgroup with the given name.
func (s *Scheduler) Cgroup(name string) (*Cgroup, bool) {
	cg, ok := s.cgroups[name]
	return cg, ok
}

// Running returns the id of the currently running task, or "" when idle.
func (s *Scheduler) Running() string {
	if s.running != nil {
		return s.running.id
	}
	return ""
}

// Steps returns the number of scheduling steps executed.
func (s *Scheduler) Steps() int64 { return s.steps }

// Trace returns the per-step running task id ("-" when idle).
func (s *Scheduler) Trace() []string { return s.trace }

// Enqueue moves a Created task into the Ready queue.
func (s *Scheduler) Enqueue(id string) error {
	t, ok := s.tasks[id]
	if !ok {
		return ErrNoTask
	}
	if t.state != stateCreated {
		return ErrInvalidState
	}
	t.state = stateReady
	return nil
}

// Suspend moves a Ready or Running task into the Suspended state.
func (s *Scheduler) Suspend(id string) error {
	t, ok := s.tasks[id]
	if !ok {
		return ErrNoTask
	}
	if t.state != stateReady && t.state != stateRunning {
		return ErrInvalidState
	}
	t.state = stateSuspended
	if s.running == t {
		s.running = nil
	}
	return nil
}

// Resume moves a Suspended task back to Ready.
func (s *Scheduler) Resume(id string) error {
	t, ok := s.tasks[id]
	if !ok {
		return ErrNoTask
	}
	if t.state != stateSuspended {
		return ErrInvalidState
	}
	t.state = stateReady
	return nil
}

// Kill terminates a task and releases its memory back to its cgroup.
func (s *Scheduler) Kill(id string) error {
	t, ok := s.tasks[id]
	if !ok {
		return ErrNoTask
	}
	if t.state == stateTerminated {
		return ErrInvalidState
	}
	s.terminate(t)
	return nil
}
