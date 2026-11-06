// Abate —— 止息混沌 · 文件级依赖图构建（精层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// DepGraph 是项目内文件级依赖图：Edges[a] 为 a 直接依赖的文件集合。
type DepGraph struct {
	Root   string
	Files  []string
	Edges  map[string]map[string]bool
	Parsed int
}

// BuildDepGraph 遍历项目构建文件级依赖图；maxFiles<=0 不限扫描数量。
func BuildDepGraph(target string, maxFiles int) *DepGraph {
	g := &DepGraph{Root: target, Edges: map[string]map[string]bool{}}
	files := collectCodeFiles(target, maxFiles)
	sort.Strings(files)
	g.Files = files
	dirIndex := indexDirs(target, files)
	module := detectModuleName(target)
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		g.Parsed++
		for _, imp := range ExtractImports(f, string(data)) {
			for _, dep := range resolveImport(target, f, imp, dirIndex, module) {
				if dep == f {
					continue
				}
				if g.Edges[f] == nil {
					g.Edges[f] = map[string]bool{}
				}
				g.Edges[f][dep] = true
			}
		}
	}
	return g
}

// collectCodeFiles 收集项目内非测试代码文件。
func collectCodeFiles(target string, maxFiles int) []string {
	var files []string
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
		if !IsCodeFile(path) || IsTestFile(path) {
			return nil
		}
		if maxFiles > 0 && len(files) >= maxFiles {
			return filepath.SkipAll
		}
		files = append(files, path)
		return nil
	})
	return files
}

// indexDirs 建立目录相对路径（斜杠分隔）到文件的索引。
func indexDirs(target string, files []string) map[string][]string {
	idx := map[string][]string{}
	for _, f := range files {
		rel, err := filepath.Rel(target, f)
		if err != nil {
			continue
		}
		dir := filepath.ToSlash(filepath.Dir(rel))
		idx[dir] = append(idx[dir], f)
	}
	return idx
}

// detectModuleName 从 go.mod 读取模块名（非 Go 项目返回空串）。
func detectModuleName(target string) string {
	data, err := os.ReadFile(filepath.Join(target, "go.mod"))
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(t, "module "))
		}
	}
	return ""
}
