// Abate —— 止息混沌 · 缺陷测试插件
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

// DefectTestPlugin 注入已知缺陷验证检测率插件。
type DefectTestPlugin struct{}

func (DefectTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "defect-test", Version: "0.1.0", Protocol: "1"}
}

func (DefectTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (DefectTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	vc := VetCommand(pt)
	if vc == nil {
		return TestResult{Name: "defect-test", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	cmd := exec.CommandContext(ctx, vc[0], vc[1:]...)
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	passCount := strings.Count(string(out), "--- PASS")
	return TestResult{
		Name:     "defect-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d pass, err=%v", passCount, err != nil),
		Metrics:   map[string]float64{"pass": float64(passCount)},
	}
}

func (DefectTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(DefectTestPlugin{})
}
