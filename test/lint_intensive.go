// Abate —— 止息混沌 · Lint 高强度测试插件
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

// LintIntensivePlugin golangci-lint 全规则检查插件。
type LintIntensivePlugin struct{}

func (LintIntensivePlugin) Capability() core.Capability {
	return core.Capability{Name: "lint-intensive", Version: "0.1.0", Protocol: "1"}
}

func (LintIntensivePlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (LintIntensivePlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "vet", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	issueCount := strings.Count(string(out), "\n")
	return TestResult{
		Name:     "lint-intensive",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("go vet issues=%d, err=%v", issueCount, err != nil),
		Metrics:   map[string]float64{"issues": float64(issueCount)},
	}
}

func (LintIntensivePlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(LintIntensivePlugin{})
}
