// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package mut

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/llm"
)

// 编译时接口检查。
var _ LLMClient = (*llm.OpenAIClient)(nil)
var _ LLMClient = (*llm.ClaudeClient)(nil)

// TestLLMIntegration_OpenAIWithFlywheel 验证 OpenAIClient 注入 Flywheel 后端到端工作。
func TestLLMIntegration_OpenAIWithFlywheel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"choices":[{"message":{"content":"func TestKill(t *testing.T) { t.Fail() }"}}]}`)
	}))
	defer srv.Close()

	client := llm.NewOpenAIClient(llm.Config{APIKey: "k", BaseURL: srv.URL})
	runner := &mockRunner{passMap: map[string]bool{"func TestKill(t *testing.T) { t.Fail() }": false}}
	f, err := NewFlywheel(client, runner, FlywheelConfig{MaxRetries: 1, MaxMutants: 10})
	if err != nil {
		t.Fatalf("NewFlywheel: %v", err)
	}

	m := core.Mutation{File: "f.go", Line: 1, Class: core.OpArithmetic, From: "+", To: "-"}
	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
	if !results[0].Killed {
		t.Error("mutant should be killed")
	}
}

// TestLLMIntegration_ClaudeWithFlywheel 验证 ClaudeClient 注入 Flywheel 后端到端工作。
func TestLLMIntegration_ClaudeWithFlywheel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"content":[{"type":"text","text":"func TestKill(t *testing.T) { t.Fail() }"}]}`)
	}))
	defer srv.Close()

	client := llm.NewClaudeClient(llm.Config{APIKey: "k", BaseURL: srv.URL})
	runner := &mockRunner{passMap: map[string]bool{"func TestKill(t *testing.T) { t.Fail() }": false}}
	f, err := NewFlywheel(client, runner, FlywheelConfig{MaxRetries: 1, MaxMutants: 10})
	if err != nil {
		t.Fatalf("NewFlywheel: %v", err)
	}

	m := core.Mutation{File: "f.go", Line: 1, Class: core.OpArithmetic, From: "+", To: "-"}
	results, err := f.Process([]core.Mutation{m}, "source")
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
	if !results[0].Killed {
		t.Error("mutant should be killed")
	}
}
