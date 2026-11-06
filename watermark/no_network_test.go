// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package watermark

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestNoNetworkCalls 以 AST 静态检查卡死：watermark 包禁止任何网络调用
// 与文件 I/O —— 水印与亮牌为纯内存计算，禁止联网心跳校验（任务四）。
//
// TestNoNetworkCalls statically enforces via AST that the watermark package
// performs no network calls and no file I/O — watermark and banner are pure
// in-memory computation; phone-home heartbeat checks are forbidden.
func TestNoNetworkCalls(t *testing.T) {
	// 禁止导入的网络与 I/O 包前缀。
	// Forbidden network and I/O package import prefixes.
	forbiddenImports := []string{
		"net/http", "net", "os", "io", "crypto/tls",
	}
	// 禁止出现的调用选择子（防御性：即使导入被绕过也拦截）。
	// Forbidden selector calls (defensive: caught even if imports slip in).
	forbiddenCalls := []string{
		"http.Get", "http.Post", "http.Head", "http.NewRequest",
		"net.Dial", "tls.Dial", "os.Open", "os.ReadFile", "os.WriteFile",
		"os.Create", "os.OpenFile", "os.Remove", "os.Stat",
	}

	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("readdir: %v", err)
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		path := filepath.Join(".", e.Name())
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, src, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		for _, imp := range file.Imports {
			ipath := strings.Trim(imp.Path.Value, "\"")
		for _, p := range forbiddenImports {
			if ipath == p {
				t.Errorf("%s: forbidden import %q (no network / no I/O)", path, ipath)
			}
		}
		}
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
		}
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			var sb strings.Builder
			if x, ok := sel.X.(*ast.Ident); ok {
				sb.WriteString(x.Name)
				sb.WriteString(".")
			}
			sb.WriteString(sel.Sel.Name)
			name := sb.String()
			for _, f := range forbiddenCalls {
				if name == f {
					t.Errorf("%s: forbidden call %s (no network / no I/O)", path, name)
				}
			}
		}
		return true
		})
	}
}
