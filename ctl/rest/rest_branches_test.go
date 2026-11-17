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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestMiddlewareStringHelpers covers the dependency-free byte/string helpers
// that back client-IP parsing — every branch is a real, observable behavior.
func TestMiddlewareStringHelpers(t *testing.T) {
	if indexByte("abc,bcd", ',') != 3 || indexByte("abc", ',') != -1 {
		t.Error("indexByte mismatch")
	}
	if lastIndexByte("a:b:c", ':') != 3 || lastIndexByte("abc", ':') != -1 {
		t.Error("lastIndexByte mismatch")
	}
	for _, s := range []string{"  x  ", "\tx\n", " \r\n ", "x"} {
		want := strings.TrimSpace(s)
		if got := trimSpace(s); got != want {
			t.Errorf("trimSpace(%q) = %q, want %q", s, got, want)
		}
	}
	for _, b := range []byte{' ', '\t', '\n', '\r', '\f', '\v'} {
		if !isSpace(b) {
			t.Errorf("isSpace(%q) = false", b)
		}
	}
	if isSpace('x') || isSpace('0') {
		t.Error("isSpace accepted a non-whitespace byte")
	}
	if maxInt(3, 7) != 7 || maxInt(7, 3) != 7 || maxInt(5, 5) != 5 {
		t.Error("maxInt mismatch")
	}
}

// TestClientIPBranches proves forwarded-header and remote-address parsing,
// including the single-address preference and the multi-address split.
func TestClientIPBranches(t *testing.T) {
	mk := func(fwd, remote string) *http.Request {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		if fwd != "" {
			req.Header.Set("X-Forwarded-For", fwd)
		}
		req.RemoteAddr = remote
		return req
	}
	cases := []struct{ fwd, remote, want string }{
		{" 1.2.3.4 ", "9.9.9.9:1", "1.2.3.4"},
		{"1.2.3.4, 5.6.7.8", "9.9.9.9:1", "1.2.3.4"},
		{"", "10.0.0.1:4242", "10.0.0.1"},
		{"", "no-colon", "no-colon"},
	}
	for _, c := range cases {
		if got := clientIP(mk(c.fwd, c.remote)); got != c.want {
			t.Errorf("clientIP(fwd=%q, remote=%q) = %q, want %q", c.fwd, c.remote, got, c.want)
		}
	}
}

// TestIPLimiterBranches proves the deterministic token bucket: capacity default,
// per-client isolation, exhaustion, and the global refill after resetStep.
func TestIPLimiterBranches(t *testing.T) {
	lim := newIPLimiter(0) // exercises the <=0 default branch
	if lim.cap != 100 {
		t.Errorf("default capacity = %d, want 100", lim.cap)
	}
	small := newIPLimiter(2)
	if small.resetStep != maxInt(1024, 16) {
		t.Errorf("resetStep = %d", small.resetStep)
	}
	for i := 0; i < 2; i++ {
		if !small.allow("a") {
			t.Fatalf("request %d for client a rejected before exhaustion", i)
		}
	}
	if small.allow("a") {
		t.Error("exhausted client a was not rejected")
	}
	if !small.allow("b") {
		t.Error("a fresh client b must have its own bucket")
	}
	// Force a global refill by driving served up to a resetStep boundary.
	refill := newIPLimiter(1)
	_ = refill.allow("c")
	for refill.served < refill.resetStep-1 {
		refill.buckets["c"] = 0
		refill.served++
	}
	refill.buckets["c"] = 0
	if !refill.allow("c") {
		t.Error("global refill did not restore the token bucket")
	}
}

