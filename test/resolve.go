// Abate —— 止息混沌 · 导入声明解析（精层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"os"
	"path/filepath"
	"strings"
)

// resolveImport 将一条导入声明解析为项目内文件列表。
// 策略优先级：相对路径 > 模块前缀剥离 > 目录名后缀匹配。
func resolveImport(target, from, imp string, dirIndex map[string][]string, module string) []string {
	imp = strings.TrimSpace(imp)
	if imp == "" {
		return nil
	}
	// 策略 1：相对路径导入（./ ../）— JS/TS/Rust/Python 常见。
	if strings.HasPrefix(imp, ".") {
		return resolveRelative(from, imp)
	}
	// 策略 2：Go 模块前缀剥离 — module/foo/bar → foo/bar。
	if module != "" && strings.HasPrefix(imp, module+"/") {
		tail := strings.TrimPrefix(imp, module+"/")
		return resolveByDirIndex(target, tail, dirIndex)
	}
	// 策略 2b：Go 模块自身（imp == module）。
	if module != "" && imp == module {
		return dirIndex["."]
	}
	// 策略 3：Java/Kotlin 包路径 — com.foo.bar → 尝试 com/foo/bar 目录。
	if strings.Contains(imp, ".") && !strings.Contains(imp, "/") {
		slash := strings.ReplaceAll(imp, ".", "/")
		return resolveByDirIndex(target, slash, dirIndex)
	}
	// 策略 4：路径式导入 — 直接按目录索引匹配。
	return resolveByDirIndex(target, imp, dirIndex)
}

// resolveRelative 解析相对路径导入，尝试补扩展名和目录匹配。
func resolveRelative(from, imp string) []string {
	base := filepath.Dir(from)
	p := filepath.Clean(filepath.Join(base, filepath.FromSlash(imp)))
	// 精确文件匹配。
	if IsCodeFile(p) {
		if _, err := os.Stat(p); err == nil {
			return []string{p}
		}
	}
	// 尝试补扩展名。
	for _, ext := range []string{".go", ".kt", ".java", ".ts", ".tsx", ".js", ".jsx", ".py", ".rs", ".swift", ".rb", ".php", ".cs", ".scala", ".dart", ".c", ".cpp", ".h", ".hpp"} {
		q := p + ext
		if _, err := os.Stat(q); err == nil {
			return []string{q}
		}
	}
	// 尝试 index.ts / index.js 等约定。
	for _, idx := range []string{"index.ts", "index.tsx", "index.js", "index.jsx"} {
		q := filepath.Join(p, idx)
		if _, err := os.Stat(q); err == nil {
			return []string{q}
		}
	}
	return nil
}

// resolveByDirIndex 按目录索引匹配导入路径。
// 对于 Kotlin/Java，导入串含类名（com.foo.Bar），需剥离末尾段后再匹配目录。
func resolveByDirIndex(target, imp string, dirIndex map[string][]string) []string {
	rel := filepath.ToSlash(imp)
	// 精确目录匹配。
	if fs, ok := dirIndex[rel]; ok {
		return fs
	}
	// 剥离末尾段（类名）后匹配目录。
	pkgPath := rel
	if idx := strings.LastIndex(rel, "/"); idx > 0 {
		pkgPath = rel[:idx]
	}
	if fs, ok := dirIndex[pkgPath]; ok {
		return fs
	}
	// 后缀匹配：imp 为 foo/bar，索引中有 pkg/foo/bar。
	var out []string
	for dir, fs := range dirIndex {
		if dir == rel || dir == pkgPath ||
			strings.HasSuffix(dir, "/"+rel) ||
			strings.HasSuffix(dir, "/"+pkgPath) {
			out = append(out, fs...)
		}
	}
	return out
}
