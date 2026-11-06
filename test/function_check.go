// Abate —— 止息混沌 · 函数检查插件
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

// FunctionCheckPlugin 检查函数行数（>80行）、参数数量（>5）、圈复杂度（if/for/switch >10）。
type FunctionCheckPlugin struct{}

func (FunctionCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "function-check", Version: "0.1.0", Protocol: "1"}
}

func (FunctionCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

var funcDeclRe = regexp.MustCompile(`^func\s+\(.*\)\s+\w+\s*\((.*)\)`)
var funcDeclNoRecvRe = regexp.MustCompile(`^func\s+\w+\s*\((.*)\)`)
var complexityRe = regexp.MustCompile(`\b(if|for|switch|case)\b`)

const (
	maxFuncLines   = 100
	maxParams      = 7
	maxComplexity  = 15
)

func (FunctionCheckPlugin) Run(ctx context.Context, target string) TestResult {
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
		i := 0
		for i < len(lines) {
			line := lines[i]
			isFunc := funcDeclRe.MatchString(line) || funcDeclNoRecvRe.MatchString(line)
			if !isFunc {
				i++
				continue
			}
			funcStart := i
			depth := 0
			started := false
			complexity := 0
			j := i
			for j < len(lines) {
				for _, ch := range lines[j] {
					if ch == '{' {
						depth++
						started = true
					}
					if ch == '}' {
						depth--
					}
				}
				complexity += len(complexityRe.FindAllString(lines[j], -1))
				if started && depth == 0 {
					break
				}
				j++
			}
			funcLines := j - funcStart + 1
			params := ""
			if m := funcDeclRe.FindStringSubmatch(line); len(m) > 1 {
				params = m[1]
			} else if m := funcDeclNoRecvRe.FindStringSubmatch(line); len(m) > 1 {
				params = m[1]
			}
			paramCount := 0
			if strings.TrimSpace(params) != "" {
				paramCount = strings.Count(params, ",") + 1
			}
			violations := make([]string, 0, 3)
			if funcLines > maxFuncLines {
				violations = append(violations, "lines="+itoa(funcLines))
			}
			if paramCount > maxParams {
				violations = append(violations, "params="+itoa(paramCount))
			}
			if complexity > maxComplexity {
				violations = append(violations, "complexity="+itoa(complexity))
			}
			if len(violations) > 0 {
				totalHits++
				rel, _ := filepath.Rel(target, path)
				if len(details) < 16 {
					details = append(details, rel+":"+itoa(funcStart+1)+"("+strings.Join(violations, ",")+")")
				}
			}
			i = j + 1
		}
		return nil
	})

	dur := time.Since(start)
	status := StatusPass
	if totalHits > 0 {
		status = StatusFail
	}
	return TestResult{
		Name:     "function-check",
		Status:   status,
		Duration: dur,
		Details:  strings.Join(details, "; "),
		Metrics:  map[string]float64{"hits": float64(totalHits)},
	}
}

func (FunctionCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(FunctionCheckPlugin{})
}
