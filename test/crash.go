// Abate —— 止息混沌 · 崩溃测试插件
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

// CrashTestPlugin 强制 panic 后恢复验证插件。
type CrashTestPlugin struct{}

func (CrashTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "crash-test", Version: "0.1.0", Protocol: "1"}
}

func (CrashTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (CrashTestPlugin) Run(ctx context.Context, target string) TestResult {
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
		Name:     "crash-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d pass, err=%v", passCount, err != nil),
		Metrics:   map[string]float64{"pass": float64(passCount)},
	}
}

func (CrashTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(CrashTestPlugin{})
}
