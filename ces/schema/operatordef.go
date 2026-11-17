/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

// Operator family identifiers used by CES-VALID target-domain matching.
const (
	FamilyMutation = "mutation"
	FamilyChaos    = "chaos"
)

// OperatorDef is an inline or referenced perturbation operator definition.
//
// Family is mutation or chaos; TargetKind records the required target domain
// (source for mutation, simulator for chaos). A non-invertible operator must
// declare a Compensation strategy or CES-VALID-METHOD-002 rejects it.
type OperatorDef struct {
	ID           string            `json:"id" yaml:"id"`
	Family       string            `json:"family" yaml:"family"`
	Params       map[string]string `json:"params,omitempty" yaml:"params,omitempty"`
	Targets      []string          `json:"targets" yaml:"targets"`
	Invertible   bool              `json:"invertible" yaml:"invertible"`
	Compensation string            `json:"compensation,omitempty" yaml:"compensation,omitempty"`
	TargetKind   string            `json:"targetKind,omitempty" yaml:"targetKind,omitempty"`
}
