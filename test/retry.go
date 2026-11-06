// Abate —— 止息混沌 · 瞬态失败重试策略
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"time"
)

// RetryPolicy 瞬态失败重试策略。
type RetryPolicy struct {
	// MaxAttempts 最大尝试次数（含首次执行）。
	MaxAttempts int
	// Backoff 每次重试前的退避时长。
	Backoff time.Duration
}

// DefaultRetryPolicy 默认策略：最多 2 次尝试，500ms 退避。
var DefaultRetryPolicy = RetryPolicy{MaxAttempts: 2, Backoff: 500 * time.Millisecond}

// runFn 单次执行的抽象，便于测试替换。
type runFn func(ctx context.Context) TestResult

// WithRetry 按策略执行：仅瞬态归因的失败才重试，其余直接返回。
// 重试成功时在 Metrics 标注 retried=1，保留最终结果。
func WithRetry(ctx context.Context, policy RetryPolicy, fn runFn) TestResult {
	if policy.MaxAttempts < 1 {
		policy.MaxAttempts = 1
	}
	var res TestResult
	for attempt := 1; attempt <= policy.MaxAttempts; attempt++ {
		res = fn(ctx)
		if res.Status != StatusFail {
			return res
		}
		if ClassifyFailure(res) != FailureTransient {
			return res
		}
		if attempt == policy.MaxAttempts {
			break
		}
		select {
		case <-ctx.Done():
			return res
		case <-time.After(policy.Backoff):
		}
	}
	return res
}

// AnnotateClassification 为结果补充归因标注：失败结果写入 failure_kind 指标。
func AnnotateClassification(res TestResult) TestResult {
	if res.Status != StatusFail {
		return res
	}
	kind := ClassifyFailure(res)
	if res.Metrics == nil {
		res.Metrics = map[string]float64{}
	}
	res.Metrics["failure_kind"] = kind.Metric()
	if res.Details != "" {
		res.Details += " "
	}
	res.Details += "cause=" + kind.Label()
	return res
}
