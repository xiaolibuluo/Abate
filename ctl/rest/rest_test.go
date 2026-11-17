/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// newTestServer builds a server with a fixed token for the test harness.
func newTestServer(t *testing.T) (*Server, *httptest.Server, string) {
	t.Helper()
	cfg := DefaultServerConfig()
	cfg.Auth = "test-token"
	server, err := New(cfg, ServerOptions{})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	ts := httptest.NewServer(server.Handler())
	t.Cleanup(ts.Close)
	return server, ts, "test-token"
}

// do issues one request and returns the status code plus body. The response body
// is fully consumed and closed inside this helper, so callers never hold an open
// *http.Response — this keeps bodyclose satisfied without suppression and makes
// the harness surface exactly what assertions need.
func do(t *testing.T, ts *httptest.Server, method, path, token string,
	body []byte) (int, []byte) {
	t.Helper()
	req, err := http.NewRequest(method, ts.URL+path, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	data, _ := io.ReadAll(resp.Body)
	status := resp.StatusCode
	_ = resp.Body.Close()
	return status, data
}

// TestHealthReadyVersionNoAuth proves the system probes need no token.
func TestHealthReadyVersionNoAuth(t *testing.T) {
	t.Parallel()
	_, ts, _ := newTestServer(t)
	for _, path := range []string{"/api/v1/health", "/api/v1/ready", "/api/v1/version"} {
		status, data := do(t, ts, "GET", path, "", nil)
		if status != http.StatusOK {
			t.Fatalf("%s = %d, want 200 (%s)", path, status, data)
		}
	}
}

// TestOpenAPIUnauthenticated proves the contract endpoint is public.
func TestOpenAPIUnauthenticated(t *testing.T) {
	t.Parallel()
	_, ts, _ := newTestServer(t)
	status, data := do(t, ts, "GET", "/openapi.json", "", nil)
	if status != http.StatusOK {
		t.Fatalf("openapi = %d, want 200", status)
	}
	var doc openAPIDoc
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("decode openapi: %v", err)
	}
	if doc.OpenAPI != "3.1.0" {
		t.Fatalf("openapi version = %q", doc.OpenAPI)
	}
	// Distinct paths collapse in the document; count operations instead to prove
	// every documented route is present without drift.
	wantOps := len(documentation)
	gotOps := 0
	for _, item := range doc.Paths {
		if item.Get != nil {
			gotOps++
		}
		if item.Post != nil {
			gotOps++
		}
		if item.Delete != nil {
			gotOps++
		}
	}
	if gotOps != wantOps {
		t.Fatalf("openapi operations = %d, want %d", gotOps, wantOps)
	}
}

// TestAuthRequired proves every protected endpoint rejects a missing token.
func TestAuthRequired(t *testing.T) {
	t.Parallel()
	_, ts, _ := newTestServer(t)
	protected := []string{
		"/api/v1/checks", "/api/v1/experiments", "/api/v1/plugins",
	}
	for _, path := range protected {
		status, _ := do(t, ts, "GET", path, "", nil)
		if status != http.StatusUnauthorized {
			t.Fatalf("%s = %d, want 401", path, status)
		}
	}
}

// TestAuthWrongTokenRejected proves a bad bearer token is refused.
func TestAuthWrongTokenRejected(t *testing.T) {
	t.Parallel()
	_, ts, _ := newTestServer(t)
	status, _ := do(t, ts, "GET", "/api/v1/checks", "wrong", nil)
	if status != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", status)
	}
}

// TestCreateCheckValidatesLanguage proves an unsupported language is 400.
func TestCreateCheckValidatesLanguage(t *testing.T) {
	t.Parallel()
	_, ts, token := newTestServer(t)
	body, _ := json.Marshal(multiCheckRequest{Language: "cobol", Paths: []string{"."}})
	status, _ := do(t, ts, "POST", "/api/v1/checks", token, body)
	if status != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", status)
	}
}