// TestMiddlewareChains exercises recover/log/rate-limit/auth wrappers through
// real handlers, covering panic recovery, status recording and 429/401 paths.
func TestMiddlewareChains(t *testing.T) {
	// recoverMiddleware converts a panic into a 500 problem document.
	rec := httptest.NewRecorder()
	recoverMiddleware(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		panic("boom")
	})).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("recovered panic status = %d", rec.Code)
	}

	// logMiddleware records status and byte count without a wall clock.
	logRec := httptest.NewRecorder()
	logMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("hi"))
	})).ServeHTTP(logRec, httptest.NewRequest(http.MethodGet, "/", nil))
	if logRec.Code != http.StatusCreated {
		t.Errorf("log middleware status = %d", logRec.Code)
	}

	// rateLimitMiddleware rejects an over-quota client with 429.
	lim := newIPLimiter(1)
	limited := rateLimitMiddleware(lim, http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	r1 := httptest.NewRecorder()
	limited.ServeHTTP(r1, httptest.NewRequest(http.MethodGet, "/", nil))
	r2 := httptest.NewRecorder()
	limited.ServeHTTP(r2, httptest.NewRequest(http.MethodGet, "/", nil))
	if r1.Code != http.StatusOK || r2.Code != http.StatusTooManyRequests {
		t.Errorf("rate limit codes = %d, %d", r1.Code, r2.Code)
	}
	if r2.Header().Get("Retry-After") == "" {
		t.Error("429 response missing Retry-After")
	}

	// authMiddleware exempts probes, rejects bad tokens, accepts good ones.
	authed := authMiddleware("sek", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	probe := httptest.NewRecorder()
	authed.ServeHTTP(probe, httptest.NewRequest(http.MethodGet, "/api/v1/health", nil))
	denied := httptest.NewRecorder()
	authed.ServeHTTP(denied, httptest.NewRequest(http.MethodGet, "/api/v1/checks", nil))
	allowed := httptest.NewRequest(http.MethodGet, "/api/v1/checks", nil)
	allowed.Header.Set("Authorization", "Bearer sek")
	okRec := httptest.NewRecorder()
	authed.ServeHTTP(okRec, allowed)
	if probe.Code != http.StatusOK || denied.Code != http.StatusUnauthorized || okRec.Code != http.StatusOK {
		t.Errorf("auth codes = probe:%d denied:%d allowed:%d", probe.Code, denied.Code, okRec.Code)
	}
}

// TestStatusRecorderBranches proves Write/WriteHeader only record the first
// status and accumulate byte counts.
func TestStatusRecorderBranches(t *testing.T) {
	base := httptest.NewRecorder()
	sr := &statusRecorder{ResponseWriter: base}
	sr.WriteHeader(http.StatusAccepted)
	sr.WriteHeader(http.StatusTeapot) // must not overwrite the first status
	if sr.status != http.StatusAccepted {
		t.Errorf("status = %d, want 202", sr.status)
	}
	n, _ := sr.Write([]byte("abc"))
	if n != 3 || sr.bytes != 3 {
		t.Errorf("write count n=%d bytes=%d", n, sr.bytes)
	}
	// A Write before any WriteHeader defaults the status to 200.
	sr2 := &statusRecorder{ResponseWriter: httptest.NewRecorder()}
	_, _ = sr2.Write([]byte("x"))
	if sr2.status != http.StatusOK {
		t.Errorf("implicit status = %d, want 200", sr2.status)
	}
}

// TestProblemWriters exercises every problem+json writer and routeParam,
// asserting status codes, content type and body shape.
func TestProblemWriters(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/x", nil)
	writers := []struct {
		name   string
		status int
		call   func(http.ResponseWriter, *http.Request)
	}{
		{"bad request", http.StatusBadRequest, func(w http.ResponseWriter, r *http.Request) { writeBadRequest(w, r, "d") }},
		{"not found", http.StatusNotFound, writeNotFound},
		{"method not allowed", http.StatusMethodNotAllowed, writeMethodNotAllowed},
		{"conflict", http.StatusConflict, func(w http.ResponseWriter, r *http.Request) { writeConflict(w, r, "d") }},
		{"unauthorized", http.StatusUnauthorized, writeUnauthorized},
		{"too many", http.StatusTooManyRequests, writeTooManyRequests},
		{"internal", http.StatusInternalServerError, func(w http.ResponseWriter, r *http.Request) { writeInternalError(w, r, "d") }},
	}
	for _, w := range writers {
		rec := httptest.NewRecorder()
		w.call(rec, req)
		if rec.Code != w.status {
			t.Errorf("%s: status = %d, want %d", w.name, rec.Code, w.status)
		}
		if ct := rec.Header().Get("Content-Type"); ct != "application/problem+json" {
			t.Errorf("%s: content-type = %q", w.name, ct)
		}
		var p Problem
		if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
			t.Errorf("%s: invalid problem json: %v", w.name, err)
		}
		if p.Status != w.status {
			t.Errorf("%s: body status = %d", w.name, p.Status)
		}
	}

	// routeParam and withParams round-trip; missing/nil contexts yield "".
	ctx := withParams(req.Context(), map[string]string{"id": "t1"})
	if got := routeParam(ctx, "id"); got != "t1" {
		t.Errorf("routeParam = %q, want t1", got)
	}
	if got := routeParam(req.Context(), "id"); got != "" {
		t.Errorf("routeParam on bare context = %q, want empty", got)
	}
}

// TestWriteJSONAndReadBody cover the JSON response helper (including its error
// branch with an unmarshalable value) and the bounded body reader.
func TestWriteJSONAndReadBody(t *testing.T) {
	rec := httptest.NewRecorder()
	writeJSON(rec, http.StatusOK, map[string]string{"k": "v"})
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"k":"v"`) {
		t.Errorf("writeJSON = %d %q", rec.Code, rec.Body.String())
	}
	// A channel cannot be marshaled -> writeInternalError path (req is nil-safe).
	errRec := httptest.NewRecorder()
	writeJSON(errRec, http.StatusOK, make(chan int))
	if errRec.Code != http.StatusInternalServerError {
		t.Errorf("unmarshalable value status = %d", errRec.Code)
	}

	// readBody: nil body, normal body, and a body exceeding the cap.
	if got := readBody(httptest.NewRequest(http.MethodPost, "/", nil)); len(got) != 0 {
		t.Errorf("empty body should read as no bytes, got %v", got)
	}
	full := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("payload")))
	if string(readBody(full)) != "payload" {
		t.Error("readBody did not return the payload")
	}
	big := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(make([]byte, maxRequestBody+10)))
	if len(readBody(big)) != maxRequestBody {
		t.Errorf("readBody did not cap at %d", maxRequestBody)
	}
}
