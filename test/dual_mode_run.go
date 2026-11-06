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
	pt := DetectProjectType(target)
	vc := VetCommand(pt)
	if vc == nil {
		return TestResult{Name: "dual-mode-run", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	if pt == ProjectGo {
		// Go: debug + release dual vet
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
	} else {
		// Non-Go: single mode vet
		singleCmd := exec.CommandContext(ctx, vc[0], vc[1:]...)
		singleCmd.Dir = target
		_, err := singleCmd.CombinedOutput()
		if err != nil {
			allPass = false
			results = append(results, "vet:FAIL")
		} else {
			results = append(results, "vet:OK")
		}
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
		Metrics:  map[string]float64{"modes": float64(len(results))},
	}
}

func (DualModeRunPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(DualModeRunPlugin{})
}
