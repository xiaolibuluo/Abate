/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

// CES-VALID rule identifiers (03-ces-spec.md, module CES-VALID section 4).
// The specification narrative says "22 rules" but its itemized list contains
// 25; this implementation realizes every listed rule without inventing or
// dropping any, and reports the discrepancy in the delivery record.
const (
	RuleSteady001   = "CES-VALID-STEADY-001"
	RuleSteady002   = "CES-VALID-STEADY-002"
	RuleSteady003   = "CES-VALID-STEADY-003"
	RuleSteady004   = "CES-VALID-STEADY-004"
	RuleSteady005   = "CES-VALID-STEADY-005"
	RuleMethod001   = "CES-VALID-METHOD-001"
	RuleMethod002   = "CES-VALID-METHOD-002"
	RuleMethod003   = "CES-VALID-METHOD-003"
	RuleMethod004   = "CES-VALID-METHOD-004"
	RuleProbe001    = "CES-VALID-PROBE-001"
	RuleProbe002    = "CES-VALID-PROBE-002"
	RuleProbe003    = "CES-VALID-PROBE-003"
	RuleProbe004    = "CES-VALID-PROBE-004"
	RuleProbe005    = "CES-VALID-PROBE-005"
	RuleRollback001 = "CES-VALID-ROLLBACK-001"
	RuleRollback002 = "CES-VALID-ROLLBACK-002"
	RuleRollback003 = "CES-VALID-ROLLBACK-003"
	RuleSchedule001 = "CES-VALID-SCHEDULE-001"
	RuleSchedule002 = "CES-VALID-SCHEDULE-002"
	RuleSchedule003 = "CES-VALID-SCHEDULE-003"
	RuleSchedule004 = "CES-VALID-SCHEDULE-004"
	RuleTarget001   = "CES-VALID-TARGET-001"
	RuleTarget002   = "CES-VALID-TARGET-002"
	RuleTarget003   = "CES-VALID-TARGET-003"
	RuleTarget004   = "CES-VALID-TARGET-004"
)

// allRuleIDs returns every rule id in a stable, lexicographic order.
func allRuleIDs() []string {
	return []string{
		RuleMethod001, RuleMethod002, RuleMethod003, RuleMethod004,
		RuleProbe001, RuleProbe002, RuleProbe003, RuleProbe004, RuleProbe005,
		RuleRollback001, RuleRollback002, RuleRollback003,
		RuleSchedule001, RuleSchedule002, RuleSchedule003, RuleSchedule004,
		RuleSteady001, RuleSteady002, RuleSteady003, RuleSteady004, RuleSteady005,
		RuleTarget001, RuleTarget002, RuleTarget003, RuleTarget004,
	}
}
