// Abate —— 止息混沌 · 文件头检查插件
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

// FileHeaderCheckPlugin 检查每个 .go 文件头是否包含版权声明和 AGPLv3 声明。
type FileHeaderCheckPlugin struct{}

func (FileHeaderCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "file-header-check", Version: "0.1.0", Protocol: "1"}
}

func (FileHeaderCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (FileHeaderCheckPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	totalFiles := 0
	missingHeader := 0
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
		totalFiles++
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		header := string(data)
		if len(header) > 512 {
			header = header[:512]
		}
		hasCopyright := strings.Contains(header, "Copyright")
		hasAGPL := strings.Contains(header, "AGPL")
		if !hasCopyright || !hasAGPL {
			missingHeader++
			rel, _ := filepath.Rel(target, path)
			if len(details) < 16 {
				details = append(details, rel)
			}
		}
		return nil
	})

	dur := time.Since(start)
	status := StatusPass
	if missingHeader > 0 {
		status = StatusFail
	}
	return TestResult{
		Name:     "file-header-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(details, "; "),
		Metrics:  map[string]float64{"total": float64(totalFiles), "missing": float64(missingHeader)},
	}
}

func (FileHeaderCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(FileHeaderCheckPlugin{})
}
