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
	"testing"
)

// TestHubSubcodesHandler 验证 /api/v1/hub/subcodes 端点返回种子注册表。
func TestHubSubcodesHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hub/subcodes", nil)
	rec := httptest.NewRecorder()
	hubSubcodesHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}

	var body struct {
		Version string `json:"version"`
		Count   int    `json:"count"`
		Codes   []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Params string `json:"params"`
		} `json:"codes"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if body.Version != "v0.1" {
		t.Errorf("version = %q, want v0.1", body.Version)
	}
	if body.Count != 20 {
		t.Errorf("count = %d, want 20", body.Count)
	}
	if len(body.Codes) != 20 {
		t.Errorf("codes len = %d, want 20", len(body.Codes))
	}
	if body.Codes[0].Code != "F1.1" {
		t.Errorf("first code = %q, want F1.1", body.Codes[0].Code)
	}
}
