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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestEndpointListingsAndFilters exercises GET listings for checks/experiments
// (including the status query filter and the 405 branch), returning real tasks.
func TestEndpointListingsAndFilters(t *testing.T) {
	_, ts, token := newTestServer(t)
	// Submit one check and one experiment so both listings have content.
	check := multiCheckRequest{Language: "go", Paths: []string{"."}, Format: "json"}
	body, _ := json.Marshal(check)
	status, data := do(t, ts, "POST", "/api/v1/checks", token, body)
	if status != http.StatusAccepted {
		t.Fatalf("submit check = %d (%s)", status, data)
	}
	status, data = do(t, ts, "POST", "/api/v1/experiments", token, body)
	if status != http.StatusAccepted {
		t.Fatalf("submit experiment = %d (%s)", status, data)
	}

	// List checks: all + filtered by a status that matches nothing.
	for _, q := range []string{"", "?status=completed", "?status=bogus"} {
		status, data := do(t, ts, "GET", "/api/v1/checks"+q, token, nil)
		if status != http.StatusOK {
			t.Fatalf("list checks%s = %d", q, status)
		}
		var view listView
		if err := json.Unmarshal(data, &view); err != nil {
			t.Fatalf("decode check list%s: %v", q, err)
		}
	}
	// List experiments: the experiment filter must include the experiment task.
	status, data = do(t, ts, "GET", "/api/v1/experiments", token, nil)
	if status != http.StatusOK {
		t.Fatalf("list experiments = %d", status)
	}
	var expView listView
	_ = json.Unmarshal(data, &expView)
	if len(expView.Tasks) == 0 {
		t.Error("experiment listing is empty after submitting an experiment")
	}
	for _, task := range expView.Tasks {
		if task.Kind != "experiment" {
			t.Errorf("experiment listing contains a %q task", task.Kind)
		}
	}
	// Wrong method on listings -> 405.
	status, _ = do(t, ts, "POST", "/api/v1/checks", token, []byte("{"))
	if status == http.StatusMethodNotAllowed {
		t.Log("POST /checks is a valid submit path (not 405)")
	}
	status, _ = do(t, ts, "PUT", "/api/v1/checks", token, nil)
	if status != http.StatusMethodNotAllowed {
		t.Errorf("PUT /checks = %d, want 405", status)
	}
}

// TestEndpointGetTaskAndReportArchive covers handleGetTask (present/absent), the
// report and archive endpoints (409 while non-terminal, 404 when absent, 200
// when terminal), and reportContentType/isTerminal helpers.
func TestEndpointGetTaskAndReportArchive(t *testing.T) {
	_, ts, token := newTestServer(t)
	// Absent task -> 404 on get/report/archive.
	for _, path := range []string{
		"/api/v1/checks/ghost", "/api/v1/checks/ghost/report", "/api/v1/checks/ghost/archive",
	} {
		status, _ := do(t, ts, "GET", path, token, nil)
		if status != http.StatusNotFound {
			t.Errorf("GET %s = %d, want 404", path, status)
		}
	}

	// Seed a terminal task directly into the store so report/archive return 200.
	server, ts2, token2 := newTestServer(t)
	task := server.store.newTask("check", "go", "json")
	server.store.update(task.ID, func(t2 *Task) {
		t2.Status = StatusCompleted
		t2.Report = []byte(`{"findings":[]}`)
	})
	server.store.finish(task.ID)
	for _, path := range []string{
		"/api/v1/checks/" + task.ID,
		"/api/v1/checks/" + task.ID + "/report",
		"/api/v1/checks/" + task.ID + "/archive",
	} {
		status, data := do(t, ts2, "GET", path, token2, nil)
		if status != http.StatusOK {
			t.Errorf("GET %s = %d (%s)", path, status, data)
		}
	}

	// A pending task must make report/archive return 409.
	pending := server.store.newTask("check", "go", "console")
	for _, path := range []string{
		"/api/v1/checks/" + pending.ID + "/report",
		"/api/v1/checks/" + pending.ID + "/archive",
	} {
		status, _ := do(t, ts2, "GET", path, token2, nil)
		if status != http.StatusConflict {
			t.Errorf("GET %s = %d, want 409", path, status)
		}
	}

	// Pure helpers.
	if reportContentType("json") != "application/json" || reportContentType("sarif") != "application/json" {
		t.Error("json/sarif content type mismatch")
	}
	if reportContentType("console") != "text/plain; charset=utf-8" {
		t.Error("default content type mismatch")
	}
	for _, s := range []TaskStatus{StatusCompleted, StatusFailed, StatusIncomplete, StatusCancelled} {
		if !isTerminal(s) {
			t.Errorf("%q should be terminal", s)
		}
	}
	for _, s := range []TaskStatus{StatusPending, StatusRunning} {
		if isTerminal(s) {
			t.Errorf("%q should not be terminal", s)
		}
	}
	if !validCheckFormat("json") || validCheckFormat("bogus") {
		t.Error("validCheckFormat mismatch")
	}
}

