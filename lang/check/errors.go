/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import "errors"

// Sentinel errors of the check package. Every failure wraps exactly one of
// these values so callers can classify an invocation with errors.Is.
var (
	// ErrInvalidOptions reports an Options value that failed validation.
	ErrInvalidOptions = errors.New("check: invalid options")
	// ErrInvalidFormat reports an unsupported report format.
	ErrInvalidFormat = errors.New("check: invalid format")
	// ErrNoPaths reports an empty expanded source path list.
	ErrNoPaths = errors.New("check: no source paths")
	// ErrReadSource reports a source buffer that could not be read.
	ErrReadSource = errors.New("check: cannot read source")
	// ErrFrontendFailed reports a contract failure inside the frontend.
	ErrFrontendFailed = errors.New("check: frontend parse failed")
	// ErrParseFailed reports error-severity diagnostics for a source file.
	ErrParseFailed = errors.New("check: source has syntax diagnostics")
	// ErrNoFrontend reports a checker built without a frontend.
	ErrNoFrontend = errors.New("check: nil frontend")
	// ErrNoRules reports a check with no active rule after filtering.
	ErrNoRules = errors.New("check: no active rules")
)
