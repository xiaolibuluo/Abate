/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"errors"
	"strings"

	"abate/ces/parse"
	"abate/ces/schema"
	"abate/kernel/algebra"
)

// compileMethod reuses the CES parser's expression compiler, which assembles
// the operator tree through KRN-ALG (depth guard and concurrent conflict
// detection are enforced inside that compiler).
func compileMethod(exp *schema.Experiment) (algebra.Operator, error) {
	return parse.CompileExpr(exp.Spec.Method.Expression, exp.Spec.Method.Operators)
}

// expressionHasProduct reports whether an expression contains a concurrent product.
func expressionHasProduct(expr string) bool {
	return strings.Contains(expr, "⊗")
}

// compileFailure inspects a compiler error and classifies deterministic causes.
func compileFailure(err error) (depth bool, conflict bool, message string) {
	if err == nil {
		return false, false, ""
	}
	var pe *parse.ParseError
	if !errors.As(err, &pe) {
		return false, false, err.Error()
	}
	if pe.RuleID == parse.RuleExprDepth ||
		strings.Contains(pe.Message, algebra.ErrTooDeep.Error()) {
		return true, false, pe.Message
	}
	if strings.Contains(pe.Message, algebra.ErrTargetConflict.Error()) ||
		strings.Contains(pe.Message, algebra.ErrIncompatible.Error()) {
		return false, true, pe.Message
	}
	return false, false, pe.Message
}
