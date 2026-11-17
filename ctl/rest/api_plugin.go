/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"encoding/json"
	"errors"
	"net/http"
)

// handleListPlugins returns the installed plugin inventory.
func (s *Server) handleListPlugins(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		writeMethodNotAllowed(w, req)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"plugins": s.plugins.List(),
	})
}

// handleInstallPlugin registers a plugin, rejecting an unsigned upload to honor
// the trust-anchor invariant (PLAN-07 section 6).
func (s *Server) handleInstallPlugin(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeMethodNotAllowed(w, req)
		return
	}
	var upload PluginUpload
	if err := json.Unmarshal(readBody(req), &upload); err != nil {
		writeBadRequest(w, req, errBadJSON.Error())
		return
	}
	if upload.Signature == "" {
		writeBadRequest(w, req, "plugin signature is required")
		return
	}
	info, err := s.plugins.Install(upload)
	if err != nil {
		if errors.Is(err, errPluginExists) {
			writeConflict(w, req, err.Error())
			return
		}
		writeBadRequest(w, req, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, info)
}

// handleRemovePlugin deletes a plugin by identifier.
func (s *Server) handleRemovePlugin(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		writeMethodNotAllowed(w, req)
		return
	}
	id := routeParam(req.Context(), "id")
	if err := s.plugins.Remove(id); err != nil {
		writeNotFound(w, req)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// handleUpgradePlugin replaces a plugin's metadata under a verified signature.
func (s *Server) handleUpgradePlugin(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeMethodNotAllowed(w, req)
		return
	}
	id := routeParam(req.Context(), "id")
	var upload PluginUpload
	if err := json.Unmarshal(readBody(req), &upload); err != nil {
		writeBadRequest(w, req, errBadJSON.Error())
		return
	}
	upload.ID = id
	if upload.Signature == "" {
		writeBadRequest(w, req, "plugin signature is required")
		return
	}
	info, err := s.plugins.Upgrade(upload)
	if err != nil {
		if errors.Is(err, errPluginMissing) {
			writeNotFound(w, req)
			return
		}
		writeBadRequest(w, req, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, info)
}
