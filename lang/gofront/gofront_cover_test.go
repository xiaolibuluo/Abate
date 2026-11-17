/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"errors"
	"go/ast"
	"go/scanner"
	"go/token"
	"go/types"
	"testing"

	"abate/lang/air"
)

// emptyView builds a view over an empty syntax tree for boundary checks.
func emptyView() *GoView {
	return newGoView(&ast.File{}, token.NewFileSet(), nil)
}

// TestCoverViewBoundaries exercises nil and zero-value view branches.
func TestCoverViewBoundaries(t *testing.T) {
	t.Parallel()
	var nilView *GoView
	if nilView.PackageName() != "" || nilView.NodeCount() != 0 ||
		nilView.Imports() != nil {
		t.Fatal("nil view accessors must be inert")
	}
	view := emptyView()
	if view.PackageName() != "" || view.Imports() != nil || view.NodeCount() != 0 {
		t.Fatal("empty view accessors drifted")
	}
	if pos := view.position(token.NoPos); pos.Canonical() != "0/0/0/0" {
		t.Fatalf("NoPos conversion: %s", pos.Canonical())
	}
	if pos := view.span(token.NoPos, token.NoPos); pos.EndOffset != 0 {
		t.Fatalf("NoPos span: %s", pos.String())
	}
	if node := view.synthetic(air.KindRegion, "synthetic"); node == nil ||
		node.Kind() != air.KindRegion || air.AttributeOf(node) != "synthetic" {
		t.Fatal("synthetic node construction failed")
	}
	view.bind(nil, view.synthetic(air.KindRegion, "x"))
	view.bind(&ast.Ident{}, nil)
	if _, ok := view.lookup(nil); ok {
		t.Fatal("lookup must reject a nil key")
	}
	if _, ok := nilView.lookup(&ast.Ident{}); ok {
		t.Fatal("lookup on a nil view must fail")
	}
	nilView.bind(&ast.Ident{}, view.synthetic(air.KindRegion, "x"))
	if pos := nilView.position(token.Pos(1)); pos.Canonical() != "0/0/0/0" {
		t.Fatal("nil view position must be the zero position")
	}
	if span := nilView.span(token.Pos(1), token.Pos(2)); span.EndOffset != 0 {
		t.Fatal("nil view span must be zero-width")
	}
	if root := mapFile(view); root == nil || root.Kind() != air.KindPackage {
		t.Fatal("mapFile must tolerate a file without a package name")
	}
	if len(appendNode(nil, nil)) != 0 {
		t.Fatal("appendNode must filter nil nodes")
	}
}

// TestCoverDamagedSyntax drives the explicit degraded-mapping branches.
func TestCoverDamagedSyntax(t *testing.T) {
	t.Parallel()
	view := emptyView()
	badDecl := mapDecl(view, &ast.BadDecl{})
	if badDecl.Kind() != air.KindRegion ||
		badDecl.Sem(air.SemPartial) != "true" {
		t.Fatal("BadDecl was not mapped to a degraded node")
	}
	if badStmt := mapStmt(view, &ast.BadStmt{}); badStmt.Kind() != air.KindRegion {
		t.Fatal("BadStmt was not mapped to a degraded node")
	}
	if badExpr := mapExpr(view, &ast.BadExpr{}); badExpr.Kind() != air.KindRegion {
		t.Fatal("BadExpr was not mapped to a degraded node")
	}
	if badType := mapType(view, &ast.BadExpr{}); badType.Kind() != air.KindRegion {
		t.Fatal("bad type was not mapped to a degraded node")
	}
}

