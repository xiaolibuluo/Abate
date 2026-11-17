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

	"abate/lang/air"
)

// mapType maps a Go type expression onto the AIR type family.
//
// Named types retain an inner identifier node so the lexical token stream is
// still covered by AIR. Pointer types are StarExpr in go/ast, so they are
// handled here rather than in the expression table.
func mapType(v *GoView, expr ast.Expr) air.Node {
	if isNilASTNode(expr) {
		return nil
	}
	n := mapTypeSwitch(v, expr)
	v.bind(expr, n)
	return n
}

// mapTypeSwitch maps named, pointer, container and leaf type expressions.
func mapTypeSwitch(v *GoView, expr ast.Expr) air.Node {
	switch x := expr.(type) {
	case *ast.Ident:
		return v.newNode(air.KindNamed, x.Pos(), x.End(), x.Name,
			mapExpr(v, x))
	case *ast.SelectorExpr:
		n := v.newNode(air.KindNamed, x.Pos(), x.End(), x.Sel.Name,
			mapExpr(v, x))
		air.Annotate(n, air.SemNote, "qualified-type")
		return n
	case *ast.StarExpr:
		return v.newNode(air.KindPointer, x.Pos(), x.End(), "*",
			mapType(v, x.X))
	case *ast.ArrayType:
		kind := air.KindArray
		if x.Len == nil {
			kind = air.KindSliceType
		}
		return v.newNode(kind, x.Pos(), x.End(), "",
			mapExpr(v, x.Len), mapType(v, x.Elt))
	case *ast.MapType:
		return v.newNode(air.KindMap, x.Pos(), x.End(), "",
			mapType(v, x.Key), mapType(v, x.Value))
	case *ast.ChanType:
		n := v.newNode(air.KindChanType, x.Pos(), x.End(), "",
			mapType(v, x.Value))
		air.Annotate(n, air.SemChanDir, channelDirection(x.Dir))
		return n
	case *ast.Ellipsis:
		return v.newNode(air.KindSpread, x.Pos(), x.End(), "...",
			mapType(v, x.Elt))
	case *ast.ParenExpr:
		return v.newNode(air.KindRegion, x.Pos(), x.End(), "()",
			mapType(v, x.X))
	case *ast.BadExpr:
		return damaged(v, x.Pos(), x.End(), "bad-type")
	default:
		return mapDerivedType(v, x)
	}
}

// Maps signature, composite, generic and operator-derived type expressions.
func mapDerivedType(v *GoView, expr ast.Expr) air.Node {
	switch x := expr.(type) {
	case *ast.FuncType:
		children := make([]air.Node, 0, 4)
		if x.TypeParams != nil {
			children = appendNode(children, mapTypeParams(v, x.TypeParams))
		}
		children = append(children, mapFieldList(v, x.Params)...)
		children = append(children, mapFieldList(v, x.Results)...)
		// The func keyword and the declaration name sit outside the signature.
		// Starting at the type parameter or parameter list keeps sibling
		// intervals disjoint while still covering the whole signature.
		start := x.Pos()
		switch {
		case x.TypeParams != nil && x.TypeParams.Opening.IsValid():
			start = x.TypeParams.Opening
		case x.Params != nil && x.Params.Opening.IsValid():
			start = x.Params.Opening
		}
		return v.newNode(air.KindFuncType, start, x.End(), "", children...)
	case *ast.StructType:
		return v.newNode(air.KindStructType, x.Pos(), x.End(), "",
			mapFieldList(v, x.Fields)...)
	case *ast.InterfaceType:
		return v.newNode(air.KindInterfaceType, x.Pos(), x.End(), "",
			mapFieldList(v, x.Methods)...)
	case *ast.IndexExpr:
		n := v.newNode(air.KindGenericInst, x.Pos(), x.End(), "",
			mapExpr(v, x))
		return n
	case *ast.IndexListExpr:
		return v.newNode(air.KindGenericInst, x.Pos(), x.End(), "",
			mapExpr(v, x))
	case *ast.BinaryExpr:
		if x.Op == token.OR {
			return v.newNode(air.KindUnion, x.Pos(), x.End(), "|",
				mapType(v, x.X), mapType(v, x.Y))
		}
		return mapExpr(v, expr)
	case *ast.UnaryExpr:
		if x.Op == token.TILDE {
			return v.newNode(air.KindRef, x.Pos(), x.End(), "~",
				mapType(v, x.X))
		}
		return mapExpr(v, expr)
	default:
		return v.synthetic(air.KindNamed, "unknown-type")
	}
}

// channelDirection renders a Go channel direction as a neutral annotation.
func channelDirection(dir ast.ChanDir) string {
	switch dir {
	case ast.SEND:
		return "send"
	case ast.RECV:
		return "receive"
	default:
		return "both"
	}
}
