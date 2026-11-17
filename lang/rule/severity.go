/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

// Severity ranks a finding's impact on the check gate.
type Severity string

// Severity values used by Specifications and Findings.
const (
	// SeverityError marks a finding that must fail the check gate.
	SeverityError Severity = "error"
	// SeverityWarning marks a finding that should be reviewed.
	SeverityWarning Severity = "warning"
	// SeverityInfo marks an advisory finding.
	SeverityInfo Severity = "info"
)

// Valid reports whether the severity is one of the allocated values.
func (s Severity) Valid() bool {
	switch s {
	case SeverityError, SeverityWarning, SeverityInfo:
		return true
	default:
		return false
	}
}

// Rank returns the numeric weight of the severity, with error ranked highest.
func (s Severity) Rank() int {
	switch s {
	case SeverityError:
		return 3
	case SeverityWarning:
		return 2
	case SeverityInfo:
		return 1
	default:
		return 0
	}
}

// String returns the stable wire text of the severity.
func (s Severity) String() string {
	return string(s)
}
