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
	"strings"

	"abate/lang/air"
)

// mapFile maps a whole parsed file onto a package root node.
func mapFile(v *GoView) air.Node {
	children := make([]air.Node, 0, 8)
	name := ""
	if v.file.Name != nil {
		name = v.file.Name.Name
		ident := v.newNode(air.KindIdent, v.file.Name.Pos(),
			v.file.Name.End(), name)
		v.bind(v.file.Name, ident)
		children = append(children, ident)
	}
	children = append(children, mapComments(v)...)
	for _, decl := range v.file.Decls {
		children = appendNode(children, mapDecl(v, decl))
	}
	orderChildren(children)

	end := token.Pos(0)
	if tokenFile := v.fset.File(v.file.Pos()); tokenFile != nil {
		end = token.Pos(tokenFile.Base() + tokenFile.Size())
	} else if v.file.Name != nil {
		end = v.file.End()
	}
	root := v.newNode(air.KindPackage, v.file.Pos(), end, name, children...)
	v.bind(v.file, root)
	return root
}

// mapDecl dispatches a top-level declaration to its AIR family.
func mapDecl(v *GoView, decl ast.Decl) air.Node {
	if isNilASTNode(decl) {
		return damaged(v, token.NoPos, token.NoPos, "unknown-decl")
	}
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return mapFuncDecl(v, d)
	case *ast.GenDecl:
		return mapGenDecl(v, d)
	case *ast.BadDecl:
		return damaged(v, d.Pos(), d.End(), "bad-decl")
	default:
		return v.synthetic(air.KindRegion, "unknown-decl")
	}
}

// damaged builds an explicit degraded node for syntax the parser could not
// recover, so a broken file still yields a complete AIR tree.
func damaged(v *GoView, pos, end token.Pos, text string) air.Node {
	n := v.newNode(air.KindRegion, pos, end, text)
	air.Annotate(n, air.SemPartial, "true")
	return n
}

// mapGenDecl maps import, const, var, and type declaration groups.
func mapGenDecl(v *GoView, decl *ast.GenDecl) air.Node {
	kind := air.KindDeclStmt
	switch decl.Tok {
	case token.IMPORT:
		kind = air.KindImport
	case token.CONST:
		kind = air.KindConstDecl
	case token.VAR:
		kind = air.KindVarDecl
	case token.TYPE:
		kind = air.KindTypeDecl
	}
	children := make([]air.Node, 0, len(decl.Specs))
	for index, spec := range decl.Specs {
		children = appendNode(children, mapSpec(v, spec, decl.Tok, index))
	}
	n := v.newNode(kind, decl.Pos(), decl.End(), decl.Tok.String(), children...)
	v.bind(decl, n)
	return n
}

// mapFuncDecl maps a function or method declaration.
func mapFuncDecl(v *GoView, decl *ast.FuncDecl) air.Node {
	kind := air.KindFuncDecl
	if decl.Recv != nil {
		kind = air.KindMethodDecl
	}
	children := make([]air.Node, 0, 4)
	children = append(children, mapFieldList(v, decl.Recv)...)
	if decl.Name != nil {
		children = appendNode(children, mapExpr(v, decl.Name))
	}
	if decl.Type != nil {
		children = appendNode(children, mapType(v, decl.Type))
		if decl.Type.TypeParams != nil {
			air.Annotate(children[len(children)-1], air.SemTypeParam, "generic")
		}
	}
	if !isNilASTNode(decl.Body) {
		children = appendNode(children, mapStmt(v, decl.Body))
	}
	name := ""
	if decl.Name != nil {
		name = decl.Name.Name
	}
	// A FuncType starts at the func keyword, before the declaration name, so
	// children must be offset-ordered to keep intervals non-overlapping.
	orderChildren(children)
	n := v.newNode(kind, decl.Pos(), decl.End(), name, children...)
	v.bind(decl, n)
	return n
}

// mapComments maps every comment, doc group, and //go: directive by position.
func mapComments(v *GoView) []air.Node {
	out := make([]air.Node, 0, len(v.file.Comments))
	docGroups := make(map[*ast.CommentGroup]bool)
	for _, decl := range v.file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Doc != nil {
				docGroups[d.Doc] = true
			}
		case *ast.FuncDecl:
			if d.Doc != nil {
				docGroups[d.Doc] = true
			}
		}
	}
	for _, group := range v.file.Comments {
		if isNilASTNode(group) {
			continue
		}
		for _, comment := range group.List {
			if isNilASTNode(comment) {
				continue
			}
			kind := air.KindComment
			if strings.HasPrefix(comment.Text, "//go:") {
				kind = air.KindDirective
			} else if docGroups[group] {
				kind = air.KindDoc
			}
			n := v.newNode(kind, comment.Pos(), comment.End(), comment.Text)
			if kind == air.KindDirective {
				air.Annotate(n, air.SemDirective, comment.Text)
			}
			v.bind(comment, n)
			out = append(out, n)
		}
	}
	return out
}
