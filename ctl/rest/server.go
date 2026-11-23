/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

// errQueueFull reports the task queue is at capacity (maps to HTTP 429).
var errQueueFull = errors.New("rest: task queue is full")

// Server is the Abate local control API. It is a pure orchestration layer: it
// reuses the language-neutral checker, an injectable task runner and an
// injectable plugin manager, and never embeds business logic or network I/O.
type Server struct {
	cfg     ServerConfig
	router  *router
	store   *taskStore
	runner  TaskRunner
	plugins PluginManager
	limiter *ipLimiter
	sem     chan struct{}
	handler http.Handler
	http    *http.Server
	ln      net.Listener
}

// New builds a server from a configuration, validating the auth/network
// invariant and mounting every control endpoint.
func New(cfg ServerConfig, opts ServerOptions) (*Server, error) {
	cfg, err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	runner := opts.Runner
	if runner == nil {
		runner = checkRunner{}
	}
	plugins := opts.Plugins
	if plugins == nil {
		plugins = newStaticPluginManager()
	}
	s := &Server{
		cfg:     cfg,
		router:  newRouter(),
		store:   newTaskStore(),
		runner:  runner,
		plugins: plugins,
		limiter: newIPLimiter(cfg.RateLimit),
		sem:     make(chan struct{}, cfg.MaxConcurrent),
	}
	s.mountRoutes()
	s.handler = s.buildHandler()
	return s, nil
}

// mountRoutes registers every control endpoint in a fixed, deterministic order.
func (s *Server) mountRoutes() {
	const base = "/api/" + MetaAPIVersion
	s.router.Handle("GET", base+"/health", s.handleHealth)
	s.router.Handle("GET", base+"/ready", s.handleReady)
	s.router.Handle("GET", base+"/version", s.handleVersion)
	s.router.Handle("POST", base+"/checks", s.handleCreateCheck)
	s.router.Handle("GET", base+"/checks", s.handleListChecks)
	s.router.Handle("GET", base+"/checks/{id}", s.handleGetTask)
	s.router.Handle("GET", base+"/checks/{id}/report", s.handleExperimentReport)
	s.router.Handle("GET", base+"/checks/{id}/archive", s.handleExperimentArchive)
	s.router.Handle("POST", base+"/experiments", s.handleCreateExperiment)
	s.router.Handle("GET", base+"/experiments", s.handleListExperiments)
	s.router.Handle("GET", base+"/experiments/{id}", s.handleGetTask)
	s.router.Handle("POST", base+"/experiments/{id}/cancel", s.handleCancelExperiment)
	s.router.Handle("GET", base+"/experiments/{id}/report", s.handleExperimentReport)
	s.router.Handle("GET", base+"/experiments/{id}/archive", s.handleExperimentArchive)
	s.router.Handle("GET", base+"/plugins", s.handleListPlugins)
	s.router.Handle("POST", base+"/plugins", s.handleInstallPlugin)
	s.router.Handle("DELETE", base+"/plugins/{id}", s.handleRemovePlugin)
	s.router.Handle("POST", base+"/plugins/{id}/upgrade", s.handleUpgradePlugin)
	s.router.Handle("GET", "/openapi.json", s.handleOpenAPI)
}

// buildHandler wraps the router in the fixed middleware chain.
func (s *Server) buildHandler() http.Handler {
	h := http.Handler(s.router)
	h = rateLimitMiddleware(s.limiter, h)
	h = authMiddleware(s.cfg.Auth, h)
	h = logMiddleware(h)
	h = recoverMiddleware(h)
	h = http.TimeoutHandler(h, time.Duration(s.cfg.RequestTimeoutSeconds)*time.Second,
		"request timeout")
	return h
}

// Handler returns the fully wrapped request handler, usable with httptest.
func (s *Server) Handler() http.Handler { return s.handler }

// ListenAndServe binds the configured address and serves until shutdown.
func (s *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.cfg.Listen)
	if err != nil {
		return fmt.Errorf("%w: %w", errBindFailed, err)
	}
	s.ln = ln
	s.http = &http.Server{
		Handler:           s.handler,
		ReadHeaderTimeout: time.Duration(s.cfg.RequestTimeoutSeconds) * time.Second,
		WriteTimeout:      120 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	return s.http.Serve(ln)
}

// Shutdown gracefully stops the server, draining in-flight requests.
func (s *Server) Shutdown(ctx context.Context) error {
	if s.http == nil {
		return nil
	}
	return s.http.Shutdown(ctx)
}

// submit reserves a task and launches it under the concurrency limit.
func (s *Server) submit(req multiCheckRequest, kind string) (*Task, error) {
	if s.store.activeCount() >= s.cfg.QueueSize {
		return nil, errQueueFull
	}
	task := s.store.newTask(kind, req.Language, req.Format)
	s.runTask(task.ID, req)
	return task, nil
}

// runTask executes a task asynchronously, updating its status deterministically.
func (s *Server) runTask(id string, req multiCheckRequest) {
	go func() {
		s.sem <- struct{}{}
		defer func() { <-s.sem; s.store.finish(id) }()
		s.store.update(id, func(t *Task) {
			if t.Cancelled {
				t.Status = StatusCancelled
			} else {
				t.Status = StatusRunning
			}
		})
		if snap, ok := s.store.get(id); ok && snap.Cancelled {
			return
		}
		result, runErr := s.runner.Run(req)
		s.store.update(id, func(t *Task) {
			if t.Cancelled {
				t.Status = StatusIncomplete
				if runErr == nil {
					t.Report = result.Result.Render(req.Format)
					t.FingerprintHex = result.Result.FingerprintHex()
					t.Summary = toSummaryView(result.Summary)
				}
				return
			}
			if runErr != nil {
				t.Status = StatusFailed
				t.Detail = runErr.Error()
				return
			}
			t.Status = StatusCompleted
			t.FingerprintHex = result.Result.FingerprintHex()
			t.Summary = toSummaryView(result.Summary)
			t.Report = result.Result.Render(req.Format)
		})
	}()
}

// errBindFailed reports a failure to bind the listen address.
var errBindFailed = errors.New("rest: cannot bind listen address")
