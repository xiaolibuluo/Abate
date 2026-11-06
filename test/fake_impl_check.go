// Abate —— 止息混沌 · 假实现检查插件
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

// FakeImplCheckPlugin 扫描源码中显式标记的假实现（函数名含 Fake/Mock/Stub/Placeholder）。
type FakeImplCheckPlugin struct{}

func (FakeImplCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "fake-impl-check", Version: "0.1.0", Protocol: "1"}
}

func (FakeImplCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func isFakeMarker(line string) bool {
	lower := strings.ToLower(line)
	if strings.Contains(lower, "func ") {
		if strings.Contains(lower, "fake") || strings.Contains(lower, "mock") || strings.Contains(lower, "stub") || strings.Contains(lower, "placeholder") {
			return true
		}
	}
	return false
}

func (FakeImplCheckPlugin) Run(ctx context.Context, target string) TestResult {
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
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}
		rel, _ := filepath.Rel(target, path)
		if strings.HasPrefix(rel, "test/") {
			return nil
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			if isFakeMarker(line) {
				totalHits++
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
		Name:     "fake-impl-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(details, "; "),
		Metrics:  map[string]float64{"hits": float64(totalHits)},
	}
}

func (FakeImplCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(FakeImplCheckPlugin{})
}
