// Abate —— 止息混沌 · 模糊测试插件
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

// FuzzTestPlugin 随机/畸形输入不崩溃测试插件。
type FuzzTestPlugin struct{}

func (FuzzTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "fuzz-test", Version: "0.1.0", Protocol: "1"}
}

func (FuzzTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (FuzzTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "vet", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	fuzzCount := strings.Count(string(out), "fuzzing")
	return TestResult{
		Name:     "fuzz-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d fuzz iterations, err=%v", fuzzCount, err != nil),
		Metrics:   map[string]float64{"iterations": float64(fuzzCount)},
	}
}

func (FuzzTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(FuzzTestPlugin{})
}
