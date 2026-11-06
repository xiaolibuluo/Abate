//go:build matrix

// Abate —— 止息混沌 · 全矩阵测试入口
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权
//
// 仅在显式指定 -tags=matrix 时运行：go test ./test/ -tags=matrix -run TestMatrix -v -timeout 8m

package test

import (
	"context"
	"testing"
	"time"
)

func TestMatrix(t *testing.T) {
	runner := NewRunner(DefaultRegistry, 60*time.Second)
	report := runner.RunMatrix(context.Background(), "/root/abate")
	for _, res := range report.Results {
		icon := "PASS"
		if res.Status == StatusFail {
			icon = "FAIL"
		}
		if res.Status == StatusSkip {
			icon = "SKIP"
		}
		t.Logf("%s %-24s %8v  %s", icon, res.Name, res.Duration, res.Details)
	}
	t.Logf("\nTotal: %d  Pass: %d  Fail: %d  Skip: %d", report.Total, report.Passed, report.Failed, report.Skipped)
	if report.Failed > 0 {
		t.Errorf("matrix had %d failures out of %d total", report.Failed, report.Total)
	}
}
