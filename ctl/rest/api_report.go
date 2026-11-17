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

// reportContentType maps a report format to its response content type.
func reportContentType(format string) string {
	switch format {
	case "json", "sarif":
		return "application/json"
	default:
		return "text/plain; charset=utf-8"
	}
}

// handleExperimentReport streams the rendered report of a completed experiment.
// A still-running experiment returns 409 so clients poll the task status first.
func (s *Server) handleExperimentReport(w http.ResponseWriter, req *http.Request) {
	id := routeParam(req.Context(), "id")
	task, ok := s.store.get(id)
	if !ok {
		writeNotFound(w, req)
		return
	}
	if !isTerminal(task.Status) {
		writeConflict(w, req, "experiment is not ready; poll the task status")
		return
	}
	w.Header().Set("Content-Type", reportContentType(task.Format))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(task.Report)
}

// handleExperimentArchive returns the experiment archive payload. The archive is
// the deterministic report bytes; this endpoint mirrors the CLI replay/verify
// artifact so the three entry points share one serialized form.
func (s *Server) handleExperimentArchive(w http.ResponseWriter, req *http.Request) {
	id := routeParam(req.Context(), "id")
	task, ok := s.store.get(id)
	if !ok {
		writeNotFound(w, req)
		return
	}
	if !isTerminal(task.Status) {
		writeConflict(w, req, "experiment is not ready; poll the task status")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(task.Report)
}

// isTerminal reports whether a task has reached a stable lifecycle state.
func isTerminal(status TaskStatus) bool {
	switch status {
	case StatusCompleted, StatusFailed, StatusIncomplete, StatusCancelled:
		return true
	default:
		return false
	}
}
