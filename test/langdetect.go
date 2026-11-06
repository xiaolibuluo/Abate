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
	// Go
	if strings.HasSuffix(base, "_test.go") {
		return true
	}
	// Kotlin / Java
	if strings.HasSuffix(base, "test.kt") || strings.HasSuffix(base, "test.java") {
		return true
	}
	// Python
	if strings.HasPrefix(base, "test_") && strings.HasSuffix(base, ".py") {
		return true
	}
	if strings.HasSuffix(base, "_test.py") {
		return true
	}
	// JS / TS / TSX / JSX
	for _, ext := range []string{".js", ".ts", ".tsx", ".jsx"} {
		if strings.HasSuffix(base, ".test"+ext) || strings.HasSuffix(base, ".spec"+ext) {
			return true
		}
	}
	// Rust
	if strings.HasSuffix(base, "_test.rs") || strings.HasSuffix(base, "tests.rs") {
		return true
	}
	// C / C++ — Google Test 约定
	if strings.HasSuffix(base, "_test.c") || strings.HasSuffix(base, "_test.cpp") {
		return true
	}
	// Swift — XCTest 约定
	if strings.HasSuffix(base, "tests.swift") || strings.HasSuffix(base, "_test.swift") {
		return true
	}
	// Ruby — RSpec / Minitest 约定
	if strings.HasSuffix(base, "_test.rb") || strings.HasSuffix(base, "_spec.rb") {
		return true
	}
	// PHP — PHPUnit 约定
	if strings.HasSuffix(base, "test.php") {
		return true
	}
	// C# — xUnit / NUnit 约定
	if strings.HasSuffix(base, "test.cs") {
		return true
	}
	// Scala — ScalaTest 约定
	if strings.HasSuffix(base, "test.scala") || strings.HasSuffix(base, "spec.scala") {
		return true
	}
	// Dart — flutter_test 约定
	if strings.HasSuffix(base, "_test.dart") {
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
	ProjectSwift     // Package.swift
	ProjectRuby      // Gemfile
	ProjectPhp       // composer.json
	ProjectDotNet    // *.csproj / *.sln
	ProjectCpp       // CMakeLists.txt / Makefile
	ProjectScala     // build.sbt
	ProjectDart      // pubspec.yaml
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
	if fileExists(filepath.Join(target, "Package.swift")) {
		return ProjectSwift
	}
	if fileExists(filepath.Join(target, "Gemfile")) {
		return ProjectRuby
	}
	if fileExists(filepath.Join(target, "composer.json")) {
		return ProjectPhp
	}
	if hasFileWithExt(target, ".csproj") || hasFileWithExt(target, ".sln") {
		return ProjectDotNet
	}
	if fileExists(filepath.Join(target, "CMakeLists.txt")) || fileExists(filepath.Join(target, "Makefile")) {
		return ProjectCpp
	}
	if fileExists(filepath.Join(target, "build.sbt")) {
		return ProjectScala
	}
	if fileExists(filepath.Join(target, "pubspec.yaml")) {
		return ProjectDart
	}
	return ProjectUnknown
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

// hasFileWithExt 检查目录下是否存在指定扩展名的文件（仅扫一层）。
func hasFileWithExt(target, ext string) bool {
	entries, err := os.ReadDir(target)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(strings.ToLower(e.Name()), ext) {
			return true
		}
	}
	return false
}
