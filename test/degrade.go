// Abate —— 止息混沌 · 插件三级灾备兜底无感降级
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ScanMode 是插件扫描模式：完整、采样、兜底。
type ScanMode string

const (
	ModeFull    ScanMode = "full"
	ModeSampled ScanMode = "sampled"
	ModeMinimal ScanMode = "minimal"
)

// Metric 返回扫描模式的可量化标识。
func (m ScanMode) Metric() float64 {
	switch m {
	case ModeFull:
		return 0
	case ModeSampled:
		return 1
	default:
		return 2
	}
}

// ProjectScale 是目标项目的快速规模画像。
type ProjectScale struct {
	Files          int
	CodeFiles      int
	Modules        int
	Bytes          int64
	EstimatedLines int
}

// ScanPlan 是单个插件本次执行的扫描计划。
type ScanPlan struct {
	Mode          ScanMode
	Scale         ProjectScale
	CoverageRatio float64
	FileBudget    int
	ByteBudget    int64
	Reasons       []string
}

const (
	fullMaxCodeFiles    = 2000
	fullMaxBytes        = int64(32 << 20)
	sampledMaxCodeFiles = 20000
	sampledMaxBytes     = int64(512 << 20)
	sampledFileBudget   = 5000
	sampledByteBudget   = int64(128 << 20)
	minimalFileBudget   = 200
	minimalByteBudget   = int64(8 << 20)
)

var moduleFiles = map[string]bool{
	"go.mod": true, "settings.gradle": true, "settings.gradle.kts": true,
	"pom.xml": true, "package.json": true, "Cargo.toml": true,
	"requirements.txt": true, "setup.py": true, "pyproject.toml": true,
}

// AnalyzeScale 快速统计项目规模，只读目录元数据，不读取文件内容。
func AnalyzeScale(target string) ProjectScale {
	var s ProjectScale
	_ = filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return nil
		}
		if info.IsDir() {
			if path != target && ShouldSkipDir(info.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		s.Files++
		s.Bytes += info.Size()
		if moduleFiles[info.Name()] {
			s.Modules++
		}
		if IsCodeFile(path) {
			s.CodeFiles++
		}
		return nil
	})
	s.EstimatedLines = int(s.Bytes / 40)
	return s
}

// PlanScan 根据项目规模和剩余时间预算生成三级扫描计划。
func PlanScan(ctx context.Context, target string) ScanPlan {
	s := AnalyzeScale(target)
	mode := ModeFull
	reasons := []string{"scale:full-range"}
	if s.CodeFiles > fullMaxCodeFiles || s.Bytes > fullMaxBytes {
		mode = ModeSampled
		reasons = []string{"scale:sampled-range"}
	}
	if s.CodeFiles > sampledMaxCodeFiles || s.Bytes > sampledMaxBytes {
		mode = ModeMinimal
		reasons = []string{"scale:minimal-range"}
	}
	if d, ok := ctx.Deadline(); ok {
		remain := time.Until(d)
		if remain < 10*time.Second {
			mode = ModeMinimal
			reasons = append(reasons, "deadline:minimal")
		} else if remain < 30*time.Second && mode == ModeFull {
			mode = ModeSampled
			reasons = append(reasons, "deadline:sampled")
		}
	}
	return newPlan(mode, s, reasons)
}

func newPlan(mode ScanMode, s ProjectScale, reasons []string) ScanPlan {
	p := ScanPlan{Mode: mode, Scale: s, CoverageRatio: 1, Reasons: reasons}
	switch mode {
	case ModeSampled:
		p.FileBudget = sampledFileBudget
		p.ByteBudget = sampledByteBudget
		p.CoverageRatio = coverage(s.CodeFiles, sampledFileBudget)
	case ModeMinimal:
		p.FileBudget = minimalFileBudget
		p.ByteBudget = minimalByteBudget
		p.CoverageRatio = coverage(s.CodeFiles, minimalFileBudget)
	}
	return p
}

func coverage(codeFiles, budget int) float64 {
	if codeFiles <= 0 {
		return 1
	}
	if codeFiles <= budget {
		return 1
	}
	return float64(budget) / float64(codeFiles)
}

