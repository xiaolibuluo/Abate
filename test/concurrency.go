// Abate —— 止息混沌 · 并发测试插件
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

// ConcurrencyTestPlugin 并发数据竞争检测插件。
type ConcurrencyTestPlugin struct{}

func (ConcurrencyTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "concurrency-test", Version: "0.1.0", Protocol: "1"}
}

func (ConcurrencyTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (ConcurrencyTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	vc := VetCommand(pt)
	if vc == nil {
		return TestResult{Name: "concurrency-test", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	cmd := exec.CommandContext(ctx, vc[0], vc[1:]...)
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	raceCount := strings.Count(string(out), "DATA RACE")
	if err != nil || raceCount > 0 {
		status = StatusFail
	}
	return TestResult{
		Name:     "concurrency-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%d races, err=%v", raceCount, err != nil),
		Metrics:   map[string]float64{"races": float64(raceCount)},
	}
}

func (ConcurrencyTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(ConcurrencyTestPlugin{})
}
