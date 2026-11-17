/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goVET4Rules closes the vet semantic subset with the appends, nilfunc, bools,
// copylocks, unreachable and httpresponse analyzers. Comparisons and empty
// blocks are expressed directly over AIR binary and block nodes.
func goVET4Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-VET-034", Reference: "cmd/vet appends: cap of a composite literal",
			Severity: rule.SeverityInfo, Category: rule.CategoryPerformance,
			Pattern: `(Call Ident=cap $.. CompositeLit)`,
			Message: "cap of a freshly built literal is a constant; use a named length",
		},
		{
			ID: "GO-VET-035", Reference: "cmd/vet appends: len of a composite literal",
			Severity: rule.SeverityInfo, Category: rule.CategoryPerformance,
			Pattern: `(Call Ident=len $.. CompositeLit)`,
			Message: "len of a freshly built literal is a constant; use a named length",
		},
		{
			ID: "GO-VET-036", Reference: "cmd/vet nilfunc: equality comparison with nil",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Binary=== $.. NilLit)`,
			Message: "verify that a nil comparison is on a nilable type",
		},
		{
			ID: "GO-VET-037", Reference: "cmd/vet nilfunc: inequality comparison with nil",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Binary=!= $.. NilLit)`,
			Message: "verify that a nil comparison is on a nilable type",
		},
		{
			ID: "GO-VET-038", Reference: "cmd/vet bools: comparison with a boolean literal",
			Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
			Pattern: `(Binary=== $.. BoolLit)`,
			Message: "compare booleans directly instead of against a literal",
		},
		{
			ID: "GO-VET-039", Reference: "cmd/vet copylocks: reflect copy review",
			Severity: rule.SeverityInfo, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector+Copy) $..)`,
			Message: "verify that a copied value does not contain a mutex",
		},
		{
			ID: "GO-VET-040", Reference: "cmd/vet copylocks: mutex Lock review",
			Severity: rule.SeverityInfo, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector=Lock) $..)`,
			Message: "pair every Lock with an unlocked path",
		},
		{
			ID: "GO-VET-041", Reference: "cmd/vet copylocks: mutex Unlock review",
			Severity: rule.SeverityInfo, Category: rule.CategoryConcurrency,
			Pattern: `(Call (Selector=Unlock) $..)`,
			Message: "verify Unlock is reached on every branch",
		},
		{
			ID: "GO-VET-042", Reference: "cmd/vet httpresponse: close review",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Close) $..)`,
			Message: "check the error returned by Close and close exactly once",
		},
		{
			ID: "GO-VET-043", Reference: "cmd/vet unreachable: empty block",
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Pattern: `(Block !ANY)`,
			Message: "an empty block usually hides a missing statement",
		},
		{
			ID: "GO-VET-044", Reference: "cmd/vet httpresponse: deferred body close",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			// The analyzer targets the immediate defer r.Close() statement.
			// Requiring the call as the direct defer operand excludes a Close
			// buried in a deferred closure, such as a WASM teardown helper.
			Pattern: `(Defer (Call (Selector+Close)))`,
			Message: "defer the response body close immediately after a successful request",
		},
	}
}
