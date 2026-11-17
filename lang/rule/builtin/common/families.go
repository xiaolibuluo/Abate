/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import "abate/lang/rule"

// family is the shared metadata of one defect-class theme.
//
// All five variants of a family share the same severity, category, message,
// mutation declaration and AIR fixture pair, so one fixture table proves all
// sixty generated rules can hit and can stay silent.
type family struct {
	name     string
	severity rule.Severity
	category rule.Category
	message  string
	mutation rule.MutationDecl
	positive string
	negative string
}

// familyCount is the number of declared cross-language themes.
const familyCount = 12

// families rebuilds the ordered theme table on every call.
func families() []family {
	return []family{
		{
			name: "naming", severity: rule.SeverityWarning,
			category: rule.CategoryNaming,
			message:  "identifier uses an underscore instead of a canonical compound name",
			mutation: rule.MutationDecl{Family: "M-NAM", Kind: "rename", Params: []string{"canonical"}},
			positive: "Ident:bad_name", negative: "Ident:goodName",
		},
		{
			name: "empty-capture", severity: rule.SeverityWarning,
			category: rule.CategoryErrorProne,
			message:  "captured binding has an empty attribute text",
			mutation: rule.MutationDecl{Family: "M-LIT", Kind: "replace", Params: []string{"placeholder"}},
			positive: "Ident:", negative: "Ident:value",
		},
		{
			name: "magic-number", severity: rule.SeverityInfo,
			category: rule.CategoryStyle,
			message:  "unnamed numeric literal should be bound to a named constant",
			mutation: rule.MutationDecl{Family: "M-LIT", Kind: "increment", Params: []string{"one"}},
			positive: "NumLit:1000", negative: "NumLit:1",
		},
		{
			name: "error-ignored", severity: rule.SeverityError,
			category: rule.CategoryErrorProne,
			message:  "error result is discarded by a blank binding",
			mutation: rule.MutationDecl{Family: "M-STMT", Kind: "insert", Params: []string{"check"}},
			positive: "Ident:_", negative: "Ident:err",
		},
		{
			name: "comparison-operator", severity: rule.SeverityWarning,
			category: rule.CategoryErrorProne,
			message:  "non-strict equality operator should be replaced by a strict one",
			mutation: rule.MutationDecl{Family: "M-OP", Kind: "swap", Params: []string{"strict"}},
			positive: "Binary:==", negative: "Binary:===",
		},
		{
			name: "bare-dereference", severity: rule.SeverityError,
			category: rule.CategoryErrorProne,
			message:  "pointer dereference is not guarded by a null check",
			mutation: rule.MutationDecl{Family: "M-COND", Kind: "negate", Params: []string{"guard"}},
			positive: "Star:unsafe", negative: "Star:safe",
		},
		{
			name: "bool-literal-compare", severity: rule.SeverityWarning,
			category: rule.CategoryStyle,
			message:  "boolean expression is compared against a literal true",
			mutation: rule.MutationDecl{Family: "M-OP", Kind: "drop", Params: []string{"operand"}},
			positive: "Binary:flag==true", negative: "Binary:flag!=true",
		},
		{
			name: "constant-naming", severity: rule.SeverityWarning,
			category: rule.CategoryNaming,
			message:  "constant name should follow the canonical casing convention",
			mutation: rule.MutationDecl{Family: "M-NAM", Kind: "rename", Params: []string{"canonical"}},
			positive: "Ident:MAX_SIZE", negative: "Ident:MaxSize",
		},
		{
			name: "long-parameters", severity: rule.SeverityInfo,
			category: rule.CategoryBoundary,
			message:  "function signature declares too many positional parameters",
			mutation: rule.MutationDecl{Family: "M-STMT", Kind: "split", Params: []string{"record"}},
			positive: "FuncDecl:many_arguments", negative: "FuncDecl:fewArgs",
		},
		{
			name: "unused-binding", severity: rule.SeverityWarning,
			category: rule.CategoryErrorProne,
			message:  "binding is prefixed as unused but never consumed",
			mutation: rule.MutationDecl{Family: "M-NAM", Kind: "rename", Params: []string{"used"}},
			positive: "Ident:_unused", negative: "Ident:used",
		},
		{
			name: "comment-placeholder", severity: rule.SeverityInfo,
			category: rule.CategoryStyle,
			message:  "comment still carries an unresolved placeholder marker",
			mutation: rule.MutationDecl{Family: "M-LIT", Kind: "replace", Params: []string{"documented"}},
			positive: "Comment:stub", negative: "Comment:explain",
		},
		{
			name: "boundary-compare", severity: rule.SeverityWarning,
			category: rule.CategoryBoundary,
			message:  "boundary comparison against a literal zero needs a named limit",
			mutation: rule.MutationDecl{Family: "M-OP", Kind: "swap", Params: []string{"limit"}},
			positive: "Binary:index<=0", negative: "Binary:index<count",
		},
	}
}
