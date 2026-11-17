/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csbuiltin

import "abate/lang/rule"

// csRules declares the twelve C#/.NET-specific naming, style and idiom rules
// layered on top of the sixty shared AIR structural templates.
//
// Every pattern matches a general AIR node (Ident, NumLit, Optional) by its
// attribute text or kind, so the rules work with the real csfront frontend as
// well as synthetic AIR fixtures. C# constructs that the frontend distinguishes
// only through a semantic annotation are matched on their distinctive keyword
// identifier (record, async, from, get, sealed, ...).
func csRules() []rule.Spec {
	return []rule.Spec{
		{
			ID:        "CS-DGN-001",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Naming",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryNaming,
			Pattern:   `Ident~_`,
			Message:   "use PascalCase for type and member names; avoid snake_case identifiers",
			Mutation: rule.MutationDecl{
				Family: "M-NAM", Kind: "rename", Params: []string{"PascalCase"},
			},
		},
		{
			ID:        "CS-DGN-002",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Naming Constants",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryNaming,
			Pattern:   `Ident/^[A-Z0-9_]{3,}$/`,
			Message:   "use PascalCase for constants; avoid SCREAMING_CAPS identifiers",
			Mutation: rule.MutationDecl{
				Family: "M-NAM", Kind: "rename", Params: []string{"PascalCase"},
			},
		},
		{
			ID:        "CS-DGN-003",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Async",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=async`,
			Message:   "append the Async suffix to asynchronous methods for discoverability",
		},
		{
			ID:        "CS-DGN-004",
			Reference: "docs/reference/dotnet-design.md — C# Coding Conventions / LINQ",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=from`,
			Message:   "review LINQ queries for readability and deferred-execution pitfalls",
		},
		{
			ID:        "CS-DGN-005",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Properties",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Optional`,
			Message:   "review nullable references; enforce null checks or use non-nullable types",
		},
		{
			ID:        "CS-DGN-006",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Records",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=record`,
			Message:   "keep record types immutable; avoid adding mutable state after the primary constructor",
		},
		{
			ID:        "CS-DGN-007",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Exceptions",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=Exception`,
			Message:   "avoid catching the base System.Exception; catch the most specific type",
		},
		{
			ID:        "CS-DGN-008",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Namespaces",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=using`,
			Message:   "organize using directives; place imports inside the namespace and remove unused ones",
		},
		{
			ID:        "CS-DGN-009",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Magic Numbers",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `NumLit/^[0-9]{4,}$/`,
			Message:   "avoid magic numbers; bind large literals to a named constant",
			Mutation: rule.MutationDecl{
				Family: "M-LIT", Kind: "replace", Params: []string{"constant"},
			},
		},
		{
			ID:        "CS-DGN-010",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / ToString",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=ToString`,
			Message:   "override ToString deliberately; document the contract and keep it allocation-free",
		},
		{
			ID:        "CS-DGN-011",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Properties",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=get`,
			Message:   "review property accessors; prefer read-only or init-only where mutable state is unneeded",
		},
		{
			ID:        "CS-DGN-012",
			Reference: "docs/reference/dotnet-design.md — .NET Framework Design Guidelines / Inheritance",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=sealed`,
			Message:   "review sealed classes; enable extension through interfaces when sealing is not required",
		},
	}
}
