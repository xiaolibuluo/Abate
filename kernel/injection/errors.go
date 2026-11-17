/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import "errors"

// Sentinel errors of the chaos injection kernel.
var (
	// ErrNilWorld reports an injector built without a simulation world.
	ErrNilWorld = errors.New("injection: world must not be nil")

	// ErrNilOperator reports a nil operator passed to Inject or Rollback.
	ErrNilOperator = errors.New("injection: operator must not be nil")

	// ErrWrongFamily reports an operator outside the C family.
	ErrWrongFamily = errors.New("injection: operator is not a chaos-family operator")

	// ErrUnknownOperator reports an operator outside the twelve registered faults.
	ErrUnknownOperator = errors.New("injection: operator is not registered")

	// ErrUnknownTarget reports a fault target absent from the simulation world.
	ErrUnknownTarget = errors.New("injection: unknown fault target")

	// ErrNotInvertible reports an operator with no declared dual.
	ErrNotInvertible = errors.New("injection: operator has no dual")

	// ErrNotDual reports a rollback step that did not receive an Inverse result.
	ErrNotDual = errors.New("injection: rollback requires the dual operator")

	// ErrInvalidParams reports malformed fault parameters.
	ErrInvalidParams = errors.New("injection: malformed fault parameters")

	// ErrCanceled reports a canceled or missing run context.
	ErrCanceled = errors.New("injection: run context is canceled")

	// ErrRollbackIncomplete reports a baseline that was not fully restored.
	ErrRollbackIncomplete = errors.New("injection: baseline not restored")

	// ErrDependencyTimeout reports a virtual dependency timeout fault.
	ErrDependencyTimeout = errors.New("injection: dependency timed out")

	// ErrDependencyFailure reports a virtual dependency error fault.
	ErrDependencyFailure = errors.New("injection: dependency returned an error")
)
