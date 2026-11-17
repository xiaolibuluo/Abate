/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

// terminate releases a task's resources and marks it Terminated.
func (s *Scheduler) terminate(t *Task) {
	if cg, ok := s.cgroups[t.cgroup]; ok {
		cg.memUsed -= t.mem
		if cg.memUsed < 0 {
			cg.memUsed = 0
		}
		cg.tasks--
		if cg.tasks < 0 {
			cg.tasks = 0
		}
	}
	t.mem = 0
	t.state = stateTerminated
	if s.running == t {
		s.running = nil
	}
	for i, x := range s.order {
		if x == t.id {
			s.order = append(s.order[:i], s.order[i+1:]...)
			break
		}
	}
}

// Allocate charges bytes to a task's cgroup, triggering an OOM kill when the
// cgroup memory limit would be exceeded.
func (s *Scheduler) Allocate(id string, bytes int64) error {
	t, ok := s.tasks[id]
	if !ok {
		return ErrNoTask
	}
	if t.state == stateTerminated {
		return ErrInvalidState
	}
	cg, ok := s.cgroups[t.cgroup]
	if !ok {
		return ErrInvalidCgroup
	}
	if cg.memUsed+bytes > cg.memLimit {
		s.terminate(t)
		t.oom = true
		return ErrOOM
	}
	t.mem += bytes
	cg.memUsed += bytes
	return nil
}

// Step performs one weighted-fair scheduling decision and advances virtual time.
func (s *Scheduler) Step() {
	var best *Task
	for _, id := range s.order {
		t := s.tasks[id]
		if t.state != stateReady && t.state != stateRunning {
			continue
		}
		if best == nil {
			best = t
			continue
		}
		bp, tp := best.priority(s), t.priority(s)
		if tp < bp || (tp == bp && t.id < best.id) {
			best = t
		}
	}
	if best == nil {
		s.running = nil
		_, _ = s.clk.Advance(s.quantum)
		s.steps++
		s.trace = append(s.trace, "-")
		return
	}
	if s.running != nil && s.running != best && s.running.state == stateRunning {
		s.running.state = stateReady
	}
	best.state = stateRunning
	s.running = best
	best.vruntime += int64(s.quantum)
	_, _ = s.clk.Advance(s.quantum)
	s.steps++
	s.trace = append(s.trace, best.id)
}

// Run executes the given number of scheduling steps.
func (s *Scheduler) Run(steps int) {
	for i := 0; i < steps; i++ {
		s.Step()
	}
}

// InjectFault applies a scheduler-level fault to a task or cgroup.
func (s *Scheduler) InjectFault(f Fault) error {
	switch f.Kind {
	case FaultOOM:
		t, ok := s.tasks[f.Target]
		if !ok {
			return ErrNoTask
		}
		s.terminate(t)
		t.oom = true
		return nil
	case FaultSuspend:
		return s.Suspend(f.Target)
	case FaultThrottle:
		cg, ok := s.cgroups[f.Target]
		if !ok {
			return ErrNoCgroup
		}
		cg.SetShares(f.Value)
		return nil
	default:
		return ErrUnknownFault
	}
}

// RecoverFault restores state affected by a prior fault.
func (s *Scheduler) RecoverFault(f Fault) error {
	switch f.Kind {
	case FaultThrottle:
		cg, ok := s.cgroups[f.Target]
		if !ok {
			return ErrNoCgroup
		}
		cg.shares = cg.baseShares
		return nil
	default:
		return ErrUnknownFault
	}
}
