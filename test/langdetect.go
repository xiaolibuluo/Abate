// Abate —— 止息混沌 · 全语言检测工具
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"os"
	"path/filepath"
	"strings"
)

// SupportedCodeExtensions 支持的代码文件扩展名
var SupportedCodeExtensions = map[string]bool{
	".go":   true,
	".kt":   true,
	".java": true,
	".py":   true,
	".js":   true,
	".ts":   true,
	".tsx":  true,
	".jsx":  true,
	".rs":   true,
	".c":    true,
	".cpp":  true,
	".h":    true,
	".hpp":  true,
	".swift": true,
	".rb":   true,
	".php":  true,
	".cs":   true,
	".scala": true,
	".dart": true,
}

// IsCodeFile 判断是否为支持的代码文件
func IsCodeFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return SupportedCodeExtensions[ext]
}

// IsTestFile 判断是否为测试文件（跨语言）
func IsTestFile(path string) bool {
	base := strings.ToLower(filepath.Base(path))
	if strings.HasSuffix(base, "_test.go") {
		return true
	}
	if strings.HasSuffix(base, "test.kt") {
		return true
	}
	if strings.HasSuffix(base, "test.java") {
		return true
	}
	if strings.HasPrefix(base, "test_") && strings.HasSuffix(base, ".py") {
		return true
	}
	if strings.HasSuffix(base, ".test.js") || strings.HasSuffix(base, ".spec.js") {
		return true
	}
	if strings.HasSuffix(base, ".test.ts") || strings.HasSuffix(base, ".spec.ts") {
		return true
	}
	if strings.HasSuffix(base, "_test.rs") {
		return true
	}
	if strings.HasSuffix(base, "tests.rs") {
		return true
	}
	return false
}

// ProjectType 项目类型
 type ProjectType int

const (
	ProjectUnknown ProjectType = iota
	ProjectGo
	ProjectGradle
	ProjectMaven
	ProjectNpm
	ProjectCargo
	ProjectPython
)

// DetectProjectType 自动检测项目类型
func DetectProjectType(target string) ProjectType {
	if fileExists(filepath.Join(target, "go.mod")) {
		return ProjectGo
	}
	if fileExists(filepath.Join(target, "settings.gradle.kts")) || fileExists(filepath.Join(target, "settings.gradle")) {
		return ProjectGradle
	}
	if fileExists(filepath.Join(target, "pom.xml")) {
		return ProjectMaven
	}
	if fileExists(filepath.Join(target, "package.json")) {
		return ProjectNpm
	}
	if fileExists(filepath.Join(target, "Cargo.toml")) {
		return ProjectCargo
	}
	if fileExists(filepath.Join(target, "requirements.txt")) || fileExists(filepath.Join(target, "setup.py")) || fileExists(filepath.Join(target, "pyproject.toml")) {
		return ProjectPython
	}
	return ProjectUnknown
}

// BuildCommand 返回项目类型对应的构建命令
func BuildCommand(pt ProjectType) []string {
	switch pt {
	case ProjectGo:
		return []string{"go", "build", "./..."}
	case ProjectGradle:
		return []string{"gradle", "compileDebugKotlin", "--no-daemon"}
	case ProjectMaven:
		return []string{"mvn", "compile", "-q"}
	case ProjectNpm:
		return []string{"npm", "run", "build"}
	case ProjectCargo:
		return []string{"cargo", "build"}
	case ProjectPython:
		return []string{"python3", "-m", "compileall", "."}
	default:
		return nil
	}
}

// VetCommand 返回项目类型对应的静态检查命令
func VetCommand(pt ProjectType) []string {
	switch pt {
	case ProjectGo:
		return []string{"go", "vet", "./..."}
	case ProjectGradle:
		return []string{"gradle", "lint", "--no-daemon"}
	case ProjectMaven:
		return []string{"mvn", "checkstyle:check", "-q"}
	case ProjectNpm:
		return []string{"npm", "run", "lint"}
	case ProjectCargo:
		return []string{"cargo", "clippy"}
	case ProjectPython:
		return []string{"python3", "-m", "pyflakes", "."}
	default:
		return nil
	}
}

// TestCommand 返回项目类型对应的测试命令
func TestCommand(pt ProjectType) []string {
	switch pt {
	case ProjectGo:
		return []string{"go", "test", "./..."}
	case ProjectGradle:
		return []string{"gradle", "test", "--no-daemon"}
	case ProjectMaven:
		return []string{"mvn", "test", "-q"}
	case ProjectNpm:
		return []string{"npm", "test"}
	case ProjectCargo:
		return []string{"cargo", "test"}
	case ProjectPython:
		return []string{"python3", "-m", "pytest"}
	default:
		return nil
	}
}

// ShouldSkipDir 判断是否应跳过该目录
func ShouldSkipDir(name string) bool {
	return name == ".git" || name == "vendor" || name == "node_modules" ||
		name == "build" || name == ".gradle" || name == "target" ||
		name == "__pycache__" || name == ".idea" || name == ".vscode"
}

// ShouldSkipFile 判断是否应跳过该文件（测试文件或插件自身）
func ShouldSkipFile(path string, target string) bool {
	if IsTestFile(path) {
		return true
	}
	rel, err := filepath.Rel(target, path)
	if err != nil {
		return false
	}
	if strings.HasPrefix(rel, "test/") {
		return true
	}
	return false
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
