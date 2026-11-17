/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"errors"
	"slices"
	"strings"

	"abate/ces/schema"
)

// ErrUnnormalized is returned when a document has not passed normalization.
var ErrUnnormalized = errors.New("validate: experiment is not normalized")

// builtinRegistry assembles every listed rule into a fresh registry, so
// validation carries no global mutable state and is safe for concurrent use.
func builtinRegistry() (*Registry, error) {
	reg := NewRegistry()
	groups := [][]Rule{
		steadyRules(), methodRules(), probeRules(),
		rollbackRules(), scheduleRules(), targetRules(),
	}
	for _, group := range groups {
		for _, rule := range group {
			if err := reg.Register(rule); err != nil {
				return nil, err
			}
		}
	}
	return reg, nil
}

// isNormalized reports the observable invariants produced by CES-PARSE
// Normalize: a namespace, a rollback strategy, an exclusive schedule and a
// derived non-zero seed.
func isNormalized(exp *schema.Experiment) bool {
	if exp.Metadata.Namespace == "" || exp.Spec.Rollback.Strategy == "" ||
		exp.Spec.Seed == 0 {
		return false
	}
	return scheduleModeCount(exp.Spec.Schedule) == 1
}

// Validate checks a normalized experiment against every CES-VALID rule and
// returns all violations, aggregated and deterministically ordered. A nil or
// unnormalized document yields ErrUnnormalized instead of a partial report.
func Validate(exp *schema.Experiment) ([]Violation, error) {
	if exp == nil || !isNormalized(exp) {
		return nil, ErrUnnormalized
	}
	reg, err := builtinRegistry()
	if err != nil {
		return nil, err
	}
	out := reg.run(exp)
	sortViolations(out)
	return out, nil
}

// sortViolations imposes a total, deterministic order on the aggregate list.
func sortViolations(v []Violation) {
	slices.SortFunc(v, func(a, b Violation) int {
		if c := strings.Compare(a.RuleID, b.RuleID); c != 0 {
			return c
		}
		if c := strings.Compare(a.FieldPath, b.FieldPath); c != 0 {
			return c
		}
		if c := strings.Compare(a.Message, b.Message); c != 0 {
			return c
		}
		return strings.Compare(string(a.Severity), string(b.Severity))
	})
}

// RegisteredIDs returns the lexicographically ordered identifiers of all
// builtin rules, for rule-id uniqueness checks and TCK cross-referencing.
func RegisteredIDs() ([]string, error) {
	reg, err := builtinRegistry()
	if err != nil {
		return nil, err
	}
	return reg.IDs(), nil
}
