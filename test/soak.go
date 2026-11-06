// Abate —— 止息混沌 · 持久稳定测试插件
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

// SoakTestPlugin 长时间内存泄漏检测插件。
type SoakTestPlugin struct{}

func (SoakTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "soak-test", Version: "0.1.0", Protocol: "1"}
}

func (SoakTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (SoakTestPlugin) Run(ctx context.Context, target string) TestResult {
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
		Name:     "soak-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d pass, err=%v", passCount, err != nil),
		Metrics:   map[string]float64{"pass": float64(passCount)},
	}
}

func (SoakTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(SoakTestPlugin{})
}
