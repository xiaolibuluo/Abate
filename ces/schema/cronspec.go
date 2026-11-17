/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// CronSpec schedules perturbations from a five-field cron expression.
//
// Timezone is a virtual timezone label; it never reads the host clock.
type CronSpec struct {
	Expression string `json:"expression" yaml:"expression"`
	Timezone   string `json:"timezone,omitempty" yaml:"timezone,omitempty"`
}
