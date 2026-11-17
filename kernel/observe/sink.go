/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

// Sink is the uniform consumer port for verifiers, recorders and reporters.
//
// A failing sink must not block the remaining sinks or the scheduling loop;
// the pipeline records the failure as a meta observation instead.
type Sink interface {
	Observe(Observation) error
}
