// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package llm 提供 LLM 客户端实现，对接 OpenAI 和 Anthropic Claude API。
// Package llm provides LLM client implementations for OpenAI and Anthropic Claude APIs.
package llm

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Config 是 LLM 客户端公共配置。
type Config struct {
	APIKey  string
	Model   string
	BaseURL string
}

// BaseClient 提供公共 HTTP 请求能力。
type BaseClient struct {
	cfg    Config
	httpDo func(req *http.Request) (*http.Response, error)
}

// httpDoFn 可在测试中替换以避免真实网络请求。
var httpDoFn = func(c *BaseClient, req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

// newRequestFn 可在测试中替换以触发请求构造错误。
var newRequestFn = http.NewRequest

// readAllFn 可在测试中替换以触发响应读取错误。
var readAllFn = io.ReadAll

// doRequest 发送 HTTP 请求并返回响应体字节。
func (c *BaseClient) doRequest(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	req, err := newRequestFn(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("llm: create request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := httpDoFn(c, req)
	if err != nil {
		return nil, fmt.Errorf("llm: http request: %w", err)
	}
	defer resp.Body.Close()
	data, err := readAllFn(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("llm: read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("llm: API error %d: %s", resp.StatusCode, string(data))
	}
	return data, nil
}

// buildPrompt 根据变异体和源码构造补测生成提示词。
func buildPrompt(file string, line int, class, from, to, source string) string {
	return fmt.Sprintf(
		"You are a Go test expert. A mutation was applied to the following Go source file.\n\n"+
			"File: %s\nLine: %d\nMutation class: %s\nOriginal: %q\nMutated: %q\n\n"+
			"Source code:\n```go\n%s\n```\n\n"+
			"Write a Go test that would FAIL against the mutated code but PASS against the original. "+
			"Return ONLY the test code, no explanation.",
		file, line, class, from, to, source,
	)
}
