/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsbuiltin

import "abate/lang/rule"

// jsRules declares the twelve JavaScript/TypeScript-specific naming, style and
// idiom rules layered on top of the sixty shared AIR structural templates.
//
// Every pattern matches a general AIR node (Binary, Ident, NumLit, Comment,
// VarDecl, With, Unary, TypeAssert, EnumDecl) by its attribute text, so the
// rules work with the real jsfront frontend as well as synthetic AIR fixtures.
func jsRules() []rule.Spec {
	return []rule.Spec{
		{
			ID:        "JS-ES-001",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Comparison Operators",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Binary/^==$/`,
			Message:   "use the strict equality operator '===' instead of '=='",
			Mutation:  rule.MutationDecl{Family: "M-OP", Kind: "swap", Params: []string{"==="}},
		},
		{
			ID:        "JS-ES-002",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Comparison Operators",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Binary/^!=$/`,
			Message:   "use the strict inequality operator '!==' instead of '!='",
			Mutation:  rule.MutationDecl{Family: "M-OP", Kind: "swap", Params: []string{"!=="}},
		},
		{
			ID:        "JS-ES-003",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Variables",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `VarDecl`,
			Message:   "avoid var; declare with const and fall back to let",
			Mutation:  rule.MutationDecl{Family: "M-VAR", Kind: "replace", Params: []string{"const"}},
		},
		{
			ID:        "JS-ES-004",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Blocks",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `With`,
			Message:   "avoid the with statement; it makes identifier resolution ambiguous",
		},
		{
			ID:        "JS-ES-005",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Standard Library",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=eval`,
			Message:   "avoid the eval built-in; it is a code-injection and performance risk",
		},
		{
			ID:        "JS-ES-006",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Numbers",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `NumLit/^[0-9]{4,}$/`,
			Message:   "avoid magic numbers; bind large literals to a named constant",
			Mutation:  rule.MutationDecl{Family: "M-LIT", Kind: "replace", Params: []string{"constant"}},
		},
		{
			ID:        "JS-ES-007",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Comments",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Comment~eslint`,
			Message:   "avoid blanket 'eslint-disable' suppression comments; fix the underlying warning",
		},
		{
			ID:        "JS-ES-008",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Comparison Operators",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Binary=instanceof`,
			Message:   "prefer a type predicate or property check over 'instanceof'",
		},
		{
			ID:        "JS-ES-009",
			Reference: "docs/reference/js-airbnb.md — Airbnb JavaScript Style Guide / Standard Library",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=alert`,
			Message:   "avoid alert/confirm/prompt; use a structured UI or logging layer",
		},
		{
			ID:        "JS-TS-001",
			Reference: "docs/reference/js-ts-official.md — TypeScript Handbook / Narrowing",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Unary=void`,
			Message:   "avoid the void operator; return a value or use a side-effect call directly",
		},
		{
			ID:        "JS-TS-002",
			Reference: "docs/reference/js-ts-official.md — TypeScript Handbook / Type Assertions",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `TypeAssert`,
			Message:   "minimize type assertions ('as'); prefer narrowing or a typed signature",
		},
		{
			ID:        "JS-TS-003",
			Reference: "docs/reference/js-ts-official.md — TypeScript Handbook / Enums",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `EnumDecl`,
			Message:   "prefer a union of string literals or 'as const' over a runtime enum",
		},
	}
}
