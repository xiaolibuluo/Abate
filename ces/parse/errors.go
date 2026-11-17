/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import "fmt"

// Rule identifiers for every deterministic parser failure path.
const (
	RuleUnknownField       = "CES-PARSE-001"
	RuleUnsupportedVersion = "CES-PARSE-002"
	RuleInvalidFormat      = "CES-PARSE-003"
	RuleEmptyDocument      = "CES-PARSE-004"
	RuleMalformedDocument  = "CES-PARSE-005"
	RuleMissingField       = "CES-PARSE-006"
	RuleDanglingProbeRef   = "CES-PARSE-007"
	RuleScheduleExclusive  = "CES-PARSE-008"
	RuleExprSyntax         = "CES-PARSE-009"
	RuleExprDepth          = "CES-PARSE-010"
	RuleExprUnknownOp      = "CES-PARSE-011"
	RuleExprNotInvertible  = "CES-PARSE-012"
	RuleInputTooLarge      = "CES-PARSE-013"
)

// ParseError is a positioned, deterministic failure of the CES parser.
//
// Its message never contains a pointer address or a map iteration order, so
// the same bad input always yields byte-identical diagnostics.
//
// The name ParseError is mandated by the CES specification (03-ces-spec.md,
// module CES-PARSE section 3); it is kept verbatim so the public contract
// matches the architecture document, hence the narrow revive stutter waiver.
type ParseError struct { //nolint:revive // CES spec mandates the ParseError type name.
	Line      int
	Column    int
	FieldPath string
	RuleID    string
	Message   string
}

// Error renders the structured error into one deterministic line.
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s %s:%d:%d %s", e.RuleID, e.FieldPath, e.Line, e.Column, e.Message)
}

// newError builds a document-level error at the default position.
func newError(rule, path, msg string) *ParseError {
	return &ParseError{Line: 1, Column: 1, FieldPath: path, RuleID: rule, Message: msg}
}

// newExprError builds an expression error at a rune column.
func newExprError(rule, msg string, col int) *ParseError {
	return &ParseError{
		Line: 1, Column: col + 1,
		FieldPath: "$.spec.method.expression", RuleID: rule, Message: msg,
	}
}
