/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

import (
	"errors"
	"fmt"
	"strings"

	"abate/ces/parse"
	"abate/ces/schema"
	"abate/ces/validate"
)

// AbateDriver is the bundled reference implementation of the TCK Driver.
//
// It adapts Abate's own parse/normalize/validate pipeline, so running the kit
// against it is the project's "eat our own dog food" bootstrap: a regression in
// any CES rule surfaces immediately as a TCK failure. Third parties implement
// their own Driver without importing Abate internals.
type AbateDriver struct{}

// driverName returns a stable, string-only label for a Driver value.
func driverName(d Driver) string {
	return fmt.Sprintf("%T", d)
}

// toFormat maps a textual format hint onto the parser's Format type.
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

// Parse decodes and normalizes a CES document using Abate's pipeline.
func (AbateDriver) Parse(data []byte, format string) (ParseResult, error) {
	exp, err := parse.Parse(data, toFormat(format))
	if err != nil {
		return ParseResult{OK: false, RuleID: ruleIDOf(err), Error: err.Error()}, err
	}
	if _, nerr := parse.Normalize(exp); nerr != nil {
		return ParseResult{OK: false, RuleID: ruleIDOf(nerr), Error: nerr.Error()}, nerr
	}
	return ParseResult{OK: true}, nil
}

// Validate checks a CES document through Abate's validation rule registry.
func (AbateDriver) Validate(data []byte) ([]Violation, error) {
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
	out := make([]Violation, 0, len(viols))
	for _, v := range viols {
		out = append(out, Violation{
			RuleID:    v.RuleID,
			FieldPath: v.FieldPath,
			Message:   v.Message,
			Severity:  string(v.Severity),
		})
	}
	return out, nil
}

// ruleIDOf extracts the CES-PARSE rule identifier from a parse/normalize error,
// returning the empty string for non-structured errors.
func ruleIDOf(err error) string {
	var pe *parse.ParseError
	if errors.As(err, &pe) {
		return pe.RuleID
	}
	return ""
}

// Compile-time assertions that AbateDriver satisfies the Driver contract and that
// the schema import is used (the violation mapping relies on validate.Violation).
var (
	_ Driver = AbateDriver{}
	_        = schema.APIVersion
)
