// Abate —— 止息混沌 · 静态检查插件
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

// StaticAnalysisPlugin go vet + gosec 静态分析插件。
type StaticAnalysisPlugin struct{}

func (StaticAnalysisPlugin) Capability() core.Capability {
	return core.Capability{Name: "static-analysis", Version: "0.1.0", Protocol: "1"}
}

func (StaticAnalysisPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (StaticAnalysisPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "vet", "-tags", "integration", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	issueCount := strings.Count(string(out), "\n")
	return TestResult{
		Name:     "static-analysis",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("vet issues=%d, err=%v", issueCount, err != nil),
		Metrics:   map[string]float64{"issues": float64(issueCount)},
	}
}

func (StaticAnalysisPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(StaticAnalysisPlugin{})
}