// TestCoverDiagnosticsHelpers covers non-ErrorList failures and caps.
func TestCoverDiagnosticsHelpers(t *testing.T) {
	t.Parallel()
	if got := collectDiagnostics("x.go", nil, 10); got != nil {
		t.Fatal("a nil parse error must produce no diagnostics")
	}
	if got := collectDiagnostics("x.go", errors.New("boom"), 0); got != nil {
		t.Fatal("a non-positive cap must produce no diagnostics")
	}
	plain := collectDiagnostics("x.go", errors.New("boom"), 10)
	if len(plain) != 1 || plain[0].Code != "LANG-GO-E-001" ||
		plain[0].Message != "boom" {
		t.Fatalf("non-ErrorList diagnostic: %+v", plain)
	}
	list := scanner.ErrorList{
		{Pos: token.Position{Filename: "x.go", Line: 1, Column: 2}, Msg: "first"},
		{Pos: token.Position{Filename: "x.go", Line: 2, Column: 3}, Msg: "second"},
	}
	capped := collectDiagnostics("x.go", list, 1)
	if len(capped) != 1 || capped[0].Line != 1 || capped[0].Column != 2 {
		t.Fatalf("capped ErrorList: %+v", capped)
	}
}

// TestCoverMemoryImporter covers the in-memory import cache and name fallback.
func TestCoverMemoryImporter(t *testing.T) {
	t.Parallel()
	importer := newMemoryImporter()
	first, err := importer.Import("fmt")
	if err != nil || first == nil || first.Name() != "fmt" {
		t.Fatalf("Import(fmt): %v %v", first, err)
	}
	second, err := importer.Import("fmt")
	if err != nil || first != second {
		t.Fatal("the importer must cache packages")
	}
	nested, err := importer.Import("example.com/a/b")
	if err != nil || nested.Name() != "b" {
		t.Fatalf("nested import name: %v %v", nested, err)
	}
	empty, err := importer.Import("")
	if err != nil || empty.Name() != "unknown" {
		t.Fatalf("empty import path: %v %v", empty, err)
	}
}

// TestCoverSemanticHelpers covers the checker's guard and annotation branches.
func TestCoverSemanticHelpers(t *testing.T) {
	t.Parallel()
	annotateTypes(nil, token.NewFileSet(), nil, "p")
	view := emptyView()
	annotateTypes(view, token.NewFileSet(), nil, "")
	if hasSideEffect(nil) {
		t.Fatal("a nil body cannot have side effects")
	}
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{
		&ast.Ident{Name: "missing"}: {},
	}}
	annotateTypeFacts(view, info)
	view.file.Decls = []ast.Decl{&ast.FuncDecl{Body: &ast.BlockStmt{}}}
	annotatePurity(view)
	if direction := channelDirection(ast.SEND); direction != "send" {
		t.Fatalf("send direction: %q", direction)
	}
	if direction := channelDirection(ast.RECV); direction != "receive" {
		t.Fatalf("receive direction: %q", direction)
	}
	if direction := channelDirection(0); direction != "both" {
		t.Fatalf("bidirectional channel: %q", direction)
	}
	if literalKind(token.ILLEGAL) != air.KindStrLit {
		t.Fatal("literalKind must fall back to the string family")
	}
	if identKind("iota") != air.KindIdent || identKind("true") != air.KindBoolLit {
		t.Fatal("identifier classification drifted")
	}
}

// TestCoverRegenerateNil checks the empty-input paths of the token helpers.
func TestCoverRegenerateNil(t *testing.T) {
	t.Parallel()
	if got := Regenerate(nil); len(got) != 0 {
		t.Fatalf("Regenerate(nil) = %d tokens", len(got))
	}
	if err := VerifyTokenFidelity(nil, air.NewUnit(air.NewPackageRef("p", nil))); err != nil {
		t.Fatalf("an empty source has no tokens to compare: %v", err)
	}
	if tokens, err := ScanTokens(nil); err != nil || len(tokens) != 0 {
		t.Fatalf("ScanTokens(nil) = %d, %v", len(tokens), err)
	}
}

// TestCoverFrontendDefaults covers the zero-value diagnostic cap fallback.
func TestCoverFrontendDefaults(t *testing.T) {
	t.Parallel()
	var frontend GoFrontend
	unit, diagnostics, err := frontend.Parse(air.SourceUnit{
		Language: "go",
		Files: []air.SourceFile{
			air.NewSourceFile("x.go", []byte("package p\n\nvar A = 1\n")),
		},
	})
	if err != nil || unit == nil || len(diagnostics) != 0 || unit.NodeCount() == 0 {
		t.Fatalf("zero-value frontend: %+v %v %v", unit, diagnostics, err)
	}
}
