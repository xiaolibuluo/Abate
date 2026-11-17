/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// SteadyState is one hypothesis that must hold before and after a perturbation.
//
// ProbeRef names a probe declared in the experiment, HoldFor is the minimum
// virtual duration the hypothesis must remain satisfied, and Tolerance
// qualifies how closely the observed value must match Expectation.
type SteadyState struct {
	ProbeRef    string          `json:"probeRef" yaml:"probeRef"`
	Expectation Expectation     `json:"expectation" yaml:"expectation"`
	Tolerance   Tolerance       `json:"tolerance" yaml:"tolerance"`
	HoldFor     VirtualDuration `json:"holdFor" yaml:"holdFor"`
}
