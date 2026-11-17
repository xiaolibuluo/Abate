/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "fmt"

// Severity values used by structured frontend diagnostics.
const (
	// SeverityError marks a finding that prevents a faithful parse.
	SeverityError = "error"
	// SeverityWarning marks a recoverable finding.
	SeverityWarning = "warning"
	// SeverityInfo marks an informational finding.
	SeverityInfo = "info"
)

// MaxDiagnostics bounds the number of diagnostics a frontend may emit, which
// prevents an error storm on a badly damaged source file.
const MaxDiagnostics = 100

// Diagnostic is a structured, language-tagged parsing message.
type Diagnostic struct {
	// Line is the one-based line of the finding.
	Line int
	// Column is the one-based column of the finding.
	Column int
	// Severity is one of SeverityError, SeverityWarning, or SeverityInfo.
	Severity string
	// Code follows the LANG-<LANG>-E/W-xxx convention.
	Code string
	// Message is the human-readable explanation.
	Message string
	// Path is the virtual source path the finding belongs to.
	Path string
}

// String renders the diagnostic in a stable one-line form.
func (d Diagnostic) String() string {
	return fmt.Sprintf("%s:%d:%d %s %s: %s",
		d.Path, d.Line, d.Column, d.Severity, d.Code, d.Message)
}

// IsError reports whether the diagnostic has error severity.
func (d Diagnostic) IsError() bool {
	return d.Severity == SeverityError
}
