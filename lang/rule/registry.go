/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"slices"
	"strings"
)

// Registry holds compiled rules keyed by namespaced identifier.
//
// Loading is validation: schema conformance, a non-empty official reference,
// identifier uniqueness and pattern compilability are all checked before a
// rule becomes visible. A bad rule is rejected at load time and never reaches
// the checking pipeline.
type Registry struct {
	rules map[RuleID]Rule
}

// NewRegistry returns an empty registry ready to load rules.
func NewRegistry() *Registry {
	return &Registry{rules: make(map[RuleID]Rule, 16)}
}

// Register validates and adds one rule.
func (r *Registry) Register(spec Spec) error {
	if err := r.check(spec, nil); err != nil {
		return err
	}
	compiled, err := Compile(spec)
	if err != nil {
		return err
	}
	if r.rules == nil {
		r.rules = make(map[RuleID]Rule, 16)
	}
	r.rules[spec.ID] = compiled
	return nil
}

// Load validates and adds a batch atomically: either every rule is registered
// or the registry is left untouched.
func (r *Registry) Load(specs []Spec) error {
	staged := make(map[RuleID]Rule, len(specs))
	for _, spec := range specs {
		if err := r.check(spec, staged); err != nil {
			return err
		}
		compiled, err := Compile(spec)
		if err != nil {
			return err
		}
		staged[spec.ID] = compiled
	}
	if r.rules == nil {
		r.rules = make(map[RuleID]Rule, len(staged))
	}
	ids := make([]RuleID, 0, len(staged))
	for id := range staged {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	for _, id := range ids {
		r.rules[id] = staged[id]
	}
	return nil
}

// For returns the rules of a language in ascending identifier order. An empty
// language selects every registered rule.
func (r *Registry) For(language string) []Rule {
	want := strings.ToLower(strings.TrimSpace(language))
	out := make([]Rule, 0, len(r.rules))
	for _, id := range r.IDs() {
		current := r.rules[id]
		if want != "" && current.Spec().Language() != want {
			continue
		}
		out = append(out, current)
	}
	return out
}

// Lookup returns the rule registered under an identifier.
func (r *Registry) Lookup(id RuleID) (Rule, bool) {
	current, ok := r.rules[id]
	return current, ok
}

// IDs returns every registered identifier in ascending order.
func (r *Registry) IDs() []RuleID {
	ids := make([]RuleID, 0, len(r.rules))
	for id := range r.rules {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	return ids
}

// Len returns the number of registered rules.
func (r *Registry) Len() int {
	return len(r.rules)
}

// check applies the four load-time validations against both the registry and
// an optional staging set used by atomic batch loading.
func (r *Registry) check(spec Spec, staged map[RuleID]Rule) error {
	if strings.TrimSpace(spec.Reference) == "" {
		return fmt.Errorf("%w: %s", ErrEmptyReference, spec.ID)
	}
	if err := spec.Validate(); err != nil {
		return err
	}
	if _, exists := r.rules[spec.ID]; exists {
		return fmt.Errorf("%w: %s", ErrDuplicateID, spec.ID)
	}
	if _, exists := staged[spec.ID]; exists {
		return fmt.Errorf("%w: %s", ErrDuplicateID, spec.ID)
	}
	if _, err := CompilePattern(spec.Pattern); err != nil {
		return fmt.Errorf("%w: %s: %w", ErrBadPattern, spec.ID, err)
	}
	return nil
}
