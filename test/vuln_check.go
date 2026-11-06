// Abate —— 止息混沌 · 漏洞检查插件
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

// VulnCheckPlugin 静态安全扫描：硬编码密钥、SQL 拼接、命令注入模式检测。
type VulnCheckPlugin struct{}

func (VulnCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "vuln-check", Version: "0.1.0", Protocol: "1"}
}

func (VulnCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

var (
	hardcodedSecretRe = regexp.MustCompile(`(?i)(password|secret|token|apikey|api_key)\s*[:=]\s*"[^"]{8,}"`)
	sqlConcatRe       = regexp.MustCompile(`fmt\.Sprintf\s*\(\s*"SELECT`)
	cmdExecRe         = regexp.MustCompile(`exec\.Command\s*\(\s*"sh"\s*,\s*"-c"\s*,\s*fmt`)
	evalRe            = regexp.MustCompile(`(?i)\beval\s*\(`)
)

func (VulnCheckPlugin) Run(ctx context.Context, target string) TestResult {
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
			if hardcodedSecretRe.MatchString(line) {
				matched = true
			}
			if sqlConcatRe.MatchString(line) {
				matched = true
			}
			if cmdExecRe.MatchString(line) {
				matched = true
			}
			if evalRe.MatchString(line) {
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
		Name:     "vuln-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(details, "; "),
		Metrics:  map[string]float64{"hits": float64(totalHits)},
	}
}

func (VulnCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(VulnCheckPlugin{})
}
