/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

// Exit codes of the check pipeline, aligned with PLAN-07 section 4 and
// LNG-CHK section 3: clean, business failure (a gate violation), and an
// internal error strictly separated from a violation.
const (
	// ExitClean reports a run with zero findings.
	ExitClean = 0
	// ExitViolations reports a run with at least one finding.
	ExitViolations = 1
	// ExitInternal reports a frontend, source or option failure.
	ExitInternal = 2
)

// ExitCode maps a result and an error to the check exit code.
func ExitCode(result Result, err error) int {
	if err != nil {
		return ExitInternal
	}
	if result.HasViolations() {
		return ExitViolations
	}
	return ExitClean
}
