/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// ContinuousSpec schedules a perturbation that persists for a virtual span.
//
// Duration must be positive and is bounded by the scheduler budget.
type ContinuousSpec struct {
	Duration VirtualDuration `json:"duration" yaml:"duration"`
}
