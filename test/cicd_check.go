// Abate —— 止息混沌 · CI/CD 检查插件
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

// CicdCheckPlugin 模拟完整 CI/CD 流水线插件。
type CicdCheckPlugin struct{}

func (CicdCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "cicd-check", Version: "0.1.0", Protocol: "1"}
}

func (CicdCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (CicdCheckPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	steps := []string{"build", "vet", "test"}
	results := make([]string, 0, len(steps))
	allPass := true
	for _, step := range steps {
		var cmd *exec.Cmd
		switch step {
		case "build":
			cmd = exec.CommandContext(ctx, "go", "build", "./...")
		case "vet":
			cmd = exec.CommandContext(ctx, "go", "vet", "./...")
		case "test":
			cmd = exec.CommandContext(ctx, "go", "vet", "./...")
		}
		cmd.Dir = target
		out, err := cmd.CombinedOutput()
		if err != nil {
			allPass = false
			results = append(results, fmt.Sprintf("%s:FAIL", step))
		} else {
			results = append(results, fmt.Sprintf("%s:OK", step))
		}
		_ = out
	}
	dur := time.Since(start)
	status := StatusPass
	if !allPass {
		status = StatusFail
	}
	return TestResult{
		Name:     "cicd-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(results, " "),
		Metrics:   map[string]float64{"steps": float64(len(steps))},
	}
}

func (CicdCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(CicdCheckPlugin{})
}
