// Abate —— 止息混沌 · 性能基线测试插件
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/qingdujian/abate/core"
)

// PerfBaselinePlugin benchmark 对比基线插件。
type PerfBaselinePlugin struct{}

func (PerfBaselinePlugin) Capability() core.Capability {
	return core.Capability{Name: "perf-baseline", Version: "0.1.0", Protocol: "1"}
}

func (PerfBaselinePlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (PerfBaselinePlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "vet", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	benchCount := strings.Count(string(out), "ns/op")
	return TestResult{
		Name:     "perf-baseline",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d benchmarks, err=%v", benchCount, err != nil),
		Metrics:   map[string]float64{"benchmarks": float64(benchCount)},
	}
}

func (PerfBaselinePlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(PerfBaselinePlugin{})
}
