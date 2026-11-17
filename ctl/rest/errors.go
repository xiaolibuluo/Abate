/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import "errors"

// Control-plane error classes. They mirror the CLI four-state exit semantics
// and the REST problem+json status mapping (validation→400, not-found→404,
// conflict→409, internal→500).
var (
	errUnsupportedLanguage = errors.New("rest: unsupported language")
	errNoPaths             = errors.New("rest: no source paths")
	errNoRules             = errors.New("rest: no active rules")
	errReadSource          = errors.New("rest: cannot read source")
	errFrontendFailed      = errors.New("rest: frontend construction failed")
	errRuleLoadFailed      = errors.New("rest: rule load failed")
	errParseFailed         = errors.New("rest: parse failed")
	errPluginInvalid       = errors.New("rest: invalid plugin")
	errPluginExists        = errors.New("rest: plugin already exists")
	errPluginMissing       = errors.New("rest: plugin not found")
	errBadJSON             = errors.New("rest: malformed request body")
)

// validCheckFormat reports whether a report format is implemented.
func validCheckFormat(format string) bool {
	switch format {
	case "console", "json", "sarif":
		return true
	default:
		return false
	}
}
