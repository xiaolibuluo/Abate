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
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func runMCP(t *testing.T, input string) string {
	var out bytes.Buffer
	handler := &DefaultToolHandler{}
	srv := NewMCPServer(strings.NewReader(input), &out, handler)
	_ = srv.Serve()
	return out.String()
}

func parseResponses(t *testing.T, output string) []JSONRPCResponse {
	var responses []JSONRPCResponse
	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		if line == "" {
			continue
		}
		var resp JSONRPCResponse
		if err := json.Unmarshal([]byte(line), &resp); err != nil {
			t.Fatalf("parse response: %v", err)
		}
		responses = append(responses, resp)
	}
	return responses
}

func TestMCP_Initialize(t *testing.T) {
	out := runMCP(t, `{"jsonrpc":"2.0","id":1,"method":"initialize"}`+"\n")
	resps := parseResponses(t, out)
	if len(resps) != 1 {
		t.Fatalf("responses = %d", len(resps))
	}
	result := resps[0].Result.(map[string]any)
	if result["protocolVersion"] != "2024-11-05" {
		t.Errorf("protocol = %v", result["protocolVersion"])
	}
}

func TestMCP_ToolsList(t *testing.T) {
	out := runMCP(t, `{"jsonrpc":"2.0","id":2,"method":"tools/list"}`+"\n")
	resps := parseResponses(t, out)
	result := resps[0].Result.(map[string]any)
	tools := result["tools"].([]any)
	if len(tools) != 4 {
		t.Errorf("tools = %d, want 4", len(tools))
	}
}

func TestMCP_ToolsCall_GenerateMutants(t *testing.T) {
	args := `{"filename":"test.go","source":"package main\nfunc add(a,b int) int { return a+b }","seed":42}`
	req := `{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"generate_mutants","arguments":` + args + `}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error != nil {
		t.Fatalf("error: %v", resps[0].Error)
	}
	result := resps[0].Result.(map[string]any)
	content := result["content"].([]any)
	if len(content) != 1 {
		t.Fatalf("content = %d", len(content))
	}
}

func TestMCP_ToolsCall_ListFaultClasses(t *testing.T) {
	req := `{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"list_fault_classes","arguments":{}}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error != nil {
		t.Fatalf("error: %v", resps[0].Error)
	}
}

func TestMCP_ToolsCall_UnknownTool(t *testing.T) {
	req := `{"jsonrpc":"2.0","id":5,"method":"tools/call","params":{"name":"nonexistent","arguments":{}}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error == nil {
		t.Fatal("expected error for unknown tool")
	}
}

func TestMCP_ToolsCall_InvalidParams(t *testing.T) {
	req := `{"jsonrpc":"2.0","id":6,"method":"tools/call","params":"invalid"}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error == nil {
		t.Fatal("expected error for invalid params")
	}
}

func TestMCP_MethodNotFound(t *testing.T) {
	out := runMCP(t, `{"jsonrpc":"2.0","id":7,"method":"unknown/method"}`+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error == nil {
		t.Fatal("expected error for unknown method")
	}
}

func TestMCP_ParseError(t *testing.T) {
	out := runMCP(t, `not json`+"\n")
	resps := parseResponses(t, out)
	if len(resps) != 1 || resps[0].Error == nil {
		t.Fatal("expected parse error")
	}
}

func TestMCP_EmptyLine(t *testing.T) {
	out := runMCP(t, "\n\n"+`{"jsonrpc":"2.0","id":8,"method":"initialize"}`+"\n")
	resps := parseResponses(t, out)
	if len(resps) != 1 {
		t.Fatalf("responses = %d, want 1 (empty lines skipped)", len(resps))
	}
}

func TestMCP_ToolsCall_GenerateMutants_InvalidArgs(t *testing.T) {
	req := `{"jsonrpc":"2.0","id":9,"method":"tools/call","params":{"name":"generate_mutants","arguments":"invalid"}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error == nil {
		t.Fatal("expected error for invalid args")
	}
}

func TestMCP_ToolsCall_GenerateMutants_BadSource(t *testing.T) {
	args := `{"filename":"test.go","source":"not go code!!!","seed":42}`
	req := `{"jsonrpc":"2.0","id":10,"method":"tools/call","params":{"name":"generate_mutants","arguments":` + args + `}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error == nil {
		t.Fatal("expected error for bad source")
	}
}

func TestMCP_ToolsCall_ApplyMutant(t *testing.T) {
	args := `{"source":"package main\n\nfunc add(a,b int) int { return a+b }","mutant":{"file":"test.go","line":3,"col":33,"class":"M1","from":"+","to":"-"}}`
	req := `{"jsonrpc":"2.0","id":11,"method":"tools/call","params":{"name":"apply_mutant","arguments":` + args + `}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error != nil {
		t.Fatalf("error: %v", resps[0].Error)
	}
}

func TestMCP_ToolsCall_ApplyMutant_InvalidArgs(t *testing.T) {
	req := `{"jsonrpc":"2.0","id":12,"method":"tools/call","params":{"name":"apply_mutant","arguments":"invalid"}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error == nil {
		t.Fatal("expected error for invalid args")
	}
}

func TestMCP_ToolsCall_ApplyMutant_BadMutant(t *testing.T) {
	args := `{"source":"package main\nfunc f() {}","mutant":{"file":"","line":0,"col":0,"class":"","from":"","to":""}}`
	req := `{"jsonrpc":"2.0","id":13,"method":"tools/call","params":{"name":"apply_mutant","arguments":` + args + `}}`
	out := runMCP(t, req+"\n")
	resps := parseResponses(t, out)
	if resps[0].Error == nil {
		t.Fatal("expected error for bad mutant")
	}
}

func TestMCP_EOF(t *testing.T) {
	var out bytes.Buffer
	handler := &DefaultToolHandler{}
	srv := NewMCPServer(strings.NewReader(""), &out, handler)
	if err := srv.Serve(); err != nil {
		t.Fatalf("Serve with EOF: %v", err)
	}
	if out.Len() != 0 {
		t.Error("output should be empty on EOF")
	}
}

type errorReader struct{}

func (errorReader) Read(_ []byte) (int, error) {
	return 0, bytes.ErrTooLarge
}

func TestMCP_ReadError(t *testing.T) {
	var out bytes.Buffer
	handler := &DefaultToolHandler{}
	srv := NewMCPServer(errorReader{}, &out, handler)
	if err := srv.Serve(); err == nil {
		t.Fatal("expected read error")
	}
}
