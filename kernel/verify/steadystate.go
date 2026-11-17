/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

// SteadyState is one hypothesis over the observations of a single probe.
//
// HoldFor is the minimum virtual nanosecond span the hypothesis must hold
// continuously before the verdict can be HELD.
type SteadyState struct {
	ProbeID     string
	Expectation Expectation
	Tolerance   Tolerance
	HoldFor     int64
}
