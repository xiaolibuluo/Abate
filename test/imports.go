// Abate —— 止息混沌 · 多语言导入提取（精层）
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

import (
	"path/filepath"
	"strings"
)

// ExtractImports 按文件语言提取导入声明（原始模块串）。
// 覆盖 SupportedCodeExtensions 声明的全部 19 种扩展名。
func ExtractImports(file, content string) []string {
	switch strings.ToLower(filepath.Ext(file)) {
	case ".go":
		return parseGoImports(content)
	case ".kt", ".java":
		return parseJvmImports(content)
	case ".js", ".ts", ".tsx", ".jsx":
		return parseJsImports(content)
	case ".py":
		return parsePyImports(content)
	case ".rs":
		return parseRustImports(content)
	case ".c", ".cpp", ".h", ".hpp":
		return parseCImports(content)
	case ".swift":
		return parseSwiftImports(content)
	case ".rb":
		return parseRubyImports(content)
	case ".php":
		return parsePhpImports(content)
	case ".cs":
		return parseCsImports(content)
	case ".scala":
		return parseScalaImports(content)
	case ".dart":
		return parseDartImports(content)
	default:
		return nil
	}
}

// parseGoImports 解析 go 源文件 import 声明（含块导入）。
func parseGoImports(content string) []string {
	var out []string
	inBlock := false
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if inBlock {
			if t == ")" {
				inBlock = false
				continue
			}
			if s := quoted(t); s != "" {
				out = append(out, s)
			}
			continue
		}
		if strings.HasPrefix(t, "import") && strings.HasSuffix(t, "(") {
			inBlock = true
			continue
		}
		if strings.HasPrefix(t, "import ") {
			if s := quoted(t); s != "" {
				out = append(out, s)
			}
		}
	}
	return out
}

// parseJvmImports 解析 Kotlin/Java import 声明。
func parseJvmImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if !strings.HasPrefix(t, "import ") {
			continue
		}
		t = strings.TrimPrefix(t, "import ")
		t = strings.TrimPrefix(t, "static ")
		t = strings.TrimSpace(strings.TrimSuffix(t, ";"))
		if t != "" {
			out = append(out, t)
		}
	}
	return out
}

// parseJsImports 解析 JS/TS 的 import/require/from 声明。
func parseJsImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if idx := strings.LastIndex(t, " from "); idx >= 0 {
			if s := quoted(t[idx+len(" from "):]); s != "" {
				out = append(out, s)
				continue
			}
		}
		if strings.Contains(t, "require(") || strings.HasPrefix(t, "import ") {
			if s := quoted(t); s != "" {
				out = append(out, s)
			}
		}
	}
	return out
}

// parsePyImports 解析 Python import/from 声明。
func parsePyImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "import ") || strings.HasPrefix(t, "from ") {
			if fs := strings.Fields(t); len(fs) > 1 {
				out = append(out, strings.Trim(fs[1], ","))
			}
		}
	}
	return out
}

// parseRustImports 解析 Rust use/mod 声明。
func parseRustImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "use ") {
			t = strings.TrimSuffix(strings.TrimPrefix(t, "use "), ";")
			if t = strings.TrimSpace(t); t != "" {
				out = append(out, t)
			}
			continue
		}
		if strings.HasPrefix(t, "mod ") {
			t = strings.TrimSuffix(strings.TrimPrefix(t, "mod "), ";")
			if t = strings.TrimSpace(t); t != "" {
				out = append(out, t)
			}
		}
	}
	return out
}

// parseCImports 解析 C/C++ #include 声明（引号内为项目内依赖）。
func parseCImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "#include") {
			if s := quoted(t); s != "" {
				out = append(out, s)
			}
		}
	}
	return out
}

// parseSwiftImports 解析 Swift import 声明。
func parseSwiftImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "import ") {
			if t = strings.TrimSpace(strings.TrimPrefix(t, "import ")); t != "" {
				out = append(out, t)
			}
		}
	}
	return out
}

// parseRubyImports 解析 Ruby require/require_relative 声明。
func parseRubyImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		for _, kw := range []string{"require_relative ", "require "} {
			if strings.HasPrefix(t, kw) {
				if s := quoted(t); s != "" {
					out = append(out, s)
				}
				break
			}
		}
	}
	return out
}

// parsePhpImports 解析 PHP use/require/include 声明。
func parsePhpImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "use ") {
			t = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(t, "use "), ";"))
			if t != "" {
				out = append(out, t)
			}
			continue
		}
		for _, kw := range []string{"require_once ", "include_once ", "require ", "include "} {
			if strings.HasPrefix(t, kw) {
				if s := quoted(t); s != "" {
					out = append(out, s)
				}
				break
			}
		}
	}
	return out
}

// parseCsImports 解析 C# using 声明。
func parseCsImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "using ") {
			t = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(t, "using "), ";"))
			if t != "" && !strings.Contains(t, " ") {
				out = append(out, t)
			}
		}
	}
	return out
}

// parseScalaImports 解析 Scala import 声明（含通配符 _ 后缀）。
func parseScalaImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		if strings.HasPrefix(t, "import ") {
			t = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(t, "import "), "_"))
			if t != "" {
				out = append(out, t)
			}
		}
	}
	return out
}

// parseDartImports 解析 Dart import/export/part 声明。
func parseDartImports(content string) []string {
	var out []string
	for _, line := range strings.Split(content, "\n") {
		t := strings.TrimSpace(line)
		for _, kw := range []string{"import ", "export ", "part "} {
			if strings.HasPrefix(t, kw) {
				if s := quoted(t); s != "" {
					out = append(out, s)
				}
				break
			}
		}
	}
	return out
}

// quoted 提取串中第一段引号内容（双引号优先）。
func quoted(s string) string {
	for _, q := range [2]byte{'"', '\''} {
		i := strings.IndexByte(s, q)
		if i < 0 {
			continue
		}
		j := strings.IndexByte(s[i+1:], q)
		if j < 0 {
			continue
		}
		return s[i+1 : i+1+j]
	}
	return ""
}
