/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"fmt"
	"strings"

	"abate/lang/rule"
)

// Template is one language-neutral AIR structural rule declaration.
//
// Every field is pure data whose type mirrors rule.Spec. Positive and Negative
// are compact AIR fixture fragments ("Kind:text" tokens joined by a pipe);
// Reference optionally overrides the citation produced by ReferenceFor.
type Template struct {
	// Family names the defect-class theme this template belongs to.
	Family string
	// Pattern is the AIR structural pattern source string.
	Pattern string
	// Where holds the predicate clauses filtering matched captures.
	Where []string
	// Message is the human-readable violation template.
	Message string
	// Mutation optionally declares one derived mutation operator.
	Mutation rule.MutationDecl
	// Severity ranks the finding produced by the rule.
	Severity rule.Severity
	// Category groups the rule by defect class.
	Category rule.Category
	// Reference optionally overrides the generated official citation.
	Reference string
	// Positive is the AIR fixture fragment the rule must detect.
	Positive string
	// Negative is the AIR fixture fragment the rule must not detect.
	Negative string
}

// Key renders the declaration into a byte-stable deduplication key.
//
// The key covers every declaration field, so two templates belonging to
// different families or carrying different fixtures are never merged.
func (t Template) Key() string {
	parts := []string{
		t.Family, t.Pattern, strings.Join(t.Where, "\x1f"),
		t.Message, t.Severity.String(), t.Category.String(),
		t.Mutation.Canonical(), t.Reference, t.Positive, t.Negative,
	}
	return strings.Join(parts, "\x00")
}

// Fixtures reports whether the template carries both fixture sides.
func (t Template) Fixtures() bool {
	return t.Positive != "" && t.Negative != ""
}

// toSpec assembles one rule.Spec from the template and its generated identity.
func (t Template) toSpec(namespace string, ordinal int, citation string) rule.Spec {
	reference := citation
	if t.Reference != "" {
		reference = t.Reference
	}
	return rule.Spec{
		ID:        rule.RuleID(fmt.Sprintf("%s-%s-%03d", namespace, sourceAIR, ordinal)),
		Reference: reference,
		Severity:  t.Severity,
		Category:  t.Category,
		Pattern:   t.Pattern,
		Where:     append([]string(nil), t.Where...),
		Message:   t.Message,
		Mutation:  t.Mutation,
	}
}

// newTemplate merges one theme family with one of its five variants.
func newTemplate(f family, v variant) Template {
	return Template{
		Family:   f.name,
		Pattern:  v.pattern,
		Where:    append([]string(nil), v.where...),
		Message:  f.message,
		Mutation: f.mutation,
		Severity: f.severity,
		Category: f.category,
		Positive: f.positive,
		Negative: f.negative,
	}
}
