/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"strings"

	"abate/ces/schema"
	"abate/kernel/algebra"
)

// CompileExpr parses a CES operator expression into an algebra operator tree.
//
// An operator:// reference must be declared in ops; a bare identifier is an
// inline definition whose family is inferred from its m-/c- prefix.
func CompileExpr(expr string, ops []schema.OperatorDef) (algebra.Operator, error) {
	tokens, err := lex(expr)
	if err != nil {
		return nil, newExprError(RuleExprSyntax, err.Error(), 0)
	}
	p := &exprParser{tokens: tokens}
	ast, err := p.parse()
	if err != nil {
		return nil, err
	}
	return compileNode(ast, ops, 0)
}

// compileNode lowers one expression node into an algebra operator.
func compileNode(n *exprNode, ops []schema.OperatorDef, depth int) (algebra.Operator, error) {
	if depth > algebra.MaxDepth {
		return nil, newExprError(RuleExprDepth, "operator tree exceeds max depth", n.pos)
	}
	switch n.kind {
	case nodeIdentity:
		return algebra.Identity(""), nil
	case nodeOperand:
		if n.ref {
			return resolveRefOperand(n.text, ops, n.pos)
		}
		return resolveOperand(n.text, ops, n.pos)
	case nodeInverse:
		inner, err := compileNode(n.left, ops, depth+1)
		if err != nil {
			return nil, err
		}
		inv, err := algebra.Inverse(inner)
		if err != nil {
			return nil, newExprError(RuleExprNotInvertible, "operator has no dual", n.pos)
		}
		return inv, nil
	case nodeCompose:
		return combine(n, ops, depth, algebra.Compose)
	case nodeProduct:
		return combine(n, ops, depth, algebra.Concurrent)
	default:
		return nil, newExprError(RuleExprSyntax, "unknown expression node", n.pos)
	}
}

// combine lowers a binary node with the given algebra join operation.
func combine(n *exprNode, ops []schema.OperatorDef, depth int,
	join func(algebra.Operator, algebra.Operator) (algebra.Operator, error),
) (algebra.Operator, error) {
	left, err := compileNode(n.left, ops, depth+1)
	if err != nil {
		return nil, err
	}
	right, err := compileNode(n.right, ops, depth+1)
	if err != nil {
		return nil, err
	}
	joined, err := join(left, right)
	if err != nil {
		return nil, newExprError(RuleExprSyntax,
			"operator composition failed: "+err.Error(), n.pos)
	}
	return joined, nil
}

// resolveOperand synthesizes an inline operator, declaring it on first use.
func resolveOperand(id string, ops []schema.OperatorDef, pos int) (algebra.Operator, error) {
	for _, def := range ops {
		if def.ID == id {
			return buildOperator(def, pos)
		}
	}
	family := algebra.FamilyChaos
	if strings.HasPrefix(id, "m-") {
		family = algebra.FamilyMutation
	}
	return algebra.NewBasic(algebra.OperatorID(id), family,
		algebra.Params{}, []algebra.TargetKey{algebra.TargetKey(id)}, true), nil
}

// resolveRefOperand resolves an operator:// reference that must be declared.
func resolveRefOperand(id string, ops []schema.OperatorDef, pos int) (algebra.Operator, error) {
	for _, def := range ops {
		if def.ID == id {
			return buildOperator(def, pos)
		}
	}
	return nil, newExprError(RuleExprUnknownOp,
		"operator:// reference is not declared: "+id, pos)
}

// buildOperator materializes a declared CES operator definition.
func buildOperator(def schema.OperatorDef, pos int) (algebra.Operator, error) {
	var family algebra.Family
	switch def.Family {
	case schema.FamilyMutation:
		family = algebra.FamilyMutation
	case schema.FamilyChaos:
		family = algebra.FamilyChaos
	default:
		return nil, newExprError(RuleExprSyntax,
			"unknown operator family: "+def.Family, pos)
	}
	targets := make([]algebra.TargetKey, 0, len(def.Targets))
	for _, t := range def.Targets {
		targets = append(targets, algebra.TargetKey(t))
	}
	return algebra.NewBasic(algebra.OperatorID(def.ID), family,
		algebra.Params(def.Params), targets, def.Invertible), nil
}
