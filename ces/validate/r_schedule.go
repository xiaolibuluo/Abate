/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import "abate/ces/schema"

// scheduleRules returns the four schedule rules in registration order.
func scheduleRules() []Rule {
	return []Rule{
		newRule(RuleSchedule001, checkSchedule001),
		newRule(RuleSchedule002, checkSchedule002),
		newRule(RuleSchedule003, checkSchedule003),
		newRule(RuleSchedule004, checkSchedule004),
	}
}

// scheduleModeCount counts the non-nil scheduling modes deterministically.
func scheduleModeCount(s schema.Schedule) int {
	n := 0
	if s.Once != nil {
		n++
	}
	if s.Cron != nil {
		n++
	}
	if s.Continuous != nil {
		n++
	}
	return n
}

// checkSchedule001 enforces exactly one scheduling mode.
func checkSchedule001(exp *schema.Experiment) []Violation {
	if scheduleModeCount(exp.Spec.Schedule) != 1 {
		return []Violation{errorAt(RuleSchedule001, "spec.schedule",
			"exactly one of once, cron or continuous must be set")}
	}
	return nil
}

// checkSchedule002 validates the five-field cron expression and virtual timezone.
func checkSchedule002(exp *schema.Experiment) []Violation {
	c := exp.Spec.Schedule.Cron
	if c == nil {
		return nil
	}
	var out []Violation
	if !validCron(c.Expression) {
		out = append(out, errorAt(RuleSchedule002, "spec.schedule.cron.expression",
			"cron expression must have five valid fields"))
	}
	if !validVirtualTimezone(c.Timezone) {
		out = append(out, errorAt(RuleSchedule002, "spec.schedule.cron.timezone",
			"timezone must be UTC or a Virtual/* label"))
	}
	return out
}

// checkSchedule003 requires a positive continuous duration.
func checkSchedule003(exp *schema.Experiment) []Violation {
	c := exp.Spec.Schedule.Continuous
	if c != nil && c.Duration <= 0 {
		return []Violation{errorAt(RuleSchedule003,
			"spec.schedule.continuous.duration",
			"continuous duration must be positive")}
	}
	return nil
}

// checkSchedule004 requires once.At to be a future virtual time (> 0).
func checkSchedule004(exp *schema.Experiment) []Violation {
	o := exp.Spec.Schedule.Once
	if o != nil && o.At <= 0 {
		return []Violation{errorAt(RuleSchedule004, "spec.schedule.once.at",
			"once.at must be a future virtual time")}
	}
	return nil
}
