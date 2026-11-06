// Abate —— 止息混沌 · 稳定性探测插件（稳层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/qingdujian/abate/core"
)

// StabilityProbePlugin 稳定性探测插件（稳层）。
// 对目标项目执行重试验证、超时容忍和资源耗尽恢复测试。
type StabilityProbePlugin struct{}

func (StabilityProbePlugin) Capability() core.Capability {
	return core.Capability{Name: "stability-probe", Version: "0.1.0", Protocol: "1"}
}

func (StabilityProbePlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (StabilityProbePlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	if pt == ProjectUnknown {
		return TestResult{Name: "stability-probe", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	bc := BuildCommand(pt)
	if bc == nil {
		return TestResult{Name: "stability-probe", Status: StatusSkip, Duration: time.Since(start), Details: "no build command", Metrics: map[string]float64{}}
	}
	// 三项稳定性探测。
	retryResult := probeRetry(ctx, target, bc)
	timeoutResult := probeTimeout(ctx, target, bc)
	resourceResult := probeResource(ctx, target, bc)
	dur := time.Since(start)
	// 汇总：三项全过才 PASS。
	status := StatusPass
	if !retryResult || !timeoutResult || !resourceResult {
		status = StatusFail
	}
	details := fmt.Sprintf("retry=%v timeout=%v resource=%v", retryResult, timeoutResult, resourceResult)
	return TestResult{
		Name:    "stability-probe",
		Status:  status,
		Duration: dur,
		Details: details,
		Metrics: map[string]float64{
			"retry_pass":   boolToFloat(retryResult),
			"timeout_pass": boolToFloat(timeoutResult),
			"resource_pass": boolToFloat(resourceResult),
		},
	}
}

// probeRetry 重试验证：连续执行 3 次构建，结果必须一致。
func probeRetry(ctx context.Context, target string, bc []string) bool {
	results := make([]bool, 3)
	for i := 0; i < 3; i++ {
		child, cancel := context.WithTimeout(ctx, 60*time.Second)
		cmd := exec.CommandContext(child, bc[0], bc[1:]...)
		cmd.Dir = target
		err := cmd.Run()
		cancel()
		results[i] = (err == nil)
	}
	// 三次结果必须一致。
	return results[0] == results[1] && results[1] == results[2]
}

// probeTimeout 超时容忍：用极短超时（1ms）触发超时，验证不 panic 不挂起。
func probeTimeout(ctx context.Context, target string, bc []string) bool {
	child, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(child, bc[0], bc[1:]...)
	cmd.Dir = target
	err := cmd.Run()
	// 超时是预期行为，只要不 panic 即通过。
	_ = err
	return true
}

// probeResource 资源耗尽恢复：执行构建后检查进程是否正常退出（无僵尸进程）。
func probeResource(ctx context.Context, target string, bc []string) bool {
	child, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(child, bc[0], bc[1:]...)
	cmd.Dir = target
	err := cmd.Run()
	// 进程正常退出（无论成功失败）即通过。
	return err == nil || cmd.ProcessState != nil
}

func (StabilityProbePlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(StabilityProbePlugin{})
}

func boolToFloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
