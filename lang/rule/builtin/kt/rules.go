/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktbuiltin

import "abate/lang/rule"

// ktRules declares the twelve Kotlin-specific naming, style and idiom rules
// layered on top of the shared AIR structural templates.
//
// Every pattern matches a general AIR node (Ident, NumLit, Unary, Optional) by
// its attribute text or kind, so the rules work with the real ktfront frontend
// as well as synthetic AIR fixtures. Kotlin constructs that the frontend emits
// only as a distinctive keyword leaf (suspend, data, sealed, companion,
// lateinit, var, inline, when) are matched on that leaf identifier text.
func ktRules() []rule.Spec {
	return []rule.Spec{
		{
			ID:        "KT-DGN-001",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Naming",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryNaming,
			Pattern:   `Ident~_`,
			Message:   "use camelCase for function and property names; avoid snake_case identifiers",
			Mutation: rule.MutationDecl{
				Family: "M-NAM", Kind: "rename", Params: []string{"camelCase"},
			},
		},
		{
			ID:        "KT-DGN-002",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Inline",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=inline`,
			Message:   "review inline functions; restrict them to hot-path lambda parameters and mind cross-module inlining limits",
		},
		{
			ID:        "KT-DGN-003",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Coroutines",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=suspend`,
			Message:   "append the Async suffix to suspending functions for discoverability",
		},
		{
			ID:        "KT-DGN-004",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / when",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=when`,
			Message:   "review when expressions for exhaustiveness and branch coverage",
		},
		{
			ID:        "KT-DGN-005",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Null Safety",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Unary=!!`,
			Message:   "avoid the !! non-null assertion; use safe calls or explicit null checks",
		},
		{
			ID:        "KT-DGN-006",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Null Safety",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Optional`,
			Message:   "review nullable types; prefer non-nullable declarations where feasible",
		},
		{
			ID:        "KT-DGN-007",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Magic Numbers",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `NumLit/^[0-9]{4,}$/`,
			Message:   "avoid magic numbers; bind large literals to a named constant",
			Mutation: rule.MutationDecl{
				Family: "M-LIT", Kind: "replace", Params: []string{"constant"},
			},
		},
		{
			ID:        "KT-DGN-008",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Immutability",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=var`,
			Message:   "prefer val (immutable) over var unless mutation is required",
		},
		{
			ID:        "KT-DGN-009",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / lateinit",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=lateinit`,
			Message:   "review lateinit properties; ensure they are initialized before first access",
		},
		{
			ID:        "KT-DGN-010",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Data Classes",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=data`,
			Message:   "keep data classes immutable; avoid adding mutable state after the primary constructor",
		},
		{
			ID:        "KT-DGN-011",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Sealed Classes",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=sealed`,
			Message:   "review sealed classes; exploit when exhaustiveness for algebraic data types",
		},
		{
			ID:        "KT-DGN-012",
			Reference: "docs/reference/kotlin-conventions.md — Kotlin Coding Conventions / Companion Objects",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=companion`,
			Message:   "review companion objects; expose public API through top-level declarations or extensions",
		},
	}
}
