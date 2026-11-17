/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"strings"
)

// Spec is the table-driven declaration of one rule.
//
// A Specification is pure data: Pattern is a structural pattern source string,
// Where holds predicate clauses, and Mutation optionally declares an M-family
// operator derivation. Loading performs schema validation, reference checking,
// identifier uniqueness and pattern compilation before a Rule is registered.
type Spec struct {
	// ID is the namespaced identifier in LANG-SOURCE-NNN form.
	ID RuleID
	// Reference cites the official specification entry backing this rule.
	Reference string
	// Severity ranks the finding produced by this rule.
	Severity Severity
	// Category groups the rule by defect class.
	Category Category
	// Pattern is a non-empty structural pattern source string.
	Pattern string
	// Where holds predicate clauses that filter matched captures.
	Where []string
	// Message is the human-readable violation template.
	Message string
	// Mutation optionally declares a derived mutation operator.
	Mutation MutationDecl
}

// Validate checks the schema-level invariants of a Specification.
//
// Reference presence, identifier uniqueness and pattern compilability are
// checked separately by the Registry so each rejection names its own sentinel.
func (s Spec) Validate() error {
	if _, err := ParseRuleID(string(s.ID)); err != nil {
		return err
	}
	if strings.TrimSpace(s.Reference) == "" {
		return fmt.Errorf("%w: %s", ErrEmptyReference, s.ID)
	}
	if !s.Severity.Valid() {
		return fmt.Errorf("%w: bad severity %q on %s", ErrInvalidSpec, s.Severity, s.ID)
	}
	if !s.Category.Valid() {
		return fmt.Errorf("%w: bad category %q on %s", ErrInvalidSpec, s.Category, s.ID)
	}
	if strings.TrimSpace(s.Pattern) == "" {
		return fmt.Errorf("%w: empty pattern on %s", ErrBadPattern, s.ID)
	}
	if strings.TrimSpace(s.Message) == "" {
		return fmt.Errorf("%w: empty message on %s", ErrInvalidSpec, s.ID)
	}
	if !s.Mutation.Valid() {
		return fmt.Errorf("%w: bad mutation on %s", ErrInvalidSpec, s.ID)
	}
	return nil
}

// Language returns the lowercase language tag derived from the identifier.
func (s Spec) Language() string {
	return s.ID.Language()
}

// Canonical renders the immutable declaration fields in a byte-stable form.
func (s Spec) Canonical() string {
	parts := []string{
		string(s.ID), s.Reference, s.Severity.String(), s.Category.String(),
		s.Pattern, strings.Join(s.Where, "\x1f"), s.Message, s.Mutation.Canonical(),
	}
	return strings.Join(parts, "\x00")
}
