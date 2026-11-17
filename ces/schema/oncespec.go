/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// OnceSpec schedules a single perturbation at a future virtual time.
//
// At is measured from the start of the experiment and must be positive.
type OnceSpec struct {
	At VirtualDuration `json:"at" yaml:"at"`
}
