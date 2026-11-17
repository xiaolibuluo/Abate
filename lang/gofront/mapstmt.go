/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"go/ast"

	"abate/lang/air"
)

// mapStmt maps a Go statement onto its AIR counterpart.
func mapStmt(v *GoView, stmt ast.Stmt) air.Node {
	if isNilASTNode(stmt) {
		return nil
	}
	n := mapStmtSwitch(v, stmt)
	v.bind(stmt, n)
	return n
}

// mapStmtSwitch maps simple statement kinds and delegates other categories.
func mapStmtSwitch(v *GoView, stmt ast.Stmt) air.Node {
	switch s := stmt.(type) {
	case *ast.AssignStmt:
		children := make([]air.Node, 0, len(s.Lhs)+len(s.Rhs))
		for _, lhs := range s.Lhs {
			children = appendNode(children, mapExpr(v, lhs))
		}
		for _, rhs := range s.Rhs {
			children = appendNode(children, mapExpr(v, rhs))
		}
		return v.boundOperator(air.KindAssign, s.Pos(), s.End(),
			s.Tok.String(), children...)
	case *ast.ExprStmt:
		return v.newNode(air.KindExprStmt, s.Pos(), s.End(), "",
			mapExpr(v, s.X))
	case *ast.DeclStmt:
		return v.newNode(air.KindDeclStmt, s.Pos(), s.End(), "",
			mapDecl(v, s.Decl))
	case *ast.DeferStmt:
		return v.newNode(air.KindDefer, s.Pos(), s.End(), "",
			mapExpr(v, s.Call))
	case *ast.GoStmt:
		return v.newNode(air.KindGoStmt, s.Pos(), s.End(), "",
			mapExpr(v, s.Call))
	case *ast.SendStmt:
		n := v.boundOperator(air.KindChanOp, s.Pos(), s.End(), "<-",
			mapExpr(v, s.Chan), mapExpr(v, s.Value))
		air.Annotate(n, air.SemChanDir, "send")
		return n
	case *ast.ReturnStmt:
		children := make([]air.Node, 0, len(s.Results))
		for _, item := range s.Results {
			children = appendNode(children, mapExpr(v, item))
		}
		return v.newNode(air.KindReturn, s.Pos(), s.End(), "", children...)
	case *ast.IncDecStmt:
		return v.boundOperator(air.KindIncDec, s.Pos(), s.End(),
			s.Tok.String(), mapExpr(v, s.X))
	case *ast.EmptyStmt:
		n := v.newNode(air.KindRegion, s.Pos(), s.End(), ";")
		air.Annotate(n, air.SemNote, "empty")
		return n
	case *ast.BadStmt:
		return damaged(v, s.Pos(), s.End(), "bad-stmt")
	default:
		return mapControlStmt(v, s)
	}
}

// Maps control-flow statement kinds and delegates clause statements.
func mapControlStmt(v *GoView, stmt ast.Stmt) air.Node {
	switch s := stmt.(type) {
	case *ast.BlockStmt:
		children := make([]air.Node, 0, len(s.List))
		for _, item := range s.List {
			children = appendNode(children, mapStmt(v, item))
		}
		return v.newNode(air.KindBlock, s.Pos(), s.End(), "", children...)
	case *ast.IfStmt:
		return v.newNode(air.KindIf, s.Pos(), s.End(), "",
			mapStmt(v, s.Init), mapExpr(v, s.Cond),
			mapStmt(v, s.Body), mapStmt(v, s.Else))
	case *ast.ForStmt:
		return v.newNode(air.KindFor, s.Pos(), s.End(), "",
			mapStmt(v, s.Init), mapExpr(v, s.Cond),
			mapStmt(v, s.Post), mapStmt(v, s.Body))
	case *ast.RangeStmt:
		return v.boundOperator(air.KindRangeFor, s.Pos(), s.End(),
			s.Tok.String(), mapExpr(v, s.Key), mapExpr(v, s.Value),
			mapExpr(v, s.X), mapStmt(v, s.Body))
	case *ast.SwitchStmt:
		return v.newNode(air.KindSwitch, s.Pos(), s.End(), "",
			mapStmt(v, s.Init), mapExpr(v, s.Tag), mapStmt(v, s.Body))
	case *ast.TypeSwitchStmt:
		n := v.newNode(air.KindSwitch, s.Pos(), s.End(), "",
			mapStmt(v, s.Init), mapStmt(v, s.Assign), mapStmt(v, s.Body))
		air.Annotate(n, air.SemNote, "type-switch")
		return n
	case *ast.SelectStmt:
		return v.newNode(air.KindSelect, s.Pos(), s.End(), "",
			mapStmt(v, s.Body))
	default:
		return mapClauseStmt(v, s)
	}
}

// Maps clause and label statement kinds and reports unknown statements.
func mapClauseStmt(v *GoView, stmt ast.Stmt) air.Node {
	switch s := stmt.(type) {
	case *ast.CaseClause:
		children := make([]air.Node, 0, len(s.List)+len(s.Body))
		for _, item := range s.List {
			children = appendNode(children, mapExpr(v, item))
		}
		for _, item := range s.Body {
			children = appendNode(children, mapStmt(v, item))
		}
		return v.newNode(air.KindCase, s.Pos(), s.End(), "", children...)
	case *ast.CommClause:
		children := make([]air.Node, 0, len(s.Body)+1)
		children = appendNode(children, mapStmt(v, s.Comm))
		for _, item := range s.Body {
			children = appendNode(children, mapStmt(v, item))
		}
		return v.newNode(air.KindCase, s.Pos(), s.End(), "", children...)
	case *ast.BranchStmt:
		children := make([]air.Node, 0, 1)
		if s.Label != nil {
			label := v.newNode(air.KindIdent, s.Label.Pos(),
				s.Label.End(), s.Label.Name)
			v.bind(s.Label, label)
			children = append(children, label)
		}
		n := v.newNode(air.KindBranch, s.Pos(), s.End(),
			s.Tok.String(), children...)
		air.Annotate(n, air.SemLabel, s.Tok.String())
		return n
	case *ast.LabeledStmt:
		if s.Label == nil {
			return damaged(v, s.Pos(), s.End(), "unknown-stmt")
		}
		label := v.newNode(air.KindIdent, s.Label.Pos(), s.Label.End(), s.Label.Name)
		v.bind(s.Label, label)
		return v.newNode(air.KindLabel, s.Pos(), s.End(), s.Label.Name,
			label, mapStmt(v, s.Stmt))
	default:
		return damaged(v, s.Pos(), s.End(), "unknown-stmt")
	}
}
