/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"context"
	"encoding/json"
	"net/http"
)

// Problem is the RFC 9457 problem+json payload emitted by the control API.
//
// The Instance field carries a deterministic request identifier (a counter, not
// a random UUID) so error reports stay reproducible across runs.
type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

// problemTypes map internal error classes to stable type URIs.
const (
	problemValidation = "about:blank#validation"
	problemNotFound   = "about:blank#not-found"
	problemConflict   = "about:blank#conflict"
	problemAuth       = "about:blank#unauthorized"
	problemRate       = "about:blank#too-many-requests"
	problemInternal   = "about:blank#internal"
)

// paramContextKey is the context key under which path parameters are stored.
type paramContextKey struct{}

// withParams returns a context carrying the decoded path parameters.
func withParams(ctx context.Context, params map[string]string) context.Context {
	return context.WithValue(ctx, paramContextKey{}, params)
}

// routeParam reads a path parameter previously bound by the router.
func routeParam(ctx context.Context, name string) string {
	params, ok := ctx.Value(paramContextKey{}).(map[string]string)
	if !ok || params == nil {
		return ""
	}
	return params[name]
}

// writeProblem serializes a problem document with the correct content type and
// status code. The token, if any, is never echoed into the body or logs.
func writeProblem(w http.ResponseWriter, status int, problemType, title, detail, instance string) {
	body, err := json.Marshal(Problem{
		Type: problemType, Title: title, Status: status,
		Detail: detail, Instance: instance,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func writeBadRequest(w http.ResponseWriter, req *http.Request, detail string) {
	writeProblem(w, http.StatusBadRequest, problemValidation,
		"Bad Request", detail, requestInstance(req))
}

func writeNotFound(w http.ResponseWriter, req *http.Request) {
	writeProblem(w, http.StatusNotFound, problemNotFound,
		"Not Found", "the requested resource does not exist", requestInstance(req))
}

func writeMethodNotAllowed(w http.ResponseWriter, req *http.Request) {
	writeProblem(w, http.StatusMethodNotAllowed, problemValidation,
		"Method Not Allowed", "the method is not permitted on this resource",
		requestInstance(req))
}

func writeConflict(w http.ResponseWriter, req *http.Request, detail string) {
	writeProblem(w, http.StatusConflict, problemConflict,
		"Conflict", detail, requestInstance(req))
}

func writeUnauthorized(w http.ResponseWriter, req *http.Request) {
	writeProblem(w, http.StatusUnauthorized, problemAuth,
		"Unauthorized", "a valid bearer token is required", requestInstance(req))
}

func writeTooManyRequests(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Retry-After", "1")
	writeProblem(w, http.StatusTooManyRequests, problemRate,
		"Too Many Requests", "the request rate limit has been exceeded",
		requestInstance(req))
}

func writeInternalError(w http.ResponseWriter, req *http.Request, detail string) {
	writeProblem(w, http.StatusInternalServerError, problemInternal,
		"Internal Server Error", detail, requestInstance(req))
}

// writeJSON serializes v as a deterministic JSON document with the given status.
func writeJSON(w http.ResponseWriter, status int, v any) {
	body, err := json.Marshal(v)
	if err != nil {
		writeInternalError(w, nil, "response encoding failed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}
