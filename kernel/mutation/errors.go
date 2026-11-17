/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import "errors"

// Sentinel errors of the mutation kernel.
var (
	// ErrNilTarget reports a nil Target passed to the engine.
	ErrNilTarget = errors.New("mutation: target must not be nil")

	// ErrNilSuite reports a nil Suite passed to the engine.
	ErrNilSuite = errors.New("mutation: suite must not be nil")

	// ErrNilRegistry reports a nil operator registry.
	ErrNilRegistry = errors.New("mutation: registry must not be nil")

	// ErrNoNodes reports a target exposing no mutable nodes.
	ErrNoNodes = errors.New("mutation: target exposes no mutable nodes")

	// ErrNoTests reports a suite exposing no test cases.
	ErrNoTests = errors.New("mutation: suite exposes no test cases")

	// ErrNoMutants reports that no operator applies to any node.
	ErrNoMutants = errors.New("mutation: no applicable mutation points")

	// ErrUnknownOperator reports an operator outside the registered families.
	ErrUnknownOperator = errors.New("mutation: operator is not registered")

	// ErrCanceled reports a canceled or missing run context.
	ErrCanceled = errors.New("mutation: run context is canceled")

	// ErrNilKill reports a missing kill predicate passed to the minimizer.
	ErrNilKill = errors.New("mutation: kill predicate must not be nil")
)
