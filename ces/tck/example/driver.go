/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package example

import (
	"errors"
	"strings"

	"abate/ces/parse"
	"abate/ces/tck"
	"abate/ces/validate"
)

// Driver is the example third-party implementation of the TCK Driver contract.
//
// It adapts the public parse/normalize/validate pipeline; a real integrator
// would substitute its own engine behind the same interface. The interface is
// the portable boundary — the kit never sees internal types.
type Driver struct{}

// toFormat maps a textual hint onto the parser's Format enum.
func toFormat(hint string) parse.Format {
	switch strings.ToLower(strings.TrimSpace(hint)) {
	case "json":
		return parse.FormatJSON
	case "yaml":
		return parse.FormatYAML
	default:
		return parse.FormatAuto
	}
}

// Parse decodes, normalizes and reports acceptance of a CES document.
func (Driver) Parse(data []byte, format string) (tck.ParseResult, error) {
	exp, err := parse.Parse(data, toFormat(format))
	if err != nil {
		return tck.ParseResult{OK: false, RuleID: ruleIDOf(err), Error: err.Error()}, err
	}
	if _, nerr := parse.Normalize(exp); nerr != nil {
		return tck.ParseResult{OK: false, RuleID: ruleIDOf(nerr), Error: nerr.Error()}, nerr
	}
	return tck.ParseResult{OK: true}, nil
}

// Validate checks an already-decoded CES document and returns every violation.
func (Driver) Validate(data []byte) ([]tck.Violation, error) {
	exp, err := parse.Parse(data, parse.FormatAuto)
	if err != nil {
		return nil, err
	}
	norm, nerr := parse.Normalize(exp)
	if nerr != nil {
		return nil, nerr
	}
	viols, verr := validate.Validate(norm)
	if verr != nil {
		return nil, verr
	}
	out := make([]tck.Violation, 0, len(viols))
	for _, v := range viols {
		out = append(out, tck.Violation{
			RuleID:    v.RuleID,
			FieldPath: v.FieldPath,
			Message:   v.Message,
			Severity:  string(v.Severity),
		})
	}
	return out, nil
}

// ruleIDOf extracts the CES-PARSE rule identifier from a structured error.
func ruleIDOf(err error) string {
	var pe *parse.ParseError
	if errors.As(err, &pe) {
		return pe.RuleID
	}
	return ""
}

// Compile-time assertion that the example driver satisfies the TCK contract.
var _ tck.Driver = Driver{}
