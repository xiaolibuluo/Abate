// Abate —— 止息混沌 · 三级灾备降级框架测试
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/qingdujian/abate/core"
)

// TestPlanScanFullMode 验证小规模项目进入完整模式。
func TestPlanScanFullMode(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module x\n"), 0644); err != nil {
		t.Fatal(err)
	}
	plan := PlanScan(context.Background(), dir)
	if plan.Mode != ModeFull {
		t.Fatalf("want full, got %s", plan.Mode)
	}
	if plan.CoverageRatio != 1 {
		t.Fatalf("want coverage 1, got %.2f", plan.CoverageRatio)
	}
}

// TestPlanScanSampledByFiles 验证代码文件数超阈值时降级为采样模式且覆盖率低于 1。
func TestPlanScanSampledByFiles(t *testing.T) {
	dir := t.TempDir()
	for i := 0; i < sampledFileBudget+100; i++ {
		p := filepath.Join(dir, fmt.Sprintf("f%05d.go", i))
		if err := os.WriteFile(p, []byte("package x\n"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	plan := PlanScan(context.Background(), dir)
	if plan.Mode != ModeSampled {
		t.Fatalf("want sampled, got %s", plan.Mode)
	}
	if plan.CoverageRatio >= 1 {
		t.Fatalf("sampled coverage should be < 1, got %.2f", plan.CoverageRatio)
	}
}

// TestPlanScanMinimalByDeadline 验证剩余时间不足时直接进入兜底模式。
func TestPlanScanMinimalByDeadline(t *testing.T) {
	dir := t.TempDir()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	plan := PlanScan(ctx, dir)
	if plan.Mode != ModeMinimal {
		t.Fatalf("want minimal, got %s", plan.Mode)
	}
}

// slowPlugin 故意不尊重 ctx，模拟超时场景。
type slowPlugin struct{ name string }

func (p slowPlugin) Capability() core.Capability { return core.Capability{Name: p.name} }
func (slowPlugin) Setup(ctx context.Context) (func(), error) { return nil, nil }
func (p slowPlugin) Run(ctx context.Context, target string) TestResult {
	time.Sleep(5 * time.Second)
	return TestResult{Name: p.name, Status: StatusPass}
}
func (slowPlugin) Teardown(ctx context.Context) error { return nil }

// TestRunPluginSafelyDegrades 验证超时后自动降级到兜底模式并输出降级指标。
func TestRunPluginSafelyDegrades(t *testing.T) {
	dir := t.TempDir()
	res := RunPluginSafely(context.Background(), slowPlugin{name: "slow"}, dir, 2*time.Second)
	mode, ok := res.Metrics["scan_mode"]
	if !ok || mode != ModeMinimal.Metric() {
		t.Fatalf("want minimal mode metric, got %+v", res.Metrics)
	}
	if res.Metrics["degraded"] != 1 {
		t.Fatalf("want degraded=1, got %+v", res.Metrics)
	}
	if res.Status != StatusSkip {
		t.Fatalf("empty dir minimal probe should skip, got %s", res.Status)
	}
}

// panicPlugin 故意 panic，验证隔离恢复。
type panicPlugin struct{}

func (panicPlugin) Capability() core.Capability { return core.Capability{Name: "panic"} }
func (panicPlugin) Setup(ctx context.Context) (func(), error) { return nil, nil }
func (panicPlugin) Run(ctx context.Context, target string) TestResult {
	panic("boom")
}
func (panicPlugin) Teardown(ctx context.Context) error { return nil }

// TestRunPluginSafelyPanicRecovered 验证插件 panic 被隔离捕获并转为 FAIL。
func TestRunPluginSafelyPanicRecovered(t *testing.T) {
	dir := t.TempDir()
	res := RunPluginSafely(context.Background(), panicPlugin{}, dir, 10*time.Second)
	if res.Status != StatusFail {
		t.Fatalf("want FAIL on panic, got %s", res.Status)
	}
}
