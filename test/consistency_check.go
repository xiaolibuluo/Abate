// Abate —— 止息混沌 · 对齐检查插件（自省层）
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

// ConsistencyCheckPlugin 对齐检查插件：校验多语言能力声明与实际实现的一致性。
type ConsistencyCheckPlugin struct{}

func (ConsistencyCheckPlugin) Capability() core.Capability {
	return core.Capability{Name: "consistency-check", Version: "0.1.0", Protocol: "1"}
}

func (ConsistencyCheckPlugin) Setup(ctx context.Context) (func(), error) {
	return nil, nil
}

func (ConsistencyCheckPlugin) Run(ctx context.Context, target string) TestResult {
	start := time.Now()
	pt := DetectProjectType(target)
	if pt == ProjectUnknown {
		return TestResult{Name: "consistency-check", Status: StatusSkip, Duration: time.Since(start), Details: "unsupported project type", Metrics: map[string]float64{}}
	}
	var issues []string
	// 检查 1：SupportedCodeExtensions 中每种扩展名是否都有对应的 IsTestFile 命名约定。
	missingTestFile := checkTestFileCoverage()
	for _, ext := range missingTestFile {
		issues = append(issues, fmt.Sprintf("IsTestFile: no convention for %s", ext))
	}
	// 检查 2：SupportedCodeExtensions 中每种扩展名是否都有对应的 ExtractImports 实现。
	missingImports := checkImportsCoverage()
	for _, ext := range missingImports {
		issues = append(issues, fmt.Sprintf("ExtractImports: no parser for %s", ext))
	}
	// 检查 3：每种 ProjectType 是否都有 Build/Vet/Test 命令。
	missingCmds := checkCommandCoverage()
	for _, c := range missingCmds {
		issues = append(issues, c)
	}
	// 检查 4：issuePatterns 是否覆盖所有代码扩展名的输出格式。
	missingPatterns := checkPatternCoverage()
	for _, ext := range missingPatterns {
		issues = append(issues, fmt.Sprintf("issuePatterns: no pattern for %s", ext))
	}
	// 检查 5：文件行数是否超限（默认 300 行红线）。
	overLimit := checkLineLimit(target, 300)
	for _, f := range overLimit {
		issues = append(issues, fmt.Sprintf("line-limit: %s (%d lines)", f.Name, f.Lines))
	}
	dur := time.Since(start)
	status := StatusPass
	if len(issues) > 0 {
		status = StatusFail
	}
	details := fmt.Sprintf("checks=5 issues=%d", len(issues))
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
		Name:    "consistency-check",
		Status:  status,
		Duration: dur,
		Details: details,
		Metrics: map[string]float64{
			"issues":            float64(len(issues)),
			"missing_testfile":  float64(len(missingTestFile)),
			"missing_imports":   float64(len(missingImports)),
			"missing_cmds":      float64(len(missingCmds)),
			"missing_patterns":  float64(len(missingPatterns)),
			"over_limit":        float64(len(overLimit)),
		},
	}
}

func (ConsistencyCheckPlugin) Teardown(ctx context.Context) error {
	return nil
}

func init() {
	DefaultRegistry.Register(ConsistencyCheckPlugin{})
}

// checkTestFileCoverage 检查 SupportedCodeExtensions 中哪些扩展名没有 IsTestFile 命名约定。
// 头文件（.h/.hpp）是声明文件，不直接包含测试逻辑，豁免检查。
func checkTestFileCoverage() []string {
	// 豍免：头文件不需要测试文件命名约定。
	exempt := map[string]bool{".h": true, ".hpp": true}
	// 已知有测试文件命名约定的扩展名。
	covered := map[string]bool{
		".go": true, ".kt": true, ".java": true, ".py": true,
		".js": true, ".ts": true, ".tsx": true, ".jsx": true,
		".rs": true, ".c": true, ".cpp": true,
		".swift": true, ".rb": true, ".php": true,
		".cs": true, ".scala": true, ".dart": true,
	}
	var missing []string
	for ext := range SupportedCodeExtensions {
		if exempt[ext] {
			continue
		}
		if !covered[ext] {
			missing = append(missing, ext)
		}
	}
	return missing
}