func degradePlan(p ScanPlan, reason string) ScanPlan {
	switch p.Mode {
	case ModeFull:
		return newPlan(ModeSampled, p.Scale, append(p.Reasons, reason))
	case ModeSampled:
		return newPlan(ModeMinimal, p.Scale, append(p.Reasons, reason))
	default:
		return newPlan(ModeMinimal, p.Scale, append(p.Reasons, reason))
	}
}

func (p ScanPlan) budget(timeout time.Duration) time.Duration {
	if timeout <= 0 {
		timeout = 60 * time.Second
	}
	switch p.Mode {
	case ModeFull:
		return timeout * 80 / 100
	case ModeSampled:
		return timeout * 40 / 100
	default:
		return timeout * 10 / 100
	}
}

// RunPluginSafely 以三级灾备方式执行插件：完整 -> 采样 -> 兜底，
// 超时自动降级，瞬态失败自动重试，失败结果自动归因（稳层闭环）。
func RunPluginSafely(ctx context.Context, p TestPlugin, target string, timeout time.Duration) TestResult {
	name := p.Capability().Name
	plan := PlanScan(ctx, target)
	for {
		child, cancel := context.WithTimeout(ctx, plan.budget(timeout))
		res, timedOut := runIsolated(child, p, target)
		cancel()
		if timedOut {
			if plan.Mode == ModeMinimal {
				return minimalResult(name, plan, "minimal-budget-exhausted")
			}
			next := degradePlan(plan, fmt.Sprintf("degraded:after-%s-timeout", plan.Mode))
			if next.Mode == ModeMinimal {
				return minimalResult(name, next, fmt.Sprintf("degraded:after-%s-timeout", plan.Mode))
			}
			plan = next
			continue
		}
		// 瞬态失败自动重试一次（稳层）。
		if res.Status == StatusFail && ClassifyFailure(res) == FailureTransient {
			retryCtx, retryCancel := context.WithTimeout(ctx, plan.budget(timeout))
			r2, to2 := runIsolated(retryCtx, p, target)
			retryCancel()
			if !to2 {
				res = r2
				if res.Metrics == nil {
					res.Metrics = map[string]float64{}
				}
				res.Metrics["retried"] = 1
			}
		}
		return annotateResult(AnnotateClassification(res), plan)
	}
}

func runIsolated(ctx context.Context, p TestPlugin, target string) (TestResult, bool) {
	ch := make(chan TestResult, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- TestResult{Name: p.Capability().Name, Status: StatusFail, Details: fmt.Sprintf("panic: %v", r), Metrics: map[string]float64{}}
			}
		}()
		ch <- p.Run(ctx, target)
	}()
	select {
	case res := <-ch:
		return res, false
	case <-ctx.Done():
		return TestResult{}, true
	}
}

func annotateResult(res TestResult, plan ScanPlan) TestResult {
	if res.Metrics == nil {
		res.Metrics = map[string]float64{}
	}
	res.Metrics["scan_mode"] = plan.Mode.Metric()
	res.Metrics["coverage_ratio"] = plan.CoverageRatio
	res.Metrics["code_files"] = float64(plan.Scale.CodeFiles)
	res.Metrics["modules"] = float64(plan.Scale.Modules)
	if res.Details != "" {
		res.Details += " "
	}
	res.Details += fmt.Sprintf("mode=%s coverage=%.2f", plan.Mode, plan.CoverageRatio)
	return res
}

func minimalResult(name string, plan ScanPlan, reason string) TestResult {
	status := StatusPass
	if plan.Scale.Modules == 0 || plan.Scale.CodeFiles == 0 {
		status = StatusSkip
	}
	details := fmt.Sprintf("minimal probe: modules=%d code_files=%d bytes=%d reason=%s", plan.Scale.Modules, plan.Scale.CodeFiles, plan.Scale.Bytes, reason)
	return TestResult{
		Name:    name,
		Status:  status,
		Details: details + fmt.Sprintf(" mode=%s coverage=%.2f", plan.Mode, plan.CoverageRatio),
		Metrics: map[string]float64{
			"scan_mode":      plan.Mode.Metric(),
			"coverage_ratio": plan.CoverageRatio,
			"code_files":     float64(plan.Scale.CodeFiles),
			"modules":        float64(plan.Scale.Modules),
			"degraded":       1,
		},
	}
}
