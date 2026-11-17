/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"errors"
	"fmt"
	"slices"

	"abate/ces/schema"
)

// Sentinel errors of the validate package.
var (
	// ErrDuplicateRule reports two rules registered under the same identifier.
	ErrDuplicateRule = errors.New("validate: duplicate rule id")
	// ErrNilRule reports a nil rule passed to Register.
	ErrNilRule = errors.New("validate: nil rule")
)

// Registry is a collection of uniquely identified rules.
type Registry struct {
	rules map[string]Rule
}

// NewRegistry returns an empty rule registry.
func NewRegistry() *Registry {
	return &Registry{rules: make(map[string]Rule)}
}

// Register adds a rule, rejecting a nil rule or a duplicate identifier.
func (r *Registry) Register(rule Rule) error {
	if rule == nil {
		return ErrNilRule
	}
	id := rule.ID()
	if _, exists := r.rules[id]; exists {
		return fmt.Errorf("%w: %s", ErrDuplicateRule, id)
	}
	r.rules[id] = rule
	return nil
}

// IDs returns every registered identifier in lexicographic order.
func (r *Registry) IDs() []string {
	ids := make([]string, 0, len(r.rules))
	for id := range r.rules {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	return ids
}

// run executes every rule in lexicographic identifier order and aggregates
// every violation, so a single document yields a complete, ordered report.
func (r *Registry) run(exp *schema.Experiment) []Violation {
	var out []Violation
	for _, id := range r.IDs() {
		out = append(out, r.rules[id].Check(exp)...)
	}
	return out
}
