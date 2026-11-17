/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// ProbeSpec carries the union of the four kind-specific probe configurations.
//
// The HTTP group uses URL, Method, Headers and ExpectStatus; the shell group
// uses Command, Timeout and ExpectExitCode; the metric group uses Query and
// Expect; the PromQL group uses PromQL, Range and Expect. CES-VALID checks
// that the fields belonging to the declared kind are present and legal.
type ProbeSpec struct {
	URL            string            `json:"url,omitempty" yaml:"url,omitempty"`
	Method         string            `json:"method,omitempty" yaml:"method,omitempty"`
	Headers        map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	ExpectStatus   int               `json:"expectStatus,omitempty" yaml:"expectStatus,omitempty"`
	Command        string            `json:"command,omitempty" yaml:"command,omitempty"`
	Timeout        VirtualDuration   `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	ExpectExitCode int               `json:"expectExitCode,omitempty" yaml:"expectExitCode,omitempty"`
	Query          string            `json:"query,omitempty" yaml:"query,omitempty"`
	Expect         string            `json:"expect,omitempty" yaml:"expect,omitempty"`
	PromQL         string            `json:"promql,omitempty" yaml:"promql,omitempty"`
	Range          VirtualDuration   `json:"range,omitempty" yaml:"range,omitempty"`
}