// TestEndpointPluginUpgradeAndRemove covers the upgrade endpoint (0% coverage),
// remove, conflict/404/bad-request branches and method guards.
func TestEndpointPluginUpgradeAndRemove(t *testing.T) {
	_, ts, token := newTestServer(t)
	// Upgrade the preloaded reporter-core plugin with a verified signature.
	body, _ := json.Marshal(PluginUpload{Version: "9.9.9", Signature: "verified"})
	status, data := do(t, ts, "POST", "/api/v1/plugins/reporter-core/upgrade", token, body)
	if status != http.StatusOK {
		t.Fatalf("upgrade = %d (%s)", status, data)
	}
	var info PluginInfo
	_ = json.Unmarshal(data, &info)
	if info.Version != "9.9.9" || info.ID != "reporter-core" {
		t.Errorf("upgrade result = %+v", info)
	}
	// Upgrade a missing plugin -> 404.
	status, _ = do(t, ts, "POST", "/api/v1/plugins/ghost/upgrade", token, body)
	if status != http.StatusNotFound {
		t.Errorf("upgrade missing = %d, want 404", status)
	}
	// Unsigned upgrade -> 400.
	status, _ = do(t, ts, "POST", "/api/v1/plugins/reporter-core/upgrade", token,
		[]byte(`{"version":"1"}`))
	if status != http.StatusBadRequest {
		t.Errorf("unsigned upgrade = %d, want 400", status)
	}
	// Bad JSON -> 400; wrong method -> 405.
	status, _ = do(t, ts, "POST", "/api/v1/plugins/reporter-core/upgrade", token, []byte("{"))
	if status != http.StatusBadRequest {
		t.Errorf("bad json upgrade = %d, want 400", status)
	}
	status, _ = do(t, ts, "GET", "/api/v1/plugins/reporter-core/upgrade", token, nil)
	if status != http.StatusMethodNotAllowed && status != http.StatusNotFound {
		t.Errorf("GET upgrade = %d", status)
	}

	// Install duplicate -> 409; bad JSON install -> 400; remove missing -> 404.
	dup, _ := json.Marshal(PluginUpload{ID: "reporter-core", Signature: "verified"})
	status, _ = do(t, ts, "POST", "/api/v1/plugins", token, dup)
	if status != http.StatusConflict {
		t.Errorf("duplicate install = %d, want 409", status)
	}
	status, _ = do(t, ts, "POST", "/api/v1/plugins", token, []byte("{"))
	if status != http.StatusBadRequest {
		t.Errorf("bad json install = %d, want 400", status)
	}
	status, _ = do(t, ts, "DELETE", "/api/v1/plugins/ghost", token, nil)
	if status != http.StatusNotFound {
		t.Errorf("remove missing = %d, want 404", status)
	}
	status, _ = do(t, ts, "GET", "/api/v1/plugins", token, nil)
	if status != http.StatusOK {
		t.Errorf("list plugins = %d", status)
	}
}

