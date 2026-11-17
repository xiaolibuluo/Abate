/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"abate/security/redact"
)

// captureLog redirects the process-wide logger to buf and restores it through
// t.Cleanup. The test never calls t.Parallel because log state is global.
func captureLog(t *testing.T, buf *bytes.Buffer) {
	t.Helper()
	previous := log.Writer()
	log.SetOutput(buf)
	t.Cleanup(func() { log.SetOutput(previous) })
}

// TestRecoverMiddlewareRedactsPanicValue proves T9 at the panic log point: a
// recovered panic value carrying a Bearer credential must never be printed raw.
func TestRecoverMiddlewareRedactsPanicValue(t *testing.T) {
	var buf bytes.Buffer
	captureLog(t, &buf)

	const canary = "CANARY-PANIC-9f8e7d6c"
	handler := recoverMiddleware(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		panic("Bearer " + canary)
	}))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/experiments", nil)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("panic recovery status = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
	got := buf.String()
	if strings.Contains(got, canary) {
		t.Fatalf("raw credential leaked into panic log: %q", got)
	}
	if !strings.Contains(got, redact.Placeholder) {
		t.Fatalf("panic log is missing the redaction marker: %q", got)
	}
}

// TestLogMiddlewareRedactsURLPath proves T9 at the access log point: a secret
// embedded in the request path is redacted while the route structure survives.
func TestLogMiddlewareRedactsURLPath(t *testing.T) {
	var buf bytes.Buffer
	captureLog(t, &buf)

	const canary = "CANARY-PATH-9f8e7d6c"
	path := "/v1/experiments/token=" + canary + "/detail"
	handler := logMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	handler.ServeHTTP(rec, req)

	got := buf.String()
	if strings.Contains(got, canary) {
		t.Fatalf("raw credential leaked into access log: %q", got)
	}
	if !strings.Contains(got, redact.Placeholder) {
		t.Fatalf("access log is missing the redaction marker: %q", got)
	}
	if !strings.Contains(got, "/detail") {
		t.Fatalf("route structure was destroyed by redaction: %q", got)
	}
}

// TestLogMiddlewareNeverEmitsRawQuery proves that a credential carried in the
// query string cannot reach the access log, which only prints the redacted path.
func TestLogMiddlewareNeverEmitsRawQuery(t *testing.T) {
	var buf bytes.Buffer
	captureLog(t, &buf)

	const canary = "CANARY-QUERY-9f8e7d6c"
	target := "/v1/token=" + canary + "?secret=" + canary
	handler := logMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	handler.ServeHTTP(rec, req)

	got := buf.String()
	if strings.Contains(got, canary) {
		t.Fatalf("raw query credential leaked into access log: %q", got)
	}
	if !strings.Contains(got, redact.Placeholder) {
		t.Fatalf("access log is missing the redaction marker: %q", got)
	}
}
