/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"abate/security/redact"
)

// statusRecorder captures the response status and byte count for logging.
type statusRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

// WriteHeader records the first status code written.
func (s *statusRecorder) WriteHeader(code int) {
	if s.status == 0 {
		s.status = code
	}
	s.ResponseWriter.WriteHeader(code)
}

// Write counts the bytes written for the structured log line.
func (s *statusRecorder) Write(b []byte) (int, error) {
	if s.status == 0 {
		s.status = http.StatusOK
	}
	n, err := s.ResponseWriter.Write(b)
	s.bytes += n
	return n, err
}

// recoverMiddleware converts a handler panic into a 500 problem document so a
// single bad request cannot crash the long-lived server.
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("rest: panic recovered on %s %s: %s",
					req.Method, redact.Path(req.URL.Path), redact.Text(fmt.Sprint(r)))
				writeInternalError(w, req, "an internal error occurred")
			}
		}()
		next.ServeHTTP(w, req)
	})
}

// logMiddleware emits one structured line per request. It deliberately omits the
// clock: the control plane stays free of wall-clock dependence in its outputs.
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		rec := &statusRecorder{ResponseWriter: w}
		next.ServeHTTP(rec, req)
		if rec.status == 0 {
			rec.status = http.StatusOK
		}
		log.Printf("rest: %s %s %d %d", req.Method, redact.Path(req.URL.Path), rec.status, rec.bytes)
	})
}

// authMiddleware enforces the bearer token except for exempt system probes.
func authMiddleware(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if authExempt(req.URL.Path) {
			next.ServeHTTP(w, req)
			return
		}
		if !verifyBearer(req, token) {
			writeUnauthorized(w, req)
			return
		}
		next.ServeHTTP(w, req)
	})
}

// ipLimiter is a deterministic, time-free per-client rate limiter.
//
// Each client IP receives a fixed token bucket of size cap. Tokens are consumed
// per request; when exhausted the request is rejected with 429. To avoid a
// permanent lockout without depending on the wall clock, every global resetStep
// served requests refills all buckets to full. The mechanism is reproducible:
// identical request streams yield identical acceptance decisions.
type ipLimiter struct {
	mu        sync.Mutex
	buckets   map[string]int
	cap       int
	served    int
	resetStep int
}

// newIPLimiter returns a limiter with the given per-IP token capacity.
func newIPLimiter(capacity int) *ipLimiter {
	if capacity <= 0 {
		capacity = 100
	}
	return &ipLimiter{
		buckets:   make(map[string]int),
		cap:       capacity,
		resetStep: maxInt(1024, capacity*8),
	}
}

// allow reports whether a request from client may proceed, consuming one token.
func (l *ipLimiter) allow(client string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.served++
	if l.served%l.resetStep == 0 {
		for k := range l.buckets {
			l.buckets[k] = l.cap
		}
	}
	remaining, ok := l.buckets[client]
	if !ok {
		remaining = l.cap
	}
	if remaining <= 0 {
		return false
	}
	l.buckets[client] = remaining - 1
	return true
}

// rateLimitMiddleware rejects over-quota clients with 429 + Retry-After.
func rateLimitMiddleware(limiter *ipLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		client := clientIP(req)
		if !limiter.allow(client) {
			writeTooManyRequests(w, req)
			return
		}
		next.ServeHTTP(w, req)
	})
}

// clientIP returns the originating client IP, preferring the forwarded header
// only when it carries a single address, else the remote host.
func clientIP(req *http.Request) string {
	if fwd := req.Header.Get("X-Forwarded-For"); fwd != "" {
		if idx := indexByte(fwd, ','); idx >= 0 {
			return trimSpace(fwd[:idx])
		}
		return trimSpace(fwd)
	}
	host := req.RemoteAddr
	if idx := lastIndexByte(host, ':'); idx >= 0 {
		return host[:idx]
	}
	return host
}

// maxInt returns the larger of two integers without depending on math.Max.
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// indexByte returns the first index of c in s, or -1.
func indexByte(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// lastIndexByte returns the last index of c in s, or -1.
func lastIndexByte(s string, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// trimSpace trims ASCII whitespace from both ends.
func trimSpace(s string) string {
	start, end := 0, len(s)
	for start < end && isSpace(s[start]) {
		start++
	}
	for end > start && isSpace(s[end-1]) {
		end--
	}
	return s[start:end]
}

// isSpace reports whether b is an ASCII whitespace byte.
func isSpace(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r', '\f', '\v':
		return true
	default:
		return false
	}
}
