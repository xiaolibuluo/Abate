/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"errors"
	"fmt"
	"slices"

	"abate/kernel/algebra"
)

// ErrInvalidDerived reports a derived operator that fails registry admission.
var ErrInvalidDerived = errors.New("mutation: invalid derived operator")

// NewRegistryWith builds the seven-class registry plus a frozen set of derived
// operators. The base registry is only touched during construction: extra
// operators are validated, sorted by identifier and indexed by their "from"
// token, so the returned registry stays immutable and safe for concurrent reads.
func NewRegistryWith(extra []algebra.Operator) (*Registry, error) {
	base := NewRegistry()
	if len(extra) == 0 {
		return base, nil
	}
	derived := make([]algebra.Operator, 0, len(extra))
	byToken := make(map[string][]algebra.Operator, len(extra))
	seen := make(map[algebra.OperatorID]bool, len(extra))
	for _, op := range extra {
		if err := validateDerived(base, op); err != nil {
			return nil, err
		}
		id := op.ID()
		if seen[id] {
			return nil, fmt.Errorf("%w: duplicate identifier %s", ErrInvalidDerived, id)
		}
		seen[id] = true
		derived = append(derived, op)
		token := op.Params()["from"]
		byToken[token] = append(byToken[token], op)
	}
	slices.SortFunc(derived, byOperatorID)
	for _, token := range sortedTokens(byToken) {
		slices.SortFunc(byToken[token], byOperatorID)
	}
	base.derived = derived
	base.byToken = byToken
	return base, nil
}

// NewEngineWithRegistry binds one seed to an explicitly supplied registry.
func NewEngineWithRegistry(seed uint64, reg *Registry) (*Engine, error) {
	if reg == nil {
		return nil, ErrNilRegistry
	}
	return &Engine{seed: seed, registry: reg}, nil
}

// validateDerived admits only well-formed M-family derived operators.
func validateDerived(base *Registry, op algebra.Operator) error {
	if op == nil {
		return fmt.Errorf("%w: nil operator", ErrInvalidDerived)
	}
	if op.ID() == "" {
		return fmt.Errorf("%w: empty identifier", ErrInvalidDerived)
	}
	if op.Family() != algebra.FamilyMutation {
		return fmt.Errorf("%w: family %s on %s", ErrInvalidDerived,
			op.Family(), op.ID())
	}
	params := op.Params()
	if !slices.Contains(base.names, params["class"]) {
		return fmt.Errorf("%w: class %q on %s", ErrInvalidDerived,
			params["class"], op.ID())
	}
	for _, key := range []string{"from", "to", "derived-from"} {
		if params[key] == "" {
			return fmt.Errorf("%w: missing %q on %s", ErrInvalidDerived,
				key, op.ID())
		}
	}
	return nil
}

// sortedTokens returns the derived-token index keys in ascending order so that
// construction never depends on map traversal order.
func sortedTokens(index map[string][]algebra.Operator) []string {
	keys := make([]string, 0, len(index))
	for key := range index {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
