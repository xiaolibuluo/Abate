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
	"fmt"
	"strings"

	"github.com/qingdujian/abate/core"
)

// DefaultClaudeURL 是 Anthropic API 默认端点。
const DefaultClaudeURL = "https://api.anthropic.com/v1/messages"

// claudeMarshalFn 可在测试中替换以触发序列化错误。
var claudeMarshalFn = json.Marshal

// ClaudeClient 实现 mut.LLMClient 接口，对接 Anthropic Messages API。
type ClaudeClient struct {
	BaseClient
}

// NewClaudeClient 构造 Claude 客户端。
func NewClaudeClient(cfg Config) *ClaudeClient {
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultClaudeURL
	}
	if cfg.Model == "" {
		cfg.Model = "claude-sonnet-4-20250514"
	}
	return &ClaudeClient{BaseClient: BaseClient{cfg: cfg}}
}

// claudeRequest 是 Anthropic Messages 请求体。
type claudeRequest struct {
	Model    string         `json:"model"`
	MaxTokens int           `json:"max_tokens"`
	Messages []claudeMessage `json:"messages"`
}

type claudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// claudeResponse 是 Anthropic Messages 响应体。
type claudeResponse struct {
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
}

// GenerateTest 实现 mut.LLMClient 接口。
func (c *ClaudeClient) GenerateTest(mutant core.Mutation, source, feedback string) (string, error) {
	prompt := buildPrompt(mutant.File, mutant.Line, string(mutant.Class), mutant.From, mutant.To, source, feedback)
	reqBody := claudeRequest{
		Model:     c.cfg.Model,
		MaxTokens: 4096,
		Messages: []claudeMessage{
			{Role: "user", Content: prompt},
		},
	}
	data, err := claudeMarshalFn(reqBody)
	if err != nil {
		return "", fmt.Errorf("claude: marshal request: %w", err)
	}
	headers := map[string]string{
		"Content-Type":      "application/json",
		"x-api-key":         c.cfg.APIKey,
		"anthropic-version": "2023-06-01",
	}
	respData, err := c.doRequest("POST", c.cfg.BaseURL, headers, data)
	if err != nil {
		return "", err
	}
	var resp claudeResponse
	if err := json.Unmarshal(respData, &resp); err != nil {
		return "", fmt.Errorf("claude: unmarshal response: %w", err)
	}
	if len(resp.Content) == 0 {
		return "", fmt.Errorf("claude: no content in response")
	}
	return strings.TrimSpace(resp.Content[0].Text), nil
}
