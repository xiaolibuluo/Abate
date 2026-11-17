/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package metrics

// Robustness holds the independent robustness counters of one run. Each cause
// is a separate column; a panic localizer is set only when the run did not
// finish, and an empty one means the report is complete.
type Robustness struct {
	PanicLocalizer          string
	TimeoutCount            int
	OOMCount                int
	DeterminismFailureCount int
	ProcessedFiles          int
	RegisteredFiles         int
}

// Complete reports whether every registered file was processed. A gap is a
// zero-tolerance condition and never a reason to skip or exclude a file.
func (r Robustness) Complete() bool {
	return r.ProcessedFiles == r.RegisteredFiles
}
