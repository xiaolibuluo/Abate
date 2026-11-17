/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"strconv"
	"sync"
)

// taskStore is the in-memory, concurrency-safe task registry of the server.
//
// The identifier sequence is a monotonic counter (no random UUID), so two runs
// with identical inputs produce identical task identifiers — a determinism
// requirement for reproducible CI traces.
type taskStore struct {
	mu     sync.Mutex
	tasks  map[string]*Task
	seq    uint64
	active int
}

// newTaskStore returns an empty store.
func newTaskStore() *taskStore {
	return &taskStore{tasks: make(map[string]*Task)}
}

// newTask reserves and returns a task with the next deterministic identifier.
func (s *taskStore) newTask(kind, language, format string) *Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	s.active++
	task := &Task{
		ID:       "t" + strconv.FormatUint(s.seq, 10),
		Kind:     kind,
		Language: language,
		Format:   format,
		Status:   StatusPending,
	}
	s.tasks[task.ID] = task
	return task
}

// get returns a task copy by identifier.
func (s *taskStore) get(id string) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task, ok := s.tasks[id]
	if !ok {
		return Task{}, false
	}
	return *task, true
}

// update applies a mutation to a task under the store lock.
func (s *taskStore) update(id string, fn func(*Task)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if task, ok := s.tasks[id]; ok {
		fn(task)
	}
}

// finish decrements the active counter when a task leaves the running set.
func (s *taskStore) finish(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; ok {
		s.active--
	}
}

// activeCount returns the number of in-flight tasks (queue occupancy).
func (s *taskStore) activeCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active
}

// list returns tasks filtered by status, ordered by identifier ascending so the
// listing is deterministic regardless of map iteration order.
func (s *taskStore) list(status string) []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		if status != "" && task.Status != TaskStatus(status) {
			continue
		}
		out = append(out, *task)
	}
	sortTasks(out)
	return out
}

// sortTasks orders tasks by identifier ascending.
func sortTasks(tasks []Task) {
	for i := 1; i < len(tasks); i++ {
		for j := i; j > 0 && tasks[j-1].ID > tasks[j].ID; j-- {
			tasks[j-1], tasks[j] = tasks[j], tasks[j-1]
		}
	}
}
