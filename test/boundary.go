// Abate —— 止息混沌 · 边界条件测试插件（固层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/qingdujian/abate/core"
)

// BoundaryTestPlugin 边界条件检测插件（固层）。
type BoundaryTestPlugin struct{}

func (BoundaryTestPlugin) Capability() core.Capability {
	return core.Capability{Name: "boundary-test", Version: "0.1.0", Protocol: "1"}
}

func (BoundaryTestPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (BoundaryTestPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	if pt == ProjectUnknown {
		return TestResult{Name: "boundary-test", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	var issues []string
	var maxDepth int
	var emptyFiles, hugeFiles, specialNames int
	maxHuge := int64(1 << 20) // 1MB
	_ = filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return nil
		}
		if info.IsDir() {
			if path != target && ShouldSkipDir(info.Name()) {
				return filepath.SkipDir
			}
			rel, _ := filepath.Rel(target, path)
			depth := strings.Count(filepath.ToSlash(rel), "/")
			if depth > maxDepth {
				maxDepth = depth
			}
			if depth > 10 {
				issues = append(issues, fmt.Sprintf("deep-nest: %s (%d)", rel, depth))
			}
			return nil
		}
		if !IsCodeFile(path) {
			return nil
		}
		name := info.Name()
		if info.Size() == 0 {
			emptyFiles++
			issues = append(issues, fmt.Sprintf("empty-file: %s", name))
		}
		if info.Size() > maxHuge {
			hugeFiles++
			issues = append(issues, fmt.Sprintf("huge-file: %s (%dKB)", name, info.Size()>>10))
		}
		if hasSpecialChar(name) {
			specialNames++
			issues = append(issues, fmt.Sprintf("special-name: %s", name))
		}
		return nil
	})
	dur := time.Since(start)
	status := StatusPass
	if emptyFiles > 0 || hugeFiles > 0 {
		status = StatusFail
	}
	details := fmt.Sprintf("maxDepth=%d emptyFiles=%d hugeFiles=%d specialNames=%d issues=%d", maxDepth, emptyFiles, hugeFiles, specialNames, len(issues))
	if len(issues) > 0 {
		shown := issues
		if len(shown) > 5 {
			shown = shown[:5]
		}
		details += " | " + strings.Join(shown, "; ")
		if len(issues) > 5 {
			details += fmt.Sprintf(" ... +%d more", len(issues)-5)
		}
	}
	return TestResult{
		Name:    "boundary-test",
		Status:  status,
		Duration: dur,
		Details: details,
		Metrics: map[string]float64{
			"max_depth":     float64(maxDepth),
			"empty_files":   float64(emptyFiles),
			"huge_files":    float64(hugeFiles),
			"special_names": float64(specialNames),
		},
	}
}

func (BoundaryTestPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(BoundaryTestPlugin{})
}

// hasSpecialChar 判断文件名是否含特殊字符（空格、中文、Unicode）。
func hasSpecialChar(name string) bool {
	for _, r := range name {
		if r > 127 || r == ' ' {
			return true
		}
	}
	return false
}
