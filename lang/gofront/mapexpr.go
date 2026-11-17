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
	"abate/lang/frontkit"
)

// mapExpr maps any Go expression onto its AIR counterpart and records the
// binding so the go/types annotation pass can enrich the very same node.
func mapExpr(v *GoView, expr ast.Expr) air.Node {
	if isNilASTNode(expr) {
		return nil
	}
	n := mapExprSwitch(v, expr)
	v.bind(expr, n)
	return n
}

// mapExprSwitch contains the exhaustive expression mapping table.
func mapExprSwitch(v *GoView, expr ast.Expr) air.Node {
	switch x := expr.(type) {
	case *ast.Ident:
		return v.newNode(identKind(x.Name), x.Pos(), x.End(), x.Name)
	case *ast.BasicLit:
		return v.newNode(literalKind(x.Kind), x.Pos(), x.End(), x.Value)
	case *ast.BinaryExpr:
		return v.boundOperator(air.KindBinary, x.Pos(), x.End(), x.Op.String(),
			mapExpr(v, x.X), mapExpr(v, x.Y))
	case *ast.UnaryExpr:
		return mapUnary(v, x)
	case *ast.CallExpr:
		children := make([]air.Node, 0, len(x.Args)+1)
		children = appendNode(children, mapExpr(v, x.Fun))
		for _, arg := range x.Args {
			children = appendNode(children, mapExpr(v, arg))
		}
		n := v.newNode(air.KindCall, x.Pos(), x.End(), "", children...)
		if x.Ellipsis.IsValid() {
			air.Annotate(n, air.SemNote, "variadic-call")
		}
		return n
	case *ast.SelectorExpr:
		sel := v.newNode(air.KindIdent, x.Sel.Pos(), x.Sel.End(), x.Sel.Name)
		v.bind(x.Sel, sel)
		n := v.newNode(air.KindSelector, x.Pos(), x.End(), x.Sel.Name,
			mapExpr(v, x.X), sel)
		return n
	case *ast.IndexExpr:
		return v.newNode(air.KindIndex, x.Pos(), x.End(), "",
			mapExpr(v, x.X), mapExpr(v, x.Index))
	case *ast.IndexListExpr:
		children := make([]air.Node, 0, len(x.Indices)+1)
		children = appendNode(children, mapExpr(v, x.X))
		for _, index := range x.Indices {
			children = appendNode(children, mapExpr(v, index))
		}
		n := v.newNode(air.KindGenericInst, x.Pos(), x.End(), "", children...)
		air.Annotate(n, air.SemNote, "instantiation")
		return n
	case *ast.SliceExpr:
		n := v.newNode(air.KindSlice, x.Pos(), x.End(), "",
			mapExpr(v, x.X), mapExpr(v, x.Low), mapExpr(v, x.High),
			mapExpr(v, x.Max))
		return n
	case *ast.StarExpr:
		return v.boundOperator(air.KindStar, x.Pos(), x.End(), "*",
			mapExpr(v, x.X))
	case *ast.ParenExpr:
		n := v.newNode(air.KindRegion, x.Pos(), x.End(), "()",
			mapExpr(v, x.X))
		air.Annotate(n, air.SemNote, "parenthesized")
		return n
	case *ast.CompositeLit:
		children := make([]air.Node, 0, len(x.Elts)+1)
		children = appendNode(children, mapType(v, x.Type))
		for _, elt := range x.Elts {
			children = appendNode(children, mapExpr(v, elt))
		}
		return v.newNode(air.KindCompositeLit, x.Pos(), x.End(), "", children...)
	case *ast.KeyValueExpr:
		n := v.boundOperator(air.KindAssign, x.Pos(), x.End(), ":",
			mapExpr(v, x.Key), mapExpr(v, x.Value))
		air.Annotate(n, air.SemNote, "key-value")
		return n
	case *ast.FuncLit:
		return v.newNode(air.KindFuncLit, x.Pos(), x.End(), "func",
			mapType(v, x.Type), mapStmt(v, x.Body))
	case *ast.TypeAssertExpr:
		children := make([]air.Node, 0, 2)
		children = appendNode(children, mapExpr(v, x.X))
		if x.Type != nil {
			children = appendNode(children, mapType(v, x.Type))
		}
		return v.newNode(air.KindTypeAssert, x.Pos(), x.End(), "", children...)
	case *ast.Ellipsis:
		return v.newNode(air.KindSpread, x.Pos(), x.End(), "...",
			mapType(v, x.Elt))
	case *ast.BadExpr:
		return damaged(v, x.Pos(), x.End(), "bad-expr")
	case *ast.ArrayType, *ast.StructType, *ast.FuncType,
		*ast.InterfaceType, *ast.MapType, *ast.ChanType:
		return mapType(v, expr)
	default:
		return damaged(v, expr.Pos(), expr.End(), "unknown-expr")
	}
}

// mapUnary maps unary operators, splitting channel receives into ChanOp.
func mapUnary(v *GoView, expr *ast.UnaryExpr) air.Node {
	child := mapExpr(v, expr.X)
	if expr.Op == token.ARROW {
		n := v.boundOperator(air.KindChanOp, expr.Pos(), expr.End(), "<-", child)
		air.Annotate(n, air.SemChanDir, "receive")
		return n
	}
	kind := air.KindUnary
	if expr.Op == token.AND {
		kind = air.KindAddr
	}
	return v.boundOperator(kind, expr.Pos(), expr.End(), expr.Op.String(), child)
}

// boundOperator routes one Go operator node through the shared frontkit entry.
func (v *GoView) boundOperator(kind air.NodeKind, pos, end token.Pos,
	text string, children ...air.Node) air.Node {
	start, endPos := v.spanAt(pos, end)
	tok := frontkit.Token{
		Kind:      frontkit.TokenOp,
		Text:      text,
		Offset:    start.Offset,
		EndOffset: endPos.Offset,
		Line:      start.Line,
		Column:    start.Column,
	}
	return frontkit.BindOperator(kind, start, endPos, tok, children...)
}

// identKind maps the three predeclared value identifiers to literal kinds.
func identKind(name string) air.NodeKind {
	switch name {
	case "true", "false":
		return air.KindBoolLit
	case "nil":
		return air.KindNilLit
	default:
		return air.KindIdent
	}
}

// literalKind maps a Go token to the neutral AIR literal family.
func literalKind(kind token.Token) air.NodeKind {
	switch kind {
	case token.INT, token.FLOAT, token.IMAG:
		return air.KindNumLit
	case token.CHAR:
		return air.KindCharLit
	case token.STRING:
		return air.KindStrLit
	default:
		return air.KindStrLit
	}
}
