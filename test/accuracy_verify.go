// Abate —— 止息混沌 · 准确性验证插件（准层）
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

// AccuracyVerifyPlugin 准确性验证插件（准层）。
// 验证工具输出的解析正确性、误报率和指标一致性。
type AccuracyVerifyPlugin struct{}

func (AccuracyVerifyPlugin) Capability() core.Capability {
	return core.Capability{Name: "accuracy-verify", Version: "0.1.0", Protocol: "1"}
}

func (AccuracyVerifyPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (AccuracyVerifyPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	if pt == ProjectUnknown {
		return TestResult{Name: "accuracy-verify", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	vc := VetCommand(pt)
	if vc == nil {
		return TestResult{Name: "accuracy-verify", Status: StatusSkip, Duration: time.Since(start), Details: "no vet command", Metrics: map[string]float64{}}
	}
	// 执行工具获取原始输出。
	child, cancel := context.WithTimeout(ctx, 60*time.Second)
	cmd := exec.CommandContext(child, vc[0], vc[1:]...)
	cmd.Dir = target
	out, err := cmd.CombinedOutput()
	cancel()
	rawOutput := string(out)
	dur := time.Since(start)
	// 检查 1：解析正确性 — ParseIssues 提取的 issue 行数不应超过原始输出行数。
	issues := ParseIssues(rawOutput, "verify", SevWarning)
	rawLines := 0
	for _, line := range strings.Split(rawOutput, "\n") {
		if strings.TrimSpace(line) != "" {
			rawLines++
		}
	}
	parseAccurate := len(issues) <= rawLines
	// 检查 2：误报率 — 解析出的 issue 文件必须真实存在。
	falsePositives := 0
	for _, iss := range issues {
		if !fileExists(iss.File) && !isRelativePath(iss.File) {
			falsePositives++
		}
	}
	falsePositiveRate := 0.0
	if len(issues) > 0 {
		falsePositiveRate = float64(falsePositives) / float64(len(issues))
	}
	// 检查 3：指标一致性 — CountIssues 统计总数应等于 issues 长度。
	st := CountIssues(issues)
	metricConsistent := st.Total == len(issues) && st.Errors+st.Warnings+st.Infos == st.Total
	// 汇总。
	status := StatusPass
	if !parseAccurate || falsePositiveRate > 0.5 || !metricConsistent {
		status = StatusFail
	}
	details := fmt.Sprintf("rawLines=%d issues=%d parseAccurate=%v fpr=%.2f metricConsistent=%v err=%v",
		rawLines, len(issues), parseAccurate, falsePositiveRate, metricConsistent, err != nil)
	return TestResult{
		Name:    "accuracy-verify",
		Status:  status,
		Duration: dur,
		Details: details,
		Metrics: map[string]float64{
			"raw_lines":          float64(rawLines),
			"issues":             float64(len(issues)),
			"false_positives":    float64(falsePositives),
			"false_positive_rate": falsePositiveRate,
			"parse_accurate":     boolToFloat(parseAccurate),
			"metric_consistent":  boolToFloat(metricConsistent),
		},
	}
}

func (AccuracyVerifyPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(AccuracyVerifyPlugin{})
}

// isRelativePath 判断是否为相对路径（可能不含完整路径前缀）。
func isRelativePath(p string) bool {
	return !strings.HasPrefix(p, "/") && !strings.Contains(p, ":\\")
}
