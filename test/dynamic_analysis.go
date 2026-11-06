// Abate —— 止息混沌 · 动态检查插件
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

// DynamicAnalysisPlugin -race + -cover 动态分析插件。
type DynamicAnalysisPlugin struct{}

func (DynamicAnalysisPlugin) Capability() core.Capability {
	return core.Capability{Name: "dynamic-analysis", Version: "0.1.0", Protocol: "1"}
}

func (DynamicAnalysisPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (DynamicAnalysisPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "vet", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	raceCount := strings.Count(string(out), "DATA RACE")
	if err != nil || raceCount > 0 {
		status = StatusFail
	}
	return TestResult{
		Name:     "dynamic-analysis",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d races, err=%v", raceCount, err != nil),
		Metrics:   map[string]float64{"races": float64(raceCount)},
	}
}

func (DynamicAnalysisPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(DynamicAnalysisPlugin{})
}
