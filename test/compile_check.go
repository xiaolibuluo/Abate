// Abate —— 止息混沌 · 编译检查插件（自省层）
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

// CompileCheckPlugin 编译检查插件：实际编译目标项目，捕获编译错误。
type CompileCheckPlugin struct{}

func (CompileCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "compile-check", Version: "0.1.0", Protocol: "1"}
}

func (CompileCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (CompileCheckPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	if pt == ProjectUnknown {
		return TestResult{Name: "compile-check", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	// Go 项目用 go build ./... 捕获全部编译错误。
	if pt == ProjectGo {
		return runGoBuild(ctx, target, start)
	}
	// 非 Go 项目用 BuildCommand。
	bc := BuildCommand(pt)
	if bc == nil {
		return TestResult{Name: "compile-check", Status: StatusSkip, Duration: time.Since(start), Details: "no build command", Metrics: map[string]float64{}}
	}
	cmd := exec.CommandContext(ctx, bc[0], bc[1:]...)
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	if err == nil {
		return TestResult{Name: "compile-check", Status: StatusPass, Duration: dur, Details: "build:OK", Metrics: map[string]float64{"errors": 0}}
	}
	issues := ParseIssues(string(out), "build", SevError)
	st := CountIssues(issues)
	details := fmt.Sprintf("build:FAIL errors=%d warnings=%d", st.Errors, st.Warnings)
	if top := FormatIssues(issues, 5); top != "" {
		details += " | " + top
	}
	return TestResult{Name: "compile-check", Status: StatusFail, Duration: dur, Details: details, Metrics: map[string]float64{"errors": float64(st.Errors), "warnings": float64(st.Warnings)}}
}

// runGoBuild 执行 go build ./... 并解析编译错误。
func runGoBuild(ctx context.Context, target string, start time.Time) TestResult {
	cmd := exec.CommandContext(ctx, "go", "build", "./...")
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	dur := time.Since(start)
	if err == nil {
		// 额外检查 go vet 捕获未使用导入等问题。
		vetCmd := exec.CommandContext(ctx, "go", "vet", "./...")
		vetCmd.Dir = target
		vetOut, vetErr := vetCmd.CombinedOutput()
		if vetErr == nil {
			return TestResult{Name: "compile-check", Status: StatusPass, Duration: dur, Details: "build:OK vet:OK", Metrics: map[string]float64{"errors": 0}}
		}
		issues := ParseIssues(string(vetOut), "vet", SevWarning)
		st := CountIssues(issues)
		details := fmt.Sprintf("build:OK vet:WARN issues=%d", st.Total)
		if top := FormatIssues(issues, 5); top != "" {
			details += " | " + top
		}
		return TestResult{Name: "compile-check", Status: StatusFail, Duration: dur, Details: details, Metrics: map[string]float64{"errors": 0, "warnings": float64(st.Total)}}
	}
	// 解析编译错误。
	lines := strings.Split(string(out), "\n")
	errorCount := 0
	var errorLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "error") || strings.Contains(line, "undefined") ||
			strings.Contains(line, "cannot") || strings.Contains(line, "syntax") {
			errorCount++
			if len(errorLines) < 5 {
				errorLines = append(errorLines, line)
			}
		}
	}
	details := fmt.Sprintf("build:FAIL errors=%d", errorCount)
	if len(errorLines) > 0 {
		details += " | " + strings.Join(errorLines, "; ")
	}
	return TestResult{Name: "compile-check", Status: StatusFail, Duration: dur, Details: details, Metrics: map[string]float64{"errors": float64(errorCount)}}
}

func (CompileCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(CompileCheckPlugin{})
}
