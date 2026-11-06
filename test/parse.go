// Abate —— 止息混沌 · 多语言工具输出解析器（准层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Severity 问题严重级别。
type Severity string

const (
	SevError   Severity = "error"
	SevWarning Severity = "warning"
	SevInfo    Severity = "info"
)

// Metric 返回严重级别的可量化标识。
func (s Severity) Metric() float64 {
	switch s {
	case SevError:
		return 3
	case SevWarning:
		return 2
	default:
		return 1
	}
}

// Issue 是从工具输出中提取的结构化问题。
type Issue struct {
	File     string
	Line     int
	Severity Severity
	Message  string
	Tool     string
}

// Location 返回 文件:行号 形式的定位串。
func (i Issue) Location() string {
	if i.Line > 0 {
		return fmt.Sprintf("%s:%d", i.File, i.Line)
	}
	return i.File
}

// issuePatterns 各语言工具的通用输出模式：file:line[:col]: [severity:] message
var issuePatterns = []*regexp.Regexp{
	// go vet / golangci: file.go:12:3: message
	regexp.MustCompile(`^(\S+\.go):(\d+)(?::\d+)?:\s*(?:fatal:)?\s*(?:error:)?\s*(.+)$`),
	// gradle lint / android lint: file.kt:12: Error: message
	regexp.MustCompile(`^(\S+\.(?:kt|java)):(\d+):\s*(Error|Warning|Information):\s*(.+)$`),
	// eslint / tsc: /path/file.js:12:5: error: message
	regexp.MustCompile(`^(\S+\.(?:js|ts|tsx|jsx)):(\d+):\d+:\s*(error|warning)?:?\s*(.+)$`),
	// clippy / rustc: file.rs:12:5: message [lint]
	regexp.MustCompile(`^(\S+\.rs):(\d+):\d+:\s*(error|warning)?:?\s*(.+)$`),
	// python: file.py:12:30: E501 message
	regexp.MustCompile(`^(\S+\.py):(\d+):\d+:\s*([EWF]\d+)\s*(.+)$`),
	// 通用兜底: file.ext:12: message
	regexp.MustCompile(`^(\S+\.[a-zA-Z0-9]+):(\d+):\s*(.+)$`),
}

// ParseIssues 从工具输出中提取结构化问题列表。
// tool 标注来源工具，severityHint 提供无法从行内判断级别时的默认级别。
func ParseIssues(output, tool string, severityHint Severity) []Issue {
	var issues []Issue
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			continue
		}
		if iss, ok := parseIssueLine(line, tool, severityHint); ok {
			issues = append(issues, iss)
		}
	}
	return issues
}

func parseIssueLine(line, tool string, hint Severity) (Issue, bool) {
	for _, re := range issuePatterns {
		m := re.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		file := m[1]
		if !IsCodeFile(file) {
			continue
		}
		ln, err := strconv.Atoi(m[2])
		if err != nil || ln <= 0 {
			continue
		}
		iss := Issue{File: filepath.Clean(file), Line: ln, Tool: tool, Severity: hint}
		// 模式分组：部分含显式级别，部分不含。
		switch len(m) {
		case 5:
			iss.Severity = normalizeSeverity(m[3])
			iss.Message = strings.TrimSpace(m[4])
		case 4:
			iss.Message = strings.TrimSpace(m[3])
		default:
			continue
		}
		if iss.Message == "" {
			continue
		}
		return iss, true
	}
	return Issue{}, false
}

func normalizeSeverity(s string) Severity {
	switch strings.ToLower(s) {
	case "error", "fatal", "e":
		return SevError
	case "warning", "w":
		return SevWarning
	default:
		// python 风格代码：E 系 error，W/F 系 warning，其余 info。
		if len(s) == 4 {
			switch s[0] {
			case 'E', 'F':
				return SevError
			case 'W':
				return SevWarning
			}
		}
		return SevInfo
	}
}

// IssueStats 按严重级别统计问题数量。
type IssueStats struct {
	Errors   int
	Warnings int
	Infos    int
	Total    int
}

// CountIssues 统计问题列表。
func CountIssues(issues []Issue) IssueStats {
	var st IssueStats
	for _, iss := range issues {
		st.Total++
		switch iss.Severity {
		case SevError:
			st.Errors++
		case SevWarning:
			st.Warnings++
		default:
			st.Infos++
		}
	}
	return st
}

// FormatIssues 将问题列表格式化为最多 limit 行的摘要文本。
func FormatIssues(issues []Issue, limit int) string {
	if len(issues) == 0 {
		return ""
	}
	if limit <= 0 {
		limit = 5
	}
	var b strings.Builder
	for i, iss := range issues {
		if i >= limit {
			fmt.Fprintf(&b, "... +%d more", len(issues)-limit)
			break
		}
		if i > 0 {
			b.WriteString("; ")
		}
		fmt.Fprintf(&b, "%s[%s] %s", iss.Location(), iss.Severity, truncate(iss.Message, 60))
	}
	return b.String()
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
