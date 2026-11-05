// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。网络交互使用同样触发 AGPLv3 Copyleft。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details. Network use triggers AGPLv3 copyleft.

// Package web 实现 Abate Web 控制台：REST API + 内嵌静态资源。
// 零依赖红线：仅标准库 net/http + embed，无任何三方 HTTP 框架。
// AGPLv3 说明：本控制台通过网络向用户提供服务，对应源码依 AGPLv3 公开。
//
// Package web implements the Abate web console: REST API + embedded static
// assets. Zero-dependency red line: stdlib net/http + embed only. Under
// AGPLv3, this console serves users over a network and its corresponding
// source is published accordingly.
package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/dsl"
)

//go:embed static
var assets embed.FS

// Version 由构建期注入（-X main.version=...）。
// Version is injected at build time.
var Version = "dev"

// ExperimentView 是实验的 API 视图。
// ExperimentView is the API view of one experiment.
type ExperimentView struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Seed     int64  `json:"seed"`
	Created  string `json:"created"`
	ScoreBP  int64  `json:"scoreBp,omitempty"`
}

// Handler 是控制台 HTTP 处理器集合。
// Handler is the console HTTP handler set.
type Handler struct {
	experiments []core.Experiment
	startedAt   time.Time
}

// NewHandler 构造处理器并挂载演示数据（真实数据由 store 包接管后替换）。
// NewHandler constructs the handler with demo data until the store package
// takes over persistence.
func NewHandler() *Handler {
	demo := core.Experiment{
		ID: "demo-latency", Seed: 42, Criticality: 4,
		Target: core.Target{Kind: core.TargetService, Name: "order-service"},
		Fault: core.Fault{Class: core.F1Latency,
			Params: map[string]string{"delay": "500ms"}},
		Blast:  core.BlastRadius{MaxInstances: 1, MaxDurationMs: 300000},
		Steady: []core.SLIAssert{{SLI: "http_success_rate", Assert: ">= 99.9%"}},
	}
	return &Handler{experiments: []core.Experiment{demo}, startedAt: time.Now()}
}

// Routes 注册全部路由到给定 ServeMux。
// Routes registers all routes onto the given ServeMux.
func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", h.health)
	mux.HandleFunc("GET /api/version", h.version)
	mux.HandleFunc("GET /api/experiments", h.listExperiments)
	mux.HandleFunc("POST /api/experiments", h.createExperiment)
	mux.HandleFunc("GET /api/experiments/", h.getExperiment)
	mux.HandleFunc("GET /api/spec/clauses", h.specClauses)
	mux.HandleFunc("GET /api/v1/hub/subcodes", hubSubcodesHandler)
	// mustSubFS 提取 static 子树：使 GET / 命中 index.html 而非目录列表。
	// mustSubFS extracts the static subtree so GET / serves index.html.
	mux.Handle("GET /", http.FileServerFS(mustSubFS(assets, "static")))
	return mux
}

// mustSubFS 提取嵌入文件系统的子树；失败即 panic——
// 编译期内嵌资源失败属于不可恢复的构建缺陷，启动时立即暴露。
// mustSubFS extracts an embedded FS subtree; failure panics, since a broken
// compile-time embedding is an unrecoverable build defect best surfaced at startup.
func mustSubFS(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic("web: embedded subtree " + dir + " missing: " + err.Error())
	}
	return sub
}

// Listen 启动 HTTP 服务；addr 形如 ":8080"。
// Listen starts the HTTP server on addr (e.g. ":8080").
func (h *Handler) Listen(addr string) error {
	srv := &http.Server{
		Addr:              addr,
		Handler:           h.Routes(),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	return srv.ListenAndServe()
}

// writeJSON 统一 JSON 响应（确定性 Content-Type 与状态码）。
// writeJSON writes a uniform JSON response.
func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

// writeError 统一错误响应（CES-2.1.2 结构化错误风格）。
// writeError writes a uniform structured error response.
func writeError(w http.ResponseWriter, code int, path, reason string) {
	writeJSON(w, code, map[string]string{
		"path": path, "reason": reason,
	})
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) version(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"name":    "Abate",
		"version": Version,
		"spec":    "CES v0.1",
		"license": "AGPL-3.0-only + commercial",
		"vendor":  "小黎的部落 www.qingdujian.com",
	})
}

func (h *Handler) listExperiments(w http.ResponseWriter, _ *http.Request) {
	views := make([]ExperimentView, 0, len(h.experiments))
	for _, e := range h.experiments {
		views = append(views, ExperimentView{
			ID: e.ID, Status: "Pending", Seed: e.Seed,
			Created: h.startedAt.UTC().Format(time.RFC3339),
		})
	}
	writeJSON(w, http.StatusOK, views)
}

func (h *Handler) createExperiment(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 0, 1<<20)
	buf := make([]byte, 32<<10)
	for {
		n, err := r.Body.Read(buf)
		body = append(body, buf[:n]...)
		// 大小检查必须先于 err 判断：末块随 EOF 返回时不得跳过超限判定。
		// The size check must precede the err break: the final chunk arrives
		// with EOF and must not skip the limit check.
		if len(body) > 1<<20 {
			writeError(w, http.StatusRequestEntityTooLarge,
				"body", "experiment DSL exceeds 1 MiB")
			return
		}
		if err != nil {
			break
		}
	}
	e, err := dsl.Parse(string(body))
	if err != nil {
		writeError(w, http.StatusBadRequest, "dsl", err.Error())
		return
	}
	h.experiments = append(h.experiments, e)
	writeJSON(w, http.StatusCreated, ExperimentView{
		ID: e.ID, Status: "Pending", Seed: e.Seed,
		Created: time.Now().UTC().Format(time.RFC3339),
	})
}

func (h *Handler) getExperiment(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/experiments/")
	if id == "" || strings.Contains(id, "/") {
		writeError(w, http.StatusNotFound, "id", "experiment not found")
		return
	}
	for _, e := range h.experiments {
		if e.ID == id {
			writeJSON(w, http.StatusOK, e)
			return
		}
	}
	writeError(w, http.StatusNotFound, "id", fmt.Sprintf("experiment %q not found", id))
}

// specClauses 返回 CES 条款目录摘要（规范网页展示用）。
// specClauses returns a summary of CES clauses for spec display.
func (h *Handler) specClauses(w http.ResponseWriter, _ *http.Request) {
	clauses := []map[string]string{
		{"id": "CES-1", "title": "术语与故障分类学 Terminology & Fault Taxonomy"},
		{"id": "CES-2", "title": "实验模型 DSL Experiment Model"},
		{"id": "CES-3", "title": "执行语义与状态机 Execution Semantics"},
		{"id": "CES-4", "title": "可靠性评分算法 Scoring Algorithm"},
		{"id": "CES-5", "title": "观测与判定 Observation & Verdict"},
		{"id": "CES-6", "title": "报告格式 Report Format"},
		{"id": "CES-7", "title": "插件与适配器协议 Plugin Protocol"},
		{"id": "CES-8", "title": "一致性认证 Conformance"},
	}
	writeJSON(w, http.StatusOK, clauses)
}
