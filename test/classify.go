// Abate —— 止息混沌 · 失败归因分类器
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import "strings"

// FailureKind 失败归因类别：代码问题、环境问题、瞬态问题。
type FailureKind string

const (
	FailureCode      FailureKind = "code"
	FailureEnv       FailureKind = "env"
	FailureTransient FailureKind = "transient"
)

// Metric 返回归因类别的可量化标识。
func (k FailureKind) Metric() float64 {
	switch k {
	case FailureCode:
		return 0
	case FailureEnv:
		return 1
	case FailureTransient:
		return 2
	}
	return 0
}

// Label 返回归因类别的人类可读标签。
func (k FailureKind) Label() string {
	switch k {
	case FailureEnv:
		return "env-issue"
	case FailureTransient:
		return "transient"
	default:
		return "code-issue"
	}
}

// envPatterns 环境问题特征：工具缺失、权限不足、网络不通、资源耗尽。
var envPatterns = []string{
	"command not found",
	"no such file or directory",
	"permission denied",
	"connection refused",
	"network is unreachable",
	"out of memory",
	"oom-kill",
	"exec format error",
	"is not installed",
	"sdk location not found",
	"gradle version",
	"java_home",
}

// transientPatterns 瞬态问题特征：可自动重试恢复。
var transientPatterns = []string{
	"resource temporarily unavailable",
	"text file busy",
	"i/o timeout",
	"broken pipe",
	"connection reset",
	"connection closed",
	"temporary failure",
	"eagain",
}

// ClassifyFailure 对失败结果归因：瞬态 > 环境 > 代码。
// 瞬态优先匹配，因为瞬态错误可能同时携带环境特征。
func ClassifyFailure(res TestResult) FailureKind {
	if res.Status != StatusFail {
		return ""
	}
	d := strings.ToLower(res.Details)
	for _, p := range transientPatterns {
		if strings.Contains(d, p) {
			return FailureTransient
		}
	}
	for _, p := range envPatterns {
		if strings.Contains(d, p) {
			return FailureEnv
		}
	}
	return FailureCode
}
