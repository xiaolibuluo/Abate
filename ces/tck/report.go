/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

import "fmt"

// Report is the compliance verdict for one TCK run against one Driver.
type Report struct {
	// Driver names the implementation under test.
	Driver string
	// Total is the number of cases executed.
	Total int
	// Passed is the number of cases that matched the expected outcome.
	Passed int
	// Failed is the number of cases whose outcome diverged from expected.
	Failed int
	// Skipped is the number of cases not applicable to the driver.
	Skipped int
	// Details carries one entry per executed case, in case order.
	Details []CaseResult
	// Compliant is true only when Failed == 0 and Skipped == 0.
	Compliant bool
}

// summary renders a single-line compliance summary for logs and reports.
func (r Report) summary() string {
	return fmt.Sprintf("TCK[%s]: %d/%d passed, %d failed, %d skipped, compliant=%t",
		r.Driver, r.Passed, r.Total, r.Failed, r.Skipped, r.Compliant)
}
