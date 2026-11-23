/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import "errors"

// Sentinel errors of the real go test mutation backend. Every rejection
// wraps exactly one of these values so callers can classify failures with
// errors.Is.
var (
	// ErrNoGoFiles reports a workspace without any non-test Go source file.
	ErrNoGoFiles = errors.New("gomutate: no non-test Go source files")
	// ErrNoTests reports a workspace whose packages declare no test function.
	ErrNoTests = errors.New("gomutate: no test functions discovered")
	// ErrListTests reports a failure of the go test -list discovery step.
	ErrListTests = errors.New("gomutate: test discovery failed")
	// ErrCoverage reports a failure of the per-test coverage step.
	ErrCoverage = errors.New("gomutate: coverage collection failed")
	// ErrWorkspace reports a failure while materializing a mutated workspace.
	ErrWorkspace = errors.New("gomutate: workspace materialization failed")
	// ErrUnknownNode reports a mutation point that names no discovered node.
	ErrUnknownNode = errors.New("gomutate: unknown mutation node")
	// ErrInvalidRoot reports an unreadable or empty workspace root.
	ErrInvalidRoot = errors.New("gomutate: invalid workspace root")
)
