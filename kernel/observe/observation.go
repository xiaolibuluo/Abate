/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

import "encoding/json"

// Probe status identifiers of a normalized observation.
const (
	StatusOK       = "ok"
	StatusDegraded = "degraded"
	StatusFailed   = "failed"
	StatusTimeout  = "timeout"
)

// Observation is one normalized, totally ordered record of the stream.
//
// At is an integer nanosecond virtual time and Seq is assigned by the
// scheduler. Status is one of ok, degraded, failed or timeout.
type Observation struct {
	At      int64  `json:"at"`
	Seq     uint64 `json:"seq"`
	ProbeID string `json:"probeId"`
	Value   Value  `json:"value"`
	Status  string `json:"status"`
}

// Encode returns the canonical JSON encoding used for stream hashing.
func (o Observation) Encode() ([]byte, error) {
	return json.Marshal(o)
}
