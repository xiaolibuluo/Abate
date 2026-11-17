/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import "errors"

// ErrUnsupportedLanguage reports a tag or path that names no supported language.
var ErrUnsupportedLanguage = errors.New("unsupported language")

// ErrLexFailed reports that the token scanner could not classify the source.
var ErrLexFailed = errors.New("lexical analysis failed")

// ErrParseDepth reports that the recursive-descent depth limit was exceeded.
var ErrParseDepth = errors.New("parse depth limit exceeded")

// ErrDiagnosticLimit reports that the diagnostic sink reached air.MaxDiagnostics.
var ErrDiagnosticLimit = errors.New("diagnostic limit exceeded")

// ErrUnknownToken reports a token kind outside the frozen TokenKind enum.
var ErrUnknownToken = errors.New("unknown token kind")

// ErrBadDialect reports a dialect spelling that cannot be normalized.
var ErrBadDialect = errors.New("invalid dialect")

// ErrNilSpec reports a nil LanguageSpec passed to a frontend constructor.
var ErrNilSpec = errors.New("nil language spec")
