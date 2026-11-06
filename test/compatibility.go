// Abate —— 止息混沌 · 兼容性测试插件
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

// CompatibilityTestPlugin Go 版本兼容性验证插件。
type CompatibilityTestPlugin struct{}

func (CompatibilityTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "compatibility-test", Version: "0.1.0", Protocol: "1"}
}

func (CompatibilityTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (CompatibilityTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "version")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	if err != nil {
		return TestResult{Name: "compatibility-test", Status: StatusFail, Duration: time.Since(start), Details: fmt.Sprintf("go version: %s", err)}
	}
	goVer := strings.TrimSpace(string(out))
	cmd2 := exec.CommandContext(ctx, "go", "build", "./...")
	cmd2.Dir = target
	_, err2 := cmd2.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err2 != nil {
		status = StatusFail
	}
	return TestResult{
		Name:     "compatibility-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%s, build err=%v", goVer, err2 != nil),
		Metrics:   map[string]float64{},
	}
}

func (CompatibilityTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(CompatibilityTestPlugin{})
}
