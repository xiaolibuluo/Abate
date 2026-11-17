/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

import (
	"errors"
	"fmt"
	"slices"
)

// errDuplicateOperator reports a second registration under the same identifier.
var errDuplicateOperator = errors.New("mutate: duplicate operator")

// Registry holds the ordered set of mutation operators. Comment, whitespace
// and formatting operators are rejected by Register and by OperatorSpec.
type Registry struct {
	specs map[string]OperatorSpec
	order []string
}

// NewRegistry returns an empty registry ready for registration.
func NewRegistry() Registry {
	return Registry{specs: make(map[string]OperatorSpec)}
}

// DefaultRegistry builds the three P0 operators and the five P1 operators
// declared by ci/e2e_plan.json.
func DefaultRegistry() (Registry, error) {
	registry := NewRegistry()
	builders := []func(*Registry) error{
		RegisterComparison,
		RegisterArithmetic,
		RegisterDigit,
		RegisterCondition,
		RegisterLiteral,
		RegisterReturn,
		RegisterBlock,
		RegisterCall,
	}
	for _, build := range builders {
		if err := build(&registry); err != nil {
			return Registry{}, err
		}
	}
	return registry, nil
}

// Register validates and stores one operator. Duplicate identifiers and
// comment, whitespace or formatting families are refused.
func (r *Registry) Register(spec OperatorSpec) error {
	if err := spec.Validate(); err != nil {
		return err
	}
	if _, found := r.specs[spec.ID]; found {
		return fmt.Errorf("%w: %s", errDuplicateOperator, spec.ID)
	}
	if r.specs == nil {
		r.specs = make(map[string]OperatorSpec)
	}
	r.specs[spec.ID] = spec
	r.order = append(r.order, spec.ID)
	slices.Sort(r.order)
	return nil
}

// All returns every operator in ascending identifier order.
func (r Registry) All() []OperatorSpec {
	out := make([]OperatorSpec, 0, len(r.order))
	for _, id := range r.order {
		out = append(out, r.specs[id])
	}
	return out
}

// Lookup returns one operator by identifier.
func (r Registry) Lookup(id string) (OperatorSpec, bool) {
	spec, found := r.specs[id]
	return spec, found
}

// IDs returns the registered identifiers in ascending order.
func (r Registry) IDs() []string {
	return slices.Clone(r.order)
}
