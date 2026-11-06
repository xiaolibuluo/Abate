// Abate —— 止息混沌 · 单元测试插件
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

// UnitTestPlugin 函数级隔离测试插件。
type UnitTestPlugin struct{}

func (UnitTestPlugin) Capability() core.Capability {
	return core.Capability{
		Name:    "unit-test",
		Version: "0.1.0",
		Protocol: "1",
	}
}

func (UnitTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (UnitTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	vc := VetCommand(pt)
	if vc == nil {
		return TestResult{Name: "unit-test", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	cmd := exec.CommandContext(ctx, vc[0], vc[1:]...)
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	lines := strings.Split(string(out), "\n")
	okCount := 0
	failCount := 0
	for _, l := range lines {
		if strings.HasPrefix(l, "ok") {
			okCount++
		}
		if strings.HasPrefix(l, "FAIL") {
			failCount++
		}
	}
	status := StatusPass
	if err != nil || failCount > 0 {
		status = StatusFail
	}
	return TestResult{
		Name:    "unit-test",
		Status:  status,
		Duration: dur,
		Details:  fmt.Sprintf("%d ok, %d fail", okCount, failCount),
		Metrics:  map[string]float64{"ok": float64(okCount), "fail": float64(failCount)},
	}
}

func (UnitTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	// 注册到默认注册表。
	DefaultRegistry.Register(UnitTestPlugin{})
}
