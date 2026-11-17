/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cbuiltin

import "abate/lang/rule"

// cRules declares the twelve C and C++ official-specification rules layered on
// top of the sixty shared AIR templates. Every entry is a direct rule.Spec with
// an explicit official reference, a structural pattern over generic AIR nodes,
// and a derived M-family mutation operator.
func cRules() []rule.Spec {
	return []rule.Spec{
		{
			ID:        "C-CERT-001",
			Reference: "docs/reference/c-cert.md — CERT C / Preprocessor",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Directive=define`,
			Message:   "prefer a typed constant or inline function over a #define macro",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "replace-macro", Params: []string{"const"},
			},
		},
		{
			ID:        "C-CERT-002",
			Reference: "docs/reference/c-cert.md — CERT C / Numeric",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `NumLit/^[0-9]+$/`,
			Message:   "avoid magic numbers; name the constant with an enum or const",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "name-literal", Params: []string{"kConstant"},
			},
		},
		{
			ID:        "C-CERT-003",
			Reference: "docs/reference/c-cert.md — CERT C / MEM",
			Severity:  rule.SeverityError,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident~alloc`,
			Message:   "always check the return value of an allocator for NULL",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "check-null", Params: []string{"NULL"},
			},
		},
		{
			ID:        "C-CERT-004",
			Reference: "docs/reference/c-cert.md — CERT C / EXP",
			Severity:  rule.SeverityError,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Star`,
			Message:   "review raw pointer dereference for bounds and lifetime safety",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "review-deref", Params: []string{"safe"},
			},
		},
		{
			ID:        "C-CERT-005",
			Reference: "docs/reference/c-cert.md — CERT C / EXP",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `NilLit`,
			Message:   "avoid the NULL literal; use nullptr in C++ or an explicit check",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "replace-null", Params: []string{"nullptr"},
			},
		},
		{
			ID:        "C-CERT-006",
			Reference: "docs/reference/c-cert.md — CERT C / STR",
			Severity:  rule.SeverityError,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident~str`,
			Message:   "review str* family usage; prefer sized variants or safe APIs",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "use-sized", Params: []string{"snprintf"},
			},
		},
		{
			ID:        "C-CERT-007",
			Reference: "docs/reference/c-cert.md — CERT C / FIO",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=printf`,
			Message:   "audit printf for format-string vulnerabilities; prefer a safe API",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "audit-format", Params: []string{"fputs"},
			},
		},
		{
			ID:        "C-CERT-008",
			Reference: "docs/reference/c-cert.md — CERT C / EXP",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Cast=sizeof`,
			Message:   "do not apply sizeof to an expression with side effects",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "split-sizeof", Params: []string{"value"},
			},
		},
		{
			ID:        "C-CERT-009",
			Reference: "docs/reference/c-cert.md — CERT C / EXP",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `IncDec`,
			Message:   "prefer prefix increment and decrement where the value is unused",
			Mutation: rule.MutationDecl{
				Family: "M-CERT", Kind: "prefix-inc", Params: []string{"++x"},
			},
		},
		{
			ID:        "C-CPP-001",
			Reference: "docs/reference/cpp-core.md — C++ Core Guidelines / R.11",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Unary=new`,
			Message:   "avoid raw new/delete; use a standard smart pointer instead",
			Mutation: rule.MutationDecl{
				Family: "M-CPP", Kind: "use-smart-ptr", Params: []string{"unique_ptr"},
			},
		},
		{
			ID:        "C-CPP-002",
			Reference: "docs/reference/cpp-core.md — C++ Core Guidelines / E.2",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Throw`,
			Message:   "avoid throwing exceptions across language or module boundaries",
			Mutation: rule.MutationDecl{
				Family: "M-CPP", Kind: "use-error-code", Params: []string{"status"},
			},
		},
		{
			ID:        "C-CPP-003",
			Reference: "docs/reference/cpp-core.md — C++ Core Guidelines / Type",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Cast=reinterpret_cast`,
			Message:   "avoid reinterpret_cast; prefer a typed, ownership-aware design",
			Mutation: rule.MutationDecl{
				Family: "M-CPP", Kind: "avoid-cast", Params: []string{"variant"},
			},
		},
	}
}
