/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pybuiltin

import "abate/lang/rule"

// pyRules declares the twelve Python-specific naming, style and idiom rules
// layered on top of the sixty shared AIR structural templates.
//
// Every pattern matches a general AIR node (Ident, NumLit, Binary, Comment,
// ClassDecl, Star, NilLit, Named) by its attribute text, so the rules work with
// the real pyfront frontend as well as synthetic AIR fixtures.
func pyRules() []rule.Spec {
	return []rule.Spec{
		{
			ID:        "PY-PEP8-001",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Naming Conventions",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryNaming,
			Pattern:   `Ident/[A-Z]/`,
			Message:   "identifier uses mixed case; Python style prefers snake_case (lower_case_with_underscores)",
			Mutation:  rule.MutationDecl{Family: "M-NAM", Kind: "rename", Params: []string{"snake_case"}},
		},
		{
			ID:        "PY-PEP8-002",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Naming Conventions / Class Names",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryNaming,
			Pattern:   `ClassDecl~_`,
			Message:   "class name should use CapWords (CamelCase); avoid snake_case",
			Mutation:  rule.MutationDecl{Family: "M-NAM", Kind: "rename", Params: []string{"CapWords"}},
		},
		{
			ID:        "PY-PEP8-003",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Naming Conventions",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident/^[lIO]$/`,
			Message:   "avoid the ambiguous single-letter identifiers l, O and I",
		},
		{
			ID:        "PY-PEP8-004",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Programming Recommendations",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `NumLit/^[0-9]{4,}$/`,
			Message:   "avoid magic numbers; bind large literals to a named constant",
			Mutation:  rule.MutationDecl{Family: "M-LIT", Kind: "replace", Params: []string{"constant"}},
		},
		{
			ID:        "PY-PEP8-005",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Comments",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Directive~noqa`,
			Message:   "avoid blanket '# noqa' suppression comments; fix the underlying warning",
		},
		{
			ID:        "PY-PEP8-006",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Comments",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `Directive~ignore`,
			Message:   "avoid blanket '# type: ignore' comments; fix the type error",
		},
		{
			ID:        "PY-PEP8-007",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Imports",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Star=*`,
			Message:   "avoid wildcard imports 'from module import *'",
			Mutation:  rule.MutationDecl{Family: "M-IMP", Kind: "replace", Params: []string{"explicit"}},
		},
		{
			ID:        "PY-PEP8-008",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Programming Recommendations",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=eval`,
			Message:   "avoid the eval built-in; it is a code-injection and performance risk",
		},
		{
			ID:        "PY-PEP8-009",
			Reference: "docs/reference/python-pep8.md — PEP 8 / Programming Recommendations",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=exec`,
			Message:   "avoid the exec built-in; it is a code-injection and performance risk",
		},
		{
			ID:        "PY-PEP484-001",
			Reference: "docs/reference/python-pep8.md — PEP 484 / Comparisons to None",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `(Binary/^==$/ $.. NilLit=None)`,
			Message:   "compare with None using 'is' rather than '=='",
			Mutation:  rule.MutationDecl{Family: "M-OP", Kind: "swap", Params: []string{"is"}},
		},
		{
			ID:        "PY-PEP484-002",
			Reference: "docs/reference/python-pep8.md — PEP 484 / Comparisons to None",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `(Binary/^!=$/ $.. NilLit=None)`,
			Message:   "compare with None using 'is not' rather than '!='",
			Mutation:  rule.MutationDecl{Family: "M-OP", Kind: "swap", Params: []string{"is not"}},
		},
		{
			ID:        "PY-PEP484-003",
			Reference: "docs/reference/python-pep8.md — PEP 484 / Annotating Any",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Named=Any`,
			Message:   "avoid the typing.Any escape hatch; prefer a precise type annotation",
		},
	}
}
