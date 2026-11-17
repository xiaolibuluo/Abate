/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"fmt"
	"slices"
	"strings"

	"abate/kernel/algebra"
	"abate/lang/rule"
)

// Deriver maps compiled rules to derived mutation operator specifications.
type Deriver struct{}

// NewDeriver returns a stateless deriver.
func NewDeriver() Deriver {
	return Deriver{}
}

// Derive produces exactly one specification per rule, ordered by source rule
// identifier. An unknown rule class is rejected with ErrUnknownTransform.
func (Deriver) Derive(rules []rule.Rule) ([]MutationSpec, error) {
	if len(rules) == 0 {
		return nil, ErrEmptyRuleSet
	}
	specs := make([]MutationSpec, 0, len(rules))
	for _, current := range rules {
		if current == nil {
			return nil, fmt.Errorf("%w: nil rule", ErrInvalidSpec)
		}
		decl := current.Spec()
		transform, ok := transformFor(decl)
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrUnknownTransform, decl.ID)
		}
		row, ok := transformEntryFor(transform)
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrUnknownTransform, transform)
		}
		name := "derived-" + strings.ToLower(string(decl.ID))
		specs = append(specs, MutationSpec{
			SourceRuleID: decl.ID,
			OperatorID:   algebra.FormatOperatorID(algebra.FamilyMutation, name, "1"),
			TargetKind:   row.targetKind.String(),
			Transform:    transform,
		})
	}
	slices.SortFunc(specs, func(a, b MutationSpec) int {
		return strings.Compare(string(a.SourceRuleID), string(b.SourceRuleID))
	})
	return specs, nil
}
