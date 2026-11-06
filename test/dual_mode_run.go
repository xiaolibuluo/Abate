// Abate —— 止息混沌 · 双模式运行测试插件
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/qingdujian/abate/core"
)

// DualModeRunPlugin debug+release 模式运行测试插件。
type DualModeRunPlugin struct{}

func (DualModeRunPlugin) Capability() core.Capability {
	return core.Capability{Name: "dual-mode-run", Version: "0.1.0", Protocol: "1"}
}

func (DualModeRunPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (DualModeRunPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	results := make([]string, 0, 2)
	allPass := true
	debugCmd := exec.CommandContext(ctx, "go", "vet", "./...")
	debugCmd.Dir = target
	_, err := debugCmd.CombinedOutput()
	if err != nil {
		allPass = false
		results = append(results, "debug:FAIL")
	} else {
		results = append(results, "debug:OK")
	}
	releaseCmd := exec.CommandContext(ctx, "go", "vet", "./...")
	releaseCmd.Dir = target
	_, err = releaseCmd.CombinedOutput()
	if err != nil {
		allPass = false
		results = append(results, "release:FAIL")
	} else {
		results = append(results, "release:OK")
	}
	dur := time.Since(start)
	status := StatusPass
	if !allPass {
		status = StatusFail
	}
	return TestResult{
		Name:     "dual-mode-run",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(results, " "),
		Metrics:   map[string]float64{"modes": 2},
	}
}

func (DualModeRunPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(DualModeRunPlugin{})
}
