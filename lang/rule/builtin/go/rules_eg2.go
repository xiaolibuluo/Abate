/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goEG2Rules declares the Effective Go function, control-flow and error
// rules. The patterns combine node kinds with child patterns, deep matching,
// negation and captures so each entry detects a concrete construct.
func goEG2Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-EG-012", Reference: "Effective Go, Names: Getters",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `FuncDecl^Get`,
			Message: "a getter should be named for the property, not prefixed with Get",
			Mutation: rule.MutationDecl{
				Family: "M-NAM", Kind: "rename", Params: []string{"getter"},
			},
		},
		{
			ID: "GO-EG-013", Reference: "Effective Go, Semicolons: omit explicit semicolons",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `Region=;`,
			Message: "the lexer inserts semicolons; do not write an empty statement",
		},
		{
			ID: "GO-EG-014", Reference: "Effective Go, Errors: Panic",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call Ident=panic $arg)`,
			Message: "avoid panic($arg) in library code; return an error instead",
			Mutation: rule.MutationDecl{
				Family: "M-ERR", Kind: "replace", Params: []string{"error"},
			},
		},
		{
			ID: "GO-EG-015", Reference: "Effective Go, Printing: prefer the fmt package",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(Call Ident=println $..)`,
			Message: "use fmt.Println instead of the builtin println",
		},
		{
			ID: "GO-EG-016", Reference: "Effective Go, Data: Allocation with new",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `(Call Ident=new $..)`,
			Message: "prefer a composite literal or &T{} over new(T)",
		},
		{
			ID: "GO-EG-017", Reference: "Effective Go, Control structures: Defer",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(For $.. Defer)`,
			Message: "defer inside a loop accumulates calls until the function returns",
			Mutation: rule.MutationDecl{
				Family: "M-CTRL", Kind: "hoist", Params: []string{"defer"},
			},
		},
		{
			ID: "GO-EG-018", Reference: "Effective Go, Functions: Named result parameters",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(Return !ANY)`,
			Message: "naked returns are unclear in long functions; name the results",
		},
		{
			ID: "GO-EG-019", Reference: "Effective Go, Concurrency: Channels",
			Severity: rule.SeverityWarning, Category: rule.CategoryConcurrency,
			Pattern: `(Select (Block !ANY))`,
			Message: "an empty select blocks forever; use a signal channel instead",
		},
		{
			ID: "GO-EG-020", Reference: "Effective Go, Control structures: If and For",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `Branch=goto`,
			Message: "avoid goto; structured control flow is easier to verify",
		},
		{
			ID: "GO-EG-021", Reference: "Effective Go, Methods: Interfaces and other types",
			Severity: rule.SeverityInfo, Category: rule.CategoryBoundary,
			Pattern: `(InterfaceType !ANY)`,
			Message: "prefer the predeclared any over an empty interface type",
		},
		{
			ID: "GO-EG-022", Reference: "Effective Go, Initialization: The init function",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `FuncDecl=init`,
			Message: "prefer explicit initialization over package init side effects",
		},
	}
}
