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
	"abate/lang/rule"
)

// MutationSpec is one derived M-family mutation operator.
type MutationSpec struct {
	// SourceRuleID is the rule the operator was derived from.
	SourceRuleID rule.RuleID
	// OperatorID is the globally unique m-derived-* operator identifier.
	OperatorID algebra.OperatorID
	// TargetKind is the AIR node kind wire name the operator acts on.
	TargetKind string
	// Transform is the derivation strategy applied to the rule.
	Transform Transform
}

// Canonical renders the specification in a byte-stable form used by the
// determinism checks.
func (s MutationSpec) Canonical() string {
	return strings.Join([]string{
		string(s.SourceRuleID), string(s.OperatorID), s.TargetKind,
		string(s.Transform),
	}, "\x00")
}

// Operator builds the algebra operator described by this specification.
func (s MutationSpec) Operator() (algebra.Operator, error) {
	row, ok := transformEntryFor(s.Transform)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnknownTransform, s.Transform)
	}
	params := algebra.Params{
		"class":        row.class,
		"from":         row.from,
		"to":           row.to,
		"derived-from": string(s.SourceRuleID),
	}
	targets := []algebra.TargetKey{algebra.TargetKey(row.from)}
	return algebra.NewBasic(s.OperatorID, algebra.FamilyMutation,
		params, targets, row.invertible), nil
}
