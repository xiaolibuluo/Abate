// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package llm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestNewClaudeClientDefaults 验证默认 URL 和 Model。
func TestNewClaudeClientDefaults(t *testing.T) {
	c := NewClaudeClient(Config{APIKey: "test-key"})
	if c.cfg.BaseURL != DefaultClaudeURL {
		t.Errorf("BaseURL = %q, want %q", c.cfg.BaseURL, DefaultClaudeURL)
	}
	if c.cfg.Model != "claude-sonnet-4-20250514" {
		t.Errorf("Model = %q, want claude-sonnet-4-20250514", c.cfg.Model)
	}
}

// TestNewClaudeClientCustomConfig 验证自定义配置。
func TestNewClaudeClientCustomConfig(t *testing.T) {
	c := NewClaudeClient(Config{APIKey: "k", Model: "claude-3", BaseURL: "http://custom"})
	if c.cfg.BaseURL != "http://custom" {
		t.Errorf("BaseURL = %q, want http://custom", c.cfg.BaseURL)
	}
	if c.cfg.Model != "claude-3" {
		t.Errorf("Model = %q, want claude-3", c.cfg.Model)
	}
}

// TestClaudeGenerateTestSuccess 验证成功路径。
func TestClaudeGenerateTestSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-api-key") != "test-key" {
			t.Errorf("api key header = %q", r.Header.Get("x-api-key"))
		}
		if r.Header.Get("anthropic-version") != "2023-06-01" {
			t.Errorf("anthropic-version = %q", r.Header.Get("anthropic-version"))
		}
		resp := claudeResponse{Content: []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}{{Type: "text", Text: "  func TestMutant(t *testing.T) {}  "}}}
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	c := NewClaudeClient(Config{APIKey: "test-key", BaseURL: srv.URL})
	result, err := c.GenerateTest(sampleMutation(), "package main")
	if err != nil {
		t.Fatalf("GenerateTest: %v", err)
	}
	if result != "func TestMutant(t *testing.T) {}" {
		t.Errorf("result = %q", result)
	}
}

// TestClaudeGenerateTestHTTPError 验证 HTTP 请求失败。
func TestClaudeGenerateTestHTTPError(t *testing.T) {
	orig := httpDoFn
	httpDoFn = func(c *BaseClient, req *http.Request) (*http.Response, error) {
		return nil, errors.New("connection refused")
	}
	defer func() { httpDoFn = orig }()

	c := NewClaudeClient(Config{APIKey: "k"})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "connection refused") {
		t.Errorf("err = %v, want connection refused", err)
	}
}

// TestClaudeGenerateTestAPIError 验证非 2xx 状态码。
func TestClaudeGenerateTestAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, "rate limited")
	}))
	defer srv.Close()

	c := NewClaudeClient(Config{APIKey: "bad", BaseURL: srv.URL})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "429") {
		t.Errorf("err = %v, want 429", err)
	}
}

// TestClaudeGenerateTestBadJSON 验证 JSON 解析失败。
func TestClaudeGenerateTestBadJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "not json at all")
	}))
	defer srv.Close()

	c := NewClaudeClient(Config{APIKey: "k", BaseURL: srv.URL})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "unmarshal") {
		t.Errorf("err = %v, want unmarshal error", err)
	}
}

// TestClaudeMarshalError 验证 json.Marshal 失败路径。
func TestClaudeMarshalError(t *testing.T) {
	orig := claudeMarshalFn
	claudeMarshalFn = func(v any) ([]byte, error) {
		return nil, errors.New("marshal failed")
	}
	defer func() { claudeMarshalFn = orig }()

	c := NewClaudeClient(Config{APIKey: "k"})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "marshal") {
		t.Errorf("err = %v, want marshal error", err)
	}
}

// TestClaudeGenerateTestEmptyContent 验证空 content 响应。
func TestClaudeGenerateTestEmptyContent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(claudeResponse{})
	}))
	defer srv.Close()

	c := NewClaudeClient(Config{APIKey: "k", BaseURL: srv.URL})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "no content") {
		t.Errorf("err = %v, want no content", err)
	}
}
