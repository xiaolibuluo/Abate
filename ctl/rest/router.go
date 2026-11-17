/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"net/http"
	"strings"
)

// route is one registered (method, path) handler with optional {param} segments.
type route struct {
	method   string
	segments []string
	handler  http.HandlerFunc
}

// router is a deterministic method+path dispatcher.
//
// It uses an exact, ordered segment match so registration order never affects
// dispatch: static segments must match literally and {name} segments bind one
// URL path component. No regular-expression or map-ordering dependency exists.
type router struct {
	routes []route
}

// newRouter returns an empty router ready for registration.
func newRouter() *router {
	return &router{routes: make([]route, 0, 24)}
}

// Handle registers one handler for a method and a slash-separated path. Path
// segments wrapped in {braces} are treated as single-component parameters.
func (r *router) Handle(method, path string, handler http.HandlerFunc) {
	trimmed := strings.Trim(path, "/")
	segments := make([]string, 0)
	if trimmed != "" {
		segments = strings.Split(trimmed, "/")
	}
	r.routes = append(r.routes, route{
		method: method, segments: segments, handler: handler,
	})
}

// match resolves a request to a registered handler and returns the extracted
// path parameters. A nil handler signals no route.
func (r *router) match(method, path string) (http.HandlerFunc, map[string]string) {
	segments := splitPath(path)
	for _, rt := range r.routes {
		if rt.method != method {
			continue
		}
		if len(rt.segments) != len(segments) {
			continue
		}
		params := make(map[string]string)
		ok := true
		for i, want := range rt.segments {
			got := segments[i]
			if strings.HasPrefix(want, "{") && strings.HasSuffix(want, "}") {
				params[want[1:len(want)-1]] = got
				continue
			}
			if want != got {
				ok = false
				break
			}
		}
		if ok {
			return rt.handler, params
		}
	}
	return nil, nil
}

// splitPath returns the non-empty slash-separated segments of a request path.
func splitPath(path string) []string {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		return []string{}
	}
	return strings.Split(trimmed, "/")
}

// ServeHTTP dispatches a request, returning 405 for a known path with the wrong
// method and 404 for an unknown path, both as RFC 9457 problem documents.
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, params := r.match(req.Method, req.URL.Path)
	if handler == nil {
		handler, _ = r.match("GET", req.URL.Path)
		if handler != nil {
			writeMethodNotAllowed(w, req)
			return
		}
		writeNotFound(w, req)
		return
	}
	ctx := withParams(req.Context(), params)
	handler(w, req.WithContext(ctx))
}
