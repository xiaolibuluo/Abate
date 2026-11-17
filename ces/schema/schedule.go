/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Schedule selects exactly one of the three supported scheduling modes.
//
// CES-VALID-SCHEDULE-001 enforces mutual exclusion: exactly one pointer must
// be non-nil. Keeping them as pointers makes "unset" distinguishable from a
// zero-valued specification.
type Schedule struct {
	Once       *OnceSpec       `json:"once,omitempty" yaml:"once,omitempty"`
	Cron       *CronSpec       `json:"cron,omitempty" yaml:"cron,omitempty"`
	Continuous *ContinuousSpec `json:"continuous,omitempty" yaml:"continuous,omitempty"`
}
