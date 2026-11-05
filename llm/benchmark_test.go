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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qingdujian/abate/core"
)

// BenchmarkBuildPrompt 基准测试：提示词构造。
func BenchmarkBuildPrompt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buildPrompt("main.go", 42, "M1", "+", "-", "package main\nfunc add(a,b int) int { return a+b }")
	}
}

// BenchmarkNewOpenAIClient 基准测试：OpenAI 客户端构造。
func BenchmarkNewOpenAIClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewOpenAIClient(Config{APIKey: "k", Model: "gpt-4o"})
	}
}

// BenchmarkNewClaudeClient 基准测试：Claude 客户端构造。
func BenchmarkNewClaudeClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewClaudeClient(Config{APIKey: "k", Model: "claude-sonnet-4-20250514"})
	}
}

// BenchmarkOpenAIGenerateTest 基准测试：OpenAI GenerateTest 端到端（mock server）。
func BenchmarkOpenAIGenerateTest(b *testing.B) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"choices":[{"message":{"content":"func TestKill(t *testing.T) { t.Fail() }"}}]}`)
	}))
	defer srv.Close()

	client := NewOpenAIClient(Config{APIKey: "k", BaseURL: srv.URL})
	mutant := core.Mutation{File: "f.go", Line: 1, Class: "M1", From: "+", To: "-"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.GenerateTest(mutant, "package main")
	}
}

// BenchmarkClaudeGenerateTest 基准测试：Claude GenerateTest 端到端（mock server）。
func BenchmarkClaudeGenerateTest(b *testing.B) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"content":[{"type":"text","text":"func TestKill(t *testing.T) { t.Fail() }"}]}`)
	}))
	defer srv.Close()

	client := NewClaudeClient(Config{APIKey: "k", BaseURL: srv.URL})
	mutant := core.Mutation{File: "f.go", Line: 1, Class: "M1", From: "+", To: "-"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.GenerateTest(mutant, "package main")
	}
}
