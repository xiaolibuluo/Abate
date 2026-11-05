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

// DefaultOpenAIURL 是 OpenAI API 默认端点。
const DefaultOpenAIURL = "https://api.openai.com/v1/chat/completions"

// openAIMarshalFn 可在测试中替换以触发序列化错误。
var openAIMarshalFn = json.Marshal

// OpenAIClient 实现 mut.LLMClient 接口，对接 OpenAI Chat Completions API。
type OpenAIClient struct {
	BaseClient
}

// NewOpenAIClient 构造 OpenAI 客户端。
func NewOpenAIClient(cfg Config) *OpenAIClient {
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultOpenAIURL
	}
	if cfg.Model == "" {
		cfg.Model = "gpt-4o"
	}
	return &OpenAIClient{BaseClient: BaseClient{cfg: cfg}}
}

// openAIRequest 是 Chat Completions 请求体。
type openAIRequest struct {
	Model    string          `json:"model"`
	Messages []openAIMessage `json:"messages"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// openAIResponse 是 Chat Completions 响应体。
type openAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// GenerateTest 实现 mut.LLMClient 接口。
func (c *OpenAIClient) GenerateTest(mutant core.Mutation, source string) (string, error) {
	prompt := buildPrompt(mutant.File, mutant.Line, string(mutant.Class), mutant.From, mutant.To, source)
	reqBody := openAIRequest{
		Model: c.cfg.Model,
		Messages: []openAIMessage{
			{Role: "user", Content: prompt},
		},
	}
	data, err := openAIMarshalFn(reqBody)
	if err != nil {
		return "", fmt.Errorf("openai: marshal request: %w", err)
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.cfg.APIKey,
	}
	respData, err := c.doRequest("POST", c.cfg.BaseURL, headers, data)
	if err != nil {
		return "", err
	}
	var resp openAIResponse
	if err := json.Unmarshal(respData, &resp); err != nil {
		return "", fmt.Errorf("openai: unmarshal response: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("openai: no choices in response")
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}
