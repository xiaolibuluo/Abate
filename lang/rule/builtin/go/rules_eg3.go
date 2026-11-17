/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goEG3Rules declares the Effective Go data, initialization, method and
// concurrency rules. Several entries use a where clause to refine a broad
// node kind with a value or annotation constraint.
func goEG3Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-EG-023", Reference: "Effective Go, Data: Append",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(ExprStmt (Call Ident=append $..))`,
			Message: "append returns a possibly new slice; the result must be used",
			Mutation: rule.MutationDecl{
				Family: "M-CALL", Kind: "drop-result", Params: []string{"append"},
			},
		},
		{
			ID: "GO-EG-024", Reference: "Effective Go, Initialization: Constants",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(VarDecl $.. NumLit/^[0-9]{4,}$/)`,
			Message: "give large numeric literals a named constant",
		},
		{
			ID: "GO-EG-025", Reference: "Effective Go, Errors: error strings",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(Call (Selector+New) $.. StrLit/^.[A-Z]/)`,
			Message: "error strings should not be capitalized",
		},
		{
			ID: "GO-EG-026", Reference: "Effective Go, Concurrency: Goroutines",
			Severity: rule.SeverityError, Category: rule.CategoryConcurrency,
			Pattern: `(RangeFor $.. Defer)`,
			Message: "defer inside a range loop accumulates calls until return",
		},
		{
			ID: "GO-EG-027", Reference: "Effective Go, Errors: Recover",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call Ident=recover $..)`,
			Message: "recover only works when called directly inside a deferred function",
		},
		{
			ID: "GO-EG-028", Reference: "Effective Go, Methods: Pointers vs. Values",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `(MethodDecl $.. Ident=this)`,
			Message: "receiver names should be short type abbreviations, not this",
		},
		{
			ID: "GO-EG-029", Reference: "Effective Go, Data: Slices",
			Severity: rule.SeverityWarning, Category: rule.CategoryPerformance,
			Pattern: `(For $.. Binary=+)`,
			Message: "growing a string with + inside a loop is quadratic; use strings.Builder",
		},
		{
			ID: "GO-EG-030", Reference: "Effective Go, Commentary: comment sentences",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `Doc`, Where: []string{`regex:^// [a-z]`},
			Message: "doc comments should begin with a capitalized sentence",
		},
		{
			ID: "GO-EG-031", Reference: "Effective Go, Data: Allocation with make",
			Severity: rule.SeverityInfo, Category: rule.CategoryPerformance,
			Pattern: `(Call Ident=make $.. NumLit=0)`,
			Message: "pass a capacity hint to make when the final size is known",
		},
		{
			ID: "GO-EG-032", Reference: "Effective Go, Data: Composite literals",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `(StructType !ANY)`,
			Message: "an empty struct type should carry a documented purpose",
		},
		{
			ID: "GO-EG-033", Reference: "Effective Go, Concurrency: Goroutines",
			Severity: rule.SeverityError, Category: rule.CategoryConcurrency,
			Pattern: `(GoStmt $.. (Call Ident=panic $arg))`,
			Message: "a panic inside a goroutine cannot be recovered by its parent",
		},
	}
}
