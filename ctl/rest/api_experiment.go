/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"net/http"
)

// handleCreateExperiment submits an experiment task and returns 202 with location.
// Experiments share the check backend but are modeled as independent async tasks
// with their own lifecycle and cancel path, matching the PLAN-07 resource model.
func (s *Server) handleCreateExperiment(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeMethodNotAllowed(w, req)
		return
	}
	parsed, err := decodeCheckRequest(readBody(req))
	if err != nil {
		writeBadRequest(w, req, err.Error())
		return
	}
	task, submitErr := s.submit(parsed, "experiment")
	if submitErr != nil {
		writeTooManyRequests(w, req)
		return
	}
	location := "/api/" + MetaAPIVersion + "/experiments/" + task.ID
	w.Header().Set("Location", location)
	writeJSON(w, http.StatusAccepted, s.mustGet(task.ID))
}

// handleListExperiments returns experiment tasks, optionally filtered by status.
func (s *Server) handleListExperiments(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		writeMethodNotAllowed(w, req)
		return
	}
	status := req.URL.Query().Get("status")
	experiments := make([]Task, 0)
	for _, task := range s.store.list(status) {
		if task.Kind == "experiment" {
			experiments = append(experiments, task)
		}
	}
	writeJSON(w, http.StatusOK, listView{Tasks: experiments})
}

// handleCancelExperiment requests graceful cancellation of a running experiment.
// The in-flight task is marked cancelled; the worker records the archive and
// transitions to incomplete rather than dropping the partial result.
func (s *Server) handleCancelExperiment(w http.ResponseWriter, req *http.Request) {
	id := routeParam(req.Context(), "id")
	task, ok := s.store.get(id)
	if !ok {
		writeNotFound(w, req)
		return
	}
	if task.Status == StatusCompleted || task.Status == StatusFailed {
		writeConflict(w, req, "experiment is already terminal")
		return
	}
	s.store.update(id, func(t *Task) {
		t.Cancelled = true
		if t.Status == StatusPending {
			t.Status = StatusCancelled
		}
	})
	writeJSON(w, http.StatusOK, s.mustGet(id))
}

// mustGet returns a task by id, falling back to zero only for internal callers
// that have just confirmed presence.
func (s *Server) mustGet(id string) Task {
	task, _ := s.store.get(id)
	return task
}
