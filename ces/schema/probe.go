/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Probe is a periodic observation source declared in an experiment.
//
// ID is unique within the document and is referenced by SteadyState. Interval
// must be positive; the pipeline collects probes in ID order so registration
// order can never affect the observation stream.
type Probe struct {
	ID       string          `json:"id" yaml:"id"`
	Kind     ProbeKind       `json:"kind" yaml:"kind"`
	Spec     ProbeSpec       `json:"spec" yaml:"spec"`
	Interval VirtualDuration `json:"interval" yaml:"interval"`
}
