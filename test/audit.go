// Abate —— 止息混沌 · 全链路审计（严层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"fmt"
	"os"
	"strings"
)

// AuditFinding 是审计发现的问题。
type AuditFinding struct {
	File     string
	Line     int
	Category string // header / license / secret / todo / fixme
	Message  string
}

// AuditResult 是审计汇总。
type AuditResult struct {
	Total    int
	Headers  int
	Licenses int
	Secrets  int
	Todos    int
	Fixmes   int
	Findings []AuditFinding
}

// secretPatterns 敏感信息特征。
var secretPatterns = []string{
	"password =", "passwd =", "pwd =",
	"api_key =", "apikey =", "api-key =",
	"secret =", "token =", "access_key =",
	"private_key", "BEGIN RSA PRIVATE",
	"BEGIN PRIVATE KEY",
	"AKIA", // AWS access key prefix
}

// RunAudit 对项目执行全链路审计。
func RunAudit(target string, maxFiles int) AuditResult {
	files := collectCodeFiles(target, maxFiles)
	var result AuditResult
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		content := string(data)
		lines := strings.Split(content, "\n")
		// 文件头检查：前 5 行须含 Copyright 或 AGPL。
		if !hasHeader(lines) {
			result.Headers++
			result.Findings = append(result.Findings, AuditFinding{File: f, Category: "header", Message: "missing copyright/license header"})
		}
		// 逐行扫描敏感信息和 TODO/FIXME。
		for i, line := range lines {
			lower := strings.ToLower(line)
			for _, p := range secretPatterns {
				if strings.Contains(lower, p) {
					result.Secrets++
					result.Findings = append(result.Findings, AuditFinding{File: f, Line: i + 1, Category: "secret", Message: "potential secret: " + p})
					break
				}
			}
			if strings.Contains(line, "TODO") {
				result.Todos++
			}
			if strings.Contains(line, "FIXME") || strings.Contains(line, "HACK") || strings.Contains(line, "XXX") {
				result.Fixmes++
				result.Findings = append(result.Findings, AuditFinding{File: f, Line: i + 1, Category: "fixme", Message: strings.TrimSpace(line)})
			}
		}
	}
	result.Total = len(files)
	return result
}

// hasHeader 检查文件前几行是否含版权声明。
func hasHeader(lines []string) bool {
	n := 5
	if len(lines) < n {
		n = len(lines)
	}
	for i := 0; i < n; i++ {
		if strings.Contains(lines[i], "Copyright") || strings.Contains(lines[i], "AGPL") {
			return true
		}
	}
	return false
}

// FormatAudit 将审计结果格式化为摘要文本。
func FormatAudit(r AuditResult, limit int) string {
	if r.Total == 0 {
		return "no files audited"
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("files=%d headers_missing=%d secrets=%d todos=%d fixmes=%d",
		r.Total, r.Headers, r.Secrets, r.Todos, r.Fixmes))
	if len(r.Findings) > 0 {
		b.WriteString(" | ")
		for i, fnd := range r.Findings {
			if i >= limit {
				b.WriteString(fmt.Sprintf("... +%d more", len(r.Findings)-limit))
				break
			}
			if i > 0 {
				b.WriteString("; ")
			}
			b.WriteString(fmt.Sprintf("%s:%d[%s]", shortFile("", fnd.File), fnd.Line, fnd.Category))
		}
	}
	return b.String()
}
