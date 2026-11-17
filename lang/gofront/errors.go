/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import "errors"

// Sentinel errors of the gofront package.
var (
	// ErrUnsupportedLanguage reports a SourceUnit tagged with another language.
	ErrUnsupportedLanguage = errors.New("gofront: unsupported language")
	// ErrUnknownEditPath reports an incremental edit for a file not in the unit.
	ErrUnknownEditPath = errors.New("gofront: edit path not in source unit")
	// ErrTokenFidelity reports a token stream that AIR could not reproduce.
	ErrTokenFidelity = errors.New("gofront: token fidelity check failed")
	// ErrNoPackageClause reports a source file without a usable package clause.
	ErrNoPackageClause = errors.New("gofront: missing package clause")
)
