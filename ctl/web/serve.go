/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package web

import (
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"abate/ctl/rest"
)

// webHelp documents the web command and its copyable examples.
const webHelp = `Usage: abate web [flags]

Serve the local control API together with a self-contained browser console from
a single process. The console is served same-origin so no CORS configuration is
needed and the API is never exposed separately.

Flags:
  --listen host:port  Bind address (default 127.0.0.1:8430).
  --token string      Bearer token required for API access. Mandatory for a
                      non-loopback bind address; the server refuses to start
                      exposed and unauthenticated.
  --rest string       External REST base URL. When set, /api is reverse-proxied
                      to this base while the console is still served locally.
  --static dir        Serve static assets from a directory instead of the
                      embedded build (development only).
  --help, -h          Print this help text.

Examples:
  abate web
  abate web --listen 127.0.0.1:8443 --token s3cr3t
  abate web --rest http://10.0.0.5:8420
`

// writeWebHelp prints the deterministic web usage text.
func writeWebHelp(w io.Writer) {
	_, _ = fmt.Fprint(w, webHelp)
}

// Run starts the web command: it builds the request handler from the parsed
// configuration and serves until the listener fails.
func Run(args []string, stdout, stderr io.Writer) int {
	if helpRequested(args) {
		writeWebHelp(stdout)
		return exitWebOK
	}
	cfg, code := ParseConfig(args, stderr)
	if code != exitWebOK {
		return code
	}
	api, err := buildAPIHandler(cfg)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate web: %v\n", err)
		return exitWebUsageError
	}
	root := BuildHandler(api, cfg)
	ln, err := net.Listen("tcp", cfg.Listen)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "abate web: cannot bind %s: %v\n", cfg.Listen, err)
		return exitWebInternal
	}
	srv := &http.Server{
		Handler:           root,
		ReadHeaderTimeout: 5 * time.Second,
	}
	_, _ = fmt.Fprintf(stdout, "abate web: listening on %s\n", cfg.Listen)
	if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
		_, _ = fmt.Fprintf(stderr, "abate web: %v\n", err)
		return exitWebInternal
	}
	return exitWebOK
}

// buildAPIHandler returns the handler for /api and /openapi.json. When RestBase is
// empty it embeds the internal REST server; otherwise it reverse-proxies to it.
func buildAPIHandler(cfg Config) (http.Handler, error) {
	if strings.TrimSpace(cfg.RestBase) == "" {
		srv, err := rest.New(rest.ServerConfig{
			Listen: cfg.Listen,
			Auth:   cfg.Token,
		}, rest.ServerOptions{})
		if err != nil {
			return nil, err
		}
		return srv.Handler(), nil
	}
	target, err := url.Parse(cfg.RestBase)
	if err != nil {
		return nil, fmt.Errorf("web: invalid --rest %q: %w", cfg.RestBase, err)
	}
	return newReverseProxy(target), nil
}

// BuildHandler composes the API handler with the embedded console. Non-API routes
// fall back to the single-page console (index.html) so client-side routing works.
func BuildHandler(api http.Handler, cfg Config) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/", api)
	mux.Handle("/openapi.json", api)
	mux.HandleFunc("/", spaHandler(spaRoot(cfg)))
	return mux
}

// spaRoot returns the filesystem backing the console assets.
func spaRoot(cfg Config) fs.FS {
	if strings.TrimSpace(cfg.StaticDir) != "" {
		return os.DirFS(cfg.StaticDir)
	}
	sub, err := fs.Sub(assets, "dist")
	if err != nil {
		// embed.FS is a compile-time constant; this can only fail on a programming
		// mistake, never at runtime for a correctly built binary.
		panic("web: cannot open embedded dist: " + err.Error())
	}
	return sub
}

// spaHandler serves the console assets and falls back to index.html for any path
// that does not name a real file, enabling client-side routing. The fallback uses
// a pre-read copy of index.html rather than http.FileServer so a deep client-side
// route can never trigger FileServer's directory-trailing-slash redirect loop.
func spaHandler(root fs.FS) http.HandlerFunc {
	fileServer := http.FileServer(http.FS(root))
	index, _ := fs.ReadFile(root, "index.html")
	return func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(path.Clean("/"+r.URL.Path), "/")
		if name == "" || name == "." {
			name = "index.html"
		}
		if f, err := root.Open(name); err == nil {
			_ = f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(index)
	}
}

// newReverseProxy returns an HTTP reverse proxy to target for /api routing.
func newReverseProxy(target *url.URL) http.Handler {
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.Host = target.Host
		},
	}
}
