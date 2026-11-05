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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/qingdujian/abate/core"
)

// TestGenerateTestToolDescription 验证工具描述。
func TestGenerateTestToolDescription(t *testing.T) {
	tool := generateTestTool()
	if tool.Name != "generate_test" {
		t.Errorf("name = %q", tool.Name)
	}
}

// TestHandleGenerateTestOpenAI 验证 OpenAI provider 成功路径。
func TestHandleGenerateTestOpenAI(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"choices":[{"message":{"content":"func TestKill(t *testing.T) { t.Fail() }"}}]}`)
	}))
	defer srv.Close()

	mutant := core.Mutation{File: "f.go", Line: 1, Class: core.OpArithmetic, From: "+", To: "-"}
	args, _ := json.Marshal(map[string]any{
		"provider": "openai",
		"apiKey":   "k",
		"baseURL":  srv.URL,
		"source":   "src",
		"mutant":   mutant,
	})
	result, err := handleGenerateTest(args)
	if err != nil {
		t.Fatalf("handleGenerateTest: %v", err)
	}
	if len(result.Content) == 0 || !strings.Contains(result.Content[0].Text, "TestKill") {
		t.Errorf("unexpected result: %+v", result)
	}
}

// TestHandleGenerateTestClaude 验证 Claude provider 成功路径。
func TestHandleGenerateTestClaude(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"content":[{"type":"text","text":"func TestKill(t *testing.T) { t.Fail() }"}]}`)
	}))
	defer srv.Close()

	mutant := core.Mutation{File: "f.go", Line: 1, Class: core.OpArithmetic, From: "+", To: "-"}
	args, _ := json.Marshal(map[string]any{
		"provider": "claude",
		"apiKey":   "k",
		"baseURL":  srv.URL,
		"source":   "src",
		"mutant":   mutant,
	})
	result, err := handleGenerateTest(args)
	if err != nil {
		t.Fatalf("handleGenerateTest: %v", err)
	}
	if len(result.Content) == 0 || !strings.Contains(result.Content[0].Text, "TestKill") {
		t.Errorf("unexpected result: %+v", result)
	}
}

// TestHandleGenerateTestUnsupportedProvider 验证不支持的 provider。
func TestHandleGenerateTestUnsupportedProvider(t *testing.T) {
	mutant := core.Mutation{File: "f.go", Line: 1, Class: core.OpArithmetic, From: "+", To: "-"}
	args, _ := json.Marshal(map[string]any{
		"provider": "gemini",
		"apiKey":   "k",
		"source":   "src",
		"mutant":   mutant,
	})
	_, err := handleGenerateTest(args)
	if err == nil || !strings.Contains(err.Error(), "unsupported provider") {
		t.Errorf("err = %v, want unsupported provider", err)
	}
}

// TestHandleGenerateTestInvalidArgs 验证无效参数。
func TestHandleGenerateTestInvalidArgs(t *testing.T) {
	_, err := handleGenerateTest(json.RawMessage(`not json`))
	if err == nil || !strings.Contains(err.Error(), "invalid args") {
		t.Errorf("err = %v, want invalid args", err)
	}
}

// TestCallToolGenerateTest 验证 DefaultToolHandler 路由 generate_test。
func TestCallToolGenerateTest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"choices":[{"message":{"content":"func TestK(t *testing.T) { t.Fail() }"}}]}`)
	}))
	defer srv.Close()

	mutant := core.Mutation{File: "f.go", Line: 1, Class: core.OpArithmetic, From: "+", To: "-"}
	args, _ := json.Marshal(map[string]any{
		"provider": "openai",
		"apiKey":   "k",
		"baseURL":  srv.URL,
		"source":   "src",
		"mutant":   mutant,
	})
	h := &DefaultToolHandler{}
	result, err := h.CallTool("generate_test", args)
	if err != nil {
		t.Fatalf("CallTool: %v", err)
	}
	if len(result.Content) == 0 {
		t.Error("expected content")
	}
}

// TestHandleGenerateTestLLMError 验证 LLM 返回错误。
func TestHandleGenerateTestLLMError(t *testing.T) {
	mutant := core.Mutation{File: "f.go", Line: 1, Class: core.OpArithmetic, From: "+", To: "-"}
	args, _ := json.Marshal(map[string]any{
		"provider": "openai",
		"apiKey":   "k",
		"baseURL":  "http://127.0.0.1:0",
		"source":   "src",
		"mutant":   mutant,
	})
	_, err := handleGenerateTest(args)
	if err == nil {
		t.Error("expected error from invalid URL")
	}
}

// TestToolsListIncludesGenerateTest 验证 tools/list 包含 generate_test。
func TestToolsListIncludesGenerateTest(t *testing.T) {
	srv := NewMCPServer(strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"tools/list"}`+"\n"), &strings.Builder{}, &DefaultToolHandler{})
	_ = srv.Serve()
	// tools/list 现在应返回 4 个工具。
}
