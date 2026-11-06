// Abate —— 止息混沌 · 兼容性测试插件
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/qingdujian/abate/core"
)

// CompatibilityTestPlugin 多语言兼容性验证插件，自动检测项目类型并构建。
type CompatibilityTestPlugin struct{}

func (CompatibilityTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "compatibility-test", Version: "0.1.0", Protocol: "1"}
}

func (CompatibilityTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (CompatibilityTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	ptName := projectTypeName(pt)
	buildCmd := BuildCommand(pt)
	if buildCmd == nil {
		return TestResult{
			Name:    "compatibility-test",
			Status:  StatusFail,
			Duration: time.Since(start),
			Details: fmt.Sprintf("unknown project type at %s", target),
		}
	}
	cmd := exec.CommandContext(ctx, buildCmd[0], buildCmd[1:]...)
	cmd.Dir = target
	_, err := cmd.CombinedOutput()
	dur := time.Since(start)
	status := StatusPass
	if err != nil {
		status = StatusFail
	}
	return TestResult{
		Name:     "compatibility-test",
		Status:   status,
		Duration: dur,
		Details:  fmt.Sprintf("%s, build err=%v", ptName, err != nil),
		Metrics:  map[string]float64{},
	}
}

func (CompatibilityTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(CompatibilityTestPlugin{})
}

func projectTypeName(pt ProjectType) string {
	switch pt {
	case ProjectGo:
		return "Go"
	case ProjectGradle:
		return "Gradle/Android"
	case ProjectMaven:
		return "Maven/Java"
	case ProjectNpm:
		return "Node.js/npm"
	case ProjectCargo:
		return "Rust/Cargo"
	case ProjectPython:
		return "Python"
	default:
		return "Unknown"
	}
}
