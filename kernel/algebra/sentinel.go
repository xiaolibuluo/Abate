/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import "errors"

// MaxDepth is the maximum operator tree depth accepted by Compose and Concurrent.
const MaxDepth = 64

// separator joins canonical fields and pairs without leaking map order.
const separator = "\x00"

// Sentinel errors of the algebra package.
//
//nolint:gochecknoglobals // sentinel errors are required for the errors.Is contract.
var (
	// ErrNotInvertible reports an operator that has no dual.
	ErrNotInvertible = errors.New("algebra: operator is not invertible")

	// ErrIncompatible reports two operators whose target domains are incompatible.
	ErrIncompatible = errors.New("algebra: incompatible operator domains")

	// ErrTargetConflict reports two concurrent operators sharing a target.
	ErrTargetConflict = errors.New("algebra: concurrent operators share targets")

	// ErrTooDeep reports an operator tree deeper than MaxDepth.
	ErrTooDeep = errors.New("algebra: operator tree exceeds max depth")

	// ErrNilOperator reports a nil operator passed to an algebra operation.
	ErrNilOperator = errors.New("algebra: operator must not be nil")
)
