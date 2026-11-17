/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import "abate/ces/validate"

// Exit codes consumed by CI, shared by every command.
const (
	exitOK = iota
	exitBusinessFailure
	exitInternalError
	exitUsageError
)

// blocking reports whether validation results must fail a run.
//
// Error-severity violations always block. The decision reads validate output
// only and never re-implements a validation rule inside the control plane.
func blocking(vs []validate.Violation) bool {
	for _, v := range vs {
		if v.Severity == validate.SeverityError {
			return true
		}
	}
	return false
}

// blockingStrict additionally promotes warnings to blocking failures.
func blockingStrict(vs []validate.Violation) bool {
	for _, v := range vs {
		if v.Severity == validate.SeverityError || v.Severity == validate.SeverityWarning {
			return true
		}
	}
	return false
}
