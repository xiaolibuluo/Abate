/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
	"strings"

	"abate/lang/air"
)

// memoryImporter resolves every import purely in memory.
//
// It never touches the real filesystem, never calls go/importer.Default, and
// returns an incomplete but complete-marked package for dependencies that were
// not supplied in the SourceUnit. Type errors on those packages are collected
// by the checker's error callback and simply degrade to missing annotations.
type memoryImporter struct {
	packages map[string]*types.Package
}

// newMemoryImporter returns an empty in-memory importer.
func newMemoryImporter() *memoryImporter {
	return &memoryImporter{packages: make(map[string]*types.Package)}
}

// Import returns the package for path, creating an empty one on first use.
func (m *memoryImporter) Import(path string) (*types.Package, error) {
	if pkg, ok := m.packages[path]; ok {
		return pkg, nil
	}
	name := path
	if slash := strings.LastIndex(path, "/"); slash >= 0 {
		name = path[slash+1:]
	}
	if name == "" {
		name = "unknown"
	}
	pkg := types.NewPackage(path, name)
	pkg.MarkComplete()
	m.packages[path] = pkg
	return pkg, nil
}

// annotateTypes runs a best-effort in-memory type check and enriches AIR.
//
// All findings are attached to the exact nodes produced by the mapping layer.
// The checker's error callback is intentionally quiet: a missing dependency
// must not turn a syntactically valid file into a parse failure. Files with a
// nil pointer are filtered before go/types walks the syntax tree.
func annotateTypes(v *GoView, fset *token.FileSet, files []*ast.File,
	packageName string) {
	if v == nil || v.file == nil || packageName == "" {
		return
	}
	safeFiles := make([]*ast.File, 0, len(files))
	for _, file := range files {
		if file != nil {
			safeFiles = append(safeFiles, file)
		}
	}
	info := &types.Info{
		Types:  make(map[ast.Expr]types.TypeAndValue),
		Defs:   make(map[*ast.Ident]types.Object),
		Uses:   make(map[*ast.Ident]types.Object),
		Scopes: make(map[ast.Node]*types.Scope),
	}
	conf := types.Config{
		Importer:                 newMemoryImporter(),
		Error:                    func(error) {},
		DisableUnusedImportCheck: true,
	}
	_, _ = conf.Check(packageName, fset, safeFiles, info)
	v.info = info
	annotateTypeFacts(v, info)
	annotatePurity(v)
}

// annotateTypeFacts attaches resolved types and scope chains to nodes.
func annotateTypeFacts(v *GoView, info *types.Info) {
	for expr, fact := range info.Types {
		target, ok := v.lookup(expr)
		if !ok || fact.Type == nil {
			continue
		}
		air.Annotate(target, air.SemTypeRef, types.TypeString(fact.Type, nil))
	}
	for ident, object := range info.Defs {
		target, ok := v.lookup(ident)
		if !ok || object == nil {
			continue
		}
		air.Annotate(target, air.SemScopeID, scopeChain(object.Parent()))
		if _, isVar := object.(*types.Var); isVar {
			air.Annotate(target, air.SemMutability, "mutable")
		}
	}
	for ident, object := range info.Uses {
		target, ok := v.lookup(ident)
		if !ok || object == nil {
			continue
		}
		air.Annotate(target, air.SemScopeID, scopeChain(object.Parent()))
	}
}

// scopeChain renders a scope chain as a stable slash-separated position list.
func scopeChain(scope *types.Scope) string {
	parts := make([]string, 0, 4)
	for current := scope; current != nil; current = current.Parent() {
		parts = append(parts, strconv.Itoa(int(current.Pos())))
	}
	return strings.Join(parts, "/")
}

// annotatePurity applies a syntax-based purity heuristic to declarations.
func annotatePurity(v *GoView) {
	for _, decl := range v.file.Decls {
		if isNilASTNode(decl) {
			continue
		}
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		target, ok := v.lookup(fn)
		if !ok {
			continue
		}
		if isNilASTNode(fn.Body) {
			air.Annotate(target, air.SemPurity, "pure")
			continue
		}
		if hasSideEffect(fn.Body) {
			air.Annotate(target, air.SemPurity, "side-effect")
		} else {
			air.Annotate(target, air.SemPurity, "pure")
		}
	}
}

// hasSideEffect reports whether a function body contains effectful syntax.
//
// This is a documented, conservative heuristic: any assignment, increment,
// call, goroutine, defer, or channel send is treated as a side effect. It is
// syntax-only by design so it stays identical for every parse of the source.
func hasSideEffect(body ast.Node) bool {
	if isNilASTNode(body) {
		return false
	}
	found := false
	ast.Inspect(body, func(node ast.Node) bool {
		if found {
			return false
		}
		switch node.(type) {
		case *ast.AssignStmt, *ast.IncDecStmt, *ast.GoStmt,
			*ast.DeferStmt, *ast.SendStmt, *ast.CallExpr:
			found = true
			return false
		default:
			return true
		}
	})
	return found
}
