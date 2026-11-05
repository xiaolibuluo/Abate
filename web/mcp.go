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
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/qingdujian/abate/core"
	"github.com/qingdujian/abate/mut"
)

// MCPTool 描述一个 MCP 工具。
type MCPTool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema any    `json:"inputSchema"`
}

// MCPCallParams 是 tools/call 请求参数。
type MCPCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

// MCPContent 是工具返回内容块。
type MCPContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// MCPResult 是 tools/call 返回结果。
type MCPResult struct {
	Content []MCPContent `json:"content"`
}

// JSONRPCRequest 是 JSON-RPC 2.0 请求。
type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

// JSONRPCResponse 是 JSON-RPC 2.0 响应。
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result,omitempty"`
	Error   *JSONRPCError   `json:"error,omitempty"`
}

// JSONRPCError 是 JSON-RPC 2.0 错误。
type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// MCPServer 是 Abate MCP Server，通过 stdio 暴露变异测试能力。
// MCPServer is the Abate MCP Server exposing mutation testing over stdio.
type MCPServer struct {
	reader  *bufio.Reader
	writer  io.Writer
	tools   []MCPTool
	handler MCPToolHandler
}

// MCPToolHandler 处理工具调用。
type MCPToolHandler interface {
	CallTool(name string, args json.RawMessage) (MCPResult, error)
}

// NewMCPServer 构造 MCP Server。
func NewMCPServer(r io.Reader, w io.Writer, handler MCPToolHandler) *MCPServer {
	return &MCPServer{
		reader:  bufio.NewReader(r),
		writer:  w,
		handler: handler,
		tools:   defaultMCPTools(),
	}
}

// defaultMCPTools 返回 Abate 暴露的 MCP 工具列表。
func defaultMCPTools() []MCPTool {
	return []MCPTool{
		{
			Name:        "generate_mutants",
			Description: "Generate Go source mutants (M1+M2) from source code with deterministic seed",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"filename": map[string]any{"type": "string", "description": "Source file name"},
					"source":   map[string]any{"type": "string", "description": "Go source code"},
					"seed":     map[string]any{"type": "integer", "description": "Deterministic seed"},
				},
				"required": []string{"filename", "source", "seed"},
			},
		},
		{
			Name:        "apply_mutant",
			Description: "Apply a mutant to source code and return the mutated source",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"source":  map[string]any{"type": "string", "description": "Original Go source code"},
					"mutant":  map[string]any{"type": "object", "description": "Mutation to apply"},
				},
				"required": []string{"source", "mutant"},
			},
		},
		{
			Name:        "list_fault_classes",
			Description: "List all CES fault classes (F1-F6)",
			InputSchema: map[string]any{"type": "object", "properties": map[string]any{}},
		},
		generateTestTool(),
	}
}

// Serve 启动 MCP Server 主循环，读取 JSON-RPC 请求并响应。
func (s *MCPServer) Serve() error {
	for {
		line, err := s.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var req JSONRPCRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			s.writeError(nil, -32700, "Parse error")
			continue
		}
		s.handleRequest(&req)
	}
}

// handleRequest 处理单个 JSON-RPC 请求。
func (s *MCPServer) handleRequest(req *JSONRPCRequest) {
	switch req.Method {
	case "initialize":
		s.writeResult(req.ID, map[string]any{
			"protocolVersion": "2024-11-05",
			"capabilities":   map[string]any{"tools": map[string]any{}},
			"serverInfo":     map[string]any{"name": "abate-mcp", "version": Version},
		})
	case "tools/list":
		s.writeResult(req.ID, map[string]any{"tools": s.tools})
	case "tools/call":
		var params MCPCallParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			s.writeError(req.ID, -32602, "Invalid params")
			return
		}
		result, err := s.handler.CallTool(params.Name, params.Arguments)
		if err != nil {
			s.writeError(req.ID, -32603, err.Error())
			return
		}
		s.writeResult(req.ID, result)
	default:
		s.writeError(req.ID, -32601, fmt.Sprintf("Method not found: %s", req.Method))
	}
}

// writeResult 写入成功响应。
func (s *MCPServer) writeResult(id json.RawMessage, result any) {
	resp := JSONRPCResponse{JSONRPC: "2.0", ID: id, Result: result}
	data, _ := json.Marshal(resp)
	fmt.Fprintf(s.writer, "%s\n", data)
}

// writeError 写入错误响应。
func (s *MCPServer) writeError(id json.RawMessage, code int, msg string) {
	resp := JSONRPCResponse{
		JSONRPC: "2.0", ID: id,
		Error: &JSONRPCError{Code: code, Message: msg},
	}
	data, _ := json.Marshal(resp)
	fmt.Fprintf(s.writer, "%s\n", data)
}

// DefaultToolHandler 是默认工具处理器，实现 MCPToolHandler 接口。
type DefaultToolHandler struct{}

// CallTool 处理工具调用。
func (h *DefaultToolHandler) CallTool(name string, args json.RawMessage) (MCPResult, error) {
	switch name {
	case "generate_mutants":
		return h.handleGenerateMutants(args)
	case "apply_mutant":
		return h.handleApplyMutant(args)
	case "list_fault_classes":
		return h.handleListFaultClasses()
	case "generate_test":
		return handleGenerateTest(args)
	default:
		return MCPResult{}, fmt.Errorf("unknown tool: %s", name)
	}
}

type generateMutantsArgs struct {
	Filename string `json:"filename"`
	Source   string `json:"source"`
	Seed     int64  `json:"seed"`
}

func (h *DefaultToolHandler) handleGenerateMutants(args json.RawMessage) (MCPResult, error) {
	var a generateMutantsArgs
	if err := json.Unmarshal(args, &a); err != nil {
		return MCPResult{}, fmt.Errorf("invalid args: %w", err)
	}
	mutants, err := mut.Generate(a.Filename, a.Source, a.Seed)
	if err != nil {
		return MCPResult{}, err
	}
	data, _ := json.MarshalIndent(mutants, "", "  ")
	return MCPResult{Content: []MCPContent{{Type: "text", Text: string(data)}}}, nil
}

type applyMutantArgs struct {
	Source string         `json:"source"`
	Mutant core.Mutation  `json:"mutant"`
}

func (h *DefaultToolHandler) handleApplyMutant(args json.RawMessage) (MCPResult, error) {
	var a applyMutantArgs
	if err := json.Unmarshal(args, &a); err != nil {
		return MCPResult{}, fmt.Errorf("invalid args: %w", err)
	}
	result, err := mut.Apply(a.Source, a.Mutant)
	if err != nil {
		return MCPResult{}, err
	}
	return MCPResult{Content: []MCPContent{{Type: "text", Text: result}}}, nil
}

func (h *DefaultToolHandler) handleListFaultClasses() (MCPResult, error) {
	classes := core.FaultClasses()
	data, _ := json.MarshalIndent(classes, "", "  ")
	return MCPResult{Content: []MCPContent{{Type: "text", Text: string(data)}}}, nil
}
