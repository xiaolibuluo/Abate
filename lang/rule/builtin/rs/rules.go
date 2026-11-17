/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsbuiltin

import "abate/lang/rule"

// rsRules declares the twelve Rust official-specification rules layered on top
// of the sixty shared AIR templates. Every entry is a direct rule.Spec with an
// explicit official reference, a structural pattern over generic AIR nodes, and
// a derived M-family mutation operator.
func rsRules() []rule.Spec {
	return []rule.Spec{
		{
			ID:        "RS-API-001",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Error Handling",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=unwrap`,
			Message:   "avoid unwrap(); propagate the error with ? or expect",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "use-question", Params: []string{"?"},
			},
		},
		{
			ID:        "RS-API-002",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Error Handling",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=expect`,
			Message:   "expect() still panics; prefer explicit error handling",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "use-question", Params: []string{"?"},
			},
		},
		{
			ID:        "RS-API-003",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Panics",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=panic`,
			Message:   "avoid panic!() in library code; return a Result instead",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "return-result", Params: []string{"Result"},
			},
		},
		{
			ID:        "RS-API-004",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / unfinished-macros",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=todo`,
			Message:   "todo!() panics at runtime; implement the branch or document it",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "implement", Params: []string{"todo"},
			},
		},
		{
			ID:        "RS-API-005",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / unfinished-macros",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=unimplemented`,
			Message:   "unimplemented!() panics at runtime; provide a real body",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "implement", Params: []string{"unimplemented"},
			},
		},
		{
			ID:        "RS-API-006",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Clone",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryPerformance,
			Pattern:   `Ident=clone`,
			Message:   "review .clone(); pass by reference or copy where possible",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "use-reference", Params: []string{"&T"},
			},
		},
		{
			ID:        "RS-API-007",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Unsafe",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `@Note`,
			Message:   "review unsafe blocks; justify every raw pointer and FFI use",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "justify-unsafe", Params: []string{"unsafe"},
			},
		},
		{
			ID:        "RS-API-008",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Lifetimes",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `GenericParam`,
			Message:   "review explicit lifetimes; elide them when inference works",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "elide-lifetime", Params: []string{"'a"},
			},
		},
		{
			ID:        "RS-API-009",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Mutability",
			Severity:  rule.SeverityInfo,
			Category:  rule.CategoryStyle,
			Pattern:   `@Mutability`,
			Message:   "prefer immutable bindings; only use mut where mutation occurs",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "make-immutable", Params: []string{"let"},
			},
		},
		{
			ID:        "RS-API-010",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Casting",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Cast=as`,
			Message:   "avoid as casts; use TryFrom or checked conversions",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "use-tryfrom", Params: []string{"TryFrom"},
			},
		},
		{
			ID:        "RS-API-011",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Debugging",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryStyle,
			Pattern:   `Ident=dbg`,
			Message:   "remove dbg!() debug macros before shipping",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "remove-debug", Params: []string{"dbg"},
			},
		},
		{
			ID:        "RS-API-012",
			Reference: "docs/reference/rust-api.md — Rust API Guidelines / Control Flow",
			Severity:  rule.SeverityWarning,
			Category:  rule.CategoryErrorProne,
			Pattern:   `Ident=unreachable`,
			Message:   "unreachable!() is undefined behaviour if ever hit; prove it",
			Mutation: rule.MutationDecl{
				Family: "M-RS", Kind: "prove-unreachable", Params: []string{"unreachable"},
			},
		},
	}
}
