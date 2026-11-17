/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

type (
	// Outcome is the measured result of one real frontend parse. Kind carries
	// the independent timeout, memory, contract and success classification;
	// AllocDeltaBytes is the TotalAlloc increment of that parse; the two AIR
	// identifiers are filled by the runner before the drilldown record is made.
	Outcome struct {
		Kind            string
		AllocDeltaBytes uint64
		Diagnostics     int
		Errors          int
		AIRIDBefore     string
		AIRIDAfter      string
	}

	// ProbeBudget is the libFuzzer-aligned resource envelope of one parse.
	ProbeBudget struct {
		TimeoutMillis int64
		MemLimitBytes uint64
	}
)
