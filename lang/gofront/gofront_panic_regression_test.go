/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"go/ast"
	"testing"

	"abate/lang/air"
)

// TestIsNilASTNodeDetectsTypedNil proves plain nil and typed-nil AST pointers are rejected while real nodes pass.
func TestIsNilASTNodeDetectsTypedNil(t *testing.T) {
	t.Parallel()
	var empty ast.Node
	if !isNilASTNode(empty) {
		t.Fatal("plain nil ast.Node must be treated as nil")
	}
	var block *ast.BlockStmt
	if !isNilASTNode(block) {
		t.Fatal("typed-nil *ast.BlockStmt must be treated as nil")
	}
	if isNilASTNode(&ast.BlockStmt{}) {
		t.Fatal("a real AST node must not be treated as nil")
	}
}

// TestTruncatedGoSourcesReturnDiagnosticsOrTree proves two truncate-half gVisor-shaped inputs return a diagnostic tree or an explicit parse error instead of panicking.
func TestTruncatedGoSourcesReturnDiagnosticsOrTree(t *testing.T) {
	cases := []struct {
		name string
		path string
		src  string
	}{
		{"replica", "replica.go", "package replica\n\nfunc (r *Replica) syncFrom(peer string) error"},
		{"stats_test", "stats_test.go", "package stats\n\nfunc TestReplicaStats(t *testing.T)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source := air.SourceUnit{
				Language: "go",
				Files: []air.SourceFile{
					air.NewSourceFile(tc.path, []byte(tc.src)),
				},
			}
			unit, diagnostics, err := NewGoFrontend().Parse(source)
			if err != nil {
				return
			}
			if unit == nil {
				t.Fatal("Parse returned a nil unit without an error")
			}
			if len(unit.Roots) != 1 {
				t.Fatalf("Parse returned %d roots, want 1", len(unit.Roots))
			}
			if len(diagnostics) == 0 && unit.NodeCount() == 0 {
				t.Fatal("Parse returned neither diagnostics nor a diagnostic AIR tree")
			}
		})
	}
}
