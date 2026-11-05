// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// validDSL 是一个通过 CES 第 2 章校验的合法实验 DSL。
const validDSL = `id: exp-web-test
seed: 7
criticality: 3
target:
  kind: service
  name: pay
fault:
  type: F3
  loss: 30%
blastRadius:
  maxInstances: 1
  maxDurationMs: 60000
steadyState:
  - sli: http_success_rate
    assert: ">= 99%"
`

// newServer 构造测试用 HTTP 服务器。
func newServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(NewHandler().Routes())
	t.Cleanup(srv.Close)
	return srv
}

// getJSON 发起 GET 并解码 JSON 响应。
func getJSON(t *testing.T, url string, out any) int {
	t.Helper()
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("GET %s: %v", url, err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		t.Fatalf("decode %s: %v", url, err)
	}
	return resp.StatusCode
}

// TestNewHandler 验证处理器构造与演示数据装载。
func TestNewHandler(t *testing.T) {
	h := NewHandler()
	if h == nil || len(h.experiments) == 0 {
		t.Fatal("NewHandler must preload demo experiment")
	}
	if h.startedAt.IsZero() {
		t.Error("startedAt must be set")
	}
}

// TestHealth 验证健康检查端点。
func TestHealth(t *testing.T) {
	srv := newServer(t)
	var body map[string]string
	if code := getJSON(t, srv.URL+"/api/health", &body); code != 200 {
		t.Fatalf("health code=%d", code)
	}
	if body["status"] != "ok" {
		t.Errorf("health body=%v", body)
	}
}

// TestVersion 验证版本端点包含品牌与双证标识。
func TestVersion(t *testing.T) {
	srv := newServer(t)
	var body map[string]string
	if code := getJSON(t, srv.URL+"/api/version", &body); code != 200 {
		t.Fatalf("version code=%d", code)
	}
	for _, k := range []string{"name", "version", "spec", "license", "vendor"} {
		if body[k] == "" {
			t.Errorf("version body missing %q: %v", k, body)
		}
	}
	if body["name"] != "Abate" {
		t.Errorf("name=%q", body["name"])
	}
}

// TestListExperiments 验证实验列表端点返回演示数据。
func TestListExperiments(t *testing.T) {
	srv := newServer(t)
	var views []ExperimentView
	if code := getJSON(t, srv.URL+"/api/experiments", &views); code != 200 {
		t.Fatalf("list code=%d", code)
	}
	if len(views) == 0 || views[0].ID != "demo-latency" {
		t.Errorf("views=%v", views)
	}
}

// TestCreateAndGetExperiment 验证提交合法 DSL 并可回读。
func TestCreateAndGetExperiment(t *testing.T) {
	srv := newServer(t)
	resp, err := http.Post(srv.URL+"/api/experiments",
		"text/plain; charset=utf-8", strings.NewReader(validDSL))
	if err != nil {
		t.Fatalf("POST: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("create code=%d, want 201", resp.StatusCode)
	}
	var view ExperimentView
	if err := json.NewDecoder(resp.Body).Decode(&view); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if view.ID != "exp-web-test" || view.Status != "Pending" {
		t.Errorf("view=%+v", view)
	}

	var detail map[string]any
	if code := getJSON(t, srv.URL+"/api/experiments/exp-web-test", &detail); code != 200 {
		t.Fatalf("get code=%d", code)
	}
	if detail["id"] != "exp-web-test" {
		t.Errorf("detail=%v", detail)
	}
}

// TestCreateExperimentInvalid 验证非法 DSL 被拒绝（400 结构化错误）。
func TestCreateExperimentInvalid(t *testing.T) {
	srv := newServer(t)
	resp, err := http.Post(srv.URL+"/api/experiments",
		"text/plain", strings.NewReader("id: bad\n"))
	if err != nil {
		t.Fatalf("POST: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("invalid DSL code=%d, want 400", resp.StatusCode)
	}
}

// TestCreateExperimentTooLarge 验证超大请求体被拒绝（413）。
func TestCreateExperimentTooLarge(t *testing.T) {
	srv := newServer(t)
	big := strings.NewReader("id: " + strings.Repeat("x", (1<<20)+8) + "\n")
	resp, err := http.Post(srv.URL+"/api/experiments", "text/plain", big)
	if err != nil {
		t.Fatalf("POST: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Errorf("oversized body code=%d, want 413", resp.StatusCode)
	}
}

// TestGetExperimentNotFound 验证未命中与非法 ID 的 404 分支。
func TestGetExperimentNotFound(t *testing.T) {
	srv := newServer(t)
	var body map[string]string
	if code := getJSON(t, srv.URL+"/api/experiments/nope", &body); code != 404 {
		t.Errorf("missing id code=%d, want 404", code)
	}
	if code := getJSON(t, srv.URL+"/api/experiments/", &body); code != 404 {
		t.Errorf("empty id code=%d, want 404", code)
	}
	if code := getJSON(t, srv.URL+"/api/experiments/a/b", &body); code != 404 {
		t.Errorf("nested id code=%d, want 404", code)
	}
}

// TestSpecClauses 验证规范章节端点返回八章目录。
func TestSpecClauses(t *testing.T) {
	srv := newServer(t)
	var clauses []map[string]string
	if code := getJSON(t, srv.URL+"/api/spec/clauses", &clauses); code != 200 {
		t.Fatalf("clauses code=%d", code)
	}
	if len(clauses) != 8 {
		t.Errorf("len(clauses)=%d, want 8", len(clauses))
	}
}

// TestStaticIndex 验证根路径返回内嵌控制台页面（fs.Sub 语义）。
func TestStaticIndex(t *testing.T) {
	srv := newServer(t)
	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatalf("GET /: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("GET / code=%d", resp.StatusCode)
	}
	buf := make([]byte, 4096)
	n, _ := resp.Body.Read(buf)
	if !strings.Contains(string(buf[:n]), "Abate") {
		t.Error("index page must contain brand")
	}
}

// TestMustSubFS 验证嵌入子树提取：正常路径可用，非法路径 panic（构建缺陷立即暴露）。
func TestMustSubFS(t *testing.T) {
	if sub := mustSubFS(assets, "static"); sub == nil {
		t.Fatal("mustSubFS(static) must succeed")
	}
	defer func() {
		if recover() == nil {
			t.Error("invalid dir must panic")
		}
	}()
	mustSubFS(assets, "../escape")
}

// TestListenBadAddr 验证非法监听地址立即返回错误。
func TestListenBadAddr(t *testing.T) {
	if err := NewHandler().Listen("127.0.0.1:not-a-port"); err == nil {
		t.Error("Listen on bad addr must fail")
	}
}
