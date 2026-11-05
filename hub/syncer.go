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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DefaultHubURL 是 ChaosHub 数据源默认端点。
const DefaultHubURL = "https://hub.qingdujian.com/api/v1/subcodes"

// syncPayload 是 ChaosHub 数据源的 JSON 响应格式。
type syncPayload struct {
	Version string    `json:"version"`
	Codes   []SubCode `json:"codes"`
}

// Syncer 从 ChaosHub 数据源同步子编码表。
type Syncer struct {
	url    string
	apiKey string
	httpDo func(req *http.Request) (*http.Response, error)
}

// httpDoFn 可在测试中替换以避免真实网络请求。
var httpDoFn = func(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

// readAllFn 可在测试中替换以触发读取错误。
var readAllFn = io.ReadAll

// NewSyncer 构造同步器。
func NewSyncer(url, apiKey string) *Syncer {
	if url == "" {
		url = DefaultHubURL
	}
	return &Syncer{url: url, apiKey: apiKey}
}

// Sync 执行远程同步，返回拉取到的注册表。
func (s *Syncer) Sync() (*Registry, error) {
	req, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return nil, fmt.Errorf("hub: create request: %w", err)
	}
	if s.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+s.apiKey)
	}
	resp, err := httpDoFn(req)
	if err != nil {
		return nil, fmt.Errorf("hub: http request: %w", err)
	}
	defer resp.Body.Close()
	data, err := readAllFn(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("hub: read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("hub: API error %d: %s", resp.StatusCode, string(data))
	}
	var payload syncPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("hub: unmarshal response: %w", err)
	}
	reg := NewRegistry(payload.Version)
	for _, sc := range payload.Codes {
		reg.codes[sc.Code] = sc
	}
	return reg, nil
}
