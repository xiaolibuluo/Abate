/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goCRC3Rules declares the remaining Go Code Review Comments rules on import
// paths, whitespace, error punctuation and non-deterministic library calls.
func goCRC3Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-CRC-023", Reference: "Go Code Review Comments, Imports",
			Severity: rule.SeverityInfo, Category: rule.CategoryBoundary,
			Pattern: `Import/^..*[A-Z]/`,
			Message: "import paths should be lowercase",
		},
		{
			ID: "GO-CRC-024", Reference: "Go Code Review Comments, Gofmt",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `StrLit`, Where: []string{`regex:\s"$`},
			Message: "avoid trailing whitespace inside string literals",
		},
		{
			ID: "GO-CRC-025", Reference: "Go Code Review Comments, Error Strings",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(Call (Selector+New) $.. StrLit/^..*\..$/)`,
			Message: "error strings should not end with punctuation",
		},
		{
			ID: "GO-CRC-026", Reference: "Go Code Review Comments, Handle Errors",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+MustCompile) $..)`,
			// The rule targets library code; a _test.go file is test code,
			// where a fixed pattern parsed once is the standard idiom.
			Where:   []string{`path-not-regex:_test\.go$`},
			Message: "prefer the error-returning form over a Must helper in libraries",
		},
		{
			ID: "GO-CRC-027", Reference: "Go Code Review Comments, Synchronous Functions",
			Severity: rule.SeverityWarning, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector+Now) $..)`,
			Message: "read the clock through an injected source, not directly",
		},
		{
			ID: "GO-CRC-028", Reference: "Go Code Review Comments, Contexts",
			Severity: rule.SeverityInfo, Category: rule.CategoryBoundary,
			Pattern: `(Call (Selector+Getenv) $..)`,
			Message: "process environment reads should be injected as configuration",
		},
		{
			ID: "GO-CRC-029", Reference: "Go Code Review Comments, Synchronous Functions",
			Severity: rule.SeverityError, Category: rule.CategoryBoundary,
			Pattern: `(Call (Selector+Exit) $..)`,
			Message: "library code must not terminate the process; return an error",
		},
		{
			ID: "GO-CRC-030", Reference: "Go Code Review Comments, Useful Test Failures",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `(Call (Selector+Fatal) $..)`,
			Message: "test failures should name the actual and expected values",
		},
		{
			ID: "GO-CRC-031", Reference: "Go Code Review Comments, Contexts",
			Severity: rule.SeverityWarning, Category: rule.CategoryBoundary,
			Pattern: `(Call (Selector+Setenv) $..)`,
			Message: "mutating the process environment leaks global state",
		},
		{
			ID: "GO-CRC-032", Reference: "Go Code Review Comments, Handle Errors",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+ReadFile) $..)`,
			Message: "check both the bytes and the error from a file read",
		},
		{
			ID: "GO-CRC-033", Reference: "Go Code Review Comments, Contexts",
			Severity: rule.SeverityInfo, Category: rule.CategoryBoundary,
			Pattern: `(Call (Selector+Post) $.. StrLit)`,
			Message: "network calls should carry an explicit context and timeout",
		},
	}
}
