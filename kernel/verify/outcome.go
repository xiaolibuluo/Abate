/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

// Outcome is the result of one steady-state assertion.
type Outcome string

// Outcome values of a steady-state assertion.
const (
	OutcomeHeld         Outcome = "HELD"
	OutcomeViolated     Outcome = "VIOLATED"
	OutcomeInconclusive Outcome = "INCONCLUSIVE"
)
