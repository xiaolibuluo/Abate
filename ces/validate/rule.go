/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import "abate/ces/schema"

// Rule is one deterministic, side-effect-free validation rule.
type Rule interface {
	// ID returns the globally unique rule identifier.
	ID() string
	// Check returns all violations this rule finds in exp.
	Check(exp *schema.Experiment) []Violation
}

// funcRule adapts a plain function into a Rule.
type funcRule struct {
	id    string
	check func(*schema.Experiment) []Violation
}

// ID returns the adapted rule identifier.
func (r funcRule) ID() string { return r.id }

// Check delegates to the adapted function.
func (r funcRule) Check(exp *schema.Experiment) []Violation {
	return r.check(exp)
}

// newRule builds a Rule from an identifier and a check function.
func newRule(id string, check func(*schema.Experiment) []Violation) Rule {
	return funcRule{id: id, check: check}
}

// violation builds one deterministic violation record.
func violation(id, path, msg string, sev Severity) Violation {
	return Violation{RuleID: id, FieldPath: path, Message: msg, Severity: sev}
}

// errorAt builds one deterministic error-severity violation record.
func errorAt(id, path, msg string) Violation {
	return violation(id, path, msg, SeverityError)
}

// warningAt builds one deterministic warning-severity violation record.
func warningAt(id, path, msg string) Violation {
	return violation(id, path, msg, SeverityWarning)
}
