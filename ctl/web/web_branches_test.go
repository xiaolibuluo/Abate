/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package web

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestBuildAPIHandlerBranches covers the embedded-REST path (empty RestBase),
// the reverse-proxy path (non-empty RestBase) and the invalid-URL error branch.
func TestBuildAPIHandlerBranches(t *testing.T) {
	// Embedded REST server: loopback needs no token.
	embedded, err := buildAPIHandler(Config{Listen: "127.0.0.1:8430"})
	if err != nil || embedded == nil {
		t.Fatalf("embedded handler: err=%v handler=%v", err, embedded)
	}
	// The embedded handler must answer the public health probe.
	rec := httptest.NewRecorder()
	embedded.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/health", nil))
	if rec.Code != http.StatusOK {
		t.Errorf("embedded health = %d", rec.Code)
	}

	// Reverse proxy to an external REST base.
	proxied, err := buildAPIHandler(Config{RestBase: "http://127.0.0.1:9999", Listen: DefaultListen})
	if err != nil || proxied == nil {
		t.Fatalf("proxy handler: err=%v handler=%v", err, proxied)
	}

	// Invalid URL must be rejected with a descriptive error.
	if _, err := buildAPIHandler(Config{RestBase: "://bad", Listen: DefaultListen}); err == nil {
		t.Error("invalid --rest URL must fail")
	}
}

// TestNewReverseProxyDirector proves the proxy rewrites scheme/host/host header
// to the configured upstream while preserving the request path.
func TestNewReverseProxyDirector(t *testing.T) {
	// Exercise the Director by proxying to a local httptest upstream: the Host
	// header must be rewritten to the target and the request path preserved.
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The Director rewrites r.Host to the target host and preserves the path.
		_, _ = io.WriteString(w, "upstream-ok:"+r.URL.Path+":"+r.Host)
	}))
	t.Cleanup(upstream.Close)
	upstreamURL, _ := url.Parse(upstream.URL)
	proxy := newReverseProxy(upstreamURL)
	rec := httptest.NewRecorder()
	proxy.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "http://placeholder/api/v1/health", nil))
	body := rec.Body.String()
	if rec.Code != http.StatusOK || !strings.Contains(body, "/api/v1/health") ||
		!strings.Contains(body, upstreamURL.Host) {
		t.Errorf("proxy response = %d %q (want path + host %q)", rec.Code, body, upstreamURL.Host)
	}
}

// TestSpaRootBranches covers the StaticDir override and the embedded dist path,
// and proves spaHandler serves a real file from an injected directory as well as
// falling back to index.html for unknown client-side routes.
func TestSpaRootBranches(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte("INDEX-CONSOLE"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "app.js"), []byte("APP-JS"), 0o600); err != nil {
		t.Fatal(err)
	}
	// StaticDir override returns an os.DirFS rooted at the given directory.
	root := spaRoot(Config{StaticDir: dir})
	if _, err := root.Open("index.html"); err != nil {
		t.Fatalf("static root cannot open index.html: %v", err)
	}
	// Embedded fallback returns the dist sub-filesystem with the console.
	emb := spaRoot(Config{})
	if _, err := emb.Open("index.html"); err != nil {
		t.Fatalf("embedded root cannot open index.html: %v", err)
	}

	h := spaHandler(root)
	// A real file is served by the file server.
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/app.js", nil))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "APP-JS") {
		t.Errorf("real file = %d %q", rec.Code, rec.Body.String())
	}
	// Root path resolves to index.html.
	rec2 := httptest.NewRecorder()
	h.ServeHTTP(rec2, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec2.Code != http.StatusOK {
		t.Errorf("root = %d", rec2.Code)
	}
	// Unknown deep route falls back to the pre-read index.html.
	rec3 := httptest.NewRecorder()
	h.ServeHTTP(rec3, httptest.NewRequest(http.MethodGet, "/deep/client/route", nil))
	if rec3.Code != http.StatusOK || !strings.Contains(rec3.Body.String(), "INDEX-CONSOLE") {
		t.Errorf("spa fallback = %d %q", rec3.Code, rec3.Body.String())
	}
	if ct := rec3.Header().Get("Content-Type"); !strings.Contains(ct, "text/html") {
		t.Errorf("fallback content-type = %q", ct)
	}
}

// TestRunErrorBranches covers Run's failure paths that do not block: an invalid
// --rest URL makes buildAPIHandler fail (usage error), and an out-of-range port
// makes net.Listen fail (internal error). Both are deterministic and offline.
func TestRunErrorBranches(t *testing.T) {
	// Invalid --rest URL: buildAPIHandler returns an error -> usage error.
	var stdout, stderr bytes.Buffer
	code := Run([]string{"--rest", "://bad", "--listen", "127.0.0.1:0"}, &stdout, &stderr)
	if code != exitWebUsageError {
		t.Errorf("invalid rest code = %d, want %d", code, exitWebUsageError)
	}
	if !strings.Contains(stderr.String(), "invalid --rest") {
		t.Errorf("missing invalid-rest diagnostic: %q", stderr.String())
	}

	// Out-of-range loopback port: buildAPIHandler succeeds, net.Listen fails.
	var out2, err2 bytes.Buffer
	code = Run([]string{"--listen", "127.0.0.1:99999"}, &out2, &err2)
	if code != exitWebInternal {
		t.Errorf("bad port code = %d, want %d", code, exitWebInternal)
	}
	if !strings.Contains(err2.String(), "cannot bind") {
		t.Errorf("missing bind diagnostic: %q", err2.String())
	}

	// --static pointing at a dir with index still composes BuildHandler fine.
	dir := t.TempDir()
	_ = os.WriteFile(filepath.Join(dir, "index.html"), []byte("CONSOLE"), 0o600)
	h := BuildHandler(fakeAPI(), Config{Listen: DefaultListen, StaticDir: dir})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusOK {
		t.Errorf("BuildHandler with static dir = %d", rec.Code)
	}
}
