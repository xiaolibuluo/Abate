// Abate —— 止息混沌 · 占位检查插件
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

// PlaceholderCheckPlugin 扫描源码中的占位标记（TODO/FIXME/HACK/XXX/PLACEHOLDER）。
type PlaceholderCheckPlugin struct{}

func (PlaceholderCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "placeholder-check", Version: "0.1.0", Protocol: "1"}
}

func (PlaceholderCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

var placeholderRe = regexp.MustCompile(`(?i)\b(TODO|FIXME|HACK|XXX|PLACEHOLDER)\b`)

func (PlaceholderCheckPlugin) Run(ctx context.Context, target string) TestResult {
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
			if placeholderRe.MatchString(line) {
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
		Name:     "placeholder-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(details, "; "),
		Metrics:  map[string]float64{"hits": float64(totalHits)},
	}
}

func (PlaceholderCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(PlaceholderCheckPlugin{})
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
