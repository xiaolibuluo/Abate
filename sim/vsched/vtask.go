/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package vsched

// Task lifecycle states. They are untyped integer constants so the simulator can
// expose them without introducing a second exported type in this file.
const (
	stateCreated = iota
	stateReady
	stateRunning
	stateSuspended
	stateTerminated
)

// priorityScale converts vruntime into a comparable scheduling priority. A
// higher cgroup share lowers the value, so weighted tasks are selected first.
const priorityScale = int64(1024)

// Task is a virtual schedulable unit attached to a cgroup.
type Task struct {
	id       string
	state    int
	cgroup   string
	vruntime int64
	mem      int64
	oom      bool
}

// ID returns the task id.
func (t *Task) ID() string { return t.id }

// State returns the current lifecycle state constant.
func (t *Task) State() int { return t.state }

// Cgroup returns the name of the owning cgroup.
func (t *Task) Cgroup() string { return t.cgroup }

// VRuntime returns the accumulated virtual runtime in nanoseconds.
func (t *Task) VRuntime() int64 { return t.vruntime }

// MemBytes returns the bytes currently charged to the task.
func (t *Task) MemBytes() int64 { return t.mem }

// OOM reports whether the task was killed by an out-of-memory fault.
func (t *Task) OOM() bool { return t.oom }

// priority returns the weighted scheduling priority used by the scheduler.
func (t *Task) priority(s *Scheduler) int64 {
	w := int64(1)
	if cg, ok := s.cgroups[t.cgroup]; ok && cg.shares > 0 {
		w = cg.shares
	}
	return t.vruntime * priorityScale / w
}
