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

// versionView is the GET /version payload.
type versionView struct {
	Version   string `json:"version"`
	Schema    string `json:"schema"`
	Build     string `json:"build"`
	License   string `json:"license"`
	Source    string `json:"source"`
	Watermark string `json:"watermark"`
	API       string `json:"api"`
}

// healthView is the liveness probe payload.
type healthView struct {
	Status string `json:"status"`
}

// readyView is the readiness probe payload.
type readyView struct {
	Status     string   `json:"status"`
	Components []string `json:"components"`
}

// handleHealth reports liveness: the process can answer requests.
func (*Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, healthView{Status: "ok"})
}

// handleReady reports readiness: the control plane and its ports are initialized.
func (*Server) handleReady(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, readyView{
		Status:     "ok",
		Components: []string{"checker", "store", "runner", "plugins"},
	})
}

// handleVersion reports build provenance. It is exempt from authentication so
// operators can discover the build before presenting a token.
func (*Server) handleVersion(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, versionView{
		Version:   MetaVersion,
		Schema:    MetaSchema,
		Build:     MetaBuild,
		License:   MetaLicense,
		Source:    MetaSource,
		Watermark: MetaWatermark,
		API:       MetaAPIVersion,
	})
}
