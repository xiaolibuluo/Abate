// Abate —— 止息混沌 · 测试矩阵编排器
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// MatrixReport 是矩阵测试的汇总报告。
type MatrixReport struct {
	Total   int
	Passed  int
	Failed  int
	Skipped int
	Results []TestResult
}

// Runner 是测试矩阵编排器。
type Runner struct {
	registry *Registry
	timeout  time.Duration
}

// NewRunner 构造编排器。
func NewRunner(r *Registry, timeout time.Duration) *Runner {
	return &Runner{registry: r, timeout: timeout}
}

// RunMatrix 串行运行所有已注册插件，返回矩阵报告。
func (rn *Runner) RunMatrix(ctx context.Context, target string) MatrixReport {
	names := rn.registry.List()
	report := MatrixReport{Total: len(names)}
	for _, name := range names {
		p, ok := rn.registry.Lookup(name)
		if !ok {
			report.Results = append(report.Results, TestResult{
				Name: name, Status: StatusSkip, Details: "not found in registry",
			})
			report.Skipped++
			continue
		}
		result := rn.runOne(ctx, p, target)
		report.Results = append(report.Results, result)
		icon := "PASS"
		if result.Status == StatusFail {
			icon = "FAIL"
		}
		if result.Status == StatusSkip {
			icon = "SKIP"
		}
		fmt.Printf("[%s] %-24s %8v  %s\n", icon, result.Name, result.Duration, result.Details)
		switch result.Status {
		case StatusPass:
			report.Passed++
		case StatusFail:
			report.Failed++
		case StatusSkip:
			report.Skipped++
	}
	}
	return report
}

// runOne 运行单个插件，包含 Setup → Run → Teardown 全生命周期。
func (rn *Runner) runOne(ctx context.Context, p TestPlugin, target string) TestResult {
	cap := p.Capability()
	name := cap.Name

	child, cancel := context.WithTimeout(ctx, rn.timeout)
	defer cancel()

	cleanup, err := p.Setup(child)
	if err != nil {
		return TestResult{Name: name, Status: StatusFail, Details: fmt.Sprintf("setup: %s", err)}
	}
	if cleanup != nil {
		defer cleanup()
	}

	start := time.Now()
	result := p.Run(child, target)
	result.Duration = time.Since(start)

	if err := p.Teardown(child); err != nil {
		if result.Status == StatusPass {
			result.Status = StatusFail
			result.Details = fmt.Sprintf("teardown: %s", err)
		}
	}

	return result
}

// FormatText 将矩阵报告格式化为文本表格。
func (r MatrixReport) FormatText() string {
	var b strings.Builder
	fmt.Fprintf(&b, "=== Test Matrix Report ===\n")
	fmt.Fprintf(&b, "Total: %d  Pass: %d  Fail: %d  Skip: %d\n", r.Total, r.Passed, r.Failed, r.Skipped)
	for _, res := range r.Results {
	icon := "✅"
	if res.Status == StatusFail {
		icon = "❌"
	}
	if res.Status == StatusSkip {
		icon = "⏭️"
	}
	fmt.Fprintf(&b, "  %s %-24s %8v  %s\n", icon, res.Name, res.Duration, res.Details)
	}
	if r.Failed == 0 {
		fmt.Fprintf(&b, "=== ALL PASS ===\n")
	} else {
		fmt.Fprintf(&b, "=== %d FAILED ===\n", r.Failed)
	}
	return b.String()
}
