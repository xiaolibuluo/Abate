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
	"fmt"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/llm"
	"github.com/qingdujian/abate/mut"
)

// generateTestTool 返回 generate_test 工具描述。
func generateTestTool() MCPTool {
	return MCPTool{
		Name:        "generate_test",
		Description: "Generate a Go test to kill a survived mutant using an LLM client",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"provider": map[string]any{"type": "string", "description": "LLM provider: openai or claude"},
				"apiKey":   map[string]any{"type": "string", "description": "API key"},
				"model":    map[string]any{"type": "string", "description": "Model name (optional)"},
				"baseURL":  map[string]any{"type": "string", "description": "Custom API base URL (optional)"},
				"source":   map[string]any{"type": "string", "description": "Original Go source code"},
				"mutant":   map[string]any{"type": "object", "description": "Mutation to kill"},
			},
			"required": []string{"provider", "apiKey", "source", "mutant"},
		},
	}
}

type generateTestArgs struct {
	Provider string         `json:"provider"`
	APIKey   string         `json:"apiKey"`
	Model    string         `json:"model"`
	BaseURL  string         `json:"baseURL"`
	Source   string         `json:"source"`
	Mutant   core.Mutation  `json:"mutant"`
}

// handleGenerateTest 处理 generate_test 工具调用。
func handleGenerateTest(args json.RawMessage) (MCPResult, error) {
	var a generateTestArgs
	if err := json.Unmarshal(args, &a); err != nil {
		return MCPResult{}, fmt.Errorf("invalid args: %w", err)
	}
	var client mut.LLMClient
	cfg := llm.Config{APIKey: a.APIKey, Model: a.Model, BaseURL: a.BaseURL}
	switch a.Provider {
	case "openai":
		client = llm.NewOpenAIClient(cfg)
	case "claude":
		client = llm.NewClaudeClient(cfg)
	default:
		return MCPResult{}, fmt.Errorf("unsupported provider: %s", a.Provider)
	}
	testCode, err := client.GenerateTest(a.Mutant, a.Source)
	if err != nil {
		return MCPResult{}, err
	}
	return MCPResult{Content: []MCPContent{{Type: "text", Text: testCode}}}, nil
}
