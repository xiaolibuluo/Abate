/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import "errors"

// Sentinel errors of the Go built-in rule set.
var (
	// ErrBadSource reports a fixture or source buffer that cannot be parsed.
	ErrBadSource = errors.New("gobuiltin: source cannot be parsed")
	// ErrNilRegistry reports a nil registry handed to LoadBuiltins.
	ErrNilRegistry = errors.New("gobuiltin: nil registry")
)
