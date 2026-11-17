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
	"strings"
	"testing"
)

// fakeAPI is a stand-in for the REST server handler.
func fakeAPI() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"path":`+"`"+r.URL.Path+"`"+`}`)
	})
}

func newTestServer(t *testing.T, api http.Handler, cfg Config) *httptest.Server {
	t.Helper()
	h := BuildHandler(api, cfg)
	return httptest.NewServer(h)
}

func TestParseConfigDefaults(t *testing.T) {
	t.Parallel()
	cfg, code := ParseConfig(nil, io.Discard)
	if code != exitWebOK {
		t.Fatalf("code = %d, want %d", code, exitWebOK)
	}
	if cfg.Listen != DefaultListen {
		t.Fatalf("listen = %q, want %q", cfg.Listen, DefaultListen)
	}
	if cfg.RestBase != DefaultRestBase {
		t.Fatalf("rest = %q, want %q", cfg.RestBase, DefaultRestBase)
	}
}

func TestParseConfigFlags(t *testing.T) {
	t.Parallel()
	cfg, code := ParseConfig([]string{"--listen", "127.0.0.1:9000", "--token", "sek", "--rest", "http://x:1"}, io.Discard)
	if code != exitWebOK {
		t.Fatalf("code = %d, want %d", code, exitWebOK)
	}
	if cfg.Listen != "127.0.0.1:9000" || cfg.Token != "sek" || cfg.RestBase != "http://x:1" {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}

func TestParseConfigEmptyListenRejected(t *testing.T) {
	t.Parallel()
	var stderr bytes.Buffer
	cfg, code := ParseConfig([]string{"--listen", "   "}, &stderr)
	if code != exitWebUsageError {
		t.Fatalf("code = %d, want %d", code, exitWebUsageError)
	}
	if cfg.Listen != "" {
		t.Fatalf("listen should be empty, got %q", cfg.Listen)
	}
	if stderr.Len() == 0 {
		t.Fatal("expected help text on stderr")
	}
}

func TestParseConfigHelpIsNotError(t *testing.T) {
	t.Parallel()
	_, code := ParseConfig([]string{"--help"}, io.Discard)
	// --help is handled by Run, not ParseConfig; here it is an unknown flag, so we
	// only assert ParseConfig does not panic and returns a deterministic code.
	if code != exitWebUsageError {
		t.Fatalf("code = %d, want %d", code, exitWebUsageError)
	}
}

func TestBuildHandlerServesConsole(t *testing.T) {
	t.Parallel()
	srv := newTestServer(t, fakeAPI(), Config{Listen: DefaultListen})
	t.Cleanup(srv.Close)

	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatalf("get /: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	if !strings.Contains(string(body), "Abate Console") {
		t.Fatalf("console not served: %q", string(body))
	}
}

func TestBuildHandlerSPARoutingFallback(t *testing.T) {
	t.Parallel()
	srv := newTestServer(t, fakeAPI(), Config{Listen: DefaultListen})
	t.Cleanup(srv.Close)

	// A deep client-side route must fall back to index.html.
	resp, err := http.Get(srv.URL + "/checks/abc")
	if err != nil {
		t.Fatalf("get spa route: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	if !strings.Contains(string(body), "Abate Console") {
		t.Fatalf("spa fallback failed: %q", string(body))
	}
}

func TestBuildHandlerDelegatesAPI(t *testing.T) {
	t.Parallel()
	srv := newTestServer(t, fakeAPI(), Config{Listen: DefaultListen})
	t.Cleanup(srv.Close)

	for _, p := range []string{"/api/v1/health", "/openapi.json"} {
		resp, err := http.Get(srv.URL + p)
		if err != nil {
			t.Fatalf("get %s: %v", p, err)
		}
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if !strings.Contains(string(body), p) {
			t.Fatalf("api not delegated for %s: %q", p, string(body))
		}
	}
}

func TestRunHelpWritesUsage(t *testing.T) {
	t.Parallel()
	var stdout, stderr bytes.Buffer
	code := Run([]string{"--help"}, &stdout, &stderr)
	if code != exitWebOK {
		t.Fatalf("code = %d, want %d", code, exitWebOK)
	}
	if !strings.Contains(stdout.String(), "Usage: abate web") {
		t.Fatalf("help not printed: %q", stdout.String())
	}
}

func TestRunRejectsEmptyListen(t *testing.T) {
	t.Parallel()
	var stdout, stderr bytes.Buffer
	code := Run([]string{"--listen", ""}, &stdout, &stderr)
	if code != exitWebUsageError {
		t.Fatalf("code = %d, want %d", code, exitWebUsageError)
	}
}
