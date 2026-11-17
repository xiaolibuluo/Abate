/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import "errors"

// Sentinel errors of the rule package. Every failure wraps exactly one of
// these values so callers can classify a load or match failure with errors.Is.
var (
	// ErrInvalidSpec reports a Specification that fails schema validation.
	ErrInvalidSpec = errors.New("rule: invalid spec")
	// ErrEmptyReference reports a rule without an official reference.
	ErrEmptyReference = errors.New("rule: empty reference")
	// ErrDuplicateID reports two rules sharing one namespaced identifier.
	ErrDuplicateID = errors.New("rule: duplicate rule id")
	// ErrBadPattern reports a pattern that cannot be compiled.
	ErrBadPattern = errors.New("rule: pattern does not compile")
	// ErrInvalidRuleID reports a malformed namespaced rule identifier.
	ErrInvalidRuleID = errors.New("rule: invalid rule id")
	// ErrInvalidPredicate reports a where clause that cannot be parsed.
	ErrInvalidPredicate = errors.New("rule: invalid predicate")
	// ErrUnknownLanguage reports a language filter with no registered rules.
	ErrUnknownLanguage = errors.New("rule: unknown language")
	// ErrNilUnit reports a nil AIR unit handed to a rule.
	ErrNilUnit = errors.New("rule: nil air unit")
	// ErrEmptyRuleSet reports a registry loaded without any rule.
	ErrEmptyRuleSet = errors.New("rule: empty rule set")
)
