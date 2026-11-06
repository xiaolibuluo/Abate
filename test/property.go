// Abate —— 止息混沌 · 属性测试插件
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

// PropertyTestPlugin 随机输入验证不变式插件。
type PropertyTestPlugin struct{}

func (PropertyTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "property-test", Version: "0.1.0", Protocol: "1"}
}

func (PropertyTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (PropertyTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "test", "-run", "Property", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	passCount := strings.Count(string(out), "PASS")
	return TestResult{
		Name:     "property-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d property checks, err=%v", passCount, err != nil),
		Metrics:   map[string]float64{"checks": float64(passCount)},
	}
}

func (PropertyTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(PropertyTestPlugin{})
}
