// Abate —— 止息混沌 · 逻辑检查插件
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/qingdujian/abate/core"
)

// LogicCheckPlugin 扫描源码中的逻辑缺陷模式（忽略 error、恒真恒假条件、空 if 体）。
type LogicCheckPlugin struct{}

func (LogicCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "logic-check", Version: "0.1.0", Protocol: "1"}
}

func (LogicCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

var (
	ignoredErrRe  = regexp.MustCompile(`^\s*_\s*=\s*\w+\(.*\)$`)
	alwaysTrueRe  = regexp.MustCompile(`if\s+(true|1\s*==\s*1|1\s*!=\s*0)\s*\{`)
	alwaysFalseRe = regexp.MustCompile(`if\s+(false|0\s*==\s*1|1\s*==\s*0)\s*\{`)
	emptyIfRe     = regexp.MustCompile(`\{\s*\}`)
)

func (LogicCheckPlugin) Run(ctx context.Context, target string) TestResult {
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
			matched := false
			if ignoredErrRe.MatchString(line) {
				matched = true
			}
			if alwaysTrueRe.MatchString(line) {
				matched = true
			}
			if alwaysFalseRe.MatchString(line) {
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
		Name:     "logic-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(details, "; "),
		Metrics:  map[string]float64{"hits": float64(totalHits)},
	}
}

func (LogicCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(LogicCheckPlugin{})
}
