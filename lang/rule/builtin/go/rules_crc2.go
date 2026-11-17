/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goCRC2Rules declares the error-handling, import, flow and receiver rules of
// the Go Code Review Comments. Each pattern targets a distinct AIR construct.
func goCRC2Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-CRC-012", Reference: "Go Code Review Comments, Error Strings",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(Call (Selector+Errorf) $.. StrLit/^.[A-Z]/)`,
			Message: "fmt.Errorf messages should not be capitalized",
		},
		{
			ID: "GO-CRC-013", Reference: "Go Code Review Comments, Import Blank",
			Severity: rule.SeverityInfo, Category: rule.CategoryBoundary,
			Pattern: `(Import $.. Ident=_)`,
			Message: "blank imports belong in main or test packages; document the side effect",
		},
		{
			ID: "GO-CRC-014", Reference: "Go Code Review Comments, Gofmt",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(Call Ident=print $..)`,
			Message: "use fmt.Print instead of the builtin print",
		},
		{
			ID: "GO-CRC-015", Reference: "Go Code Review Comments, Handle Errors",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Assign Ident=_ $value)`,
			Message: "discarding a result with _ hides an error; handle it explicitly",
			Mutation: rule.MutationDecl{
				Family: "M-ERR", Kind: "insert-check", Params: []string{"error"},
			},
		},
		{
			ID: "GO-CRC-016", Reference: "Go Code Review Comments, Import Dot",
			Severity: rule.SeverityWarning, Category: rule.CategoryBoundary,
			Pattern: `(Import $.. Ident=.)`,
			Message: "dot imports make provenance unclear; use a qualified import",
		},
		{
			ID: "GO-CRC-017", Reference: "Go Code Review Comments, Naked Returns",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(FuncDecl $.. (Return !ANY))`,
			Message: "naked returns inside a function body harm readability",
		},
		{
			ID: "GO-CRC-018", Reference: "Go Code Review Comments, Receiver Names",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `(MethodDecl $.. Ident=self)`,
			Message: "receiver names should be short type abbreviations, not self",
		},
		{
			ID: "GO-CRC-019", Reference: "Go Code Review Comments, Comment Sentences",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `Comment`, Where: []string{`regex:^// [a-z]`},
			Message: "comments should be complete sentences beginning with a capital",
		},
		{
			ID: "GO-CRC-020", Reference: "Go Code Review Comments, Pass Values",
			Severity: rule.SeverityInfo, Category: rule.CategoryPerformance,
			Pattern: `(Call (Selector+DeepEqual) $..)`,
			Message: "prefer a typed comparison over reflect.DeepEqual",
		},
		{
			ID: "GO-CRC-021", Reference: "Go Code Review Comments, Synchronous Functions",
			Severity: rule.SeverityWarning, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector+Sleep) $..)`,
			Message: "avoid sleeping; use a deterministic injected clock or channel",
		},
		{
			ID: "GO-CRC-022", Reference: "Go Code Review Comments, Receiver Type",
			Severity: rule.SeverityInfo, Category: rule.CategoryBoundary,
			Pattern: `(MethodDecl $.. (Pointer $..))`,
			Message: "keep pointer and value receivers consistent across a type",
		},
	}
}
