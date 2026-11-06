// Abate —— 止息混沌 · 桩检查插件
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/qingdujian/abate/core"
)

// StubCheckPlugin 扫描源码中的桩函数（panic not implemented、直接 return nil）。
type StubCheckPlugin struct{}

func (StubCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "stub-check", Version: "0.1.0", Protocol: "1"}
}

func (StubCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (StubCheckPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	totalHits := 0
	details := make([]string, 0, 16)

	_ = filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "vendor" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			lower := strings.ToLower(trimmed)
			matched := false
			if strings.Contains(lower, "not implemented") {
				matched = true
			}
			if trimmed == "return nil" || trimmed == "return nil, nil" {
				matched = true
			}
			if matched {
				totalHits++
				rel, _ := filepath.Rel(target, path)
				if len(details) < 16 {
					details = append(details, rel+":"+itoa(i+1))
				}
			}
		}
		return nil
	})

	dur := time.Since(start)
	status := StatusPass
	if totalHits > 0 {
		status = StatusFail
	}
	return TestResult{
		Name:     "stub-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(details, "; "),
		Metrics:  map[string]float64{"hits": float64(totalHits)},
	}
}

func (StubCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(StubCheckPlugin{})
}
