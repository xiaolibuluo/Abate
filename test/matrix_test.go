//go:build matrix

// Abate —— 止息混沌 · 全矩阵测试入口
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权
//
// 仅在显式指定 -tags=matrix 时运行：go test ./test/ -tags=matrix -run TestMatrix -v -timeout 10m

package test

import (
	"context"
	"testing"
	"time"
)

func TestMatrix(t *testing.T) {
	runner := NewRunner(DefaultRegistry, 5*time.Minute)
	report := runner.RunMatrix(context.Background(), "..")
	t.Log("\n" + report.FormatText())
	if report.Failed > 0 {
		t.Errorf("matrix had %d failures out of %d total", report.Failed, report.Total)
	}
}
