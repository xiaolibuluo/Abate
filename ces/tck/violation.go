/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

// Violation is one neutral validation failure returned by Driver.Validate.
type Violation struct {
	// RuleID is the CES-VALID rule identifier (e.g. CES-VALID-STEADY-001).
	RuleID string
	// FieldPath locates the offending field (e.g. spec.steadyState[0].holdFor).
	FieldPath string
	// Message is a byte-stable human-readable description.
	Message string
	// Severity is "error" or "warning".
	Severity string
}
