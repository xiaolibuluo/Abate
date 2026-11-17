/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "abate/lang/rule"

// goEG1Rules declares the Effective Go naming rules. Every entry carries a
// structural pattern over AIR identifier or package nodes and an official
// reference to the Effective Go "Names" section.
func goEG1Rules() []rule.Spec {
	return []rule.Spec{
		{
			ID: "GO-EG-001", Reference: "Effective Go, Names: MixedCaps",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident/^[A-Za-z][A-Za-z0-9]*_[A-Za-z0-9_]+$/`,
			// The Go toolchain mandates the _test suffix on external test
			// package names, so MixedCaps never applies to that identifier.
			Where:   []string{`text-not-regex:_test$`},
			Message: "use MixedCaps rather than an underscore in an identifier",
			Mutation: rule.MutationDecl{
				Family: "M-NAM", Kind: "rename", Params: []string{"MixedCaps"},
			},
		},
		{
			ID: "GO-EG-002", Reference: "Effective Go, Names: MixedCaps (avoid ALL_CAPS)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident/^[A-Z0-9_]{3,}$/`,
			Message: "avoid ALL_CAPS identifiers; use MixedCaps",
			Mutation: rule.MutationDecl{
				Family: "M-NAM", Kind: "rename", Params: []string{"MixedCaps"},
			},
		},
		{
			ID: "GO-EG-003", Reference: "Effective Go, Names: Initialisms (ID)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident+Id`,
			Message: "write ID with both letters capitalized",
		},
		{
			ID: "GO-EG-004", Reference: "Effective Go, Names: Initialisms (URL)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Url`,
			Message: "write URL in all capitals",
		},
		{
			ID: "GO-EG-005", Reference: "Effective Go, Names: Initialisms (HTTP)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Http`,
			Message: "write HTTP in all capitals",
		},
		{
			ID: "GO-EG-006", Reference: "Effective Go, Names: Initialisms (JSON)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Json`,
			Message: "write JSON in all capitals",
		},
		{
			ID: "GO-EG-007", Reference: "Effective Go, Names: Initialisms (XML)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Xml`,
			Message: "write XML in all capitals",
		},
		{
			ID: "GO-EG-008", Reference: "Effective Go, Names: Initialisms (API)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Api`,
			Message: "write API in all capitals",
		},
		{
			ID: "GO-EG-009", Reference: "Effective Go, Names: Initialisms (SQL)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Sql`,
			Message: "write SQL in all capitals",
		},
		{
			ID: "GO-EG-010", Reference: "Effective Go, Names: Initialisms (UUID)",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Ident~Uuid`,
			Message: "write UUID in all capitals",
		},
		{
			ID: "GO-EG-011", Reference: "Effective Go, Names: Package names",
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Pattern: `Package~_`,
			// The Go toolchain mandates the _test suffix on external test
			// package names, so it is not an underscore naming defect.
			Where:   []string{`text-not-regex:_test$`},
			Message: "package names should be short lowercase words without underscores",
		},
	}
}