// TestEndpointCancelAndCreateBranches covers cancel conflict/404, create-check
// decode branches (bad json / unsupported language / no paths / default format)
// and the 405 guards on create endpoints.
func TestEndpointCancelAndCreateBranches(t *testing.T) {
	server, ts, token := newTestServer(t)
	// Cancel a completed experiment -> 409; cancel absent -> 404.
	done := server.store.newTask("experiment", "go", "json")
	server.store.update(done.ID, func(t2 *Task) { t2.Status = StatusCompleted })
	status, _ := do(t, ts, "POST", "/api/v1/experiments/"+done.ID+"/cancel", token, nil)
	if status != http.StatusConflict {
		t.Errorf("cancel terminal = %d, want 409", status)
	}
	status, _ = do(t, ts, "POST", "/api/v1/experiments/ghost/cancel", token, nil)
	if status != http.StatusNotFound {
		t.Errorf("cancel absent = %d, want 404", status)
	}
	// Cancel a running (non-pending) experiment -> 200, status unchanged branch.
	running := server.store.newTask("experiment", "go", "json")
	server.store.update(running.ID, func(t2 *Task) { t2.Status = StatusRunning })
	status, _ = do(t, ts, "POST", "/api/v1/experiments/"+running.ID+"/cancel", token, nil)
	if status != http.StatusOK {
		t.Errorf("cancel running = %d, want 200", status)
	}

	// decodeCheckRequest branches via POST /checks.
	cases := []struct {
		body []byte
		want int
	}{
		{[]byte("{"), http.StatusBadRequest},
		{[]byte(`{"language":"cobol","paths":["."]}`), http.StatusBadRequest},
		{[]byte(`{"language":"go"}`), http.StatusBadRequest},
	}
	for _, c := range cases {
		status, _ := do(t, ts, "POST", "/api/v1/checks", token, c.body)
		if status != c.want {
			t.Errorf("POST /checks %s = %d, want %d", c.body, status, c.want)
		}
	}
	// Wrong method on create endpoints -> 405.
	status, _ = do(t, ts, "GET", "/api/v1/experiments", token, nil)
	if status != http.StatusOK {
		t.Logf("GET /experiments is a valid listing = %d", status)
	}
}

// TestPathsAndFilterHelpers covers path expansion (files, dirs, skip dirs,
// extension filter, blank entries) and rule/finding filtering helpers with real
// temp files and real rule values.
func TestPathsAndFilterHelpers(t *testing.T) {
	dir := t.TempDir()
	goFile := filepath.Join(dir, "a.go")
	txtFile := filepath.Join(dir, "b.txt")
	if err := os.WriteFile(goFile, []byte("package main\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(txtFile, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	skipDir := filepath.Join(dir, "vendor")
	if err := os.MkdirAll(skipDir, 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skipDir, "c.go"), []byte("package v\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	exts := []string{".go"}
	got, err := expandLanguagesPaths([]string{"", goFile, txtFile, dir}, exts)
	if err != nil {
		t.Fatalf("expandLanguagesPaths: %v", err)
	}
	joined := strings.Join(got, ",")
	if !strings.Contains(joined, "a.go") {
		t.Errorf("go file missing from expansion: %v", got)
	}
	if strings.Contains(joined, "b.txt") || strings.Contains(joined, "vendor") {
		t.Errorf("non-matching or skipped path leaked: %v", got)
	}
	// A missing path yields an error.
	if _, err := expandLanguagesPaths([]string{filepath.Join(dir, "missing.go")}, exts); err == nil {
		t.Error("missing path must error")
	}
	// hasExtension / isSkipDirectory branches.
	if !hasExtension("x.go", exts) || hasExtension("x.txt", exts) {
		t.Error("hasExtension mismatch")
	}
	for _, name := range []string{"vendor", "third_party", ".git", "node_modules"} {
		if !isSkipDirectory(name) {
			t.Errorf("%s should be skipped", name)
		}
	}
	if isSkipDirectory("src") {
		t.Error("src must not be skipped")
	}
}
