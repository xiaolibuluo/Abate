// Abate —— 止息混沌 · 严层审计插件
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"time"

	"github.com/qingdujian/abate/core"
)

// AuditPlugin 全链路审计插件（严层）。
// 零容忍策略：发现敏感信息即 FAIL，缺头/TODO 仅 WARN。
type AuditPlugin struct{}

func (AuditPlugin) Capability() core.Capability {
	return core.Capability{Name: "audit", Version: "0.1.0", Protocol: "1"}
}

func (AuditPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (AuditPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	if pt == ProjectUnknown {
		return TestResult{Name: "audit", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	plan := PlanScan(ctx, target)
	maxFiles := 0
	if plan.Mode != ModeFull {
		maxFiles = plan.FileBudget
	}
	r := RunAudit(target, maxFiles)
	dur := time.Since(start)
	// 零容忍：敏感信息 > 0 即 FAIL。
	status := StatusPass
	if r.Secrets > 0 {
		status = StatusFail
	}
	details := FormatAudit(r, 5)
	return TestResult{
		Name:    "audit",
		Status:  status,
		Duration: dur,
		Details: fmt.Sprintf("%s mode=%s coverage=%.2f", details, plan.Mode, plan.CoverageRatio),
		Metrics: map[string]float64{
			"files":    float64(r.Total),
			"headers":  float64(r.Headers),
			"secrets":  float64(r.Secrets),
			"todos":    float64(r.Todos),
			"fixmes":   float64(r.Fixmes),
		},
	}
}

func (AuditPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(AuditPlugin{})
}
