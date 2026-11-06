// Abate —— 止息混沌 · 测试矩阵编排器
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/qingdujian/abate/log"
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
	logger   log.Logger
}

// NewRunner 构造编排器。
func NewRunner(r *Registry, timeout time.Duration) *Runner {
	return &Runner{
		registry: r,
		timeout:  timeout,
		logger:   log.Nop(),
	}
}

// SetLogger 注入日志器，实现插件执行全链路日志。
func (rn *Runner) SetLogger(l log.Logger) {
	if l == nil {
		l = log.Nop()
	}
	rn.logger = l.With(log.F("component", "runner"))
}

// RunMatrix 串行运行所有已注册插件，返回矩阵报告。
func (rn *Runner) RunMatrix(ctx context.Context, target string) MatrixReport {
	names := rn.registry.List()
	report := MatrixReport{Total: len(names)}
	rn.logger.Info("matrix started",
		log.F("total", len(names)),
		log.F("target", target),
		log.F("timeout", rn.timeout.String()),
	)
	matrixStart := time.Now()

	for i, name := range names {
		rn.logger.Debug("processing plugin",
			log.F("index", i+1),
			log.F("total", len(names)),
			log.F("plugin", name),
		)

		p, ok := rn.registry.Lookup(name)
		if !ok {
			rn.logger.Warn("plugin not found in registry", log.F("plugin", name))
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

		rn.logger.Info("plugin completed",
			log.F("plugin", result.Name),
			log.F("status", string(result.Status)),
			log.F("duration", result.Duration.String()),
			log.F("details", result.Details),
		)

		switch result.Status {
		case StatusPass:
			report.Passed++
		case StatusFail:
			report.Failed++
		case StatusSkip:
			report.Skipped++
		}
	}

	rn.logger.Info("matrix completed",
		log.F("total", report.Total),
		log.F("passed", report.Passed),
		log.F("failed", report.Failed),
		log.F("skipped", report.Skipped),
		log.F("duration", time.Since(matrixStart).String()),
	)

	return report
}

// runOne 运行单个插件，包含 Setup -> Run -> Teardown 全生命周期。
func (rn *Runner) runOne(ctx context.Context, p TestPlugin, target string) TestResult {
	cap := p.Capability()
	name := cap.Name
	pluginLogger := rn.logger.With(log.F("plugin", name))

	pluginLogger.Debug("setup phase started")
	child, cancel := context.WithTimeout(ctx, rn.timeout)
	defer cancel()

	setupStart := time.Now()
	cleanup, err := p.Setup(child)
	setupDur := time.Since(setupStart)
	if err != nil {
		pluginLogger.Error("setup failed",
			log.F("err", err.Error()),
			log.F("setup_duration", setupDur.String()),
		)
		return TestResult{Name: name, Status: StatusFail, Details: fmt.Sprintf("setup: %s", err)}
	}
	pluginLogger.Debug("setup completed", log.F("setup_duration", setupDur.String()))

	if cleanup != nil {
		defer cleanup()
	}

	pluginLogger.Debug("run phase started", log.F("target", target))
	runStart := time.Now()
	result := p.Run(child, target)
	result.Duration = time.Since(runStart)
	pluginLogger.Debug("run phase completed",
		log.F("run_duration", result.Duration.String()),
		log.F("status", string(result.Status)),
	)

	pluginLogger.Debug("teardown phase started")
	teardownStart := time.Now()
	if err := p.Teardown(child); err != nil {
		teardownDur := time.Since(teardownStart)
		pluginLogger.Error("teardown failed",
			log.F("err", err.Error()),
			log.F("teardown_duration", teardownDur.String()),
		)
		if result.Status == StatusPass {
			result.Status = StatusFail
			result.Details = fmt.Sprintf("teardown: %s", err)
		}
	} else {
		pluginLogger.Debug("teardown completed",
			log.F("teardown_duration", time.Since(teardownStart).String()),
		)
	}

	return result
}

// FormatText 将矩阵报告格式化为文本表格。
func (r MatrixReport) FormatText() string {
	var b strings.Builder
	fmt.Fprintf(&b, "=== Test Matrix Report ===\n")
	fmt.Fprintf(&b, "Total: %d  Pass: %d  Fail: %d  Skip: %d\n", r.Total, r.Passed, r.Failed, r.Skipped)
	for _, res := range r.Results {
		icon := "PASS"
		if res.Status == StatusFail {
			icon = "FAIL"
		}
		if res.Status == StatusSkip {
			icon = "SKIP"
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