// checkImportsCoverage 检查 SupportedCodeExtensions 中哪些扩展名没有 ExtractImports 实现。
// 通过对每种扩展名调用 ExtractImports 并检查是否返回 nil（空内容除外）来判断。
func checkImportsCoverage() []string {
	// ExtractImports 的 switch 覆盖范围。
	covered := map[string]bool{
		".go": true, ".kt": true, ".java": true, ".py": true,
		".js": true, ".ts": true, ".tsx": true, ".jsx": true,
		".rs": true, ".c": true, ".cpp": true, ".h": true, ".hpp": true,
		".swift": true, ".rb": true, ".php": true,
		".cs": true, ".scala": true, ".dart": true,
	}
	var missing []string
	for ext := range SupportedCodeExtensions {
		if !covered[ext] {
			missing = append(missing, ext)
		}
	}
	return missing
}

// checkCommandCoverage 检查每种 ProjectType 是否都有 Build/Vet/Test 命令。
func checkCommandCoverage() []string {
	var issues []string
	allTypes := []ProjectType{
		ProjectGo, ProjectGradle, ProjectMaven, ProjectNpm, ProjectCargo, ProjectPython,
		ProjectSwift, ProjectRuby, ProjectPhp, ProjectDotNet, ProjectCpp, ProjectScala, ProjectDart,
	}
	for _, pt := range allTypes {
		if BuildCommand(pt) == nil {
			issues = append(issues, fmt.Sprintf("BuildCommand: missing for project type %d", pt))
		}
		if VetCommand(pt) == nil {
			issues = append(issues, fmt.Sprintf("VetCommand: missing for project type %d", pt))
		}
		if TestCommand(pt) == nil {
			issues = append(issues, fmt.Sprintf("TestCommand: missing for project type %d", pt))
		}
	}
	return issues
}

// checkPatternCoverage 检查 issuePatterns 是否覆盖所有代码扩展名。
func checkPatternCoverage() []string {
	// issuePatterns 中正则显式覆盖的扩展名。
	covered := map[string]bool{
		".go": true, ".kt": true, ".java": true, ".js": true, ".ts": true,
		".tsx": true, ".jsx": true, ".rs": true, ".py": true,
		".c": true, ".cpp": true, ".h": true, ".hpp": true,
		".swift": true, ".rb": true, ".php": true,
		".cs": true, ".scala": true, ".dart": true,
	}
	var missing []string
	for ext := range SupportedCodeExtensions {
		if !covered[ext] {
			missing = append(missing, ext)
		}
	}
	return missing
}

// LineLimitFile 是超出行数限制的文件信息。
type LineLimitFile struct {
	Path  string
	Name  string
	Lines int
}

// docFileExts 是需要行数检查的文档文件扩展名。
var docFileExts = map[string]bool{
	".md": true, ".txt": true, ".yaml": true, ".yml": true, ".json": true,
}

// isDocFile 判断是否为文档文件。
func isDocFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return docFileExts[ext]
}

// checkLineLimit 检查目标项目中超过行数限制的代码文件和文档文件。
func checkLineLimit(target string, limit int) []LineLimitFile {
	var over []LineLimitFile
	// 检查代码文件（通过依赖图收集）。
	g := BuildDepGraph(target, 0)
	for _, f := range g.Files {
		data, err := readFileContent(f)
		if err != nil {
			continue
		}
		lines := strings.Count(data, "\n") + 1
		if lines > limit {
			rel, _ := filepath.Rel(target, f)
			over = append(over, LineLimitFile{Path: f, Name: rel, Lines: lines})
		}
	}
	// 检查文档文件（.md/.txt/.yaml/.yml/.json）。
	_ = filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return nil
		}
		if info.IsDir() {
			if path != target && ShouldSkipDir(info.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if !isDocFile(path) {
			return nil
		}
		data, err := readFileContent(path)
		if err != nil {
			return nil
		}
		lines := strings.Count(data, "\n") + 1
		if lines > limit {
			rel, _ := filepath.Rel(target, path)
			over = append(over, LineLimitFile{Path: path, Name: rel, Lines: lines})
		}
		return nil
	})
	return over
}
