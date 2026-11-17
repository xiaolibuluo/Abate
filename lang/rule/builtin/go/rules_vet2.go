/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goVET2Rules declares the vet appends, assign, stringintconv, waitgroup and
// lostcancel semantic subset. Each entry fixes the number or kind of call
// arguments that the corresponding analyzer inspects.
func goVET2Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-VET-012", Reference: "cmd/vet appends: append with no values",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call Ident=append $only !ANY)`,
			Message: "append with no values is a no-op; check the call",
		},
		{
			ID: "GO-VET-013", Reference: "cmd/vet assign: copy with one argument",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call Ident=copy $only !ANY)`,
			Message: "copy needs both a destination and a source",
		},
		{
			ID: "GO-VET-014", Reference: "cmd/vet stringintconv: integer to string conversion",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call Ident=string $.. NumLit)`,
			Message: "string(65) yields a rune, not the decimal text; use strconv.Itoa",
		},
		{
			ID: "GO-VET-015", Reference: "cmd/vet stringintconv: string to integer conversion",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call Ident=int $.. StrLit)`,
			Message: "converting a string to int is a type error; use strconv.Atoi",
		},
		{
			ID: "GO-VET-016", Reference: "cmd/vet composite: make without a length",
			Severity: rule.SeverityWarning, Category: rule.CategoryPerformance,
			Pattern: `(Call Ident=make $only !ANY)`,
			Message: "make without a length yields an empty slice; pass a capacity hint",
		},
		{
			ID: "GO-VET-017", Reference: "cmd/vet unusedresult: Replace with a zero count",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Replace) $.. NumLit=0)`,
			Message: "strings.Replace with n=0 performs no replacement",
		},
		{
			ID: "GO-VET-018", Reference: "cmd/vet unusedresult: SplitN with a zero count",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Split) $.. NumLit=0)`,
			Message: "a zero split count returns no substrings",
		},
		{
			ID: "GO-VET-019", Reference: "cmd/vet waitgroup: Wait before Add",
			Severity: rule.SeverityInfo, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector+Wait) $..)`,
			Message: "verify Add is called before Wait on a WaitGroup",
		},
		{
			ID: "GO-VET-020", Reference: "cmd/vet waitgroup: Add with a zero delta",
			Severity: rule.SeverityWarning, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector+Add) $.. NumLit=0)`,
			Message: "Add(0) does not register any goroutine",
		},
		{
			ID: "GO-VET-021", Reference: "cmd/vet lostcancel: zero-duration Sleep",
			Severity: rule.SeverityInfo, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector+Sleep) $.. NumLit=0)`,
			Message: "a zero-duration Sleep never yields; use a channel instead",
		},
		{
			ID: "GO-VET-022", Reference: "cmd/vet lostcancel: context Done channel",
			Severity: rule.SeverityInfo, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector+Done) $..)`,
			Message: "select on ctx.Done() and always call the cancel function",
		},
	}
}
