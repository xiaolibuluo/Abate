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
	"strconv"

	"abate/lang/air"
)

// mapSpec dispatches one declaration spec.
func mapSpec(v *GoView, spec ast.Spec, tok token.Token, index int) air.Node {
	if isNilASTNode(spec) {
		return v.synthetic(air.KindRegion, "unknown-spec")
	}
	switch s := spec.(type) {
	case *ast.ImportSpec:
		return mapImportSpec(v, s)
	case *ast.ValueSpec:
		return mapValueSpec(v, s, tok, index)
	case *ast.TypeSpec:
		return mapTypeSpec(v, s)
	default:
		return v.synthetic(air.KindRegion, "unknown-spec")
	}
}

// mapImportSpec maps one import path and its optional alias.
func mapImportSpec(v *GoView, spec *ast.ImportSpec) air.Node {
	if spec == nil {
		return v.synthetic(air.KindImport, "")
	}
	children := make([]air.Node, 0, 2)
	if spec.Name != nil {
		children = appendNode(children, mapExpr(v, spec.Name))
	}
	text := ""
	if spec.Path != nil {
		text = spec.Path.Value
		lit := v.newNode(air.KindStrLit, spec.Path.Pos(),
			spec.Path.End(), spec.Path.Value)
		v.bind(spec.Path, lit)
		children = append(children, lit)
	}
	n := v.newNode(air.KindImport, spec.Pos(), spec.End(), text, children...)
	v.bind(spec, n)
	return n
}

// mapValueSpec maps a const or var binding, preserving iota position.
func mapValueSpec(v *GoView, spec *ast.ValueSpec, tok token.Token,
	index int) air.Node {
	if spec == nil {
		return v.synthetic(air.KindVarDecl, "")
	}
	kind := air.KindVarDecl
	if tok == token.CONST {
		kind = air.KindConstDecl
	}
	children := make([]air.Node, 0, len(spec.Names)+len(spec.Values)+1)
	for _, name := range spec.Names {
		children = appendNode(children, mapExpr(v, name))
	}
	if spec.Type != nil {
		children = appendNode(children, mapType(v, spec.Type))
	}
	for _, value := range spec.Values {
		children = appendNode(children, mapExpr(v, value))
	}
	text := ""
	if len(spec.Names) > 0 {
		text = spec.Names[0].Name
	}
	n := v.newNode(kind, spec.Pos(), spec.End(), text, children...)
	if containsIota(spec.Values) {
		air.Annotate(n, air.SemIota, strconv.Itoa(index))
	}
	v.bind(spec, n)
	return n
}

// mapTypeSpec maps a named type, its type parameters, and its underlying type.
func mapTypeSpec(v *GoView, spec *ast.TypeSpec) air.Node {
	if spec == nil {
		return v.synthetic(air.KindTypeDecl, "")
	}
	children := make([]air.Node, 0, 3)
	text := ""
	if spec.Name != nil {
		text = spec.Name.Name
		children = appendNode(children, mapExpr(v, spec.Name))
	}
	if spec.TypeParams != nil {
		children = appendNode(children, mapTypeParams(v, spec.TypeParams))
	}
	if spec.Type != nil {
		children = appendNode(children, mapType(v, spec.Type))
	}
	n := v.newNode(air.KindTypeDecl, spec.Pos(), spec.End(), text, children...)
	if spec.Assign.IsValid() {
		air.Annotate(n, air.SemNote, "type-alias")
	}
	v.bind(spec, n)
	return n
}

// mapTypeParams wraps a generic type parameter list in a GenericParam node.
func mapTypeParams(v *GoView, fields *ast.FieldList) air.Node {
	if fields == nil {
		return v.synthetic(air.KindGenericParam, "")
	}
	children := make([]air.Node, 0, len(fields.List)*2)
	for _, field := range fields.List {
		for _, name := range field.Names {
			param := mapExpr(v, name)
			air.Annotate(param, air.SemTypeParam, "type-parameter")
			children = appendNode(children, param)
		}
		if field.Type != nil {
			children = appendNode(children, mapType(v, field.Type))
		}
	}
	return v.newNode(air.KindGenericParam, fields.Pos(), fields.End(),
		"", children...)
}

// mapFieldList maps struct fields, interface methods, and parameter lists.
func mapFieldList(v *GoView, fields *ast.FieldList) []air.Node {
	if fields == nil {
		return nil
	}
	out := make([]air.Node, 0, len(fields.List))
	for _, field := range fields.List {
		children := make([]air.Node, 0, len(field.Names)+2)
		for _, name := range field.Names {
			children = appendNode(children, mapExpr(v, name))
		}
		if field.Type != nil {
			children = appendNode(children, mapType(v, field.Type))
		}
		if field.Tag != nil {
			tag := v.newNode(air.KindAttribute, field.Tag.Pos(),
				field.Tag.End(), field.Tag.Value)
			air.Annotate(tag, air.SemTag, field.Tag.Value)
			v.bind(field.Tag, tag)
			children = append(children, tag)
		}
		n := v.newNode(air.KindField, field.Pos(), field.End(), "", children...)
		v.bind(field, n)
		out = append(out, n)
	}
	return out
}

// containsIota reports whether a value list references the iota constant.
func containsIota(values []ast.Expr) bool {
	found := false
	for _, value := range values {
		if isNilASTNode(value) {
			continue
		}
		ast.Inspect(value, func(node ast.Node) bool {
			if ident, ok := node.(*ast.Ident); ok && ident.Name == "iota" {
				found = true
				return false
			}
			return !found
		})
		if found {
			break
		}
	}
	return found
}
