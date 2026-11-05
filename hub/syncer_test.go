// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package hub

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestNewSyncerDefaults 验证默认 URL。
func TestNewSyncerDefaults(t *testing.T) {
	s := NewSyncer("", "")
	if s.url != DefaultHubURL {
		t.Errorf("url = %q, want %q", s.url, DefaultHubURL)
	}
}

// TestSyncerSuccess 验证成功同步。
func TestSyncerSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("auth = %q", r.Header.Get("Authorization"))
		}
		fmt.Fprint(w, `{"version":"v0.2","codes":[{"code":"F1.1","name":"延迟","params":"delayMs"}]}`)
	}))
	defer srv.Close()

	s := NewSyncer(srv.URL, "test-key")
	reg, err := s.Sync()
	if err != nil {
		t.Fatalf("Sync: %v", err)
	}
	if reg.Version() != "v0.2" {
		t.Errorf("version = %q, want v0.2", reg.Version())
	}
	if reg.Count() != 1 {
		t.Errorf("count = %d, want 1", reg.Count())
	}
	sc, _ := reg.Lookup("F1.1")
	if sc.Name != "延迟" {
		t.Errorf("name = %q", sc.Name)
	}
}

// TestSyncerHTTPError 验证 HTTP 请求失败。
func TestSyncerHTTPError(t *testing.T) {
	orig := httpDoFn
	httpDoFn = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("network down")
	}
	defer func() { httpDoFn = orig }()

	s := NewSyncer("", "")
	_, err := s.Sync()
	if err == nil || !strings.Contains(err.Error(), "network down") {
		t.Errorf("err = %v, want network down", err)
	}
}

// TestSyncerAPIError 验证非 2xx 状态码。
func TestSyncerAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "unauthorized")
	}))
	defer srv.Close()

	s := NewSyncer(srv.URL, "bad")
	_, err := s.Sync()
	if err == nil || !strings.Contains(err.Error(), "401") {
		t.Errorf("err = %v, want 401", err)
	}
}

// TestSyncerBadJSON 验证 JSON 解析失败。
func TestSyncerBadJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "not json")
	}))
	defer srv.Close()

	s := NewSyncer(srv.URL, "")
	_, err := s.Sync()
	if err == nil || !strings.Contains(err.Error(), "unmarshal") {
		t.Errorf("err = %v, want unmarshal", err)
	}
}

// TestSyncerReadAllError 验证响应读取失败。
func TestSyncerReadAllError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	orig := readAllFn
	readAllFn = func(r io.Reader) ([]byte, error) {
		return nil, errors.New("read failed")
	}
	defer func() { readAllFn = orig }()

	s := NewSyncer(srv.URL, "")
	_, err := s.Sync()
	if err == nil || !strings.Contains(err.Error(), "read response") {
		t.Errorf("err = %v, want read response", err)
	}
}

// TestSyncerNewRequestError 验证请求构造失败。
func TestSyncerNewRequestError(t *testing.T) {
	s := NewSyncer("http://[::1]:namedport", "")
	_, err := s.Sync()
	if err == nil || !strings.Contains(err.Error(), "create request") {
		t.Errorf("err = %v, want create request", err)
	}
}
