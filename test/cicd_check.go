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

// CicdCheckPlugin 模拟完整 CI/CD 流水线插件，自动检测项目类型。
type CicdCheckPlugin struct{}

func (CicdCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "cicd-check", Version: "0.1.0", Protocol: "1"}
}

func (CicdCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (CicdCheckPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	results := make([]string, 0, 3)
	allPass := true

	steps := []struct {
		name string
		cmd  []string
	}{
		{"build", BuildCommand(pt)},
		{"vet", VetCommand(pt)},
		{"test", TestCommand(pt)},
	}

	for _, step := range steps {
		if step.cmd == nil {
			results = append(results, fmt.Sprintf("%s:SKIP", step.name))
			continue
		}
		cmd := exec.CommandContext(ctx, step.cmd[0], step.cmd[1:]...)
		cmd.Dir = target
		_, err := cmd.CombinedOutput()
		if err != nil {
			allPass = false
			results = append(results, fmt.Sprintf("%s:FAIL", step.name))
		} else {
			results = append(results, fmt.Sprintf("%s:OK", step.name))
		}
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
		Metrics:  map[string]float64{"steps": float64(len(steps))},
	}
}

func (CicdCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(CicdCheckPlugin{})
}
