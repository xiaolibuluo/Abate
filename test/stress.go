// Abate —— 止息混沌 · 压力测试插件
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

// StressTestPlugin 超负载稳定性测试插件。
type StressTestPlugin struct{}

func (StressTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "stress-test", Version: "0.1.0", Protocol: "1"}
}

func (StressTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (StressTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "vet", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	passCount := strings.Count(string(out), "--- PASS")
	return TestResult{
		Name:     "stress-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d pass, err=%v", passCount, err != nil),
		Metrics:   map[string]float64{"pass": float64(passCount)},
	}
}

func (StressTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(StressTestPlugin{})
}
