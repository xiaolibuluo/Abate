/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"fmt"
	"strings"

	"abate/kernel/algebra"
	"abate/kernel/mutation"
)

// Operators validates the specifications and builds their algebra operators.
// A specification is available only when the mutation registry indexes it by
// its "from" token and accepts it as a registered M-family operator.
func Operators(specs []MutationSpec) ([]algebra.Operator, error) {
	if len(specs) == 0 {
		return nil, ErrEmptyRuleSet
	}
	seen := make(map[algebra.OperatorID]bool, len(specs))
	ops := make([]algebra.Operator, 0, len(specs))
	for _, spec := range specs {
		if !spec.Transform.Valid() {
			return nil, fmt.Errorf("%w: %s", ErrUnknownTransform, spec.Transform)
		}
		if !strings.HasPrefix(string(spec.OperatorID), derivedPrefix) {
			return nil, fmt.Errorf("%w: %s", ErrUnknownOperator, spec.OperatorID)
		}
		if seen[spec.OperatorID] {
			return nil, fmt.Errorf("%w: duplicate %s", ErrInvalidSpec, spec.OperatorID)
		}
		seen[spec.OperatorID] = true
		op, err := spec.Operator()
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	if err := checkAvailability(ops); err != nil {
		return nil, err
	}
	return ops, nil
}

// Validate reports whether every specification is available in a mutation
// experiment without returning the built operators.
func Validate(specs []MutationSpec) error {
	_, err := Operators(specs)
	return err
}

// checkAvailability admits operators only when the registry can discover and
// validate each of them through the public Judge path.
func checkAvailability(ops []algebra.Operator) error {
	reg, err := mutation.NewRegistryWith(ops)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUnknownOperator, err)
	}
	for _, op := range ops {
		params := op.Params()
		found := false
		for _, candidate := range reg.For(params["from"]) {
			if candidate.ID() == op.ID() && reg.Valid(candidate) {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("%w: undiscoverable %s", ErrUnknownOperator, op.ID())
		}
	}
	return nil
}
