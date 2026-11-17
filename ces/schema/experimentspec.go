/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// ExperimentSpec holds the executable body of a CES document.
//
// SteadyState must contain at least one hypothesis, Method describes the
// perturbation algebra, and Rollback plus Schedule are always required.
// Seed and Timeout are optional: normalization fills their defaults.
type ExperimentSpec struct {
	Target      Target          `json:"target" yaml:"target"`
	SteadyState []SteadyState   `json:"steadyState" yaml:"steadyState"`
	Method      Method          `json:"method" yaml:"method"`
	Probes      []Probe         `json:"probes,omitempty" yaml:"probes,omitempty"`
	Rollback    Rollback        `json:"rollback" yaml:"rollback"`
	Schedule    Schedule        `json:"schedule" yaml:"schedule"`
	Timeout     VirtualDuration `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Seed        uint64          `json:"seed,omitempty" yaml:"seed,omitempty"`
}