// TestCheckEndToEnd runs a real check on a temp Go file and fetches the report.
func TestCheckEndToEnd(t *testing.T) {
	t.Parallel()
	_, ts, token := newTestServer(t)
	dir := t.TempDir()
	src := filepath.Join(dir, "sample.go")
	if err := os.WriteFile(src, []byte("package main\n\nvar bad_name = 1\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	body, _ := json.Marshal(multiCheckRequest{
		Language: "go", Paths: []string{src}, Format: "json",
	})
	status, data := do(t, ts, "POST", "/api/v1/checks", token, body)
	if status != http.StatusAccepted {
		t.Fatalf("submit = %d, want 202 (%s)", status, data)
	}
	var task Task
	if err := json.Unmarshal(data, &task); err != nil {
		t.Fatal(err)
	}
	if task.Status != StatusPending && task.Status != StatusRunning && task.Status != StatusCompleted {
		t.Fatalf("initial status = %q", task.Status)
	}
	// Poll until terminal without sleeping: the in-process runner completes a
	// tiny check in microseconds, so a bounded spin is both fast and
	// deterministic (no wall-clock dependency the linter forbids).
	var deadlineHit bool
	for i := 0; i < 2000; i++ {
		status, data = do(t, ts, "GET", "/api/v1/checks/"+task.ID, token, nil)
		if status != http.StatusOK {
			t.Fatalf("get = %d (%s)", status, data)
		}
		_ = json.Unmarshal(data, &task)
		if isTerminal(task.Status) {
			deadlineHit = false
			break
		}
		deadlineHit = true
	}
	if deadlineHit {
		t.Fatalf("task %s never terminated, last status %q", task.ID, task.Status)
	}
	if task.Status != StatusCompleted {
		t.Fatalf("final status = %q, detail %q", task.Status, task.Detail)
	}
	if task.Summary.Total < 1 {
		t.Fatalf("expected at least one finding, got %d", task.Summary.Total)
	}
	// Fetch the report.
	status, data = do(t, ts, "GET", "/api/v1/checks/"+task.ID+"/report", token, nil)
	if status != http.StatusOK {
		t.Fatalf("report = %d", status)
	}
	if len(data) == 0 {
		t.Fatal("report body is empty")
	}
}

// TestExperimentCancel proves cancel moves a pending task to a terminal state.
func TestExperimentCancel(t *testing.T) {
	t.Parallel()
	_, ts, token := newTestServer(t)
	body, _ := json.Marshal(multiCheckRequest{Language: "go", Paths: []string{"."}, Format: "console"})
	status, data := do(t, ts, "POST", "/api/v1/experiments", token, body)
	if status != http.StatusAccepted {
		t.Fatalf("submit = %d (%s)", status, data)
	}
	var task Task
	_ = json.Unmarshal(data, &task)
	status, _ = do(t, ts, "POST", "/api/v1/experiments/"+task.ID+"/cancel", token, nil)
	if status != http.StatusOK {
		t.Fatalf("cancel = %d", status)
	}
	_, data = do(t, ts, "GET", "/api/v1/experiments/"+task.ID, token, nil)
	_ = json.Unmarshal(data, &task)
	if !task.Cancelled {
		t.Fatalf("task %s not marked cancelled", task.ID)
	}
}

// TestPluginInstallRejectsUnsigned proves the trust-anchor invariant holds.
func TestPluginInstallRejectsUnsigned(t *testing.T) {
	t.Parallel()
	_, ts, token := newTestServer(t)
	body, _ := json.Marshal(PluginUpload{ID: "my-plugin", Type: "reporter", Version: "1.0.0"})
	status, _ := do(t, ts, "POST", "/api/v1/plugins", token, body)
	if status != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400 for unsigned", status)
	}
}

// TestPluginLifecycle proves install (signed), list and remove all work.
func TestPluginLifecycle(t *testing.T) {
	t.Parallel()
	_, ts, token := newTestServer(t)
	upload := PluginUpload{ID: "signed", Type: "reporter", Version: "2.0.0", Signature: "verified"}
	body, _ := json.Marshal(upload)
	status, data := do(t, ts, "POST", "/api/v1/plugins", token, body)
	if status != http.StatusCreated {
		t.Fatalf("install = %d (%s)", status, data)
	}
	status, data = do(t, ts, "GET", "/api/v1/plugins", token, nil)
	if status != http.StatusOK {
		t.Fatalf("list = %d", status)
	}
	var listing struct {
		Plugins []PluginInfo `json:"plugins"`
	}
	_ = json.Unmarshal(data, &listing)
	found := false
	for _, p := range listing.Plugins {
		if p.ID == "signed" {
			found = true
		}
	}
	if !found {
		t.Fatal("installed plugin not present in listing")
	}
	status, _ = do(t, ts, "DELETE", "/api/v1/plugins/signed", token, nil)
	if status != http.StatusNoContent {
		t.Fatalf("remove = %d", status)
	}
}

// TestUnknownRouteReturns404 proves an unregistered path yields problem+json 404.
func TestUnknownRouteReturns404(t *testing.T) {
	t.Parallel()
	_, ts, token := newTestServer(t)
	status, _ := do(t, ts, "GET", "/api/v1/nope", token, nil)
	if status != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", status)
	}
}

// TestConfigRejectsRemoteWithoutAuth proves the auth invariant at startup.
func TestConfigRejectsRemoteWithoutAuth(t *testing.T) {
	t.Parallel()
	cfg := DefaultServerConfig()
	cfg.Listen = "0.0.0.0:8420"
	if _, err := New(cfg, ServerOptions{}); err == nil {
		t.Fatal("New accepted a remote bind without auth")
	}
}
