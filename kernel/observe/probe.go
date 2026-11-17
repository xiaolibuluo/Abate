/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

// Probe kind identifiers of the simulator-side observation sources.
const (
	ProbeHTTP   = "http"
	ProbeShell  = "shell"
	ProbeMetric = "metric"
	ProbeProm   = "prom"
)

// Probe is a deterministic observation source inside the simulator.
//
// Collect receives the current virtual time and returns a tagged value plus
// its status. A probe declares timeout itself; the pipeline never reads a real
// clock, so the result stays reproducible under replay.
type Probe interface {
	ID() string
	Kind() string
	Collect(at int64) (Value, string)
}
