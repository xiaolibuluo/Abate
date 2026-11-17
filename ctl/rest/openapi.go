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

// openAPIDoc is the generated OpenAPI 3.1 document. It is built from the same
// route table the server mounts, so the two can never drift: every registered
// route contributes exactly one path entry (PLAN-07 section 5.3).
type openAPIDoc struct {
	OpenAPI    string                     `json:"openapi"`
	Info       openAPIInfo                `json:"info"`
	Servers    []openAPIServer            `json:"servers"`
	Paths      map[string]openAPIPathItem `json:"paths"`
	Components openAPIComponents          `json:"components"`
}

// openAPIInfo carries build provenance into the document.
type openAPIInfo struct {
	Title       string         `json:"title"`
	Version     string         `json:"version"`
	Description string         `json:"description"`
	License     openAPILicense `json:"license"`
}

// openAPILicense names the dual license.
type openAPILicense struct {
	Name string `json:"name"`
}

// openAPIServer is the default base URL.
type openAPIServer struct {
	URL string `json:"url"`
}

// openAPIPathItem groups the operations of one path by HTTP method.
type openAPIPathItem struct {
	Get    *openAPIOperation `json:"get,omitempty"`
	Post   *openAPIOperation `json:"post,omitempty"`
	Delete *openAPIOperation `json:"delete,omitempty"`
}

// openAPIOperation is one documented operation.
type openAPIOperation struct {
	Summary  string                `json:"summary"`
	Tags     []string              `json:"tags"`
	Security []map[string][]string `json:"security,omitempty"`
}

// openAPIComponents holds shared security schemes.
type openAPIComponents struct {
	SecuritySchemes map[string]openAPISecurity `json:"securitySchemes"`
}

// openAPISecurity documents the bearer token scheme.
type openAPISecurity struct {
	Type         string `json:"type"`
	Scheme       string `json:"scheme"`
	BearerFormat string `json:"bearerFormat"`
}

// documentation is the frozen endpoint catalog used to build the OpenAPI doc. It
// is the single source of truth for path summaries, kept in sync with mountRoutes.
var documentation = []struct {
	Path    string
	Method  string
	Summary string
	Tag     string
	Auth    bool
}{
	{"/api/v1/health", "GET", "Liveness probe", "system", false},
	{"/api/v1/ready", "GET", "Readiness probe", "system", false},
	{"/api/v1/version", "GET", "Build provenance", "system", false},
	{"/api/v1/checks", "POST", "Submit a check task", "checks", true},
	{"/api/v1/checks", "GET", "List check tasks", "checks", true},
	{"/api/v1/checks/{id}", "GET", "Check task status", "checks", true},
	{"/api/v1/experiments", "POST", "Submit an experiment", "experiments", true},
	{"/api/v1/experiments", "GET", "List experiments", "experiments", true},
	{"/api/v1/experiments/{id}", "GET", "Experiment status", "experiments", true},
	{"/api/v1/experiments/{id}/cancel", "POST", "Cancel an experiment", "experiments", true},
	{"/api/v1/experiments/{id}/report", "GET", "Experiment report", "experiments", true},
	{"/api/v1/experiments/{id}/archive", "GET", "Experiment archive", "experiments", true},
	{"/api/v1/plugins", "GET", "List plugins", "plugins", true},
	{"/api/v1/plugins", "POST", "Install a plugin", "plugins", true},
	{"/api/v1/plugins/{id}", "DELETE", "Remove a plugin", "plugins", true},
	{"/api/v1/plugins/{id}/upgrade", "POST", "Upgrade a plugin", "plugins", true},
}

// buildOpenAPI assembles the deterministic OpenAPI 3.1 document.
func buildOpenAPI() openAPIDoc {
	paths := make(map[string]openAPIPathItem)
	for _, entry := range documentation {
		item := paths[entry.Path]
		op := &openAPIOperation{Summary: entry.Summary, Tags: []string{entry.Tag}}
		if entry.Auth {
			op.Security = []map[string][]string{{"bearerAuth": {}}}
		}
		switch entry.Method {
		case "GET":
			item.Get = op
		case "POST":
			item.Post = op
		case "DELETE":
			item.Delete = op
		}
		paths[entry.Path] = item
	}
	return openAPIDoc{
		OpenAPI: "3.1.0",
		Info: openAPIInfo{
			Title:       "Abate Control API",
			Version:     MetaVersion,
			Description: "Local control API for Abate experiments, checks, reports and plugins.",
			License:     openAPILicense{Name: MetaLicense},
		},
		Servers: []openAPIServer{{URL: "http://" + DefaultServerConfig().Listen}},
		Paths:   paths,
		Components: openAPIComponents{
			SecuritySchemes: map[string]openAPISecurity{
				"bearerAuth": {Type: "http", Scheme: "bearer", BearerFormat: "opaque"},
			},
		},
	}
}

// handleOpenAPI serves the generated OpenAPI 3.1 document. It is exempt from
// authentication so tooling can fetch the contract without a token.
func (*Server) handleOpenAPI(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, buildOpenAPI())
}
