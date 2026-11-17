/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

// Verdict is the decision for one steady-state hypothesis.
//
// FirstSeq and Timeline record the first and every violating observation, so
// a replay can be re-judged and compared byte for byte.
type Verdict struct {
	ProbeID  string   `json:"probeId"`
	Outcome  Outcome  `json:"outcome"`
	Reason   string   `json:"reason"`
	FirstSeq uint64   `json:"firstSeq"`
	Timeline []uint64 `json:"timeline,omitempty"`
}
