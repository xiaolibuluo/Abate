/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import "errors"

// Sentinel errors of the derivation pipeline. Every rejection wraps exactly
// one of these values so callers can classify a failure with errors.Is.
var (
	// ErrEmptyRuleSet reports a derivation requested without any rule.
	ErrEmptyRuleSet = errors.New("derive: empty rule set")
	// ErrEmptySuite reports a derivation experiment without test cases.
	ErrEmptySuite = errors.New("derive: empty test suite")
	// ErrNilUnit reports a nil AIR unit handed to the adapter.
	ErrNilUnit = errors.New("derive: nil air unit")
	// ErrInvalidSpec reports a malformed derived mutation specification.
	ErrInvalidSpec = errors.New("derive: invalid mutation spec")
	// ErrUnknownTransform reports a rule class outside the five-row table.
	ErrUnknownTransform = errors.New("derive: unknown transform")
	// ErrUnknownOperator reports an operator the mutation registry rejects.
	ErrUnknownOperator = errors.New("derive: operator is not registered")
	// ErrNoDerivedPoints reports a run whose report has no derived mutants.
	ErrNoDerivedPoints = errors.New("derive: no derived mutation points")
)
