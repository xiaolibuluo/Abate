// Abate —— 止息混沌 · 集成测试插件
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

// IntegrationTestPlugin 多包协作场景验证插件。
type IntegrationTestPlugin struct{}

func (IntegrationTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "integration-test", Version: "0.1.0", Protocol: "1"}
}

func (IntegrationTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (IntegrationTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "test", "-tags", "integration", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	okCount := strings.Count(string(out), "ok")
	return TestResult{
		Name:     "integration-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d packages, err=%v", okCount, err != nil),
		Metrics:   map[string]float64{"packages": float64(okCount)},
	}
}

func (IntegrationTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(IntegrationTestPlugin{})
}
