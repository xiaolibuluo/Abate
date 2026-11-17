/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

// decodeCheckRequest parses and validates a task creation body.
func decodeCheckRequest(body []byte) (multiCheckRequest, error) {
	var req multiCheckRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return multiCheckRequest{}, errBadJSON
	}
	req.Language = strings.ToLower(strings.TrimSpace(req.Language))
	if languageByTag(req.Language) == nil {
		return multiCheckRequest{}, errUnsupportedLanguage
	}
	if len(req.Paths) == 0 {
		return multiCheckRequest{}, errNoPaths
	}
	if req.Format == "" {
		req.Format = "console"
	}
	return req, nil
}

// handleCreateCheck submits a check task and returns 202 with its location.
func (s *Server) handleCreateCheck(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeMethodNotAllowed(w, req)
		return
	}
	body := readBody(req)
	parsed, err := decodeCheckRequest(body)
	if err != nil {
		writeBadRequest(w, req, err.Error())
		return
	}
	task, submitErr := s.submit(parsed, "check")
	if submitErr != nil {
		writeTooManyRequests(w, req)
		return
	}
	location := "/api/" + MetaAPIVersion + "/checks/" + task.ID
	w.Header().Set("Location", location)
	writeJSON(w, http.StatusAccepted, s.mustGet(task.ID))
}

// handleListChecks returns the check tasks, optionally filtered by status.
func (s *Server) handleListChecks(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		writeMethodNotAllowed(w, req)
		return
	}
	status := req.URL.Query().Get("status")
	writeJSON(w, http.StatusOK, listView{Tasks: s.store.list(status)})
}

// handleGetTask returns one task by identifier, or 404 when absent.
func (s *Server) handleGetTask(w http.ResponseWriter, req *http.Request) {
	id := routeParam(req.Context(), "id")
	task, ok := s.store.get(id)
	if !ok {
		writeNotFound(w, req)
		return
	}
	writeJSON(w, http.StatusOK, task)
}

// listView wraps a task listing in a deterministic envelope.
type listView struct {
	Tasks []Task `json:"tasks"`
}
