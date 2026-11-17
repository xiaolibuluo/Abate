/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"errors"
	"runtime"
	"strings"
)

// ServerConfig holds the tunable parameters of a control server.
//
// The zero value is not usable: New validates Listen and Auth, rejecting a
// remote bind address without a token so the server can never start exposed and
// unauthenticated (PLAN-07 section 4).
type ServerConfig struct {
	// Listen is the host:port to bind. Defaults to 127.0.0.1:8420.
	Listen string
	// Auth is the required Bearer token. It is mandatory for every endpoint
	// except /health, /ready and /version, even on localhost.
	Auth string
	// MaxConcurrent caps simultaneously executing experiment/check tasks.
	MaxConcurrent int
	// RequestTimeout bounds the per-request handler duration.
	RequestTimeoutSeconds int
	// QueueSize bounds the pending task queue; a fuller queue returns 429.
	QueueSize int
	// RateLimit is the tokens-per-second ceiling per client IP.
	RateLimit int
}

// DefaultServerConfig returns the frozen default configuration.
func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Listen:                "127.0.0.1:8420",
		MaxConcurrent:         runtime.NumCPU(),
		RequestTimeoutSeconds: 30,
		QueueSize:             100,
		RateLimit:             100,
	}
}

// ErrRemoteWithoutAuth reports a remote bind address paired with an empty token.
var ErrRemoteWithoutAuth = errors.New("rest: remote bind requires a non-empty auth token")

// ErrInvalidListen reports a missing or malformed bind address.
var ErrInvalidListen = errors.New("rest: invalid listen address")

// Validate normalizes the configuration and enforces the auth/network invariant.
func (c ServerConfig) Validate() (ServerConfig, error) {
	if strings.TrimSpace(c.Listen) == "" {
		return ServerConfig{}, ErrInvalidListen
	}
	if c.MaxConcurrent <= 0 {
		c.MaxConcurrent = runtime.NumCPU()
	}
	if c.RequestTimeoutSeconds <= 0 {
		c.RequestTimeoutSeconds = 30
	}
	if c.QueueSize <= 0 {
		c.QueueSize = 100
	}
	if c.RateLimit <= 0 {
		c.RateLimit = 100
	}
	host := c.Listen
	if idx := strings.LastIndex(c.Listen, ":"); idx >= 0 {
		host = c.Listen[:idx]
	}
	if host != "127.0.0.1" && host != "localhost" && host != "::1" {
		if strings.TrimSpace(c.Auth) == "" {
			return ServerConfig{}, ErrRemoteWithoutAuth
		}
	}
	return c, nil
}

// authExempt reports whether a path may be served without a bearer token.
func authExempt(path string) bool {
	switch path {
	case "/api/" + MetaAPIVersion + "/health",
		"/api/" + MetaAPIVersion + "/ready",
		"/api/" + MetaAPIVersion + "/version",
		"/openapi.json":
		return true
	default:
		return false
	}
}
