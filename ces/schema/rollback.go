/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Rollback declares the compensation policy of an experiment.
//
// Strategy defaults to auto during normalization. Timeout is required when
// Strategy is timeout. VerifySteadyState re-runs the steady-state assertions
// after the dual operators have been applied.
type Rollback struct {
	Strategy          Strategy        `json:"strategy,omitempty" yaml:"strategy,omitempty"`
	Timeout           VirtualDuration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	VerifySteadyState bool            `json:"verifySteadyState" yaml:"verifySteadyState"`
}
