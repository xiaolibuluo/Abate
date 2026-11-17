/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

// Severity classifies how strongly a violation blocks execution.
type Severity string

// Severity levels of a validation violation.
const (
	// SeverityError marks a violation that must be fixed before execution.
	SeverityError Severity = "error"
	// SeverityWarning marks an advisory violation, promoted to an error by strict mode.
	SeverityWarning Severity = "warning"
)
