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
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/qingdujian/abate/core"
)

func sampleMutation() core.Mutation {
	return core.Mutation{
		File:  "main.go",
		Line:  10,
		Class: "M1",
		From:  "+",
		To:    "-",
	}
}

// TestNewOpenAIClientDefaults 验证默认 URL 和 Model。
func TestNewOpenAIClientDefaults(t *testing.T) {
	c := NewOpenAIClient(Config{APIKey: "test-key"})
	if c.cfg.BaseURL != DefaultOpenAIURL {
		t.Errorf("BaseURL = %q, want %q", c.cfg.BaseURL, DefaultOpenAIURL)
	}
	if c.cfg.Model != "gpt-4o" {
		t.Errorf("Model = %q, want gpt-4o", c.cfg.Model)
	}
}

// TestNewOpenAIClientCustomConfig 验证自定义配置。
func TestNewOpenAIClientCustomConfig(t *testing.T) {
	c := NewOpenAIClient(Config{APIKey: "k", Model: "gpt-4", BaseURL: "http://custom"})
	if c.cfg.BaseURL != "http://custom" {
		t.Errorf("BaseURL = %q, want http://custom", c.cfg.BaseURL)
	}
	if c.cfg.Model != "gpt-4" {
		t.Errorf("Model = %q, want gpt-4", c.cfg.Model)
	}
}

// TestOpenAIGenerateTestSuccess 验证成功路径。
func TestOpenAIGenerateTestSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("auth header = %q", r.Header.Get("Authorization"))
		}
		fmt.Fprint(w, `{"choices":[{"message":{"content":"  func TestMutant(t *testing.T) {}  "}}]}`)
	}))
	defer srv.Close()

	c := NewOpenAIClient(Config{APIKey: "test-key", BaseURL: srv.URL})
	result, err := c.GenerateTest(sampleMutation(), "package main")
	if err != nil {
		t.Fatalf("GenerateTest: %v", err)
	}
	if result != "func TestMutant(t *testing.T) {}" {
		t.Errorf("result = %q", result)
	}
}

// TestOpenAIGenerateTestHTTPError 验证 HTTP 请求失败。
func TestOpenAIGenerateTestHTTPError(t *testing.T) {
	orig := httpDoFn
	httpDoFn = func(c *BaseClient, req *http.Request) (*http.Response, error) {
		return nil, errors.New("network unreachable")
	}
	defer func() { httpDoFn = orig }()

	c := NewOpenAIClient(Config{APIKey: "k"})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "network unreachable") {
		t.Errorf("err = %v, want network unreachable", err)
	}
}

// TestOpenAIGenerateTestAPIError 验证非 2xx 状态码。
func TestOpenAIGenerateTestAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid api key")
	}))
	defer srv.Close()

	c := NewOpenAIClient(Config{APIKey: "bad", BaseURL: srv.URL})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "401") {
		t.Errorf("err = %v, want 401", err)
	}
}

// TestOpenAIGenerateTestBadJSON 验证 JSON 解析失败。
func TestOpenAIGenerateTestBadJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "not json")
	}))
	defer srv.Close()

	c := NewOpenAIClient(Config{APIKey: "k", BaseURL: srv.URL})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "unmarshal") {
		t.Errorf("err = %v, want unmarshal error", err)
	}
}

// TestOpenAIGenerateTestEmptyChoices 验证空 choices 响应。
func TestOpenAIGenerateTestEmptyChoices(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(openAIResponse{})
	}))
	defer srv.Close()

	c := NewOpenAIClient(Config{APIKey: "k", BaseURL: srv.URL})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "no choices") {
		t.Errorf("err = %v, want no choices", err)
	}
}

// TestDoRequestNewRequestError 验证 http.NewRequest 失败路径。
func TestDoRequestNewRequestError(t *testing.T) {
	orig := newRequestFn
	newRequestFn = func(method, url string, body io.Reader) (*http.Request, error) {
		return nil, errors.New("new request failed")
	}
	defer func() { newRequestFn = orig }()

	c := NewOpenAIClient(Config{APIKey: "k"})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "create request") {
		t.Errorf("err = %v, want create request", err)
	}
}

// TestDoRequestReadAllError 验证 io.ReadAll 失败路径。
func TestDoRequestReadAllError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	orig := readAllFn
	readAllFn = func(r io.Reader) ([]byte, error) {
		return nil, errors.New("read failed")
	}
	defer func() { readAllFn = orig }()

	c := NewOpenAIClient(Config{APIKey: "k", BaseURL: srv.URL})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "read response") {
		t.Errorf("err = %v, want read response", err)
	}
}

// TestOpenAIMarshalError 验证 json.Marshal 失败路径。
func TestOpenAIMarshalError(t *testing.T) {
	orig := openAIMarshalFn
	openAIMarshalFn = func(v any) ([]byte, error) {
		return nil, errors.New("marshal failed")
	}
	defer func() { openAIMarshalFn = orig }()

	c := NewOpenAIClient(Config{APIKey: "k"})
	_, err := c.GenerateTest(sampleMutation(), "source")
	if err == nil || !strings.Contains(err.Error(), "marshal") {
		t.Errorf("err = %v, want marshal error", err)
	}
}

// TestBuildPrompt 验证提示词包含关键信息。
func TestBuildPrompt(t *testing.T) {
	p := buildPrompt("f.go", 5, "M1", "+", "-", "src")
	if !strings.Contains(p, "f.go") || !strings.Contains(p, "M1") || !strings.Contains(p, "src") {
		t.Error("prompt missing key info")
	}
}
