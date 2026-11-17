/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goVET1Rules declares the vet printf and errorsas semantic subset. The
// patterns require an exact child sequence, so a formatting call without the
// argument that matches its directive is detected through a negative pattern.
func goVET1Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-VET-001", Reference: "cmd/vet printf: Printf directive without an argument",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Printf) StrLit~% !ANY)`,
			Message: "Printf format directive has no matching argument",
		},
		{
			ID: "GO-VET-002", Reference: "cmd/vet printf: Sprintf directive without an argument",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Sprintf) StrLit~% !ANY)`,
			Message: "Sprintf format directive has no matching argument",
		},
		{
			ID: "GO-VET-003", Reference: "cmd/vet printf: Errorf directive without an argument",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Errorf) StrLit~% !ANY)`,
			Message: "Errorf format directive has no matching argument",
		},
		{
			ID: "GO-VET-004", Reference: "cmd/vet printf: Fatalf directive without an argument",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Fatalf) StrLit~% !ANY)`,
			Message: "Fatalf format directive has no matching argument",
		},
		{
			ID: "GO-VET-005", Reference: "cmd/vet printf: Logf directive without an argument",
			Severity: rule.SeverityError, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Logf) StrLit~% !ANY)`,
			Message: "Logf format directive has no matching argument",
		},
		{
			ID: "GO-VET-006", Reference: "cmd/vet printf: Println with a format directive",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Println) $.. StrLit~%)`,
			Message: "Println does not interpret formatting directives",
		},
		{
			ID: "GO-VET-007", Reference: "cmd/vet printf: Print with a format directive",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Print) $.. StrLit~%)`,
			Message: "Print does not interpret formatting directives",
		},
		{
			ID: "GO-VET-008", Reference: "cmd/vet printf: Sprint with a format directive",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Sprint) $.. StrLit~%)`,
			Message: "Sprint does not interpret formatting directives",
		},
		{
			ID: "GO-VET-009", Reference: "cmd/vet printf: Errorf should wrap with %w",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Errorf) $.. StrLit~%v)`,
			Message: "use %w instead of %v when wrapping an error",
		},
		{
			ID: "GO-VET-010", Reference: "cmd/vet errorsas: errors.As target must be a pointer",
			Severity: rule.SeverityWarning, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+As) $..)`,
			Message: "verify that the errors.As target is a pointer to an error type",
		},
		{
			ID: "GO-VET-011", Reference: "cmd/vet errorsas: errors.Unwrap chain",
			Severity: rule.SeverityInfo, Category: rule.CategoryErrorProne,
			Pattern: `(Call (Selector+Unwrap) $..)`,
			Message: "prefer errors.Is or errors.As over a manual unwrap chain",
		},
	}
}
